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

// HostSolutionsApplyFilterSpec contains fields that describe a filter that to be used for applying the desired
// specification of solutions with deployment type DeploymentType.EVERY_HOST_PINNED to a given cluster.
type HostSolutionsApplyFilterSpec struct {

	// Solutions are specific solutions within the cluster to be considered during the
	// execution of the apply operation.
	//
	// If unset or empty, all solutions are applied.
	Solutions *[]string `json:"solutions,omitempty"`

	// Hosts on which solutions that are specified by this structure need
	// to be applied.
	//
	// If unset or empty, the solutions are applied on all of the hosts in the cluster.
	Hosts *[]string `json:"hosts,omitempty"`
}

// ClusterSolutionsApplyFilterSpec contains fields that describe a filter that to be used for applying the
// desired specification of solutions with deployment type DeploymentType.CLUSTER_VM_SET to a given cluster.
type ClusterSolutionsApplyFilterSpec struct {

	// Solutions are specific solutions within the cluster to be considered during the
	// execution of the apply operation.
	//
	// If unset or empty, all solutions are applied.
	Solutions []string `json:"solutions,omitempty"`

	// Hosts on which solutions that are specified by this structure need
	// to be applied.
	//
	// If unset or empty, the solutions are applied on all of the hosts in the cluster.
	Hosts []string `json:"hosts,omitempty"`
}

// ApplySpec contains fields that describe a specification to be used for applying the desired solution
// specification to a given cluster.
type ApplySpec struct {

	// HostSolutions is the apply filter for solutions with deployment type
	// DeploymentType.EVERY_HOST_PINNED.
	//
	// If unset or empty and ClusterSolutions is unset or empty,
	// all solutions are applied on the cluster.
	HostSolutions *HostSolutionsApplyFilterSpec `json:"host_solutions,omitempty"`

	// ClusterSolutions is the apply filter for solutions with deployment type
	// DeploymentType.CLUSTER_VM_SET.
	//
	// If unset or empty and HostSolutions is unset or empty, all solutions are applied on
	// the cluster.
	ClusterSolutions *ClusterSolutionsApplyFilterSpec `json:"cluster_solutions,omitempty"`
}

const (
	// Success indicates the apply operation completed successfully.
	Success Status = "SUCCESS"

	// Error indicates the apply operation encountered an error.
	Error Status = "ERROR"
)

// ApplyStatus contains fields that describe the status of an apply operation.
type ApplyStatus struct {
	Status Status `json:"status"`

	// StartTime is the vLCM system time when the operation started.
	StartTime time.Time `json:"start_time"`

	// EndTime is the vLCM system time when the operation completed.
	EndTime time.Time `json:"end_time"`
}

// HostApplyStatus contains fields that describe the apply status for a specific host.
type HostApplyStatus struct {

	// Status is the aggregated apply status of the solutions on the host.
	//
	// Unset if the apply operation is not completed for the specified host.
	Status ApplyStatus `json:"status"`

	// SolutionStatus is the apply status of the different solutions on the host.
	SolutionStatus map[string]ApplyStatus `json:"solution_statuses"`
}

// HostSolutionsApplyStatus contains fields that describe the apply status of solutions with deployment
// type DeploymentType.EVERY_HOST_PINNED.
type HostSolutionsApplyStatus struct {

	// Status is the aggregated apply status of the solutions.
	//
	// Unset if the apply operation is not completed for solutions with deployment type
	// DeploymentType.EVERY_HOST_PINNED.
	Status ApplyStatus `json:"status"`

	// HostStatuses is the apply status of the hosts that were part of the apply operation.
	HostStatuses map[string]HostApplyStatus `json:"hostStatuses"`
}

// ClusterSolutionApplyStatus contains fields that describe the apply status for a specific solution.
type ClusterSolutionApplyStatus struct {

	// Status is the aggregated apply status for the deployment units of the solution.
	//
	// Unset if the apply operation is not completed for the specified deployment unit.
	Status ApplyStatus `json:"status"`

	// DeploymentUnitStatuses is the apply status for the different deployment units of the solution.
	DeploymentUnitStatuses map[string]ApplyStatus `json:"deployment_unit_statuses"`
}

// ClusterSolutionsApplyStatus contains fields that describe the apply status of solutions with deployment
// type DeploymentType.CLUSTER_VM_SET.
type ClusterSolutionsApplyStatus struct {

	// Status is the aggregated apply status of the solutions.
	//
	// Unset if the apply operation is not completed for solutions with deployment type
	// DeploymentType.CLUSTER_VM_SET.
	Status ApplyStatus `json:"status"`

	// SolutionStatuses is the apply status of the solutions that were part of the apply operation.
	SolutionStatuses map[string]ClusterSolutionApplyStatus `json:"solution_statuses"`
}

// ApplyResult contains fields that describe the result of an apply operation.
type ApplyResult struct {

	// ApplyStatus is the aggregated status of an apply operation.
	//
	// Unset if the apply operation is in progress.
	ApplyStatus ApplyStatus `json:"status"`

	// HostSolutionStatus is the apply status of all solutions with deployment type
	// DeploymentType.EVERY_HOST_PINNED that were part of the apply operation.
	HostSolutionStatus HostSolutionsApplyStatus `json:"host_solutions_status"`

	// ClusterSolutionStatus is the apply status of all solutions with deployment type
	// DeploymentType.CLUSTER_VM_SET that were part of the apply operation.
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
