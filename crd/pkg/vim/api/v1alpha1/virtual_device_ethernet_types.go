// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// +kubebuilder:validation:Enum=VirtualE1000e;VirtualE1000e;VirtualPCNet32;VirtualSriovEthernetCard;VirtualVmxnet;VirtualVmxnet2;VirtualVmxnet3;VirtualVmxnet3Vrdma

// EthernetCardType identifies the concrete type of a virtual Ethernet
// card in a virtual machine.
type EthernetCardType string

const (
	// EthernetCardTypeVirtualE1000 is an Intel E1000 Ethernet adapter.
	EthernetCardTypeVirtualE1000 = EthernetCardType(VirtualDeviceTypeE1000)

	// EthernetCardTypeVirtualE1000e is an Intel E1000e Ethernet adapter.
	EthernetCardTypeVirtualE1000e = EthernetCardType(VirtualDeviceTypeE1000e)

	// EthernetCardTypeVirtualPCNet32 is an AMD Lance PCNet32 Ethernet
	// adapter.
	EthernetCardTypeVirtualPCNet32 = EthernetCardType(VirtualDeviceTypePCNet32)

	// EthernetCardTypeVirtualSriovEthernetCard is an SR-IOV enabled
	// Ethernet adapter.
	EthernetCardTypeVirtualSriovEthernetCard = EthernetCardType(VirtualDeviceTypeSriovEthernetCard)

	// EthernetCardTypeVirtualVmxnet is a VMware Vmxnet Ethernet adapter.
	EthernetCardTypeVirtualVmxnet = EthernetCardType(VirtualDeviceTypeVmxnet)

	// EthernetCardTypeVirtualVmxnet2 is a VMware Vmxnet2 Ethernet adapter.
	EthernetCardTypeVirtualVmxnet2 = EthernetCardType(VirtualDeviceTypeVmxnet2)

	// EthernetCardTypeVirtualVmxnet3 is a VMware Vmxnet3 Ethernet adapter.
	EthernetCardTypeVirtualVmxnet3 = EthernetCardType(VirtualDeviceTypeVmxnet3)

	// EthernetCardTypeVirtualVmxnet3Vrdma is a VMware Vmxnet3 VRDMA
	// Ethernet adapter.
	EthernetCardTypeVirtualVmxnet3Vrdma = EthernetCardType(VirtualDeviceTypeVmxnet3Vrdma)
)

// +kubebuilder:validation:XValidation:rule="[has(self.vmxnet3), has(self.sriov)].filter(x, x).size() <= 1",message="at most one of vmxnet3 or sriov may be specified"

// VirtualEthernetCard represents a virtual Ethernet card in a virtual machine.
// It corresponds to vim.vm.device.VirtualEthernetCard.
type VirtualEthernetCard struct {
	// +optional

	// AddressType is the MAC address assignment type.
	// Valid values: "Manual", "Generated", "Assigned".
	AddressType string `json:"addressType,omitempty"`

	// +optional

	// MacAddress is the MAC address assigned to the virtual network adapter.
	MacAddress string `json:"macAddress,omitempty"`

	// +optional

	// WakeOnLanEnabled indicates whether wake-on-LAN is enabled on this
	// virtual network adapter.
	WakeOnLanEnabled *bool `json:"wakeOnLanEnabled,omitempty"`

	// +optional

	// ResourceAllocation describes the network resource requirements of this
	// virtual Ethernet card.
	ResourceAllocation *VirtualEthernetCardResourceAllocation `json:"resourceAllocation,omitempty"`

	// +optional

	// ExternalId is an identifier assigned by an external management plane or
	// controller.
	ExternalId string `json:"externalId,omitempty"`

	// +optional

	// UptCompatibilityEnabled indicates whether UPT (Universal Pass-through)
	// compatibility is enabled on this network adapter.
	//
	// Deprecated: As of vSphere API 8.0, VMDirectPath Gen 2 is no longer
	// supported and there is no replacement.
	UptCompatibilityEnabled *bool `json:"uptCompatibilityEnabled,omitempty"`

	// +optional

	// SubnetId is the ID of the subnet the virtual network adapter connects
	// to. Set only when the adapter is connected to a subnet.
	SubnetId string `json:"subnetId,omitempty"`

	// +optional

	// Vmxnet3 contains Vmxnet3-specific data when this card is a Vmxnet3
	// adapter.
	Vmxnet3 *VirtualVmxnet3 `json:"vmxnet3,omitempty"`

	// +optional

	// Sriov contains SR-IOV-specific data when this card is an SR-IOV
	// Ethernet adapter.
	Sriov *VirtualSriovEthernetCard `json:"sriov,omitempty"`
}

// VirtualEthernetCardResourceAllocation describes the network resource
// requirements of a virtual Ethernet card.
// It corresponds to vim.vm.device.VirtualEthernetCard.ResourceAllocation.
type VirtualEthernetCardResourceAllocation struct {
	// +optional

	// Reservation is the guaranteed network bandwidth in Mbits/sec.
	// Reservation must not exceed Limit when Limit is set.
	Reservation *int64 `json:"reservation,omitempty"`

	// Share describes the relative network bandwidth weight during resource
	// contention.
	Share SharesInfo `json:"share"`

	// +optional

	// Limit is the maximum network bandwidth in Mbits/sec.
	// Set to -1 to indicate no limit.
	Limit *int64 `json:"limit,omitempty"`
}

// VirtualSriovEthernetCard represents an SR-IOV enabled virtual Ethernet
// adapter in a virtual machine.
// It corresponds to vim.vm.device.VirtualSriovEthernetCard.
type VirtualSriovEthernetCard struct {
	// +optional

	// AllowGuestOSMtuChange indicates whether MTU can be changed from the
	// guest OS.
	AllowGuestOSMtuChange *bool `json:"allowGuestOSMtuChange,omitempty"`

	// +optional

	// SriovBacking contains SR-IOV passthrough backing information.
	// Mutually exclusive with DvxBacking.
	SriovBacking *VirtualSriovEthernetCardSriovBackingInfo `json:"sriovBacking,omitempty"`

	// +optional

	// DvxBacking contains DVX backing information for DVX-based SR-IOV
	// devices. Mutually exclusive with SriovBacking.
	DvxBacking *VirtualPCIPassthroughDvxBackingInfo `json:"dvxBacking,omitempty"`
}

// VirtualVmxnet3StrictLatencyConfig contains strict latency configuration for
// a Vmxnet3 adapter.
// It corresponds to vim.vm.device.VirtualVmxnet3StrictLatencyConfig.
type VirtualVmxnet3StrictLatencyConfig struct {
	// +optional

	// Allowed indicates whether strict latency configuration is permitted on
	// this adapter.
	Allowed *bool `json:"allowed,omitempty"`

	// +optional

	// MeasureLatency indicates whether latency measurement is enabled.
	MeasureLatency *bool `json:"measureLatency,omitempty"`

	// +optional

	// MaxTxQueues is the number of transmit queues (1-32).
	MaxTxQueues int32 `json:"maxTxQueues,omitempty"`

	// +optional

	// MaxRxQueues is the number of receive queues (1-32).
	MaxRxQueues int32 `json:"maxRxQueues,omitempty"`

	// +optional

	// TxDataRingDescSize is the transmit data ring descriptor size
	// (128-2048, multiple of 64).
	TxDataRingDescSize int32 `json:"txDataRingDescSize,omitempty"`

	// +optional

	// RxDataRingDescSize is the receive data ring descriptor size
	// (128-2048, multiple of 64).
	RxDataRingDescSize int32 `json:"rxDataRingDescSize,omitempty"`

	// +optional

	// DisableOffload is the type of offload disable operation.
	DisableOffload string `json:"disableOffload,omitempty"`
}

// VirtualVmxnet3 represents a Vmxnet3 virtual Ethernet adapter in a virtual
// machine.
// It corresponds to vim.vm.device.VirtualVmxnet3.
type VirtualVmxnet3 struct {
	// +optional

	// Uptv2Enabled indicates whether UPTv2 (Uniform Pass-through version 2)
	// compatibility is enabled on this network adapter.
	Uptv2Enabled *bool `json:"uptv2Enabled,omitempty"`

	// +optional

	// StrictLatencyConfig contains strict latency configuration parameters
	// for this adapter.
	StrictLatencyConfig *VirtualVmxnet3StrictLatencyConfig `json:"strictLatencyConfig,omitempty"`

	// +optional

	// Vrdma contains VRDMA-specific data when this adapter is a Vmxnet3
	// VRDMA adapter.
	Vrdma *VirtualVmxnet3Vrdma `json:"vrdma,omitempty"`
}

// VirtualVmxnet3Vrdma represents a Vmxnet3 VRDMA virtual Ethernet adapter in
// a virtual machine.
// It corresponds to vim.vm.device.VirtualVmxnet3Vrdma.
type VirtualVmxnet3Vrdma struct {
	// +optional

	// DeviceProtocol is the VRDMA device protocol.
	DeviceProtocol string `json:"deviceProtocol,omitempty"`
}
