// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.FolderFlag

	types.ClusterConfigSpecEx
}

func init() {
	cli.Register("cluster.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "CLUSTER"
}

func (cmd *create) Description() string {
	return `Create CLUSTER in datacenter.

The cluster is added to the folder specified by the 'folder' flag. If not given,
this defaults to the host folder in the specified or default datacenter.

Examples:
  govc cluster.create ClusterA
  govc cluster.create -folder /dc2/test-folder ClusterB`
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	folder, err := cmd.FolderOrDefault("host")
	if err != nil {
		return err
	}

	_, err = folder.CreateCluster(ctx, f.Arg(0), cmd.ClusterConfigSpecEx)

	return err
}
