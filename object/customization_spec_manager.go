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
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// CustomizationSpecManager a customization spec client
type CustomizationSpecManager struct {
	Common
}

// NewCustomizationSpecManager creates a new customization spec client
func NewCustomizationSpecManager(c *vim25.Client) *CustomizationSpecManager {
	cs := CustomizationSpecManager{
		Common: NewCommon(c, *c.ServiceContent.CustomizationSpecManager),
	}

	return &cs
}

// DoesCustomizationSpecExist returns true when the customization spec exists
func (cs CustomizationSpecManager) DoesCustomizationSpecExist(ctx context.Context, name string) (bool, error) {
	req := types.DoesCustomizationSpecExist{
		This: cs.Reference(),
		Name: name,
	}

	res, err := methods.DoesCustomizationSpecExist(ctx, cs.c, &req)

	if err != nil {
		return false, err
	}

	return res.Returnval, nil
}

// GetCustomizationSpec gets a customization spec by name
func (cs CustomizationSpecManager) GetCustomizationSpec(ctx context.Context, name string) (*types.CustomizationSpecItem, error) {
	req := types.GetCustomizationSpec{
		This: cs.Reference(),
		Name: name,
	}

	res, err := methods.GetCustomizationSpec(ctx, cs.c, &req)

	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// CreateCustomizationSpec creates a customization spec
func (cs CustomizationSpecManager) CreateCustomizationSpec(ctx context.Context, item types.CustomizationSpecItem) error {
	req := types.CreateCustomizationSpec{
		This: cs.Reference(),
		Item: item,
	}

	_, err := methods.CreateCustomizationSpec(ctx, cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// OverwriteCustomizationSpec overwrites a customization spec
func (cs CustomizationSpecManager) OverwriteCustomizationSpec(ctx context.Context, item types.CustomizationSpecItem) error {
	req := types.OverwriteCustomizationSpec{
		This: cs.Reference(),
		Item: item,
	}

	_, err := methods.OverwriteCustomizationSpec(ctx, cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCustomizationSpec deletes a customization spec
func (cs CustomizationSpecManager) DeleteCustomizationSpec(ctx context.Context, name string) error {
	req := types.DeleteCustomizationSpec{
		This: cs.Reference(),
		Name: name,
	}

	_, err := methods.DeleteCustomizationSpec(ctx, cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// DuplicateCustomizationSpec duplicates a customization spec
func (cs CustomizationSpecManager) DuplicateCustomizationSpec(ctx context.Context, name string, newName string) error {
	req := types.DuplicateCustomizationSpec{
		This:    cs.Reference(),
		Name:    name,
		NewName: newName,
	}

	_, err := methods.DuplicateCustomizationSpec(ctx, cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RenameCustomizationSpec renames a customization spec
func (cs CustomizationSpecManager) RenameCustomizationSpec(ctx context.Context, name string, newName string) error {
	req := types.RenameCustomizationSpec{
		This:    cs.Reference(),
		Name:    name,
		NewName: newName,
	}

	_, err := methods.RenameCustomizationSpec(ctx, cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// CustomizationSpecItemToXml converts the given spec to xml
func (cs CustomizationSpecManager) CustomizationSpecItemToXml(ctx context.Context, item types.CustomizationSpecItem) (string, error) {
	req := types.CustomizationSpecItemToXml{
		This: cs.Reference(),
		Item: item,
	}

	res, err := methods.CustomizationSpecItemToXml(ctx, cs.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

// XmlToCustomizationSpecItem converts xml to a customization spec item
func (cs CustomizationSpecManager) XmlToCustomizationSpecItem(ctx context.Context, xml string) (*types.CustomizationSpecItem, error) {
	req := types.XmlToCustomizationSpecItem{
		This:        cs.Reference(),
		SpecItemXml: xml,
	}

	res, err := methods.XmlToCustomizationSpecItem(ctx, cs.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}
