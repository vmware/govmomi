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

type set struct {
	*option.Set
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.option.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.Set = &option.Set{}
	cmd.Set.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.Set.ClientFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.Set.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *set) Description() string {
	return option.SetDescription + `

Examples:
  govc host.option.set Config.HostAgent.plugins.solo.enableMob true
  govc host.option.set Config.HostAgent.log.level verbose
  govc host.option.set Config.HostAgent.vmacore.soap.sessionTimeout 90`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	m, err := host.ConfigManager().OptionManager(ctx)
	if err != nil {
		return err
	}

	return cmd.Update(ctx, f, m)
}
