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

type detach struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("disk.detach", &detach{})
}

func (cmd *detach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *detach) Usage() string {
	return "ID"
}

func (cmd *detach) Description() string {
	return `Detach disk ID from VM.

See also: govc device.remove

Examples:
  govc disk.detach -vm $vm ID`
}

func (cmd *detach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	id := f.Arg(0)

	return vm.DetachDisk(ctx, id)
}
