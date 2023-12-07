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

package object_test

import (
	"context"
	"strings"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

// Datastore should implement the Reference interface.
var _ object.Reference = object.Datastore{}

func TestDatastoreFindInventoryPath(t *testing.T) {
	model := simulator.VPX()
	model.Datacenter = 2
	model.Datastore = 2

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		refs := simulator.Map.All("Datastore")

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
