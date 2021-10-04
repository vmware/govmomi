/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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
)

type update struct {
	*flags.ClientFlag

	name, desc string
}

func init() {
	cli.Register("library.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.name, "n", "", "Library or item name")
	f.StringVar(&cmd.desc, "d", "", "Library or item description")
}

func (cmd *update) Usage() string {
	return "PATH"
}

func (cmd *update) Description() string {
	return `Update library or item PATH.

Examples:
  govc library.update -d "new library description" -n "new-name" my-library
  govc library.update -d "new item description" -n "new-item-name" my-library/my-item`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)

	res, err := flags.ContentLibraryResult(ctx, c, "", f.Arg(0))
	if err != nil {
		return err
	}

	switch t := res.GetResult().(type) {
	case library.Library:
		lib := &library.Library{
			ID:          t.ID,
			Name:        cmd.name,
			Description: cmd.desc,
		}
		t.Patch(lib)
		return m.UpdateLibrary(ctx, &t)
	case library.Item:
		item := &library.Item{
			ID:          t.ID,
			Name:        cmd.name,
			Description: cmd.desc,
		}
		t.Patch(item)
		return m.UpdateLibraryItem(ctx, item)
	default:
		return fmt.Errorf("%q is a %T", f.Arg(0), t)
	}
}
