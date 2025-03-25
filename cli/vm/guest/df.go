// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type df struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("guest.df", &df{})
}

func (cmd *df) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *df) Description() string {
	return `Report file system disk space usage.

Examples:
  govc guest.df -vm $name`
}

type dfResult []types.GuestDiskInfo

func (r dfResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", "Filesystem", "Size", "Used", "Avail", "Use%")
	for _, disk := range r {
		used := disk.Capacity - disk.FreeSpace
		use := 100.0 * float32(used) / float32(disk.Capacity)
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%.0f%%\n", disk.DiskPath,
			units.ByteSize(disk.Capacity), units.ByteSize(used), units.ByteSize(disk.FreeSpace), use)
	}
	return tw.Flush()
}

func (cmd *df) Run(ctx context.Context, f *flag.FlagSet) error {
	obj, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	var vm mo.VirtualMachine
	err = obj.Properties(ctx, obj.Reference(), []string{"guest.disk"}, &vm)
	if err != nil {
		return err
	}

	return cmd.WriteResult(dfResult(vm.Guest.Disk))
}
