// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type tools struct {
	*flags.ClientFlag
	*flags.SearchFlag

	mount   bool
	upgrade bool
	options string
	unmount bool
}

func init() {
	cli.Register("vm.guest.tools", &tools{})
}

func (cmd *tools) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.BoolVar(&cmd.mount, "mount", false, "Mount tools CD installer in the guest")
	f.BoolVar(&cmd.upgrade, "upgrade", false, "Upgrade tools in the guest")
	f.StringVar(&cmd.options, "options", "", "Installer options")
	f.BoolVar(&cmd.unmount, "unmount", false, "Unmount tools CD installer in the guest")
}

func (cmd *tools) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *tools) Usage() string {
	return "VM..."
}

func (cmd *tools) Description() string {
	return `Manage guest tools in VM.

Examples:
  govc vm.guest.tools -mount VM
  govc vm.guest.tools -unmount VM
  govc vm.guest.tools -upgrade -options "opt1 opt2" VM`
}

func (cmd *tools) Upgrade(ctx context.Context, vm *object.VirtualMachine) error {
	task, err := vm.UpgradeTools(ctx, cmd.options)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}

func (cmd *tools) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		switch {
		case cmd.mount:
			err = vm.MountToolsInstaller(ctx)
			if err != nil {
				return err
			}
		case cmd.upgrade:
			err = cmd.Upgrade(ctx, vm)
			if err != nil {
				return err
			}
		case cmd.unmount:
			err = vm.UnmountToolsInstaller(ctx)
			if err != nil {
				return err
			}
		default:
			return flag.ErrHelp
		}
	}

	return nil
}
