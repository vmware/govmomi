/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package types

import (
	"testing"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestVirtualMachineConfigInfoToConfigSpec(t *testing.T) {
	testCases := []struct {
		name string
		conf VirtualMachineConfigInfo
		spec VirtualMachineConfigSpec
		fail bool
	}{
		{
			name: "default value",
			conf: VirtualMachineConfigInfo{},
			spec: VirtualMachineConfigSpec{},
		},
		{
			name: "matching names",
			conf: VirtualMachineConfigInfo{
				Name: "Hello, world.",
			},
			spec: VirtualMachineConfigSpec{
				Name: "Hello, world.",
			},
		},
		{
			name: "matching nics",
			conf: VirtualMachineConfigInfo{
				Name: "Hello, world.",
				Hardware: VirtualHardware{
					Device: []BaseVirtualDevice{
						&VirtualVmxnet3{
							VirtualVmxnet: VirtualVmxnet{
								VirtualEthernetCard: VirtualEthernetCard{
									VirtualDevice: VirtualDevice{
										Key: 3,
									},
									MacAddress: "00:11:22:33:44:55:66:77",
								},
							},
						},
					},
				},
			},
			spec: VirtualMachineConfigSpec{
				Name: "Hello, world.",
				DeviceChange: []BaseVirtualDeviceConfigSpec{
					&VirtualDeviceConfigSpec{
						Operation:     VirtualDeviceConfigSpecOperationAdd,
						FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
						Device: &VirtualVmxnet3{
							VirtualVmxnet: VirtualVmxnet{
								VirtualEthernetCard: VirtualEthernetCard{
									VirtualDevice: VirtualDevice{
										Key: 3,
									},
									MacAddress: "00:11:22:33:44:55:66:77",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "nics with different mac addresses",
			fail: true,
			conf: VirtualMachineConfigInfo{
				Name: "Hello, world.",
				Hardware: VirtualHardware{
					Device: []BaseVirtualDevice{
						&VirtualVmxnet3{
							VirtualVmxnet: VirtualVmxnet{
								VirtualEthernetCard: VirtualEthernetCard{
									VirtualDevice: VirtualDevice{
										Key: 3,
									},
									MacAddress: "00:11:22:33:44:55:66:77",
								},
							},
						},
					},
				},
			},
			spec: VirtualMachineConfigSpec{
				Name: "Hello, world.",
				DeviceChange: []BaseVirtualDeviceConfigSpec{
					&VirtualDeviceConfigSpec{
						Operation:     VirtualDeviceConfigSpecOperationAdd,
						FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
						Device: &VirtualVmxnet3{
							VirtualVmxnet: VirtualVmxnet{
								VirtualEthernetCard: VirtualEthernetCard{
									VirtualDevice: VirtualDevice{
										Key: 3,
									},
									MacAddress: "00:11:22:33:44:55:66:88",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "vAppConfig",
			conf: VirtualMachineConfigInfo{
				Name: "Hello, world.",
				VAppConfig: &VmConfigInfo{
					InstallBootRequired: false,
					IpAssignment:        VAppIPAssignmentInfo{},
					Product: []VAppProductInfo{
						{
							Key:  1,
							Name: "P1",
						},
					},
				},
			},
			spec: VirtualMachineConfigSpec{
				Name: "Hello, world.",
				VAppConfig: &VmConfigSpec{
					InstallBootRequired: NewBool(false),
					IpAssignment:        &VAppIPAssignmentInfo{},
					Product: []VAppProductSpec{
						{
							ArrayUpdateSpec: ArrayUpdateSpec{
								Operation: ArrayUpdateOperationAdd,
							},
							Info: &VAppProductInfo{
								Key:  1,
								Name: "P1",
							},
						},
					},
				},
			},
		},
		{
			name: "really big config",
			conf: VirtualMachineConfigInfo{
				Name:    "vm-001",
				GuestId: "otherGuest",
				Files:   VirtualMachineFileInfo{VmPathName: "[datastore1]"},
				Hardware: VirtualHardware{
					NumCPU:   1,
					MemoryMB: 128,
					Device: []BaseVirtualDevice{
						&VirtualLsiLogicController{
							VirtualSCSIController: VirtualSCSIController{
								SharedBus: VirtualSCSISharingNoSharing,
								VirtualController: VirtualController{
									BusNumber: 0,
									VirtualDevice: VirtualDevice{
										Key: 1000,
									},
								},
							},
						},
						&VirtualDisk{
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
						&VirtualE1000{
							VirtualEthernetCard: VirtualEthernetCard{
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
							},
						},
					},
				},
				ExtraConfig: []BaseOptionValue{
					&OptionValue{Key: "bios.bootOrder", Value: "ethernet0"},
				},
			},
			spec: VirtualMachineConfigSpec{
				Name:     "vm-001",
				GuestId:  "otherGuest",
				Files:    &VirtualMachineFileInfo{VmPathName: "[datastore1]"},
				NumCPUs:  1,
				MemoryMB: 128,
				DeviceChange: []BaseVirtualDeviceConfigSpec{
					&VirtualDeviceConfigSpec{
						Operation:     VirtualDeviceConfigSpecOperationAdd,
						FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
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
						Operation:     VirtualDeviceConfigSpecOperationAdd,
						FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
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
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			e, a := tc.spec, tc.conf.ToConfigSpec()
			ed, err := xml.MarshalIndent(e, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal expected ConfigSpec: %v", err)
			}
			ad, err := xml.MarshalIndent(a, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal actual   ConfigSpec: %v", err)
			}
			eds, ads := string(ed), string(ad)
			if eds != ads && !tc.fail {
				t.Errorf("unexpected error: \n\n"+
					"exp=%+v\n\nact=%+v\n\n"+
					"exp.s=%s\n\nact.s=%s\n\n", e, a, eds, ads)
			} else if eds == ads && tc.fail {
				t.Errorf("expected error did not occur: \n\n"+
					"exp=%+v\n\nact=%+v\n\n"+
					"exp.s=%s\n\nact.s=%s\n\n", e, a, eds, ads)
			}
		})
	}
}
