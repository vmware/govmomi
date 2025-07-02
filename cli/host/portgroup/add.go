// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package portgroup

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.HostSystemFlag

	spec types.HostPortGroupSpec
}

func init() {
	cli.Register("host.portgroup.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.StringVar(&cmd.spec.VswitchName, "vswitch", "", "vSwitch Name")
	f.Var(flags.NewInt32(&cmd.spec.VlanId), "vlan", "VLAN ID")
}

func (cmd *add) Description() string {
	return `Add portgroup to HOST.

Examples:
  govc host.portgroup.add -vswitch vSwitch0 -vlan 3201 bridge`
}

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	cmd.spec.Name = f.Arg(0)

	return ns.AddPortGroup(ctx, cmd.spec)
}
