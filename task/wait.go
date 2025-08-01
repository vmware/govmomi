// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"context"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/progress"
	"github.com/vmware/govmomi/vim25/types"
)

type taskProgress struct {
	info *types.TaskInfo
}

func (t taskProgress) Percentage() float32 {
	return float32(t.info.Progress)
}

func (t taskProgress) Detail() string {
	return ""
}

func (t taskProgress) Error() error {
	if t.info.Error != nil {
		return Error{t.info.Error, t.info.Description}
	}

	return nil
}

type taskCallback struct {
	ch   chan<- progress.Report
	info *types.TaskInfo
	err  error
}

func (t *taskCallback) fn(pc []types.PropertyChange) bool {
	for _, c := range pc {
		if c.Name != "info" {
			continue
		}

		if c.Op != types.PropertyChangeOpAssign {
			continue
		}

		if c.Val == nil {
			continue
		}

		ti := c.Val.(types.TaskInfo)
		t.info = &ti
	}

	// t.info could be nil if pc can't satisfy the rules above
	if t.info == nil {
		return false
	}

	pr := taskProgress{t.info}

	// Store copy of error, so Wait() can return it as well.
	t.err = pr.Error()

	switch t.info.State {
	case types.TaskInfoStateQueued, types.TaskInfoStateRunning:
		if t.ch != nil {
			// Don't care if this is dropped
			select {
			case t.ch <- pr:
			default:
			}
		}
		return false
	case types.TaskInfoStateSuccess, types.TaskInfoStateError:
		if t.ch != nil {
			// Last one must always be delivered
			t.ch <- pr
		}
		return true
	default:
		panic("unknown state: " + t.info.State)
	}
}

// WaitEx waits for a task to finish with either success or failure. It does so
// by waiting for the "info" property of task managed object to change. The
// function returns when it finds the task in the "success" or "error" state.
// In the former case, the return value is nil. In the latter case the return
// value is an instance of this package's Error struct.
//
// Any error returned while waiting for property changes causes the function to
// return immediately and propagate the error.
//
// If the progress.Sinker argument is specified, any progress updates for the
// task are sent here. The completion percentage is passed through directly.
// The detail for the progress update is set to an empty string. If the task
// finishes in the error state, the error instance is passed through as well.
// Note that this error is the same error that is returned by this function.
func WaitEx(
	ctx context.Context,
	ref types.ManagedObjectReference,
	pc *property.Collector,
	s progress.Sinker) (*types.TaskInfo, error) {

	cb := &taskCallback{}

	// Include progress sink if specified
	if s != nil {
		cb.ch = s.Sink()
		defer close(cb.ch)
	}

	filter := &property.WaitFilter{
		WaitOptions: property.WaitOptions{
			PropagateMissing: true,
		},
	}
	filter.Add(ref, ref.Type, []string{"info"})

	if err := property.WaitForUpdatesEx(
		ctx,
		pc,
		filter,
		func(updates []types.ObjectUpdate) bool {
			for _, update := range updates {
				// Only look at updates for the expected task object.
				if update.Obj.Value == ref.Value && update.Obj.Type == ref.Type {
					if cb.fn(update.ChangeSet) {
						return true
					}
				}
			}
			return false
		}); err != nil {

		return nil, err
	}

	return cb.info, cb.err
}
