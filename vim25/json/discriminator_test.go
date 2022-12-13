// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type DS1 struct {
	F1 interface{} `json:"f1"`
}

type DS2 struct {
	F1 noop1 `json:"f1"`
}

type DS3 struct {
	F1 string `json:"f1"`
}

type DS3Noop1 DS3

func (v DS3Noop1) noop1() {}

type DS4 struct {
	F1 string      `json:"f1"`
	F2 interface{} `json:"f2"`
}

type DS4Noop1 DS4

func (v DS4Noop1) noop1() {}

type DS5 struct {
	F1 string `json:"f1"`
	F2 noop1  `json:"f2"`
}

type DS5Noop1 DS5

func (v DS5Noop1) noop1() {}

type DS6 struct {
	F1 emptyIface `json:"f1"`
}

type DS7 struct {
	F1 noop2 `json:"f1"`
}

var discriminatorTests = []struct {
	obj interface{}
	str string
	tf  string
	vf  string
	af  string
}{
	// primitive values stored in interface with 0 methods
	{obj: DS1{F1: uint(1)}, str: `{"_t":"DS1","f1":{"_t":"uint","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: uint8(1)}, str: `{"_t":"DS1","f1":{"_t":"uint8","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: uint16(1)}, str: `{"_t":"DS1","f1":{"_t":"uint16","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: uint32(1)}, str: `{"_t":"DS1","f1":{"_t":"uint32","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: uint64(1)}, str: `{"_t":"DS1","f1":{"_t":"uint64","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: uintptr(1)}, str: `{"_t":"DS1","f1":{"_t":"uintptr","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS6{F1: int(-1)}, str: `{"_t":"DS6","f1":{"_t":"int","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: int8(-1)}, str: `{"_t":"DS6","f1":{"_t":"int8","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: int16(-1)}, str: `{"_t":"DS6","f1":{"_t":"int16","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: int32(-1)}, str: `{"_t":"DS6","f1":{"_t":"int32","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: int64(-1)}, str: `{"_t":"DS6","f1":{"_t":"int64","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS1{F1: float32(-1.0)}, str: `{"_t":"DS1","f1":{"_t":"float32","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: float64(1.0)}, str: `{"_t":"DS1","f1":{"_t":"float64","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: float32(-1.1)}, str: `{"_t":"DS1","f1":{"_t":"float32","_v":-1.1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: float64(1.1)}, str: `{"_t":"DS1","f1":{"_t":"float64","_v":1.1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS1{F1: "hello"}, str: `{"_t":"DS1","f1":{"_t":"string","_v":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: true}, str: `{"_t":"DS1","f1":{"_t":"bool","_v":true}}`, tf: "_t", vf: "_v", af: "_a"},

	// address of primitive values stored in interface with 0 methods
	{obj: DS1{F1: addrOfUint(1)}, str: `{"_t":"DS1","f1":{"_t":"uint","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfUint8(1)}, str: `{"_t":"DS1","f1":{"_t":"uint8","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfUint16(1)}, str: `{"_t":"DS1","f1":{"_t":"uint16","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfUint32(1)}, str: `{"_t":"DS1","f1":{"_t":"uint32","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfUint64(1)}, str: `{"_t":"DS1","f1":{"_t":"uint64","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfUintptr(1)}, str: `{"_t":"DS1","f1":{"_t":"uintptr","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS6{F1: addrOfInt(-1)}, str: `{"_t":"DS6","f1":{"_t":"int","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: addrOfInt8(-1)}, str: `{"_t":"DS6","f1":{"_t":"int8","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: addrOfInt16(-1)}, str: `{"_t":"DS6","f1":{"_t":"int16","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: addrOfInt32(-1)}, str: `{"_t":"DS6","f1":{"_t":"int32","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS6{F1: addrOfInt64(-1)}, str: `{"_t":"DS6","f1":{"_t":"int64","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS1{F1: addrOfFloat32(-1.0)}, str: `{"_t":"DS1","f1":{"_t":"float32","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfFloat64(1.0)}, str: `{"_t":"DS1","f1":{"_t":"float64","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfFloat32(-1.1)}, str: `{"_t":"DS1","f1":{"_t":"float32","_a":true,"_v":-1.1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfFloat64(1.1)}, str: `{"_t":"DS1","f1":{"_t":"float64","_a":true,"_v":1.1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS1{F1: addrOfString("hello")}, str: `{"_t":"DS1","f1":{"_t":"string","_a":true,"_v":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: addrOfBool(true)}, str: `{"_t":"DS1","f1":{"_t":"bool","_a":true,"_v":true}}`, tf: "_t", vf: "_v", af: "_a"},

	// primitive values stored in interface with >0 methods
	{obj: DS2{F1: uintNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"uintNoop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: uint8Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint8Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: uint16Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint16Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: uint32Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint32Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: uint64Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint64Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: uintptrNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"uintptrNoop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: intNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"intNoop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: int8Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int8Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: int16Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int16Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: int32Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int32Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: int64Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int64Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: float32Noop(-1.0)}, str: `{"_t":"DS2","f1":{"_t":"float32Noop","_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: float64Noop(1.0)}, str: `{"_t":"DS2","f1":{"_t":"float64Noop","_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: float32Noop(-1.1)}, str: `{"_t":"DS2","f1":{"_t":"float32Noop","_v":-1.1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: float64Noop(1.1)}, str: `{"_t":"DS2","f1":{"_t":"float64Noop","_v":1.1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: stringNoop("hello")}, str: `{"_t":"DS2","f1":{"_t":"stringNoop","_v":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: boolNoop(true)}, str: `{"_t":"DS2","f1":{"_t":"boolNoop","_v":true}}`, tf: "_t", vf: "_v", af: "_a"},

	// address of primitive values stored in interface with >0 methods
	{obj: DS2{F1: addrOfUintNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"uintNoop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfUint8Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint8Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfUint16Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint16Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfUint32Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint32Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfUint64Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"uint64Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfUintptrNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"uintptrNoop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: addrOfIntNoop(1)}, str: `{"_t":"DS2","f1":{"_t":"intNoop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfInt8Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int8Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfInt16Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int16Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfInt32Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int32Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfInt64Noop(1)}, str: `{"_t":"DS2","f1":{"_t":"int64Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: addrOfFloat32Noop(-1.0)}, str: `{"_t":"DS2","f1":{"_t":"float32Noop","_a":true,"_v":-1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfFloat64Noop(1.0)}, str: `{"_t":"DS2","f1":{"_t":"float64Noop","_a":true,"_v":1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfFloat32Noop(-1.1)}, str: `{"_t":"DS2","f1":{"_t":"float32Noop","_a":true,"_v":-1.1}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfFloat64Noop(1.1)}, str: `{"_t":"DS2","f1":{"_t":"float64Noop","_a":true,"_v":1.1}}`, tf: "_t", vf: "_v", af: "_a"},

	{obj: DS2{F1: addrOfStringNoop("hello")}, str: `{"_t":"DS2","f1":{"_t":"stringNoop","_a":true,"_v":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: addrOfBoolNoop(true)}, str: `{"_t":"DS2","f1":{"_t":"boolNoop","_a":true,"_v":true}}`, tf: "_t", vf: "_v", af: "_a"},

	// address of primitive values stored in interface with >0 methods with empty byAddrFieldName
	//
	// Unmarshaling the JSON into the object causes the deocder to check to see if the JSON objects
	// can be stored in DS7.F1, finding out that they cannot due to all the types implementing
	// DS7.F1 by-address. Thus the decoder will then check to see if the value can be assigned to
	// DS7.F1 by-address, which will work. This illustrates that even if the byAddrFieldName is
	// omitted from JSON, the decoder still attempts to store the object in an interface if a
	// pointer to the object implements the interface.
	{obj: DS7{F1: addrOfUintNoop(1)}, str: `{"_t":"DS7","f1":{"_t":"uintNoop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfUint8Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"uint8Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfUint16Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"uint16Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfUint32Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"uint32Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfUint64Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"uint64Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfUintptrNoop(1)}, str: `{"_t":"DS7","f1":{"_t":"uintptrNoop","_v":1}}`, tf: "_t", vf: "_v"},

	{obj: DS7{F1: addrOfIntNoop(1)}, str: `{"_t":"DS7","f1":{"_t":"intNoop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfInt8Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"int8Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfInt16Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"int16Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfInt32Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"int32Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfInt64Noop(1)}, str: `{"_t":"DS7","f1":{"_t":"int64Noop","_v":1}}`, tf: "_t", vf: "_v"},

	{obj: DS7{F1: addrOfFloat32Noop(-1.0)}, str: `{"_t":"DS7","f1":{"_t":"float32Noop","_v":-1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfFloat64Noop(1.0)}, str: `{"_t":"DS7","f1":{"_t":"float64Noop","_v":1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfFloat32Noop(-1.1)}, str: `{"_t":"DS7","f1":{"_t":"float32Noop","_v":-1.1}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfFloat64Noop(1.1)}, str: `{"_t":"DS7","f1":{"_t":"float64Noop","_v":1.1}}`, tf: "_t", vf: "_v"},

	{obj: DS7{F1: addrOfStringNoop("hello")}, str: `{"_t":"DS7","f1":{"_t":"stringNoop","_v":"hello"}}`, tf: "_t", vf: "_v"},
	{obj: DS7{F1: addrOfBoolNoop(true)}, str: `{"_t":"DS7","f1":{"_t":"boolNoop","_v":true}}`, tf: "_t", vf: "_v"},

	// struct value stored in interface with 0 methods
	{obj: DS1{F1: DS3{F1: "hello"}}, str: `{"_t":"DS1","f1":{"_t":"DS3","f1":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: DS4{F1: "hello", F2: int(1)}}, str: `{"_t":"DS1","f1":{"_t":"DS4","f1":"hello","f2":{"_t":"int","_v":1}}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: DS4{F1: "hello", F2: DS3{F1: "world"}}}, str: `{"_t":"DS1","f1":{"_t":"DS4","f1":"hello","f2":{"_t":"DS3","f1":"world"}}}`, tf: "_t", vf: "_v", af: "_a"},

	// struct value stored in interface with >0 methods
	{obj: DS2{F1: DS3Noop1{F1: "hello"}}, str: `{"_t":"DS2","f1":{"_t":"DS3Noop1","f1":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: DS4Noop1{F1: "hello", F2: int(1)}}, str: `{"_t":"DS2","f1":{"_t":"DS4Noop1","f1":"hello","f2":{"_t":"int","_v":1}}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: DS5Noop1{F1: "hello", F2: DS3Noop1{F1: "world"}}}, str: `{"_t":"DS2","f1":{"_t":"DS5Noop1","f1":"hello","f2":{"_t":"DS3Noop1","f1":"world"}}}`, tf: "_t", vf: "_v", af: "_a"},

	// address of struct value stored in interface with 0 methods
	{obj: DS1{F1: &DS3{F1: "hello"}}, str: `{"_t":"DS1","f1":{"_t":"DS3","_a":true,"f1":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS1{F1: DS4{F1: "hello", F2: &DS3{F1: "world"}}}, str: `{"_t":"DS1","f1":{"_t":"DS4","f1":"hello","f2":{"_t":"DS3","_a":true,"f1":"world"}}}`, tf: "_t", vf: "_v", af: "_a"},

	// address of struct value stored in interface with >0 methods
	{obj: DS2{F1: DS3Noop1{F1: "hello"}}, str: `{"_t":"DS2","f1":{"_t":"DS3Noop1","f1":"hello"}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: DS4Noop1{F1: "hello", F2: int(1)}}, str: `{"_t":"DS2","f1":{"_t":"DS4Noop1","f1":"hello","f2":{"_t":"int","_v":1}}}`, tf: "_t", vf: "_v", af: "_a"},
	{obj: DS2{F1: DS5Noop1{F1: "hello", F2: DS3Noop1{F1: "world"}}}, str: `{"_t":"DS2","f1":{"_t":"DS5Noop1","f1":"hello","f2":{"_t":"DS3Noop1","f1":"world"}}}`, tf: "_t", vf: "_v", af: "_a"},
}

func discriminatorToTypeFn(discriminator string) (reflect.Type, bool) {
	switch discriminator {
	case "DS1":
		return reflect.TypeOf(DS1{}), true
	case "DS2":
		return reflect.TypeOf(DS2{}), true
	case "DS3":
		return reflect.TypeOf(DS3{}), true
	case "DS3Noop1":
		return reflect.TypeOf(DS3Noop1{}), true
	case "DS4":
		return reflect.TypeOf(DS4{}), true
	case "DS4Noop1":
		return reflect.TypeOf(DS4Noop1{}), true
	case "DS5":
		return reflect.TypeOf(DS5{}), true
	case "DS5Noop1":
		return reflect.TypeOf(DS5Noop1{}), true
	case "DS6":
		return reflect.TypeOf(DS6{}), true
	case "DS7":
		return reflect.TypeOf(DS7{}), true
	case "uintNoop":
		return reflect.TypeOf(uintNoop(0)), true
	case "uint8Noop":
		return reflect.TypeOf(uint8Noop(0)), true
	case "uint16Noop":
		return reflect.TypeOf(uint16Noop(0)), true
	case "uint32Noop":
		return reflect.TypeOf(uint32Noop(0)), true
	case "uint64Noop":
		return reflect.TypeOf(uint64Noop(0)), true
	case "uintptrNoop":
		return reflect.TypeOf(uintptrNoop(0)), true
	case "intNoop":
		return reflect.TypeOf(intNoop(0)), true
	case "int8Noop":
		return reflect.TypeOf(int8Noop(0)), true
	case "int16Noop":
		return reflect.TypeOf(int16Noop(0)), true
	case "int32Noop":
		return reflect.TypeOf(int32Noop(0)), true
	case "int64Noop":
		return reflect.TypeOf(int64Noop(0)), true
	case "float32Noop":
		return reflect.TypeOf(float32Noop(0)), true
	case "float64Noop":
		return reflect.TypeOf(float64Noop(0)), true
	case "boolNoop":
		return reflect.TypeOf(boolNoop(true)), true
	case "stringNoop":
		return reflect.TypeOf(stringNoop("")), true
	default:
		return nil, false
	}
}

func TestDiscriminator(t *testing.T) {
	t.Run("Encode", testDiscriminatorEncode)
	t.Run("Decode", testDiscriminatorDecode)
}

func testDiscriminatorEncode(t *testing.T) {
	for _, tc := range discriminatorTests {
		tc := tc
		t.Run("", func(t *testing.T) {
			var w bytes.Buffer
			enc := NewEncoder(&w)
			enc.SetDiscriminator(tc.tf, tc.vf, tc.af)
			if err := enc.Encode(tc.obj); err != nil {
				t.Errorf("unexpected error: %v", err)
			} else if a, e := w.String(), tc.str; a != e+"\n" {
				t.Errorf("mismatch: e=%s, a=%s", e, a)
			}
		})
	}
}


func testDiscriminatorDecode(t *testing.T) {
	for _, tc := range discriminatorTests {
		tc := tc
		t.Run("", func(t *testing.T) {
			typ := reflect.TypeOf(tc.obj)
			val := reflect.New(typ)
			obj := val.Elem().Interface()

			dec := NewDecoder(strings.NewReader(tc.str))
			dec.SetDiscriminator(tc.tf, tc.vf, tc.af, discriminatorToTypeFn)
			if err := dec.Decode(&obj); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			a, e := obj, tc.obj
			if diff := cmp.Diff(a, e); diff != "" {
				t.Errorf("mismatch: e=%+v, a=%+v, diff=%s", e, a, diff)
			}
		})
	}
}

func addrOfUint(v uint) *uint {
	return &v
}
func addrOfUint8(v uint8) *uint8 {
	return &v
}
func addrOfUint16(v uint16) *uint16 {
	return &v
}
func addrOfUint32(v uint32) *uint32 {
	return &v
}
func addrOfUint64(v uint64) *uint64 {
	return &v
}
func addrOfUintptr(v uintptr) *uintptr {
	return &v
}

func addrOfInt(v int) *int {
	return &v
}
func addrOfInt8(v int8) *int8 {
	return &v
}
func addrOfInt16(v int16) *int16 {
	return &v
}
func addrOfInt32(v int32) *int32 {
	return &v
}
func addrOfInt64(v int64) *int64 {
	return &v
}

func addrOfFloat32(v float32) *float32 {
	return &v
}
func addrOfFloat64(v float64) *float64 {
	return &v
}

func addrOfBool(v bool) *bool {
	return &v
}
func addrOfString(v string) *string {
	return &v
}

type emptyIface interface{}

type noop1 interface {
	noop1()
}

type noop2 interface {
	noop2()
}

type uintNoop uint

func (v uintNoop) noop1()  {}
func (v *uintNoop) noop2() {}

type uint8Noop uint8

func (v uint8Noop) noop1()  {}
func (v *uint8Noop) noop2() {}

type uint16Noop uint16

func (v uint16Noop) noop1()  {}
func (v *uint16Noop) noop2() {}

type uint32Noop uint32

func (v uint32Noop) noop1()  {}
func (v *uint32Noop) noop2() {}

type uint64Noop uint64

func (v uint64Noop) noop1()  {}
func (v *uint64Noop) noop2() {}

type uintptrNoop uintptr

func (v uintptrNoop) noop1()  {}
func (v *uintptrNoop) noop2() {}

type intNoop int

func (v intNoop) noop1()  {}
func (v *intNoop) noop2() {}

type int8Noop int8

func (v int8Noop) noop1()  {}
func (v *int8Noop) noop2() {}

type int16Noop int16

func (v int16Noop) noop1()  {}
func (v *int16Noop) noop2() {}

type int32Noop int32

func (v int32Noop) noop1()  {}
func (v *int32Noop) noop2() {}

type int64Noop int64

func (v int64Noop) noop1()  {}
func (v *int64Noop) noop2() {}

type float32Noop float32

func (v float32Noop) noop1()  {}
func (v *float32Noop) noop2() {}

type float64Noop float64

func (v float64Noop) noop1()  {}
func (v *float64Noop) noop2() {}

type stringNoop string

func (v stringNoop) noop1()  {}
func (v *stringNoop) noop2() {}

type boolNoop bool

func (v boolNoop) noop1()  {}
func (v *boolNoop) noop2() {}

func addrOfUintNoop(v uintNoop) *uintNoop {
	return &v
}
func addrOfUint8Noop(v uint8Noop) *uint8Noop {
	return &v
}
func addrOfUint16Noop(v uint16Noop) *uint16Noop {
	return &v
}
func addrOfUint32Noop(v uint32Noop) *uint32Noop {
	return &v
}
func addrOfUint64Noop(v uint64Noop) *uint64Noop {
	return &v
}
func addrOfUintptrNoop(v uintptrNoop) *uintptrNoop {
	return &v
}

func addrOfIntNoop(v intNoop) *intNoop {
	return &v
}
func addrOfInt8Noop(v int8Noop) *int8Noop {
	return &v
}
func addrOfInt16Noop(v int16Noop) *int16Noop {
	return &v
}
func addrOfInt32Noop(v int32Noop) *int32Noop {
	return &v
}
func addrOfInt64Noop(v int64Noop) *int64Noop {
	return &v
}

func addrOfFloat32Noop(v float32Noop) *float32Noop {
	return &v
}
func addrOfFloat64Noop(v float64Noop) *float64Noop {
	return &v
}

func addrOfBoolNoop(v boolNoop) *boolNoop {
	return &v
}
func addrOfStringNoop(v stringNoop) *stringNoop {
	return &v
}
