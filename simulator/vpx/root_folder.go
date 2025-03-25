// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vpx

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var RootFolder = mo.Folder{
	ManagedEntity: mo.ManagedEntity{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Folder", Value: "group-d1"},
			Value:          nil,
			AvailableField: nil,
		},
		Parent:        (*types.ManagedObjectReference)(nil),
		CustomValue:   nil,
		OverallStatus: "green",
		ConfigStatus:  "green",
		ConfigIssue:   nil,
		EffectiveRole: []int32{-1},
		Permission: []types.Permission{
			{
				DynamicData: types.DynamicData{},
				Entity:      &types.ManagedObjectReference{Type: "Folder", Value: "group-d1"},
				Principal:   "VSPHERE.LOCAL\\Administrator",
				Group:       false,
				RoleId:      -1,
				Propagate:   true,
			},
			{
				DynamicData: types.DynamicData{},
				Entity:      &types.ManagedObjectReference{Type: "Folder", Value: "group-d1"},
				Principal:   "VSPHERE.LOCAL\\Administrators",
				Group:       true,
				RoleId:      -1,
				Propagate:   true,
			},
		},
		Name:                "Datacenters",
		DisabledMethod:      nil,
		RecentTask:          nil,
		DeclaredAlarmState:  nil,
		AlarmActionsEnabled: (*bool)(nil),
		Tag:                 nil,
	},
	ChildType:   []string{"Folder", "Datacenter"},
	ChildEntity: nil,
}
