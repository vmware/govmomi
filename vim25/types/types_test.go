/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestManagedObjectReference(t *testing.T) {

	testCases := []struct {
		name    string
		obj     ManagedObjectReference
		expXML  string
		expJSON string
	}{
		{
			name: "with server GUID",
			obj: ManagedObjectReference{
				Type:       "fake",
				Value:      "fake",
				ServerGUID: "fake",
			},
			expXML:  `<ManagedObjectReference type="fake" serverGuid="fake">fake</ManagedObjectReference>`,
			expJSON: `{"_typeName":"ManagedObjectReference","type":"fake","value":"fake","serverGuid":"fake"}`,
		},
		{
			name: "sans server GUID",
			obj: ManagedObjectReference{
				Type:  "fake",
				Value: "fake",
			},
			expXML:  `<ManagedObjectReference type="fake">fake</ManagedObjectReference>`,
			expJSON: `{"_typeName":"ManagedObjectReference","type":"fake","value":"fake"}`,
		},
	}

	for i := range testCases {
		tc := testCases[i] // capture the test case

		t.Run(tc.name, func(t *testing.T) {
			t.Run("xml", func(t *testing.T) {
				act, err := xml.Marshal(tc.obj)
				if err != nil {
					t.Fatal(err)
				}
				if e, a := tc.expXML, string(act); e != a {
					t.Fatalf("failed to marshal MoRef to XML: exp=%s, act=%s", e, a)
				}
			})
			t.Run("json", func(t *testing.T) {
				var w bytes.Buffer
				enc := NewJSONEncoder(&w)
				if err := enc.Encode(tc.obj); err != nil {
					t.Fatal(err)
				}
				assert.JSONEq(t, tc.expJSON, w.String(),
					"failed to marshal MoRef to JSON")
			})
		})
	}
}

func TestVirtualMachineConfigSpec(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}
}

func TestVirtualMachineAffinityInfo(t *testing.T) {
	// See https://github.com/vmware/govmomi/issues/1008
	in := VirtualMachineAffinityInfo{
		AffinitySet: []int32{0, 1, 2, 3},
	}

	b, err := xml.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	var out VirtualMachineAffinityInfo

	err = xml.Unmarshal(b, &out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(in, out) {
		t.Errorf("%#v vs %#v", in, out)
	}
}
