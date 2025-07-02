// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package floppy

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("device.clock.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *add) Description() string {
	return `Add precision clock device to VM.

Examples:
  govc device.clock.add -vm $vm
  govc device.info clock-*`
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	d := &types.VirtualPrecisionClock{
		VirtualDevice: types.VirtualDevice{
			Backing: &types.VirtualPrecisionClockSystemClockBackingInfo{
				Protocol: string(types.HostDateTimeInfoProtocolPtp), // TODO: ntp option
			},
		},
	}

	err = vm.AddDevice(ctx, d)
	if err != nil {
		return err
	}

	// output name of device we just created
	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	devices = devices.SelectByType(d)

	name := devices.Name(devices[len(devices)-1])

	fmt.Println(name)

	return nil
}
