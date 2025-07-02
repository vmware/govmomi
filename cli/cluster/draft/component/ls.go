// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"context"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type lsResult map[string]clusters.SettingsComponentInfo

func (r lsResult) Write(w io.Writer) error {
	return nil
}

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId string
	draftId   string
}

func init() {
	cli.Register("cluster.draft.component.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")

}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *ls) Usage() string {
	return "CLUSTER"
}

func (cmd *ls) Description() string {
	return `Displays the list of components in a software draft.

Examples:
  govc cluster.draft.component.ls -cluster-id=domain-c21 -draft-id=13`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	if d, err := dm.ListSoftwareDraftComponents(cmd.clusterId, cmd.draftId); err != nil {
		return err
	} else {
		if !cmd.All() {
			cmd.JSON = true
		}
		return cmd.WriteResult(lsResult(d))
	}
}
