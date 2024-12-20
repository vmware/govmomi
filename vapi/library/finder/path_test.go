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
	"testing"

	"github.com/stretchr/testify/assert"

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

	testCases := []struct {
		name                             string
		nilDatacenter                    bool
		topLevelDirectoryCreateSupported *bool
	}{
		{
			name:                             "Nil datacenter and nil topLevelCreate",
			nilDatacenter:                    true,
			topLevelDirectoryCreateSupported: nil,
		},
		{
			name:                             "Nil datacenter and false topLevelCreate",
			nilDatacenter:                    true,
			topLevelDirectoryCreateSupported: types.New(false),
		},
		{
			name:                             "Nil datacenter and true topLevelCreate",
			nilDatacenter:                    true,
			topLevelDirectoryCreateSupported: types.New(true),
		},
		{
			name:                             "Non-nil datacenter and nil topLevelCreate",
			nilDatacenter:                    false,
			topLevelDirectoryCreateSupported: nil,
		},
		{
			name:                             "Non-Nil datacenter and false topLevelCreate",
			nilDatacenter:                    false,
			topLevelDirectoryCreateSupported: types.New(false),
		},
		{
			name:                             "Non-Nil datacenter and true topLevelCreate",
			nilDatacenter:                    false,
			topLevelDirectoryCreateSupported: types.New(true),
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			simulator.Test(func(ctx context.Context, vc *vim25.Client) {

				vf := find.NewFinder(vc)
				rc := rest.NewClient(vc)
				lf := finder.NewPathFinder(library.NewManager(rc), vc)

				dc, err := vf.Datacenter(ctx, "*")
				if !assert.NoError(t, err) || !assert.NotNil(t, dc) {
					t.FailNow()
				}

				ds, err := vf.Datastore(ctx, "*")
				if !assert.NoError(t, err) || !assert.NotNil(t, ds) {
					t.FailNow()
				}

				var (
					dsName string
					dsURL  string
					moDS   mo.Datastore
				)
				if !assert.NoError(
					t,
					ds.Properties(
						ctx,
						ds.Reference(),
						[]string{"name", "summary"},
						&moDS)) {
					t.FailNow()
				}

				dsName = moDS.Name
				dsURL = moDS.Summary.Url

				storage := []library.Storage{
					{
						StorageBacking: library.StorageBacking{
							DatastoreID: ds.Reference().Value,
							Type:        "DATASTORE",
						},
						StorageURIs: []string{
							fmt.Sprintf("%s/contentlib-${lib_id}/${item_id}/${file_1_name}_${file_1_id}.iso", dsURL),
							fmt.Sprintf("%s/contentlib-${lib_id}/${item_id}/${file_2_name}_${file_2_id}.iso?serverId=${server_id}", dsURL),
						},
					},
				}

				var fsType string
				if v := tc.topLevelDirectoryCreateSupported; v != nil && *v {
					fsType = string(types.HostFileSystemVolumeFileSystemTypeOTHER)
				} else {
					fsType = string(types.HostFileSystemVolumeFileSystemTypeVsan)
				}

				simulator.Map.WithLock(
					simulator.SpoofContext(),
					ds.Reference(),
					func() {
						ds := simulator.Map.Get(ds.Reference()).(*simulator.Datastore)
						ds.Summary.Type = fsType
						ds.Capability.TopLevelDirectoryCreateSupported = tc.topLevelDirectoryCreateSupported
					})

				if !assert.NoError(
					t,
					lf.ResolveLibraryItemStorage(ctx, dc, storage)) {

					t.FailNow()
				}

				assert.Len(t, storage, 1)
				assert.Len(t, storage[0].StorageURIs, 2)

				for _, s := range storage {
					for _, u := range s.StorageURIs {
						var path object.DatastorePath
						path.FromString(u)
						assert.Equal(t, path.Datastore, dsName)
						assert.NotContains(t, u, "?")
					}
				}
			})
		})
	}
}
