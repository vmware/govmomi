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

package esx

import "github.com/vmware/govmomi/vim25/types"

var VirtualNicManagerNetConfig = []types.VirtualNicManagerNetConfig{
	{
		NicType:            "management",
		MultiSelectAllowed: true,
		CandidateVnic: []types.HostVirtualNic{
			{
				Device:    "vmk0",
				Key:       "management.key-vim.host.VirtualNic-vmk0",
				Portgroup: "Management Network",
				Spec: types.HostVirtualNicSpec{
					Ip: &types.HostIpConfig{
						Dhcp:       true,
						IpAddress:  "127.0.0.1",
						SubnetMask: "255.0.0.0",
					},
					Mac:                 "00:0c:29:81:d8:a0",
					Portgroup:           "Management Network",
					Mtu:                 1500,
					TsoEnabled:          types.NewBool(true),
					NetStackInstanceKey: "defaultTcpipStack",
					SystemOwned:         types.NewBool(false),
				},
			},
		},
		SelectedVnic: []string{"management.key-vim.host.VirtualNic-vmk0"},
	},
}
