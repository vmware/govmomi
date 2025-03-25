// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type HostVirtualNicManager struct {
	mo.HostVirtualNicManager

	Host *mo.HostSystem
}

func NewHostVirtualNicManager(host *mo.HostSystem) *HostVirtualNicManager {
	return &HostVirtualNicManager{
		Host: host,
		HostVirtualNicManager: mo.HostVirtualNicManager{
			Info: types.HostVirtualNicManagerInfo{
				NetConfig: esx.VirtualNicManagerNetConfig,
			},
		},
	}
}

func (m *HostVirtualNicManager) QueryNetConfig(req *types.QueryNetConfig) soap.HasFault {
	body := new(methods.QueryNetConfigBody)

	for _, c := range m.Info.NetConfig {
		if c.NicType == req.NicType {
			body.Res = &types.QueryNetConfigResponse{
				Returnval: &c,
			}
			return body
		}
	}

	body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: req.NicType})

	return body
}
