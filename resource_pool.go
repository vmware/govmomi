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

type ResourcePool struct {
	types.ManagedObjectReference

	InventoryPath string

	c *Client
}

func NewResourcePool(c *Client, ref types.ManagedObjectReference) *ResourcePool {
	return &ResourcePool{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (p ResourcePool) Reference() types.ManagedObjectReference {
	return p.ManagedObjectReference
}

func (p ResourcePool) ImportVApp(spec types.BaseImportSpec, folder *Folder, host *HostSystem) (*HttpNfcLease, error) {
	req := types.ImportVApp{
		This: p.Reference(),
		Spec: spec,
	}

	if folder != nil {
		ref := folder.Reference()
		req.Folder = &ref
	}

	if host != nil {
		ref := host.Reference()
		req.Host = &ref
	}

	res, err := methods.ImportVApp(p.c, &req)
	if err != nil {
		return nil, err
	}

	return NewHttpNfcLease(p.c, res.Returnval), nil
}

func (p ResourcePool) Create(name string, spec types.ResourceConfigSpec) (*ResourcePool, error) {
	req := types.CreateResourcePool{
		This: p.Reference(),
		Name: name,
		Spec: spec,
	}

	res, err := methods.CreateResourcePool(p.c, &req)
	if err != nil {
		return nil, err
	}

	return NewResourcePool(p.c, res.Returnval), nil
}

func (p ResourcePool) UpdateConfig(name string, config *types.ResourceConfigSpec) error {
	req := types.UpdateConfig{
		This:   p.Reference(),
		Name:   name,
		Config: config,
	}

	if config != nil && config.Entity == nil {
		ref := p.Reference()

		// Create copy of config so changes won't leak back to the caller
		newConfig := *config
		newConfig.Entity = &ref
		req.Config = &newConfig
	}

	_, err := methods.UpdateConfig(p.c, &req)
	return err
}

func (p ResourcePool) DestroyChildren() error {
	req := types.DestroyChildren{
		This: p.Reference(),
	}

	_, err := methods.DestroyChildren(p.c, &req)
	return err
}

func (p ResourcePool) Destroy() (*Task, error) {
	req := types.Destroy_Task{
		This: p.Reference(),
	}

	res, err := methods.Destroy_Task(p.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(p.c, res.Returnval), nil
}
