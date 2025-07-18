// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"errors"
	"math"
	"strconv"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	ErrKeyNameNotFound = errors.New("key name not found")
)

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

func NewCustomFieldsManager(c *vim25.Client) *CustomFieldsManager {
	m := CustomFieldsManager{
		Common: NewCommon(c, *c.ServiceContent.CustomFieldsManager),
	}

	return &m
}

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

func (m CustomFieldsManager) Remove(ctx context.Context, key int32) error {
	req := types.RemoveCustomFieldDef{
		This: m.Reference(),
		Key:  key,
	}

	_, err := methods.RemoveCustomFieldDef(ctx, m.c, &req)
	return err
}

func (m CustomFieldsManager) Rename(ctx context.Context, key int32, name string) error {
	req := types.RenameCustomFieldDef{
		This: m.Reference(),
		Key:  key,
		Name: name,
	}

	_, err := methods.RenameCustomFieldDef(ctx, m.c, &req)
	return err
}

func (m CustomFieldsManager) Set(ctx context.Context, entity types.ManagedObjectReference, key int32, value string) error {
	req := types.SetField{
		This:   m.Reference(),
		Entity: entity,
		Key:    key,
		Value:  value,
	}

	_, err := methods.SetField(ctx, m.c, &req)
	return err
}

type CustomFieldDefList []types.CustomFieldDef

func (m CustomFieldsManager) Field(ctx context.Context) (CustomFieldDefList, error) {
	var fm mo.CustomFieldsManager

	err := m.Properties(ctx, m.Reference(), []string{"field"}, &fm)
	if err != nil {
		return nil, err
	}

	return fm.Field, nil
}

func (m CustomFieldsManager) FindKey(ctx context.Context, name string) (int32, error) {
	field, err := m.Field(ctx)
	if err != nil {
		return -1, err
	}

	for _, def := range field {
		if def.Name == name {
			return def.Key, nil
		}
	}

	k, err := strconv.ParseInt(name, 10, 32)
	if err != nil {
		return -1, ErrKeyNameNotFound
	}

	if k >= math.MinInt32 && k <= math.MaxInt32 {
		return int32(k), nil
	}

	return -1, ErrKeyNameNotFound
}

func (l CustomFieldDefList) ByKey(key int32) *types.CustomFieldDef {
	for _, def := range l {
		if def.Key == key {
			return &def
		}
	}
	return nil
}
