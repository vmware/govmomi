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

func TestDistributedVirtualSwitchEthernetCardBackingInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		obj := simulator.Map(ctx).Any("DistributedVirtualSwitch").(*simulator.DistributedVirtualSwitch)

		dvs := object.NewDistributedVirtualSwitch(c, obj.Self)

		_, err := dvs.EthernetCardBackingInfo(ctx)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
