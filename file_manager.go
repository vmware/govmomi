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
	"github.com/vmware/govmomi/vim25/tasks"
	"github.com/vmware/govmomi/vim25/types"
)

type FileManager struct {
	c *Client
}

// DeleteDatastoreFile deletes the specified file or folder from the datastore.
func (f FileManager) DeleteDatastoreFile(name string, dc *Datacenter) error {
	req := types.DeleteDatastoreFile_Task{
		This: *f.c.ServiceContent.FileManager,
		Name: name,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	task, err := tasks.DeleteDatastoreFile(f.c, &req)
	if err != nil {
		return err
	}

	return f.c.waitForTask(task)
}

// MakeDirectory creates a folder using the specified name.
func (f FileManager) MakeDirectory(name string, dc *Datacenter, createParentDirectories bool) error {
	req := types.MakeDirectory{
		This: *f.c.ServiceContent.FileManager,
		Name: name,
		CreateParentDirectories: createParentDirectories,
	}

	if dc != nil {
		ref := dc.Reference()
		req.Datacenter = &ref
	}

	_, err := methods.MakeDirectory(f.c, &req)
	return err
}
