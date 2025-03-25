// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type chown struct {
	*GuestFlag
}

func init() {
	cli.Register("guest.chown", &chown{})
}

func (cmd *chown) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)
}

func (cmd *chown) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *chown) Usage() string {
	return "UID[:GID] FILE"
}

func (cmd *chown) Description() string {
	return `Change FILE UID and GID on VM.

Examples:
  govc guest.chown -vm $name UID[:GID] /var/log/foo.log`
}

func (cmd *chown) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	var attr types.GuestPosixFileAttributes

	ids := strings.SplitN(f.Arg(0), ":", 2)
	if len(ids) == 0 {
		return flag.ErrHelp
	}

	ownerIDStr := ids[0]
	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		return err
	}

	attr.OwnerId = new(int32)
	*attr.OwnerId = int32(ownerID)

	if len(ids) == 2 {
		groupIDStr := ids[1]
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			return err
		}

		attr.GroupId = new(int32)
		*attr.GroupId = int32(groupID)
	}

	return m.ChangeFileAttributes(ctx, cmd.Auth(), f.Arg(1), &attr)
}
