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
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type shutdown struct {
	*flags.HostSystemFlag
	force  bool
	reboot bool
}

func init() {
	cli.Register("host.shutdown", &shutdown{})
}

func (cmd *shutdown) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(&cmd.force, "f", false, "Force shutdown when host is not in maintenance mode")
	f.BoolVar(&cmd.reboot, "r", false, "Reboot host")
}

func (cmd *shutdown) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *shutdown) Usage() string {
	return `HOST...`
}

func (cmd *shutdown) Description() string {
	return `Shutdown HOST.`
}

func (cmd *shutdown) Shutdown(ctx context.Context, host *object.HostSystem) error {
	req := types.ShutdownHost_Task{
		This:  host.Reference(),
		Force: cmd.force,
	}

	res, err := methods.ShutdownHost_Task(ctx, host.Client(), &req)
	if err != nil {
		return err
	}

	task := object.NewTask(host.Client(), res.Returnval)

	logger := cmd.ProgressLogger(fmt.Sprintf("%s shutdown... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *shutdown) Reboot(ctx context.Context, host *object.HostSystem) error {
	req := types.RebootHost_Task{
		This:  host.Reference(),
		Force: cmd.force,
	}

	res, err := methods.RebootHost_Task(ctx, host.Client(), &req)
	if err != nil {
		return err
	}

	task := object.NewTask(host.Client(), res.Returnval)

	logger := cmd.ProgressLogger(fmt.Sprintf("%s reboot... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *shutdown) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	s := cmd.Shutdown
	if cmd.reboot {
		s = cmd.Reboot
	}

	for _, host := range hosts {
		err = s(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
