/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package object

import (
	"errors"
	"strconv"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

var (
	// ErrKeyNameNotFound returned when a key by the specified name can't be found
	ErrKeyNameNotFound = errors.New("key name not found")
)

// CustomFieldsManager represents a client for custom fields
type CustomFieldsManager struct {
	Common
}

// GetCustomFieldsManager wraps NewCustomFieldsManager, returning ErrNotSupported
// when the client is not connected to a vCenter instance.
func GetCustomFieldsManager(c *vim25.Client) (*CustomFieldsManager, error) {
	if c.ServiceContent.CustomFieldsManager == nil {
		return nil, ErrNotSupported
	}
	return NewCustomFieldsManager(c), nil
}

// NewCustomFieldsManager creates a new custom fields client
func NewCustomFieldsManager(c *vim25.Client) *CustomFieldsManager {
	m := CustomFieldsManager{
		Common: NewCommon(c, *c.ServiceContent.CustomFieldsManager),
	}

	return &m
}

// Add a custom field
func (m CustomFieldsManager) Add(ctx context.Context, name string, moType string, fieldDefPolicy *types.PrivilegePolicyDef, fieldPolicy *types.PrivilegePolicyDef) (*types.CustomFieldDef, error) {
	req := types.AddCustomFieldDef{
		This:           m.Reference(),
		Name:           name,
		MoType:         moType,
		FieldDefPolicy: fieldDefPolicy,
		FieldPolicy:    fieldPolicy,
	}

	res, err := methods.AddCustomFieldDef(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// Remove a custom field
func (m CustomFieldsManager) Remove(ctx context.Context, key int) error {
	req := types.RemoveCustomFieldDef{
		This: m.Reference(),
		Key:  key,
	}

	_, err := methods.RemoveCustomFieldDef(ctx, m.c, &req)
	return err
}

// Rename a custom field
func (m CustomFieldsManager) Rename(ctx context.Context, key int, name string) error {
	req := types.RenameCustomFieldDef{
		This: m.Reference(),
		Key:  key,
		Name: name,
	}

	_, err := methods.RenameCustomFieldDef(ctx, m.c, &req)
	return err
}

// Set a custom field value
func (m CustomFieldsManager) Set(ctx context.Context, entity types.ManagedObjectReference, key int, value string) error {
	req := types.SetField{
		This:   m.Reference(),
		Entity: entity,
		Key:    key,
		Value:  value,
	}

	_, err := methods.SetField(ctx, m.c, &req)
	return err
}

// Field gets a custom field definition
func (m CustomFieldsManager) Field(ctx context.Context) ([]types.CustomFieldDef, error) {
	var fm mo.CustomFieldsManager

	err := m.Properties(ctx, m.Reference(), []string{"field"}, &fm)
	if err != nil {
		return nil, err
	}

	return fm.Field, nil
}

// FindKey finds a key id by key name
func (m CustomFieldsManager) FindKey(ctx context.Context, key string) (int, error) {
	field, err := m.Field(ctx)
	if err != nil {
		return -1, err
	}

	for _, def := range field {
		if def.Name == key {
			return def.Key, nil
		}
	}

	k, err := strconv.Atoi(key)
	if err == nil {
		// assume literal int key
		return k, nil
	}

	return -1, ErrKeyNameNotFound
}
