// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualHardwareSpec contains the complete hardware configuration of a
// virtual machine.
// It corresponds to vim.vm.VirtualHardware.
type VirtualHardwareSpec struct {
	// +required

	// NumCPU is the number of virtual CPUs present in this virtual
	// machine.
	NumCPU int32 `json:"numCPU"`

	// +required

	// Memory is the memory size.
	Memory resource.Quantity `json:"memory"`

	// +optional

	// NumCoresPerSocket is the number of cores used to distribute virtual
	// CPUs among sockets in this virtual machine.
	NumCoresPerSocket *int32 `json:"numCoresPerSocket,omitempty"`

	// +optional

	// SimultaneousThreads is the number of SMT (Simultaneous
	// Multithreading) threads per core.
	SimultaneousThreads *int32 `json:"simultaneousThreads,omitempty"`

	// +optional

	// VirtualICH7MPresent indicates whether this virtual machine has a
	// Virtual Intel I/O Controller Hub 7.
	VirtualICH7MPresent *bool `json:"virtualICH7MPresent,omitempty"`

	// +optional

	// VirtualSMCPresent indicates whether this virtual machine has a
	// System Management Controller.
	VirtualSMCPresent *bool `json:"virtualSMCPresent,omitempty"`

	// +optional

	// AutoCoresPerSocket indicates whether cores per socket is
	// automatically determined.
	AutoCoresPerSocket *bool `json:"autoCoresPerSocket,omitempty"`

	// +optional

	// MotherboardLayout is the motherboard layout for this virtual
	// machine. Default is "i440bxHostBridge".
	MotherboardLayout string `json:"motherboardLayout,omitempty"`

	// +optional

	Devices VirtualDevices `json:"devices,omitempty"`
}

// VirtualMachineConfigInfoSpec encapsulates the configuration settings and
// virtual hardware for a virtual machine.
// It corresponds to vim.vm.ConfigInfo.
type VirtualMachineConfigInfoSpec struct {
	// +required

	// Name is the display name of the virtual machine.
	Name string `json:"name"`

	// +required

	// GuestFullName is the full name of the guest operating system for
	// the virtual machine. For example: Windows 2000 Professional.
	GuestFullName string `json:"guestFullName"`

	// +required

	// Version is the version string for this virtual machine.
	HardwareVersion string `json:"hardwareVersion"`

	// +required

	// UUID is the 128-bit SMBIOS UUID of a virtual machine represented as
	// a hexadecimal string in "12345678-abcd-1234-cdef-123456789abc"
	// format.
	UUID string `json:"uuid"`

	// +optional

	// InstanceUUID is the VirtualCenter-specific 128-bit UUID of a
	// virtual machine, represented as a hexadecimal string.
	InstanceUUID string `json:"instanceUuid,omitempty"`

	// +optional

	// NPIVNodeWorldWideName is a list of 64-bit node World Wide Names
	// (WWN). These are paired with NPIVPortWorldWideName to be used by
	// NPIV VPORTs instantiated for the virtual machine on the physical
	// HBAs of the host.
	NPIVNodeWorldWideName []int64 `json:"npivNodeWorldWideName,omitempty"`

	// +optional

	// NPIVPortWorldWideName is a list of 64-bit port World Wide Names
	// (WWN).
	NPIVPortWorldWideName []int64 `json:"npivPortWorldWideName,omitempty"`

	// +optional

	// NPIVWorldWideNameType is the source that provides/generates the
	// assigned WWNs.
	NPIVWorldWideNameType string `json:"npivWorldWideNameType,omitempty"`

	// +optional

	// NPIVDesiredNodeWwns is the desired number of NPIV node WWNs to be
	// extended from the original list.
	NPIVDesiredNodeWwns *int32 `json:"npivDesiredNodeWwns,omitempty"`

	// +optional

	// NPIVDesiredPortWwns is the desired number of NPIV port WWNs to be
	// extended from the original list.
	NPIVDesiredPortWwns *int32 `json:"npivDesiredPortWwns,omitempty"`

	// +optional

	// NPIVTemporaryDisabled disables the NPIV capability on the virtual
	// machine temporarily. When set, NPIV VPORTs will not be instantiated
	// but WWNs in the configuration are preserved.
	NPIVTemporaryDisabled *bool `json:"npivTemporaryDisabled,omitempty"`

	// +optional

	// NPIVOnNonRdmDisks indicates whether NPIV can be enabled on the
	// virtual machine with non-RDM disks.
	NPIVOnNonRdmDisks *bool `json:"npivOnNonRdmDisks,omitempty"`

	// +optional

	// LocationID is a hash incorporating the virtual machine's config
	// file location and the UUID of the host assigned to run the virtual
	// machine.
	LocationID string `json:"locationId,omitempty"`

	// +required

	// Template indicates whether the virtual machine is a template.
	Template bool `json:"template"`

	// +required

	// GuestID is the guest operating system configured on a virtual
	// machine.
	GuestID string `json:"guestId"`

	// +required

	// AlternateGuestName is used as the display name for the operating
	// system if GuestID is "other" or "other-64".
	AlternateGuestName string `json:"alternateGuestName"`

	// +optional

	// Annotation is a description for the virtual machine.
	Annotation string `json:"annotation,omitempty"`

	// +required

	// Files describes the locations of files associated with the virtual
	// machine.
	Files VirtualMachineFileInfo `json:"files"`

	// +optional

	// Tools is the VMware Tools configuration for the virtual machine.
	Tools *ToolsConfigInfo `json:"tools,omitempty"`

	// +required

	// DefaultPowerOps defines the configured defaults for power operations
	// on the virtual machine.
	DefaultPowerOps VirtualMachineDefaultPowerOpInfo `json:"defaultPowerOps"`

	// +optional

	// RebootPowerOff indicates whether the next reboot will result in a
	// power off.
	RebootPowerOff *bool `json:"rebootPowerOff,omitempty"`

	// +required

	// Hardware contains the processor, memory, and virtual device
	// configuration of the virtual machine.
	Hardware VirtualHardwareSpec `json:"hardware"`

	// +optional
	// +listType=atomic

	// VCPUConfig is the per-vCPU configuration. The array is indexed by
	// vCPU number.
	VCPUConfig []VirtualMachineVcpuConfig `json:"vcpuConfig,omitempty"`

	// +optional

	// CPUAllocation defines the resource limits for CPU.
	CPUAllocation *ResourceAllocationInfo `json:"cpuAllocation,omitempty"`

	// +optional

	// MemoryAllocation defines the resource limits for memory.
	MemoryAllocation *ResourceAllocationInfo `json:"memoryAllocation,omitempty"`

	// +optional

	// LatencySensitivity describes the latency-sensitivity of the virtual
	// machine.
	LatencySensitivity *LatencySensitivity `json:"latencySensitivity,omitempty"`

	// +optional

	// MemoryHotAddEnabled indicates whether memory can be added while
	// the virtual machine is running.
	MemoryHotAddEnabled *bool `json:"memoryHotAddEnabled,omitempty"`

	// +optional

	// CPUHotAddEnabled indicates whether virtual processors can be added
	// while this virtual machine is running.
	CPUHotAddEnabled *bool `json:"cpuHotAddEnabled,omitempty"`

	// +optional

	// CPUHotRemoveEnabled indicates whether virtual processors can be
	// removed while this virtual machine is running.
	CPUHotRemoveEnabled *bool `json:"cpuHotRemoveEnabled,omitempty"`

	// +optional

	// HotPlugMemoryLimit is the maximum amount of memory that
	// can be added to a running virtual machine.
	HotPlugMemoryLimit *resource.Quantity `json:"hotPlugMemoryLimit,omitempty"`

	// +optional

	// HotPlugMemoryIncrementSize is the memory increment size
	// for adding memory to a running virtual machine.
	HotPlugMemoryIncrementSize *resource.Quantity `json:"hotPlugMemoryIncrementSize,omitempty"`

	// +optional

	// CPUAffinity defines the CPU affinity settings for this virtual
	// machine.
	CPUAffinity *VirtualMachineAffinityInfo `json:"cpuAffinity,omitempty"`

	// +optional

	// MemoryAffinity defines the memory affinity settings for this
	// virtual machine.
	//
	// Deprecated: since vSphere 6.0.
	MemoryAffinity *VirtualMachineAffinityInfo `json:"memoryAffinity,omitempty"`

	// +optional
	// +listType=atomic

	// ExtraConfig contains additional configuration information for the
	// virtual machine.
	ExtraConfig []OptionValue `json:"extraConfig,omitempty"`

	// +optional
	// +listType=atomic

	// CPUFeatureMask specifies CPU feature compatibility masks that
	// override the defaults from the virtual machine's guest OS
	// descriptor.
	CPUFeatureMask []HostCPUIDInfo `json:"cpuFeatureMask,omitempty"`

	// +optional

	// SwapPlacement is the virtual machine swapfile placement policy.
	SwapPlacement string `json:"swapPlacement,omitempty"`

	// +optional

	// BootOptions defines the boot-time behavior of the virtual machine.
	BootOptions *VirtualMachineBootOptions `json:"bootOptions,omitempty"`

	// +optional

	// RepConfig contains the vSphere Replication settings for this
	// virtual machine.
	RepConfig *ReplicationConfigSpec `json:"repConfig,omitempty"`

	// +optional

	// VAssertsEnabled indicates whether user-configured virtual asserts
	// will be triggered during virtual machine replay.
	VAssertsEnabled *bool `json:"vAssertsEnabled,omitempty"`

	// +optional

	// ChangeTrackingEnabled indicates whether changed block tracking for
	// this VM's disks is active.
	ChangeTrackingEnabled *bool `json:"changeTrackingEnabled,omitempty"`

	// +optional

	// Firmware is the firmware type for this virtual machine.
	Firmware *Firmware `json:"firmware,omitempty"`

	// +optional

	// MaxMKSConnections is the maximum number of active remote display
	// connections that the virtual machine will support.
	MaxMKSConnections *int32 `json:"maxMksConnections,omitempty"`

	// +optional

	// GuestAutoLockEnabled indicates whether the guest operating system
	// will logout any active sessions when there are no remote display
	// connections open to the virtual machine.
	GuestAutoLockEnabled *bool `json:"guestAutoLockEnabled,omitempty"`

	// +optional

	// ManagedBy specifies that this VM is managed by a VC extension.
	ManagedBy *ManagedByInfo `json:"managedBy,omitempty"`

	// +optional

	// MemoryReservationLockedToMax indicates whether memory resource
	// reservation for this virtual machine will always be equal to the
	// virtual machine's memory size.
	MemoryReservationLockedToMax *bool `json:"memoryReservationLockedToMax,omitempty"`

	// +optional

	// InitialOverhead contains resource overhead values used only for
	// admission control when determining if a host has sufficient
	// resources to power on this virtual machine.
	InitialOverhead *VirtualMachineConfigInfoOverheadInfo `json:"initialOverhead,omitempty"`

	// +optional

	// NestedHVEnabled indicates whether this VM is configured to use
	// nested hardware-assisted virtualization.
	NestedHVEnabled *bool `json:"nestedHVEnabled,omitempty"`

	// +optional

	// VPMCEnabled indicates whether this VM has virtual CPU performance
	// counters enabled.
	VPMCEnabled *bool `json:"vPMCEnabled,omitempty"`

	// +optional

	// ScheduledHardwareUpgradeInfo contains the configuration of
	// scheduled hardware upgrades and the result of the last attempt.
	ScheduledHardwareUpgradeInfo *ScheduledHardwareUpgradeInfo `json:"scheduledHardwareUpgradeInfo,omitempty"`

	// +optional

	// VFlashCacheReservation specifies the total vFlash resource
	// reservation for the vFlash caches associated with this VM's virtual
	// disks, in bytes.
	//
	// Deprecated: since vSphere 7.0 because vFlash Read Cache end of
	// availability.
	VFlashCacheReservation *int64 `json:"vFlashCacheReservation,omitempty"`

	// +optional

	// VMXConfigChecksum is a checksum of the VMX configuration file.
	VMXConfigChecksum []byte `json:"vmxConfigChecksum,omitempty"`

	// +optional

	// MessageBusTunnelEnabled indicates whether tunneling of clients from
	// the guest VM into the common message bus on the host network is
	// allowed.
	MessageBusTunnelEnabled *bool `json:"messageBusTunnelEnabled,omitempty"`

	// +optional

	// VMStorageObjectID is the virtual machine object identifier for
	// object-based storage systems.
	VMStorageObjectID string `json:"vmStorageObjectId,omitempty"`

	// +optional

	// SwapStorageObjectID is the virtual machine swap object identifier
	// for object-based storage systems.
	SwapStorageObjectID string `json:"swapStorageObjectId,omitempty"`

	// +optional

	// KeyId contains the virtual machine's cryptographic key options.
	KeyId *CryptoKeyId `json:"keyId,omitempty"`

	// +optional

	// MigrateEncryption describes whether encrypted vMotion is required
	// for this VM.
	MigrateEncryption string `json:"migrateEncryption,omitempty"`

	// +optional

	// SgxInfo describes the virtual SGX configuration of the virtual
	// machine.
	SgxInfo *VirtualMachineSgxInfo `json:"sgxInfo,omitempty"`

	// +optional

	// FTEncryptionMode describes whether encrypted Fault Tolerance is
	// required for this VM.
	FTEncryptionMode string `json:"ftEncryptionMode,omitempty"`

	// +optional

	// SEVEnabled indicates whether SEV (Secure Encrypted Virtualization)
	// is enabled.
	SEVEnabled *bool `json:"sevEnabled,omitempty"`

	// +optional

	// NumaInfo describes the virtual NUMA topology of the virtual machine.
	NumaInfo *VirtualMachineVirtualNumaInfo `json:"numaInfo,omitempty"`

	// +optional

	// PMEMFailoverEnabled indicates whether VMs configured to use PMem
	// will be failed over to other hosts by HA.
	//
	// Deprecated: as of vSphere 9.0 APIs with no replacement.
	PMEMFailoverEnabled *bool `json:"pmemFailoverEnabled,omitempty"`

	// +optional

	// VMXStatsCollectionEnabled indicates whether VMXStats collection is
	// enabled for this VM.
	VMXStatsCollectionEnabled *bool `json:"vmxStatsCollectionEnabled,omitempty"`

	// +optional

	// VMOpNotificationToAppEnabled indicates whether operation
	// notification to applications is enabled.
	VMOpNotificationToAppEnabled *bool `json:"vmOpNotificationToAppEnabled,omitempty"`

	// +optional

	// VMOpNotificationTimeout is the operation notification timeout in
	// seconds.
	VMOpNotificationTimeout *int64 `json:"vmOpNotificationTimeout,omitempty"`

	// +optional

	// DeviceSwap reports the current status of the device swap feature.
	DeviceSwap *VirtualMachineVirtualDeviceSwap `json:"deviceSwap,omitempty"`

	// +optional

	// DeviceGroups describes the assignable hardware device groups for
	// this virtual machine.
	DeviceGroups *VirtualMachineVirtualDeviceGroups `json:"deviceGroups,omitempty"`

	// +optional

	// FixedPassthruHotPlugEnabled indicates whether support to add and
	// remove fixed passthrough devices when the VM is running is enabled.
	FixedPassthruHotPlugEnabled *bool `json:"fixedPassthruHotPlugEnabled,omitempty"`

	// +optional

	// MetroFTEnabled indicates whether FT Metro Cluster is enabled for
	// the VM.
	MetroFTEnabled *bool `json:"metroFtEnabled,omitempty"`

	// +optional

	// MetroFTHostGroup indicates the Host Group for FT Metro Cluster
	// enabled virtual machines.
	MetroFTHostGroup string `json:"metroFtHostGroup,omitempty"`

	// +optional

	// TDXEnabled indicates whether TDX (Trust Domain Extensions) is
	// enabled.
	TDXEnabled *bool `json:"tdxEnabled,omitempty"`

	// +optional

	// SEVSNPEnabled indicates whether SEV-SNP (Secure Encrypted
	// Virtualization Secure Nested Paging) is enabled.
	SEVSNPEnabled *bool `json:"sevSnpEnabled,omitempty"`
}

type VirtualMachineConfigInfoDatastoreUrlPair struct {
	// +required

	Name string `json:"name"`

	// +required

	URL string `json:"url"`
}

// VirtualMachineFileInfo contains the locations of virtual machine files
// other than virtual disk files.
// It corresponds to vim.vm.FileInfo.
type VirtualMachineFileInfo struct {
	// +optional

	// VMPathName is the path to the configuration file for the virtual
	// machine, e.g., the .vmx file. This also implicitly defines the
	// configuration directory.
	VMPathName string `json:"vmPathName,omitempty"`

	// +optional

	// SnapshotDirectory is the path to the directory that holds suspend
	// and snapshot files belonging to the virtual machine.
	SnapshotDirectory string `json:"snapshotDirectory,omitempty"`

	// +optional

	// SuspendDirectory is the path to the directory for suspend files.
	SuspendDirectory string `json:"suspendDirectory,omitempty"`

	// +optional

	// LogDirectory is the path to the directory for log files. Defaults
	// to the same directory as the configuration file.
	LogDirectory string `json:"logDirectory,omitempty"`

	// +optional

	// FTMetadataDirectory is the path to the directory for fault
	// tolerance metadata files.
	FTMetadataDirectory string `json:"ftMetadataDirectory,omitempty"`
}

// VirtualMachineDefaultPowerOpInfo holds the configured defaults for power
// operations on a virtual machine.
// It corresponds to vim.vm.DefaultPowerOpInfo.
type VirtualMachineDefaultPowerOpInfo struct {
	// +optional

	// PowerOffType is the advisory default power-off type.
	PowerOffType string `json:"powerOffType,omitempty"`

	// +optional

	// SuspendType is the advisory default suspend type.
	SuspendType string `json:"suspendType,omitempty"`

	// +optional

	// ResetType is the advisory default reset type.
	ResetType string `json:"resetType,omitempty"`

	// +optional

	// DefaultPowerOffType is the default power-off operation: "soft",
	// "hard", or "preset".
	DefaultPowerOffType string `json:"defaultPowerOffType,omitempty"`

	// +optional

	// DefaultSuspendType is the default suspend operation: "soft",
	// "hard", or "preset".
	DefaultSuspendType string `json:"defaultSuspendType,omitempty"`

	// +optional

	// DefaultResetType is the default reset operation: "soft", "hard",
	// or "preset".
	DefaultResetType string `json:"defaultResetType,omitempty"`

	// +optional

	// StandbyAction describes the behavior of the virtual machine when
	// it receives the S1 ACPI call.
	StandbyAction string `json:"standbyAction,omitempty"`
}

// ResourceAllocationInfo specifies the reserved resource requirement and
// the limit (maximum allowed usage) for a given kind of resource.
// It corresponds to vim.ResourceAllocationInfo.
type ResourceAllocationInfo struct {
	// +optional

	// Reservation is the guaranteed amount of resource available to the
	// virtual machine. Units are MHz for CPU.
	Reservation *int64 `json:"reservation,omitempty"`

	// +optional

	// ExpandableReservation indicates whether the reservation can grow
	// beyond the specified value.
	ExpandableReservation *bool `json:"expandableReservation,omitempty"`

	// +optional

	// Limit is the maximum allowed resource usage. Set to -1 for no
	// fixed limit. Units are MHz for CPU.
	Limit *int64 `json:"limit,omitempty"`

	// +optional

	// Shares describes the resource shares used in case of contention.
	Shares *SharesInfo `json:"shares,omitempty"`

	// +optional

	// OverheadLimit is the maximum allowed overhead resource reservation.
	// Units are MB.
	OverheadLimit *int64 `json:"overheadLimit,omitempty"`
}

// VirtualMachineAffinityInfo specifies the scheduling affinity for a
// virtual machine.
// It corresponds to vim.vm.AffinityInfo.
type VirtualMachineAffinityInfo struct {
	// +optional

	// AffinitySet is the list of processors that may be used by the
	// virtual machine. An empty list removes any existing affinity.
	AffinitySet []int32 `json:"affinitySet,omitempty"`
}

// HostCPUIDInfo describes the CPU features of a host or the CPU feature
// requirements of a virtual machine.
// It corresponds to vim.host.CpuIdInfo.
type HostCPUIDInfo struct {
	// +required

	// Level is the level (EAX input to CPUID).
	Level int32 `json:"level"`

	// +optional

	// Vendor is the CPU vendor for which this mask applies.
	Vendor string `json:"vendor,omitempty"`

	// +optional

	// EAX represents the required EAX bits as a formatted bit mask string.
	EAX string `json:"eax,omitempty"`

	// +optional

	// EBX represents the required EBX bits as a formatted bit mask string.
	EBX string `json:"ebx,omitempty"`

	// +optional

	// ECX represents the required ECX bits as a formatted bit mask string.
	ECX string `json:"ecx,omitempty"`

	// +optional

	// EDX represents the required EDX bits as a formatted bit mask string.
	EDX string `json:"edx,omitempty"`
}

// VirtualMachineVirtualDeviceSwap reports the current status of the
// device swap feature for a virtual machine.
// It corresponds to vim.vm.VirtualDeviceSwap.
type VirtualMachineVirtualDeviceSwap struct {
	// +optional

	// LsiToPvscsi describes the status of LSI Logic to ParaVirtual SCSI
	// controller swap.
	LsiToPvscsi *VirtualMachineVirtualDeviceSwapDeviceSwapInfo `json:"lsiToPvscsi,omitempty"`
}

// VirtualMachineVirtualDeviceSwapDeviceSwapInfo contains information
// about a specific device swap operation.
// It corresponds to vim.vm.VirtualDeviceSwap.DeviceSwapInfo.
type VirtualMachineVirtualDeviceSwapDeviceSwapInfo struct {
	// +optional

	// Enabled indicates whether the swap operation is enabled for this
	// virtual machine.
	Enabled *bool `json:"enabled,omitempty"`

	// +optional

	// Applicable indicates whether the swap operation is applicable to
	// this virtual machine.
	Applicable *bool `json:"applicable,omitempty"`

	// +optional

	// Status is the status of the operation.
	Status string `json:"status,omitempty"`
}

// VirtualMachineVirtualDeviceGroups contains information about vendor
// device groups used by a virtual machine.
// It corresponds to vim.vm.VirtualDeviceGroups.
type VirtualMachineVirtualDeviceGroups struct {
	// +optional

	// DeviceGroup is the list of device groups used by this VM.
	DeviceGroup []VirtualMachineVirtualDeviceGroupsDeviceGroup `json:"deviceGroup,omitempty"`
}

// VirtualMachineVirtualDeviceGroupsDeviceGroup describes a device group
// in a virtual machine.
// It corresponds to vim.vm.VirtualDeviceGroups.DeviceGroup.
type VirtualMachineVirtualDeviceGroupsDeviceGroup struct {
	// +required

	// GroupInstanceKey is a unique integer identifying this device group.
	GroupInstanceKey int32 `json:"groupInstanceKey"`

	// +optional

	// DeviceInfo provides a label and summary for the device group.
	DeviceInfo *VirtualDeviceDescription `json:"deviceInfo,omitempty"`
}

// ToolsConfigInfo contains settings for VMware Tools running in the
// guest operating system.
// It corresponds to vim.vm.ToolsConfigInfo.
type ToolsConfigInfo struct {
	// +optional

	// ToolsVersion is the version of VMware Tools installed in the guest.
	ToolsVersion *int32 `json:"toolsVersion,omitempty"`

	// +optional

	// AfterPowerOn indicates whether scripts run after the virtual
	// machine powers on.
	AfterPowerOn *bool `json:"afterPowerOn,omitempty"`

	// +optional

	// AfterResume indicates whether scripts run after the virtual machine
	// resumes.
	AfterResume *bool `json:"afterResume,omitempty"`

	// +optional

	// BeforeGuestReboot indicates whether scripts run before the virtual
	// machine reboots.
	BeforeGuestReboot *bool `json:"beforeGuestReboot,omitempty"`

	// +optional

	// BeforeGuestShutdown indicates whether scripts run before the
	// virtual machine powers off.
	BeforeGuestShutdown *bool `json:"beforeGuestShutdown,omitempty"`

	// +optional

	// BeforeGuestStandby indicates whether scripts run before the
	// virtual machine suspends.
	BeforeGuestStandby *bool `json:"beforeGuestStandby,omitempty"`

	// +optional

	// ToolsUpgradePolicy is the upgrade policy setting for VMware Tools.
	ToolsUpgradePolicy string `json:"toolsUpgradePolicy,omitempty"`

	// +optional

	// PendingCustomization is the filename of a pending customization
	// package on the host.
	PendingCustomization string `json:"pendingCustomization,omitempty"`

	// +optional

	// CustomizationKeyId is the ID of the key used to encrypt the
	// customization package attached to this VM.
	CustomizationKeyId *CryptoKeyId `json:"customizationKeyId,omitempty"`

	// +optional

	// SyncTimeWithHost indicates whether VMware Tools periodically
	// synchronizes guest time with host time.
	SyncTimeWithHost *bool `json:"syncTimeWithHost,omitempty"`

	// +optional

	// SyncTimeWithHostAllowed indicates whether VMware Tools is allowed
	// to synchronize guest time with host time.
	SyncTimeWithHostAllowed *bool `json:"syncTimeWithHostAllowed,omitempty"`

	// +optional

	// ToolsInstallType is the installation type of VMware Tools in the
	// guest operating system.
	ToolsInstallType string `json:"toolsInstallType,omitempty"`

	// +optional

	// LastInstallInfo describes the status of the last tools upgrade
	// attempt.
	LastInstallInfo *ToolsConfigInfoToolsLastInstallInfo `json:"lastInstallInfo,omitempty"`
}

// ToolsConfigInfoToolsLastInstallInfo describes the status of the last
// VMware Tools upgrade attempt.
// It corresponds to vim.vm.ToolsConfigInfo.ToolsLastInstallInfo.
type ToolsConfigInfoToolsLastInstallInfo struct {
	// +required

	// Counter is the number of attempts made to upgrade VMware Tools on
	// this virtual machine.
	Counter int32 `json:"counter"`

	// +optional

	// Fault describes the error from the last upgrade attempt, if any.
	Fault *LocalizedMethodFault `json:"fault,omitempty"`
}

// LocalizedMethodFault describes a method fault with a localized message.
// It corresponds to vmodl.LocalizedMethodFault.
type LocalizedMethodFault struct {
	// +optional

	// LocalizedMessage is a human-readable description of the fault.
	LocalizedMessage string `json:"localizedMessage,omitempty"`
}

// LatencySensitivity specifies the latency-sensitivity of a virtual
// machine or vCPU.
// It corresponds to vim.LatencySensitivity.
type LatencySensitivity struct {
	// +required

	// Level is the nominal latency-sensitivity level.
	Level string `json:"level"`

	// +optional

	// Sensitivity is the custom absolute latency-sensitivity value in
	// microseconds. Only used when Level is "custom".
	//
	// Deprecated: as of vSphere 5.5.
	Sensitivity *int32 `json:"sensitivity,omitempty"`
}

// VirtualMachineVcpuConfig is the per-vCPU configuration.
// It corresponds to vim.vm.VcpuConfig.
type VirtualMachineVcpuConfig struct {
	// +optional

	// LatencySensitivity is the latency-sensitivity specification for
	// this vCPU.
	LatencySensitivity *LatencySensitivity `json:"latencySensitivity,omitempty"`
}

// ManagedByInfo contains information about the extension responsible for
// the lifecycle of a virtual machine.
// It corresponds to vim.ext.ManagedByInfo.
type ManagedByInfo struct {
	// +required

	// ExtensionKey is the key of the extension managing this entity.
	ExtensionKey string `json:"extensionKey"`

	// +required

	// Type is the managed entity type as defined by the managing extension.
	Type string `json:"type"`
}

// VirtualMachineConfigInfoOverheadInfo contains resource overhead values
// used for admission control when powering on the virtual machine.
// It corresponds to vim.vm.ConfigInfo.OverheadInfo.
type VirtualMachineConfigInfoOverheadInfo struct {
	// +optional

	// InitialMemoryReservation is the memory overhead required to power
	// on the virtual machine, in bytes.
	InitialMemoryReservation *int64 `json:"initialMemoryReservation,omitempty"`

	// +optional

	// InitialSwapReservation is the disk space required for swap when
	// powering on the virtual machine, in bytes.
	InitialSwapReservation *int64 `json:"initialSwapReservation,omitempty"`
}

// ScheduledHardwareUpgradeInfo contains settings for scheduled hardware
// upgrades and the result of the last upgrade attempt.
// It corresponds to vim.vm.ScheduledHardwareUpgradeInfo.
type ScheduledHardwareUpgradeInfo struct {
	// +optional

	// UpgradePolicy is the scheduled hardware upgrade policy setting.
	UpgradePolicy string `json:"upgradePolicy,omitempty"`

	// +optional

	// VersionKey is the target hardware version key for the next
	// scheduled upgrade.
	VersionKey string `json:"versionKey,omitempty"`

	// +optional

	// ScheduledHardwareUpgradeStatus is the status of the last scheduled
	// hardware upgrade attempt.
	ScheduledHardwareUpgradeStatus string `json:"scheduledHardwareUpgradeStatus,omitempty"`

	// +optional

	// Fault describes the failure of the last scheduled hardware upgrade
	// attempt, if any.
	Fault *LocalizedMethodFault `json:"fault,omitempty"`
}

// ReplicationConfigSpec encapsulates the replication configuration
// parameters for a virtual machine.
// It corresponds to vim.vm.ReplicationConfigSpec.
type ReplicationConfigSpec struct {
	// +required

	// Generation is a number reflecting the freshness of the
	// replication configuration.
	Generation int64 `json:"generation"`

	// +required

	// VMReplicationId uniquely identifies the replicated VM between
	// primary and secondary sites.
	VMReplicationId string `json:"vmReplicationId"`

	// +required

	// Destination is the IP address of the HBR server in the secondary
	// site.
	Destination string `json:"destination"`

	// +required

	// Port is the port on the HBR server in the secondary site.
	Port int32 `json:"port"`

	// +required

	// RPO is the Recovery Point Objective in minutes.
	RPO int64 `json:"rpo"`

	// +required

	// QuiesceGuestEnabled indicates whether to quiesce the guest file
	// system or applications before creating a consistent replica.
	QuiesceGuestEnabled bool `json:"quiesceGuestEnabled"`

	// +required

	// Paused indicates whether the VM or group has been paused for
	// replication.
	Paused bool `json:"paused"`

	// +required

	// OppUpdatesEnabled indicates whether to perform opportunistic
	// updates between consistent replicas.
	OppUpdatesEnabled bool `json:"oppUpdatesEnabled"`

	// +optional

	// NetCompressionEnabled indicates whether compression is used when
	// sending replication traffic over the network.
	NetCompressionEnabled *bool `json:"netCompressionEnabled,omitempty"`

	// +optional

	// NetEncryptionEnabled indicates whether encryption is used when
	// sending replication traffic over the network.
	NetEncryptionEnabled *bool `json:"netEncryptionEnabled,omitempty"`

	// +optional

	// EncryptionDestination is the IP address of the encryption
	// tunnelling agent on the secondary site.
	EncryptionDestination string `json:"encryptionDestination,omitempty"`

	// +optional

	// EncryptionPort is the port of the encryption tunnelling agent.
	EncryptionPort *int32 `json:"encryptionPort,omitempty"`

	// +optional

	// RemoteCertificateThumbprint is the SHA256 thumbprint of the remote
	// server certificate, used when encryption is enabled.
	RemoteCertificateThumbprint string `json:"remoteCertificateThumbprint,omitempty"`

	// +optional

	// DataSetsReplicationEnabled indicates whether DataSets files are
	// replicated.
	DataSetsReplicationEnabled *bool `json:"dataSetsReplicationEnabled,omitempty"`

	// +optional

	// Disk is the set of virtual disks configured for replication.
	Disk []ReplicationInfoDiskSettings `json:"disk,omitempty"`
}

// ReplicationInfoDiskSettings encapsulates the replication properties of
// a replicated virtual disk.
// It corresponds to vim.vm.ReplicationConfigSpec.DiskSettings.
type ReplicationInfoDiskSettings struct {
	// +required

	// Key is the device key of the disk in the VM's configuration.
	Key int32 `json:"key"`

	// +required

	// DiskReplicationId uniquely identifies the replicated disk between
	// primary and secondary sites.
	DiskReplicationId string `json:"diskReplicationId"`
}

// VirtualMachineSgxInfo describes the virtual SGX (Software Guard
// Extensions) configuration of a virtual machine.
// It corresponds to vim.vm.SgxInfo.
type VirtualMachineSgxInfo struct {
	// +required

	// EpcSize is the size of the virtual Enclave Page Cache in megabytes.
	EpcSize int64 `json:"epcSize"`

	// +optional

	// FlcMode is the Flexible Launch Control mode. Defaults to
	// "unlocked" if unset.
	FlcMode string `json:"flcMode,omitempty"`

	// +optional

	// LePubKeyHash is the SHA256 digest of the provider launch enclave's
	// SIGSTRUCT.MODULUS. Only used when FlcMode is "locked".
	LePubKeyHash string `json:"lePubKeyHash,omitempty"`

	// +optional

	// RequireAttestation indicates whether the virtual machine requires
	// remote attestation.
	RequireAttestation *bool `json:"requireAttestation,omitempty"`
}

// VirtualMachineVirtualNumaInfo describes the virtual NUMA topology
// configuration of a virtual machine.
// It corresponds to vim.vm.VirtualNumaInfo.
type VirtualMachineVirtualNumaInfo struct {
	// +optional

	// AutoCoresPerNumaNode indicates whether the cores per NUMA node
	// are determined automatically.
	AutoCoresPerNumaNode *bool `json:"autoCoresPerNumaNode,omitempty"`

	// +optional

	// CoresPerNumaNode is the number of cores per virtual NUMA node.
	CoresPerNumaNode *int32 `json:"coresPerNumaNode,omitempty"`

	// +optional

	// VnumaOnCpuHotaddExposed indicates whether the virtual NUMA
	// topology is exposed when CPU hot-add is enabled.
	VnumaOnCpuHotaddExposed *bool `json:"vnumaOnCpuHotaddExposed,omitempty"`
}

// VirtualMachineConfigInfoStatus encapsulates the observed state of the
// configuration settings and virtual hardware for a virtual machine.
// It corresponds to vim.vm.ConfigInfo.
type VirtualMachineConfigInfoStatus struct {
	// +required

	// ChangeVersion is a unique identifier for a given version of the
	// configuration. Each change to the configuration updates this value.
	ChangeVersion string `json:"changeVersion"`

	// +required

	// Modified is the last time a virtual machine's configuration was
	// modified.
	Modified metav1.Time `json:"modified"`

	// +optional

	// CreateDate is the time the virtual machine's configuration was
	// created.
	CreateDate *metav1.Time `json:"createDate,omitempty"`

	// +optional

	// DatastoreURL describes the observed set of datastores on which this VM
	// is stored, as well as the URL identification for each datastore.
	// Changes to datastores do not generate property updates on this
	// property. However, when this property is retrieved it returns the
	// current datastore information.
	DatastoreURL []VirtualMachineConfigInfoDatastoreUrlPair `json:"datastoreURL,omitempty"`

	// +optional
	// +listType=atomic

	// VMXRuntimeConfig contains properties established when the VM powers
	// on and examined when the VM is resumed to ensure compatibility with
	// the suspended device state. Only populated while the VM is powered on.
	VMXRuntimeConfig []OptionValue `json:"vmxRuntimeConfig,omitempty"`
}
