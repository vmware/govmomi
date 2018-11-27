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
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.DatacenterFlag
}

func init() {
	cli.Register("library.item.file.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Description() string {
	return `Display library information.

Examples:
  govc library.item.file.info library_name
  govc library.item.file.info library_name item_name
  govc library.item.file.info library_name item_name -json | jq .`
}

type infoResult []library.File

func (t infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, item := range t {
		fmt.Fprintf(tw, "Name:\t%s\n", item.Name)
		fmt.Fprintf(tw, "  Size:\t%d\n", item.Size)
		fmt.Fprintf(tw, "  Version:\t%s\n", item.Version)
		fmt.Fprintf(tw, "  Checksum Info:\n")
		fmt.Fprintf(tw, "    Algorithm:\t%s\n", item.Checksum.Algorithm)
		fmt.Fprintf(tw, "    Checksum:\t%s\n", item.Checksum.Checksum)
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var res infoResult
		var items []library.Item
		var err error

		if f.NArg() != 2 {
			return flag.ErrHelp
		}

		var library *library.Library
		arg := f.Arg(0)
		name := f.Arg(1)
		library, err = m.GetLibraryByName(ctx, arg)
		if err != nil {
			return err
		}
		items, err = m.GetLibraryItems(ctx, library.ID)
		if err != nil {
			return err
		}
		if len(items) == 0 {
			return fmt.Errorf("No library items found")
		}

		for _, item := range items {
			if item.Name == name || item.ID == name {
				res, err = m.ListLibraryItemFiles(ctx, item.ID)
			}

		}

		return cmd.WriteResult(res)
	})
}
