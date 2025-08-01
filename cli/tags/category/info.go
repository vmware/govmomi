// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package category

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("tags.category.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *info) Usage() string {
	return "[NAME]"
}

func (cmd *info) Description() string {
	return `Display category info.

If NAME is provided, display info for only that category.
Otherwise display info for all categories.

Examples:
  govc tags.category.info
  govc tags.category.info k8s-zone`
}

type infoResult []tags.Category

func (t infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, item := range t {
		fmt.Fprintf(tw, "Name:\t%s\n", item.Name)
		fmt.Fprintf(tw, "  ID:\t%s\n", item.ID)
		fmt.Fprintf(tw, "  Description:\t%s\n", item.Description)
		fmt.Fprintf(tw, "  Cardinality:\t%s\n", item.Cardinality)
		fmt.Fprintf(tw, "  AssociableTypes:\t%s\n", item.AssociableTypes)
		fmt.Fprintf(tw, "  UsedBy: \t%s\n", item.UsedBy)
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	arg := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := tags.NewManager(c)
	var res infoResult

	if f.NArg() == 1 {
		cat, cerr := m.GetCategory(ctx, arg)
		if cerr != nil {
			return cerr
		}
		res = append(res, *cat)
	} else {
		res, err = m.GetCategories(ctx)
		if err != nil {
			return err
		}
	}

	return cmd.WriteResult(res)
}
