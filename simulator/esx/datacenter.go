// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// Datacenter is the default template for Datacenter properties.
// Capture method:
// govc datacenter.info -dump
var Datacenter = mo.Datacenter{
	ManagedEntity: mo.ManagedEntity{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Datacenter", Value: "ha-datacenter"},
			Value:          nil,
			AvailableField: nil,
		},
		Parent:              (*types.ManagedObjectReference)(nil),
		CustomValue:         nil,
		OverallStatus:       "",
		ConfigStatus:        "",
		ConfigIssue:         nil,
		EffectiveRole:       nil,
		Permission:          nil,
		Name:                "ha-datacenter",
		DisabledMethod:      nil,
		RecentTask:          nil,
		DeclaredAlarmState:  nil,
		TriggeredAlarmState: nil,
		AlarmActionsEnabled: (*bool)(nil),
		Tag:                 nil,
	},
	VmFolder:        types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-vm"},
	HostFolder:      types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-host"},
	DatastoreFolder: types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-datastore"},
	NetworkFolder:   types.ManagedObjectReference{Type: "Folder", Value: "ha-folder-network"},
	Datastore:       nil,
	Network: []types.ManagedObjectReference{
		{Type: "Network", Value: "HaNetwork-VM Network"},
	},
	Configuration: types.DatacenterConfigInfo{},
}
