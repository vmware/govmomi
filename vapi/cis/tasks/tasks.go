// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tasks

import (
	"context"
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

func (c *Manager) WaitForCompletion(ctx context.Context, taskId string) (string, error) {
	ticker := time.NewTicker(time.Second * time.Duration(c.pollingInterval))

	for {
		select {
		case <-ctx.Done():
		case <-ticker.C:
			taskInfo, err := c.getTaskInfo(taskId)
			status := taskInfo["status"].(string)
			if err != nil {
				return status, err
			}

			if status != "RUNNING" {
				return status, nil
			}
		}
	}
}

func (c *Manager) getTaskInfo(taskId string) (map[string]interface{}, error) {
	path := c.Resource(TasksPath).WithSubpath(taskId)
	req := path.Request(http.MethodGet)
	var res map[string]interface{}
	return res, c.Do(context.Background(), req, &res)
}
