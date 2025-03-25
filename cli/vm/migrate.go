// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type migrate struct {
	*flags.FolderFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.DatastoreFlag
	*flags.NetworkFlag
	*flags.VirtualMachineFlag

	priority types.VirtualMachineMovePriority
	spec     types.VirtualMachineRelocateSpec
}

func init() {
	cli.Register("vm.migrate", &migrate{})
}

func (cmd *migrate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)

	f.StringVar((*string)(&cmd.priority), "priority", string(types.VirtualMachineMovePriorityDefaultPriority), "The task priority")
}

func (cmd *migrate) Process(ctx context.Context) error {
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *migrate) Usage() string {
	return "VM..."
}

func (cmd *migrate) Description() string {
	return `Migrates VM to a specific resource pool, host or datastore.

Examples:
  govc vm.migrate -host another-host vm-1 vm-2 vm-3
  govc vm.migrate -pool another-pool vm-1 vm-2 vm-3
  govc vm.migrate -ds another-ds vm-1 vm-2 vm-3`
}

func (cmd *migrate) relocate(ctx context.Context, vm *object.VirtualMachine) error {
	spec := cmd.spec

	if cmd.NetworkFlag.IsSet() {
		dev, err := cmd.NetworkFlag.Device()
		if err != nil {
			return err
		}

		devices, err := vm.Device(ctx)
		if err != nil {
			return err
		}

		net := devices.SelectByType((*types.VirtualEthernetCard)(nil))
		if len(net) != 1 {
			return fmt.Errorf("-net specified, but %s has %d nics", vm.Name(), len(net))
		}
		cmd.NetworkFlag.Change(net[0], dev)

		spec.DeviceChange = append(spec.DeviceChange, &types.VirtualDeviceConfigSpec{
			Device:    net[0],
			Operation: types.VirtualDeviceConfigSpecOperationEdit,
		})
	}

	if cmd.VirtualMachineFlag.Spec {
		return cmd.VirtualMachineFlag.WriteAny(spec)
	}

	task, err := vm.Relocate(ctx, spec, cmd.priority)
	if err != nil {
		return err
	}

	logger := cmd.DatastoreFlag.ProgressLogger(fmt.Sprintf("migrating %s... ", vm.Reference()))
	_, err = task.WaitForResult(ctx, logger)
	if err != nil {
		return err
	}

	logger.Wait()

	return nil
}

func (cmd *migrate) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachineFlag.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	folder, err := cmd.FolderIfSpecified()
	if err != nil {
		return err
	}

	if folder != nil {
		ref := folder.Reference()
		cmd.spec.Folder = &ref
	}

	host, err := cmd.HostSystemFlag.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	if host != nil {
		ref := host.Reference()
		cmd.spec.Host = &ref
	}

	pool, err := cmd.ResourcePoolFlag.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}

	if pool == nil && host != nil {
		pool, err = host.ResourcePool(ctx)
		if err != nil {
			return err
		}
	}

	if pool != nil {
		ref := pool.Reference()
		cmd.spec.Pool = &ref
	}

	ds, err := cmd.DatastoreFlag.DatastoreIfSpecified()
	if err != nil {
		return err
	}

	if ds != nil {
		ref := ds.Reference()
		cmd.spec.Datastore = &ref
	}

	for _, vm := range vms {
		err = cmd.relocate(ctx, vm)
		if err != nil {
			return err
		}
	}

	return nil
}
