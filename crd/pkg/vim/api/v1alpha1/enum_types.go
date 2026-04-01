// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// +kubebuilder:validation:Enum=BIOS;EFI

// Firmware identifies the firmware used to boot a virtual machine.
// It corresponds to
// vim.vm.GuestOsDescriptor.GuestOsDescriptorFirmwareType.
type Firmware string

const (
	// FirmwareBIOS indicates traditional BIOS firmware.
	FirmwareBIOS Firmware = "BIOS"

	// FirmwareEFI indicates UEFI/EFI firmware.
	FirmwareEFI Firmware = "EFI"
)

// +kubebuilder:validation:Enum=Deprecated;Experimental;Legacy;Supported;TechPreview;Terminated;Unsupported

// SupportLevel indicates the support level for a guest OS descriptor.
// It corresponds to
// vim.vm.GuestOsDescriptor.GuestOsDescriptorSupportLevel.
type SupportLevel string

const (
	// SupportLevelDeprecated indicates a deprecated guest OS that will be
	// removed in a future release.
	SupportLevelDeprecated SupportLevel = "Deprecated"

	// SupportLevelExperimental indicates an experimental guest OS that is
	// not yet fully validated.
	SupportLevelExperimental SupportLevel = "Experimental"

	// SupportLevelLegacy indicates a legacy guest OS with limited support.
	SupportLevelLegacy SupportLevel = "Legacy"

	// SupportLevelSupported indicates a fully supported guest OS.
	SupportLevelSupported SupportLevel = "Supported"

	// SupportLevelTechPreview indicates a guest OS available as a
	// technology preview.
	SupportLevelTechPreview SupportLevel = "TechPreview"

	// SupportLevelTerminated indicates a guest OS whose support has ended.
	SupportLevelTerminated SupportLevel = "Terminated"

	// SupportLevelUnsupported indicates a guest OS that is not supported.
	SupportLevelUnsupported SupportLevel = "Unsupported"
)

// HostDateTimeInfoProtocol describes types of time synchronization protocols.
// It corresponds to vim.host.DateTimeInfo.Protocol.
type HostDateTimeInfoProtocol string

const (
	// HostDateTimeInfoProtocolNTP indicates Network Time Protocol (NTP).
	HostDateTimeInfoProtocolNTP HostDateTimeInfoProtocol = "NTP"

	// HostDateTimeInfoProtocolPTP indicates Precision Time Protocol (PTP).
	HostDateTimeInfoProtocolPTP HostDateTimeInfoProtocol = "PTP"
)

// +kubebuilder:validation:Enum=Native512;Emulated512;Native4k;SoftwareEmulated4k;Unknown

// SCSIDiskType identifies the type of a SCSI disk drive.
// It corresponds to vim.host.ScsiDisk.ScsiDiskType.
type SCSIDiskType string

const (
	// SCSIDiskTypeNative512 indicates a 512 native sector size drive.
	SCSIDiskTypeNative512 SCSIDiskType = "Native512"

	// SCSIDiskTypeEmulated512 indicates a 4K sector size drive in 512
	// emulation mode.
	SCSIDiskTypeEmulated512 SCSIDiskType = "Emulated512"

	// SCSIDiskTypeNative4k indicates a 4K native sector size drive.
	SCSIDiskTypeNative4k SCSIDiskType = "Native4k"

	// SCSIDiskTypeSoftwareEmulated4k indicates a software emulated 4K
	// drive.
	SCSIDiskTypeSoftwareEmulated4k SCSIDiskType = "SoftwareEmulated4k"

	// SCSIDiskTypeUnknown indicates an unknown disk type.
	SCSIDiskTypeUnknown SCSIDiskType = "Unknown"
)
