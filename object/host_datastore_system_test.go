// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/test"
)

func TestHostDatastoreSystemResignatureUnresolvedVmfsVolume(t *testing.T) {
	c := test.NewAuthenticatedClient(t)
	host := NewHostSystem(c, esx.HostSystem.Reference())

	ctx := context.Background()

	hds, err := host.ConfigManager().DatastoreSystem(ctx)
	if err != nil {
		t.Error(err)
	}
	hss, err := host.ConfigManager().StorageSystem(ctx)
	if err != nil {
		t.Error(err)
	}

	paths, err := hss.QueryUnresolvedVmfsVolumes(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(paths) == 0 {
		t.Skip("No unresolved vmfs volumes found")
	}

	task, err := hds.ResignatureUnresolvedVmfsVolumes(ctx, []string{paths[0].Extent[0].DevicePath})
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
