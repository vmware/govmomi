// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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

var myTypes = map[string]reflect.Type{
	"MyType":      reflect.TypeOf(MyType{}),
	"ValueType":   reflect.TypeOf(ValueType{}),
	"PointerType": reflect.TypeOf(PointerType{}),
}

func MyTypes(name string) (reflect.Type, bool) {
	t, ok := myTypes[name]
	return t, ok
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
		{Value: ParseTime("2009-10-04T01:35:58+00:00")},
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
		dec.TypeFunc = MyTypes
		err = dec.Decode(&r2)
		if err != nil {
			t.Fatalf("Unmarshal: %s", err)
		}

		switch r1.Values[0].(type) {
		case time.Time:
			if !r1.Values[0].(time.Time).Equal(r2.Values[0].(time.Time)) {
				t.Errorf("Expected: %#v, actual: %#v", r1, r2)
			}
		default:
			if !reflect.DeepEqual(r1, r2) {
				t.Errorf("Expected: %#v, actual: %#v", r1, r2)
			}
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
	dec.TypeFunc = MyTypes
	err = dec.Decode(&r2)
	if err != nil {
		t.Fatalf("Unmarshal: %s", err)
	}

	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("expected: %#v, actual: %#v", r1, r2)
	}
}

type test3iface interface {
	Value() string
}

type test3a struct {
	V string `xml:",chardata"`
}

func (t test3a) Value() string { return t.V }

type test3b struct {
	V string `xml:",chardata"`
}

func (t test3b) Value() string { return t.V }

func TestUnmarshalInterfaceWithoutTypeAttr(t *testing.T) {
	var r struct {
		XMLName Name         `xml:"root"`
		Values  []test3iface `xml:"value,typeattr"`
	}

	b := `
	<root xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<value xsi:type="test3a">A</value>
	<value>B</value>
	</root>
	`

	fn := func(name string) (reflect.Type, bool) {
		switch name {
		case "test3a":
			return reflect.TypeOf(test3a{}), true
		case "test3iface":
			return reflect.TypeOf(test3b{}), true
		default:
			return nil, false
		}
	}

	dec := NewDecoder(bytes.NewReader([]byte(b)))
	dec.TypeFunc = fn
	err := dec.Decode(&r)
	if err != nil {
		t.Fatalf("Unmarshal: %s", err)
	}

	if len(r.Values) != 2 {
		t.Errorf("Expected 2 values")
	}

	exps := []struct {
		Typ reflect.Type
		Val string
	}{
		{
			Typ: reflect.TypeOf(test3a{}),
			Val: "A",
		},
		{
			Typ: reflect.TypeOf(test3b{}),
			Val: "B",
		},
	}

	for i, e := range exps {
		if val := r.Values[i].Value(); val != e.Val {
			t.Errorf("Expected: %s, got: %s", e.Val, val)
		}

		if typ := reflect.TypeOf(r.Values[i]); typ.Name() != e.Typ.Name() {
			t.Errorf("Expected: %s, got: %s", e.Typ.Name(), typ.Name())
		}
	}
}

// https://github.com/vmware/govmomi/issues/246
func TestNegativeValuesUnsignedFields(t *testing.T) {
	type T struct {
		I   string
		O   any
		U8  uint8  `xml:"u8"`
		U16 uint16 `xml:"u16"`
		U32 uint32 `xml:"u32"`
		U64 uint64 `xml:"u64"`
	}

	var tests = []T{
		{I: "<T><u8>-128</u8></T>", O: uint8(0x80)},
		{I: "<T><u8>-1</u8></T>", O: uint8(0xff)},
		{I: "<T><u16>-32768</u16></T>", O: uint16(0x8000)},
		{I: "<T><u16>-1</u16></T>", O: uint16(0xffff)},
		{I: "<T><u32>-2147483648</u32></T>", O: uint32(0x80000000)},
		{I: "<T><u32>-1</u32></T>", O: uint32(0xffffffff)},
		{I: "<T><u64>-9223372036854775808</u64></T>", O: uint64(0x8000000000000000)},
		{I: "<T><u64>-1</u64></T>", O: uint64(0xffffffffffffffff)},
	}

	for _, test := range tests {
		err := Unmarshal([]byte(test.I), &test)
		if err != nil {
			t.Errorf("Unmarshal error: %v", err)
			continue
		}

		var expected = test.O
		var actual any
		switch reflect.ValueOf(test.O).Type().Kind() {
		case reflect.Uint8:
			actual = test.U8
		case reflect.Uint16:
			actual = test.U16
		case reflect.Uint32:
			actual = test.U32
		case reflect.Uint64:
			actual = test.U64
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual: %v, expected: %v", actual, expected)
		}
	}
}
