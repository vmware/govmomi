// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk_test

import (
	"io"
	"os"
	"testing"

	"github.com/vmware/govmomi/units"
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

	w := io.Discard
	if testing.Verbose() {
		w = os.Stderr
	}

	if err := di.Descriptor.Write(w); err != nil {
		t.Error(err)
	}

	cap := di.Descriptor.Capacity()
	hdr := di.Capacity
	if cap != hdr {
		t.Errorf("descriptor capacity %d != header capacity %d", cap, hdr)
	}

	scap := units.ByteSize(di.Capacity).String()
	if scap != "30.0MB" {
		t.Errorf("capacity=%s", scap)
	}

	_, err = di.OVF()
	if err != nil {
		t.Fatal(err)
	}

	di, err = vmdk.Stat("../govc/test/images/ttylinux-pc_i486-16.1.ovf")
	if err != vmdk.ErrInvalidFormat {
		t.Fatalf("error=%s", err)
	}
}

func TestDiskInvalid(t *testing.T) {
	_, err := vmdk.Stat("import_test.go")
	if err != vmdk.ErrInvalidFormat {
		t.Errorf("expected ErrInvalidFormat: %s", err)
	}
}
