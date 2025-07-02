// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestVirtualMachinePropertyExtraConfig(t *testing.T) {
	key := "guestinfo.test"
	val := func(v types.AnyType) types.OptionValue {
		return types.OptionValue{Key: key, Value: v}
	}
	f := func(item string) string {
		return (&mo.Field{Path: "config.extraConfig", Key: key, Item: item}).String()
	}

	tests := []types.PropertyChange{
		{Name: f(""), Val: val("111"), Op: types.PropertyChangeOpAdd},
		{Name: f(""), Val: val("222"), Op: types.PropertyChangeOpAssign},
		{Name: f("key"), Val: val("333"), Op: types.PropertyChangeOpAssign},
		{Name: f("value"), Val: val("444"), Op: types.PropertyChangeOpAssign},
		{Name: f(""), Val: val(""), Op: types.PropertyChangeOpRemove},
	}

	expect := map[string]string{
		f("key"):   key,
		f("value"): "444",
	}

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		kind := []string{"VirtualMachine"}

		m := view.NewManager(c)
		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			t.Fatal(err)
		}
		defer v.Destroy(ctx)

		refs, err := v.Find(ctx, kind, property.Match{})
		if err != nil {
			t.Fatal(err)
		}
		vm := object.NewVirtualMachine(c, refs[0])

		pc := property.DefaultCollector(c)

		for _, test := range tests {
			t.Logf("%s: %s", test.Op, test.Name)
			update := make(chan bool)
			parked := sync.OnceFunc(func() { update <- true })

			var change *types.PropertyChange
			cb := func(updates []types.ObjectUpdate) bool {
				parked()
				update := updates[0]
				if update.Kind != types.ObjectUpdateKindModify {
					return false
				}
				change = &update.ChangeSet[0]
				if change.Op != test.Op {
					t.Logf("ignore: change Op=%s, test Op=%s", change.Op, test.Op)
					return false
				}
				return true
			}

			filter := new(property.WaitFilter)
			filter.Add(v.Reference(), kind[0], []string{test.Name}, v.TraversalSpec())
			go func() {
				werr := property.WaitForUpdates(ctx, pc, filter, cb)
				if werr != nil {
					t.Log(werr)
				}
				update <- true
			}()
			<-update // wait until above go func is parked in WaitForUpdatesEx()

			opt := test.Val.(types.OptionValue)
			spec := types.VirtualMachineConfigSpec{
				ExtraConfig: []types.BaseOptionValue{&opt},
			}
			task, err := vm.Reconfigure(ctx, spec)
			if err != nil {
				t.Fatal(err)
			}
			if err := task.Wait(ctx); err != nil {
				t.Fatal(err)
			}
			<-update // wait until update is received (cb returns true)

			if change == nil {
				t.Fatal("no change")
			}

			if change.Name != test.Name {
				t.Errorf("Name: %s", change.Name)
			}

			if change.Op != test.Op {
				t.Errorf("Op: %s", change.Op)
			}

			if change.Op == types.PropertyChangeOpRemove {
				if change.Val != nil {
					t.Errorf("Val: %#v", change.Val)
				}
				continue
			}

			switch change.Val.(type) {
			case types.OptionValue:
				if !reflect.DeepEqual(change.Val, test.Val) {
					t.Errorf("change.Val: %#v", change.Val)
					t.Errorf("test.Val:   %#v", test.Val)
				}
			case string:
				if expect[change.Name] != change.Val {
					t.Errorf("Val: %s", change.Val)
				}
			default:
				t.Errorf("unexpected type: %T", change.Val)
			}
		}
	})
}

func TestVirtualMachinePropertyDevice(t *testing.T) {
	key := int32(3000)

	f := func(item string) string {
		return (&mo.Field{Path: "config.hardware.device", Key: key, Item: item}).String()
	}

	info := types.Description{Label: "cdrom-3000", Summary: "cdrom-3000"}

	tests := []types.PropertyChange{
		{Name: f(""), Val: nil, Op: types.PropertyChangeOpAdd},
		{Name: f(""), Val: nil, Op: types.PropertyChangeOpAssign},
		{Name: f("deviceInfo"), Val: info, Op: types.PropertyChangeOpAssign},
		{Name: f("deviceInfo.label"), Val: info.Label, Op: types.PropertyChangeOpAssign},
		{Name: f(""), Val: nil, Op: types.PropertyChangeOpRemove},
	}

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		kind := []string{"VirtualMachine"}

		m := view.NewManager(c)
		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			t.Fatal(err)
		}
		defer v.Destroy(ctx)

		refs, err := v.Find(ctx, kind, property.Match{})
		if err != nil {
			t.Fatal(err)
		}
		vm := object.NewVirtualMachine(c, refs[0])

		pc := property.DefaultCollector(c)

		for _, test := range tests {
			t.Logf("%s: %s", test.Op, test.Name)
			update := make(chan bool)
			parked := sync.OnceFunc(func() { update <- true })

			var change *types.PropertyChange
			cb := func(updates []types.ObjectUpdate) bool {
				parked()
				update := updates[0]
				if update.Kind != types.ObjectUpdateKindModify {
					return false
				}
				change = &update.ChangeSet[0]
				if change.Op != test.Op {
					t.Logf("ignore: change Op=%s, test Op=%s", change.Op, test.Op)
					return false
				}
				return true
			}

			filter := new(property.WaitFilter)
			filter.Add(v.Reference(), kind[0], []string{test.Name}, v.TraversalSpec())
			go func() {
				werr := property.WaitForUpdates(ctx, pc, filter, cb)
				if werr != nil {
					t.Log(werr)
				}
				update <- true
			}()
			<-update // wait until above go func is parked in WaitForUpdatesEx()

			device, err := vm.Device(ctx)
			if err != nil {
				t.Fatal(err)
			}

			switch test.Op {
			case types.PropertyChangeOpAdd:
				ide, err := device.FindIDEController("")
				if err != nil {
					t.Fatal(err)
				}
				cdrom, err := device.CreateCdrom(ide)
				if err != nil {
					t.Fatal(err)
				}
				cdrom.GetVirtualDevice().Key = key
				if err = vm.AddDevice(ctx, cdrom); err != nil {
					t.Fatal(err)
				}
			case types.PropertyChangeOpAssign:
				cdrom := device.FindByKey(key)
				if err = vm.EditDevice(ctx, cdrom); err != nil {
					t.Fatal(err)
				}
			case types.PropertyChangeOpRemove:
				cdrom := device.FindByKey(key)
				if err = vm.RemoveDevice(ctx, false, cdrom); err != nil {
					t.Fatal(err)
				}
			}
			<-update // wait until update is received (cb returns true)

			if change == nil {
				t.Fatal("no change")
			}

			if change.Name != test.Name {
				t.Errorf("Name: %s", change.Name)
			}

			if change.Op != test.Op {
				t.Errorf("Op: %s", change.Op)
			}

			if change.Op == types.PropertyChangeOpRemove {
				if change.Val != nil {
					t.Errorf("Val: %#v", change.Val)
				}
				continue
			}

			if test.Val != nil {
				if !reflect.DeepEqual(change.Val, test.Val) {
					t.Errorf("change.Val: %#v", change.Val)
					t.Errorf("test.Val:   %#v", test.Val)
				}
			} else {
				if _, ok := change.Val.(types.VirtualCdrom); !ok {
					t.Errorf("unexpected type: %T", change.Val)
				}
			}
		}
	})
}
