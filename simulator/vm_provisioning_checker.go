/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
