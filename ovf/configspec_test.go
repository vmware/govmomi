/*
Copyright (c) 2015-2024 VMware, Inc. All Rights Reserved.

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

package ovf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/types"
)

func TestEnvelopeToConfigSpec(t *testing.T) {
	e := testEnvelope(t, "fixtures/haproxy-vsphere.ovf")

	cs, err := e.ToConfigSpec()

	var w bytes.Buffer
	enc := types.NewJSONEncoder(&w)
	enc.SetIndent("", "    ")
	assert.NoError(t, enc.Encode(cs))
	t.Logf("\n\nconfigSpec=%s\n\n", w.String())

	assert.NoError(t, err)
	assert.NotEmpty(t, cs)

	assert.Equal(t, "haproxy", cs.Name)
	assert.Equal(t, int32(2), cs.NumCPUs)
	assert.Equal(t, int32(2), cs.NumCoresPerSocket)
	assert.Equal(t, int64(4096), cs.MemoryMB)
	assert.Equal(t, "vmx-13", cs.Version)

	if assert.Len(t, cs.DeviceChange, 21) {

		var scsiController1Key int32
		if d, ok := cs.DeviceChange[0].GetVirtualDeviceConfigSpec().Device.(*types.VirtualLsiLogicController); assert.True(t, ok) {
			scsiController1Key = d.Key
			assert.Equal(t, int32(0), d.BusNumber)
			if assert.NotNil(t, d.SlotInfo) {
				si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
				assert.True(t, ok)
				assert.Equal(t, int32(128), si.PciSlotNumber)
			}
		}

		if d, ok := cs.DeviceChange[1].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
			assert.Equal(t, scsiController1Key, d.ControllerKey)
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.True(t, ok)
			assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
			if assert.NotNil(t, db.ThinProvisioned) {
				assert.True(t, *db.ThinProvisioned)
			}
			assert.Equal(t, int64(20*1024*1024*1024), d.CapacityInBytes)
		}

		if bd, ok := cs.DeviceChange[2].GetVirtualDeviceConfigSpec().Device.(types.BaseVirtualEthernetCard); assert.True(t, ok) {
			d := bd.GetVirtualEthernetCard()
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(2), *d.UnitNumber)
			}
			if assert.NotNil(t, d.Connectable) {
				assert.True(t, d.Connectable.AllowGuestControl)
			}
			if assert.NotNil(t, d.WakeOnLanEnabled) {
				assert.False(t, *d.WakeOnLanEnabled)
			}
			if assert.NotNil(t, d.SlotInfo) {
				si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
				assert.True(t, ok)
				assert.Equal(t, int32(160), si.PciSlotNumber)
			}
		}

		if bd, ok := cs.DeviceChange[3].GetVirtualDeviceConfigSpec().Device.(types.BaseVirtualEthernetCard); assert.True(t, ok) {
			d := bd.GetVirtualEthernetCard()
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(3), *d.UnitNumber)
			}
			if assert.NotNil(t, d.Connectable) {
				assert.True(t, d.Connectable.AllowGuestControl)
			}
			if assert.NotNil(t, d.WakeOnLanEnabled) {
				assert.False(t, *d.WakeOnLanEnabled)
			}
			if assert.NotNil(t, d.SlotInfo) {
				si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
				assert.True(t, ok)
				assert.Equal(t, int32(192), si.PciSlotNumber)
			}
			if assert.NotNil(t, d.UptCompatibilityEnabled) {
				assert.False(t, *d.UptCompatibilityEnabled)
			}
		}

		if d, ok := cs.DeviceChange[4].GetVirtualDeviceConfigSpec().Device.(*types.VirtualMachineVideoCard); assert.True(t, ok) {
			if assert.NotNil(t, d.Enable3DSupport) {
				assert.False(t, *d.Enable3DSupport)
			}
			assert.Equal(t, int64(262144), d.GraphicsMemorySizeInKB)
			if assert.NotNil(t, d.UseAutoDetect) {
				assert.False(t, *d.UseAutoDetect)
			}
			assert.Equal(t, int64(4096), d.VideoRamSizeInKB)
			assert.Equal(t, int32(1), d.NumDisplays)
			assert.Equal(t, "automatic", d.Use3dRenderer)
		}

		var ideControllerKey int32
		if d, ok := cs.DeviceChange[5].GetVirtualDeviceConfigSpec().Device.(*types.VirtualIDEController); assert.True(t, ok) {
			ideControllerKey = d.Key
			assert.Equal(t, int32(1), d.BusNumber)
		}

		if d, ok := cs.DeviceChange[6].GetVirtualDeviceConfigSpec().Device.(*types.VirtualIDEController); assert.True(t, ok) {
			assert.Equal(t, int32(0), d.BusNumber)
		}

		if d, ok := cs.DeviceChange[7].GetVirtualDeviceConfigSpec().Device.(*types.VirtualCdrom); assert.True(t, ok) {
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			assert.Equal(t, ideControllerKey, d.ControllerKey)
		}

		var sioControllerKey int32
		if d, ok := cs.DeviceChange[8].GetVirtualDeviceConfigSpec().Device.(*types.VirtualSIOController); assert.True(t, ok) {
			sioControllerKey = d.Key
		}

		if d, ok := cs.DeviceChange[9].GetVirtualDeviceConfigSpec().Device.(*types.VirtualFloppy); assert.True(t, ok) {
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			assert.Equal(t, sioControllerKey, d.ControllerKey)
		}

		if d, ok := cs.DeviceChange[10].GetVirtualDeviceConfigSpec().Device.(*types.VirtualMachineVMCIDevice); assert.True(t, ok) {
			if assert.NotNil(t, d.AllowUnrestrictedCommunication) {
				assert.False(t, *d.AllowUnrestrictedCommunication)
			}
		}

		var scsiController2Key int32
		if d, ok := cs.DeviceChange[11].GetVirtualDeviceConfigSpec().Device.(*types.ParaVirtualSCSIController); assert.True(t, ok) {
			scsiController2Key = d.Key
			assert.Equal(t, int32(1), d.BusNumber)
			assert.Nil(t, d.SlotInfo)
		}

		if d, ok := cs.DeviceChange[12].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
			assert.Equal(t, scsiController2Key, d.ControllerKey)
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.True(t, ok)
			assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
			if assert.NotNil(t, db.ThinProvisioned) {
				assert.True(t, *db.ThinProvisioned)
			}
			assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
		}

		var sataController1Key int32
		if d, ok := cs.DeviceChange[13].GetVirtualDeviceConfigSpec().Device.(*types.VirtualAHCIController); assert.True(t, ok) {
			sataController1Key = d.Key
			assert.Equal(t, int32(0), d.BusNumber)
			assert.Nil(t, d.SlotInfo)
		}

		if d, ok := cs.DeviceChange[14].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
			assert.Equal(t, sataController1Key, d.ControllerKey)
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.True(t, ok)
			assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
			if assert.NotNil(t, db.ThinProvisioned) {
				assert.True(t, *db.ThinProvisioned)
			}
			assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
		}

		var sataController2Key int32
		if d, ok := cs.DeviceChange[15].GetVirtualDeviceConfigSpec().Device.(*types.VirtualAHCIController); assert.True(t, ok) {
			sataController2Key = d.Key
			assert.Equal(t, int32(1), d.BusNumber)
			assert.Nil(t, d.SlotInfo)
		}

		if d, ok := cs.DeviceChange[16].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
			assert.Equal(t, sataController2Key, d.ControllerKey)
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.True(t, ok)
			assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
			if assert.NotNil(t, db.ThinProvisioned) {
				assert.True(t, *db.ThinProvisioned)
			}
			assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
		}

		var nvmeController1Key int32
		if d, ok := cs.DeviceChange[17].GetVirtualDeviceConfigSpec().Device.(*types.VirtualNVMEController); assert.True(t, ok) {
			nvmeController1Key = d.Key
			assert.Equal(t, int32(0), d.BusNumber)
			assert.Nil(t, d.SlotInfo)
		}

		if d, ok := cs.DeviceChange[18].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
			assert.Equal(t, nvmeController1Key, d.ControllerKey)
			if assert.NotNil(t, d.UnitNumber) {
				assert.Equal(t, int32(0), *d.UnitNumber)
			}
			db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.True(t, ok)
			assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
			if assert.NotNil(t, db.ThinProvisioned) {
				assert.True(t, *db.ThinProvisioned)
			}
			assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
		}

		if d, ok := cs.DeviceChange[19].GetVirtualDeviceConfigSpec().Device.(*types.VirtualUSBController); assert.True(t, ok) {
			if assert.NotNil(t, d.AutoConnectDevices) {
				assert.True(t, *d.AutoConnectDevices)
			}
			if assert.NotNil(t, d.EhciEnabled) {
				assert.False(t, *d.EhciEnabled)
			}
		}

		if d, ok := cs.DeviceChange[20].GetVirtualDeviceConfigSpec().Device.(*types.VirtualUSBXHCIController); assert.True(t, ok) {
			if assert.NotNil(t, d.AutoConnectDevices) {
				assert.True(t, *d.AutoConnectDevices)
			}
		}
	}

	assert.ElementsMatch(t, cs.ExtraConfig, []types.BaseOptionValue{
		&types.OptionValue{
			Key:   "guest_rpc.auth.cloud-init.set",
			Value: "FALSE",
		},
	})

	if assert.NotNil(t, cs.Flags) {
		if assert.NotNil(t, cs.Flags.VbsEnabled) {
			assert.False(t, *cs.Flags.VbsEnabled)
		}
		if assert.NotNil(t, cs.Flags.VvtdEnabled) {
			assert.False(t, *cs.Flags.VvtdEnabled)
		}
	}

	assert.Equal(t, "bios", cs.Firmware)

	if assert.NotNil(t, cs.BootOptions) {
		if assert.NotNil(t, cs.BootOptions.EfiSecureBootEnabled) {
			assert.False(t, *cs.BootOptions.EfiSecureBootEnabled)
		}
	}

	if assert.NotNil(t, cs.CpuHotAddEnabled) {
		assert.False(t, *cs.CpuHotAddEnabled)
	}
	if assert.NotNil(t, cs.CpuHotRemoveEnabled) {
		assert.False(t, *cs.CpuHotRemoveEnabled)
	}
	if assert.NotNil(t, cs.MemoryHotAddEnabled) {
		assert.False(t, *cs.MemoryHotAddEnabled)
	}
	if assert.NotNil(t, cs.NestedHVEnabled) {
		assert.False(t, *cs.NestedHVEnabled)
	}
	if assert.NotNil(t, cs.VirtualICH7MPresent) {
		assert.False(t, *cs.VirtualICH7MPresent)
	}

	assert.Equal(t, int32(1), cs.SimultaneousThreads)

	if assert.NotNil(t, cs.VPMCEnabled) {
		assert.False(t, *cs.VPMCEnabled)
	}

	if assert.NotNil(t, cs.CpuAllocation) {
		if assert.NotNil(t, cs.CpuAllocation.Shares) {
			assert.Equal(t, &types.SharesInfo{
				Shares: 2000,
				Level:  types.SharesLevelNormal,
			}, cs.CpuAllocation.Shares)
		}
	}

	if assert.NotNil(t, cs.PowerOpInfo) {
		assert.Equal(t, &types.VirtualMachineDefaultPowerOpInfo{
			PowerOffType:  "soft",
			ResetType:     "soft",
			SuspendType:   "hard",
			StandbyAction: "checkpoint",
		}, cs.PowerOpInfo)
	}

	if assert.NotNil(t, cs.Tools) {
		assert.Equal(t, &types.ToolsConfigInfo{
			SyncTimeWithHost:        types.NewBool(false),
			SyncTimeWithHostAllowed: types.NewBool(true),
			AfterPowerOn:            types.NewBool(true),
			AfterResume:             types.NewBool(true),
			BeforeGuestShutdown:     types.NewBool(true),
			BeforeGuestStandby:      types.NewBool(true),
			ToolsUpgradePolicy:      "manual",
		}, cs.Tools)
	}

	if va, ok := cs.VAppConfig.(*types.VAppConfigSpec); assert.True(t, ok) {
		if assert.Len(t, va.Product, 4) {
			assert.ElementsMatch(t,
				[]types.VAppProductSpec{
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppProductInfo{
							Key:         0,
							Name:        "HAProxy for the Load Balancer API v0.2.0",
							Vendor:      "VMware Inc.",
							Version:     "v0.2.0",
							FullVersion: "v0.2.0",
							ProductUrl:  "https://vmware.com",
							VendorUrl:   "https://vmware.com",
						},
					},
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppProductInfo{
							Key:     1,
							ClassId: "appliance",
						},
					},
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppProductInfo{
							Key:     2,
							ClassId: "network",
						},
					},
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppProductInfo{
							Key:     3,
							ClassId: "loadbalance",
						},
					},
				},
				va.Product,
			)
		}
		if assert.Len(t, va.Property, 6) {
			assert.ElementsMatch(t,
				[]types.VAppPropertySpec{

					// default
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              0,
							Category:         "Load Balancer API",
							Id:               "BUILD_TIMESTAMP",
							Type:             "string",
							UserConfigurable: types.NewBool(false),
							DefaultValue:     "1615488399",
						},
					},
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              1,
							Category:         "Load Balancer API",
							Id:               "BUILD_DATE",
							Type:             "string",
							UserConfigurable: types.NewBool(false),
							DefaultValue:     "2021-03-11T18:46:39Z",
						},
					},

					// appliance
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              2,
							Category:         "1. Appliance Configuration",
							ClassId:          "appliance",
							Id:               "root_pwd",
							Label:            "1.1. Root Password",
							Type:             "string",
							UserConfigurable: types.NewBool(true),
							Description:      "The initial password of the root user. Subsequent changes of password should be performed in operating system. (6-128 characters)",
						},
					},

					// network
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              3,
							Category:         "2. Network Config",
							ClassId:          "network",
							Id:               "hostname",
							Label:            "2.1. Host Name",
							Type:             "string",
							UserConfigurable: types.NewBool(true),
							DefaultValue:     "haproxy.local",
							Description:      "The host name. A fully-qualified domain name is also supported.",
						},
					},

					// loadbalance
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              4,
							Category:         "3. Load Balancing",
							ClassId:          "loadbalance",
							Id:               "service_ip_range",
							Label:            "3.1. Load Balancer IP Ranges, comma-separated in CIDR format (Eg 1.2.3.4/28,5.6.7.8/28)",
							Type:             "string",
							UserConfigurable: types.NewBool(true),
							Description:      "The IP ranges the load balancer will use for Kubernetes Services and Control Planes. The Appliance will currently respond to ALL the IPs in these ranges whether they're assigned or not. As such, these ranges must not overlap with the IPs assigned for the appliance or any other VMs on the network.",
						},
					},
					{
						ArrayUpdateSpec: types.ArrayUpdateSpec{
							Operation: types.ArrayUpdateOperationAdd,
						},
						Info: &types.VAppPropertyInfo{
							Key:              5,
							Category:         "3. Load Balancing",
							ClassId:          "loadbalance",
							Id:               "dataplane_port",
							Label:            "3.2. Dataplane API Management Port",
							Type:             "int",
							UserConfigurable: types.NewBool(true),
							DefaultValue:     "5556",
							Description:      "Specifies the port on which the Dataplane API will be advertized on the Management Network.",
						},
					},
				},
				va.Property,
			)
		}
	}
}
