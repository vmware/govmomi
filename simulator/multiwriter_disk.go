// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

// AttachMultiWriterDisk adds a multi-writer VirtualDisk to vm backed by
// sharedPath (VirtualDiskFlatVer2BackingInfo with Sharing set to
// sharingMultiWriter). The disk is placed on vm's first ParaVirtualSCSI
// controller.
//
// For the first VM in a shared-disk group, pass create=true so vcsim creates
// the VMDK at sharedPath (1 GiB by default). For each additional VM that
// should participate in the same multi-writer group, call again with the
// same sharedPath and create=false so the existing backing is attached
// (object.diskFileOperation treats zero capacity as attach).
//
// After that, VirtualMachine.FetchVmGroupForMultiwriterDisks on vcsim will
// return peers that share the same backing FileName string.
// The fetch request's diskIds field is ignored by vcsim (real VC filters by it).
//
// sharedPath is typically "[LocalDS_0] name.vmdk" at the datastore root. If you
// use nested segments (e.g. "[LocalDS_0] dir/name.vmdk"), dir must already exist;
// vcsim does not create parent directories when create is true.
func AttachMultiWriterDisk(ctx context.Context, vm *object.VirtualMachine, sharedPath string, create bool) (diskKey int32, err error) {
	devs, err := vm.Device(ctx)
	if err != nil {
		return 0, fmt.Errorf("device list: %w", err)
	}
	scsiList := devs.SelectByType((*types.ParaVirtualSCSIController)(nil))
	if len(scsiList) == 0 {
		return 0, fmt.Errorf("vm %s has no ParaVirtualSCSIController", vm.Reference().Value)
	}
	scsi := scsiList[0].(types.BaseVirtualController)

	var capKB int64
	if create {
		capKB = int64(units.GB) / units.KB
	}
	disk := &types.VirtualDisk{
		CapacityInKB: capKB,
		VirtualDevice: types.VirtualDevice{
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: sharedPath,
				},
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: types.NewBool(true),
				Sharing:         string(types.VirtualDiskSharingSharingMultiWriter),
			},
		},
	}
	(object.VirtualDeviceList{disk}).AssignController(disk, scsi)

	cs, err := (object.VirtualDeviceList{disk}).ConfigSpec(types.VirtualDeviceConfigSpecOperationAdd)
	if err != nil {
		return 0, fmt.Errorf("config spec: %w", err)
	}

	task, err := vm.Reconfigure(ctx, types.VirtualMachineConfigSpec{DeviceChange: cs})
	if err != nil {
		return 0, fmt.Errorf("reconfigure: %w", err)
	}
	if err := task.Wait(ctx); err != nil {
		return 0, fmt.Errorf("reconfigure task: %w", err)
	}

	devs, err = vm.Device(ctx)
	if err != nil {
		return 0, fmt.Errorf("device list after reconfigure: %w", err)
	}
	wantSharing := string(types.VirtualDiskSharingSharingMultiWriter)
	for _, d := range devs.SelectByType((*types.VirtualDisk)(nil)) {
		vd := d.(*types.VirtualDisk)
		b, ok := vd.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
		if !ok {
			continue
		}
		if b.FileName == sharedPath && b.Sharing == wantSharing {
			return vd.Key, nil
		}
	}
	return 0, fmt.Errorf("multi-writer disk with path %q not found on vm %s", sharedPath, vm.Reference().Value)
}
