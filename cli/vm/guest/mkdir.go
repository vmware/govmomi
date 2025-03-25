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

type mkdir struct {
	*GuestFlag

	createParents bool
}

func init() {
	cli.Register("guest.mkdir", &mkdir{})
}

func (cmd *mkdir) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.createParents, "p", false, "Create intermediate directories as needed")
}

func (cmd *mkdir) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *mkdir) Usage() string {
	return "PATH"
}

func (cmd *mkdir) Description() string {
	return `Create directory PATH in VM.

Examples:
  govc guest.mkdir -vm $name /tmp/logs
  govc guest.mkdir -vm $name -p /tmp/logs/foo/bar`
}

func (cmd *mkdir) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	err = m.MakeDirectory(ctx, cmd.Auth(), f.Arg(0), cmd.createParents)

	// ignore EEXIST if -p flag is given
	if err != nil && cmd.createParents {
		if fault.Is(err, &types.FileAlreadyExists{}) {
			return nil
		}
	}

	return err
}
