// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package maintenance

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type exit struct {
	*flags.HostSystemFlag

	timeout int32
}

func init() {
	cli.Register("host.maintenance.exit", &exit{})
}

func (cmd *exit) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.Var(flags.NewInt32(&cmd.timeout), "timeout", "Timeout")
}

func (cmd *exit) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *exit) Usage() string {
	return "HOST..."
}

func (cmd *exit) Description() string {
	return `Take HOST out of maintenance mode.

This blocks if any concurrent running maintenance-only host configurations operations are being performed.
For example, if VMFS volumes are being upgraded.

The 'timeout' flag is the number of seconds to wait for the exit maintenance mode to succeed.
If the timeout is less than or equal to zero, there is no timeout.`
}

func (cmd *exit) ExitMaintenanceMode(ctx context.Context, host *object.HostSystem) error {
	task, err := host.ExitMaintenanceMode(ctx, cmd.timeout)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s exiting maintenance mode... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *exit) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.ExitMaintenanceMode(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
