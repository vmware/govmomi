// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const applianceLoggingForwardingPath = "/appliance/logging/forwarding"

// Manager provides convenience methods to configure appliance logging forwarding.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Forwarding represents configuration for log message forwarding.
type Forwarding struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

func (m *Manager) getManagerResource() *rest.Resource {
	return m.Resource(applianceLoggingForwardingPath)
}

// Forwarding returns all logging forwarding config.
func (m *Manager) Forwarding(ctx context.Context) ([]Forwarding, error) {
	r := m.getManagerResource()
	var res []Forwarding
	return res, m.Do(ctx, r.Request(http.MethodGet), &res)
}
