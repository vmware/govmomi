// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package subscriber

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.subscriber.ls", &ls{})
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
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List library subscriptions.

Examples:
  govc library.subscriber.ls library-name`
}

type lsResultsWriter []library.SubscriberSummary

func (r lsResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "Subscription ID", "Library Name", "Library ID", "vCenter")

	for _, i := range r {
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", i.SubscriptionID, i.LibraryName, i.LibraryID, i.LibraryVcenterHostname)
	}

	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	lib, err := flags.ContentLibrary(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}
	m := library.NewManager(c)

	s, err := m.ListSubscribers(ctx, lib)
	if err != nil {
		return err
	}

	return cmd.WriteResult(lsResultsWriter(s))
}
