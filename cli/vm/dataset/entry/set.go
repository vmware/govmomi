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

type set struct {
	*flags.VirtualMachineFlag
	dataSet string
}

func init() {
	cli.Register("vm.dataset.entry.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.StringVar(&cmd.dataSet, "dataset", "", "Data set name or ID")
}

func (cmd *set) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *set) Usage() string {
	return "KEY VALUE"
}

func (cmd *set) Description() string {
	return `Set the value of a data set entry.

Examples:
  govc vm.dataset.entry.set -vm $vm -dataset com.example.project2 somekey somevalue`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}
	entryKey := f.Arg(0)
	entryValue := f.Arg(1)

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
	return mgr.SetEntry(ctx, vmId, dataSetId, entryKey, entryValue)
}
