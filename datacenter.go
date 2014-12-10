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
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type DatacenterFolders struct {
	VmFolder        *Folder
	HostFolder      *Folder
	DatastoreFolder *Folder
	NetworkFolder   *Folder
}

type Datacenter struct {
	types.ManagedObjectReference

	c *Client
}

func NewDatacenter(c *Client, ref types.ManagedObjectReference) *Datacenter {
	return &Datacenter{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (d Datacenter) Reference() types.ManagedObjectReference {
	return d.ManagedObjectReference
}

func (d *Datacenter) Folders() (*DatacenterFolders, error) {
	var md mo.Datacenter

	ps := []string{"vmFolder", "hostFolder", "datastoreFolder", "networkFolder"}
	err := d.c.Properties(d.Reference(), ps, &md)
	if err != nil {
		return nil, err
	}

	df := &DatacenterFolders{
		VmFolder:        NewFolder(d.c, md.VmFolder),
		HostFolder:      NewFolder(d.c, md.HostFolder),
		DatastoreFolder: NewFolder(d.c, md.DatastoreFolder),
		NetworkFolder:   NewFolder(d.c, md.NetworkFolder),
	}

	return df, nil
}

func (d Datacenter) Destroy() (*Task, error) {
	req := types.Destroy_Task{
		This: d.Reference(),
	}

	res, err := methods.Destroy_Task(d.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(d.c, res.Returnval), nil
}
