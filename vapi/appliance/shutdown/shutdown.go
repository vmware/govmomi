/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0.
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shutdown

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	Path     = "/api/appliance/shutdown"
	Action   = "action"
	Cancel   = "cancel"
	PowerOff = "poweroff"
	Reboot   = "reboot"
)

// Manager provides convenience methods for shutdown Action
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Cancel cancels pending shutdown action.
func (m *Manager) Cancel(ctx context.Context) error {
	r := m.Resource(Path).WithParam(Action, Cancel)

	return m.Do(ctx, r.Request(http.MethodPost), nil)
}

// Spec represents request body class for an operation.
type Spec struct {
	Delay  int    `json:"delay"`
	Reason string `json:"reason"`
}

// Config defines shutdown configuration returned by the Shutdown.get operation
type Config struct {
	Action       string `json:"action"`
	Reason       string `json:"reason"`
	ShutdownTime string `json:"shutdown_time"`
}

// Get returns details about the pending shutdown action.
func (m *Manager) Get(ctx context.Context) (Config, error) {
	r := m.Resource(Path)

	var c Config

	return c, m.Do(ctx, r.Request(http.MethodGet), &c)
}

// PowerOff powers off the appliance.
func (m *Manager) PowerOff(ctx context.Context, reason string, delay int) error {
	r := m.Resource(Path).WithParam(Action, PowerOff)

	return m.Do(ctx, r.Request(http.MethodPost, Spec{
		Delay:  delay,
		Reason: reason,
	}), nil)
}

// Reboot reboots the appliance
func (m *Manager) Reboot(ctx context.Context, reason string, delay int) error {
	r := m.Resource(Path).WithParam(Action, Reboot)

	return m.Do(ctx, r.Request(http.MethodPost, Spec{
		Delay:  delay,
		Reason: reason,
	}), nil)
}
