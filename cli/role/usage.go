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

type usage struct {
	*permissions.PermissionFlag
}

func init() {
	cli.Register("role.usage", &usage{})
}

func (cmd *usage) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = permissions.NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)
}

func (cmd *usage) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *usage) Usage() string {
	return "NAME..."
}

func (cmd *usage) Description() string {
	return `List usage for role NAME.

Examples:
  govc role.usage
  govc role.usage Admin`
}

func (cmd *usage) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	if f.NArg() == 0 {
		cmd.List.Permissions, err = m.RetrieveAllPermissions(ctx)
		if err != nil {
			return err
		}
	} else {
		for _, name := range f.Args() {
			role, err := cmd.Role(name)
			if err != nil {
				return err
			}

			perms, err := m.RetrieveRolePermissions(ctx, role.RoleId)
			if err != nil {
				return err
			}

			cmd.List.Add(perms)
		}
	}

	return cmd.WriteResult(&cmd.List)
}
