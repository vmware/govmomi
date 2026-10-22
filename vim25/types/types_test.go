// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"bytes"
	"reflect"
	"strings"
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

func TestHostBusAdapterVersionsDecode(t *testing.T) {
	for _, tc := range []struct {
		name            string
		elements        string
		driverVersion   string
		firmwareVersion string
	}{
		{"both", `<driverVersion>1.2.3</driverVersion><firmwareVersion>4.5.6</firmwareVersion>`, "1.2.3", "4.5.6"},
		{"driver only", `<driverVersion>1.2.3</driverVersion>`, "1.2.3", ""},
		{"firmware only", `<firmwareVersion>4.5.6</firmwareVersion>`, "", "4.5.6"},
		{"neither", "", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := `<HostStorageDeviceInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><hostBusAdapter xsi:type="HostFibreChannelHba"><device>vmhba1</device><driver>lpfc</driver>` + tc.elements + `</hostBusAdapter></HostStorageDeviceInfo>`
			var info HostStorageDeviceInfo
			dec := xml.NewDecoder(strings.NewReader(input))
			dec.TypeFunc = TypeFunc()
			if err := dec.Decode(&info); err != nil {
				t.Fatal(err)
			}
			if len(info.HostBusAdapter) != 1 {
				t.Fatalf("expected one adapter, got %d", len(info.HostBusAdapter))
			}
			assert.IsType(t, &HostFibreChannelHba{}, info.HostBusAdapter[0])
			hba := info.HostBusAdapter[0].GetHostHostBusAdapter()
			assert.Equal(t, "vmhba1", hba.Device)
			assert.Equal(t, "lpfc", hba.Driver)
			assert.Equal(t, tc.driverVersion, hba.DriverVersion)
			assert.Equal(t, tc.firmwareVersion, hba.FirmwareVersion)
		})
	}

	t.Run("malformed XML", func(t *testing.T) {
		var info HostStorageDeviceInfo
		dec := xml.NewDecoder(strings.NewReader(`<HostStorageDeviceInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><hostBusAdapter xsi:type="HostFibreChannelHba"><driverVersion>1.2.3</hostBusAdapter></HostStorageDeviceInfo>`))
		dec.TypeFunc = TypeFunc()
		assert.Error(t, dec.Decode(&info))
	})
}

func TestHostBusAdapterVersionsMarshal(t *testing.T) {
	for _, tc := range []struct {
		name    string
		obj     HostHostBusAdapter
		expXML  string
		expJSON string
	}{
		{
			name:    "populated",
			obj:     HostHostBusAdapter{DriverVersion: "1.2.3", FirmwareVersion: "4.5.6"},
			expXML:  `<HostHostBusAdapter><device></device><bus>0</bus><status></status><model></model><driverVersion>1.2.3</driverVersion><firmwareVersion>4.5.6</firmwareVersion></HostHostBusAdapter>`,
			expJSON: `{"_typeName":"HostHostBusAdapter","device":"","bus":0,"status":"","model":"","driverVersion":"1.2.3","firmwareVersion":"4.5.6"}`,
		},
		{
			name:    "empty",
			expXML:  `<HostHostBusAdapter><device></device><bus>0</bus><status></status><model></model></HostHostBusAdapter>`,
			expJSON: `{"_typeName":"HostHostBusAdapter","device":"","bus":0,"status":"","model":""}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("xml", func(t *testing.T) {
				act, err := xml.Marshal(tc.obj)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.expXML, string(act))
				var out HostHostBusAdapter
				if err := xml.Unmarshal(act, &out); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.obj, out)
			})
			t.Run("json", func(t *testing.T) {
				var w bytes.Buffer
				if err := NewJSONEncoder(&w).Encode(tc.obj); err != nil {
					t.Fatal(err)
				}
				assert.JSONEq(t, tc.expJSON, w.String())
			})
		})
	}
}
