// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type rm struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId   string
	draftId     string
	componentId string
}

func init() {
	cli.Register("cluster.draft.component.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")
	f.StringVar(&cmd.componentId, "component-id", "", "The identifier of the software component.")
}

func (cmd *rm) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *rm) Usage() string {
	return "CLUSTER"
}

func (cmd *rm) Description() string {
	return `Removes a component from a software draft.

Examples:
  govc cluster.draft.component.rm -cluster-id=domain-c21 -draft-id=13 -component-id=NVD-AIE-800`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	return dm.RemoveSoftwareDraftComponents(cmd.clusterId, cmd.draftId, cmd.componentId)
}
