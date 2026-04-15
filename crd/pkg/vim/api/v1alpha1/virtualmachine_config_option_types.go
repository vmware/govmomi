// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineConfigOptionsSpec defines the desired state of
// VirtualMachineConfigOptions.
type VirtualMachineConfigOptionsSpec struct {
	// HardwareVersion is the desired hardware version for which to get the
	// config options, ex.: vmx-19.
	HardwareVersion string `json:"hardwareVersion"`
}

// VirtualMachineConfigOptionsStatus defines the observed state of
// VirtualMachineConfigOptions.
type VirtualMachineConfigOptionsStatus struct {
	// +optional

	// ObservedGeneration describes the value of the metadata.generation field
	// the last time this object was reconciled by its primary controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// +optional

	// Conditions describes any conditions associated with this object.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional

	// Description provides a human-readable description of this configuration option.
	Description string `json:"description,omitempty"`

	// +optional

	// GuestOSDescriptors lists the supported guest operating systems and their
	// configuration details.
	GuestOSDescriptors []GuestOsDescriptor `json:"guestOSDescriptors,omitempty"`

	// +optional

	// GuestOSDefaultIndex indicates the default guest operating system.
	GuestOSDefaultIndex int32 `json:"guestOSDefaultIndex,omitempty"`

	// +optional

	// HardwareOptions describes processor, memory, and virtual device options.
	HardwareOptions *VirtualHardwareOption `json:"hardwareOptions,omitempty"`

	// +optional

	// Capabilities describes the VM capabilities supported by this configuration.
	Capabilities *VirtualMachineConfigCapability `json:"capabilities,omitempty"`

	// +optional

	// StorageClassOptions describes storage class configuration options.
	StorageClassOptions *StorageClassOption `json:"storageClassOptions,omitempty"`

	// +optional

	// SupportedMonitorTypes lists the monitor types supported (e.g., "debug", "release").
	SupportedMonitorTypes []string `json:"supportedMonitorTypes,omitempty"`

	// +optional

	// SupportedOvfEnvironmentTransports lists the property transports available
	// for the OVF environment.
	SupportedOvfEnvironmentTransports []string `json:"supportedOvfEnvironmentTransports,omitempty"`

	// +optional

	// SupportedOvfInstallTransports lists the transports supported for the OVF
	// installation phase.
	SupportedOvfInstallTransports []string `json:"supportedOvfInstallTransports,omitempty"`

	// +optional

	// DefaultDevices lists the type names of virtual devices created on a
	// virtual machine by default. Clients should not create these devices
	// explicitly.
	DefaultDevices VirtualDevices `json:"defaultDevices,omitempty"`

	// +optional

	// PropertyRelations describes the relationships between properties of the
	// virtual machine config spec.
	PropertyRelations []VirtualMachinePropertyRelation `json:"propertyRelations,omitempty"`
}

// GuestOsDescriptor describes a guest operating system and its configuration
// options. It corresponds to vim.vm.GuestOsDescriptor.
type GuestOsDescriptor struct {
	// ID is the short identifier for the guest OS (e.g., "Windows2000",
	// "OtherLinux64").
	ID VirtualMachineGuestOsIdentifier `json:"id"`

	// FullName is the full descriptive name of the guest OS (e.g.,
	// "Windows 2000 Professional").
	FullName string `json:"fullName"`

	// +optional

	// Family is the family to which this guest OS belongs (e.g.,
	// "Windows", "Linux").
	Family VirtualMachineGuestOsFamily `json:"family,omitempty"`

	// +optional

	// SupportedMaxCPUs is the maximum number of processors (virtual CPUs)
	// supported for this guest.
	SupportedMaxCPUs int32 `json:"supportedMaxCPUs,omitempty"`

	// +optional

	// NumSupportedPhysicalSockets is the maximum number of sockets supported
	// for this guest.
	NumSupportedPhysicalSockets int32 `json:"numSupportedPhysicalSockets,omitempty"`

	// +optional

	// NumSupportedCoresPerSocket is the maximum number of cores per socket
	// supported for this guest.
	NumSupportedCoresPerSocket int32 `json:"numSupportedCoresPerSocket,omitempty"`

	// +optional

	// SupportedMaxMem is the maximum memory supported for this guest.
	SupportedMaxMem *resource.Quantity `json:"supportedMaxMem,omitempty"`

	// +optional

	// SupportedMinMem is the minimum memory required for this guest.
	SupportedMinMem *resource.Quantity `json:"supportedMinMem,omitempty"`

	// +optional

	// RecommendedMem is the recommended default memory size for this guest.
	RecommendedMem *resource.Quantity `json:"recommendedMem,omitempty"`

	// +optional

	// RecommendedColorDepth is the recommended default color depth for this
	// guest.
	RecommendedColorDepth int32 `json:"recommendedColorDepth,omitempty"`

	// +optional

	// RecommendedDiskSize is the recommended default disk size for this guest.
	RecommendedDiskSize *resource.Quantity `json:"recommendedDiskSize,omitempty"`

	// +optional

	// SupportedNumDisks is the number of disks supported for this guest.
	SupportedNumDisks int32 `json:"supportedNumDisks,omitempty"`

	// +optional

	// RecommendedDiskController is the recommended default disk controller
	// type for this guest.
	RecommendedDiskController VirtualControllerType `json:"recommendedDiskController,omitempty"`

	// +optional

	// RecommendedSCSIController is the recommended default SCSI controller
	// type for this guest.
	RecommendedSCSIController VirtualControllerType `json:"recommendedSCSIController,omitempty"`

	// +optional

	// RecommendedCdromController is the recommended default CD-ROM controller
	// type for this guest.
	RecommendedCdromController VirtualControllerType `json:"recommendedCdromController,omitempty"`

	// +optional

	// SupportedDiskControllers lists supported disk controller types for this
	// guest.
	SupportedDiskControllers []VirtualControllerType `json:"supportedDiskControllers,omitempty"`

	// +optional

	// RecommendedEthernetCard is the recommended default ethernet adapter type
	// for this guest.
	RecommendedEthernetCard EthernetCardType `json:"recommendedEthernetCard,omitempty"`

	// +optional

	// SupportedEthernetCards lists supported ethernet adapter types for this
	// guest.
	SupportedEthernetCards []EthernetCardType `json:"supportedEthernetCards,omitempty"`

	// +optional

	// SupportedUSBControllerList lists supported USB controller types for this
	// guest.
	SupportedUSBControllerList []VirtualUSBControllerType `json:"supportedUSBControllerList,omitempty"`

	// +optional

	// RecommendedUSBController is the recommended default USB controller type
	// for this guest.
	RecommendedUSBController VirtualUSBControllerType `json:"recommendedUSBController,omitempty"`

	// +optional

	// WakeOnLanEthernetCard lists the NIC types supported by this guest that
	// also support Wake-on-LAN.
	WakeOnLanEthernetCard []EthernetCardType `json:"wakeOnLanEthernetCard,omitempty"`

	// +optional

	// RecommendedFirmware is the recommended firmware type for this guest.
	RecommendedFirmware Firmware `json:"recommendedFirmware,omitempty"`

	// +optional

	// SupportedFirmware lists supported firmware types for this guest.
	SupportedFirmware []Firmware `json:"supportedFirmware,omitempty"`

	// +optional

	// VRAMSize describes the video RAM size limits supported by this
	// guest.
	VRAMSize *ResourceQuantityOption `json:"vRAMSize,omitempty"`

	// +optional

	// NumSupportedFloppyDevices is the maximum number of floppy devices
	// supported by this guest.
	NumSupportedFloppyDevices int32 `json:"numSupportedFloppyDevices,omitempty"`

	// +optional

	// NumRecommendedPhysicalSockets is the recommended number of sockets for
	// this guest.
	NumRecommendedPhysicalSockets int32 `json:"numRecommendedPhysicalSockets,omitempty"`

	// +optional

	// NumRecommendedCoresPerSocket is the recommended number of cores per
	// socket for this guest.
	NumRecommendedCoresPerSocket int32 `json:"numRecommendedCoresPerSocket,omitempty"`

	// +optional

	// SupportsSlaveDisk indicates whether this guest supports a disk
	// configured as a slave.
	SupportsSlaveDisk *bool `json:"supportsSlaveDisk,omitempty"`

	// +optional

	// SmcRequired indicates that this guest requires an SMC (Apple hardware).
	SmcRequired bool `json:"smcRequired,omitempty"`

	// +optional

	// SmcRecommended indicates whether SMC (Apple hardware) is recommended
	// for this guest.
	SmcRecommended bool `json:"smcRecommended,omitempty"`

	// +optional

	// Ich7mRecommended indicates whether an I/O Controller Hub is recommended
	// for this guest.
	Ich7mRecommended bool `json:"ich7mRecommended,omitempty"`

	// +optional

	// UsbRecommended indicates whether a USB controller is recommended for
	// this guest.
	UsbRecommended bool `json:"usbRecommended,omitempty"`

	// +optional

	// SupportsWakeOnLan indicates whether this guest supports Wake-on-LAN.
	SupportsWakeOnLan bool `json:"supportsWakeOnLan,omitempty"`

	// +optional

	// SupportsVMI indicates whether this guest supports the virtual machine
	// interface.
	SupportsVMI bool `json:"supportsVMI,omitempty"`

	// +optional

	// Supports3D indicates whether this guest supports 3D graphics.
	Supports3D bool `json:"supports3D,omitempty"`

	// +optional

	// Recommended3D indicates whether 3D graphics are recommended for this
	// guest.
	Recommended3D bool `json:"recommended3D,omitempty"`

	// +optional

	// SupportsSecureBoot indicates whether Secure Boot is supported for this
	// guest. Only meaningful when virtual EFI firmware is in use.
	SupportsSecureBoot *bool `json:"supportsSecureBoot,omitempty"`

	// +optional

	// DefaultSecureBoot indicates whether Secure Boot should be enabled by
	// default for this guest. Only meaningful when virtual EFI firmware is in
	// use.
	DefaultSecureBoot *bool `json:"defaultSecureBoot,omitempty"`

	// +optional

	// SupportsCpuHotAdd indicates whether CPUs can be added to this guest
	// while the virtual machine is running.
	SupportsCpuHotAdd bool `json:"supportsCpuHotAdd,omitempty"`

	// +optional

	// SupportsCpuHotRemove indicates whether CPUs can be removed from this
	// guest while the virtual machine is running.
	SupportsCpuHotRemove bool `json:"supportsCpuHotRemove,omitempty"`

	// +optional

	// SupportsMemoryHotAdd indicates whether memory can be added to this
	// guest while the virtual machine is running.
	SupportsMemoryHotAdd bool `json:"supportsMemoryHotAdd,omitempty"`

	// +optional

	// SupportsPvscsiControllerForBoot indicates whether this guest can use
	// a PVSCSI controller as the boot adapter.
	SupportsPvscsiControllerForBoot bool `json:"supportsPvscsiControllerForBoot,omitempty"`

	// +optional

	// DiskUuidEnabled indicates whether disk UUID should be enabled by
	// default for this guest.
	DiskUuidEnabled bool `json:"diskUuidEnabled,omitempty"`

	// +optional

	// SupportsHotPlugPCI indicates whether this guest supports hot-plug of
	// PCI devices.
	SupportsHotPlugPCI bool `json:"supportsHotPlugPCI,omitempty"`

	// +optional

	// SupportsTPM20 indicates whether TPM 2.0 is supported for this guest.
	SupportsTPM20 *bool `json:"supportsTPM20,omitempty"`

	// +optional

	// RecommendedTPM20 indicates whether TPM 2.0 is recommended for this
	// guest.
	RecommendedTPM20 *bool `json:"recommendedTPM20,omitempty"`

	// +optional

	// VvtdSupported indicates support for Intel Virtualization Technology for
	// Directed I/O (VT-d) for this guest.
	VvtdSupported *BoolOption `json:"vvtdSupported,omitempty"`

	// +optional

	// VbsSupported indicates support for Virtualization-based security for
	// this guest.
	VbsSupported *BoolOption `json:"vbsSupported,omitempty"`

	// +optional

	// VsgxSupported indicates support for Intel Software Guard Extensions
	// (SGX) for this guest.
	VsgxSupported *BoolOption `json:"vsgxSupported,omitempty"`

	// +optional

	// VsgxRemoteAttestationSupported indicates support for Intel SGX remote
	// attestation for this guest.
	VsgxRemoteAttestationSupported *bool `json:"vsgxRemoteAttestationSupported,omitempty"`

	// +optional

	// VwdtSupported indicates support for a virtual watchdog timer for this
	// guest.
	VwdtSupported *bool `json:"vwdtSupported,omitempty"`

	// +optional

	// PersistentMemorySupported indicates support for persistent memory
	// (virtual NVDIMM) for this guest.
	PersistentMemorySupported *bool `json:"persistentMemorySupported,omitempty"`

	// +optional

	// SupportedMinPersistentMemory is the minimum persistent memory supported
	// for this guest.
	SupportedMinPersistentMemory *resource.Quantity `json:"supportedMinPersistentMemory,omitempty"`

	// +optional

	// SupportedMaxPersistentMemory is the maximum total persistent memory
	// supported for this guest across all virtual NVDIMM devices.
	SupportedMaxPersistentMemory *resource.Quantity `json:"supportedMaxPersistentMemory,omitempty"`

	// +optional

	// RecommendedPersistentMemory is the recommended default persistent memory
	// size for this guest.
	RecommendedPersistentMemory *resource.Quantity `json:"recommendedPersistentMemory,omitempty"`

	// +optional

	// PersistentMemoryHotAddSupported indicates support for persistent memory
	// hot-add for this guest.
	PersistentMemoryHotAddSupported *bool `json:"persistentMemoryHotAddSupported,omitempty"`

	// +optional

	// PersistentMemoryHotRemoveSupported indicates support for persistent
	// memory hot-remove for this guest.
	PersistentMemoryHotRemoveSupported *bool `json:"persistentMemoryHotRemoveSupported,omitempty"`

	// +optional

	// PersistentMemoryColdGrowthSupported indicates support for virtual NVDIMM
	// cold-growth (capacity increase while powered off) for this guest.
	PersistentMemoryColdGrowthSupported *bool `json:"persistentMemoryColdGrowthSupported,omitempty"`

	// +optional

	// PersistentMemoryColdGrowthGranularity is the granularity for
	// virtual NVDIMM cold-growth operations.
	PersistentMemoryColdGrowthGranularity *ResourceQuantityOption `json:"persistentMemoryColdGrowthGranularity,omitempty"`

	// +optional

	// PersistentMemoryHotGrowthSupported indicates support for virtual NVDIMM
	// hot-growth (capacity increase while powered on) for this guest.
	PersistentMemoryHotGrowthSupported *bool `json:"persistentMemoryHotGrowthSupported,omitempty"`

	// +optional

	// PersistentMemoryHotGrowthGranularity is the granularity for
	// virtual NVDIMM hot-growth operations.
	PersistentMemoryHotGrowthGranularity *ResourceQuantityOption `json:"persistentMemoryHotGrowthGranularity,omitempty"`

	// +optional

	// SupportedForCreate indicates whether this guest OS can be selected
	// during VM creation.
	SupportedForCreate bool `json:"supportedForCreate,omitempty"`

	// +optional

	// SupportLevel indicates the support level for this guest OS.
	//
	// Valid values are:
	//   - Deprecated
	//   - Experimental
	//   - Legacy
	//   - Supported
	//   - TechPreview
	//   - Terminated
	//   - Unsupported
	SupportLevel SupportLevel `json:"supportLevel,omitempty"`
}

// VirtualMachinePropertyRelation describes a relationship between properties
// of the virtual machine config spec.
// It corresponds to vim.vm.PropertyRelation.
type VirtualMachinePropertyRelation struct {
	// Key is the target property and its value.
	Key VirtualMachineProperty `json:"key"`

	// +optional

	// Relations lists the related properties and their values.
	Relations []VirtualMachineProperty `json:"relations,omitempty"`
}

// VirtualMachineProperty represents a named virtual machine config property
// and its string value.
type VirtualMachineProperty struct {
	// Name is the property name.
	Name string `json:"name"`

	// +optional

	// Value is the property value, represented as a string.
	Value string `json:"value,omitempty"`
}

// VirtualHardwareOption describes the hardware configuration options available.
// It corresponds to vim.vm.VirtualHardwareOption.
type VirtualHardwareOption struct {
	// HardwareVersion is the virtual hardware version number.
	HardwareVersion string `json:"hardwareVersion"`

	// +optional

	// DeviceListReadonly indicates whether the set of virtual devices can be
	// changed (i.e., whether devices can be added or removed). This does not
	// preclude changing device properties.
	DeviceListReadonly bool `json:"deviceListReadonly,omitempty"`

	// +optional

	// NumCPU lists acceptable values for the number of CPUs. The first value
	// is the default. These values are typically superseded by the guest OS
	// descriptor's maximum CPU count.
	NumCPU []int32 `json:"numCPU,omitempty"`

	// +optional

	// NumCoresPerSocket describes the range of cores per socket.
	NumCoresPerSocket *IntOption `json:"numCoresPerSocket,omitempty"`

	// +optional

	// AutoCoresPerSocket describes the options for automatically distributing
	// virtual CPUs across sockets.
	AutoCoresPerSocket *BoolOption `json:"autoCoresPerSocket,omitempty"`

	// +optional

	// NumCpuReadonly indicates whether the number of virtual CPUs can be
	// changed.
	NumCpuReadonly bool `json:"numCpuReadonly,omitempty"`

	// +optional

	// NumCPUSimultaneousThreads describes the range of SMT (simultaneous
	// multi-threading) threads.
	NumCPUSimultaneousThreads *IntOption `json:"numCPUSimultaneousThreads,omitempty"`

	// +optional

	// Memory describes the memory range options.
	Memory *ResourceQuantityOption `json:"memory,omitempty"`

	// +optional

	// NumNumaNodes describes the range of NUMA nodes.
	NumNumaNodes *IntOption `json:"numNumaNodes,omitempty"`

	// +optional

	// NumPCIControllers describes the range of PCI controllers.
	NumPCIControllers *IntOption `json:"numPCIControllers,omitempty"`

	// +optional

	// NumIDEControllers describes the range of IDE controllers. Note that
	// SCSI controllers are on the PCI bus and their options are in
	// VirtualPCIControllerOption.
	NumIDEControllers *IntOption `json:"numIDEControllers,omitempty"`

	// +optional

	// NumUSBControllers describes the range of USB (1.x/2.0) controllers.
	NumUSBControllers *IntOption `json:"numUSBControllers,omitempty"`

	// +optional

	// NumUSBXHCIControllers describes the range of USB 3.0 (XHCI) controllers.
	NumUSBXHCIControllers *IntOption `json:"numUSBXHCIControllers,omitempty"`

	// +optional

	// NumSIOControllers describes the range of Super IO (SIO) controllers,
	// which control floppy drives, serial ports, and parallel ports.
	NumSIOControllers *IntOption `json:"numSIOControllers,omitempty"`

	// +optional

	// NumPS2Controllers describes the range of PS/2 controllers for keyboards
	// and mice.
	NumPS2Controllers *IntOption `json:"numPS2Controllers,omitempty"`

	// +optional

	// NumSupportedWwnPorts describes the range of NPIV WorldWidePort names
	// supported.
	NumSupportedWwnPorts *IntOption `json:"numSupportedWwnPorts,omitempty"`

	// +optional

	// NumSupportedWwnNodes describes the range of NPIV WorldWideNode names
	// supported.
	NumSupportedWwnNodes *IntOption `json:"numSupportedWwnNodes,omitempty"`

	// +optional

	// NumTPMDevices describes the range of TPM devices.
	NumTPMDevices *IntOption `json:"numTPMDevices,omitempty"`

	// +optional

	// NumNVDIMMControllers describes the range of NVDIMM controllers.
	NumNVDIMMControllers *IntOption `json:"numNVDIMMControllers,omitempty"`

	// +optional

	// NumPrecisionClockDevices describes the range of virtual precision clock
	// devices.
	NumPrecisionClockDevices *IntOption `json:"numPrecisionClockDevices,omitempty"`

	// +optional

	// NumWDTDevices describes the range of virtual watchdog timer devices.
	NumWDTDevices *IntOption `json:"numWDTDevices,omitempty"`

	// +optional

	// NumDeviceGroups describes the range of device groups supported.
	NumDeviceGroups *IntOption `json:"numDeviceGroups,omitempty"`

	// +optional

	// EpcMemory describes the Intel SGX EPC (Enclave Page Cache) memory range.
	EpcMemory *ResourceQuantityOption `json:"epcMemory,omitempty"`

	// +optional

	// AcpiHostBridgesFirmware lists supported ACPI host bridges firmware
	// types. The list is empty for hardware versions vmx-17 and older, and
	// set to ["EFI"] for vmx-18 or newer.
	AcpiHostBridgesFirmware []Firmware `json:"acpiHostBridgesFirmware,omitempty"`

	// +optional

	// DeviceGroupTypes lists supported device group types.
	DeviceGroupTypes []string `json:"deviceGroupTypes,omitempty"`

	// +optional

	// LicensingLimit lists property names limited by licensing restrictions.
	LicensingLimit []string `json:"licensingLimit,omitempty"`

	// +optional

	// VirtualDeviceOptions describes the virtual device configuration options,
	// aggregated by device type.
	VirtualDeviceOptions []VirtualDeviceOption `json:"virtualDeviceOptions,omitempty"`
}

// VirtualMachineConfigCapability describes the capabilities supported by a
// virtual machine configuration. It corresponds to vim.vm.Capability.
type VirtualMachineConfigCapability struct {
	// +optional

	// SnapshotOperationsSupported indicates whether snapshot operations are
	// supported.
	SnapshotOperationsSupported bool `json:"snapshotOperationsSupported,omitempty"`

	// +optional

	// MultipleSnapshotsSupported indicates whether multiple snapshots are
	// supported.
	MultipleSnapshotsSupported bool `json:"multipleSnapshotsSupported,omitempty"`

	// +optional

	// SnapshotConfigSupported indicates whether snapshot configuration is
	// supported.
	SnapshotConfigSupported bool `json:"snapshotConfigSupported,omitempty"`

	// +optional

	// PoweredOffSnapshotsSupported indicates whether snapshot operations are
	// supported in the powered-off state.
	PoweredOffSnapshotsSupported bool `json:"poweredOffSnapshotsSupported,omitempty"`

	// +optional

	// MemorySnapshotsSupported indicates whether memory snapshots are
	// supported.
	MemorySnapshotsSupported bool `json:"memorySnapshotsSupported,omitempty"`

	// +optional

	// RevertToSnapshotSupported indicates whether reverting to a snapshot is
	// supported.
	RevertToSnapshotSupported bool `json:"revertToSnapshotSupported,omitempty"`

	// +optional

	// QuiescedSnapshotsSupported indicates whether quiesced snapshots are
	// supported.
	QuiescedSnapshotsSupported bool `json:"quiescedSnapshotsSupported,omitempty"`

	// +optional

	// LockSnapshotsSupported indicates whether the snapshot tree can be
	// locked.
	LockSnapshotsSupported bool `json:"lockSnapshotsSupported,omitempty"`

	// +optional

	// DiskOnlySnapshotOnSuspendedVMSupported indicates whether disk-only
	// snapshots can be created while the virtual machine is suspended.
	DiskOnlySnapshotOnSuspendedVMSupported *bool `json:"diskOnlySnapshotOnSuspendedVMSupported,omitempty"`

	// +optional

	// ConsolePreferencesSupported indicates whether console preferences can
	// be set for this virtual machine.
	ConsolePreferencesSupported bool `json:"consolePreferencesSupported,omitempty"`

	// +optional

	// CpuFeatureMaskSupported indicates whether CPU feature requirement masks
	// can be set for this virtual machine.
	CpuFeatureMaskSupported bool `json:"cpuFeatureMaskSupported,omitempty"`

	// +optional

	// S1AcpiManagementSupported indicates whether ACPI S1 settings management
	// is supported.
	S1AcpiManagementSupported bool `json:"s1AcpiManagementSupported,omitempty"`

	// +optional

	// SettingScreenResolutionSupported indicates whether the console screen
	// resolution can be configured.
	SettingScreenResolutionSupported bool `json:"settingScreenResolutionSupported,omitempty"`

	// +optional

	// SettingDisplayTopologySupported indicates whether the console display
	// topology can be configured.
	SettingDisplayTopologySupported bool `json:"settingDisplayTopologySupported,omitempty"`

	// +optional

	// SettingVideoRamSizeSupported indicates whether the video RAM size can
	// be configured.
	SettingVideoRamSizeSupported bool `json:"settingVideoRamSizeSupported,omitempty"`

	// +optional

	// ToolsAutoUpdateSupported indicates whether VMware Tools auto-update is
	// supported.
	ToolsAutoUpdateSupported bool `json:"toolsAutoUpdateSupported,omitempty"`

	// +optional

	// ToolsSyncTimeSupported indicates whether asking tools to sync time with
	// the host is supported.
	ToolsSyncTimeSupported bool `json:"toolsSyncTimeSupported,omitempty"`

	// +optional

	// ToolsSyncTimeAllowSupported indicates support for allowing or
	// disallowing all tools time synchronization with the host.
	ToolsSyncTimeAllowSupported *bool `json:"toolsSyncTimeAllowSupported,omitempty"`

	// +optional

	// VmNpivWwnSupported indicates whether virtual machine NPIV WWN is
	// supported.
	VmNpivWwnSupported bool `json:"vmNpivWwnSupported,omitempty"`

	// +optional

	// NpivWwnOnNonRdmVmSupported indicates whether assigning NPIV WWN to
	// virtual machines without RDM disks is supported.
	NpivWwnOnNonRdmVmSupported bool `json:"npivWwnOnNonRdmVmSupported,omitempty"`

	// +optional

	// VmNpivWwnDisableSupported indicates whether disabling NPIV on the
	// virtual machine is supported.
	VmNpivWwnDisableSupported bool `json:"vmNpivWwnDisableSupported,omitempty"`

	// +optional

	// VmNpivWwnUpdateSupported indicates whether updating NPIV WWNs on the
	// virtual machine is supported.
	VmNpivWwnUpdateSupported bool `json:"vmNpivWwnUpdateSupported,omitempty"`

	// +optional

	// SwapPlacementSupported indicates whether a configurable swapfile
	// placement policy is supported.
	SwapPlacementSupported bool `json:"swapPlacementSupported,omitempty"`

	// +optional

	// VirtualMmuUsageSupported indicates whether the use of nested page table
	// hardware support can be explicitly set.
	VirtualMmuUsageSupported bool `json:"virtualMmuUsageSupported,omitempty"`

	// +optional

	// VirtualMmuUsageIgnored indicates that VirtualMachineFlagInfo.virtualMmuUsage
	// is ignored, always operating as if "on" was selected.
	VirtualMmuUsageIgnored *bool `json:"virtualMmuUsageIgnored,omitempty"`

	// +optional

	// VirtualExecUsageIgnored indicates that VirtualMachineFlagInfo.virtualExecUsage
	// is ignored, always operating as if "hvOn" was selected.
	VirtualExecUsageIgnored *bool `json:"virtualExecUsageIgnored,omitempty"`

	// +optional

	// DiskSharesSupported indicates whether resource settings for disks can
	// be applied to this virtual machine.
	DiskSharesSupported bool `json:"diskSharesSupported,omitempty"`

	// +optional

	// BootOptionsSupported indicates whether boot options can be configured.
	BootOptionsSupported bool `json:"bootOptionsSupported,omitempty"`

	// +optional

	// BootRetryOptionsSupported indicates whether automatic boot retry can
	// be configured.
	BootRetryOptionsSupported bool `json:"bootRetryOptionsSupported,omitempty"`

	// +optional

	// ChangeTrackingSupported indicates whether disk change tracking is
	// supported. Even when supported, it may not be available for all disk
	// types (e.g., passthru RDMs).
	ChangeTrackingSupported bool `json:"changeTrackingSupported,omitempty"`

	// +optional

	// HostBasedReplicationSupported indicates whether host-based replication
	// is supported.
	HostBasedReplicationSupported bool `json:"hostBasedReplicationSupported,omitempty"`

	// +optional

	// GuestAutoLockSupported indicates whether guest OS auto-lock and MKS
	// connection controls are supported.
	GuestAutoLockSupported bool `json:"guestAutoLockSupported,omitempty"`

	// +optional

	// MemoryReservationLockSupported indicates whether
	// memoryReservationLockedToMax may be set to true for this virtual
	// machine.
	MemoryReservationLockSupported bool `json:"memoryReservationLockSupported,omitempty"`

	// +optional

	// FeatureRequirementSupported indicates whether the featureRequirement
	// feature is supported.
	FeatureRequirementSupported bool `json:"featureRequirementSupported,omitempty"`

	// +optional

	// PoweredOnMonitorTypeChangeSupported indicates whether a monitor type
	// change is supported while the virtual machine is powered on.
	PoweredOnMonitorTypeChangeSupported bool `json:"poweredOnMonitorTypeChangeSupported,omitempty"`

	// +optional

	// SeSparseDiskSupported indicates whether the Flex-SE (space-efficient,
	// sparse) format is supported for virtual disks.
	SeSparseDiskSupported bool `json:"seSparseDiskSupported,omitempty"`

	// +optional

	// MultipleCoresPerSocketSupported indicates whether multiple cores per
	// socket is supported.
	MultipleCoresPerSocketSupported bool `json:"multipleCoresPerSocketSupported,omitempty"`

	// +optional

	// NestedHVSupported indicates whether nested hardware-assisted
	// virtualization is supported.
	NestedHVSupported bool `json:"nestedHVSupported,omitempty"`

	// +optional

	// VPMCSupported indicates whether virtualized CPU performance counters
	// are supported.
	VPMCSupported bool `json:"vpmcSupported,omitempty"`

	// +optional

	// PerVmEvcSupported indicates whether Per-VM EVC mode is supported.
	PerVmEvcSupported *bool `json:"perVmEvcSupported,omitempty"`

	// +optional

	// SecureBootSupported indicates whether Secure Boot is supported.
	SecureBootSupported *bool `json:"secureBootSupported,omitempty"`

	// +optional

	// SuspendToMemorySupported indicates whether suspending to memory is
	// supported.
	SuspendToMemorySupported *bool `json:"suspendToMemorySupported,omitempty"`

	// +optional

	// RequireSgxAttestationSupported indicates whether requiring SGX remote
	// attestation is supported.
	RequireSgxAttestationSupported *bool `json:"requireSgxAttestationSupported,omitempty"`

	// +optional

	// ChangeModeDisksSupported indicates whether change mode on virtual disks
	// is supported.
	ChangeModeDisksSupported *bool `json:"changeModeDisksSupported,omitempty"`

	// +optional

	// VendorDeviceGroupSupported indicates whether Vendor Device Groups are
	// supported.
	VendorDeviceGroupSupported *bool `json:"vendorDeviceGroupSupported,omitempty"`

	// +optional

	// SEVSupported indicates whether AMD SEV (Secure Encrypted
	// Virtualization) is supported.
	SEVSupported *bool `json:"sevSupported,omitempty"`

	// +optional

	// SEVSNPSupported indicates whether AMD SEV-SNP (Secure Encrypted
	// Virtualization - Secure Nested Paging) is supported.
	SEVSNPSupported *bool `json:"sevSnpSupported,omitempty"`

	// +optional

	// TDXSupported indicates whether Intel TDX (Trusted Domain Extensions)
	// is supported.
	TDXSupported *bool `json:"tdxSupported,omitempty"`
}

// StorageClassOption describes datastore configuration options.
type StorageClassOption struct {
	// +optional

	// UnsupportedVolumes lists volume types that are not supported for this
	// virtual machine configuration.
	UnsupportedVolumes []string `json:"unsupportedVolumes,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion:true
// +kubebuilder:subresource:status

// VirtualMachineConfigOptions is the schema for the
// VirtualMachineConfigOptions API and
// represents the desired state and observed status of a
// VirtualMachineConfigOptions resource.
type VirtualMachineConfigOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineConfigOptionsSpec   `json:"spec,omitempty"`
	Status VirtualMachineConfigOptionsStatus `json:"status,omitempty"`
}

// GetConditions returns the status conditions for the
// VirtualMachineConfigOptions.
func (p VirtualMachineConfigOptions) GetConditions() []metav1.Condition {
	return p.Status.Conditions
}

// SetConditions sets the status conditions for the
// VirtualMachineConfigOptions.
func (p *VirtualMachineConfigOptions) SetConditions(conditions []metav1.Condition) {
	p.Status.Conditions = conditions
}

// GetConditions returns the conditions for the
// VirtualMachineConfigOptionsStatus.
func (p VirtualMachineConfigOptionsStatus) GetConditions() []metav1.Condition {
	return p.Conditions
}

// SetConditions sets the conditions for the
// VirtualMachineConfigOptionsStatus.
func (p *VirtualMachineConfigOptionsStatus) SetConditions(conditions []metav1.Condition) {
	p.Conditions = conditions
}

// +kubebuilder:object:root=true

// VirtualMachineConfigOptionsList contains a list of
// VirtualMachineConfigOptions objects.
type VirtualMachineConfigOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineConfigOptions `json:"items"`
}

func init() {
	objectTypes = append(
		objectTypes,
		&VirtualMachineConfigOptions{},
		&VirtualMachineConfigOptionsList{})
}
