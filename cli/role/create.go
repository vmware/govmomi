// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package role

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/permissions"
)

type create struct {
	*permissions.PermissionFlag
}

func init() {
	cli.Register("role.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = permissions.NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Usage() string {
	return "NAME [PRIVILEGE]..."
}

func (cmd *create) Description() string {
	return `Create authorization role.

Optionally populate the role with the given PRIVILEGE(s).

Examples:
  govc role.create MyRole
  govc role.create NoDC $(govc role.ls Admin | grep -v Datacenter.)`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	_, err = m.AddRole(ctx, f.Arg(0), f.Args()[1:])
	return err
}
