// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package profile

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
)

type ls struct {
	*flags.ClientFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("gpu.host.profile.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List available vGPU profiles on host.

Examples:
  govc gpu.host.profile.ls -host hostname
  govc gpu.host.profile.ls -host hostname -json | jq -r '.profiles[]'
  govc gpu.host.profile.ls -host hostname -json | jq -r '.profiles[] | select(contains("nvidia_a40"))'`
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type lsResult struct {
	Profiles []string `json:"profiles"`
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "Available vGPU profiles:")
	for _, profile := range r.Profiles {
		fmt.Fprintf(tw, "  %s\n", profile)
	}
	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	if host == nil {
		return flag.ErrHelp
	}

	var o mo.HostSystem
	pc := property.DefaultCollector(host.Client())
	err = pc.RetrieveOne(ctx, host.Reference(), []string{"config.sharedPassthruGpuTypes"}, &o)
	if err != nil {
		return err
	}

	if o.Config == nil {
		return fmt.Errorf("failed to get host configuration")
	}

	if len(o.Config.SharedPassthruGpuTypes) == 0 {
		return fmt.Errorf("no vGPU profiles available on this host")
	}

	res := &lsResult{
		Profiles: o.Config.SharedPassthruGpuTypes,
	}

	return cmd.WriteResult(res)
}
