// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package group

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/sso"
	"github.com/vmware/govmomi/ssoadmin"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("sso.group.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Remove SSO group.

Examples:
  govc sso.group.rm NAME`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		return c.DeletePrincipal(ctx, f.Arg(0))
	})
}
