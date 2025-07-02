// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"reflect"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/vim25/types"
)

func TestEventManagerVPX(t *testing.T) {
	logEvents = testing.Verbose()
	ctx := context.Background()

	m := VPX()
	m.Datacenter = 2

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	e := event.NewManager(c.Client)
	count := m.Count()

	root := c.ServiceContent.RootFolder
	vm := m.Map().Any("VirtualMachine").(*VirtualMachine)
	host := m.Map().Get(vm.Runtime.Host.Reference()).(*HostSystem)

	vmEvents := 6 // BeingCreated + InstanceUuid + Uuid + Created + Starting + PoweredOn
	tests := []struct {
		obj    types.ManagedObjectReference
		expect int
		ids    []string
	}{
		{root, -1 * count.Machine, nil},
		{root, 1, []string{"SessionEvent"}}, // UserLoginSessionEvent
		{vm.Reference(), 0, []string{"SessionEvent"}},
		{root, count.Machine, []string{"VmCreatedEvent"}},     // concrete type
		{root, count.Machine * vmEvents, []string{"VmEvent"}}, // base type
		{vm.Reference(), 1, []string{"VmCreatedEvent"}},
		{vm.Reference(), vmEvents, nil},
		{host.Reference(), len(host.Vm), []string{"VmCreatedEvent"}},
		{host.Reference(), len(host.Vm) * vmEvents, nil},
	}

	for i, test := range tests {
		n := 0
		filter := types.EventFilterSpec{
			Entity: &types.EventFilterSpecByEntity{
				Entity:    test.obj,
				Recursion: types.EventFilterSpecRecursionOptionAll,
			},
			EventTypeId: test.ids,
			MaxCount:    100,
		}

		f := func(obj types.ManagedObjectReference, events []types.BaseEvent) error {
			n += len(events)

			qevents, qerr := e.QueryEvents(ctx, filter)
			if qerr != nil {
				t.Fatal(qerr)
			}

			if n != len(qevents) {
				t.Errorf("%d vs %d", n, len(qevents))
			}

			return nil
		}

		err = e.Events(ctx, []types.ManagedObjectReference{test.obj}, filter.MaxCount, false, false, f, test.ids...)
		if err != nil {
			t.Fatalf("%d: %s", i, err)
		}

		if test.expect < 0 {
			expect := test.expect * -1
			if n < expect {
				t.Errorf("%d: expected at least %d events, got: %d", i, expect, n)
			}
			continue
		}

		if test.expect != n {
			t.Errorf("%d: expected %d events, got: %d", i, test.expect, n)
		}
	}

	// Test that we don't panic if event ID is not defined in esx.EventInfo
	type TestHostRemovedEvent struct {
		types.HostEvent
	}
	var hre TestHostRemovedEvent
	kind := reflect.TypeOf(hre)
	types.Add(kind.Name(), kind)

	err = e.PostEvent(ctx, &hre)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEventManagerRead(t *testing.T) {
	logEvents = testing.Verbose()
	ctx := context.Background()
	m := VPX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	vc, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := vc.Logout(ctx); err != nil {
			t.Fatal(err)
		}
	}()

	spec := types.EventFilterSpec{
		Entity: &types.EventFilterSpecByEntity{
			Entity:    vc.Client.ServiceContent.RootFolder,
			Recursion: types.EventFilterSpecRecursionOptionChildren,
		},
	}
	em := event.NewManager(vc.Client)
	c, err := em.CreateCollectorForEvents(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	page, err := c.LatestPage(ctx)
	if err != nil {
		t.Fatal(err)
	}
	nevents := len(page)
	if nevents == 0 {
		t.Fatal("no recent events")
	}
	tests := []struct {
		max    int
		rewind bool
		order  bool
		read   func(context.Context, int32) ([]types.BaseEvent, error)
	}{
		{nevents, true, true, c.ReadNextEvents},
		{nevents / 3, true, true, c.ReadNextEvents},
		{nevents * 3, false, true, c.ReadNextEvents},
		{3, false, false, c.ReadPreviousEvents},
		{nevents * 3, false, true, c.ReadNextEvents},
	}

	for _, test := range tests {
		var all []types.BaseEvent
		count := 0
		for {
			events, err := test.read(ctx, int32(test.max))
			if err != nil {
				t.Fatal(err)
			}
			if len(events) == 0 {
				// expecting 0 below as we've read all events in the page
				zevents, nerr := test.read(ctx, int32(test.max))
				if nerr != nil {
					t.Fatal(nerr)
				}
				if len(zevents) != 0 {
					t.Errorf("zevents=%d", len(zevents))
				}
				break
			}
			count += len(events)
			all = append(all, events...)
		}
		if count < len(page) {
			t.Errorf("expected at least %d events, got: %d", len(page), count)
		}

		for i := 1; i < len(all); i++ {
			prev := all[i-1].GetEvent().Key
			key := all[i].GetEvent().Key
			if test.order {
				if prev > key {
					t.Errorf("key %d > %d", prev, key)
				}
			} else {
				if prev < key {
					t.Errorf("key %d < %d", prev, key)
				}
			}
		}

		if test.rewind {
			if err = c.Rewind(ctx); err != nil {
				t.Error(err)
			}
		}
	}

	// after Reset() we should only get events via ReadPreviousEvents
	if err = c.Reset(ctx); err != nil {
		t.Fatal(err)
	}

	events, err := c.ReadNextEvents(ctx, int32(nevents))
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 0 {
		t.Errorf("expected 0 events, got %d", len(events))
	}

	event := &types.GeneralEvent{Message: "vcsim"}
	event.Datacenter = &types.DatacenterEventArgument{
		Datacenter: m.Map().Any("Datacenter").Reference(),
	}
	err = em.PostEvent(ctx, event)
	if err != nil {
		t.Fatal(err)
	}

	events, err = c.ReadNextEvents(ctx, int32(nevents))
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 events, got %d", len(events))
	}

	count := 0
	for {
		events, err = c.ReadPreviousEvents(ctx, 3)
		if err != nil {
			t.Fatal(err)
		}
		if len(events) == 0 {
			break
		}
		count += len(events)
	}
	if count < nevents {
		t.Errorf("expected %d events, got %d", nevents, count)
	}
}
