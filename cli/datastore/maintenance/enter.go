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

type enter struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("datastore.maintenance.enter", &enter{})
}

func (cmd *enter) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *enter) Usage() string {
	return "DATASTORE"
}

func (cmd *enter) Description() string {
	return `Put DATASTORE in maintenance mode.

Examples:
  govc datastore.cluster.change -drs-mode automated my-datastore-cluster # automatically schedule Storage DRS migration
  govc datastore.maintenance.enter -ds my-datastore-cluster/datastore1
  # no virtual machines can be powered on and no provisioning operations can be performed on the datastore during this time
  govc datastore.maintenance.exit -ds my-datastore-cluster/datastore1`
}

func (cmd *enter) EnterMaintenanceMode(ctx context.Context, ds *object.Datastore) error {
	req := &types.DatastoreEnterMaintenanceMode{
		This: ds.Reference(),
	}
	res, err := methods.DatastoreEnterMaintenanceMode(ctx, ds.Client(), req)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s entering maintenance mode... ", ds.InventoryPath))
	defer logger.Wait()

	task := object.NewTask(ds.Client(), *res.Returnval.Task)
	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *enter) Run(ctx context.Context, f *flag.FlagSet) error {
	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	return cmd.EnterMaintenanceMode(ctx, ds)
}
