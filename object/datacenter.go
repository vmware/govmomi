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
	"fmt"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// DatacenterFolders represents the datacenter folder paths
type DatacenterFolders struct {
	VMFolder        *Folder
	HostFolder      *Folder
	DatastoreFolder *Folder
	NetworkFolder   *Folder
}

// Datacenter a datacenter client
type Datacenter struct {
	Common
}

// NewDatacenter creates a new datacenter client
func NewDatacenter(c *vim25.Client, ref types.ManagedObjectReference) *Datacenter {
	return &Datacenter{
		Common: NewCommon(c, ref),
	}
}

// Folders gets the folders for the datacenter
func (d *Datacenter) Folders(ctx context.Context) (*DatacenterFolders, error) {
	var md mo.Datacenter

	ps := []string{"name", "vmFolder", "hostFolder", "datastoreFolder", "networkFolder"}
	err := d.Properties(ctx, d.Reference(), ps, &md)
	if err != nil {
		return nil, err
	}

	df := &DatacenterFolders{
		VMFolder:        NewFolder(d.c, md.VmFolder),
		HostFolder:      NewFolder(d.c, md.HostFolder),
		DatastoreFolder: NewFolder(d.c, md.DatastoreFolder),
		NetworkFolder:   NewFolder(d.c, md.NetworkFolder),
	}

	paths := []struct {
		name string
		path *string
	}{
		{"vm", &df.VMFolder.InventoryPath},
		{"host", &df.HostFolder.InventoryPath},
		{"datastore", &df.DatastoreFolder.InventoryPath},
		{"network", &df.NetworkFolder.InventoryPath},
	}

	for _, p := range paths {
		*p.path = fmt.Sprintf("/%s/%s", md.Name, p.name)
	}

	return df, nil
}

// Destroy the datacenter
func (d Datacenter) Destroy(ctx context.Context) (*Task, error) {
	req := types.Destroy_Task{
		This: d.Reference(),
	}

	res, err := methods.Destroy_Task(ctx, d.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(d.c, res.Returnval), nil
}
