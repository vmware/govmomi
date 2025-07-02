// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package category

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type rm struct {
	*flags.ClientFlag
	force bool
}

func init() {
	cli.Register("tags.category.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.BoolVar(&cmd.force, "f", false, "Delete tag regardless of attached objects")
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Delete category NAME.

Fails if category is used by any tag, unless the '-f' flag is provided.

Examples:
  govc tags.category.rm k8s-region
  govc tags.category.rm -f k8s-zone`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	categoryID := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := tags.NewManager(c)
	cat, err := m.GetCategory(ctx, categoryID)
	if err != nil {
		return err
	}
	if !cmd.force {
		ctags, err := m.ListTagsForCategory(ctx, cat.ID)
		if err != nil {
			return err
		}
		if len(ctags) > 0 {
			return fmt.Errorf("category %s used by %d tags", categoryID, len(ctags))
		}
	}
	return m.DeleteCategory(ctx, cat)
}
