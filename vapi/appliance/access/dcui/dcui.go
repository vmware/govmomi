// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package dcui

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/api/appliance/access/dcui"

// Manager provides convenience methods to get/set enabled state of DCUI.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Get returns enabled state of Direct Console User Interface (DCUI TTY2).
func (m *Manager) Get(ctx context.Context) (bool, error) {
	r := m.Resource(Path)

	var state bool
	err := m.Do(ctx, r.Request(http.MethodGet), &state)

	return state, err
}

// Access represents the value to be set for DCUI
type Access struct {
	Enabled bool `json:"enabled"`
}

// Set enables state of Direct Console User Interface (DCUI TTY2).
func (m *Manager) Set(ctx context.Context, inp Access) error {
	r := m.Resource(Path)

	return m.Do(ctx, r.Request(http.MethodPut, inp), nil)
}
