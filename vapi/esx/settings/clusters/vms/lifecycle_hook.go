// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

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
	Vm string

	// LifecycleState is the state of the VM specified by Vm.
	LifecycleState LifecycleState

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
	DynamicUpdateProcessed bool
}

// LifecycleHookConfig contains fields that describe a VM lifecycle hook configuration.
type LifecycleHookConfig struct {

	// Timeout is the maximum time in seconds for vLCM to wait for a hook to
	// be processed by the solution. An issue is raised if the time elapsed and
	// the hook is still not processed. See {@link Solutions.IssueInfo}. The
	// issue is attached to the {@link Solutions.DeploymentInfo} structure that
	// holds the VM for which the hook was activated.
	//
	// If unset, defaults to 10 hours.
	Timeout int `json:"timeout"`
}
