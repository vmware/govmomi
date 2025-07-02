// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type update struct {
	*flags.ClientFlag

	tag tags.Tag
	cat string
}

func init() {
	cli.Register("tags.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.tag.Name, "n", "", "Name of tag")
	f.StringVar(&cmd.tag.Description, "d", "", "Description of tag")
	f.StringVar(&cmd.cat, "c", "", "Tag category")
}

func (cmd *update) Usage() string {
	return "NAME"
}

func (cmd *update) Description() string {
	return `Update tag.

Examples:
  govc tags.update -d "K8s zone US-CA1" k8s-zone-us-ca1
  govc tags.update -d "K8s zone US-CA1" -c k8s-zone us-ca1`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	arg := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := tags.NewManager(c)
	tag, err := m.GetTagForCategory(ctx, arg, cmd.cat)
	if err != nil {
		return err
	}
	tag.Patch(&cmd.tag)
	return m.UpdateTag(ctx, tag)
}
