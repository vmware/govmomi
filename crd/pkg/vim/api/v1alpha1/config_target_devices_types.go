// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

type ConfigTargetDevices struct {
	// +optional

	// CDROM describes the desired state for available CD-ROM devices.
	CDROM []VirtualMachineCdromInfo `json:"cdrom,omitempty"`

	// +optional

	// Floppy describes the desired state for available floppy devices.
	Floppy []VirtualMachineTargetInfo `json:"floppy,omitempty"`

	// +optional

	// Serial describes the desired state for available serial port
	// devices.
	Serial []VirtualMachineTargetInfo `json:"serial,omitempty"`

	// +optional

	// Parallel describes the desired state for available parallel
	// port devices.
	Parallel []VirtualMachineTargetInfo `json:"parallel,omitempty"`

	// +optional

	// Sound describes the desired state for available sound
	// devices.
	Sound []VirtualMachineTargetInfo `json:"sound,omitempty"`

	// +optional

	// USB describes the desired state for available USB devices.
	USB []VirtualMachineUSBInfo `json:"usb,omitempty"`

	// +optional

	// PCIPassthrough describes the desired state for available PCI
	// passthrough devices.
	PCIPassthrough []VirtualMachinePCIPassthroughInfo `json:"pciPassthrough,omitempty"`

	// +optional

	// DynamicPassthroughDevices describes the desired state for available
	// dynamic DirectPath PCI devices.
	DynamicPassthroughDevices []VirtualMachineDynamicPassthroughInfo `json:"dynamicPassthroughDevices,omitempty"`

	// +optional

	// SRIOV describes the desired state for available SR-IOV
	// devices.
	SRIOV []VirtualMachineSriovInfo `json:"sriov,omitempty"`

	// +optional

	// VGPUDevice describes the desired state for available vGPU device
	// capabilities.
	VGPUDevice []VirtualMachineVgpuDeviceInfo `json:"vgpuDevice,omitempty"`

	// +optional

	// VGPUProfile describes the desired state for available vGPU profile
	// attributes.
	VGPUProfile []VirtualMachineVgpuProfileInfo `json:"vgpuProfile,omitempty"`

	// +optional

	// SharedGPUPassthroughTypes describes the desired state for available
	// shared GPU passthrough types.
	SharedGPUPassthroughTypes []VirtualMachinePciSharedGpuPassthroughInfo `json:"sharedGpuPassthroughTypes,omitempty"`

	// +optional

	// SGXTargetInfo describes the desired state for Intel SGX targeting.
	SGXTargetInfo *VirtualMachineSgxTargetInfo `json:"sgxTargetInfo,omitempty"`

	// +optional

	// PrecisionClockInfo describes the desired state for host clock
	// resources used by virtual precision clocks.
	PrecisionClockInfo []VirtualMachinePrecisionClockInfo `json:"precisionClockInfo,omitempty"`

	// +optional

	// VendorDeviceGroupInfo describes the desired state for vendor device
	// groups.
	VendorDeviceGroupInfo []VirtualMachineVendorDeviceGroupInfo `json:"vendorDeviceGroupInfo,omitempty"`

	// +optional

	// DVXClassInfo describes the desired state for Device Virtualization
	// Extensions (DVX) classes.
	DVXClassInfo []VirtualMachineDvxClassInfo `json:"dvxClassInfo,omitempty"`

	// +optional

	// IDEDisks describes the desired state for physical IDE disks
	// available as raw disk targets.
	IDEDisks []VirtualMachineIdeDiskDeviceInfo `json:"ideDisks,omitempty"`

	// +optional

	// SCSIDisks describes the desired state for physical SCSI disks
	// available as raw disk mapping targets.
	SCSIDisks []VirtualMachineScsiDiskDeviceInfo `json:"scsiDisks,omitempty"`

	// +optional

	// SCSIPassthrough describes the desired state for generic SCSI
	// passthrough devices.
	SCSIPassthrough []VirtualMachineTargetInfo `json:"scsiPassthrough,omitempty"`

	// +optional

	// VFlashModule describes the desired state for available vFlash modules.
	VFlashModule []VirtualMachineVFlashModuleInfo `json:"vflashModule,omitempty"`
}

type VirtualMachineTargetInfo struct {
	// +optional

	// Name is the device name.
	Name string `json:"name,omitempty"`
}

type VirtualMachineCdromInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// Description of the physical device. This is set only by the server.
	Description string `json:"description,omitempty"`
}

// VirtualMachineUSBInfo describes a USB device backing option.
type VirtualMachineUSBInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// Description is the user-visible name of the USB device.
	Description string `json:"description,omitempty"`

	// +optional

	// Vendor is the vendor ID.
	Vendor int32 `json:"vendor,omitempty"`

	// +optional

	// Product is the product ID.
	Product int32 `json:"product,omitempty"`

	// +optional

	// PhysicalPath is the autoconnect pattern describing the physical path to
	// the specific port on the host where the USB device is attached.
	PhysicalPath string `json:"physicalPath,omitempty"`

	// +optional

	// Family is the list of device class families.
	// For possible values see VirtualMachineUsbInfoFamily.
	Family []string `json:"family,omitempty"`

	// +optional

	// Speed is the list of possible device speeds detected by the server.
	// For possible values see VirtualMachineUsbInfoSpeed.
	Speed []string `json:"speed,omitempty"`
}

// VirtualMachinePCIPassthroughInfo describes a generic PCI device that can
// be attached to a virtual machine.
type VirtualMachinePCIPassthroughInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// PciDevice describes the PCI device, including vendor, class, and
	// device identification information.
	PciDevice HostPCIDevice `json:"pciDevice"`

	// +required

	// SystemID is the ID of the system to which the PCI device is attached.
	SystemID string `json:"systemId"`
}

// VirtualMachineDynamicPassthroughInfo describes a dynamic DirectPath
// PCI device.
type VirtualMachineDynamicPassthroughInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// VendorName is the vendor name of this PCI device.
	VendorName string `json:"vendorName"`

	// +required

	// DeviceName is the device name of this PCI device.
	DeviceName string `json:"deviceName"`

	// +optional

	// CustomLabel is the custom label attached to this PCI device.
	CustomLabel string `json:"customLabel,omitempty"`

	// +required

	// VendorID is the PCI vendor ID for this device.
	VendorID int32 `json:"vendorID"`

	// +required

	// DeviceID is the PCI device ID for this device.
	DeviceID int32 `json:"deviceID"`
}

// HostPCIDevice describes a PCI device present on the host.
// It corresponds to vim.host.PciDevice.
type HostPCIDevice struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// Id is the name ID of this PCI, composed of "bus:slot.function".
	Id string `json:"id"`

	// +required

	// ClassId is the class of this PCI device.
	ClassId int32 `json:"classId"`

	// +required

	// Bus is the bus ID of this PCI device.
	Bus int32 `json:"bus"`

	// +required

	// Slot is the slot ID of this PCI device.
	Slot int32 `json:"slot"`

	// +optional

	// PhysicalSlot is the physical slot number of this PCI device.
	PhysicalSlot int32 `json:"physicalSlot,omitempty"`

	// +optional

	// SlotDescription is the description of the slot.
	SlotDescription string `json:"slotDescription,omitempty"`

	// +required

	// Function is the function ID of this PCI device.
	Function int32 `json:"function"`

	// +required

	// VendorId is the vendor ID of this PCI device.
	//
	// The vendor ID may be negative. vSphere uses an unsigned short to
	// represent it; the WSDL representation uses a signed short. Values
	// greater than 32767 are converted to two's complement. When
	// specifying VirtualPCIPassthroughDeviceBackingInfo.vendorId, use
	// the retrieved HostPciDevice.vendorId value directly.
	VendorId int32 `json:"vendorId"`

	// +required

	// SubVendorId is the subvendor ID of this PCI device.
	//
	// The subvendor ID may be negative. vSphere uses an unsigned short to
	// represent it; the WSDL representation uses a signed short. Values
	// greater than 32767 are converted to two's complement.
	SubVendorId int32 `json:"subVendorId"`

	// +required

	// VendorName is the vendor name of this PCI device.
	VendorName string `json:"vendorName"`

	// +required

	// DeviceId is the device ID of this PCI device.
	//
	// The device ID may be negative. vSphere uses an unsigned short to
	// represent it; the WSDL representation uses a signed short. Values
	// greater than 32767 are converted to two's complement. When
	// specifying VirtualPCIPassthroughDeviceBackingInfo.deviceId, use
	// the retrieved HostPciDevice.deviceId value and convert it to a
	// string.
	DeviceId int32 `json:"deviceId"`

	// +required

	// SubDeviceId is the subdevice ID of this PCI device.
	//
	// The subdevice ID may be negative. vSphere uses an unsigned short to
	// represent it; the WSDL representation uses a signed short. Values
	// greater than 32767 are converted to two's complement.
	SubDeviceId int32 `json:"subDeviceId"`

	// +optional

	// ParentBridge is the parent bridge of this PCI device.
	ParentBridge string `json:"parentBridge,omitempty"`

	// +required

	// DeviceName is the device name of this PCI device.
	DeviceName string `json:"deviceName"`

	// +optional

	// DeviceClassName is the name for the PCI device class (e.g.
	// "Host bridge", "iSCSI device", "Fibre channel HBA").
	DeviceClassName string `json:"deviceClassName,omitempty"`
}

type VirtualMachineSriovDevicePoolInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// Key is used to extend to other device types.
	Key string `json:"key,omitempty"`
}

// VirtualMachineSriovNetworkDevicePoolInfo describes a pool of SR-IOV
// devices sharing common network characteristics on the host.
// It corresponds to vim.host.SriovNetworkDevicePoolInfo.
type VirtualMachineSriovNetworkDevicePoolInfo struct {
	VirtualMachineSriovDevicePoolInfo `json:",inline"`

	// +optional

	// SwitchKey identifies the DVS switch associated with this pool.
	SwitchKey string `json:"switchKey,omitempty"`

	// +optional

	// SwitchUUID is the UUID of the distributed virtual switch associated
	// with this pool.
	SwitchUUID string `json:"switchUUID,omitempty"`
}

// VirtualMachineSriovInfo describes a SR-IOV device that can be attached to
// a virtual machine.
type VirtualMachineSriovInfo struct {
	VirtualMachinePCIPassthroughInfo `json:",inline"`

	// +required

	// VirtualFunction indicates whether the corresponding PCI device is a
	// virtual function instantiated by a SR-IOV capable device.
	VirtualFunction bool `json:"virtualFunction"`

	// +optional

	// Pnic is the name of the physical NIC that is represented by a SR-IOV
	// capable physical function.
	Pnic string `json:"pnic,omitempty"`

	// +optional

	// DevicePool is the SR-IOV device pool information.
	DevicePool *VirtualMachineSriovNetworkDevicePoolInfo `json:"devicePool,omitempty"`
}

// VirtualMachineVgpuDeviceInfo describes a PCI vGPU device and its
// capabilities.
type VirtualMachineVgpuDeviceInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// DeviceName is the vGPU device name.
	DeviceName string `json:"deviceName"`

	// +required

	// DeviceVendorID is a well-known unique identifier for the device.
	// It concatenates the 16-bit PCI vendor ID in the lower bits followed
	// by the 16-bit PCI device ID.
	DeviceVendorID int64 `json:"deviceVendorId"`

	// +required

	// MaxFbSizeInGib is the maximum framebuffer size in gibibytes.
	MaxFbSizeInGib int64 `json:"maxFbSizeInGib"`

	// +required

	// TimeSlicedCapable indicates whether the device is time-sliced capable.
	TimeSlicedCapable bool `json:"timeSlicedCapable"`

	// +required

	// MigCapable indicates whether the device is Multiple Instance GPU
	// capable.
	MigCapable bool `json:"migCapable"`

	// +required

	// ComputeProfileCapable indicates whether the device is compute profile
	// capable.
	ComputeProfileCapable bool `json:"computeProfileCapable"`

	// +required

	// QuadroProfileCapable indicates whether the device is quadro profile
	// capable.
	QuadroProfileCapable bool `json:"quadroProfileCapable"`
}

// VirtualMachineVMotionStunTimeInfo describes the VMotion stun time for a
// vGPU profile at a given migration bandwidth.
type VirtualMachineVMotionStunTimeInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// MigrationBW is the migration bandwidth in Mbps.
	MigrationBW int64 `json:"migrationBW"`

	// +required

	// StunTime is the stun time in seconds.
	StunTime int64 `json:"stunTime"`
}

// VirtualMachineVgpuProfileInfo describes a PCI vGPU profile and its
// attributes.
type VirtualMachineVgpuProfileInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// ProfileName is the vGPU profile name.
	ProfileName string `json:"profileName"`

	// +required

	// DeviceVendorID is a well-known unique identifier for the device that
	// supports this profile. It concatenates the 16-bit PCI vendor ID in
	// the lower bits followed by the 16-bit PCI device ID.
	DeviceVendorID int64 `json:"deviceVendorId"`

	// +required

	// FbSizeInGib is the profile framebuffer size in gibibytes.
	FbSizeInGib int64 `json:"fbSizeInGib"`

	// +required

	// ProfileSharing indicates how this profile is shared within the device.
	ProfileSharing string `json:"profileSharing"`

	// +required

	// ProfileClass indicates the class for this profile.
	ProfileClass string `json:"profileClass"`

	// +optional

	// StunTimeEstimates contains VMotion stun time information for this
	// profile.
	StunTimeEstimates []VirtualMachineVMotionStunTimeInfo `json:"stunTimeEstimates,omitempty"`
}

// VirtualMachinePciSharedGpuPassthroughInfo describes a shared GPU passthrough type.
type VirtualMachinePciSharedGpuPassthroughInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// VGPU describes the VGPU corresponding to this GPU passthrough device.
	VGPU string `json:"vgpu"`
}

// VirtualMachineSgxTargetInfo describes Intel Software Guard Extensions
// information.
type VirtualMachineSgxTargetInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// MaxEpcSize is the maximum size in bytes of EPC available on the
	// compute resource.
	MaxEpcSize int64 `json:"maxEpcSize"`

	// +optional

	// FlcModes are the FLC modes available in the compute resource.
	// For possible values see VirtualMachineSgxInfoFlcModes.
	FlcModes []string `json:"flcModes,omitempty"`

	// +optional

	// LePubKeyHashes are the public key hashes of the provider launch
	// enclaves available in the compute resource.
	LePubKeyHashes []string `json:"lePubKeyHashes,omitempty"`

	// +optional

	// RequireAttestationSupported indicates whether the host or cluster
	// supports requiring SGX remote attestation.
	RequireAttestationSupported bool `json:"requireAttestationSupported,omitempty"`
}

// VirtualMachinePrecisionClockInfo describes host clock resources for precision clocks.
type VirtualMachinePrecisionClockInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// SystemClockProtocol is the clock protocol (e.g., "PTP", "NTP").
	SystemClockProtocol HostDateTimeInfoProtocol `json:"systemClockProtocol,omitempty"`
}

type VirtualMachineVendorDeviceGroupInfoComponentDeviceInfoComponentType string

const (
	VirtualMachineVendorDeviceGroupInfoComponentDeviceInfoComponentTypeDVX VirtualMachineVendorDeviceGroupInfoComponentDeviceInfoComponentType = "DVX"
)

// VirtualMachineVendorDeviceGroupInfoComponentDeviceInfo describes a component
// device within a vendor device group.
type VirtualMachineVendorDeviceGroupInfoComponentDeviceInfo struct {
	// +required

	// Type is the component type.
	Type VirtualMachineVendorDeviceGroupInfoComponentDeviceInfoComponentType `json:"type"`

	// +required

	// VendorName is the name of the component device vendor.
	VendorName string `json:"vendorName"`

	// +required

	// DeviceName is the name of the component device.
	DeviceName string `json:"deviceName"`

	// +required

	// IsConfigurable indicates whether this device may be configured by the
	// user or UI.
	IsConfigurable bool `json:"isConfigurable"`

	// +required

	// Device is the VirtualDevice template for this component device.
	Device VirtualDevice `json:"device"`
}

// VirtualMachineVendorDeviceGroupInfo describes a PCI vendor device group
// device.
type VirtualMachineVendorDeviceGroupInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// DeviceGroupName is the name of the vendor device group.
	DeviceGroupName string `json:"deviceGroupName"`

	// +optional

	// DeviceGroupDescription is the description of the vendor device group.
	DeviceGroupDescription string `json:"deviceGroupDescription,omitempty"`

	// +optional

	// ComponentDeviceInfo describes the component devices of this vendor
	// device group, with one entry per component device in the device group
	// spec.
	ComponentDeviceInfo []VirtualMachineVendorDeviceGroupInfoComponentDeviceInfo `json:"componentDeviceInfo,omitempty"`
}

// VsanHostVsanDiskInfo represents additional VSAN information for a SCSI
// disk used by VSAN, mapping the physical disk to its VSAN disk.
type VsanHostVsanDiskInfo struct {
	// +required

	// VsanUUID is the disk UUID in VSAN.
	VsanUUID string `json:"vsanUuid"`

	// +required

	// FormatVersion is the VSAN file system version number.
	FormatVersion int32 `json:"formatVersion"`
}

// HostDiskDimensionsLba describes logical block addressing dimensions of a
// disk, giving the block count and block size.
type HostDiskDimensionsLba struct {
	// +required

	// Block is the number of blocks.
	Block int64 `json:"block"`

	// +required

	// BlockSize is the size of each block in bytes.
	BlockSize int32 `json:"blockSize"`
}

// HostScsiDisk describes a SCSI disk on the host.
type HostScsiDisk struct {
	// +required

	// DeviceName is the name of the host device.
	DeviceName string `json:"deviceName"`

	// +optional

	// DeviceType is the type of the host device.
	DeviceType string `json:"deviceType,omitempty"`

	// +optional

	// Key is the linkable identifier.
	Key string `json:"key,omitempty"`

	// +optional

	// UUID is the UUID of the SCSI LUN.
	UUID string `json:"uuid,omitempty"`

	// +optional

	// CanonicalName is the canonical name of the SCSI LUN.
	CanonicalName string `json:"canonicalName,omitempty"`

	// +optional

	// DisplayName is the display name of the SCSI LUN.
	DisplayName string `json:"displayName,omitempty"`

	// +optional

	// LunType is the type of SCSI LUN.
	LunType string `json:"lunType,omitempty"`

	// +optional

	// Vendor is the vendor of the SCSI LUN.
	Vendor string `json:"vendor,omitempty"`

	// +optional

	// Model is the model of the SCSI LUN.
	Model string `json:"model,omitempty"`

	// +required

	// Capacity is the size of the SCSI disk using the Logical Block
	// Addressing scheme.
	Capacity HostDiskDimensionsLba `json:"capacity"`

	// +required

	// DevicePath is the file path of the device, which can be opened to
	// create partitions on the disk.
	DevicePath string `json:"devicePath"`

	// +optional

	// Ssd indicates whether the disk is SSD backed.
	Ssd bool `json:"ssd,omitempty"`

	// +optional

	// LocalDisk indicates whether the disk is local.
	LocalDisk bool `json:"localDisk,omitempty"`

	// +optional

	// ScsiDiskType is the type of the disk drive.
	ScsiDiskType SCSIDiskType `json:"scsiDiskType,omitempty"`

	// +optional

	// EmulatedDIXDIFEnabled indicates whether the disk has emulated
	// Data Integrity Extension (DIX) / Data Integrity Field (DIF) enabled.
	EmulatedDIXDIFEnabled bool `json:"emulatedDIXDIFEnabled,omitempty"`

	// +optional

	// PhysicalLocation is the physical location of the disk, if it can be
	// determined.
	PhysicalLocation []string `json:"physicalLocation,omitempty"`

	// +optional

	// UsedByMemoryTiering indicates whether the disk is used for memory
	// tiering.
	UsedByMemoryTiering bool `json:"usedByMemoryTiering,omitempty"`

	// +optional

	// VsanDiskInfo contains additional VSAN information for this disk, if
	// the disk is used by VSAN.
	VsanDiskInfo *VsanHostVsanDiskInfo `json:"vsanDiskInfo,omitempty"`
}

// VirtualMachineScsiDiskDeviceInfo contains detailed information about a
// specific SCSI disk hardware device.
// These devices are for VirtualDisk RawDiskMappingVer1BackingInfo.
type VirtualMachineScsiDiskDeviceInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// Capacity is the size of the disk.
	Capacity *resource.Quantity `json:"capacity,omitempty"`

	// +optional

	// Disk contains detailed information about the SCSI disk.
	Disk *HostScsiDisk `json:"disk,omitempty"`

	// +optional

	// TransportHint is a transport identifier hint used to identify the
	// device. Use the Disk field for a definitive correlation to a host
	// physical disk.
	TransportHint string `json:"transportHint,omitempty"`

	// +optional

	// LunNumber is a LUN number hint used to identify the SCSI device.
	// Use the Disk field for a definitive correlation to a host physical
	// disk.
	LunNumber int32 `json:"lunNumber,omitempty"`
}

// VirtualMachineDvxClassInfo describes a Device Virtualization Extensions
// (DVX) device class.
type VirtualMachineDvxClassInfo struct {
	// +required

	// DeviceClass is the DVX device class.
	DeviceClass ElementDescription `json:"deviceClass"`

	// +required

	// VendorName is the label for the vendor name of this class.
	// The value is defined by vendors of the DVX device class as part of
	// their localizable messages.
	VendorName string `json:"vendorName"`

	// +required

	// SriovNic indicates whether the devices of this class are SR-IOV NICs.
	SriovNic bool `json:"sriovNic"`

	// +optional

	// ConfigParams are the configuration parameters for this DVX device
	// class.
	ConfigParams []OptionDef `json:"configParams,omitempty"`
}

// VirtualMachineIdeDiskDevicePartitionInfo describes the size of a single
// partition on an IDE disk.
type VirtualMachineIdeDiskDevicePartitionInfo struct {
	// +required

	// ID is the identification of the partition.
	ID int32 `json:"id"`

	// +required

	// Capacity is the size of the partition.
	Capacity int32 `json:"capacity"`
}

// VirtualMachineIdeDiskDeviceInfo contains detailed information about a
// specific IDE disk hardware device.
// These devices are for VirtualDisk RawDiskVer2BackingInfo and
// PartitionedRawDiskVer2BackingInfo backings.
type VirtualMachineIdeDiskDeviceInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +optional

	// Capacity is the size of the disk.
	Capacity *resource.Quantity `json:"capacity,omitempty"`

	// +optional

	// PartitionTable describes the partitions on this disk.
	PartitionTable []VirtualMachineIdeDiskDevicePartitionInfo `json:"partitionTable,omitempty"`
}

// HostVFlashManagerVFlashCacheConfigInfoVFlashModuleConfigOption describes
// the configuration options for a vFlash module, including supported block
// sizes, cache sizes, cache modes, and consistency types.
type HostVFlashManagerVFlashCacheConfigInfoVFlashModuleConfigOption struct {
	// +required

	// VFlashModule is the name of the vFlash module.
	VFlashModule string `json:"vFlashModule"`

	// +required

	// VFlashModuleVersion is the version of the vFlash module.
	VFlashModuleVersion string `json:"vFlashModuleVersion"`

	// +required

	// MinSupportedModuleVersion is the minimum supported version of the
	// vFlash module.
	MinSupportedModuleVersion string `json:"minSupportedModuleVersion"`

	// +required

	// CacheMode defines the supported cache modes.
	// For possible values see VirtualDiskVFlashCacheConfigInfoCacheMode.
	CacheMode ChoiceOption `json:"cacheMode"`

	// +required

	// CacheConsistencyType defines the supported cache data consistency
	// types.
	// For possible values see
	// VirtualDiskVFlashCacheConfigInfoCacheConsistencyType.
	CacheConsistencyType ChoiceOption `json:"cacheConsistencyType"`

	// +required

	// BlockSizeOption defines the range of virtual disk cache block
	// sizes.
	BlockSizeOption ResourceQuantityOption `json:"blockSizeOption"`

	// +required

	// ReservationOption defines the range of virtual disk cache sizes.
	ReservationOption ResourceQuantityOption `json:"reservationOption"`

	// +required

	// MaxDiskSize is the maximum size of a virtual disk supported.
	MaxDiskSize resource.Quantity `json:"maxDiskSize"`
}

// VirtualMachineVFlashModuleInfo contains information about a vFlash module
// on the host.
type VirtualMachineVFlashModuleInfo struct {
	VirtualMachineTargetInfo `json:",inline"`

	// +required

	// VFlashModule contains information about the vFlash module.
	VFlashModule HostVFlashManagerVFlashCacheConfigInfoVFlashModuleConfigOption `json:"vFlashModule"`
}
