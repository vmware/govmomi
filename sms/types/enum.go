// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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

func (e AlarmType) Values() []AlarmType {
	return []AlarmType{
		AlarmTypeSpaceCapacityAlarm,
		AlarmTypeCapabilityAlarm,
		AlarmTypeStorageObjectAlarm,
		AlarmTypeObjectAlarm,
		AlarmTypeComplianceAlarm,
		AlarmTypeManageabilityAlarm,
		AlarmTypeReplicationAlarm,
		AlarmTypeCertificateAlarm,
	}
}

func (e AlarmType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e BackingStoragePoolType) Values() []BackingStoragePoolType {
	return []BackingStoragePoolType{
		BackingStoragePoolTypeThinProvisioningPool,
		BackingStoragePoolTypeDeduplicationPool,
		BackingStoragePoolTypeThinAndDeduplicationCombinedPool,
	}
}

func (e BackingStoragePoolType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e BlockDeviceInterface) Values() []BlockDeviceInterface {
	return []BlockDeviceInterface{
		BlockDeviceInterfaceFc,
		BlockDeviceInterfaceIscsi,
		BlockDeviceInterfaceFcoe,
		BlockDeviceInterfaceOtherBlock,
	}
}

func (e BlockDeviceInterface) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e EntityReferenceEntityType) Values() []EntityReferenceEntityType {
	return []EntityReferenceEntityType{
		EntityReferenceEntityTypeDatacenter,
		EntityReferenceEntityTypeResourcePool,
		EntityReferenceEntityTypeStoragePod,
		EntityReferenceEntityTypeCluster,
		EntityReferenceEntityTypeVm,
		EntityReferenceEntityTypeDatastore,
		EntityReferenceEntityTypeHost,
		EntityReferenceEntityTypeVmFile,
		EntityReferenceEntityTypeScsiPath,
		EntityReferenceEntityTypeScsiTarget,
		EntityReferenceEntityTypeScsiVolume,
		EntityReferenceEntityTypeScsiAdapter,
		EntityReferenceEntityTypeNasMount,
	}
}

func (e EntityReferenceEntityType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("sms:EntityReferenceEntityType", reflect.TypeOf((*EntityReferenceEntityType)(nil)).Elem())
}

// List of possible file system interfaces
type FileSystemInterface string

const (
	FileSystemInterfaceNfs             = FileSystemInterface("nfs")
	FileSystemInterfaceOtherFileSystem = FileSystemInterface("otherFileSystem")
)

func (e FileSystemInterface) Values() []FileSystemInterface {
	return []FileSystemInterface{
		FileSystemInterfaceNfs,
		FileSystemInterfaceOtherFileSystem,
	}
}

func (e FileSystemInterface) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("sms:FileSystemInterface", reflect.TypeOf((*FileSystemInterface)(nil)).Elem())
}

type FileSystemInterfaceVersion string

const (
	FileSystemInterfaceVersionNFSV3_0 = FileSystemInterfaceVersion("NFSV3_0")
)

func (e FileSystemInterfaceVersion) Values() []FileSystemInterfaceVersion {
	return []FileSystemInterfaceVersion{
		FileSystemInterfaceVersionNFSV3_0,
	}
}

func (e FileSystemInterfaceVersion) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("sms:FileSystemInterfaceVersion", reflect.TypeOf((*FileSystemInterfaceVersion)(nil)).Elem())
}

type ManagedObjectType string

const (
	ManagedObjectTypeSmsServiceInstance = ManagedObjectType("SmsServiceInstance")
	ManagedObjectTypeSmsStorageManager  = ManagedObjectType("SmsStorageManager")
	ManagedObjectTypeSmsTask            = ManagedObjectType("SmsTask")
	ManagedObjectTypeSmsSessionManager  = ManagedObjectType("SmsSessionManager")
	ManagedObjectTypeSmsProvider        = ManagedObjectType("SmsProvider")
	ManagedObjectTypeVasaProvider       = ManagedObjectType("VasaProvider")
)

func (e ManagedObjectType) Values() []ManagedObjectType {
	return []ManagedObjectType{
		ManagedObjectTypeSmsServiceInstance,
		ManagedObjectTypeSmsStorageManager,
		ManagedObjectTypeSmsTask,
		ManagedObjectTypeSmsSessionManager,
		ManagedObjectTypeSmsProvider,
		ManagedObjectTypeVasaProvider,
	}
}

func (e ManagedObjectType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("sms:ManagedObjectType", reflect.TypeOf((*ManagedObjectType)(nil)).Elem())
}

// Profiles supported by VASA Provider.
type ProviderProfile string

const (
	// PBM profile
	ProviderProfileProfileBasedManagement = ProviderProfile("ProfileBasedManagement")
	// Replication profile
	ProviderProfileReplication = ProviderProfile("Replication")
)

func (e ProviderProfile) Values() []ProviderProfile {
	return []ProviderProfile{
		ProviderProfileProfileBasedManagement,
		ProviderProfileReplication,
	}
}

func (e ProviderProfile) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e ReplicationReplicationState) Values() []ReplicationReplicationState {
	return []ReplicationReplicationState{
		ReplicationReplicationStateSOURCE,
		ReplicationReplicationStateTARGET,
		ReplicationReplicationStateFAILEDOVER,
		ReplicationReplicationStateINTEST,
		ReplicationReplicationStateREMOTE_FAILEDOVER,
	}
}

func (e ReplicationReplicationState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e SmsAlarmStatus) Values() []SmsAlarmStatus {
	return []SmsAlarmStatus{
		SmsAlarmStatusRed,
		SmsAlarmStatusGreen,
		SmsAlarmStatusYellow,
	}
}

func (e SmsAlarmStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e SmsEntityType) Values() []SmsEntityType {
	return []SmsEntityType{
		SmsEntityTypeStorageArrayEntity,
		SmsEntityTypeStorageProcessorEntity,
		SmsEntityTypeStoragePortEntity,
		SmsEntityTypeStorageLunEntity,
		SmsEntityTypeStorageFileSystemEntity,
		SmsEntityTypeStorageCapabilityEntity,
		SmsEntityTypeCapabilitySchemaEntity,
		SmsEntityTypeCapabilityProfileEntity,
		SmsEntityTypeDefaultProfileEntity,
		SmsEntityTypeResourceAssociationEntity,
		SmsEntityTypeStorageContainerEntity,
		SmsEntityTypeStorageObjectEntity,
		SmsEntityTypeMessageCatalogEntity,
		SmsEntityTypeProtocolEndpointEntity,
		SmsEntityTypeVirtualVolumeInfoEntity,
		SmsEntityTypeBackingStoragePoolEntity,
		SmsEntityTypeFaultDomainEntity,
		SmsEntityTypeReplicationGroupEntity,
	}
}

func (e SmsEntityType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e SmsTaskState) Values() []SmsTaskState {
	return []SmsTaskState{
		SmsTaskStateQueued,
		SmsTaskStateRunning,
		SmsTaskStateSuccess,
		SmsTaskStateError,
	}
}

func (e SmsTaskState) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e StorageContainerVvolContainerTypeEnum) Values() []StorageContainerVvolContainerTypeEnum {
	return []StorageContainerVvolContainerTypeEnum{
		StorageContainerVvolContainerTypeEnumNFS,
		StorageContainerVvolContainerTypeEnumNFS4x,
		StorageContainerVvolContainerTypeEnumSCSI,
		StorageContainerVvolContainerTypeEnumNVMe,
	}
}

func (e StorageContainerVvolContainerTypeEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e ThinProvisioningStatus) Values() []ThinProvisioningStatus {
	return []ThinProvisioningStatus{
		ThinProvisioningStatusRED,
		ThinProvisioningStatusYELLOW,
		ThinProvisioningStatusGREEN,
	}
}

func (e ThinProvisioningStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VasaAuthenticationType) Values() []VasaAuthenticationType {
	return []VasaAuthenticationType{
		VasaAuthenticationTypeLoginByToken,
		VasaAuthenticationTypeUseSessionId,
	}
}

func (e VasaAuthenticationType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VasaProfile) Values() []VasaProfile {
	return []VasaProfile{
		VasaProfileBlockDevice,
		VasaProfileFileSystem,
		VasaProfileCapability,
		VasaProfilePolicy,
		VasaProfileObject,
		VasaProfileStatistics,
		VasaProfileStorageDrsBlockDevice,
		VasaProfileStorageDrsFileSystem,
	}
}

func (e VasaProfile) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VasaProviderCertificateStatus) Values() []VasaProviderCertificateStatus {
	return []VasaProviderCertificateStatus{
		VasaProviderCertificateStatusValid,
		VasaProviderCertificateStatusExpirySoftLimitReached,
		VasaProviderCertificateStatusExpiryHardLimitReached,
		VasaProviderCertificateStatusExpired,
		VasaProviderCertificateStatusInvalid,
	}
}

func (e VasaProviderCertificateStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VasaProviderProfile) Values() []VasaProviderProfile {
	return []VasaProviderProfile{
		VasaProviderProfileBlockDevice,
		VasaProviderProfileFileSystem,
		VasaProviderProfileCapability,
	}
}

func (e VasaProviderProfile) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VasaProviderStatus) Values() []VasaProviderStatus {
	return []VasaProviderStatus{
		VasaProviderStatusOnline,
		VasaProviderStatusOffline,
		VasaProviderStatusSyncError,
		VasaProviderStatusUnknown,
		VasaProviderStatusConnected,
		VasaProviderStatusDisconnected,
	}
}

func (e VasaProviderStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VpCategory) Values() []VpCategory {
	return []VpCategory{
		VpCategoryInternal,
		VpCategoryExternal,
	}
}

func (e VpCategory) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

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

func (e VpType) Values() []VpType {
	return []VpType{
		VpTypePERSISTENCE,
		VpTypeDATASERVICE,
		VpTypeUNKNOWN,
	}
}

func (e VpType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("sms:VpType", reflect.TypeOf((*VpType)(nil)).Elem())
}
