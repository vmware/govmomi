// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type markasvm struct {
	*flags.SearchFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("vm.markasvm", &markasvm{})
}

func (cmd *markasvm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)
	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *markasvm) Process(ctx context.Context) error {
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *markasvm) Usage() string {
	return "VM..."
}

func (cmd *markasvm) Description() string {
	return `Mark VM template as a virtual machine.

Examples:
  govc vm.markasvm -host host1 $name
  govc vm.markasvm -host host1 $name`
}

func (cmd *markasvm) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	pool, err := cmd.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemFlag.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	if pool == nil {
		if host == nil {
			return flag.ErrHelp
		}

		pool, err = host.ResourcePool(ctx)
		if err != nil {
			return err
		}
	}

	for _, vm := range vms {
		err := vm.MarkAsVirtualMachine(ctx, *pool, host)
		if err != nil {
			return err
		}
	}

	return nil
}
