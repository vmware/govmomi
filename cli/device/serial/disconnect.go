// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package serial

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type disconnect struct {
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("device.serial.disconnect", &disconnect{})
}

func (cmd *disconnect) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.device, "device", "", "serial port device name")
}

func (cmd *disconnect) Description() string {
	return `Disconnect service URI from serial port.

Examples:
  govc device.ls | grep serialport-
  govc device.serial.disconnect -vm $vm -device serialport-8000
  govc device.info -vm $vm serialport-*`
}

func (cmd *disconnect) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *disconnect) Run(ctx context.Context, f *flag.FlagSet) error {
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

	d, err := devices.FindSerialPort(cmd.device)
	if err != nil {
		return err
	}

	return vm.EditDevice(ctx, devices.DisconnectSerialPort(d))
}
