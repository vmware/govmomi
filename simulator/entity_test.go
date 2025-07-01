// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"testing"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func TestRename(t *testing.T) {
	m := VPX()
	m.Datacenter = 2
	m.Folder = 2

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	ctx := m.Service.Context
	dc := ctx.Map.Any("Datacenter").(*Datacenter)
	vmFolder := ctx.Map.Get(dc.VmFolder).(*Folder)

	f1 := ctx.Map.Get(vmFolder.ChildEntity[0]).(*Folder) // "F1"

	id := vmFolder.CreateFolder(ctx, &types.CreateFolder{
		This: vmFolder.Reference(),
		Name: "F2",
	}).(*methods.CreateFolderBody).Res.Returnval

	f2 := ctx.Map.Get(id).(*Folder) // "F2"

	states := []types.TaskInfoState{types.TaskInfoStateError, types.TaskInfoStateSuccess}
	name := f1.Name

	for _, expect := range states {
		id = f2.RenameTask(ctx, &types.Rename_Task{
			This:    f2.Reference(),
			NewName: name,
		}).(*methods.Rename_TaskBody).Res.Returnval

		task := ctx.Map.Get(id).(*Task)
		task.Wait()

		if task.Info.State != expect {
			t.Errorf("state=%s", task.Info.State)
		}

		name += "-uniq"
	}
}
