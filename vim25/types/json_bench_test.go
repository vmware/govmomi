// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"bytes"
	"os"
	"testing"

	"github.com/vmware/govmomi/vim25/json"
)

// BenchmarkDecodeVirtualMachineConfigInfo illustrates the performance
// difference between decoding a large vminfo object when the type
// discriminator is the first property in the object versus the last
// property.
func BenchmarkDecodeVirtualMachineConfigInfo(b *testing.B) {

	testCases := []struct {
		name string
		path string
	}{
		{
			name: "vm info w type name first",
			path: "./testdata/vminfo.json",
		},
		{
			name: "vm info w type name last",
			path: "./testdata/vminfo-typename-at-end.json",
		},
	}

	for _, tc := range testCases {
		tc := tc // capture the range variable
		b.Run(tc.name, func(b *testing.B) {
			buf, err := os.ReadFile(tc.path)
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {

				dec := json.NewDecoder(bytes.NewReader(buf))
				dec.SetDiscriminator(
					"_typeName", "_value",
					json.DiscriminatorToTypeFunc(TypeFunc()),
				)

				var obj any
				if err := dec.Decode(&obj); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
