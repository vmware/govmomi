// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package device

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/vm"
	"github.com/vmware/govmomi/vim25/types"
)

type boot struct {
	*flags.VirtualMachineFlag

	firmware string
	order    string
	types.VirtualMachineBootOptions
}

func init() {
	cli.Register("device.boot", &boot{})
}

func (cmd *boot) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.Int64Var(&cmd.BootDelay, "delay", 0, "Delay in ms before starting the boot sequence")
	f.StringVar(&cmd.order, "order", "", "Boot device order [-,floppy,cdrom,ethernet,disk]")
	f.Int64Var(&cmd.BootRetryDelay, "retry-delay", 0, "Delay in ms before a boot retry")

	cmd.BootRetryEnabled = types.NewBool(false)
	f.BoolVar(cmd.BootRetryEnabled, "retry", false, "If true, retry boot after retry-delay")

	cmd.EnterBIOSSetup = types.NewBool(false)
	f.BoolVar(cmd.EnterBIOSSetup, "setup", false, "If true, enter BIOS setup on next boot")

	f.Var(flags.NewOptionalBool(&cmd.EfiSecureBootEnabled), "secure", "Enable EFI secure boot")
	f.StringVar(&cmd.firmware, "firmware", "", vm.FirmwareUsage)
}

func (cmd *boot) Description() string {
	return `Configure VM boot settings.

Examples:
  govc device.boot -vm $vm -delay 1000 -order floppy,cdrom,ethernet,disk
  govc device.boot -vm $vm -order - # reset boot order
  govc device.boot -vm $vm -firmware efi -secure
  govc device.boot -vm $vm -firmware bios -secure=false`
}

func (cmd *boot) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *boot) Run(ctx context.Context, f *flag.FlagSet) error {
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

	if cmd.order != "" {
		o := strings.Split(cmd.order, ",")
		cmd.BootOrder = devices.BootOrder(o)
	}

	spec := types.VirtualMachineConfigSpec{
		BootOptions: &cmd.VirtualMachineBootOptions,
		Firmware:    cmd.firmware,
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
