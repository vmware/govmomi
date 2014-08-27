/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import (
	"errors"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// Task adds functionality for the Task managed object.
type Task struct {
	c   *Client
	ref types.ManagedObjectReference
}

func NewTask(c *Client, ref types.ManagedObjectReference) *Task {
	t := Task{
		c:   c,
		ref: ref,
	}

	return &t
}

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
		return errors.New(t.info.Error.LocalizedMessage)
	}

	return nil
}

type taskCallback struct {
	ch   chan<- vim25.Progress
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

func (t *Task) Wait() error {
	_, err := t.WaitForResult(nil)
	return err
}

func (t *Task) WaitForResult(ch chan<- vim25.Progress) (*types.TaskInfo, error) {
	cb := &taskCallback{ch: ch}
	if ch != nil {
		defer close(ch)
	}

	err := t.c.WaitForProperties(t.ref, []string{"info"}, cb.fn)
	if err != nil {
		return nil, err
	}

	return cb.info, cb.err
}
