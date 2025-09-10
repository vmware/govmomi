// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

// This spec contains information needed for queryActiveAlarm API to filter the result.
//
// This structure may be used only with operations rendered under `/sms`.
type AlarmFilter struct {
	types.DynamicData

	// The status of the alarm to search for.
	//
	// Should be one of
	// `SmsAlarmStatus_enum`. If not specified, all status values
	// should be considered.
	AlarmStatus string `xml:"alarmStatus,omitempty" json:"alarmStatus,omitempty"`
	// The status of the alarm to search for.
	//
	// Should be one of
	// `AlarmType_enum`. If not specified, all alarm types
	// should be considered.
	AlarmType string `xml:"alarmType,omitempty" json:"alarmType,omitempty"`
	// The entityType of interest, VASA provider should
	// return all active alarms of this type when `AlarmFilter.entityId`
	// is not set.
	//
	// See `SmsEntityType_enum`.
	EntityType string `xml:"entityType,omitempty" json:"entityType,omitempty"`
	// The identifiers of the entities of interest.
	//
	// If set, all entities must be
	// of the same `SmsEntityType_enum` and it should be set in
	// `AlarmFilter.entityType`. VASA provider can skip listing the missing entities.
	EntityId []types.AnyType `xml:"entityId,omitempty,typeattr" json:"entityId,omitempty"`
	// The page marker used for query pagination.
	//
	// This is an opaque string that
	// will be set based on the value returned by the VASA provider - see
	// `AlarmResult.pageMarker`. For initial request this should be set to
	// null, indicating request for the first page.
	PageMarker string `xml:"pageMarker,omitempty" json:"pageMarker,omitempty"`
}

func init() {
	types.Add("sms:AlarmFilter", reflect.TypeOf((*AlarmFilter)(nil)).Elem())
}

// Contains result for queryActiveAlarm API.
//
// This structure may be used only with operations rendered under `/sms`.
type AlarmResult struct {
	types.DynamicData

	// Resulting storage alarms.
	StorageAlarm []StorageAlarm `xml:"storageAlarm,omitempty" json:"storageAlarm,omitempty"`
	// The page marker used for query pagination.
	//
	// This is an opaque string that
	// will be set by the VASA provider. The client will set the same value in
	// `AlarmFilter.pageMarker` to query the next page. VP should unset
	// this value to indicate the end of page.
	PageMarker string `xml:"pageMarker,omitempty" json:"pageMarker,omitempty"`
}

func init() {
	types.Add("sms:AlarmResult", reflect.TypeOf((*AlarmResult)(nil)).Elem())
}

// Thrown if the object is already at the desired state.
//
// This is always a warning.
//
// This structure may be used only with operations rendered under `/sms`.
type AlreadyDone struct {
	SmsReplicationFault
}

func init() {
	types.Add("sms:AlreadyDone", reflect.TypeOf((*AlreadyDone)(nil)).Elem())
}

type AlreadyDoneFault AlreadyDone

func init() {
	types.Add("sms:AlreadyDoneFault", reflect.TypeOf((*AlreadyDoneFault)(nil)).Elem())
}

// A boxed array of `BackingStoragePool`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfBackingStoragePool struct {
	BackingStoragePool []BackingStoragePool `xml:"BackingStoragePool,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfBackingStoragePool", reflect.TypeOf((*ArrayOfBackingStoragePool)(nil)).Elem())
}

// A boxed array of `DatastoreBackingPoolMapping`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfDatastoreBackingPoolMapping struct {
	DatastoreBackingPoolMapping []DatastoreBackingPoolMapping `xml:"DatastoreBackingPoolMapping,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfDatastoreBackingPoolMapping", reflect.TypeOf((*ArrayOfDatastoreBackingPoolMapping)(nil)).Elem())
}

// A boxed array of `DatastorePair`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfDatastorePair struct {
	DatastorePair []DatastorePair `xml:"DatastorePair,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfDatastorePair", reflect.TypeOf((*ArrayOfDatastorePair)(nil)).Elem())
}

// A boxed array of `DeviceId`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfDeviceId struct {
	DeviceId []BaseDeviceId `xml:"DeviceId,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfDeviceId", reflect.TypeOf((*ArrayOfDeviceId)(nil)).Elem())
}

// A boxed array of `FaultDomainProviderMapping`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfFaultDomainProviderMapping struct {
	FaultDomainProviderMapping []FaultDomainProviderMapping `xml:"FaultDomainProviderMapping,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfFaultDomainProviderMapping", reflect.TypeOf((*ArrayOfFaultDomainProviderMapping)(nil)).Elem())
}

// A boxed array of `GroupOperationResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfGroupOperationResult struct {
	GroupOperationResult []BaseGroupOperationResult `xml:"GroupOperationResult,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfGroupOperationResult", reflect.TypeOf((*ArrayOfGroupOperationResult)(nil)).Elem())
}

// A boxed array of `NameValuePair`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfNameValuePair struct {
	NameValuePair []NameValuePair `xml:"NameValuePair,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfNameValuePair", reflect.TypeOf((*ArrayOfNameValuePair)(nil)).Elem())
}

// A boxed array of `PointInTimeReplicaInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfPointInTimeReplicaInfo struct {
	PointInTimeReplicaInfo []PointInTimeReplicaInfo `xml:"PointInTimeReplicaInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfPointInTimeReplicaInfo", reflect.TypeOf((*ArrayOfPointInTimeReplicaInfo)(nil)).Elem())
}

// A boxed array of `PolicyAssociation`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfPolicyAssociation struct {
	PolicyAssociation []PolicyAssociation `xml:"PolicyAssociation,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfPolicyAssociation", reflect.TypeOf((*ArrayOfPolicyAssociation)(nil)).Elem())
}

// A boxed array of `QueryReplicationPeerResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfQueryReplicationPeerResult struct {
	QueryReplicationPeerResult []QueryReplicationPeerResult `xml:"QueryReplicationPeerResult,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfQueryReplicationPeerResult", reflect.TypeOf((*ArrayOfQueryReplicationPeerResult)(nil)).Elem())
}

// A boxed array of `RecoveredDevice`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfRecoveredDevice struct {
	RecoveredDevice []RecoveredDevice `xml:"RecoveredDevice,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfRecoveredDevice", reflect.TypeOf((*ArrayOfRecoveredDevice)(nil)).Elem())
}

// A boxed array of `RecoveredDiskInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfRecoveredDiskInfo struct {
	RecoveredDiskInfo []RecoveredDiskInfo `xml:"RecoveredDiskInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfRecoveredDiskInfo", reflect.TypeOf((*ArrayOfRecoveredDiskInfo)(nil)).Elem())
}

// A boxed array of `RelatedStorageArray`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfRelatedStorageArray struct {
	RelatedStorageArray []RelatedStorageArray `xml:"RelatedStorageArray,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfRelatedStorageArray", reflect.TypeOf((*ArrayOfRelatedStorageArray)(nil)).Elem())
}

// A boxed array of `ReplicaIntervalQueryResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfReplicaIntervalQueryResult struct {
	ReplicaIntervalQueryResult []ReplicaIntervalQueryResult `xml:"ReplicaIntervalQueryResult,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfReplicaIntervalQueryResult", reflect.TypeOf((*ArrayOfReplicaIntervalQueryResult)(nil)).Elem())
}

// A boxed array of `ReplicationGroupData`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfReplicationGroupData struct {
	ReplicationGroupData []ReplicationGroupData `xml:"ReplicationGroupData,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfReplicationGroupData", reflect.TypeOf((*ArrayOfReplicationGroupData)(nil)).Elem())
}

// A boxed array of `ReplicationTargetInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfReplicationTargetInfo struct {
	ReplicationTargetInfo []ReplicationTargetInfo `xml:"ReplicationTargetInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfReplicationTargetInfo", reflect.TypeOf((*ArrayOfReplicationTargetInfo)(nil)).Elem())
}

// A boxed array of `SmsProviderInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfSmsProviderInfo struct {
	SmsProviderInfo []BaseSmsProviderInfo `xml:"SmsProviderInfo,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfSmsProviderInfo", reflect.TypeOf((*ArrayOfSmsProviderInfo)(nil)).Elem())
}

// A boxed array of `SourceGroupMemberInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfSourceGroupMemberInfo struct {
	SourceGroupMemberInfo []SourceGroupMemberInfo `xml:"SourceGroupMemberInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfSourceGroupMemberInfo", reflect.TypeOf((*ArrayOfSourceGroupMemberInfo)(nil)).Elem())
}

// A boxed array of `StorageAlarm`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageAlarm struct {
	StorageAlarm []StorageAlarm `xml:"StorageAlarm,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageAlarm", reflect.TypeOf((*ArrayOfStorageAlarm)(nil)).Elem())
}

// A boxed array of `StorageArray`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageArray struct {
	StorageArray []StorageArray `xml:"StorageArray,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageArray", reflect.TypeOf((*ArrayOfStorageArray)(nil)).Elem())
}

// A boxed array of `StorageContainer`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageContainer struct {
	StorageContainer []StorageContainer `xml:"StorageContainer,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageContainer", reflect.TypeOf((*ArrayOfStorageContainer)(nil)).Elem())
}

// A boxed array of `StorageFileSystem`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageFileSystem struct {
	StorageFileSystem []StorageFileSystem `xml:"StorageFileSystem,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageFileSystem", reflect.TypeOf((*ArrayOfStorageFileSystem)(nil)).Elem())
}

// A boxed array of `StorageFileSystemInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageFileSystemInfo struct {
	StorageFileSystemInfo []StorageFileSystemInfo `xml:"StorageFileSystemInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageFileSystemInfo", reflect.TypeOf((*ArrayOfStorageFileSystemInfo)(nil)).Elem())
}

// A boxed array of `StorageLun`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageLun struct {
	StorageLun []StorageLun `xml:"StorageLun,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageLun", reflect.TypeOf((*ArrayOfStorageLun)(nil)).Elem())
}

// A boxed array of `StoragePort`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStoragePort struct {
	StoragePort []BaseStoragePort `xml:"StoragePort,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStoragePort", reflect.TypeOf((*ArrayOfStoragePort)(nil)).Elem())
}

// A boxed array of `StorageProcessor`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfStorageProcessor struct {
	StorageProcessor []StorageProcessor `xml:"StorageProcessor,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfStorageProcessor", reflect.TypeOf((*ArrayOfStorageProcessor)(nil)).Elem())
}

// A boxed array of `SupportedVendorModelMapping`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfSupportedVendorModelMapping struct {
	SupportedVendorModelMapping []SupportedVendorModelMapping `xml:"SupportedVendorModelMapping,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfSupportedVendorModelMapping", reflect.TypeOf((*ArrayOfSupportedVendorModelMapping)(nil)).Elem())
}

// A boxed array of `TargetDeviceId`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfTargetDeviceId struct {
	TargetDeviceId []TargetDeviceId `xml:"TargetDeviceId,omitempty" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfTargetDeviceId", reflect.TypeOf((*ArrayOfTargetDeviceId)(nil)).Elem())
}

// A boxed array of `TargetGroupMemberInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/sms`.
type ArrayOfTargetGroupMemberInfo struct {
	TargetGroupMemberInfo []BaseTargetGroupMemberInfo `xml:"TargetGroupMemberInfo,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("sms:ArrayOfTargetGroupMemberInfo", reflect.TypeOf((*ArrayOfTargetGroupMemberInfo)(nil)).Elem())
}

// This exception is thrown when an error occurs while
// connecting to the vpxd service to validate the user
// credentials
//
// This structure may be used only with operations rendered under `/sms`.
type AuthConnectionFailed struct {
	types.NoPermission
}

func init() {
	types.Add("sms:AuthConnectionFailed", reflect.TypeOf((*AuthConnectionFailed)(nil)).Elem())
}

type AuthConnectionFailedFault AuthConnectionFailed

func init() {
	types.Add("sms:AuthConnectionFailedFault", reflect.TypeOf((*AuthConnectionFailedFault)(nil)).Elem())
}

// This data object represents SDRS related data associated with block device or file system.
//
// This structure may be used only with operations rendered under `/sms`.
type BackingConfig struct {
	types.DynamicData

	// Identifier for the backing pool for thin provisioning
	ThinProvisionBackingIdentifier string `xml:"thinProvisionBackingIdentifier,omitempty" json:"thinProvisionBackingIdentifier,omitempty"`
	// Identifier for the backing pool for deduplication
	DeduplicationBackingIdentifier string `xml:"deduplicationBackingIdentifier,omitempty" json:"deduplicationBackingIdentifier,omitempty"`
	// Flag to indicate whether auto-tiering optimizations are active
	AutoTieringEnabled *bool `xml:"autoTieringEnabled" json:"autoTieringEnabled,omitempty"`
	// Aggregate indication of space savings efficiency in the shared
	// deduplication pool.
	//
	// The value is between 0 and 100, higher values
	// indicating better efficiency.
	DeduplicationEfficiency int64 `xml:"deduplicationEfficiency,omitempty" json:"deduplicationEfficiency,omitempty"`
	// Frequency in seconds at which interval auto-tiering optimizations
	// are applied.
	//
	// A value of 0 indicates continuous optimization.
	PerformanceOptimizationInterval int64 `xml:"performanceOptimizationInterval,omitempty" json:"performanceOptimizationInterval,omitempty"`
}

func init() {
	types.Add("sms:BackingConfig", reflect.TypeOf((*BackingConfig)(nil)).Elem())
}

// This data object represents the backing storage pool information of block device or file system.
//
// This structure may be used only with operations rendered under `/sms`.
type BackingStoragePool struct {
	types.DynamicData

	// Unique identifier
	Uuid string `xml:"uuid" json:"uuid"`
	Type string `xml:"type" json:"type"`
	// Upper bound of the available capacity in the backing storage pool.
	CapacityInMB int64 `xml:"capacityInMB" json:"capacityInMB"`
	// Aggregate used space in the backing storage pool.
	UsedSpaceInMB int64 `xml:"usedSpaceInMB" json:"usedSpaceInMB"`
}

func init() {
	types.Add("sms:BackingStoragePool", reflect.TypeOf((*BackingStoragePool)(nil)).Elem())
}

// This exception is thrown if there is a problem with calls to the
// CertificateAuthority.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateAuthorityFault struct {
	ProviderRegistrationFault

	// Fault code returned by certificate authority.
	FaultCode int32 `xml:"faultCode" json:"faultCode"`
}

func init() {
	types.Add("sms:CertificateAuthorityFault", reflect.TypeOf((*CertificateAuthorityFault)(nil)).Elem())
}

type CertificateAuthorityFaultFault CertificateAuthorityFault

func init() {
	types.Add("sms:CertificateAuthorityFaultFault", reflect.TypeOf((*CertificateAuthorityFaultFault)(nil)).Elem())
}

// This exception is thrown if `VasaProviderInfo.retainVasaProviderCertificate`
// is true and the provider uses a certificate issued by a Certificate Authority,
// but the root certificate of the Certificate Authority is not imported to VECS truststore
// before attempting the provider registration.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateNotImported struct {
	ProviderRegistrationFault
}

func init() {
	types.Add("sms:CertificateNotImported", reflect.TypeOf((*CertificateNotImported)(nil)).Elem())
}

type CertificateNotImportedFault CertificateNotImported

func init() {
	types.Add("sms:CertificateNotImportedFault", reflect.TypeOf((*CertificateNotImportedFault)(nil)).Elem())
}

// This exception is thrown if the certificate provided by the
// provider is not trusted.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateNotTrusted struct {
	ProviderRegistrationFault

	// Certificate
	Certificate string `xml:"certificate" json:"certificate"`
}

func init() {
	types.Add("sms:CertificateNotTrusted", reflect.TypeOf((*CertificateNotTrusted)(nil)).Elem())
}

// An CertificateNotTrusted fault is thrown when an Agency's configuration
// contains OVF package URL or VIB URL for that vSphere ESX Agent Manager is not
// able to make successful SSL trust verification of the server's certificate.
//
// Reasons for this might be that the certificate provided via the API
// `AgentConfigInfo.ovfSslTrust` and `AgentConfigInfo.vibSslTrust`
// or via the script /usr/lib/vmware-eam/bin/eam-utility.py
//   - is invalid.
//   - does not match the server's certificate.
//
// If there is no provided certificate, the fault is thrown when the server's
// certificate is not trusted by the system or is invalid - @see
// `AgentConfigInfo.ovfSslTrust` and
// `AgentConfigInfo.vibSslTrust`.
// To enable Agency creation 1) provide a valid certificate used by the
// server hosting the `AgentConfigInfo.ovfPackageUrl` or
// `AgentConfigInfo.vibUrl` or 2) ensure the server's certificate is
// signed by a CA trusted by the system. Then retry the operation, vSphere
// ESX Agent Manager will retry the SSL trust verification and proceed with
// reaching the desired state.
//
// This structure may be used only with operations rendered under `/eam`.
//
// `**Since:**` vEAM API 8.2
type CertificateNotTrustedFault CertificateNotTrusted

func init() {
	types.Add("sms:CertificateNotTrustedFault", reflect.TypeOf((*CertificateNotTrustedFault)(nil)).Elem())
}

// This exception is thrown if SMS failed to
// refresh a CA signed certificate for the provider (or)
// push the latest CA root certificates and CRLs to the provider.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateRefreshFailed struct {
	types.MethodFault

	ProviderId []string `xml:"providerId,omitempty" json:"providerId,omitempty"`
}

func init() {
	types.Add("sms:CertificateRefreshFailed", reflect.TypeOf((*CertificateRefreshFailed)(nil)).Elem())
}

type CertificateRefreshFailedFault CertificateRefreshFailed

func init() {
	types.Add("sms:CertificateRefreshFailedFault", reflect.TypeOf((*CertificateRefreshFailedFault)(nil)).Elem())
}

// This exception is thrown if SMS failed to revoke CA signed certificate of the provider.
//
// This structure may be used only with operations rendered under `/sms`.
type CertificateRevocationFailed struct {
	types.MethodFault
}

func init() {
	types.Add("sms:CertificateRevocationFailed", reflect.TypeOf((*CertificateRevocationFailed)(nil)).Elem())
}

type CertificateRevocationFailedFault CertificateRevocationFailed

func init() {
	types.Add("sms:CertificateRevocationFailedFault", reflect.TypeOf((*CertificateRevocationFailedFault)(nil)).Elem())
}

// This data object represents the result of queryDatastoreBackingPoolMapping API.
//
// More than one datastore can map to the same set of BackingStoragePool.
//
// This structure may be used only with operations rendered under `/sms`.
type DatastoreBackingPoolMapping struct {
	types.DynamicData

	// Refers instances of `Datastore`.
	Datastore          []types.ManagedObjectReference `xml:"datastore" json:"datastore"`
	BackingStoragePool []BackingStoragePool           `xml:"backingStoragePool,omitempty" json:"backingStoragePool,omitempty"`
}

func init() {
	types.Add("sms:DatastoreBackingPoolMapping", reflect.TypeOf((*DatastoreBackingPoolMapping)(nil)).Elem())
}

// Deprecated as of SMS API 5.0.
//
// Datastore pair that is returned as a result of queryDrsMigrationCapabilityForPerformanceEx API.
//
// This structure may be used only with operations rendered under `/sms`.
type DatastorePair struct {
	types.DynamicData

	// Refers instance of `Datastore`.
	Datastore1 types.ManagedObjectReference `xml:"datastore1" json:"datastore1"`
	// Refers instance of `Datastore`.
	Datastore2 types.ManagedObjectReference `xml:"datastore2" json:"datastore2"`
}

func init() {
	types.Add("sms:DatastorePair", reflect.TypeOf((*DatastorePair)(nil)).Elem())
}

// Base class that represents a replicated device.
//
// This structure may be used only with operations rendered under `/sms`.
type DeviceId struct {
	types.DynamicData
}

func init() {
	types.Add("sms:DeviceId", reflect.TypeOf((*DeviceId)(nil)).Elem())
}

// Deprecated as of SMS API 5.0.
//
// This data object represents the result of queryDrsMigrationCapabilityForPerformanceEx API.
//
// This structure may be used only with operations rendered under `/sms`.
type DrsMigrationCapabilityResult struct {
	types.DynamicData

	RecommendedDatastorePair    []DatastorePair `xml:"recommendedDatastorePair,omitempty" json:"recommendedDatastorePair,omitempty"`
	NonRecommendedDatastorePair []DatastorePair `xml:"nonRecommendedDatastorePair,omitempty" json:"nonRecommendedDatastorePair,omitempty"`
}

func init() {
	types.Add("sms:DrsMigrationCapabilityResult", reflect.TypeOf((*DrsMigrationCapabilityResult)(nil)).Elem())
}

// This exception indicates there are duplicate entries in the input argument.
//
// This structure may be used only with operations rendered under `/sms`.
type DuplicateEntry struct {
	types.MethodFault
}

func init() {
	types.Add("sms:DuplicateEntry", reflect.TypeOf((*DuplicateEntry)(nil)).Elem())
}

type DuplicateEntryFault DuplicateEntry

func init() {
	types.Add("sms:DuplicateEntryFault", reflect.TypeOf((*DuplicateEntryFault)(nil)).Elem())
}

// Deprecated as of SMS API 4.0.
//
// Unique identifier of a given entity with the storage
// management service.
//
// It is similar to the VirtualCenter
// ManagedObjectReference but also identifies certain
// non-managed objects.
//
// This structure may be used only with operations rendered under `/sms`.
type EntityReference struct {
	types.DynamicData

	// Unique identifier for the entity of a given type in the
	// system.
	//
	// A VirtualCenter managed object ID can be supplied
	// here, in which case the type may be unset. Otherwise, the
	// type must be set.
	Id string `xml:"id" json:"id"`
	// Type of the entity.
	Type EntityReferenceEntityType `xml:"type,omitempty" json:"type,omitempty"`
}

func init() {
	types.Add("sms:EntityReference", reflect.TypeOf((*EntityReference)(nil)).Elem())
}

// Input to the failover or testFailover methods.
//
// This structure may be used only with operations rendered under `/sms`.
type FailoverParam struct {
	types.DynamicData

	// Whether the failover is a planned failover or not.
	//
	// Note that testFailover
	// can also be executed in an unplanned mode. When this flag is
	// set to false, the recovery VASA provider must not try to connect
	// to the primary VASA provider during the failover.
	IsPlanned bool `xml:"isPlanned" json:"isPlanned"`
	// Do not execute the (test) failover but check if the configuration
	// is correct to execute the (test) failover.
	//
	// If set to <code>true</code>, the (test)failover result is an array where
	// each element is either `GroupOperationResult` or `GroupErrorResult`.
	//
	// If set to <code>false</code>, the (test)failover result is an array where
	// each element is either `FailoverSuccessResult` or `GroupErrorResult`.
	CheckOnly bool `xml:"checkOnly" json:"checkOnly"`
	// The replication groups to failover.
	//
	// It is OK for the VASA
	// provider to successfully failover only some groups. The
	// groups that did not complete will be retried.
	ReplicationGroupsToFailover []ReplicationGroupData `xml:"replicationGroupsToFailover,omitempty" json:"replicationGroupsToFailover,omitempty"`
	// Storage policies for the devices after (test)failover.
	//
	// Failover should be done even if policies cannot be associated.
	// Test failover, however, should fail if policies cannot be associated.
	//
	// If policies cannot be associated, VASA provider can notify the client by
	// doing either or both of these:
	// 1\. Set the warning in the result for a replication group to indicate
	// such a failure to set the policy.
	// 2\. Raise a compliance alarm after the failover is done.
	//
	// If not specified, the default policies are used. Callers may reassign
	// policy later.
	PolicyAssociations []PolicyAssociation `xml:"policyAssociations,omitempty" json:"policyAssociations,omitempty"`
}

func init() {
	types.Add("sms:FailoverParam", reflect.TypeOf((*FailoverParam)(nil)).Elem())
}

// The parameters of `VasaProvider.FailoverReplicationGroup_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type FailoverReplicationGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Settings for the failover.
	FailoverParam BaseFailoverParam `xml:"failoverParam,typeattr" json:"failoverParam"`
}

func init() {
	types.Add("sms:FailoverReplicationGroupRequestType", reflect.TypeOf((*FailoverReplicationGroupRequestType)(nil)).Elem())
}

type FailoverReplicationGroup_Task FailoverReplicationGroupRequestType

func init() {
	types.Add("sms:FailoverReplicationGroup_Task", reflect.TypeOf((*FailoverReplicationGroup_Task)(nil)).Elem())
}

type FailoverReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Results of a successful failover operation.
//
// The target fault domain Id, and the device group id are inherited.
//
// This structure may be used only with operations rendered under `/sms`.
type FailoverSuccessResult struct {
	GroupOperationResult

	// Some replicators may automatically reverse replication on failover.
	//
	// Such
	// replicators must move the replication status to
	// `SOURCE`
	// In other cases, it can remain as `FAILEDOVER`.
	NewState string `xml:"newState" json:"newState"`
	// Id of the Point in Time snapshot used during failover.
	//
	// If not present,
	// latest PIT was used.
	PitId *PointInTimeReplicaId `xml:"pitId,omitempty" json:"pitId,omitempty"`
	// Optional id of the Point in Time snapshot that was automatically created before
	// failing over.
	//
	// This is recommended so users can revert back to this
	// snapshot to avoid data loss. This can be removed after the reverse
	// replication call succeeds.
	PitIdBeforeFailover *PointInTimeReplicaId `xml:"pitIdBeforeFailover,omitempty" json:"pitIdBeforeFailover,omitempty"`
	// Recovered Devices.
	//
	// This is optional because in some corner cases the
	// replication groups on the target site may not have any virtual volumes.
	RecoveredDeviceInfo []RecoveredDevice `xml:"recoveredDeviceInfo,omitempty" json:"recoveredDeviceInfo,omitempty"`
	// Time stamp of recovery.
	TimeStamp *time.Time `xml:"timeStamp" json:"timeStamp,omitempty"`
}

func init() {
	types.Add("sms:FailoverSuccessResult", reflect.TypeOf((*FailoverSuccessResult)(nil)).Elem())
}

// This spec contains information needed for `SmsStorageManager.QueryFaultDomain`
// API to filter the result.
//
// This structure may be used only with operations rendered under `/sms`.
type FaultDomainFilter struct {
	types.DynamicData

	// If specified, query for this specific provider only; else query for all
	// providers.
	ProviderId string `xml:"providerId,omitempty" json:"providerId,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainFilter", reflect.TypeOf((*FaultDomainFilter)(nil)).Elem())
}

// Information about a Fault Domain.
//
// This structure may be used only with operations rendered under `/sms`.
type FaultDomainInfo struct {
	types.FaultDomainId

	// Name of the fault domain, if not specified, the id will be used in place
	// of the name.
	//
	// Name need not be unique.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	// Description - could be a localized string.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// Identifier of the Storage Array that this Fault Domain belongs to.
	//
	// A Fault
	// Domain and all its children should report same `FaultDomainInfo.storageArrayId`. It
	// can be left unspecified. If not specified, vSphere will not support High
	// Availability feature for this Fault Domain. When specified, vSphere will
	// treat the the currently active VASA provider for `StorageArray` as
	// the active VASA provider for this Fault Domain and its children.
	// Changing High Availability support choice for a Fault Domain
	// intermittently, by sometimes specifying the storageArrayId and sometimes
	// not, will cause unexpected result and might cause VP to be in 'syncError'
	// state in vSphere.
	StorageArrayId string `xml:"storageArrayId,omitempty" json:"storageArrayId,omitempty"`
	// List of children, the entries in the array should always be
	// `FaultDomainId` and not `FaultDomainInfo`.
	//
	// The 2016 vSphere release will not support nested Fault Domains. The field
	// FaultDomainInfo#children is ignored by vSphere 2016 release.
	Children []types.FaultDomainId `xml:"children,omitempty" json:"children,omitempty"`
	// VASA provider that is actively managing this fault domain
	//
	// Refers instance of `SmsProvider`.
	Provider *types.ManagedObjectReference `xml:"provider,omitempty" json:"provider,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainInfo", reflect.TypeOf((*FaultDomainInfo)(nil)).Elem())
}

// This mapping will be set in InactiveProvider fault to notify clients
// the current active provider for the specified fault domains.
//
// This structure may be used only with operations rendered under `/sms`.
type FaultDomainProviderMapping struct {
	types.DynamicData

	// Active provider managing the fault domains
	//
	// Refers instance of `SmsProvider`.
	ActiveProvider types.ManagedObjectReference `xml:"activeProvider" json:"activeProvider"`
	// Fault domains being managed by the provider
	FaultDomainId []types.FaultDomainId `xml:"faultDomainId,omitempty" json:"faultDomainId,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainProviderMapping", reflect.TypeOf((*FaultDomainProviderMapping)(nil)).Elem())
}

// This data object represents the FC storage port.
//
// This structure may be used only with operations rendered under `/sms`.
type FcStoragePort struct {
	StoragePort

	// World Wide Name for the Port
	PortWwn string `xml:"portWwn" json:"portWwn"`
	// World Wide Name for the Node
	NodeWwn string `xml:"nodeWwn" json:"nodeWwn"`
}

func init() {
	types.Add("sms:FcStoragePort", reflect.TypeOf((*FcStoragePort)(nil)).Elem())
}

// This data object represents the FCoE storage port.
//
// This structure may be used only with operations rendered under `/sms`.
type FcoeStoragePort struct {
	StoragePort

	// World Wide Name for the Port
	PortWwn string `xml:"portWwn" json:"portWwn"`
	// World Wide Name for the Node
	NodeWwn string `xml:"nodeWwn" json:"nodeWwn"`
}

func init() {
	types.Add("sms:FcoeStoragePort", reflect.TypeOf((*FcoeStoragePort)(nil)).Elem())
}

// Error result.
//
// This structure may be used only with operations rendered under `/sms`.
type GroupErrorResult struct {
	GroupOperationResult

	// Error array, must contain at least one entry.
	Error []types.LocalizedMethodFault `xml:"error" json:"error"`
}

func init() {
	types.Add("sms:GroupErrorResult", reflect.TypeOf((*GroupErrorResult)(nil)).Elem())
}

// Replication group information.
//
// May be either a `SourceGroupInfo` or
// `TargetGroupInfo`.
//
// This structure may be used only with operations rendered under `/sms`.
type GroupInfo struct {
	types.DynamicData

	// Identifier of the group + fault domain id.
	GroupId types.ReplicationGroupId `xml:"groupId" json:"groupId"`
}

func init() {
	types.Add("sms:GroupInfo", reflect.TypeOf((*GroupInfo)(nil)).Elem())
}

// The base class for any operation on a replication group.
//
// Usually, there is an
// operation specific &lt;Operation&gt;SuccessResult
//
// This structure may be used only with operations rendered under `/sms`.
type GroupOperationResult struct {
	types.DynamicData

	// Replication group Id.
	GroupId types.ReplicationGroupId     `xml:"groupId" json:"groupId"`
	Warning []types.LocalizedMethodFault `xml:"warning,omitempty" json:"warning,omitempty"`
}

func init() {
	types.Add("sms:GroupOperationResult", reflect.TypeOf((*GroupOperationResult)(nil)).Elem())
}

// Thrown if the VASA Provider on which the call is made is currently not
// active.
//
// If the client maintains a cache of the topology of fault domains
// and replication groups, it's expected to update the cache based on the
// mapping information set in this fault.
//
// This structure may be used only with operations rendered under `/sms`.
type InactiveProvider struct {
	types.MethodFault

	// Mapping between VASA provider and the fault domains
	Mapping []FaultDomainProviderMapping `xml:"mapping,omitempty" json:"mapping,omitempty"`
}

func init() {
	types.Add("sms:InactiveProvider", reflect.TypeOf((*InactiveProvider)(nil)).Elem())
}

type InactiveProviderFault InactiveProvider

func init() {
	types.Add("sms:InactiveProviderFault", reflect.TypeOf((*InactiveProviderFault)(nil)).Elem())
}

// This fault is thrown if failed to register provider due to incorrect credentials.
//
// This structure may be used only with operations rendered under `/sms`.
type IncorrectUsernamePassword struct {
	ProviderRegistrationFault
}

func init() {
	types.Add("sms:IncorrectUsernamePassword", reflect.TypeOf((*IncorrectUsernamePassword)(nil)).Elem())
}

type IncorrectUsernamePasswordFault IncorrectUsernamePassword

func init() {
	types.Add("sms:IncorrectUsernamePasswordFault", reflect.TypeOf((*IncorrectUsernamePasswordFault)(nil)).Elem())
}

// This exception is thrown if the provider certificate is empty, malformed,
// expired, not yet valid, revoked or fails host name verification.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidCertificate struct {
	ProviderRegistrationFault

	// Provider certificate
	Certificate string `xml:"certificate" json:"certificate"`
}

func init() {
	types.Add("sms:InvalidCertificate", reflect.TypeOf((*InvalidCertificate)(nil)).Elem())
}

type InvalidCertificateFault InvalidCertificate

func init() {
	types.Add("sms:InvalidCertificateFault", reflect.TypeOf((*InvalidCertificateFault)(nil)).Elem())
}

// Thrown if the function is called at the wrong end of the replication (i.e.
//
// the failing function should be tried at the opposite end of replication).
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidFunctionTarget struct {
	SmsReplicationFault
}

func init() {
	types.Add("sms:InvalidFunctionTarget", reflect.TypeOf((*InvalidFunctionTarget)(nil)).Elem())
}

type InvalidFunctionTargetFault InvalidFunctionTarget

func init() {
	types.Add("sms:InvalidFunctionTargetFault", reflect.TypeOf((*InvalidFunctionTargetFault)(nil)).Elem())
}

// This exception is thrown if the specified storage profile is invalid.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidProfile struct {
	types.MethodFault
}

func init() {
	types.Add("sms:InvalidProfile", reflect.TypeOf((*InvalidProfile)(nil)).Elem())
}

type InvalidProfileFault InvalidProfile

func init() {
	types.Add("sms:InvalidProfileFault", reflect.TypeOf((*InvalidProfileFault)(nil)).Elem())
}

// Thrown if the replication group is not in the correct state.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidReplicationState struct {
	SmsReplicationFault

	// States where the operation would have been successful.
	DesiredState []string `xml:"desiredState,omitempty" json:"desiredState,omitempty"`
	// Current state.
	CurrentState string `xml:"currentState" json:"currentState"`
}

func init() {
	types.Add("sms:InvalidReplicationState", reflect.TypeOf((*InvalidReplicationState)(nil)).Elem())
}

type InvalidReplicationStateFault InvalidReplicationState

func init() {
	types.Add("sms:InvalidReplicationStateFault", reflect.TypeOf((*InvalidReplicationStateFault)(nil)).Elem())
}

// This exception is thrown if a specified session is invalid.
//
// This can occur if the VirtualCenter session referred to by
// the cookie has timed out or has been closed.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidSession struct {
	types.NoPermission

	// VirtualCenter session cookie that is invalid.
	SessionCookie string `xml:"sessionCookie" json:"sessionCookie"`
}

func init() {
	types.Add("sms:InvalidSession", reflect.TypeOf((*InvalidSession)(nil)).Elem())
}

type InvalidSessionFault InvalidSession

func init() {
	types.Add("sms:InvalidSessionFault", reflect.TypeOf((*InvalidSessionFault)(nil)).Elem())
}

// This exception is thrown if `VasaProviderSpec.url` is malformed.
//
// This structure may be used only with operations rendered under `/sms`.
type InvalidUrl struct {
	ProviderRegistrationFault

	// Provider `VasaProviderSpec.url`
	Url string `xml:"url" json:"url"`
}

func init() {
	types.Add("sms:InvalidUrl", reflect.TypeOf((*InvalidUrl)(nil)).Elem())
}

type InvalidUrlFault InvalidUrl

func init() {
	types.Add("sms:InvalidUrlFault", reflect.TypeOf((*InvalidUrlFault)(nil)).Elem())
}

// This data object represents the iSCSI storage port.
//
// This structure may be used only with operations rendered under `/sms`.
type IscsiStoragePort struct {
	StoragePort

	// IQN or EQI identifier
	Identifier string `xml:"identifier" json:"identifier"`
}

func init() {
	types.Add("sms:IscsiStoragePort", reflect.TypeOf((*IscsiStoragePort)(nil)).Elem())
}

// This data object represents the lun, HBA association
// for synchronous replication.
//
// This structure may be used only with operations rendered under `/sms`.
type LunHbaAssociation struct {
	types.DynamicData

	CanonicalName string                     `xml:"canonicalName" json:"canonicalName"`
	Hba           []types.HostHostBusAdapter `xml:"hba" json:"hba"`
}

func init() {
	types.Add("sms:LunHbaAssociation", reflect.TypeOf((*LunHbaAssociation)(nil)).Elem())
}

// This exception is thrown if more than one sort spec is
// specified in a list query.
//
// This structure may be used only with operations rendered under `/sms`.
type MultipleSortSpecsNotSupported struct {
	types.InvalidArgument
}

func init() {
	types.Add("sms:MultipleSortSpecsNotSupported", reflect.TypeOf((*MultipleSortSpecsNotSupported)(nil)).Elem())
}

type MultipleSortSpecsNotSupportedFault MultipleSortSpecsNotSupported

func init() {
	types.Add("sms:MultipleSortSpecsNotSupportedFault", reflect.TypeOf((*MultipleSortSpecsNotSupportedFault)(nil)).Elem())
}

// This data object represents a name value pair.
//
// This structure may be used only with operations rendered under `/sms`.
type NameValuePair struct {
	types.DynamicData

	// Name of the parameter
	ParameterName string `xml:"parameterName" json:"parameterName"`
	// Value of the parameter
	ParameterValue string `xml:"parameterValue" json:"parameterValue"`
}

func init() {
	types.Add("sms:NameValuePair", reflect.TypeOf((*NameValuePair)(nil)).Elem())
}

// This fault is thrown when backings (@link
// sms.storage.StorageLun/ @link sms.storage.StorageFileSystem)
// of the specified datastores refer to different
// VASA providers.
//
// This structure may be used only with operations rendered under `/sms`.
type NoCommonProviderForAllBackings struct {
	QueryExecutionFault
}

func init() {
	types.Add("sms:NoCommonProviderForAllBackings", reflect.TypeOf((*NoCommonProviderForAllBackings)(nil)).Elem())
}

type NoCommonProviderForAllBackingsFault NoCommonProviderForAllBackings

func init() {
	types.Add("sms:NoCommonProviderForAllBackingsFault", reflect.TypeOf((*NoCommonProviderForAllBackingsFault)(nil)).Elem())
}

// This exception is set if the replication target is not yet available.
//
// This structure may be used only with operations rendered under `/sms`.
type NoReplicationTarget struct {
	SmsReplicationFault
}

func init() {
	types.Add("sms:NoReplicationTarget", reflect.TypeOf((*NoReplicationTarget)(nil)).Elem())
}

type NoReplicationTargetFault NoReplicationTarget

func init() {
	types.Add("sms:NoReplicationTargetFault", reflect.TypeOf((*NoReplicationTargetFault)(nil)).Elem())
}

// This exception is thrown when there is no valid replica
// to be used in recovery.
//
// This may happen when a Virtual Volume
// is created on the source domain, but the replica is yet to
// be copied to the target.
//
// This structure may be used only with operations rendered under `/sms`.
type NoValidReplica struct {
	SmsReplicationFault

	// Identifier of the device which does not have a valid
	// replica.
	//
	// This is the identifier on the target site.
	// This may not be set if the ReplicationGroup does not
	// have even a single valid replica.
	DeviceId BaseDeviceId `xml:"deviceId,omitempty,typeattr" json:"deviceId,omitempty"`
}

func init() {
	types.Add("sms:NoValidReplica", reflect.TypeOf((*NoValidReplica)(nil)).Elem())
}

type NoValidReplicaFault NoValidReplica

func init() {
	types.Add("sms:NoValidReplicaFault", reflect.TypeOf((*NoValidReplicaFault)(nil)).Elem())
}

// This exception is thrown if the VASA Provider on which the call is made
// does not support this operation.
//
// This structure may be used only with operations rendered under `/sms`.
type NotSupportedByProvider struct {
	types.MethodFault
}

func init() {
	types.Add("sms:NotSupportedByProvider", reflect.TypeOf((*NotSupportedByProvider)(nil)).Elem())
}

type NotSupportedByProviderFault NotSupportedByProvider

func init() {
	types.Add("sms:NotSupportedByProviderFault", reflect.TypeOf((*NotSupportedByProviderFault)(nil)).Elem())
}

// This exception is set if the replication peer is not reachable.
//
// For prepareFailover, it is the target that is not reachable.
// For other functions, it is the source that is not reachable.
//
// This structure may be used only with operations rendered under `/sms`.
type PeerNotReachable struct {
	SmsReplicationFault
}

func init() {
	types.Add("sms:PeerNotReachable", reflect.TypeOf((*PeerNotReachable)(nil)).Elem())
}

type PeerNotReachableFault PeerNotReachable

func init() {
	types.Add("sms:PeerNotReachableFault", reflect.TypeOf((*PeerNotReachableFault)(nil)).Elem())
}

// Identity of the Point in Time Replica object.
//
// This structure may be used only with operations rendered under `/sms`.
type PointInTimeReplicaId struct {
	types.DynamicData

	// ID of the Point In Time replica.
	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("sms:PointInTimeReplicaId", reflect.TypeOf((*PointInTimeReplicaId)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/sms`.
type PointInTimeReplicaInfo struct {
	types.DynamicData

	// Id of the PIT replica.
	//
	// Note that this id is always used
	// in combination with the `ReplicationGroupId`, hence must be
	// unique to the `ReplicationGroupId`.
	Id PointInTimeReplicaId `xml:"id" json:"id"`
	// Name of the PIT replica.
	//
	// This may be a localized string
	// in a language as chosen by the VASA provider.
	PitName string `xml:"pitName" json:"pitName"`
	// Time when the snapshot was taken.
	//
	// Time stamps are maintained by
	// the Replication provider, note that this carries time zone information
	// as well.
	TimeStamp time.Time `xml:"timeStamp" json:"timeStamp"`
	// VASA provider managed tags associated with the replica.
	Tags []string `xml:"tags,omitempty" json:"tags,omitempty"`
}

func init() {
	types.Add("sms:PointInTimeReplicaInfo", reflect.TypeOf((*PointInTimeReplicaInfo)(nil)).Elem())
}

// Describes the policy association object.
//
// This structure may be used only with operations rendered under `/sms`.
type PolicyAssociation struct {
	types.DynamicData

	// The source device id.
	//
	// The corresponding recovered device
	// gets the specified <code>policyId</code>.
	Id BaseDeviceId `xml:"id,typeattr" json:"id"`
	// Policy id.
	PolicyId string `xml:"policyId" json:"policyId"`
	// Datastore object.
	//
	// Refers instance of `Datastore`.
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("sms:PolicyAssociation", reflect.TypeOf((*PolicyAssociation)(nil)).Elem())
}

// The parameters of `VasaProvider.PrepareFailoverReplicationGroup_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type PrepareFailoverReplicationGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// List of replication group IDs.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
}

func init() {
	types.Add("sms:PrepareFailoverReplicationGroupRequestType", reflect.TypeOf((*PrepareFailoverReplicationGroupRequestType)(nil)).Elem())
}

type PrepareFailoverReplicationGroup_Task PrepareFailoverReplicationGroupRequestType

func init() {
	types.Add("sms:PrepareFailoverReplicationGroup_Task", reflect.TypeOf((*PrepareFailoverReplicationGroup_Task)(nil)).Elem())
}

type PrepareFailoverReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Input to the promoteReplicationGroup method.
//
// This structure may be used only with operations rendered under `/sms`.
type PromoteParam struct {
	types.DynamicData

	// Specifies whether the promote operation is a planned one.
	//
	// When this flag is set to false, the recovery VASA provider must not
	// try to connect to the primary VASA provider during promote.
	IsPlanned bool `xml:"isPlanned" json:"isPlanned"`
	// The replication groups to promote.
	//
	// It is legal for the VASA
	// provider to successfully promote only some groups. The
	// groups that did not succeed will be retried.
	//
	// The identifiers of the Virtual Volumes do not change after the
	// promote operation.
	ReplicationGroupsToPromote []types.ReplicationGroupId `xml:"replicationGroupsToPromote,omitempty" json:"replicationGroupsToPromote,omitempty"`
}

func init() {
	types.Add("sms:PromoteParam", reflect.TypeOf((*PromoteParam)(nil)).Elem())
}

// The parameters of `VasaProvider.PromoteReplicationGroup_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type PromoteReplicationGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Specifies an array of replication group IDs whose
	// in-test devices (`INTEST`) need to be
	// promoted to failover `FAILEDOVER` state.
	PromoteParam PromoteParam `xml:"promoteParam" json:"promoteParam"`
}

func init() {
	types.Add("sms:PromoteReplicationGroupRequestType", reflect.TypeOf((*PromoteReplicationGroupRequestType)(nil)).Elem())
}

type PromoteReplicationGroup_Task PromoteReplicationGroupRequestType

func init() {
	types.Add("sms:PromoteReplicationGroup_Task", reflect.TypeOf((*PromoteReplicationGroup_Task)(nil)).Elem())
}

type PromoteReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// This exception is thrown if the VASA Provider on which the call is made
// is currently busy.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderBusy struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderBusy", reflect.TypeOf((*ProviderBusy)(nil)).Elem())
}

type ProviderBusyFault ProviderBusy

func init() {
	types.Add("sms:ProviderBusyFault", reflect.TypeOf((*ProviderBusyFault)(nil)).Elem())
}

// This fault is thrown if the Storage Monitoring Service failed to connect to the VASA provider.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderConnectionFailed struct {
	types.RuntimeFault
}

func init() {
	types.Add("sms:ProviderConnectionFailed", reflect.TypeOf((*ProviderConnectionFailed)(nil)).Elem())
}

type ProviderConnectionFailedFault ProviderConnectionFailed

func init() {
	types.Add("sms:ProviderConnectionFailedFault", reflect.TypeOf((*ProviderConnectionFailedFault)(nil)).Elem())
}

// This fault is thrown when a VASA provider cannot be found for the specified
// entities.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderNotFound struct {
	QueryExecutionFault
}

func init() {
	types.Add("sms:ProviderNotFound", reflect.TypeOf((*ProviderNotFound)(nil)).Elem())
}

type ProviderNotFoundFault ProviderNotFound

func init() {
	types.Add("sms:ProviderNotFoundFault", reflect.TypeOf((*ProviderNotFoundFault)(nil)).Elem())
}

// This exception is thrown if the VASA Provider is out of resource to satisfy
// a provisioning request.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderOutOfProvisioningResource struct {
	types.MethodFault

	// Identifier of the provisioning resource.
	ProvisioningResourceId string `xml:"provisioningResourceId" json:"provisioningResourceId"`
	// Currently available.
	AvailableBefore int64 `xml:"availableBefore,omitempty" json:"availableBefore,omitempty"`
	// Necessary for provisioning.
	AvailableAfter int64 `xml:"availableAfter,omitempty" json:"availableAfter,omitempty"`
	// Total amount (free + used).
	Total int64 `xml:"total,omitempty" json:"total,omitempty"`
	// This resource limitation is transient, i.e.
	//
	// the resource
	// will be available after some time.
	IsTransient *bool `xml:"isTransient" json:"isTransient,omitempty"`
}

func init() {
	types.Add("sms:ProviderOutOfProvisioningResource", reflect.TypeOf((*ProviderOutOfProvisioningResource)(nil)).Elem())
}

type ProviderOutOfProvisioningResourceFault ProviderOutOfProvisioningResource

func init() {
	types.Add("sms:ProviderOutOfProvisioningResourceFault", reflect.TypeOf((*ProviderOutOfProvisioningResourceFault)(nil)).Elem())
}

// This exception is thrown if the VASA Provider on which the call is made
// is out of resource.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderOutOfResource struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderOutOfResource", reflect.TypeOf((*ProviderOutOfResource)(nil)).Elem())
}

type ProviderOutOfResourceFault ProviderOutOfResource

func init() {
	types.Add("sms:ProviderOutOfResourceFault", reflect.TypeOf((*ProviderOutOfResourceFault)(nil)).Elem())
}

// This fault is thrown if failed to register provider to storage
// management service.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderRegistrationFault struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderRegistrationFault", reflect.TypeOf((*ProviderRegistrationFault)(nil)).Elem())
}

type ProviderRegistrationFaultFault BaseProviderRegistrationFault

func init() {
	types.Add("sms:ProviderRegistrationFaultFault", reflect.TypeOf((*ProviderRegistrationFaultFault)(nil)).Elem())
}

// Thrown if a failure occurs when synchronizing the service
// cache with provider information.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderSyncFailed struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderSyncFailed", reflect.TypeOf((*ProviderSyncFailed)(nil)).Elem())
}

type ProviderSyncFailedFault BaseProviderSyncFailed

func init() {
	types.Add("sms:ProviderSyncFailedFault", reflect.TypeOf((*ProviderSyncFailedFault)(nil)).Elem())
}

// This exception is thrown if the VASA Provider on which the call is made is
// currently not available, e.g.
//
// VASA Provider is in offline state. This error
// usually means the provider is temporarily unavailable due to network outage, etc.
// The client is expected to wait for some time and retry the same call.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderUnavailable struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderUnavailable", reflect.TypeOf((*ProviderUnavailable)(nil)).Elem())
}

type ProviderUnavailableFault ProviderUnavailable

func init() {
	types.Add("sms:ProviderUnavailableFault", reflect.TypeOf((*ProviderUnavailableFault)(nil)).Elem())
}

// This fault is thrown if failed to unregister provider from storage
// management service.
//
// This structure may be used only with operations rendered under `/sms`.
type ProviderUnregistrationFault struct {
	types.MethodFault
}

func init() {
	types.Add("sms:ProviderUnregistrationFault", reflect.TypeOf((*ProviderUnregistrationFault)(nil)).Elem())
}

type ProviderUnregistrationFaultFault ProviderUnregistrationFault

func init() {
	types.Add("sms:ProviderUnregistrationFaultFault", reflect.TypeOf((*ProviderUnregistrationFaultFault)(nil)).Elem())
}

// This exception is thrown if the storage management service
// fails to register with the VirtualCenter proxy during
// initialization.
//
// This structure may be used only with operations rendered under `/sms`.
type ProxyRegistrationFailed struct {
	types.RuntimeFault
}

func init() {
	types.Add("sms:ProxyRegistrationFailed", reflect.TypeOf((*ProxyRegistrationFailed)(nil)).Elem())
}

type ProxyRegistrationFailedFault ProxyRegistrationFailed

func init() {
	types.Add("sms:ProxyRegistrationFailedFault", reflect.TypeOf((*ProxyRegistrationFailedFault)(nil)).Elem())
}

type QueryAboutInfo QueryAboutInfoRequestType

func init() {
	types.Add("sms:QueryAboutInfo", reflect.TypeOf((*QueryAboutInfo)(nil)).Elem())
}

type QueryAboutInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QueryAboutInfoRequestType", reflect.TypeOf((*QueryAboutInfoRequestType)(nil)).Elem())
}

type QueryAboutInfoResponse struct {
	Returnval SmsAboutInfo `xml:"returnval" json:"returnval"`
}

type QueryActiveAlarm QueryActiveAlarmRequestType

func init() {
	types.Add("sms:QueryActiveAlarm", reflect.TypeOf((*QueryActiveAlarm)(nil)).Elem())
}

// The parameters of `VasaProvider.QueryActiveAlarm`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryActiveAlarmRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Filter criteria for the alarm state.
	AlarmFilter *AlarmFilter `xml:"alarmFilter,omitempty" json:"alarmFilter,omitempty"`
}

func init() {
	types.Add("sms:QueryActiveAlarmRequestType", reflect.TypeOf((*QueryActiveAlarmRequestType)(nil)).Elem())
}

type QueryActiveAlarmResponse struct {
	Returnval *AlarmResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryArray QueryArrayRequestType

func init() {
	types.Add("sms:QueryArray", reflect.TypeOf((*QueryArray)(nil)).Elem())
}

type QueryArrayAssociatedWithLun QueryArrayAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryArrayAssociatedWithLun", reflect.TypeOf((*QueryArrayAssociatedWithLun)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryArrayAssociatedWithLun`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryArrayAssociatedWithLunRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `ScsiLun.canonicalName`
	// of ScsiLun
	CanonicalName string `xml:"canonicalName" json:"canonicalName"`
}

func init() {
	types.Add("sms:QueryArrayAssociatedWithLunRequestType", reflect.TypeOf((*QueryArrayAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryArrayAssociatedWithLunResponse struct {
	Returnval *StorageArray `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// The parameters of `SmsStorageManager.QueryArray`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryArrayRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// List of `SmsProviderInfo.uid` for the VASA
	// provider objects.
	ProviderId []string `xml:"providerId,omitempty" json:"providerId,omitempty"`
}

func init() {
	types.Add("sms:QueryArrayRequestType", reflect.TypeOf((*QueryArrayRequestType)(nil)).Elem())
}

type QueryArrayResponse struct {
	Returnval []StorageArray `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryAssociatedBackingStoragePool QueryAssociatedBackingStoragePoolRequestType

func init() {
	types.Add("sms:QueryAssociatedBackingStoragePool", reflect.TypeOf((*QueryAssociatedBackingStoragePool)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryAssociatedBackingStoragePool`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryAssociatedBackingStoragePoolRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Unique identifier of a StorageLun or StorageFileSystem.
	EntityId string `xml:"entityId,omitempty" json:"entityId,omitempty"`
	// Entity type of the entity specified using entityId. This can be either
	// StorageLun or StorageFileSystem.
	EntityType string `xml:"entityType,omitempty" json:"entityType,omitempty"`
}

func init() {
	types.Add("sms:QueryAssociatedBackingStoragePoolRequestType", reflect.TypeOf((*QueryAssociatedBackingStoragePoolRequestType)(nil)).Elem())
}

type QueryAssociatedBackingStoragePoolResponse struct {
	Returnval []BackingStoragePool `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryDatastoreBackingPoolMapping QueryDatastoreBackingPoolMappingRequestType

func init() {
	types.Add("sms:QueryDatastoreBackingPoolMapping", reflect.TypeOf((*QueryDatastoreBackingPoolMapping)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryDatastoreBackingPoolMapping`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryDatastoreBackingPoolMappingRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array containing references to `Datastore` objects.
	//
	// Refers instances of `Datastore`.
	Datastore []types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("sms:QueryDatastoreBackingPoolMappingRequestType", reflect.TypeOf((*QueryDatastoreBackingPoolMappingRequestType)(nil)).Elem())
}

type QueryDatastoreBackingPoolMappingResponse struct {
	Returnval []DatastoreBackingPoolMapping `xml:"returnval" json:"returnval"`
}

type QueryDatastoreCapability QueryDatastoreCapabilityRequestType

func init() {
	types.Add("sms:QueryDatastoreCapability", reflect.TypeOf((*QueryDatastoreCapability)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryDatastoreCapability`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryDatastoreCapabilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// reference to `Datastore`
	//
	// Refers instance of `Datastore`.
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("sms:QueryDatastoreCapabilityRequestType", reflect.TypeOf((*QueryDatastoreCapabilityRequestType)(nil)).Elem())
}

type QueryDatastoreCapabilityResponse struct {
	Returnval *StorageCapability `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryDrsMigrationCapabilityForPerformance QueryDrsMigrationCapabilityForPerformanceRequestType

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformance", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformance)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceEx QueryDrsMigrationCapabilityForPerformanceExRequestType

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceEx", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceEx)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryDrsMigrationCapabilityForPerformanceEx`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryDrsMigrationCapabilityForPerformanceExRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array containing references to `Datastore` objects.
	//
	// Refers instances of `Datastore`.
	Datastore []types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceExRequestType", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceExRequestType)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceExResponse struct {
	Returnval DrsMigrationCapabilityResult `xml:"returnval" json:"returnval"`
}

// The parameters of `SmsStorageManager.QueryDrsMigrationCapabilityForPerformance`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryDrsMigrationCapabilityForPerformanceRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Reference to the source `Datastore`
	//
	// Refers instance of `Datastore`.
	SrcDatastore types.ManagedObjectReference `xml:"srcDatastore" json:"srcDatastore"`
	// Reference to the destination `Datastore`
	//
	// Refers instance of `Datastore`.
	DstDatastore types.ManagedObjectReference `xml:"dstDatastore" json:"dstDatastore"`
}

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceRequestType", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceRequestType)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceResponse struct {
	Returnval bool `xml:"returnval" json:"returnval"`
}

// This exception is thrown if a failure occurs while
// processing a query request.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryExecutionFault struct {
	types.MethodFault
}

func init() {
	types.Add("sms:QueryExecutionFault", reflect.TypeOf((*QueryExecutionFault)(nil)).Elem())
}

type QueryExecutionFaultFault BaseQueryExecutionFault

func init() {
	types.Add("sms:QueryExecutionFaultFault", reflect.TypeOf((*QueryExecutionFaultFault)(nil)).Elem())
}

type QueryFaultDomain QueryFaultDomainRequestType

func init() {
	types.Add("sms:QueryFaultDomain", reflect.TypeOf((*QueryFaultDomain)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryFaultDomain`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryFaultDomainRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// spec for the query operation.
	Filter *FaultDomainFilter `xml:"filter,omitempty" json:"filter,omitempty"`
}

func init() {
	types.Add("sms:QueryFaultDomainRequestType", reflect.TypeOf((*QueryFaultDomainRequestType)(nil)).Elem())
}

type QueryFaultDomainResponse struct {
	Returnval []types.FaultDomainId `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryFileSystemAssociatedWithArray QueryFileSystemAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryFileSystemAssociatedWithArray", reflect.TypeOf((*QueryFileSystemAssociatedWithArray)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryFileSystemAssociatedWithArray`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryFileSystemAssociatedWithArrayRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryFileSystemAssociatedWithArrayRequestType", reflect.TypeOf((*QueryFileSystemAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryFileSystemAssociatedWithArrayResponse struct {
	Returnval []StorageFileSystem `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryHostAssociatedWithLun QueryHostAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryHostAssociatedWithLun", reflect.TypeOf((*QueryHostAssociatedWithLun)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryHostAssociatedWithLun`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryHostAssociatedWithLunRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageLun.uuid` for the StorageLun
	// object.
	Scsi3Id string `xml:"scsi3Id" json:"scsi3Id"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryHostAssociatedWithLunRequestType", reflect.TypeOf((*QueryHostAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryHostAssociatedWithLunResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryLunAssociatedWithArray QueryLunAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryLunAssociatedWithArray", reflect.TypeOf((*QueryLunAssociatedWithArray)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryLunAssociatedWithArray`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryLunAssociatedWithArrayRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryLunAssociatedWithArrayRequestType", reflect.TypeOf((*QueryLunAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryLunAssociatedWithArrayResponse struct {
	Returnval []StorageLun `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryLunAssociatedWithPort QueryLunAssociatedWithPortRequestType

func init() {
	types.Add("sms:QueryLunAssociatedWithPort", reflect.TypeOf((*QueryLunAssociatedWithPort)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryLunAssociatedWithPort`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryLunAssociatedWithPortRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StoragePort.uuid` for the StoragePort
	// object.
	PortId string `xml:"portId" json:"portId"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryLunAssociatedWithPortRequestType", reflect.TypeOf((*QueryLunAssociatedWithPortRequestType)(nil)).Elem())
}

type QueryLunAssociatedWithPortResponse struct {
	Returnval []StorageLun `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryNfsDatastoreAssociatedWithFileSystem QueryNfsDatastoreAssociatedWithFileSystemRequestType

func init() {
	types.Add("sms:QueryNfsDatastoreAssociatedWithFileSystem", reflect.TypeOf((*QueryNfsDatastoreAssociatedWithFileSystem)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryNfsDatastoreAssociatedWithFileSystem`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryNfsDatastoreAssociatedWithFileSystemRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageFileSystem.uuid` for the
	// StorageFileSystem object
	FileSystemId string `xml:"fileSystemId" json:"fileSystemId"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryNfsDatastoreAssociatedWithFileSystemRequestType", reflect.TypeOf((*QueryNfsDatastoreAssociatedWithFileSystemRequestType)(nil)).Elem())
}

type QueryNfsDatastoreAssociatedWithFileSystemResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// This exception is thrown if the specified entity and related
// entity type combination for a list query is not supported.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryNotSupported struct {
	types.InvalidArgument

	// Entity type.
	EntityType EntityReferenceEntityType `xml:"entityType,omitempty" json:"entityType,omitempty"`
	// Related entity type.
	RelatedEntityType EntityReferenceEntityType `xml:"relatedEntityType" json:"relatedEntityType"`
}

func init() {
	types.Add("sms:QueryNotSupported", reflect.TypeOf((*QueryNotSupported)(nil)).Elem())
}

type QueryNotSupportedFault QueryNotSupported

func init() {
	types.Add("sms:QueryNotSupportedFault", reflect.TypeOf((*QueryNotSupportedFault)(nil)).Elem())
}

type QueryPointInTimeReplica QueryPointInTimeReplicaRequestType

func init() {
	types.Add("sms:QueryPointInTimeReplica", reflect.TypeOf((*QueryPointInTimeReplica)(nil)).Elem())
}

// Describes the search criteria for the PiT query.
//
// If none of the fields
// is set, or if the number of PiT replicas is too large, VASA provider can
// return `QueryPointInTimeReplicaSummaryResult`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPointInTimeReplicaParam struct {
	types.DynamicData

	// Specifies the replica time span that vSphere is interested in.
	ReplicaTimeQueryParam *ReplicaQueryIntervalParam `xml:"replicaTimeQueryParam,omitempty" json:"replicaTimeQueryParam,omitempty"`
	// Only the replicas that match the given name are requested.
	//
	// A regexp according to http://www.w3.org/TR/xmlschema-2/#regexs.
	PitName string `xml:"pitName,omitempty" json:"pitName,omitempty"`
	// Only the replicas with tags that match the given tag(s) are requested.
	//
	// Each entry may be a regexp according to http://www.w3.org/TR/xmlschema-2/#regexs.
	Tags []string `xml:"tags,omitempty" json:"tags,omitempty"`
	// This field is hint for the preferred type of return results.
	//
	// It can be either true for `QueryPointInTimeReplicaSuccessResult` or
	// false for `QueryPointInTimeReplicaSummaryResult`.
	// If not set, VP may choose the appropriate type, as described in
	// <code>ReplicaQueryIntervalParam</code>.
	PreferDetails *bool `xml:"preferDetails" json:"preferDetails,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaParam", reflect.TypeOf((*QueryPointInTimeReplicaParam)(nil)).Elem())
}

// The parameters of `VasaProvider.QueryPointInTimeReplica`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPointInTimeReplicaRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// List of replication group IDs.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
	// Search criteria specification for all the groups.
	QueryParam *QueryPointInTimeReplicaParam `xml:"queryParam,omitempty" json:"queryParam,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaRequestType", reflect.TypeOf((*QueryPointInTimeReplicaRequestType)(nil)).Elem())
}

type QueryPointInTimeReplicaResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

// Return type for successful
// `VasaProvider.QueryPointInTimeReplica`
// operation.
//
// If the VASA provider decides that there are too many to return,
// it could set the result of some of the groups to `TooMany`
// fault or `QueryPointInTimeReplicaSummaryResult`.
//
// vSphere will then query for these groups separately.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPointInTimeReplicaSuccessResult struct {
	GroupOperationResult

	// Information about the available replicas.
	ReplicaInfo []PointInTimeReplicaInfo `xml:"replicaInfo,omitempty" json:"replicaInfo,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaSuccessResult", reflect.TypeOf((*QueryPointInTimeReplicaSuccessResult)(nil)).Elem())
}

// Summary of the available replicas.
//
// Mostly useful for CDP type replicators.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPointInTimeReplicaSummaryResult struct {
	GroupOperationResult

	// A series of query results.
	//
	// No special ordering is assumed by vSphere.
	IntervalResults []ReplicaIntervalQueryResult `xml:"intervalResults,omitempty" json:"intervalResults,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaSummaryResult", reflect.TypeOf((*QueryPointInTimeReplicaSummaryResult)(nil)).Elem())
}

type QueryPortAssociatedWithArray QueryPortAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithArray", reflect.TypeOf((*QueryPortAssociatedWithArray)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryPortAssociatedWithArray`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPortAssociatedWithArrayRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithArrayRequestType", reflect.TypeOf((*QueryPortAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithArrayResponse struct {
	Returnval []BaseStoragePort `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type QueryPortAssociatedWithLun QueryPortAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithLun", reflect.TypeOf((*QueryPortAssociatedWithLun)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryPortAssociatedWithLun`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPortAssociatedWithLunRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageLun.uuid` for the StorageLun
	// object.
	Scsi3Id string `xml:"scsi3Id" json:"scsi3Id"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithLunRequestType", reflect.TypeOf((*QueryPortAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithLunResponse struct {
	Returnval BaseStoragePort `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type QueryPortAssociatedWithProcessor QueryPortAssociatedWithProcessorRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithProcessor", reflect.TypeOf((*QueryPortAssociatedWithProcessor)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryPortAssociatedWithProcessor`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryPortAssociatedWithProcessorRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageProcessor.uuid` for the
	// StorageProcessor object.
	ProcessorId string `xml:"processorId" json:"processorId"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithProcessorRequestType", reflect.TypeOf((*QueryPortAssociatedWithProcessorRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithProcessorResponse struct {
	Returnval []BaseStoragePort `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type QueryProcessorAssociatedWithArray QueryProcessorAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryProcessorAssociatedWithArray", reflect.TypeOf((*QueryProcessorAssociatedWithArray)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryProcessorAssociatedWithArray`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryProcessorAssociatedWithArrayRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryProcessorAssociatedWithArrayRequestType", reflect.TypeOf((*QueryProcessorAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryProcessorAssociatedWithArrayResponse struct {
	Returnval []StorageProcessor `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryProvider QueryProviderRequestType

func init() {
	types.Add("sms:QueryProvider", reflect.TypeOf((*QueryProvider)(nil)).Elem())
}

type QueryProviderInfo QueryProviderInfoRequestType

func init() {
	types.Add("sms:QueryProviderInfo", reflect.TypeOf((*QueryProviderInfo)(nil)).Elem())
}

type QueryProviderInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QueryProviderInfoRequestType", reflect.TypeOf((*QueryProviderInfoRequestType)(nil)).Elem())
}

type QueryProviderInfoResponse struct {
	Returnval BaseSmsProviderInfo `xml:"returnval,typeattr" json:"returnval"`
}

type QueryProviderRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QueryProviderRequestType", reflect.TypeOf((*QueryProviderRequestType)(nil)).Elem())
}

type QueryProviderResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryReplicationGroup QueryReplicationGroupRequestType

func init() {
	types.Add("sms:QueryReplicationGroup", reflect.TypeOf((*QueryReplicationGroup)(nil)).Elem())
}

type QueryReplicationGroupInfo QueryReplicationGroupInfoRequestType

func init() {
	types.Add("sms:QueryReplicationGroupInfo", reflect.TypeOf((*QueryReplicationGroupInfo)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryReplicationGroupInfo`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryReplicationGroupInfoRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	RgFilter ReplicationGroupFilter       `xml:"rgFilter" json:"rgFilter"`
}

func init() {
	types.Add("sms:QueryReplicationGroupInfoRequestType", reflect.TypeOf((*QueryReplicationGroupInfoRequestType)(nil)).Elem())
}

type QueryReplicationGroupInfoResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

// The parameters of `VasaProvider.QueryReplicationGroup`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryReplicationGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// List of replication group IDs.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationGroupRequestType", reflect.TypeOf((*QueryReplicationGroupRequestType)(nil)).Elem())
}

type QueryReplicationGroupResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

// Information about the replication groups.
//
// Information about both the source
// groups and the target groups is returned.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryReplicationGroupSuccessResult struct {
	GroupOperationResult

	// Information about the replication group.
	//
	// May be either
	// `SourceGroupInfo` or `TargetGroupInfo`.
	RgInfo BaseGroupInfo `xml:"rgInfo,typeattr" json:"rgInfo"`
}

func init() {
	types.Add("sms:QueryReplicationGroupSuccessResult", reflect.TypeOf((*QueryReplicationGroupSuccessResult)(nil)).Elem())
}

type QueryReplicationPeer QueryReplicationPeerRequestType

func init() {
	types.Add("sms:QueryReplicationPeer", reflect.TypeOf((*QueryReplicationPeer)(nil)).Elem())
}

// The parameters of `VasaProvider.QueryReplicationPeer`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryReplicationPeerRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// An optional list of source fault domain ID.
	FaultDomainId []types.FaultDomainId `xml:"faultDomainId,omitempty" json:"faultDomainId,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationPeerRequestType", reflect.TypeOf((*QueryReplicationPeerRequestType)(nil)).Elem())
}

type QueryReplicationPeerResponse struct {
	Returnval []QueryReplicationPeerResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// Information about the replication peers of a VASA provider.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryReplicationPeerResult struct {
	types.DynamicData

	// Source fault domain id, must correspond to an id from the input.
	SourceDomain types.FaultDomainId `xml:"sourceDomain" json:"sourceDomain"`
	// Target fault domains for the given source, fault domain ID's are globally
	// unique.
	//
	// There can be one or more target domains for a given source.
	TargetDomain []types.FaultDomainId `xml:"targetDomain,omitempty" json:"targetDomain,omitempty"`
	// Error must be set when targetDomain field is not set.
	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
	// Optional warning messages.
	Warning []types.LocalizedMethodFault `xml:"warning,omitempty" json:"warning,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationPeerResult", reflect.TypeOf((*QueryReplicationPeerResult)(nil)).Elem())
}

type QuerySessionManager QuerySessionManagerRequestType

func init() {
	types.Add("sms:QuerySessionManager", reflect.TypeOf((*QuerySessionManager)(nil)).Elem())
}

type QuerySessionManagerRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QuerySessionManagerRequestType", reflect.TypeOf((*QuerySessionManagerRequestType)(nil)).Elem())
}

type QuerySessionManagerResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type QuerySmsTaskInfo QuerySmsTaskInfoRequestType

func init() {
	types.Add("sms:QuerySmsTaskInfo", reflect.TypeOf((*QuerySmsTaskInfo)(nil)).Elem())
}

type QuerySmsTaskInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QuerySmsTaskInfoRequestType", reflect.TypeOf((*QuerySmsTaskInfoRequestType)(nil)).Elem())
}

type QuerySmsTaskInfoResponse struct {
	Returnval SmsTaskInfo `xml:"returnval" json:"returnval"`
}

type QuerySmsTaskResult QuerySmsTaskResultRequestType

func init() {
	types.Add("sms:QuerySmsTaskResult", reflect.TypeOf((*QuerySmsTaskResult)(nil)).Elem())
}

type QuerySmsTaskResultRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QuerySmsTaskResultRequestType", reflect.TypeOf((*QuerySmsTaskResultRequestType)(nil)).Elem())
}

type QuerySmsTaskResultResponse struct {
	Returnval types.AnyType `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type QueryStorageContainer QueryStorageContainerRequestType

func init() {
	types.Add("sms:QueryStorageContainer", reflect.TypeOf((*QueryStorageContainer)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryStorageContainer`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryStorageContainerRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageContainerSpec`
	ContainerSpec *StorageContainerSpec `xml:"containerSpec,omitempty" json:"containerSpec,omitempty"`
}

func init() {
	types.Add("sms:QueryStorageContainerRequestType", reflect.TypeOf((*QueryStorageContainerRequestType)(nil)).Elem())
}

type QueryStorageContainerResponse struct {
	Returnval *StorageContainerResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type QueryStorageManager QueryStorageManagerRequestType

func init() {
	types.Add("sms:QueryStorageManager", reflect.TypeOf((*QueryStorageManager)(nil)).Elem())
}

type QueryStorageManagerRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:QueryStorageManagerRequestType", reflect.TypeOf((*QueryStorageManagerRequestType)(nil)).Elem())
}

type QueryStorageManagerResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type QueryVmfsDatastoreAssociatedWithLun QueryVmfsDatastoreAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryVmfsDatastoreAssociatedWithLun", reflect.TypeOf((*QueryVmfsDatastoreAssociatedWithLun)(nil)).Elem())
}

// The parameters of `SmsStorageManager.QueryVmfsDatastoreAssociatedWithLun`.
//
// This structure may be used only with operations rendered under `/sms`.
type QueryVmfsDatastoreAssociatedWithLunRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `StorageLun.uuid` for the StorageLun object
	Scsi3Id string `xml:"scsi3Id" json:"scsi3Id"`
	// `StorageArray.uuid` for the StorageArray
	// object.
	ArrayId string `xml:"arrayId" json:"arrayId"`
}

func init() {
	types.Add("sms:QueryVmfsDatastoreAssociatedWithLunRequestType", reflect.TypeOf((*QueryVmfsDatastoreAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryVmfsDatastoreAssociatedWithLunResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// Represents the device after the failover.
//
// Even though many of the fields in this structure are
// marked optional, it is important for VASA provider to
// make sure that the recovery of the entire ReplicationGroup succeeds
// atomically. The only valid scenario when there is a device specific
// recovery failure is when there is no valid replica for the Virtual Volume
// (e.g. Virtual Volume was just added to the ReplicationGroup).
//
// This structure may be used only with operations rendered under `/sms`.
type RecoveredDevice struct {
	types.DynamicData

	// Identifier of the device which was the target of replication before
	// failover.
	TargetDeviceId *ReplicaId `xml:"targetDeviceId,omitempty" json:"targetDeviceId,omitempty"`
	// Identifier of the target device after test or failover.
	RecoveredDeviceId BaseDeviceId `xml:"recoveredDeviceId,omitempty,typeattr" json:"recoveredDeviceId,omitempty"`
	// Identifier of the source of the replication data before the failover
	// stopped the replication.
	SourceDeviceId BaseDeviceId `xml:"sourceDeviceId,typeattr" json:"sourceDeviceId"`
	// Informational messages.
	Info []string `xml:"info,omitempty" json:"info,omitempty"`
	// Datastore for the newly surfaced device.
	//
	// Refers instance of `Datastore`.
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
	// Only to be filled in if the `RecoveredDevice.recoveredDeviceId` is `VirtualMachineId`.
	RecoveredDiskInfo []RecoveredDiskInfo `xml:"recoveredDiskInfo,omitempty" json:"recoveredDiskInfo,omitempty"`
	// Virtual Volume specific recovery error.
	//
	// This should be rare.
	Error *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
	// Warnings.
	Warnings []types.LocalizedMethodFault `xml:"warnings,omitempty" json:"warnings,omitempty"`
}

func init() {
	types.Add("sms:RecoveredDevice", reflect.TypeOf((*RecoveredDevice)(nil)).Elem())
}

// Describes the recovered disks for a given virtual machine.
//
// Only applicable for VAIO based replicators. Upon recovery,
// all the replicated disks must be attached to the virtual machine,
// i.e. the VMX file must refer to the correct file paths. Device
// keys must be preserved and non-replicated disks can refer to
// non-existent file names.
// Array based replicators can ignore this class.
//
// This structure may be used only with operations rendered under `/sms`.
type RecoveredDiskInfo struct {
	types.DynamicData

	// Virtual disk key.
	//
	// Note that disk device
	// keys must not change after recovery - in other words, the device
	// key is the same on both the source and target sites.
	//
	// For example, if a VMDK d1 is being replicated to d1', and d1 is attached as device
	// 2001 to the source VM, the recovered VM should have d1' attached
	// as 2001.
	DeviceKey int32 `xml:"deviceKey" json:"deviceKey"`
	// URL of the datastore that disk was recovered to.
	DsUrl string `xml:"dsUrl" json:"dsUrl"`
	// Full pathname of the disk.
	DiskPath string `xml:"diskPath" json:"diskPath"`
}

func init() {
	types.Add("sms:RecoveredDiskInfo", reflect.TypeOf((*RecoveredDiskInfo)(nil)).Elem())
}

// Information about member virtual volumes in a ReplicationGroup
// on the target after failover or testFailoverStart.
//
// This must include information about all the vSphere managed snapshots in
// the ReplicationGroup.
//
// This structure may be used only with operations rendered under `/sms`.
type RecoveredTargetGroupMemberInfo struct {
	TargetGroupMemberInfo

	// Identifier of the target device after test or failover.
	RecoveredDeviceId BaseDeviceId `xml:"recoveredDeviceId,omitempty,typeattr" json:"recoveredDeviceId,omitempty"`
}

func init() {
	types.Add("sms:RecoveredTargetGroupMemberInfo", reflect.TypeOf((*RecoveredTargetGroupMemberInfo)(nil)).Elem())
}

// The parameters of `SmsStorageManager.RegisterProvider_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type RegisterProviderRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `SmsProviderSpec`
	// containing parameters needed to register the
	// provider
	ProviderSpec BaseSmsProviderSpec `xml:"providerSpec,typeattr" json:"providerSpec"`
}

func init() {
	types.Add("sms:RegisterProviderRequestType", reflect.TypeOf((*RegisterProviderRequestType)(nil)).Elem())
}

type RegisterProvider_Task RegisterProviderRequestType

func init() {
	types.Add("sms:RegisterProvider_Task", reflect.TypeOf((*RegisterProvider_Task)(nil)).Elem())
}

type RegisterProvider_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Indicates whether the provider has been marked as active for the given array
// for the session context.
//
// SMS uses `StorageArray.priority` value to mark a provider
// as active among the ones that are registered with SMS and manage this array.
//
// This structure may be used only with operations rendered under `/sms`.
type RelatedStorageArray struct {
	types.DynamicData

	// `StorageArray.uuid` of StorageArray
	ArrayId string `xml:"arrayId" json:"arrayId"`
	// This field indicates whether the provider is currently serving data for the StorageArray
	Active bool `xml:"active" json:"active"`
	// Manageability status of StorageArray on VASA provider, if true it is manageable.
	Manageable bool `xml:"manageable" json:"manageable"`
	// Deprecated as of SMS API 6.0, replaced by `VasaProviderInfo.priority`.
	//
	// `StorageArray.priority` of StorageArray
	// For VASA 1.0 providers, this field is set to -1.
	Priority int32 `xml:"priority" json:"priority"`
}

func init() {
	types.Add("sms:RelatedStorageArray", reflect.TypeOf((*RelatedStorageArray)(nil)).Elem())
}

// Identifier of the replication target device.
//
// For Virtual Volumes, this could be the same as a Virtual Volume
// Id, for VMDK's this could be an FCD uuid, or some other ID
// made up by the replicator. This identifier is opaque to vSphere and
// hence does not have any distinguishing type. This can be used
// to identify the replica without the accompanying source device id
// (though there are no such uses in the current API).
//
// Since this an opaque type, the recovered device id at
// `RecoveredTargetGroupMemberInfo.recoveredDeviceId`
// should be filled in even if the values are the same.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicaId struct {
	types.DynamicData

	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("sms:ReplicaId", reflect.TypeOf((*ReplicaId)(nil)).Elem())
}

// Summarizes the collection of replicas in one time interval.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicaIntervalQueryResult struct {
	types.DynamicData

	// Beginning of interval (inclusive).
	FromDate time.Time `xml:"fromDate" json:"fromDate"`
	// End of interval (exclusive).
	ToDate time.Time `xml:"toDate" json:"toDate"`
	// Number of Point in Time replicas available for recovery.
	//
	// TODO: Do we want to have also ask for number of 'special'
	// PiTs e.g. those that are consistent?
	Number int32 `xml:"number" json:"number"`
}

func init() {
	types.Add("sms:ReplicaIntervalQueryResult", reflect.TypeOf((*ReplicaIntervalQueryResult)(nil)).Elem())
}

// Defines the parameters for a Point In Time replica (PiT) query.
//
// vSphere will not set all the three fields.
//
// In other words, the following combinations of fields are allowed:
//   - All the three fields are omitted.
//   - `ReplicaQueryIntervalParam.fromDate` and `ReplicaQueryIntervalParam.toDate` are set.
//   - `ReplicaQueryIntervalParam.fromDate` and `ReplicaQueryIntervalParam.number` are set.
//   - `ReplicaQueryIntervalParam.toDate` and `ReplicaQueryIntervalParam.number` are set.
//
// When all the fields are omitted, VASA provider should return
// `QueryPointInTimeReplicaSummaryResult`.
// But, returned result can be either `QueryPointInTimeReplicaSuccessResult`
// or `QueryPointInTimeReplicaSummaryResult` based on value i.e true or false
// respectively for field `QueryPointInTimeReplicaParam.preferDetails`.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicaQueryIntervalParam struct {
	types.DynamicData

	// Return all PiTs including and later than <code>fromDate</code>.
	FromDate *time.Time `xml:"fromDate" json:"fromDate,omitempty"`
	// Return all PiTs earlier than <code>toDate</code>.
	ToDate *time.Time `xml:"toDate" json:"toDate,omitempty"`
	// Return information for only <code>number</code> entries.
	Number int32 `xml:"number,omitempty" json:"number,omitempty"`
}

func init() {
	types.Add("sms:ReplicaQueryIntervalParam", reflect.TypeOf((*ReplicaQueryIntervalParam)(nil)).Elem())
}

// Describes one Replication Group's data.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicationGroupData struct {
	types.DynamicData

	// Replication group to failover.
	GroupId types.ReplicationGroupId `xml:"groupId" json:"groupId"`
	// The PIT that should be used for (test)Failover.
	//
	// Use the latest if not specified.
	PitId *PointInTimeReplicaId `xml:"pitId,omitempty" json:"pitId,omitempty"`
}

func init() {
	types.Add("sms:ReplicationGroupData", reflect.TypeOf((*ReplicationGroupData)(nil)).Elem())
}

// This spec contains information needed for `SmsStorageManager.QueryReplicationGroupInfo`
// API to filter the result.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicationGroupFilter struct {
	types.DynamicData

	// Query for the given replication groups from their associated providers.
	//
	// The groupId cannot be null or empty.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
}

func init() {
	types.Add("sms:ReplicationGroupFilter", reflect.TypeOf((*ReplicationGroupFilter)(nil)).Elem())
}

// Information about each replication target.
//
// This structure may be used only with operations rendered under `/sms`.
type ReplicationTargetInfo struct {
	types.DynamicData

	// Id of the target replication group (including the fault domain ID).
	TargetGroupId types.ReplicationGroupId `xml:"targetGroupId" json:"targetGroupId"`
	// Description of the replication agreement.
	//
	// This could be used to describe the characteristics of the replication
	// relationship between the source and the target (e.g. RPO, Replication
	// Mode, and other such properties). It is expected that VASA provider
	// will localize the string before sending to vSphere.
	ReplicationAgreementDescription string `xml:"replicationAgreementDescription,omitempty" json:"replicationAgreementDescription,omitempty"`
}

func init() {
	types.Add("sms:ReplicationTargetInfo", reflect.TypeOf((*ReplicationTargetInfo)(nil)).Elem())
}

// The parameters of `VasaProvider.ReverseReplicateGroup_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type ReverseReplicateGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of replication groups (currently in
	// `FAILEDOVER` state) that need to be reversed.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
}

func init() {
	types.Add("sms:ReverseReplicateGroupRequestType", reflect.TypeOf((*ReverseReplicateGroupRequestType)(nil)).Elem())
}

type ReverseReplicateGroup_Task ReverseReplicateGroupRequestType

func init() {
	types.Add("sms:ReverseReplicateGroup_Task", reflect.TypeOf((*ReverseReplicateGroup_Task)(nil)).Elem())
}

type ReverseReplicateGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Represents the result of a successful reverse replication action.
//
// The newly
// established replication relation might have a different source group ID and
// different set of target group IDs. This means that the replication topology
// will need to be discovered again by the DR orchestration programs (SRM/CAM).
// However, we assume that after the reverse replication, the new source fault
// domain id remains the same as the old (i.e. before failover) fault domain id.
//
// This structure may be used only with operations rendered under `/sms`.
type ReverseReplicationSuccessResult struct {
	GroupOperationResult

	// The replication group ID of the newly created source group.
	//
	// FaultDomainId
	// must remain the same.
	NewGroupId types.DeviceGroupId `xml:"newGroupId" json:"newGroupId"`
}

func init() {
	types.Add("sms:ReverseReplicationSuccessResult", reflect.TypeOf((*ReverseReplicationSuccessResult)(nil)).Elem())
}

// This exception is thrown if the storage management service
// has not yet been initialized successfully and therefore is
// not ready to process requests.
//
// This structure may be used only with operations rendered under `/sms`.
type ServiceNotInitialized struct {
	types.RuntimeFault
}

func init() {
	types.Add("sms:ServiceNotInitialized", reflect.TypeOf((*ServiceNotInitialized)(nil)).Elem())
}

type ServiceNotInitializedFault ServiceNotInitialized

func init() {
	types.Add("sms:ServiceNotInitializedFault", reflect.TypeOf((*ServiceNotInitializedFault)(nil)).Elem())
}

// This data object type describes system information.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsAboutInfo struct {
	types.DynamicData

	Name           string `xml:"name" json:"name"`
	FullName       string `xml:"fullName" json:"fullName"`
	Vendor         string `xml:"vendor" json:"vendor"`
	ApiVersion     string `xml:"apiVersion" json:"apiVersion"`
	InstanceUuid   string `xml:"instanceUuid" json:"instanceUuid"`
	VasaApiVersion string `xml:"vasaApiVersion,omitempty" json:"vasaApiVersion,omitempty"`
}

func init() {
	types.Add("sms:SmsAboutInfo", reflect.TypeOf((*SmsAboutInfo)(nil)).Elem())
}

// The common base type for all SMS faults.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsFault struct {
	types.MethodFault
}

func init() {
	types.Add("sms:SmsFault", reflect.TypeOf((*SmsFault)(nil)).Elem())
}

type SmsFaultFault SmsFault

func init() {
	types.Add("sms:SmsFaultFault", reflect.TypeOf((*SmsFaultFault)(nil)).Elem())
}

// Thrown when login fails due to token not provided or token could not be
// validated.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsInvalidLogin struct {
	types.MethodFault
}

func init() {
	types.Add("sms:SmsInvalidLogin", reflect.TypeOf((*SmsInvalidLogin)(nil)).Elem())
}

type SmsInvalidLoginFault SmsInvalidLogin

func init() {
	types.Add("sms:SmsInvalidLoginFault", reflect.TypeOf((*SmsInvalidLoginFault)(nil)).Elem())
}

// Information about Storage Monitoring Service (SMS)
// providers.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsProviderInfo struct {
	types.DynamicData

	// Unique identifier
	Uid string `xml:"uid" json:"uid"`
	// Name
	Name string `xml:"name" json:"name"`
	// Description of the provider
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// Version of the provider
	Version string `xml:"version,omitempty" json:"version,omitempty"`
}

func init() {
	types.Add("sms:SmsProviderInfo", reflect.TypeOf((*SmsProviderInfo)(nil)).Elem())
}

// Specification for Storage Monitoring Service (SMS)
// providers.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsProviderSpec struct {
	types.DynamicData

	// Name
	// The maximum length of the name is 275 characters.
	Name string `xml:"name" json:"name"`
	// Description of the provider
	Description string `xml:"description,omitempty" json:"description,omitempty"`
}

func init() {
	types.Add("sms:SmsProviderSpec", reflect.TypeOf((*SmsProviderSpec)(nil)).Elem())
}

// The parameters of `SmsStorageManager.SmsRefreshCACertificatesAndCRLs_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsRefreshCACertificatesAndCRLsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `SmsProviderInfo.uid` for providers
	ProviderId []string `xml:"providerId,omitempty" json:"providerId,omitempty"`
}

func init() {
	types.Add("sms:SmsRefreshCACertificatesAndCRLsRequestType", reflect.TypeOf((*SmsRefreshCACertificatesAndCRLsRequestType)(nil)).Elem())
}

type SmsRefreshCACertificatesAndCRLs_Task SmsRefreshCACertificatesAndCRLsRequestType

func init() {
	types.Add("sms:SmsRefreshCACertificatesAndCRLs_Task", reflect.TypeOf((*SmsRefreshCACertificatesAndCRLs_Task)(nil)).Elem())
}

type SmsRefreshCACertificatesAndCRLs_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Base class for all Replication faults.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsReplicationFault struct {
	types.MethodFault
}

func init() {
	types.Add("sms:SmsReplicationFault", reflect.TypeOf((*SmsReplicationFault)(nil)).Elem())
}

type SmsReplicationFaultFault BaseSmsReplicationFault

func init() {
	types.Add("sms:SmsReplicationFaultFault", reflect.TypeOf((*SmsReplicationFaultFault)(nil)).Elem())
}

// A ResourceInUse fault indicating that some error has occurred because
// some resources are in use.
//
// Information about the devices that are in
// use may be supplied.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsResourceInUse struct {
	types.ResourceInUse

	// The list of device Ids that are in use.
	DeviceIds []BaseDeviceId `xml:"deviceIds,omitempty,typeattr" json:"deviceIds,omitempty"`
}

func init() {
	types.Add("sms:SmsResourceInUse", reflect.TypeOf((*SmsResourceInUse)(nil)).Elem())
}

type SmsResourceInUseFault SmsResourceInUse

func init() {
	types.Add("sms:SmsResourceInUseFault", reflect.TypeOf((*SmsResourceInUseFault)(nil)).Elem())
}

// This data object type contains all information about a task.
//
// This structure may be used only with operations rendered under `/sms`.
type SmsTaskInfo struct {
	types.DynamicData

	// The unique key for the task.
	Key string `xml:"key" json:"key"`
	// The managed object that represents this task.
	//
	// Refers instance of `SmsTask`.
	Task types.ManagedObjectReference `xml:"task" json:"task"`
	// Managed Object to which the operation applies.
	Object *types.ManagedObjectReference `xml:"object,omitempty" json:"object,omitempty"`
	// If the task state is "error", then this property contains the fault code.
	Error *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
	// If the task state is "success", then this property may be used
	// to hold a return value.
	Result types.AnyType `xml:"result,omitempty,typeattr" json:"result,omitempty"`
	// Time stamp when the task started running.
	StartTime *time.Time `xml:"startTime" json:"startTime,omitempty"`
	// Time stamp when the task was completed (whether success or failure).
	CompletionTime *time.Time `xml:"completionTime" json:"completionTime,omitempty"`
	// Runtime status of the task.
	//
	// Possible values are `SmsTaskState_enum`
	State string `xml:"state" json:"state"`
	// If the task state is "running", then this property contains a
	// progress measurement, expressed as percentage completed, from 0 to 100.
	//
	// If this property is not set, then the command does not report progress.
	Progress int32 `xml:"progress,omitempty" json:"progress,omitempty"`
}

func init() {
	types.Add("sms:SmsTaskInfo", reflect.TypeOf((*SmsTaskInfo)(nil)).Elem())
}

// Replication group details on the source.
//
// We do not assume the same
// Replication Group id on all the sites. This is returned as answer to
// queryReplicationGroup.
//
// This structure may be used only with operations rendered under `/sms`.
type SourceGroupInfo struct {
	GroupInfo

	// Name of the replication group, may be edited after creating the
	// Replication Group, not unique.
	//
	// May be a localized string. Some vendors may
	// choose to use name as the group id, to support this, vSphere will not
	// allow the name to be modified - even if vSphere creates/manages the
	// Replication Group.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	// Description the Replication Group, may be edited after creating the
	// Replication Group.
	//
	// May be a localized string.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// State of the replication group on the source.
	State string `xml:"state" json:"state"`
	// Information about the target Replication Groups.
	Replica []ReplicationTargetInfo `xml:"replica,omitempty" json:"replica,omitempty"`
	// Information about the member virtual volumes and their replicas.
	MemberInfo []SourceGroupMemberInfo `xml:"memberInfo,omitempty" json:"memberInfo,omitempty"`
}

func init() {
	types.Add("sms:SourceGroupInfo", reflect.TypeOf((*SourceGroupInfo)(nil)).Elem())
}

// Represents a member virtual volume of a replication group on the source end
// of the replication arrow.
//
// This structure may be used only with operations rendered under `/sms`.
type SourceGroupMemberInfo struct {
	types.DynamicData

	// Identifier of the source device.
	//
	// May be a Virtual Volume, a Virtual Disk or a Virtual Machine
	DeviceId BaseDeviceId `xml:"deviceId,typeattr" json:"deviceId"`
	// Target devices, key'ed by the fault domain id.
	//
	// TODO: It is not clear if we
	// really need this information, since the target side query can return the
	// target -&gt; source relation information.
	TargetId []TargetDeviceId `xml:"targetId,omitempty" json:"targetId,omitempty"`
}

func init() {
	types.Add("sms:SourceGroupMemberInfo", reflect.TypeOf((*SourceGroupMemberInfo)(nil)).Elem())
}

// This data object represents the storage alarm.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageAlarm struct {
	types.DynamicData

	// Monotonically increasing sequence number which
	// VP will maintain.
	AlarmId int64 `xml:"alarmId" json:"alarmId"`
	// The type of Alarm.
	//
	// Must be one of the string values from
	// `AlarmType_enum`
	// Note that for VMODL VP implementation this field must be populated with one
	// of the values from `vasa.data.notification.AlarmType`
	AlarmType string `xml:"alarmType" json:"alarmType"`
	// Container identifier
	ContainerId string `xml:"containerId,omitempty" json:"containerId,omitempty"`
	// The unique identifier of the object impacted by the Alarm.
	//
	// From VASA version 3 onwards, a non-null `StorageAlarm.alarmObject`
	// will override this member.
	// This field is made optional from VASA3. Either this or
	// `StorageAlarm.alarmObject` must be set.
	ObjectId string `xml:"objectId,omitempty" json:"objectId,omitempty"`
	// The type of object impacted by the Alarm.
	//
	// Must be one of the string values
	// from `SmsEntityType_enum`
	// Note that for VMODL VP implementation this field must be populated with one
	// of the values from `vasa.data.notification.EntityType`
	ObjectType string `xml:"objectType" json:"objectType"`
	// Current status of the object.
	//
	// Must be one of the string values from
	// `SmsAlarmStatus_enum`
	Status string `xml:"status" json:"status"`
	// Time-stamp when the alarm occurred in VP context
	AlarmTimeStamp time.Time `xml:"alarmTimeStamp" json:"alarmTimeStamp"`
	// Pre-defined message for system-defined event
	MessageId string `xml:"messageId" json:"messageId"`
	// List of parameters (name/value) to be passed as input for message
	ParameterList []NameValuePair `xml:"parameterList,omitempty" json:"parameterList,omitempty"`
	// The ID of the object on which the alarm is raised; this is an object,
	// since ID's may not always be strings.
	//
	// vSphere will first use
	// `StorageAlarm.alarmObject` if set, and if not uses `StorageAlarm.objectId`.
	AlarmObject types.AnyType `xml:"alarmObject,omitempty,typeattr" json:"alarmObject,omitempty"`
}

func init() {
	types.Add("sms:StorageAlarm", reflect.TypeOf((*StorageAlarm)(nil)).Elem())
}

// This data object represents the storage array.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageArray struct {
	types.DynamicData

	// Name
	Name string `xml:"name" json:"name"`
	// Unique identifier
	Uuid string `xml:"uuid" json:"uuid"`
	// Storage array Vendor Id
	VendorId string `xml:"vendorId" json:"vendorId"`
	// Model Id
	ModelId string `xml:"modelId" json:"modelId"`
	// Storage array firmware
	Firmware string `xml:"firmware,omitempty" json:"firmware,omitempty"`
	// List of alternate storage array names
	AlternateName []string `xml:"alternateName,omitempty" json:"alternateName,omitempty"`
	// Supported block-device interfaces
	SupportedBlockInterface []string `xml:"supportedBlockInterface,omitempty" json:"supportedBlockInterface,omitempty"`
	// Supported file-system interfaces
	SupportedFileSystemInterface []string `xml:"supportedFileSystemInterface,omitempty" json:"supportedFileSystemInterface,omitempty"`
	// List of supported profiles
	SupportedProfile []string `xml:"supportedProfile,omitempty" json:"supportedProfile,omitempty"`
	// Deprecated as of SMS API 6.0, replaced by `VasaProviderInfo.priority`.
	//
	// Priority level of the provider for the given array within the session context.
	//
	// SMS will use this value to pick a provider among the ones that are registered
	// with SMS and manage this array. Once the provider is chosen, SMS will communicate
	// with it to get the data related to this array.
	// Valid range: 0 to 255.
	Priority int32 `xml:"priority,omitempty" json:"priority,omitempty"`
	// Required for NVMe-oF arrays and optional otherwise.
	//
	// Transport information to address the array's discovery service.
	DiscoverySvc []types.VASAStorageArrayDiscoverySvcInfo `xml:"discoverySvc,omitempty" json:"discoverySvc,omitempty"`
}

func init() {
	types.Add("sms:StorageArray", reflect.TypeOf((*StorageArray)(nil)).Elem())
}

// This data object represents the storage capability.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageCapability struct {
	types.DynamicData

	Uuid        string `xml:"uuid" json:"uuid"`
	Name        string `xml:"name" json:"name"`
	Description string `xml:"description" json:"description"`
}

func init() {
	types.Add("sms:StorageCapability", reflect.TypeOf((*StorageCapability)(nil)).Elem())
}

// This data object represents the storage container.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageContainer struct {
	types.DynamicData

	// Unique identifier
	Uuid string `xml:"uuid" json:"uuid"`
	// Name of the container
	Name string `xml:"name" json:"name"`
	// Maximum allowed capacity of the Virtual Volume in MBs
	MaxVvolSizeInMB int64 `xml:"maxVvolSizeInMB" json:"maxVvolSizeInMB"`
	// `SmsProviderInfo.uid` for providers that reports the storage container.
	ProviderId []string `xml:"providerId" json:"providerId"`
	ArrayId    []string `xml:"arrayId" json:"arrayId"`
	// Represents type of VVOL container, the supported values are listed in
	// `StorageContainerVvolContainerTypeEnum_enum`.
	//
	// If the storage array is not capable of supporting mixed PEs for a storage container,
	// the VVOL VASA provider sets this property to the supported endpoint type
	VvolContainerType string `xml:"vvolContainerType,omitempty" json:"vvolContainerType,omitempty"`
	// Indicates if this storage container is stretched
	Stretched *bool `xml:"stretched" json:"stretched,omitempty"`
}

func init() {
	types.Add("sms:StorageContainer", reflect.TypeOf((*StorageContainer)(nil)).Elem())
}

// This data object represents information about storage containers and related providers.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageContainerResult struct {
	types.DynamicData

	// `StorageContainer` objects
	StorageContainer []StorageContainer `xml:"storageContainer,omitempty" json:"storageContainer,omitempty"`
	// `SmsProviderInfo` corresponding to providers that
	// report these storage containers
	ProviderInfo []BaseSmsProviderInfo `xml:"providerInfo,omitempty,typeattr" json:"providerInfo,omitempty"`
}

func init() {
	types.Add("sms:StorageContainerResult", reflect.TypeOf((*StorageContainerResult)(nil)).Elem())
}

// This data object represents the specification to query
// storage containers retrieved from VASA providers.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageContainerSpec struct {
	types.DynamicData

	ContainerId []string `xml:"containerId,omitempty" json:"containerId,omitempty"`
}

func init() {
	types.Add("sms:StorageContainerSpec", reflect.TypeOf((*StorageContainerSpec)(nil)).Elem())
}

// This data object represents the storage file-system.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageFileSystem struct {
	types.DynamicData

	// Unique identifier
	Uuid string `xml:"uuid" json:"uuid"`
	// Information about the file system
	Info                    []StorageFileSystemInfo `xml:"info" json:"info"`
	NativeSnapshotSupported bool                    `xml:"nativeSnapshotSupported" json:"nativeSnapshotSupported"`
	ThinProvisioningStatus  string                  `xml:"thinProvisioningStatus" json:"thinProvisioningStatus"`
	Type                    string                  `xml:"type" json:"type"`
	Version                 string                  `xml:"version" json:"version"`
	// Backing config information
	BackingConfig *BackingConfig `xml:"backingConfig,omitempty" json:"backingConfig,omitempty"`
}

func init() {
	types.Add("sms:StorageFileSystem", reflect.TypeOf((*StorageFileSystem)(nil)).Elem())
}

// This data object represents information about the storage
// file-system.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageFileSystemInfo struct {
	types.DynamicData

	// Server Name
	FileServerName string `xml:"fileServerName" json:"fileServerName"`
	// File Path
	FileSystemPath string `xml:"fileSystemPath" json:"fileSystemPath"`
	// IP address
	IpAddress string `xml:"ipAddress,omitempty" json:"ipAddress,omitempty"`
}

func init() {
	types.Add("sms:StorageFileSystemInfo", reflect.TypeOf((*StorageFileSystemInfo)(nil)).Elem())
}

// This data object represents the storage lun.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageLun struct {
	types.DynamicData

	// Unique Indentfier
	Uuid string `xml:"uuid" json:"uuid"`
	// Identifier reported by vSphere(ESX) for this LUN
	VSphereLunIdentifier string `xml:"vSphereLunIdentifier" json:"vSphereLunIdentifier"`
	// Display Name which appears in storage array management
	// console
	VendorDisplayName string `xml:"vendorDisplayName" json:"vendorDisplayName"`
	// Capacity In MB
	CapacityInMB int64 `xml:"capacityInMB" json:"capacityInMB"`
	// Used space in MB for a thin provisioned LUN
	UsedSpaceInMB int64 `xml:"usedSpaceInMB" json:"usedSpaceInMB"`
	// Indicates whether the LUN is thin provisioned
	LunThinProvisioned bool `xml:"lunThinProvisioned" json:"lunThinProvisioned"`
	// Alternate identifiers associated with the LUN
	AlternateIdentifier []string `xml:"alternateIdentifier,omitempty" json:"alternateIdentifier,omitempty"`
	// Indicates whether Storage DRS is permitted to manage
	// performance between this LUN and other LUNs from the same
	// array.
	DrsManagementPermitted bool   `xml:"drsManagementPermitted" json:"drsManagementPermitted"`
	ThinProvisioningStatus string `xml:"thinProvisioningStatus" json:"thinProvisioningStatus"`
	// Backing config information
	BackingConfig *BackingConfig `xml:"backingConfig,omitempty" json:"backingConfig,omitempty"`
}

func init() {
	types.Add("sms:StorageLun", reflect.TypeOf((*StorageLun)(nil)).Elem())
}

// This data object represents the storage port.
//
// This structure may be used only with operations rendered under `/sms`.
type StoragePort struct {
	types.DynamicData

	// Unique identifier
	Uuid string `xml:"uuid" json:"uuid"`
	// Storage Port Type
	Type string `xml:"type" json:"type"`
	// Other identifiers which can help identify storage port
	AlternateName []string `xml:"alternateName,omitempty" json:"alternateName,omitempty"`
}

func init() {
	types.Add("sms:StoragePort", reflect.TypeOf((*StoragePort)(nil)).Elem())
}

// This data object represents the storage processor.
//
// This structure may be used only with operations rendered under `/sms`.
type StorageProcessor struct {
	types.DynamicData

	// Unique Identifier
	Uuid string `xml:"uuid" json:"uuid"`
	// List of alternate identifiers
	AlternateIdentifer []string `xml:"alternateIdentifer,omitempty" json:"alternateIdentifer,omitempty"`
}

func init() {
	types.Add("sms:StorageProcessor", reflect.TypeOf((*StorageProcessor)(nil)).Elem())
}

// Mapping between the supported vendorID and corresponding
// modelID
//
// This structure may be used only with operations rendered under `/sms`.
type SupportedVendorModelMapping struct {
	types.DynamicData

	// SCSI Vendor ID
	VendorId string `xml:"vendorId,omitempty" json:"vendorId,omitempty"`
	// SCSI Model ID
	ModelId string `xml:"modelId,omitempty" json:"modelId,omitempty"`
}

func init() {
	types.Add("sms:SupportedVendorModelMapping", reflect.TypeOf((*SupportedVendorModelMapping)(nil)).Elem())
}

// This exception is thrown if a sync operation is invoked
// while another sync invocation is in progress.
//
// This structure may be used only with operations rendered under `/sms`.
type SyncInProgress struct {
	ProviderSyncFailed
}

func init() {
	types.Add("sms:SyncInProgress", reflect.TypeOf((*SyncInProgress)(nil)).Elem())
}

type SyncInProgressFault SyncInProgress

func init() {
	types.Add("sms:SyncInProgressFault", reflect.TypeOf((*SyncInProgressFault)(nil)).Elem())
}

// Throw if an synchronization is ongoing.
//
// This structure may be used only with operations rendered under `/sms`.
type SyncOngoing struct {
	SmsReplicationFault

	// Task identifier of the ongoing sync (@see sms.TaskInfo#key).
	//
	// Refers instance of `SmsTask`.
	Task types.ManagedObjectReference `xml:"task" json:"task"`
}

func init() {
	types.Add("sms:SyncOngoing", reflect.TypeOf((*SyncOngoing)(nil)).Elem())
}

type SyncOngoingFault SyncOngoing

func init() {
	types.Add("sms:SyncOngoingFault", reflect.TypeOf((*SyncOngoingFault)(nil)).Elem())
}

// The parameters of `VasaProvider.SyncReplicationGroup_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type SyncReplicationGroupRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// List of replication group IDs.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
	// Localized name for the point-in-time snapshot created.
	PitName string `xml:"pitName" json:"pitName"`
}

func init() {
	types.Add("sms:SyncReplicationGroupRequestType", reflect.TypeOf((*SyncReplicationGroupRequestType)(nil)).Elem())
}

// Result object for a replication group that was successfully synchronized.
//
// This structure may be used only with operations rendered under `/sms`.
type SyncReplicationGroupSuccessResult struct {
	GroupOperationResult

	// Creation time of the PIT
	TimeStamp time.Time `xml:"timeStamp" json:"timeStamp"`
	// PIT id.
	//
	// If the VASA provider does not support PIT, this can be
	// left unset.
	//
	// A PIT created as a result of the <code>syncReplicationGroup</code>
	// may or may not have the same retention policy as other PITs. A VASA provider
	// can choose to delete such a PIT after a successful <code>testFailoverStop</code>
	PitId   *PointInTimeReplicaId `xml:"pitId,omitempty" json:"pitId,omitempty"`
	PitName string                `xml:"pitName,omitempty" json:"pitName,omitempty"`
}

func init() {
	types.Add("sms:SyncReplicationGroupSuccessResult", reflect.TypeOf((*SyncReplicationGroupSuccessResult)(nil)).Elem())
}

type SyncReplicationGroup_Task SyncReplicationGroupRequestType

func init() {
	types.Add("sms:SyncReplicationGroup_Task", reflect.TypeOf((*SyncReplicationGroup_Task)(nil)).Elem())
}

type SyncReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Represents a replication target device, since the replication group id can
// be the same in all the domains, this is keyed by the fault domain id.
//
// This structure may be used only with operations rendered under `/sms`.
type TargetDeviceId struct {
	types.DynamicData

	// ID of the fault domain.
	DomainId types.FaultDomainId `xml:"domainId" json:"domainId"`
	// ID of the target device.
	DeviceId ReplicaId `xml:"deviceId" json:"deviceId"`
}

func init() {
	types.Add("sms:TargetDeviceId", reflect.TypeOf((*TargetDeviceId)(nil)).Elem())
}

// Information about the replication target group.
//
// This is returned as answer
// to queryReplicationGroup before failover or testFailoverStart.
//
// This does not have to include the
// snapshot objects in the ReplicationGroup, however, see also
// `RecoveredTargetGroupMemberInfo`.
//
// This structure may be used only with operations rendered under `/sms`.
type TargetGroupInfo struct {
	GroupInfo

	// Replication source information.
	SourceInfo TargetToSourceInfo `xml:"sourceInfo" json:"sourceInfo"`
	// Replication state of the group on the replication target.
	State string `xml:"state" json:"state"`
	// Member device information.
	//
	// When the ReplicationGroup is either in `FAILEDOVER`
	// or `INTEST`, this
	// should be `RecoveredTargetGroupMemberInfo`.
	// Otherwise, this should be `TargetGroupMemberInfo`
	Devices []BaseTargetGroupMemberInfo `xml:"devices,omitempty,typeattr" json:"devices,omitempty"`
	// Whether the VASA provider is capable of executing
	// `VasaProvider.PromoteReplicationGroup_Task` for this
	// ReplicationGroup.
	//
	// False if not set. Note that this setting is per
	// ReplicationGroup per Target domain.
	IsPromoteCapable *bool `xml:"isPromoteCapable" json:"isPromoteCapable,omitempty"`
	// Name of Replication Group.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

func init() {
	types.Add("sms:TargetGroupInfo", reflect.TypeOf((*TargetGroupInfo)(nil)).Elem())
}

// Information about member virtual volumes in a ReplicationGroup
// on the target when the state is `TARGET`.
//
// This need not include information about all the snapshots in
// the ReplicationGroup.
//
// This structure may be used only with operations rendered under `/sms`.
type TargetGroupMemberInfo struct {
	types.DynamicData

	// Identifier of the replica device.
	ReplicaId ReplicaId `xml:"replicaId" json:"replicaId"`
	// Source device, since the device id can be the same in all the domains,
	// this needs to supplemented with the domain id to identify the device.
	SourceId BaseDeviceId `xml:"sourceId,typeattr" json:"sourceId"`
	// Datastore of the target device.
	//
	// This may be used by CAM/SRM
	// to notify the administrators to setup access paths for the hosts
	// to access the recovered devices.
	//
	// Refers instance of `Datastore`.
	TargetDatastore types.ManagedObjectReference `xml:"targetDatastore" json:"targetDatastore"`
}

func init() {
	types.Add("sms:TargetGroupMemberInfo", reflect.TypeOf((*TargetGroupMemberInfo)(nil)).Elem())
}

// Information about each replication target.
//
// This structure may be used only with operations rendered under `/sms`.
type TargetToSourceInfo struct {
	types.DynamicData

	// Id of the source CG id (including the fault domain ID).
	SourceGroupId types.ReplicationGroupId `xml:"sourceGroupId" json:"sourceGroupId"`
	// Description of the replication agreement.
	//
	// This could be used to describe the characteristics of the replication
	// relationship between the source and the target (e.g. RPO, Replication
	// Mode, and other such properties). It is expected that VASA provider
	// will localize the string before sending to vSphere.
	ReplicationAgreementDescription string `xml:"replicationAgreementDescription,omitempty" json:"replicationAgreementDescription,omitempty"`
}

func init() {
	types.Add("sms:TargetToSourceInfo", reflect.TypeOf((*TargetToSourceInfo)(nil)).Elem())
}

// Input to testFailover method.
//
// This structure may be used only with operations rendered under `/sms`.
type TestFailoverParam struct {
	FailoverParam
}

func init() {
	types.Add("sms:TestFailoverParam", reflect.TypeOf((*TestFailoverParam)(nil)).Elem())
}

// The parameters of `VasaProvider.TestFailoverReplicationGroupStart_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type TestFailoverReplicationGroupStartRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Settings for the failover.
	TestFailoverParam TestFailoverParam `xml:"testFailoverParam" json:"testFailoverParam"`
}

func init() {
	types.Add("sms:TestFailoverReplicationGroupStartRequestType", reflect.TypeOf((*TestFailoverReplicationGroupStartRequestType)(nil)).Elem())
}

type TestFailoverReplicationGroupStart_Task TestFailoverReplicationGroupStartRequestType

func init() {
	types.Add("sms:TestFailoverReplicationGroupStart_Task", reflect.TypeOf((*TestFailoverReplicationGroupStart_Task)(nil)).Elem())
}

type TestFailoverReplicationGroupStart_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// The parameters of `VasaProvider.TestFailoverReplicationGroupStop_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type TestFailoverReplicationGroupStopRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of replication groups that need to stop test.
	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty" json:"groupId,omitempty"`
	// \- if true, VP should force-unbind all Virtual Volumes
	// and move the RG from INTEST to TARGET state. If false, VP will report all the
	// Virtual Volumes which need to be cleaned up before a failover operation
	// can be triggered. The default value will be false.
	Force bool `xml:"force" json:"force"`
}

func init() {
	types.Add("sms:TestFailoverReplicationGroupStopRequestType", reflect.TypeOf((*TestFailoverReplicationGroupStopRequestType)(nil)).Elem())
}

type TestFailoverReplicationGroupStop_Task TestFailoverReplicationGroupStopRequestType

func init() {
	types.Add("sms:TestFailoverReplicationGroupStop_Task", reflect.TypeOf((*TestFailoverReplicationGroupStop_Task)(nil)).Elem())
}

type TestFailoverReplicationGroupStop_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// This exception is thrown if the request exceeds the maximum number of
// elements in batch that the VASA Provider can support for the specific API.
//
// This structure may be used only with operations rendered under `/sms`.
type TooMany struct {
	types.MethodFault

	// Maximum number of elements in batch that the VASA Provider can support
	// for the specific API.
	//
	// If the value is not specified (zero) or invalid
	// (negative), client will assume the default value is 1.
	MaxBatchSize int64 `xml:"maxBatchSize,omitempty" json:"maxBatchSize,omitempty"`
}

func init() {
	types.Add("sms:TooMany", reflect.TypeOf((*TooMany)(nil)).Elem())
}

type TooManyFault TooMany

func init() {
	types.Add("sms:TooManyFault", reflect.TypeOf((*TooManyFault)(nil)).Elem())
}

// The parameters of `SmsStorageManager.UnregisterProvider_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type UnregisterProviderRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `SmsProviderInfo.uid` for
	// the provider
	ProviderId string `xml:"providerId" json:"providerId"`
}

func init() {
	types.Add("sms:UnregisterProviderRequestType", reflect.TypeOf((*UnregisterProviderRequestType)(nil)).Elem())
}

type UnregisterProvider_Task UnregisterProviderRequestType

func init() {
	types.Add("sms:UnregisterProvider_Task", reflect.TypeOf((*UnregisterProvider_Task)(nil)).Elem())
}

type UnregisterProvider_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// The parameters of `SmsStorageManager.UpgradeVASAProvider_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type UpgradeVASAProviderRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `VASAProviderUpgradeSpec` containing parameter to upgrade the
	// VASA Provider. If spec is for non VVOL VASA Provider, then exception is thrown.
	UpgradeSpec VASAProviderUpgradeSpec `xml:"upgradeSpec" json:"upgradeSpec"`
}

func init() {
	types.Add("sms:UpgradeVASAProviderRequestType", reflect.TypeOf((*UpgradeVASAProviderRequestType)(nil)).Elem())
}

type UpgradeVASAProvider_Task UpgradeVASAProviderRequestType

func init() {
	types.Add("sms:UpgradeVASAProvider_Task", reflect.TypeOf((*UpgradeVASAProvider_Task)(nil)).Elem())
}

type UpgradeVASAProvider_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// VASA Provider Specification for upgrade.
//
// This structure may be used only with operations rendered under `/sms`.
type VASAProviderUpgradeSpec struct {
	types.DynamicData

	// `SmsProviderInfo.uid` for the provider
	ProviderUid string `xml:"providerUid" json:"providerUid"`
	// VASA Provider username
	Username string `xml:"username" json:"username"`
	// VASA Provider password
	Password string `xml:"password" json:"password"`
}

func init() {
	types.Add("sms:VASAProviderUpgradeSpec", reflect.TypeOf((*VASAProviderUpgradeSpec)(nil)).Elem())
}

// Identity of a virtual volume for policy API purposes.
//
// For the sake of
// brevity, let us use VVolId. This works because the class is defined as a part
// of the policy package.
//
// WSDL names do not have this feature, but WSDL names are usually prefixed with
// the package name any way.
//
// This structure may be used only with operations rendered under `/sms`.
type VVolId struct {
	DeviceId

	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("sms:VVolId", reflect.TypeOf((*VVolId)(nil)).Elem())
}

// Information about VASA (vStorage APIs for Storage Awareness) providers.
//
// This structure may be used only with operations rendered under `/sms`.
type VasaProviderInfo struct {
	SmsProviderInfo

	// Provider URL
	Url string `xml:"url" json:"url"`
	// Provider certificate
	Certificate string `xml:"certificate,omitempty" json:"certificate,omitempty"`
	// The operational state of VASA Provider.
	Status string `xml:"status,omitempty" json:"status,omitempty"`
	// A fault that describes the cause of the current operational status.
	StatusFault *types.LocalizedMethodFault `xml:"statusFault,omitempty" json:"statusFault,omitempty"`
	// Supported VASA(vStorage APIs for Storage Awareness) version
	VasaVersion string `xml:"vasaVersion,omitempty" json:"vasaVersion,omitempty"`
	// Namespace to categorize storage capabilities provided by
	// arrays managed by the provider
	Namespace string `xml:"namespace,omitempty" json:"namespace,omitempty"`
	// Time stamp to indicate when last sync operation was completed
	// successfully.
	LastSyncTime string `xml:"lastSyncTime,omitempty" json:"lastSyncTime,omitempty"`
	// List containing mapping between the supported vendorID and
	// corresponding modelID
	SupportedVendorModelMapping []SupportedVendorModelMapping `xml:"supportedVendorModelMapping,omitempty" json:"supportedVendorModelMapping,omitempty"`
	// Deprecated as of SMS API 3.0, use `StorageArray.supportedProfile`.
	//
	// List of supported profiles
	SupportedProfile []string `xml:"supportedProfile,omitempty" json:"supportedProfile,omitempty"`
	// List of supported profiles at provider level.
	//
	// Must be one of the string
	// values from `ProviderProfile_enum`.
	SupportedProviderProfile []string `xml:"supportedProviderProfile,omitempty" json:"supportedProviderProfile,omitempty"`
	// List containing mapping between storage arrays reported by the provider
	// and information such as whether the provider is considered active for them.
	RelatedStorageArray []RelatedStorageArray `xml:"relatedStorageArray,omitempty" json:"relatedStorageArray,omitempty"`
	// Provider identifier reported by the provider which is unique within
	// the provider namespace.
	ProviderId string `xml:"providerId,omitempty" json:"providerId,omitempty"`
	// Provider certificate expiry date.
	CertificateExpiryDate string `xml:"certificateExpiryDate,omitempty" json:"certificateExpiryDate,omitempty"`
	// Provider certificate status
	// This field holds values from `VasaProviderCertificateStatus_enum`
	CertificateStatus string `xml:"certificateStatus,omitempty" json:"certificateStatus,omitempty"`
	// Service location for the VASA endpoint that SMS is using
	// to communicate with the provider.
	ServiceLocation string `xml:"serviceLocation,omitempty" json:"serviceLocation,omitempty"`
	// Indicates the type of deployment supported by the provider.
	//
	// If true, it is an active/passive deployment and the provider needs to be
	// activated explicitly using activateProviderEx() VASA API.
	// If false, it is an active/active deployment and provider does not need any
	// explicit activation to respond to VASA calls.
	NeedsExplicitActivation *bool `xml:"needsExplicitActivation" json:"needsExplicitActivation,omitempty"`
	// Maximum number of elements in batch APIs that the VASA Provider can support.
	//
	// This value is common to all batch APIs supported by the provider. However,
	// for each specific API, the provider may still throw or return `TooMany`
	// fault in which a different value of maxBatchSize can be specified.
	// If the value is not specified (zero) or invalid (negative), client will
	// assume there's no common limit for the number of elements that can be
	// handled in all batch APIs.
	MaxBatchSize int64 `xml:"maxBatchSize,omitempty" json:"maxBatchSize,omitempty"`
	// Indicate whether the provider wants to retain its certificate after bootstrapping.
	//
	// If true, SMS will not provision a VMCA signed certificate for the provider
	// and all certificate life cycle management workflows are disabled for this provider certificate.
	// If false, SMS will provision a VMCA signed certificate for the provider and
	// all certificate life cycle management workflows are enabled for this provider certificate.
	RetainVasaProviderCertificate *bool `xml:"retainVasaProviderCertificate" json:"retainVasaProviderCertificate,omitempty"`
	// Indicates if this provider is independent of arrays.
	//
	// Default value for this flag is false, which means this provider supports
	// arrays. Arrays will be queried for this provider during sync. If this flag
	// is set to true, arrays will not be synced for this provider and array
	// related API will not be invoked on this provider.
	ArrayIndependentProvider *bool `xml:"arrayIndependentProvider" json:"arrayIndependentProvider,omitempty"`
	// Type of this VASA provider.
	//
	// This field will be equal to one of the values of `VpType_enum`.
	Type string `xml:"type,omitempty" json:"type,omitempty"`
	// This field indicates the category of the provider and will be equal to one of the values of
	// `VpCategory_enum`.
	Category string `xml:"category,omitempty" json:"category,omitempty"`
	// Priority level of the provider within a VASA HA group.
	//
	// For a stand-alone
	// provider which does not participate in VASA HA, this field will be ignored.
	//
	// The priority value is an integer with valid range from 0 to 255.
	Priority int32 `xml:"priority,omitempty" json:"priority,omitempty"`
	// Unique identifier of a VASA HA group.
	//
	// Providers should report this
	// identifier to utilize HA feature supported by vSphere. Different providers
	// reporting the same <code>failoverGroupId</code> will be treated as an HA
	// group. Failover/failback will be done within one group.
	FailoverGroupId string `xml:"failoverGroupId,omitempty" json:"failoverGroupId,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderInfo", reflect.TypeOf((*VasaProviderInfo)(nil)).Elem())
}

type VasaProviderReconnectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:VasaProviderReconnectRequestType", reflect.TypeOf((*VasaProviderReconnectRequestType)(nil)).Elem())
}

type VasaProviderReconnect_Task VasaProviderReconnectRequestType

func init() {
	types.Add("sms:VasaProviderReconnect_Task", reflect.TypeOf((*VasaProviderReconnect_Task)(nil)).Elem())
}

type VasaProviderReconnect_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VasaProviderRefreshCertificateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:VasaProviderRefreshCertificateRequestType", reflect.TypeOf((*VasaProviderRefreshCertificateRequestType)(nil)).Elem())
}

type VasaProviderRefreshCertificate_Task VasaProviderRefreshCertificateRequestType

func init() {
	types.Add("sms:VasaProviderRefreshCertificate_Task", reflect.TypeOf((*VasaProviderRefreshCertificate_Task)(nil)).Elem())
}

type VasaProviderRefreshCertificate_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VasaProviderRevokeCertificateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("sms:VasaProviderRevokeCertificateRequestType", reflect.TypeOf((*VasaProviderRevokeCertificateRequestType)(nil)).Elem())
}

type VasaProviderRevokeCertificate_Task VasaProviderRevokeCertificateRequestType

func init() {
	types.Add("sms:VasaProviderRevokeCertificate_Task", reflect.TypeOf((*VasaProviderRevokeCertificate_Task)(nil)).Elem())
}

type VasaProviderRevokeCertificate_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// VASA(vStorage APIs for Storage Awareness) provider
// specification
//
// This structure may be used only with operations rendered under `/sms`.
type VasaProviderSpec struct {
	SmsProviderSpec

	// Username
	// The maximum length of the username is 255 characters.
	Username string `xml:"username" json:"username"`
	// Password
	// The maximum length of the password is 255 characters.
	Password string `xml:"password" json:"password"`
	// URL
	Url string `xml:"url" json:"url"`
	// Certificate
	Certificate string `xml:"certificate,omitempty" json:"certificate,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderSpec", reflect.TypeOf((*VasaProviderSpec)(nil)).Elem())
}

// The parameters of `VasaProvider.VasaProviderSync_Task`.
//
// This structure may be used only with operations rendered under `/sms`.
type VasaProviderSyncRequestType struct {
	This    types.ManagedObjectReference `xml:"_this" json:"_this"`
	ArrayId string                       `xml:"arrayId,omitempty" json:"arrayId,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderSyncRequestType", reflect.TypeOf((*VasaProviderSyncRequestType)(nil)).Elem())
}

type VasaProviderSync_Task VasaProviderSyncRequestType

func init() {
	types.Add("sms:VasaProviderSync_Task", reflect.TypeOf((*VasaProviderSync_Task)(nil)).Elem())
}

type VasaProviderSync_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

// Represents a virtual disk with a UUID (aka FCD).
//
// Virtual Volume VASA providers can ignore this class.
//
// This structure may be used only with operations rendered under `/sms`.
type VasaVirtualDiskId struct {
	DeviceId

	// See VIM documentation for more details on first class storage - which is
	// new in 2016.
	DiskId string `xml:"diskId" json:"diskId"`
}

func init() {
	types.Add("sms:VasaVirtualDiskId", reflect.TypeOf((*VasaVirtualDiskId)(nil)).Elem())
}

// Represents a virtual disk.
//
// Ideally a UUID, since we do not yet have an FCD,
// let us use VM's UUID + diskKey.
// Virtual Volume VASA providers can ignore this class.
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualDiskKey struct {
	DeviceId

	// The vmInstanceUUID is unique to a VM.
	//
	// See
	// http://pubs.vmware.com/vsphere-60/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.vm.ConfigInfo.html
	VmInstanceUUID string `xml:"vmInstanceUUID" json:"vmInstanceUUID"`
	DeviceKey      int32  `xml:"deviceKey" json:"deviceKey"`
}

func init() {
	types.Add("sms:VirtualDiskKey", reflect.TypeOf((*VirtualDiskKey)(nil)).Elem())
}

// Identifies a VirtualDisk uniquely using the disk key, vCenter managed object id of the VM it
// belongs to and the corresponding vCenter UUID.
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualDiskMoId struct {
	DeviceId

	// The vCenter UUID - this is informational only, and may not always be set.
	VcUuid string `xml:"vcUuid,omitempty" json:"vcUuid,omitempty"`
	// The managed object id that corresponds to vCenter's ManagedObjectReference#key for this VM.
	VmMoid string `xml:"vmMoid" json:"vmMoid"`
	// The disk key that corresponds to the VirtualDevice#key in vCenter.
	DiskKey string `xml:"diskKey" json:"diskKey"`
}

func init() {
	types.Add("sms:VirtualDiskMoId", reflect.TypeOf((*VirtualDiskMoId)(nil)).Elem())
}

// Identifies a virtual machine by its VMX file path.
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualMachineFilePath struct {
	VirtualMachineId

	// The vCenter UUID - this is informational only,
	// and may not always be set.
	VcUuid string `xml:"vcUuid,omitempty" json:"vcUuid,omitempty"`
	// Datastore URL, which is globally unique (name is not).
	DsUrl string `xml:"dsUrl" json:"dsUrl"`
	// Full path name from the URL onwards.
	//
	// When the vmxPath is returned after failover, the VMX file
	// should be fixed up to contain correct target filenames for all replicated
	// disks. For non-replicated disks, the target filenames can contain
	// any arbitrary path. For better security, it is recommended to
	// set these disks pointed to a random string (e.g. UUID).
	VmxPath string `xml:"vmxPath" json:"vmxPath"`
}

func init() {
	types.Add("sms:VirtualMachineFilePath", reflect.TypeOf((*VirtualMachineFilePath)(nil)).Elem())
}

// Abstracts the identity of a virtual machine.
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualMachineId struct {
	DeviceId
}

func init() {
	types.Add("sms:VirtualMachineId", reflect.TypeOf((*VirtualMachineId)(nil)).Elem())
}

// Identifies a VirtualMachine uniquely using its vCenter managed object id and the corresponding
// vCenter UUID
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualMachineMoId struct {
	VirtualMachineId

	// The vCenter UUID - this is informational only, and may not always be set.
	VcUuid string `xml:"vcUuid,omitempty" json:"vcUuid,omitempty"`
	// The managed object id that corresponds to vCenter's ManagedObjectReference#key for this VM.
	VmMoid string `xml:"vmMoid" json:"vmMoid"`
}

func init() {
	types.Add("sms:VirtualMachineMoId", reflect.TypeOf((*VirtualMachineMoId)(nil)).Elem())
}

// Identifies a virtual machine by its vmInstanceUUID
//
// This structure may be used only with operations rendered under `/sms`.
type VirtualMachineUUID struct {
	VirtualMachineId

	// The vmInstanceUUID is unique to a VM.
	//
	// See
	// http://pubs.vmware.com/vsphere-60/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.vm.ConfigInfo.html
	VmInstanceUUID string `xml:"vmInstanceUUID" json:"vmInstanceUUID"`
}

func init() {
	types.Add("sms:VirtualMachineUUID", reflect.TypeOf((*VirtualMachineUUID)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("sms:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
