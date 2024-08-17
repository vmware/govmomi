/*
Copyright (c) 2015-2024 VMware, Inc. All Rights Reserved.

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

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

// Datacenter should implement the Reference interface.
var _ object.Reference = object.Datacenter{}

func TestDatacenterFolders(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 2
	model.Folder = 1

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		search := object.NewSearchIndex(c)
		finder := find.NewFinder(c)

		dc, err := finder.Datacenter(ctx, "DC1")
		if err != nil {
			t.Fatal(err)
		}

		f, err := dc.Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		folders := []*object.Folder{
			f.DatastoreFolder,
			f.HostFolder,
			f.NetworkFolder,
			f.VmFolder,
		}

		for _, folder := range folders {
			p := "/F0/DC1/" + folder.Name()
			if p != folder.InventoryPath {
				t.Errorf("InventoryPath=%s", folder.InventoryPath)
			}

			ref, err := search.FindByInventoryPath(ctx, folder.InventoryPath)
			if err != nil {
				t.Fatal(err)
			}
			if ref == nil {
				t.Errorf("invalid InventoryPath: %s", folder.InventoryPath)
			}
		}
	}, model)
}
