// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"strconv"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type chmod struct {
	*GuestFlag
}

func init() {
	cli.Register("guest.chmod", &chmod{})
}

func (cmd *chmod) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)
}

func (cmd *chmod) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *chmod) Usage() string {
	return "MODE FILE"
}

func (cmd *chmod) Description() string {
	return `Change FILE MODE on VM.

Examples:
  govc guest.chmod -vm $name 0644 /var/log/foo.log`
}

func (cmd *chmod) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	var attr types.GuestPosixFileAttributes

	attr.Permissions, err = strconv.ParseInt(f.Arg(0), 0, 64)
	if err != nil {
		return err
	}

	return m.ChangeFileAttributes(ctx, cmd.Auth(), f.Arg(1), &attr)
}
