// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
)

type ls struct {
	*flags.OutputFlag
	*flags.ClientFlag

	key      string
	prefix   string
	snapshot string
}

func init() {
	cli.Register("disk.metadata.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.key, "K", "", "Get value for key only")
	f.StringVar(&cmd.prefix, "p", "", "Key filter prefix")
	f.StringVar(&cmd.snapshot, "s", "", "Snapshot ID")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "ID"
}

func (cmd *ls) Description() string {
	return `List metadata for disk ID.

Examples:
  govc disk.metadata.ls 9b06a8b-d047-4d3c-b15b-43ea9608b1a6`
}

type lsResult []types.KeyValue

func (r lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	for _, data := range r {
		fmt.Fprintf(tw, "%s\t%s\n", data.Key, data.Value)
	}
	return tw.Flush()
}

func (r lsResult) Dump() any {
	return []types.KeyValue(r)
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vc, err := vslm.NewClient(ctx, c)
	if err != nil {
		return err
	}

	m := vslm.NewGlobalObjectManager(vc)

	id := types.ID{Id: f.Arg(0)}
	var data []types.KeyValue
	var sid *types.ID
	if cmd.snapshot != "" {
		sid = &types.ID{Id: cmd.snapshot}
	}

	if cmd.key != "" {
		val, err := m.RetrieveMetadataValue(ctx, id, sid, cmd.key)
		if err != nil {
			return err
		}
		data = []types.KeyValue{{Key: cmd.key, Value: val}}
	} else {
		data, err = m.RetrieveMetadata(ctx, id, sid, cmd.prefix)
		if err != nil {
			return err
		}
	}

	return cmd.WriteResult(lsResult(data))
}
