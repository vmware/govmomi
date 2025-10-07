// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/types"
)

type clusterHooksPath types.ManagedObjectReference

const (
	HooksPath = clusters.BasePath + "/%s/vms/lifecycle-hooks"
)

func (c clusterHooksPath) String() string {
	cid := types.ManagedObjectReference(c).Value
	return fmt.Sprintf(HooksPath, cid)
}

// LifecycleState contains the different VM lifecycle states a solution can
// hook into. See LifecycleHooks and SolutionSpec.
type LifecycleState string

const (
	// PostProvisioning reached once immediately after a VM is created.
	PostProvisioning LifecycleState = "POST_PROVISIONING"

	// PostPowerOn is post VM power-on, reached immediately after every VM power-on.
	PostPowerOn LifecycleState = "POST_POWER_ON"
)

// LifecycleHookInfo contains fields that describe a VM lifecycle hook that is activated for a given VM.
type LifecycleHookInfo struct {

	// Vm is the identifier of the VM for which the hook is activated.
	Vm string `json:"vm"`

	// LifecycleState is the state of the VM specified by Vm.
	LifecycleState LifecycleState `json:"lifecycle_state"`

	// Configuration of the hook.
	// LifecycleHookConfig Configuration

	// HookActivated is the vLCM system time when the hook is activated.
	// DateTime HookActivated

	// DynamicUpdateProcessed represents if the DynamicUpdateSpec given with
	// LcycleHooks.ProcessDynamicUpdate is applied successfully for
	// the LifecycleState of the given member vm.
	//
	// Defaults to False.
	//
	// See LifecycleHooks#processDynamicUpdate about how to process the
	// dynamic update for a given LifecycleState.
	DynamicUpdateProcessed bool `json:"dynamic_update_processed"`
}

// LifecycleHookConfig contains fields that describe a VM lifecycle hook configuration.
type LifecycleHookConfig struct {

	// Timeout is the maximum time in seconds for vLCM to wait for a hook to
	// be processed by the solution. An issue is raised if the time elapsed and
	// the hook is still not processed. See Solutions.IssueInfo. The
	// issue is attached to the Solutions.DeploymentInfo structure that
	// holds the VM for which the hook was activated.
	//
	// If unset, defaults to 10 hours.
	Timeout *int `json:"timeout,omitempty"`
}

type HookListResult struct {
	Hooks []LifecycleHookInfo `json:"hooks"`
}

type ProcessedHookSpec struct {
	LifecycleState LifecycleState `json:"lifecycle_state"`

	ProcessedSuccessfully bool `json:"processed_successfully"`

	Vm string `json:"vm"`
}

type DynamicUpdateSpec struct {
	Vm string `json:"vm"`

	Solution string `json:"solution"`

	LifecycleState LifecycleState `json:"lifecycle_state"`

	AlternativeVmSpec *AlternativeVmSpec `json:"alternative_vm_spec,omitempty"`
}

func (m *Manager) ListHooks(ctx context.Context, cluster types.ManagedObjectReference, solution string) (*HookListResult, error) {
	p := clusterHooksPath(cluster).String()
	url := m.Resource(p).WithSubpath(solution)
	var res HookListResult

	if err := m.Do(ctx, url.Request(http.MethodGet), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *Manager) MarkAsProcessed(ctx context.Context, cluster types.ManagedObjectReference, spec *ProcessedHookSpec) (*HookListResult, error) {
	p := clusterHooksPath(cluster).String()
	url := m.Resource(p).WithParam("action", "mark-as-processed")

	var errMsg rest.Error
	if err := m.Do(ctx, url.Request(http.MethodPost, spec), &errMsg); err != nil {
		if len(errMsg.Messages) > 0 {
			return nil, &errMsg.Messages[0]
		}
		return nil, err
	}

	return nil, nil
}

func (m *Manager) ProcessDynamicUpdate(ctx context.Context, cluster types.ManagedObjectReference, spec *DynamicUpdateSpec) error {
	p := clusterHooksPath(cluster).String()
	url := m.Resource(p).WithParam("action", "process-dynamic-update")

	var errMsg rest.Error
	if err := m.Do(ctx, url.Request(http.MethodPost, spec), &errMsg); err != nil {
		if len(errMsg.Messages) > 0 {
			return &errMsg.Messages[0]
		}
		return err
	}

	return nil
}
