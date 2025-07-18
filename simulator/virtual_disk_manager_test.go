// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"path"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

func TestVirtualDiskManager(t *testing.T) {
	ctx := context.Background()

	m := ESX()
	defer m.Remove()
	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	dm := object.NewVirtualDiskManager(c.Client)
	fm := object.NewFileManager(c.Client)

	spec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: 1024 * 1024,
	}

	name := "[LocalDS_0] disks/disk1.vmdk"

	for i, fail := range []bool{true, false, true} {
		task, err := dm.CreateVirtualDisk(ctx, name, nil, spec)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if fail {
			if err == nil {
				t.Error("expected error") // disk1 already exists
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}

		if i == 0 {
			err = fm.MakeDirectory(ctx, path.Dir(name), nil, true)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	qname := name
	for i, fail := range []bool{false, true, true} {
		if i == 1 {
			spec.CapacityKb = 0
		}

		if i == 2 {
			qname += "_missing_file"
		}
		task, err := dm.ExtendVirtualDisk(ctx, qname, nil, spec.CapacityKb*2, nil)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if fail {
			if err == nil {
				t.Error("expected error")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}

	qname = name
	for _, fail := range []bool{false, true} {
		id, err := dm.QueryVirtualDiskUuid(ctx, qname, nil)
		if fail {
			if err == nil {
				t.Error("expected error")
			}
		} else {
			if err != nil {
				t.Error(err)
			}

			_, err = uuid.Parse(id)
			if err != nil {
				t.Error(err)
			}
		}
		qname += "-enoent"
	}

	old := name
	name = strings.Replace(old, "disk1", "disk2", 1)

	for _, fail := range []bool{false, true} {
		task, err := dm.MoveVirtualDisk(ctx, old, nil, name, nil, false)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if fail {
			if err == nil {
				t.Error("expected error") // disk1 no longer exists
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}

	for _, fail := range []bool{false, true} {
		task, err := dm.CopyVirtualDisk(ctx, name, nil, old, nil, &types.VirtualDiskSpec{}, false)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if fail {
			if err == nil {
				t.Error("expected error") // disk1 exists again
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}

	for _, fail := range []bool{false, true} {
		task, err := dm.DeleteVirtualDisk(ctx, name, nil)
		if err != nil {
			t.Fatal(err)
		}

		err = task.Wait(ctx)
		if fail {
			if err == nil {
				t.Error("expected error") // disk2 no longer exists
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}
}
