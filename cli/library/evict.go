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

type evict struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("library.evict", &evict{})
}

func (cmd *evict) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *evict) Usage() string {
	return "LIBRARY NAME | ITEM NAME"
}

func (cmd *evict) Description() string {
	return `Evict library NAME or item NAME.

Examples:
  govc library.evict subscribed-library
  govc library.evict subscribed-library/item`
}

func (cmd *evict) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

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
		return m.EvictSubscribedLibrary(ctx, &t)
	case library.Item:
		return m.EvictSubscribedLibraryItem(ctx, &t)
	default:
		return fmt.Errorf("%q is a %T", f.Arg(0), t)
	}
}
