// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/disk"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.DatastoreFlag
	long bool
}

func init() {
	cli.Register("disk.snapshot.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Usage() string {
	return "ID"
}

func (cmd *ls) Description() string {
	return `List snapshots for disk ID on DS.

Examples:
  govc disk.snapshot.ls -l 9b06a8b-d047-4d3c-b15b-43ea9608b1a6`
}

type lsResult struct {
	Info *types.VStorageObjectSnapshotInfo `json:"info"`
	cmd  *ls
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, o := range r.Info.Snapshots {
		_, _ = fmt.Fprintf(tw, "%s\t%s", o.Id.Id, o.Description)
		if r.cmd.long {
			created := o.CreateTime.Format(time.Stamp)
			_, _ = fmt.Fprintf(tw, "\t%s", created)
		}
		_, _ = fmt.Fprintln(tw)
	}

	return tw.Flush()
}

func (r *lsResult) Dump() any {
	return r.Info
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := disk.NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	snapshots, err := m.RetrieveSnapshotInfo(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	info := &types.VStorageObjectSnapshotInfo{Snapshots: snapshots}
	res := lsResult{Info: info, cmd: cmd}

	return cmd.WriteResult(&res)
}
