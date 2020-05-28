/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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
