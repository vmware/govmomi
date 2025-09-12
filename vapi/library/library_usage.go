// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"net/http"
	"time"

	"github.com/vmware/govmomi/vapi/internal"
)

// Usage provides methods to get usage information on content library.
type Usage struct {
	ID           string     `json:"usage,omitempty"`
	ResourceUrn  string     `json:"resource_urn,omitempty"`
	AdditionTime *time.Time `json:"addition_time,omitempty"`
}

// UsageSummary contains commonly used information about the usage of a content library.
type UsageSummary struct {
	ID          string `json:"usage,omitempty"`
	ResourceUrn string `json:"resource_urn,omitempty"`
}

// UsageList lists the usage by resource(s) on a content library.
type UsageList struct {
	LibraryUsageList []UsageSummary `json:"library_usage_list,omitempty"`
}

// GetLibraryUsage retrieves the library usage information for a given usage identifier.
func (c *Manager) GetLibraryUsage(ctx context.Context, libraryID, usageID string) (Usage, error) {
	url := c.Resource(internal.APILibraryPath).
		WithSubpath(libraryID).
		WithSubpath(internal.LibraryUsages).
		WithSubpath(usageID)
	var res Usage
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ListLibraryUsage retrieves the list of resources currently using a content library.
// A content library can be safely deleted if no usage is present for that library.
func (c *Manager) ListLibraryUsage(ctx context.Context, libraryID string) (UsageList, error) {
	url := c.Resource(internal.APILibraryPath).WithSubpath(libraryID).WithSubpath(internal.LibraryUsages)
	var res UsageList
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// RemoveLibraryUsage removes a resource usage on a content library.
func (c *Manager) RemoveLibraryUsage(ctx context.Context, libraryID string, usageID string) error {
	url := c.Resource(internal.APILibraryPath).
		WithSubpath(libraryID).
		WithSubpath(internal.LibraryUsages).
		WithSubpath(usageID)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// AddUsage defines the information required to add a resource usage on a content library.
type AddUsage struct {
	ResourceUrn string `json:"resource_urn,omitempty"`
}

// AddLibraryUsage adds a resource usage on a content library.
func (c *Manager) AddLibraryUsage(ctx context.Context, libraryID string, addUsage AddUsage) (string, error) {
	url := c.Resource(internal.APILibraryPath).
		WithSubpath(libraryID).
		WithSubpath(internal.LibraryUsages)

	var res struct {
		Value string `json:"value,omitempty"`
	}
	return res.Value, c.Do(ctx, url.Request(http.MethodPost, addUsage), &res)
}
