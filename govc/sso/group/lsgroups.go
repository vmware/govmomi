/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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
	"fmt"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/sso"
	"github.com/vmware/govmomi/ssoadmin"
)

type id struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("sso.group.lsgroups", &id{})
}

func (cmd *id) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *id) Usage() string {
	return "Parent Group NAME"
}

func (cmd *id) Description() string {
	return `List SSO groups, which are members of a local group

Examples:
  govc sso.group.lsgroups
  govc sso.group.lsgroups syncusers
  govc sso.group.lsgroups -json syncusers`
}

func (cmd *id) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *id) Run(ctx context.Context, f *flag.FlagSet) error {
	arg := f.Arg(0)
	if arg == "" {
		arg = "users@vsphere.local"
	}
	search := f.Arg(1)

	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		group, err := c.FindGroup(ctx, arg)
		if err != nil {
			return err
		}
		if group == nil {
			return fmt.Errorf("group %q not found", arg)
		}

		info, err := c.FindGroupsInGroup(ctx, arg, search)
		if err != nil {
			return err
		}

		return cmd.WriteResult(groupResult(info))
	})
}
