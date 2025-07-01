// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestDefaultESX(t *testing.T) {
	s := New(NewServiceInstance(NewContext(), esx.ServiceContent, esx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()

	client, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(client.Client, false)

	dc, err := finder.DatacenterOrDefault(ctx, "")
	if err != nil {
		t.Fatal(err)
	}

	finder.SetDatacenter(dc)

	host, err := finder.HostSystemOrDefault(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}

	if host.Name() != esx.HostSystem.Summary.Config.Name {
		t.Fail()
	}

	pool, err := finder.ResourcePoolOrDefault(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}

	if pool.Name() != "Resources" {
		t.Fail()
	}
}

func TestDefaultVPX(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		for _, e := range Map(ctx).All("HostSystem") {
			host := e.(*HostSystem)
			// issue #3221
			if host.Config.Host != host.Self {
				t.Errorf("config.host=%s", host.Config.Host)
			}
		}
	})
}

func TestMaintenanceMode(t *testing.T) {
	ctx := context.Background()
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c := m.Service.client()

	hs := m.Map().Get(esx.HostSystem.Reference()).(*HostSystem)
	host := object.NewHostSystem(c, hs.Self)

	task, err := host.EnterMaintenanceMode(ctx, 1, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if hs.Runtime.InMaintenanceMode != true {
		t.Fatal("expect InMaintenanceMode is true; got false")
	}

	task, err = host.ExitMaintenanceMode(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if hs.Runtime.InMaintenanceMode != false {
		t.Fatal("expect InMaintenanceMode is false; got true")
	}
}

func TestNewHostSystem(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	hs := NewHostSystem(m.Service.Context, esx.HostSystem)

	assert.Equal(t, &hs.Runtime, hs.Summary.Runtime, "expected pointer to runtime in summary")
	assert.False(t, esx.AdvancedOptions[0] == hs.Config.Option[0], "expected each host to have it's own advanced options")
}

func TestDestroyHostSystem(t *testing.T) {
	ctx := context.Background()
	m := VPX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c := m.Service.client()

	vm := m.Map().Any("VirtualMachine").(*VirtualMachine)
	hs := m.Map().Get(*vm.Runtime.Host).(*HostSystem)
	host := object.NewHostSystem(c, hs.Self)

	vms := []*VirtualMachine{}
	for _, vmref := range hs.Vm {
		vms = append(vms, m.Map().Get(vmref).(*VirtualMachine))
	}

	task, err := host.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err == nil {
		t.Fatal("expect err because host with vms cannot be destroyed")
	}

	for _, vmref := range hs.Vm {
		vm := m.Map().Get(vmref).(*VirtualMachine)
		vmo := object.NewVirtualMachine(c, vm.Self)

		task, err := vmo.PowerOff(ctx)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if err != nil {
			t.Fatal(err)
		}

		task, err = vmo.Destroy(ctx)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}

	task, err = host.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	hs2 := m.Map().Get(esx.HostSystem.Reference())
	if hs2 != nil {
		t.Fatal("host should have been destroyed")
	}
}

func TestDisconnect(t *testing.T) {
	ctx := context.Background()
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c := m.Service.client()

	hs := m.Map().Get(esx.HostSystem.Reference()).(*HostSystem)
	host := object.NewHostSystem(c, hs.Self)

	task, err := host.Disconnect(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if hs.Runtime.ConnectionState != types.HostSystemConnectionStateDisconnected {
		t.Fatalf("expect ConnectionState to be %s; got %s",
			types.HostSystemConnectionStateDisconnected, hs.Runtime.ConnectionState)
	}

	task, err = host.Reconnect(ctx, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if hs.Runtime.ConnectionState != types.HostSystemConnectionStateConnected {
		t.Fatalf("expect ConnectionState to be %s; got %s",
			types.HostSystemConnectionStateConnected, hs.Runtime.ConnectionState)
	}
}
