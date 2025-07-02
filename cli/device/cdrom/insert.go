// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cdrom

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
	cli.Register("device.cdrom.insert", &insert{})
}

func (cmd *insert) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "CD-ROM device name")
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
	return "ISO"
}

func (cmd *insert) Description() string {
	return `Insert media on datastore into CD-ROM device.

If device is not specified, the first CD-ROM device is used.

Examples:
  govc device.cdrom.insert -vm vm-1 -device cdrom-3000 images/boot.iso
  govc device.cdrom.insert -vm vm-1 library:/boot/linux/ubuntu.iso # Content Library ISO`
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

	c, err := devices.FindCdrom(cmd.device)
	if err != nil {
		return err
	}

	iso, err := cmd.FileBacking(ctx, f.Arg(0), false)
	if err != nil {
		return err
	}

	return vm.EditDevice(ctx, devices.InsertIso(c, iso))
}
