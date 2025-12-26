// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import "github.com/vmware/govmomi/vim25/types"

const (
	VirtualMachineDefaultDevicePCIControllerKey = int32(100)
	VirtualMachineDefaultDevicePS2ControllerKey = int32(300)
)

var VirtualMachineDefaultDevicePCIController = &types.VirtualPCIController{
	VirtualController: types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: VirtualMachineDefaultDevicePCIControllerKey,
			DeviceInfo: &types.Description{
				Label:   "PCI controller 0",
				Summary: "PCI controller 0",
			},
		},
		Device: []int32{
			VirtualMachineDefaultDeviceVideoCard.Key,
			VirtualMachineDefaultDeviceVMCIDevice.Key,
		},
	},
}

var VirtualMachineDefaultDeviceIDEControllerBus0 = &types.VirtualIDEController{
	VirtualController: types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: 200,
			DeviceInfo: &types.Description{
				Label:   "IDE 0",
				Summary: "IDE 0",
			},
		},
		BusNumber: 0,
	},
}

var VirtualMachineDefaultDeviceIDEControllerBus1 = &types.VirtualIDEController{
	VirtualController: types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: 201,
			DeviceInfo: &types.Description{
				Label:   "IDE 1",
				Summary: "IDE 1",
			},
		},
		BusNumber: 1,
	},
}

var VirtualMachineDefaultDevicePS2Controller = &types.VirtualPS2Controller{
	VirtualController: types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: VirtualMachineDefaultDevicePS2ControllerKey,
			DeviceInfo: &types.Description{
				Label:   "PS2 controller 0",
				Summary: "PS2 controller 0",
			},
		},
		Device: []int32{
			VirtualMachineDefaultDeviceVirtualKeyboard.Key,
			VirtualMachineDefaultDeviceVirtualPointingDevice.Key,
		},
	},
}

var VirtualMachineDefaultDeviceSIOController = &types.VirtualSIOController{
	VirtualController: types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: 400,
			DeviceInfo: &types.Description{
				Label:   "SIO controller 0",
				Summary: "SIO controller 0",
			},
		},
	},
}
var VirtualMachineDefaultDeviceVirtualKeyboard = &types.VirtualKeyboard{
	VirtualDevice: types.VirtualDevice{
		Key: 600,
		DeviceInfo: &types.Description{
			Label:   "Keyboard ",
			Summary: "Keyboard",
		},
		ControllerKey: VirtualMachineDefaultDevicePS2ControllerKey,
		UnitNumber:    types.NewInt32(0),
	},
}
var VirtualMachineDefaultDeviceVirtualPointingDevice = &types.VirtualPointingDevice{
	VirtualDevice: types.VirtualDevice{
		Key: 700,
		DeviceInfo: &types.Description{
			Label:   "Pointing device",
			Summary: "Pointing device; Device",
		},
		Backing: &types.VirtualPointingDeviceDeviceBackingInfo{
			VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
				UseAutoDetect: types.NewBool(false),
			},
			HostPointingDevice: "autodetect",
		},
		ControllerKey: VirtualMachineDefaultDevicePS2ControllerKey,
		UnitNumber:    types.NewInt32(1),
	},
}
var VirtualMachineDefaultDeviceVideoCard = &types.VirtualMachineVideoCard{
	VirtualDevice: types.VirtualDevice{
		Key: 500,
		DeviceInfo: &types.Description{
			Label:   "Video card ",
			Summary: "Video card",
		},
		ControllerKey: VirtualMachineDefaultDevicePCIControllerKey,
		UnitNumber:    types.NewInt32(0),
	},
	VideoRamSizeInKB:       4096,
	NumDisplays:            1,
	UseAutoDetect:          types.NewBool(false),
	Enable3DSupport:        types.NewBool(false),
	Use3dRenderer:          "automatic",
	GraphicsMemorySizeInKB: 262144,
}

var VirtualMachineDefaultDeviceVMCIDevice = &types.VirtualMachineVMCIDevice{
	VirtualDevice: types.VirtualDevice{
		Key: 12000,
		DeviceInfo: &types.Description{
			Label:   "VMCI device",
			Summary: "Device on the virtual machine PCI bus that provides support for the virtual machine communication interface",
		},
		ControllerKey: VirtualMachineDefaultDevicePCIControllerKey,
		UnitNumber:    types.NewInt32(17),
	},
	Id:                             -1,
	AllowUnrestrictedCommunication: types.NewBool(false),
	FilterEnable:                   types.NewBool(true),
}

// VirtualDevice is the default set of VirtualDevice types created for a VirtualMachine
// Capture method:
//
//	govc vm.create foo
//	govc object.collect -s -dump vm/foo config.hardware.device
var VirtualDevice = []types.BaseVirtualDevice{
	VirtualMachineDefaultDevicePCIController,
	VirtualMachineDefaultDeviceIDEControllerBus0,
	VirtualMachineDefaultDeviceIDEControllerBus1,
	VirtualMachineDefaultDevicePS2Controller,
	VirtualMachineDefaultDeviceSIOController,
	VirtualMachineDefaultDeviceVirtualKeyboard,
	VirtualMachineDefaultDeviceVirtualPointingDevice,
	VirtualMachineDefaultDeviceVideoCard,
	VirtualMachineDefaultDeviceVMCIDevice,
}

// EthernetCard template for types.VirtualEthernetCard
var EthernetCard = types.VirtualE1000{
	VirtualEthernetCard: types.VirtualEthernetCard{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         4000,
			Backing: &types.VirtualEthernetCardNetworkBackingInfo{
				VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
					VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
					DeviceName:               "VM Network",
					UseAutoDetect:            types.NewBool(false),
				},
				Network:           (*types.ManagedObjectReference)(nil),
				InPassthroughMode: types.NewBool(false),
			},
			Connectable: &types.VirtualDeviceConnectInfo{
				DynamicData:       types.DynamicData{},
				StartConnected:    true,
				AllowGuestControl: true,
				Connected:         false,
				Status:            "untried",
			},
			SlotInfo: &types.VirtualDevicePciBusSlotInfo{
				VirtualDeviceBusSlotInfo: types.VirtualDeviceBusSlotInfo{},
				PciSlotNumber:            32,
			},
			ControllerKey: 100,
			UnitNumber:    types.NewInt32(7),
		},
		AddressType:      "generated",
		MacAddress:       "",
		WakeOnLanEnabled: types.NewBool(true),
	},
}
