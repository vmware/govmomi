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

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type attach struct {
	*flags.VirtualMachineFlag
	*flags.DatastoreFlag
}

func init() {
	cli.Register("disk.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *attach) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.DatastoreFlag.Process(ctx)
}

func (cmd *attach) Usage() string {
	return "ID"
}

func (cmd *attach) Description() string {
	return `Attach disk ID on VM.

See also: govc vm.disk.attach

Examples:
  govc disk.attach -vm $vm ID
  govc disk.attach -vm $vm -ds $ds ID`
}

func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if ds == nil {
		var props mo.VirtualMachine
		err = vm.Properties(ctx, vm.Reference(), []string{"datastore"}, &props)
		if err != nil {
			return err
		}
		if len(props.Datastore) != 1 {
			ds, err = cmd.Datastore() // likely results in MultipleFoundError
			if err != nil {
				return err
			}
		}
	}

	id := f.Arg(0)

	return vm.AttachDisk(ctx, id, ds, 0, nil)
}
