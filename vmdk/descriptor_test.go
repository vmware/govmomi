// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/vmdk"
)

func TestDescriptor(t *testing.T) {
	desc := &vmdk.Descriptor{
		Version:  1,
		Encoding: "UTF-8",
		CID:      123,
		Type:     "vmfs",
		Extent: []vmdk.Extent{{
			Type:       "VMFS",
			Permission: "RW",
			Size:       1024,
			Info:       "test-flat.vmdk",
		}},
		DDB: map[string]string{
			"adapterType":      "lsilogic",
			"virtualHWVersion": "14",
		},
	}

	var buf bytes.Buffer

	err := desc.Write(&buf)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := vmdk.ParseDescriptor(&buf)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(desc, parsed) {
		t.Error("not equal")
	}
}
