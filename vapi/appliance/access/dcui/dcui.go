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
