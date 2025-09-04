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

// The {@name Status} {@term enumerated type} defines the status values that can be reported for an operation.
type Status string

const (

	// The operation is in pending state.
	Pending Status = "PENDING"

	// The operation is in progress.
	Running Status = "RUNNING"

	// The operation is blocked.
	Blocked Status = "BLOCKED"

	// The operation completed successfully.
	Succeeded Status = "SUCCEEDED"

	// The operation failed.
	Failed Status = "FAILED"
)

// The {@name Progress} {@term structure} contains information describe the progress of an operation.
type Progress struct {
	/**
	 * Total amount of the work for the operation.
	 */
	Total uint64 `json:"total"`

	/**
	 * The amount of work completed for the operation. The value can only be
	 * incremented.
	 */
	Completed uint64 `json:"completed"`

	/**
	 * Message about the work progress.
	 */
	Message rest.LocalizableMessage `json:"message"`
}

type Info struct {
	/**
	 * Description of the operation associated with the task.
	 */
	Description rest.LocalizableMessage `json:"description"`

	/**
	 * Identifier of the service containing the operation.
	 */
	Service string `json:"service"`

	/**
	 * Identifier of the operation associated with the task.
	 */
	Operation string `json:"operation"`

	/**
	 * Parent of the current task.
	 *
	 * @field.optional This {@term field} will be {@term unset} if the
	 * task has no parent.
	 */
	Parent string `json:"parent"`

	/**
	 * Identifier of the target created by the operation or an existing one
	 * the operation performed on.
	 *
	 * @field.optional This {@term field} will be {@term unset} if the
	 * operation has no target or multiple targets.
	 */
	Target map[string]string `json:"target"`

	/**
	 * Status of the operation associated with the task.
	 */
	Status Status `json:"status"`

	/**
	 * Flag to indicate whether or not the operation can be cancelled.
	 * The value may change as the operation progresses.
	 */
	Cancelable bool `json:"cancelable"`

	/**
	 * Description of the error if the operation status is "FAILED".
	 *
	 * @field.optional If {@term unset} the description of why the operation
	 * failed will be included in the result of the operation
	 * (see {@link Info#result}).
	 */
	Error rest.Error `json:"error"`

	/**
	 * Time when the operation is started.
	 */
	Start time.Time `json:"start_time"`

	/**
	 * Time when the operation is completed.
	 */
	End time.Time `json:"end_time"`

	/**
	 * Name of the user who performed the operation.
	 *
	 * @field.optional This {@term field} will be {@term unset} if the
	 * operation is performed by the system.
	 */
	User string `json:"user"`

	/**
	 * Progress of the operation.
	 */
	Progress Progress `json:"progress"`

	/**
	 * Result of the operation.
	 * <p>
	 * If an operation reports partial results before it completes, this
	 * {@term field} could be {@term set} before the {@link CommonInfo#status}
	 * has the value {@link Status#SUCCEEDED}. The value could change as the
	 * operation progresses.
	 *
	 * @field.optional This {@term field} will be {@term unset} if the
	 * operation does not return a result or if the result is not available
	 * at the current step of the operation.
	 */
	Result json.RawMessage `json:"result"`
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
