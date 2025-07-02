// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestVirtualMachineConfigSpec(t *testing.T) {

	t.Run("marshal to xml", func(t *testing.T) {
		spec := VirtualMachineConfigSpec{
			Name:     "vm-001",
			GuestId:  "otherGuest",
			Files:    &VirtualMachineFileInfo{VmPathName: "[datastore1]"},
			NumCPUs:  1,
			MemoryMB: 128,
			DeviceChange: []BaseVirtualDeviceConfigSpec{
				&VirtualDeviceConfigSpec{
					Operation: VirtualDeviceConfigSpecOperationAdd,
					Device: &VirtualLsiLogicController{VirtualSCSIController{
						SharedBus: VirtualSCSISharingNoSharing,
						VirtualController: VirtualController{
							BusNumber: 0,
							VirtualDevice: VirtualDevice{
								Key: 1000,
							},
						},
					}},
				},
				&VirtualDeviceConfigSpec{
					Operation:     VirtualDeviceConfigSpecOperationAdd,
					FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
					Device: &VirtualDisk{
						VirtualDevice: VirtualDevice{
							Key:           0,
							ControllerKey: 1000,
							UnitNumber:    new(int32), // zero default value
							Backing: &VirtualDiskFlatVer2BackingInfo{
								DiskMode:        string(VirtualDiskModePersistent),
								ThinProvisioned: NewBool(true),
								VirtualDeviceFileBackingInfo: VirtualDeviceFileBackingInfo{
									FileName: "[datastore1]",
								},
							},
						},
						CapacityInKB: 4000000,
					},
				},
				&VirtualDeviceConfigSpec{
					Operation: VirtualDeviceConfigSpecOperationAdd,
					Device: &VirtualE1000{VirtualEthernetCard{
						VirtualDevice: VirtualDevice{
							Key: 0,
							DeviceInfo: &Description{
								Label:   "Network Adapter 1",
								Summary: "VM Network",
							},
							Backing: &VirtualEthernetCardNetworkBackingInfo{
								VirtualDeviceDeviceBackingInfo: VirtualDeviceDeviceBackingInfo{
									DeviceName: "VM Network",
								},
							},
						},
						AddressType: string(VirtualEthernetCardMacTypeGenerated),
					}},
				},
			},
			ExtraConfig: []BaseOptionValue{
				&OptionValue{Key: "bios.bootOrder", Value: "ethernet0"},
			},
		}

		_, err := xml.MarshalIndent(spec, "", " ")
		assert.NoError(t, err)
	})

	t.Run("EnsureDisksHaveControllers", func(t *testing.T) {

		const vmx20 = "vmx-20"

		testCases := []struct {
			name               string
			configSpec         *VirtualMachineConfigSpec
			existingDevices    []BaseVirtualDevice
			err                error
			expectedConfigSpec *VirtualMachineConfigSpec
			panicMsg           string
		}{
			{
				name:               "nil configSpec arg should panic",
				configSpec:         nil,
				existingDevices:    nil,
				err:                nil,
				expectedConfigSpec: nil,
				panicMsg:           "configSpec is nil",
			},
			{
				name: "do nothing if configSpec has no disks",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
				},
				existingDevices: nil,
				err:             nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
				},
			},
			{
				name: "do nothing if configSpec has a disk, but it is being removed",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: nil,
				err:             nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device:    &VirtualDisk{},
						},
					},
				},
			},
			{
				name: "do nothing if configSpec has a disk, but it is already attached to a controller from existing devices",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&ParaVirtualSCSIController{
						VirtualSCSIController: VirtualSCSIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           1000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk and there is a nil base device change in ConfigSpec and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						BaseVirtualDeviceConfigSpec(nil),
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						BaseVirtualDeviceConfigSpec(nil),
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk and there is a nil device change in ConfigSpec and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						nilVirtualDeviceConfigSpec{},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						nilVirtualDeviceConfigSpec{},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			{
				name: "attach disk to new PVSCSI controller when adding disk and there is a nil base device in ConfigSpec and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Device: BaseVirtualDevice(nil),
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Device: BaseVirtualDevice(nil),
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk and there is a nil device in ConfigSpec and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Device: nilVirtualDevice{},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Device: nilVirtualDevice{},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk and existing device includes PCI controller and nil base device",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					BaseVirtualDevice(nil),
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk and existing device includes PCI controller and nil device",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					nilVirtualDevice{},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk while removing SATA controller and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk that references SATA controller being removed and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 15000,
								},
							},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationRemove,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk that references non-existent controller and existing device includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1000,
								},
							},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk sans controller and no existing devices",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: nil,
				err:             nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -2,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualPCIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										Key: -1,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: -1,
											Key:           -2,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding PCI controller and disk sans controller and no existing devices",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualPCIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										Key: -50,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: nil,
				err:             nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualPCIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										Key: -50,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -51,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: -50,
											Key:           -51,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to new PVSCSI controller when adding disk sans controller and existing devices only includes PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			//
			// SCSI (in existing devices)
			//
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and PVSCSI controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&ParaVirtualSCSIController{
						VirtualSCSIController: VirtualSCSIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           1000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and Bus Logic controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualBusLogicController{
						VirtualSCSIController: VirtualSCSIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           1000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and LSI Logic controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualLsiLogicController{
						VirtualSCSIController: VirtualSCSIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           1000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and LSI Logic SAS controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualLsiLogicSASController{
						VirtualSCSIController: VirtualSCSIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           1000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and SCSI controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualSCSIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								ControllerKey: 100,
								Key:           1000,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 1000,
								},
							},
						},
					},
				},
			},

			//
			// SATA (in existing devices)
			//
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and SATA controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualAHCIController{
						VirtualSATAController: VirtualSATAController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           15000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 15000,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and AHCI controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualAHCIController{
						VirtualSATAController: VirtualSATAController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									ControllerKey: 100,
									Key:           15000,
								},
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 15000,
								},
							},
						},
					},
				},
			},

			//
			// NVME (in existing devices)
			//
			{
				name: "attach disk to existing controller when adding disk and existing devices includes PCI and NVME controllers",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
					&VirtualNVMEController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								ControllerKey: 100,
								Key:           31000,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: 31000,
								},
							},
						},
					},
				},
			},

			//
			// SCSI (in ConfigSpec)
			//
			{
				name: "attach disk to PVSCSI controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to Bus Logic controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to LSI Logic controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to LSI Logic SAS controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualLsiLogicSASController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualLsiLogicSASController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to SCSI controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualSCSIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           -10,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualSCSIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           -10,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},

			//
			// SATA (in ConfigSpec)
			//
			{
				name: "attach disk to SATA controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},
			{
				name: "attach disk to AHCI controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -10,
										},
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},

			//
			// NVME (in ConfigSpec)
			//
			{
				name: "attach disk to NVME controller in ConfigSpec when adding disk and existing devices includes only PCI controller",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           -10,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: []BaseVirtualDevice{
					&VirtualPCIController{
						VirtualController: VirtualController{
							VirtualDevice: VirtualDevice{
								Key: 100,
							},
						},
					},
				},
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           -10,
									},
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -10,
								},
							},
						},
					},
				},
			},

			//
			// First SCSI HBA is full
			//
			{
				name: "attach disk to new PVSCSI controller when adding disk and existing devices includes PCI controller and LSI Logic controller with no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
										BusNumber: 1,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			//
			// First and second SCSI HBAs are full
			//
			{
				name: "attach disk to new PVSCSI controller when adding disk and existing devices includes PCI controller and two existing SCSI controllers have no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
										BusNumber: 0,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),

					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
										BusNumber: 3,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 1,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			//
			// First, second, and third SCSI HBAs are full
			//
			{
				name: "attach disk to new PVSCSI controller when adding disk and existing devices includes PCI controller and three existing SCSI controllers have no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
										BusNumber: 3,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),

					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
										BusNumber: 0,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),

					//
					// SCSI HBA 3 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1002,
										},
										BusNumber: 1,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1002)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 2,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			//
			// First, second, and third SCSI HBAs are full (different bus number)
			//
			{
				name: "attach disk to new PVSCSI controller (bus number three) when adding disk and existing devices includes PCI controller and three existing SCSI controllers have no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
										BusNumber: 2,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),

					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
										BusNumber: 0,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),

					//
					// SCSI HBA 3 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualBusLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1002,
										},
										BusNumber: 1,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1002)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 3,
									},
									HotAddRemove: NewBool(true),
									SharedBus:    VirtualSCSISharingNoSharing,
								},
							},
						},
					},
				},
			},

			//
			// All SCSI HBAs are full
			//
			{
				name: "attach disk to new SATA controller when adding disk and existing devices includes PCI controller and there are already four SCSI controllers with no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),
					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),
					//
					// SCSI HBA 3 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualSCSIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           1002,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1002)...,
					),
					//
					// SCSI HBA 4 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicSASController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1003,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1003)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           -1,
										},
										BusNumber: 0,
									},
								},
							},
						},
					},
				},
			},

			{
				name: "attach disk to new NVME controller when adding disk and existing devices includes PCI controller and there are already four SCSI and SATA controllers with no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),
					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),
					//
					// SCSI HBA 3 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualSCSIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           1002,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1002)...,
					),
					//
					// SCSI HBA 4 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicSASController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1003,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1003)...,
					),

					//
					// SATA HBA 1 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15000)...,
					),
					//
					// SATA HBA 2 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15001,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15001)...,
					),
					//
					// SATA HBA 3 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15002,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15002)...,
					),
					//
					// SATA HBA 4 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15003,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15003)...,
					),
				),
				err: nil,
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualDisk{
								VirtualDevice: VirtualDevice{
									ControllerKey: -1,
								},
							},
						},
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device: &VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           -1,
									},
									BusNumber: 0,
								},
								SharedBus: string(VirtualNVMEControllerSharingNoSharing),
							},
						},
					},
				},
			},
			{
				name: "return an error that there are no available controllers when adding a disk and existing devices includes PCI controller and there are already four SCSI, SATA, and NVME controllers with no free slots",
				configSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
				existingDevices: ensureDiskControllerJoinSlices(
					[]BaseVirtualDevice{
						&VirtualPCIController{
							VirtualController: VirtualController{
								VirtualDevice: VirtualDevice{
									Key: 100,
								},
							},
						},
					},

					//
					// SCSI HBA 1 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&ParaVirtualSCSIController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1000,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1000)...,
					),
					//
					// SCSI HBA 2 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1001,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1001)...,
					),
					//
					// SCSI HBA 3 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualSCSIController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           1002,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1002)...,
					),
					//
					// SCSI HBA 4 / 16 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualLsiLogicSASController{
								VirtualSCSIController: VirtualSCSIController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           1003,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(16, 1003)...,
					),

					//
					// SATA HBA 1 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15000,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15000)...,
					),
					//
					// SATA HBA 2 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15001,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15001)...,
					),
					//
					// SATA HBA 3 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualSATAController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           15002,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15002)...,
					),
					//
					// SATA HBA 4 / 30 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualAHCIController{
								VirtualSATAController: VirtualSATAController{
									VirtualController: VirtualController{
										VirtualDevice: VirtualDevice{
											ControllerKey: 100,
											Key:           15003,
										},
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(30, 15003)...,
					),

					//
					// NVME HBA 1 / 15 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           31000,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(15, 31000)...,
					),
					//
					// NVME HBA 2 / 15 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           31001,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(15, 31001)...,
					),
					//
					// NVME HBA 3 / 15 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           31002,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(15, 31002)...,
					),
					//
					// NVME HBA 4 / 15 disks
					//
					append(
						[]BaseVirtualDevice{
							&VirtualNVMEController{
								VirtualController: VirtualController{
									VirtualDevice: VirtualDevice{
										ControllerKey: 100,
										Key:           31003,
									},
								},
							},
						},
						ensureDiskControllerGenerateVirtualDisks(15, 31003)...,
					),
				),
				err: fmt.Errorf("no controllers available"),
				expectedConfigSpec: &VirtualMachineConfigSpec{
					Version: vmx20,
					DeviceChange: []BaseVirtualDeviceConfigSpec{
						&VirtualDeviceConfigSpec{
							Operation: VirtualDeviceConfigSpecOperationAdd,
							Device:    &VirtualDisk{},
						},
					},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				if tc.panicMsg != "" {
					assert.PanicsWithValue(t, tc.panicMsg, func() { _ = tc.configSpec.EnsureDisksHaveControllers(tc.existingDevices...) })
				} else {
					var err error
					assert.NotPanics(t, func() { err = tc.configSpec.EnsureDisksHaveControllers(tc.existingDevices...) })
					assert.Equal(t, tc.err, err)
					assert.Equal(t, tc.expectedConfigSpec, tc.configSpec)
				}
			})
		}
	})
}

type nilVirtualDeviceConfigSpec struct{}

func (n nilVirtualDeviceConfigSpec) GetVirtualDeviceConfigSpec() *VirtualDeviceConfigSpec {
	return nil
}

type nilVirtualDevice struct{}

func (n nilVirtualDevice) GetVirtualDevice() *VirtualDevice {
	return nil
}

func ensureDiskControllerJoinSlices[T any](a []T, b ...[]T) []T {
	for i := range b {
		a = append(a, b[i]...)
	}
	return a
}

func ensureDiskControllerGenerateVirtualDisks(
	numDisks int, controllerKey int32) []BaseVirtualDevice {

	devices := make([]BaseVirtualDevice, numDisks)
	for i := range devices {
		devices[i] = &VirtualDisk{
			VirtualDevice: VirtualDevice{
				ControllerKey: controllerKey,
			},
		}
	}
	return devices
}
