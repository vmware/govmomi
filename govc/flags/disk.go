/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package flags

import (
	"flag"
	"path/filepath"

	"github.com/vmware/govmomi/vim25/types"
)

type DiskFlag struct {
	*DatastoreFlag

	name string
	path string
}

func (f *DiskFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "disk", "", "Disk path name")
}

func (f *DiskFlag) Process() error {
	return nil
}

func (f *DiskFlag) IsSet() bool {
	return f.name != ""
}

func (f *DiskFlag) Path() (string, error) {
	return f.DatastorePath(f.name)
}

func (f *DiskFlag) Copy(name string) (types.BaseVirtualDevice, error) {
	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	src, err := f.Path()
	if err != nil {
		return nil, err
	}

	// TODO: support cross-datacenter
	dst, err := f.DatastorePath(name)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(dst)
	// ignore mkdir errors, copy will error if dir does not exist
	_ = c.FileManager().MakeDirectory(dir, dc, false)

	// TODO: adpater, type options
	spec := &types.VirtualDiskSpec{
		AdapterType: "lsiLogic",
		DiskType:    "thin",
	}

	err = c.VirtualDiskManager().CopyVirtualDisk(c, src, dc, dst, dc, spec, false)
	if err != nil {
		return nil, err
	}

	return &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Key:           -1,
			ControllerKey: -1,
			UnitNumber:    -1,
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: dst,
				},
			},
		},
	}, nil
}

func (f *DiskFlag) Controller() (types.BaseVirtualDevice, error) {
	// TODO: adapter option
	return &types.VirtualLsiLogicController{types.VirtualSCSIController{
		SharedBus: types.VirtualSCSISharingNoSharing,
		VirtualController: types.VirtualController{
			BusNumber: 0,
			VirtualDevice: types.VirtualDevice{
				Key: -1,
			},
		},
	}}, nil
}

func (f *DiskFlag) Cdrom(path string) (types.BaseVirtualDevice, error) {
	return &types.VirtualCdrom{
		VirtualDevice: types.VirtualDevice{
			Key:           -2,
			ControllerKey: 200,
			UnitNumber:    -1,
			Backing: &types.VirtualCdromIsoBackingInfo{
				types.VirtualDeviceFileBackingInfo{
					FileName: path,
				},
			},
		},
	}, nil
}
