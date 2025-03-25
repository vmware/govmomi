// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"net/http"
	"time"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// Session is used to create an initial update or download session
type Session struct {
	ClientProgress            int64                    `json:"client_progress,omitempty"`
	ErrorMessage              *rest.LocalizableMessage `json:"error_message,omitempty"`
	ExpirationTime            *time.Time               `json:"expiration_time,omitempty"`
	ID                        string                   `json:"id,omitempty"`
	LibraryItemContentVersion string                   `json:"library_item_content_version,omitempty"`
	LibraryItemID             string                   `json:"library_item_id,omitempty"`
	State                     string                   `json:"state,omitempty"`
}

// CreateLibraryItemUpdateSession creates a new library item
func (c *Manager) CreateLibraryItemUpdateSession(ctx context.Context, session Session) (string, error) {
	url := c.Resource(internal.LibraryItemUpdateSession)
	spec := struct {
		CreateSpec Session `json:"create_spec"`
	}{session}
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// GetLibraryItemUpdateSession gets the update session information with status
func (c *Manager) GetLibraryItemUpdateSession(ctx context.Context, id string) (*Session, error) {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id)
	var res Session
	return &res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ListLibraryItemUpdateSession gets the list of update sessions
func (c *Manager) ListLibraryItemUpdateSession(ctx context.Context) ([]string, error) {
	url := c.Resource(internal.LibraryItemUpdateSession)
	var res []string
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// CancelLibraryItemUpdateSession cancels an update session
func (c *Manager) CancelLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id).WithAction("cancel")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// CompleteLibraryItemUpdateSession completes an update session
func (c *Manager) CompleteLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id).WithAction("complete")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// DeleteLibraryItemUpdateSession deletes an update session
func (c *Manager) DeleteLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// FailLibraryItemUpdateSession fails an update session
func (c *Manager) FailLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id).WithAction("fail")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// KeepAliveLibraryItemUpdateSession keeps an inactive update session alive.
func (c *Manager) KeepAliveLibraryItemUpdateSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemUpdateSession).WithID(id).WithAction("keep-alive")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// WaitOnLibraryItemUpdateSession blocks until the update session is no longer
// in the ACTIVE state.
func (c *Manager) WaitOnLibraryItemUpdateSession(
	ctx context.Context, sessionID string,
	interval time.Duration, intervalCallback func()) error {

	// Wait until the upload operation is complete to return.
	for {
		session, err := c.GetLibraryItemUpdateSession(ctx, sessionID)
		if err != nil {
			return err
		}

		if session.State != "ACTIVE" {
			if session.State == "ERROR" {
				return session.ErrorMessage
			}
			return nil
		}
		time.Sleep(interval)
		if intervalCallback != nil {
			intervalCallback()
		}
	}
}

// CreateLibraryItemDownloadSession creates a new library item
func (c *Manager) CreateLibraryItemDownloadSession(ctx context.Context, session Session) (string, error) {
	url := c.Resource(internal.LibraryItemDownloadSession)
	spec := struct {
		CreateSpec Session `json:"create_spec"`
	}{session}
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// GetLibraryItemDownloadSession gets the download session information with status
func (c *Manager) GetLibraryItemDownloadSession(ctx context.Context, id string) (*Session, error) {
	url := c.Resource(internal.LibraryItemDownloadSession).WithID(id)
	var res Session
	return &res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ListLibraryItemDownloadSession gets the list of download sessions
func (c *Manager) ListLibraryItemDownloadSession(ctx context.Context) ([]string, error) {
	url := c.Resource(internal.LibraryItemDownloadSession)
	var res []string
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// CancelLibraryItemDownloadSession cancels an download session
func (c *Manager) CancelLibraryItemDownloadSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemDownloadSession).WithID(id).WithAction("cancel")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// DeleteLibraryItemDownloadSession deletes an download session
func (c *Manager) DeleteLibraryItemDownloadSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemDownloadSession).WithID(id)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// FailLibraryItemDownloadSession fails an download session
func (c *Manager) FailLibraryItemDownloadSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemDownloadSession).WithID(id).WithAction("fail")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}

// KeepAliveLibraryItemDownloadSession keeps an inactive download session alive.
func (c *Manager) KeepAliveLibraryItemDownloadSession(ctx context.Context, id string) error {
	url := c.Resource(internal.LibraryItemDownloadSession).WithID(id).WithAction("keep-alive")
	return c.Do(ctx, url.Request(http.MethodPost), nil)
}
