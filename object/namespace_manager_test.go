// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestDatastoreNamespaceManager(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewDatastoreNamespaceManager(c)

		finder := find.NewFinder(c)
		dc, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatal(err)
		}

		ds, err := finder.DefaultDatastore(ctx)
		if err != nil {
			t.Fatal(err)
		}

		store := simulator.Map(ctx).Get(ds.Reference()).(*simulator.Datastore)

		name := "foo"

		dir, err := m.CreateDirectory(ctx, ds, name, "")
		if !fault.Is(err, &types.CannotCreateFile{}) {
			t.Errorf("err=%v", err)
		}

		store.Summary.Type = string(types.HostFileSystemVolumeFileSystemTypeVsan)
		store.Capability.TopLevelDirectoryCreateSupported = types.NewBool(true)

		dir, err = m.CreateDirectory(ctx, ds, name, "")
		if err != nil {
			t.Errorf("err=%v", err)
		}

		err = m.DeleteDirectory(ctx, dc, name)
		if !fault.Is(err, &types.InvalidDatastorePath{}) {
			t.Errorf("err=%v", err)
		}

		err = m.DeleteDirectory(ctx, dc, dir)
		if err != nil {
			t.Errorf("delete %s, err=%v", dir, err)
		}
	})
}
