// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type update struct {
	*AccountFlag
}

func init() {
	cli.Register("host.account.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AccountFlag, ctx = newAccountFlag(ctx)
	cmd.AccountFlag.Register(ctx, f)
}

func (cmd *update) Description() string {
	return `Update local account on HOST.

Examples:
  govc host.account.update -id root -password password-for-esx60`
}

func (cmd *update) Process(ctx context.Context) error {
	if err := cmd.AccountFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.AccountFlag.HostAccountManager(ctx)
	if err != nil {
		return err
	}
	return m.Update(ctx, &cmd.HostAccountSpec)
}
