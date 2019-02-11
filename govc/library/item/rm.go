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

package item

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type rm struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.item.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *rm) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rm) Description() string {
	return `List library items.

Examples:
  govc library.item.rm library_name item_name
  govc library.item.rm library_name item_name -json | jq .`
}

type rmResult []library.Item

func (r rmResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i].Name)
	}
	return nil
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var err error
		var removedItems []library.Item

		if f.NArg() != 2 {
			return flag.ErrHelp
		}

		arg := f.Arg(0)
		itemName := f.Arg(1)
		library, err := m.GetLibraryByName(ctx, arg)
		if err != nil {
			return err
		}

		res, err := m.GetLibraryItems(ctx, library.ID)
		if err != nil {
			return err
		}

		for _, r := range res {
			if r.Name == itemName || r.ID == itemName {
				removedItems = append(removedItems, r)
			}
		}

		if len(removedItems) == 0 {
			fmt.Printf("Library %s item %s not found\n", arg, itemName)
			return nil
		}

		if len(removedItems) != 1 {
			fmt.Printf("Too many items (%s) found in Library %s\n", itemName, arg)
			return nil
		}

		err = m.DeleteLibraryItem(ctx, &removedItems[0])
		if err != nil {
			return err
		}
		fmt.Printf("Removed item %s\n", itemName)

		return nil
	})
}
