// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

// VirtualDeviceBackingInfoType identifies the concrete runtime type of a
// virtual device's backing information. The values correspond to vSphere
// API class names (e.g. vim.vm.device.VirtualDisk.FlatVer2BackingInfo).
type VirtualDeviceBackingInfoType string

const (
	// CD-ROM backing types
	VirtualDeviceBackingInfoTypeCdromATAPI             VirtualDeviceBackingInfoType = "VirtualCdromAtapiBackingInfo"
	VirtualDeviceBackingInfoTypeCdromISO               VirtualDeviceBackingInfoType = "VirtualCdromIsoBackingInfo"
	VirtualDeviceBackingInfoTypeCdromPassthrough       VirtualDeviceBackingInfoType = "VirtualCdromPassthroughBackingInfo"
	VirtualDeviceBackingInfoTypeCdromRemoteATAPI       VirtualDeviceBackingInfoType = "VirtualCdromRemoteAtapiBackingInfo"
	VirtualDeviceBackingInfoTypeCdromRemotePassthrough VirtualDeviceBackingInfoType = "VirtualCdromRemotePassthroughBackingInfo"

	// Disk backing types
	VirtualDeviceBackingInfoTypeDiskFlatVer1           VirtualDeviceBackingInfoType = "VirtualDiskFlatVer1BackingInfo"
	VirtualDeviceBackingInfoTypeDiskFlatVer2           VirtualDeviceBackingInfoType = "VirtualDiskFlatVer2BackingInfo"
	VirtualDeviceBackingInfoTypeDiskLocalPMem          VirtualDeviceBackingInfoType = "VirtualDiskLocalPMemBackingInfo"
	VirtualDeviceBackingInfoTypeDiskPartitionedRawVer2 VirtualDeviceBackingInfoType = "VirtualDiskPartitionedRawDiskVer2BackingInfo"
	VirtualDeviceBackingInfoTypeDiskRawMappingVer1     VirtualDeviceBackingInfoType = "VirtualDiskRawDiskMappingVer1BackingInfo"
	VirtualDeviceBackingInfoTypeDiskRawVer2            VirtualDeviceBackingInfoType = "VirtualDiskRawDiskVer2BackingInfo"
	VirtualDeviceBackingInfoTypeDiskSeSparse           VirtualDeviceBackingInfoType = "VirtualDiskSeSparseBackingInfo"
	VirtualDeviceBackingInfoTypeDiskSparseVer1         VirtualDeviceBackingInfoType = "VirtualDiskSparseVer1BackingInfo"
	VirtualDeviceBackingInfoTypeDiskSparseVer2         VirtualDeviceBackingInfoType = "VirtualDiskSparseVer2BackingInfo"

	// Ethernet card backing types
	VirtualDeviceBackingInfoTypeEthernetCardDistributedVirtualPort VirtualDeviceBackingInfoType = "VirtualEthernetCardDistributedVirtualPortBackingInfo"
	VirtualDeviceBackingInfoTypeEthernetCardLegacyNetwork          VirtualDeviceBackingInfoType = "VirtualEthernetCardLegacyNetworkBackingInfo"
	VirtualDeviceBackingInfoTypeEthernetCardNetwork                VirtualDeviceBackingInfoType = "VirtualEthernetCardNetworkBackingInfo"
	VirtualDeviceBackingInfoTypeEthernetCardOpaqueNetwork          VirtualDeviceBackingInfoType = "VirtualEthernetCardOpaqueNetworkBackingInfo"

	// Floppy backing types
	VirtualDeviceBackingInfoTypeFloppyDevice       VirtualDeviceBackingInfoType = "VirtualFloppyDeviceBackingInfo"
	VirtualDeviceBackingInfoTypeFloppyImage        VirtualDeviceBackingInfoType = "VirtualFloppyImageBackingInfo"
	VirtualDeviceBackingInfoTypeFloppyRemoteDevice VirtualDeviceBackingInfoType = "VirtualFloppyRemoteDeviceBackingInfo"

	// NVDIMM backing type
	VirtualDeviceBackingInfoTypeNVDIMM VirtualDeviceBackingInfoType = "VirtualNVDIMMBackingInfo"

	// Parallel port backing types
	VirtualDeviceBackingInfoTypeParallelPortDevice VirtualDeviceBackingInfoType = "VirtualParallelPortDeviceBackingInfo"
	VirtualDeviceBackingInfoTypeParallelPortFile   VirtualDeviceBackingInfoType = "VirtualParallelPortFileBackingInfo"

	// PCI passthrough backing types
	VirtualDeviceBackingInfoTypePCIPassthroughDevice VirtualDeviceBackingInfoType = "VirtualPCIPassthroughDeviceBackingInfo"

	// TODO(akutz) While this type does extend VirtualDeviceBackingInfo, it is
	//             not used as a value for a VirtualDevice's backing field.
	//             Rather it is used with the VirtualSriovEthernetCard's
	//             dvxBacking field.
	// VirtualDeviceBackingInfoTypePCIPassthroughDvx     VirtualDeviceBackingInfoType = "VirtualPCIPassthroughDvxBackingInfo"

	VirtualDeviceBackingInfoTypePCIPassthroughDynamic VirtualDeviceBackingInfoType = "VirtualPCIPassthroughDynamicBackingInfo"
	VirtualDeviceBackingInfoTypePCIPassthroughPlugin  VirtualDeviceBackingInfoType = "VirtualPCIPassthroughPluginBackingInfo"
	VirtualDeviceBackingInfoTypePCIPassthroughVmiop   VirtualDeviceBackingInfoType = "VirtualPCIPassthroughVmiopBackingInfo"

	// Pointing device backing type
	VirtualDeviceBackingInfoTypePointingDeviceDevice VirtualDeviceBackingInfoType = "VirtualPointingDeviceDeviceBackingInfo"

	// Precision clock backing type
	VirtualDeviceBackingInfoTypePrecisionClockSystemClock VirtualDeviceBackingInfoType = "VirtualPrecisionClockSystemClockBackingInfo"

	// SCSI passthrough backing type
	VirtualDeviceBackingInfoTypeSCSIPassthroughDevice VirtualDeviceBackingInfoType = "VirtualSCSIPassthroughDeviceBackingInfo"

	// Serial port backing types
	VirtualDeviceBackingInfoTypeSerialPortDevice    VirtualDeviceBackingInfoType = "VirtualSerialPortDeviceBackingInfo"
	VirtualDeviceBackingInfoTypeSerialPortFile      VirtualDeviceBackingInfoType = "VirtualSerialPortFileBackingInfo"
	VirtualDeviceBackingInfoTypeSerialPortPipe      VirtualDeviceBackingInfoType = "VirtualSerialPortPipeBackingInfo"
	VirtualDeviceBackingInfoTypeSerialPortThinPrint VirtualDeviceBackingInfoType = "VirtualSerialPortThinPrintBackingInfo"
	VirtualDeviceBackingInfoTypeSerialPortURI       VirtualDeviceBackingInfoType = "VirtualSerialPortURIBackingInfo"

	// Sound card backing type
	VirtualDeviceBackingInfoTypeSoundCardDevice VirtualDeviceBackingInfoType = "VirtualSoundCardDeviceBackingInfo"

	// SR-IOV Ethernet card backing type
	// TODO(akutz) While this type does extend VirtualDeviceBackingInfo, it is
	//             not used as a value for a VirtualDevice's backing field.
	//             Rather it is used with the VirtualSriovEthernetCard's
	//             sriovBacking field.
	//VirtualDeviceBackingInfoTypeSriovEthernetCardSriov VirtualDeviceBackingInfoType = "VirtualSriovEthernetCardSriovBackingInfo"

	// USB backing types
	VirtualDeviceBackingInfoTypeUSBRemoteClient VirtualDeviceBackingInfoType = "VirtualUSBRemoteClientBackingInfo"
	VirtualDeviceBackingInfoTypeUSBRemoteHost   VirtualDeviceBackingInfoType = "VirtualUSBRemoteHostBackingInfo"
	VirtualDeviceBackingInfoTypeUSBUSB          VirtualDeviceBackingInfoType = "VirtualUSBUSBBackingInfo"
)

// +kubebuilder:validation:XValidation:rule="[has(self.device), has(self.ethernetCardDistributedVirtualPort), has(self.ethernetCardOpaqueNetwork), has(self.file), has(self.pciPassthroughPlugin), has(self.pipe), has(self.precisionClock), has(self.remoteDevice), has(self.uri)].filter(x, x).size() <= 1",message="at most one of device, ethernetCardDistributedVirtualPort, ethernetCardOpaqueNetwork, file, pciPassthroughPlugin, pipe, precisionClock, remoteDevice, or uri may be specified"

// VirtualDeviceBackingInfo maps the polymorphic virtual device backing
// information hierarchy to a Kubernetes-compatible flat structure.
// It corresponds to vim.vm.device.VirtualDevice.BackingInfo.
type VirtualDeviceBackingInfo struct {

	// Type is the type of the virtual device backing information.
	Type VirtualDeviceBackingInfoType `json:"type"`

	// +optional

	// Device contains host device backing information.
	Device *VirtualDeviceDeviceBackingInfo `json:"device,omitempty"`

	// +optional

	// EthernetCardDistributedVirtualPort contains DVS port backing
	// information for virtual Ethernet cards.
	EthernetCardDistributedVirtualPort *VirtualEthernetCardDistributedVirtualPortBackingInfo `json:"ethernetCardDistributedVirtualPort,omitempty"`

	// +optional

	// EthernetCardOpaqueNetwork contains opaque network backing
	// information for virtual Ethernet cards.
	EthernetCardOpaqueNetwork *VirtualEthernetCardOpaqueNetworkBackingInfo `json:"ethernetCardOpaqueNetwork,omitempty"`

	// +optional

	// File contains host file backing information.
	File *VirtualDeviceFileBackingInfo `json:"file,omitempty"`

	// +optional

	// PCIPassthroughPlugin contains plugin-based PCI passthrough backing
	// information.
	PCIPassthroughPlugin *VirtualPCIPassthroughPluginBackingInfo `json:"pciPassthroughPlugin,omitempty"`

	// +optional

	// Pipe contains named pipe backing information.
	Pipe *VirtualDevicePipeBackingInfo `json:"pipe,omitempty"`

	// +optional

	// PrecisionClock contains system clock backing information for
	// precision clock devices.
	PrecisionClock *VirtualPrecisionClockSystemClockBackingInfo `json:"precisionClock,omitempty"`

	// +optional

	// RemoteDevice contains remote device backing information.
	RemoteDevice *VirtualDeviceRemoteDeviceBackingInfo `json:"remoteDevice,omitempty"`

	// +optional

	// URI contains network URI backing information.
	URI *VirtualDeviceURIBackingInfo `json:"uri,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="[has(self.cdromPassthrough), has(self.diskRawDiskVer2), has(self.ethernetCardNetwork), has(self.pciPassthroughDevice), has(self.pciPassthroughDynamic), has(self.pointingDeviceDevice), has(self.usbRemoteHost)].filter(x, x).size() <= 1",message="at most one of cdromPassthrough, diskRawDiskVer2, ethernetCardNetwork, pciPassthroughDevice, pciPassthroughDynamic, pointingDeviceDevice, or usbRemoteHost may be specified"

// VirtualDeviceDeviceBackingInfo contains information about a host device that
// backs a virtual device.
// It corresponds to vim.vm.device.VirtualDevice.DeviceBackingInfo.
type VirtualDeviceDeviceBackingInfo struct {
	// +optional

	// DeviceName is the name of the device on the host system.
	DeviceName string `json:"deviceName,omitempty"`

	// +optional

	// UseAutoDetect indicates whether the device should be auto-detected
	// instead of directly specified. When true, DeviceName is ignored.
	UseAutoDetect *bool `json:"useAutoDetect,omitempty"`

	// +optional

	CdromPassthrough *VirtualCdromPassthroughBackingInfo `json:"cdromPassthrough,omitempty"`

	// +optional

	DiskRawDiskVer2 *VirtualDiskRawDiskVer2BackingInfo `json:"diskRawDiskVer2,omitempty"`

	// +optional

	EthernetCardNetwork *VirtualEthernetCardNetworkBackingInfo `json:"ethernetCardNetwork,omitempty"`

	// +optional

	PCIPassthroughDevice *VirtualPCIPassthroughDeviceBackingInfo `json:"pciPassthroughDevice,omitempty"`

	// +optional

	PCIPassthroughDynamic *VirtualPCIPassthroughDynamicBackingInfo `json:"pciPassthroughDynamic,omitempty"`

	// +optional

	PointingDeviceDevice *VirtualPointingDeviceDeviceBackingInfo `json:"pointingDeviceDevice,omitempty"`

	// +optional

	USBRemoteHost *VirtualUSBRemoteHostBackingInfo `json:"usbRemoteHost,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="[has(self.cdromRemotePassthrough), has(self.usbRemoteClient)].filter(x, x).size() <= 1",message="at most one of cdromRemotePassthrough or usbRemoteClient may be specified"

// VirtualDeviceRemoteDeviceBackingInfo contains information about a remote
// device that backs a virtual device.
// It corresponds to vim.vm.device.VirtualDevice.RemoteDeviceBackingInfo.
type VirtualDeviceRemoteDeviceBackingInfo struct {
	// DeviceName is the name of the device on the remote system.
	DeviceName string `json:"deviceName"`

	// +optional

	// UseAutoDetect indicates whether the device should be auto-detected
	// instead of directly specified. When true, DeviceName is ignored.
	UseAutoDetect *bool `json:"useAutoDetect,omitempty"`

	// +optional

	CdromRemotePassthrough *VirtualCdromRemotePassthroughBackingInfo `json:"cdromRemotePassthrough,omitempty"`

	// +optional

	USBRemoteClient *VirtualUSBRemoteClientBackingInfo `json:"usbRemoteClient,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="[has(self.nvdimm), has(self.virtualDiskFlatVer1), has(self.virtualDiskFlatVer2), has(self.virtualDiskLocalPMem), has(self.virtualDiskRawDiskMappingVer1), has(self.virtualDiskSeSparse), has(self.virtualDiskSparseVer1), has(self.virtualDiskSparseVer2)].filter(x, x).size() <= 1",message="at most one of nvdimm, virtualDiskFlatVer1, virtualDiskFlatVer2, virtualDiskLocalPMem, virtualDiskRawDiskMappingVer1, virtualDiskSeSparse, virtualDiskSparseVer1, or virtualDiskSparseVer2 may be specified"

// VirtualDeviceFileBackingInfo contains information about a host file that
// backs a virtual device.
// It corresponds to vim.vm.device.VirtualDevice.FileBackingInfo.
type VirtualDeviceFileBackingInfo struct {
	// FileName is the filename for the host file used in this backing.
	FileName string `json:"fileName"`

	// +optional

	// BackingObjectId is the backing object's durable and immutable
	// identifier.
	BackingObjectId *string `json:"backingObjectId,omitempty"`

	// +optional

	NVDIMM *VirtualNVDIMMBackingInfo `json:"nvdimm,omitempty"`

	// +optional

	VirtualDiskFlatVer1 *VirtualDiskFlatVer1BackingInfo `json:"virtualDiskFlatVer1,omitempty"`

	// +optional

	VirtualDiskFlatVer2 *VirtualDiskFlatVer2BackingInfo `json:"virtualDiskFlatVer2,omitempty"`

	// +optional

	VirtualDiskLocalPMem *VirtualDiskLocalPMemBackingInfo `json:"virtualDiskLocalPMem,omitempty"`

	// +optional

	VirtualDiskRawDiskMappingVer1 *VirtualDiskRawDiskMappingVer1BackingInfo `json:"virtualDiskRawDiskMappingVer1,omitempty"`

	// +optional

	VirtualDiskSeSparse *VirtualDiskSeSparseBackingInfo `json:"virtualDiskSeSparse,omitempty"`

	// +optional

	VirtualDiskSparseVer1 *VirtualDiskSparseVer1BackingInfo `json:"virtualDiskSparseVer1,omitempty"`

	// +optional

	VirtualDiskSparseVer2 *VirtualDiskSparseVer2BackingInfo `json:"virtualDiskSparseVer2,omitempty"`
}

// VirtualDevicePipeBackingInfo contains information about a named pipe that
// backs a virtual device.
// It corresponds to vim.vm.device.VirtualDevice.PipeBackingInfo.
type VirtualDevicePipeBackingInfo struct {
	// PipeName is the pipe name for the host pipe associated with this
	// backing.
	PipeName string `json:"pipeName"`

	// +optional

	SerialPort *VirtualSerialPortPipeBackingInfo `json:"serialPort,omitempty"`
}

// VirtualDeviceURIBackingInfo contains information about a network URI that
// backs a virtual device.
// It corresponds to vim.vm.device.VirtualDevice.URIBackingInfo.
type VirtualDeviceURIBackingInfo struct {
	// ServiceURI identifies the local host or a remote system on the network.
	ServiceURI string `json:"serviceURI"`

	// Direction is the connection direction.
	Direction string `json:"direction"`

	// +optional

	// ProxyURI identifies a proxy service providing network access to
	// ServiceURI.
	ProxyURI string `json:"proxyURI,omitempty"`

	// +optional

	SerialPort *VirtualSerialPortURIBackingInfo `json:"serialPort,omitempty"`
}

// VirtualCdromPassthroughBackingInfo defines device pass-through backing for a
// virtual CD-ROM.
// It corresponds to vim.vm.device.VirtualCdrom.PassthroughBackingInfo.
type VirtualCdromPassthroughBackingInfo struct {
	// Exclusive indicates whether the virtual machine has exclusive CD-ROM
	// device access.
	Exclusive bool `json:"exclusive"`
}

// VirtualCdromRemotePassthroughBackingInfo defines remote pass-through device
// backing for a virtual CD-ROM.
// It corresponds to vim.vm.device.VirtualCdrom.RemotePassthroughBackingInfo.
type VirtualCdromRemotePassthroughBackingInfo struct {
	// Exclusive indicates whether the virtual machine has exclusive CD-ROM
	// device access.
	Exclusive bool `json:"exclusive"`
}

// VirtualPCIPassthroughDeviceBackingInfo contains information about the host
// PCI device backing for a PCI passthrough device.
// It corresponds to vim.vm.device.VirtualPCIPassthrough.DeviceBackingInfo.
type VirtualPCIPassthroughDeviceBackingInfo struct {
	// Id is the PCI name ID, composed of "bus:slot.function".
	Id string `json:"id"`

	// DeviceId is the PCI device ID.
	DeviceId string `json:"deviceId"`

	// SystemId is the ID of the system the PCI device is attached to.
	SystemId string `json:"systemId"`

	// VendorId is the PCI vendor ID.
	VendorId int16 `json:"vendorId"`
}

// VirtualPCIPassthroughDynamicBackingInfo contains information about the
// Dynamic DirectPath PCI device backing.
// It corresponds to vim.vm.device.VirtualPCIPassthrough.DynamicBackingInfo.
type VirtualPCIPassthroughDynamicBackingInfo struct {
	// +optional

	// AllowedDevice lists the allowed PCI devices for this Dynamic DirectPath
	// device.
	AllowedDevice []VirtualPCIPassthroughAllowedDevice `json:"allowedDevice,omitempty"`

	// +optional

	// CustomLabel is an optional label that the device must also have set.
	CustomLabel string `json:"customLabel,omitempty"`

	// +optional

	// AssignedId is the ID of the device assigned when the VM is powered on.
	AssignedId string `json:"assignedId,omitempty"`
}

// VirtualPCIPassthroughDvxBackingInfo defines DVX (Device Virtualization
// Extensions) backing for a PCI passthrough device.
// It corresponds to vim.vm.device.VirtualPCIPassthrough.DvxBackingInfo.
type VirtualPCIPassthroughDvxBackingInfo struct {
	// +optional

	// DeviceClass is the device class that backs this DVX device.
	DeviceClass string `json:"deviceClass,omitempty"`

	// +optional

	// ConfigParams contains configuration parameters for this device class.
	ConfigParams []OptionValue `json:"configParams,omitempty"`
}

// VirtualPCIPassthroughPluginBackingInfo is a base backing type for
// plugin-based PCI passthrough devices.
// It corresponds to vim.vm.device.VirtualPCIPassthrough.PluginBackingInfo.
type VirtualPCIPassthroughPluginBackingInfo struct {
	// +optional

	Vmiop *VirtualPCIPassthroughVmiopBackingInfo `json:"vmiop,omitempty"`
}

// VirtualPCIPassthroughVmiopBackingInfo contains information about a VMIOP
// plugin-based PCI passthrough device (typically vGPU).
// It corresponds to vim.vm.device.VirtualPCIPassthrough.VmiopBackingInfo.
type VirtualPCIPassthroughVmiopBackingInfo struct {
	// +optional

	// Vgpu is the vGPU configuration type exposed by the VMIOP plugin.
	Vgpu string `json:"vgpu,omitempty"`

	// +optional

	// VgpuMigrateDataSize is the expected size of the vGPU device state
	// during migration.
	VgpuMigrateDataSize *resource.Quantity `json:"vgpuMigrateDataSize,omitempty"`

	// +optional

	// MigrateSupported indicates whether the vGPU device supports migration.
	MigrateSupported *bool `json:"migrateSupported,omitempty"`

	// +optional

	// EnhancedMigrateCapability indicates whether the vGPU has enhanced
	// migration features for sub-second downtime.
	EnhancedMigrateCapability *bool `json:"enhancedMigrateCapability,omitempty"`
}

// VirtualPointingDeviceDeviceBackingInfo defines host mouse device backing for
// a virtual pointing device.
// It corresponds to vim.vm.device.VirtualPointingDevice.DeviceBackingInfo.
type VirtualPointingDeviceDeviceBackingInfo struct {
	// HostPointingDevice defines the mouse type used to interact with the
	// host mouse.
	HostPointingDevice string `json:"hostPointingDevice"`
}

// VirtualSerialPortPipeBackingInfo defines named pipe backing for a virtual
// serial port.
// It corresponds to vim.vm.device.VirtualSerialPort.PipeBackingInfo.
type VirtualSerialPortPipeBackingInfo struct {
	// Endpoint is the role of the virtual machine as an endpoint for the
	// pipe ("client" or "server").
	Endpoint string `json:"endpoint"`

	// +optional

	// NoRxLoss enables optimized data transfer over the pipe to prevent data
	// overrun.
	NoRxLoss *bool `json:"noRxLoss,omitempty"`
}

// VirtualSerialPortThinPrintBackingInfo defines ThinPrint device backing for
// a virtual serial port.
// It corresponds to vim.vm.device.VirtualSerialPort.ThinPrintBackingInfo.
type VirtualSerialPortThinPrintBackingInfo struct{}

// VirtualSerialPortURIBackingInfo defines network URI backing for a virtual
// serial port.
// It corresponds to vim.vm.device.VirtualSerialPort.URIBackingInfo.
type VirtualSerialPortURIBackingInfo struct {
	VirtualDeviceURIBackingInfo `json:",inline"`
}

// VirtualUSBRemoteClientBackingInfo identifies a USB device on a remote client
// host.
// It corresponds to vim.vm.device.VirtualUSB.RemoteClientBackingInfo.
type VirtualUSBRemoteClientBackingInfo struct {
	// Hostname is the name of the remote client host where the physical USB
	// device resides.
	Hostname string `json:"hostname"`
}

// VirtualUSBRemoteHostBackingInfo identifies a USB device on a specific ESX
// host, supporting persistent access across vMotion.
// It corresponds to vim.vm.device.VirtualUSB.RemoteHostBackingInfo.
type VirtualUSBRemoteHostBackingInfo struct {
	// Hostname is the name of the ESX host to which the physical USB device
	// is attached.
	Hostname string `json:"hostname"`
}

// VirtualEthernetCardNetworkBackingInfo defines standard network backing for
// a virtual Ethernet card.
// It corresponds to vim.vm.device.VirtualEthernetCard.NetworkBackingInfo.
type VirtualEthernetCardNetworkBackingInfo struct {
	// +optional

	Network *ManagedObjectReference `json:"network,omitempty"`
}
// ManagementNetwork describes a management network accessible to virtual
// machines on a host.
type ManagementNetwork struct {
	// +optional

	// Name is the name of the network.
	Name string `json:"name,omitempty"`

	// +optional

	// Type is the type of the network.
	Type string `json:"type,omitempty"`
}

// DistributedVirtualSwitchPortConnection describes a connection to a
// distributed virtual switch port or portgroup.
// It corresponds to vim.dvs.PortConnection.
type DistributedVirtualSwitchPortConnection struct {
	// SwitchUuid is the UUID of the distributed virtual switch.
	SwitchUuid string `json:"switchUuid"`

	// +optional

	// PortgroupKey is the key of the portgroup. Specify this to connect to a
	// portgroup rather than a specific port.
	PortgroupKey string `json:"portgroupKey,omitempty"`

	// +optional

	// PortKey is the key of the specific port. Specify this to connect to a
	// particular port rather than a portgroup.
	PortKey string `json:"portKey,omitempty"`

	// +optional

	// ConnectionCookie is a unique identifier for this port connection
	// instance, assigned by the server.
	ConnectionCookie *int32 `json:"connectionCookie,omitempty"`
}

// VirtualEthernetCardDistributedVirtualPortBackingInfo defines backing for a
// virtual Ethernet card connected to a distributed virtual switch port or
// portgroup.
// It corresponds to vim.vm.device.VirtualEthernetCard.DistributedVirtualPortBackingInfo.
type VirtualEthernetCardDistributedVirtualPortBackingInfo struct {
	// Port is the distributed virtual switch port or portgroup connection.
	Port DistributedVirtualSwitchPortConnection `json:"port"`
}

// VirtualEthernetCardOpaqueNetworkBackingInfo defines backing for a virtual
// Ethernet card connected to an opaque network.
// It corresponds to vim.vm.device.VirtualEthernetCard.OpaqueNetworkBackingInfo.
type VirtualEthernetCardOpaqueNetworkBackingInfo struct {
	// OpaqueNetworkId is the opaque network identifier.
	OpaqueNetworkId string `json:"opaqueNetworkId"`

	// OpaqueNetworkType is the opaque network type.
	OpaqueNetworkType string `json:"opaqueNetworkType"`
}

// VirtualSriovEthernetCardSriovBackingInfo contains information about the
// SR-IOV physical function and virtual function backing for a passthrough NIC.
// It corresponds to vim.vm.device.VirtualSriovEthernetCard.SriovBackingInfo.
type VirtualSriovEthernetCardSriovBackingInfo struct {
	// +optional

	// PhysicalFunctionBacking is the physical function backing for this
	// device.
	PhysicalFunctionBacking *VirtualPCIPassthroughDeviceBackingInfo `json:"physicalFunctionBacking,omitempty"`

	// +optional

	// VirtualFunctionBacking is the virtual function backing for this
	// device.
	VirtualFunctionBacking *VirtualPCIPassthroughDeviceBackingInfo `json:"virtualFunctionBacking,omitempty"`

	// +optional

	// VirtualFunctionIndex is the index of the assigned virtual function.
	VirtualFunctionIndex int32 `json:"virtualFunctionIndex,omitempty"`
}

// VirtualNVDIMMBackingInfo contains information about the file backing for a
// virtual NVDIMM device.
// It corresponds to vim.vm.device.VirtualNVDIMM.BackingInfo.
type VirtualNVDIMMBackingInfo struct {
	// +optional

	// ChangeId is the change ID of the virtual NVDIMM for the corresponding
	// snapshot, used to track incremental changes.
	ChangeId string `json:"changeId,omitempty"`

	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields

	Parent *VirtualNVDIMMBackingInfo `json:"parent,omitempty"`
}

// VirtualPrecisionClockSystemClockBackingInfo contains information about using
// the host system clock as the backing reference clock for a virtual precision
// clock device.
// It corresponds to vim.vm.device.VirtualPrecisionClock.SystemClockBackingInfo.
type VirtualPrecisionClockSystemClockBackingInfo struct {
	// +optional

	// Protocol is the time synchronization protocol used to discipline the
	// system clock (e.g. "ptp", "ntp").
	Protocol string `json:"protocol,omitempty"`
}
