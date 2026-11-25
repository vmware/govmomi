// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package tagging provides a client for the vSphere 9.0+ tagging APIs under
// "/api/vcenter/tagging". It supersedes the "/rest/com/vmware/cis/tagging"
// based package github.com/vmware/govmomi/vapi/tags.
package tagging

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// Manager extends rest.Client, adding tagging related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// DefaultListPageSize is the page size used by the "/api/vcenter/tagging"
// list operations, when the request does not specify one.
const DefaultListPageSize = 20

// TagListItem is an item returned by ListTags.
type TagListItem struct {
	Tag  string  `json:"tag"`
	Info TagInfo `json:"info"`
}

// TagInfo is the information about a tag in a TagListItem.
type TagInfo struct {
	Name        string   `json:"name,omitempty"`
	Category    string   `json:"category,omitempty"`
	Description string   `json:"description,omitempty"`
	UsedBy      []string `json:"used_by,omitempty"`
}

// TagListResult is a single page of the tags list response.
type TagListResult struct {
	Marker string        `json:"marker,omitempty"`
	Items  []TagListItem `json:"items"`
}

// CategoryListItem is an item returned by ListCategories.
type CategoryListItem struct {
	CategoryID string       `json:"category_id"`
	Info       CategoryInfo `json:"info"`
}

// CategoryInfo is the information about a category in a CategoryListItem.
type CategoryInfo struct {
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	Cardinality     string   `json:"cardinality,omitempty"`
	AssociableTypes []string `json:"associable_types,omitempty"`
	UsedBy          []string `json:"used_by,omitempty"`
}

// CategoryListResult is a single page of the categories list response.
type CategoryListResult struct {
	Marker string             `json:"marker,omitempty"`
	Items  []CategoryListItem `json:"items"`
}

// list walks all pages of the given list operation, calling add for each page.
// The marker returned by the server is followed until it is empty, so callers
// never deal with paging.
func (c *Manager) list(ctx context.Context, path string, names []string, page func(*http.Request) (marker string, err error)) error {
	marker := ""
	for {
		url := c.Resource(path)
		for _, name := range names {
			url.WithParam("names", name)
		}
		if marker != "" {
			url.WithParam("marker", marker)
		}

		var err error
		marker, err = page(url.Request(http.MethodGet))
		if err != nil {
			return err
		}
		if marker == "" {
			return nil
		}
	}
}

// ListTags returns all tags using the "GET /api/vcenter/tagging/tags" API (vSphere 9.0+).
// If names are specified, only tags with a name equal to one of names are returned.
func (c *Manager) ListTags(ctx context.Context, names ...string) ([]TagListItem, error) {
	var items []TagListItem
	err := c.list(ctx, internal.APITagPath, names, func(req *http.Request) (string, error) {
		var res TagListResult
		if err := c.Do(ctx, req, &res); err != nil {
			return "", err
		}
		items = append(items, res.Items...)
		return res.Marker, nil
	})
	return items, err
}

// ListTagsForCategory returns the tags in the category with the given name.
// The category is looked up first (by name, using ListCategories) and its ID is
// matched against each tag's category, as the tags API has no category filter.
// If names are specified, only tags with a name equal to one of names are returned.
// An error is returned if there is no category with the given name.
// Tag names are only unique within a category, so using both category and names
// identifies at most one tag.
func (c *Manager) ListTagsForCategory(ctx context.Context, category string, names ...string) ([]TagListItem, error) {
	categories, err := c.ListCategories(ctx, category)
	if err != nil {
		return nil, err
	}

	var id string
	for _, cat := range categories {
		if cat.Info.Name == category { // names filter is exact, but be explicit
			id = cat.CategoryID
			break
		}
	}
	if id == "" {
		return nil, fmt.Errorf("category %q not found", category)
	}

	items, err := c.ListTags(ctx, names...)
	if err != nil {
		return nil, err
	}

	var res []TagListItem
	for _, item := range items {
		if item.Info.Category == id {
			res = append(res, item)
		}
	}
	return res, nil
}

// ListCategories returns all categories using the "GET /api/vcenter/tagging/categories" API (vSphere 9.0+).
// If names are specified, only categories with a name equal to one of names are returned.
func (c *Manager) ListCategories(ctx context.Context, names ...string) ([]CategoryListItem, error) {
	var items []CategoryListItem
	err := c.list(ctx, internal.APICategoryPath, names, func(req *http.Request) (string, error) {
		var res CategoryListResult
		if err := c.Do(ctx, req, &res); err != nil {
			return "", err
		}
		items = append(items, res.Items...)
		return res.Marker, nil
	})
	return items, err
}
