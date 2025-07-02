// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vswitch

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.HostSystemFlag

	nic  string
	spec types.HostVirtualSwitchSpec
}

func init() {
	cli.Register("host.vswitch.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.spec.NumPorts = 128 // default
	f.Var(flags.NewInt32(&cmd.spec.NumPorts), "ports", "Number of ports")
	f.Var(flags.NewInt32(&cmd.spec.Mtu), "mtu", "MTU")
	f.StringVar(&cmd.nic, "nic", "", "Bridge nic device")
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	if cmd.nic != "" {
		cmd.spec.Bridge = &types.HostVirtualSwitchBondBridge{
			NicDevice: []string{cmd.nic},
		}
	}

	return ns.AddVirtualSwitch(ctx, f.Arg(0), &cmd.spec)
}
