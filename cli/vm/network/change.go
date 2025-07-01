// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type change struct {
	*flags.VirtualMachineFlag
	*flags.NetworkFlag
}

func init() {
	cli.Register("vm.network.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Usage() string {
	return "DEVICE"
}

func (cmd *change) Description() string {
	return `Change network DEVICE configuration.

Note that '-net' is currently required with '-net.address', even when not changing the VM network.

Examples:
  govc vm.network.change -vm $vm -net PG2 ethernet-0
  govc vm.network.change -vm $vm -net PG2 -net.address 00:00:0f:2e:5d:69 ethernet-0 # set to manual MAC address
  govc vm.network.change -vm $vm -net PG2 -net.address - ethernet-0 # set to generated MAC address
  govc device.info -vm $vm ethernet-*`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return errors.New("please specify a vm")
	}

	name := f.Arg(0)

	if name == "" {
		return errors.New("please specify a device name")
	}

	// Set network if specified as extra argument.
	if f.NArg() > 1 {
		err = cmd.NetworkFlag.Set(f.Arg(1))
		if err != nil {
			return fmt.Errorf("couldn't set specified network %v",
				err)
		}
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	net := devices.Find(name)

	if net == nil {
		return fmt.Errorf("device '%s' not found", name)
	}

	dev, err := cmd.NetworkFlag.Device()
	if err != nil {
		return err
	}

	cmd.NetworkFlag.Change(net, dev)

	return vm.EditDevice(ctx, net)
}
