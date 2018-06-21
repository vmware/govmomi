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

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
	id string
}

func init() {
	cli.Register("tags.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
	f.StringVar(&cmd.id, "i", "", "ID for category")

}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List all tags, or list tags for category.

Examples:
  govc tags.ls
  govc tags.ls -i CATEGORYID`
}

type getResult []string

func (r getResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		if cmd.id == "" {
			tagSlice, err := c.ListTags(ctx)
			if err != nil {
				return err
			}
			result := getResult(tagSlice)
			cmd.WriteResult(result)
		} else {
			tagSlice, err := c.ListTagsForCategory(ctx, cmd.id)
			if err != nil {
				return err
			}
			result := getResult(tagSlice)
			cmd.WriteResult(result)
		}
		return nil
	})
}
