// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package about

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type about struct {
	*flags.ClientFlag
	*flags.OutputFlag

	Long bool
	c    bool
}

func init() {
	cli.Register("about", &about{})
}

func (cmd *about) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.Long, "l", false, "Include service content")
	f.BoolVar(&cmd.c, "c", false, "Include client info")
}

func (cmd *about) Description() string {
	return `Display About info for HOST.

System information including the name, type, version, and build number.

Examples:
  govc about
  govc about -json | jq -r .about.productLineId`
}

func (cmd *about) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *about) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	res := infoResult{
		a: &c.ServiceContent.About,
	}

	if cmd.Long {
		res.Content = &c.ServiceContent
	} else {
		res.About = res.a
	}

	if cmd.c {
		res.Client = c.Client
	}

	return cmd.WriteResult(&res)
}

type infoResult struct {
	Content *types.ServiceContent `json:"content,omitempty"`
	About   *types.AboutInfo      `json:"about,omitempty"`
	Client  *soap.Client          `json:"client,omitempty"`
	a       *types.AboutInfo
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "FullName:\t%s\n", r.a.FullName)
	fmt.Fprintf(tw, "Name:\t%s\n", r.a.Name)
	fmt.Fprintf(tw, "Vendor:\t%s\n", r.a.Vendor)
	fmt.Fprintf(tw, "Version:\t%s\n", r.a.Version)
	fmt.Fprintf(tw, "Build:\t%s\n", r.a.Build)
	fmt.Fprintf(tw, "OS type:\t%s\n", r.a.OsType)
	fmt.Fprintf(tw, "API type:\t%s\n", r.a.ApiType)
	fmt.Fprintf(tw, "API version:\t%s\n", r.a.ApiVersion)
	fmt.Fprintf(tw, "Product ID:\t%s\n", r.a.ProductLineId)
	fmt.Fprintf(tw, "UUID:\t%s\n", r.a.InstanceUuid)
	return tw.Flush()
}

func (r *infoResult) Dump() any {
	if r.Content != nil {
		return r.Content
	}
	return r.About
}
