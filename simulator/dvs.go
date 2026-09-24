// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type VmwareDistributedVirtualSwitch struct {
	mo.VmwareDistributedVirtualSwitch

	types.FetchDVPortsResponse
}

func (s *VmwareDistributedVirtualSwitch) eventArgument() *types.DvsEventArgument {
	return &types.DvsEventArgument{
		EntityEventArgument: types.EntityEventArgument{
			Name: s.Name,
		},
		Dvs: s.Self,
	}
}

func (s *VmwareDistributedVirtualSwitch) event(ctx *Context) types.DvsEvent {
	return types.DvsEvent{
		Event: types.Event{
			Datacenter: datacenterEventArgument(ctx, s),
			Dvs:        s.eventArgument(),
		},
	}
}

func (s *VmwareDistributedVirtualSwitch) AddDVPortgroupTask(ctx *Context, c *types.AddDVPortgroup_Task) soap.HasFault {
	task := CreateTask(s, "addDVPortgroup", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		return nil, s.addDVPortgroups(ctx, c.Spec)
	})

	return &methods.AddDVPortgroup_TaskBody{
		Res: &types.AddDVPortgroup_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *VmwareDistributedVirtualSwitch) addDVPortgroups(ctx *Context, specs []types.DVPortgroupConfigSpec) types.BaseMethodFault {
	f := ctx.Map.getEntityParent(s, "Folder").(*Folder)

	portgroups := s.Portgroup
	portgroupNames := s.Summary.PortgroupName

	for _, spec := range specs {
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
				return &types.DuplicateName{
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
			SubnetId:                     spec.SubnetId,
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

	return nil
}

func (s *VmwareDistributedVirtualSwitch) ReconfigureDvsTask(ctx *Context, req *types.ReconfigureDvs_Task) soap.HasFault {
	task := CreateTask(s, "reconfigureDvs", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		spec := req.Spec.GetDVSConfigSpec()

		members := s.Summary.HostMember
		configHosts := slices.Clone(s.Config.GetDVSConfigInfo().Host)

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

				// Mirror what a real vCenter does on the host side when it joins
				// a DVS: record a HostProxySwitch (with its uplink pnic backing)
				// on HostSystem.Config.Network.ProxySwitch. Collectors that read
				// the host's own proxySwitch list (rather than just the DVS's
				// Summary.HostMember) need this to know which physical NICs back
				// this switch on this host.
				proxySwitch := types.HostProxySwitch{
					DvsUuid: s.Uuid,
					DvsName: s.Name,
					Key:     s.Self.Value,
					Pnic:    dvsUplinkPnics(host, member.Backing),
				}
				ctx.Update(host, []types.PropertyChange{
					{Name: "config.network.proxySwitch", Val: append(host.Config.Network.ProxySwitch, proxySwitch)},
				})

				members = append(members, member.Host)
				// A real vCenter always keeps DVSConfigInfo.Host in sync
				// with Summary.HostMember -- a client reading config.host
				// for host membership (rather than summary.hostMember)
				// otherwise sees an empty list regardless of actual
				// membership.
				hostRef := member.Host
				configHosts = append(configHosts, types.DistributedVirtualSwitchHostMember{
					Config: types.DistributedVirtualSwitchHostMemberConfigInfo{
						Host: &hostRef,
					},
				})
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
				configHosts = slices.DeleteFunc(configHosts, func(m types.DistributedVirtualSwitchHostMember) bool {
					return m.Config.Host != nil && *m.Config.Host == member.Host
				})

				proxySwitches := slices.Clone(host.Config.Network.ProxySwitch)
				proxySwitches = slices.DeleteFunc(proxySwitches, func(ps types.HostProxySwitch) bool {
					return ps.DvsUuid == s.Uuid
				})
				ctx.Update(host, []types.PropertyChange{
					{Name: "config.network.proxySwitch", Val: proxySwitches},
				})

				ctx.postEvent(&types.DvsHostLeftEvent{
					DvsEvent: s.event(ctx),
					HostLeft: *host.eventArgument(),
				})
			case types.ConfigSpecOperationEdit:
				return nil, &types.NotSupported{}
			}
		}

		config := s.Config.GetDVSConfigInfo()
		config.Host = configHosts

		ctx.Update(s, []types.PropertyChange{
			{Name: "summary.hostMember", Val: members},
			{Name: "summary.numHosts", Val: int32(len(members))},
			{Name: "config", Val: config},
		})

		// Invalidate FetchDVPorts cache: host membership changes affect uplink ports
		s.FetchDVPortsResponse.Returnval = nil

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

// dvsUplinkPnics returns the physical NIC device keys to back a host's proxy
// switch for a DVS it is joining. If the caller specified pnics via
// DistributedVirtualSwitchHostMemberPnicBacking, those are used as-is;
// otherwise it falls back to the first pnic on the host not already claimed
// by a standard vSwitch or another DVS proxy switch, matching what a real
// vCenter picks when a client adds a host to a DVS without an explicit
// uplink assignment.
func dvsUplinkPnics(host *HostSystem, backing types.BaseDistributedVirtualSwitchHostMemberBacking) []string {
	if b, ok := backing.(*types.DistributedVirtualSwitchHostMemberPnicBacking); ok {
		pnics := make([]string, 0, len(b.PnicSpec))
		for _, p := range b.PnicSpec {
			pnics = append(pnics, "key-vim.host.PhysicalNic-"+p.PnicDevice)
		}
		return pnics
	}

	used := make(map[string]bool)
	for _, vs := range host.Config.Network.Vswitch {
		for _, key := range vs.Pnic {
			used[key] = true
		}
	}
	for _, ps := range host.Config.Network.ProxySwitch {
		for _, key := range ps.Pnic {
			used[key] = true
		}
	}

	for _, pnic := range host.Config.Network.Pnic {
		if !used[pnic.Key] {
			return []string{pnic.Key}
		}
	}

	return nil
}

func (s *VmwareDistributedVirtualSwitch) FetchDVPorts(ctx *Context, req *types.FetchDVPorts) soap.HasFault {
	body := &methods.FetchDVPortsBody{}
	body.Res = &types.FetchDVPortsResponse{
		Returnval: s.dvPortgroups(ctx, req.Criteria),
	}
	return body
}

func (s *VmwareDistributedVirtualSwitch) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
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

func (s *VmwareDistributedVirtualSwitch) dvPortgroups(ctx *Context, criteria *types.DistributedVirtualSwitchPortCriteria) []types.DistributedVirtualPort {
	res := s.FetchDVPortsResponse.Returnval
	if len(res) != 0 {
		return res
	}

	uplinkPg := s.uplinkPortgroup(ctx)

	for _, ref := range s.Portgroup {
		pg := ctx.Map.Get(ref).(*DistributedVirtualPortgroup)

		if uplinkPg != nil && pg.Self == uplinkPg.Self {
			res = append(res, s.uplinkPorts(ctx, pg)...)
			continue
		}

		res = append(res, s.regularPorts(ctx, pg)...)
	}

	// filter ports by criteria
	res = s.filterDVPorts(res, criteria)

	return res
}

// uplinkPortgroup returns the DVS's auto-created uplink portgroup. It's
// identified by name rather than a stored reference: CreateDVSTask (folder.go)
// creates it via a nested AddDVPortgroupTask, which -- like all vcsim tasks --
// completes asynchronously in its own goroutine after the creating task's
// lock is released, so there's no point during DVS creation itself where the
// new portgroup's reference is reliably available yet to store.
func (s *VmwareDistributedVirtualSwitch) uplinkPortgroup(ctx *Context) *DistributedVirtualPortgroup {
	name := s.Name + "-DVUplinks" + strings.TrimPrefix(s.Self.Value, "dvs")

	for _, ref := range s.Portgroup {
		if pg, ok := ctx.Map.Get(ref).(*DistributedVirtualPortgroup); ok && pg.Name == name {
			return pg
		}
	}

	return nil
}

// uplinkPorts generates one DistributedVirtualPort per (host, pnic) pair
// backing this DVS's uplink portgroup, with Connectee populated -- mirroring
// what a real vCenter reports for physical NIC uplinks. A real vCenter's
// uplink ports are host-scoped: the same numbered port exists once per host,
// each instance connected to that host's own pnic (see
// DistributedVirtualSwitchPortCriteria.Host). vcsim tracks the pnic backing
// via HostSystem.Config.Network.ProxySwitch, populated when a host joins the
// DVS (see ReconfigureDvsTask).
func (s *VmwareDistributedVirtualSwitch) uplinkPorts(ctx *Context, pg *DistributedVirtualPortgroup) []types.DistributedVirtualPort {
	var ports []types.DistributedVirtualPort

	for _, hostRef := range s.Summary.HostMember {
		host, ok := ctx.Map.Get(hostRef).(*HostSystem)
		if !ok {
			continue
		}

		var pnics []string
		for _, ps := range host.Config.Network.ProxySwitch {
			if ps.DvsUuid == s.Uuid {
				pnics = ps.Pnic
				break
			}
		}

		for i, pnicKey := range pnics {
			key := strconv.Itoa(i)

			connectedEntity := hostRef
			ports = append(ports, types.DistributedVirtualPort{
				DvsUuid:      s.Uuid,
				Key:          key,
				PortgroupKey: pg.Key,
				Connectee: &types.DistributedVirtualSwitchPortConnectee{
					ConnectedEntity: &connectedEntity,
					NicKey:          pnicKey,
					Type:            string(types.DistributedVirtualSwitchPortConnecteeConnecteeTypePnic),
				},
				Config: types.DVPortConfigInfo{
					Setting: pg.Config.DefaultPortConfig,
				},
			})
		}
	}

	return ports
}

// regularPorts generates ports for a non-uplink portgroup. It mirrors what a
// real vCenter does when a VM's NIC connects to a DVPortgroup: one of the
// portgroup's ports gets claimed and its Connectee set to that VM's vNIC
// (DistributedVirtualSwitchPortConnectee, NicType "vmVnic") -- confirmed
// against a real vCenter's recorded/replayed inventory, which sets Connectee
// this way on every connected port, not just uplink ones. PortKeys with no
// VM claiming them stay disconnected, matching a real vCenter's unused port
// capacity.
func (s *VmwareDistributedVirtualSwitch) regularPorts(ctx *Context, pg *DistributedVirtualPortgroup) []types.DistributedVirtualPort {
	claimed := make(map[string]types.DistributedVirtualPort, len(pg.PortKeys))
	keys := slices.Clone(pg.PortKeys)
	nextKey := len(pg.PortKeys)

	for _, vmRef := range pg.Vm {
		vm, ok := ctx.Map.Get(vmRef).(*VirtualMachine)
		if !ok {
			continue
		}

		for _, d := range vm.Config.Hardware.Device {
			card, ok := d.(types.BaseVirtualEthernetCard)
			if !ok {
				continue
			}

			nic := card.GetVirtualEthernetCard()
			b, ok := nic.Backing.(*types.VirtualEthernetCardDistributedVirtualPortBackingInfo)
			if !ok || b.Port.PortgroupKey != pg.Key {
				continue
			}

			var key string
			if len(keys) > 0 {
				key = keys[0]
				keys = keys[1:]
			} else {
				key = strconv.Itoa(nextKey)
				nextKey++
			}

			connectedEntity := vmRef
			claimed[key] = types.DistributedVirtualPort{
				DvsUuid:      s.Uuid,
				Key:          key,
				PortgroupKey: pg.Key,
				Connectee: &types.DistributedVirtualSwitchPortConnectee{
					ConnectedEntity: &connectedEntity,
					NicKey:          strconv.Itoa(int(nic.Key)),
					Type:            string(types.DistributedVirtualSwitchPortConnecteeConnecteeTypeVmVnic),
				},
				Config: types.DVPortConfigInfo{
					Setting: pg.Config.DefaultPortConfig,
				},
			}
		}
	}

	for i := len(pg.PortKeys); i < nextKey; i++ {
		pg.PortKeys = append(pg.PortKeys, strconv.Itoa(i))
	}
	if int32(nextKey) > pg.Config.NumPorts {
		pg.Config.NumPorts = int32(nextKey)
	}

	ports := make([]types.DistributedVirtualPort, 0, len(pg.PortKeys))
	for _, key := range pg.PortKeys {
		if p, ok := claimed[key]; ok {
			ports = append(ports, p)
			continue
		}

		ports = append(ports, types.DistributedVirtualPort{
			DvsUuid:      s.Uuid,
			Key:          key,
			PortgroupKey: pg.Key,
			Config: types.DVPortConfigInfo{
				Setting: pg.Config.DefaultPortConfig,
			},
		})
	}

	return ports
}

func (s *VmwareDistributedVirtualSwitch) filterDVPorts(
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

func (s *VmwareDistributedVirtualSwitch) filterDVPortsByPortgroupKey(
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
			if slices.Contains(criteria.PortgroupKey, p.PortgroupKey) {
				filtered = append(filtered, p)
			}
		}
		return filtered
	}

	// outside portgroup keys
	filtered := []types.DistributedVirtualPort{}

	for _, p := range ports {
		found := slices.Contains(criteria.PortgroupKey, p.PortgroupKey)

		if !found {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (s *VmwareDistributedVirtualSwitch) filterDVPortsByPortKey(
	ports []types.DistributedVirtualPort,
	criteria *types.DistributedVirtualSwitchPortCriteria,
) []types.DistributedVirtualPort {
	if len(criteria.PortKey) == 0 {
		return ports
	}

	filtered := []types.DistributedVirtualPort{}

	for _, p := range ports {
		if slices.Contains(criteria.PortKey, p.Key) {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func (s *VmwareDistributedVirtualSwitch) filterDVPortsByConnected(
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
