// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/namespace"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	long bool
}

func init() {
	cli.Register("namespace.cluster.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List namepace enabled clusters.

Examples:
  govc namespace.cluster.ls
  govc namespace.cluster.ls -l
  govc namespace.cluster.ls -json | jq .`
}

type lsWriter struct {
	cmd     *ls
	Cluster []namespace.ClusterSummary `json:"cluster"`
}

func (r *lsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, cluster := range r.Cluster {
		path, err := r.path(cluster)
		if err != nil {
			return err
		}
		fmt.Fprintf(tw, "%s", path)
		if r.cmd.long {
			fmt.Fprintf(tw, "\t%s", cluster.ConfigStatus)
			fmt.Fprintf(tw, "\t%s", cluster.KubernetesStatus)
		}
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}

func (r *lsWriter) path(cluster namespace.ClusterSummary) (string, error) {
	ref := cluster.Reference()

	c, err := r.cmd.Client()
	if err != nil {
		return "", err
	}

	finder := find.NewFinder(c, false)
	obj, err := finder.ObjectReference(context.Background(), ref)
	if err != nil {
		return "", err
	}

	return obj.(*object.ClusterComputeResource).InventoryPath, nil
}

func (r *lsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Cluster)
}

func (r *lsWriter) Dump() any {
	return r.Cluster
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	clusters, err := m.ListClusters(ctx)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&lsWriter{cmd, clusters})
}
