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

type DatacenterFolders struct {
	VmFolder        Folder
	HostFolder      Folder
	DatastoreFolder Folder
	NetworkFolder   Folder
}

type Datacenter struct {
	types.ManagedObjectReference
}

func NewDatacenter(ref types.ManagedObjectReference) *Datacenter {
	return &Datacenter{
		ManagedObjectReference: ref,
	}
}

func (d Datacenter) Reference() types.ManagedObjectReference {
	return d.ManagedObjectReference
}

func (d *Datacenter) Folders(c *Client) (*DatacenterFolders, error) {
	var md mo.Datacenter

	ps := []string{"vmFolder", "hostFolder", "datastoreFolder", "networkFolder"}
	err := c.Properties(d.Reference(), ps, &md)
	if err != nil {
		return nil, err
	}

	df := &DatacenterFolders{
		VmFolder:        Folder{md.VmFolder},
		HostFolder:      Folder{md.HostFolder},
		DatastoreFolder: Folder{md.DatastoreFolder},
		NetworkFolder:   Folder{md.NetworkFolder},
	}

	return df, nil
}
