// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

type clusterTransitionPath types.ManagedObjectReference

const (
	TransitionPath = clusters.BasePath + "/%s/vms/transition"
)

func (c clusterTransitionPath) String() string {
	cid := types.ManagedObjectReference(c).Value
	return fmt.Sprintf(TransitionPath, cid)
}

// ValidationResult contains fields that describe a validation result.
type ValidationResult struct {
	// Notifications associated with the validation.
	Notifications rest.Notifications `json:"notifications"`
}

// VmSelectionType defines the different selection types for VM selection.
type VmSelectionType string

const (
	// Select System VMs that have a specific property configured in the VM
	// extra configuration. The property has a
	// key='com.vmware.vim.eam.selection'. The value represents a unique
	// identifier used for VM selection and is provided by the client.
	VmSelectionTypeVmExtraConfig VmSelectionType = "VM_EXTRA_CONFIG"
)

// The VmSelectionSpec structure contains fields
// to describe the criteria used to select System VMs to which an
// AlternativeVmSpec configuration is applied.
type VmSelectionSpec struct {
	// Selection type for this VmSelectionSpec.
	SelectionType VmSelectionType `json:"selection_type"`

	// Unique VM extra configuration property value. The recommended usage is
	// with an UUID.
	//
	// See VmSelectionTypeVmExtraConfig for more details.
	ExtraConfigValue string `json:"extra_config_value"`
}

// TransitionSpec contains fields that describe the specification for transitioning a System VM Solution.
type TransitionSpec struct {
	// Cluster to transition from.
	SourceCluster types.ManagedObjectReference `json:"source_cluster"`

	// Target desired solution specification in vLCM.
	Solution *SolutionSpec `json:"solution"`
}

// EnableSpec contains fields that describe specification for enablement of EAM managed solution in vLCM.
type EnableSpec struct {
	// Identifier of the solution in EAM (EAM agency).
	EamAgencyID string `json:"eam_agency_id"`

	// Target desired solution specification in vLCM.
	Solution *SolutionSpec `json:"solution"`
}

// MultiSourceEnableSpec contains fields that describe specification for enablement of multiple EAM managed solutions into single vLCM managed solution.
// Supported only for solutions with deployment type DeploymentType#CLUSTER_VM_SET.
type MultiSourceEnableSpec struct {
	// EAM Agency identifiers.
	EamAgencyIDs []string `json:"eam_agency_ids"`

	// Target desired solution specification in vLCM.
	// The given SolutionSpec should not contain any AlternativeVmSpecs. See SourceAlternativeVmSpecs
	// about how to configure AlternativeVmSpecs.
	Solution *SolutionSpec `json:"solution"`

	// Relation between System VMs and their respective VmSelectionSpecs.
	// Provided VM IDs must be part of the solution being transitioned and
	// must exist in the cluster where the solution is installed.
	// Provided VmSelectionSpecs must be present in the applied desired state as part of the
	// ClusterSolutionSpec#alternativeVmSpecs.
	//
	// If unset, no VmSelectionSpecs are applied on the source agencies' System VMs during the enablement.
	SourceVmSelectionSpecs map[string]*VmSelectionSpec `json:"source_vm_selection_specs,omitempty"`

	// Cluster module to be reused for transitioned System VMs. Used to express
	// VM-VM anti affinity relation between System VMs in the vSphere Cluster.
	// The module must exist for the cluster where the solution is installed.
	//
	// If unset, no cluster module is reused. vLCM creates a new module if needed.
	ClusterModule string `json:"cluster_module,omitempty"`
}

// Enable enables an EAM managed solution in vLCM. The solution specification is
// validated before the enablement is started.
//
// The enablement only transfers ownership of the solution from EAM to LCCM
// and sets the desired state in LCCM. The new desired state is not applied,
// the solution system VMs are untouched.
//
// The following happens once the operation is started:
// - A removal of the corresponding agency in EAM is triggered.
//
// The following happens once the operation is completed:
// - The corresponding agency in EAM can no longer be controlled through the EAM API.
// - The management of the desired solution specification can be done only through vLCM. See Solutions
//
// @param cluster Identifier of the cluster.
// @param solution Identifier of the solution.
// @param spec Enablement specification.
//
// @throws Error If there is an internal error. The accompanying error message will give more details about the failure.
// @throws InvalidArgument If the validation of the cluster desired specification fails, the solution associated with solution already exists in vLCM for the cluster, or the agency associated with the eamAgencyID exists, but is not in the scope of the cluster. The value of the data field of Error contains more information. It is a structure that contains all the fields defined in ValidationResult.
// @throws NotAllowedInCurrentState If System VMs are disabled on the cluster via internal ESX Agent Manager (EAM) API (eam.EsxtAgentManager#disable).
// @throws Unsupported If the cluster associated with cluster is not managed by vLCM.
// @throws NotFound If there is no cluster associated with the cluster or if there is no EAM agency associated with the eamAgencyID field of the spec in the system.
// @throws ServiceUnavailable If the service is not available.
// @throws Unauthenticated If the caller is not authenticated.
// @throws Unauthorized If the user doesn't have the required privileges.
func (m *Manager) Enable(ctx context.Context, cluster types.ManagedObjectReference, solution string, spec *EnableSpec) error {
	p := clusterTransitionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution).WithParam("action", "enable").WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodPost, spec), &taskId); err != nil {
		return err
	}

	_, err := tasks.NewManager(m.Client).WaitForRunningOrError(ctx, taskId)
	return err
}

// MultiSourceEnable enables multiple EAM managed solutions in vLCM as a single solution. The
// solution specification is validated before the enablement is started.
//
// The enablement only transfers ownership of the solutions from EAM to LCCM
// and sets the desired state in LCCM. The new desired state is not applied,
// the solution system VMs are untouched.
//
// The following happens once the operation is started:
// - A removal of the corresponding agencies in EAM is triggered.
//
// The following happens once the operation is completed:
// - The corresponding agencies in EAM can no longer be controlled through the EAM API.
// - The management of the desired solution specification can be done only through vLCM. See Solutions
//
// Supported only for solutions with deployment type DeploymentType#CLUSTER_VM_SET.
//
// @param cluster Identifier of the cluster.
// @param solution Identifier of the solution.
// @param spec Enablement specification.
//
// @throws Error If there is an internal error. The accompanying error message will give more details about the failure.
// @throws InvalidArgument If the validation of the cluster desired specification fails, the solution associated with solution already exists in vLCM for the cluster, or the agencies associated with the eamAgencyIDs exist, but are not in the scope of the cluster. The value of the data field of Error contains more information. It is a structure that contains all the fields defined in ValidationResult.
// @throws NotAllowedInCurrentState If System VMs are disabled on the cluster via internal ESX Agent Manager (EAM) API (eam.EsxtAgentManager#disable).
// @throws Unsupported If the cluster associated with cluster is not managed by vLCM.
// @throws NotFound If there is no cluster associated with the cluster or if there is no EAM agency associated with any of the IDs in eamAgencyIDs field of the spec in the system.
// @throws ServiceUnavailable If the service is not available.
// @throws Unauthenticated If the caller is not authenticated.
// @throws Unauthorized If the user doesn't have the required privileges.
func (m *Manager) MultiSourceEnable(ctx context.Context, cluster types.ManagedObjectReference, solution string, spec *MultiSourceEnableSpec) error {
	p := clusterTransitionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution).WithParam("action", "multi-source-enable").WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodPost, spec), &taskId); err != nil {
		return err
	}

	_, err := tasks.NewManager(m.Client).WaitForRunningOrError(ctx, taskId)
	return err
}

// Transition transitions a System VM Solution desired state to a target cluster. The
// solution specification is validated before the transition is started.
//
// The operation only initiates the transition. The target desired state is
// not applied and the solution system VMs remain untouched. A consecutive
// Solutions#apply operation is needed to complete the transition.
//
// Once the operation is completed:
// - The desired state for the solution is set to the target cluster.
// - The solution can be managed only on the target cluster.
//
// @param cluster Identifier of the target cluster.
// @param solution Identifier of the solution.
// @param spec Transition specification.
//
// @throws Error If there is an internal error. The accompanying error message will give more details about the failure.
// @throws InvalidArgument If the validation of the cluster desired specification fails, the solution associated with solution does not exist in vLCM for sourceCluster field, the solution associated with solution already exists in vLCM for the new target cluster associated with cluster, or the target cluster is not DRS enabled. The value of the data field of Error contains more information. It is a structure that contains all the fields defined in ValidationResult.
// @throws NotAllowedInCurrentState If System VMs are disabled on the cluster via internal ESX Agent Manager (EAM) API (eam.EsxAgentManager#disable).
// @throws Unsupported If the cluster associated with cluster is not managed by vLCM.
// @throws NotFound If there is no cluster associated with the cluster.
// @throws ServiceUnavailable If the service is not available.
// @throws Unauthenticated If the caller is not authenticated.
// @throws Unauthorized If the user doesn't have the required privileges.
func (m *Manager) Transition(ctx context.Context, cluster types.ManagedObjectReference, solution string, spec *TransitionSpec) error {
	p := clusterTransitionPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution).WithParam("action", "transition").WithParam("vmw-task", "true")
	var taskId string

	if err := m.Do(ctx, url.Request(http.MethodPost, spec), &taskId); err != nil {
		return err
	}

	_, err := tasks.NewManager(m.Client).WaitForRunningOrError(ctx, taskId)
	return err
}
