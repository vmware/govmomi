// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type create struct {
	*flags.VirtualMachineFlag

	description string
	memory      bool
	quiesce     bool
}

func init() {
	cli.Register("snapshot.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.memory, "m", true, "Include memory state")
	f.BoolVar(&cmd.quiesce, "q", false, "Quiesce guest file system")
	f.StringVar(&cmd.description, "d", "", "Snapshot description")
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Create snapshot of VM with NAME.

Examples:
  govc snapshot.create -vm my-vm happy-vm-state`
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
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

	task, err := vm.CreateSnapshot(ctx, f.Arg(0), cmd.description, cmd.memory, cmd.quiesce)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
