// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
		datastoreMap                     map[string]mo.Datastore
		topLevelDirectoryCreateSupported *bool
	}{
		{
			name:                             "Nil datacenter and nil topLevelCreate",
			nilDatacenter:                    true,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: nil,
		},
		{
			name:                             "Nil datacenter and false topLevelCreate",
			nilDatacenter:                    true,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: types.New(false),
		},
		{
			name:                             "Nil datacenter and true topLevelCreate",
			nilDatacenter:                    true,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: types.New(true),
		},
		{
			name:                             "Non-nil datacenter and nil topLevelCreate",
			nilDatacenter:                    false,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: nil,
		},
		{
			name:                             "Non-Nil datacenter and false topLevelCreate",
			nilDatacenter:                    false,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: types.New(false),
		},
		{
			name:                             "Non-Nil datacenter and true topLevelCreate",
			nilDatacenter:                    false,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: types.New(true),
		},
		{
			name:                             "Nil datastoreMap",
			nilDatacenter:                    true,
			datastoreMap:                     nil,
			topLevelDirectoryCreateSupported: nil,
		},
		{
			name:                             "Non-Nil datastoreMap and true topLevelCreate",
			nilDatacenter:                    true,
			datastoreMap:                     map[string]mo.Datastore{},
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
						[]string{"name", "summary.url"},
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

				sctx := ctx.(*simulator.Context)
				sctx.Map.WithLock(
					sctx,
					ds.Reference(),
					func() {
						ds := sctx.Map.Get(ds.Reference()).(*simulator.Datastore)
						ds.Summary.Type = fsType
						ds.Capability.TopLevelDirectoryCreateSupported = tc.topLevelDirectoryCreateSupported
					})

				nilDSM := tc.datastoreMap == nil

				if !assert.NoError(
					t,
					lf.ResolveLibraryItemStorage(
						ctx,
						dc,
						tc.datastoreMap,
						storage)) {

					t.FailNow()
				}

				assert.Len(t, storage, 1)
				assert.Len(t, storage[0].StorageURIs, 2)

				if nilDSM {
					assert.Nil(t, tc.datastoreMap)
				} else if assert.NotNil(t, tc.datastoreMap) {
					if assert.Len(t, tc.datastoreMap, 1) {
						dsv := ds.Reference().Value
						if assert.Contains(t, tc.datastoreMap, dsv) {
							ds := tc.datastoreMap[dsv]
							assert.Equal(t, ds.Name, dsName)
							assert.Equal(t, ds.Summary.Url, dsURL)
							assert.Equal(t, ds.Capability.TopLevelDirectoryCreateSupported, tc.topLevelDirectoryCreateSupported)
						}
					}
				}

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
