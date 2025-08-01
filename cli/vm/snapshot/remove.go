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

type remove struct {
	*flags.VirtualMachineFlag

	recursive   bool
	consolidate bool
}

func init() {
	cli.Register("snapshot.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.recursive, "r", false, "Remove snapshot children")
	f.BoolVar(&cmd.consolidate, "c", true, "Consolidate disks")
}

func (cmd *remove) Usage() string {
	return "NAME"
}

func (cmd *remove) Description() string {
	return `Remove snapshot of VM with given NAME.

NAME can be the snapshot name, tree path, moid or '*' to remove all snapshots.

Examples:
  govc snapshot.remove -vm my-vm happy-vm-state`
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
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

	if f.Arg(0) == "*" {
		task, err = vm.RemoveAllSnapshot(ctx, &cmd.consolidate)
	} else {
		task, err = vm.RemoveSnapshot(ctx, f.Arg(0), cmd.recursive, &cmd.consolidate)
	}

	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
