// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package property_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestWaitForUpdatesEx(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 1
	model.Cluster = 0
	model.Pool = 0
	model.Machine = 1
	model.Autostart = false

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		// Set up the finder and get a VM.
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("default datacenter not found: %s", err)
		}
		finder.SetDatacenter(datacenter)
		vmList, err := finder.VirtualMachineList(ctx, "*")
		if len(vmList) == 0 {
			t.Fatal("vmList == 0")
		}
		vm := vmList[0]

		pc, err := property.DefaultCollector(c).Create(ctx)
		if err != nil {
			t.Fatalf("failed to create new property collector: %s", err)
		}

		// Start a goroutine to wait for power state changes to the VM.
		chanResult := make(chan any)
		cancelCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			defer close(chanResult)
			if err := property.WaitForUpdatesEx(
				cancelCtx,
				pc,
				&property.WaitFilter{
					CreateFilter: types.CreateFilter{
						Spec: getDatacenterToVMFolderFilter(datacenter),
					},
					WaitOptions: property.WaitOptions{
						Options: &types.WaitOptions{
							MaxWaitSeconds: addrOf(int32(3)),
						},
					},
				},
				func(updates []types.ObjectUpdate) bool {
					return waitForPowerStateChanges(
						cancelCtx,
						vm,
						chanResult,
						updates,
						types.VirtualMachinePowerStatePoweredOn)
				},
			); err != nil {
				chanResult <- err
				return
			}
		}()

		// Power on the VM to cause a property change.
		if _, err := vm.PowerOn(ctx); err != nil {
			t.Fatalf("error while powering on vm: %s", err)
		}

		select {
		case <-time.After(3 * time.Second):
			t.Fatalf("timed out while waiting for property update")
		case result := <-chanResult:
			switch tResult := result.(type) {
			case types.VirtualMachinePowerState:
				if tResult != types.VirtualMachinePowerStatePoweredOn {
					t.Fatalf("unexpected power state: %s", tResult)
				}
			case error:
				t.Fatalf("error while waiting for updates: %s", tResult)
			}
		}
	}, model)
}

func TestWaitForUpdatesExEmptyUpdateSet(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 1
	model.Cluster = 0
	model.Pool = 0
	model.Machine = 0
	model.Autostart = false

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		// Set up the finder and get a VM.
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("default datacenter not found: %s", err.Error())
		}
		finder.SetDatacenter(datacenter)
		vmList, err := finder.VirtualMachineList(ctx, "*")
		if len(vmList) != 0 {
			t.Fatalf("vmList != 0")
		}

		pc, err := property.DefaultCollector(c).Create(ctx)
		if err != nil {
			t.Fatalf("failed to create new property collector: %s", err.Error())
		}

		// Start a goroutine to wait for updates on any VMs.
		// Since there are no VMs in the filter set, we expect to
		// receive an empty update set.
		chanResult := make(chan error)
		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			defer close(chanResult)
			_ = pc.WaitForUpdatesEx(
				cancelCtx,
				&property.WaitOptions{},
				func(updates []types.ObjectUpdate) bool {
					var err error
					if len(updates) > 0 {
						err = fmt.Errorf("unexpected update")
					}
					chanResult <- err
					cancel()
					return true
				})
		}()

		select {
		case <-ctx.Done():
			t.Fatalf("timed out while waiting for updates")
		case err := <-chanResult:
			if err != nil {
				t.Fatalf("error while waiting for updates: %s", err.Error())
			}
		}
	}, model)
}

func TestRetrievePropertiesOneAtATime(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 1
	model.Cluster = 0
	model.Pool = 0
	model.Machine = 3
	model.Autostart = false

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("default datacenter not found: %s", err)
		}
		finder.SetDatacenter(datacenter)
		pc := property.DefaultCollector(c)

		resp, err := pc.RetrieveProperties(ctx, types.RetrieveProperties{
			SpecSet: []types.PropertyFilterSpec{
				getDatacenterToVMFolderFilter(datacenter),
			},
		}, 1)
		if err != nil {
			t.Fatalf("failed to retrieve properties one object at a time: %s", err)
		}

		vmRefs := map[types.ManagedObjectReference]struct{}{}
		for i := range resp.Returnval {
			oc := resp.Returnval[i]
			vmRefs[oc.Obj] = struct{}{}
		}

		if a, e := len(vmRefs), 3; a != 3 {
			t.Fatalf("unexpected number of vms: a=%d, e=%d", a, e)
		}

	}, model)
}

func waitForPowerStateChanges(
	ctx context.Context,
	vm *object.VirtualMachine,
	chanResult chan any,
	updates []types.ObjectUpdate,
	expectedPowerState types.VirtualMachinePowerState) bool {

	for _, u := range updates {
		if ctx.Err() != nil {
			return false
		}
		if u.Obj != vm.Reference() {
			continue
		}
		for _, cs := range u.ChangeSet {
			if cs.Name == "runtime.powerState" {
				if cs.Val == expectedPowerState {
					select {
					case <-ctx.Done():
						// No-op
					default:
						chanResult <- cs.Val
					}
					return true
				}
			}
		}
	}
	return false
}

func getDatacenterToVMFolderFilter(dc *object.Datacenter) types.PropertyFilterSpec {
	// Define a wait filter that looks for updates to VM power
	// states for VMs under the specified datacenter.
	return types.PropertyFilterSpec{
		ObjectSet: []types.ObjectSpec{
			{
				Obj:  dc.Reference(),
				Skip: addrOf(true),
				SelectSet: []types.BaseSelectionSpec{
					// Datacenter --> VM folder
					&types.TraversalSpec{
						SelectionSpec: types.SelectionSpec{
							Name: "dcToVMFolder",
						},
						Type: "Datacenter",
						Path: "vmFolder",
						SelectSet: []types.BaseSelectionSpec{
							&types.SelectionSpec{
								Name: "visitFolders",
							},
						},
					},
					// Folder --> children (folder / VM)
					&types.TraversalSpec{
						SelectionSpec: types.SelectionSpec{
							Name: "visitFolders",
						},
						Type: "Folder",
						// Folder --> children (folder / VM)
						Path: "childEntity",
						SelectSet: []types.BaseSelectionSpec{
							// Folder --> child folder
							&types.SelectionSpec{
								Name: "visitFolders",
							},
						},
					},
				},
			},
		},
		PropSet: []types.PropertySpec{
			{
				Type:    "VirtualMachine",
				PathSet: []string{"runtime.powerState"},
			},
		},
	}
}
