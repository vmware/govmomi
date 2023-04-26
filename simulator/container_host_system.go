/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"fmt"
	"strings"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

const KiB = 1024
const MiB = 1024 * KiB
const GiB = 1024 * MiB
const TiB = 1024 * GiB
const Pib = 1024 * TiB

const KB = 1000
const MB = 1000 * KB
const GB = 1000 * MB
const TB = 1000 * GB
const PB = 1000 * TB

type simHost struct {
	host *HostSystem
	c    *container
}

// createSimulationHost inspects the provided HostSystem and creates a simHost binding for it if
// the vm.Config.ExtraConfig set contains a key "RUN.container".
// If the ExtraConfig set does not contain that key, this returns nil.
// Methods on the simHost type are written to check for nil object so the return from this call can be blindly
// assigned and invoked without the caller caring about whether a binding for a backing container was warranted.
func createSimulationHost(ctx *Context, host *HostSystem) (*simHost, error) {
	sh := &simHost{
		host: host,
	}

	advOpts := ctx.Map.Get(host.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	fault := advOpts.QueryOptions(&types.QueryOptions{Name: "RUN.container"}).(*methods.QueryOptionsBody).Fault()
	if fault != nil {
		if _, ok := fault.VimFault().(*types.InvalidName); ok {
			return nil, nil
		}
		return nil, fmt.Errorf("errror retrieving container backing from host config manager: %+v", fault.VimFault())
	}

	// assemble env
	var dockerEnv []string
	var dockerVol []string
	var dockerNet []string
	var symlinkCmds [][]string

	var err error

	hName := host.Summary.Config.Name
	hUuid := host.Summary.Hardware.Uuid
	containerName := constructContainerName(hName, hUuid)

	for i := range host.Config.FileSystemVolume.MountInfo {
		info := &host.Config.FileSystemVolume.MountInfo[i]
		name := info.Volume.GetHostFileSystemVolume().Name

		// NOTE: if we ever need persistence cross-invocation we can look at encoding the disk info as a label
		labels := []string{"name=" + name, "container=" + containerName, deleteWithContainer}
		dockerUuid, err := createVolume("", labels, nil)
		if err != nil {
			return nil, err
		}

		uuid := volumeIDtoHostVolumeUUID(dockerUuid)
		name = strings.Replace(name, uuidToken, uuid, -1)

		switch vol := info.Volume.(type) {
		case *types.HostVmfsVolume:
			vol.BlockSizeMb = 1
			vol.BlockSize = KiB
			vol.UnmapGranularity = KiB
			vol.UnmapPriority = "low"
			vol.MajorVersion = 6
			vol.Version = "6.82"
			vol.Uuid = uuid
			vol.HostFileSystemVolume.Name = name
			for e := range vol.Extent {
				vol.Extent[e].DiskName = "____simulated_volume_____"
				if vol.Extent[e].Partition == 0 {
					// HACK: this should be unique within the diskname, but for now this will suffice
					//  partitions start at 1
					vol.Extent[e].Partition = int32(e + 1)
				}
			}
			vol.Ssd = types.NewBool(true)
			vol.Local = types.NewBool(true)
		case *types.HostVfatVolume:
			vol.HostFileSystemVolume.Name = name
		}

		info.VStorageSupport = "vStorageUnsupported"

		info.MountInfo.Path = "/vmfs/volumes/" + uuid
		info.MountInfo.Mounted = types.NewBool(true)
		info.MountInfo.Accessible = types.NewBool(true)
		if info.MountInfo.AccessMode == "" {
			info.MountInfo.AccessMode = "readWrite"
		}

		opt := "rw"
		if info.MountInfo.AccessMode == "readOnly" {
			opt = "ro"
		}

		dockerVol = append(dockerVol, fmt.Sprintf("%s:/vmfs/volumes/%s:%s", dockerUuid, uuid, opt))
		symlinkCmds = append(symlinkCmds, []string{"ln", "-s", fmt.Sprintf("/vmfs/volumes/%s", uuid), fmt.Sprintf("/vmfs/volumes/%s", name)})
		if strings.HasPrefix(name, "OSDATA") {
			symlinkCmds = append(symlinkCmds, []string{"mkdir", "-p", "/var/lib/vmware"})
			symlinkCmds = append(symlinkCmds, []string{"ln", "-s", fmt.Sprintf("/vmfs/volumes/%s", uuid), "/var/lib/vmware/osdata"})
		}
	}

	// TODO: extract the underlay's from a topology config
	// create a bridge for each broadcast domain a pnic is connected to
	dockerNet = append(dockerNet, defaultUnderlayBridgeName)

	// TODO: add in vSwitches if we know them at this point

	// - a pnic does not have an IP so this is purely a connectivity statement, not a network identity
	// ? how is this underlay topology expressed? Initially we can assume a flat topology with all hosts on the same broadcast domain

	// if there's a DVS that doesn't have a bridge, create the bridge

	sh.c, err = create(ctx, hName, hUuid, dockerNet, dockerVol, nil, dockerEnv, "alpine", []string{"sleep", "infinity"})
	if err != nil {
		return nil, err
	}

	err = sh.c.start(ctx)
	if err != nil {
		return nil, err
	}

	// create symlinks from /vmfs/volumes/ for the Volume Name - the direct mount (path) is only the uuid
	// ? can we do this via a script in the ESX image? are the volume names exposed in any manner instead the host? They must be because these mounts exist
	// but where does that come from? Chicken and egg problem? ConfigStore?
	for _, symlink := range symlinkCmds {
		_, err := sh.c.exec(ctx, symlink)
		if err != nil {
			return nil, err
		}
	}

	return sh, nil
}

// remove destroys the container associated with the host and any volumes with labels specifying their lifecycle
// is coupled with the container
func (sh *simHost) remove(ctx *Context) error {
	if sh == nil {
		return nil
	}

	return sh.c.remove(ctx)
}

// volumeIDtoHostVolumeUUID takes the 64 char docker uuid and converts it into a 32char ESX form of 8-8-4-12
// Perhaps we should do this using an md5 rehash, but instead we just take the first 32char for ease of cross-reference.
func volumeIDtoHostVolumeUUID(id string) string {
	return fmt.Sprintf("%s-%s-%s-%s", id[0:8], id[8:16], id[16:20], id[20:32])
}

// By reference to physical system, partition numbering tends to work out like this:
// 1. EFI System (100 MB)
// Free space (1.97 MB)
// 5. Basic Data (4 GB) (bootbank1)
// 6. Basic Data (4 GB) (bootbank2)
// 7. VMFSL (119.9 GB)  (os-data)
// 8. VMFS (1 TB)       (datastore1)
// I assume the jump from 1 -> 5 harks back to the primary/logical partitions from MBT days
const uuidToken = "%__UUID__%"

var defaultSimVolumes = []types.HostFileSystemMountInfo{
	{
		MountInfo: types.HostMountInfo{
			AccessMode: "readWrite",
		},
		Volume: &types.HostVmfsVolume{
			HostFileSystemVolume: types.HostFileSystemVolume{
				Type:     "VMFS",
				Name:     "datastore1",
				Capacity: 1 * TiB,
			},
			Extent: []types.HostScsiDiskPartition{
				{
					Partition: 8,
				},
			},
		},
	},
	{
		MountInfo: types.HostMountInfo{
			AccessMode: "readWrite",
		},
		Volume: &types.HostVmfsVolume{
			HostFileSystemVolume: types.HostFileSystemVolume{
				Type:     "OTHER",
				Name:     "OSDATA-%__UUID__%",
				Capacity: 128 * GiB,
			},
			Extent: []types.HostScsiDiskPartition{
				{
					Partition: 7,
				},
			},
		},
	},
	{
		MountInfo: types.HostMountInfo{
			AccessMode: "readOnly",
		},
		Volume: &types.HostVfatVolume{
			HostFileSystemVolume: types.HostFileSystemVolume{
				Type:     "OTHER",
				Name:     "BOOTBANK1",
				Capacity: 4 * GiB,
			},
		},
	},
	{
		MountInfo: types.HostMountInfo{
			AccessMode: "readOnly",
		},
		Volume: &types.HostVfatVolume{
			HostFileSystemVolume: types.HostFileSystemVolume{
				Type:     "OTHER",
				Name:     "BOOTBANK2",
				Capacity: 4 * GiB,
			},
		},
	},
}

const defaultUnderlayBridgeName = "vcsim-underlay"
