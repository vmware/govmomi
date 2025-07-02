// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package group

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*InfoFlag
}

func init() {
	cli.Register("cluster.group.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.InfoFlag, ctx = NewInfoFlag(ctx)
	cmd.InfoFlag.Register(ctx, f)
}

func (cmd *change) Process(ctx context.Context) error {
	if cmd.name == "" {
		return flag.ErrHelp
	}
	return cmd.InfoFlag.Process(ctx)
}

func (cmd *change) Usage() string {
	return `NAME...`
}

func (cmd *change) Description() string {
	return `Set cluster group members.

Examples:
  govc cluster.group.change -name my_group vm_a vm_b vm_c # set
  govc cluster.group.change -name my_group vm_a vm_b vm_c $(govc cluster.group.ls -name my_group) vm_d # add
  govc cluster.group.ls -name my_group | grep -v vm_b | xargs govc cluster.group.change -name my_group vm_a vm_b vm_c # remove`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	update := types.ArrayUpdateSpec{Operation: types.ArrayUpdateOperationEdit}
	group, err := cmd.Group(ctx)
	if err != nil {
		return err
	}

	refs, err := cmd.ObjectList(ctx, group.kind, f.Args())
	if err != nil {
		return err
	}

	*group.refs = refs

	return cmd.Apply(ctx, update, group.info)
}
