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
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type get struct {
	*flags.ClientFlag
	*flags.OutputFlag
	name string
}

func init() {
	cli.Register("tags.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
	f.StringVar(&cmd.name, "n", "", "Name of category")
}

func (cmd *get) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}
func (cmd *get) Usage() string {
	return "TAGID"
}

func (cmd *get) Description() string {
	return `Get tags by tags' ID, or Get tags for category by tag name and category ID. 

Examples:
  govc tags.get TAGID
  govc tags.get -n TAGNAME CATEGORYID `
}

type getTagName []tags.Tag

func (t getTagName) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, item := range t {
		fmt.Fprintf(tw, "Name: %s\nID: %s\nDescription: %s\nCategoryID: %s\nUsedBy: %s\n",
			item.Name, item.ID, item.Description, item.CategoryID, item.UsedBy)
	}

	return tw.Flush()
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	id := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		if cmd.name == "" {
			tag, err := c.GetTag(ctx, id)
			if err != nil {
				return err
			}
			fmt.Printf("Name: %s\nID: %s\nDescription: %s\nCategoryID: %s\nUsedBy: %s\n",
				tag.Name, tag.ID, tag.Description, tag.CategoryID, tag.UsedBy)
		} else {
			tagSlice, err := c.GetTagByNameForCategory(ctx, cmd.name, id)
			if err != nil {
				return err
			}
			result := getTagName(tagSlice)
			cmd.WriteResult(result)
		}
		return nil

	})
}
