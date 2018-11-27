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
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
)

// UpdateSession is used to create an initial update session
type UpdateSession struct {
	CreateSpec struct {
		ID                        string `json:"id,omitempty"`
		LibraryItemID             string `json:"library_item_id,omitempty"`
		LibraryItemContentVersion string `json:"library_item_content_version,omitempty"`
		// ErrorMessage              struct {
		//	ID             string   `json:"id,omitempty"`
		//	DefaultMessage string   `json:"default_message,omitempty"`
		//	Args           []string `json:"args,omitempty"`
		// } `json:"error_message,omitempty"`
		ClientProgress int64  `json:"client_progress,omitempty"`
		State          string `json:"state,omitempty"`
		// ExpirationTime time.Time `json:"expiration_time,omitempty"`
	} `json:"create_spec"`
}

// CreateLibraryItemUpdateSession creates a new library item
func (c *Manager) CreateLibraryItemUpdateSession(ctx context.Context, session UpdateSession) (string, error) {
	url := internal.URL(c, internal.LibraryItemUpdateSession)
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, session), &res)
}

// GetLibraryItemUpdateSession gets the update session information with status
func (c *Manager) GetLibraryItemUpdateSession(ctx context.Context, id string) (*UpdateSession, error) {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id)
	var res UpdateSession
	return &res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ListLibraryItemUpdateSession gets the list of update sessions
func (c *Manager) ListLibraryItemUpdateSession(ctx context.Context) (*[]string, error) {
	url := internal.URL(c, internal.LibraryItemUpdateSession)
	var res []string
	return &res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// CancelLibraryItemUpdateSession cancels an update session
func (c *Manager) CancelLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id).WithAction("cancel")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// CompleteLibraryItemUpdateSession completes an update session
func (c *Manager) CompleteLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id).WithAction("complete")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// DeleteLibraryItemUpdateSession completes an update session
func (c *Manager) DeleteLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// FailLibraryItemUpdateSession completes an update session
func (c *Manager) FailLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id).WithAction("fail")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// KeepAliveLibraryItemUpdateSession completes an update session
func (c *Manager) KeepAliveLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := internal.URL(c, internal.LibraryItemUpdateSession).WithID(id).WithAction("keep-alive")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// UpdateFile is used to add a file using an update session
type UpdateFile struct {
	FileSpec struct {
		Name       string `json:"name,omitempty"`
		SourceType string `json:"source_type,omitempty"`
		// SourceEndpoint struct {
		//	URI                      string `json:"uri,omitempty"`
		//	SSLCertificateThumbprint string `json:"ssl_certificate_thumbprint,omitempty"`
		// } `json:"source_endpoint,omitempty"`
		Size         int64 `json:"size,omitempty"`
		ChecksumInfo struct {
			Algorithm string `json:"algorithm,omitempty"`
			Checksum  string `json:"checksum,omitempty"`
		} `json:"checksum_info,omitempty"`
	} `json:"file_spec"`
}

// UpdateFileInfo is returned from adding a file when using an update session
type UpdateFileInfo struct {
	Name             string `json:"name,omitempty"`
	SourceType       string `json:"source_type,omitempty"`
	Status           string `json:"status,omitempty"`
	BytesTransferred int64  `json:"bytes_transferred,omitempty"`
	Size             int64  `json:"size,omitempty"`
	ChecksumInfo     struct {
		Algorithm string `json:"algorithm,omitempty"`
		Checksum  string `json:"checksum,omitempty"`
	} `json:"checksum_info,omitempty"`
	SourceEndpoint struct {
		URI                      string `json:"uri,omitempty"`
		SSLCertificateThumbprint string `json:"ssl_certificate_thumbprint,omitempty"`
	} `json:"source_endpoint,omitempty"`
	UploadEndpoint struct {
		URI                      string `json:"uri,omitempty"`
		SSLCertificateThumbprint string `json:"ssl_certificate_thumbprint,omitempty"`
	} `json:"upload_endpoint,omitempty"`
}

// AddLibraryItemFile adds a file
func (c *Manager) AddLibraryItemFile(ctx context.Context, sessionID string, updateFile UpdateFile) (*UpdateFileInfo, error) {
	url := internal.URL(c, internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("add")
	var res UpdateFileInfo
	return &res, c.Do(ctx, url.Request(http.MethodPost, updateFile), &res)
}

// GetLibraryItemFile retrieves information about a specific file
func (c *Manager) GetLibraryItemFile(ctx context.Context, sessionID string, filename string) (*UpdateFileInfo, error) {
	type FileName struct {
		FileName string `json:"file_name,omitempty"`
	}
	url := internal.URL(c, internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("get")
	var res UpdateFileInfo
	var fileName = FileName{FileName: filename}
	return &res, c.Do(ctx, url.Request(http.MethodPost, fileName), &res)
}
