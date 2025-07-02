// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package role

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/permissions"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*permissions.PermissionFlag
}

func init() {
	cli.Register("role.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = permissions.NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.PermissionFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return "[NAME]"
}

func (cmd *ls) Description() string {
	return `List authorization roles.

If NAME is provided, list privileges for the role.

Examples:
  govc role.ls
  govc role.ls Admin`
}

type lsRoleList object.AuthorizationRoleList

func (rl lsRoleList) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, role := range rl {
		fmt.Fprintf(tw, "%s\t%s\n", role.Name, role.Info.GetDescription().Summary)
	}

	return tw.Flush()
}

type lsRole types.AuthorizationRole

func (r lsRole) Write(w io.Writer) error {
	for _, p := range r.Privilege {
		fmt.Println(p)
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 1 {
		return flag.ErrHelp
	}

	_, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	if f.NArg() == 1 {
		role, err := cmd.Role(f.Arg(0))
		if err != nil {
			return err
		}

		return cmd.WriteResult(lsRole(*role))
	}

	return cmd.WriteResult(lsRoleList(cmd.Roles))
}
