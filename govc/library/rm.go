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

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
)

type rm struct {
	*flags.ClientFlag
	force bool
}

func init() {
	cli.Register("library.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Delete library or item NAME.

Examples:
  govc library.rm /library_name
  govc library.rm /library_name/item_name`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)

		res, err := finder.NewFinder(m).Find(ctx, f.Arg(0))
		if err != nil {
			return err
		}

		if len(res) != 1 {
			return fmt.Errorf("%q matches %d items", f.Arg(0), len(res))
		}

		switch t := res[0].GetResult().(type) {
		case library.Library:
			return m.DeleteLibrary(ctx, &t)
		case library.Item:
			return m.DeleteLibraryItem(ctx, &t)
		default:
			return fmt.Errorf("%q is a %T", f.Arg(0), t)
		}
	})
}
