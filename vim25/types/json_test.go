// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var serializationTests = []struct {
	name      string
	file      string
	data      any
	goType    reflect.Type
	expDecErr string
}{
	{
		name:   "vminfo",
		file:   "./testdata/vminfo.json",
		data:   &vmInfoObjForTests,
		goType: reflect.TypeOf(VirtualMachineConfigInfo{}),
	},
	{
		name:   "retrieveResult",
		file:   "./testdata/retrieveResult.json",
		data:   &retrieveResultForTests,
		goType: reflect.TypeOf(RetrieveResult{}),
	},
	{
		name:      "vminfo-invalid-type-name-value",
		file:      "./testdata/vminfo-invalid-type-name-value.json",
		data:      &vmInfoObjForTests,
		goType:    reflect.TypeOf(VirtualMachineConfigInfo{}),
		expDecErr: `json: cannot unmarshal bool into Go struct field VirtualMachineConfigInfo.extraConfig of type string`,
	},
}

func TestSerialization(t *testing.T) {
	for _, test := range serializationTests {
		t.Run(test.name+" Decode", func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			dec := NewJSONDecoder(f)

			ee := test.expDecErr
			data := reflect.New(test.goType).Interface()
			if err := dec.Decode(data); err != nil {
				if ee != err.Error() {
					t.Errorf("expected error mismatch: e=%v, a=%v", ee, err)
				} else if ee == "" {
					t.Errorf("unexpected error: %v", err)
				}
			} else if ee != "" {
				t.Errorf("expected error did not occur: %v", ee)
			} else {
				a, e := data, test.data
				if diff := cmp.Diff(a, e); diff != "" {
					t.Errorf("mismatched %v: %s", test.name, diff)
				}
			}
		})

		t.Run(test.name+" Encode", func(t *testing.T) {
			if test.expDecErr != "" {
				t.Skip("skipping due to expected decode error")
			}

			expJSON, err := os.ReadFile(test.file)
			if err != nil {
				t.Fatal(err)
			}

			var w bytes.Buffer
			_ = w
			enc := NewJSONEncoder(&w)

			if err := enc.Encode(test.data); err != nil {
				t.Fatal(err)
			}

			expected, actual := string(expJSON), w.String()
			assert.JSONEq(t, expected, actual)
		})
	}

	t.Run("ConfigSpec", func(t *testing.T) {
		t.Run("Encode", func(t *testing.T) {

			var testCases = []struct {
				name        string
				data        any
				expected    string
				expectPanic bool
			}{
				{
					name:     "nil ConfigSpec",
					data:     (*VirtualMachineConfigSpec)(nil),
					expected: "null",
				},
				{
					name: "ConfigSpec with nil OptionValue value",
					data: &VirtualMachineConfigSpec{
						ExtraConfig: []BaseOptionValue{
							&OptionValue{
								Key:   "key1",
								Value: nil,
							},
						},
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec","extraConfig":[{"_typeName":"OptionValue","key":"key1","value":null}]}`,
				},
				{
					name: "ConfigSpec with nil OptionValue interface value",
					data: &VirtualMachineConfigSpec{
						ExtraConfig: []BaseOptionValue{
							&OptionValue{
								Key:   "key1",
								Value: (any)(nil),
							},
						},
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec","extraConfig":[{"_typeName":"OptionValue","key":"key1","value":null}]}`,
				},
				{
					name: "ConfigSpec with nil element in OptionValues",
					data: &VirtualMachineConfigSpec{
						ExtraConfig: []BaseOptionValue{
							(*OptionValue)(nil),
						},
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec","extraConfig":[null]}`,
				},
				{
					name: "ConfigSpec with nil ToolsConfigInfo",
					data: &VirtualMachineConfigSpec{
						Tools: (*ToolsConfigInfo)(nil),
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec"}`,
				},
				{
					name: "ConfigSpec with nil vAppConfig",
					data: &VirtualMachineConfigSpec{
						VAppConfig: nil,
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec"}`,
				},
				{
					name: "ConfigSpec with nil pointer vAppConfig ",
					data: &VirtualMachineConfigSpec{
						VAppConfig: (*VmConfigSpec)(nil),
					},
					expected: `{"_typeName":"VirtualMachineConfigSpec","vAppConfig":null}`,
				},
			}

			for i := range testCases {
				tc := testCases[i]
				t.Run(tc.name, func(t *testing.T) {
					var w bytes.Buffer
					enc := NewJSONEncoder(&w)

					var panicErr any

					func() {
						defer func() {
							panicErr = recover()
						}()
						if err := enc.Encode(tc.data); err != nil {
							t.Fatal(err)
						}
					}()

					if tc.expectPanic && panicErr == nil {
						t.Fatalf("did not panic, w=%v", w.String())
					} else if tc.expectPanic && panicErr != nil {
						t.Logf("expected panic occurred: %v\n", panicErr)
					} else if !tc.expectPanic && panicErr != nil {
						t.Fatalf("unexpected panic occurred: %v\n", panicErr)
					} else if a, e := w.String(), tc.expected+"\n"; a != e {
						t.Fatalf("act=%s != exp=%s", a, e)
					} else {
						t.Log(a)
					}
				})
			}
		})
	})
}

func TestOptionValueSerialization(t *testing.T) {
	tv, e := time.Parse(time.RFC3339Nano, "2022-12-12T11:48:35.473645Z")
	if e != nil {
		t.Log("Cannot parse test timestamp. This is coding error.")
		t.Fail()
		return
	}
	options := []struct {
		name    string
		wire    string
		binding OptionValue
	}{
		{
			name: "boolean",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "boolean","_value": true}
			}`,
			binding: OptionValue{Key: "option1", Value: true},
		},
		{
			name: "byte",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "byte","_value": 16}
			}`,
			binding: OptionValue{Key: "option1", Value: uint8(16)},
		},
		{
			name: "short",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "short","_value": 300}
			}`,
			binding: OptionValue{Key: "option1", Value: int16(300)},
		},
		{
			name: "int",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "int","_value": 300}}`,
			binding: OptionValue{Key: "option1", Value: int32(300)},
		},
		{
			name: "long",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "long","_value": 300}}`,
			binding: OptionValue{Key: "option1", Value: int64(300)},
		},
		{
			name: "float",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "float","_value": 30.5}}`,
			binding: OptionValue{Key: "option1", Value: float32(30.5)},
		},
		{
			name: "double",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "double","_value": 12.2}}`,
			binding: OptionValue{Key: "option1", Value: float64(12.2)},
		},
		{
			name: "string",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "string","_value": "test"}}`,
			binding: OptionValue{Key: "option1", Value: "test"},
		},
		{
			name: "dateTime", // time.Time
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "dateTime","_value": "2022-12-12T11:48:35.473645Z"}}`,
			binding: OptionValue{Key: "option1", Value: tv},
		},
		{
			name: "binary", // []byte
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "binary","_value": "SGVsbG8="}}`,
			binding: OptionValue{Key: "option1", Value: []byte("Hello")},
		},
		// during serialization we have no way to guess that a string is to be
		// converted to uri. Using net.URL solves this. It is a breaking change.
		// See https://github.com/vmware/govmomi/pull/3123
		// {
		// 	name: "anyURI", // string
		// 	wire: `{"_typeName": "OptionValue","key": "option1",
		// 		"value": {"_typeName": "anyURI","_value": "http://hello"}}`,
		// 	binding: OptionValue{Key: "option1", Value: "test"},
		// },
		{
			name: "enum",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "CustomizationNetBIOSMode","_value": "enableNetBIOS"}}`,
			binding: OptionValue{Key: "option1", Value: CustomizationNetBIOSModeEnableNetBIOS},
		},
		// There is no ArrayOfCustomizationNetBIOSMode type emitted i.e. no enum
		// array types are emitted in govmomi.
		// We can process these in the serialization logic i.e. discover or
		// prepend the "ArrayOf" prefix
		// {
		// 	name: "array of enum",
		// 	wire: `{
		//		"_typeName": "OptionValue",
		//		"key": "option1",
		// 		"value": {"_typeName": "ArrayOfCustomizationNetBIOSMode",
		//                "_value": ["enableNetBIOS"]}}`,
		// 	binding: OptionValue{Key: "option1",
		//		Value: []CustomizationNetBIOSMode{
		//			CustomizationNetBIOSModeEnableNetBIOS
		//		}},
		// },

		// array of struct is weird. Do we want to unmarshal this as
		// []ClusterHostRecommendation directly? Why do we want to use
		// ArrayOfClusterHostRecommendation wrapper?
		// if SOAP does it then I guess back compat is a big reason.
		{
			name: "array of struct",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "ArrayOfClusterHostRecommendation","_value": [
					{
						"_typeName":"ClusterHostRecommendation",
						"host": {
							"_typeName": "ManagedObjectReference",
							"type": "HostSystem",
							"value": "host-42"
						},
						"rating":42
					}]}}`,
			binding: OptionValue{
				Key: "option1",
				Value: ArrayOfClusterHostRecommendation{
					ClusterHostRecommendation: []ClusterHostRecommendation{
						{
							Host: ManagedObjectReference{
								Type:  "HostSystem",
								Value: "host-42",
							},
							Rating: 42,
						},
					},
				},
			},
		},
	}

	for _, opt := range options {
		t.Run("Serialize "+opt.name, func(t *testing.T) {
			r := strings.NewReader(opt.wire)
			dec := NewJSONDecoder(r)
			v := OptionValue{}
			e := dec.Decode(&v)
			if e != nil {
				assert.Fail(t, "Cannot read json", "json %v err %v", opt.wire, e)
				return
			}
			assert.Equal(t, opt.binding, v)
		})

		t.Run("De-serialize "+opt.name, func(t *testing.T) {
			var w bytes.Buffer
			enc := NewJSONEncoder(&w)
			enc.Encode(opt.binding)
			s := w.String()
			assert.JSONEq(t, opt.wire, s)
		})
	}
}

var vmInfoObjForTests = VirtualMachineConfigInfo{
	ChangeVersion:         "2022-12-12T11:48:35.473645Z",
	Modified:              mustParseTime(time.RFC3339, "1970-01-01T00:00:00Z"),
	Name:                  "test",
	GuestFullName:         "VMware Photon OS (64-bit)",
	Version:               "vmx-20",
	Uuid:                  "422ca90b-853b-1101-3350-759f747730cc",
	CreateDate:            addrOfMustParseTime(time.RFC3339, "2022-12-12T11:47:24.685785Z"),
	InstanceUuid:          "502cc2a5-1f06-2890-6d70-ba2c55c5c2b7",
	NpivTemporaryDisabled: NewBool(true),
	LocationId:            "Earth",
	Template:              false,
	GuestId:               "vmwarePhoton64Guest",
	AlternateGuestName:    "",
	Annotation:            "Hello, world.",
	Files: VirtualMachineFileInfo{
		VmPathName:        "[datastore1] test/test.vmx",
		SnapshotDirectory: "[datastore1] test/",
		SuspendDirectory:  "[datastore1] test/",
		LogDirectory:      "[datastore1] test/",
	},
	Tools: &ToolsConfigInfo{
		ToolsVersion:            1,
		AfterPowerOn:            NewBool(true),
		AfterResume:             NewBool(true),
		BeforeGuestStandby:      NewBool(true),
		BeforeGuestShutdown:     NewBool(true),
		BeforeGuestReboot:       nil,
		ToolsUpgradePolicy:      "manual",
		SyncTimeWithHostAllowed: NewBool(true),
		SyncTimeWithHost:        NewBool(false),
		LastInstallInfo: &ToolsConfigInfoToolsLastInstallInfo{
			Counter: 0,
		},
	},
	Flags: VirtualMachineFlagInfo{
		EnableLogging:            NewBool(true),
		UseToe:                   NewBool(false),
		RunWithDebugInfo:         NewBool(false),
		MonitorType:              "release",
		HtSharing:                "any",
		SnapshotDisabled:         NewBool(false),
		SnapshotLocked:           NewBool(false),
		DiskUuidEnabled:          NewBool(false),
		SnapshotPowerOffBehavior: "powerOff",
		RecordReplayEnabled:      NewBool(false),
		FaultToleranceType:       "unset",
		CbrcCacheEnabled:         NewBool(false),
		VvtdEnabled:              NewBool(false),
		VbsEnabled:               NewBool(false),
	},
	DefaultPowerOps: VirtualMachineDefaultPowerOpInfo{
		PowerOffType:        "soft",
		SuspendType:         "hard",
		ResetType:           "soft",
		DefaultPowerOffType: "soft",
		DefaultSuspendType:  "hard",
		DefaultResetType:    "soft",
		StandbyAction:       "checkpoint",
	},
	RebootPowerOff: NewBool(false),
	Hardware: VirtualHardware{
		NumCPU:              1,
		NumCoresPerSocket:   NewInt32(1),
		AutoCoresPerSocket:  NewBool(true),
		MemoryMB:            2048,
		VirtualICH7MPresent: NewBool(false),
		VirtualSMCPresent:   NewBool(false),
		Device: []BaseVirtualDevice{
			&VirtualIDEController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 200,
						DeviceInfo: &Description{
							Label:   "IDE 0",
							Summary: "IDE 0",
						},
					},
					BusNumber: 0,
				},
			},
			&VirtualIDEController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 201,
						DeviceInfo: &Description{
							Label:   "IDE 1",
							Summary: "IDE 1",
						},
					},
					BusNumber: 1,
				},
			},
			&VirtualPS2Controller{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 300,
						DeviceInfo: &Description{
							Label:   "PS2 controller 0",
							Summary: "PS2 controller 0",
						},
					},
					BusNumber: 0,
					Device:    []int32{600, 700},
				},
			},
			&VirtualPCIController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 100,
						DeviceInfo: &Description{
							Label:   "PCI controller 0",
							Summary: "PCI controller 0",
						},
					},
					BusNumber: 0,
					Device:    []int32{500, 12000, 14000, 1000, 15000, 4000},
				},
			},
			&VirtualSIOController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 400,
						DeviceInfo: &Description{
							Label:   "SIO controller 0",
							Summary: "SIO controller 0",
						},
					},
					BusNumber: 0,
				},
			},
			&VirtualKeyboard{
				VirtualDevice: VirtualDevice{
					Key: 600,
					DeviceInfo: &Description{
						Label:   "Keyboard",
						Summary: "Keyboard",
					},
					ControllerKey: 300,
					UnitNumber:    NewInt32(0),
				},
			},
			&VirtualPointingDevice{
				VirtualDevice: VirtualDevice{
					Key:        700,
					DeviceInfo: &Description{Label: "Pointing device", Summary: "Pointing device; Device"},
					Backing: &VirtualPointingDeviceDeviceBackingInfo{
						VirtualDeviceDeviceBackingInfo: VirtualDeviceDeviceBackingInfo{
							UseAutoDetect: NewBool(false),
						},
						HostPointingDevice: "autodetect",
					},
					ControllerKey: 300,
					UnitNumber:    NewInt32(1),
				},
			},
			&VirtualMachineVideoCard{
				VirtualDevice: VirtualDevice{
					Key:           500,
					DeviceInfo:    &Description{Label: "Video card ", Summary: "Video card"},
					ControllerKey: 100,
					UnitNumber:    NewInt32(0),
				},
				VideoRamSizeInKB:       4096,
				NumDisplays:            1,
				UseAutoDetect:          NewBool(false),
				Enable3DSupport:        NewBool(false),
				Use3dRenderer:          "automatic",
				GraphicsMemorySizeInKB: 262144,
			},
			&VirtualMachineVMCIDevice{
				VirtualDevice: VirtualDevice{
					Key: 12000,
					DeviceInfo: &Description{
						Label: "VMCI device",
						Summary: "Device on the virtual machine PCI " +
							"bus that provides support for the " +
							"virtual machine communication interface",
					},
					ControllerKey: 100,
					UnitNumber:    NewInt32(17),
				},
				Id:                             -1,
				AllowUnrestrictedCommunication: NewBool(false),
				FilterEnable:                   NewBool(true),
			},
			&ParaVirtualSCSIController{
				VirtualSCSIController: VirtualSCSIController{
					VirtualController: VirtualController{
						VirtualDevice: VirtualDevice{
							Key: 1000,
							DeviceInfo: &Description{
								Label:   "SCSI controller 0",
								Summary: "VMware paravirtual SCSI",
							},
							ControllerKey: 100,
							UnitNumber:    NewInt32(3),
						},
						Device: []int32{2000},
					},
					HotAddRemove:       NewBool(true),
					SharedBus:          "noSharing",
					ScsiCtlrUnitNumber: 7,
				},
			},
			&VirtualAHCIController{
				VirtualSATAController: VirtualSATAController{
					VirtualController: VirtualController{
						VirtualDevice: VirtualDevice{
							Key: 15000,
							DeviceInfo: &Description{
								Label:   "SATA controller 0",
								Summary: "AHCI",
							},
							ControllerKey: 100,
							UnitNumber:    NewInt32(24),
						},
						Device: []int32{16000},
					},
				},
			},
			&VirtualCdrom{
				VirtualDevice: VirtualDevice{
					Key: 16000,
					DeviceInfo: &Description{
						Label:   "CD/DVD drive 1",
						Summary: "Remote device",
					},
					Backing: &VirtualCdromRemotePassthroughBackingInfo{
						VirtualDeviceRemoteDeviceBackingInfo: VirtualDeviceRemoteDeviceBackingInfo{
							UseAutoDetect: NewBool(false),
						},
					},
					Connectable:   &VirtualDeviceConnectInfo{AllowGuestControl: true, Status: "untried"},
					ControllerKey: 15000,
					UnitNumber:    NewInt32(0),
				},
			},
			&VirtualDisk{
				VirtualDevice: VirtualDevice{
					Key: 2000,
					DeviceInfo: &Description{
						Label:   "Hard disk 1",
						Summary: "4,194,304 KB",
					},
					Backing: &VirtualDiskFlatVer2BackingInfo{
						VirtualDeviceFileBackingInfo: VirtualDeviceFileBackingInfo{
							BackingObjectId: "1",
							FileName:        "[datastore1] test/test.vmdk",
							Datastore: &ManagedObjectReference{
								Type:  "Datastore",
								Value: "datastore-21",
							},
						},
						DiskMode:               "persistent",
						Split:                  NewBool(false),
						WriteThrough:           NewBool(false),
						ThinProvisioned:        NewBool(false),
						EagerlyScrub:           NewBool(false),
						Uuid:                   "6000C298-df15-fe89-ddcb-8ea33329595d",
						ContentId:              "e4e1a794c6307ce7906a3973fffffffe",
						ChangeId:               "",
						Parent:                 nil,
						DeltaDiskFormat:        "",
						DigestEnabled:          NewBool(false),
						DeltaGrainSize:         0,
						DeltaDiskFormatVariant: "",
						Sharing:                "sharingNone",
						KeyId:                  nil,
					},
					ControllerKey: 1000,
					UnitNumber:    NewInt32(0),
				},
				CapacityInKB:    4194304,
				CapacityInBytes: 4294967296,
				Shares:          &SharesInfo{Shares: 1000, Level: "normal"},
				StorageIOAllocation: &StorageIOAllocationInfo{
					Limit:       NewInt64(-1),
					Shares:      &SharesInfo{Shares: 1000, Level: "normal"},
					Reservation: NewInt32(0),
				},
				DiskObjectId:               "1-2000",
				NativeUnmanagedLinkedClone: NewBool(false),
			},
			&VirtualVmxnet3{
				VirtualVmxnet: VirtualVmxnet{
					VirtualEthernetCard: VirtualEthernetCard{
						VirtualDevice: VirtualDevice{
							Key: 4000,
							DeviceInfo: &Description{
								Label:   "Network adapter 1",
								Summary: "VM Network",
							},
							Backing: &VirtualEthernetCardNetworkBackingInfo{
								VirtualDeviceDeviceBackingInfo: VirtualDeviceDeviceBackingInfo{
									DeviceName:    "VM Network",
									UseAutoDetect: NewBool(false),
								},
								Network: &ManagedObjectReference{
									Type:  "Network",
									Value: "network-27",
								},
							},
							Connectable: &VirtualDeviceConnectInfo{
								MigrateConnect: "unset",
								StartConnected: true,
								Status:         "untried",
							},
							ControllerKey: 100,
							UnitNumber:    NewInt32(7),
						},
						AddressType:      "assigned",
						MacAddress:       "00:50:56:ac:4d:ed",
						WakeOnLanEnabled: NewBool(true),
						ResourceAllocation: &VirtualEthernetCardResourceAllocation{
							Reservation: NewInt64(0),
							Share: SharesInfo{
								Shares: 50,
								Level:  "normal",
							},
							Limit: NewInt64(-1),
						},
						UptCompatibilityEnabled: NewBool(true),
					},
				},
				Uptv2Enabled: NewBool(false),
			},
			&VirtualUSBXHCIController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 14000,
						DeviceInfo: &Description{
							Label:   "USB xHCI controller ",
							Summary: "USB xHCI controller",
						},
						SlotInfo: &VirtualDevicePciBusSlotInfo{
							PciSlotNumber: -1,
						},
						ControllerKey: 100,
						UnitNumber:    NewInt32(23),
					},
				},

				AutoConnectDevices: NewBool(false),
			},
		},
		MotherboardLayout:   "i440bxHostBridge",
		SimultaneousThreads: 1,
	},
	CpuAllocation: &ResourceAllocationInfo{
		Reservation:           NewInt64(0),
		ExpandableReservation: NewBool(false),
		Limit:                 NewInt64(-1),
		Shares: &SharesInfo{
			Shares: 1000,
			Level:  SharesLevelNormal,
		},
	},
	MemoryAllocation: &ResourceAllocationInfo{
		Reservation:           NewInt64(0),
		ExpandableReservation: NewBool(false),
		Limit:                 NewInt64(-1),
		Shares: &SharesInfo{
			Shares: 20480,
			Level:  SharesLevelNormal,
		},
	},
	LatencySensitivity: &LatencySensitivity{
		Level: LatencySensitivitySensitivityLevelNormal,
	},
	MemoryHotAddEnabled: NewBool(false),
	CpuHotAddEnabled:    NewBool(false),
	CpuHotRemoveEnabled: NewBool(false),
	ExtraConfig: []BaseOptionValue{
		&OptionValue{Key: "nvram", Value: "test.nvram"},
		&OptionValue{Key: "svga.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge0.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge4.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge4.virtualDev", Value: "pcieRootPort"},
		&OptionValue{Key: "pciBridge4.functions", Value: "8"},
		&OptionValue{Key: "pciBridge5.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge5.virtualDev", Value: "pcieRootPort"},
		&OptionValue{Key: "pciBridge5.functions", Value: "8"},
		&OptionValue{Key: "pciBridge6.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge6.virtualDev", Value: "pcieRootPort"},
		&OptionValue{Key: "pciBridge6.functions", Value: "8"},
		&OptionValue{Key: "pciBridge7.present", Value: "TRUE"},
		&OptionValue{Key: "pciBridge7.virtualDev", Value: "pcieRootPort"},
		&OptionValue{Key: "pciBridge7.functions", Value: "8"},
		&OptionValue{Key: "hpet0.present", Value: "TRUE"},
		&OptionValue{Key: "RemoteDisplay.maxConnections", Value: "-1"},
		&OptionValue{Key: "sched.cpu.latencySensitivity", Value: "normal"},
		&OptionValue{Key: "vmware.tools.internalversion", Value: "0"},
		&OptionValue{Key: "vmware.tools.requiredversion", Value: "12352"},
		&OptionValue{Key: "migrate.hostLogState", Value: "none"},
		&OptionValue{Key: "migrate.migrationId", Value: "0"},
		&OptionValue{Key: "migrate.hostLog", Value: "test-36f94569.hlog"},
		&OptionValue{
			Key:   "viv.moid",
			Value: "c5b34aa9-d962-4a74-b7d2-b83ec683ba1b:vm-28:lIgQ2t7v24n2nl3N7K3m6IHW2OoPF4CFrJd5N+Tdfio=",
		},
	},
	DatastoreUrl: []VirtualMachineConfigInfoDatastoreUrlPair{
		{
			Name: "datastore1",
			Url:  "/vmfs/volumes/63970ed8-4abddd2a-62d7-02003f49c37d",
		},
	},
	SwapPlacement: "inherit",
	BootOptions: &VirtualMachineBootOptions{
		EnterBIOSSetup:       NewBool(false),
		EfiSecureBootEnabled: NewBool(false),
		BootDelay:            1,
		BootRetryEnabled:     NewBool(false),
		BootRetryDelay:       10000,
		NetworkBootProtocol:  "ipv4",
	},
	FtInfo:                       nil,
	RepConfig:                    nil,
	VAppConfig:                   nil,
	VAssertsEnabled:              NewBool(false),
	ChangeTrackingEnabled:        NewBool(false),
	Firmware:                     "bios",
	MaxMksConnections:            -1,
	GuestAutoLockEnabled:         NewBool(true),
	ManagedBy:                    nil,
	MemoryReservationLockedToMax: NewBool(false),
	InitialOverhead: &VirtualMachineConfigInfoOverheadInfo{
		InitialMemoryReservation: 214446080,
		InitialSwapReservation:   2541883392,
	},
	NestedHVEnabled: NewBool(false),
	VPMCEnabled:     NewBool(false),
	ScheduledHardwareUpgradeInfo: &ScheduledHardwareUpgradeInfo{
		UpgradePolicy:                  "never",
		ScheduledHardwareUpgradeStatus: "none",
	},
	ForkConfigInfo:         nil,
	VFlashCacheReservation: 0,
	VmxConfigChecksum: []uint8{
		0x69, 0xf7, 0xa7, 0x9e,
		0xd1, 0xc2, 0x21, 0x4b,
		0x6c, 0x20, 0x77, 0x0a,
		0x94, 0x94, 0x99, 0xee,
		0x17, 0x5d, 0xdd, 0xa3,
	},
	MessageBusTunnelEnabled: NewBool(false),
	GuestIntegrityInfo: &VirtualMachineGuestIntegrityInfo{
		Enabled: NewBool(false),
	},
	MigrateEncryption: "opportunistic",
	SgxInfo: &VirtualMachineSgxInfo{
		FlcMode:            "unlocked",
		RequireAttestation: NewBool(false),
	},
	ContentLibItemInfo:      nil,
	FtEncryptionMode:        "ftEncryptionOpportunistic",
	GuestMonitoringModeInfo: &VirtualMachineGuestMonitoringModeInfo{},
	SevEnabled:              NewBool(false),
	NumaInfo: &VirtualMachineVirtualNumaInfo{
		AutoCoresPerNumaNode:    NewBool(true),
		VnumaOnCpuHotaddExposed: NewBool(false),
	},
	PmemFailoverEnabled:          NewBool(false),
	VmxStatsCollectionEnabled:    NewBool(true),
	VmOpNotificationToAppEnabled: NewBool(false),
	VmOpNotificationTimeout:      -1,
	DeviceSwap: &VirtualMachineVirtualDeviceSwap{
		LsiToPvscsi: &VirtualMachineVirtualDeviceSwapDeviceSwapInfo{
			Enabled:    NewBool(true),
			Applicable: NewBool(false),
			Status:     "none",
		},
	},
	Pmem:         nil,
	DeviceGroups: &VirtualMachineVirtualDeviceGroups{},
}

var retrieveResultForTests = RetrieveResult{
	Token: "",
	Objects: []ObjectContent{

		{

			DynamicData: DynamicData{},
			Obj: ManagedObjectReference{

				Type:  "Folder",
				Value: "group-d1",
			},
			PropSet: []DynamicProperty{
				{

					Name: "alarmActionsEnabled",
					Val:  true,
				},
				{

					Name: "availableField",
					Val: ArrayOfCustomFieldDef{

						CustomFieldDef: []CustomFieldDef{},
					},
				},

				{

					Name: "childEntity",
					Val: ArrayOfManagedObjectReference{
						ManagedObjectReference: []ManagedObjectReference{},
					},
				},
				{
					Name: "childType",
					Val: ArrayOfString{
						String: []string{
							"Folder",
							"Datacenter"},
					},
				},
				{
					Name: "configIssue",
					Val: ArrayOfEvent{
						Event: []BaseEvent{},
					},
				},
				{
					Name: "configStatus",
					Val:  ManagedEntityStatusGray},
				{
					Name: "customValue",
					Val: ArrayOfCustomFieldValue{
						CustomFieldValue: []BaseCustomFieldValue{},
					},
				},
				{
					Name: "declaredAlarmState",
					Val: ArrayOfAlarmState{
						AlarmState: []AlarmState{
							{
								Key: "alarm-328.group-d1",
								Entity: ManagedObjectReference{
									Type:  "Folder",
									Value: "group-d1"},
								Alarm: ManagedObjectReference{
									Type:  "Alarm",
									Value: "alarm-328"},
								OverallStatus: "gray",
								Time:          time.Date(2023, time.January, 14, 8, 57, 35, 279575000, time.UTC),
								Acknowledged:  NewBool(false),
							},
							{
								Key: "alarm-327.group-d1",
								Entity: ManagedObjectReference{
									Type:  "Folder",
									Value: "group-d1"},
								Alarm: ManagedObjectReference{
									Type:  "Alarm",
									Value: "alarm-327"},
								OverallStatus: "green",
								Time:          time.Date(2023, time.January, 14, 8, 56, 40, 83607000, time.UTC),
								Acknowledged:  NewBool(false),
								EventKey:      756,
							},
							{
								DynamicData: DynamicData{},
								Key:         "alarm-326.group-d1",
								Entity: ManagedObjectReference{
									Type:  "Folder",
									Value: "group-d1"},
								Alarm: ManagedObjectReference{
									Type:  "Alarm",
									Value: "alarm-326"},
								OverallStatus: "green",
								Time: time.Date(2023,
									time.January,
									14,
									8,
									56,
									35,
									82616000,
									time.UTC),
								Acknowledged: NewBool(false),
								EventKey:     751,
							},
						},
					},
				},
				{
					Name: "disabledMethod",
					Val: ArrayOfString{
						String: []string{},
					},
				},
				{
					Name: "effectiveRole",
					Val: ArrayOfInt{
						Int: []int32{-1},
					},
				},
				{
					Name: "name",
					Val:  "Datacenters"},
				{
					Name: "overallStatus",
					Val:  ManagedEntityStatusGray},
				{
					Name: "permission",
					Val: ArrayOfPermission{
						Permission: []Permission{
							{
								Entity: &ManagedObjectReference{
									Value: "group-d1",
									Type:  "Folder",
								},
								Principal: "VSPHERE.LOCAL\\vmware-vsm-2bd917c6-e084-4d1f-988d-a68f7525cc94",
								Group:     false,
								RoleId:    1034,
								Propagate: true},
							{
								Entity: &ManagedObjectReference{
									Value: "group-d1",
									Type:  "Folder",
								},
								Principal: "VSPHERE.LOCAL\\topologysvc-2bd917c6-e084-4d1f-988d-a68f7525cc94",
								Group:     false,
								RoleId:    1024,
								Propagate: true},
							{
								Entity: &ManagedObjectReference{
									Value: "group-d1",
									Type:  "Folder",
								},
								Principal: "VSPHERE.LOCAL\\vpxd-extension-2bd917c6-e084-4d1f-988d-a68f7525cc94",
								Group:     false,
								RoleId:    -1,
								Propagate: true},
						},
					},
				},
				{
					Name: "recentTask",
					Val: ArrayOfManagedObjectReference{
						ManagedObjectReference: []ManagedObjectReference{
							{
								Type:  "Task",
								Value: "task-186"},
							{
								Type:  "Task",
								Value: "task-187"},
							{
								Type:  "Task",
								Value: "task-188"},
						},
					},
				},
				{
					Name: "tag",
					Val: ArrayOfTag{
						Tag: []Tag{},
					},
				},
				{
					Name: "triggeredAlarmState",
					Val: ArrayOfAlarmState{
						AlarmState: []AlarmState{},
					},
				},
				{
					Name: "value",
					Val: ArrayOfCustomFieldValue{
						CustomFieldValue: []BaseCustomFieldValue{},
					},
				},
			},
			MissingSet: nil,
		},
	},
}

func mustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

func addrOfMustParseTime(layout, value string) *time.Time {
	t := mustParseTime(layout, value)
	return &t
}
