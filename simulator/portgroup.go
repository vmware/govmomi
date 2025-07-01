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

type DistributedVirtualPortgroup struct {
	mo.DistributedVirtualPortgroup
}

func (p *DistributedVirtualPortgroup) event(ctx *Context) types.DVPortgroupEvent {
	dvs := ctx.Map.Get(*p.Config.DistributedVirtualSwitch).(*DistributedVirtualSwitch)

	return types.DVPortgroupEvent{
		Event: types.Event{
			Datacenter: datacenterEventArgument(ctx, p),
			Net: &types.NetworkEventArgument{
				EntityEventArgument: types.EntityEventArgument{
					Name: p.Name,
				},
				Network: p.Self,
			},
			Dvs: dvs.eventArgument(),
		},
	}
}

func (s *DistributedVirtualPortgroup) RenameTask(ctx *Context, req *types.Rename_Task) soap.HasFault {
	canDup := s.DistributedVirtualPortgroup.Config.BackingType == string(types.DistributedVirtualPortgroupBackingTypeNsx)

	return RenameTask(ctx, s, req, canDup)
}

func (s *DistributedVirtualPortgroup) ReconfigureDVPortgroupTask(ctx *Context, req *types.ReconfigureDVPortgroup_Task) soap.HasFault {
	task := CreateTask(s, "reconfigureDvPortgroup", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		s.Config.DefaultPortConfig = req.Spec.DefaultPortConfig
		s.Config.NumPorts = req.Spec.NumPorts
		s.Config.AutoExpand = req.Spec.AutoExpand
		s.Config.Type = req.Spec.Type
		s.Config.Description = req.Spec.Description
		s.Config.DynamicData = req.Spec.DynamicData
		s.Config.Name = req.Spec.Name
		s.Config.Policy = req.Spec.Policy
		s.Config.PortNameFormat = req.Spec.PortNameFormat
		s.Config.VmVnicNetworkResourcePoolKey = req.Spec.VmVnicNetworkResourcePoolKey
		s.Config.LogicalSwitchUuid = req.Spec.LogicalSwitchUuid
		s.Config.BackingType = req.Spec.BackingType

		return nil, nil
	})

	return &methods.ReconfigureDVPortgroup_TaskBody{
		Res: &types.ReconfigureDVPortgroup_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *DistributedVirtualPortgroup) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
	task := CreateTask(s, "destroy", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		vswitch := ctx.Map.Get(*s.Config.DistributedVirtualSwitch).(*DistributedVirtualSwitch)
		ctx.Map.RemoveReference(ctx, vswitch, &vswitch.Portgroup, s.Reference())
		ctx.Map.removeString(ctx, vswitch, &vswitch.Summary.PortgroupName, s.Name)

		f := ctx.Map.getEntityParent(vswitch, "Folder").(*Folder)
		folderRemoveChild(ctx, &f.Folder, s.Reference())
		ctx.postEvent(&types.DVPortgroupDestroyedEvent{DVPortgroupEvent: s.event(ctx)})

		return nil, nil
	})

	return &methods.Destroy_TaskBody{
		Res: &types.Destroy_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}

}
