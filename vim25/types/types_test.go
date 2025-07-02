// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
