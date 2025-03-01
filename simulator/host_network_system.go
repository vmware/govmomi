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

type HostNetworkSystem struct {
	mo.HostNetworkSystem

	Host *mo.HostSystem

	types.QueryNetworkHintResponse
}

func NewHostNetworkSystem(host *mo.HostSystem) *HostNetworkSystem {
	return &HostNetworkSystem{
		Host: host,
		HostNetworkSystem: mo.HostNetworkSystem{
			NetworkInfo: &types.HostNetworkInfo{
				Vswitch: []types.HostVirtualSwitch{
					{
						Name:      "vSwitch0",
						Portgroup: []string{"VM Network"},
					},
				},
				Portgroup: host.Config.Network.Portgroup,
			},
		},
	}
}

func (s *HostNetworkSystem) init(r *Registry) {
	for _, obj := range r.objects {
		if h, ok := obj.(*HostSystem); ok {
			if h.ConfigManager.NetworkSystem.Value == s.Self.Value {
				s.Host = &h.HostSystem
			}
		}
	}
}

func (s *HostNetworkSystem) folder(ctx *Context) *Folder {
	f := ctx.Map.getEntityDatacenter(s.Host).NetworkFolder
	return ctx.Map.Get(f).(*Folder)
}

func (s *HostNetworkSystem) AddVirtualSwitch(c *types.AddVirtualSwitch) soap.HasFault {
	r := &methods.AddVirtualSwitchBody{}

	for _, vswitch := range s.NetworkInfo.Vswitch {
		if vswitch.Name == c.VswitchName {
			r.Fault_ = Fault("", &types.AlreadyExists{Name: c.VswitchName})
			return r
		}
	}

	s.NetworkInfo.Vswitch = append(s.NetworkInfo.Vswitch, types.HostVirtualSwitch{
		Name: c.VswitchName,
	})

	r.Res = &types.AddVirtualSwitchResponse{}

	return r
}

func (s *HostNetworkSystem) RemoveVirtualSwitch(c *types.RemoveVirtualSwitch) soap.HasFault {
	r := &methods.RemoveVirtualSwitchBody{}

	vs := s.NetworkInfo.Vswitch

	for i, v := range vs {
		if v.Name == c.VswitchName {
			s.NetworkInfo.Vswitch = append(vs[:i], vs[i+1:]...)
			r.Res = &types.RemoveVirtualSwitchResponse{}
			return r
		}
	}

	r.Fault_ = Fault("", &types.NotFound{})

	return r
}

func (s *HostNetworkSystem) AddPortGroup(ctx *Context, c *types.AddPortGroup) soap.HasFault {
	var vswitch *types.HostVirtualSwitch

	r := &methods.AddPortGroupBody{}

	if c.Portgrp.Name == "" {
		r.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "name"})
		return r
	}

	for i := range s.NetworkInfo.Vswitch {
		if s.NetworkInfo.Vswitch[i].Name == c.Portgrp.VswitchName {
			vswitch = &s.NetworkInfo.Vswitch[i]
			break
		}
	}

	if vswitch == nil {
		r.Fault_ = Fault("", &types.NotFound{})
		return r
	}

	network := &mo.Network{}
	network.Name = c.Portgrp.Name
	network.Entity().Name = network.Name

	folder := s.folder(ctx)

	if obj := ctx.Map.FindByName(c.Portgrp.Name, folder.ChildEntity); obj != nil {
		r.Fault_ = Fault("", &types.DuplicateName{
			Name:   c.Portgrp.Name,
			Object: obj.Reference(),
		})

		return r
	}

	folderPutChild(ctx, &folder.Folder, network)

	vswitch.Portgroup = append(vswitch.Portgroup, c.Portgrp.Name)

	s.NetworkInfo.Portgroup = append(s.NetworkInfo.Portgroup, types.HostPortGroup{
		Key:  "key-vim.host.PortGroup-" + c.Portgrp.Name,
		Port: nil,
		Spec: c.Portgrp,
	})

	r.Res = &types.AddPortGroupResponse{}

	return r
}

func (s *HostNetworkSystem) RemovePortGroup(ctx *Context, c *types.RemovePortGroup) soap.HasFault {
	var vswitch *types.HostVirtualSwitch

	r := &methods.RemovePortGroupBody{}

	for i, v := range s.NetworkInfo.Vswitch {
		for j, pg := range v.Portgroup {
			if pg == c.PgName {
				vswitch = &s.NetworkInfo.Vswitch[i]
				vswitch.Portgroup = append(vswitch.Portgroup[:j], vswitch.Portgroup[j+1:]...)
			}
		}
	}

	if vswitch == nil {
		r.Fault_ = Fault("", &types.NotFound{})
		return r
	}

	folder := s.folder(ctx)
	e := ctx.Map.FindByName(c.PgName, folder.ChildEntity)
	folderRemoveChild(ctx, &folder.Folder, e.Reference())

	for i, pg := range s.NetworkInfo.Portgroup {
		if pg.Spec.Name == c.PgName {
			var portgroup = s.NetworkInfo.Portgroup
			s.NetworkInfo.Portgroup = append(portgroup[:i], portgroup[i+1:]...)
		}
	}

	r.Res = &types.RemovePortGroupResponse{}

	return r
}

func (s *HostNetworkSystem) UpdateNetworkConfig(req *types.UpdateNetworkConfig) soap.HasFault {
	s.NetworkConfig = &req.Config

	return &methods.UpdateNetworkConfigBody{
		Res: &types.UpdateNetworkConfigResponse{
			Returnval: types.HostNetworkConfigResult{},
		},
	}
}

func (s *HostNetworkSystem) QueryNetworkHint(req *types.QueryNetworkHint) soap.HasFault {
	if s.Host.Runtime.ConnectionState != types.HostSystemConnectionStateConnected {
		return &methods.QueryNetworkHintBody{
			Fault_: Fault("", &types.HostNotConnected{}),
		}
	}

	return &methods.QueryNetworkHintBody{
		Res: &types.QueryNetworkHintResponse{
			Returnval: s.QueryNetworkHintResponse.Returnval,
		},
	}
}
