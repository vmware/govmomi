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

package tags

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type create struct {
	*flags.ClientFlag
	description string
}

func init() {
	cli.Register("tags.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.description, "d", "", "Description of tag")
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Usage() string {
	return "NAME CATEGORYID"
}

func (cmd *create) Description() string {
	return ` Create tag. This command will output the ID you just created.

Examples:
  govc tags.create -d "description" NAME CATEGORYID`
}

func withClient(ctx context.Context, cmd *flags.ClientFlag, f func(*tags.RestClient) error) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}
	tagsURL := vc.URL()
	tagsURL.User = cmd.Userinfo()

	c := tags.NewClient(tagsURL, !cmd.IsSecure(), "")
	if err != nil {
		return err
	}

	if err = c.Login(ctx); err != nil {
		return err
	}
	defer c.Logout(ctx)

	return f(c)
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	name := f.Arg(0)
	id := f.Arg(1)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		id, err := c.CreateTagIfNotExist(ctx, name, cmd.description, id)
		if err != nil {
			return err
		}
		fmt.Println(*id)
		return nil

	})
}
