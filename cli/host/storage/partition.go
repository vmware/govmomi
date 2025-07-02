// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package storage

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

type partition struct {
	*flags.HostSystemFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("host.storage.partition", &partition{})
}

func (cmd *partition) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *partition) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *partition) Usage() string {
	return "DEVICE_PATH"
}

func (cmd *partition) Description() string {
	return `Show partition table for device at DEVICE_PATH.`
}

func (cmd *partition) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return fmt.Errorf("specify device path")
	}

	path := f.Args()[0]

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	ss, err := host.ConfigManager().StorageSystem(ctx)
	if err != nil {
		return err
	}

	var hss mo.HostStorageSystem
	err = ss.Properties(ctx, ss.Reference(), nil, &hss)
	if err != nil {
		return nil
	}

	info, err := ss.RetrieveDiskPartitionInfo(ctx, path)
	if err != nil {
		return err
	}

	return cmd.WriteResult(partitionInfo(*info))
}

type partitionInfo types.HostDiskPartitionInfo

func (p partitionInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Table format: %s\n", p.Spec.PartitionFormat)
	fmt.Fprintf(tw, "Number of sectors: %d\n", p.Spec.TotalSectors)
	fmt.Fprintf(tw, "\n")

	fmt.Fprintf(tw, "Number\t")
	fmt.Fprintf(tw, "Start\t")
	fmt.Fprintf(tw, "End\t")
	fmt.Fprintf(tw, "Size\t")
	fmt.Fprintf(tw, "Type\t")
	fmt.Fprintf(tw, "\n")

	for _, e := range p.Spec.Partition {
		sectors := e.EndSector - e.StartSector

		fmt.Fprintf(tw, "%d\t", e.Partition)
		fmt.Fprintf(tw, "%d\t", e.StartSector)
		fmt.Fprintf(tw, "%d\t", e.EndSector)
		fmt.Fprintf(tw, "%s\t", units.ByteSize(sectors*512))
		fmt.Fprintf(tw, "%s\t", e.Type)
		fmt.Fprintf(tw, "\n")
	}

	return tw.Flush()
}
