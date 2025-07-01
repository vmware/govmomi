// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package offline

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/esx/settings/depots"
)

type create struct {
	*flags.ClientFlag

	spec depots.SettingsDepotsOfflineCreateSpec
}

func init() {
	cli.Register("vlcm.depot.offline.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.spec.Location, "l", "", "The URL to the depot contents. Only applicable when source-type is PULL")
	f.StringVar(&cmd.spec.Description, "d", "", "An optional description")
	f.StringVar(&cmd.spec.OwnerData, "owner-data", "", "Optional data about the depot's owner")
	f.StringVar(&cmd.spec.FileId, "file-id", "", "File identifier. Only used when source-type is PUSH")
	f.StringVar(&cmd.spec.SourceType, "source-type", string(depots.SourceTypePull), "The depot source type. The default is PULL")
}

func (cmd *create) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Usage() string {
	return "VLCM"
}

func (cmd *create) Description() string {
	return `Creates an offline image depot.

Execution will block the terminal for the duration of the task.

Examples:
  govc vlcm.depot.offline.create -l=<https://your.server.com/filepath>
  govc vlcm.depot.offline.create -l=<https://your.server.com/filepath> -source-type=PULL
  govc vlcm.depot.offline.create -file-id=<your file identifier> -source-type=PUSH
  govc vlcm.depot.offline.create -l=<https://your.server.com/filepath> -source-type=PULL -d="This is a depot used for testing" -owner-data="After all, why not? Why shouldn't I keep it?"`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := depots.NewManager(rc)

	if taskId, err := dm.CreateOfflineDepot(cmd.spec); err != nil {
		return err
	} else if _, err = tasks.NewManager(rc).WaitForCompletion(ctx, taskId); err != nil {
		return err
	} else {
		return nil
	}
}
