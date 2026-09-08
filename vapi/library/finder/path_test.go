// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package finder_test

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	"github.com/vmware/govmomi/vim25/soap"
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

				dsName := "LocalDS_0"
				if v := tc.topLevelDirectoryCreateSupported; v != nil && *v == false {
					dsName = "vsanDatastore"

					err := enableVSAN(ctx, vf)
					if !assert.NoError(t, err) {
						t.FailNow()
					}
				}

				dc, err := vf.Datacenter(ctx, "*")
				if !assert.NoError(t, err) || !assert.NotNil(t, dc) {
					t.FailNow()
				}

				ds, err := vf.Datastore(ctx, dsName)
				if !assert.NoError(t, err) || !assert.NotNil(t, ds) {
					t.FailNow()
				}

				var (
					dsURL string
					moDS  mo.Datastore
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

// TestListLibraryItemStorageResolvesToDatastorePath verifies that the
// storage_uris reported by ListLibraryItemStorage for a real, uploaded
// library item file are rooted at the datastore's own summary.url, the way a
// real vCenter reports them, and not vcsim's internal on-disk temp
// directory. Regression test for the vcsim change in PR #4114
// ("vcsim: report a real vCenter-shaped Url for local datastores" and
// "vcsim: export Datastore.Path() and fix vapi/simulator's content library
// path"), which made Datastore.Summary.Url a synthesized "ds://" value
// without updating vapi/simulator's storage_uris to match, causing
// ResolveLibraryItemStorage's TrimPrefix(uri, ds.Summary.Url) to silently
// fail to strip the prefix and leak vcsim's real temp directory into the
// resolved path.
func TestListLibraryItemStorageResolvesToDatastorePath(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)
		if err := rc.Login(ctx, simulator.DefaultLogin); !assert.NoError(t, err) {
			t.FailNow()
		}

		vf := find.NewFinder(vc)
		dc, err := vf.Datacenter(ctx, "*")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		vf.SetDatacenter(dc)

		ds, err := vf.Datastore(ctx, "LocalDS_0")
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		m := library.NewManager(rc)

		libID, err := m.CreateLibrary(ctx, library.Library{
			Name: "test-lib",
			Type: "LOCAL",
			Storage: []library.StorageBacking{
				{
					DatastoreID: ds.Reference().Value,
					Type:        "DATASTORE",
				},
			},
		})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		itemID, err := m.CreateLibraryItem(ctx, library.Item{
			Name:      "test-item",
			Type:      "OVF",
			LibraryID: libID,
		})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		sessionID, err := m.CreateLibraryItemUpdateSession(
			ctx, library.Session{LibraryItemID: itemID})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		diskPath := "../../library/testdata/ttylinux-pc_i486-16.1-disk1.vmdk"
		f, err := os.Open(filepath.Clean(diskPath))
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		defer f.Close()

		fi, err := f.Stat()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		update, err := m.AddLibraryItemFile(ctx, sessionID, library.UpdateFile{
			Name:       "disk1.vmdk",
			SourceType: "PUSH",
			Size:       fi.Size(),
		})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		u, err := url.Parse(update.UploadEndpoint.URI)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		p := soap.DefaultUpload
		p.ContentLength = fi.Size()
		if !assert.NoError(t, m.Client.Upload(ctx, f, u, &p)) {
			t.FailNow()
		}

		if !assert.NoError(
			t, m.CompleteLibraryItemUpdateSession(ctx, sessionID)) {

			t.FailNow()
		}

		storage, err := m.ListLibraryItemStorage(ctx, itemID)
		if !assert.NoError(t, err) || !assert.Len(t, storage, 1) {
			t.FailNow()
		}
		if !assert.Len(t, storage[0].StorageURIs, 1) {
			t.FailNow()
		}

		// Before resolving, the reported storage_uris must already be shaped
		// like a real vCenter's datastore URL, not vcsim's local temp dir.
		rawURI := storage[0].StorageURIs[0]
		assert.NotContains(t, rawURI, os.TempDir())
		assert.NotContains(t, rawURI, "govcsim")

		lf := finder.NewPathFinder(m, vc)
		if !assert.NoError(
			t, lf.ResolveLibraryItemStorage(ctx, dc, nil, storage)) {

			t.FailNow()
		}

		resolved := storage[0].StorageURIs[0]
		assert.False(
			t,
			strings.Contains(resolved, os.TempDir()) ||
				strings.Contains(resolved, "govcsim"),
			"resolved storage URI leaked vcsim's internal path: %s", resolved)

		var dp object.DatastorePath
		assert.True(t, dp.FromString(resolved), "not a datastore path: %s", resolved)
		assert.Equal(t, "LocalDS_0", dp.Datastore)
		assert.True(
			t,
			strings.HasPrefix(dp.Path, "contentlib-"+libID+"/"+itemID+"/"),
			"unexpected relative path: %s", dp.Path)
	})
}

// TODO(dougm) consider vSAN enablement via simulator.Model
func enableVSAN(ctx context.Context, vf *find.Finder) error {
	cluster, err := vf.DefaultClusterComputeResource(ctx)
	if err != nil {
		return err
	}

	task, err := cluster.Reconfigure(ctx, &types.ClusterConfigSpecEx{
		VsanConfig: &types.VsanClusterConfigInfo{
			Enabled: types.NewBool(true),
		},
	}, true)

	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
