// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"github.com/vmware/govmomi/vim25/types"
)

/*
Source:
  - https://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_ResourceAllocationSettingData.xsd
  - https://schemas.dmtf.org/wbem/cim-html/2/CIM_ResourceAllocationSettingData.html
*/
type CIMResourceType uint16

// Please note, the iota pattern is not used to ensure these constants remain
// affixed to an explicit value.
const (
	Other              CIMResourceType = 1
	ComputerSystem     CIMResourceType = 2
	Processor          CIMResourceType = 3
	Memory             CIMResourceType = 4
	IdeController      CIMResourceType = 5
	ParallelScsiHba    CIMResourceType = 6
	FcHba              CIMResourceType = 7
	IScsiHba           CIMResourceType = 8
	IbHba              CIMResourceType = 9
	EthernetAdapter    CIMResourceType = 10
	OtherNetwork       CIMResourceType = 11
	IoSlot             CIMResourceType = 12
	IoDevice           CIMResourceType = 13
	FloppyDrive        CIMResourceType = 14
	CdDrive            CIMResourceType = 15
	DvdDrive           CIMResourceType = 16
	DiskDrive          CIMResourceType = 17
	TapeDrive          CIMResourceType = 18
	StorageExtent      CIMResourceType = 19
	OtherStorage       CIMResourceType = 20
	SerialPort         CIMResourceType = 21
	ParallelPort       CIMResourceType = 22
	UsbController      CIMResourceType = 23
	Graphics           CIMResourceType = 24
	Ieee1394           CIMResourceType = 25
	PartitionableUnit  CIMResourceType = 26
	BasePartitionable  CIMResourceType = 27
	PowerSupply        CIMResourceType = 28
	CoolingDevice      CIMResourceType = 29
	EthernetSwitchPort CIMResourceType = 30
	LogicalDisk        CIMResourceType = 31
	StorageVolume      CIMResourceType = 32
	EthernetConnection CIMResourceType = 33
)

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_VirtualSystemSettingData.xsd
*/

type CIMVirtualSystemSettingData struct {
	ElementName string `xml:"ElementName" json:"elementName"`
	InstanceID  string `xml:"InstanceID" json:"instanceID"`

	AutomaticRecoveryAction              *uint8   `xml:"AutomaticRecoveryAction" json:"automaticRecoveryAction,omitempty"`
	AutomaticShutdownAction              *uint8   `xml:"AutomaticShutdownAction" json:"automaticShutdownAction,omitempty"`
	AutomaticStartupAction               *uint8   `xml:"AutomaticStartupAction" json:"automaticStartupAction,omitempty"`
	AutomaticStartupActionDelay          *string  `xml:"AutomaticStartupActionDelay>Interval" json:"automaticStartupActionDelay,omitempty"`
	AutomaticStartupActionSequenceNumber *uint16  `xml:"AutomaticStartupActionSequenceNumber" json:"automaticStartupActionSequenceNumber,omitempty"`
	Caption                              *string  `xml:"Caption" json:"caption,omitempty"`
	ConfigurationDataRoot                *string  `xml:"ConfigurationDataRoot" json:"configurationDataRoot,omitempty"`
	ConfigurationFile                    *string  `xml:"ConfigurationFile" json:"configurationFile,omitempty"`
	ConfigurationID                      *string  `xml:"ConfigurationID" json:"configurationID,omitempty"`
	CreationTime                         *string  `xml:"CreationTime" json:"creationTime,omitempty"`
	Description                          *string  `xml:"Description" json:"description,omitempty"`
	LogDataRoot                          *string  `xml:"LogDataRoot" json:"logDataRoot,omitempty"`
	Notes                                []string `xml:"Notes" json:"notes,omitempty"`
	RecoveryFile                         *string  `xml:"RecoveryFile" json:"recoveryFile,omitempty"`
	SnapshotDataRoot                     *string  `xml:"SnapshotDataRoot" json:"snapshotDataRoot,omitempty"`
	SuspendDataRoot                      *string  `xml:"SuspendDataRoot" json:"suspendDataRoot,omitempty"`
	SwapFileDataRoot                     *string  `xml:"SwapFileDataRoot" json:"swapFileDataRoot,omitempty"`
	VirtualSystemIdentifier              *string  `xml:"VirtualSystemIdentifier" json:"virtualSystemIdentifier,omitempty"`
	VirtualSystemType                    *string  `xml:"VirtualSystemType" json:"virtualSystemType,omitempty"`
}

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_ResourceAllocationSettingData.xsd
*/

type CIMResourceAllocationSettingData struct {
	ElementName string `xml:"ElementName" json:"elementName,omitempty"`
	InstanceID  string `xml:"InstanceID" json:"instanceID,omitempty"`

	ResourceType      *CIMResourceType `xml:"ResourceType" json:"resourceType,omitempty"`
	OtherResourceType *string          `xml:"OtherResourceType" json:"otherResourceType,omitempty"`
	ResourceSubType   *string          `xml:"ResourceSubType" json:"resourceSubType,omitempty"`

	AddressOnParent       *string  `xml:"AddressOnParent" json:"addressOnParent,omitempty"`
	Address               *string  `xml:"Address" json:"address,omitempty"`
	AllocationUnits       *string  `xml:"AllocationUnits" json:"allocationUnits,omitempty"`
	AutomaticAllocation   *bool    `xml:"AutomaticAllocation" json:"automaticAllocation,omitempty"`
	AutomaticDeallocation *bool    `xml:"AutomaticDeallocation" json:"automaticDeallocation,omitempty"`
	Caption               *string  `xml:"Caption" json:"caption,omitempty"`
	Connection            []string `xml:"Connection" json:"connection,omitempty"`
	ConsumerVisibility    *uint16  `xml:"ConsumerVisibility" json:"consumerVisibility,omitempty"`
	Description           *string  `xml:"Description" json:"description,omitempty"`
	HostResource          []string `xml:"HostResource" json:"hostResource,omitempty"`
	Limit                 *uint64  `xml:"Limit" json:"limit,omitempty"`
	MappingBehavior       *uint    `xml:"MappingBehavior" json:"mappingBehavior,omitempty"`
	Parent                *string  `xml:"Parent" json:"parent,omitempty"`
	PoolID                *string  `xml:"PoolID" json:"poolID,omitempty"`
	Reservation           *uint64  `xml:"Reservation" json:"reservation,omitempty"`
	VirtualQuantity       *uint    `xml:"VirtualQuantity" json:"virtualQuantity,omitempty"`
	VirtualQuantityUnits  *string  `xml:"VirtualQuantityUnits" json:"virtualQuantityUnits,omitempty"`
	Weight                *uint    `xml:"Weight" json:"weight,omitempty"`
}

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_StorageAllocationSettingData.xsd
*/
type CIMStorageAllocationSettingData struct {
	ElementName string `xml:"ElementName" json:"elementName"`
	InstanceID  string `xml:"InstanceID" json:"instanceID"`

	ResourceType      *CIMResourceType `xml:"ResourceType" json:"resourceType,omitempty"`
	OtherResourceType *string          `xml:"OtherResourceType" json:"otherResourceType,omitempty"`
	ResourceSubType   *string          `xml:"ResourceSubType" json:"resourceSubType,omitempty"`

	Access                       *uint16         `xml:"Access" json:"access,omitempty"`
	Address                      *string         `xml:"Address" json:"address,omitempty"`
	AddressOnParent              *string         `xml:"AddressOnParent" json:"addressOnParent,omitempty"`
	AllocationUnits              *string         `xml:"AllocationUnits" json:"allocationUnits,omitempty"`
	AutomaticAllocation          *bool           `xml:"AutomaticAllocation" json:"automaticAllocation,omitempty"`
	AutomaticDeallocation        *bool           `xml:"AutomaticDeallocation" json:"automaticDeallocation,omitempty"`
	Caption                      *string         `xml:"Caption" json:"caption,omitempty"`
	ChangeableType               *uint16         `xml:"ChangeableType" json:"changeableType,omitempty"`
	ComponentSetting             []types.AnyType `xml:"ComponentSetting" json:"componentSetting,omitempty"`
	ConfigurationName            *string         `xml:"ConfigurationName" json:"configurationName,omitempty"`
	Connection                   []string        `xml:"Connection" json:"connection,omitempty"`
	ConsumerVisibility           *uint16         `xml:"ConsumerVisibility" json:"consumerVisibility,omitempty"`
	Description                  *string         `xml:"Description" json:"description,omitempty"`
	Generation                   *uint64         `xml:"Generation" json:"generation,omitempty"`
	HostExtentName               *string         `xml:"HostExtentName" json:"hostExtentName,omitempty"`
	HostExtentNameFormat         *uint16         `xml:"HostExtentNameFormat" json:"hostExtentNameFormat,omitempty"`
	HostExtentNameNamespace      *uint16         `xml:"HostExtentNameNamespace" json:"hostExtentNameNamespace,omitempty"`
	HostExtentStartingAddress    *uint64         `xml:"HostExtentStartingAddress" json:"hostExtentStartingAddress,omitempty"`
	HostResource                 []string        `xml:"HostResource" json:"hostResource,omitempty"`
	HostResourceBlockSize        *uint64         `xml:"HostResourceBlockSize" json:"hostResourceBlockSize,omitempty"`
	Limit                        *uint64         `xml:"Limit" json:"limit,omitempty"`
	MappingBehavior              *uint           `xml:"MappingBehavior" json:"mappingBehavior,omitempty"`
	OtherHostExtentNameFormat    *string         `xml:"OtherHostExtentNameFormat" json:"otherHostExtentNameFormat,omitempty"`
	OtherHostExtentNameNamespace *string         `xml:"OtherHostExtentNameNamespace" json:"otherHostExtentNameNamespace,omitempty"`
	Parent                       *string         `xml:"Parent" json:"parent,omitempty"`
	PoolID                       *string         `xml:"PoolID" json:"poolID,omitempty"`
	Reservation                  *uint64         `xml:"Reservation" json:"reservation,omitempty"`
	SoID                         *string         `xml:"SoID" json:"soID,omitempty"`
	SoOrgID                      *string         `xml:"SoOrgID" json:"soOrgID,omitempty"`
	VirtualQuantity              *uint           `xml:"VirtualQuantity" json:"virtualQuantity,omitempty"`
	VirtualQuantityUnits         *string         `xml:"VirtualQuantityUnits" json:"virtualQuantityUnits,omitempty"`
	VirtualResourceBlockSize     *uint64         `xml:"VirtualResourceBlockSize" json:"virtualResourceBlockSize,omitempty"`
	Weight                       *uint           `xml:"Weight" json:"weight,omitempty"`
}
