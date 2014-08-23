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

type OvfManager struct {
	c *Client
}

// CreateDescriptor wraps methods.CreateDescriptor
func (o OvfManager) CreateDescriptor(obj Reference, cdp types.OvfCreateDescriptorParams) (*types.OvfCreateDescriptorResult, error) {
	req := types.CreateDescriptor{
		This: *o.c.ServiceContent.OvfManager,
		Obj:  obj.Reference(),
		Cdp:  cdp,
	}

	res, err := methods.CreateDescriptor(o.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// CreateImportSpec wraps methods.CreateImportSpec
func (o OvfManager) CreateImportSpec(ovfDescriptor string, resourcePool Reference, datastore Reference, cisp types.OvfCreateImportSpecParams) (*types.OvfCreateImportSpecResult, error) {
	req := types.CreateImportSpec{
		This:          *o.c.ServiceContent.OvfManager,
		OvfDescriptor: ovfDescriptor,
		ResourcePool:  resourcePool.Reference(),
		Datastore:     datastore.Reference(),
		Cisp:          cisp,
	}

	res, err := methods.CreateImportSpec(o.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// ParseDescriptor wraps methods.ParseDescriptor
func (o OvfManager) ParseDescriptor(ovfDescriptor string, pdp types.OvfParseDescriptorParams) (*types.OvfParseDescriptorResult, error) {
	req := types.ParseDescriptor{
		This:          *o.c.ServiceContent.OvfManager,
		OvfDescriptor: ovfDescriptor,
		Pdp:           pdp,
	}

	res, err := methods.ParseDescriptor(o.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// ValidateHost wraps methods.ValidateHost
func (o OvfManager) ValidateHost(ovfDescriptor string, host Reference, vhp types.OvfValidateHostParams) (*types.OvfValidateHostResult, error) {
	req := types.ValidateHost{
		This:          *o.c.ServiceContent.OvfManager,
		OvfDescriptor: ovfDescriptor,
		Host:          host.Reference(),
		Vhp:           vhp,
	}

	res, err := methods.ValidateHost(o.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}
