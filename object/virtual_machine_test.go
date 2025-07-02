// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"testing"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

// VirtualMachine should implement the Reference interface.
var _ Reference = VirtualMachine{}

// pretty.Printf generated
var snapshot = &types.VirtualMachineSnapshotInfo{
	DynamicData:     types.DynamicData{},
	CurrentSnapshot: &types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-11"},
	RootSnapshotList: []types.VirtualMachineSnapshotTree{
		{
			DynamicData:    types.DynamicData{},
			Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-1"},
			Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
			Name:           "root",
			Description:    "",
			Id:             1,
			CreateTime:     time.Now(),
			State:          "poweredOn",
			Quiesced:       false,
			BackupManifest: "",
			ChildSnapshotList: []types.VirtualMachineSnapshotTree{
				{
					DynamicData:       types.DynamicData{},
					Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-2"},
					Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:              "child",
					Description:       "",
					Id:                2,
					CreateTime:        time.Now(),
					State:             "poweredOn",
					Quiesced:          false,
					BackupManifest:    "",
					ChildSnapshotList: nil,
					ReplaySupported:   types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-3"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "child",
					Description:    "",
					Id:             3,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:    types.DynamicData{},
							Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-9"},
							Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:           "grandkid",
							Description:    "",
							Id:             9,
							CreateTime:     time.Now(),
							State:          "poweredOn",
							Quiesced:       false,
							BackupManifest: "",
							ChildSnapshotList: []types.VirtualMachineSnapshotTree{
								{
									DynamicData:       types.DynamicData{},
									Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-10"},
									Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
									Name:              "great",
									Description:       "",
									Id:                10,
									CreateTime:        time.Now(),
									State:             "poweredOn",
									Quiesced:          false,
									BackupManifest:    "",
									ChildSnapshotList: nil,
									ReplaySupported:   types.NewBool(false),
								},
							},
							ReplaySupported: types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-5"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "voodoo",
					Description:    "",
					Id:             5,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:       types.DynamicData{},
							Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-11"},
							Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:              "child",
							Description:       "",
							Id:                11,
							CreateTime:        time.Now(),
							State:             "poweredOn",
							Quiesced:          false,
							BackupManifest:    "",
							ChildSnapshotList: nil,
							ReplaySupported:   types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-6"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "better",
					Description:    "",
					Id:             6,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:    types.DynamicData{},
							Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-7"},
							Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:           "best",
							Description:    "",
							Id:             7,
							CreateTime:     time.Now(),
							State:          "poweredOn",
							Quiesced:       false,
							BackupManifest: "",
							ChildSnapshotList: []types.VirtualMachineSnapshotTree{
								{
									DynamicData:       types.DynamicData{},
									Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-8"},
									Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
									Name:              "betterer",
									Description:       "",
									Id:                8,
									CreateTime:        time.Now(),
									State:             "poweredOn",
									Quiesced:          false,
									BackupManifest:    "",
									ChildSnapshotList: nil,
									ReplaySupported:   types.NewBool(false),
								},
							},
							ReplaySupported: types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
			},
			ReplaySupported: types.NewBool(false),
		},
	},
}

func TestVirtualMachineSnapshotMap(t *testing.T) {
	m := make(snapshotMap)
	m.add("", snapshot.RootSnapshotList)

	tests := []struct {
		name   string
		expect int
	}{
		{"enoent", 0},
		{"root", 1},
		{"child", 3},
		{"root/child", 2},
		{"root/voodoo/child", 1},
		{"2-snapshot-6", 1},
	}

	for _, test := range tests {
		s := m[test.name]

		if len(s) != test.expect {
			t.Errorf("%s: %d != %d", test.name, len(s), test.expect)
		}
	}
}

func TestDiskFileOperation(t *testing.T) {
	backing := &types.VirtualDiskFlatVer2BackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: "[datastore1] data/disk1.vmdk",
		},
		Parent: nil,
	}

	parent := &types.VirtualDiskFlatVer2BackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: "[datastore1] data/parent.vmdk",
		},
	}

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Backing: backing,
		},
	}

	op := types.VirtualDeviceConfigSpecOperationAdd
	fop := types.VirtualDeviceConfigSpecFileOperationCreate

	res := diskFileOperation(op, fop, disk)
	if res != "" {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInKB = 1
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInKB = 0
	disk.CapacityInBytes = 1
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInBytes = 0
	backing.Parent = parent
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}
}
