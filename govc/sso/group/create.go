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

type create struct {
	*flags.ClientFlag

	types.AdminGroupDetails
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.AdminGroupDetails.Description, "d", "", "Group description")
}

func init() {
	cli.Register("sso.group.create", &create{})
}

func (cmd *create) Description() string {
	return `Create SSO group.

Examples:
  govc sso.group.create NAME`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		return c.CreateGroup(ctx, f.Arg(0), cmd.AdminGroupDetails)
	})
}
