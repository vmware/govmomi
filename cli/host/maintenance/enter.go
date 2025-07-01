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

type enter struct {
	*flags.HostSystemFlag

	timeout  int32
	evacuate bool
}

func init() {
	cli.Register("host.maintenance.enter", &enter{})
}

func (cmd *enter) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.Var(flags.NewInt32(&cmd.timeout), "timeout", "Timeout")
	f.BoolVar(&cmd.evacuate, "evacuate", false, "Evacuate powered off VMs")
}

func (cmd *enter) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *enter) Usage() string {
	return "HOST..."
}

func (cmd *enter) Description() string {
	return `Put HOST in maintenance mode.

While this task is running and when the host is in maintenance mode,
no VMs can be powered on and no provisioning operations can be performed on the host.`
}

func (cmd *enter) EnterMaintenanceMode(ctx context.Context, host *object.HostSystem) error {
	task, err := host.EnterMaintenanceMode(ctx, cmd.timeout, cmd.evacuate, nil) // TODO: spec param
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s entering maintenance mode... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *enter) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.EnterMaintenanceMode(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
