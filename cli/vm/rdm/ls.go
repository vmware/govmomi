// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rdm

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
	cli.Register("vm.rdm.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List available devices that could be attach to VM with RDM.

Examples:
  govc vm.rdm.ls -vm VM`
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

	res := infoResult{
		Disks: vmConfigOptions.ScsiDisk,
	}
	return cmd.WriteResult(&res)
}

type infoResult struct {
	Disks []types.VirtualMachineScsiDiskDeviceInfo `json:"disks"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	for _, disk := range r.Disks {
		fmt.Fprintf(tw, "Name:\t%s\n", disk.Name)
		fmt.Fprintf(tw, "  Device name:\t%s\n", disk.Disk.DeviceName)
		fmt.Fprintf(tw, "  Device path:\t%s\n", disk.Disk.DevicePath)
		fmt.Fprintf(tw, "  Canonical Name:\t%s\n", disk.Disk.CanonicalName)

		var uids []string
		for _, descriptor := range disk.Disk.Descriptor {
			uids = append(uids, descriptor.Id)
		}

		fmt.Fprintf(tw, "  UIDS:\t%s\n", strings.Join(uids, " ,"))
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
