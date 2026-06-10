// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// +kubebuilder:validation:Enum=ipv4;ipv6

// VirtualMachineBootOptionsNetworkBootProtocolType enumerates the protocols
// available for PXE network boot or NetBoot.
// It corresponds to vim.vm.BootOptions.NetworkBootProtocolType.
type VirtualMachineBootOptionsNetworkBootProtocolType string

const (
	// VirtualMachineBootOptionsNetworkBootProtocolTypeIPv4 indicates
	// IPv4 boot protocol.
	VirtualMachineBootOptionsNetworkBootProtocolTypeIPv4 VirtualMachineBootOptionsNetworkBootProtocolType = "IPv4"

	// VirtualMachineBootOptionsNetworkBootProtocolTypeIPv6 indicates
	// IPv6 boot protocol.
	VirtualMachineBootOptionsNetworkBootProtocolTypeIPv6 VirtualMachineBootOptionsNetworkBootProtocolType = "IPv6"
)

// +kubebuilder:validation:Enum=Cdrom;Disk;Ethernet;Floppy

// VirtualMachineBootOptionsBootableDeviceType identifies the type of a
// bootable device in the boot order.
// It corresponds to the concrete subtypes of
// vim.vm.BootOptions.BootableDevice.
type VirtualMachineBootOptionsBootableDeviceType string

const (
	// VirtualMachineBootOptionsBootableDeviceTypeCdrom indicates a
	// CD-ROM device. The first CD-ROM with bootable media found is used.
	VirtualMachineBootOptionsBootableDeviceTypeCdrom VirtualMachineBootOptionsBootableDeviceType = "Cdrom"

	// VirtualMachineBootOptionsBootableDeviceTypeDisk indicates a
	// virtual disk device.
	VirtualMachineBootOptionsBootableDeviceTypeDisk VirtualMachineBootOptionsBootableDeviceType = "Disk"

	// VirtualMachineBootOptionsBootableDeviceTypeEthernet indicates an
	// Ethernet adapter device. PXE boot is attempted from this device.
	VirtualMachineBootOptionsBootableDeviceTypeEthernet VirtualMachineBootOptionsBootableDeviceType = "Ethernet"

	// VirtualMachineBootOptionsBootableDeviceTypeFloppy indicates a
	// floppy device.
	VirtualMachineBootOptionsBootableDeviceTypeFloppy VirtualMachineBootOptionsBootableDeviceType = "Floppy"
)

// VirtualMachineBootOptionsBootableDevice represents a device in the boot
// order for a virtual machine.
// It corresponds to vim.vm.BootOptions.BootableDevice and its subtypes
// (BootableCdromDevice, BootableDiskDevice, BootableEthernetDevice,
// BootableFloppyDevice).
type VirtualMachineBootOptionsBootableDevice struct {
	// Type identifies the concrete bootable device type.
	Type VirtualMachineBootOptionsBootableDeviceType `json:"type"`

	// +optional

	// DeviceKey is the device key of the bootable device.
	// Only applicable when Type is Disk or Ethernet; it references the
	// key property of the corresponding virtual device.
	DeviceKey *int32 `json:"deviceKey,omitempty"`
}

// VirtualMachineBootOptions defines the boot-time behavior of a virtual
// machine.
// It corresponds to vim.vm.BootOptions.
type VirtualMachineBootOptions struct {
	// +optional

	// BootDelay is the delay in milliseconds before starting the boot
	// sequence. The boot delay specifies a time interval between virtual
	// machine power on or restart and the beginning of the boot sequence.
	BootDelay *int64 `json:"bootDelay,omitempty"`

	// +optional

	// EnterBIOSSetup indicates that the virtual machine automatically
	// enters BIOS setup the next time it boots. The virtual machine
	// resets this flag to false so that subsequent boots proceed
	// normally.
	EnterBIOSSetup *bool `json:"enterBIOSSetup,omitempty"`

	// +optional

	// EFISecureBootEnabled indicates whether the virtual machine's
	// firmware will perform signature checks of any EFI images loaded
	// during startup, and will refuse to start any images that do not
	// pass those signature checks.
	EFISecureBootEnabled *bool `json:"efiSecureBootEnabled,omitempty"`

	// +optional

	// BootRetryEnabled indicates whether a virtual machine that fails to
	// boot will try again after the BootRetryDelay time period has
	// expired. When false, the virtual machine waits indefinitely for
	// you to initiate boot retry.
	BootRetryEnabled *bool `json:"bootRetryEnabled,omitempty"`

	// +optional

	// BootRetryDelay is the delay in milliseconds before a boot retry.
	// The virtual machine uses this value only if BootRetryEnabled is
	// true.
	BootRetryDelay *int64 `json:"bootRetryDelay,omitempty"`

	// +optional

	// BootOrder is the boot order. Listed devices are used for booting.
	// After the list is exhausted, the default BIOS boot device
	// algorithm is used. The order of entries is significant: the first
	// device is tried first, then the second, and so on.
	BootOrder []VirtualMachineBootOptionsBootableDevice `json:"bootOrder,omitempty"`

	// +optional

	// NetworkBootProtocol is the protocol to attempt during PXE network
	// boot or NetBoot.
	NetworkBootProtocol VirtualMachineBootOptionsNetworkBootProtocolType `json:"networkBootProtocol,omitempty"`
}
