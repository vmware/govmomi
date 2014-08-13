/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package xml

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type MyType struct {
	Value string
}

func TestMarshalWithEmptyInterface(t *testing.T) {
	var r1, r2 struct {
		XMLName Name          `xml:"root"`
		Values  []interface{} `xml:"value,typeattr"`
	}

	var tests = []struct {
		Value interface{}
	}{
		{Value: bool(true)},
		{Value: int8(-8)},
		{Value: int16(-16)},
		{Value: int32(-32)},
		{Value: int64(-64)},
		{Value: uint8(8)},
		{Value: uint16(16)},
		{Value: uint32(32)},
		{Value: uint64(64)},
		{Value: float32(32.0)},
		{Value: float64(64.0)},
		{Value: string("string")},
		{Value: time.Now()},
		{Value: []byte("bytes")},
		{Value: MyType{Value: "v"}},
	}

	for _, test := range tests {
		r1.XMLName.Local = "root"
		r1.Values = []interface{}{test.Value}
		r2.XMLName = Name{}
		r2.Values = nil

		b, err := Marshal(r1)
		if err != nil {
			t.Fatalf("Marshal: %s", err)
		}

		dec := NewDecoder(bytes.NewReader(b))
		dec.AddType(reflect.TypeOf(MyType{}))
		err = dec.Decode(&r2)
		if err != nil {
			t.Fatalf("Unmarshal: %s", err)
		}

		if !reflect.DeepEqual(r1, r2) {
			t.Errorf("Expected: %#v, actual: %#v", r1, r2)
		}
	}
}

type VIntf interface {
	V() string
}

type ValueType struct {
	Value string `xml:",chardata"`
}

type PointerType struct {
	Value string `xml:",chardata"`
}

func (t ValueType) V() string {
	return t.Value
}

func (t *PointerType) V() string {
	return t.Value
}

func TestMarshalWithInterface(t *testing.T) {
	var r1, r2 struct {
		XMLName Name    `xml:"root"`
		Values  []VIntf `xml:"value,typeattr"`
	}

	r1.XMLName.Local = "root"
	r1.Values = []VIntf{
		ValueType{"v1"},
		&PointerType{"v2"},
	}

	b, err := Marshal(r1)
	if err != nil {
		t.Fatalf("Marshal: %s", err)
	}

	dec := NewDecoder(bytes.NewReader(b))
	dec.AddType(reflect.TypeOf(ValueType{}))
	dec.AddType(reflect.TypeOf(PointerType{}))
	err = dec.Decode(&r2)
	if err != nil {
		t.Fatalf("Unmarshal: %s", err)
	}

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("expected: %#v, actual: %#v", r1, r2)
	}
}
