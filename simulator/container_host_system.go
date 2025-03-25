// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"strings"

	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	advOptPrefixPnicToUnderlayPrefix = "RUN.underlay."
	advOptContainerBackingImage      = "RUN.container"
	defaultUnderlayBridgeName        = "vcsim-underlay"
)

type simHost struct {
	host *HostSystem
	c    *container
}

// createSimHostMounts iterates over the provide filesystem mount info, creating docker volumes. It does _not_ delete volumes
// already created if creation of one fails.
// Returns:
// volume mounts: mount options suitable to pass directly to docker
// exec commands: a set of commands to run in the sim host after creation
// error: if construction of the above outputs fails
func createSimHostMounts(ctx *Context, containerName string, mounts []types.HostFileSystemMountInfo) ([]string, [][]string, error) {
	var dockerVol []string
	var symlinkCmds [][]string

	for i := range mounts {
		info := &mounts[i]
		name := info.Volume.GetHostFileSystemVolume().Name

		// NOTE: if we ever need persistence cross-invocation we can look at encoding the disk info as a label
		labels := []string{"name=" + name, "container=" + containerName, deleteWithContainer}
		dockerUuid, err := createVolume("", labels, nil)
		if err != nil {
			return nil, nil, err
		}

		uuid := volumeIDtoHostVolumeUUID(dockerUuid)
		name = strings.Replace(name, uuidToken, uuid, -1)

		switch vol := info.Volume.(type) {
		case *types.HostVmfsVolume:
			vol.BlockSizeMb = 1
			vol.BlockSize = units.KB
			vol.UnmapGranularity = units.KB
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

		// create symlinks from /vmfs/volumes/ for the Volume Name - the direct mount (path) is only the uuid
		// ? can we do this via a script in the ESX image instead of via exec?
		// ? are the volume names exposed in any manner inside the host? They must be because these mounts exist but where does that come from? Chicken and egg problem? ConfigStore?
		symlinkCmds = append(symlinkCmds, []string{"ln", "-s", fmt.Sprintf("/vmfs/volumes/%s", uuid), fmt.Sprintf("/vmfs/volumes/%s", name)})
		if strings.HasPrefix(name, "OSDATA") {
			symlinkCmds = append(symlinkCmds, []string{"mkdir", "-p", "/var/lib/vmware"})
			symlinkCmds = append(symlinkCmds, []string{"ln", "-s", fmt.Sprintf("/vmfs/volumes/%s", uuid), "/var/lib/vmware/osdata"})
		}
	}

	return dockerVol, symlinkCmds, nil
}

// createSimHostNetworks creates the networks for the host if not already created. Because we expect multiple hosts on the same network to act as a cluster
// it's likely that only the first host will create networks.
// This includes:
// * bridge network per-pNIC
// * bridge network per-DVS
//
// Returns:
// * array of networks to attach to
// * array of commands to run
// * error
//
// TODO: implement bridge network per DVS - not needed until container backed VMs are "created" on container backed "hosts"
func createSimHostNetworks(ctx *Context, containerName string, networkInfo *types.HostNetworkInfo, advOpts *OptionManager) ([]string, [][]string, error) {
	var dockerNet []string
	var cmds [][]string

	existingNets := make(map[string]string)

	// a pnic does not have an IP so this is purely a connectivity statement, not a network identity, however this is not how docker works
	// so we're going to end up with a veth (our pnic) that does have an IP assigned. That IP will end up being used in a NetConfig structure associated
	// with the pNIC. See HostSystem.getNetConfigInterface.
	for i := range networkInfo.Pnic {
		pnicName := networkInfo.Pnic[i].Device

		bridge := getPnicUnderlay(advOpts, pnicName)

		if pnic, attached := existingNets[bridge]; attached {
			return nil, nil, fmt.Errorf("cannot attach multiple pNICs to the same underlay: %s and %s both attempting to connect to %s for %s", pnic, pnicName, bridge, containerName)
		}

		_, err := createBridge(bridge)
		if err != nil {
			return nil, nil, err
		}

		dockerNet = append(dockerNet, bridge)
		existingNets[bridge] = pnicName
	}

	return dockerNet, cmds, nil
}

func getPnicUnderlay(advOpts *OptionManager, pnicName string) string {
	queryRes := advOpts.QueryOptions(&types.QueryOptions{Name: advOptPrefixPnicToUnderlayPrefix + pnicName}).(*methods.QueryOptionsBody).Res
	return queryRes.Returnval[0].GetOptionValue().Value.(string)
}

// createSimulationHostcreates a simHost binding if the host.ConfigManager.AdvancedOption set contains a key "RUN.container".
// If the set does not contain that key, this returns nil.
// Methods on the simHost type are written to check for nil object so the return from this call can be blindly
// assigned and invoked without the caller caring about whether a binding for a backing container was warranted.
//
// The created simhost is based off of the details of the supplied host system.
// VMFS locations are created based on FileSystemMountInfo
// Bridge networks are created to simulate underlay networks - one per pNIC. You cannot connect two pNICs to the same underlay.
//
// On Network connectivity - initially this is using docker network constructs. This means we cannot easily use nested "ip netns" so we cannot
// have a perfect representation of the ESX structure: pnic(veth)->vswtich(bridge)->{vmk,vnic}(veth)
// Instead we have the following:
// * bridge network per underlay - everything connects directly to the underlay
// * VMs/CRXs connect to the underlay dictated by the Uplink pNIC attached to their vSwitch
// * hostd vmknic gets the "host" container IP - we don't currently support multiple vmknics with different IPs
// * no support for mocking VLANs
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

	var execCmds [][]string

	var err error

	hName := host.Summary.Config.Name
	hUuid := host.Summary.Hardware.Uuid
	containerName := constructContainerName(hName, hUuid)

	// create volumes and mounts
	dockerVol, volCmds, err := createSimHostMounts(ctx, containerName, host.Config.FileSystemVolume.MountInfo)
	if err != nil {
		return nil, err
	}
	execCmds = append(execCmds, volCmds...)

	// create networks
	dockerNet, netCmds, err := createSimHostNetworks(ctx, containerName, host.Config.Network, advOpts)
	if err != nil {
		return nil, err
	}
	execCmds = append(execCmds, netCmds...)

	// create the container
	sh.c, err = create(ctx, hName, hUuid, dockerNet, dockerVol, nil, dockerEnv, "alpine:3.20.3", []string{"sleep", "infinity"})
	if err != nil {
		return nil, err
	}

	// start the container
	err = sh.c.start(ctx)
	if err != nil {
		return nil, err
	}

	// run post-creation steps
	for _, cmd := range execCmds {
		_, err := sh.c.exec(ctx, cmd)
		if err != nil {
			return nil, err
		}
	}

	_, detail, err := sh.c.inspect()
	if err != nil {
		return nil, err
	}
	for i := range host.Config.Network.Pnic {
		pnic := &host.Config.Network.Pnic[i]
		bridge := getPnicUnderlay(advOpts, pnic.Device)
		settings := detail.NetworkSettings.Networks[bridge]

		// it doesn't really make sense at an ESX level to set this information as IP bindings are associated with
		// vnics (VMs) or vmknics (daemons such as hostd).
		// However it's a useful location to stash this info in a manner that can be retrieved at a later date.
		pnic.Spec.Ip.IpAddress = settings.IPAddress
		pnic.Spec.Ip.SubnetMask = prefixToMask(settings.IPPrefixLen)

		pnic.Mac = settings.MacAddress
	}

	// update the active "management" nicType with the container IP for vmnic0
	netconfig, err := host.getNetConfigInterface(ctx, "management")
	if err != nil {
		return nil, err
	}
	netconfig.vmk.Spec.Ip.IpAddress = netconfig.uplink.Spec.Ip.IpAddress
	netconfig.vmk.Spec.Ip.SubnetMask = netconfig.uplink.Spec.Ip.SubnetMask
	netconfig.vmk.Spec.Mac = netconfig.uplink.Mac

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
				Capacity: 1 * units.TB,
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
				Capacity: 128 * units.GB,
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
				Capacity: 4 * units.GB,
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
				Capacity: 4 * units.GB,
			},
		},
	},
}
