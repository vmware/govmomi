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
