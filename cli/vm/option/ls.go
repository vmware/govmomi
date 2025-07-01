// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package option

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	flags.EnvBrowser
}

func init() {
	cli.Register("vm.option.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.EnvBrowser.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List VM config option keys for CLUSTER.

Examples:
  govc vm.option.ls -cluster C0`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	b, err := cmd.Browser(ctx)
	if err != nil {
		return err
	}

	opts, err := b.QueryConfigOptionDescriptor(ctx)
	if err != nil {
		return err
	}

	return cmd.VirtualMachineFlag.WriteResult(&lsResult{opts})
}

type lsResult struct {
	opts []types.VirtualMachineConfigOptionDescriptor
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, d := range r.opts {
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", d.Key, d.Description)
	}

	return tw.Flush()
}

func (r *lsResult) Dump() any {
	return r.opts
}

func (r *lsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.opts)
}
