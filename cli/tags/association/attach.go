// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package association

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25/types"
)

type attach struct {
	*flags.DatacenterFlag
	cat string
}

func init() {
	cli.Register("tags.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.cat, "c", "", "Tag category")
}

func (cmd *attach) Usage() string {
	return "NAME PATH"
}

func (cmd *attach) Description() string {
	return `Attach tag NAME to object PATH.

Examples:
  govc tags.attach k8s-region-us /dc1
  govc tags.attach -c k8s-region us-ca1 /dc1/host/cluster1`
}

func convertPath(ctx context.Context, c *rest.Client, cmd *flags.DatacenterFlag, managedObj string) (*types.ManagedObjectReference, error) {
	client, err := cmd.ClientFlag.Client()
	if err != nil {
		return nil, err
	}

	ref := client.ServiceContent.RootFolder

	switch managedObj {
	case "", "-":
	default:
		ref, err = cmd.ManagedObject(ctx, managedObj)
		if err != nil {
			m := library.NewManager(c)
			res, _ := finder.NewFinder(m).Find(ctx, managedObj)
			if len(res) != 1 {
				return nil, err
			}
			switch t := res[0].GetResult().(type) {
			case library.Library:
				ref = types.ManagedObjectReference{Type: "com.vmware.content.Library", Value: t.ID}
			case library.Item:
				ref = types.ManagedObjectReference{Type: "com.vmware.content.library.Item", Value: t.ID}
			default:
				return nil, err
			}
		}
	}
	return &ref, nil
}

func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
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
	return m.AttachTag(ctx, tag.ID, ref)
}
