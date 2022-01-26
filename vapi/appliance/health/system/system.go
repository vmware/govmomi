/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package system

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	HealthPath          = "/appliance/health/system"
	HealthLastCheckPath = "/appliance/health/system/lastcheck"
)

// Manager provides convenience methods to get overall health of the system.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Get gets overall health of system.
func (m *Manager) Get(ctx context.Context) (string, error) {
	r := m.Resource(HealthPath)

	var status string
	return status, m.Do(ctx, r.Request(http.MethodGet), &status)
}

// LastCheck gets last check timestamp of the health of the system.
func (m *Manager) LastCheck(ctx context.Context) (string, error) {
	r := m.Resource(HealthLastCheckPath)

	var status string
	return status, m.Do(ctx, r.Request(http.MethodGet), &status)
}
