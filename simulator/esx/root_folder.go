// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// RootFolder is the default template for the ServiceContent rootFolder property.
// Capture method:
//
//	govc folder.info -dump /
var RootFolder = mo.Folder{
	ManagedEntity: mo.ManagedEntity{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-root"},
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
				Entity:      &types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-root"},
				Principal:   "vpxuser",
				Group:       false,
				RoleId:      -1,
				Propagate:   true,
			},
			{
				DynamicData: types.DynamicData{},
				Entity:      &types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-root"},
				Principal:   "dcui",
				Group:       false,
				RoleId:      -1,
				Propagate:   true,
			},
			{
				DynamicData: types.DynamicData{},
				Entity:      &types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-root"},
				Principal:   "root",
				Group:       false,
				RoleId:      -1,
				Propagate:   true,
			},
		},
		Name:                "ha-folder-root",
		DisabledMethod:      nil,
		RecentTask:          nil,
		DeclaredAlarmState:  nil,
		TriggeredAlarmState: nil,
		AlarmActionsEnabled: (*bool)(nil),
		Tag:                 nil,
	},
	ChildType:   []string{"Datacenter"},
	ChildEntity: nil,
}
