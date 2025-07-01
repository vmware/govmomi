// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package entry

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	govc "github.com/vmware/govmomi/cli/vm/dataset"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type get struct {
	*flags.VirtualMachineFlag
	dataSet string
}

func init() {
	cli.Register("vm.dataset.entry.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.StringVar(&cmd.dataSet, "dataset", "", "Data set name or ID")
}

func (cmd *get) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *get) Usage() string {
	return "KEY"
}

func (cmd *get) Description() string {
	return `Read the value of a data set entry.

Examples:
  govc vm.dataset.entry.get -vm $vm -dataset com.example.project2 somekey`
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
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
	entryValue, err := mgr.GetEntry(ctx, vmId, dataSetId, entryKey)
	if err != nil {
		return err
	}

	fmt.Println(entryValue)
	return nil
}
