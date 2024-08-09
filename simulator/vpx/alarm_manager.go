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
