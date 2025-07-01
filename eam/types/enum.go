// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

// Deprecated as of vSphere 9.0. Please refer to vLCM System VMs APIs.
//
// Defines if the deployed VMs needs to run on different hosts.
type AgencyVMPlacementPolicyVMAntiAffinity string

const (
	// Denotes no specific VM anti-affinity policy.
	AgencyVMPlacementPolicyVMAntiAffinityNone = AgencyVMPlacementPolicyVMAntiAffinity("none")
	// Best effort is made the VMs to run on different hosts as long as
	// this does not impact the ability of the host to satisfy current CPU
	// or memory requirements for virtual machines on the system.
	//
	// NOTE: Currently not supported - i.e. the agency configuration is
	// considered as invalid.
	AgencyVMPlacementPolicyVMAntiAffinitySoft = AgencyVMPlacementPolicyVMAntiAffinity("soft")
)

func (e AgencyVMPlacementPolicyVMAntiAffinity) Values() []AgencyVMPlacementPolicyVMAntiAffinity {
	return []AgencyVMPlacementPolicyVMAntiAffinity{
		AgencyVMPlacementPolicyVMAntiAffinityNone,
		AgencyVMPlacementPolicyVMAntiAffinitySoft,
	}
}

func (e AgencyVMPlacementPolicyVMAntiAffinity) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:AgencyVMPlacementPolicyVMAntiAffinity", reflect.TypeOf((*AgencyVMPlacementPolicyVMAntiAffinity)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM System VMs APIs.
//
// Defines if the deployed VM is affinied to run on the same host it is
// deployed on.
type AgencyVMPlacementPolicyVMDataAffinity string

const (
	// Denotes no specific VM data affinity policy.
	AgencyVMPlacementPolicyVMDataAffinityNone = AgencyVMPlacementPolicyVMDataAffinity("none")
	// Best effort is made the VM to run on the same host it is deployed on
	// as long as this does not impact the ability of the host to satisfy
	// current CPU or memory requirements for virtual machines on the
	// system.
	//
	// NOTE: Currently not supported - i.e. the agency configuration is
	// considered as invalid.
	AgencyVMPlacementPolicyVMDataAffinitySoft = AgencyVMPlacementPolicyVMDataAffinity("soft")
)

func (e AgencyVMPlacementPolicyVMDataAffinity) Values() []AgencyVMPlacementPolicyVMDataAffinity {
	return []AgencyVMPlacementPolicyVMDataAffinity{
		AgencyVMPlacementPolicyVMDataAffinityNone,
		AgencyVMPlacementPolicyVMDataAffinitySoft,
	}
}

func (e AgencyVMPlacementPolicyVMDataAffinity) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:AgencyVMPlacementPolicyVMDataAffinity", reflect.TypeOf((*AgencyVMPlacementPolicyVMDataAffinity)(nil)).Elem())
}

type AgentConfigInfoAuthenticationScheme string

const (
	AgentConfigInfoAuthenticationSchemeNONE              = AgentConfigInfoAuthenticationScheme("NONE")
	AgentConfigInfoAuthenticationSchemeVMWARE_SESSION_ID = AgentConfigInfoAuthenticationScheme("VMWARE_SESSION_ID")
)

func (e AgentConfigInfoAuthenticationScheme) Values() []AgentConfigInfoAuthenticationScheme {
	return []AgentConfigInfoAuthenticationScheme{
		AgentConfigInfoAuthenticationSchemeNONE,
		AgentConfigInfoAuthenticationSchemeVMWARE_SESSION_ID,
	}
}

func (e AgentConfigInfoAuthenticationScheme) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:AgentConfigInfoAuthenticationScheme", reflect.TypeOf((*AgentConfigInfoAuthenticationScheme)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// Defines the type of disk provisioning for the target Agent VMs.
type AgentConfigInfoOvfDiskProvisioning string

const (
	// Denotes no specific type for disk provisioning.
	//
	// Disks will be
	// provisioned as defaulted by vSphere.
	AgentConfigInfoOvfDiskProvisioningNone = AgentConfigInfoOvfDiskProvisioning("none")
	// Disks will be provisioned with only used space allocated.
	AgentConfigInfoOvfDiskProvisioningThin = AgentConfigInfoOvfDiskProvisioning("thin")
	// Disks will be provisioned with full size allocated.
	AgentConfigInfoOvfDiskProvisioningThick = AgentConfigInfoOvfDiskProvisioning("thick")
)

func (e AgentConfigInfoOvfDiskProvisioning) Values() []AgentConfigInfoOvfDiskProvisioning {
	return []AgentConfigInfoOvfDiskProvisioning{
		AgentConfigInfoOvfDiskProvisioningNone,
		AgentConfigInfoOvfDiskProvisioningThin,
		AgentConfigInfoOvfDiskProvisioningThick,
	}
}

func (e AgentConfigInfoOvfDiskProvisioning) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:AgentConfigInfoOvfDiskProvisioning", reflect.TypeOf((*AgentConfigInfoOvfDiskProvisioning)(nil)).Elem())
}

// Represents the state of the VM lifecycle.
type AgentVmHookVmState string

const (
	// The VM is provisioned and not powered-on.
	AgentVmHookVmStateProvisioned = AgentVmHookVmState("provisioned")
	// The VM is powered on.
	AgentVmHookVmStatePoweredOn = AgentVmHookVmState("poweredOn")
	// The VM is about to be powered on as part of a VM upgrade workflow.
	AgentVmHookVmStatePrePowerOn = AgentVmHookVmState("prePowerOn")
)

func (e AgentVmHookVmState) Values() []AgentVmHookVmState {
	return []AgentVmHookVmState{
		AgentVmHookVmStateProvisioned,
		AgentVmHookVmStatePoweredOn,
		AgentVmHookVmStatePrePowerOn,
	}
}

func (e AgentVmHookVmState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:AgentVmHookVmState", reflect.TypeOf((*AgentVmHookVmState)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// The <code>GoalState</code> enumeration defines the goal of the entity.
type EamObjectRuntimeInfoGoalState string

const (
	// The entity should be fully deployed and active.
	//
	// If the entity is an
	// `Agency`, it should install VIBs and deploy and power on all agent
	// virtual machines. If the entity is an `Agent`, its VIB should be installed and its
	// agent virtual machine should be deployed and powered on.
	EamObjectRuntimeInfoGoalStateEnabled = EamObjectRuntimeInfoGoalState("enabled")
	// The entity should be fully deployed but inactive.
	//
	// f the entity is an
	// `Agency`, the behavior is similar to the <code>enabled</code> goal state, but
	// agents are not powered on (if they have been powered on they are powered
	// off).
	EamObjectRuntimeInfoGoalStateDisabled = EamObjectRuntimeInfoGoalState("disabled")
	// The entity should be completely removed from the vCenter Server.
	//
	// If the entity is an
	// `Agency`, no more VIBs or agent virtual machines are deployed. All installed VIBs
	// installed by the `Agency` are uninstalled and any deployed agent virtual machines
	// are powered off (if they have been powered on) and deleted.
	// If the entity is an `Agent`, its VIB is uninstalled and the virtual machine is
	// powered off and deleted.
	EamObjectRuntimeInfoGoalStateUninstalled = EamObjectRuntimeInfoGoalState("uninstalled")
)

func (e EamObjectRuntimeInfoGoalState) Values() []EamObjectRuntimeInfoGoalState {
	return []EamObjectRuntimeInfoGoalState{
		EamObjectRuntimeInfoGoalStateEnabled,
		EamObjectRuntimeInfoGoalStateDisabled,
		EamObjectRuntimeInfoGoalStateUninstalled,
	}
}

func (e EamObjectRuntimeInfoGoalState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:EamObjectRuntimeInfoGoalState", reflect.TypeOf((*EamObjectRuntimeInfoGoalState)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM APIs.
//
// <code>Status</code> defines a health value that denotes how well the entity
// conforms to the goal state.
type EamObjectRuntimeInfoStatus string

const (
	// The entity is in perfect compliance with the goal state.
	EamObjectRuntimeInfoStatusGreen = EamObjectRuntimeInfoStatus("green")
	// The entity is actively working to reach the desired goal state.
	EamObjectRuntimeInfoStatusYellow = EamObjectRuntimeInfoStatus("yellow")
	// The entity has reached an issue which prevents it from reaching the desired goal
	// state.
	//
	// To remediate any offending issues, look at `EamObjectRuntimeInfo.issue`
	// and use either `EamObject.Resolve` or
	// `EamObject.ResolveAll`.
	EamObjectRuntimeInfoStatusRed = EamObjectRuntimeInfoStatus("red")
)

func (e EamObjectRuntimeInfoStatus) Values() []EamObjectRuntimeInfoStatus {
	return []EamObjectRuntimeInfoStatus{
		EamObjectRuntimeInfoStatusGreen,
		EamObjectRuntimeInfoStatusYellow,
		EamObjectRuntimeInfoStatusRed,
	}
}

func (e EamObjectRuntimeInfoStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:EamObjectRuntimeInfoStatus", reflect.TypeOf((*EamObjectRuntimeInfoStatus)(nil)).Elem())
}

// Deprecated as of vSphere 9.0. Please refer to vLCM Image APIs.
//
// <code>MaintenanceModePolicy</code> defines how ESX Agent Manager is going
// to put into maintenance mode hosts which are part of a cluster not managed
type EsxAgentManagerMaintenanceModePolicy string

const (
	// Only a single host at a time will be put into maintenance mode.
	EsxAgentManagerMaintenanceModePolicySingleHost = EsxAgentManagerMaintenanceModePolicy("singleHost")
	// Hosts will be put into maintenance mode simultaneously.
	//
	// If vSphere DRS
	// is enabled, its recommendations will be used. Otherwise, it will be
	// attempted to put in maintenance mode simultaneously as many host as
	// possible.
	EsxAgentManagerMaintenanceModePolicyMultipleHosts = EsxAgentManagerMaintenanceModePolicy("multipleHosts")
)

func (e EsxAgentManagerMaintenanceModePolicy) Values() []EsxAgentManagerMaintenanceModePolicy {
	return []EsxAgentManagerMaintenanceModePolicy{
		EsxAgentManagerMaintenanceModePolicySingleHost,
		EsxAgentManagerMaintenanceModePolicyMultipleHosts,
	}
}

func (e EsxAgentManagerMaintenanceModePolicy) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:EsxAgentManagerMaintenanceModePolicy", reflect.TypeOf((*EsxAgentManagerMaintenanceModePolicy)(nil)).Elem())
	types.AddMinAPIVersionForType("eam:EsxAgentManagerMaintenanceModePolicy", "7.4")
}

// Supported types of hooks for agents.
type HooksHookType string

const (
	// Hook raised for an agent immediately after a Virtual Machine was
	// created.
	HooksHookTypePOST_PROVISIONING = HooksHookType("POST_PROVISIONING")
	// Hook raised for an agent immediately after a Virtual Machine was
	// powered on.
	HooksHookTypePOST_POWER_ON = HooksHookType("POST_POWER_ON")
)

func (e HooksHookType) Values() []HooksHookType {
	return []HooksHookType{
		HooksHookTypePOST_PROVISIONING,
		HooksHookTypePOST_POWER_ON,
	}
}

func (e HooksHookType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:HooksHookType", reflect.TypeOf((*HooksHookType)(nil)).Elem())
}

type ManagedObjectTypes string

const (
	ManagedObjectTypesAgency          = ManagedObjectTypes("Agency")
	ManagedObjectTypesAgent           = ManagedObjectTypes("Agent")
	ManagedObjectTypesEamObject       = ManagedObjectTypes("EamObject")
	ManagedObjectTypesEsxAgentManager = ManagedObjectTypes("EsxAgentManager")
	ManagedObjectTypesEamTask         = ManagedObjectTypes("EamTask")
)

func (e ManagedObjectTypes) Values() []ManagedObjectTypes {
	return []ManagedObjectTypes{
		ManagedObjectTypesAgency,
		ManagedObjectTypesAgent,
		ManagedObjectTypesEamObject,
		ManagedObjectTypesEsxAgentManager,
		ManagedObjectTypesEamTask,
	}
}

func (e ManagedObjectTypes) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:ManagedObjectTypes", reflect.TypeOf((*ManagedObjectTypes)(nil)).Elem())
}

// Reasons solution is not valid for application.
type SolutionsInvalidReason string

const (
	// The OVF descriptor provided in the VM source is invalid.
	SolutionsInvalidReasonINVALID_OVF_DESCRIPTOR = SolutionsInvalidReason("INVALID_OVF_DESCRIPTOR")
	// The provided VM source is inaccessible from ESX Agent Manager.
	SolutionsInvalidReasonINACCESSBLE_VM_SOURCE = SolutionsInvalidReason("INACCESSBLE_VM_SOURCE")
	// The provided networks are not suitable for application purposes.
	SolutionsInvalidReasonINVALID_NETWORKS = SolutionsInvalidReason("INVALID_NETWORKS")
	// The provided datastores are not suitable for application purposes.
	SolutionsInvalidReasonINVALID_DATASTORES = SolutionsInvalidReason("INVALID_DATASTORES")
	// The provided resource pool is not accessible or part of the cluster.
	SolutionsInvalidReasonINVALID_RESOURCE_POOL = SolutionsInvalidReason("INVALID_RESOURCE_POOL")
	// The provided folder is inaccessible or not part of the same datacenter
	// with the cluster.
	SolutionsInvalidReasonINVALID_FOLDER = SolutionsInvalidReason("INVALID_FOLDER")
	// The provided OVF properties are insufficient to satisfy the required
	// user configurable properties in the VM described in the vmSource.
	SolutionsInvalidReasonINVALID_PROPERTIES = SolutionsInvalidReason("INVALID_PROPERTIES")
	// The legacy agency requested for transition is not valid/cannot be
	// mapped to systm Virtual Machines solution.
	SolutionsInvalidReasonINVALID_TRANSITION = SolutionsInvalidReason("INVALID_TRANSITION")
)

func (e SolutionsInvalidReason) Values() []SolutionsInvalidReason {
	return []SolutionsInvalidReason{
		SolutionsInvalidReasonINVALID_OVF_DESCRIPTOR,
		SolutionsInvalidReasonINACCESSBLE_VM_SOURCE,
		SolutionsInvalidReasonINVALID_NETWORKS,
		SolutionsInvalidReasonINVALID_DATASTORES,
		SolutionsInvalidReasonINVALID_RESOURCE_POOL,
		SolutionsInvalidReasonINVALID_FOLDER,
		SolutionsInvalidReasonINVALID_PROPERTIES,
		SolutionsInvalidReasonINVALID_TRANSITION,
	}
}

func (e SolutionsInvalidReason) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:SolutionsInvalidReason", reflect.TypeOf((*SolutionsInvalidReason)(nil)).Elem())
}

// Describes possible reasons a solution is non compliant.
type SolutionsNonComplianceReason string

const (
	// There is ongoing work to acheive the desired state.
	SolutionsNonComplianceReasonWORKING = SolutionsNonComplianceReason("WORKING")
	// ESX Agent Manager has ecnountered am issue attempting to acheive the
	// desired state.
	SolutionsNonComplianceReasonISSUE = SolutionsNonComplianceReason("ISSUE")
	// ESX Agent Manager is awaiting user input to continue attempting to
	// acheive the desired state.
	SolutionsNonComplianceReasonIN_HOOK = SolutionsNonComplianceReason("IN_HOOK")
	// An obsoleted spec is currently in application for this solution.
	//
	// This state should take precedence over:
	//   - `WORKING`
	//   - `ISSUE`
	//   - `IN_HOOK`
	SolutionsNonComplianceReasonOBSOLETE_SPEC = SolutionsNonComplianceReason("OBSOLETE_SPEC")
	// Application for this solutiona has never been requested with
	// `Solutions.Apply`.
	SolutionsNonComplianceReasonNO_SPEC = SolutionsNonComplianceReason("NO_SPEC")
)

func (e SolutionsNonComplianceReason) Values() []SolutionsNonComplianceReason {
	return []SolutionsNonComplianceReason{
		SolutionsNonComplianceReasonWORKING,
		SolutionsNonComplianceReasonISSUE,
		SolutionsNonComplianceReasonIN_HOOK,
		SolutionsNonComplianceReasonOBSOLETE_SPEC,
		SolutionsNonComplianceReasonNO_SPEC,
	}
}

func (e SolutionsNonComplianceReason) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:SolutionsNonComplianceReason", reflect.TypeOf((*SolutionsNonComplianceReason)(nil)).Elem())
}

// Virtual Machine deployment optimization strategies.
type SolutionsVMDeploymentOptimization string

const (
	// Utilizes all cloning methods available, will create initial snapshots
	// on the Virtual Machines.
	SolutionsVMDeploymentOptimizationALL_CLONES = SolutionsVMDeploymentOptimization("ALL_CLONES")
	// Utilize only full copy cloning menthods, will create initial snapshots
	// on the Virtual Machines.
	SolutionsVMDeploymentOptimizationFULL_CLONES_ONLY = SolutionsVMDeploymentOptimization("FULL_CLONES_ONLY")
	// Virtual Machiness will not be cloned from pre-existing deployment.
	SolutionsVMDeploymentOptimizationNO_CLONES = SolutionsVMDeploymentOptimization("NO_CLONES")
)

func (e SolutionsVMDeploymentOptimization) Values() []SolutionsVMDeploymentOptimization {
	return []SolutionsVMDeploymentOptimization{
		SolutionsVMDeploymentOptimizationALL_CLONES,
		SolutionsVMDeploymentOptimizationFULL_CLONES_ONLY,
		SolutionsVMDeploymentOptimizationNO_CLONES,
	}
}

func (e SolutionsVMDeploymentOptimization) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:SolutionsVMDeploymentOptimization", reflect.TypeOf((*SolutionsVMDeploymentOptimization)(nil)).Elem())
}

// Provisioning types for system Virtual Machines.
type SolutionsVMDiskProvisioning string

const (
	// Disks will be provisioned with only used space allocated.
	SolutionsVMDiskProvisioningTHIN = SolutionsVMDiskProvisioning("THIN")
	// Disks will be provisioned with full size allocated.
	SolutionsVMDiskProvisioningTHICK = SolutionsVMDiskProvisioning("THICK")
)

func (e SolutionsVMDiskProvisioning) Values() []SolutionsVMDiskProvisioning {
	return []SolutionsVMDiskProvisioning{
		SolutionsVMDiskProvisioningTHIN,
		SolutionsVMDiskProvisioningTHICK,
	}
}

func (e SolutionsVMDiskProvisioning) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:SolutionsVMDiskProvisioning", reflect.TypeOf((*SolutionsVMDiskProvisioning)(nil)).Elem())
}

// Defines the DRS placement policies applied on the VMs.
type SolutionsVmPlacementPolicy string

const (
	// VMs are anti-affined to each other.
	SolutionsVmPlacementPolicyVM_VM_ANTI_AFFINITY = SolutionsVmPlacementPolicy("VM_VM_ANTI_AFFINITY")
)

func (e SolutionsVmPlacementPolicy) Values() []SolutionsVmPlacementPolicy {
	return []SolutionsVmPlacementPolicy{
		SolutionsVmPlacementPolicyVM_VM_ANTI_AFFINITY,
	}
}

func (e SolutionsVmPlacementPolicy) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("eam:SolutionsVmPlacementPolicy", reflect.TypeOf((*SolutionsVmPlacementPolicy)(nil)).Elem())
}
