// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cluster"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
	moduleID string
}

func init() {
	cli.Register("cluster.module.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.moduleID, "id", "", "Module ID")
}

func (cmd *ls) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List cluster modules.

When -id is specified, that module's members are listed.

Examples:
  govc cluster.module.ls
  govc cluster.module.ls -json | jq .
  govc cluster.module.ls -id module_id`
}

type lsResult []cluster.ModuleSummary

func (r lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, i := range r {
		fmt.Fprintf(tw, "%s\t%s\n", i.Cluster, i.Module)
	}

	return tw.Flush()
}

func (cmd *ls) List(ctx context.Context, m *cluster.Manager) error {
	var res lsResult

	res, err := m.ListModules(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(res)
}

type lsMemberResult []types.ManagedObjectReference

func (r lsMemberResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, i := range r {
		fmt.Fprintf(tw, "%s\t%s\n", i.Reference().Type, i.Reference().Value)
	}

	return tw.Flush()
}

func (cmd *ls) ListMembers(ctx context.Context, m *cluster.Manager, moduleID string) error {
	var res lsMemberResult

	res, err := m.ListModuleMembers(ctx, moduleID)
	if err != nil {
		return err
	}

	return cmd.WriteResult(res)
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 0 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := cluster.NewManager(c)

	if cmd.moduleID == "" {
		return cmd.List(ctx, m)
	}

	return cmd.ListMembers(ctx, m, cmd.moduleID)
}
