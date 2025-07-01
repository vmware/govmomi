// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
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
