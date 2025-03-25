// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type register struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.FolderFlag

	name     string
	template bool
}

func init() {
	cli.Register("vm.register", &register{})
}

func (cmd *register) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	f.StringVar(&cmd.name, "name", "", "Name of the VM")
	f.BoolVar(&cmd.template, "template", false, "Mark VM as template")
}

func (cmd *register) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *register) Usage() string {
	return "VMX"
}

func (cmd *register) Description() string {
	return `Add an existing VM to the inventory.

VMX is a path to the vm config file, relative to DATASTORE.

Examples:
  govc vm.register path/name.vmx
  govc vm.register -template -host $host path/name.vmx`
}

func (cmd *register) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	pool, err := cmd.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemFlag.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	if cmd.template {
		if pool != nil || host == nil {
			return flag.ErrHelp
		}
	} else if pool == nil {
		if host != nil {
			pool, err = host.ResourcePool(ctx)
			if err != nil {
				return err
			}
		} else {
			// neither -host nor -pool were specified, so use the default pool (ESX)
			pool, err = cmd.ResourcePool()
			if err != nil {
				return err
			}
		}
	}

	folder, err := cmd.FolderFlag.Folder()
	if err != nil {
		return err
	}

	path, err := cmd.DatastorePath(f.Arg(0))
	if err != nil {
		return err
	}

	task, err := folder.RegisterVM(ctx, path, cmd.name, cmd.template, pool, host)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
