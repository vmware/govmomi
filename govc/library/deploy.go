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

package library

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vcenter"
	"github.com/vmware/govmomi/vim25/types"
)

type deploy struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.FolderFlag
}

func init() {
	cli.Register("library.deploy", &deploy{})
}

func (cmd *deploy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *deploy) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.FolderFlag.Process(ctx)
}

func (cmd *deploy) Usage() string {
	return "TEMPLATE VM_NAME"
}

func (cmd *deploy) Description() string {
	return `Deploy library OVF template.

Examples:
  govc library.deploy /library_name/ovf_template vm_name`
}

func (cmd *deploy) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	name := f.Arg(1)

	return cmd.DatastoreFlag.WithRestClient(ctx, func(c *rest.Client) error {
		m := vcenter.NewManager(c)

		res, err := finder.NewFinder(library.NewManager(c)).Find(ctx, path)
		if err != nil {
			return err
		}
		if len(res) != 1 {
			return fmt.Errorf("%q matches %d items", path, len(res))
		}
		item, ok := res[0].GetResult().(library.Item)
		if !ok {
			return fmt.Errorf("%q is a %T", path, item)
		}

		ds, err := cmd.Datastore()
		if err != nil {
			return err
		}
		rp, err := cmd.ResourcePool()
		if err != nil {
			return err
		}
		folder, err := cmd.Folder()
		if err != nil {
			return err
		}

		deploy := vcenter.Deploy{
			DeploymentSpec: vcenter.DeploymentSpec{
				Name:               name,
				DefaultDatastoreID: ds.Reference().Value,
				AcceptAllEULA:      true,
			},
			Target: vcenter.Target{
				ResourcePoolID: rp.Reference().Value,
				FolderID:       folder.Reference().Value,
			},
		}

		d, err := m.DeployLibraryItem(ctx, item.ID, deploy)
		if err != nil {
			return err
		}

		if !d.Succeeded {
			return errors.New(d.Error.Error())
		}

		finder, err := cmd.FolderFlag.Finder(false)
		if err != nil {
			return err
		}
		ref, err := finder.ObjectReference(ctx, types.ManagedObjectReference{
			Type:  "VirtualMachine",
			Value: d.ResourceID.ID,
		})
		if err != nil {
			return err
		}

		vm := ref.(*object.VirtualMachine)

		fmt.Println(vm.InventoryPath)

		return nil
	})
}
