// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package dataset

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type ls struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vm.dataset.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List datasets.

Examples:
  govc vm.dataset.ls -vm $vm
  govc vm.dataset.ls -vm $vm -json | jq '.[].description'`
}

type lsResult []dataset.Summary

func (r lsResult) Write(w io.Writer) error {
	for _, d := range r {
		fmt.Fprintln(w, d.Name)
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
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
	l, err := mgr.ListDataSets(ctx, vmId)
	if err != nil {
		return err
	}
	return cmd.WriteResult(lsResult(l))
}
