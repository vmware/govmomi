// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package floppy

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type insert struct {
	*flags.DatastoreFlag
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("device.floppy.insert", &insert{})
}

func (cmd *insert) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "Floppy device name")
}

func (cmd *insert) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *insert) Usage() string {
	return "IMG"
}

func (cmd *insert) Description() string {
	return `Insert IMG on datastore into floppy device.

If device is not specified, the first floppy device is used.

Examples:
  govc device.floppy.insert -vm vm-1 vm-1/config.img`
}

func (cmd *insert) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil || f.NArg() != 1 {
		return flag.ErrHelp
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	c, err := devices.FindFloppy(cmd.device)
	if err != nil {
		return err
	}

	img, err := cmd.DatastorePath(f.Arg(0))
	if err != nil {
		return nil
	}

	return vm.EditDevice(ctx, devices.InsertImg(c, img))
}
