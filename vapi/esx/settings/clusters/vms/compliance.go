// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

// ComplianceStatus describes the possible compliance status.
type ComplianceStatus string

const (

	// ComplianceStatusCompliant indicates the status is compliant with the desired solution specification.
	ComplianceStatusCompliant ComplianceStatus = "COMPLIANT"

	// ComplianceStatusNonCompliant indicates the status is non-compliant with the desired solution specification.
	ComplianceStatusNonCompliant ComplianceStatus = "NON_COMPLIANT"

	// ComplianceStatusIncompatible indicates the target state is incompatible with the system.
	ComplianceStatusIncompatible ComplianceStatus = "INCOMPATIBLE"
)

type IssueType string

const (

	// VmPoweredOn indicates the System VM is expected to be powered-off, but it is powered-on.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM power-off the System VM.
	VmPoweredOn IssueType = "VM_POWERED_ON"

	// VmPoweredOff indicates the System VM is expected to be powered-on, but it is powered-off.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM power-on the System VM.
	VmPoweredOff IssueType = "VM_POWERED_OFF"

	// VmSuspended indicates the System VM is expected to be powered-on, but it is suspended.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM power-on the System VM.
	VmSuspended IssueType = "VM_SUSPENDED"

	// VmInaccessible indicates the System VM is expected to be connected, but it is inaccessible.
	// To remediate the deployment, remove the VM and re-invoke Solutions.Apply to have vLCM
	// redeploy the System VM or do the necessary changes to ensure that the connection state
	// of the VM is vim.ConnectionState.connected.
	// NOTE: When the HA is enabled on the cluster this may be transient state and automatically remediated.
	VmInaccessible IssueType = "VM_INACCESSIBLE"

	// VmCorrupted indicates the System VM is corrupted. To remediate the deployment,
	// re-invoke Solutions.Apply to have vLCM delete the System VM and redeploy it.
	VmCorrupted IssueType = "VM_CORRUPTED"

	// VmNotDeployed indicates the System VM has not been deployed because of an unexpected error.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	VmNotDeployed IssueType = "VM_NOT_DEPLOYED"

	// VmNotRemoved indicates the System VM has not been deployed because of an unexpected error.
	// To remediate the deployment, manually remove the VM or re-invoke Solutions.Apply to have vLCM retry the System VM removal.
	VmNotRemoved IssueType = "VM_NOT_REMOVED"

	// VmDatastoreMissing indicates the System VM has not been deployed because the configured datastore
	// for it is missing on the host. To unblock the System VM deployment, provide a proper datastore in the SolutionSpec.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	VmDatastoreMissing IssueType = "VM_DATASTORE_MISSING"

	// VmNetworkMissing indicates the System VM has not been deployed because the configured network
	// for it is missing on the host. To unblock the System VM deployment provide a proper network in the SolutionSpec.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	VmNetworkMissing IssueType = "VM_NETWORK_MISSING"

	// OvfInvalidFormat indicates the System VM has not been deployed because the provided
	// OvfResource part of SolutionSpec contains an invalid OVF. To unblock the System VM deployment
	// provide a new OvfResource with valid OVF. To remediate the deployment, re-invoke Solutions.Apply
	// to have vLCM redeploy the System VM.
	OvfInvalidFormat IssueType = "OVF_INVALID_FORMAT"

	// OvfInvalidProperty indicates the System VM is expected to be deployed or reconfigured, but an OVF
	// property is either missing or has an invalid value. To unblock the System VM deployment or reconfiguration,
	// provide a new ovfDescriptorProperties in the SolutionSpec. To remediate the deployment,
	// re-invoke Solutions.Apply to have vLCM redeploy or reconfigure the System VM.
	OvfInvalidProperty IssueType = "OVF_INVALID_PROPERTY"

	// OvfCannotAccess indicates the System VM has not been deployed because vLCM is not able to
	// access the OVF package. To unblock the System VM deployment, provide a proper accessible OvfResource.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	OvfCannotAccess IssueType = "OVF_CANNOT_ACCESS"

	// OvfCertificateNotTrusted indicates the System VM has not been deployed because vLCM is not able to
	// make successful SSL trust verification of the server certificate when establishing connection to the provided OVF package.
	// To unblock the System VM deployment, provide a valid OvfResource.Certificate or ensure the server certificate is
	// signed by a CA trusted by the system. To remediate the deployment, re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	OvfCertificateNotTrusted IssueType = "OVF_CERTIFICATE_NOT_TRUSTED"

	// InsufficientSpace indicates the System VM has not been deployed because the configured System VM
	// datastore does not have enough free space. To unblock the System VM deployment, make enough free space
	// on the datastore or provide a new datastore in the SolutionSpec. To remediate the deployment,
	// re-invoke Solutions.Apply to have vLCM redeploy the System VM.
	InsufficientSpace IssueType = "INSUFFICIENT_SPACE"

	// InsufficientResources indicates the System VM has not been powered-on because the cluster does not
	// have enough free CPU or memory resources. To unblock the System VM power-on, make enough CPU and memory resources available.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM power-on the System VM.
	InsufficientResources IssueType = "INSUFFICIENT_RESOURCES"

	// HostInMaintenanceMode indicates a System VM operation has not been initiated because the host is in
	// maintenance mode. To unblock the System VM operation, move the host out of maintenance mode.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	HostInMaintenanceMode IssueType = "HOST_IN_MAINTENANCE_MODE"

	// HostInPartialMaintenanceMode indicates a System VM operation has not been initiated because the host is in
	// partial maintenance mode. To unblock the System VM operation, move the host out of partial maintenance mode.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	HostInPartialMaintenanceMode IssueType = "HOST_IN_PARTIAL_MAINTENANCE_MODE"

	// HostInStandbyMode indicates a System VM operation has not been initiated because the host is in
	// stand by mode. To unblock the System VM operation, move the host out of stand by mode.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	HostInStandbyMode IssueType = "HOST_IN_STAND_BY_MODE"

	// HostNotReachable indicates a System VM operation has not been initiated because the host is not
	// reachable from vCenter Server. Any operation on the affected host is not possible.
	// Typical reasons are disconnected or powered-off host. To unblock the System VM operation,
	// reconnect and powered-on the host. To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	HostNotReachable IssueType = "HOST_NOT_REACHABLE"

	// VmInvalidConfig indicates a System VM operation has not been initiated because the System VM
	// has an invalid configuration. To unblock the System VM operation, inspect and correct the System VM configuration as necessary.
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	VmInvalidConfig IssueType = "VM_INVALID_CONFIG"

	// VmsDisabled indicates System VMs are disabled on the cluster via internal ESX Agent
	// Manager (EAM) API (eam.EsxtAgentManager.disable). The present System VMs in the cluster are powered-off.
	// No new System VMs are created. Modifications of the desired specification are not permitted.
	//
	// This issue cannot be remediated via vLCM API. Remediation requires the System VMs to be enabled
	// on the cluster via internal EAM API (eam.EsxAgentManager.enable). As result, the present System VMs
	// are powered-on, Modifications of the desired specification are permitted.
	//
	// Enabling and disabling the System VMs operations on the cluster is operated by vSAN Cluster Shutdown
	// and Start-up workflows. Refer to vSAN Cluster Shutdown and Start-up documentation.
	//
	// NOTE: In future versions of vLCM, Enabling and disabling the System VM operations will happen via internal vLCM APIs.
	VmsDisabled IssueType = "VMS_DISABLED"

	// VmLifecycleHookTimedOut indicates the System VM deployment is not completed because the System VM
	// lifecycle hook has not been processed in the configured LifecycleHookConfig.Timeout.
	// To remediate the deployment, re-invoke Solutions.Apply and process again the VM lifecycle hook.
	VmLifecycleHookTimedOut IssueType = "VM_LIFECYCLE_HOOK_TIMED_OUT"

	// VmLifecycleHookFailed indicates the System VM deployment is not completed because the System VM
	// lifecycle hook has been failed by the client. To remediate the deployment, re-invoke Solutions.Apply
	// and process again the VM lifecycle hook.
	VmLifecycleHookFailed IssueType = "VM_LIFECYCLE_HOOK_FAILED"

	// VmProtected indicates the System VM deployment is not completed because the System VM is
	// protected from modifications (example: VM is in a process of vSphere HA recovery).
	// To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	VmProtected IssueType = "VM_PROTECTED"

	// VmLifecycleHookDynamicUpdateFailed indicates the System VM deployment is not completed because the System VM
	// lifecycle hook dynamic update has failed. To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	VmLifecycleHookDynamicUpdateFailed IssueType = "VM_LIFECYCLE_HOOK_DYNAMIC_UPDATE_FAILED"

	// ClusterTransitionFailed indicates the System VM deployment is not completed because the System VM
	// failed to transition to the target cluster. To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	ClusterTransitionFailed IssueType = "CLUSTER_TRANSITION_FAILED"

	// TransitionFailed indicates the System VM deployment is not completed because the Agent failed
	// to transition to a System VM. To remediate the deployment, re-invoke Solutions.Apply to have vLCM retry the operation.
	TransitionFailed IssueType = "TRANSITION_FAILED"
)

// IssueInfo contains fields that describe an issue that blocks the system to reach the desired
// specification of a given VM deployment.
type IssueInfo struct {

	// Type of the issue.
	Type IssueType `json:"type"`

	// Notifications provides additional information about the issue.
	Notifications rest.Notifications `json:"notifications"`
}

// Status defines how well a deployment conforms to the desired specification that is specified by
// the solutionInfo.
type Status string

const (

	// StatusNotApplied indicates the desired specification of the solution has never been applied.
	StatusNotApplied Status = "NOT_APPLIED"

	// StatusInProgress indicates the system is actively working to reach the desired specification.
	StatusInProgress Status = "IN_PROGRESS"

	// StatusCompliant indicates the deployment is in full compliance with the desired specification.
	StatusCompliant Status = "COMPLIANT"

	// StatusIssue indicates the system has hit issues that do not allow the deployment to reach
	// the desired specification. See Solutions.DeploymentInfo.Issues.
	StatusIssue Status = "ISSUE"

	// StatusInLifecycleHook indicates the system is waiting on an activated VM lifecycle hook to be
	// processed by the solution in order to continue attempting to reach the desired specification.
	// See Solutions.DeploymentInfo.LifecycleHook.
	StatusInLifecycleHook Status = "IN_LIFECYCLE_HOOK"

	// StatusBlocked indicates the system is blocked from reaching the desired specification.
	// For example, this can occur if RemediationPolicy.SEQUENTIAL is set and another deployment is in ISSUE status.
	StatusBlocked Status = "BLOCKED"

	// StatusObsoleteSpec indicates the current desired specification of the solution is newer than the applied.
	//
	// This state should take precedence over:
	//   - Status.BLOCKED
	//   - Status.IN_PROGRESS
	//   - Status.ISSUE
	//   - Status.IN_LIFECYCLE_HOOK
	StatusObsoleteSpec Status = "OBSOLETE_SPEC"
)

// DeploymentInfo contains fields that describe the state of a single VM deployment of a solution.
type DeploymentInfo struct {

	// Status is the compliance status of the deployment.
	Status Status `json:"status"`

	// Vm is the identifier of the currently deployed VM. More information about the
	// runtime state of the VM can be observed through the VIM API.
	//
	// This field is unset if:
	//   - The VM deployment is not started yet.
	//   - There are issues specified by the Issues field that prevents the VM to be deployed.
	Vm *string `json:"vm,omitempty"`

	// ReplacementVm is the identifier of the VM that is going to replace the current deployed VM.
	// More information about the runtime state of the VM can be observed through the VIM API.
	//
	// This field is unset if there is no ongoing VM upgrade for the current VM deployment.
	ReplacementVm *string `json:"replacement_vm,omitempty"`

	// Issues is a list of IssueInfo which do not allow the deployment to reach
	// the desired specification specified by the SolutionInfo. In order to remediate these issues
	// an apply operation Solutions.Apply needs to be initiated.
	Issues []IssueInfo `json:"issues"`

	// LifecycleHook is the activated VM lifecycle hook for the VM specified by the Vm field
	// that the system is waiting to be processed by the solution in order to continue attempting to reach the desired specification.
	//
	// This field is unset if there is no activated hook for the VM.
	LifecycleHook *LifecycleHookInfo `json:"lifecycle_hook,omitempty"`

	// SolutionInfo describes the current desired solution specification of the deployment.
	SolutionInfo SolutionInfo `json:"solution_info"`
}

// DeploymentCompliance contains fields that describe the compliance of a given VM deployment.
// See DeploymentInfo.
type DeploymentCompliance struct {

	// Status is the compliance status of the deployment.
	Status ComplianceStatus `json:"status"`

	// Notifications describing the compliance result.
	Notifications rest.Notifications `json:"notifications"`

	// Deployment is the current VM deployment.
	Deployment DeploymentInfo `json:"deployment"`
}

// HostCompliance contains fields that describe the compliance for a specific host.
type HostCompliance struct {

	// Status is the aggregated compliance status for all solutions for which compliance check was requested.
	Status ComplianceStatus `json:"status"`

	// Compliances for the solutions for which a compliance check was requested.
	Compliances map[string]DeploymentCompliance `json:"compliances"`
}

// HostSolutionsCompliance contains fields that describe the compliance of solutions with deployment
// type DeploymentType.EVERY_HOST_PINNED.
type HostSolutionsCompliance struct {

	// Status is the aggregated compliance status for all solutions for which a compliance check was requested.
	Status ComplianceStatus `json:"status"`

	// Compliances is the compliance status of the hosts that were part of the check compliance operation.
	Compliances map[string]HostCompliance `json:"compliances"`
}

// ClusterSolutionCompliance contains fields that describe the compliance for a specific solution.
type ClusterSolutionCompliance struct {
	// Status is the aggregated compliance status for all deployment units for which a compliance check was requested.
	Status ComplianceStatus `json:"status"`

	// Compliances is the compliance status for the deployment units for which a compliance check was requested.
	Compliances map[string]DeploymentCompliance `json:"compliances"`
}

// ClusterSolutionsCompliance contains fields that describe the compliance of solutions with deployment
// type DeploymentType.CLUSTER_VM_SET.
type ClusterSolutionsCompliance struct {

	// Status is the aggregated compliance status for all solutions for which a compliance check was requested.
	Status ComplianceStatus `json:"status"`

	// Compliances for the solutions for which a compliance check was requested.
	Compliances map[string]ClusterSolutionCompliance `json:"compliances"`
}

// ClusterCompliance contains fields that describe the result of the compliance
// Solutions.CheckCompliance operation.
type ClusterCompliance struct {

	// Status is the aggregated status of the compliance check operation.
	Status ComplianceStatus `json:"status"`

	// HostSolutionsStatus is the compliance status of all solutions with deployment type
	// DeploymentType.EVERY_HOST_PINNED that were part of the Solutions.CheckCompliance operation.
	HostSolutionsStatus HostSolutionsCompliance `json:"host_solutions_status"`

	// ClusterSolutionsStatus is the compliance status of all solutions with deployment type
	// DeploymentType.CLUSTER_VM_SET that were part of the Solutions.CheckCompliance operation.
	ClusterSolutionsStatus ClusterSolutionsCompliance `json:"cluster_solutions_status"`
}

// CheckComplianceFilterSpec contains fields that describe a filter for compliance check in a given cluster.
type CheckComplianceFilterSpec struct {

	// Solutions are identifiers of the solutions that to be checked for compliance.
	//
	// If unset, the compliance is checked for all solutions in the cluster.
	Solutions *[]string `json:"solutions,omitempty"`

	// Hosts are identifiers of the hosts that to be checked for compliance of all
	// solutions with deployment type DeploymentType.EVERY_HOST_PINNED.
	//
	// If unset or empty and DeploymentUnits is unset or empty, the compliance is checked
	// for all hosts in the cluster.
	Hosts *[]string `json:"hosts,omitempty"`

	// DeploymentUnits are identifiers of the deployment units that to be checked for compliance
	// of all solutions with deployment type DeploymentType.CLUSTER_VM_SET.
	//
	// The deployment unit represents a single VM instance deployment.
	//
	// If unset or empty and Hosts is unset or empty, the compliance is checked for
	// all deployment units in the cluster.
	DeploymentUnits *[]string `json:"deployment_units,omitempty"`
}

func (m *Manager) CheckCompliance(ctx context.Context, cluster types.ManagedObjectReference, filterSpec *CheckComplianceFilterSpec) (*ClusterCompliance, error) {
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p).WithParam("action", "check-compliance").WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodPost, filterSpec), &taskId); err != nil {
		return nil, err
	}

	if len(taskId) == 0 {
		return nil, errors.New("no task returned")
	}

	task, err := tasks.NewManager(m.Client).WaitForCompletion(ctx, taskId)
	if err != nil {
		return nil, err
	}

	var compliance ClusterCompliance
	if err := json.Unmarshal(task.Result, &compliance); err != nil {
		return nil, err
	}

	return &compliance, err
}
