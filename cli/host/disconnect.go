// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package host

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type disconnect struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.disconnect", &disconnect{})
}

func (cmd *disconnect) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *disconnect) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *disconnect) Description() string {
	return `Disconnect HOST from vCenter.`
}

func (cmd *disconnect) Disconnect(ctx context.Context, host *object.HostSystem) error {
	task, err := host.Disconnect(ctx)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s disconnecting... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *disconnect) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.Disconnect(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
