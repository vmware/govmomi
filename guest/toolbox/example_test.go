// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/guest/toolbox"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleClient_Run() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		if !test.HasDocker() {
			fmt.Println("Linux")
			return nil
		}
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}
		err = simulator.RunContainer(ctx, c, vm, "nginx")
		if err != nil {
			return err
		}

		tools, err := toolbox.NewClient(ctx, c, vm, &types.NamePasswordAuthentication{
			Username: "user",
			Password: "pass",
		})

		cmd := &exec.Cmd{
			Path:   "uname",
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		return tools.Run(ctx, cmd)
	})
	// Output:
	// Linux
}
