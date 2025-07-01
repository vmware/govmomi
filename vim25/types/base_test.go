// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestAnyType(t *testing.T) {
	x := func(s string) []byte {
		s = `<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">` + s
		s += `</root>`
		return []byte(s)
	}

	tests := []struct {
		Input []byte
		Value any
	}{
		{
			Input: x(`<name xsi:type="xsd:string">test</name>`),
			Value: "test",
		},
		{
			Input: x(`<name xsi:type="ArrayOfString"><string>AA</string><string>BB</string></name>`),
			Value: ArrayOfString{String: []string{"AA", "BB"}},
		},
	}

	for _, test := range tests {
		var r struct {
			A any `xml:"name,typeattr"`
		}

		dec := xml.NewDecoder(bytes.NewReader(test.Input))
		dec.TypeFunc = TypeFunc()

		err := dec.Decode(&r)
		if err != nil {
			t.Fatalf("Decode: %s", err)
		}

		if !reflect.DeepEqual(r.A, test.Value) {
			t.Errorf("Expected: %#v, actual: %#v", r.A, test.Value)
		}
	}
}
