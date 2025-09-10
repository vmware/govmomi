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

// The {@name ComplianceStatus} {@term enumerated type} describes the possible
// compliance status.
type ComplianceStatus string

const (

	//
	// The status is compliant with the desired solution specification.
	///
	ComplianceStatusCompliant ComplianceStatus = "COMPLIANT"

	//
	// The status is <b>non</b>-compliant with the desired solution
	// specification.
	///
	ComplianceStatusNonCompliant ComplianceStatus = "NON_COMPLIANT"

	//
	// Target state is incompatible with the system.
	///
	ComplianceStatusIncompatible ComplianceStatus = "INCOMPATIBLE"
)

type IssueType string

const (

	//
	// The System VM is expected to be powered-off, but it is powered-on.
	// To remediate the deployment, re-invoke {@link Solutions#apply} to
	// have vLCM power-off the System VM.
	///
	VmPoweredOn IssueType = "VM_POWERED_ON"

	//
	// The System VM is expected to be powered-on, but it is powered-off.
	// To remediate the deployment, re-invoke {@link Solutions#apply} to
	// have vLCM power-on the System VM.
	///
	VmPoweredOff IssueType = "VM_POWERED_OFF"

	//
	// The System VM is expected to be powered-on, but it is suspended. To
	// remediate the deployment, re-invoke {@link Solutions#apply} to have
	// vLCM power-on the System VM.
	///
	VmSuspended IssueType = "VM_SUSPENDED"

	//
	// The System VM is expected to be connected, but it is inaccessible.
	// To remediate the deployment, remove the VM and re-invoke
	// {@link Solutions#apply} to have vLCM redeploy the System VM or do
	// the necessary changes to ensure that the connection state of the VM
	// is (vim.ConnectionState#connected).
	// <p>
	// NOTE: When the HA is enabled on the cluster this may be transient
	// state and automatically remediated.
	///
	VmInaccessible IssueType = "VM_INACCESSIBLE"

	//
	// The System VM is corrupted. To remediate the deployment, re-invoke
	// {@link Solutions#apply} to have vLCM delete the System VM and
	// redeploys it.
	///
	VmCorrupted IssueType = "VM_CORRUPTED"

	//
	// The System VM has not been deployed because of an unexpected error.
	// To remediate the deployment, re-invoke {@link Solutions#apply} to
	// have vLCM redeploy the System VM.
	///
	VmNotDeployed IssueType = "VM_NOT_DEPLOYED"

	//
	// The System VM has not been deployed because of an unexpected error.
	// To remediate the deployment, manually remove the VM or re-invoke
	// {@link Solutions#apply} to have vLCM retry the System VM removal.
	///
	VmNotRemoved IssueType = "VM_NOT_REMOVED"

	//
	// The System VM has not been deployed because the configured datastore
	// for it is missing on the host. To unblock the System VM deployment,
	// provide a proper datastore in the {@name SolutionSpec}. To
	// remediate the deployment, re-invoke {@link Solutions#apply} to have
	// vLCM to redeploy the System VM.
	///
	VmDatastoreMissing IssueType = "VM_DATASTORE_MISSING"

	//
	// The System VM has not been deployed because the configured network
	// for it is missing on the host. To unblock the System VM deployment
	// provide a proper network in the {@name SolutionSpec}. To remediate
	// the deployment, re-invoke {@link Solutions#apply} to have vLCM
	// redeploy the System VM.
	///
	VmNetworkMissing IssueType = "VM_NETWORK_MISSING"

	//
	// The System VM has not been deployed because the provided
	// {@name OvfResource} part of {@name SolutionSpec} contains an invalid
	// OVF. To unblock the System VM deployment provide a new
	// {@name OvfResource} with valid OVF. To remediate the deployment,
	// re-invoke {@link Solutions#apply} to have vLCM redeploy the System
	// VM.
	///
	OvfInvalidFormat IssueType = "OVF_INVALID_FORMAT"

	//
	// The System VM is expected to be deployed or reconfigured, but an OVF
	// property is either missing or has an invalid value. To unblock the
	// System VM deployment or reconfiguration, provide a new
	// {@name ovfDescriptorProperties} in the {@name SolutionSpec}. To
	// remediate the deployment, re-invoke {@link Solutions#apply} to have
	// vLCM redeploy or reconfigure the System VM.
	///
	OvfInvalidProperty IssueType = "OVF_INVALID_PROPERTY"

	//
	// The System VM has not been deployed because vLCM is not able to
	// access the OVF package. To unblock the System VM deployment,
	// provide a proper accessible {@name OvfResource}. To remediate the
	// deployment, re-invoke {@link Solutions#apply} to have vLCM redeploy
	// the System VM.
	///
	OvfCannotAccess IssueType = "OVF_CANNOT_ACCESS"

	//
	// The System VM has not been deployed because vLCM is not able to
	// make successful SSL trust verification of the server certificate
	// when establishing connection to the provided OVF package. To unblock
	// the System VM deployment, provide a valid
	// {@name OvfResource#certificate} or ensure the server certificate is
	// signed by a CA trusted by the system. To remediate the deployment,
	// re-invoke {@link Solutions#apply} to have vLCM redeploy the System
	// VM.
	///
	OvfCertificateNotTrusted IssueType = "OVF_CERTIFICATE_NOT_TRUSTED"

	//
	// The System VM has not been deployed because the configured System VM
	// datastore does not have enough free space. To unblock the System VM
	// deployment, make enough free space on the datastore or provide a
	// new datastore in the {@name SolutionSpec}. To remediate the
	// deployment, re-invoke {@link Solutions#apply} to have vLCM redeploy
	// the System VM.
	///
	InsufficientSpace IssueType = "INSUFFICIENT_SPACE"

	//
	// The System VM has not been powered-on because the cluster does not
	// have enough free CPU or memory resources. To unblock the System VM
	// power-on, make enough CPU and memory resources available. To
	// remediate the deployment, re-invoke {@link Solutions#apply} to have
	// vLCM power-on the System VM.
	///
	InsufficientResources IssueType = "INSUFFICIENT_RESOURCES"

	//
	// A System VM operation has not been initiated because the host is in
	// maintenance mode. To unblock the System VM operation, move the host
	// out of maintenance mode. To remediate the deployment, re-invoke
	// {@link Solutions#apply} to have vLCM retry the operation.
	///
	HostInMaintenanceMode IssueType = "HOST_IN_MAINTENANCE_MODE"

	//
	// A System VM operation has not been initiated because the host is in
	// partial maintenance mode. To unblock the System VM operation, move
	// the host out of partial maintenance mode. To remediate the
	// deployment, re-invoke {@link Solutions#apply} to have vLCM retry the
	// operation.
	///
	HostInPartialMaintenanceMode IssueType = "HOST_IN_PARTIAL_MAINTENANCE_MODE"

	//
	// A System VM operation has not been initiated because the host is in
	// stand by mode. To unblock the System VM operation, move the host
	// out of stand by mode. To remediate the deployment, re-invoke
	// {@link Solutions#apply} to have vLCM retry the operation.
	///
	HostInStandbyMode IssueType = "HOST_IN_STAND_BY_MODE"

	//
	// A System VM operation has not been initiated because the host is not
	// reachable from vCenter Server. Any operation on the affected host is
	// not possible. Typical reasons are disconnected or powered-off host.
	// To unblock the System VM operation, reconnect and powered-on the
	// host. To remediate the deployment, re-invoke {@link Solutions#apply}
	// to have vLCM retry the operation.
	///
	HostNotReachable IssueType = "HOST_NOT_REACHABLE"

	//
	// A System VM operation has not been initiated because the System VM
	// has an invalid configuration. To unblock the System VM operation,
	// inspect and correct the System VM configuration as necessary. To
	// remediate the deployment, re-invoke {@link Solutions#apply} to have
	// vLCM retry the operation.
	///
	VmInvalidConfig IssueType = "VM_INVALID_CONFIG"

	//
	// System VMs are disabled on the cluster via internal ESX Agent
	// Manager (EAM) API (eam.EsxtAgentManager#disable). The present
	// System VMs in the cluster are powered-off. No new System VMs are
	// created. Modifications of the desired speification are not
	// permitted.
	// <p>
	// This issue cannot be remediated via vLCM API. Remediation requires
	// the System VMs to be enabled on the cluster via internal EAM API
	// (eam.EsxAgentManager#enable). As result, the present System
	// VMs are powered-on, Modifications of the desired specification are
	// permitted.
	// <p>
	// Enabling and disabling the System VMs operations on the cluster is
	// operated by vSAN Cluster Shutdown and Start-up workflows. Refer to
	// vSAN Cluster Shutdown and Start-up documentation.
	// <p>
	// NOTE: In future versions of vLCM, Enabling and disabling the System
	// VM operations will happen via internal vLCM APIs.
	///
	VmsDisabled IssueType = "VMS_DISABLED"

	//
	// The System VM deployment is not completed because the System VM
	// lifecycle hook has not been processed in the configured
	// {@link LifecycleHookConfig#timeout}. To remediate the deployment,
	// re-invoke {@link Solutions#apply} and process again the VM lifecycle
	// hook.
	///
	VmLifecycleHookTimedOut IssueType = "VM_LIFECYCLE_HOOK_TIMED_OUT"

	//
	// The System VM deployment is not completed because the System VM
	// lifecycle hook has been failed by the client. To remediate the
	// deployment, re-invoke {@link Solutions#apply} and process again the
	// VM lifecycle hook.
	///
	VmLifecycleHookFailed IssueType = "VM_LIFECYCLE_HOOK_FAILED"

	//
	// The System VM deployment is not completed because the System VM is
	// protected from modifications (example: VM is in a process of vSphere
	// HA recovery). To remediate the deployment, re-invoke
	// {@link Solutions#apply} to have vLCM retry the operation.
	///
	VmProtected IssueType = "VM_PROTECTED"

	//
	// The System VM deployment is not completed because the System VM
	// lifecycle hook dynamic update has failed. To remediate the
	// deployment, re-invoke {@link Solutions#apply} to have vLCM retry the
	// operation.
	///
	VmLifecycleHookDynamicUpdateFailed IssueType = "VM_LIFECYCLE_HOOK_DYNAMIC_UPDATE_FAILED"

	//
	// The System VM deployment is not completed because the System VM
	// failed to transition to the target cluster. To remediate the
	// deployment, re-invoke {@link Solutions#apply} to have vLCM retry the
	// operation.
	///
	ClusterTransitionFailed IssueType = "CLUSTER_TRANSITION_FAILED"

	//
	// The System VM deployment is not completed because the Agent failed
	// to transition to a System VM. To remediate the deployment, re-invoke
	// {@link Solutions#apply} to have vLCM retry the operation.
	///
	TransitionFailed IssueType = "TRANSITION_FAILED"
)

// The {@name IssueInfo} {@term structure} contains {@term fields} that
// describe an issue that blocks the system to reach the desired
// specification of a given VM deployment.
// /
type IssueInfo struct {

	//
	// The {@name IssueType} {@term enumerated type} defines the type of the
	// issues.
	///
	//
	// Type of the issue.
	///
	Type IssueType `json:"type"`

	//
	// Provides additional information about the issue.
	///
	Notifications rest.Notifications `json:"notifications"`
}

// The {@name Status} {@term enumerated type} defines how well a
// deployment conforms to the desired specification that is specified by
// the {@name #solutionInfo}.
///

type Status string

const (

	//
	// The desired specification of the solution has never been applied.
	///
	StatusNotApplied Status = "NOT_APPLIED"

	//
	// The system is actively working to reach the desired specification.
	///
	StatusInProgress Status = "IN_PROGRESS"

	//
	// The deployment is in full compliance with the desired specification.
	///
	StatusCompliant Status = "COMPLIANT"

	//
	// The system has hit issues that do not allow the deployment to reach
	// the desired specification. See
	// {@link Solutions.DeploymentInfo#issues}.
	///
	StatusIssue Status = "ISSUE"

	//
	// The system is waiting on an activated VM lifecycle hook to be
	// processed by the solution in order to continue attempting to reach
	// the desired specification. See
	// {@link Solutions.DeploymentInfo#lifecycleHook}.
	///
	StatusInLifecycleHook Status = "IN_LIFECYCLE_HOOK"

	//
	// The system is blocked from reaching the desired specification. For
	// example, this can occur if {@link RemediationPolicy#SEQUENTIAL} is
	// set and another deployment is in {@name #ISSUE} status.
	///
	StatusBlocked Status = "BLOCKED"

	//
	// The current desired specification of the solution is newer than the
	// applied.
	//
	// <ul>
	// This state should take precedence over:
	// <feature name="MultipleClustersPerVsphereZone">
	// <li>{@name Status#BLOCKED}</li></feature>
	// <li>{@name Status#IN_PROGRESS}</li>
	// <li>{@name Status#ISSUE}</li>
	// <li>{@name Status#IN_LIFECYCLE_HOOK}</li>
	// </ul>
	///
	StatusObsoleteSpec Status = "OBSOLETE_SPEC"
)

// The {@name DeploymentInfo} {@term structure} contains {@term fields} that
// describe the state of a single VM deployment of a solution.
// /
type DeploymentInfo struct {

	//
	// Compliance status of the deployment.
	///
	Status Status `json:"status"`

	//
	// Identifier of the currently deployed VM. More information about the
	// runtime state of the VM can be observed through the VIM API.
	//
	// @field.optional This field is {@term unset} if:
	//                 <ul>
	//                 <li>The VM deployment is not started yet.</li>
	//                 <li>There are issues specified by the {@name #issues}
	//                 that prevents the VM to be deployed.</li>
	//                 </ul>
	///
	Vm *string `json:"vm,omitempty"`

	//
	// Identifier of the VM that is going to replace the current deployed VM.
	// More information about the runtime state of the VM can be observed
	// through the VIM API.
	//
	// @field.optional This field is {@term unset} if there is no ongoing VM
	//                 upgrade for the current VM deployment.
	///
	ReplacementVm *string `json:"replacement_vm,omitempty"`

	//
	// List of {@link IssueInfo} which do not allow the deployment to reach
	// the desired specification specified by the {@name #solutionInfo}. In
	// order to remediate these issues an apply operation
	// {@link Solutions#apply} need to be initiated.
	///
	Issues []IssueInfo `json:"issues"`

	//
	// The activated VM lifecycle hook for the VM specified by the {@name #vm}
	// that the system is waiting to be processed by the solution in order to
	// continue attempting to reach the desired specification.
	//
	// @field.optional This field is {@term unset} if there is no activated
	//                 hook for the VM.
	///
	LifecycleHook *LifecycleHookInfo `json:"lifecycle_hook,omitempty"`

	//
	// Describes the current desired solution specification of the deployment.
	///
	SolutionInfo SolutionInfo `json:"solution_info"`
}

// The {@name DeploymentCompliance} {@term structure} contains {@term fields}
// that describe the compliance of a given VM deployment. See
// {@link DeploymentInfo}
// /
type DeploymentCompliance struct {

	//
	// Compliance status of the deployment.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Notifications describing the compliance result.
	///
	Notifications rest.Notifications `json:"notifications"`

	//
	// The current VM deployment.
	///
	Deployment DeploymentInfo `json:"deployment"`
}

// The {@name HostCompliance} {@term structure} contains {@term fields} that
// describe the compliance for a specific host.
// /
type HostCompliance struct {

	//
	// Aggregated compliance status for all solutions for which compliance
	// check was requested.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Compliance for the solutions for which a compliance check was
	// requested.
	///
	Compliances map[string]DeploymentCompliance `json:"compliances"`
}

// The {@name HostSolutionsCompliance} {@term structure} contains
// {@term fields} that describe the compliance of solutions with deployment
// type {@link DeploymentType#EVERY_HOST_PINNED}.
// /
type HostSolutionsCompliance struct {

	//
	// Aggregated compliance status for all solutions for which a compliance
	// check was requested.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Compliance status of the hosts that were part of the check compliance
	// {@term operation}.
	///
	Compliances map[string]HostCompliance `json:"compliances"`
}

// The {@name ClusterSolutionCompliance} {@term structure} contains
// {@term fields} that describe the compliance for a specific solution.
// /
type ClusterSolutionCompliance struct {
	//
	// Aggregated compliance status for all deployment units for which a
	// compliance check was requested.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Compliance status for the deployment units for which a compliance check
	// was requested.
	///
	Compliances map[string]DeploymentCompliance `json:"compliances"`
}

// The {@name ClusterSolutionsCompliance} {@term structure} contains
// {@term fields} that describe the compliance of solutions with deployment
// type {@link DeploymentType#CLUSTER_VM_SET}.
// /
type ClusterSolutionsCompliance struct {

	//
	// Aggregated compliance status for all solutions for which a compliance
	// check was requested.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Compliance for the solutions for which a compliance check was
	// requested.
	///
	Compliances map[string]ClusterSolutionCompliance `json:"compliances"`
}

// The {@name ClusterCompliance} {@term structure} contains {@term fields}
// that describe the result of the compliance
// {@link Solutions#checkCompliance} {@term operation}.
// /
type ClusterCompliance struct {

	//
	// Aggregated status of the compliance check {@term operation}.
	///
	Status ComplianceStatus `json:"status"`

	//
	// Compliance status of all solutions with deployment type
	// {@link DeploymentType#EVERY_HOST_PINNED} that were part of the
	// {@link Solutions#checkCompliance} {@term operation}.
	///
	HostSolutionsStatus HostSolutionsCompliance `json:"host_solutions_status"`

	//
	// Compliance status of all solutions with deployment type
	// {@link DeploymentType#CLUSTER_VM_SET} that were part of the
	// {@link Solutions#checkCompliance} {@term operation}.
	///
	ClusterSolutionsStatus ClusterSolutionsCompliance `json:"cluster_solutions_status"`
}

// The {@name CheckComplianceFilterSpec} {@term structure} contains {@term
// fields} that describe a filter for compliance check in a given cluster.
type CheckComplianceFilterSpec struct {

	/**
	 * Identifiers of the solutions that to be checked for compliance.
	 *
	 * @field.optional If {@term unset}, the compliance is checked for all
	 *                 solutions in the cluster.
	 */
	Solutions []string `json:"solutions"`

	/**
	 * Identifiers of the hosts that to be checked for compliance of all
	 * solutions with deployment type
	 * {@link DeploymentType#EVERY_HOST_PINNED}.
	 *
	 * @field.optional If {@term unset} or empty and {#member deploymentUnits}
	 *                 is {@term unset} or empty, the compliance is checked
	 *                 for all hosts in the cluster.
	 */
	Hosts []string `json:"hosts"`

	/**
	 * Identifiers of the deployment units that to be checked for compliance
	 * of all solutions with deployment type
	 * {@link DeploymentType#CLUSTER_VM_SET}.
	 * <p>
	 * The deployment unit represents a single VM instance deployment.
	 *
	 * @field.optional If {@term unset} or empty and {#member hosts} is
	 *                 {@term unset} or empty, the compliance is checked for
	 *                 all deployment units in the cluster.
	 */
	DeploymentUnits []string `json:"deployment_units"`
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
