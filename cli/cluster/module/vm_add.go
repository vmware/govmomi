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

type addVMs struct {
	*flags.SearchFlag
	moduleID string
}

func init() {
	cli.Register("cluster.module.vm.add", &addVMs{})
}

func (cmd *addVMs) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.StringVar(&cmd.moduleID, "id", "", "Module ID")
}

func (cmd *addVMs) Process(ctx context.Context) error {
	return cmd.SearchFlag.Process(ctx)
}

func (cmd *addVMs) Usage() string {
	return `VM...`
}

func (cmd *addVMs) Description() string {
	return `Add VM(s) to a cluster module.

Examples:
  govc cluster.module.vm.add -id module_id $vm...`
}

func (cmd *addVMs) Run(ctx context.Context, f *flag.FlagSet) error {
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

	allAdded, err := cluster.NewManager(c).AddModuleMembers(ctx, cmd.moduleID, refs...)
	if err != nil {
		return err
	}

	if !allAdded {
		return fmt.Errorf("a VM is already a member of the module or not within the module's cluster")
	}

	return nil
}
