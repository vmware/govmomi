// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package group

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*InfoFlag

	long bool
}

func init() {
	cli.Register("cluster.group.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.InfoFlag, ctx = NewInfoFlag(ctx)
	cmd.InfoFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Process(ctx context.Context) error {
	return cmd.InfoFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List cluster groups and group members.

Examples:
  govc cluster.group.ls -cluster my_cluster
  govc cluster.group.ls -cluster my_cluster -l | grep ClusterHostGroup
  govc cluster.group.ls -cluster my_cluster -name my_group`
}

type groupResult []string

func (r groupResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}

	return nil
}

type groupResultLong []types.BaseClusterGroupInfo

func (r groupResultLong) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for i := range r {
		info := r[i].GetClusterGroupInfo()
		kind := fmt.Sprintf("%T", r[i])
		kind = strings.SplitN(kind, ".", 2)[1]
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", kind, info.Name)
	}

	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var res groupResult

	if cmd.name == "" {
		groups, err := cmd.Groups(ctx)
		if err != nil {
			return err
		}

		if cmd.long {
			return cmd.WriteResult(groupResultLong(groups))
		}

		for _, g := range groups {
			res = append(res, g.GetClusterGroupInfo().Name)
		}
	} else {
		group, err := cmd.Group(ctx)
		if err != nil {
			return err
		}

		names, err := cmd.Names(ctx, *group.refs)
		if err != nil {
			return err
		}

		for _, ref := range *group.refs {
			res = append(res, names[ref])
		}
	}

	return cmd.WriteResult(res)
}
