/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

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

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type instantclone struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.NetworkFlag
	*flags.FolderFlag
	*flags.VirtualMachineFlag

	name        string
	extraConfig extraConfig

	Client         *vim25.Client
	Datacenter     *object.Datacenter
	Datastore      *object.Datastore
	ResourcePool   *object.ResourcePool
	Folder         *object.Folder
	VirtualMachine *object.VirtualMachine
}

func init() {
	cli.Register("vm.instantclone", &instantclone{})
}

func (cmd *instantclone) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.Var(&cmd.extraConfig, "e", "ExtraConfig. <key>=<value>")
}

func (cmd *instantclone) Usage() string {
	return "NAME"
}

func (cmd *instantclone) Description() string {
	return `Instant Clone VM to NAME.

Examples:
  govc vm.instantclone -vm source-vm new-vm
  # Configure ExtraConfig variables on a guest VM:
  govc vm.instantclone -vm source-vm -e guestinfo.ipaddress=192.168.0.1 -e guestinfo.netmask=255.255.255.0 new-vm
  # Read the variable set above inside the guest:
  vmware-rpctool "info-get guestinfo.ipaddress"
  vmware-rpctool "info-get guestinfo.netmask"`
}

func (cmd *instantclone) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *instantclone) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error

	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	cmd.name = f.Arg(0)
	if cmd.name == "" {
		return flag.ErrHelp
	}

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	cmd.Datacenter, err = cmd.DatacenterFlag.Datacenter()
	if err != nil {
		return err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return err
	}

	cmd.Folder, err = cmd.FolderFlag.Folder()
	if err != nil {
		return err
	}

	cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
	if err != nil {
		return err
	}

	cmd.VirtualMachine, err = cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}

	if cmd.VirtualMachine == nil {
		return flag.ErrHelp
	}

	_, err = cmd.instantcloneVM(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *instantclone) instantcloneVM(ctx context.Context) (*object.VirtualMachine, error) {
	relocateSpec := types.VirtualMachineRelocateSpec{}

	if cmd.NetworkFlag.IsSet() {
		devices, err := cmd.VirtualMachine.Device(ctx)
		if err != nil {
			return nil, err
		}

		// prepare virtual device config spec for network card
		configSpecs := []types.BaseVirtualDeviceConfigSpec{}

		op := types.VirtualDeviceConfigSpecOperationAdd
		card, derr := cmd.NetworkFlag.Device()
		if derr != nil {
			return nil, derr
		}
		// search for the first network card of the source
		for _, device := range devices {
			if _, ok := device.(types.BaseVirtualEthernetCard); ok {
				op = types.VirtualDeviceConfigSpecOperationEdit
				// set new backing info
				cmd.NetworkFlag.Change(device, card)
				card = device
				break
			}
		}

		configSpecs = append(configSpecs, &types.VirtualDeviceConfigSpec{
			Operation: op,
			Device:    card,
		})

		relocateSpec.DeviceChange = configSpecs
	}

	if cmd.FolderFlag.IsSet() {
		folderref := cmd.Folder.Reference()
		relocateSpec.Folder = &folderref
	}

	if cmd.ResourcePoolFlag.IsSet() {
		poolref := cmd.ResourcePool.Reference()
		relocateSpec.Pool = &poolref
	}

	if cmd.DatastoreFlag.IsSet() {
		datastoreref := cmd.Datastore.Reference()
		relocateSpec.Datastore = &datastoreref
	}

	instantcloneSpec := &types.VirtualMachineInstantCloneSpec{
		Name:     cmd.name,
		Location: relocateSpec,
	}

	if len(cmd.extraConfig) > 0 {
		instantcloneSpec.Config = cmd.extraConfig
	}

	task, err := cmd.VirtualMachine.InstantClone(ctx, *instantcloneSpec)
	if err != nil {
		return nil, err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Instant Cloning %s to %s...", cmd.VirtualMachine.InventoryPath, cmd.name))
	defer logger.Wait()

	info, err := task.WaitForResult(ctx, logger)
	if err != nil {
		return nil, err
	}

	return object.NewVirtualMachine(cmd.Client, info.Result.(types.ManagedObjectReference)), nil
}
