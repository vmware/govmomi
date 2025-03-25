// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func RenameTask(ctx *Context, e mo.Entity, r *types.Rename_Task, dup ...bool) soap.HasFault {
	task := CreateTask(e, "rename", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		obj := ctx.Map.Get(r.This).(mo.Entity).Entity()

		canDup := len(dup) == 1 && dup[0]
		if parent, ok := asFolderMO(ctx.Map.Get(*obj.Parent)); ok && !canDup {
			if ctx.Map.FindByName(r.NewName, parent.ChildEntity) != nil {
				return nil, &types.InvalidArgument{InvalidProperty: "name"}
			}
		}

		ctx.Update(e, []types.PropertyChange{{Name: "name", Val: r.NewName}})

		return nil, nil
	})

	return &methods.Rename_TaskBody{
		Res: &types.Rename_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
