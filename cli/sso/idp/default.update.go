// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package idp

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/sso"
	"github.com/vmware/govmomi/ssoadmin"
)

type didpupd struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("sso.idp.default.update", &didpupd{})
}

func (cmd *didpupd) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *didpupd) Usage() string {
	return "NAME"
}

func (cmd *didpupd) Description() string {
	return `Set SSO default identity provider source.

Examples:
  govc sso.idp.default.update NAME`
}

func (cmd *didpupd) Run(ctx context.Context, f *flag.FlagSet) error {
	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		return c.SetDefaultDomains(ctx, f.Arg(0))
	})
}
