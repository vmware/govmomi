// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package draft

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type commit struct {
	*flags.ClientFlag

	clusterId string
	draftId   string
}

func init() {
	cli.Register("cluster.draft.commit", &commit{})
}

func (cmd *commit) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")
}

func (cmd *commit) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *commit) Usage() string {
	return "CLUSTER"
}

func (cmd *commit) Description() string {
	return `Commits the provided software draft.

Execution will block the terminal for the duration of the task.

Examples:
  govc cluster.draft.commit -cluster-id=domain-c21 -draft-id=13`
}

func (cmd *commit) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	if taskId, err := dm.CommitSoftwareDraft(cmd.clusterId, cmd.draftId, clusters.SettingsClustersSoftwareDraftsCommitSpec{}); err != nil {
		return err
	} else if _, err = tasks.NewManager(rc).WaitForCompletion(ctx, taskId); err != nil {
		return err
	} else {
		return nil
	}
}
