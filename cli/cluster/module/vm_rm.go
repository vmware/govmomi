// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cluster"
	"github.com/vmware/govmomi/vim25/mo"
)

type rmVMs struct {
	*flags.SearchFlag
	moduleID string
}

func init() {
	cli.Register("cluster.module.vm.rm", &rmVMs{})
}

func (cmd *rmVMs) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.StringVar(&cmd.moduleID, "id", "", "Module ID")
}

func (cmd *rmVMs) Process(ctx context.Context) error {
	return cmd.SearchFlag.Process(ctx)
}

func (cmd *rmVMs) Usage() string {
	return `VM...`
}

func (cmd *rmVMs) Description() string {
	return `Remove VM(s) from a cluster module.

Examples:
  govc cluster.module.vm.rm -id module_id $vm...`
}

func (cmd *rmVMs) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() < 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	refs := make([]mo.Reference, 0, len(vms))
	for _, vm := range vms {
		refs = append(refs, vm.Reference())
	}

	allRemoved, err := cluster.NewManager(c).RemoveModuleMembers(ctx, cmd.moduleID, refs...)
	if err != nil {
		return err
	}

	if !allRemoved {
		return fmt.Errorf("a VM was not a member of the module")
	}

	return nil
}
