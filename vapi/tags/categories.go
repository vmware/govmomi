// Copyright 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tags

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	CategoryURL      = "/com/vmware/cis/tagging/category"
	ErrAlreadyExists = "already_exists"
)

type CategoryCreateSpec struct {
	CreateSpec CategoryCreate `json:"create_spec"`
}

type CategoryUpdateSpec struct {
	UpdateSpec CategoryUpdate `json:"update_spec,omitempty"`
}

type CategoryCreate struct {
	AssociableTypes []string `json:"associable_types"`
	Cardinality     string   `json:"cardinality"`
	Description     string   `json:"description"`
	Name            string   `json:"name"`
}

type CategoryUpdate struct {
	AssociableTypes []string `json:"associable_types,omitempty"`
	Cardinality     string   `json:"cardinality,omitempty"`
	Description     string   `json:"description,omitempty"`
	Name            string   `json:"name,omitempty"`
}

type Category struct {
	ID              string   `json:"id"`
	Description     string   `json:"description"`
	Name            string   `json:"name"`
	Cardinality     string   `json:"cardinality"`
	AssociableTypes []string `json:"associable_types"`
	UsedBy          []string `json:"used_by"`
}

func (c *RestClient) CreateCategoryIfNotExist(ctx context.Context, name string, description string, categoryType string, multiValue bool) (*string, error) {
	categories, err := c.GetCategoriesByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		var multiValueStr string
		if multiValue {
			multiValueStr = "MULTIPLE"
		} else {
			multiValueStr = "SINGLE"
		}
		categoryCreate := CategoryCreate{[]string{categoryType}, multiValueStr, description, name}
		spec := CategoryCreateSpec{categoryCreate}
		id, err := c.CreateCategory(ctx, &spec)
		if err != nil {
			// in case there are two docker daemon try to create inventory category, query the category once again
			if strings.Contains(err.Error(), "ErrAlreadyExists") {
				if categories, err = c.GetCategoriesByName(ctx, name); err != nil {
					return nil, fmt.Errorf("Failed to get inventory category for %s", err)

				}
			} else {
				return nil, fmt.Errorf("Failed to create inventory category for %s", err)

			}
		} else {
			return id, nil
		}
	}
	if categories != nil {
		return &categories[0].ID, nil
	}
	// should not happen
	return nil, fmt.Errorf("Failed to create inventory for it's existed, but could not query back. Please check system")

}

func (c *RestClient) CreateCategory(ctx context.Context, spec *CategoryCreateSpec) (*string, error) {
	stream, _, status, err := c.call(ctx, "POST", CategoryURL, spec, nil)

	if status != http.StatusOK || err != nil {
		return nil, fmt.Errorf("Create category failed with status code: %d, error message: %s", status, err)

	}

	type RespValue struct {
		Value string
	}

	var pID RespValue
	if err := json.NewDecoder(stream).Decode(&pID); err != nil {
		return nil, fmt.Errorf("Decode response body failed for: %s", err)

	}
	return &(pID.Value), nil
}

func (c *RestClient) GetCategory(ctx context.Context, id string) (*Category, error) {

	stream, _, status, err := c.call(ctx, "GET", fmt.Sprintf("%s/id:%s", CategoryURL, id), nil, nil)

	if status != http.StatusOK || err != nil {
		return nil, fmt.Errorf("Get category failed with status code: %d, error message: %s", status, err)

	}

	type RespValue struct {
		Value Category
	}

	var pCategory RespValue
	if err := json.NewDecoder(stream).Decode(&pCategory); err != nil {
		return nil, fmt.Errorf("Decode response body failed for: %s", err)

	}
	return &(pCategory.Value), nil
}

func (c *RestClient) UpdateCategory(ctx context.Context, id string, spec *CategoryUpdateSpec) error {
	_, _, status, err := c.call(ctx, "PATCH", fmt.Sprintf("%s/id:%s", CategoryURL, id), spec, nil)

	if status != http.StatusOK || err != nil {
		return fmt.Errorf("Update category failed with status code: %d, error message: %s", status, err)
	}

	return nil
}

func (c *RestClient) DeleteCategory(ctx context.Context, id string) error {

	_, _, status, err := c.call(ctx, "DELETE", fmt.Sprintf("%s/id:%s", CategoryURL, id), nil, nil)

	if status != http.StatusOK || err != nil {
		return fmt.Errorf("Delete category failed with status code: %d, error message: %s", status, err)

	}
	return nil
}

func (c *RestClient) ListCategories(ctx context.Context) ([]string, error) {

	stream, _, status, err := c.call(ctx, "GET", CategoryURL, nil, nil)

	if status != http.StatusOK || err != nil {
		return nil, fmt.Errorf("Get categories failed with status code: %d, error message: %s", status, err)

	}

	type Categories struct {
		Value []string
	}

	var pCategories Categories
	if err := json.NewDecoder(stream).Decode(&pCategories); err != nil {
		return nil, fmt.Errorf("Decode response body failed for: %s", err)

	}
	return pCategories.Value, nil
}

func (c *RestClient) GetCategoriesByName(ctx context.Context, name string) ([]Category, error) {
	categoryIds, err := c.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("Get category failed for: %s", err)

	}

	var categories []Category
	for _, cID := range categoryIds {
		category, err := c.GetCategory(ctx, cID)
		if err != nil {
			return nil, fmt.Errorf("Get category %s failed for %s", cID, err)
		}
		if category.Name == name {
			categories = append(categories, *category)
		}
	}
	return categories, nil
}
