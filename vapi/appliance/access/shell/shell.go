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
