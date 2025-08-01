// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package maintenance

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type exit struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("datastore.maintenance.exit", &exit{})
}

func (cmd *exit) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *exit) Usage() string {
	return "DATASTORE"
}

func (cmd *exit) Description() string {
	return `Take DATASTORE out of maintenance mode.`
}

func (cmd *exit) ExitMaintenanceMode(ctx context.Context, ds *object.Datastore) error {
	req := &types.DatastoreExitMaintenanceMode_Task{
		This: ds.Reference(),
	}
	res, err := methods.DatastoreExitMaintenanceMode_Task(ctx, ds.Client(), req)
	if err != nil {
		return err
	}

	task := object.NewTask(ds.Client(), res.Returnval)

	logger := cmd.ProgressLogger(fmt.Sprintf("%s exiting maintenance mode... ", ds.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *exit) Run(ctx context.Context, f *flag.FlagSet) error {
	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	return cmd.ExitMaintenanceMode(ctx, ds)
}
