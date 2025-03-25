// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
)

// Checksum provides checksum information on library item files.
type Checksum struct {
	Algorithm string `json:"algorithm,omitempty"`
	Checksum  string `json:"checksum"`
}

// File provides methods to get information on library item files.
type File struct {
	Cached           *bool     `json:"cached,omitempty"`
	Checksum         *Checksum `json:"checksum_info,omitempty"`
	Name             string    `json:"name,omitempty"`
	Size             *int64    `json:"size,omitempty"`
	Version          string    `json:"version,omitempty"`
	DownloadEndpoint string    `json:"file_download_endpoint,omitempty"`
}

// ListLibraryItemFiles returns a list of all the files for a library item.
func (c *Manager) ListLibraryItemFiles(ctx context.Context, id string) ([]File, error) {
	url := c.Resource(internal.LibraryItemFilePath).WithParam("library_item_id", id)
	var res []File
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetLibraryItemFile returns a file with the provided name for a library item.
func (c *Manager) GetLibraryItemFile(ctx context.Context, id, fileName string) (*File, error) {
	url := c.Resource(internal.LibraryItemFilePath).WithID(id).WithAction("get")
	spec := struct {
		Name string `json:"name"`
	}{fileName}
	var res File
	return &res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}
