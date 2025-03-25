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

type VmProvisioningChecker struct {
	mo.VirtualMachineProvisioningChecker
}

func (c *VmProvisioningChecker) CheckRelocateTask(
	ctx *Context,
	r *types.CheckRelocate_Task) soap.HasFault {

	task := CreateTask(c, "checkRelocate", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		check := types.CheckResult{
			Vm: &r.Vm,
		}

		return types.ArrayOfCheckResult{
			CheckResult: []types.CheckResult{check},
		}, nil
	})

	return &methods.CheckRelocate_TaskBody{
		Res: &types.CheckRelocate_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
