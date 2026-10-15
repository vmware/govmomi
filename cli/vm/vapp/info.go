// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapp

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.vapp.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *info) Description() string {
	return `List VM vApp properties (OVF properties).

Examples:
  govc vm.vapp.info -vm $vm
  govc vm.vapp.info -vm $vm -json | jq .`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var mvm mo.VirtualMachine
	if err := vm.Properties(ctx, vm.Reference(), []string{"config.vAppConfig"}, &mvm); err != nil {
		return err
	}

	var props []types.VAppPropertyInfo
	if mvm.Config != nil && mvm.Config.VAppConfig != nil {
		props = mvm.Config.VAppConfig.GetVmConfigInfo().Property
	}

	return cmd.WriteResult(&infoResult{props})
}

type infoResult struct {
	Property []types.VAppPropertyInfo
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	_, _ = fmt.Fprintln(tw, "Key\tId\tLabel\tType\tValue")
	for _, p := range r.Property {
		_, _ = fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t%s\n", p.Key, p.Id, p.Label, p.Type, p.Value)
	}

	return tw.Flush()
}

func (r *infoResult) Dump() any {
	return r.Property
}

func (r *infoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Property)
}
