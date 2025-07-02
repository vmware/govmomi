// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/namespace"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type infoResult struct {
	namespace.NamespacesInstanceInfo

	ctx context.Context
	cmd *info
}

func (r infoResult) Write(w io.Writer) error {
	c, err := r.cmd.Client()
	if err != nil {
		return err
	}

	rc, err := r.cmd.RestClient()
	if err != nil {
		return err
	}
	l := library.NewManager(rc)

	pc, err := r.cmd.PbmClient()
	if err != nil {
		return err
	}

	var ids []string
	for _, s := range r.StorageSpecs {
		ids = append(ids, s.Policy)
	}
	m, err := pc.ProfileMap(r.ctx, ids...)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	cluster := types.ManagedObjectReference{Type: "ClusterComputeResource", Value: r.ClusterId}
	path, err := find.InventoryPath(r.ctx, c, cluster)
	if err != nil {
		return err
	}

	fmt.Fprintf(tw, "Cluster:\t%s\n", path)
	fmt.Fprintf(tw, "Status:\t%s\n", r.ConfigStatus)
	fmt.Fprintf(tw, "VM Classes:\t%s\n", strings.Join(r.VmServiceSpec.VmClasses, ","))

	for _, s := range r.VmServiceSpec.ContentLibraries {
		info, err := l.GetLibraryByID(r.ctx, s)
		if err != nil {
			return err
		}
		fmt.Fprintf(tw, "Content Library:\t%s\n", info.Name)
	}

	for _, s := range r.StorageSpecs {
		fmt.Fprintf(tw, "Storage Policy:\t%s\n", m.Name[s.Policy].GetPbmProfile().Name)
	}

	return tw.Flush()
}

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("namespace.info", &info{})
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
	return "NAME"
}

func (cmd *info) Description() string {
	return `Displays the details of a vSphere Namespace.

Examples:
  govc namespace.info test-namespace`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	nm := namespace.NewManager(rc)

	d, err := nm.GetNamespace(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	return cmd.WriteResult(&infoResult{d, ctx, cmd})
}
