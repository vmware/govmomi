// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	long bool
}

func init() {
	cli.Register("namespace.service.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List all registered vSphere Supervisor Services.

Examples:
  govc namespace.service.ls
  govc namespace.service.ls -l
  govc namespace.service.ls -json | jq .`
}

type lsWriter struct {
	cmd     *ls
	Service []namespace.SupervisorServiceSummary `json:"service"`
}

func (r *lsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, svc := range r.Service {
		fmt.Fprintf(tw, "%s", svc.ID)
		if r.cmd.long {
			fmt.Fprintf(tw, "\t%s", svc.State)
			fmt.Fprintf(tw, "\t%s", svc.Name)
		}
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}

func (r *lsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Service)
}

func (r *lsWriter) Dump() any {
	return r.Service
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	supervisorservices, err := m.ListSupervisorServices(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsWriter{cmd, supervisorservices})
}
