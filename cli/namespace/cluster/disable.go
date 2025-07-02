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
	"github.com/vmware/govmomi/vapi/namespace"
)

type disableCluster struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("namespace.cluster.disable", &disableCluster{})
}

func (cmd *disableCluster) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *disableCluster) Description() string {
	return `Disables vSphere Namespaces on the specified cluster.

Examples:
  govc namespace.cluster.disable -cluster "Workload-Cluster"`
}

func (cmd *disableCluster) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	// Cluster object reference lookup
	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	clusterId := cluster.Reference().Value

	err = m.DisableCluster(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("error disabling cluster: %s", err)
	}

	return nil
}
