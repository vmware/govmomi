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

type eject struct {
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("device.cdrom.eject", &eject{})
}

func (cmd *eject) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "CD-ROM device name")
}

func (cmd *eject) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *eject) Description() string {
	return `Eject media from CD-ROM device.

If device is not specified, the first CD-ROM device is used.

Examples:
  govc device.cdrom.eject -vm vm-1
  govc device.cdrom.eject -vm vm-1 -device floppy-1`
}

func (cmd *eject) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
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

	return vm.EditDevice(ctx, devices.EjectIso(c))
}
