// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// DownloadFile is the specification for the downloadsession
// operations file:add, file:get, and file:list.
type DownloadFile struct {
	BytesTransferred int64                    `json:"bytes_transferred"`
	Checksum         *Checksum                `json:"checksum_info,omitempty"`
	DownloadEndpoint *TransferEndpoint        `json:"download_endpoint,omitempty"`
	ErrorMessage     *rest.LocalizableMessage `json:"error_message,omitempty"`
	Name             string                   `json:"name"`
	Size             int64                    `json:"size,omitempty"`
	Status           string                   `json:"status"`
}

// GetLibraryItemDownloadSessionFile retrieves information about a specific file that is a part of an download session.
func (c *Manager) GetLibraryItemDownloadSessionFile(ctx context.Context, sessionID string, name string) (*DownloadFile, error) {
	url := c.Resource(internal.LibraryItemDownloadSessionFile).WithID(sessionID).WithAction("get")
	spec := struct {
		Name string `json:"file_name"`
	}{name}
	var res DownloadFile
	err := c.Do(ctx, url.Request(http.MethodPost, spec), &res)
	if err != nil {
		return nil, err
	}
	if res.Status == "ERROR" {
		return nil, res.ErrorMessage
	}
	return &res, nil
}

// ListLibraryItemDownloadSessionFile retrieves information about a specific file that is a part of an download session.
func (c *Manager) ListLibraryItemDownloadSessionFile(ctx context.Context, sessionID string) ([]DownloadFile, error) {
	url := c.Resource(internal.LibraryItemDownloadSessionFile).WithParam("download_session_id", sessionID)
	var res []DownloadFile
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// PrepareLibraryItemDownloadSessionFile retrieves information about a specific file that is a part of an download session.
func (c *Manager) PrepareLibraryItemDownloadSessionFile(ctx context.Context, sessionID string, name string) (*DownloadFile, error) {
	url := c.Resource(internal.LibraryItemDownloadSessionFile).WithID(sessionID).WithAction("prepare")
	spec := struct {
		Name string `json:"file_name"`
	}{name}
	var res DownloadFile
	return &res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}
