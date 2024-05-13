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
