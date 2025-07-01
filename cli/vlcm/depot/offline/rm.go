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

type rm struct {
	*flags.ClientFlag

	depotId string
}

func init() {
	cli.Register("vlcm.depot.offline.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.depotId, "depot-id", "", "The identifier of the depot. Use the 'ls' command to see the list of depots.")
}

func (cmd *rm) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *rm) Usage() string {
	return "VLCM"
}

func (cmd *rm) Description() string {
	return `Deletes an offline image depot.

Execution will block the terminal for the duration of the task.

Examples:
  govc vlcm.depot.offline.rm -depot-id=<your depot's identifier>`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := depots.NewManager(rc)

	if taskId, err := dm.DeleteOfflineDepot(cmd.depotId); err != nil {
		return err
	} else if _, err = tasks.NewManager(rc).WaitForCompletion(ctx, taskId); err != nil {
		return err
	} else {
		return nil
	}
}
