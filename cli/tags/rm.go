// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags

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

	cat   string
	force bool
}

func init() {
	cli.Register("tags.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.cat, "c", "", "Tag category")
	f.BoolVar(&cmd.force, "f", false, "Delete tag regardless of attached objects")
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Delete tag NAME.

Fails if tag is attached to any object, unless the '-f' flag is provided.

Examples:
  govc tags.rm k8s-zone-us-ca1
  govc tags.rm -f -c k8s-zone us-ca2`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	tagID := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := tags.NewManager(c)
	if !cmd.force {
		objs, err := m.ListAttachedObjects(ctx, tagID)
		if err != nil {
			return err
		}
		if len(objs) > 0 {
			return fmt.Errorf("tag %s has %d attached objects", tagID, len(objs))
		}
	}
	tag, err := m.GetTagForCategory(ctx, tagID, cmd.cat)
	if err != nil {
		return err
	}
	return m.DeleteTag(ctx, tag)
}
