// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vim25/types"
)

//	The {@name HostSolutionsApplyFilterSpec} {@term structure} contains {@term
//
// fields} that describe a filter that to be used for applying the desired
// specification of solutions with deployment type {@link
// DeploymentType#EVERY_HOST_PINNED} to a given cluster.
type HostSolutionsApplyFilterSpec struct {

	/**
	 * Specific solutions within the cluster to be considered during the
	 * execution of the apply {@term operation}.
	 *
	 * @field.optional if {@term unset} or empty, all solutions are applied.
	 */
	Solutions []string `json:"solutions"`

	/**
	 * Hosts on which solutions that are specified by this structure need
	 * to be applied.
	 *
	 * @field.optional if {@term unset} or empty, the solutions are applied on
	 *                 all of the hosts in the cluster.
	 */
	Hosts []string `json:"hosts"`
}

/**
 * The {@name ClusterSolutionsApplyFilterSpec} {@term structure} contains
 * {@term fields} that describe a filter that to be used for applying the
 * desired specification of solutions with deployment type
 * {@link DeploymentType#CLUSTER_VM_SET} to a given cluster.
 */
type ClusterSolutionsApplyFilterSpec struct {

	/**
	 * Specific solutions within the cluster to be considered during the
	 * execution of the apply {@term operation}.
	 *
	 * @field.optional if {@term unset} or empty, all solutions are applied.
	 */
	Solutions []string `json:"solutions,omitempty"`

	/**
	 * Hosts on which solutions that are specified by this structure need
	 * to be applied.
	 *
	 * @field.optional if {@term unset} or empty, the solutions are applied on
	 *                 all of the hosts in the cluster.
	 */
	Hosts []string `json:"hosts,omitempty"`
}

// The {@name ApplySpec} {@term structure} contains {@term fields} that
// describe a specification to be used for applying the desired solution
// specification to a given cluster.
type ApplySpec struct {

	/**
	 * Apply filter for solutions with deployment type
	 * {@link DeploymentType#EVERY_HOST_PINNED}.
	 *
	 * @field.optional if {@term unset} or empty and
	 *                 {#member clusterSolutions} is {@term unset} or empty,
	 *                 all solutions are applied on the cluster.
	 */
	HostSolutions *HostSolutionsApplyFilterSpec `json:"host_solutions,omitempty"`

	/**
	 * Apply filter for solutions with deployment type
	 * {@link DeploymentType#CLUSTER_VM_SET}.
	 *
	 * @field.optional if {@term unset} or empty and {#member  hostSolutions}
	 *                 is {@term unset} or empty, all solutions are applied on
	 *                 the cluster.
	 */
	ClusterSolutions *ClusterSolutionsApplyFilterSpec `json:"cluster_solutions,omitempty"`
}

const (
	/**
	 * The {@name Status} {@term enumerated type} contains the status codes of
	 * an {@link Solutions#apply} {@term operation}.
	 */

	/**
	 * The apply {@term operation} completed successfully.
	 */
	Success Status = "SUCCESS"

	/**
	 * The apply {@term operation} encountered an error.
	 */
	Error Status = "ERROR"
)

/**
 * The {@name ApplyStatus} {@term structure} contains {@term fields} that
 * describe the status of an {@link #apply} {@term operation}.
 */
type ApplyStatus struct {
	Status Status `json:"status"`

	/**
	 * The vLCM system time when the {@term operation} started.
	 */
	StartTime time.Time `json:"start_time"`

	/**
	 * The vLCM system time when the {@term operation} completed.
	 */
	EndTime time.Time `json:"end_time"`
}

/**
 * The {@name HostApplyStatus} {@term structure} contains {@term fields} that
 * describe the apply status for a specific host.
 */
type HostApplyStatus struct {

	/**
	 * Aggregated apply status of the solutions on the host.
	 *
	 * @field.optional {@term unset} if the apply {@term operation} is not
	 *                 completed for the specified host.
	 */
	Status ApplyStatus `json:"status"`

	/**
	 * The apply status of the different solutions on the host.
	 */
	SolutionStatus map[string]ApplyStatus `json:"solution_statuses"`
}

/**
 * The {@name HostSolutionsApplyStatus} {@term structure} contains
 * {@term fields} that describe the apply status of solutions with deployment
 * type {@link DeploymentType#EVERY_HOST_PINNED}.
 */
type HostSolutionsApplyStatus struct {

	/**
	 * Aggregated apply status of the solutions.
	 *
	 * @field.optional {@term unset} if the apply {@term operation} is not
	 *                 completed for solutions with deployment type
	 *                 {@link DeploymentType#EVERY_HOST_PINNED}.
	 */
	Status ApplyStatus `json:"status"`

	/**
	 * The apply status of the hosts that were part of the apply
	 * {@term operation}.
	 */
	HostStatuses map[string]HostApplyStatus `json:"hostStatuses"`
}

/**
 * The {@name ClusterSolutionApplyStatus} {@term structure} contains
 * {@term fields} that describe the apply status for a specific solution.
 */
type ClusterSolutionApplyStatus struct {

	/**
	 * Aggregated apply status for the deployment units of the solution.
	 *
	 * @field.optional {@term unset} if the apply {@term operation} is not
	 *                 completed for the specified deployment unit.
	 */
	Status ApplyStatus `json:"status"`

	/**
	 * The apply status for the different deployment units of the solution.
	 */
	DeploymentUnitStatuses map[string]ApplyStatus `json:"deployment_unit_statuses"`
}

/**
 * The {@name ClusterSolutionsApplyStatus} {@term structure} contains
 * {@term fields} that describe the apply status of solutions with deployment
 * type {@link DeploymentType#CLUSTER_VM_SET}.
 */
type ClusterSolutionsApplyStatus struct {

	/**
	 * Aggregated apply status of the solutions.
	 *
	 * @field.optional {@term unset} if the apply {@term operation} is not
	 *                 completed for solutions with deployment type
	 *                 {@link DeploymentType#CLUSTER_VM_SET}.
	 */
	Status ApplyStatus `json:"status"`

	/**
	 * The apply status of the solutions that were part of the apply
	 * {@term operation}.
	 */
	SolutionStatuses map[string]ClusterSolutionApplyStatus `json:"solution_statuses"`
}

// The {@name ApplyResult} {@term structure} contains {@term fields} that
// describe the result of an {@link #apply} {@term operation}.
type ApplyResult struct {

	/**
	 * Aggregated status of an apply {@term operation}.
	 *
	 * @field.optional {@term unset} if the apply {@term operation} is in
	 *                 progress.
	 */
	ApplyStatus ApplyStatus `json:"status"`

	/**
	 * The apply status of all solutions with deployment type
	 * {@link DeploymentType#EVERY_HOST_PINNED} that were part of the apply
	 * {@term operation}.
	 */
	HostSolutionStatus HostSolutionsApplyStatus `json:"host_solutions_status"`

	/**
	 * The apply status of all solutions with deployment type
	 * {@link DeploymentType#CLUSTER_VM_SET} that were part of the apply
	 * {@term operation}.
	 */
	ClusterSolutionStatus ClusterSolutionsApplyStatus `json:"cluster_solutions_status"`
}

// Apply applies the given solution and returns the taskid.
func (m *Manager) Apply(ctx context.Context, cluster types.ManagedObjectReference, applySpec *ApplySpec) (string, error) {
	p := clusterSolutionPath(cluster).String()
	url := m.Resource(p).WithParam("action", "apply").WithParam("vmw-task", "true")
	var taskId string

	return taskId, m.Do(ctx, url.Request(http.MethodPost, applySpec), &taskId)
}

// ApplyWaitForCompletion waits for the apply task to complete and returns the
// result from the apply task.
func (m *Manager) ApplyWaitForCompletion(ctx context.Context, taskId string) (*ApplyResult, error) {
	task, err := tasks.NewManager(m.Client).WaitForCompletion(ctx, taskId)
	if err != nil {
		return nil, err
	}

	var result ApplyResult
	if err := json.Unmarshal(task.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
