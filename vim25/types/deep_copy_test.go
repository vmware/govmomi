// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestMustDeepCopy(t *testing.T) {
	newConfigSpec := func() VirtualMachineConfigSpec {
		return VirtualMachineConfigSpec{
			Name:     "vm-001",
			GuestId:  "otherGuest",
			Files:    &VirtualMachineFileInfo{VmPathName: "[datastore1]"},
			NumCPUs:  1,
			MemoryMB: 128,
			DeviceChange: []BaseVirtualDeviceConfigSpec{
				&VirtualDeviceConfigSpec{
					Operation: VirtualDeviceConfigSpecOperationAdd,
					Device: &VirtualLsiLogicController{VirtualSCSIController{
						SharedBus: VirtualSCSISharingNoSharing,
						VirtualController: VirtualController{
							BusNumber: 0,
							VirtualDevice: VirtualDevice{
								Key: 1000,
							},
						},
					}},
				},
				&VirtualDeviceConfigSpec{
					Operation:     VirtualDeviceConfigSpecOperationAdd,
					FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
					Device: &VirtualDisk{
						VirtualDevice: VirtualDevice{
							Key:           0,
							ControllerKey: 1000,
							UnitNumber:    NewInt32(10),
							Backing: &VirtualDiskFlatVer2BackingInfo{
								DiskMode:        string(VirtualDiskModePersistent),
								ThinProvisioned: NewBool(true),
								VirtualDeviceFileBackingInfo: VirtualDeviceFileBackingInfo{
									FileName: "[datastore1]",
								},
							},
						},
						CapacityInKB: 4000000,
					},
				},
				&VirtualDeviceConfigSpec{
					Operation: VirtualDeviceConfigSpecOperationAdd,
					Device: &VirtualE1000{VirtualEthernetCard{
						VirtualDevice: VirtualDevice{
							Key: 0,
							DeviceInfo: &Description{
								Label:   "Network Adapter 1",
								Summary: "VM Network",
							},
							Backing: &VirtualEthernetCardNetworkBackingInfo{
								VirtualDeviceDeviceBackingInfo: VirtualDeviceDeviceBackingInfo{
									DeviceName: "VM Network",
								},
							},
						},
						AddressType: string(VirtualEthernetCardMacTypeGenerated),
					}},
				},
			},
			ExtraConfig: []BaseOptionValue{
				&OptionValue{Key: "bios.bootOrder", Value: "ethernet0"},
			},
		}
	}

	t.Run("a string", func(t *testing.T) {
		t.Parallel()
		var dst AnyType
		assert.NotPanics(t, func() {
			dst = MustDeepCopy("hello")
		})
		assert.Equal(t, "hello", dst)
	})

	t.Run("a *uint8", func(t *testing.T) {
		t.Parallel()
		var dst AnyType
		assert.NotPanics(t, func() {
			dst = MustDeepCopy(New(uint8(42)))
		})
		assert.Equal(t, &[]uint8{42}[0], dst)
	})

	t.Run("a VirtualMachineConfigSpec", func(t *testing.T) {
		t.Parallel()
		var dst AnyType
		assert.NotPanics(t, func() {
			dst = MustDeepCopy(newConfigSpec())
		})
		assert.Equal(t, newConfigSpec(), dst)
	})

	t.Run("a *VirtualMachineConfigSpec", func(t *testing.T) {
		t.Parallel()
		var dst AnyType
		assert.NotPanics(t, func() {
			dst = MustDeepCopy(New(newConfigSpec()))
		})
		assert.Equal(t, New(newConfigSpec()), dst)
	})

	t.Run("a **VirtualMachineConfigSpec", func(t *testing.T) {
		t.Parallel()

		var dst AnyType
		exp := newConfigSpec()
		ptrExp := &exp
		ptrPtrExp := &ptrExp

		assert.NotPanics(t, func() {
			dst = MustDeepCopy(New(New(newConfigSpec())))
		})

		assert.Equal(t, ptrPtrExp, dst)

		exp.Name += "-not-equal"
		assert.NotEqual(t, ptrPtrExp, dst)
	})

	t.Run("a VirtualMachineConfigSpec with nil DeviceChange vs empty", func(t *testing.T) {
		t.Parallel()

		t.Run("src is nil, exp is nil", func(t *testing.T) {
			t.Parallel()
			var dst AnyType
			exp, src := newConfigSpec(), newConfigSpec()
			exp.DeviceChange = nil
			src.DeviceChange = nil
			assert.NotPanics(t, func() {
				dst = MustDeepCopy(src)
			})
			assert.Equal(t, exp, dst, cmp.Diff(exp, dst))
		})

		t.Run("src is empty, exp is nil", func(t *testing.T) {
			t.Parallel()
			var dst AnyType
			exp, src := newConfigSpec(), newConfigSpec()
			exp.DeviceChange = nil
			src.DeviceChange = []BaseVirtualDeviceConfigSpec{}
			assert.NotPanics(t, func() {
				dst = MustDeepCopy(src)
			})
			assert.Equal(t, exp, dst, cmp.Diff(exp, dst))
		})

	})
}
