// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package entry

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	govc "github.com/vmware/govmomi/cli/vm/dataset"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type rm struct {
	*flags.VirtualMachineFlag
	dataSet string
}

func init() {
	cli.Register("vm.dataset.entry.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.StringVar(&cmd.dataSet, "dataset", "", "Data set name or ID")
}

func (cmd *rm) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *rm) Usage() string {
	return "KEY"
}

func (cmd *rm) Description() string {
	return `Delete data set entry.

Examples:
  govc vm.dataset.entry.rm -vm $vm -dataset com.example.project2 somekey`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	entryKey := f.Arg(0)

	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}
	vmId := vm.Reference().Value

	if cmd.dataSet == "" {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	mgr := dataset.NewManager(c)
	dataSetId, err := govc.FindDataSetId(ctx, mgr, vmId, cmd.dataSet)
	if err != nil {
		return err
	}
	return mgr.DeleteEntry(ctx, vmId, dataSetId, entryKey)
}
