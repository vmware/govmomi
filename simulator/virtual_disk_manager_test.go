// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vmdk"
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

func TestVirtualDiskEncryptDuringCopy(t *testing.T) {
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

	// Helper function to read VMDK descriptor from disk
	readVMDKDescriptor := func(dsPath string) (*vmdk.Descriptor, error) {
		// Convert datastore path to filesystem path
		// The Model creates datastores under {tempdir}/{dc_name}-{ds_name}/
		// For ESX model, dc_name is "ha-datacenter"
		dsPath = strings.TrimPrefix(dsPath, "[LocalDS_0] ")

		// Try both possible paths (ESX uses ha-datacenter)
		paths := []string{
			filepath.Join(m.dir, "ha-datacenter-LocalDS_0", dsPath),
			filepath.Join(m.dir, "DC0-LocalDS_0", dsPath),
		}

		var lastErr error
		for _, fullPath := range paths {
			f, err := os.Open(fullPath)
			if err != nil {
				lastErr = err
				continue
			}
			defer f.Close()
			return vmdk.ParseDescriptor(f)
		}

		return nil, lastErr
	}

	// Create source disk without encryption
	sourceName := "[LocalDS_0] disks/source_unencrypted.vmdk"
	destName := "[LocalDS_0] disks/dest_encrypted.vmdk"

	// Create the directory
	err = fm.MakeDirectory(ctx, path.Dir(sourceName), nil, true)
	if err != nil {
		t.Fatal(err)
	}

	// Create source disk
	sourceSpec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: 1024 * 1024,
	}

	task, err := dm.CreateVirtualDisk(ctx, sourceName, nil, sourceSpec)
	if err != nil {
		t.Fatal(err)
	}
	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	const (
		keyID      = "test-key-123"
		providerID = "test-provider"
	)

	// Copy disk with encryption
	destSpec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: 1024 * 1024,
		Crypto: &types.CryptoSpecEncrypt{
			CryptoKeyId: types.CryptoKeyId{
				KeyId: keyID,
				ProviderId: &types.KeyProviderId{
					Id: providerID,
				},
			},
		},
	}

	task, err = dm.CopyVirtualDisk(ctx, sourceName, nil, destName, nil, destSpec, false)
	if err != nil {
		t.Fatal(err)
	}
	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the destination disk was created with encryption
	// Read the VMDK descriptor to check for crypto keys
	desc, err := readVMDKDescriptor(destName)
	if err != nil {
		t.Fatalf("Failed to read VMDK descriptor: %v", err)
	}

	// Check that crypto keys were set in the descriptor
	assert.NotNil(t, desc.EncryptionKeys)
	assert.Equal(t, crypto.KeyLocatorTypeList, desc.EncryptionKeys.Type)
	assert.Len(t, desc.EncryptionKeys.List, 1)
	assert.Equal(t, crypto.KeyLocatorTypePair, desc.EncryptionKeys.List[0].Type)
	assert.NotNil(t, desc.EncryptionKeys.List[0].Pair)
	assert.NotNil(t, desc.EncryptionKeys.List[0].Pair.Locker)
	assert.Equal(t, crypto.KeyLocatorTypeFQID, desc.EncryptionKeys.List[0].Pair.Locker.Type)
	assert.NotNil(t, desc.EncryptionKeys.List[0].Pair.Locker.Indirect)
	assert.Equal(t, crypto.KeyLocatorTypeFQID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.Type)
	assert.Equal(t, keyID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.FQID.KeyID)
	assert.Equal(t, providerID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.FQID.KeyServerID)

	// Test copy with different encryption operations
	testCases := []struct {
		name       string
		destName   string
		cryptoSpec types.BaseCryptoSpec
		keyID      string
		providerID string
	}{
		{
			name:     "deep recrypt",
			destName: "[LocalDS_0] disks/dest_deep_recrypt.vmdk",
			cryptoSpec: &types.CryptoSpecDeepRecrypt{
				NewKeyId: types.CryptoKeyId{
					KeyId: "deep-recrypt-key",
					ProviderId: &types.KeyProviderId{
						Id: "deep-provider",
					},
				},
			},
			keyID:      "deep-recrypt-key",
			providerID: "deep-provider",
		},
		{
			name:     "shallow recrypt",
			destName: "[LocalDS_0] disks/dest_shallow_recrypt.vmdk",
			cryptoSpec: &types.CryptoSpecShallowRecrypt{
				NewKeyId: types.CryptoKeyId{
					KeyId: "shallow-recrypt-key",
					ProviderId: &types.KeyProviderId{
						Id: "shallow-provider",
					},
				},
			},
			keyID:      "shallow-recrypt-key",
			providerID: "shallow-provider",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			spec := &types.FileBackedVirtualDiskSpec{
				VirtualDiskSpec: types.VirtualDiskSpec{
					AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
					DiskType:    string(types.VirtualDiskTypeThin),
				},
				CapacityKb: 1024 * 1024,
				Crypto:     tc.cryptoSpec,
			}

			task, err := dm.CopyVirtualDisk(ctx, sourceName, nil, tc.destName, nil, spec, false)
			if err != nil {
				t.Fatal(err)
			}
			err = task.Wait(ctx)
			if err != nil {
				t.Fatal(err)
			}

			// Verify encryption metadata by reading VMDK descriptor
			desc, err := readVMDKDescriptor(tc.destName)
			if err != nil {
				t.Fatalf("Failed to read VMDK descriptor: %v", err)
			}

			// Verify encryption metadata in the disk
			assert.NotNil(t, desc.EncryptionKeys)
			assert.Equal(t, crypto.KeyLocatorTypeList, desc.EncryptionKeys.Type)
			assert.Len(t, desc.EncryptionKeys.List, 1)
			assert.Equal(t, crypto.KeyLocatorTypePair, desc.EncryptionKeys.List[0].Type)
			assert.NotNil(t, desc.EncryptionKeys.List[0].Pair)
			assert.NotNil(t, desc.EncryptionKeys.List[0].Pair.Locker)
			assert.Equal(t, crypto.KeyLocatorTypeFQID, desc.EncryptionKeys.List[0].Pair.Locker.Type)
			assert.NotNil(t, desc.EncryptionKeys.List[0].Pair.Locker.Indirect)
			assert.Equal(t, crypto.KeyLocatorTypeFQID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.Type)
			assert.Equal(t, tc.keyID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.FQID.KeyID)
			assert.Equal(t, tc.providerID, desc.EncryptionKeys.List[0].Pair.Locker.Indirect.FQID.KeyServerID)
		})
	}

	// Clean up
	for _, name := range []string{sourceName, destName} {
		task, err := dm.DeleteVirtualDisk(ctx, name, nil)
		if err != nil {
			continue
		}
		_ = task.Wait(ctx)
	}
	for _, tc := range testCases {
		task, err := dm.DeleteVirtualDisk(ctx, tc.destName, nil)
		if err != nil {
			continue
		}
		_ = task.Wait(ctx)
	}
}
