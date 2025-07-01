// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pci

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("device.pci.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List allowed PCI passthrough devices that could be attach to VM.

Examples:
  govc device.pci.ls -vm VM`
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}

	vmConfigOptions, err := queryConfigTarget(ctx, vm)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&infoResult{PciDevices: vmConfigOptions.PciPassthrough})
}

type infoResult struct {
	PciDevices []types.BaseVirtualMachinePciPassthroughInfo `json:"pciDevices"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "System ID\tAddress\tDevice Name\n")
	for _, d := range r.PciDevices {
		info := d.GetVirtualMachinePciPassthroughInfo()
		pd := info.PciDevice
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", info.SystemId, pd.Id, pd.VendorName, pd.DeviceName)
	}
	return tw.Flush()
}

func queryConfigTarget(ctx context.Context, m *object.VirtualMachine) (*types.ConfigTarget, error) {
	b, err := m.EnvironmentBrowser(ctx)
	if err != nil {
		return nil, err
	}
	return b.QueryConfigTarget(ctx, nil)
}
