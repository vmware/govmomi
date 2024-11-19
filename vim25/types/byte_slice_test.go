/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
	"testing"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestByteSlice(t *testing.T) {
	in := &ArrayOfByte{
		Byte: []byte("xmhell"),
	}

	res, err := xml.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	var out ArrayOfByte
	if err := xml.Unmarshal(res, &out); err != nil {
		t.Logf("%s", string(res))
		t.Fatal(err)
	}

	if !bytes.Equal(in.Byte, out.Byte) {
		t.Errorf("out=%#v", out.Byte)
	}
}

func TestSignedByteSlice(t *testing.T) {
	//  int8: <byte>4</byte><byte>-80</byte><byte>-79</byte><byte>-78</byte>
	// uint8: <byte>4</byte><byte>176</byte><byte>177</byte><byte>178</byte>
	in := &ArrayOfByte{
		Byte: []uint8{0x4, 0xb0, 0xb1, 0xb2},
	}

	res, err := xml.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	var out struct {
		Byte []int8 `xml:"byte,omitempty" json:"_value"`
	}

	// ByteSlice.MarshalXML must output signed byte, otherwise this fails with:
	// strconv.ParseInt: parsing "176": value out of range
	if err := xml.Unmarshal(res, &out); err != nil {
		t.Logf("%s", string(res))
		t.Fatal(err)
	}

	for i := range in.Byte {
		if in.Byte[i] != byte(out.Byte[i]) {
			t.Errorf("out=%x", out.Byte[i])
		}
	}
}
