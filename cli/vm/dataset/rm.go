// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package dataset

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type rm struct {
	*flags.VirtualMachineFlag
	force bool
}

func init() {
	cli.Register("vm.dataset.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.BoolVar(&cmd.force, "force", false, "Delete the data set even if it has entries")
}

func (cmd *rm) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Delete data set.

Fails if the data set has entries, unless the '-force' flag is set.

Examples:
  govc vm.dataset.rm -vm $vm com.example.project2
  govc vm.dataset.rm -vm $vm -force=true com.example.project3`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}
	vmId := vm.Reference().Value

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	mgr := dataset.NewManager(c)

	id, err := FindDataSetId(ctx, mgr, vmId, f.Arg(0))
	if err != nil {
		return err
	}

	err = mgr.DeleteDataSet(ctx, vmId, id, cmd.force)
	if err != nil {
		return err
	}

	return nil
}
