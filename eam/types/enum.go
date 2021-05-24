/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type AgencyVMPlacementPolicyVMAntiAffinity string

const (
	AgencyVMPlacementPolicyVMAntiAffinityNone = AgencyVMPlacementPolicyVMAntiAffinity("none")
	AgencyVMPlacementPolicyVMAntiAffinitySoft = AgencyVMPlacementPolicyVMAntiAffinity("soft")
)

func init() {
	types.Add("eam:AgencyVMPlacementPolicyVMAntiAffinity", reflect.TypeOf((*AgencyVMPlacementPolicyVMAntiAffinity)(nil)).Elem())
}

type AgencyVMPlacementPolicyVMDataAffinity string

const (
	AgencyVMPlacementPolicyVMDataAffinityNone = AgencyVMPlacementPolicyVMDataAffinity("none")
	AgencyVMPlacementPolicyVMDataAffinitySoft = AgencyVMPlacementPolicyVMDataAffinity("soft")
)

func init() {
	types.Add("eam:AgencyVMPlacementPolicyVMDataAffinity", reflect.TypeOf((*AgencyVMPlacementPolicyVMDataAffinity)(nil)).Elem())
}

type AgentConfigInfoOvfDiskProvisioning string

const (
	AgentConfigInfoOvfDiskProvisioningNone  = AgentConfigInfoOvfDiskProvisioning("none")
	AgentConfigInfoOvfDiskProvisioningThin  = AgentConfigInfoOvfDiskProvisioning("thin")
	AgentConfigInfoOvfDiskProvisioningThick = AgentConfigInfoOvfDiskProvisioning("thick")
)

func init() {
	types.Add("eam:AgentConfigInfoOvfDiskProvisioning", reflect.TypeOf((*AgentConfigInfoOvfDiskProvisioning)(nil)).Elem())
}

type AgentVmHookVmState string

const (
	AgentVmHookVmStateProvisioned = AgentVmHookVmState("provisioned")
	AgentVmHookVmStatePoweredOn   = AgentVmHookVmState("poweredOn")
	AgentVmHookVmStatePrePowerOn  = AgentVmHookVmState("prePowerOn")
)

func init() {
	types.Add("eam:AgentVmHookVmState", reflect.TypeOf((*AgentVmHookVmState)(nil)).Elem())
}

type EamObjectRuntimeInfoGoalState string

const (
	EamObjectRuntimeInfoGoalStateEnabled     = EamObjectRuntimeInfoGoalState("enabled")
	EamObjectRuntimeInfoGoalStateDisabled    = EamObjectRuntimeInfoGoalState("disabled")
	EamObjectRuntimeInfoGoalStateUninstalled = EamObjectRuntimeInfoGoalState("uninstalled")
)

func init() {
	types.Add("eam:EamObjectRuntimeInfoGoalState", reflect.TypeOf((*EamObjectRuntimeInfoGoalState)(nil)).Elem())
}

type EamObjectRuntimeInfoStatus string

const (
	EamObjectRuntimeInfoStatusGreen  = EamObjectRuntimeInfoStatus("green")
	EamObjectRuntimeInfoStatusYellow = EamObjectRuntimeInfoStatus("yellow")
	EamObjectRuntimeInfoStatusRed    = EamObjectRuntimeInfoStatus("red")
)

func init() {
	types.Add("eam:EamObjectRuntimeInfoStatus", reflect.TypeOf((*EamObjectRuntimeInfoStatus)(nil)).Elem())
}

type EsxAgentManagerMaintenanceModePolicy string

const (
	EsxAgentManagerMaintenanceModePolicySingleHost    = EsxAgentManagerMaintenanceModePolicy("singleHost")
	EsxAgentManagerMaintenanceModePolicyMultipleHosts = EsxAgentManagerMaintenanceModePolicy("multipleHosts")
)

func init() {
	types.Add("eam:EsxAgentManagerMaintenanceModePolicy", reflect.TypeOf((*EsxAgentManagerMaintenanceModePolicy)(nil)).Elem())
}
