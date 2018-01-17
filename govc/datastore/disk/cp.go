/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package disk

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
)

type cp struct {
	*flags.DatastoreFlag

	spec
}

func init() {
	cli.Register("datastore.disk.cp", &cp{})
}

func (cmd *cp) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.spec.Register(ctx, f)
}

func (cmd *cp) Usage() string {
	return "SRC DST"
}

func (cmd *cp) Description() string {
	return `Copy SRC to DST disk on DS.

Examples:
  govc datastore.disk.cp disks/disk1.vmdk disks/disk2.vmdk`
}

func (cmd *cp) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	dc, err := cmd.Datacenter()
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	m := object.NewVirtualDiskManager(ds.Client())
	src := ds.Path(f.Arg(0))
	dst := ds.Path(f.Arg(1))

	task, err := m.CopyVirtualDisk(ctx, src, dc, dst, dc, &cmd.spec.VirtualDiskSpec, cmd.force)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Copying %s to %s...", src, dst))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
