/*
Copyright (c) 2017-2023 VMware, Inc. All Rights Reserved.

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

package esx

import (
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

// HostConfigInfo is the default template for the HostSystem config property.
// Capture method:
// govc object.collect -s -dump HostSystem:ha-host config.fileSystemVolume
//   - slightly modified for uuids and DiskName
var HostFileSystemVolumeInfo = types.HostFileSystemVolumeInfo{
	VolumeTypeList: []string{"VMFS", "NFS", "NFS41", "vsan", "VVOL", "VFFS", "OTHER", "PMEM"},
	MountInfo: []types.HostFileSystemMountInfo{
		{
			MountInfo: types.HostMountInfo{
				Path:               "/vmfs/volumes/deadbeef-01234567-89ab-cdef00000003",
				AccessMode:         "readWrite",
				Mounted:            types.NewBool(true),
				Accessible:         types.NewBool(true),
				InaccessibleReason: "",
				MountFailedReason:  "",
			},
			Volume: &types.HostVmfsVolume{
				HostFileSystemVolume: types.HostFileSystemVolume{
					Type:     "VMFS",
					Name:     "datastore1",
					Capacity: 3.5 * units.TB,
				},
				BlockSizeMb:        1,
				BlockSize:          units.KB,
				UnmapGranularity:   units.KB,
				UnmapPriority:      "low",
				UnmapBandwidthSpec: (*types.VmfsUnmapBandwidthSpec)(nil),
				MaxBlocks:          61 * units.MB,
				MajorVersion:       6,
				Version:            "6.82",
				Uuid:               "deadbeef-01234567-89ab-cdef00000003",
				Extent: []types.HostScsiDiskPartition{
					{
						DiskName:  "____simulated_volumes_____",
						Partition: 8,
					},
				},
				VmfsUpgradable:   false,
				ForceMountedInfo: (*types.HostForceMountedInfo)(nil),
				Ssd:              types.NewBool(true),
				Local:            types.NewBool(true),
				ScsiDiskType:     "",
			},
			VStorageSupport: "vStorageUnsupported",
		},
		{
			MountInfo: types.HostMountInfo{
				Path:               "/vmfs/volumes/deadbeef-01234567-89ab-cdef00000002",
				AccessMode:         "readWrite",
				Mounted:            types.NewBool(true),
				Accessible:         types.NewBool(true),
				InaccessibleReason: "",
				MountFailedReason:  "",
			},
			Volume: &types.HostVmfsVolume{
				HostFileSystemVolume: types.HostFileSystemVolume{
					Type:     "OTHER",
					Name:     "OSDATA-deadbeef-01234567-89ab-cdef00000002",
					Capacity: 128 * units.GB,
				},
				BlockSizeMb:        1,
				BlockSize:          units.KB,
				UnmapGranularity:   0,
				UnmapPriority:      "",
				UnmapBandwidthSpec: (*types.VmfsUnmapBandwidthSpec)(nil),
				MaxBlocks:          256 * units.KB,
				MajorVersion:       1,
				Version:            "1.00",
				Uuid:               "deadbeef-01234567-89ab-cdef00000002",
				Extent: []types.HostScsiDiskPartition{
					{
						DiskName:  "____simulated_volumes_____",
						Partition: 7,
					},
				},
				VmfsUpgradable:   false,
				ForceMountedInfo: (*types.HostForceMountedInfo)(nil),
				Ssd:              types.NewBool(true),
				Local:            types.NewBool(true),
				ScsiDiskType:     "",
			},
			VStorageSupport: "vStorageUnsupported",
		},
		{
			MountInfo: types.HostMountInfo{
				Path:               "/vmfs/volumes/deadbeef-01234567-89ab-cdef00000001",
				AccessMode:         "readOnly",
				Mounted:            types.NewBool(true),
				Accessible:         types.NewBool(true),
				InaccessibleReason: "",
				MountFailedReason:  "",
			},
			Volume: &types.HostVfatVolume{
				HostFileSystemVolume: types.HostFileSystemVolume{
					Type:     "OTHER",
					Name:     "BOOTBANK1",
					Capacity: 4 * units.GB,
				},
			},
			VStorageSupport: "",
		},
		{
			MountInfo: types.HostMountInfo{
				Path:               "/vmfs/volumes/deadbeef-01234567-89ab-cdef00000000",
				AccessMode:         "readOnly",
				Mounted:            types.NewBool(true),
				Accessible:         types.NewBool(true),
				InaccessibleReason: "",
				MountFailedReason:  "",
			},
			Volume: &types.HostVfatVolume{
				HostFileSystemVolume: types.HostFileSystemVolume{
					Type:     "OTHER",
					Name:     "BOOTBANK2",
					Capacity: 4 * units.GB,
				},
			},
			VStorageSupport: "",
		},
	},
}
