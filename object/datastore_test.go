// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"strings"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
)

// Datastore should implement the Reference interface.
var _ object.Reference = object.Datastore{}

func TestDatastoreFindInventoryPath(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 2
	model.Datastore = 2

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		refs := simulator.Map(ctx).All("Datastore")

		for _, obj := range refs {
			ds := object.NewDatastore(c, obj.Reference())

			upload := func() error {
				p := soap.DefaultUpload
				r := strings.NewReader(ds.Name())
				p.ContentLength = r.Size()
				return ds.Upload(ctx, r, "name.txt", &p)

			}

			// with InventoryPath not set
			if err := upload(); err != nil {
				t.Error(err)
			}

			if ds.InventoryPath != "" {
				t.Error("InventoryPath should still be empty")
			}

			// Set InventoryPath and DatacenterPath
			err := ds.FindInventoryPath(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.HasPrefix(ds.InventoryPath, ds.DatacenterPath) {
				t.Errorf("InventoryPath=%s, DatacenterPath=%s", ds.InventoryPath, ds.DatacenterPath)
			}

			// with InventoryPath set
			if err := upload(); err != nil {
				t.Error(err)
			}

			tests := []struct {
				path, kind string
			}{
				{ds.InventoryPath, "Datastore"},
				{ds.DatacenterPath, "Datacenter"},
			}

			for _, test := range tests {
				ref, err := object.NewSearchIndex(c).FindByInventoryPath(ctx, test.path)
				if err != nil {
					t.Fatal(err)
				}

				if ref == nil {
					t.Fatalf("failed to find %s", test.path)
				}
				if ref.Reference().Type != test.kind {
					t.Errorf("%s is not a %s", test.path, test.kind)
				}
			}
		}
	}, model)
}

func TestDatastoreInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		obj, err := find.NewFinder(c).DefaultDatastore(ctx)
		if err != nil {
			t.Fatal(err)
		}
		pc := property.DefaultCollector(c)

		props := []string{
			"info.url",
			"info.name",
		}

		var ds mo.Datastore
		err = pc.RetrieveOne(ctx, obj.Reference(), props, &ds)
		if err != nil {
			t.Fatal(err)
		}

		info := ds.Info.GetDatastoreInfo()
		if info.Url == "" {
			t.Error("no info.url")
		}
		if info.Name == "" {
			t.Error("no info.name")
		}
	})
}
