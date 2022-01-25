/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package staged

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/appliance/update/staged"

// Manager provides convenience methods to set/get status of the staged update.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Description represents description of the update.
type Description struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

//Info contains information about the staged update
type Info struct {
	Description     Description `json:"description"`
	Name            string      `json:"name"`
	Priority        Priority    `json:"priority"`
	RebootRequired  bool        `json:"reboot_required"`
	ReleaseDate     string      `json:"release_date"`
	Severity        Severity    `json:"severity"`
	Size            int         `json:"size"`
	StagingComplete bool        `json:"staging_complete"`
	UpdateType      UpdateType  `json:"update_type"`
	Version         string      `json:"version"`
}

//Get Gets the automatic update checking and staging policy.
func (m *Manager) Get(ctx context.Context) (Info, error) {
	r := m.Resource(Path)
	var info Info

	return info, m.Do(ctx, r.Request(http.MethodGet), &info)
}

// Delete Deletes the staged update.
func (m *Manager) Delete(ctx context.Context) error {
	r := m.Resource(Path)

	return m.Do(ctx, r.Request(http.MethodDelete), nil)
}
