// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package permissions

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type set struct {
	*PermissionFlag

	types.Permission

	role string
}

func init() {
	cli.Register("permissions.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)

	f.StringVar(&cmd.Principal, "principal", "", "User or group for which the permission is defined")
	f.BoolVar(&cmd.Group, "group", false, "True, if principal refers to a group name; false, for a user name")
	f.BoolVar(&cmd.Propagate, "propagate", true, "Whether or not this permission propagates down the hierarchy to sub-entities")
	f.StringVar(&cmd.role, "role", "Admin", "Permission role name")
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *set) Usage() string {
	return "[PATH]..."
}

func (cmd *set) Description() string {
	return `Set the permissions managed entities.

Examples:
  govc permissions.set -principal root -role Admin
  govc permissions.set -principal $USER@vsphere.local -role Admin /dc1/host/cluster1`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	refs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	role, err := cmd.Role(cmd.role)
	if err != nil {
		return err
	}

	cmd.Permission.RoleId = role.RoleId

	perms := []types.Permission{cmd.Permission}

	for _, ref := range refs {
		err = m.SetEntityPermissions(ctx, ref, perms)
		if err != nil {
			return err
		}
	}

	return nil
}
