// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"fmt"

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

// CIMOSType represents the CIM (Common Information Model) OSType enumeration
// values. These values are defined by the DMTF CIM schema and used in various
// management standards. Please refer to the following URL for more information:
// https://learn.microsoft.com/en-us/windows/win32/cimwin32prov/cim-operatingsystem
type CIMOSType int

const (
	CIMOSTypeUnknown                          CIMOSType = 0
	CIMOSTypeOther                            CIMOSType = 1
	CIMOSTypeMACOS                            CIMOSType = 2
	CIMOSTypeATTUNIX                          CIMOSType = 3
	CIMOSTypeDGUX                             CIMOSType = 4
	CIMOSTypeDECNT                            CIMOSType = 5
	CIMOSTypeTru64UNIX                        CIMOSType = 6
	CIMOSTypeOpenVMS                          CIMOSType = 7
	CIMOSTypeHPUX                             CIMOSType = 8
	CIMOSTypeAIX                              CIMOSType = 9
	CIMOSTypeMVS                              CIMOSType = 10
	CIMOSTypeOS400                            CIMOSType = 11
	CIMOSTypeOS2                              CIMOSType = 12
	CIMOSTypeJavaMachine                      CIMOSType = 13
	CIMOSTypeMSDOS                            CIMOSType = 14
	CIMOSTypeWIN3x                            CIMOSType = 15
	CIMOSTypeWIN95                            CIMOSType = 16
	CIMOSTypeWIN98                            CIMOSType = 17
	CIMOSTypeWINNT                            CIMOSType = 18
	CIMOSTypeWINCE                            CIMOSType = 19
	CIMOSTypeNCR3000                          CIMOSType = 20
	CIMOSTypeNetWare                          CIMOSType = 21
	CIMOSTypeOSF                              CIMOSType = 22
	CIMOSTypeDCOS                             CIMOSType = 23
	CIMOSTypeReliantUNIX                      CIMOSType = 24
	CIMOSTypeSCOUnixWare                      CIMOSType = 25
	CIMOSTypeSCOOpenServer                    CIMOSType = 26
	CIMOSTypeSequent                          CIMOSType = 27
	CIMOSTypeIRIX                             CIMOSType = 28
	CIMOSTypeSolaris                          CIMOSType = 29
	CIMOSTypeSunOS                            CIMOSType = 30
	CIMOSTypeU6000                            CIMOSType = 31
	CIMOSTypeASERIES                          CIMOSType = 32
	CIMOSTypeTandemNSK                        CIMOSType = 33
	CIMOSTypeTandemNT                         CIMOSType = 34
	CIMOSTypeBS2000                           CIMOSType = 35
	CIMOSTypeLINUX                            CIMOSType = 36
	CIMOSTypeLynx                             CIMOSType = 37
	CIMOSTypeXENIX                            CIMOSType = 38
	CIMOSTypeVMESA                            CIMOSType = 39
	CIMOSTypeInteractive                      CIMOSType = 40
	CIMOSTypeBSDUNIX                          CIMOSType = 41
	CIMOSTypeFreeBSD                          CIMOSType = 42
	CIMOSTypeNetBSD                           CIMOSType = 43
	CIMOSTypeGNUHurd                          CIMOSType = 44
	CIMOSTypeOS9                              CIMOSType = 45
	CIMOSTypeMACHKernel                       CIMOSType = 46
	CIMOSTypeInferno                          CIMOSType = 47
	CIMOSTypeQNX                              CIMOSType = 48
	CIMOSTypeEPOC                             CIMOSType = 49
	CIMOSTypeIXWorks                          CIMOSType = 50
	CIMOSTypeVxWorks                          CIMOSType = 51
	CIMOSTypeMiNT                             CIMOSType = 52
	CIMOSTypeBeOS                             CIMOSType = 53
	CIMOSTypeHPMPE                            CIMOSType = 54
	CIMOSTypeNextStep                         CIMOSType = 55
	CIMOSTypePalmPilot                        CIMOSType = 56
	CIMOSTypeRhapsody                         CIMOSType = 57
	CIMOSTypeWindows2000                      CIMOSType = 58
	CIMOSTypeDedicated                        CIMOSType = 59
	CIMOSTypeOS390                            CIMOSType = 60
	CIMOSTypeVSE                              CIMOSType = 61
	CIMOSTypeTPF                              CIMOSType = 62
	CIMOSTypeWindowsMe                        CIMOSType = 63
	CIMOSTypeCalderaOpenUNIX                  CIMOSType = 64
	CIMOSTypeOpenBSD                          CIMOSType = 65
	CIMOSTypeNotApplicable                    CIMOSType = 66
	CIMOSTypeWindowsXP                        CIMOSType = 67
	CIMOSTypezOS                              CIMOSType = 68
	CIMOSTypeMicrosoftWindowsServer2003       CIMOSType = 69
	CIMOSTypeMicrosoftWindowsServer2003_64Bit CIMOSType = 70
	CIMOSTypeWindowsXP64Bit                   CIMOSType = 71
	CIMOSTypeWindowsXPEmbedded                CIMOSType = 72
	CIMOSTypeWindowsVista                     CIMOSType = 73
	CIMOSTypeWindowsVista64Bit                CIMOSType = 74
	CIMOSTypeWindowsEmbeddedforPointofService CIMOSType = 75
	CIMOSTypeMicrosoftWindowsServer2008       CIMOSType = 76
	CIMOSTypeMicrosoftWindowsServer2008_64Bit CIMOSType = 77
	CIMOSTypeFreeBSD64Bit                     CIMOSType = 78
	CIMOSTypeRedHatEnterpriseLinux            CIMOSType = 79
	CIMOSTypeRedHatEnterpriseLinux64Bit       CIMOSType = 80
	CIMOSTypeSolaris64Bit                     CIMOSType = 81
	CIMOSTypeSUSE                             CIMOSType = 82
	CIMOSTypeSUSE64Bit                        CIMOSType = 83
	CIMOSTypeSLES                             CIMOSType = 84
	CIMOSTypeSLES64Bit                        CIMOSType = 85
	CIMOSTypeNovellOES                        CIMOSType = 86
	CIMOSTypeNovellLinuxDesktop               CIMOSType = 87
	CIMOSTypeSunJavaDesktopSystem             CIMOSType = 88
	CIMOSTypeMandriva                         CIMOSType = 89
	CIMOSTypeMandriva64Bit                    CIMOSType = 90
	CIMOSTypeTurboLinux                       CIMOSType = 91
	CIMOSTypeTurboLinux64Bit                  CIMOSType = 92
	CIMOSTypeUbuntu                           CIMOSType = 93
	CIMOSTypeUbuntu64Bit                      CIMOSType = 94
	CIMOSTypeDebian                           CIMOSType = 95
	CIMOSTypeDebian64Bit                      CIMOSType = 96
	CIMOSTypeLinux24x                         CIMOSType = 97
	CIMOSTypeLinux24x64Bit                    CIMOSType = 98
	CIMOSTypeLinux26x                         CIMOSType = 99
	CIMOSTypeLinux26x64Bit                    CIMOSType = 100
	CIMOSTypeLinux64Bit                       CIMOSType = 101
	CIMOSTypeOther64Bit                       CIMOSType = 102
	CIMOSTypeMicrosoftWindowsServer2008R2     CIMOSType = 103
	CIMOSTypeVMwareESXi                       CIMOSType = 104
	CIMOSTypeMicrosoftWindows7                CIMOSType = 105
	CIMOSTypeCentOS32bit                      CIMOSType = 106
	CIMOSTypeCentOS64bit                      CIMOSType = 107
	CIMOSTypeOracle32bit                      CIMOSType = 108
	CIMOSTypeOracle64bit                      CIMOSType = 109
	CIMOSTypeeComStation32bitx                CIMOSType = 110
	CIMOSTypeMicrosoftWindowsServer2011       CIMOSType = 111
	CIMOSTypeMicrosoftWindowsServer2012       CIMOSType = 112
	CIMOSTypeMicrosoftWindows8                CIMOSType = 113
	CIMOSTypeMicrosoftWindows81               CIMOSType = 114
	CIMOSTypeMicrosoftWindowsServer2012R2     CIMOSType = 115
)

// cimGuestIDEntry maps a (cimId, cimVersion) pair to a vSphere GuestId.
// Entries keyed as "cimId" are bare (no version); entries keyed as
// "cimId-version" are version-specific.
type cimGuestIDEntry struct {
	cimID   int
	version string
	guestID types.VirtualMachineGuestOsIdentifier
}

// cimGuestIDTable maps CIM OS type integers (from the DMTF CIM schema) and
// optional version strings to vSphere GuestId identifiers. The lookup
// algorithm (CIMOSTypeToGuestID) first tries (cimId, version), then falls
// back to (cimId, "").
var cimGuestIDTable = []cimGuestIDEntry{
	// macOS
	{2, "", types.VirtualMachineGuestOsIdentifierDarwinGuest},
	{2, "10", types.VirtualMachineGuestOsIdentifierDarwin10Guest},
	{2, "11", types.VirtualMachineGuestOsIdentifierDarwin11Guest},
	{2, "19", types.VirtualMachineGuestOsIdentifierDarwin19_64Guest},
	{2, "20", types.VirtualMachineGuestOsIdentifierDarwin20_64Guest},
	{2, "21", types.VirtualMachineGuestOsIdentifierDarwin21_64Guest},
	{2, "22", types.VirtualMachineGuestOsIdentifierDarwin22_64Guest},
	{2, "23", types.VirtualMachineGuestOsIdentifierDarwin23_64Guest},

	// Debian (all versioned; no bare entry → bare cimId returns "")
	{95, "4", types.VirtualMachineGuestOsIdentifierDebian4Guest},
	{96, "4", types.VirtualMachineGuestOsIdentifierDebian4_64Guest},
	{95, "5", types.VirtualMachineGuestOsIdentifierDebian5Guest},
	{96, "5", types.VirtualMachineGuestOsIdentifierDebian5_64Guest},
	{95, "6", types.VirtualMachineGuestOsIdentifierDebian6Guest},
	{96, "6", types.VirtualMachineGuestOsIdentifierDebian6_64Guest},
	{95, "7", types.VirtualMachineGuestOsIdentifierDebian7Guest},
	{96, "7", types.VirtualMachineGuestOsIdentifierDebian7_64Guest},
	{95, "8", types.VirtualMachineGuestOsIdentifierDebian8Guest},
	{96, "8", types.VirtualMachineGuestOsIdentifierDebian8_64Guest},
	{95, "9", types.VirtualMachineGuestOsIdentifierDebian9Guest},
	{96, "9", types.VirtualMachineGuestOsIdentifierDebian9_64Guest},
	{95, "10", types.VirtualMachineGuestOsIdentifierDebian10Guest},
	{96, "10", types.VirtualMachineGuestOsIdentifierDebian10_64Guest},
	{95, "11", types.VirtualMachineGuestOsIdentifierDebian11Guest},
	{96, "11", types.VirtualMachineGuestOsIdentifierDebian11_64Guest},
	{95, "12", types.VirtualMachineGuestOsIdentifierDebian12Guest},
	{96, "12", types.VirtualMachineGuestOsIdentifierDebian12_64Guest},
	{95, "13", types.VirtualMachineGuestOsIdentifierDebian13Guest},
	{96, "13", types.VirtualMachineGuestOsIdentifierDebian13_64Guest},
	{96, "1", types.VirtualMachineGuestOsIdentifierPardus_64Guest},

	// DOS
	{14, "", types.VirtualMachineGuestOsIdentifierDosGuest},

	// FreeBSD
	{42, "", types.VirtualMachineGuestOsIdentifierFreebsdGuest},
	{78, "", types.VirtualMachineGuestOsIdentifierFreebsd64Guest},
	{42, "11", types.VirtualMachineGuestOsIdentifierFreebsd11Guest},
	{78, "11", types.VirtualMachineGuestOsIdentifierFreebsd11_64Guest},
	{42, "12", types.VirtualMachineGuestOsIdentifierFreebsd12Guest},
	{78, "12", types.VirtualMachineGuestOsIdentifierFreebsd12_64Guest},
	{42, "13", types.VirtualMachineGuestOsIdentifierFreebsd13Guest},
	{78, "13", types.VirtualMachineGuestOsIdentifierFreebsd13_64Guest},
	{42, "14", types.VirtualMachineGuestOsIdentifierFreebsd14Guest},
	{78, "14", types.VirtualMachineGuestOsIdentifierFreebsd14_64Guest},
	{42, "15", types.VirtualMachineGuestOsIdentifierFreebsd15Guest},
	{78, "15", types.VirtualMachineGuestOsIdentifierFreebsd15_64Guest},

	// Mandriva
	{89, "", types.VirtualMachineGuestOsIdentifierMandrivaGuest},
	{90, "", types.VirtualMachineGuestOsIdentifierMandriva64Guest},

	// NetWare (all versioned; no bare entry → bare cimId returns "")
	{21, "4", types.VirtualMachineGuestOsIdentifierNetware4Guest},
	{21, "5", types.VirtualMachineGuestOsIdentifierNetware5Guest},
	{21, "6", types.VirtualMachineGuestOsIdentifierNetware6Guest},

	// Novell
	{87, "", types.VirtualMachineGuestOsIdentifierNld9Guest},
	{86, "", types.VirtualMachineGuestOsIdentifierOesGuest},

	// SCO
	{26, "5", types.VirtualMachineGuestOsIdentifierOpenServer5Guest},
	{26, "6", types.VirtualMachineGuestOsIdentifierOpenServer6Guest},

	// OS/2
	{12, "", types.VirtualMachineGuestOsIdentifierOs2Guest},

	// Linux generic
	{97, "", types.VirtualMachineGuestOsIdentifierOther24xLinuxGuest},
	{98, "", types.VirtualMachineGuestOsIdentifierOther24xLinux64Guest},
	{99, "", types.VirtualMachineGuestOsIdentifierOther26xLinuxGuest},
	{100, "", types.VirtualMachineGuestOsIdentifierOther26xLinux64Guest},
	{1, "", types.VirtualMachineGuestOsIdentifierOtherGuest},
	{101, "", types.VirtualMachineGuestOsIdentifierOtherLinux64Guest},
	{102, "", types.VirtualMachineGuestOsIdentifierOtherGuest64},
	{36, "", types.VirtualMachineGuestOsIdentifierOtherLinuxGuest},

	// RHEL (bare entry for 32-bit; no bare entry for 64-bit → versioned only)
	{79, "", types.VirtualMachineGuestOsIdentifierRedhatGuest},
	{79, "2", types.VirtualMachineGuestOsIdentifierRhel2Guest},
	{79, "3", types.VirtualMachineGuestOsIdentifierRhel3Guest},
	{80, "3", types.VirtualMachineGuestOsIdentifierRhel3_64Guest},
	{79, "4", types.VirtualMachineGuestOsIdentifierRhel4Guest},
	{80, "4", types.VirtualMachineGuestOsIdentifierRhel4_64Guest},
	{79, "5", types.VirtualMachineGuestOsIdentifierRhel5Guest},
	{80, "5", types.VirtualMachineGuestOsIdentifierRhel5_64Guest},
	{79, "6", types.VirtualMachineGuestOsIdentifierRhel6Guest},
	{80, "6", types.VirtualMachineGuestOsIdentifierRhel6_64Guest},
	{79, "7", types.VirtualMachineGuestOsIdentifierRhel7Guest},
	{80, "7", types.VirtualMachineGuestOsIdentifierRhel7_64Guest},
	{80, "8", types.VirtualMachineGuestOsIdentifierRhel8_64Guest},
	{80, "9", types.VirtualMachineGuestOsIdentifierRhel9_64Guest},
	{80, "10", types.VirtualMachineGuestOsIdentifierRhel10_64Guest},

	// Sun Java Desktop
	{88, "", types.VirtualMachineGuestOsIdentifierSjdsGuest},

	// SLES
	{84, "", types.VirtualMachineGuestOsIdentifierSlesGuest},
	{85, "", types.VirtualMachineGuestOsIdentifierSles64Guest},
	{84, "10", types.VirtualMachineGuestOsIdentifierSles10Guest},
	{85, "10", types.VirtualMachineGuestOsIdentifierSles10_64Guest},
	{84, "11", types.VirtualMachineGuestOsIdentifierSles11Guest},
	{85, "11", types.VirtualMachineGuestOsIdentifierSles11_64Guest},
	{84, "12", types.VirtualMachineGuestOsIdentifierSles12Guest},
	{85, "12", types.VirtualMachineGuestOsIdentifierSles12_64Guest},
	{85, "15", types.VirtualMachineGuestOsIdentifierSles15_64Guest},
	{85, "16", types.VirtualMachineGuestOsIdentifierSles16_64Guest},

	// SUSE
	{82, "", types.VirtualMachineGuestOsIdentifierSuseGuest},
	{83, "", types.VirtualMachineGuestOsIdentifierSuse64Guest},
	{82, "11", types.VirtualMachineGuestOsIdentifierOpensuseGuest},
	{83, "11", types.VirtualMachineGuestOsIdentifierOpensuse64Guest},

	// Solaris (all versioned; no bare entry → bare cimId returns "")
	{29, "6", types.VirtualMachineGuestOsIdentifierSolaris6Guest},
	{29, "7", types.VirtualMachineGuestOsIdentifierSolaris7Guest},
	{29, "8", types.VirtualMachineGuestOsIdentifierSolaris8Guest},
	{29, "9", types.VirtualMachineGuestOsIdentifierSolaris9Guest},
	{29, "10", types.VirtualMachineGuestOsIdentifierSolaris10Guest},
	{81, "10", types.VirtualMachineGuestOsIdentifierSolaris10_64Guest},
	{81, "11", types.VirtualMachineGuestOsIdentifierSolaris11_64Guest},

	// TurboLinux
	{91, "", types.VirtualMachineGuestOsIdentifierTurboLinuxGuest},
	{92, "", types.VirtualMachineGuestOsIdentifierTurboLinux64Guest},

	// Ubuntu
	{93, "", types.VirtualMachineGuestOsIdentifierUbuntuGuest},
	{94, "", types.VirtualMachineGuestOsIdentifierUbuntu64Guest},

	// SCO UnixWare
	{25, "7", types.VirtualMachineGuestOsIdentifierUnixWare7Guest},

	// Windows Desktop (legacy)
	{15, "", types.VirtualMachineGuestOsIdentifierWin31Guest},
	{16, "", types.VirtualMachineGuestOsIdentifierWin95Guest},
	{17, "", types.VirtualMachineGuestOsIdentifierWin98Guest},
	{18, "", types.VirtualMachineGuestOsIdentifierWinNTGuest},
	{63, "", types.VirtualMachineGuestOsIdentifierWinMeGuest},
	{67, "", types.VirtualMachineGuestOsIdentifierWinXPProGuest},
	{71, "", types.VirtualMachineGuestOsIdentifierWinXPPro64Guest},
	{73, "", types.VirtualMachineGuestOsIdentifierWinVistaGuest},
	{74, "", types.VirtualMachineGuestOsIdentifierWinVista64Guest},
	{105, "", types.VirtualMachineGuestOsIdentifierWindows7Guest},

	// Windows Server
	// win2000ServGuest is the only non-starred bare entry for cimId=58
	{58, "", types.VirtualMachineGuestOsIdentifierWin2000ServGuest},
	// winNetStandardGuest is the only non-starred bare entry for cimId=69/70
	{69, "", types.VirtualMachineGuestOsIdentifierWinNetStandardGuest},
	{70, "", types.VirtualMachineGuestOsIdentifierWinNetStandard64Guest},
	// WinServer2008 (76/77) entries are all starred → no bare entry → returns ""
	{103, "", types.VirtualMachineGuestOsIdentifierWindows7Server64Guest},
	{112, "", types.VirtualMachineGuestOsIdentifierWindows8Server64Guest},
	{115, "", types.VirtualMachineGuestOsIdentifierWindows9Server64Guest},

	// CentOS (bare entry maps to generic centos/centos64, not versioned centos7)
	{106, "", types.VirtualMachineGuestOsIdentifierCentosGuest},
	{107, "", types.VirtualMachineGuestOsIdentifierCentos64Guest},
	{106, "6", types.VirtualMachineGuestOsIdentifierCentos6Guest},
	{107, "6", types.VirtualMachineGuestOsIdentifierCentos6_64Guest},
	{106, "7", types.VirtualMachineGuestOsIdentifierCentos7Guest},
	{107, "7", types.VirtualMachineGuestOsIdentifierCentos7_64Guest},
	{107, "8", types.VirtualMachineGuestOsIdentifierCentos8_64Guest},
	{107, "9", types.VirtualMachineGuestOsIdentifierCentos9_64Guest},

	// Oracle Linux
	{108, "", types.VirtualMachineGuestOsIdentifierOracleLinuxGuest},
	{109, "", types.VirtualMachineGuestOsIdentifierOracleLinux64Guest},
	{108, "6", types.VirtualMachineGuestOsIdentifierOracleLinux6Guest},
	{109, "6", types.VirtualMachineGuestOsIdentifierOracleLinux6_64Guest},
	{108, "7", types.VirtualMachineGuestOsIdentifierOracleLinux7Guest},
	{109, "7", types.VirtualMachineGuestOsIdentifierOracleLinux7_64Guest},
	{109, "8", types.VirtualMachineGuestOsIdentifierOracleLinux8_64Guest},
	{109, "9", types.VirtualMachineGuestOsIdentifierOracleLinux9_64Guest},
	{109, "10", types.VirtualMachineGuestOsIdentifierOracleLinux10_64Guest},

	// eComStation
	{110, "", types.VirtualMachineGuestOsIdentifierEComStationGuest},
	{110, "2", types.VirtualMachineGuestOsIdentifierEComStation2Guest},

	// VMware ESXi (all versioned; no bare entry → bare cimId returns "")
	{104, "4", types.VirtualMachineGuestOsIdentifierVmkernelGuest},
	{104, "5", types.VirtualMachineGuestOsIdentifierVmkernel5Guest},
	{104, "6", types.VirtualMachineGuestOsIdentifierVmkernel6Guest},
	{104, "65", types.VirtualMachineGuestOsIdentifierVmkernel65Guest},
	{104, "7", types.VirtualMachineGuestOsIdentifierVmkernel7Guest},
	{104, "8", types.VirtualMachineGuestOsIdentifierVmkernel8Guest},
	{104, "9", types.VirtualMachineGuestOsIdentifierVmkernel9Guest},
}

// cimGuestIDIndex is a two-level index built from cimGuestIDTable.
// Keys are "cimId" (bare) or "cimId-version" (versioned).
var cimGuestIDIndex map[string]types.VirtualMachineGuestOsIdentifier

func init() {
	cimGuestIDIndex = make(map[string]types.VirtualMachineGuestOsIdentifier, len(cimGuestIDTable))
	for _, e := range cimGuestIDTable {
		key := cimGuestIDKey(e.cimID, e.version)
		cimGuestIDIndex[key] = e.guestID
	}
}

func cimGuestIDKey(cimID int, version string) string {
	if version == "" {
		return fmt.Sprintf("%d", cimID)
	}
	return fmt.Sprintf("%d-%s", cimID, version)
}

// CIMOSTypeToGuestID translates a CIM OSType value and optional version string
// to a vSphere Guest OS identifier. It first tries the exact (cimId, version)
// pair, then falls back to the bare cimId with no version. Returns an empty
// string when no mapping exists (e.g. CIMOSTypeUnknown, WINCE, NetWare without
// a version, WinServer2008, or other CIM types with no vSphere equivalent).
func CIMOSTypeToGuestID(osType CIMOSType, version string) types.VirtualMachineGuestOsIdentifier {
	id := int(osType)

	// Try exact (cimId, version) match first.
	if version != "" {
		if g, ok := cimGuestIDIndex[cimGuestIDKey(id, version)]; ok {
			return g
		}
	}

	// Fall back to bare cimId.
	if g, ok := cimGuestIDIndex[cimGuestIDKey(id, "")]; ok {
		return g
	}

	return ""
}

// GuestIDToCIMOSType translates a vSphere Guest OS identifier to a CIM
// OSType value.
func GuestIDToCIMOSType[T ~string](guestID T) CIMOSType {
	switch types.VirtualMachineGuestOsIdentifier(guestID) {
	// Windows Desktop
	case types.VirtualMachineGuestOsIdentifierDosGuest:
		return CIMOSTypeMSDOS
	case types.VirtualMachineGuestOsIdentifierWin31Guest:
		return CIMOSTypeWIN3x
	case types.VirtualMachineGuestOsIdentifierWin95Guest:
		return CIMOSTypeWIN95
	case types.VirtualMachineGuestOsIdentifierWin98Guest:
		return CIMOSTypeWIN98
	case types.VirtualMachineGuestOsIdentifierWinMeGuest:
		return CIMOSTypeWindowsMe
	case types.VirtualMachineGuestOsIdentifierWinNTGuest:
		return CIMOSTypeWINNT
	case types.VirtualMachineGuestOsIdentifierWin2000ProGuest:
		return CIMOSTypeWindows2000
	case types.VirtualMachineGuestOsIdentifierWin2000ServGuest:
		return CIMOSTypeWindows2000
	case types.VirtualMachineGuestOsIdentifierWin2000AdvServGuest:
		return CIMOSTypeWindows2000
	case types.VirtualMachineGuestOsIdentifierWinXPHomeGuest:
		return CIMOSTypeWindowsXP
	case types.VirtualMachineGuestOsIdentifierWinXPProGuest:
		return CIMOSTypeWindowsXP
	case types.VirtualMachineGuestOsIdentifierWinXPPro64Guest:
		return CIMOSTypeWindowsXP64Bit
	case types.VirtualMachineGuestOsIdentifierWinVistaGuest:
		return CIMOSTypeWindowsVista
	case types.VirtualMachineGuestOsIdentifierWinVista64Guest:
		return CIMOSTypeWindowsVista64Bit
	case types.VirtualMachineGuestOsIdentifierWindows7Guest:
		return CIMOSTypeMicrosoftWindows7
	case types.VirtualMachineGuestOsIdentifierWindows7_64Guest:
		return CIMOSTypeMicrosoftWindows7
	case types.VirtualMachineGuestOsIdentifierWindows8Guest:
		return CIMOSTypeMicrosoftWindows8
	case types.VirtualMachineGuestOsIdentifierWindows8_64Guest:
		return CIMOSTypeMicrosoftWindows8
	// TODO(akutz) The following guest IDs do not exist.
	/*
		case types.VirtualMachineGuestOsIdentifierWindows81Guest:
			return CIMOSTypeMicrosoftWindows81
		case types.VirtualMachineGuestOsIdentifierWindows81_64Guest:
			return CIMOSTypeMicrosoftWindows81
	*/
	case types.VirtualMachineGuestOsIdentifierWindows9Guest:
		return CIMOSTypeOther // Windows 10/11 - no specific CIM type
	case types.VirtualMachineGuestOsIdentifierWindows9_64Guest:
		return CIMOSTypeOther64Bit
	case types.VirtualMachineGuestOsIdentifierWindows11_64Guest:
		return CIMOSTypeOther64Bit
	case types.VirtualMachineGuestOsIdentifierWindows12_64Guest:
		return CIMOSTypeOther64Bit
	case types.VirtualMachineGuestOsIdentifierWindowsHyperVGuest:
		return CIMOSTypeOther64Bit

	// Windows Server
	case types.VirtualMachineGuestOsIdentifierWinNetEnterpriseGuest:
		return CIMOSTypeMicrosoftWindowsServer2003
	case types.VirtualMachineGuestOsIdentifierWinNetDatacenterGuest:
		return CIMOSTypeMicrosoftWindowsServer2003
	case types.VirtualMachineGuestOsIdentifierWinNetBusinessGuest:
		return CIMOSTypeMicrosoftWindowsServer2003
	case types.VirtualMachineGuestOsIdentifierWinNetStandardGuest:
		return CIMOSTypeMicrosoftWindowsServer2003
	case types.VirtualMachineGuestOsIdentifierWinNetWebGuest:
		return CIMOSTypeMicrosoftWindowsServer2003
	case types.VirtualMachineGuestOsIdentifierWinNetEnterprise64Guest:
		return CIMOSTypeMicrosoftWindowsServer2003_64Bit
	case types.VirtualMachineGuestOsIdentifierWinNetDatacenter64Guest:
		return CIMOSTypeMicrosoftWindowsServer2003_64Bit
	case types.VirtualMachineGuestOsIdentifierWinNetStandard64Guest:
		return CIMOSTypeMicrosoftWindowsServer2003_64Bit
	case types.VirtualMachineGuestOsIdentifierWinLonghornGuest:
		return CIMOSTypeMicrosoftWindowsServer2008
	case types.VirtualMachineGuestOsIdentifierWinLonghorn64Guest:
		return CIMOSTypeMicrosoftWindowsServer2008_64Bit
	case types.VirtualMachineGuestOsIdentifierWindows7Server64Guest:
		return CIMOSTypeMicrosoftWindowsServer2008R2
	case types.VirtualMachineGuestOsIdentifierWindows8Server64Guest:
		return CIMOSTypeMicrosoftWindowsServer2012
	case types.VirtualMachineGuestOsIdentifierWindows9Server64Guest:
		return CIMOSTypeMicrosoftWindowsServer2012R2
	case types.VirtualMachineGuestOsIdentifierWindows2019srv_64Guest:
		return CIMOSTypeOther64Bit // No specific CIM type for 2019+
	case types.VirtualMachineGuestOsIdentifierWindows2019srvNext_64Guest:
		return CIMOSTypeOther64Bit
	case types.VirtualMachineGuestOsIdentifierWindows2022srvNext_64Guest:
		return CIMOSTypeOther64Bit

	// Linux
	case types.VirtualMachineGuestOsIdentifierAsianux3Guest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierAsianux3_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAsianux4Guest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierAsianux4_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAsianux5_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAsianux7_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAsianux8_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAsianux9_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierCentos6Guest:
		return CIMOSTypeCentOS32bit
	case types.VirtualMachineGuestOsIdentifierCentos6_64Guest:
		return CIMOSTypeCentOS64bit
	case types.VirtualMachineGuestOsIdentifierCentos7Guest:
		return CIMOSTypeCentOS32bit
	case types.VirtualMachineGuestOsIdentifierCentos7_64Guest:
		return CIMOSTypeCentOS64bit
	case types.VirtualMachineGuestOsIdentifierCentos8_64Guest:
		return CIMOSTypeCentOS64bit
	case types.VirtualMachineGuestOsIdentifierCentos9_64Guest:
		return CIMOSTypeCentOS64bit
	case types.VirtualMachineGuestOsIdentifierCoreos64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierDebian4Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian4_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian5Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian5_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian6Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian6_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian7Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian7_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian8Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian8_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian9Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian9_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian10Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian10_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian11Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian11_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierDebian12Guest:
		return CIMOSTypeDebian
	case types.VirtualMachineGuestOsIdentifierDebian12_64Guest:
		return CIMOSTypeDebian64Bit
	case types.VirtualMachineGuestOsIdentifierFedoraGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierFedora64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierFreebsd11Guest:
		return CIMOSTypeFreeBSD
	case types.VirtualMachineGuestOsIdentifierFreebsd11_64Guest:
		return CIMOSTypeFreeBSD64Bit
	case types.VirtualMachineGuestOsIdentifierFreebsd12Guest:
		return CIMOSTypeFreeBSD
	case types.VirtualMachineGuestOsIdentifierFreebsd12_64Guest:
		return CIMOSTypeFreeBSD64Bit
	case types.VirtualMachineGuestOsIdentifierFreebsd13Guest:
		return CIMOSTypeFreeBSD
	case types.VirtualMachineGuestOsIdentifierFreebsd13_64Guest:
		return CIMOSTypeFreeBSD64Bit
	case types.VirtualMachineGuestOsIdentifierFreebsd14Guest:
		return CIMOSTypeFreeBSD
	case types.VirtualMachineGuestOsIdentifierFreebsd14_64Guest:
		return CIMOSTypeFreeBSD64Bit
	case types.VirtualMachineGuestOsIdentifierFreebsdGuest:
		return CIMOSTypeFreeBSD
	case types.VirtualMachineGuestOsIdentifierFreebsd64Guest:
		return CIMOSTypeFreeBSD64Bit
	case types.VirtualMachineGuestOsIdentifierGenericLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierMandrakeGuest:
		return CIMOSTypeMandriva
	case types.VirtualMachineGuestOsIdentifierMandriva64Guest:
		return CIMOSTypeMandriva64Bit
	case types.VirtualMachineGuestOsIdentifierMandrivaGuest:
		return CIMOSTypeMandriva
	// TODO(akutz) The following guest IDs do not exist.
	/*
		case types.VirtualMachineGuestOsIdentifierNetbsd7Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierNetbsd7_64Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierNetbsd8Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierNetbsd8_64Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierNetbsd9Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierNetbsd9_64Guest:
			return CIMOSTypeNetBSD
		case types.VirtualMachineGuestOsIdentifierOpenbsd7Guest:
			return CIMOSTypeOpenBSD
		case types.VirtualMachineGuestOsIdentifierOpenbsd7_64Guest:
			return CIMOSTypeOpenBSD
		case types.VirtualMachineGuestOsIdentifierOpenbsd8Guest:
			return CIMOSTypeOpenBSD
		case types.VirtualMachineGuestOsIdentifierOpenbsd8_64Guest:
			return CIMOSTypeOpenBSD
	*/
	case types.VirtualMachineGuestOsIdentifierOpensuse64Guest:
		return CIMOSTypeSUSE64Bit
	case types.VirtualMachineGuestOsIdentifierOpensuseGuest:
		return CIMOSTypeSUSE
	case types.VirtualMachineGuestOsIdentifierOracleLinux6Guest:
		return CIMOSTypeOracle32bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux6_64Guest:
		return CIMOSTypeOracle64bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux7Guest:
		return CIMOSTypeOracle32bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux7_64Guest:
		return CIMOSTypeOracle64bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux8_64Guest:
		return CIMOSTypeOracle64bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux9_64Guest:
		return CIMOSTypeOracle64bit
	case types.VirtualMachineGuestOsIdentifierOracleLinuxGuest:
		return CIMOSTypeOracle32bit
	case types.VirtualMachineGuestOsIdentifierOracleLinux64Guest:
		return CIMOSTypeOracle64bit
	case types.VirtualMachineGuestOsIdentifierOther24xLinux64Guest:
		return CIMOSTypeLinux24x64Bit
	case types.VirtualMachineGuestOsIdentifierOther24xLinuxGuest:
		return CIMOSTypeLinux24x
	case types.VirtualMachineGuestOsIdentifierOther26xLinux64Guest:
		return CIMOSTypeLinux26x64Bit
	case types.VirtualMachineGuestOsIdentifierOther26xLinuxGuest:
		return CIMOSTypeLinux26x
	case types.VirtualMachineGuestOsIdentifierOther3xLinux64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierOther3xLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierOther4xLinux64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierOther4xLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierOther5xLinux64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierOther5xLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierOther6xLinux64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierOther6xLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierOtherLinux64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierOtherLinuxGuest:
		return CIMOSTypeLINUX
	case types.VirtualMachineGuestOsIdentifierRedhatGuest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel2Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel3Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel3_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel4Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel4_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel5Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel5_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel6Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel6_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel7Guest:
		return CIMOSTypeRedHatEnterpriseLinux
	case types.VirtualMachineGuestOsIdentifierRhel7_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel8_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRhel9_64Guest:
		return CIMOSTypeRedHatEnterpriseLinux64Bit
	case types.VirtualMachineGuestOsIdentifierRockylinux_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierSles10Guest:
		return CIMOSTypeSLES
	case types.VirtualMachineGuestOsIdentifierSles10_64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSles11Guest:
		return CIMOSTypeSLES
	case types.VirtualMachineGuestOsIdentifierSles11_64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSles12Guest:
		return CIMOSTypeSLES
	case types.VirtualMachineGuestOsIdentifierSles12_64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSles15_64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSles16_64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSlesGuest:
		return CIMOSTypeSLES
	case types.VirtualMachineGuestOsIdentifierSles64Guest:
		return CIMOSTypeSLES64Bit
	case types.VirtualMachineGuestOsIdentifierSuse64Guest:
		return CIMOSTypeSUSE64Bit
	case types.VirtualMachineGuestOsIdentifierSuseGuest:
		return CIMOSTypeSUSE
	case types.VirtualMachineGuestOsIdentifierTurboLinux64Guest:
		return CIMOSTypeTurboLinux64Bit
	case types.VirtualMachineGuestOsIdentifierTurboLinuxGuest:
		return CIMOSTypeTurboLinux
	case types.VirtualMachineGuestOsIdentifierUbuntu64Guest:
		return CIMOSTypeUbuntu64Bit
	case types.VirtualMachineGuestOsIdentifierUbuntuGuest:
		return CIMOSTypeUbuntu
	case types.VirtualMachineGuestOsIdentifierUnixWare7Guest:
		return CIMOSTypeSCOUnixWare

	// macOS
	case types.VirtualMachineGuestOsIdentifierDarwin10Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin10_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin11Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin11_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin12_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin13_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin14_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin15_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin16_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin17_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin18_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin19_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin20_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin21_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin22_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin23_64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwin64Guest:
		return CIMOSTypeMACOS
	case types.VirtualMachineGuestOsIdentifierDarwinGuest:
		return CIMOSTypeMACOS

	// Solaris
	case types.VirtualMachineGuestOsIdentifierSolaris10Guest:
		return CIMOSTypeSolaris
	case types.VirtualMachineGuestOsIdentifierSolaris10_64Guest:
		return CIMOSTypeSolaris64Bit
	case types.VirtualMachineGuestOsIdentifierSolaris11_64Guest:
		return CIMOSTypeSolaris64Bit
	case types.VirtualMachineGuestOsIdentifierSolaris6Guest:
		return CIMOSTypeSolaris
	case types.VirtualMachineGuestOsIdentifierSolaris7Guest:
		return CIMOSTypeSolaris
	case types.VirtualMachineGuestOsIdentifierSolaris8Guest:
		return CIMOSTypeSolaris
	case types.VirtualMachineGuestOsIdentifierSolaris9Guest:
		return CIMOSTypeSolaris

	// Netware
	case types.VirtualMachineGuestOsIdentifierNetware4Guest:
		return CIMOSTypeNetWare
	case types.VirtualMachineGuestOsIdentifierNetware5Guest:
		return CIMOSTypeNetWare
	case types.VirtualMachineGuestOsIdentifierNetware6Guest:
		return CIMOSTypeNetWare
	case types.VirtualMachineGuestOsIdentifierNld9Guest:
		return CIMOSTypeNovellLinuxDesktop
	case types.VirtualMachineGuestOsIdentifierOesGuest:
		return CIMOSTypeNovellOES

	// VMware
	case types.VirtualMachineGuestOsIdentifierVmkernelGuest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmkernel5Guest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmkernel6Guest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmkernel65Guest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmkernel7Guest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmkernel8Guest:
		return CIMOSTypeVMwareESXi
	case types.VirtualMachineGuestOsIdentifierVmwarePhoton64Guest:
		return CIMOSTypeLinux64Bit

	// OS/2
	case types.VirtualMachineGuestOsIdentifierEComStationGuest:
		return CIMOSTypeeComStation32bitx
	case types.VirtualMachineGuestOsIdentifierEComStation2Guest:
		return CIMOSTypeeComStation32bitx
	case types.VirtualMachineGuestOsIdentifierOs2Guest:
		return CIMOSTypeOS2

	// Other
	case types.VirtualMachineGuestOsIdentifierOtherGuest:
		return CIMOSTypeOther
	case types.VirtualMachineGuestOsIdentifierOtherGuest64:
		return CIMOSTypeOther64Bit
	case types.VirtualMachineGuestOsIdentifierAmazonlinux2_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierAmazonlinux3_64Guest:
		return CIMOSTypeLinux64Bit
	case types.VirtualMachineGuestOsIdentifierCrxPod1Guest:
		return CIMOSTypeOther

	// Default to Unknown if not found
	default:
		return CIMOSTypeUnknown
	}
}
