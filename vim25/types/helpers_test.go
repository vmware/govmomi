// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

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
					Property: []VAppPropertyInfo{
						{
							Key:              int32(0),
							ClassId:          "Class-1",
							InstanceId:       "InstanceId-1",
							Id:               "property0-key",
							Category:         "User-Configurable",
							Label:            "userConfigurable",
							Type:             "string",
							UserConfigurable: NewBool(true),
							DefaultValue:     "",
							Value:            "",
							Description:      "A user-configurable property",
						},
						{
							Key:              int32(1),
							ClassId:          "Class-1",
							InstanceId:       "InstanceId-1",
							Id:               "property1-key",
							Category:         "Non-User-Configurable",
							Label:            "non-userConfigurable",
							Type:             "string",
							UserConfigurable: NewBool(false),
							DefaultValue:     "foo",
							Value:            "",
							Description:      "A non user-configurable property",
						},
						{
							Key:              int32(2),
							ClassId:          "Class-1",
							InstanceId:       "InstanceId-1",
							Id:               "property2-key",
							Category:         "Non-User-Configurable",
							Label:            "nil-userConfigurable",
							Type:             "string",
							UserConfigurable: nil,
							DefaultValue:     "bar",
							Value:            "",
							Description:      "A nil user-configurable property",
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
					Property: []VAppPropertySpec{
						{
							ArrayUpdateSpec: ArrayUpdateSpec{
								Operation: ArrayUpdateOperationAdd,
							},
							Info: &VAppPropertyInfo{
								Key:              int32(0),
								ClassId:          "Class-1",
								InstanceId:       "InstanceId-1",
								Id:               "property0-key",
								Category:         "User-Configurable",
								Label:            "userConfigurable",
								Type:             "string",
								UserConfigurable: NewBool(true),
								DefaultValue:     "",
								Value:            "",
								Description:      "A user-configurable property",
							},
						},
						{
							ArrayUpdateSpec: ArrayUpdateSpec{
								Operation: ArrayUpdateOperationAdd,
							},
							Info: &VAppPropertyInfo{
								Key:              int32(1),
								ClassId:          "Class-1",
								InstanceId:       "InstanceId-1",
								Id:               "property1-key",
								Category:         "Non-User-Configurable",
								Label:            "non-userConfigurable",
								Type:             "string",
								UserConfigurable: NewBool(false),
								DefaultValue:     "foo",
								Value:            "",
								Description:      "A non user-configurable property",
							},
						},
						{
							ArrayUpdateSpec: ArrayUpdateSpec{
								Operation: ArrayUpdateOperationAdd,
							},
							Info: &VAppPropertyInfo{
								Key:              int32(2),
								ClassId:          "Class-1",
								InstanceId:       "InstanceId-1",
								Id:               "property2-key",
								Category:         "Non-User-Configurable",
								Label:            "nil-userConfigurable",
								Type:             "string",
								UserConfigurable: NewBool(false),
								DefaultValue:     "bar",
								Value:            "",
								Description:      "A nil user-configurable property",
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

type toStringTestCase struct {
	name     string
	in       any
	expected string
}

func newToStringTestCases[T any](in T, expected string) []toStringTestCase {
	return newToStringTestCasesWithTestCaseName(
		in, expected, reflect.TypeOf(in).Name())
}

func newToStringTestCasesWithTestCaseName[T any](
	in T, expected, testCaseName string) []toStringTestCase {

	return []toStringTestCase{
		{
			name:     testCaseName,
			in:       in,
			expected: expected,
		},
		{
			name:     "*" + testCaseName,
			in:       &[]T{in}[0],
			expected: expected,
		},
		{
			name:     "(any)(" + testCaseName + ")",
			in:       (any)(in),
			expected: expected,
		},
		{
			name:     "(any)(*" + testCaseName + ")",
			in:       (any)(&[]T{in}[0]),
			expected: expected,
		},
		{
			name:     "(any)((*" + testCaseName + ")(nil))",
			in:       (any)((*T)(nil)),
			expected: "null",
		},
	}
}

type toStringTypeWithErr struct {
	errOnCall []int
	callCount *int
	doPanic   bool
}

func (t toStringTypeWithErr) String() string {
	return "{}"
}

func (t toStringTypeWithErr) MarshalJSON() ([]byte, error) {
	defer func() {
		*t.callCount++
	}()
	if !slices.Contains(t.errOnCall, *t.callCount) {
		return []byte{'{', '}'}, nil
	}
	if t.doPanic {
		panic(fmt.Errorf("marshal json panic'd"))
	}
	return nil, fmt.Errorf("marshal json failed")
}

func TestToString(t *testing.T) {
	const (
		helloWorld = "Hello, world."
	)

	testCases := []toStringTestCase{
		{
			name:     "nil",
			in:       nil,
			expected: "null",
		},
	}

	testCases = append(testCases, newToStringTestCases(
		"Hello, world.", "Hello, world.")...)

	testCases = append(testCases, newToStringTestCasesWithTestCaseName(
		byte(1), "1", "byte")...)
	testCases = append(testCases, newToStringTestCasesWithTestCaseName(
		'a', "97", "rune")...)

	testCases = append(testCases, newToStringTestCases(
		true, "true")...)

	testCases = append(testCases, newToStringTestCases(
		complex(float32(1), float32(4)), "(1+4i)")...)
	testCases = append(testCases, newToStringTestCases(
		complex(float64(1), float64(4)), "(1+4i)")...)

	testCases = append(testCases, newToStringTestCases(
		float32(1.1), "1.1")...)
	testCases = append(testCases, newToStringTestCases(
		float64(1.1), "1.1")...)

	testCases = append(testCases, newToStringTestCases(
		int(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		int8(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		int16(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		int32(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		int64(1), "1")...)

	testCases = append(testCases, newToStringTestCases(
		uint(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		uint8(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		uint16(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		uint32(1), "1")...)
	testCases = append(testCases, newToStringTestCases(
		uint64(1), "1")...)

	testCases = append(testCases, newToStringTestCases(
		VirtualMachineConfigSpec{},
		`{"_typeName":"VirtualMachineConfigSpec"}`)...)
	testCases = append(testCases, newToStringTestCasesWithTestCaseName(
		VirtualMachineConfigSpec{
			VAppConfig: (*VmConfigSpec)(nil),
		},
		`{"_typeName":"VirtualMachineConfigSpec","vAppConfig":null}`,
		"VirtualMachineConfigSpec w nil iface")...)

	testCases = append(testCases, toStringTestCase{
		name:     "MarshalJSON returns error on special encode",
		in:       toStringTypeWithErr{callCount: new(int), errOnCall: []int{0}},
		expected: "{}",
	})
	testCases = append(testCases, toStringTestCase{
		name:     "MarshalJSON returns error on special and stdlib encode",
		in:       toStringTypeWithErr{callCount: new(int), errOnCall: []int{0, 1}},
		expected: "{}",
	})
	testCases = append(testCases, toStringTestCase{
		name:     "MarshalJSON panics on special encode",
		in:       toStringTypeWithErr{callCount: new(int), doPanic: true, errOnCall: []int{0}},
		expected: "{}",
	})
	testCases = append(testCases, toStringTestCase{
		name:     "MarshalJSON panics on special and stdlib encode",
		in:       toStringTypeWithErr{callCount: new(int), doPanic: true, errOnCall: []int{0, 1}},
		expected: "{}",
	})

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, ToString(tc.in))
		})
	}
}
