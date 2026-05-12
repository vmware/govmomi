// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

// VirtualDisk represents a virtual disk device in a virtual machine.
// It corresponds to vim.vm.device.VirtualDisk.
type VirtualDisk struct {
	// +optional

	// Capacity is the capacity of this virtual disk. If the disk is on
	// a Virtual Volume datastore the disk size must be a multiple of a
	// megabyte.
	Capacity *resource.Quantity `json:"capacity,omitempty"`

	// +optional

	// StorageIOAllocation describes storage I/O resource allocation for this
	// virtual disk.
	StorageIOAllocation *StorageIOAllocationInfo `json:"storageIOAllocation,omitempty"`

	// +optional

	// DiskObjectId is a durable and immutable identifier for the virtual disk.
	DiskObjectId string `json:"diskObjectId,omitempty"`

	// +optional

	// Iofilter lists the IDs of IO Filters associated with the virtual disk.
	Iofilter []string `json:"iofilter,omitempty"`

	// +optional

	// VDiskId is the ID of the virtual disk as a first-class entity.
	VDiskId string `json:"vDiskId,omitempty"`

	// +optional

	// VDiskVersion is the disk descriptor version of the virtual disk.
	VDiskVersion int32 `json:"vDiskVersion,omitempty"`

	// +optional

	// VirtualDiskFormat indicates the disk format.
	// Valid values include "native_4k" and "native_512".
	VirtualDiskFormat string `json:"virtualDiskFormat,omitempty"`

	// +optional

	// NativeUnmanagedLinkedClone indicates whether this disk is a linked clone
	// from an unmanaged delta disk whose parent chain is unavailable.
	NativeUnmanagedLinkedClone *bool `json:"nativeUnmanagedLinkedClone,omitempty"`

	// +optional

	// GuestReadOnly indicates whether the disk is presented to the guest in
	// read-only mode.
	GuestReadOnly *bool `json:"guestReadOnly,omitempty"`
}

// VirtualDiskFlatVer1BackingInfo defines flat-format (GSX Server 2.x) file
// backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.FlatVer1BackingInfo.
type VirtualDiskFlatVer1BackingInfo struct {
	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// Split indicates whether the virtual disk is stored in multiple 2 GB
	// files.
	Split *bool `json:"split,omitempty"`

	// +optional

	// WriteThrough indicates whether writes go directly to the file system
	// or are buffered.
	WriteThrough *bool `json:"writeThrough,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskFlatVer1BackingInfo `json:"parent,omitempty"`
}

// VirtualDiskFlatVer2BackingInfo defines flat-format (ESX Server 2.x/3.x)
// file backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.FlatVer2BackingInfo.
type VirtualDiskFlatVer2BackingInfo struct {
	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// Split indicates whether the virtual disk is stored in multiple 2 GB
	// files.
	Split *bool `json:"split,omitempty"`

	// +optional

	// WriteThrough indicates whether writes go directly to the file system
	// or are buffered.
	WriteThrough *bool `json:"writeThrough,omitempty"`

	// +optional

	// ThinProvisioned indicates whether the virtual disk is thin-provisioned.
	ThinProvisioned *bool `json:"thinProvisioned,omitempty"`

	// +optional

	// EagerlyScrub indicates whether the backing file is completely scrubbed
	// at creation time.
	EagerlyScrub *bool `json:"eagerlyScrub,omitempty"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional

	// ChangeId is the change ID for tracking incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskFlatVer2BackingInfo `json:"parent,omitempty"`

	// +optional

	// DeltaDiskFormat is the format of the delta disk.
	DeltaDiskFormat string `json:"deltaDiskFormat,omitempty"`

	// +optional

	// DigestEnabled indicates whether the disk backing has digest file
	// enabled.
	DigestEnabled *bool `json:"digestEnabled,omitempty"`

	// +optional

	// DeltaGrainSize is the grain size for seSparseFormat delta disks.
	DeltaGrainSize *resource.Quantity `json:"deltaGrainSize,omitempty"`

	// +optional

	// DeltaDiskFormatVariant provides more detailed information for the delta
	// disk format.
	DeltaDiskFormatVariant string `json:"deltaDiskFormatVariant,omitempty"`

	// +optional

	// Sharing is the sharing mode of the virtual disk.
	Sharing string `json:"sharing,omitempty"`

	// +optional

	// KeyId is the encryption key identifier for this disk backing.
	KeyId *CryptoKeyId `json:"keyId,omitempty"`
}

// VirtualDiskLocalPMemBackingInfo defines persistent memory (PMem) file
// backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.LocalPMemBackingInfo.
type VirtualDiskLocalPMemBackingInfo struct {
	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// VolumeUUID is the persistent memory volume UUID associating this
	// virtual disk with a specific host.
	VolumeUUID string `json:"volumeUUID,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`
}

// VirtualDiskRawDiskMappingVer1BackingInfo defines raw device mapping (RDM)
// file backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.RawDiskMappingVer1BackingInfo.
type VirtualDiskRawDiskMappingVer1BackingInfo struct {
	// +optional

	// LunUuid is the unique identifier of the LUN accessed by the RDM.
	LunUuid string `json:"lunUuid,omitempty"`

	// +optional

	// DeviceName is the host-specific device the LUN is accessed through.
	DeviceName string `json:"deviceName,omitempty"`

	// +optional

	// CompatibilityMode is the RDM compatibility mode.
	CompatibilityMode string `json:"compatibilityMode,omitempty"`

	// +optional

	// DiskMode is the disk persistence mode. Only applicable in virtual
	// compatibility mode.
	DiskMode string `json:"diskMode,omitempty"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional

	// ChangeId is the change ID for tracking incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskRawDiskMappingVer1BackingInfo `json:"parent,omitempty"`

	// +optional

	// DeltaDiskFormat is the format of the delta disk.
	DeltaDiskFormat string `json:"deltaDiskFormat,omitempty"`

	// +optional

	// DeltaGrainSize is the grain size for seSparseFormat delta disks.
	DeltaGrainSize *resource.Quantity `json:"deltaGrainSize,omitempty"`

	// +optional

	// Sharing is the sharing mode of the virtual disk.
	Sharing string `json:"sharing,omitempty"`
}

// VirtualDiskRawDiskVer2BackingInfo defines raw host device (VMware Server)
// backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.RawDiskVer2BackingInfo.
type VirtualDiskRawDiskVer2BackingInfo struct {
	// DescriptorFileName is the name of the raw disk descriptor file.
	DescriptorFileName string `json:"descriptorFileName"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// ChangeId is the change ID for tracking incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional

	// Sharing is the sharing mode of the virtual disk.
	Sharing string `json:"sharing,omitempty"`

	// +optional

	Partitioned *VirtualDiskPartitionedRawDiskVer2BackingInfo `json:"partitioned,omitempty"`
}

// VirtualDiskPartitionedRawDiskVer2BackingInfo defines partitioned raw host
// device backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.PartitionedRawDiskVer2BackingInfo.
type VirtualDiskPartitionedRawDiskVer2BackingInfo struct {
	// Partition is the list of partition indexes on the physical disk drive.
	Partition []int32 `json:"partition"`
}

// VirtualDiskSeSparseBackingInfo defines space-efficient sparse format file
// backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.SeSparseBackingInfo.
//
// Deprecated: As of vSphere API 9.0, use VirtualDiskFlatVer2BackingInfo with
// ThinProvisioned set to true.
type VirtualDiskSeSparseBackingInfo struct {
	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// WriteThrough indicates whether writes go directly to the file system
	// or are buffered.
	WriteThrough *bool `json:"writeThrough,omitempty"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional

	// ChangeId is the change ID for tracking incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskSeSparseBackingInfo `json:"parent,omitempty"`

	// +optional

	// DeltaDiskFormat is the format of the delta disk.
	DeltaDiskFormat string `json:"deltaDiskFormat,omitempty"`

	// +optional

	// DigestEnabled indicates whether the disk backing has digest file
	// enabled.
	DigestEnabled *bool `json:"digestEnabled,omitempty"`

	// +optional

	// GrainSize is the grain size.
	GrainSize *resource.Quantity `json:"grainSize,omitempty"`

	// +optional

	// KeyId is the encryption key identifier for this disk backing.
	KeyId *CryptoKeyId `json:"keyId,omitempty"`
}

// VirtualDiskSparseVer1BackingInfo defines sparse-format (GSX Server 2.x)
// file backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.SparseVer1BackingInfo.
type VirtualDiskSparseVer1BackingInfo struct {
	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// Split indicates whether the virtual disk is stored in multiple 2 GB
	// files.
	Split *bool `json:"split,omitempty"`

	// +optional

	// WriteThrough indicates whether writes go directly to the file system
	// or are buffered.
	WriteThrough *bool `json:"writeThrough,omitempty"`

	// +optional

	// SpaceUsed is the space in use for this sparse disk.
	SpaceUsed *resource.Quantity `json:"spaceUsed,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskSparseVer1BackingInfo `json:"parent,omitempty"`
}

// VirtualDiskSparseVer2BackingInfo defines sparse-format (VMware Server) file
// backing for a virtual disk.
// It corresponds to vim.vm.device.VirtualDisk.SparseVer2BackingInfo.
type VirtualDiskSparseVer2BackingInfo struct {
	VirtualDeviceFileBackingInfo `json:",inline"`

	// DiskMode is the disk persistence mode.
	DiskMode string `json:"diskMode"`

	// +optional

	// Split indicates whether the virtual disk is stored in multiple 2 GB
	// files.
	Split *bool `json:"split,omitempty"`

	// +optional

	// WriteThrough indicates whether writes go directly to the file system
	// or are buffered.
	WriteThrough *bool `json:"writeThrough,omitempty"`

	// +optional

	// SpaceUsed is the space in use for this sparse disk.
	SpaceUsed *resource.Quantity `json:"spaceUsed,omitempty"`

	// +optional

	// Uuid is the disk UUID.
	Uuid string `json:"uuid,omitempty"`

	// +optional

	// ContentId indicates the logical contents of the disk backing and its
	// parents.
	ContentId string `json:"contentId,omitempty"`

	// +optional

	// ChangeId is the change ID for tracking incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	// Parent is the parent backing if this is a delta disk backing.
	Parent *VirtualDiskSparseVer2BackingInfo `json:"parent,omitempty"`

	// +optional

	// KeyId is the encryption key identifier for this disk backing.
	KeyId *CryptoKeyId `json:"keyId,omitempty"`
}

// SharesInfo describes a resource share allocation.
// It corresponds to vim.SharesInfo.
type SharesInfo struct {
	// Shares is the number of shares allocated. Only meaningful when
	// Level is set to "custom".
	Shares int32 `json:"shares"`

	// +kubebuilder:validation:Enum=custom;high;low;normal

	// Level is the allocation level. Levels map to pre-determined numeric
	// share values.
	Level string `json:"level"`
}

// StorageIOAllocationInfo describes storage I/O resource allocation for a
// virtual disk.
// It corresponds to vim.StorageIOAllocationInfo.
type StorageIOAllocationInfo struct {
	// +optional

	// Limit is the maximum number of I/O operations per second.
	// Set to -1 to indicate no limit.
	Limit *int64 `json:"limit,omitempty"`

	// +optional

	// Shares describes the I/O shares used during resource contention.
	Shares *SharesInfo `json:"shares,omitempty"`

	// +optional

	// Reservation is the guaranteed number of I/O operations per second.
	Reservation *int32 `json:"reservation,omitempty"`
}
