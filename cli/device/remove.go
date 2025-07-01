// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type remove struct {
	*flags.VirtualMachineFlag
	keepFiles bool
}

func init() {
	cli.Register("device.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.BoolVar(&cmd.keepFiles, "keep", false, "Keep files in datastore")
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "DEVICE..."
}

func (cmd *remove) Description() string {
	return `Remove DEVICE from VM.

Examples:
  govc device.remove -vm $name cdrom-3000
  govc device.remove -vm $name disk-1000
  govc device.remove -vm $name -keep disk-*`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
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

	for _, name := range f.Args() {
		device := match(name, devices)
		if len(device) == 0 {
			return fmt.Errorf("device '%s' not found", name)
		}

		if err = vm.RemoveDevice(ctx, cmd.keepFiles, device...); err != nil {
			return err
		}
	}

	return nil
}
