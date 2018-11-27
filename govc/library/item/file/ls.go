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

package file

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

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.item.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `List library items.

Examples:
  govc library.item.ls
  govc library.item.ls library_name
  govc library.item.ls -json | jq .
  govc library.item.ls library_name -json | jq .`
}

type lsResult []library.Item

func (r lsResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i].Name)
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var res lsResult
		var err error

		if f.NArg() != 1 {
			return flag.ErrHelp
		}

		arg := f.Arg(0)
		library, err := m.GetLibraryByName(ctx, arg)
		if err != nil {
			return err
		}

		res, err = m.GetLibraryItems(ctx, library.ID)

		if err != nil {
			return err
		}
		return cmd.WriteResult(res)
	})
}
