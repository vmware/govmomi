// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/device"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type promote struct {
	*flags.VirtualMachineFlag

	unlink bool
}

func init() {
	cli.Register("vm.disk.promote", &promote{})
}

func (cmd *promote) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.unlink, "unlink", true, "Unlink")
}

func (cmd *promote) Usage() string {
	return "DISK..."
}

func (cmd *promote) Description() string {
	return `Promote VM disk.

Examples:
  govc device.info -vm $name disk-*
  govc vm.disk.promote -vm $name disk-1000-0
  govc vm.disk.promote -vm $name disk-*`
}

func (cmd *promote) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	devices = devices.SelectByType((*types.VirtualDisk)(nil))

	devices, err = device.Match(devices, f.Args())
	if err != nil {
		return err
	}

	disks := make([]types.VirtualDisk, len(devices))
	for i := range devices {
		disks[i] = *(devices[i].(*types.VirtualDisk))
	}

	task, err := vm.PromoteDisks(ctx, cmd.unlink, disks)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger("Promoting disks...")
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
