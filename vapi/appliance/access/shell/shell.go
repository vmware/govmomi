// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package shell

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/api/appliance/access/shell"

// Manager provides convenience methods to get/set enabled state of BASH.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Get returns enabled state of BASH, that is, access to BASH from within the controlled CLI.
func (m *Manager) Get(ctx context.Context) (Access, error) {
	r := m.Resource(Path)

	var status Access
	return status, m.Do(ctx, r.Request(http.MethodGet), &status)
}

// Access represents shell configuration.
type Access struct {
	Enabled bool `json:"enabled"`
	Timeout int  `json:"timeout"`
}

// Set enables state of BASH, that is access to BASH from within the controlled CLI.
func (m *Manager) Set(ctx context.Context, inp Access) error {
	r := m.Resource(Path)

	return m.Do(ctx, r.Request(http.MethodPut, inp), nil)
}
