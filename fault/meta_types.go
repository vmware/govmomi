// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fault

import "github.com/vmware/govmomi/vim25/types"

// IsAlreadyPoweredOffError verifies that the error is an InvalidPowerState
// error and returns true if the existing state from the error is powered off
func IsAlreadyPoweredOffError(err any) bool {
	var existingState types.VirtualMachinePowerState

	In(err, func(
		fault types.BaseMethodFault,
		localizedMessage string,
		localizableMessages []types.LocalizableMessage) bool {
		if invalidPowerState, ok := fault.(*types.InvalidPowerState); ok {
			existingState = invalidPowerState.ExistingState

			return true
		}

		return false
	})

	return existingState == types.VirtualMachinePowerStatePoweredOff
}

// IsTransientError checks whether the error type indicates an error that is
// likely to resolve without explicit action in calling logic.
// Some of those are highly transient, such as TaskInProgress. Others are
// potentially longer term, such as HostCommunication; they are inherently
// transient but may not resolve in a short time frame.
func IsTransientError(err any) bool {
	return Is(err, &types.TaskInProgress{}) ||
		Is(err, &types.NetworkDisruptedAndConfigRolledBack{}) ||
		Is(err, &types.VAppTaskInProgress{}) ||
		Is(err, &types.FailToLockFaultToleranceVMs{}) ||
		Is(err, &types.HostCommunication{})
}
