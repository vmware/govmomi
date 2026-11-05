// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"strconv"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type DistributedVirtualPortgroup struct {
	mo.DistributedVirtualPortgroup
}

func (p *DistributedVirtualPortgroup) event(ctx *Context) types.DVPortgroupEvent {
	dvs := ctx.Map.Get(*p.Config.DistributedVirtualSwitch).(*VmwareDistributedVirtualSwitch)

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
		config := s.Config
		spec := &req.Spec
		copyNonEmptyValue(&config.NumPorts, &spec.NumPorts)
		copyNonEmptyValue(&config.AutoExpand, &spec.AutoExpand)
		copyNonEmptyValue(&config.Type, &spec.Type)
		copyNonEmptyValue(&config.Description, &spec.Description)
		copyNonEmptyValue(&config.Name, &spec.Name)
		copyNonEmptyValue(&config.Policy, &spec.Policy)
		copyNonEmptyValue(&config.PortNameFormat, &spec.PortNameFormat)
		copyNonEmptyValue(&config.VmVnicNetworkResourcePoolKey, &spec.VmVnicNetworkResourcePoolKey)
		copyNonEmptyValue(&config.LogicalSwitchUuid, &spec.LogicalSwitchUuid)
		copyNonEmptyValue(&config.SubnetId, &spec.SubnetId)
		copyNonEmptyValue(&config.BackingType, &spec.BackingType)
		if spec.DefaultPortConfig != nil {
			config.DefaultPortConfig = spec.DefaultPortConfig
			current, currentOK := s.Config.DefaultPortConfig.(*types.VMwareDVSPortSetting)
			setting, settingOK := spec.DefaultPortConfig.(*types.VMwareDVSPortSetting)
			if currentOK && settingOK {
				// Port settings are partial updates: an IPFIX-only spec must retain VLAN,
				// security and teaming policies. Decode into a copy to retain omitted fields.
				var merged types.VMwareDVSPortSetting
				deepCopy(current, &merged)
				deepCopy(setting, &merged)
				config.DefaultPortConfig = &merged
			}
		}
		version, _ := strconv.Atoi(config.ConfigVersion)
		config.ConfigVersion = strconv.Itoa(version + 1)
		ctx.Update(s, []types.PropertyChange{{Name: "config", Val: config}})
		// Invalidate FetchDVPorts cache: DVPG reconfiguration may change port structure
		if dvs, ok := ctx.Map.Get(*s.Config.DistributedVirtualSwitch).(*VmwareDistributedVirtualSwitch); ok {
			dvs.FetchDVPortsResponse.Returnval = nil
		}

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
		vswitch := ctx.Map.Get(*s.Config.DistributedVirtualSwitch).(*VmwareDistributedVirtualSwitch)
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
