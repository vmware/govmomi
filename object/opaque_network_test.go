// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestOpaqueNetworkEthernetCardBackingInfo(t *testing.T) {
	model := simulator.VPX()
	model.OpaqueNetwork = 1 // Create 1 NSX backed OpaqueNetwork per DC

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)

		net, err := finder.Network(ctx, "DC0_NSX0")
		if err != nil {
			t.Fatal(err)
		}

		nsx, ok := net.(*object.OpaqueNetwork)
		if !ok {
			t.Fatalf("network type=%T", net)
		}

		summary, err := nsx.Summary(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if net.Reference() != *summary.Network {
			t.Fatal()
		}

		_, err = uuid.Parse(summary.OpaqueNetworkId)
		if err != nil {
			t.Errorf("parsing %q: %s", summary.OpaqueNetworkId, err)
		}

		_, err = net.EthernetCardBackingInfo(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}, model)
}
