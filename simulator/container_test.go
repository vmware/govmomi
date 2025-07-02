// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
