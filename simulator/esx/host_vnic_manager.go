// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
