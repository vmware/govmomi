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

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	long bool
	id   bool
}

func init() {
	cli.Register("volume.snapshot.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.id, "i", false, "List snapshot ID and volume ID only")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "[ID]..."
}

func (cmd *ls) Description() string {
	return `List all snapshots of volume ID.

Use a list of volume IDs to list all snapshots of multiple volumes at once.

Examples:
  govc volume.snapshot.ls df86393b-5ae0-4fca-87d0-b692dbc67d45
  govc volume.snapshot.ls -i df86393b-5ae0-4fca-87d0-b692dbc67d45
  govc volume.snapshot.ls -l df86393b-5ae0-4fca-87d0-b692dbc67d45
  govc volume.snapshot.ls -l $(govc volume.ls -i)`
}

type lsResult struct {
	Entries []*types.CnsSnapshotQueryResultEntry `json:"entries"`
	cmd     *ls
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(r.cmd.Out, 2, 0, 2, ' ', 0)
	if r.cmd.id {
		for _, e := range r.Entries {
			fmt.Fprintf(tw, "%s\t%s", e.Snapshot.SnapshotId.Id, e.Snapshot.VolumeId.Id)
			fmt.Fprintln(tw)
		}
		return tw.Flush()
	}

	for _, e := range r.Entries {
		fmt.Fprintf(tw, "%s\t%s", e.Snapshot.SnapshotId.Id, e.Snapshot.Description)
		if r.cmd.long {
			fmt.Fprintf(tw, "\t%s\t%s", e.Snapshot.VolumeId.Id, e.Snapshot.CreateTime.Format(time.Stamp))
		}
		fmt.Fprintln(tw)
	}
	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(f.Args()) < 1 {
		return flag.ErrHelp
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	result := lsResult{cmd: cmd}

	for _, id := range f.Args() {
		spec := types.CnsSnapshotQueryFilter{
			SnapshotQuerySpecs: []types.CnsSnapshotQuerySpec{{
				VolumeId: types.CnsVolumeId{
					Id: id,
				},
			}},
		}

		for {
			task, err := c.QuerySnapshots(ctx, spec)
			if err != nil {
				return err
			}

			info, err := cns.GetTaskInfo(ctx, task)
			if err != nil {
				return err
			}

			res, err := cns.GetQuerySnapshotsTaskResult(ctx, info)
			if err != nil {
				return err
			}

			for i, e := range res.Entries {
				if e.Error != nil {
					return errors.New(e.Error.LocalizedMessage)
				}
				result.Entries = append(result.Entries, &res.Entries[i])
			}

			if res.Cursor.Offset == res.Cursor.TotalRecords {
				break
			}

			spec.Cursor = &res.Cursor
		}
	}

	return cmd.WriteResult(&result)
}
