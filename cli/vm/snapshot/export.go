// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type export struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("snapshot.export", &export{})
}

func (cmd *export) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *export) Usage() string {
	return "NAME"
}

func (cmd *export) Description() string {
	return `Export snapshot of VM with given NAME.

NAME can be the snapshot name, tree path, or managed object ID.

Examples:
  govc snapshot.export -vm my-vm my-snapshot`
}

func (cmd *export) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *export) Run(ctx context.Context, f *flag.FlagSet) error {
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

	s, err := vm.FindSnapshot(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	l, err := vm.ExportSnapshot(ctx, s)
	if err != nil {
		return err
	}

	o, err := l.Properties(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(o)
}
