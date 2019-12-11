/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"testing"

	"github.com/google/uuid"
)

func TestEncodeDMI(t *testing.T) {
	id := uuid.Must(uuid.Parse("423e36c8-723d-2b24-8590-0eb1d9180aa2"))

	tests := []struct {
		f      func(uuid.UUID) string
		expect string
	}{
		{productSerial, "VMware-42 3e 36 c8 72 3d 2b 24-85 90 0e b1 d9 18 0a a2"},
		{productUUID, "C8363E42-3D72-242B-8590-0EB1D9180AA2"},
	}

	for _, test := range tests {
		val := test.f(id)
		if val != test.expect {
			t.Errorf("%q != %q", val, test.expect)
		}
	}
}
