// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestDatacenterCreateFolders(t *testing.T) {
	// For this test we only want the RootFolder, 1 Datacenter and its child folders
	models := []Model{
		{
			ServiceContent: esx.ServiceContent,
			RootFolder:     esx.RootFolder,
		},
		{
			ServiceContent: vpx.ServiceContent,
			RootFolder:     vpx.RootFolder,
			Datacenter:     1,
		},
	}

	for _, model := range models {
		_ = model.Create()

		dc := model.Map().Any("Datacenter").(*Datacenter)

		folders := []types.ManagedObjectReference{
			dc.VmFolder,
			dc.HostFolder,
			dc.DatastoreFolder,
			dc.NetworkFolder,
		}

		for _, ref := range folders {
			if ref.Type == "" || ref.Value == "" {
				t.Errorf("invalid moref=%#v", ref)
			}

			e := model.Map().Get(ref).(mo.Entity)

			if e.Entity().Name == "" {
				t.Error("empty name")
			}

			if *e.Entity().Parent != dc.Self {
				t.Fail()
			}

			f, ok := e.(*Folder)
			if !ok {
				t.Fatalf("unexpected type (%T) for %#v", e, ref)
			}

			if model.Map().IsVPX() {
				if len(f.ChildType) < 2 {
					t.Fail()
				}
			} else {
				if len(f.ChildType) != 1 {
					t.Fail()
				}
			}
		}
	}
}

func TestDatacenterPowerOnMultiVMs(t *testing.T) {
	ctx := context.Background()

	m := VPX()
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

	finder := find.NewFinder(c.Client, false)
	dc, err := finder.DefaultDatacenter(ctx)
	if err != nil {
		t.Fatal(err)
	}

	finder.SetDatacenter(dc)

	vms, err := finder.VirtualMachineList(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}

	// Default inventory could change in future to many VMs, ensure we have
	// at least these many VMs to test.
	numTestVMs := 2
	if len(vms) < numTestVMs {
		t.Fatalf("Need at least %v VMs in a datacenter for this test", numTestVMs)
	}
	testVMs := []types.ManagedObjectReference{}
	for _, vm := range vms[:numTestVMs] {
		testVMs = append(testVMs, vm.Reference())
	}

	// Ensure VMs are powered off first before testing multi-VM power-on.
	for _, vm := range vms[:numTestVMs] {
		task, err := vm.PowerOff(ctx)
		if err != nil {
			t.Fatal(err)
		}
		err = task.Wait(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}

	// real VC ignores unknown VM refs
	unknown := types.ManagedObjectReference{Type: "VirtualMachine", Value: "unknown"}

	dcTask, err := dc.PowerOnVM(ctx, append(testVMs, unknown))
	if err != nil {
		t.Fatal(err)
	}
	info, err := dcTask.WaitForResult(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	switch dcResult := info.Result.(type) {
	case types.ClusterPowerOnVmResult:
		if len(dcResult.Attempted) != len(testVMs) {
			t.Fatalf("Unexpected per-vm tasks in results, found %v, expected %v",
				len(dcResult.Attempted), len(testVMs))
		}
		for i, vmResult := range dcResult.Attempted {
			if vmResult.Task == nil {
				t.Fatalf("Found per-vm task nil for VM #%v", i)
			}
			vmTask := object.NewTask(c.Client, *vmResult.Task)
			err := vmTask.Wait(ctx)
			if err != nil {
				t.Fatalf("%v", err)
			}
		}
	default:
		t.Fatalf("Unexpected result type %T returned for DC PowerOnMultiVM", dcResult)
	}
}
