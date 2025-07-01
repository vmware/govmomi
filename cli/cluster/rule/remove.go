// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rule

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*InfoFlag
}

func init() {
	cli.Register("cluster.rule.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.InfoFlag, ctx = NewInfoFlag(ctx)
	cmd.InfoFlag.Register(ctx, f)
}

func (cmd *remove) Process(ctx context.Context) error {
	if cmd.name == "" {
		return flag.ErrHelp
	}
	return cmd.InfoFlag.Process(ctx)
}

func (cmd *remove) Description() string {
	return `Remove cluster rule.

Examples:
  govc cluster.group.remove -cluster my_cluster -name my_rule`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	rule, err := cmd.Rule(ctx)
	if err != nil {
		return err
	}

	update := types.ArrayUpdateSpec{
		Operation: types.ArrayUpdateOperationRemove,
		RemoveKey: rule.info.GetClusterRuleInfo().Key,
	}

	return cmd.Apply(ctx, update, nil)
}
