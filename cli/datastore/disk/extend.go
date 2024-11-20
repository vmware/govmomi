/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

	"github.com/vmware/govmomi/units"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type extend struct {
	*flags.DatastoreFlag
	Bytes     units.ByteSize
	EagerZero bool
}

func init() {
	cli.Register("datastore.disk.extend", &extend{})
}

func (cmd *extend) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.Var(&cmd.Bytes, "size", "New capacity for the disk")
	f.BoolVar(&cmd.EagerZero, "eagerZero", false, "If true, the extended part of the disk will be explicitly filled with zeroes")
}

func (cmd *extend) Process(ctx context.Context) error {
	return cmd.DatastoreFlag.Process(ctx)
}

func (cmd *extend) Usage() string {
	return "VMDK"
}

func (cmd *extend) Description() string {
	return `Extend VMDK on DS.

Examples:
  govc datastore.disk.extend disks/disk1.vmdk -size=24G
  govc datastore.disk.extend disks/disk1.vmdk -size=24G -eagerZero=true`
}

func (cmd *extend) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
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
	path := ds.Path(f.Arg(0))
	task, err := m.ExtendVirtualDisk(ctx, path, dc, int64(cmd.Bytes/1024), &cmd.EagerZero)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Extending %s...", path))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
