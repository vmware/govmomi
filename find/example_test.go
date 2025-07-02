// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package find_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func ExampleMultipleFoundError() {
	model := simulator.VPX()
	model.Portgroup = 2

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)

		_, err := finder.Network(ctx, "DC0_DVPG*")
		_, ok := err.(*find.MultipleFoundError) // returns DC0_DVPG{0,1}
		if !ok {
			return errors.New("expected error")
		}

		net0, err := finder.Network(ctx, "DC0_DVPG0")
		if err != nil {
			return err
		}
		fmt.Println(net0.GetInventoryPath())

		net1, err := finder.Network(ctx, "DC0_DVPG1")
		if err != nil {
			return err
		}
		fmt.Println(net1.GetInventoryPath())

		return nil
	}, model)
	// Output:
	// /DC0/network/DC0_DVPG0
	// /DC0/network/DC0_DVPG1
}
