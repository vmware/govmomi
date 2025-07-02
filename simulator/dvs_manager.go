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

type DistributedVirtualSwitchManager struct {
	mo.DistributedVirtualSwitchManager
}

func (m *DistributedVirtualSwitchManager) DVSManagerLookupDvPortGroup(ctx *Context, req *types.DVSManagerLookupDvPortGroup) soap.HasFault {
	body := &methods.DVSManagerLookupDvPortGroupBody{}

	for _, obj := range ctx.Map.All("DistributedVirtualSwitch") {
		dvs := obj.(*DistributedVirtualSwitch)
		if dvs.Uuid == req.SwitchUuid {
			for _, ref := range dvs.Portgroup {
				pg := ctx.Map.Get(ref).(*DistributedVirtualPortgroup)
				if pg.Key == req.PortgroupKey {
					body.Res = &types.DVSManagerLookupDvPortGroupResponse{
						Returnval: &ref,
					}
					return body
				}
			}
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}
