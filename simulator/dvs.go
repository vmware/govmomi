// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type DistributedVirtualSwitch struct {
	mo.DistributedVirtualSwitch

	types.FetchDVPortsResponse
}

func (s *DistributedVirtualSwitch) eventArgument() *types.DvsEventArgument {
	return &types.DvsEventArgument{
		EntityEventArgument: types.EntityEventArgument{
			Name: s.Name,
		},
		Dvs: s.Self,
	}
}

func (s *DistributedVirtualSwitch) event(ctx *Context) types.DvsEvent {
	return types.DvsEvent{
		Event: types.Event{
			Datacenter: datacenterEventArgument(ctx, s),
			Dvs:        s.eventArgument(),
		},
	}
}

func (s *DistributedVirtualSwitch) AddDVPortgroupTask(ctx *Context, c *types.AddDVPortgroup_Task) soap.HasFault {
	task := CreateTask(s, "addDVPortgroup", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		f := ctx.Map.getEntityParent(s, "Folder").(*Folder)

		portgroups := s.Portgroup
		portgroupNames := s.Summary.PortgroupName

		for _, spec := range c.Spec {
			pg := &DistributedVirtualPortgroup{}
			pg.Name = spec.Name
			pg.Entity().Name = pg.Name

			// Standard AddDVPortgroupTask() doesn't allow duplicate names, but NSX 3.0 does create some DVPGs with the same name.
			// Allow duplicate names using this prefix so we can reproduce and test this condition.
			if strings.HasPrefix(pg.Name, "NSX-") || spec.BackingType == string(types.DistributedVirtualPortgroupBackingTypeNsx) {
				if spec.LogicalSwitchUuid == "" {
					spec.LogicalSwitchUuid = uuid.New().String()
				}
				if spec.SegmentId == "" {
					spec.SegmentId = fmt.Sprintf("/infra/segments/vnet_%s", uuid.New().String())
				}

			} else {
				if obj := ctx.Map.FindByName(pg.Name, f.ChildEntity); obj != nil {
					return nil, &types.DuplicateName{
						Name:   pg.Name,
						Object: obj.Reference(),
					}
				}
			}

			folderPutChild(ctx, &f.Folder, pg)

			pg.Key = pg.Self.Value
			pg.Config = types.DVPortgroupConfigInfo{
				Key:                          pg.Key,
				Name:                         pg.Name,
				NumPorts:                     spec.NumPorts,
				DistributedVirtualSwitch:     &s.Self,
				DefaultPortConfig:            spec.DefaultPortConfig,
				Description:                  spec.Description,
				Type:                         spec.Type,
				Policy:                       spec.Policy,
				PortNameFormat:               spec.PortNameFormat,
				Scope:                        spec.Scope,
				VendorSpecificConfig:         spec.VendorSpecificConfig,
				ConfigVersion:                spec.ConfigVersion,
				AutoExpand:                   spec.AutoExpand,
				VmVnicNetworkResourcePoolKey: spec.VmVnicNetworkResourcePoolKey,
				LogicalSwitchUuid:            spec.LogicalSwitchUuid,
				SegmentId:                    spec.SegmentId,
				BackingType:                  spec.BackingType,
			}

			if pg.Config.LogicalSwitchUuid != "" {
				if pg.Config.BackingType == "" {
					pg.Config.BackingType = "nsx"
				}
			}

			if pg.Config.DefaultPortConfig == nil {
				pg.Config.DefaultPortConfig = &types.VMwareDVSPortSetting{
					Vlan: new(types.VmwareDistributedVirtualSwitchVlanIdSpec),
					UplinkTeamingPolicy: &types.VmwareUplinkPortTeamingPolicy{
						Policy: &types.StringPolicy{
							Value: "loadbalance_srcid",
						},
						ReversePolicy: &types.BoolPolicy{
							Value: types.NewBool(true),
						},
						NotifySwitches: &types.BoolPolicy{
							Value: types.NewBool(true),
						},
						RollingOrder: &types.BoolPolicy{
							Value: types.NewBool(true),
						},
					},
				}
			}

			if pg.Config.Policy == nil {
				pg.Config.Policy = &types.VMwareDVSPortgroupPolicy{
					DVPortgroupPolicy: types.DVPortgroupPolicy{
						BlockOverrideAllowed:               true,
						ShapingOverrideAllowed:             false,
						VendorConfigOverrideAllowed:        false,
						LivePortMovingAllowed:              false,
						PortConfigResetAtDisconnect:        true,
						NetworkResourcePoolOverrideAllowed: types.NewBool(false),
						TrafficFilterOverrideAllowed:       types.NewBool(false),
					},
					VlanOverrideAllowed:           false,
					UplinkTeamingOverrideAllowed:  false,
					SecurityPolicyOverrideAllowed: false,
					IpfixOverrideAllowed:          types.NewBool(false),
				}
			}

			for i := 0; i < int(spec.NumPorts); i++ {
				pg.PortKeys = append(pg.PortKeys, strconv.Itoa(i))
			}

			portgroups = append(portgroups, pg.Self)
			portgroupNames = append(portgroupNames, pg.Name)

			for _, h := range s.Summary.HostMember {
				pg.Host = append(pg.Host, h)

				host := ctx.Map.Get(h).(*HostSystem)
				ctx.Map.AppendReference(ctx, host, &host.Network, pg.Reference())

				parent := ctx.Map.Get(*host.HostSystem.Parent)
				computeNetworks := append(hostParent(ctx, &host.HostSystem).Network, pg.Reference())
				ctx.Update(parent, []types.PropertyChange{
					{Name: "network", Val: computeNetworks},
				})
			}

			ctx.postEvent(&types.DVPortgroupCreatedEvent{
				DVPortgroupEvent: pg.event(ctx),
			})
		}

		ctx.Update(s, []types.PropertyChange{
			{Name: "portgroup", Val: portgroups},
			{Name: "summary.portgroupName", Val: portgroupNames},
		})

		return nil, nil
	})

	return &methods.AddDVPortgroup_TaskBody{
		Res: &types.AddDVPortgroup_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *DistributedVirtualSwitch) ReconfigureDvsTask(ctx *Context, req *types.ReconfigureDvs_Task) soap.HasFault {
	task := CreateTask(s, "reconfigureDvs", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		spec := req.Spec.GetDVSConfigSpec()

		members := s.Summary.HostMember

		for _, member := range spec.Host {
			h := ctx.Map.Get(member.Host)
			if h == nil {
				return nil, &types.ManagedObjectNotFound{Obj: member.Host}
			}

			host := h.(*HostSystem)

			switch types.ConfigSpecOperation(member.Operation) {
			case types.ConfigSpecOperationAdd:
				if FindReference(s.Summary.HostMember, member.Host) != nil {
					return nil, &types.AlreadyExists{Name: host.Name}
				}

				hostNetworks := append(host.Network, s.Portgroup...)
				ctx.Update(host, []types.PropertyChange{
					{Name: "network", Val: hostNetworks},
				})
				members = append(members, member.Host)
				parent := ctx.Map.Get(*host.HostSystem.Parent)

				var pgs []types.ManagedObjectReference
				for _, ref := range s.Portgroup {
					pg := ctx.Map.Get(ref).(*DistributedVirtualPortgroup)
					pgs = append(pgs, ref)

					pgHosts := append(pg.Host, member.Host)
					ctx.Update(pg, []types.PropertyChange{
						{Name: "host", Val: pgHosts},
					})

					cr := hostParent(ctx, &host.HostSystem)
					if FindReference(cr.Network, ref) == nil {
						computeNetworks := append(cr.Network, ref)
						ctx.Update(parent, []types.PropertyChange{
							{Name: "network", Val: computeNetworks},
						})
					}
				}

				ctx.postEvent(&types.DvsHostJoinedEvent{
					DvsEvent:   s.event(ctx),
					HostJoined: *host.eventArgument(),
				})
			case types.ConfigSpecOperationRemove:
				for _, ref := range host.Vm {
					vm := ctx.Map.Get(ref).(*VirtualMachine)
					if pg := FindReference(vm.Network, s.Portgroup...); pg != nil {
						return nil, &types.ResourceInUse{
							Type: pg.Type,
							Name: pg.Value,
						}
					}
				}

				RemoveReference(&members, member.Host)

				ctx.postEvent(&types.DvsHostLeftEvent{
					DvsEvent: s.event(ctx),
					HostLeft: *host.eventArgument(),
				})
			case types.ConfigSpecOperationEdit:
				return nil, &types.NotSupported{}
			}
		}

		ctx.Update(s, []types.PropertyChange{
			{Name: "summary.hostMember", Val: members},
		})

		ctx.postEvent(&types.DvsReconfiguredEvent{
			DvsEvent:   s.event(ctx),
			ConfigSpec: spec,
		})

		return nil, nil
	})

	return &methods.ReconfigureDvs_TaskBody{
		Res: &types.ReconfigureDvs_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *DistributedVirtualSwitch) FetchDVPorts(ctx *Context, req *types.FetchDVPorts) soap.HasFault {
	body := &methods.FetchDVPortsBody{}
	body.Res = &types.FetchDVPortsResponse{
		Returnval: s.dvPortgroups(ctx, req.Criteria),
	}
	return body
}

func (s *DistributedVirtualSwitch) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
	task := CreateTask(s, "destroy", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		// TODO: should return ResourceInUse fault if any VM is using a port on this switch
		// and past that, remove refs from each host.Network, etc
		f := ctx.Map.getEntityParent(s, "Folder").(*Folder)
		folderRemoveChild(ctx, &f.Folder, s.Reference())
		ctx.postEvent(&types.DvsDestroyedEvent{DvsEvent: s.event(ctx)})
		return nil, nil
	})

	return &methods.Destroy_TaskBody{
		Res: &types.Destroy_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *DistributedVirtualSwitch) dvPortgroups(ctx *Context, criteria *types.DistributedVirtualSwitchPortCriteria) []types.DistributedVirtualPort {
	res := s.FetchDVPortsResponse.Returnval
	if len(res) != 0 {
		return res
	}

	for _, ref := range s.Portgroup {
		pg := ctx.Map.Get(ref).(*DistributedVirtualPortgroup)

		for _, key := range pg.PortKeys {
			res = append(res, types.DistributedVirtualPort{
				DvsUuid:      s.Uuid,
				Key:          key,
				PortgroupKey: pg.Key,
				Config: types.DVPortConfigInfo{
					Setting: pg.Config.DefaultPortConfig,
				},
			})
		}
	}

	// filter ports by criteria
	res = s.filterDVPorts(res, criteria)

	return res
}

func (s *DistributedVirtualSwitch) filterDVPorts(
	ports []types.DistributedVirtualPort,
	criteria *types.DistributedVirtualSwitchPortCriteria,
) []types.DistributedVirtualPort {
	if criteria == nil {
		return ports
	}

	ports = s.filterDVPortsByPortgroupKey(ports, criteria)
	ports = s.filterDVPortsByPortKey(ports, criteria)
	ports = s.filterDVPortsByConnected(ports, criteria)

	return ports
}

func (s *DistributedVirtualSwitch) filterDVPortsByPortgroupKey(
	ports []types.DistributedVirtualPort,
	criteria *types.DistributedVirtualSwitchPortCriteria,
) []types.DistributedVirtualPort {
	if len(criteria.PortgroupKey) == 0 || criteria.Inside == nil {
		return ports
	}

	// inside portgroup keys
	if *criteria.Inside {
		filtered := []types.DistributedVirtualPort{}

		for _, p := range ports {
			for _, pgk := range criteria.PortgroupKey {
				if p.PortgroupKey == pgk {
					filtered = append(filtered, p)
					break
				}
			}
		}
		return filtered
	}

	// outside portgroup keys
	filtered := []types.DistributedVirtualPort{}

	for _, p := range ports {
		found := false
		for _, pgk := range criteria.PortgroupKey {
			if p.PortgroupKey == pgk {
				found = true
				break
			}
		}

		if !found {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (s *DistributedVirtualSwitch) filterDVPortsByPortKey(
	ports []types.DistributedVirtualPort,
	criteria *types.DistributedVirtualSwitchPortCriteria,
) []types.DistributedVirtualPort {
	if len(criteria.PortKey) == 0 {
		return ports
	}

	filtered := []types.DistributedVirtualPort{}

	for _, p := range ports {
		for _, pk := range criteria.PortKey {
			if p.Key == pk {
				filtered = append(filtered, p)
				break
			}
		}
	}

	return filtered
}

func (s *DistributedVirtualSwitch) filterDVPortsByConnected(
	ports []types.DistributedVirtualPort,
	criteria *types.DistributedVirtualSwitchPortCriteria,
) []types.DistributedVirtualPort {
	if criteria.Connected == nil {
		return ports
	}

	filtered := []types.DistributedVirtualPort{}

	for _, p := range ports {
		connected := p.Connectee != nil
		if connected == *criteria.Connected {
			filtered = append(filtered, p)
		}
	}

	return filtered
}
