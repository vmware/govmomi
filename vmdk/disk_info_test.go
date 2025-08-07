// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk_test

import (
	"bytes"
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vmdk"
)

func TestGetVirtualDiskInfoByUUID(t *testing.T) {

	type testCase struct {
		name             string
		ctx              context.Context
		client           *vim25.Client
		mo               mo.VirtualMachine
		fetchProperties  bool
		diskUUID         string
		diskInfo         vmdk.VirtualDiskInfo
		excludeSnapshots bool
		err              string
	}

	t.Run("w cached properties", func(t *testing.T) {

		const (
			deviceKey     = 1000
			diskUUID      = "123"
			fileName      = "[datastore] path/to.vmdk"
			tenGiBInBytes = 10 * 1024 * 1024 * 1024
		)

		getDisk := func(backing types.BaseVirtualDeviceBackingInfo) *types.VirtualDisk {
			return &types.VirtualDisk{
				VirtualDevice: types.VirtualDevice{
					Key:     deviceKey,
					Backing: backing,
				},
				CapacityInBytes: tenGiBInBytes,
			}
		}

		getDiskInfo := func() vmdk.VirtualDiskInfo {
			return vmdk.VirtualDiskInfo{
				CapacityInBytes: tenGiBInBytes,
				DeviceKey:       deviceKey,
				FileName:        fileName,
				Size:            (1 * 1024 * 1024 * 1024) + 950,
				UniqueSize:      (5 * 1024 * 1024) + 100,
			}
		}

		getEncryptedDiskInfo := func(pid, kid string) vmdk.VirtualDiskInfo {
			return vmdk.VirtualDiskInfo{
				CapacityInBytes: tenGiBInBytes,
				DeviceKey:       deviceKey,
				FileName:        fileName,
				Size:            (1 * 1024 * 1024 * 1024) + 950,
				UniqueSize:      (5 * 1024 * 1024) + 100,
				CryptoKey: vmdk.VirtualDiskCryptoKey{
					KeyID:      kid,
					ProviderID: pid,
				},
			}
		}

		getLayoutEx := func() *types.VirtualMachineFileLayoutEx {
			return &types.VirtualMachineFileLayoutEx{
				Disk: []types.VirtualMachineFileLayoutExDiskLayout{
					{
						Key: 1000,
						Chain: []types.VirtualMachineFileLayoutExDiskUnit{
							{
								FileKey: []int32{
									4,
									5,
								},
							},
						},
					},
				},
				File: []types.VirtualMachineFileLayoutExFileInfo{
					{
						Key:        4,
						Size:       1 * 1024 * 1024 * 1024, // 1 GiB
						UniqueSize: 5 * 1024 * 1024,        // 500 MiB
					},
					{
						Key:        5,
						Size:       950,
						UniqueSize: 100,
					},
				},
			}
		}

		testCases := []testCase{
			{
				name: "diskUUID is empty",
				err:  "diskUUID is empty",
			},
			{
				name: "no matching disks",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{},
						},
					},
					LayoutEx: &types.VirtualMachineFileLayoutEx{
						File: []types.VirtualMachineFileLayoutExFileInfo{},
						Disk: []types.VirtualMachineFileLayoutExDiskLayout{},
					},
				},
				diskUUID: diskUUID,
				err:      "disk not found with uuid \"123\"",
			},
			{
				name: "one disk w VirtualDiskFlatVer2BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskFlatVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getDiskInfo(),
			},
			{
				name: "one encrypted disk w VirtualDiskFlatVer2BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskFlatVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
									KeyId: &types.CryptoKeyId{
										KeyId: "my-key-id",
										ProviderId: &types.KeyProviderId{
											Id: "my-provider-id",
										},
									},
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getEncryptedDiskInfo("my-provider-id", "my-key-id"),
			},
			{
				name: "one disk w VirtualDiskSeSparseBackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskSeSparseBackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getDiskInfo(),
			},
			{
				name: "one encrypted disk w VirtualDiskSeSparseBackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskSeSparseBackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
									KeyId: &types.CryptoKeyId{
										KeyId: "my-key-id",
										ProviderId: &types.KeyProviderId{
											Id: "my-provider-id",
										},
									},
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getEncryptedDiskInfo("my-provider-id", "my-key-id"),
			},
			{
				name: "one disk w VirtualDiskRawDiskMappingVer1BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskRawDiskMappingVer1BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getDiskInfo(),
			},
			{
				name: "one disk w VirtualDiskSparseVer2BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskSparseVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getDiskInfo(),
			},
			{
				name: "one encrypted disk w VirtualDiskSparseVer2BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskSparseVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
									KeyId: &types.CryptoKeyId{
										KeyId: "my-key-id",
										ProviderId: &types.KeyProviderId{
											Id: "my-provider-id",
										},
									},
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getEncryptedDiskInfo("my-provider-id", "my-key-id"),
			},
			{
				name: "one disk w VirtualDiskRawDiskVer2BackingInfo",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskRawDiskVer2BackingInfo{
									DescriptorFileName: fileName,
									Uuid:               diskUUID,
								}),
							},
						},
					},
					LayoutEx: getLayoutEx(),
				},
				diskUUID: diskUUID,
				diskInfo: getDiskInfo(),
			},
			{
				name: "one disk w multiple chain entries",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskFlatVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: &types.VirtualMachineFileLayoutEx{
						Disk: []types.VirtualMachineFileLayoutExDiskLayout{
							{
								Key: deviceKey,
								Chain: []types.VirtualMachineFileLayoutExDiskUnit{
									{
										FileKey: []int32{
											4,
											5,
										},
									},
									{
										FileKey: []int32{
											6,
											7,
										},
									},
									{
										FileKey: []int32{
											8,
										},
									},
								},
							},
						},
						File: []types.VirtualMachineFileLayoutExFileInfo{
							{
								Key:        4,
								Size:       1 * 1024 * 1024 * 1024, // 1 GiB
								UniqueSize: 5 * 1024 * 1024,        // 500 MiB
							},
							{
								Key:        5,
								Size:       950,
								UniqueSize: 100,
							},
							{
								Key:        6,
								Size:       500,
								UniqueSize: 100,
							},
							{
								Key:        7,
								Size:       500,
								UniqueSize: 200,
							},
							{
								Key:        8,
								Size:       1000,
								UniqueSize: 300,
							},
						},
					},
				},
				diskUUID: diskUUID,
				diskInfo: vmdk.VirtualDiskInfo{
					CapacityInBytes: tenGiBInBytes,
					DeviceKey:       deviceKey,
					FileName:        fileName,
					Size:            (1 * 1024 * 1024 * 1024) + 950 + 500 + 500 + 1000,
					UniqueSize:      (5 * 1024 * 1024) + 100 + 100 + 200 + 300,
				},
			},
			{
				name: "one disk w/o chain entries, it should return 0 for size and unique size",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskFlatVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: &types.VirtualMachineFileLayoutEx{
						Disk: []types.VirtualMachineFileLayoutExDiskLayout{
							{
								Key:   deviceKey,
								Chain: []types.VirtualMachineFileLayoutExDiskUnit{},
							},
						},
						File: []types.VirtualMachineFileLayoutExFileInfo{
							{
								Key:        1,
								Size:       100,
								UniqueSize: 100,
							},
						},
					},
				},
				diskUUID: diskUUID,
				diskInfo: vmdk.VirtualDiskInfo{
					CapacityInBytes: tenGiBInBytes,
					DeviceKey:       deviceKey,
					FileName:        fileName,
					Size:            0,
					UniqueSize:      0,
				},
			},
			{
				name: "one disk w multiple chain entries and includeSnapshots is false, it should exclude the snapshot delta disks, only the size offile keys (8, 9) should be included",
				mo: mo.VirtualMachine{
					Config: &types.VirtualMachineConfigInfo{
						Hardware: types.VirtualHardware{
							Device: []types.BaseVirtualDevice{
								getDisk(&types.VirtualDiskFlatVer2BackingInfo{
									VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
										FileName: fileName,
									},
									Uuid: diskUUID,
								}),
							},
						},
					},
					LayoutEx: &types.VirtualMachineFileLayoutEx{
						Disk: []types.VirtualMachineFileLayoutExDiskLayout{
							{
								Key: deviceKey,
								Chain: []types.VirtualMachineFileLayoutExDiskUnit{
									{
										FileKey: []int32{
											4,
											5,
										},
									},
									{
										FileKey: []int32{
											6,
											7,
										},
									},
									{
										FileKey: []int32{
											8,
											9,
										},
									},
								},
							},
						},
						File: []types.VirtualMachineFileLayoutExFileInfo{
							{
								Key:        4,
								Size:       1 * 1024 * 1024 * 1024, // 1 GiB
								UniqueSize: 5 * 1024 * 1024,        // 500 MiB
							},
							{
								Key:        5,
								Size:       950,
								UniqueSize: 100,
							},
							{
								Key:        6,
								Size:       500,
								UniqueSize: 100,
							},
							{
								Key:        7,
								Size:       500,
								UniqueSize: 200,
							},
							{
								Key:        8,
								Size:       1000,
								UniqueSize: 300,
							},
							{
								Key:        9,
								Size:       1000,
								UniqueSize: 300,
							},
						},
					},
				},
				excludeSnapshots: true,
				diskUUID:         diskUUID,
				diskInfo: vmdk.VirtualDiskInfo{
					CapacityInBytes: tenGiBInBytes,
					DeviceKey:       deviceKey,
					FileName:        fileName,
					Size:            1000 + 1000,
					UniqueSize:      300 + 300,
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var ctx context.Context
				dii, err := vmdk.GetVirtualDiskInfoByUUID(
					ctx, nil, tc.mo, false, tc.excludeSnapshots, tc.diskUUID)

				if tc.err != "" {
					assert.EqualError(t, err, tc.err)
				} else {
					assert.NoError(t, err)
				}

				assert.Equal(t, tc.diskInfo, dii)
			})
		}
	})

	t.Run("fetch properties", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pc := &propertyCollectorWithFault{}
			pc.Self = c.ServiceContent.PropertyCollector
			simulator.Map(ctx).Put(pc)

			finder := find.NewFinder(c, true)
			datacenter, err := finder.DefaultDatacenter(ctx)
			if err != nil {
				t.Fatalf("default datacenter not found: %s", err)
			}
			finder.SetDatacenter(datacenter)
			vmList, err := finder.VirtualMachineList(ctx, "*")
			if err != nil {
				t.Fatalf("failed to get vm list: %s", err)
			}
			if len(vmList) == 0 {
				t.Fatal("vmList == 0")
			}
			vm := vmList[0]

			var moVM mo.VirtualMachine
			if err := vm.Properties(
				ctx,
				vm.Reference(),
				[]string{"config", "layoutEx"},
				&moVM); err != nil {

				t.Fatal(err)
			}

			devs := object.VirtualDeviceList(moVM.Config.Hardware.Device)
			disks := devs.SelectByType(&types.VirtualDisk{})
			if len(disks) == 0 {
				t.Fatal("disks == 0")
			}

			var (
				diskUUID     string
				datastoreRef types.ManagedObjectReference
				disk         = disks[0].(*types.VirtualDisk)
				diskBacking  = disk.Backing
				diskInfo     = vmdk.VirtualDiskInfo{
					CapacityInBytes: disk.CapacityInBytes,
					DeviceKey:       disk.Key,
				}
			)

			switch tb := disk.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				diskUUID = tb.Uuid
				diskInfo.FileName = tb.FileName
				datastoreRef = *tb.Datastore
			default:
				t.Fatalf("unsupported disk backing: %T", disk.Backing)
			}

			datastore := object.NewDatastore(c, datastoreRef)
			var moDatastore mo.Datastore
			if err := datastore.Properties(
				ctx,
				datastore.Reference(),
				[]string{"info"},
				&moDatastore); err != nil {

				t.Fatal(err)
			}

			datastorePath := moDatastore.Info.GetDatastoreInfo().Url
			var diskPath object.DatastorePath
			if !diskPath.FromString(diskInfo.FileName) {
				t.Fatalf("invalid disk file name: %q", diskInfo.FileName)
			}

			const vmdkSize = 500

			assert.NoError(t, os.WriteFile(
				path.Join(datastorePath, diskPath.Path),
				bytes.Repeat([]byte{1}, vmdkSize),
				os.ModeAppend))
			assert.NoError(t, vm.RefreshStorageInfo(ctx))

			diskInfo.Size = vmdkSize
			diskInfo.UniqueSize = vmdkSize

			testCases := []testCase{
				{
					name:     "ctx is nil",
					ctx:      nil,
					client:   c,
					diskUUID: diskUUID,
					err:      "ctx is nil",
				},
				{
					name:     "client is nil",
					ctx:      context.Background(),
					client:   nil,
					diskUUID: diskUUID,
					err:      "client is nil",
				},
				{
					name:     "failed to retrieve properties",
					ctx:      context.Background(),
					client:   c,
					diskUUID: diskUUID,
					mo: mo.VirtualMachine{
						ManagedEntity: mo.ManagedEntity{
							ExtensibleManagedObject: mo.ExtensibleManagedObject{
								Self: vm.Reference(),
							},
						},
					},
					err: "failed to retrieve properties: ServerFaultCode: InvalidArgument",
				},
				{
					name:            "fetchProperties is false but cached properties are missing",
					ctx:             context.Background(),
					client:          c,
					diskUUID:        diskUUID,
					diskInfo:        diskInfo,
					fetchProperties: false,
					mo: mo.VirtualMachine{
						ManagedEntity: mo.ManagedEntity{
							ExtensibleManagedObject: mo.ExtensibleManagedObject{
								Self: vm.Reference(),
							},
						},
					},
				},
				{
					name:            "fetchProperties is true and cached properties are stale",
					ctx:             context.Background(),
					client:          c,
					diskUUID:        diskUUID,
					diskInfo:        diskInfo,
					fetchProperties: true,
					mo: mo.VirtualMachine{
						ManagedEntity: mo.ManagedEntity{
							ExtensibleManagedObject: mo.ExtensibleManagedObject{
								Self: vm.Reference(),
							},
						},
						Config: &types.VirtualMachineConfigInfo{
							Hardware: types.VirtualHardware{
								Device: []types.BaseVirtualDevice{
									&types.VirtualDisk{
										VirtualDevice: types.VirtualDevice{
											Key:     diskInfo.DeviceKey,
											Backing: diskBacking,
										},
										CapacityInBytes: 500,
									},
								},
							},
						},
					},
				},
			}

			for i := range testCases {
				tc := testCases[i]
				t.Run(tc.name, func(t *testing.T) {
					if strings.HasPrefix(tc.err, "failed to retrieve properties:") {
						propertyCollectorShouldFault = true
						defer func() {
							propertyCollectorShouldFault = false
						}()
					}

					dii, err := vmdk.GetVirtualDiskInfoByUUID(
						tc.ctx,
						tc.client,
						tc.mo,
						tc.fetchProperties,
						tc.excludeSnapshots,
						tc.diskUUID)

					if tc.err != "" {
						assert.EqualError(t, err, tc.err)
					} else {
						assert.NoError(t, err)
					}
					assert.Equal(t, tc.diskInfo, dii)
				})
			}

		})
	})
}

var propertyCollectorShouldFault bool

type propertyCollectorWithFault struct {
	simulator.PropertyCollector
}

func (pc *propertyCollectorWithFault) RetrievePropertiesEx(
	ctx *simulator.Context,
	req *types.RetrievePropertiesEx) soap.HasFault {

	if propertyCollectorShouldFault {
		return &methods.RetrievePropertiesExBody{
			Fault_: simulator.Fault("", &types.InvalidArgument{}),
		}
	}

	return pc.PropertyCollector.RetrievePropertiesEx(ctx, req)
}
