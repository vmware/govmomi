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

type add struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId        string
	draftId          string
	componentId      string
	componentVersion string
}

func init() {
	cli.Register("cluster.draft.component.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")
	f.StringVar(&cmd.componentId, "component-id", "", "The identifier of the software component.")
	f.StringVar(&cmd.componentVersion, "component-version", "", "The version of the software component.")
}

func (cmd *add) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *add) Usage() string {
	return "CLUSTER"
}

func (cmd *add) Description() string {
	return `Adds a new component to the software draft.

Examples:
  govc cluster.draft.component.add -cluster-id=domain-c21 -draft-id=13 -component-id=NVD-AIE-800 -component-version=550.54.10-1OEM.800.1.0.20613240`
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	spec := clusters.SoftwareComponentsUpdateSpec{}
	spec.ComponentsToSet = make(map[string]string)
	spec.ComponentsToSet[cmd.componentId] = cmd.componentVersion
	return dm.UpdateSoftwareDraftComponents(cmd.clusterId, cmd.draftId, spec)
}
