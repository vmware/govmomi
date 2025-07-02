// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type mv struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("cluster.mv", &mv{})
}

func (cmd *mv) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *mv) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *mv) Description() string {
	return `Move HOST to CLUSTER.

The hosts are moved to the cluster specified by the 'cluster' flag.

Examples:
  govc cluster.mv -cluster ClusterA host1 host2`
}

func (cmd *mv) Move(ctx context.Context, cluster *object.ClusterComputeResource, hosts []*object.HostSystem) error {
	task, err := cluster.MoveInto(ctx, hosts...)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("moving %d hosts to cluster %s... ", len(hosts), cluster.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *mv) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	var hosts []*object.HostSystem

	for _, path := range f.Args() {
		list, err := finder.HostSystemList(ctx, path)
		if err != nil {
			return err
		}
		hosts = append(hosts, list...)
	}

	return cmd.Move(ctx, cluster, hosts)
}
