/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package session

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type rm struct {
	*flags.ClientFlag

	cancel bool
}

func init() {
	cli.Register("library.session.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.cancel, "f", false, "Cancel session if active")
}

func (cmd *rm) Description() string {
	return `Remove a library item update session.

Examples:
  govc library.session.rm session_id`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		cancel := m.CancelLibraryItemUpdateSession
		remove := m.DeleteLibraryItemUpdateSession

		_, err := m.GetLibraryItemUpdateSession(ctx, id)
		if err != nil {
			cancel = m.CancelLibraryItemDownloadSession
			remove = m.DeleteLibraryItemDownloadSession
		}

		if cmd.cancel {
			err := cancel(ctx, id)
			if err != nil {
				return nil
			}
		}
		return remove(ctx, id)
	})
}
