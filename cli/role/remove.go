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

type remove struct {
	*permissions.PermissionFlag

	force bool
}

func init() {
	cli.Register("role.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = permissions.NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)

	f.BoolVar(&cmd.force, "force", false, "Force removal if role is in use")
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "NAME"
}

func (cmd *remove) Description() string {
	return `Remove authorization role.

Examples:
  govc role.remove MyRole
  govc role.remove MyRole -force`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	role, err := cmd.Role(f.Arg(0))
	if err != nil {
		return err
	}

	return m.RemoveRole(ctx, role.RoleId, !cmd.force)
}
