/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package maintenance

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
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
