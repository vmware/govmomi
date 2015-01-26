/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type CustomizationSpecManager struct {
	c *Client
}

func (cs CustomizationSpecManager) DoesCustomizationSpecExist(name string) (bool, error) {
	req := types.DoesCustomizationSpecExist{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Name: name,
	}

	res, err := methods.DoesCustomizationSpecExist(cs.c, &req)

	if err != nil {
		return false, err
	}

	return res.Returnval, nil
}

func (cs CustomizationSpecManager) GetCustomizationSpec(name string) (*types.CustomizationSpecItem, error) {
	req := types.GetCustomizationSpec{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Name: name,
	}

	res, err := methods.GetCustomizationSpec(cs.c, &req)

	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (cs CustomizationSpecManager) CreateCustomizationSpec(item types.CustomizationSpecItem) error {
	req := types.CreateCustomizationSpec{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Item: item,
	}

	_, err := methods.CreateCustomizationSpec(cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (cs CustomizationSpecManager) OverwriteCustomizationSpec(item types.CustomizationSpecItem) error {
	req := types.OverwriteCustomizationSpec{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Item: item,
	}

	_, err := methods.OverwriteCustomizationSpec(cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (cs CustomizationSpecManager) DeleteCustomizationSpec(name string) error {
	req := types.DeleteCustomizationSpec{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Name: name,
	}

	_, err := methods.DeleteCustomizationSpec(cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (cs CustomizationSpecManager) DuplicateCustomizationSpec(name string, newName string) error {
	req := types.DuplicateCustomizationSpec{
		This:    *cs.c.ServiceContent.CustomizationSpecManager,
		Name:    name,
		NewName: newName,
	}

	_, err := methods.DuplicateCustomizationSpec(cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (cs CustomizationSpecManager) RenameCustomizationSpec(name string, newName string) error {
	req := types.RenameCustomizationSpec{
		This:    *cs.c.ServiceContent.CustomizationSpecManager,
		Name:    name,
		NewName: newName,
	}

	_, err := methods.RenameCustomizationSpec(cs.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (cs CustomizationSpecManager) CustomizationSpecItemToXml(item types.CustomizationSpecItem) (string, error) {
	req := types.CustomizationSpecItemToXml{
		This: *cs.c.ServiceContent.CustomizationSpecManager,
		Item: item,
	}

	res, err := methods.CustomizationSpecItemToXml(cs.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (cs CustomizationSpecManager) XmlToCustomizationSpecItem(xml string) (*types.CustomizationSpecItem, error) {
	req := types.XmlToCustomizationSpecItem{
		This:        *cs.c.ServiceContent.CustomizationSpecManager,
		SpecItemXml: xml,
	}

	res, err := methods.XmlToCustomizationSpecItem(cs.c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}
