/*
Copyright (c) 2021-2023 VMware, Inc. All Rights Reserved.

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

type HooksHookType string

const (
	HooksHookTypePOST_PROVISIONING = HooksHookType("POST_PROVISIONING")
	HooksHookTypePOST_POWER_ON     = HooksHookType("POST_POWER_ON")
)

func init() {
	types.Add("eam:HooksHookType", reflect.TypeOf((*HooksHookType)(nil)).Elem())
}

type SolutionsInvalidReason string

const (
	SolutionsInvalidReasonINVALID_OVF_DESCRIPTOR = SolutionsInvalidReason("INVALID_OVF_DESCRIPTOR")
	SolutionsInvalidReasonINACCESSBLE_VM_SOURCE  = SolutionsInvalidReason("INACCESSBLE_VM_SOURCE")
	SolutionsInvalidReasonINVALID_NETWORKS       = SolutionsInvalidReason("INVALID_NETWORKS")
	SolutionsInvalidReasonINVALID_DATASTORES     = SolutionsInvalidReason("INVALID_DATASTORES")
	SolutionsInvalidReasonINVALID_RESOURCE_POOL  = SolutionsInvalidReason("INVALID_RESOURCE_POOL")
	SolutionsInvalidReasonINVALID_FOLDER         = SolutionsInvalidReason("INVALID_FOLDER")
	SolutionsInvalidReasonINVALID_PROPERTIES     = SolutionsInvalidReason("INVALID_PROPERTIES")
	SolutionsInvalidReasonINVALID_TRANSITION     = SolutionsInvalidReason("INVALID_TRANSITION")
)

func init() {
	types.Add("eam:SolutionsInvalidReason", reflect.TypeOf((*SolutionsInvalidReason)(nil)).Elem())
}

type SolutionsNonComplianceReason string

const (
	SolutionsNonComplianceReasonWORKING       = SolutionsNonComplianceReason("WORKING")
	SolutionsNonComplianceReasonISSUE         = SolutionsNonComplianceReason("ISSUE")
	SolutionsNonComplianceReasonIN_HOOK       = SolutionsNonComplianceReason("IN_HOOK")
	SolutionsNonComplianceReasonOBSOLETE_SPEC = SolutionsNonComplianceReason("OBSOLETE_SPEC")
	SolutionsNonComplianceReasonNO_SPEC       = SolutionsNonComplianceReason("NO_SPEC")
)

func init() {
	types.Add("eam:SolutionsNonComplianceReason", reflect.TypeOf((*SolutionsNonComplianceReason)(nil)).Elem())
}

type SolutionsVMDeploymentOptimization string

const (
	SolutionsVMDeploymentOptimizationALL_CLONES       = SolutionsVMDeploymentOptimization("ALL_CLONES")
	SolutionsVMDeploymentOptimizationFULL_CLONES_ONLY = SolutionsVMDeploymentOptimization("FULL_CLONES_ONLY")
	SolutionsVMDeploymentOptimizationNO_CLONES        = SolutionsVMDeploymentOptimization("NO_CLONES")
)

func init() {
	types.Add("eam:SolutionsVMDeploymentOptimization", reflect.TypeOf((*SolutionsVMDeploymentOptimization)(nil)).Elem())
}

type SolutionsVMDiskProvisioning string

const (
	SolutionsVMDiskProvisioningTHIN  = SolutionsVMDiskProvisioning("THIN")
	SolutionsVMDiskProvisioningTHICK = SolutionsVMDiskProvisioning("THICK")
)

func init() {
	types.Add("eam:SolutionsVMDiskProvisioning", reflect.TypeOf((*SolutionsVMDiskProvisioning)(nil)).Elem())
}
