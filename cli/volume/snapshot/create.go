/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

	info, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return err
	}

	if res, ok := info.Result.(types.CnsVolumeOperationBatchResult); ok {
		if vres, ok := res.VolumeResults[0].(*types.CnsSnapshotCreateResult); ok {
			if vres.CnsSnapshotOperationResult.Fault != nil {
				return errors.New(vres.CnsSnapshotOperationResult.Fault.LocalizedMessage)
			}
			return cmd.WriteResult(&createResult{VolumeResults: vres, cmd: cmd})
		}
	}

	return nil
}
