// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"strings"
	"testing"
)

func TestReadManifest(t *testing.T) {
	mf := `SHA1(ttylinux-pc_i486-16.1.ovf)= 344a536fab6782622b6beb923798e84134bd4cbd
SHA1(ttylinux-pc_i486-16.1-disk1.vmdk)= ed64564a37366bfe1c93af80e2ead0cbd398c3d3`

	sums, err := ReadManifest(strings.NewReader(mf))
	if err != nil {
		t.Error(err)
	}

	if len(sums) != 2 {
		t.Error(err)
	}
}
