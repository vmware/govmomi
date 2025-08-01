// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/vmware/govmomi/vapi/internal"
)

// Category provides methods to create, read, update, delete, and enumerate
// categories.
type Category struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	Cardinality     string   `json:"cardinality,omitempty"`
	AssociableTypes []string `json:"associable_types,omitempty"`
	UsedBy          []string `json:"used_by,omitempty"`
	CategoryID      string   `json:"category_id,omitempty"`
}

func (c *Category) hasType(kind string) bool {
	for _, k := range c.AssociableTypes {
		if kind == k {
			return true
		}
	}
	return false
}

// Patch merges Category changes from the given src.
// AssociableTypes can only be appended to and cannot shrink.
func (c *Category) Patch(src *Category) {
	if src.Name != "" {
		c.Name = src.Name
	}
	if src.Description != "" {
		c.Description = src.Description
	}
	if src.Cardinality != "" {
		c.Cardinality = src.Cardinality
	}
	// Note that in order to append to AssociableTypes any existing types must be included in their original order.
	for _, kind := range src.AssociableTypes {
		if !c.hasType(kind) {
			c.AssociableTypes = append(c.AssociableTypes, kind)
		}
	}
}

// CreateCategory creates a new category and returns the category ID.
func (c *Manager) CreateCategory(ctx context.Context, category *Category) (string, error) {
	// create avoids the annoyance of CreateTag requiring field keys to be included in the request,
	// even though the field value can be empty.
	type create struct {
		Name            string   `json:"name"`
		Description     string   `json:"description"`
		Cardinality     string   `json:"cardinality"`
		AssociableTypes []string `json:"associable_types"`
		CategoryID      string   `json:"category_id,omitempty"`
	}
	spec := struct {
		Category create `json:"create_spec"`
	}{
		Category: create{
			Name:            category.Name,
			Description:     category.Description,
			Cardinality:     category.Cardinality,
			AssociableTypes: category.AssociableTypes,
			CategoryID:      category.CategoryID,
		},
	}
	if spec.Category.AssociableTypes == nil {
		// otherwise create fails with invalid_argument
		spec.Category.AssociableTypes = []string{}
	}
	url := c.Resource(internal.CategoryPath)
	var res string
	if err := c.Do(ctx, url.Request(http.MethodPost, spec), &res); err != nil {
		return "", err
	}
	return res, nil
}

// UpdateCategory updates one or more of the AssociableTypes, Cardinality,
// Description and Name fields.
func (c *Manager) UpdateCategory(ctx context.Context, category *Category) error {
	spec := struct {
		Category Category `json:"update_spec"`
	}{
		Category: Category{
			AssociableTypes: category.AssociableTypes,
			Cardinality:     category.Cardinality,
			Description:     category.Description,
			Name:            category.Name,
		},
	}
	url := c.Resource(internal.CategoryPath).WithID(category.ID)
	return c.Do(ctx, url.Request(http.MethodPatch, spec), nil)
}

// DeleteCategory deletes a category.
func (c *Manager) DeleteCategory(ctx context.Context, category *Category) error {
	url := c.Resource(internal.CategoryPath).WithID(category.ID)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// GetCategory fetches the category information for the given identifier.
// The id parameter can be a Category ID or Category Name.
func (c *Manager) GetCategory(ctx context.Context, id string) (*Category, error) {
	if isName(id) {
		cat, err := c.GetCategories(ctx)
		if err != nil {
			return nil, err
		}

		for i := range cat {
			if cat[i].Name == id {
				return &cat[i], nil
			}
		}
	}
	url := c.Resource(internal.CategoryPath).WithID(id)
	var res Category
	if err := c.Do(ctx, url.Request(http.MethodGet), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListCategories returns all category IDs in the system.
func (c *Manager) ListCategories(ctx context.Context) ([]string, error) {
	url := c.Resource(internal.CategoryPath)
	var res []string
	if err := c.Do(ctx, url.Request(http.MethodGet), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetCategories fetches a list of category information in the system.
func (c *Manager) GetCategories(ctx context.Context) ([]Category, error) {
	ids, err := c.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("list categories: %s", err)
	}

	var categories []Category
	for _, id := range ids {
		category, err := c.GetCategory(ctx, id)
		if err != nil {
			if strings.Contains(err.Error(), http.StatusText(http.StatusNotFound)) {
				continue // deleted since last fetch
			}
			return nil, fmt.Errorf("get category %s: %v", id, err)
		}
		categories = append(categories, *category)
	}

	return categories, nil
}
