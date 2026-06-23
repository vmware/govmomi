// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// VirtualDiskOption describes the options for a virtual disk device.
// It corresponds to vim.vm.device.VirtualDiskOption.
type VirtualDiskOption struct {
	// Capacity describes the minimum, maximum, and default disk capacity.
	Capacity ResourceQuantityOption `json:"capacity"`

	// +optional

	// IOAllocation describes the options for storage I/O resource allocation.
	//
	// Deprecated: As of vSphere 8.0 U3, there is no replacement.
	IOAllocation *StorageIOAllocationOption `json:"ioAllocation,omitempty"`

	// +optional

	// VFlashCacheConfigOption describes the options for vFlash cache
	// configuration.
	//
	// Deprecated: since vSphere 7.0 because vFlash Read Cache end of
	// availability.
	VFlashCacheConfigOption *VirtualDiskOptionVFlashCacheConfigOption `json:"vFlashCacheConfigOption,omitempty"`
}

// StorageIOAllocationOption describes the options for storage I/O resource
// allocation. It corresponds to vim.StorageResourceManager.IOAllocationOption.
//
// Deprecated: As of vSphere 8.0 U3, there is no replacement.
type StorageIOAllocationOption struct {
	// Limit describes the range of values allowed for the storage IO limit.
	Limit LongOption `json:"limit"`

	// Shares describes the range of values allowed for IO shares.
	Shares SharesOption `json:"shares"`
}

// +kubebuilder:validation:Enum=Custom;High;Low;Normal

// SharesLevel is the predefined allocation level for resource shares.
// It corresponds to vim.SharesInfo.Level.
type SharesLevel string

const (
	// SharesLevelCustom indicates a custom numeric shares value.
	SharesLevelCustom SharesLevel = "Custom"

	// SharesLevelHigh indicates a high (above normal) share allocation.
	SharesLevelHigh SharesLevel = "High"

	// SharesLevelLow indicates a low (below normal) share allocation.
	SharesLevelLow SharesLevel = "Low"

	// SharesLevelNormal indicates the default share allocation.
	SharesLevelNormal SharesLevel = "Normal"
)

// SharesOption describes the options for shares-based resource allocation.
// It corresponds to vim.SharesOption.
type SharesOption struct {
	// DefaultLevel is the default shares allocation level.
	DefaultLevel SharesLevel `json:"defaultLevel"`

	// Shares describes the range of share values.
	Shares IntOption `json:"shares"`
}

// VirtualDiskOptionVFlashCacheConfigOption describes the options for vFlash
// cache configuration. It corresponds to
// vim.vm.device.VirtualDiskOption.VFlashCacheConfigOption.
//
// Deprecated: since vSphere 7.0 because vFlash Read Cache end of availability.
type VirtualDiskOptionVFlashCacheConfigOption struct {
	// BlockSize describes the cache block size range.
	BlockSize ResourceQuantityOption `json:"blockSize"`

	// CacheConsistencyType describes the available cache data consistency
	// types.
	CacheConsistencyType ChoiceOption `json:"cacheConsistencyType"`

	// CacheMode describes the available cache modes.
	CacheMode ChoiceOption `json:"cacheMode"`

	// Reservation describes the cache reservation range.
	Reservation ResourceQuantityOption `json:"reservation"`
}
