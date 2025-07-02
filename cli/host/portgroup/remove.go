// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package portgroup

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type remove struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.portgroup.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *remove) Description() string {
	return `Remove portgroup from HOST.

Examples:
  govc host.portgroup.remove bridge`
}

func (cmd *remove) Usage() string {
	return "NAME"
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	return ns.RemovePortGroup(ctx, f.Arg(0))
}
