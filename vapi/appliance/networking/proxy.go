// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package networking

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const applianceProxyConfigPath = "/appliance/networking/proxy"
const applianceNoProxyConfigPath = "/appliance/networking/noproxy"

// Manager provides convenience methods to configure appliance proxy.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Proxy represents configuration for specific proxy - ftp, http, https.
type Proxy struct {
	Server   string `json:"server,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

// ProxyList represents configuration for vcenter proxy.
type ProxyList struct {
	Ftp   Proxy `json:"ftp,omitempty"`
	Http  Proxy `json:"http,omitempty"`
	Https Proxy `json:"https,omitempty"`
}

// ProxyList returns all Proxy configuration.
func (m *Manager) ProxyList(ctx context.Context) (*ProxyList, error) {
	var res ProxyList
	var rawRes []struct {
		Key   string
		Value Proxy
	}

	r := m.Resource(applianceProxyConfigPath)
	err := m.Do(ctx, r.Request(http.MethodGet), &rawRes)
	if err != nil {
		return &res, err
	}

	for _, c := range rawRes {
		switch c.Key {
		case "http":
			res.Http = c.Value
		case "https":
			res.Https = c.Value
		case "ftp":
			res.Ftp = c.Value
		}
	}

	return &res, nil
}

// NoProxy returns all excluded servers for proxying.
func (m *Manager) NoProxy(ctx context.Context) ([]string, error) {
	r := m.Resource(applianceNoProxyConfigPath)
	var res []string
	return res, m.Do(ctx, r.Request(http.MethodGet), &res)
}
