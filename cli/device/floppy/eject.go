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

type eject struct {
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("device.floppy.eject", &eject{})
}

func (cmd *eject) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "Floppy device name")
}

func (cmd *eject) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *eject) Description() string {
	return `Eject image from floppy device.

If device is not specified, the first floppy device is used.

Examples:
  govc device.floppy.eject -vm vm-1`
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

	c, err := devices.FindFloppy(cmd.device)
	if err != nil {
		return err
	}

	return vm.EditDevice(ctx, devices.EjectImg(c))
}
