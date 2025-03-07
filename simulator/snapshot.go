// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"os"
	"path"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualMachineSnapshot struct {
	mo.VirtualMachineSnapshot
	DataSets map[string]*DataSet
}

func (v *VirtualMachineSnapshot) createSnapshotFiles(ctx *Context) types.BaseMethodFault {
	vm := ctx.Map.Get(v.Vm).(*VirtualMachine)

	snapshotDirectory := vm.Config.Files.SnapshotDirectory
	if snapshotDirectory == "" {
		snapshotDirectory = vm.Config.Files.VmPathName
	}

	index := 1
	for {
		fileName := fmt.Sprintf("%s-Snapshot%d.vmsn", vm.Name, index)
		f, err := vm.createFile(ctx, snapshotDirectory, fileName, false)
		if err != nil {
			switch err.(type) {
			case *types.FileAlreadyExists:
				index++
				continue
			default:
				return err
			}
		}

		_ = f.Close()

		p, _ := parseDatastorePath(snapshotDirectory)
		vm.useDatastore(ctx, p.Datastore)
		datastorePath := object.DatastorePath{
			Datastore: p.Datastore,
			Path:      path.Join(p.Path, fileName),
		}

		dataLayoutKey := vm.addFileLayoutEx(ctx, datastorePath, 0)
		vm.addSnapshotLayout(ctx, v.Self, dataLayoutKey)
		vm.addSnapshotLayoutEx(ctx, v.Self, dataLayoutKey, -1)

		return nil
	}
}

func (v *VirtualMachineSnapshot) removeSnapshotFiles(ctx *Context) types.BaseMethodFault {
	// TODO: also remove delta disks that were created when snapshot was taken

	vm := ctx.Map.Get(v.Vm).(*VirtualMachine)

	for idx, sLayout := range vm.Layout.Snapshot {
		if sLayout.Key == v.Self {
			vm.Layout.Snapshot = append(vm.Layout.Snapshot[:idx], vm.Layout.Snapshot[idx+1:]...)
			break
		}
	}

	for idx, sLayoutEx := range vm.LayoutEx.Snapshot {
		if sLayoutEx.Key == v.Self {
			for _, file := range vm.LayoutEx.File {
				if file.Key == sLayoutEx.DataKey || file.Key == sLayoutEx.MemoryKey {
					p, fault := parseDatastorePath(file.Name)
					if fault != nil {
						return fault
					}

					host := ctx.Map.Get(*vm.Runtime.Host).(*HostSystem)
					datastore := ctx.Map.FindByName(p.Datastore, host.Datastore).(*Datastore)
					dFilePath := datastore.resolve(ctx, p.Path)

					_ = os.Remove(dFilePath)
				}
			}

			vm.LayoutEx.Snapshot = append(vm.LayoutEx.Snapshot[:idx], vm.LayoutEx.Snapshot[idx+1:]...)
		}
	}

	vm.RefreshStorageInfo(ctx, nil)

	return nil
}

func (v *VirtualMachineSnapshot) RemoveSnapshotTask(ctx *Context, req *types.RemoveSnapshot_Task) soap.HasFault {
	task := CreateTask(v.Vm, "removeSnapshot", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		var changes []types.PropertyChange

		vm := ctx.Map.Get(v.Vm).(*VirtualMachine)
		ctx.WithLock(vm, func() {
			if vm.Snapshot.CurrentSnapshot != nil && *vm.Snapshot.CurrentSnapshot == req.This {
				parent := findParentSnapshotInTree(vm.Snapshot.RootSnapshotList, req.This)
				changes = append(changes, types.PropertyChange{Name: "snapshot.currentSnapshot", Val: parent})
			}

			rootSnapshots := removeSnapshotInTree(vm.Snapshot.RootSnapshotList, req.This, req.RemoveChildren)
			changes = append(changes, types.PropertyChange{Name: "snapshot.rootSnapshotList", Val: rootSnapshots})

			rootSnapshotRefs := make([]types.ManagedObjectReference, len(rootSnapshots))
			for i, rs := range rootSnapshots {
				rootSnapshotRefs[i] = rs.Snapshot
			}
			changes = append(changes, types.PropertyChange{Name: "rootSnapshot", Val: rootSnapshotRefs})

			if len(rootSnapshots) == 0 {
				changes = []types.PropertyChange{
					{Name: "snapshot", Val: nil},
					{Name: "rootSnapshot", Val: nil},
				}
			}

			ctx.Map.Get(req.This).(*VirtualMachineSnapshot).removeSnapshotFiles(ctx)

			ctx.Update(vm, changes)
		})

		ctx.Map.Remove(ctx, req.This)

		return nil, nil
	})

	return &methods.RemoveSnapshot_TaskBody{
		Res: &types.RemoveSnapshot_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (v *VirtualMachineSnapshot) RevertToSnapshotTask(ctx *Context, req *types.RevertToSnapshot_Task) soap.HasFault {
	task := CreateTask(v.Vm, "revertToSnapshot", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		vm := ctx.Map.Get(v.Vm).(*VirtualMachine)

		ctx.WithLock(vm, func() {
			vm.DataSets = copyDataSetsForVmClone(v.DataSets)
			ctx.Update(vm, []types.PropertyChange{
				{Name: "snapshot.currentSnapshot", Val: v.Self},
			})
		})

		return nil, nil
	})

	return &methods.RevertToSnapshot_TaskBody{
		Res: &types.RevertToSnapshot_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
