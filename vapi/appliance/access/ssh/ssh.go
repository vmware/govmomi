// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ssh

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/api/appliance/access/ssh"

// Manager provides convenience methods to get/set enabled state of SSH-based controlled CLI.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Get returns enabled state of the SSH-based controlled CLI.
func (m *Manager) Get(ctx context.Context) (bool, error) {
	r := m.Resource(Path)

	var state bool
	err := m.Do(ctx, r.Request(http.MethodGet), &state)

	return state, err
}

// Access represents the value to be set for SSH
type Access struct {
	Enabled bool `json:"enabled"`
}

// Set enables state of the SSH-based controlled CLI.
func (m *Manager) Set(ctx context.Context, inp Access) error {
	r := m.Resource(Path)

	return m.Do(ctx, r.Request(http.MethodPut, inp), nil)
}
