// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package group

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/sso"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/types"
)

type update struct {
	*flags.ClientFlag

	d string
	a string
	r string
	g bool
}

func init() {
	cli.Register("sso.group.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.d, "d", "", "Group description")
	f.StringVar(&cmd.a, "a", "", "Add user/group to group")
	f.StringVar(&cmd.r, "r", "", "Remove user/group from group")
	f.BoolVar(&cmd.g, "g", false, "Add/Remove group from group")
}

func (cmd *update) Description() string {
	return `Update SSO group.

Examples:
  govc sso.group.update -d "Group description" NAME
  govc sso.group.update -a user1 NAME
  govc sso.group.update -r user2 NAME
  govc sso.group.update -g -a group1 NAME
  govc sso.group.update -g -r group2 NAME`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	id := f.Arg(0)

	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		if cmd.d != "" {
			err := c.UpdateGroup(ctx, id, types.AdminGroupDetails{Description: cmd.d})
			if err != nil {
				return err
			}
		}

		if cmd.a != "" {
			if cmd.g {
				group, err := c.FindGroup(ctx, cmd.a)
				if err != nil {
					return err
				}
				if group == nil {
					return fmt.Errorf("group %q not found", cmd.a)
				}
				if err = c.AddGroupsToGroup(ctx, id, group.Id); err != nil {
					return err
				}
			} else {
				user, err := c.FindUser(ctx, cmd.a)
				if err != nil {
					return err
				}
				if user == nil {
					return fmt.Errorf("user %q not found", cmd.a)
				}
				if err = c.AddUsersToGroup(ctx, id, user.Id); err != nil {
					return err
				}
			}
		}

		if cmd.r != "" {
			var pid types.PrincipalId
			if cmd.g {
				group, err := c.FindGroup(ctx, cmd.r)
				if err != nil {
					return err
				}
				if group == nil {
					return fmt.Errorf("group %q not found", cmd.r)
				}
				pid = group.Id
			} else {
				user, err := c.FindUser(ctx, cmd.r)
				if err != nil {
					return err
				}
				if user == nil {
					return fmt.Errorf("user %q not found", cmd.r)
				}
				pid = user.Id
			}

			if err := c.RemoveUsersFromGroup(ctx, id, pid); err != nil {
				return err
			}
		}

		return nil
	})
}
