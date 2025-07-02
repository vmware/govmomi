// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type attach struct {
	*flags.VirtualMachineFlag
	*flags.DatastoreFlag
}

func init() {
	cli.Register("disk.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *attach) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.DatastoreFlag.Process(ctx)
}

func (cmd *attach) Usage() string {
	return "ID"
}

func (cmd *attach) Description() string {
	return `Attach disk ID on VM.

See also: govc vm.disk.attach

Examples:
  govc disk.attach -vm $vm ID
  govc disk.attach -vm $vm -ds $ds ID`
}

func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	m, err := NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	id := f.Arg(0)

	return m.AttachDisk(ctx, vm, id)
}
