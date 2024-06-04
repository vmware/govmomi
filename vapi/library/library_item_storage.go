/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
)

// Storage is an expanded form of library.File that includes details about the
// storage backing for a file in a library item
type Storage struct {
	Checksum       Checksum       `json:"checksum_info,omitempty"`
	StorageBacking StorageBacking `json:"storage_backing"`
	StorageURIs    []string       `json:"storage_uris"`
	Name           string         `json:"name"`
	Size           int64          `json:"size"`
	Cached         bool           `json:"cached"`
	Version        string         `json:"version"`
}

// ListLibraryItemStorage returns a list of all the storage for a library item.
func (c *Manager) ListLibraryItemStorage(ctx context.Context, id string) ([]Storage, error) {
	url := c.Resource(internal.LibraryItemStoragePath).WithParam("library_item_id", id)
	var res []Storage
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetLibraryItemStorage returns the storage for a specific file in a library item.
func (c *Manager) GetLibraryItemStorage(ctx context.Context, id, fileName string) ([]Storage, error) {
	url := c.Resource(internal.LibraryItemStoragePath).WithID(id).WithAction("get")
	spec := struct {
		Name string `json:"file_name"`
	}{fileName}
	var res []Storage
	return res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}
