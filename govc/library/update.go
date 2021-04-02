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

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type update struct {
	*flags.ClientFlag

	library.Library
}

func init() {
	cli.Register("library.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.Library.Name, "n", "", "Library name")
	f.StringVar(&cmd.Library.Description, "d", "", "Library description")
}

func (cmd *update) Usage() string {
	return "NAME"
}

func (cmd *update) Description() string {
	return `Update library.

Examples:
  govc library.update -d "new description" -n "new-name" current-name`
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
	dst, err := flags.ContentLibrary(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}

	cmd.ID = dst.ID
	dst.Patch(&cmd.Library)

	return m.UpdateLibrary(ctx, dst)
}
