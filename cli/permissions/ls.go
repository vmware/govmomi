// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package permissions

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type ls struct {
	*PermissionFlag

	inherited bool
}

func init() {
	cli.Register("permissions.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)

	f.BoolVar(&cmd.inherited, "a", true, "Include inherited permissions defined by parent entities")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return "[PATH]..."
}

func (cmd *ls) Description() string {
	return `List the permissions defined on or effective on managed entities.

Examples:
  govc permissions.ls
  govc permissions.ls /dc1/host/cluster1`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	refs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	for _, ref := range refs {
		perms, err := m.RetrieveEntityPermissions(ctx, ref, cmd.inherited)
		if err != nil {
			return err
		}

		cmd.List.Add(perms)
	}

	return cmd.WriteResult(&cmd.List)
}
