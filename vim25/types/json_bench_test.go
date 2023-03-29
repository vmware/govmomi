/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

				var obj interface{}
				if err := dec.Decode(&obj); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
