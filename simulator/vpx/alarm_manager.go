// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vpx

import (
	"time"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// Alarm captured using:
// govc alarm.info -dump -n alarm.VmErrorAlarm -n alarm.HostErrorAlarm
var Alarm = []mo.Alarm{
	{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Alarm", Value: "alarm-384", ServerGUID: ""},
			Value:          nil,
			AvailableField: nil,
		},
		Info: types.AlarmInfo{
			AlarmSpec: types.AlarmSpec{
				Name:        "vcsim VM Alarm",
				SystemName:  "",
				Description: "vcsim alarm for Virtual Machines",
				Enabled:     true,
				Expression: &types.OrAlarmExpression{
					AlarmExpression: types.AlarmExpression{},
					Expression: []types.BaseAlarmExpression{
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "EventEx",
							EventTypeId:     "vcsim.vm.success",
							ObjectType:      "VirtualMachine",
							Status:          "green",
						},
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "EventEx",
							EventTypeId:     "vcsim.vm.failure",
							ObjectType:      "VirtualMachine",
							Status:          "yellow",
						},
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "EventEx",
							EventTypeId:     "vcsim.vm.fatal",
							ObjectType:      "VirtualMachine",
							Status:          "red",
						},
					},
				},
				Action:          nil,
				ActionFrequency: 0,
				Setting: &types.AlarmSetting{
					ToleranceRange:     0,
					ReportingFrequency: 300,
				},
			},
			Key:              "",
			Alarm:            types.ManagedObjectReference{Type: "Alarm", Value: "alarm-384", ServerGUID: ""},
			Entity:           types.ManagedObjectReference{Type: "Folder", Value: "group-d1", ServerGUID: ""},
			LastModifiedTime: time.Now(),
			LastModifiedUser: "VSPHERE.LOCAL\\Administrator",
			CreationEventId:  0,
		},
	},
	{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Alarm", Value: "alarm-11", ServerGUID: ""},
			Value:          nil,
			AvailableField: nil,
		},
		Info: types.AlarmInfo{
			AlarmSpec: types.AlarmSpec{
				Name:        "Host error",
				SystemName:  "alarm.HostErrorAlarm",
				Description: "Default alarm to monitor host error and warning events",
				Enabled:     true,
				Expression: &types.OrAlarmExpression{
					AlarmExpression: types.AlarmExpression{},
					Expression: []types.BaseAlarmExpression{
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "GeneralHostErrorEvent",
							EventTypeId:     "",
							ObjectType:      "HostSystem",
							Status:          "",
						},
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "GeneralHostWarningEvent",
							EventTypeId:     "",
							ObjectType:      "HostSystem",
							Status:          "",
						},
					},
				},
				Action: &types.GroupAlarmAction{
					AlarmAction: types.AlarmAction{},
					Action: []types.BaseAlarmAction{
						&types.AlarmTriggeringAction{
							AlarmAction: types.AlarmAction{},
							Action:      &types.SendSNMPAction{},
							TransitionSpecs: []types.AlarmTriggeringActionTransitionSpec{
								{
									StartState: "yellow",
									FinalState: "red",
									Repeats:    true,
								},
							},
							Green2yellow: false,
							Yellow2red:   false,
							Red2yellow:   false,
							Yellow2green: false,
						},
					},
				},
				ActionFrequency: 0,
				Setting: &types.AlarmSetting{
					ToleranceRange:     0,
					ReportingFrequency: 300,
				},
			},
			Key:              "",
			Alarm:            types.ManagedObjectReference{Type: "Alarm", Value: "alarm-11", ServerGUID: ""},
			Entity:           types.ManagedObjectReference{Type: "Folder", Value: "group-d1", ServerGUID: ""},
			LastModifiedTime: time.Now(),
			LastModifiedUser: "",
			CreationEventId:  0,
		},
	},
	{
		ExtensibleManagedObject: mo.ExtensibleManagedObject{
			Self:           types.ManagedObjectReference{Type: "Alarm", Value: "alarm-12", ServerGUID: ""},
			Value:          nil,
			AvailableField: nil,
		},
		Info: types.AlarmInfo{
			AlarmSpec: types.AlarmSpec{
				Name:        "Virtual machine error",
				SystemName:  "alarm.VmErrorAlarm",
				Description: "Default alarm to monitor virtual machine error and warning events",
				Enabled:     true,
				Expression: &types.OrAlarmExpression{
					AlarmExpression: types.AlarmExpression{},
					Expression: []types.BaseAlarmExpression{
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "GeneralVmErrorEvent",
							EventTypeId:     "",
							ObjectType:      "VirtualMachine",
							Status:          "",
						},
						&types.EventAlarmExpression{
							AlarmExpression: types.AlarmExpression{},
							Comparisons:     nil,
							EventType:       "GeneralVmWarningEvent",
							EventTypeId:     "",
							ObjectType:      "VirtualMachine",
							Status:          "",
						},
					},
				},
				Action: &types.GroupAlarmAction{
					AlarmAction: types.AlarmAction{},
					Action: []types.BaseAlarmAction{
						&types.AlarmTriggeringAction{
							AlarmAction: types.AlarmAction{},
							Action:      &types.SendSNMPAction{},
							TransitionSpecs: []types.AlarmTriggeringActionTransitionSpec{
								{
									StartState: "yellow",
									FinalState: "red",
									Repeats:    true,
								},
							},
							Green2yellow: false,
							Yellow2red:   false,
							Red2yellow:   false,
							Yellow2green: false,
						},
					},
				},
				ActionFrequency: 0,
				Setting: &types.AlarmSetting{
					ToleranceRange:     0,
					ReportingFrequency: 300,
				},
			},
			Key:              "",
			Alarm:            types.ManagedObjectReference{Type: "Alarm", Value: "alarm-12", ServerGUID: ""},
			Entity:           types.ManagedObjectReference{Type: "Folder", Value: "group-d1", ServerGUID: ""},
			LastModifiedTime: time.Now(),
			LastModifiedUser: "",
			CreationEventId:  0,
		},
	},
}
