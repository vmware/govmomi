// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type tags struct {
	*flags.ClientFlag
	types.VslmTagEntry
	attach bool
}

func init() {
	cli.Register("disk.tags.attach", &tags{attach: true})
	cli.Register("disk.tags.detach", &tags{attach: false})
}

func (cmd *tags) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.ParentCategoryName, "c", "", "Tag category")
}

func (cmd *tags) Usage() string {
	return "NAME ID"
}

func (cmd *tags) name() string {
	if cmd.attach {
		return "attach"
	}
	return "detach"
}

func (cmd *tags) Description() string {
	if cmd.attach {
		return `Attach tag NAME to disk ID.

Examples:
  govc disk.tags.attach -c k8s-region k8s-region-us $id`
	}

	return `Detach tag NAME from disk ID.

Examples:
  govc disk.tags.detach -c k8s-region k8s-region-us $id`
}

func (cmd *tags) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := NewManager(ctx, c, nil)
	if err != nil {
		return err
	}

	cmd.TagName = f.Arg(0)
	method := m.DetachTag
	if cmd.attach {
		method = m.AttachTag
	}
	return method(ctx, f.Arg(1), cmd.VslmTagEntry)
}
