// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fault_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/types"
)

func TestAlreadyPoweredOff(t *testing.T) {
	var (
		errFalse any
		errTrue  any
	)

	errFalse = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.InvalidPowerState{
				ExistingState:  types.VirtualMachinePowerStateSuspended,
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
			},
			LocalizedMessage: "vm must be powered on to power off",
		},
	}

	errTrue = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.InvalidPowerState{
				ExistingState:  types.VirtualMachinePowerStatePoweredOff,
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
			},
			LocalizedMessage: "vm must be powered on to power off",
		},
	}

	assert.False(t, fault.IsAlreadyPoweredOffError(errFalse))
	assert.True(t, fault.IsAlreadyPoweredOffError(errTrue))
}
