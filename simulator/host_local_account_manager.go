// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// As of vSphere API 5.1, local groups operations are deprecated, so it's not supported here.

type HostLocalAccountManager struct {
	mo.HostLocalAccountManager
}

func (h *HostLocalAccountManager) CreateUser(ctx *Context, req *types.CreateUser) soap.HasFault {
	spec := req.User.GetHostAccountSpec()
	userDirectory := ctx.Map.UserDirectory()

	found := userDirectory.search(true, false, compareFunc(spec.Id, true))
	if len(found) > 0 {
		return &methods.CreateUserBody{
			Fault_: Fault("", &types.AlreadyExists{}),
		}
	}

	userDirectory.addUser(spec.Id)

	return &methods.CreateUserBody{
		Res: &types.CreateUserResponse{},
	}
}

func (h *HostLocalAccountManager) RemoveUser(ctx *Context, req *types.RemoveUser) soap.HasFault {
	userDirectory := ctx.Map.UserDirectory()

	found := userDirectory.search(true, false, compareFunc(req.UserName, true))

	if len(found) == 0 {
		return &methods.RemoveUserBody{
			Fault_: Fault("", &types.UserNotFound{}),
		}
	}

	userDirectory.removeUser(req.UserName)

	return &methods.RemoveUserBody{
		Res: &types.RemoveUserResponse{},
	}
}

func (h *HostLocalAccountManager) UpdateUser(req *types.UpdateUser) soap.HasFault {
	return &methods.CreateUserBody{
		Res: &types.CreateUserResponse{},
	}
}
