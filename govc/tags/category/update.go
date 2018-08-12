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

package category

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type update struct {
	*flags.ClientFlag
	name        string
	description string
	types       string
	multi       string
}

func init() {
	cli.Register("tags.category.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.name, "n", "", "Name of category")
	f.StringVar(&cmd.description, "d", "", "Description")
	f.StringVar(&cmd.types, "t", "", "Associable object types")
	f.StringVar(&cmd.multi, "m", "", "Allow multiple tags per object")
}

func (cmd *update) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *update) Usage() string {
	return "ID"
}

func (cmd *update) Description() string {
	return `Update category.

Cardinality can be either "SINGLE" or "MULTIPLE."

Examples:
  govc tags.category.update -n "name" -d "description" -t "associable_types" -m "cardinality" ID`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	id := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		category := new(tags.CategoryUpdateSpec)
		categoryTemp := new(tags.Category)
		if cmd.name != "" {
			categoryTemp.Name = cmd.name
		}
		if cmd.description != "" {
			categoryTemp.Description = cmd.description
		}
		if cmd.types != "" {
			typesField := strings.Split(cmd.types, ",")
			categoryTemp.AssociableTypes = typesField
		}
		if cmd.multi != "" {
			categoryTemp.Cardinality = cmd.multi
		}

		category.UpdateSpec = *categoryTemp
		return c.UpdateCategory(ctx, id, category)
	})
}
