// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type revert struct {
	*flags.VirtualMachineFlag

	suppressPowerOn bool
}

func init() {
	cli.Register("snapshot.revert", &revert{})
}

func (cmd *revert) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.suppressPowerOn, "s", false, "Suppress power on")
}

func (cmd *revert) Usage() string {
	return "[NAME]"
}

func (cmd *revert) Description() string {
	return `Revert to snapshot of VM with given NAME.

If NAME is not provided, revert to the current snapshot.
Otherwise, NAME can be the snapshot name, tree path or moid.

Examples:
  govc snapshot.revert -vm my-vm happy-vm-state`
}

func (cmd *revert) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *revert) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 1 {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var task *object.Task

	if f.NArg() == 1 {
		task, err = vm.RevertToSnapshot(ctx, f.Arg(0), cmd.suppressPowerOn)
	} else {
		task, err = vm.RevertToCurrentSnapshot(ctx, cmd.suppressPowerOn)
	}

	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
