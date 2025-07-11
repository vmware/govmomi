// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"path"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type DatacenterFolders struct {
	VmFolder        *Folder
	HostFolder      *Folder
	DatastoreFolder *Folder
	NetworkFolder   *Folder
}

type Datacenter struct {
	Common
}

func NewDatacenter(c *vim25.Client, ref types.ManagedObjectReference) *Datacenter {
	return &Datacenter{
		Common: NewCommon(c, ref),
	}
}

func (d *Datacenter) Folders(ctx context.Context) (*DatacenterFolders, error) {
	var md mo.Datacenter

	ps := []string{"name", "vmFolder", "hostFolder", "datastoreFolder", "networkFolder"}
	err := d.Properties(ctx, d.Reference(), ps, &md)
	if err != nil {
		return nil, err
	}

	df := &DatacenterFolders{
		VmFolder:        NewFolder(d.c, md.VmFolder),
		HostFolder:      NewFolder(d.c, md.HostFolder),
		DatastoreFolder: NewFolder(d.c, md.DatastoreFolder),
		NetworkFolder:   NewFolder(d.c, md.NetworkFolder),
	}

	paths := []struct {
		name string
		path *string
	}{
		{"vm", &df.VmFolder.InventoryPath},
		{"host", &df.HostFolder.InventoryPath},
		{"datastore", &df.DatastoreFolder.InventoryPath},
		{"network", &df.NetworkFolder.InventoryPath},
	}

	dcPath := d.InventoryPath
	if dcPath == "" {
		dcPath = "/" + md.Name
	}
	for _, p := range paths {
		*p.path = path.Join(dcPath, p.name)
	}

	return df, nil
}

func (d Datacenter) Destroy(ctx context.Context) (*Task, error) {
	req := types.Destroy_Task{
		This: d.Reference(),
	}

	res, err := methods.Destroy_Task(ctx, d.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(d.c, res.Returnval), nil
}

// PowerOnVM powers on multiple virtual machines with a single vCenter call.
// If called against ESX, serially powers on the list of VMs and the returned *Task will always be nil.
func (d Datacenter) PowerOnVM(ctx context.Context, vm []types.ManagedObjectReference, option ...types.BaseOptionValue) (*Task, error) {
	if d.Client().IsVC() {
		req := types.PowerOnMultiVM_Task{
			This:   d.Reference(),
			Vm:     vm,
			Option: option,
		}

		res, err := methods.PowerOnMultiVM_Task(ctx, d.c, &req)
		if err != nil {
			return nil, err
		}

		return NewTask(d.c, res.Returnval), nil
	}

	for _, ref := range vm {
		obj := NewVirtualMachine(d.Client(), ref)
		task, err := obj.PowerOn(ctx)
		if err != nil {
			return nil, err
		}

		err = task.Wait(ctx)
		if err != nil {
			// Ignore any InvalidPowerState fault, as it indicates the VM is already powered on
			if !fault.Is(err, &types.InvalidPowerState{}) {
				return nil, err
			}
		}
	}

	return nil, nil
}
