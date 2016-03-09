/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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
	"errors"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type create struct {
	*flags.DatastoreFlag
	*flags.OutputFlag
	*flags.VirtualMachineFlag

	controller string
	Name       string
	Bytes      units.ByteSize
	Thick      bool
	Eager      bool
	DiskMode   string
}

func init() {
	cli.Register("vm.disk.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	err := (&cmd.Bytes).Set("10G")
	if err != nil {
		panic(err)
	}

	f.StringVar(&cmd.controller, "controller", "", "Disk controller")
	f.StringVar(&cmd.Name, "name", "", "Name for new disk")
	f.Var(&cmd.Bytes, "size", "Size of new disk")
	f.BoolVar(&cmd.Thick, "thick", false, "Thick provision new disk")
	f.BoolVar(&cmd.Eager, "eager", false, "Eagerly scrub new disk")
	f.StringVar(&cmd.DiskMode, "mode", "dependent", "One of: dependent (also known as persistent), persistent, nonpersistent, undoable, independent_persistent, independent_nonpersistent, append")
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(cmd.Name) == 0 {
		return errors.New("please specify a disk name")
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return errors.New("please specify a vm")
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	devices, err := vm.Device(context.TODO())
	if err != nil {
		return err
	}

	controller, err := devices.FindDiskController(cmd.controller)
	if err != nil {
		return err
	}

	disk := devices.CreateDisk(controller, ds.Reference(), ds.Path(cmd.Name))

	existing := devices.SelectByBackingInfo(disk.Backing)

	if len(existing) > 0 {
		cmd.Log("Disk already present\n")
		return nil
	}

	backing := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)

	if cmd.Thick {
		backing.ThinProvisioned = types.NewBool(false)
	}

	backing.EagerlyScrub = types.NewBool(cmd.Eager)

	switch cmd.DiskMode {
	// NOTE: testing via the vSphere web UI has confirmed that Dependent is saved as "persistent"
	case "dependent":
	case "persistent":
		backing.DiskMode = string(types.VirtualDiskModePersistent)
	case "nonpersistent":
		backing.DiskMode = string(types.VirtualDiskModeNonpersistent)
	case "undoable":
		backing.DiskMode = string(types.VirtualDiskModeUndoable)
	case "independent_persistent":
		backing.DiskMode = string(types.VirtualDiskModeIndependent_persistent)
	case "independent_nonpersistent":
		backing.DiskMode = string(types.VirtualDiskModeIndependent_nonpersistent)
	case "append":
		backing.DiskMode = string(types.VirtualDiskModeAppend)
	default:
		return errors.New("please specify a valid mode")
	}

	cmd.Log("Creating disk\n")
	disk.CapacityInKB = int64(cmd.Bytes) / 1024
	return vm.AddDevice(context.TODO(), disk)
}
