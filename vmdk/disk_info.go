/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package vmdk

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualDiskCryptoKey struct {
	KeyID      string
	ProviderID string
}

type VirtualDiskInfo struct {
	CapacityInBytes int64
	DeviceKey       int32
	FileName        string
	Size            int64
	UniqueSize      int64
	CryptoKey       VirtualDiskCryptoKey
}

// GetVirtualDiskInfoByUUID returns information about a virtual disk identified
// by the provided UUID. This method is valid for the following backing types:
//
// - VirtualDiskFlatVer2BackingInfo
// - VirtualDiskSeSparseBackingInfo
// - VirtualDiskRawDiskMappingVer1BackingInfo
// - VirtualDiskSparseVer2BackingInfo
// - VirtualDiskRawDiskVer2BackingInfo
//
// These are the only backing types that have a UUID property for comparing the
// provided value.
func GetVirtualDiskInfoByUUID(
	ctx context.Context,
	client *vim25.Client,
	mo mo.VirtualMachine,
	fetchProperties bool,
	diskUUID string) (VirtualDiskInfo, error) {

	if diskUUID == "" {
		return VirtualDiskInfo{}, fmt.Errorf("diskUUID is empty")
	}

	switch {
	case fetchProperties,
		mo.Config == nil,
		mo.Config.Hardware.Device == nil,
		mo.LayoutEx == nil,
		mo.LayoutEx.Disk == nil,
		mo.LayoutEx.File == nil:

		if ctx == nil {
			return VirtualDiskInfo{}, fmt.Errorf("ctx is nil")
		}
		if client == nil {
			return VirtualDiskInfo{}, fmt.Errorf("client is nil")
		}

		obj := object.NewVirtualMachine(client, mo.Self)

		if err := obj.Properties(
			ctx,
			mo.Self,
			[]string{"config", "layoutEx"},
			&mo); err != nil {

			return VirtualDiskInfo{},
				fmt.Errorf("failed to retrieve properties: %w", err)
		}
	}

	// Find the disk by UUID by inspecting all of the disk backing types that
	// can have an associated UUID.
	var (
		disk      *types.VirtualDisk
		fileName  string
		cryptoKey *types.CryptoKeyId
	)
	for i := range mo.Config.Hardware.Device {
		switch tvd := mo.Config.Hardware.Device[i].(type) {
		case *types.VirtualDisk:
			switch tb := tvd.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				if tb.Uuid == diskUUID {
					disk = tvd
					fileName = tb.FileName
					cryptoKey = tb.KeyId
				}
			case *types.VirtualDiskSeSparseBackingInfo:
				if tb.Uuid == diskUUID {
					disk = tvd
					fileName = tb.FileName
					cryptoKey = tb.KeyId
				}
			case *types.VirtualDiskRawDiskMappingVer1BackingInfo:
				if tb.Uuid == diskUUID {
					disk = tvd
					fileName = tb.FileName
				}
			case *types.VirtualDiskSparseVer2BackingInfo:
				if tb.Uuid == diskUUID {
					disk = tvd
					fileName = tb.FileName
					cryptoKey = tb.KeyId
				}
			case *types.VirtualDiskRawDiskVer2BackingInfo:
				if tb.Uuid == diskUUID {
					disk = tvd
					fileName = tb.DescriptorFileName
				}
			}
		}
	}

	if disk == nil {
		return VirtualDiskInfo{},
			fmt.Errorf("disk not found with uuid %q", diskUUID)
	}

	// Build a lookup table for determining if file key belongs to this disk
	// chain.
	diskFileKeys := map[int32]struct{}{}
	for i := range mo.LayoutEx.Disk {
		if d := mo.LayoutEx.Disk[i]; d.Key == disk.Key {
			for j := range d.Chain {
				for k := range d.Chain[j].FileKey {
					diskFileKeys[d.Chain[j].FileKey[k]] = struct{}{}
				}
			}
		}
	}

	// Sum the disk's total size and unique size.
	var (
		size       int64
		uniqueSize int64
	)
	for i := range mo.LayoutEx.File {
		f := mo.LayoutEx.File[i]
		if _, ok := diskFileKeys[f.Key]; ok {
			size += f.Size
			uniqueSize += f.UniqueSize
		}
	}

	di := VirtualDiskInfo{
		CapacityInBytes: disk.CapacityInBytes,
		DeviceKey:       disk.Key,
		FileName:        fileName,
		Size:            size,
		UniqueSize:      uniqueSize,
	}

	if ck := cryptoKey; ck != nil {
		di.CryptoKey.KeyID = ck.KeyId
		if pid := ck.ProviderId; pid != nil {
			di.CryptoKey.ProviderID = pid.Id
		}
	}

	return di, nil
}
