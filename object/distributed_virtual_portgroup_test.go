// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

// DistributedVirtualPortgroup should implement the Reference interface.
var _ object.Reference = object.DistributedVirtualPortgroup{}

// DistributedVirtualPortgroup should implement the NetworkReference interface.
var _ object.NetworkReference = object.DistributedVirtualPortgroup{}

func TestDistributedVirtualPortgroupEthernetCardBackingInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		obj := simulator.Map(ctx).Any("DistributedVirtualPortgroup").(*simulator.DistributedVirtualPortgroup)

		pg := object.NewDistributedVirtualPortgroup(c, obj.Self)
		_, err := pg.EthernetCardBackingInfo(ctx)
		if err != nil {
			t.Fatal(err)
		}

		obj.Config.DistributedVirtualSwitch = nil // expect to fail if switch can't be read
		_, err = pg.EthernetCardBackingInfo(ctx)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
