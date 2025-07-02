// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package shutdown

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/shutdown"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.shutdown.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *get) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *get) Description() string {
	return `Get details about the pending shutdown action.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.get`
}

type config struct {
	shutdown.Config
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := shutdown.NewManager(c)

	s, err := m.Get(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(config{
		s,
	})
}

func (c config) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Action:%s\n", c.Action)
	fmt.Fprintf(tw, "Reason:%s\n", c.Reason)
	fmt.Fprintf(tw, "ShutDown Time:%s\n", c.ShutdownTime)

	return tw.Flush()
}
