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

type Folder struct {
	types.ManagedObjectReference
}

func (f Folder) Reference() types.ManagedObjectReference {
	return f.ManagedObjectReference
}

func (f Folder) Children(c *Client) ([]Reference, error) {
	var mf mo.Folder

	err := c.Properties(f.Reference(), []string{"childEntity"}, &mf)
	if err != nil {
		return nil, err
	}

	var rs []Reference

	for _, e := range mf.ChildEntity {
		// A folder contains managed entities, all of which are listed below.
		switch e.Type {
		case "Folder":
			rs = append(rs, Folder{ManagedObjectReference: e})
		case "Datacenter":
			rs = append(rs, Datacenter{ManagedObjectReference: e})
		case "VirtualMachine":
			panic("TODO")
		case "VirtualApp":
			panic("TODO")
		case "ComputeResource":
			panic("TODO")
		case "Network":
			rs = append(rs, Network{ManagedObjectReference: e})
		case "DistributedVirtualSwitch": // Skip
		case "DistributedVirtualPortgroup": // Skip
		case "Datastore":
			rs = append(rs, Datastore{ManagedObjectReference: e})
		default:
			panic("Unknown managed entity: " + e.Type)
		}
	}

	return rs, nil
}
