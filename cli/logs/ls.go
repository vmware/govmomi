// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type ls struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("logs.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List diagnostic log keys.

Examples:
  govc logs.ls
  govc logs.ls -host host-a`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	var host *object.HostSystem

	if c.IsVC() {
		host, err = cmd.HostSystemIfSpecified()
		if err != nil {
			return err
		}
	}

	m := object.NewDiagnosticManager(c)

	desc, err := m.QueryDescriptions(ctx, host)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, d := range desc {
		fmt.Fprintf(tw, "%s\t%s\n", d.Key, d.FileName)
	}

	return tw.Flush()
}
