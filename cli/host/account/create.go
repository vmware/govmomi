// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type create struct {
	*AccountFlag
}

func init() {
	cli.Register("host.account.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AccountFlag, ctx = newAccountFlag(ctx)
	cmd.AccountFlag.Register(ctx, f)
}

func (cmd *create) Description() string {
	return `Create local account on HOST.

Examples:
  govc host.account.create -id $USER -password password-for-esx60`
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.AccountFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.AccountFlag.HostAccountManager(ctx)
	if err != nil {
		return err
	}
	return m.Create(ctx, &cmd.HostAccountSpec)
}
