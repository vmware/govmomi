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
	"fmt"
	"io"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

type getName []tags.Category

func init() {
	cli.Register("tags.category.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *get) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *get) Usage() string {
	return "NAME"
}

func (cmd *get) Description() string {
	return `Get category by name. 
	
Will return empty if category name doesn't exist.

Examples:
  govc tags.category.get NAME`
}

func (r getName) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	name := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {
		categoriesName, err := c.GetCategoriesByName(ctx, name)
		if err != nil {
			return err
		}
		result := getName(categoriesName)
		cmd.WriteResult(result)
		return nil
	})
}
