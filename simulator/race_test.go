// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestRace(t *testing.T) {
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

	content := c.Client.ServiceContent

	wctx, cancel := context.WithCancel(ctx)
	var wg, collectors sync.WaitGroup

	nevents := -1
	em := event.NewManager(c.Client)

	wg.Add(1)
	collectors.Add(1)
	go func() {
		defer collectors.Done()

		werr := em.Events(wctx, []types.ManagedObjectReference{content.RootFolder}, 50, true, false,
			func(_ types.ManagedObjectReference, e []types.BaseEvent) error {
				if nevents == -1 {
					wg.Done() // make sure we are called at least once before cancel() below
					nevents = 0
				}

				nevents += len(e)
				return nil
			})
		if werr != nil {
			t.Error(werr)
		}
	}()

	collectors.Add(1)
	go func() {
		defer collectors.Done()

		ec, werr := em.CreateCollectorForEvents(ctx, types.EventFilterSpec{})
		if werr != nil {
			t.Error(werr)
		}

		n := 0
		for {
			events, werr := ec.ReadNextEvents(ctx, 10)
			if werr != nil {
				t.Error(werr)
			}

			n += len(events)
			if len(events) != 0 {
				continue
			}

			select {
			case <-wctx.Done():
				logf := t.Logf
				if n == 0 {
					logf = t.Errorf
				}
				logf("ReadNextEvents=%d", n)
				return
			case <-time.After(time.Millisecond * 100):
			}
		}
	}()

	ntasks := -1
	tv, err := view.NewManager(c.Client).CreateTaskView(ctx, content.TaskManager)
	if err != nil {
		t.Fatal(err)
	}

	lv, err := view.NewManager(c.Client).CreateListView(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)
	collectors.Add(1)
	go func() {
		defer collectors.Done()

		werr := tv.Collect(ctx, func(tasks []types.TaskInfo) {
			if ntasks == -1 {
				wg.Done() // make sure we are called at least once before cancel() below
				ntasks = 0
			}
			ntasks += len(tasks)
		})
		if werr != nil {
			t.Error(werr)
		}
	}()

	for i := 0; i < 2; i++ {
		spec := types.VirtualMachineConfigSpec{
			Name:    fmt.Sprintf("race-test-%d", i),
			GuestId: string(types.VirtualMachineGuestOsIdentifierOtherGuest),
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0]",
			},
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			finder := find.NewFinder(c.Client, false)
			pc := property.DefaultCollector(c.Client)
			dc, err := finder.DefaultDatacenter(ctx)
			if err != nil {
				t.Error(err)
			}

			finder.SetDatacenter(dc)

			f, err := dc.Folders(ctx)
			if err != nil {
				t.Error(err)
			}

			pool, err := finder.ResourcePool(ctx, "DC0_C0/Resources")
			if err != nil {
				t.Error(err)
			}

			for j := 0; j < 2; j++ {
				cspec := spec // copy spec and give it a unique name
				cspec.Name += fmt.Sprintf("-%d", j)

				wg.Add(1)
				go func() {
					defer wg.Done()

					task, _ := f.VmFolder.CreateVM(ctx, cspec, pool, nil)
					r, terr := task.WaitForResult(ctx, nil)
					if terr != nil {
						t.Error(terr)
					}
					_, terr = lv.Add(ctx, []types.ManagedObjectReference{r.Result.(types.ManagedObjectReference)})
					if terr != nil {
						t.Error(terr)
					}
				}()
			}

			vms, err := finder.VirtualMachineList(ctx, "*")
			if err != nil {
				t.Error(err)
			}

			for i := range vms {
				props := []string{"runtime.powerState"}
				vm := vms[i]

				wg.Add(1)
				go func() {
					defer wg.Done()

					werr := property.Wait(ctx, pc, vm.Reference(), props, func(changes []types.PropertyChange) bool {
						for _, change := range changes {
							if change.Name != props[0] {
								t.Errorf("unexpected property: %s", change.Name)
							}
							if change.Val == types.VirtualMachinePowerStatePoweredOff {
								return true
							}
						}

						wg.Add(1)
						time.AfterFunc(100*time.Millisecond, func() {
							defer wg.Done()

							_, _ = lv.Remove(ctx, []types.ManagedObjectReference{vm.Reference()})
							task, _ := vm.PowerOff(ctx)
							_ = task.Wait(ctx)
						})

						return false

					})
					if werr != nil {
						if werr != context.Canceled {
							t.Error(werr)
						}
					}
				}()
			}
		}()
	}

	wg.Wait()

	// cancel event and tasks collectors, waiting for them to complete
	cancel()
	collectors.Wait()

	t.Logf("collected %d events, %d tasks", nevents, ntasks)
	if nevents == 0 {
		t.Error("no events collected")
	}
	if ntasks == 0 {
		t.Error("no tasks collected")
	}
}

// Race VirtualMachine.Destroy vs Folder.Destroy; the latter removes the VM parent and the former should not panic in that case
func TestRaceDestroy(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	m.Folder = 1
	m.Autostart = false
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

	finder := find.NewFinder(c.Client)
	vms, err := finder.VirtualMachineList(ctx, "*")
	if err != nil {
		t.Fatal(err)
	}

	folder, err := finder.Folder(ctx, "vm/F0")
	if err != nil {
		t.Fatal(err)
	}

	notFound := false
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, vm := range vms {
			task, err := vm.Destroy(ctx)
			if err != nil {
				t.Error(err)
			}
			err = task.Wait(ctx)
			if err != nil {
				if fault.Is(err, &types.ManagedObjectNotFound{}) {
					notFound = true
				} else {
					t.Error(err)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		task, err := folder.Destroy(ctx)
		if err != nil {
			t.Error(err)
		}
		err = task.Wait(ctx)
		if err != nil {
			t.Error(err)
		}
	}()

	wg.Wait()

	if !notFound {
		t.Error("expected ManagedObjectNotFound")
	}
}

// TestWaitForUpdatesPrefixMatchRace reproduces a data race in
// PropertyFilter.matches: when a change name matches a parent prefix of the
// filter's PathSet (e.g. PathSet ["config"] matching a "config.hardware.device"
// update), the WaitForUpdatesEx goroutine re-reads the object's live property
// state via fieldValue without holding the object's lock, racing with a
// ReconfigVM_Task goroutine mutating the same VM under its lock.
func TestWaitForUpdatesPrefixMatchRace(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)

		dc, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatal(err)
		}

		finder.SetDatacenter(dc)

		folders, err := dc.Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		pool, err := finder.ResourcePool(ctx, "DC0_H0/Resources")
		if err != nil {
			t.Fatal(err)
		}

		spec := types.VirtualMachineConfigSpec{
			Name:    "prefix-match-race",
			GuestId: string(types.VirtualMachineGuestOsIdentifierOtherGuest),
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0]",
			},
		}

		task, err := folders.VmFolder.CreateVM(ctx, spec, pool, nil)
		if err != nil {
			t.Fatal(err)
		}

		info, err := task.WaitForResult(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		// The VM is intentionally left powered off: ReconfigVM_Task then
		// mutates the device list in place.
		vm := object.NewVirtualMachine(c, info.Result.(types.ManagedObjectReference))

		wctx, cancel := context.WithCancel(ctx)
		var watcher sync.WaitGroup

		watcher.Add(1)
		go func() {
			defer watcher.Done()

			pc := property.DefaultCollector(c)
			// The parent field "config" forces the PropertyFilter.matches
			// prefix path to re-read the whole config subtree on every
			// "config.*" update.
			werr := property.Wait(wctx, pc, vm.Reference(), []string{"config"}, func([]types.PropertyChange) bool {
				return false // consume updates until wctx is canceled
			})
			if werr != nil && !errors.Is(werr, context.Canceled) {
				t.Error(werr)
			}
		}()

		var writers sync.WaitGroup

		// Add and remove a NIC in a loop; the simulator assigns the
		// controller and MAC address, mutating config.hardware.device
		// during the task, after property updates have been published.
		writers.Add(1)
		go func() {
			defer writers.Done()

			for i := 0; i < 100; i++ {
				nic := &types.VirtualE1000{
					VirtualEthernetCard: types.VirtualEthernetCard{
						VirtualDevice: types.VirtualDevice{
							Backing: &types.VirtualEthernetCardNetworkBackingInfo{
								VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
									DeviceName: "VM Network",
								},
							},
						},
					},
				}

				if err := vm.AddDevice(ctx, nic); err != nil {
					t.Error(err)
					return
				}

				devices, err := vm.Device(ctx)
				if err != nil {
					t.Error(err)
					return
				}

				cards := devices.SelectByType((*types.VirtualEthernetCard)(nil))
				if len(cards) == 0 {
					t.Error("no ethernet card found")
					return
				}

				if err := vm.RemoveDevice(ctx, false, cards[len(cards)-1]); err != nil {
					t.Error(err)
					return
				}
			}
		}()

		// Reconfig ExtraConfig in a loop, mutating config.extraConfig on
		// the same VM.
		writers.Add(1)
		go func() {
			defer writers.Done()

			for i := 0; i < 100; i++ {
				rtask, err := vm.Reconfigure(ctx, types.VirtualMachineConfigSpec{
					ExtraConfig: []types.BaseOptionValue{
						&types.OptionValue{Key: "race.test", Value: fmt.Sprintf("%d", i)},
					},
				})
				if err != nil {
					t.Error(err)
					return
				}
				if err := rtask.Wait(ctx); err != nil {
					t.Error(err)
					return
				}
			}
		}()

		writers.Wait()
		cancel()
		watcher.Wait()
	})
}

func TestRaceVmRelocate(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)

		dc, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatal(err)
		}

		finder.SetDatacenter(dc)

		vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			t.Fatal(err)
		}

		folders, err := dc.Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		vmFolder := folders.VmFolder

		var failed atomic.Int32
		var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			spec := types.VirtualMachineRelocateSpec{
				Folder: types.NewReference(vmFolder.Reference()),
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				task, err := vm.Relocate(ctx, spec, types.VirtualMachineMovePriorityDefaultPriority)
				if err != nil {
					panic(err)
				}
				if err = task.Wait(ctx); err != nil {
					failed.Add(1)
				}
			}()

			vmFolder, err = vmFolder.CreateFolder(ctx, "child")
			if err != nil {
				t.Fatal("Failed to create folder", err)
			}
		}

		wg.Wait()

		if n := failed.Load(); n != 0 {
			t.Errorf("%d relocate calls failed", n)
		}
	})
}
