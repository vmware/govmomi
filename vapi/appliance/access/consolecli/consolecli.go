// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package consolecli

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/api/appliance/access/consolecli"

// Manager provides convenience methods to get/set enabled state of CLI.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Get returns enabled state of the console-based controlled CLI (TTY1).
func (m *Manager) Get(ctx context.Context) (bool, error) {
	r := m.Resource(Path)

	var status bool
	return status, m.Do(ctx, r.Request(http.MethodGet), &status)
}

// Access represents the value to be set for ConsoleCLI
type Access struct {
	Enabled bool `json:"enabled"`
}

// Set enables state of the console-based controlled CLI (TTY1).
func (m *Manager) Set(ctx context.Context, inp Access) error {
	r := m.Resource(Path)

	return m.Do(ctx, r.Request(http.MethodPut, inp), nil)
}
