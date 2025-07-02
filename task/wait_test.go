// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"testing"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// ensure task.Manager implements the mo.Reference interface
var _ mo.Reference = new(Manager)

func TestCallbackFn(t *testing.T) {
	cb := &taskCallback{}

	for _, o := range []types.PropertyChangeOp{types.PropertyChangeOpAdd, types.PropertyChangeOpRemove, types.PropertyChangeOpAssign, types.PropertyChangeOpIndirectRemove} {
		for _, s := range []types.TaskInfoState{types.TaskInfoStateQueued, types.TaskInfoStateRunning, types.TaskInfoStateSuccess, types.TaskInfoStateError} {
			c := types.PropertyChange{
				Name: "info",
				Op:   o,
				Val: types.TaskInfo{
					State: s,
				},
			}
			t.Logf("Op: %s State: %s", o, s)
			cb.fn([]types.PropertyChange{c})
		}
	}
}
