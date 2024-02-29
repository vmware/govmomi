/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
