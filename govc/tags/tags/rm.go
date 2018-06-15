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

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type rm struct {
	*flags.ClientFlag
	force bool
}

func init() {
	cli.Register("tags.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.BoolVar(&cmd.force, "f", false, "Delete tag regardless of attached objects")
}

func (cmd *rm) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rm) Usage() string {
	return "ID"
}

func (cmd *rm) Description() string {
	return `Delete tag if not attached to any object. Will delete regardless of attached object if flag is set.

Examples:
  govc tags.rm ID
  govc tags.rm -f ID`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	tagID := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		if cmd.force == false {
			return c.DeleteTagIfNoObjectAttached(ctx, tagID)
		}

		return c.DeleteTag(ctx, tagID)

	})

}
