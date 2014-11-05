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
	"path"

	"github.com/vmware/govmomi/vim25/types"
)

type DistributedVirtualPortgroup struct {
	types.ManagedObjectReference

	InventoryPath string

	c *Client
}

func NewDistributedVirtualPortgroup(c *Client, ref types.ManagedObjectReference) *DistributedVirtualPortgroup {
	return &DistributedVirtualPortgroup{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (p DistributedVirtualPortgroup) Reference() types.ManagedObjectReference {
	return p.ManagedObjectReference
}

func (p DistributedVirtualPortgroup) Name() string {
	return path.Base(p.InventoryPath)
}
