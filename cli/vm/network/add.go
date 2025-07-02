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

type add struct {
	*flags.VirtualMachineFlag
	*flags.NetworkFlag
}

func init() {
	cli.Register("vm.network.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)
}

func (cmd *add) Description() string {
	return `Add network adapter to VM.

Examples:
  govc vm.network.add -vm $vm -net "VM Network" -net.adapter e1000e
  govc vm.network.add -vm $vm -net SwitchName/PortgroupName
  govc device.info -vm $vm ethernet-*`
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return errors.New("please specify a vm")
	}

	// Set network if specified as extra argument.
	if f.NArg() > 0 {
		err = cmd.NetworkFlag.Set(f.Arg(0))
		if err != nil {
			return fmt.Errorf("couldn't set specified network %v",
				err)
		}
	}

	net, err := cmd.NetworkFlag.Device()
	if err != nil {
		return err
	}

	return vm.AddDevice(ctx, net)
}
