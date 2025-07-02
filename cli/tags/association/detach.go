// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package association

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type detach struct {
	*flags.DatacenterFlag
	cat string
}

func init() {
	cli.Register("tags.detach", &detach{})
}

func (cmd *detach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.cat, "c", "", "Tag category")
}

func (cmd *detach) Usage() string {
	return "NAME PATH"
}

func (cmd *detach) Description() string {
	return `Detach tag NAME from object PATH.

Examples:
  govc tags.detach k8s-region-us /dc1
  govc tags.detach -c k8s-region us-ca1 /dc1/host/cluster1`
}

func (cmd *detach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	tagID := f.Arg(0)
	managedObj := f.Arg(1)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	ref, err := convertPath(ctx, c, cmd.DatacenterFlag, managedObj)
	if err != nil {
		return err
	}
	m := tags.NewManager(c)
	tag, err := m.GetTagForCategory(ctx, tagID, cmd.cat)
	if err != nil {
		return err
	}
	return m.DetachTag(ctx, tag.ID, ref)
}
