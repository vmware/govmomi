/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package finder_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestResolveLibraryItemStorage(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)

		ds, err := find.NewFinder(vc).Datastore(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}

		var props mo.Datastore
		err = ds.Properties(ctx, ds.Reference(), []string{"name", "summary"}, &props)
		if err != nil {
			t.Fatal(err)
		}

		fsTypes := []string{
			string(types.HostFileSystemVolumeFileSystemTypeOTHER),
			string(types.HostFileSystemVolumeFileSystemTypeVsan),
		}

		for _, fs := range fsTypes {
			// client uses DatastoreNamespaceManager only when datastore fs is vsan/vvol
			simulator.Map.Get(ds.Reference()).(*simulator.Datastore).Summary.Type = fs

			u := props.Summary.Url

			storage := []library.Storage{
				{
					StorageBacking: library.StorageBacking{DatastoreID: ds.Reference().Value, Type: "DATASTORE"},
					StorageURIs: []string{
						fmt.Sprintf("%s/contentlib-${lib_id}/${item_id}/${file_name}_${file_id}.iso", u),
						fmt.Sprintf("%s/contentlib-${lib_id}/${item_id}/${file_name}_${file_id}.iso?serverId=${server_id}", u),
					},
				},
			}

			f := finder.NewPathFinder(library.NewManager(rc), vc)

			err = f.ResolveLibraryItemStorage(ctx, storage)
			if err != nil {
				t.Fatal(err)
			}

			var path object.DatastorePath
			for _, s := range storage {
				for _, u := range s.StorageURIs {
					path.FromString(u)
					if path.Datastore != props.Name {
						t.Errorf("failed to parse %s", u)
					}
					if strings.Contains(u, "?") {
						t.Errorf("includes query: %s", u)
					}
				}
			}
		}
	})
}
