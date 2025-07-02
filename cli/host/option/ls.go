// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package option

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/option"
)

type ls struct {
	*option.List
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.option.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.List = &option.List{}
	cmd.List.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.List.ClientFlag.Register(ctx, f)

	cmd.List.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.List.OutputFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.List.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return option.ListDescription + `

Examples:
  govc host.option.ls
  govc host.option.ls Config.HostAgent.
  govc host.option.ls Config.HostAgent.plugins.solo.enableMob`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	m, err := host.ConfigManager().OptionManager(ctx)
	if err != nil {
		return err
	}

	return cmd.Query(ctx, f, m)
}
