// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package version

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
	cli.Register("namespace.service.version.ls", &ls{})
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
	return `List all registered versions for a given vSphere Supervisor Service.

Examples:
  govc namespace.service.version.ls my-service
  govc namespace.service.version.ls -l my-service
  govc namespace.service.version.ls -json my-service | jq .`
}

func (cmd *ls) Usage() string {
	return "NAME"
}

type lsWriter struct {
	cmd      *ls
	service  string
	Versions []namespace.SupervisorServiceVersionSummary `json:"versions"`
}

func (r *lsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "%s:\n", r.service)
	for _, svc := range r.Versions {
		fmt.Fprintf(tw, "%s", svc.Version)
		if r.cmd.long {
			fmt.Fprintf(tw, "\t%s", svc.Name)
			fmt.Fprintf(tw, "\t%s", svc.State)
			fmt.Fprintf(tw, "\t%s", svc.Description)
		}
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}

func (r *lsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Versions)
}

func (r *lsWriter) Dump() any {
	return r.Versions
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	service := f.Arg(0)
	if len(service) == 0 {
		return flag.ErrHelp
	}
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	versions, err := m.ListSupervisorServiceVersions(ctx, service)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsWriter{cmd, service, versions})
}
