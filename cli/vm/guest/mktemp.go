// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
)

type mktemp struct {
	*GuestFlag

	dir    bool
	path   string
	prefix string
	suffix string
}

func init() {
	cli.Register("guest.mktemp", &mktemp{})
}

func (cmd *mktemp) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.dir, "d", false, "Make a directory instead of a file")
	f.StringVar(&cmd.path, "p", "", "If specified, create relative to this directory")
	f.StringVar(&cmd.prefix, "t", "", "Prefix")
	f.StringVar(&cmd.suffix, "s", "", "Suffix")
}

func (cmd *mktemp) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *mktemp) Description() string {
	return `Create a temporary file or directory in VM.

Examples:
  govc guest.mktemp -vm $name
  govc guest.mktemp -vm $name -d
  govc guest.mktemp -vm $name -t myprefix
  govc guest.mktemp -vm $name -p /var/tmp/$USER`
}

func (cmd *mktemp) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	mk := m.CreateTemporaryFile
	if cmd.dir {
		mk = m.CreateTemporaryDirectory
	}

	name, err := mk(ctx, cmd.Auth(), cmd.prefix, cmd.suffix, cmd.path)
	if err != nil {
		return err
	}

	fmt.Println(name)

	return nil
}
