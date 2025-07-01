// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
