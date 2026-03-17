// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/vim25/types"
)

func TestStoragegroupWithDiskConfigSpec(t *testing.T) {
	// storagegroupWithDisk.ovf: single VirtualSystem with vmw:StorageSection / vmw:StorageGroupSection,
	// one SCSI controller (buslogic), one file-backed disk (UnOS2.vmdk, vmdisk1).
	e := testEnvelope(t, "fixtures/external/ImportOvf2014/storagegroupWithDisk.ovf")
	cs, err := e.ToConfigSpec()
	require.NoError(t, err)
	require.NotEmpty(t, cs.DeviceChange)

	var controllers []types.BaseVirtualSCSIController
	var disks []*types.VirtualDisk
	for _, dc := range cs.DeviceChange {
		dev := dc.GetVirtualDeviceConfigSpec().Device
		if c, ok := dev.(types.BaseVirtualSCSIController); ok {
			controllers = append(controllers, c)
		}
		if d, ok := dev.(*types.VirtualDisk); ok {
			disks = append(disks, d)
		}
	}

	require.Len(t, controllers, 1, "OVF has one SCSI controller (buslogic)")
	_, ok := controllers[0].(*types.VirtualBusLogicController)
	assert.True(t, ok, "SCSI controller must be VirtualBusLogicController for rasd:ResourceSubType buslogic")

	require.Len(t, disks, 1, "OVF has one disk (vmdisk1, fileRef file1)")
	assert.Equal(t, int64(107373568), disks[0].CapacityInBytes,
		"DiskSection vmdisk1 ovf:capacity=107373568")
	db, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
	require.True(t, ok)
	assert.Equal(t, "UnOS2.vmdk", db.FileName,
		"file-backed disk backing FileName must be path.Base(References/File href)")
}

func TestUnmarshalExternalFixtures(t *testing.T) {
	dir := "fixtures/external"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Skipf("fixtures directory not found: %s", dir)
		return
	}

	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".ovf" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", dir, err)
	}

	for _, path := range files {
		t.Run(path, func(t *testing.T) {
			runFixtureTest(t, path, isNegativeFixture(path))
		})
	}
}
