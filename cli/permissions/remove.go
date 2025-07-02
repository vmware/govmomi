// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package permissions

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*PermissionFlag

	types.Permission
	force bool
}

func init() {
	cli.Register("permissions.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)

	f.StringVar(&cmd.Principal, "principal", "", "User or group for which the permission is defined")
	f.BoolVar(&cmd.Group, "group", false, "True, if principal refers to a group name; false, for a user name")
	f.BoolVar(&cmd.force, "f", false, "Ignore NotFound fault if permission for this entity and user or group does not exist")
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "[PATH]..."
}

func (cmd *remove) Description() string {
	return `Removes a permission rule from managed entities.

Examples:
  govc permissions.remove -principal root
  govc permissions.remove -principal $USER@vsphere.local /dc1/host/cluster1`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	refs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	for _, ref := range refs {
		err = m.RemoveEntityPermission(ctx, ref, cmd.Principal, cmd.Group)
		if err != nil {
			if cmd.force {
				if fault.Is(err, &types.NotFound{}) {
					continue
				}
			}
			return err
		}
	}

	return nil
}
