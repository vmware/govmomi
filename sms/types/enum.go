/*
Copyright (c) 2014-2022 VMware, Inc. All Rights Reserved.

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
)

func init() {
	types.Add("sms:AlarmType", reflect.TypeOf((*AlarmType)(nil)).Elem())
}

type BackingStoragePoolType string

const (
	BackingStoragePoolTypeThinProvisioningPool             = BackingStoragePoolType("thinProvisioningPool")
	BackingStoragePoolTypeDeduplicationPool                = BackingStoragePoolType("deduplicationPool")
	BackingStoragePoolTypeThinAndDeduplicationCombinedPool = BackingStoragePoolType("thinAndDeduplicationCombinedPool")
)

func init() {
	types.Add("sms:BackingStoragePoolType", reflect.TypeOf((*BackingStoragePoolType)(nil)).Elem())
}

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

type ProviderProfile string

const (
	ProviderProfileProfileBasedManagement = ProviderProfile("ProfileBasedManagement")
	ProviderProfileReplication            = ProviderProfile("Replication")
)

func init() {
	types.Add("sms:ProviderProfile", reflect.TypeOf((*ProviderProfile)(nil)).Elem())
}

type ReplicationReplicationState string

const (
	ReplicationReplicationStateSOURCE            = ReplicationReplicationState("SOURCE")
	ReplicationReplicationStateTARGET            = ReplicationReplicationState("TARGET")
	ReplicationReplicationStateFAILEDOVER        = ReplicationReplicationState("FAILEDOVER")
	ReplicationReplicationStateINTEST            = ReplicationReplicationState("INTEST")
	ReplicationReplicationStateREMOTE_FAILEDOVER = ReplicationReplicationState("REMOTE_FAILEDOVER")
)

func init() {
	types.Add("sms:ReplicationReplicationState", reflect.TypeOf((*ReplicationReplicationState)(nil)).Elem())
}

type SmsAlarmStatus string

const (
	SmsAlarmStatusRed    = SmsAlarmStatus("Red")
	SmsAlarmStatusGreen  = SmsAlarmStatus("Green")
	SmsAlarmStatusYellow = SmsAlarmStatus("Yellow")
)

func init() {
	types.Add("sms:SmsAlarmStatus", reflect.TypeOf((*SmsAlarmStatus)(nil)).Elem())
}

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

type SmsTaskState string

const (
	SmsTaskStateQueued  = SmsTaskState("queued")
	SmsTaskStateRunning = SmsTaskState("running")
	SmsTaskStateSuccess = SmsTaskState("success")
	SmsTaskStateError   = SmsTaskState("error")
)

func init() {
	types.Add("sms:SmsTaskState", reflect.TypeOf((*SmsTaskState)(nil)).Elem())
}

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

type ThinProvisioningStatus string

const (
	ThinProvisioningStatusRED    = ThinProvisioningStatus("RED")
	ThinProvisioningStatusYELLOW = ThinProvisioningStatus("YELLOW")
	ThinProvisioningStatusGREEN  = ThinProvisioningStatus("GREEN")
)

func init() {
	types.Add("sms:ThinProvisioningStatus", reflect.TypeOf((*ThinProvisioningStatus)(nil)).Elem())
}

type VasaAuthenticationType string

const (
	VasaAuthenticationTypeLoginByToken = VasaAuthenticationType("LoginByToken")
	VasaAuthenticationTypeUseSessionId = VasaAuthenticationType("UseSessionId")
)

func init() {
	types.Add("sms:VasaAuthenticationType", reflect.TypeOf((*VasaAuthenticationType)(nil)).Elem())
}

type VasaProfile string

const (
	VasaProfileBlockDevice           = VasaProfile("blockDevice")
	VasaProfileFileSystem            = VasaProfile("fileSystem")
	VasaProfileCapability            = VasaProfile("capability")
	VasaProfilePolicy                = VasaProfile("policy")
	VasaProfileObject                = VasaProfile("object")
	VasaProfileStatistics            = VasaProfile("statistics")
	VasaProfileStorageDrsBlockDevice = VasaProfile("storageDrsBlockDevice")
	VasaProfileStorageDrsFileSystem  = VasaProfile("storageDrsFileSystem")
)

func init() {
	types.Add("sms:VasaProfile", reflect.TypeOf((*VasaProfile)(nil)).Elem())
}

type VasaProviderCertificateStatus string

const (
	VasaProviderCertificateStatusValid                  = VasaProviderCertificateStatus("valid")
	VasaProviderCertificateStatusExpirySoftLimitReached = VasaProviderCertificateStatus("expirySoftLimitReached")
	VasaProviderCertificateStatusExpiryHardLimitReached = VasaProviderCertificateStatus("expiryHardLimitReached")
	VasaProviderCertificateStatusExpired                = VasaProviderCertificateStatus("expired")
	VasaProviderCertificateStatusInvalid                = VasaProviderCertificateStatus("invalid")
)

func init() {
	types.Add("sms:VasaProviderCertificateStatus", reflect.TypeOf((*VasaProviderCertificateStatus)(nil)).Elem())
}

type VasaProviderProfile string

const (
	VasaProviderProfileBlockDevice = VasaProviderProfile("blockDevice")
	VasaProviderProfileFileSystem  = VasaProviderProfile("fileSystem")
	VasaProviderProfileCapability  = VasaProviderProfile("capability")
)

func init() {
	types.Add("sms:VasaProviderProfile", reflect.TypeOf((*VasaProviderProfile)(nil)).Elem())
}

type VasaProviderStatus string

const (
	VasaProviderStatusOnline       = VasaProviderStatus("online")
	VasaProviderStatusOffline      = VasaProviderStatus("offline")
	VasaProviderStatusSyncError    = VasaProviderStatus("syncError")
	VasaProviderStatusUnknown      = VasaProviderStatus("unknown")
	VasaProviderStatusConnected    = VasaProviderStatus("connected")
	VasaProviderStatusDisconnected = VasaProviderStatus("disconnected")
)

func init() {
	types.Add("sms:VasaProviderStatus", reflect.TypeOf((*VasaProviderStatus)(nil)).Elem())
}

type VpCategory string

const (
	VpCategoryInternal = VpCategory("internal")
	VpCategoryExternal = VpCategory("external")
)

func init() {
	types.Add("sms:VpCategory", reflect.TypeOf((*VpCategory)(nil)).Elem())
}

type VpType string

const (
	VpTypePERSISTENCE = VpType("PERSISTENCE")
	VpTypeDATASERVICE = VpType("DATASERVICE")
	VpTypeUNKNOWN     = VpType("UNKNOWN")
)

func init() {
	types.Add("sms:VpType", reflect.TypeOf((*VpType)(nil)).Elem())
}
