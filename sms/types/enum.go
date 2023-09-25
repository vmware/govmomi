/*
Copyright (c) 2014-2023 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type AlarmType string

const (
	AlarmTypeSpaceCapacityAlarm = AlarmType("SpaceCapacityAlarm")
	AlarmTypeCapabilityAlarm    = AlarmType("CapabilityAlarm")
	AlarmTypeStorageObjectAlarm = AlarmType("StorageObjectAlarm")
	AlarmTypeObjectAlarm        = AlarmType("ObjectAlarm")
	AlarmTypeComplianceAlarm    = AlarmType("ComplianceAlarm")
	AlarmTypeManageabilityAlarm = AlarmType("ManageabilityAlarm")
	AlarmTypeReplicationAlarm   = AlarmType("ReplicationAlarm")
	AlarmTypeCertificateAlarm   = AlarmType("CertificateAlarm")
)

func init() {
	types.Add("sms:AlarmType", reflect.TypeOf((*AlarmType)(nil)).Elem())
}

// List of possible BackingStoragePool types
type BackingStoragePoolType string

const (
	BackingStoragePoolTypeThinProvisioningPool             = BackingStoragePoolType("thinProvisioningPool")
	BackingStoragePoolTypeDeduplicationPool                = BackingStoragePoolType("deduplicationPool")
	BackingStoragePoolTypeThinAndDeduplicationCombinedPool = BackingStoragePoolType("thinAndDeduplicationCombinedPool")
)

func init() {
	types.Add("sms:BackingStoragePoolType", reflect.TypeOf((*BackingStoragePoolType)(nil)).Elem())
}

// List of possible block device interfaces
type BlockDeviceInterface string

const (
	BlockDeviceInterfaceFc         = BlockDeviceInterface("fc")
	BlockDeviceInterfaceIscsi      = BlockDeviceInterface("iscsi")
	BlockDeviceInterfaceFcoe       = BlockDeviceInterface("fcoe")
	BlockDeviceInterfaceOtherBlock = BlockDeviceInterface("otherBlock")
)

func init() {
	types.Add("sms:BlockDeviceInterface", reflect.TypeOf((*BlockDeviceInterface)(nil)).Elem())
}

// Types of entities supported by the service.
type EntityReferenceEntityType string

const (
	EntityReferenceEntityTypeDatacenter   = EntityReferenceEntityType("datacenter")
	EntityReferenceEntityTypeResourcePool = EntityReferenceEntityType("resourcePool")
	EntityReferenceEntityTypeStoragePod   = EntityReferenceEntityType("storagePod")
	EntityReferenceEntityTypeCluster      = EntityReferenceEntityType("cluster")
	EntityReferenceEntityTypeVm           = EntityReferenceEntityType("vm")
	EntityReferenceEntityTypeDatastore    = EntityReferenceEntityType("datastore")
	EntityReferenceEntityTypeHost         = EntityReferenceEntityType("host")
	EntityReferenceEntityTypeVmFile       = EntityReferenceEntityType("vmFile")
	EntityReferenceEntityTypeScsiPath     = EntityReferenceEntityType("scsiPath")
	EntityReferenceEntityTypeScsiTarget   = EntityReferenceEntityType("scsiTarget")
	EntityReferenceEntityTypeScsiVolume   = EntityReferenceEntityType("scsiVolume")
	EntityReferenceEntityTypeScsiAdapter  = EntityReferenceEntityType("scsiAdapter")
	EntityReferenceEntityTypeNasMount     = EntityReferenceEntityType("nasMount")
)

func init() {
	types.Add("sms:EntityReferenceEntityType", reflect.TypeOf((*EntityReferenceEntityType)(nil)).Elem())
}

// List of possible file system interfaces
type FileSystemInterface string

const (
	FileSystemInterfaceNfs             = FileSystemInterface("nfs")
	FileSystemInterfaceOtherFileSystem = FileSystemInterface("otherFileSystem")
)

func init() {
	types.Add("sms:FileSystemInterface", reflect.TypeOf((*FileSystemInterface)(nil)).Elem())
}

type FileSystemInterfaceVersion string

const (
	FileSystemInterfaceVersionNFSV3_0 = FileSystemInterfaceVersion("NFSV3_0")
)

func init() {
	types.Add("sms:FileSystemInterfaceVersion", reflect.TypeOf((*FileSystemInterfaceVersion)(nil)).Elem())
}

// Profiles supported by VASA Provider.
type ProviderProfile string

const (
	// PBM profile
	ProviderProfileProfileBasedManagement = ProviderProfile("ProfileBasedManagement")
	// Replication profile
	ProviderProfileReplication = ProviderProfile("Replication")
)

func init() {
	types.Add("sms:ProviderProfile", reflect.TypeOf((*ProviderProfile)(nil)).Elem())
}

// State of the replication group at the site of the query.
//
// A replication group
// may be in different states at the source site and each of the target sites.
// Note that this state does not capture the health of the replication link. If
// necessary, that can be an additional attribute.
type ReplicationReplicationState string

const (
	// Replication Source
	ReplicationReplicationStateSOURCE = ReplicationReplicationState("SOURCE")
	// Replication target
	ReplicationReplicationStateTARGET = ReplicationReplicationState("TARGET")
	// The group failed over at this site of the query.
	//
	// It has not yet been made
	// as a source of replication.
	ReplicationReplicationStateFAILEDOVER = ReplicationReplicationState("FAILEDOVER")
	// The group is InTest.
	//
	// The testFailover devices list will be available from
	// the `TargetGroupMemberInfo`
	ReplicationReplicationStateINTEST = ReplicationReplicationState("INTEST")
	// Remote group was failed over, and this site is neither the source nor the
	// target.
	ReplicationReplicationStateREMOTE_FAILEDOVER = ReplicationReplicationState("REMOTE_FAILEDOVER")
)

func init() {
	types.Add("sms:ReplicationReplicationState", reflect.TypeOf((*ReplicationReplicationState)(nil)).Elem())
}

// Enumeration of the supported Alarm Status values
type SmsAlarmStatus string

const (
	SmsAlarmStatusRed    = SmsAlarmStatus("Red")
	SmsAlarmStatusGreen  = SmsAlarmStatus("Green")
	SmsAlarmStatusYellow = SmsAlarmStatus("Yellow")
)

func init() {
	types.Add("sms:SmsAlarmStatus", reflect.TypeOf((*SmsAlarmStatus)(nil)).Elem())
}

// Enumeration of the supported Entity Type values.
type SmsEntityType string

const (
	SmsEntityTypeStorageArrayEntity        = SmsEntityType("StorageArrayEntity")
	SmsEntityTypeStorageProcessorEntity    = SmsEntityType("StorageProcessorEntity")
	SmsEntityTypeStoragePortEntity         = SmsEntityType("StoragePortEntity")
	SmsEntityTypeStorageLunEntity          = SmsEntityType("StorageLunEntity")
	SmsEntityTypeStorageFileSystemEntity   = SmsEntityType("StorageFileSystemEntity")
	SmsEntityTypeStorageCapabilityEntity   = SmsEntityType("StorageCapabilityEntity")
	SmsEntityTypeCapabilitySchemaEntity    = SmsEntityType("CapabilitySchemaEntity")
	SmsEntityTypeCapabilityProfileEntity   = SmsEntityType("CapabilityProfileEntity")
	SmsEntityTypeDefaultProfileEntity      = SmsEntityType("DefaultProfileEntity")
	SmsEntityTypeResourceAssociationEntity = SmsEntityType("ResourceAssociationEntity")
	SmsEntityTypeStorageContainerEntity    = SmsEntityType("StorageContainerEntity")
	SmsEntityTypeStorageObjectEntity       = SmsEntityType("StorageObjectEntity")
	SmsEntityTypeMessageCatalogEntity      = SmsEntityType("MessageCatalogEntity")
	SmsEntityTypeProtocolEndpointEntity    = SmsEntityType("ProtocolEndpointEntity")
	SmsEntityTypeVirtualVolumeInfoEntity   = SmsEntityType("VirtualVolumeInfoEntity")
	SmsEntityTypeBackingStoragePoolEntity  = SmsEntityType("BackingStoragePoolEntity")
	SmsEntityTypeFaultDomainEntity         = SmsEntityType("FaultDomainEntity")
	SmsEntityTypeReplicationGroupEntity    = SmsEntityType("ReplicationGroupEntity")
)

func init() {
	types.Add("sms:SmsEntityType", reflect.TypeOf((*SmsEntityType)(nil)).Elem())
}

// List of possible states of a task.
type SmsTaskState string

const (
	// Task is put in the queue.
	SmsTaskStateQueued = SmsTaskState("queued")
	// Task is currently running.
	SmsTaskStateRunning = SmsTaskState("running")
	// Task has completed.
	SmsTaskStateSuccess = SmsTaskState("success")
	// Task has encountered an error or has been canceled.
	SmsTaskStateError = SmsTaskState("error")
)

func init() {
	types.Add("sms:SmsTaskState", reflect.TypeOf((*SmsTaskState)(nil)).Elem())
}

// List of supported VVOL Container types
type StorageContainerVvolContainerTypeEnum string

const (
	StorageContainerVvolContainerTypeEnumNFS   = StorageContainerVvolContainerTypeEnum("NFS")
	StorageContainerVvolContainerTypeEnumNFS4x = StorageContainerVvolContainerTypeEnum("NFS4x")
	StorageContainerVvolContainerTypeEnumSCSI  = StorageContainerVvolContainerTypeEnum("SCSI")
	StorageContainerVvolContainerTypeEnumNVMe  = StorageContainerVvolContainerTypeEnum("NVMe")
)

func init() {
	types.Add("sms:StorageContainerVvolContainerTypeEnum", reflect.TypeOf((*StorageContainerVvolContainerTypeEnum)(nil)).Elem())
}

// List of possible values for thin provisioning status alarm.
type ThinProvisioningStatus string

const (
	ThinProvisioningStatusRED    = ThinProvisioningStatus("RED")
	ThinProvisioningStatusYELLOW = ThinProvisioningStatus("YELLOW")
	ThinProvisioningStatusGREEN  = ThinProvisioningStatus("GREEN")
)

func init() {
	types.Add("sms:ThinProvisioningStatus", reflect.TypeOf((*ThinProvisioningStatus)(nil)).Elem())
}

// VASA provider authentication type.
type VasaAuthenticationType string

const (
	// Login using SAML token.
	VasaAuthenticationTypeLoginByToken = VasaAuthenticationType("LoginByToken")
	// Use id of an existing session that has logged-in from somewhere else.
	VasaAuthenticationTypeUseSessionId = VasaAuthenticationType("UseSessionId")
)

func init() {
	types.Add("sms:VasaAuthenticationType", reflect.TypeOf((*VasaAuthenticationType)(nil)).Elem())
}

// List of possible VASA profiles supported by Storage Array
type VasaProfile string

const (
	// Block device profile
	VasaProfileBlockDevice = VasaProfile("blockDevice")
	// File system profile
	VasaProfileFileSystem = VasaProfile("fileSystem")
	// Storage capability profile
	VasaProfileCapability = VasaProfile("capability")
	// Policy profile
	VasaProfilePolicy = VasaProfile("policy")
	// Object based storage profile
	VasaProfileObject = VasaProfile("object")
	// IO Statistics profile
	VasaProfileStatistics = VasaProfile("statistics")
	// Storage DRS specific block device profile
	VasaProfileStorageDrsBlockDevice = VasaProfile("storageDrsBlockDevice")
	// Storage DRS specific file system profile
	VasaProfileStorageDrsFileSystem = VasaProfile("storageDrsFileSystem")
)

func init() {
	types.Add("sms:VasaProfile", reflect.TypeOf((*VasaProfile)(nil)).Elem())
}

// The status of the provider certificate
type VasaProviderCertificateStatus string

const (
	// Provider certificate is valid.
	VasaProviderCertificateStatusValid = VasaProviderCertificateStatus("valid")
	// Provider certificate is within the soft limit threshold.
	VasaProviderCertificateStatusExpirySoftLimitReached = VasaProviderCertificateStatus("expirySoftLimitReached")
	// Provider certificate is within the hard limit threshold.
	VasaProviderCertificateStatusExpiryHardLimitReached = VasaProviderCertificateStatus("expiryHardLimitReached")
	// Provider certificate has expired.
	VasaProviderCertificateStatusExpired = VasaProviderCertificateStatus("expired")
	// Provider certificate is revoked, malformed or missing.
	VasaProviderCertificateStatusInvalid = VasaProviderCertificateStatus("invalid")
)

func init() {
	types.Add("sms:VasaProviderCertificateStatus", reflect.TypeOf((*VasaProviderCertificateStatus)(nil)).Elem())
}

// Deprecated as of SMS API 3.0, use `VasaProfile_enum`.
//
// Profiles supported by VASA Provider.
type VasaProviderProfile string

const (
	// Block device profile
	VasaProviderProfileBlockDevice = VasaProviderProfile("blockDevice")
	// File system profile
	VasaProviderProfileFileSystem = VasaProviderProfile("fileSystem")
	// Storage capability profile
	VasaProviderProfileCapability = VasaProviderProfile("capability")
)

func init() {
	types.Add("sms:VasaProviderProfile", reflect.TypeOf((*VasaProviderProfile)(nil)).Elem())
}

// The operational state of VASA Provider.
type VasaProviderStatus string

const (
	// VASA Provider is operating correctly.
	VasaProviderStatusOnline = VasaProviderStatus("online")
	// VASA Provider is not responding, e.g.
	//
	// communication error due to temporary
	// network outage. SMS keeps polling the provider in this state.
	VasaProviderStatusOffline = VasaProviderStatus("offline")
	// VASA Provider is connected, but sync operation failed.
	VasaProviderStatusSyncError = VasaProviderStatus("syncError")
	//
	//
	// Deprecated as of SMS API 4.0, this status is deprecated.
	//
	// VASA Provider is unreachable.
	VasaProviderStatusUnknown = VasaProviderStatus("unknown")
	// VASA Provider is connected, but has not triggered sync operation.
	VasaProviderStatusConnected = VasaProviderStatus("connected")
	// VASA Provider is disconnected, e.g.
	//
	// failed to establish a valid
	// SSL connection to the provider. SMS stops communication with the
	// provider in this state. The user can reconnect to the provider by invoking
	// `VasaProvider.VasaProviderReconnect_Task`.
	VasaProviderStatusDisconnected = VasaProviderStatus("disconnected")
)

func init() {
	types.Add("sms:VasaProviderStatus", reflect.TypeOf((*VasaProviderStatus)(nil)).Elem())
}

// A Category to indicate whether provider is of internal or external category.
//
// This classification can help selectively enable few administrative functions
// such as say unregistration of a provider.
type VpCategory string

const (
	// An internal provider category indicates the set of providers such as IOFILTERS and VSAN.
	VpCategoryInternal = VpCategory("internal")
	// An external provider category indicates the set of providers are external and not belong
	// to say either of IOFILTERS or VSAN category.
	VpCategoryExternal = VpCategory("external")
)

func init() {
	types.Add("sms:VpCategory", reflect.TypeOf((*VpCategory)(nil)).Elem())
}

// VASA Provider type.
type VpType string

const (
	// Persistence provider.
	VpTypePERSISTENCE = VpType("PERSISTENCE")
	// DataService provider.
	//
	// No storage supported for this type of provider.
	VpTypeDATASERVICE = VpType("DATASERVICE")
	// Type is unknown.
	//
	// VASA provider type can be UNKNOWN when it is undergoing sync operation.
	VpTypeUNKNOWN = VpType("UNKNOWN")
)

func init() {
	types.Add("sms:VpType", reflect.TypeOf((*VpType)(nil)).Elem())
}
