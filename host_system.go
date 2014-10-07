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
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type HostSystem struct {
	types.ManagedObjectReference

	c *Client
}

func NewHostSystem(c *Client, ref types.ManagedObjectReference) *HostSystem {
	return &HostSystem{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (h HostSystem) Reference() types.ManagedObjectReference {
	return h.ManagedObjectReference
}

func (h HostSystem) ConfigManager(c *Client) *HostConfigManager {
	return &HostConfigManager{c, h}
}

func (h HostSystem) ResourcePool(c *Client) (*ResourcePool, error) {
	var mh mo.HostSystem
	err := c.Properties(h.Reference(), []string{"parent"}, &mh)
	if err != nil {
		return nil, err
	}

	var mcr mo.ComputeResource
	err = c.Properties(*mh.Parent, []string{"resourcePool"}, &mcr)
	if err != nil {
		return nil, err
	}

	pool := NewResourcePool(*mcr.ResourcePool)
	return pool, nil
}
