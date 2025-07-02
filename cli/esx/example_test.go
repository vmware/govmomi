// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/cli/esx"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func ExampleExecutor_Run() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		host, err := find.NewFinder(c).HostSystem(ctx, "DC0_H0")
		if err != nil {
			return err
		}

		x, err := esx.NewExecutor(ctx, c, host)
		if err != nil {
			return err
		}

		res, err := x.Run(ctx, []string{"software", "vib", "list"})
		if err != nil {
			return err
		}

		for _, vib := range res.Values {
			fmt.Println(vib.Value("Name"))
		}

		return nil
	})
	// Output:
	// esx-ui
	// intelgpio
}
