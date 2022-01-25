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
	ApplianceShutDownPath = "/appliance/shutdown"
	Action                = "action"
	Cancel                = "cancel"
	PowerOff              = "poweroff"
	Reboot                = "reboot"
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

// Cancel cancels pending shutdown Action
func (m *Manager) Cancel(ctx context.Context) error {
	r := m.Resource(ApplianceShutDownPath).WithParam(Action, Cancel)

	return m.Do(ctx, r.Request(http.MethodPost), nil)
}

type Spec struct {
	Delay  int    `json:"delay"`
	Reason string `json:"reason"`
}

// Status represents the info about the pending shutdown Action
type Status struct {
	Reason       string `json:"reason,omitempty"`
	Action       string `json:"action,omitempty"`
	ShutDownTime string `json:"shutdown_time,omitempty"`
}

// Status gets details about the pending shutdown Action
func (m *Manager) Status(ctx context.Context) (Status, error) {
	r := m.Resource(ApplianceShutDownPath)

	var status Status
	err := m.Do(ctx, r.Request(http.MethodGet), &status)
	if err != nil {
		return Status{}, err
	}

	return status, nil
}

// PowerOff powers off the appliance
func (m *Manager) PowerOff(ctx context.Context, powerOffReason string, delay int) error {
	r := m.Resource(ApplianceShutDownPath).WithParam(Action, PowerOff)
	s := Spec{
		Delay:  delay,
		Reason: powerOffReason,
	}

	return m.Do(ctx, r.Request(http.MethodPost, s), nil)
}

// Reboot reboots the appliance
func (m *Manager) Reboot(ctx context.Context, rebootReason string, delay int) error {
	r := m.Resource(ApplianceShutDownPath).WithParam(Action, Reboot)
	s := Spec{
		Delay:  delay,
		Reason: rebootReason,
	}

	return m.Do(ctx, r.Request(http.MethodPost, s), nil)
}
