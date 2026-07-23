// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// SupportedDeviceType identifies a virtual device type that may be
// attached to a controller.
type SupportedDeviceType string

const (
	// SupportedDeviceTypeVirtualCdromOption identifies a CD-ROM device.
	SupportedDeviceTypeVirtualCdromOption = "VirtualCdromOption"

	// SupportedDeviceTypeVirtualDiskOption identifies a virtual disk.
	SupportedDeviceTypeVirtualDiskOption = "VirtualDiskOption"

	// SupportedDeviceTypeVirtualEthernetCardOption identifies a virtual
	// Ethernet card.
	SupportedDeviceTypeVirtualEthernetCardOption = "VirtualEthernetCardOption"

	// SupportedDeviceTypeVirtualFloppyOption identifies a floppy drive.
	SupportedDeviceTypeVirtualFloppyOption = "VirtualFloppyOption"

	// SupportedDeviceTypeVirtualKeyboardOption identifies a virtual
	// keyboard.
	SupportedDeviceTypeVirtualKeyboardOption = "VirtualKeyboardOption"

	// SupportedDeviceTypeVirtualNVDIMMOption identifies a virtual NVDIMM.
	SupportedDeviceTypeVirtualNVDIMMOption = "VirtualNVDIMMOption"

	// SupportedDeviceTypeVirtualNVMEControllerOption identifies a virtual
	// NVMe controller.
	SupportedDeviceTypeVirtualNVMEControllerOption = "VirtualNVMEControllerOption"

	// SupportedDeviceTypeVirtualParallelPortOption identifies a virtual
	// parallel port.
	SupportedDeviceTypeVirtualParallelPortOption = "VirtualParallelPortOption"

	// SupportedDeviceTypeVirtualPointingDeviceOption identifies a virtual
	// pointing device.
	SupportedDeviceTypeVirtualPointingDeviceOption = "VirtualPointingDeviceOption"

	// SupportedDeviceTypeVirtualSATAControllerOption identifies a virtual
	// SATA controller.
	SupportedDeviceTypeVirtualSATAControllerOption = "VirtualSATAControllerOption"

	// SupportedDeviceTypeVirtualSCSIControllerOption identifies a virtual
	// SCSI controller.
	SupportedDeviceTypeVirtualSCSIControllerOption = "VirtualSCSIControllerOption"

	// SupportedDeviceTypeVirtualSCSIPassthroughOption identifies a
	// virtual SCSI passthrough device.
	SupportedDeviceTypeVirtualSCSIPassthroughOption = "VirtualSCSIPassthroughOption"

	// SupportedDeviceTypeVirtualSerialPortOption identifies a virtual
	// serial port.
	SupportedDeviceTypeVirtualSerialPortOption = "VirtualSerialPortOption"

	// SupportedDeviceTypeVirtualSoundCardOption identifies a virtual
	// sound card.
	SupportedDeviceTypeVirtualSoundCardOption = "VirtualSoundCardOption"

	// SupportedDeviceTypeVirtualUSBOption identifies a virtual USB device.
	SupportedDeviceTypeVirtualUSBOption = "VirtualUSBOption"

	// SupportedDeviceTypeVirtualVideoCardOption identifies a virtual video
	// card.
	SupportedDeviceTypeVirtualVideoCardOption = "VirtualVideoCardOption"
)

// VirtualDeviceOption describes the options for a virtual device type,
// including configuration options and its relationship to other devices.
// It corresponds to vim.vm.device.VirtualDeviceOption.
type VirtualDeviceOption struct {
	// +optional

	// Type is the name of the run-time class to instantiate for this device.
	Type VirtualDeviceType `json:"type,omitempty"`

	// +optional

	// ConnectOption describes the connect options and defaults for connectable
	// devices.
	ConnectOption *VirtualDeviceConnectOption `json:"connectOption,omitempty"`

	// +optional

	// ControllerType is the data object type of the controller option that is
	// valid for controlling this device.
	ControllerType VirtualControllerType `json:"controllerType,omitempty"`

	// +optional

	// AutoAssignController indicates whether this device will be
	// auto-assigned a controller if one is required.
	AutoAssignController *BoolOption `json:"autoAssignController,omitempty"`

	// +optional

	// DefaultBackingOptionIndex is an index into the backing option list that
	// indicates the default backing.
	DefaultBackingOptionIndex int32 `json:"defaultBackingOptionIndex,omitempty"`

	// +optional

	// LicensingLimit lists property names limited by a licensing restriction
	// of the underlying product.
	LicensingLimit []string `json:"licensingLimit,omitempty"`

	// +optional

	// Deprecated indicates whether this device type is deprecated and cannot
	// be used when creating or reconfiguring a virtual machine.
	Deprecated bool `json:"deprecated,omitempty"`

	// +optional

	// PlugAndPlay indicates whether this device type can be hot-added to a
	// running virtual machine.
	PlugAndPlay bool `json:"plugAndPlay,omitempty"`

	// +optional

	// HotRemoveSupported indicates whether this device type can be
	// hot-removed from a running virtual machine.
	HotRemoveSupported bool `json:"hotRemoveSupported,omitempty"`

	// +optional

	// NumaSupported indicates whether NUMA affinity is supported for this
	// device.
	NumaSupported *bool `json:"numaSupported,omitempty"`

	// +optional

	// Controller describes options specific to virtual controller devices.
	// This field is set when the device type is a controller.
	Controller *VirtualControllerOption `json:"controller,omitempty"`

	// +optional

	// Disk describes options specific to virtual disk devices.
	// This field is set when the device type is a virtual disk.
	Disk *VirtualDiskOption `json:"disk,omitempty"`

	// +optional

	// EthernetCard describes options specific to virtual Ethernet card devices.
	// This field is set when the device type is a virtual Ethernet card.
	EthernetCard *VirtualEthernetCardOption `json:"ethernetCard,omitempty"`

	// +optional

	// SerialPort describes options specific to virtual serial port devices.
	// This field is set when the device type is a virtual serial port.
	SerialPort *VirtualSerialPortOption `json:"serialPort,omitempty"`

	// +optional

	// VideoCard describes options specific to virtual video card devices.
	// This field is set when the device type is a virtual video card.
	VideoCard *VirtualVideoCardOption `json:"videoCard,omitempty"`

	// +optional

	// NVDIMM describes options specific to virtual NVDIMM devices.
	// This field is set when the device type is a virtual NVDIMM.
	NVDIMM *VirtualNVDIMMOption `json:"nvdimm,omitempty"`

	// +optional

	// VMCI describes options specific to virtual VMCI devices.
	// This field is set when the device type is a virtual VMCI device.
	VMCI *VirtualMachineVMCIDeviceOption `json:"vmci,omitempty"`

	// +optional

	// TPM describes options specific to virtual TPM devices.
	// This field is set when the device type is a virtual TPM.
	TPM *VirtualTPMOption `json:"tpm,omitempty"`

	// +optional

	// WDT describes options specific to virtual watchdog timer devices.
	// This field is set when the device type is a virtual watchdog timer.
	WDT *VirtualWDTOption `json:"wdt,omitempty"`
}

// VirtualDeviceConnectOption describes the connect options for a connectable
// virtual device. It corresponds to vim.vm.device.VirtualDeviceOption.ConnectOption.
type VirtualDeviceConnectOption struct {
	// StartConnected indicates whether the device supports the
	// startConnected feature.
	StartConnected BoolOption `json:"startConnected"`

	// AllowGuestControl indicates whether the device can be connected or
	// disconnected from within the guest operating system.
	AllowGuestControl BoolOption `json:"allowGuestControl"`
}

// VirtualMachineVMCIDeviceOption describes the options for a VMCI device.
// It corresponds to vim.vm.VirtualMachineVMCIDeviceOption.
type VirtualMachineVMCIDeviceOption struct {
	// AllowUnrestrictedCommunication indicates support for VMCI communication
	// with all other virtual machines on the host.
	AllowUnrestrictedCommunication BoolOption `json:"allowUnrestrictedCommunication"`

	// +optional

	// FilterSpecOption describes the available options for each VMCI firewall
	// filter specification.
	FilterSpecOption *VirtualMachineVMCIDeviceOptionFilterSpecOption `json:"filterSpecOption,omitempty"`

	// +optional

	// FilterSupported indicates support for VMCI firewall filters.
	FilterSupported *BoolOption `json:"filterSupported,omitempty"`
}

// VirtualNVDIMMOption describes the options for a virtual NVDIMM device.
// It corresponds to vim.vm.device.VirtualNVDIMMOption.
type VirtualNVDIMMOption struct {
	// Capacity describes the minimum and maximum capacity.
	Capacity ResourceQuantityOption `json:"capacity"`

	// Growable indicates whether capacity growth is supported for
	// powered-off virtual machines.
	Growable bool `json:"growable"`

	// HotGrowable indicates whether capacity growth is supported for
	// powered-on virtual machines.
	HotGrowable bool `json:"hotGrowable"`

	// Granularity is the capacity growth granularity, if
	// growth is supported.
	Granularity ResourceQuantityOption `json:"granularity"`
}

// VirtualSerialPortOption describes the options for a virtual serial port.
// It corresponds to vim.vm.device.VirtualSerialPortOption.
type VirtualSerialPortOption struct {
	// YieldOnPoll indicates whether the virtual machine supports CPU yield
	// behavior when polling the virtual serial port.
	YieldOnPoll BoolOption `json:"yieldOnPoll"`
}

// VirtualTPMOption describes the options for a virtual TPM device.
// It corresponds to vim.vm.device.VirtualTPMOption.
type VirtualTPMOption struct {
	// +optional

	// SupportedFirmware lists the supported firmware types for this TPM device,
	// using GuestOsDescriptorFirmwareType enumeration values.
	SupportedFirmware []Firmware `json:"supportedFirmware,omitempty"`
}

// VirtualVideoCardOption describes the options for a virtual video card.
// It corresponds to vim.vm.device.VirtualVideoCardOption.
type VirtualVideoCardOption struct {
	// +optional

	// VideoRamSize describes the minimum, maximum, and default video
	// frame buffer size.
	VideoRamSize *ResourceQuantityOption `json:"videoRamSize,omitempty"`

	// +optional

	// NumDisplays describes the minimum, maximum, and default number of
	// displays.
	NumDisplays *IntOption `json:"numDisplays,omitempty"`

	// +optional

	// UseAutoDetect indicates whether host display settings can be used to
	// automatically determine the virtual video card display settings.
	UseAutoDetect *BoolOption `json:"useAutoDetect,omitempty"`

	// +optional

	// Support3D indicates whether the virtual video card supports 3D
	// functions.
	Support3D *BoolOption `json:"support3D,omitempty"`

	// +optional

	// Use3dRendererSupported indicates whether the virtual video card can
	// specify how to render 3D graphics.
	Use3dRendererSupported *BoolOption `json:"use3dRendererSupported,omitempty"`

	// +optional

	// GraphicsMemorySize describes the minimum, maximum, and default
	// graphics memory size.
	GraphicsMemorySize *ResourceQuantityOption `json:"graphicsMemorySize,omitempty"`

	// +optional

	// GraphicsMemorySizeSupported indicates whether the virtual video card
	// can specify the graphics memory size.
	GraphicsMemorySizeSupported *BoolOption `json:"graphicsMemorySizeSupported,omitempty"`
}

// VirtualWDTOption describes the options for a virtual watchdog timer device.
// It corresponds to vim.vm.device.VirtualWDTOption.
type VirtualWDTOption struct {
	// RunOnBoot indicates whether the "run on boot" option is settable on
	// this device.
	RunOnBoot BoolOption `json:"runOnBoot"`
}

// VirtualMachineVMCIDeviceOptionFilterSpecOption describes the available
// options for a VMCI firewall filter specification.
// It corresponds to vim.vm.VirtualMachineVMCIDeviceOption.FilterSpecOption.
type VirtualMachineVMCIDeviceOptionFilterSpecOption struct {
	// Action describes the available filter actions.
	Action ChoiceOption `json:"action"`

	// Protocol describes the available filter protocols.
	Protocol ChoiceOption `json:"protocol"`

	// Direction describes the available filter directions.
	Direction ChoiceOption `json:"direction"`

	// LowerDstPortBoundary describes the range for the lower destination
	// port boundary.
	LowerDstPortBoundary LongOption `json:"lowerDstPortBoundary"`

	// UpperDstPortBoundary describes the range for the upper destination
	// port boundary.
	UpperDstPortBoundary LongOption `json:"upperDstPortBoundary"`
}
