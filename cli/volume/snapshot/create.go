// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/types"
)

type create struct {
	*flags.ClientFlag
	*flags.OutputFlag

	id bool
}

func init() {
	cli.Register("volume.snapshot.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.id, "i", false, "Output snapshot ID and volume ID only")
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *create) Usage() string {
	return "[ID] [DESC]"
}

func (cmd *create) Description() string {
	return `Create snapshot of volume ID with description DESC.

Examples:
  govc volume.snapshot.create df86393b-5ae0-4fca-87d0-b692dbc67d45 my-snapshot
  govc volume.snapshot.create -i df86393b-5ae0-4fca-87d0-b692dbc67d45 my-snapshot`
}

type createResult struct {
	VolumeResults *types.CnsSnapshotCreateResult `json:"volumeResults"`
	cmd           *create
}

func (r *createResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(r.cmd.Out, 2, 0, 2, ' ', 0)
	if r.cmd.id {
		fmt.Fprintf(tw, "%s\t%s", r.VolumeResults.Snapshot.SnapshotId.Id,
			r.VolumeResults.Snapshot.VolumeId.Id)
	} else {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s", r.VolumeResults.Snapshot.SnapshotId.Id,
			r.VolumeResults.Snapshot.Description, r.VolumeResults.Snapshot.VolumeId.Id,
			r.VolumeResults.Snapshot.CreateTime.Format(time.Stamp))
	}
	fmt.Fprintln(tw)
	return tw.Flush()
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(f.Args()) != 2 {
		return flag.ErrHelp
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	spec := types.CnsSnapshotCreateSpec{
		VolumeId: types.CnsVolumeId{
			Id: f.Arg(0),
		},
		Description: f.Arg(1),
	}

	task, err := c.CreateSnapshots(ctx, []types.CnsSnapshotCreateSpec{spec})
	if err != nil {
		return err
	}

	info, err := cns.GetTaskInfo(ctx, task)
	if err != nil {
		return err
	}

	res, err := cns.GetTaskResult(ctx, info)
	if err != nil {
		return err
	}

	scr := res.(*types.CnsSnapshotCreateResult)
	if scr.CnsSnapshotOperationResult.Fault != nil {
		return errors.New(scr.CnsSnapshotOperationResult.Fault.LocalizedMessage)
	}

	return cmd.WriteResult(&createResult{VolumeResults: scr, cmd: cmd})
}
