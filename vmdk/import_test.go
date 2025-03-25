// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk_test

import (
	"os"
	"testing"

	"github.com/vmware/govmomi/vmdk"
)

func TestDiskInfo(t *testing.T) {
	di, err := vmdk.Stat("../govc/test/images/ttylinux-pc_i486-16.1-disk1.vmdk")
	if err != nil {
		if os.IsNotExist(err) {
			t.SkipNow()
		}
		t.Fatal(err)
	}

	_, err = di.OVF()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDiskInvalid(t *testing.T) {
	_, err := vmdk.Stat("import_test.go")
	if err != vmdk.ErrInvalidFormat {
		t.Errorf("expected ErrInvalidFormat: %s", err)
	}
}
