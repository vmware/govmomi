// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/vim25/types"
)

type mv struct {
	*GuestFlag

	noclobber bool
}

func init() {
	cli.Register("guest.mv", &mv{})
}

func (cmd *mv) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.noclobber, "n", false, "Do not overwrite an existing file")
}

func (cmd *mv) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *mv) Usage() string {
	return "SOURCE DEST"
}

func (cmd *mv) Description() string {
	return `Move (rename) files in VM.

Examples:
  govc guest.mv -vm $name /tmp/foo.sh /tmp/bar.sh
  govc guest.mv -vm $name -n /tmp/baz.sh /tmp/bar.sh`
}

func (cmd *mv) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	src := f.Arg(0)
	dst := f.Arg(1)

	err = m.MoveFile(ctx, cmd.Auth(), src, dst, !cmd.noclobber)

	if err != nil {
		if fault.Is(err, &types.NotAFile{}) {
			err = m.MoveDirectory(ctx, cmd.Auth(), src, dst)
		}
	}

	return err
}
