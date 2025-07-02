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
	"github.com/vmware/govmomi/ssoadmin/types"
)

type create struct {
	*flags.ClientFlag

	types.AdminGroupDetails
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.AdminGroupDetails.Description, "d", "", "Group description")
}

func init() {
	cli.Register("sso.group.create", &create{})
}

func (cmd *create) Description() string {
	return `Create SSO group.

Examples:
  govc sso.group.create NAME`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		return c.CreateGroup(ctx, f.Arg(0), cmd.AdminGroupDetails)
	})
}
