// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package extension

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("extension.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "[KEY]..."
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetExtensionManager(c)
	if err != nil {
		return err
	}

	var res infoResult
	exts := make(map[string]types.Extension)

	if f.NArg() == 1 {
		e, err := m.Find(ctx, f.Arg(0))
		if err != nil {
			return err
		}
		if e != nil {
			exts[f.Arg(0)] = *e
		}
	} else {
		list, err := m.List(ctx)
		if err != nil {
			return err
		}
		if f.NArg() == 0 {
			res.Extensions = list
		} else {
			for _, e := range list {
				exts[e.Key] = e
			}
		}
	}

	for _, key := range f.Args() {
		if e, ok := exts[key]; ok {
			res.Extensions = append(res.Extensions, e)
		} else {
			return fmt.Errorf("extension %s not found", key)
		}
	}

	return cmd.WriteResult(&res)
}

type infoResult struct {
	Extensions []types.Extension `json:"extensions"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, e := range r.Extensions {
		fmt.Fprintf(tw, "Name:\t%s\n", e.Key)
		fmt.Fprintf(tw, "  Version:\t%s\n", e.Version)
		fmt.Fprintf(tw, "  Description:\t%s\n", e.Description.GetDescription().Summary)
		fmt.Fprintf(tw, "  Company:\t%s\n", e.Company)
		fmt.Fprintf(tw, "  Last heartbeat time:\t%s\n", e.LastHeartbeatTime)
		fmt.Fprintf(tw, "  Subject name:\t%s\n", e.SubjectName)
		fmt.Fprintf(tw, "  Type:\t%s\n", e.Type)
	}

	return tw.Flush()
}
