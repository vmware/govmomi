/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package library

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// StorageBackings for Content Libraries
type StorageBackings struct {
	DatastoreID string `json:"datastore_id,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Library  provides methods to create, read, update, delete, and enumerate libraries.
type Library struct {
	ID          string            `json:"id,omitempty"`
	Description string            `json:"description,omitempty"`
	Name        string            `json:"name,omitempty"`
	Version     string            `json:"version,omitempty"`
	Storage     []StorageBackings `json:"storage_backings,omitempty"`
}

// Manager extends rest.Client, adding content library related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// CreateLibrary creates a new library with the given Name, Description and CategoryID.
func (c *Manager) CreateLibrary(ctx context.Context, library *Library) (string, error) {

	type create struct {
		Name        string            `json:"name"`
		Description string            `json:"description"`
		Type        string            `json:"type"`
		Storage     []StorageBackings `json:"storage_backings,omitempty"`
	}
	spec := struct {
		Library create `json:"create_spec"`
	}{
		Library: create{
			Name:        library.Name,
			Description: library.Description,
			Type:        "LOCAL",
			Storage: []StorageBackings{
				StorageBackings{
					DatastoreID: "datastore-11",
					Type:        "DATASTORE",
				},
			},
		},
	}

	url := internal.URL(c, internal.LocalLibraryPath)
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// DeleteLibrary deletes an existing library.
func (c *Manager) DeleteLibrary(ctx context.Context, library *Library) error {
	url := internal.URL(c, internal.LocalLibraryPath).WithID(library.ID)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// ListLibraries returns a list of all content library IDs in the system.
func (c *Manager) ListLibraries(ctx context.Context) ([]string, error) {
	url := internal.URL(c, internal.LibraryPath)
	var res []string
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetLibraryByID returns information on a library for the given ID.
func (c *Manager) GetLibraryByID(ctx context.Context, id string) (*Library, error) {
	url := internal.URL(c, internal.LibraryPath).WithID(id)
	var res Library
	return &res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetLibraryByName returns information on a library for the given name.
func (c *Manager) GetLibraryByName(ctx context.Context, name string) (*Library, error) {
	// Lookup by name
	libraries, err := c.GetLibraries(ctx)
	if err != nil {
		return nil, err
	}

	for i := range libraries {
		if libraries[i].Name == name {
			return &libraries[i], nil
		}
	}

	return nil, fmt.Errorf("library name (%s) not found", name)
}

// GetLibraries returns a list of all content library details in the system.
func (c *Manager) GetLibraries(ctx context.Context) ([]Library, error) {
	ids, err := c.ListLibraries(ctx)
	if err != nil {
		return nil, fmt.Errorf("get libraries failed for: %s", err)
	}

	var libraries []Library
	for _, id := range ids {
		library, err := c.GetLibraryByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("get library %s failed for %s", id, err)
		}

		libraries = append(libraries, *library)

	}
	return libraries, nil
}
