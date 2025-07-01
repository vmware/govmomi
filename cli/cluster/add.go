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

type add struct {
	*flags.ClusterFlag
	*flags.HostConnectFlag

	connect bool
	license string
}

func init() {
	cli.Register("cluster.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.HostConnectFlag, ctx = flags.NewHostConnectFlag(ctx)
	cmd.HostConnectFlag.Register(ctx, f)

	f.StringVar(&cmd.license, "license", "", "Assign license key")

	f.BoolVar(&cmd.connect, "connect", true, "Immediately connect to host")
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostConnectFlag.Process(ctx); err != nil {
		return err
	}
	if cmd.HostName == "" {
		return flag.ErrHelp
	}
	if cmd.UserName == "" {
		return flag.ErrHelp
	}
	if cmd.Password == "" {
		return flag.ErrHelp
	}
	return nil
}

func (cmd *add) Description() string {
	return `Add HOST to CLUSTER.

The host is added to the cluster specified by the 'cluster' flag.

Examples:
  thumbprint=$(govc about.cert -k -u host.example.com -thumbprint | awk '{print $2}')
  govc cluster.add -cluster ClusterA -hostname host.example.com -username root -password pass -thumbprint $thumbprint
  govc cluster.add -cluster ClusterB -hostname 10.0.6.1 -username root -password pass -noverify`
}

func (cmd *add) Add(ctx context.Context, cluster *object.ClusterComputeResource) error {
	spec := cmd.HostConnectSpec

	var license *string
	if cmd.license != "" {
		license = &cmd.license
	}

	task, err := cluster.AddHost(ctx, cmd.Spec(cluster.Client()), cmd.connect, license, nil)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("adding %s to cluster %s... ", spec.HostName, cluster.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 0 {
		return flag.ErrHelp
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	return cmd.Fault(cmd.Add(ctx, cluster))
}
