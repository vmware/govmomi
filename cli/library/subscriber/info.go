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
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.subscriber.info", &info{})
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
	return "PUBLISHED-LIBRARY SUBSCRIPTION-ID"
}

func (cmd *info) Description() string {
	return `Library subscriber info.

Examples:
  id=$(govc library.subscriber.ls | grep my-library-name | awk '{print $1}')
  govc library.subscriber.info published-library-name $id`
}

// path returns the inventory path for id, if possible
func (cmd *info) path(kind, id string) string {
	if id == "" {
		return ""
	}
	ref := types.ManagedObjectReference{Type: kind, Value: id}
	c, err := cmd.Client()
	if err == nil {
		ctx := context.Background()
		e, err := find.NewFinder(c, false).Element(ctx, ref)
		if err == nil {
			return e.Path
		}
	}
	return id
}

type infoResultsWriter struct {
	*library.Subscriber
	cmd *info
}

func (r infoResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	_, _ = fmt.Fprintf(tw, "Name:\t%s\n", r.LibraryName)
	_, _ = fmt.Fprintf(tw, "  ID:\t%s\n", r.LibraryID)
	_, _ = fmt.Fprintf(tw, "  Location:\t%s\n", r.LibraryLocation)

	p := r.cmd.path

	if r.Vcenter != nil {
		p = func(kind, id string) string { return id }
		port := ""
		if r.Vcenter.Port != 0 {
			port = fmt.Sprintf(":%d", r.Vcenter.Port)
		}
		_, _ = fmt.Fprintf(tw, "  vCenter:\t\n")
		_, _ = fmt.Fprintf(tw, "    URL:\thttps://%s%s\n", r.Vcenter.Hostname, port)
		_, _ = fmt.Fprintf(tw, "    GUID:\t%s\n", r.Vcenter.ServerGUID)
	}

	_, _ = fmt.Fprintf(tw, "  Placement:\t\n")
	_, _ = fmt.Fprintf(tw, "    Folder:\t%s\n", p("Folder", r.Placement.Folder))
	_, _ = fmt.Fprintf(tw, "    Cluster:\t%s\n", p("ClusterComputeResource", r.Placement.Cluster))
	_, _ = fmt.Fprintf(tw, "    Pool:\t%s\n", p("ResourcePool", r.Placement.ResourcePool))
	_, _ = fmt.Fprintf(tw, "    Network:\t%s\n", p("Network", r.Placement.Network))

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	lib, err := flags.ContentLibrary(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}
	m := library.NewManager(c)

	s, err := m.GetSubscriber(ctx, lib, f.Arg(1))
	if err != nil {
		return err
	}

	return cmd.WriteResult(infoResultsWriter{s, cmd})
}
