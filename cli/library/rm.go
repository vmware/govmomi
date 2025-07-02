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

type rm struct {
	*flags.ClientFlag
	force bool
}

func init() {
	cli.Register("library.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Delete library or item NAME.

Examples:
  govc library.rm /library_name
  govc library.rm library_id # Use library id if multiple libraries have the same name
  govc library.rm /library_name/item_name`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	res, err := flags.ContentLibraryResult(ctx, c, "", f.Arg(0))
	if err != nil {
		return err
	}

	switch t := res.GetResult().(type) {
	case library.Library:
		return m.DeleteLibrary(ctx, &t)
	case library.Item:
		return m.DeleteLibraryItem(ctx, &t)
	default:
		return fmt.Errorf("%q is a %T", f.Arg(0), t)
	}
}
