// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type cp struct {
	*flags.ClientFlag

	library.Item
}

func init() {
	cli.Register("library.cp", &cp{})
}

func (cmd *cp) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.Name, "n", "", "Library item name")
}

func (cmd *cp) Usage() string {
	return "SRC DST"
}

func (cmd *cp) Description() string {
	return `Copy SRC library item to DST library.
Examples:
  govc library.cp /my-content/my-item /my-other-content
  govc library.cp -n my-item2 /my-content/my-item /my-other-content`
}

func (cmd *cp) Run(ctx context.Context, f *flag.FlagSet) error {
	srcPath := f.Arg(0)
	dstPath := f.Arg(1)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	src, err := flags.ContentLibraryItem(ctx, c, srcPath)
	if err != nil {
		return err
	}

	dst, err := flags.ContentLibrary(ctx, c, dstPath)
	if err != nil {
		return err
	}

	cmd.LibraryID = dst.ID
	if cmd.Name == "" {
		cmd.Name = src.Name
	}

	id, err := m.CopyLibraryItem(ctx, src, cmd.Item)
	if err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}
