// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	json "github.com/vmware/govmomi/vim25/json"
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

type DS8 struct {
	F1 DS3 `json:"f1"`
}

type DS9 struct {
	F1 int64
}

func (d DS9) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(strconv.FormatInt(d.F1, 10))
	return b.Bytes(), nil
}
func (d *DS9) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		d.F1 = 0
		return nil
	}
	if len(s) < 1 {
		return fmt.Errorf("Cannot parse empty as int64")
	}
	v, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return fmt.Errorf("Cannot parse as int64: %v; %w", s, e)
	}
	d.F1 = v
	return nil
}

// Struct implementing UnmarshalJSON with value receiver.
type DS10 struct {
	DS9 *DS9
}

func (d DS10) UnmarshalJSON(b []byte) error {
	if d.DS9 == nil {
		return nil
	}
	return d.DS9.UnmarshalJSON(b)
}
func (d DS10) MarshalJSON() ([]byte, error) {
	if d.DS9 == nil {
		return []byte("null"), nil
	}
	return d.DS9.MarshalJSON()
}

type HexInt64 int64

func (i HexInt64) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(`"%X"`, i))
	return b.Bytes(), nil
}

func (i *HexInt64) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*i = 0
		return nil
	}
	last := len(s) - 1
	if last < 1 || s[0] != '"' || s[last] != '"' {
		return fmt.Errorf("Cannot parse as hex int64: %v", s)
	}
	v, e := strconv.ParseInt(s[1:last], 16, 64)
	if e != nil {
		return fmt.Errorf("Cannot parse as hex int64: %v; %w", s, e)
	}
	*i = HexInt64(v)
	return nil
}

func customNameWithFilter(t reflect.Type) string {
	res := json.DefaultDiscriminatorFunc(t)
	if res == "DS3" {
		return ""
	}
	return res
}

var discriminatorTests = []struct {
	obj               interface{}
	str               string
	expObj            interface{}
	expStr            string
	expEncErr         string
	expDecErr         string
	tf                string
	vf                string
	mode              json.DiscriminatorEncodeMode
	dd                bool
	discriminatorFunc json.TypeToDiscriminatorFunc
}{

	// invalid type/value combinations
	{obj: true, str: `{"_t":"string","_v":true}`, expStr: `{"_t":"bool","_v":true}`, expDecErr: `json: cannot unmarshal bool into Go value of type string`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: DS1{F1: true}, str: `{"f1":{"_t":"string","_v":true}}`, expStr: `{"f1":{"_t":"bool","_v":true}}`, expDecErr: `json: cannot unmarshal bool into Go struct field DS1.f1 of type string`},

	// encode/decode nil/null works as expected
	{obj: nil, str: `null`},
	{obj: nil, str: `null`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	// encode/decode number works as expected
	{obj: float64(3), str: `3`},

	// encode/decode empty string works as expected
	{obj: "", str: `""`},

	// encode/decode boolean works as expected
	{obj: true, str: `true`},
	{obj: false, str: `false`},

	// primitive values with root object encoded with type name
	{obj: uint(1), str: `{"_t":"uint","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: uint8(1), str: `{"_t":"uint8","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: uint16(1), str: `{"_t":"uint16","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: uint32(1), str: `{"_t":"uint32","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: uint64(1), str: `{"_t":"uint64","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: uintptr(1), str: `{"_t":"uintptr","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	{obj: int(-1), str: `{"_t":"int","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: int8(-1), str: `{"_t":"int8","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: int16(-1), str: `{"_t":"int16","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: int32(-1), str: `{"_t":"int32","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: int64(-1), str: `{"_t":"int64","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	{obj: float32(-1.0), str: `{"_t":"float32","_v":-1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: float64(1.0), str: `{"_t":"float64","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: float32(-1.1), str: `{"_t":"float32","_v":-1.1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: float64(1.1), str: `{"_t":"float64","_v":1.1}`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	{obj: "hello", str: `{"_t":"string","_v":"hello"}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: true, str: `{"_t":"bool","_v":true}`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	{obj: HexInt64(42), str: `{"_t":"HexInt64","_v":"2A"}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: DS9{F1: 42}, str: `{"_t":"DS9","_v":42}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: DS6{F1: DS9{F1: 42}}, str: `{"f1":{"_t":"DS9","_v":42}}`},
	{obj: DS9{F1: 42}, str: `42`},

	{obj: DS10{DS9: &DS9{F1: 42}}, str: `42`},
	{obj: DS6{F1: DS10{DS9: &DS9{F1: 42}}}, str: `{"f1":{"_t":"DS10","_v":42}}`, expObj: DS6{F1: DS10{DS9: nil}}},

	// primitive values stored in interface with 0 methods
	{obj: DS1{F1: uint(1)}, str: `{"f1":{"_t":"uint","_v":1}}`},
	{obj: DS1{F1: uint8(1)}, str: `{"f1":{"_t":"uint8","_v":1}}`},
	{obj: DS1{F1: uint16(1)}, str: `{"f1":{"_t":"uint16","_v":1}}`},
	{obj: DS1{F1: uint32(1)}, str: `{"f1":{"_t":"uint32","_v":1}}`},
	{obj: DS1{F1: uint64(1)}, str: `{"f1":{"_t":"uint64","_v":1}}`},
	{obj: DS1{F1: uintptr(1)}, str: `{"f1":{"_t":"uintptr","_v":1}}`},

	{obj: DS6{F1: int(-1)}, str: `{"f1":{"_t":"int","_v":-1}}`},
	{obj: DS6{F1: int8(-1)}, str: `{"f1":{"_t":"int8","_v":-1}}`},
	{obj: DS6{F1: int16(-1)}, str: `{"f1":{"_t":"int16","_v":-1}}`},
	{obj: DS6{F1: int32(-1)}, str: `{"f1":{"_t":"int32","_v":-1}}`},
	{obj: DS6{F1: int64(-1)}, str: `{"f1":{"_t":"int64","_v":-1}}`},

	{obj: DS1{F1: float32(-1.0)}, str: `{"f1":{"_t":"float32","_v":-1}}`},
	{obj: DS1{F1: float64(1.0)}, str: `{"f1":{"_t":"float64","_v":1}}`},
	{obj: DS1{F1: float32(-1.1)}, str: `{"f1":{"_t":"float32","_v":-1.1}}`},
	{obj: DS1{F1: float64(1.1)}, str: `{"f1":{"_t":"float64","_v":1.1}}`},

	{obj: DS1{F1: "hello"}, str: `{"f1":{"_t":"string","_v":"hello"}}`},
	{obj: DS1{F1: true}, str: `{"f1":{"_t":"bool","_v":true}}`},

	// address of primitive values stored in interface with 0 methods
	{obj: DS1{F1: addrOfUint(1)}, str: `{"f1":{"_t":"uint","_v":1}}`, expObj: DS1{F1: uint(1)}},
	{obj: DS1{F1: addrOfUint8(1)}, str: `{"f1":{"_t":"uint8","_v":1}}`, expObj: DS1{F1: uint8(1)}},
	{obj: DS1{F1: addrOfUint16(1)}, str: `{"f1":{"_t":"uint16","_v":1}}`, expObj: DS1{F1: uint16(1)}},
	{obj: DS1{F1: addrOfUint32(1)}, str: `{"f1":{"_t":"uint32","_v":1}}`, expObj: DS1{F1: uint32(1)}},
	{obj: DS1{F1: addrOfUint64(1)}, str: `{"f1":{"_t":"uint64","_v":1}}`, expObj: DS1{F1: uint64(1)}},
	{obj: DS1{F1: addrOfUintptr(1)}, str: `{"f1":{"_t":"uintptr","_v":1}}`, expObj: DS1{F1: uintptr(1)}},

	{obj: DS6{F1: addrOfInt(-1)}, str: `{"f1":{"_t":"int","_v":-1}}`, expObj: DS6{F1: int(-1)}},
	{obj: DS6{F1: addrOfInt8(-1)}, str: `{"f1":{"_t":"int8","_v":-1}}`, expObj: DS6{F1: int8(-1)}},
	{obj: DS6{F1: addrOfInt16(-1)}, str: `{"f1":{"_t":"int16","_v":-1}}`, expObj: DS6{F1: int16(-1)}},
	{obj: DS6{F1: addrOfInt32(-1)}, str: `{"f1":{"_t":"int32","_v":-1}}`, expObj: DS6{F1: int32(-1)}},
	{obj: DS6{F1: addrOfInt64(-1)}, str: `{"f1":{"_t":"int64","_v":-1}}`, expObj: DS6{F1: int64(-1)}},

	{obj: DS1{F1: addrOfFloat32(-1.0)}, str: `{"f1":{"_t":"float32","_v":-1}}`, expObj: DS1{F1: float32(-1.0)}},
	{obj: DS1{F1: addrOfFloat64(1.0)}, str: `{"f1":{"_t":"float64","_v":1}}`, expObj: DS1{F1: float64(1)}},
	{obj: DS1{F1: addrOfFloat32(-1.1)}, str: `{"f1":{"_t":"float32","_v":-1.1}}`, expObj: DS1{F1: float32(-1.1)}},
	{obj: DS1{F1: addrOfFloat64(1.1)}, str: `{"f1":{"_t":"float64","_v":1.1}}`, expObj: DS1{F1: float64(1.1)}},

	{obj: DS1{F1: addrOfString("hello")}, str: `{"f1":{"_t":"string","_v":"hello"}}`, expObj: DS1{F1: "hello"}},
	{obj: DS1{F1: addrOfBool(true)}, str: `{"f1":{"_t":"bool","_v":true}}`, expObj: DS1{F1: true}},

	// primitive values stored in interface with >0 methods
	{obj: DS2{F1: uintNoop(1)}, str: `{"f1":{"_t":"uintNoop","_v":1}}`},
	{obj: DS2{F1: uint8Noop(1)}, str: `{"f1":{"_t":"uint8Noop","_v":1}}`},
	{obj: DS2{F1: uint16Noop(1)}, str: `{"f1":{"_t":"uint16Noop","_v":1}}`},
	{obj: DS2{F1: uint32Noop(1)}, str: `{"f1":{"_t":"uint32Noop","_v":1}}`},
	{obj: DS2{F1: uint64Noop(1)}, str: `{"f1":{"_t":"uint64Noop","_v":1}}`},
	{obj: DS2{F1: uintptrNoop(1)}, str: `{"f1":{"_t":"uintptrNoop","_v":1}}`},

	{obj: DS2{F1: intNoop(1)}, str: `{"f1":{"_t":"intNoop","_v":1}}`},
	{obj: DS2{F1: int8Noop(1)}, str: `{"f1":{"_t":"int8Noop","_v":1}}`},
	{obj: DS2{F1: int16Noop(1)}, str: `{"f1":{"_t":"int16Noop","_v":1}}`},
	{obj: DS2{F1: int32Noop(1)}, str: `{"f1":{"_t":"int32Noop","_v":1}}`},
	{obj: DS2{F1: int64Noop(1)}, str: `{"f1":{"_t":"int64Noop","_v":1}}`},

	{obj: DS2{F1: float32Noop(-1.0)}, str: `{"f1":{"_t":"float32Noop","_v":-1}}`},
	{obj: DS2{F1: float64Noop(1.0)}, str: `{"f1":{"_t":"float64Noop","_v":1}}`},
	{obj: DS2{F1: float32Noop(-1.1)}, str: `{"f1":{"_t":"float32Noop","_v":-1.1}}`},
	{obj: DS2{F1: float64Noop(1.1)}, str: `{"f1":{"_t":"float64Noop","_v":1.1}}`},

	{obj: DS2{F1: stringNoop("hello")}, str: `{"f1":{"_t":"stringNoop","_v":"hello"}}`},
	{obj: DS2{F1: boolNoop(true)}, str: `{"f1":{"_t":"boolNoop","_v":true}}`},

	// address of primitive values stored in interface with >0 methods
	{obj: DS2{F1: addrOfUintNoop(1)}, str: `{"f1":{"_t":"uintNoop","_v":1}}`, expObj: DS2{F1: uintNoop(1)}},
	{obj: DS2{F1: addrOfUint8Noop(1)}, str: `{"f1":{"_t":"uint8Noop","_v":1}}`, expObj: DS2{F1: uint8Noop(1)}},
	{obj: DS2{F1: addrOfUint16Noop(1)}, str: `{"f1":{"_t":"uint16Noop","_v":1}}`, expObj: DS2{F1: uint16Noop(1)}},
	{obj: DS2{F1: addrOfUint32Noop(1)}, str: `{"f1":{"_t":"uint32Noop","_v":1}}`, expObj: DS2{F1: uint32Noop(1)}},
	{obj: DS2{F1: addrOfUint64Noop(1)}, str: `{"f1":{"_t":"uint64Noop","_v":1}}`, expObj: DS2{F1: uint64Noop(1)}},
	{obj: DS2{F1: addrOfUintptrNoop(1)}, str: `{"f1":{"_t":"uintptrNoop","_v":1}}`, expObj: DS2{F1: uintptrNoop(1)}},

	{obj: DS2{F1: addrOfIntNoop(1)}, str: `{"f1":{"_t":"intNoop","_v":1}}`, expObj: DS2{F1: intNoop(1)}},
	{obj: DS2{F1: addrOfInt8Noop(1)}, str: `{"f1":{"_t":"int8Noop","_v":1}}`, expObj: DS2{F1: int8Noop(1)}},
	{obj: DS2{F1: addrOfInt16Noop(1)}, str: `{"f1":{"_t":"int16Noop","_v":1}}`, expObj: DS2{F1: int16Noop(1)}},
	{obj: DS2{F1: addrOfInt32Noop(1)}, str: `{"f1":{"_t":"int32Noop","_v":1}}`, expObj: DS2{F1: int32Noop(1)}},
	{obj: DS2{F1: addrOfInt64Noop(1)}, str: `{"f1":{"_t":"int64Noop","_v":1}}`, expObj: DS2{F1: int64Noop(1)}},

	{obj: DS2{F1: addrOfFloat32Noop(-1.0)}, str: `{"f1":{"_t":"float32Noop","_v":-1}}`, expObj: DS2{F1: float32Noop(-1.0)}},
	{obj: DS2{F1: addrOfFloat64Noop(1.0)}, str: `{"f1":{"_t":"float64Noop","_v":1}}`, expObj: DS2{F1: float64Noop(1.0)}},
	{obj: DS2{F1: addrOfFloat32Noop(-1.1)}, str: `{"f1":{"_t":"float32Noop","_v":-1.1}}`, expObj: DS2{F1: float32Noop(-1.1)}},
	{obj: DS2{F1: addrOfFloat64Noop(1.1)}, str: `{"f1":{"_t":"float64Noop","_v":1.1}}`, expObj: DS2{F1: float64Noop(1.1)}},

	{obj: DS2{F1: addrOfStringNoop("hello")}, str: `{"f1":{"_t":"stringNoop","_v":"hello"}}`, expObj: DS2{F1: stringNoop("hello")}},
	{obj: DS2{F1: addrOfBoolNoop(true)}, str: `{"f1":{"_t":"boolNoop","_v":true}}`, expObj: DS2{F1: boolNoop(true)}},

	// address of primitive values stored in interface with >0 methods where only the address
	// of the value satisfies the interface.
	//
	// Unmarshaling the JSON into the object causes the decoder to check to see if the JSON objects
	// can be stored in DS7.F1, finding out that they cannot due to all the types implementing
	// DS7.F1 by-address. Thus the decoder will then check to see if the value can be assigned to
	// DS7.F1 by-address, which will work.
	{obj: DS7{F1: addrOfUintNoop(1)}, str: `{"f1":{"_t":"uintNoop","_v":1}}`},
	{obj: DS7{F1: addrOfUint8Noop(1)}, str: `{"f1":{"_t":"uint8Noop","_v":1}}`},
	{obj: DS7{F1: addrOfUint16Noop(1)}, str: `{"f1":{"_t":"uint16Noop","_v":1}}`},
	{obj: DS7{F1: addrOfUint32Noop(1)}, str: `{"f1":{"_t":"uint32Noop","_v":1}}`},
	{obj: DS7{F1: addrOfUint64Noop(1)}, str: `{"f1":{"_t":"uint64Noop","_v":1}}`},
	{obj: DS7{F1: addrOfUintptrNoop(1)}, str: `{"f1":{"_t":"uintptrNoop","_v":1}}`},

	{obj: DS7{F1: addrOfIntNoop(1)}, str: `{"f1":{"_t":"intNoop","_v":1}}`},
	{obj: DS7{F1: addrOfInt8Noop(1)}, str: `{"f1":{"_t":"int8Noop","_v":1}}`},
	{obj: DS7{F1: addrOfInt16Noop(1)}, str: `{"f1":{"_t":"int16Noop","_v":1}}`},
	{obj: DS7{F1: addrOfInt32Noop(1)}, str: `{"f1":{"_t":"int32Noop","_v":1}}`},
	{obj: DS7{F1: addrOfInt64Noop(1)}, str: `{"f1":{"_t":"int64Noop","_v":1}}`},

	{obj: DS7{F1: addrOfFloat32Noop(-1.0)}, str: `{"f1":{"_t":"float32Noop","_v":-1}}`},
	{obj: DS7{F1: addrOfFloat64Noop(1.0)}, str: `{"f1":{"_t":"float64Noop","_v":1}}`},
	{obj: DS7{F1: addrOfFloat32Noop(-1.1)}, str: `{"f1":{"_t":"float32Noop","_v":-1.1}}`},
	{obj: DS7{F1: addrOfFloat64Noop(1.1)}, str: `{"f1":{"_t":"float64Noop","_v":1.1}}`},

	{obj: DS7{F1: addrOfStringNoop("hello")}, str: `{"f1":{"_t":"stringNoop","_v":"hello"}}`},
	{obj: DS7{F1: addrOfBoolNoop(true)}, str: `{"f1":{"_t":"boolNoop","_v":true}}`},

	// struct value stored in interface with 0 methods
	{obj: DS1{F1: DS3{F1: "hello"}}, str: `{"f1":{"_t":"DS3","f1":"hello"}}`},
	{obj: DS1{F1: DS3{F1: "hello"}}, str: `{"_t":"DS1","f1":{"_t":"DS3","f1":"hello"}}`, mode: json.DiscriminatorEncodeTypeNameRootValue},

	{obj: DS1{F1: DS4{F1: "hello", F2: int(1)}}, str: `{"f1":{"_t":"DS4","f1":"hello","f2":{"_t":"int","_v":1}}}`},
	{obj: DS1{F1: DS4{F1: "hello", F2: DS3{F1: "world"}}}, str: `{"f1":{"_t":"DS4","f1":"hello","f2":{"_t":"DS3","f1":"world"}}}`},

	// struct value stored in interface with >0 methods
	{obj: DS2{F1: DS3Noop1{F1: "hello"}}, str: `{"f1":{"_t":"DS3Noop1","f1":"hello"}}`},
	{obj: DS2{F1: DS4Noop1{F1: "hello", F2: int(1)}}, str: `{"f1":{"_t":"DS4Noop1","f1":"hello","f2":{"_t":"int","_v":1}}}`},
	{obj: DS2{F1: DS5Noop1{F1: "hello", F2: DS3Noop1{F1: "world"}}}, str: `{"f1":{"_t":"DS5Noop1","f1":"hello","f2":{"_t":"DS3Noop1","f1":"world"}}}`},

	// address of struct value stored in interface with 0 methods
	{obj: DS1{F1: &DS3{F1: "hello"}}, str: `{"f1":{"_t":"DS3","f1":"hello"}}`, expObj: DS1{F1: DS3{F1: "hello"}}},
	{obj: DS1{F1: DS4{F1: "hello", F2: &DS3{F1: "world"}}}, str: `{"f1":{"_t":"DS4","f1":"hello","f2":{"_t":"DS3","f1":"world"}}}`, expObj: DS1{F1: DS4{F1: "hello", F2: DS3{F1: "world"}}}},

	// address of struct value stored in interface with >0 methods
	{obj: DS2{F1: DS3Noop1{F1: "hello"}}, str: `{"f1":{"_t":"DS3Noop1","f1":"hello"}}`},
	{obj: DS2{F1: DS4Noop1{F1: "hello", F2: int(1)}}, str: `{"f1":{"_t":"DS4Noop1","f1":"hello","f2":{"_t":"int","_v":1}}}`},
	{obj: DS2{F1: DS5Noop1{F1: "hello", F2: DS3Noop1{F1: "world"}}}, str: `{"f1":{"_t":"DS5Noop1","f1":"hello","f2":{"_t":"DS3Noop1","f1":"world"}}}`},

	// slices stored in interface with 0 methods
	{obj: DS1{F1: []int{1, 2, 3}}, str: `{"f1":{"_t":"[]int","_v":[1,2,3]}}`},
	{obj: DS1{F1: []*int{addrOfInt(1), addrOfInt(2), addrOfInt(3)}}, str: `{"f1":{"_t":"[]*int","_v":[1,2,3]}}`},
	{obj: DS1{F1: []interface{}{1, 2, 3}}, str: `{"f1":{"_t":"[]interface {}","_v":[{"_t":"int","_v":1},{"_t":"int","_v":2},{"_t":"int","_v":3}]}}`},

	// slices stored in interface with >0 methods
	{obj: DS2{F1: sliceIntNoop{1, 2, 3}}, str: `{"f1":{"_t":"sliceIntNoop","_v":[1,2,3]}}`},

	// address of slices stored in interface with >0 methods where only the
	// address of the value satisfies the interface
	{obj: DS7{F1: addrOfSliceIntNoop(sliceIntNoop{1, 2, 3})}, str: `{"f1":{"_t":"sliceIntNoop","_v":[1,2,3]}}`},

	// arrays stored in interfaces with 0 methods
	{obj: DS1{F1: [2]int{1, 2}}, str: `{"f1":{"_t":"[2]int","_v":[1,2]}}`},
	{obj: DS1{F1: [3]int{1, 2}}, str: `{"f1":{"_t":"[3]int","_v":[1,2,0]}}`},
	{obj: DS1{F1: [2]*int{addrOfInt(1), addrOfInt(2)}}, str: `{"f1":{"_t":"[2]*int","_v":[1,2]}}`},

	// arrays stored in interface with >0 methods
	{obj: DS2{F1: arrayOfTwoIntsNoop{1, 2}}, str: `{"f1":{"_t":"arrayOfTwoIntsNoop","_v":[1,2]}}`},

	// address of arrays stored in interface with >0 methods where only the
	// address of the value satisfies the interface
	{obj: DS7{F1: addrOfArrayOfTwoIntsNoop(arrayOfTwoIntsNoop{1, 2})}, str: `{"f1":{"_t":"arrayOfTwoIntsNoop","_v":[1,2]}}`},

	// maps stored in interface with 0 methods
	{obj: DS1{F1: map[string]int{"1": 1, "2": 2, "3": 3}}, str: `{"f1":{"_t":"map[string]int","1":1,"2":2,"3":3}}`},

	{obj: DS1{F1: map[string]*int{"1": addrOfInt(1), "2": addrOfInt(2), "3": addrOfInt(3)}}, str: `{"f1":{"_t":"map[string]*int","1":1,"2":2,"3":3}}`},
	{obj: DS1{F1: map[string]interface{}{"1": 1, "2": 2, "3": 3}}, str: `{"f1":{"_t":"map[string]interface {}","1":{"_t":"int","_v":1},"2":{"_t":"int","_v":2},"3":{"_t":"int","_v":3}}}`},
	// assert interface{} works during decode as well
	{obj: DS1{F1: map[string]interface{}{"1": 1, "2": 2, "3": 3}}, str: `{"f1":{"_t":"map[string]interface{}","1":{"_t":"int","_v":1},"2":{"_t":"int","_v":2},"3":{"_t":"int","_v":3}}}`, expStr: `{"f1":{"_t":"map[string]interface {}","1":{"_t":"int","_v":1},"2":{"_t":"int","_v":2},"3":{"_t":"int","_v":3}}}`},

	// maps stored in interface with >0 methods
	{obj: DS2{F1: mapStringIntNoop{"1": 1, "2": 2, "3": 3}}, str: `{"f1":{"_t":"mapStringIntNoop","1":1,"2":2,"3":3}}`},

	// address of maps stored in interface with >0 methods where only the
	// address of the value satisfies the interface
	{obj: DS7{F1: addrOfMapStringIntNoop(mapStringIntNoop{"1": 1, "2": 2, "3": 3})}, str: `{"f1":{"_t":"mapStringIntNoop","1":1,"2":2,"3":3}}`},

	// unsupported types
	{obj: DS1{F1: complex64(1.0)}, str: `{"f1":{"_t":"complex64","_v":1.0}}`, expEncErr: "json: unsupported type: complex64", expDecErr: "json: unsupported discriminator type: complex64"},
	{obj: DS1{F1: complex128(1.0)}, str: `{"f1":{"_t":"complex128","_v":1.0}}`, expEncErr: "json: unsupported type: complex128", expDecErr: "json: unsupported discriminator type: complex128"},
	{obj: DS1{F1: make(chan struct{})}, str: `{"f1":{"_t":"chan struct {}","_v":null}}`, expEncErr: "json: unsupported value: invalid kind: chan", expDecErr: "json: invalid discriminator type: chan struct {}"},
	{obj: DS1{F1: func(string) {}}, str: `{"f1":{"_t":"func(string)","_v":null}}`, expEncErr: "json: unsupported value: invalid kind: func", expDecErr: "json: invalid discriminator type: func(string)"},

	// discriminator type not a string
	{obj: DS1{}, str: `{"f1":{"_t":0,"_v":1}}`, expStr: `{"f1":null}`, expDecErr: "json: discriminator type at offset 12 is not string"},

	// discriminator not used for non-interface field
	{obj: DS8{F1: DS3{F1: "hello"}}, str: `{"f1":{"f1":"hello"}}`},
	{obj: DS8{F1: DS3{F1: "hello"}}, str: `{"_t":"DS8","f1":{"f1":"hello"}}`, mode: json.DiscriminatorEncodeTypeNameRootValue},
	{obj: DS8{F1: DS3{F1: "hello"}}, str: `{"_t":"DS8","f1":{"_t":"DS3","f1":"hello"}}`, mode: json.DiscriminatorEncodeTypeNameRootValue | json.DiscriminatorEncodeTypeNameAllObjects},
	{obj: DS8{F1: DS3{F1: "hello"}}, str: `{"_t":"DS8","f1":{"f1":"hello"}}`, mode: json.DiscriminatorEncodeTypeNameRootValue | json.DiscriminatorEncodeTypeNameAllObjects, discriminatorFunc: customNameWithFilter},

	// discriminator with full type path
	{obj: uint(1), str: `{"_t":"uint","_v":1}`, mode: json.DiscriminatorEncodeTypeNameRootValue, discriminatorFunc: json.FullName},
	{obj: DS2{F1: DS3Noop1{F1: "hello"}}, str: `{"f1":{"_t":"github.com/vmware/govmomi/vim25/json_test.DS3Noop1","f1":"hello"}}`, discriminatorFunc: json.FullName},
}

func discriminatorToTypeFn(discriminator string) (reflect.Type, bool) {
	switch discriminator {
	case "DS1":
		return reflect.TypeOf(DS1{}), true
	case "DS2":
		return reflect.TypeOf(DS2{}), true
	case "DS3":
		return reflect.TypeOf(DS3{}), true
	case "DS3Noop1", "github.com/vmware/govmomi/vim25/json_test.DS3Noop1":
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
	case "DS8":
		return reflect.TypeOf(DS8{}), true
	case "DS9":
		return reflect.TypeOf(DS9{}), true
	case "DS10":
		return reflect.TypeOf(DS10{}), true
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
	case "mapStringIntNoop":
		return reflect.TypeOf(mapStringIntNoop{}), true
	case "sliceIntNoop":
		return reflect.TypeOf(sliceIntNoop{}), true
	case "arrayOfTwoIntsNoop":
		return reflect.TypeOf(arrayOfTwoIntsNoop{}), true
	case "HexInt64":
		return reflect.TypeOf(HexInt64(0)), true
	default:
		return nil, false
	}
}

func TestDiscriminator(t *testing.T) {

	// Initialize the test case discriminator options to some defaults
	// as long as the discriminator is not disabled with dd=true.
	for i := range discriminatorTests {
		tc := &discriminatorTests[i]
		if !tc.dd {
			if tc.tf == "" {
				tc.tf = "_t"
			}
			if tc.vf == "" {
				tc.vf = "_v"
			}
		}
	}

	t.Run("Encode", testDiscriminatorEncode)
	t.Run("Decode", testDiscriminatorDecode)
}

func testDiscriminatorEncode(t *testing.T) {
	for _, tc := range discriminatorTests {
		tc := tc // capture the loop variable
		t.Run("", func(t *testing.T) {
			ee := tc.expEncErr
			var w bytes.Buffer

			enc := json.NewEncoder(&w)
			enc.SetDiscriminator(tc.tf, tc.vf, tc.mode)
			enc.SetTypeToDiscriminatorFunc(tc.discriminatorFunc)

			if err := enc.Encode(tc.obj); err != nil {
				if ee != err.Error() {
					t.Errorf("expected error mismatch: e=%v, a=%v", ee, err)
				} else if ee == "" {
					t.Errorf("unexpected error: %v", err)
				}
			} else if ee != "" {
				t.Errorf("expected error did not occur: %v", ee)
			} else {
				a, e := w.String(), tc.str
				if tc.expStr != "" {
					e = tc.expStr
				}
				if a != e+"\n" {
					t.Errorf("mismatch: e=%s, a=%s", e, a)
				}
			}
		})
	}
}

func testDiscriminatorDecode(t *testing.T) {
	for _, tc := range discriminatorTests {
		tc := tc // capture the loop variable
		t.Run("", func(t *testing.T) {
			ee := tc.expDecErr
			dec := json.NewDecoder(strings.NewReader(tc.str))
			dec.SetDiscriminator(tc.tf, tc.vf, discriminatorToTypeFn)

			var (
				err error
				obj interface{}
			)

			if tc.obj == nil || tc.mode&json.DiscriminatorEncodeTypeNameRootValue != 0 {
				err = dec.Decode(&obj)
			} else {
				switch reflect.TypeOf(tc.obj).Name() {
				case "DS1":
					var o DS1
					err = dec.Decode(&o)
					obj = o
				case "DS2":
					var o DS2
					err = dec.Decode(&o)
					obj = o
				case "DS3":
					var o DS3
					err = dec.Decode(&o)
					obj = o
				case "DS4":
					var o DS4
					err = dec.Decode(&o)
					obj = o
				case "DS5":
					var o DS5
					err = dec.Decode(&o)
					obj = o
				case "DS6":
					var o DS6
					err = dec.Decode(&o)
					obj = o
				case "DS7":
					var o DS7
					err = dec.Decode(&o)
					obj = o
				case "DS8":
					var o DS8
					err = dec.Decode(&o)
					obj = o
				case "DS9":
					var o DS9
					err = dec.Decode(&o)
					obj = o
				case "DS10":
					var o DS10
					o.DS9 = &DS9{}
					err = dec.Decode(&o)
					obj = o
				default:
					err = dec.Decode(&obj)
				}
			}

			if err != nil {
				if ee != err.Error() {
					t.Errorf("expected error mismatch: e=%v, a=%v", ee, err)
				} else if ee == "" {
					t.Errorf("unexpected error: %v", err)
				}
			} else if ee != "" {
				t.Errorf("expected error did not occur: %v", ee)
			} else {
				a, e := obj, tc.obj
				if tc.expObj != nil {
					e = tc.expObj
				}
				assertEqual(t, a, e)
			}
		})
	}
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

type mapStringIntNoop map[string]int

func (v mapStringIntNoop) noop1()  {}
func (v *mapStringIntNoop) noop2() {}

type sliceIntNoop []int

func (v sliceIntNoop) noop1()  {}
func (v *sliceIntNoop) noop2() {}

type arrayOfTwoIntsNoop [2]int

func (v arrayOfTwoIntsNoop) noop1()  {}
func (v *arrayOfTwoIntsNoop) noop2() {}

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

func addrOfMapStringIntNoop(v mapStringIntNoop) *mapStringIntNoop {
	return &v
}
func addrOfSliceIntNoop(v sliceIntNoop) *sliceIntNoop {
	return &v
}
func addrOfArrayOfTwoIntsNoop(v arrayOfTwoIntsNoop) *arrayOfTwoIntsNoop {
	return &v
}

func assertEqual(t *testing.T, a, e interface{}) {
	if !reflect.DeepEqual(a, e) {
		t.Fatalf("Actual and expected values differ.\nactual: '%#v'\nexpected: '%#v'\n", a, e)
	}
}
