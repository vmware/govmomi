// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	vlogging "github.com/vmware/govmomi/vapi/appliance/logging"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vcsa.log.forwarding.info", &info{})
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

func (cmd *info) Description() string {
	return `Retrieve the VC Appliance log forwarding configuration

Examples:
  govc vcsa.log.forwarding.info`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	fwd := vlogging.NewManager(c)

	res, err := fwd.Forwarding(ctx)
	if err != nil {
		return nil
	}

	return cmd.WriteResult(forwardingConfigResult(res))
}

type forwardingConfigResult []vlogging.Forwarding

func (res forwardingConfigResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, c := range res {
		fmt.Fprintf(tw, "Hostname:\t%s\n", c.Hostname)
		fmt.Fprintf(tw, "Port:\t%d\n", c.Port)
		fmt.Fprintf(tw, "Protocol:\t%s\n", c.Protocol)
	}

	return tw.Flush()
}
