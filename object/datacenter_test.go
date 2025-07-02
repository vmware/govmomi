// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
