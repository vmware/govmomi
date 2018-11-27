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

package update

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type create struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("library.item.update.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Description() string {
	return `Create a library item update session.

Examples:
  govc library.item.update.create library_id
  govc library.item.update.create library_id -json | jq .`
}

type createResult []library.UpdateSession

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var spec library.UpdateSession
		var err error

		if f.NArg() != 1 {
			return flag.ErrHelp
		}

		spec.CreateSpec.LibraryItemID = f.Arg(0)
		sessionID, err := m.CreateLibraryItemUpdateSession(ctx, spec)
		if err != nil {
			return err
		}

		fmt.Println(sessionID)

		return nil
	})
}
