// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List libraries, items, and files.

Examples:
  govc library.ls
  govc library.ls /lib1
  govc library.ls /lib1/item1
  govc library.ls /lib1/item1/
  govc library.ls */
  govc library.ls -json | jq .
  govc library.ls /lib1/item1 -json | jq .`
}

type lsResultsWriter []finder.FindResult

func (r lsResultsWriter) Write(w io.Writer) error {
	for _, i := range r {
		fmt.Fprintln(w, i.GetPath())
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	finder := finder.NewFinder(m)
	findResults, err := finder.Find(ctx, f.Args()...)
	if err != nil {
		return err
	}
	return cmd.WriteResult(lsResultsWriter(findResults))
}
