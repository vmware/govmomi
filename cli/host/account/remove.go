// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type remove struct {
	*AccountFlag
}

func init() {
	cli.Register("host.account.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AccountFlag, ctx = newAccountFlag(ctx)
	cmd.AccountFlag.Register(ctx, f)
}

func (cmd *remove) Description() string {
	return `Remove local account on HOST.

Examples:
  govc host.account.remove -id $USER`
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.AccountFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.AccountFlag.HostAccountManager(ctx)
	if err != nil {
		return err
	}
	return m.Remove(ctx, cmd.HostAccountSpec.Id)
}
