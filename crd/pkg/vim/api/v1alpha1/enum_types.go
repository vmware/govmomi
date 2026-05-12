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

// +kubebuilder:validation:Enum=High;Low;Normal

// LatencySensitivityLevel describes the latency-sensitivity level of a
// virtual machine or vCPU.
// It corresponds to vim.LatencySensitivity.SensitivityLevel.
type LatencySensitivityLevel string

const (
	// LatencySensitivityLevelHigh indicates high latency sensitivity.
	LatencySensitivityLevelHigh LatencySensitivityLevel = "High"

	// LatencySensitivityLevelLow indicates low latency sensitivity.
	LatencySensitivityLevelLow LatencySensitivityLevel = "Low"

	// LatencySensitivityLevelNormal indicates normal latency sensitivity.
	LatencySensitivityLevelNormal LatencySensitivityLevel = "Normal"
)

// ToVimType returns the VIM value.
func (t LatencySensitivityLevel) ToVimType() string {
	switch t {
	case LatencySensitivityLevelHigh:
		return "high"
	case LatencySensitivityLevelLow:
		return "low"
	case LatencySensitivityLevelNormal:
		return "normal"
	}
	return string(t)
}

// FromVimType parses the VIM value.
func (t *LatencySensitivityLevel) FromVimType(s string) {
	switch s {
	case "high":
		*t = LatencySensitivityLevelHigh
	case "low":
		*t = LatencySensitivityLevelLow
	case "normal":
		*t = LatencySensitivityLevelNormal
	}
	*t = LatencySensitivityLevel(s)
}

type PhysicalNICFeature string

const (
	// PhysicalNICFeatureRSS enables RSS on the vNIC, allowing
	// the guest OS to distribute incoming traffic across multiple vCPU cores
	// rather than relying on a single core, which is a major bottleneck.
	PhysicalNICFeatureRSS PhysicalNICFeature = "ReceiveSideScaling"

	PhysicalNICFeatureLRO PhysicalNICFeature = "LargeReceiveOffload"
)

// ToVimType returns the VIM value.
func (t PhysicalNICFeature) ToVimType() uint8 {
	switch t {
	case PhysicalNICFeatureRSS:
		return 4
	case PhysicalNICFeatureLRO:
		return 5
	}
	return 0
}

// FromVimType parses the VIM value.
func (t *PhysicalNICFeature) FromVimType(s uint8) {
	switch s {
	case 4:
		*t = PhysicalNICFeatureRSS
	case 5:
		*t = PhysicalNICFeatureLRO
	}
	*t = PhysicalNICFeature("")
}

// +kubebuilder:validation:Enum=PerVNIC;PerVM;PerQueue

// TxRxThreadModel describes the transmit/receive thread model.
type TxRxThreadModel string

const (
	// TxRxThreadModelPerVNIC indicates that each vNIC has its own
	// transmit/receive thread.
	TxRxThreadModelPerVNIC TxRxThreadModel = "PerVNIC"

	// TxRxThreadModelPerVM indicates that each VM has its own
	// transmit/receive thread.
	TxRxThreadModelPerVM TxRxThreadModel = "PerVM"

	// TxRxThreadModelPerQueue indicates that each queue has its own
	// transmit/receive thread.
	TxRxThreadModelPerQueue TxRxThreadModel = "PerQueue"
)

// ToVimType returns the VIM value.
func (t TxRxThreadModel) ToVimType() uint8 {
	switch t {
	case TxRxThreadModelPerVNIC:
		return 1
	case TxRxThreadModelPerVM:
		return 2
	case TxRxThreadModelPerQueue:
		return 3
	}
	return 0
}

// FromVimType parses the VIM value.
func (t *TxRxThreadModel) FromVimType(s uint8) {
	switch s {
	case 1:
		*t = TxRxThreadModelPerVNIC
	case 2:
		*t = TxRxThreadModelPerVM
	case 3:
		*t = TxRxThreadModelPerQueue
	}
	*t = TxRxThreadModel("")
}
