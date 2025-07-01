// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type rm struct {
	*flags.ClientFlag

	cancel bool
	files  bool
}

func init() {
	cli.Register("library.session.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.cancel, "f", false, "Cancel session if active")
	f.BoolVar(&cmd.files, "i", false, "Remove session item file")
}

func (cmd *rm) Description() string {
	return `Remove a library item update session.

Examples:
  govc library.session.rm session_id
  govc library.session.rm -i session_id foo.ovf`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	cancel := m.CancelLibraryItemUpdateSession
	remove := m.DeleteLibraryItemUpdateSession
	rmfile := m.RemoveLibraryItemUpdateSessionFile

	_, err = m.GetLibraryItemUpdateSession(ctx, id)
	if err != nil {
		cancel = m.CancelLibraryItemDownloadSession
		remove = m.DeleteLibraryItemDownloadSession
		rmfile = func(context.Context, string, string) error {
			return errors.New("cannot delete a download session file")
		}
	}

	if cmd.files {
		return rmfile(ctx, id, f.Arg(1))
	}

	if cmd.cancel {
		err := cancel(ctx, id)
		if err != nil {
			return nil
		}
	}
	return remove(ctx, id)
}
