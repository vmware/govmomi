// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package association

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25/mo"
)

type ls struct {
	*flags.DatacenterFlag
	r bool
	l bool
}

func init() {
	cli.Register("tags.attached.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
	f.BoolVar(&cmd.r, "r", false, "List tags attached to resource")
	f.BoolVar(&cmd.l, "l", false, "Long listing format")
}

func (cmd *ls) Usage() string {
	return "NAME"
}

func (cmd *ls) Description() string {
	return `List attached tags or objects.

Examples:
  govc tags.attached.ls k8s-region-us
  govc tags.attached.ls -json k8s-zone-us-ca1 | jq .
  govc tags.attached.ls -r /dc1/host/cluster1
  govc tags.attached.ls -json -r /dc1 | jq .`
}

type lsResult []string

func (r lsResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

type lsTagResult []tags.AttachedTags

func (r lsTagResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, item := range r {
		var ids []string
		for _, id := range item.Tags {
			ids = append(ids, id.Name)
		}
		fmt.Fprintf(tw, "%s:\t%s\n", item.ObjectID.Reference(), strings.Join(ids, ","))
	}
	return tw.Flush()
}

type lsObjectResult []tags.AttachedObjects

func (r lsObjectResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, item := range r {
		for _, id := range item.ObjectIDs {
			var ids []string
			ids = append(ids, id.Reference().String())
			fmt.Fprintf(tw, "%s:\t%s\n", item.Tag.Name, strings.Join(ids, ","))
		}

	}
	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	var res flags.OutputWriter
	m := tags.NewManager(c)

	if cmd.r {
		var refs []mo.Reference
		for _, arg := range f.Args() {
			ref, err := convertPath(ctx, c, cmd.DatacenterFlag, arg)
			if err != nil {
				return err
			}
			refs = append(refs, ref)
		}
		attached, err := m.GetAttachedTagsOnObjects(ctx, refs)
		if err != nil {
			return err
		}
		if cmd.l {
			res = lsTagResult(attached)
		} else {
			var r lsResult
			for i := range attached {
				for _, tag := range attached[i].Tags {
					r = append(r, tag.Name)
				}
			}
			res = r
		}
	} else {
		attached, err := m.GetAttachedObjectsOnTags(ctx, f.Args())
		if err != nil {
			return err
		}
		if cmd.l {
			res = lsObjectResult(attached)
		} else {
			var r lsResult
			for _, obj := range attached {
				for _, ref := range obj.ObjectIDs {
					r = append(r, ref.Reference().String())
				}
			}
			res = r
		}
	}

	return cmd.WriteResult(res)
}
