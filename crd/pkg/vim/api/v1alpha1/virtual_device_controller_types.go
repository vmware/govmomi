// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// VirtualControllerType identifies the concrete type of a virtual
// controller device in a virtual machine.
type VirtualControllerType string

const (
	// VirtualControllerTypeAHCI is an AHCI (SATA) controller.
	VirtualControllerTypeAHCI = VirtualControllerType(VirtualDeviceTypeAHCIController)

	// VirtualControllerTypeBusLogic is a BusLogic SCSI controller.
	VirtualControllerTypeBusLogic = VirtualControllerType(VirtualDeviceTypeBusLogicController)

	// VirtualControllerTypeIDE is an IDE controller.
	VirtualControllerTypeIDE = VirtualControllerType(VirtualDeviceTypeIDEController)

	// VirtualControllerTypeLsiLogic is an LSI Logic SCSI controller.
	VirtualControllerTypeLsiLogic = VirtualControllerType(VirtualDeviceTypeLsiLogicController)

	// VirtualControllerTypeLsiLogicSAS is an LSI Logic SAS controller.
	VirtualControllerTypeLsiLogicSAS = VirtualControllerType(VirtualDeviceTypeLsiLogicSASController)

	// VirtualControllerTypeNVMe is an NVMe controller.
	VirtualControllerTypeNVMe = VirtualControllerType(VirtualDeviceTypeNVMeController)

	// VirtualControllerTypeParaVirtualSCSI is a VMware Paravirtual SCSI
	// controller.
	VirtualControllerTypeParaVirtualSCSI = VirtualControllerType(VirtualDeviceTypeParaVirtualSCSIController)

	// VirtualControllerTypePCI is a PCI controller.
	VirtualControllerTypePCI = VirtualControllerType(VirtualDeviceTypePCIController)

	// VirtualControllerTypePS2 is a PS/2 controller.
	VirtualControllerTypePS2 = VirtualControllerType(VirtualDeviceTypePS2Controller)

	// VirtualControllerTypeSATA is a SATA controller.
	VirtualControllerTypeSATA = VirtualControllerType(VirtualDeviceTypeSATAController)

	// VirtualControllerTypeSIO is a Super IO (SIO) controller.
	VirtualControllerTypeSIO = VirtualControllerType(VirtualDeviceTypeSIOController)

	// VirtualControllerTypeUSB is a USB (UHCI/EHCI) controller.
	VirtualControllerTypeUSB = VirtualControllerType(VirtualDeviceTypeUSBController)

	// VirtualControllerTypeUSBXHCI is a USB 3.0 (XHCI) controller.
	VirtualControllerTypeUSBXHCI = VirtualControllerType(VirtualDeviceTypeUSBXHCIController)
)

type VirtualUSBControllerType string

const (
	// VirtualUSBControllerTypeUHCI is a USB 1.1 (UHCI) controller.
	VirtualUSBControllerTypeUHCI = VirtualUSBControllerType(VirtualDeviceTypeUSBController)

	// VirtualUSBControllerTypeEHCI is a USB 2.0 (EHCI) controller.
	VirtualUSBControllerTypeEHCI = VirtualUSBControllerType(VirtualDeviceTypeUSBController)

	// VirtualUSBControllerTypeXHCI is a USB 3.0 (XHCI) controller.
	VirtualUSBControllerTypeXHCI = VirtualUSBControllerType(VirtualDeviceTypeUSBXHCIController)
)

// VirtualController is the base data object for a device controller in a
// virtual machine.
// It corresponds to vim.vm.device.VirtualController.
type VirtualController struct {
	// BusNumber is the bus number associated with this controller.
	BusNumber int32 `json:"busNumber"`

	// +optional

	// Device is the list of keys of virtual devices controlled by this
	// controller.
	Device []int32 `json:"device,omitempty"`

	// +optional

	// NVMe contains NVMe controller-specific data when this controller is
	// an NVMe controller.
	NVMe *VirtualNVMEController `json:"nvme,omitempty"`

	// +optional

	// SCSI contains SCSI controller-specific data when this controller is
	// a SCSI controller.
	SCSI *VirtualSCSIController `json:"scsi,omitempty"`

	// +optional

	// USB contains USB HCI controller-specific data when this controller
	// is a USB (UHCI/EHCI) controller.
	USB *VirtualUSBController `json:"usb,omitempty"`

	// +optional

	// USBXHCI contains USB XHCI controller-specific data when this
	// controller is a USB 3.0 (XHCI) controller.
	USBXHCI *VirtualUSBXHCIController `json:"usbxhci,omitempty"`
}

// +kubebuilder:validation:Enum=None;Physical;Virtual

// VirtualNVMESharing describes the NVME bus sharing mode.
// It corresponds to vim.vm.device.VirtualSCSIController.Sharing.
type VirtualNVMESharing string

const (
	// VirtualNVMESharingNone disables NVME bus sharing.
	VirtualNVMESharingNone VirtualSCSISharing = "None"

	// VirtualNVMESharingPhysical enables physical NVME bus sharing.
	VirtualNVMESharingPhysical VirtualSCSISharing = "Physical"
)

// VirtualNVMEController represents an NVMe controller in a virtual machine.
// It corresponds to vim.vm.device.VirtualNVMEController.
type VirtualNVMEController struct {
	// +optional

	// SharedBus is the shared bus mode of the NVMe controller.
	SharedBus VirtualNVMESharing `json:"sharedBus,omitempty"`
}

// +kubebuilder:validation:Enum=None;Physical;Virtual

// VirtualSCSISharing describes the SCSI bus sharing mode.
// It corresponds to vim.vm.device.VirtualSCSIController.Sharing.
type VirtualSCSISharing string

const (
	// VirtualSCSISharingNone disables SCSI bus sharing.
	VirtualSCSISharingNone VirtualSCSISharing = "None"

	// VirtualSCSISharingPhysical enables physical SCSI bus sharing.
	VirtualSCSISharingPhysical VirtualSCSISharing = "Physical"

	// VirtualSCSISharingVirtualSharing enables virtual SCSI bus sharing.
	VirtualSCSISharingVirtualSharing VirtualSCSISharing = "Virtual"
)

// VirtualSCSIController represents a SCSI controller in a virtual machine.
// It corresponds to vim.vm.device.VirtualSCSIController.
type VirtualSCSIController struct {
	// +optional

	// HotAddRemove indicates whether hot-add and hot-remove of devices is
	// supported. Always true in the current implementation.
	HotAddRemove *bool `json:"hotAddRemove,omitempty"`

	// +optional
	// +kubebuilder:default=None

	// SharedBus is the SCSI bus sharing mode.
	SharedBus VirtualSCSISharing `json:"sharedBus,omitempty"`

	// +optional

	// ScsiCtlrUnitNumber is the unit number of the SCSI controller on its
	// own bus.
	ScsiCtlrUnitNumber int32 `json:"scsiCtlrUnitNumber,omitempty"`
}

// VirtualUSBController represents a USB Host Controller Interface (HCI) in a
// virtual machine.
// It corresponds to vim.vm.device.VirtualUSBController.
type VirtualUSBController struct {
	// +optional

	// AutoConnectDevices indicates whether hot-plugging of devices is enabled
	// on this controller.
	AutoConnectDevices *bool `json:"autoConnectDevices,omitempty"`

	// +optional

	// EhciEnabled indicates whether enhanced host controller interface
	// (USB 2.0) is enabled on this controller.
	EhciEnabled *bool `json:"ehciEnabled,omitempty"`
}

// VirtualUSBXHCIController represents a USB 3.0 eXtensible Host Controller
// Interface (XHCI) in a virtual machine.
// It corresponds to vim.vm.device.VirtualUSBXHCIController.
type VirtualUSBXHCIController struct {
	// +optional

	// AutoConnectDevices indicates whether hot-plugging of devices is enabled
	// on this controller.
	AutoConnectDevices *bool `json:"autoConnectDevices,omitempty"`
}
