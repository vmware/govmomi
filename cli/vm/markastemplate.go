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

type markastemplate struct {
	*flags.SearchFlag
}

func init() {
	cli.Register("vm.markastemplate", &markastemplate{})
}

func (cmd *markastemplate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)
}

func (cmd *markastemplate) Process(ctx context.Context) error {
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *markastemplate) Usage() string {
	return "VM..."
}

func (cmd *markastemplate) Description() string {
	return `Mark VM as a virtual machine template.

Examples:
  govc vm.markastemplate $name`
}

func (cmd *markastemplate) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		err := vm.MarkAsTemplate(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
