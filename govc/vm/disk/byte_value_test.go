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

package disk

import "testing"

func TestByteValue(t *testing.T) {
	var tests = []struct {
		In     string
		OutStr string
		OutInt int64
	}{
		{
			In:     "345",
			OutStr: "345B",
			OutInt: 345,
		},
		{
			In:     "345K",
			OutStr: "345KiB",
			OutInt: 345 * KiB,
		},
		{
			In:     "345kib",
			OutStr: "345KiB",
			OutInt: 345 * KiB,
		},
		{
			In:     "345KiB",
			OutStr: "345KiB",
			OutInt: 345 * KiB,
		},
		{
			In:     "345M",
			OutStr: "345MiB",
			OutInt: 345 * MiB,
		},
		{
			In:     "345G",
			OutStr: "345GiB",
			OutInt: 345 * GiB,
		},
		{
			In:     "345T",
			OutStr: "345TiB",
			OutInt: 345 * TiB,
		},
		{
			In:     "345P",
			OutStr: "345PiB",
			OutInt: 345 * PiB,
		},
	}

	v := ByteValue{}

	for _, test := range tests {
		err := v.Set(test.In)
		if err != nil {
			t.Errorf("Error: %s", err)
			continue
		}

		if v.Bytes != test.OutInt {
			t.Errorf("Int: %d", v.Bytes)
			continue
		}

		if v.String() != test.OutStr {
			t.Errorf("String: %s", test.OutStr)
			continue
		}
	}
}
