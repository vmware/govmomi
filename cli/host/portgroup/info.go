// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package portgroup

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/host/vswitch"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.portgroup.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func networkInfoPortgroup(ctx context.Context, c *flags.ClientFlag, h *flags.HostSystemFlag) ([]types.HostPortGroup, error) {
	client, err := c.Client()
	if err != nil {
		return nil, err
	}

	ns, err := h.HostNetworkSystem()
	if err != nil {
		return nil, err
	}

	var mns mo.HostNetworkSystem

	pc := property.DefaultCollector(client)
	err = pc.RetrieveOne(ctx, ns.Reference(), []string{"networkInfo.portgroup"}, &mns)
	if err != nil {
		return nil, err
	}

	return mns.NetworkInfo.Portgroup, nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	pg, err := networkInfoPortgroup(ctx, cmd.ClientFlag, cmd.HostSystemFlag)
	if err != nil {
		return err
	}

	r := &infoResult{pg}

	return cmd.WriteResult(r)
}

type infoResult struct {
	Portgroup []types.HostPortGroup `json:"portgroup"`
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for i, s := range r.Portgroup {
		if i > 0 {
			fmt.Fprintln(tw)
		}
		fmt.Fprintf(tw, "Name:\t%s\n", s.Spec.Name)
		fmt.Fprintf(tw, "Virtual switch:\t%s\n", s.Spec.VswitchName)
		fmt.Fprintf(tw, "VLAN ID:\t%d\n", s.Spec.VlanId)
		fmt.Fprintf(tw, "Active ports:\t%d\n", len(s.Port))
		vswitch.HostNetworkPolicy(tw, &s.ComputedPolicy)
	}

	return tw.Flush()
}
