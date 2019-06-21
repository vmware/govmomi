/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package group

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/sso"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/types"
)

type update struct {
	*flags.ClientFlag

	d string
	a string
	r string
}

func init() {
	cli.Register("sso.group.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.d, "d", "", "Group description")
	f.StringVar(&cmd.a, "a", "", "Add user to group")
	f.StringVar(&cmd.r, "r", "", "Remove user from group")
}

func (cmd *update) Description() string {
	return `Update SSO group.

Examples:
  govc sso.group.update -d "Group description" NAME
  govc sso.group.update -a user1 NAME
  govc sso.group.update -r user2 NAME`
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
			user, err := c.FindUser(ctx, cmd.a)
			if err != nil {
				return err
			}
			if err = c.AddUsersToGroup(ctx, id, user.Id); err != nil {
				return err
			}
		}

		if cmd.r != "" {
			user, err := c.FindUser(ctx, cmd.r)
			if err != nil {
				return err
			}
			if err = c.RemoveUsersFromGroup(ctx, id, user.Id); err != nil {
				return err
			}
		}

		return nil
	})
}
