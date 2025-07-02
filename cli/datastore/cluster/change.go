// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

func DrsBehaviorUsage() string {
	drsModes := types.StorageDrsPodConfigInfoBehavior("").Strings()

	return "Storage DRS behavior: " + strings.Join(drsModes, ", ")
}

type change struct {
	*flags.DatacenterFlag

	types.StorageDrsConfigSpec
}

func init() {
	cli.Register("datastore.cluster.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.PodConfigSpec = new(types.StorageDrsPodConfigSpec)

	f.Var(flags.NewOptionalBool(&cmd.PodConfigSpec.Enabled), "drs-enabled", "Enable Storage DRS")

	f.StringVar((*string)(&cmd.PodConfigSpec.DefaultVmBehavior), "drs-mode", "", DrsBehaviorUsage())
}

func (cmd *change) Usage() string {
	return "CLUSTER..."
}

func (cmd *change) Description() string {
	return `Change configuration of the given datastore clusters.

Examples:
  govc datastore.cluster.change -drs-enabled ClusterA
  govc datastore.cluster.change -drs-enabled=false ClusterB`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}
	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	m := object.NewStorageResourceManager(client)

	for _, path := range f.Args() {
		clusters, err := finder.DatastoreClusterList(ctx, path)
		if err != nil {
			return err
		}

		for _, cluster := range clusters {
			task, err := m.ConfigureStorageDrsForPod(ctx, cluster, cmd.StorageDrsConfigSpec, true)
			if err != nil {
				return err
			}

			_, err = task.WaitForResult(ctx, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
