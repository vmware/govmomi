// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	vapicluster "github.com/vmware/govmomi/vapi/cluster"
)

type create struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("cluster.module.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	return cmd.ClusterFlag.Process(ctx)
}

func (cmd *create) Description() string {
	return `Create cluster module.

This command will output the ID of the new module.

Examples:
  govc cluster.module.create -cluster my_cluster`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 0 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	id, err := vapicluster.NewManager(c).CreateModule(ctx, cluster.Reference())
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
