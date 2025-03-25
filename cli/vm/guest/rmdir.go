// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type rmdir struct {
	*GuestFlag

	recursive bool
}

func init() {
	cli.Register("guest.rmdir", &rmdir{})
}

func (cmd *rmdir) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.recursive, "r", false, "Recursive removal")
}

func (cmd *rmdir) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rmdir) Usage() string {
	return "PATH"
}

func (cmd *rmdir) Description() string {
	return `Remove directory PATH in VM.

Examples:
  govc guest.rmdir -vm $name /tmp/empty-dir
  govc guest.rmdir -vm $name -r /tmp/non-empty-dir`
}

func (cmd *rmdir) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	return m.DeleteDirectory(ctx, cmd.Auth(), f.Arg(0), cmd.recursive)
}
