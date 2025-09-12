// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	// TasksPath The endpoint for retrieving tasks
	TasksPath = "/api/cis/tasks"
	// The default interval in seconds for polling for task results
	defaultPollingInterval = 10
)

// Status defines the status values that can be reported for an operation.
type Status string

const (

	// Pending indicates the operation is in pending state.
	Pending Status = "PENDING"

	// Running indicates the operation is in progress.
	Running Status = "RUNNING"

	// Blocked indicates the operation is blocked.
	Blocked Status = "BLOCKED"

	// Succeeded indicates the operation completed successfully.
	Succeeded Status = "SUCCEEDED"

	// Failed indicates the operation failed.
	Failed Status = "FAILED"
)

// Progress contains information describing the progress of an operation.
type Progress struct {
	// Total is the total amount of work for the operation.
	Total uint64 `json:"total"`

	// Completed is the amount of work completed for the operation. The value can only be incremented.
	Completed uint64 `json:"completed"`

	// Message about the work progress.
	Message *rest.LocalizableMessage `json:"message,omitempty"`
}

// Info contains information about a task.
type Info struct {
	// Description of the operation associated with the task.
	Description rest.LocalizableMessage `json:"description"`

	// Service is the identifier of the service containing the operation.
	Service string `json:"service"`

	// Operation is the identifier of the operation associated with the task.
	Operation string `json:"operation"`

	// Parent of the current task.
	//
	// This field will be unset if the task has no parent.
	Parent *string `json:"parent,omitempty"`

	// Target is the identifier of the target created by the operation or an existing one
	// the operation performed on.
	//
	// This field will be unset if the operation has no target or multiple targets.
	Target map[string]string `json:"target,omitempty"`

	// Status of the operation associated with the task.
	Status Status `json:"status"`

	// Cancelable is a flag to indicate whether or not the operation can be cancelled.
	// The value may change as the operation progresses.
	Cancelable bool `json:"cancelable"`

	// Error is the description of the error if the operation status is "FAILED".
	//
	// If unset the description of why the operation failed will be included in the result of the operation
	// (see Info.Result).
	Error rest.Error `json:"error,omitempty"`

	// Start is the time when the operation is started.
	Start *time.Time `json:"start_time,omitempty"`

	// End is the time when the operation is completed.
	End *time.Time `json:"end_time,omitempty"`

	// User is the name of the user who performed the operation.
	//
	// This field will be unset if the operation is performed by the system.
	User *string `json:"user,omitempty"`

	// Progress of the operation.
	Progress Progress `json:"progress"`

	// Result of the operation.
	//
	// If an operation reports partial results before it completes, this
	// field could be set before the Status has the value SUCCEEDED. The value could change as the
	// operation progresses.
	//
	// This field will be unset if the operation does not return a result or if the result is not available
	// at the current step of the operation.
	Result json.RawMessage `json:"result,omitempty"`
}

func (t *Info) IsDone() bool {
	return t.Status != Pending && t.Status != Running
}

// Err returns an error if the task state is Failed.
func (t *Info) Err() error {
	if t.Status != Failed {
		return nil
	}
	if len(t.Error.Messages) > 0 {
		return &t.Error.Messages[0]
	}
	return errors.New(string(t.Error.ErrorType))
}

// Manager extends rest.Client, adding task related methods.
type Manager struct {
	*rest.Client
	pollingInterval int
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client:          client,
		pollingInterval: defaultPollingInterval,
	}
}

// NewManager creates a new Manager instance with the given client and custom polling interval
func NewManagerWithCustomInterval(client *rest.Client, pollingInterval int) *Manager {
	return &Manager{
		Client:          client,
		pollingInterval: pollingInterval,
	}
}

// WaitForCompletion will wait for the task status to be in anything but the
// Pending or Running state.
func (c *Manager) WaitForCompletion(ctx context.Context, taskId string) (*Info, error) {
	return c.waitForState(ctx, taskId, func(i *Info) bool { return i.IsDone() })
}

// WaitForRunningOrError will wait for the task status to be Running or Failed.
// If the task has failed, it will return the TaskInfo Error field as an error.
func (c *Manager) WaitForRunningOrError(ctx context.Context, taskId string) (*Info, error) {
	check := func(i *Info) bool {
		return i.Status == Running || i.Status == Failed
	}
	return c.waitForState(ctx, taskId, check)
}

func (c *Manager) waitForState(ctx context.Context, taskId string, check func(i *Info) bool) (*Info, error) {
	ticker := time.NewTicker(time.Second * time.Duration(c.pollingInterval))
	defer ticker.Stop()

	for {
		taskInfo, err := c.Get(ctx, taskId)
		if err != nil {
			return taskInfo, err
		}

		// Check for the state we care about.
		if check(taskInfo) {
			var err error
			if taskInfo.Err() != nil {
				err = fmt.Errorf("%s: %w", taskId, taskInfo.Err())
			}

			return taskInfo, err
		}

		// Try again.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

func (c *Manager) Get(ctx context.Context, taskId string) (*Info, error) {
	path := c.Resource(TasksPath).WithSubpath(taskId)
	req := path.Request(http.MethodGet)
	var res Info
	return &res, c.Do(ctx, req, &res)
}
