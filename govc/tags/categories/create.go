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
	types       string
	multi       bool
}

func init() {
	cli.Register("tags.category.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.description, "d", "", "description of category")
	f.StringVar(&cmd.types, "t", "", "associable_types of category")
	f.BoolVar(&cmd.multi, "m", false, "cardinality of category")
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Usage() string {
	return ` Create tag category. This command will output the ID you just created.

Examples:
  govc tags.category.create -d "description" -t "categoryType" -m(optional, if set, the cardinality will be true) NAME`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	name := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		id, err := c.CreateCategoryIfNotExist(ctx, name, cmd.description, cmd.types, cmd.multi)
		if err != nil {
			return err
		}

		fmt.Println(*id)

		return nil

	})
}
