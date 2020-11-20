/*
Copyright (c) 2014-2018 VMware, Inc. All Rights Reserved.

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
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type AlarmFilter struct {
	types.DynamicData

	AlarmStatus string          `xml:"alarmStatus,omitempty"`
	AlarmType   string          `xml:"alarmType,omitempty"`
	EntityType  string          `xml:"entityType,omitempty"`
	EntityId    []types.AnyType `xml:"entityId,omitempty,typeattr"`
	PageMarker  string          `xml:"pageMarker,omitempty"`
}

func init() {
	types.Add("sms:AlarmFilter", reflect.TypeOf((*AlarmFilter)(nil)).Elem())
}

type AlarmResult struct {
	types.DynamicData

	StorageAlarm []StorageAlarm `xml:"storageAlarm,omitempty"`
	PageMarker   string         `xml:"pageMarker,omitempty"`
}

func init() {
	types.Add("sms:AlarmResult", reflect.TypeOf((*AlarmResult)(nil)).Elem())
}

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

type ArrayOfBackingStoragePool struct {
	BackingStoragePool []BackingStoragePool `xml:"BackingStoragePool,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfBackingStoragePool", reflect.TypeOf((*ArrayOfBackingStoragePool)(nil)).Elem())
}

type ArrayOfDatastoreBackingPoolMapping struct {
	DatastoreBackingPoolMapping []DatastoreBackingPoolMapping `xml:"DatastoreBackingPoolMapping,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfDatastoreBackingPoolMapping", reflect.TypeOf((*ArrayOfDatastoreBackingPoolMapping)(nil)).Elem())
}

type ArrayOfDatastorePair struct {
	DatastorePair []DatastorePair `xml:"DatastorePair,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfDatastorePair", reflect.TypeOf((*ArrayOfDatastorePair)(nil)).Elem())
}

type ArrayOfDeviceId struct {
	DeviceId []BaseDeviceId `xml:"DeviceId,omitempty,typeattr"`
}

func init() {
	types.Add("sms:ArrayOfDeviceId", reflect.TypeOf((*ArrayOfDeviceId)(nil)).Elem())
}

type ArrayOfFaultDomainProviderMapping struct {
	FaultDomainProviderMapping []FaultDomainProviderMapping `xml:"FaultDomainProviderMapping,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfFaultDomainProviderMapping", reflect.TypeOf((*ArrayOfFaultDomainProviderMapping)(nil)).Elem())
}

type ArrayOfGroupOperationResult struct {
	GroupOperationResult []BaseGroupOperationResult `xml:"GroupOperationResult,omitempty,typeattr"`
}

func init() {
	types.Add("sms:ArrayOfGroupOperationResult", reflect.TypeOf((*ArrayOfGroupOperationResult)(nil)).Elem())
}

type ArrayOfNameValuePair struct {
	NameValuePair []NameValuePair `xml:"NameValuePair,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfNameValuePair", reflect.TypeOf((*ArrayOfNameValuePair)(nil)).Elem())
}

type ArrayOfPointInTimeReplicaInfo struct {
	PointInTimeReplicaInfo []PointInTimeReplicaInfo `xml:"PointInTimeReplicaInfo,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfPointInTimeReplicaInfo", reflect.TypeOf((*ArrayOfPointInTimeReplicaInfo)(nil)).Elem())
}

type ArrayOfPolicyAssociation struct {
	PolicyAssociation []PolicyAssociation `xml:"PolicyAssociation,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfPolicyAssociation", reflect.TypeOf((*ArrayOfPolicyAssociation)(nil)).Elem())
}

type ArrayOfQueryReplicationPeerResult struct {
	QueryReplicationPeerResult []QueryReplicationPeerResult `xml:"QueryReplicationPeerResult,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfQueryReplicationPeerResult", reflect.TypeOf((*ArrayOfQueryReplicationPeerResult)(nil)).Elem())
}

type ArrayOfRecoveredDevice struct {
	RecoveredDevice []RecoveredDevice `xml:"RecoveredDevice,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfRecoveredDevice", reflect.TypeOf((*ArrayOfRecoveredDevice)(nil)).Elem())
}

type ArrayOfRecoveredDiskInfo struct {
	RecoveredDiskInfo []RecoveredDiskInfo `xml:"RecoveredDiskInfo,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfRecoveredDiskInfo", reflect.TypeOf((*ArrayOfRecoveredDiskInfo)(nil)).Elem())
}

type ArrayOfRelatedStorageArray struct {
	RelatedStorageArray []RelatedStorageArray `xml:"RelatedStorageArray,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfRelatedStorageArray", reflect.TypeOf((*ArrayOfRelatedStorageArray)(nil)).Elem())
}

type ArrayOfReplicaIntervalQueryResult struct {
	ReplicaIntervalQueryResult []ReplicaIntervalQueryResult `xml:"ReplicaIntervalQueryResult,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfReplicaIntervalQueryResult", reflect.TypeOf((*ArrayOfReplicaIntervalQueryResult)(nil)).Elem())
}

type ArrayOfReplicationGroupData struct {
	ReplicationGroupData []ReplicationGroupData `xml:"ReplicationGroupData,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfReplicationGroupData", reflect.TypeOf((*ArrayOfReplicationGroupData)(nil)).Elem())
}

type ArrayOfReplicationTargetInfo struct {
	ReplicationTargetInfo []ReplicationTargetInfo `xml:"ReplicationTargetInfo,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfReplicationTargetInfo", reflect.TypeOf((*ArrayOfReplicationTargetInfo)(nil)).Elem())
}

type ArrayOfSmsProviderInfo struct {
	SmsProviderInfo []BaseSmsProviderInfo `xml:"SmsProviderInfo,omitempty,typeattr"`
}

func init() {
	types.Add("sms:ArrayOfSmsProviderInfo", reflect.TypeOf((*ArrayOfSmsProviderInfo)(nil)).Elem())
}

type ArrayOfSourceGroupMemberInfo struct {
	SourceGroupMemberInfo []SourceGroupMemberInfo `xml:"SourceGroupMemberInfo,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfSourceGroupMemberInfo", reflect.TypeOf((*ArrayOfSourceGroupMemberInfo)(nil)).Elem())
}

type ArrayOfStorageAlarm struct {
	StorageAlarm []StorageAlarm `xml:"StorageAlarm,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageAlarm", reflect.TypeOf((*ArrayOfStorageAlarm)(nil)).Elem())
}

type ArrayOfStorageArray struct {
	StorageArray []StorageArray `xml:"StorageArray,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageArray", reflect.TypeOf((*ArrayOfStorageArray)(nil)).Elem())
}

type ArrayOfStorageContainer struct {
	StorageContainer []StorageContainer `xml:"StorageContainer,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageContainer", reflect.TypeOf((*ArrayOfStorageContainer)(nil)).Elem())
}

type ArrayOfStorageFileSystem struct {
	StorageFileSystem []StorageFileSystem `xml:"StorageFileSystem,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageFileSystem", reflect.TypeOf((*ArrayOfStorageFileSystem)(nil)).Elem())
}

type ArrayOfStorageFileSystemInfo struct {
	StorageFileSystemInfo []StorageFileSystemInfo `xml:"StorageFileSystemInfo,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageFileSystemInfo", reflect.TypeOf((*ArrayOfStorageFileSystemInfo)(nil)).Elem())
}

type ArrayOfStorageLun struct {
	StorageLun []StorageLun `xml:"StorageLun,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageLun", reflect.TypeOf((*ArrayOfStorageLun)(nil)).Elem())
}

type ArrayOfStoragePort struct {
	StoragePort []BaseStoragePort `xml:"StoragePort,omitempty,typeattr"`
}

func init() {
	types.Add("sms:ArrayOfStoragePort", reflect.TypeOf((*ArrayOfStoragePort)(nil)).Elem())
}

type ArrayOfStorageProcessor struct {
	StorageProcessor []StorageProcessor `xml:"StorageProcessor,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfStorageProcessor", reflect.TypeOf((*ArrayOfStorageProcessor)(nil)).Elem())
}

type ArrayOfSupportedVendorModelMapping struct {
	SupportedVendorModelMapping []SupportedVendorModelMapping `xml:"SupportedVendorModelMapping,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfSupportedVendorModelMapping", reflect.TypeOf((*ArrayOfSupportedVendorModelMapping)(nil)).Elem())
}

type ArrayOfTargetDeviceId struct {
	TargetDeviceId []TargetDeviceId `xml:"TargetDeviceId,omitempty"`
}

func init() {
	types.Add("sms:ArrayOfTargetDeviceId", reflect.TypeOf((*ArrayOfTargetDeviceId)(nil)).Elem())
}

type ArrayOfTargetGroupMemberInfo struct {
	TargetGroupMemberInfo []BaseTargetGroupMemberInfo `xml:"TargetGroupMemberInfo,omitempty,typeattr"`
}

func init() {
	types.Add("sms:ArrayOfTargetGroupMemberInfo", reflect.TypeOf((*ArrayOfTargetGroupMemberInfo)(nil)).Elem())
}

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

type BackingConfig struct {
	types.DynamicData

	ThinProvisionBackingIdentifier  string `xml:"thinProvisionBackingIdentifier,omitempty"`
	DeduplicationBackingIdentifier  string `xml:"deduplicationBackingIdentifier,omitempty"`
	AutoTieringEnabled              *bool  `xml:"autoTieringEnabled"`
	DeduplicationEfficiency         int64  `xml:"deduplicationEfficiency,omitempty"`
	PerformanceOptimizationInterval int64  `xml:"performanceOptimizationInterval,omitempty"`
}

func init() {
	types.Add("sms:BackingConfig", reflect.TypeOf((*BackingConfig)(nil)).Elem())
}

type BackingStoragePool struct {
	types.DynamicData

	Uuid          string `xml:"uuid"`
	Type          string `xml:"type"`
	CapacityInMB  int64  `xml:"capacityInMB"`
	UsedSpaceInMB int64  `xml:"usedSpaceInMB"`
}

func init() {
	types.Add("sms:BackingStoragePool", reflect.TypeOf((*BackingStoragePool)(nil)).Elem())
}

type CertificateAuthorityFault struct {
	ProviderRegistrationFault

	FaultCode int32 `xml:"faultCode"`
}

func init() {
	types.Add("sms:CertificateAuthorityFault", reflect.TypeOf((*CertificateAuthorityFault)(nil)).Elem())
}

type CertificateAuthorityFaultFault CertificateAuthorityFault

func init() {
	types.Add("sms:CertificateAuthorityFaultFault", reflect.TypeOf((*CertificateAuthorityFaultFault)(nil)).Elem())
}

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

type CertificateNotTrusted struct {
	ProviderRegistrationFault

	Certificate string `xml:"certificate"`
}

func init() {
	types.Add("sms:CertificateNotTrusted", reflect.TypeOf((*CertificateNotTrusted)(nil)).Elem())
}

type CertificateNotTrustedFault CertificateNotTrusted

func init() {
	types.Add("sms:CertificateNotTrustedFault", reflect.TypeOf((*CertificateNotTrustedFault)(nil)).Elem())
}

type CertificateRefreshFailed struct {
	types.MethodFault

	ProviderId []string `xml:"providerId,omitempty"`
}

func init() {
	types.Add("sms:CertificateRefreshFailed", reflect.TypeOf((*CertificateRefreshFailed)(nil)).Elem())
}

type CertificateRefreshFailedFault CertificateRefreshFailed

func init() {
	types.Add("sms:CertificateRefreshFailedFault", reflect.TypeOf((*CertificateRefreshFailedFault)(nil)).Elem())
}

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

type DatastoreBackingPoolMapping struct {
	types.DynamicData

	Datastore          []types.ManagedObjectReference `xml:"datastore"`
	BackingStoragePool []BackingStoragePool           `xml:"backingStoragePool,omitempty"`
}

func init() {
	types.Add("sms:DatastoreBackingPoolMapping", reflect.TypeOf((*DatastoreBackingPoolMapping)(nil)).Elem())
}

type DatastorePair struct {
	types.DynamicData

	Datastore1 types.ManagedObjectReference `xml:"datastore1"`
	Datastore2 types.ManagedObjectReference `xml:"datastore2"`
}

func init() {
	types.Add("sms:DatastorePair", reflect.TypeOf((*DatastorePair)(nil)).Elem())
}

type DeviceId struct {
	types.DynamicData
}

func init() {
	types.Add("sms:DeviceId", reflect.TypeOf((*DeviceId)(nil)).Elem())
}

type DrsMigrationCapabilityResult struct {
	types.DynamicData

	RecommendedDatastorePair    []DatastorePair `xml:"recommendedDatastorePair,omitempty"`
	NonRecommendedDatastorePair []DatastorePair `xml:"nonRecommendedDatastorePair,omitempty"`
}

func init() {
	types.Add("sms:DrsMigrationCapabilityResult", reflect.TypeOf((*DrsMigrationCapabilityResult)(nil)).Elem())
}

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

type EntityReference struct {
	types.DynamicData

	Id   string                    `xml:"id"`
	Type EntityReferenceEntityType `xml:"type,omitempty"`
}

func init() {
	types.Add("sms:EntityReference", reflect.TypeOf((*EntityReference)(nil)).Elem())
}

type FailoverParam struct {
	types.DynamicData

	IsPlanned                   bool                   `xml:"isPlanned"`
	CheckOnly                   bool                   `xml:"checkOnly"`
	ReplicationGroupsToFailover []ReplicationGroupData `xml:"replicationGroupsToFailover,omitempty"`
	PolicyAssociations          []PolicyAssociation    `xml:"policyAssociations,omitempty"`
}

func init() {
	types.Add("sms:FailoverParam", reflect.TypeOf((*FailoverParam)(nil)).Elem())
}

type FailoverReplicationGroupRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	FailoverParam BaseFailoverParam            `xml:"failoverParam,typeattr"`
}

func init() {
	types.Add("sms:FailoverReplicationGroupRequestType", reflect.TypeOf((*FailoverReplicationGroupRequestType)(nil)).Elem())
}

type FailoverReplicationGroup_Task FailoverReplicationGroupRequestType

func init() {
	types.Add("sms:FailoverReplicationGroup_Task", reflect.TypeOf((*FailoverReplicationGroup_Task)(nil)).Elem())
}

type FailoverReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type FailoverSuccessResult struct {
	GroupOperationResult

	NewState            string                `xml:"newState"`
	PitId               *PointInTimeReplicaId `xml:"pitId,omitempty"`
	PitIdBeforeFailover *PointInTimeReplicaId `xml:"pitIdBeforeFailover,omitempty"`
	RecoveredDeviceInfo []RecoveredDevice     `xml:"recoveredDeviceInfo,omitempty"`
	TimeStamp           *time.Time            `xml:"timeStamp"`
}

func init() {
	types.Add("sms:FailoverSuccessResult", reflect.TypeOf((*FailoverSuccessResult)(nil)).Elem())
}

type FaultDomainFilter struct {
	types.DynamicData

	ProviderId string `xml:"providerId,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainFilter", reflect.TypeOf((*FaultDomainFilter)(nil)).Elem())
}

type FaultDomainInfo struct {
	types.FaultDomainId

	Name           string                        `xml:"name,omitempty"`
	Description    string                        `xml:"description,omitempty"`
	StorageArrayId string                        `xml:"storageArrayId,omitempty"`
	Children       []types.FaultDomainId         `xml:"children,omitempty"`
	Provider       *types.ManagedObjectReference `xml:"provider,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainInfo", reflect.TypeOf((*FaultDomainInfo)(nil)).Elem())
}

type FaultDomainProviderMapping struct {
	types.DynamicData

	ActiveProvider types.ManagedObjectReference `xml:"activeProvider"`
	FaultDomainId  []types.FaultDomainId        `xml:"faultDomainId,omitempty"`
}

func init() {
	types.Add("sms:FaultDomainProviderMapping", reflect.TypeOf((*FaultDomainProviderMapping)(nil)).Elem())
}

type FcStoragePort struct {
	StoragePort

	PortWwn string `xml:"portWwn"`
	NodeWwn string `xml:"nodeWwn"`
}

func init() {
	types.Add("sms:FcStoragePort", reflect.TypeOf((*FcStoragePort)(nil)).Elem())
}

type FcoeStoragePort struct {
	StoragePort

	PortWwn string `xml:"portWwn"`
	NodeWwn string `xml:"nodeWwn"`
}

func init() {
	types.Add("sms:FcoeStoragePort", reflect.TypeOf((*FcoeStoragePort)(nil)).Elem())
}

type GroupErrorResult struct {
	GroupOperationResult

	Error []types.LocalizedMethodFault `xml:"error"`
}

func init() {
	types.Add("sms:GroupErrorResult", reflect.TypeOf((*GroupErrorResult)(nil)).Elem())
}

type GroupInfo struct {
	types.DynamicData

	GroupId types.ReplicationGroupId `xml:"groupId"`
}

func init() {
	types.Add("sms:GroupInfo", reflect.TypeOf((*GroupInfo)(nil)).Elem())
}

type GroupOperationResult struct {
	types.DynamicData

	GroupId types.ReplicationGroupId     `xml:"groupId"`
	Warning []types.LocalizedMethodFault `xml:"warning,omitempty"`
}

func init() {
	types.Add("sms:GroupOperationResult", reflect.TypeOf((*GroupOperationResult)(nil)).Elem())
}

type InactiveProvider struct {
	types.MethodFault

	Mapping []FaultDomainProviderMapping `xml:"mapping,omitempty"`
}

func init() {
	types.Add("sms:InactiveProvider", reflect.TypeOf((*InactiveProvider)(nil)).Elem())
}

type InactiveProviderFault InactiveProvider

func init() {
	types.Add("sms:InactiveProviderFault", reflect.TypeOf((*InactiveProviderFault)(nil)).Elem())
}

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

type InvalidCertificate struct {
	ProviderRegistrationFault

	Certificate string `xml:"certificate"`
}

func init() {
	types.Add("sms:InvalidCertificate", reflect.TypeOf((*InvalidCertificate)(nil)).Elem())
}

type InvalidCertificateFault InvalidCertificate

func init() {
	types.Add("sms:InvalidCertificateFault", reflect.TypeOf((*InvalidCertificateFault)(nil)).Elem())
}

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

type InvalidReplicationState struct {
	SmsReplicationFault

	DesiredState []string `xml:"desiredState,omitempty"`
	CurrentState string   `xml:"currentState"`
}

func init() {
	types.Add("sms:InvalidReplicationState", reflect.TypeOf((*InvalidReplicationState)(nil)).Elem())
}

type InvalidReplicationStateFault InvalidReplicationState

func init() {
	types.Add("sms:InvalidReplicationStateFault", reflect.TypeOf((*InvalidReplicationStateFault)(nil)).Elem())
}

type InvalidSession struct {
	types.NoPermission

	SessionCookie string `xml:"sessionCookie"`
}

func init() {
	types.Add("sms:InvalidSession", reflect.TypeOf((*InvalidSession)(nil)).Elem())
}

type InvalidSessionFault InvalidSession

func init() {
	types.Add("sms:InvalidSessionFault", reflect.TypeOf((*InvalidSessionFault)(nil)).Elem())
}

type InvalidUrl struct {
	ProviderRegistrationFault

	Url string `xml:"url"`
}

func init() {
	types.Add("sms:InvalidUrl", reflect.TypeOf((*InvalidUrl)(nil)).Elem())
}

type InvalidUrlFault InvalidUrl

func init() {
	types.Add("sms:InvalidUrlFault", reflect.TypeOf((*InvalidUrlFault)(nil)).Elem())
}

type IscsiStoragePort struct {
	StoragePort

	Identifier string `xml:"identifier"`
}

func init() {
	types.Add("sms:IscsiStoragePort", reflect.TypeOf((*IscsiStoragePort)(nil)).Elem())
}

type LunHbaAssociation struct {
	types.DynamicData

	CanonicalName string                     `xml:"canonicalName"`
	Hba           []types.HostHostBusAdapter `xml:"hba"`
}

func init() {
	types.Add("sms:LunHbaAssociation", reflect.TypeOf((*LunHbaAssociation)(nil)).Elem())
}

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

type NameValuePair struct {
	types.DynamicData

	ParameterName  string `xml:"parameterName"`
	ParameterValue string `xml:"parameterValue"`
}

func init() {
	types.Add("sms:NameValuePair", reflect.TypeOf((*NameValuePair)(nil)).Elem())
}

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

type NoValidReplica struct {
	SmsReplicationFault

	DeviceId BaseDeviceId `xml:"deviceId,omitempty,typeattr"`
}

func init() {
	types.Add("sms:NoValidReplica", reflect.TypeOf((*NoValidReplica)(nil)).Elem())
}

type NoValidReplicaFault NoValidReplica

func init() {
	types.Add("sms:NoValidReplicaFault", reflect.TypeOf((*NoValidReplicaFault)(nil)).Elem())
}

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

type PointInTimeReplicaId struct {
	types.DynamicData

	Id string `xml:"id"`
}

func init() {
	types.Add("sms:PointInTimeReplicaId", reflect.TypeOf((*PointInTimeReplicaId)(nil)).Elem())
}

type PointInTimeReplicaInfo struct {
	types.DynamicData

	Id        PointInTimeReplicaId `xml:"id"`
	PitName   string               `xml:"pitName"`
	TimeStamp time.Time            `xml:"timeStamp"`
	Tags      []string             `xml:"tags,omitempty"`
}

func init() {
	types.Add("sms:PointInTimeReplicaInfo", reflect.TypeOf((*PointInTimeReplicaInfo)(nil)).Elem())
}

type PolicyAssociation struct {
	types.DynamicData

	Id        BaseDeviceId                 `xml:"id,typeattr"`
	PolicyId  string                       `xml:"policyId"`
	Datastore types.ManagedObjectReference `xml:"datastore"`
}

func init() {
	types.Add("sms:PolicyAssociation", reflect.TypeOf((*PolicyAssociation)(nil)).Elem())
}

type PrepareFailoverReplicationGroupRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	GroupId []types.ReplicationGroupId   `xml:"groupId,omitempty"`
}

func init() {
	types.Add("sms:PrepareFailoverReplicationGroupRequestType", reflect.TypeOf((*PrepareFailoverReplicationGroupRequestType)(nil)).Elem())
}

type PrepareFailoverReplicationGroup_Task PrepareFailoverReplicationGroupRequestType

func init() {
	types.Add("sms:PrepareFailoverReplicationGroup_Task", reflect.TypeOf((*PrepareFailoverReplicationGroup_Task)(nil)).Elem())
}

type PrepareFailoverReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type PromoteParam struct {
	types.DynamicData

	IsPlanned                  bool                       `xml:"isPlanned"`
	ReplicationGroupsToPromote []types.ReplicationGroupId `xml:"replicationGroupsToPromote,omitempty"`
}

func init() {
	types.Add("sms:PromoteParam", reflect.TypeOf((*PromoteParam)(nil)).Elem())
}

type PromoteReplicationGroupRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	PromoteParam PromoteParam                 `xml:"promoteParam"`
}

func init() {
	types.Add("sms:PromoteReplicationGroupRequestType", reflect.TypeOf((*PromoteReplicationGroupRequestType)(nil)).Elem())
}

type PromoteReplicationGroup_Task PromoteReplicationGroupRequestType

func init() {
	types.Add("sms:PromoteReplicationGroup_Task", reflect.TypeOf((*PromoteReplicationGroup_Task)(nil)).Elem())
}

type PromoteReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

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

type ProviderOutOfProvisioningResource struct {
	types.MethodFault

	ProvisioningResourceId string `xml:"provisioningResourceId"`
	AvailableBefore        int64  `xml:"availableBefore,omitempty"`
	AvailableAfter         int64  `xml:"availableAfter,omitempty"`
	Total                  int64  `xml:"total,omitempty"`
	IsTransient            *bool  `xml:"isTransient"`
}

func init() {
	types.Add("sms:ProviderOutOfProvisioningResource", reflect.TypeOf((*ProviderOutOfProvisioningResource)(nil)).Elem())
}

type ProviderOutOfProvisioningResourceFault ProviderOutOfProvisioningResource

func init() {
	types.Add("sms:ProviderOutOfProvisioningResourceFault", reflect.TypeOf((*ProviderOutOfProvisioningResourceFault)(nil)).Elem())
}

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
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QueryAboutInfoRequestType", reflect.TypeOf((*QueryAboutInfoRequestType)(nil)).Elem())
}

type QueryAboutInfoResponse struct {
	Returnval SmsAboutInfo `xml:"returnval"`
}

type QueryActiveAlarm QueryActiveAlarmRequestType

func init() {
	types.Add("sms:QueryActiveAlarm", reflect.TypeOf((*QueryActiveAlarm)(nil)).Elem())
}

type QueryActiveAlarmRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	AlarmFilter *AlarmFilter                 `xml:"alarmFilter,omitempty"`
}

func init() {
	types.Add("sms:QueryActiveAlarmRequestType", reflect.TypeOf((*QueryActiveAlarmRequestType)(nil)).Elem())
}

type QueryActiveAlarmResponse struct {
	Returnval *AlarmResult `xml:"returnval,omitempty"`
}

type QueryArray QueryArrayRequestType

func init() {
	types.Add("sms:QueryArray", reflect.TypeOf((*QueryArray)(nil)).Elem())
}

type QueryArrayAssociatedWithLun QueryArrayAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryArrayAssociatedWithLun", reflect.TypeOf((*QueryArrayAssociatedWithLun)(nil)).Elem())
}

type QueryArrayAssociatedWithLunRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	CanonicalName string                       `xml:"canonicalName"`
}

func init() {
	types.Add("sms:QueryArrayAssociatedWithLunRequestType", reflect.TypeOf((*QueryArrayAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryArrayAssociatedWithLunResponse struct {
	Returnval *StorageArray `xml:"returnval,omitempty"`
}

type QueryArrayRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ProviderId []string                     `xml:"providerId,omitempty"`
}

func init() {
	types.Add("sms:QueryArrayRequestType", reflect.TypeOf((*QueryArrayRequestType)(nil)).Elem())
}

type QueryArrayResponse struct {
	Returnval []StorageArray `xml:"returnval,omitempty"`
}

type QueryAssociatedBackingStoragePool QueryAssociatedBackingStoragePoolRequestType

func init() {
	types.Add("sms:QueryAssociatedBackingStoragePool", reflect.TypeOf((*QueryAssociatedBackingStoragePool)(nil)).Elem())
}

type QueryAssociatedBackingStoragePoolRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	EntityId   string                       `xml:"entityId,omitempty"`
	EntityType string                       `xml:"entityType,omitempty"`
}

func init() {
	types.Add("sms:QueryAssociatedBackingStoragePoolRequestType", reflect.TypeOf((*QueryAssociatedBackingStoragePoolRequestType)(nil)).Elem())
}

type QueryAssociatedBackingStoragePoolResponse struct {
	Returnval []BackingStoragePool `xml:"returnval,omitempty"`
}

type QueryDatastoreBackingPoolMapping QueryDatastoreBackingPoolMappingRequestType

func init() {
	types.Add("sms:QueryDatastoreBackingPoolMapping", reflect.TypeOf((*QueryDatastoreBackingPoolMapping)(nil)).Elem())
}

type QueryDatastoreBackingPoolMappingRequestType struct {
	This      types.ManagedObjectReference   `xml:"_this"`
	Datastore []types.ManagedObjectReference `xml:"datastore"`
}

func init() {
	types.Add("sms:QueryDatastoreBackingPoolMappingRequestType", reflect.TypeOf((*QueryDatastoreBackingPoolMappingRequestType)(nil)).Elem())
}

type QueryDatastoreBackingPoolMappingResponse struct {
	Returnval []DatastoreBackingPoolMapping `xml:"returnval"`
}

type QueryDatastoreCapability QueryDatastoreCapabilityRequestType

func init() {
	types.Add("sms:QueryDatastoreCapability", reflect.TypeOf((*QueryDatastoreCapability)(nil)).Elem())
}

type QueryDatastoreCapabilityRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Datastore types.ManagedObjectReference `xml:"datastore"`
}

func init() {
	types.Add("sms:QueryDatastoreCapabilityRequestType", reflect.TypeOf((*QueryDatastoreCapabilityRequestType)(nil)).Elem())
}

type QueryDatastoreCapabilityResponse struct {
	Returnval *StorageCapability `xml:"returnval,omitempty"`
}

type QueryDrsMigrationCapabilityForPerformance QueryDrsMigrationCapabilityForPerformanceRequestType

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformance", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformance)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceEx QueryDrsMigrationCapabilityForPerformanceExRequestType

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceEx", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceEx)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceExRequestType struct {
	This      types.ManagedObjectReference   `xml:"_this"`
	Datastore []types.ManagedObjectReference `xml:"datastore"`
}

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceExRequestType", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceExRequestType)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceExResponse struct {
	Returnval DrsMigrationCapabilityResult `xml:"returnval"`
}

type QueryDrsMigrationCapabilityForPerformanceRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	SrcDatastore types.ManagedObjectReference `xml:"srcDatastore"`
	DstDatastore types.ManagedObjectReference `xml:"dstDatastore"`
}

func init() {
	types.Add("sms:QueryDrsMigrationCapabilityForPerformanceRequestType", reflect.TypeOf((*QueryDrsMigrationCapabilityForPerformanceRequestType)(nil)).Elem())
}

type QueryDrsMigrationCapabilityForPerformanceResponse struct {
	Returnval bool `xml:"returnval"`
}

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

type QueryFaultDomainRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	Filter *FaultDomainFilter           `xml:"filter,omitempty"`
}

func init() {
	types.Add("sms:QueryFaultDomainRequestType", reflect.TypeOf((*QueryFaultDomainRequestType)(nil)).Elem())
}

type QueryFaultDomainResponse struct {
	Returnval []types.FaultDomainId `xml:"returnval,omitempty"`
}

type QueryFileSystemAssociatedWithArray QueryFileSystemAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryFileSystemAssociatedWithArray", reflect.TypeOf((*QueryFileSystemAssociatedWithArray)(nil)).Elem())
}

type QueryFileSystemAssociatedWithArrayRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryFileSystemAssociatedWithArrayRequestType", reflect.TypeOf((*QueryFileSystemAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryFileSystemAssociatedWithArrayResponse struct {
	Returnval []StorageFileSystem `xml:"returnval,omitempty"`
}

type QueryHostAssociatedWithLun QueryHostAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryHostAssociatedWithLun", reflect.TypeOf((*QueryHostAssociatedWithLun)(nil)).Elem())
}

type QueryHostAssociatedWithLunRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Scsi3Id string                       `xml:"scsi3Id"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryHostAssociatedWithLunRequestType", reflect.TypeOf((*QueryHostAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryHostAssociatedWithLunResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type QueryLunAssociatedWithArray QueryLunAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryLunAssociatedWithArray", reflect.TypeOf((*QueryLunAssociatedWithArray)(nil)).Elem())
}

type QueryLunAssociatedWithArrayRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryLunAssociatedWithArrayRequestType", reflect.TypeOf((*QueryLunAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryLunAssociatedWithArrayResponse struct {
	Returnval []StorageLun `xml:"returnval,omitempty"`
}

type QueryLunAssociatedWithPort QueryLunAssociatedWithPortRequestType

func init() {
	types.Add("sms:QueryLunAssociatedWithPort", reflect.TypeOf((*QueryLunAssociatedWithPort)(nil)).Elem())
}

type QueryLunAssociatedWithPortRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	PortId  string                       `xml:"portId"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryLunAssociatedWithPortRequestType", reflect.TypeOf((*QueryLunAssociatedWithPortRequestType)(nil)).Elem())
}

type QueryLunAssociatedWithPortResponse struct {
	Returnval []StorageLun `xml:"returnval,omitempty"`
}

type QueryNfsDatastoreAssociatedWithFileSystem QueryNfsDatastoreAssociatedWithFileSystemRequestType

func init() {
	types.Add("sms:QueryNfsDatastoreAssociatedWithFileSystem", reflect.TypeOf((*QueryNfsDatastoreAssociatedWithFileSystem)(nil)).Elem())
}

type QueryNfsDatastoreAssociatedWithFileSystemRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	FileSystemId string                       `xml:"fileSystemId"`
	ArrayId      string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryNfsDatastoreAssociatedWithFileSystemRequestType", reflect.TypeOf((*QueryNfsDatastoreAssociatedWithFileSystemRequestType)(nil)).Elem())
}

type QueryNfsDatastoreAssociatedWithFileSystemResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type QueryNotSupported struct {
	types.InvalidArgument

	EntityType        EntityReferenceEntityType `xml:"entityType,omitempty"`
	RelatedEntityType EntityReferenceEntityType `xml:"relatedEntityType"`
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

type QueryPointInTimeReplicaParam struct {
	types.DynamicData

	ReplicaTimeQueryParam *ReplicaQueryIntervalParam `xml:"replicaTimeQueryParam,omitempty"`
	PitName               string                     `xml:"pitName,omitempty"`
	Tags                  []string                   `xml:"tags,omitempty"`
	PreferDetails         *bool                      `xml:"preferDetails"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaParam", reflect.TypeOf((*QueryPointInTimeReplicaParam)(nil)).Elem())
}

type QueryPointInTimeReplicaRequestType struct {
	This       types.ManagedObjectReference  `xml:"_this"`
	GroupId    []types.ReplicationGroupId    `xml:"groupId,omitempty"`
	QueryParam *QueryPointInTimeReplicaParam `xml:"queryParam,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaRequestType", reflect.TypeOf((*QueryPointInTimeReplicaRequestType)(nil)).Elem())
}

type QueryPointInTimeReplicaResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr"`
}

type QueryPointInTimeReplicaSuccessResult struct {
	GroupOperationResult

	ReplicaInfo []PointInTimeReplicaInfo `xml:"replicaInfo,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaSuccessResult", reflect.TypeOf((*QueryPointInTimeReplicaSuccessResult)(nil)).Elem())
}

type QueryPointInTimeReplicaSummaryResult struct {
	GroupOperationResult

	IntervalResults []ReplicaIntervalQueryResult `xml:"intervalResults,omitempty"`
}

func init() {
	types.Add("sms:QueryPointInTimeReplicaSummaryResult", reflect.TypeOf((*QueryPointInTimeReplicaSummaryResult)(nil)).Elem())
}

type QueryPortAssociatedWithArray QueryPortAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithArray", reflect.TypeOf((*QueryPortAssociatedWithArray)(nil)).Elem())
}

type QueryPortAssociatedWithArrayRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithArrayRequestType", reflect.TypeOf((*QueryPortAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithArrayResponse struct {
	Returnval []BaseStoragePort `xml:"returnval,omitempty,typeattr"`
}

type QueryPortAssociatedWithLun QueryPortAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithLun", reflect.TypeOf((*QueryPortAssociatedWithLun)(nil)).Elem())
}

type QueryPortAssociatedWithLunRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Scsi3Id string                       `xml:"scsi3Id"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithLunRequestType", reflect.TypeOf((*QueryPortAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithLunResponse struct {
	Returnval BaseStoragePort `xml:"returnval,omitempty,typeattr"`
}

type QueryPortAssociatedWithProcessor QueryPortAssociatedWithProcessorRequestType

func init() {
	types.Add("sms:QueryPortAssociatedWithProcessor", reflect.TypeOf((*QueryPortAssociatedWithProcessor)(nil)).Elem())
}

type QueryPortAssociatedWithProcessorRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	ProcessorId string                       `xml:"processorId"`
	ArrayId     string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryPortAssociatedWithProcessorRequestType", reflect.TypeOf((*QueryPortAssociatedWithProcessorRequestType)(nil)).Elem())
}

type QueryPortAssociatedWithProcessorResponse struct {
	Returnval []BaseStoragePort `xml:"returnval,omitempty,typeattr"`
}

type QueryProcessorAssociatedWithArray QueryProcessorAssociatedWithArrayRequestType

func init() {
	types.Add("sms:QueryProcessorAssociatedWithArray", reflect.TypeOf((*QueryProcessorAssociatedWithArray)(nil)).Elem())
}

type QueryProcessorAssociatedWithArrayRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryProcessorAssociatedWithArrayRequestType", reflect.TypeOf((*QueryProcessorAssociatedWithArrayRequestType)(nil)).Elem())
}

type QueryProcessorAssociatedWithArrayResponse struct {
	Returnval []StorageProcessor `xml:"returnval,omitempty"`
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
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QueryProviderInfoRequestType", reflect.TypeOf((*QueryProviderInfoRequestType)(nil)).Elem())
}

type QueryProviderInfoResponse struct {
	Returnval BaseSmsProviderInfo `xml:"returnval,typeattr"`
}

type QueryProviderRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QueryProviderRequestType", reflect.TypeOf((*QueryProviderRequestType)(nil)).Elem())
}

type QueryProviderResponse struct {
	Returnval []types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type QueryReplicationGroup QueryReplicationGroupRequestType

func init() {
	types.Add("sms:QueryReplicationGroup", reflect.TypeOf((*QueryReplicationGroup)(nil)).Elem())
}

type QueryReplicationGroupInfo QueryReplicationGroupInfoRequestType

func init() {
	types.Add("sms:QueryReplicationGroupInfo", reflect.TypeOf((*QueryReplicationGroupInfo)(nil)).Elem())
}

type QueryReplicationGroupInfoRequestType struct {
	This     types.ManagedObjectReference `xml:"_this"`
	RgFilter ReplicationGroupFilter       `xml:"rgFilter"`
}

func init() {
	types.Add("sms:QueryReplicationGroupInfoRequestType", reflect.TypeOf((*QueryReplicationGroupInfoRequestType)(nil)).Elem())
}

type QueryReplicationGroupInfoResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr"`
}

type QueryReplicationGroupRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	GroupId []types.ReplicationGroupId   `xml:"groupId,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationGroupRequestType", reflect.TypeOf((*QueryReplicationGroupRequestType)(nil)).Elem())
}

type QueryReplicationGroupResponse struct {
	Returnval []BaseGroupOperationResult `xml:"returnval,omitempty,typeattr"`
}

type QueryReplicationGroupSuccessResult struct {
	GroupOperationResult

	RgInfo BaseGroupInfo `xml:"rgInfo,typeattr"`
}

func init() {
	types.Add("sms:QueryReplicationGroupSuccessResult", reflect.TypeOf((*QueryReplicationGroupSuccessResult)(nil)).Elem())
}

type QueryReplicationPeer QueryReplicationPeerRequestType

func init() {
	types.Add("sms:QueryReplicationPeer", reflect.TypeOf((*QueryReplicationPeer)(nil)).Elem())
}

type QueryReplicationPeerRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	FaultDomainId []types.FaultDomainId        `xml:"faultDomainId,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationPeerRequestType", reflect.TypeOf((*QueryReplicationPeerRequestType)(nil)).Elem())
}

type QueryReplicationPeerResponse struct {
	Returnval []QueryReplicationPeerResult `xml:"returnval,omitempty"`
}

type QueryReplicationPeerResult struct {
	types.DynamicData

	SourceDomain types.FaultDomainId          `xml:"sourceDomain"`
	TargetDomain []types.FaultDomainId        `xml:"targetDomain,omitempty"`
	Error        []types.LocalizedMethodFault `xml:"error,omitempty"`
	Warning      []types.LocalizedMethodFault `xml:"warning,omitempty"`
}

func init() {
	types.Add("sms:QueryReplicationPeerResult", reflect.TypeOf((*QueryReplicationPeerResult)(nil)).Elem())
}

type QuerySessionManager QuerySessionManagerRequestType

func init() {
	types.Add("sms:QuerySessionManager", reflect.TypeOf((*QuerySessionManager)(nil)).Elem())
}

type QuerySessionManagerRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QuerySessionManagerRequestType", reflect.TypeOf((*QuerySessionManagerRequestType)(nil)).Elem())
}

type QuerySessionManagerResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type QuerySmsTaskInfo QuerySmsTaskInfoRequestType

func init() {
	types.Add("sms:QuerySmsTaskInfo", reflect.TypeOf((*QuerySmsTaskInfo)(nil)).Elem())
}

type QuerySmsTaskInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QuerySmsTaskInfoRequestType", reflect.TypeOf((*QuerySmsTaskInfoRequestType)(nil)).Elem())
}

type QuerySmsTaskInfoResponse struct {
	Returnval SmsTaskInfo `xml:"returnval"`
}

type QuerySmsTaskResult QuerySmsTaskResultRequestType

func init() {
	types.Add("sms:QuerySmsTaskResult", reflect.TypeOf((*QuerySmsTaskResult)(nil)).Elem())
}

type QuerySmsTaskResultRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QuerySmsTaskResultRequestType", reflect.TypeOf((*QuerySmsTaskResultRequestType)(nil)).Elem())
}

type QuerySmsTaskResultResponse struct {
	Returnval types.AnyType `xml:"returnval,omitempty,typeattr"`
}

type QueryStorageContainer QueryStorageContainerRequestType

func init() {
	types.Add("sms:QueryStorageContainer", reflect.TypeOf((*QueryStorageContainer)(nil)).Elem())
}

type QueryStorageContainerRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	ContainerSpec *StorageContainerSpec        `xml:"containerSpec,omitempty"`
}

func init() {
	types.Add("sms:QueryStorageContainerRequestType", reflect.TypeOf((*QueryStorageContainerRequestType)(nil)).Elem())
}

type QueryStorageContainerResponse struct {
	Returnval *StorageContainerResult `xml:"returnval,omitempty"`
}

type QueryStorageManager QueryStorageManagerRequestType

func init() {
	types.Add("sms:QueryStorageManager", reflect.TypeOf((*QueryStorageManager)(nil)).Elem())
}

type QueryStorageManagerRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:QueryStorageManagerRequestType", reflect.TypeOf((*QueryStorageManagerRequestType)(nil)).Elem())
}

type QueryStorageManagerResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type QueryVmfsDatastoreAssociatedWithLun QueryVmfsDatastoreAssociatedWithLunRequestType

func init() {
	types.Add("sms:QueryVmfsDatastoreAssociatedWithLun", reflect.TypeOf((*QueryVmfsDatastoreAssociatedWithLun)(nil)).Elem())
}

type QueryVmfsDatastoreAssociatedWithLunRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	Scsi3Id string                       `xml:"scsi3Id"`
	ArrayId string                       `xml:"arrayId"`
}

func init() {
	types.Add("sms:QueryVmfsDatastoreAssociatedWithLunRequestType", reflect.TypeOf((*QueryVmfsDatastoreAssociatedWithLunRequestType)(nil)).Elem())
}

type QueryVmfsDatastoreAssociatedWithLunResponse struct {
	Returnval *types.ManagedObjectReference `xml:"returnval,omitempty"`
}

type RecoveredDevice struct {
	types.DynamicData

	TargetDeviceId    *ReplicaId                   `xml:"targetDeviceId,omitempty"`
	RecoveredDeviceId BaseDeviceId                 `xml:"recoveredDeviceId,omitempty,typeattr"`
	SourceDeviceId    BaseDeviceId                 `xml:"sourceDeviceId,typeattr"`
	Info              []string                     `xml:"info,omitempty"`
	Datastore         types.ManagedObjectReference `xml:"datastore"`
	RecoveredDiskInfo []RecoveredDiskInfo          `xml:"recoveredDiskInfo,omitempty"`
	Error             *types.LocalizedMethodFault  `xml:"error,omitempty"`
	Warnings          []types.LocalizedMethodFault `xml:"warnings,omitempty"`
}

func init() {
	types.Add("sms:RecoveredDevice", reflect.TypeOf((*RecoveredDevice)(nil)).Elem())
}

type RecoveredDiskInfo struct {
	types.DynamicData

	DeviceKey int32  `xml:"deviceKey"`
	DsUrl     string `xml:"dsUrl"`
	DiskPath  string `xml:"diskPath"`
}

func init() {
	types.Add("sms:RecoveredDiskInfo", reflect.TypeOf((*RecoveredDiskInfo)(nil)).Elem())
}

type RecoveredTargetGroupMemberInfo struct {
	TargetGroupMemberInfo

	RecoveredDeviceId BaseDeviceId `xml:"recoveredDeviceId,omitempty,typeattr"`
}

func init() {
	types.Add("sms:RecoveredTargetGroupMemberInfo", reflect.TypeOf((*RecoveredTargetGroupMemberInfo)(nil)).Elem())
}

type RegisterProviderRequestType struct {
	This         types.ManagedObjectReference `xml:"_this"`
	ProviderSpec BaseSmsProviderSpec          `xml:"providerSpec,typeattr"`
}

func init() {
	types.Add("sms:RegisterProviderRequestType", reflect.TypeOf((*RegisterProviderRequestType)(nil)).Elem())
}

type RegisterProvider_Task RegisterProviderRequestType

func init() {
	types.Add("sms:RegisterProvider_Task", reflect.TypeOf((*RegisterProvider_Task)(nil)).Elem())
}

type RegisterProvider_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type RelatedStorageArray struct {
	types.DynamicData

	ArrayId    string `xml:"arrayId"`
	Active     bool   `xml:"active"`
	Manageable bool   `xml:"manageable"`
	Priority   int32  `xml:"priority"`
}

func init() {
	types.Add("sms:RelatedStorageArray", reflect.TypeOf((*RelatedStorageArray)(nil)).Elem())
}

type ReplicaId struct {
	types.DynamicData

	Id string `xml:"id"`
}

func init() {
	types.Add("sms:ReplicaId", reflect.TypeOf((*ReplicaId)(nil)).Elem())
}

type ReplicaIntervalQueryResult struct {
	types.DynamicData

	FromDate time.Time `xml:"fromDate"`
	ToDate   time.Time `xml:"toDate"`
	Number   int32     `xml:"number"`
}

func init() {
	types.Add("sms:ReplicaIntervalQueryResult", reflect.TypeOf((*ReplicaIntervalQueryResult)(nil)).Elem())
}

type ReplicaQueryIntervalParam struct {
	types.DynamicData

	FromDate *time.Time `xml:"fromDate"`
	ToDate   *time.Time `xml:"toDate"`
	Number   int32      `xml:"number,omitempty"`
}

func init() {
	types.Add("sms:ReplicaQueryIntervalParam", reflect.TypeOf((*ReplicaQueryIntervalParam)(nil)).Elem())
}

type ReplicationGroupData struct {
	types.DynamicData

	GroupId types.ReplicationGroupId `xml:"groupId"`
	PitId   *PointInTimeReplicaId    `xml:"pitId,omitempty"`
}

func init() {
	types.Add("sms:ReplicationGroupData", reflect.TypeOf((*ReplicationGroupData)(nil)).Elem())
}

type ReplicationGroupFilter struct {
	types.DynamicData

	GroupId []types.ReplicationGroupId `xml:"groupId,omitempty"`
}

func init() {
	types.Add("sms:ReplicationGroupFilter", reflect.TypeOf((*ReplicationGroupFilter)(nil)).Elem())
}

type ReplicationTargetInfo struct {
	types.DynamicData

	TargetGroupId                   types.ReplicationGroupId `xml:"targetGroupId"`
	ReplicationAgreementDescription string                   `xml:"replicationAgreementDescription,omitempty"`
}

func init() {
	types.Add("sms:ReplicationTargetInfo", reflect.TypeOf((*ReplicationTargetInfo)(nil)).Elem())
}

type ReverseReplicateGroupRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	GroupId []types.ReplicationGroupId   `xml:"groupId,omitempty"`
}

func init() {
	types.Add("sms:ReverseReplicateGroupRequestType", reflect.TypeOf((*ReverseReplicateGroupRequestType)(nil)).Elem())
}

type ReverseReplicateGroup_Task ReverseReplicateGroupRequestType

func init() {
	types.Add("sms:ReverseReplicateGroup_Task", reflect.TypeOf((*ReverseReplicateGroup_Task)(nil)).Elem())
}

type ReverseReplicateGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type ReverseReplicationSuccessResult struct {
	GroupOperationResult

	NewGroupId types.DeviceGroupId `xml:"newGroupId"`
}

func init() {
	types.Add("sms:ReverseReplicationSuccessResult", reflect.TypeOf((*ReverseReplicationSuccessResult)(nil)).Elem())
}

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

type SmsAboutInfo struct {
	types.DynamicData

	Name           string `xml:"name"`
	FullName       string `xml:"fullName"`
	Vendor         string `xml:"vendor"`
	ApiVersion     string `xml:"apiVersion"`
	InstanceUuid   string `xml:"instanceUuid"`
	VasaApiVersion string `xml:"vasaApiVersion,omitempty"`
}

func init() {
	types.Add("sms:SmsAboutInfo", reflect.TypeOf((*SmsAboutInfo)(nil)).Elem())
}

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

type SmsProviderInfo struct {
	types.DynamicData

	Uid         string `xml:"uid"`
	Name        string `xml:"name"`
	Description string `xml:"description,omitempty"`
	Version     string `xml:"version,omitempty"`
}

func init() {
	types.Add("sms:SmsProviderInfo", reflect.TypeOf((*SmsProviderInfo)(nil)).Elem())
}

type SmsProviderSpec struct {
	types.DynamicData

	Name        string `xml:"name"`
	Description string `xml:"description,omitempty"`
}

func init() {
	types.Add("sms:SmsProviderSpec", reflect.TypeOf((*SmsProviderSpec)(nil)).Elem())
}

type SmsRefreshCACertificatesAndCRLsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ProviderId []string                     `xml:"providerId,omitempty"`
}

func init() {
	types.Add("sms:SmsRefreshCACertificatesAndCRLsRequestType", reflect.TypeOf((*SmsRefreshCACertificatesAndCRLsRequestType)(nil)).Elem())
}

type SmsRefreshCACertificatesAndCRLs_Task SmsRefreshCACertificatesAndCRLsRequestType

func init() {
	types.Add("sms:SmsRefreshCACertificatesAndCRLs_Task", reflect.TypeOf((*SmsRefreshCACertificatesAndCRLs_Task)(nil)).Elem())
}

type SmsRefreshCACertificatesAndCRLs_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

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

type SmsResourceInUse struct {
	types.ResourceInUse

	DeviceIds []BaseDeviceId `xml:"deviceIds,omitempty,typeattr"`
}

func init() {
	types.Add("sms:SmsResourceInUse", reflect.TypeOf((*SmsResourceInUse)(nil)).Elem())
}

type SmsResourceInUseFault SmsResourceInUse

func init() {
	types.Add("sms:SmsResourceInUseFault", reflect.TypeOf((*SmsResourceInUseFault)(nil)).Elem())
}

type SmsTaskInfo struct {
	types.DynamicData

	Key            string                        `xml:"key"`
	Task           types.ManagedObjectReference  `xml:"task"`
	Object         *types.ManagedObjectReference `xml:"object,omitempty"`
	Error          *types.LocalizedMethodFault   `xml:"error,omitempty"`
	Result         types.AnyType                 `xml:"result,omitempty,typeattr"`
	StartTime      *time.Time                    `xml:"startTime"`
	CompletionTime *time.Time                    `xml:"completionTime"`
	State          string                        `xml:"state"`
	Progress       int32                         `xml:"progress,omitempty"`
}

func init() {
	types.Add("sms:SmsTaskInfo", reflect.TypeOf((*SmsTaskInfo)(nil)).Elem())
}

type SourceGroupInfo struct {
	GroupInfo

	Name        string                  `xml:"name,omitempty"`
	Description string                  `xml:"description,omitempty"`
	State       string                  `xml:"state"`
	Replica     []ReplicationTargetInfo `xml:"replica,omitempty"`
	MemberInfo  []SourceGroupMemberInfo `xml:"memberInfo,omitempty"`
}

func init() {
	types.Add("sms:SourceGroupInfo", reflect.TypeOf((*SourceGroupInfo)(nil)).Elem())
}

type SourceGroupMemberInfo struct {
	types.DynamicData

	DeviceId BaseDeviceId     `xml:"deviceId,typeattr"`
	TargetId []TargetDeviceId `xml:"targetId,omitempty"`
}

func init() {
	types.Add("sms:SourceGroupMemberInfo", reflect.TypeOf((*SourceGroupMemberInfo)(nil)).Elem())
}

type StorageAlarm struct {
	types.DynamicData

	AlarmId        int64           `xml:"alarmId"`
	AlarmType      string          `xml:"alarmType"`
	ContainerId    string          `xml:"containerId,omitempty"`
	ObjectId       string          `xml:"objectId,omitempty"`
	ObjectType     string          `xml:"objectType"`
	Status         string          `xml:"status"`
	AlarmTimeStamp time.Time       `xml:"alarmTimeStamp"`
	MessageId      string          `xml:"messageId"`
	ParameterList  []NameValuePair `xml:"parameterList,omitempty"`
	AlarmObject    types.AnyType   `xml:"alarmObject,omitempty,typeattr"`
}

func init() {
	types.Add("sms:StorageAlarm", reflect.TypeOf((*StorageAlarm)(nil)).Elem())
}

type StorageArray struct {
	types.DynamicData

	Name                         string   `xml:"name"`
	Uuid                         string   `xml:"uuid"`
	VendorId                     string   `xml:"vendorId"`
	ModelId                      string   `xml:"modelId"`
	Firmware                     string   `xml:"firmware,omitempty"`
	AlternateName                []string `xml:"alternateName,omitempty"`
	SupportedBlockInterface      []string `xml:"supportedBlockInterface,omitempty"`
	SupportedFileSystemInterface []string `xml:"supportedFileSystemInterface,omitempty"`
	SupportedProfile             []string `xml:"supportedProfile,omitempty"`
	Priority                     int32    `xml:"priority,omitempty"`
}

func init() {
	types.Add("sms:StorageArray", reflect.TypeOf((*StorageArray)(nil)).Elem())
}

type StorageCapability struct {
	types.DynamicData

	Uuid        string `xml:"uuid"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

func init() {
	types.Add("sms:StorageCapability", reflect.TypeOf((*StorageCapability)(nil)).Elem())
}

type StorageContainer struct {
	types.DynamicData

	Uuid            string   `xml:"uuid"`
	Name            string   `xml:"name"`
	MaxVvolSizeInMB int64    `xml:"maxVvolSizeInMB"`
	ProviderId      []string `xml:"providerId"`
	ArrayId         []string `xml:"arrayId"`
}

func init() {
	types.Add("sms:StorageContainer", reflect.TypeOf((*StorageContainer)(nil)).Elem())
}

type StorageContainerResult struct {
	types.DynamicData

	StorageContainer []StorageContainer    `xml:"storageContainer,omitempty"`
	ProviderInfo     []BaseSmsProviderInfo `xml:"providerInfo,omitempty,typeattr"`
}

func init() {
	types.Add("sms:StorageContainerResult", reflect.TypeOf((*StorageContainerResult)(nil)).Elem())
}

type StorageContainerSpec struct {
	types.DynamicData

	ContainerId []string `xml:"containerId,omitempty"`
}

func init() {
	types.Add("sms:StorageContainerSpec", reflect.TypeOf((*StorageContainerSpec)(nil)).Elem())
}

type StorageFileSystem struct {
	types.DynamicData

	Uuid                    string                  `xml:"uuid"`
	Info                    []StorageFileSystemInfo `xml:"info"`
	NativeSnapshotSupported bool                    `xml:"nativeSnapshotSupported"`
	ThinProvisioningStatus  string                  `xml:"thinProvisioningStatus"`
	Type                    string                  `xml:"type"`
	Version                 string                  `xml:"version"`
	BackingConfig           *BackingConfig          `xml:"backingConfig,omitempty"`
}

func init() {
	types.Add("sms:StorageFileSystem", reflect.TypeOf((*StorageFileSystem)(nil)).Elem())
}

type StorageFileSystemInfo struct {
	types.DynamicData

	FileServerName string `xml:"fileServerName"`
	FileSystemPath string `xml:"fileSystemPath"`
	IpAddress      string `xml:"ipAddress,omitempty"`
}

func init() {
	types.Add("sms:StorageFileSystemInfo", reflect.TypeOf((*StorageFileSystemInfo)(nil)).Elem())
}

type StorageLun struct {
	types.DynamicData

	Uuid                   string         `xml:"uuid"`
	VSphereLunIdentifier   string         `xml:"vSphereLunIdentifier"`
	VendorDisplayName      string         `xml:"vendorDisplayName"`
	CapacityInMB           int64          `xml:"capacityInMB"`
	UsedSpaceInMB          int64          `xml:"usedSpaceInMB"`
	LunThinProvisioned     bool           `xml:"lunThinProvisioned"`
	AlternateIdentifier    []string       `xml:"alternateIdentifier,omitempty"`
	DrsManagementPermitted bool           `xml:"drsManagementPermitted"`
	ThinProvisioningStatus string         `xml:"thinProvisioningStatus"`
	BackingConfig          *BackingConfig `xml:"backingConfig,omitempty"`
}

func init() {
	types.Add("sms:StorageLun", reflect.TypeOf((*StorageLun)(nil)).Elem())
}

type StoragePort struct {
	types.DynamicData

	Uuid          string   `xml:"uuid"`
	Type          string   `xml:"type"`
	AlternateName []string `xml:"alternateName,omitempty"`
}

func init() {
	types.Add("sms:StoragePort", reflect.TypeOf((*StoragePort)(nil)).Elem())
}

type StorageProcessor struct {
	types.DynamicData

	Uuid               string   `xml:"uuid"`
	AlternateIdentifer []string `xml:"alternateIdentifer,omitempty"`
}

func init() {
	types.Add("sms:StorageProcessor", reflect.TypeOf((*StorageProcessor)(nil)).Elem())
}

type SupportedVendorModelMapping struct {
	types.DynamicData

	VendorId string `xml:"vendorId,omitempty"`
	ModelId  string `xml:"modelId,omitempty"`
}

func init() {
	types.Add("sms:SupportedVendorModelMapping", reflect.TypeOf((*SupportedVendorModelMapping)(nil)).Elem())
}

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

type SyncOngoing struct {
	SmsReplicationFault

	Task types.ManagedObjectReference `xml:"task"`
}

func init() {
	types.Add("sms:SyncOngoing", reflect.TypeOf((*SyncOngoing)(nil)).Elem())
}

type SyncOngoingFault SyncOngoing

func init() {
	types.Add("sms:SyncOngoingFault", reflect.TypeOf((*SyncOngoingFault)(nil)).Elem())
}

type SyncReplicationGroupRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	GroupId []types.ReplicationGroupId   `xml:"groupId,omitempty"`
	PitName string                       `xml:"pitName"`
}

func init() {
	types.Add("sms:SyncReplicationGroupRequestType", reflect.TypeOf((*SyncReplicationGroupRequestType)(nil)).Elem())
}

type SyncReplicationGroupSuccessResult struct {
	GroupOperationResult

	TimeStamp time.Time             `xml:"timeStamp"`
	PitId     *PointInTimeReplicaId `xml:"pitId,omitempty"`
	PitName   string                `xml:"pitName,omitempty"`
}

func init() {
	types.Add("sms:SyncReplicationGroupSuccessResult", reflect.TypeOf((*SyncReplicationGroupSuccessResult)(nil)).Elem())
}

type SyncReplicationGroup_Task SyncReplicationGroupRequestType

func init() {
	types.Add("sms:SyncReplicationGroup_Task", reflect.TypeOf((*SyncReplicationGroup_Task)(nil)).Elem())
}

type SyncReplicationGroup_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type TargetDeviceId struct {
	types.DynamicData

	DomainId types.FaultDomainId `xml:"domainId"`
	DeviceId ReplicaId           `xml:"deviceId"`
}

func init() {
	types.Add("sms:TargetDeviceId", reflect.TypeOf((*TargetDeviceId)(nil)).Elem())
}

type TargetGroupInfo struct {
	GroupInfo

	SourceInfo       TargetToSourceInfo          `xml:"sourceInfo"`
	State            string                      `xml:"state"`
	Devices          []BaseTargetGroupMemberInfo `xml:"devices,omitempty,typeattr"`
	IsPromoteCapable *bool                       `xml:"isPromoteCapable"`
}

func init() {
	types.Add("sms:TargetGroupInfo", reflect.TypeOf((*TargetGroupInfo)(nil)).Elem())
}

type TargetGroupMemberInfo struct {
	types.DynamicData

	ReplicaId       ReplicaId                    `xml:"replicaId"`
	SourceId        BaseDeviceId                 `xml:"sourceId,typeattr"`
	TargetDatastore types.ManagedObjectReference `xml:"targetDatastore"`
}

func init() {
	types.Add("sms:TargetGroupMemberInfo", reflect.TypeOf((*TargetGroupMemberInfo)(nil)).Elem())
}

type TargetToSourceInfo struct {
	types.DynamicData

	SourceGroupId                   types.ReplicationGroupId `xml:"sourceGroupId"`
	ReplicationAgreementDescription string                   `xml:"replicationAgreementDescription,omitempty"`
}

func init() {
	types.Add("sms:TargetToSourceInfo", reflect.TypeOf((*TargetToSourceInfo)(nil)).Elem())
}

type TestFailoverParam struct {
	FailoverParam
}

func init() {
	types.Add("sms:TestFailoverParam", reflect.TypeOf((*TestFailoverParam)(nil)).Elem())
}

type TestFailoverReplicationGroupStartRequestType struct {
	This              types.ManagedObjectReference `xml:"_this"`
	TestFailoverParam TestFailoverParam            `xml:"testFailoverParam"`
}

func init() {
	types.Add("sms:TestFailoverReplicationGroupStartRequestType", reflect.TypeOf((*TestFailoverReplicationGroupStartRequestType)(nil)).Elem())
}

type TestFailoverReplicationGroupStart_Task TestFailoverReplicationGroupStartRequestType

func init() {
	types.Add("sms:TestFailoverReplicationGroupStart_Task", reflect.TypeOf((*TestFailoverReplicationGroupStart_Task)(nil)).Elem())
}

type TestFailoverReplicationGroupStart_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type TestFailoverReplicationGroupStopRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	GroupId []types.ReplicationGroupId   `xml:"groupId,omitempty"`
	Force   bool                         `xml:"force"`
}

func init() {
	types.Add("sms:TestFailoverReplicationGroupStopRequestType", reflect.TypeOf((*TestFailoverReplicationGroupStopRequestType)(nil)).Elem())
}

type TestFailoverReplicationGroupStop_Task TestFailoverReplicationGroupStopRequestType

func init() {
	types.Add("sms:TestFailoverReplicationGroupStop_Task", reflect.TypeOf((*TestFailoverReplicationGroupStop_Task)(nil)).Elem())
}

type TestFailoverReplicationGroupStop_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type TooMany struct {
	types.MethodFault

	MaxBatchSize int64 `xml:"maxBatchSize,omitempty"`
}

func init() {
	types.Add("sms:TooMany", reflect.TypeOf((*TooMany)(nil)).Elem())
}

type TooManyFault TooMany

func init() {
	types.Add("sms:TooManyFault", reflect.TypeOf((*TooManyFault)(nil)).Elem())
}

type UnregisterProviderRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	ProviderId string                       `xml:"providerId"`
}

func init() {
	types.Add("sms:UnregisterProviderRequestType", reflect.TypeOf((*UnregisterProviderRequestType)(nil)).Elem())
}

type UnregisterProvider_Task UnregisterProviderRequestType

func init() {
	types.Add("sms:UnregisterProvider_Task", reflect.TypeOf((*UnregisterProvider_Task)(nil)).Elem())
}

type UnregisterProvider_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VVolId struct {
	DeviceId

	Id string `xml:"id"`
}

func init() {
	types.Add("sms:VVolId", reflect.TypeOf((*VVolId)(nil)).Elem())
}

type VasaProviderInfo struct {
	SmsProviderInfo

	Url                           string                        `xml:"url"`
	Certificate                   string                        `xml:"certificate,omitempty"`
	Status                        string                        `xml:"status,omitempty"`
	StatusFault                   *types.LocalizedMethodFault   `xml:"statusFault,omitempty"`
	VasaVersion                   string                        `xml:"vasaVersion,omitempty"`
	Namespace                     string                        `xml:"namespace,omitempty"`
	LastSyncTime                  string                        `xml:"lastSyncTime,omitempty"`
	SupportedVendorModelMapping   []SupportedVendorModelMapping `xml:"supportedVendorModelMapping,omitempty"`
	SupportedProfile              []string                      `xml:"supportedProfile,omitempty"`
	SupportedProviderProfile      []string                      `xml:"supportedProviderProfile,omitempty"`
	RelatedStorageArray           []RelatedStorageArray         `xml:"relatedStorageArray,omitempty"`
	ProviderId                    string                        `xml:"providerId,omitempty"`
	CertificateExpiryDate         string                        `xml:"certificateExpiryDate,omitempty"`
	CertificateStatus             string                        `xml:"certificateStatus,omitempty"`
	ServiceLocation               string                        `xml:"serviceLocation,omitempty"`
	NeedsExplicitActivation       *bool                         `xml:"needsExplicitActivation"`
	MaxBatchSize                  int64                         `xml:"maxBatchSize,omitempty"`
	RetainVasaProviderCertificate *bool                         `xml:"retainVasaProviderCertificate"`
	ArrayIndependentProvider      *bool                         `xml:"arrayIndependentProvider"`
	Type                          string                        `xml:"type,omitempty"`
	Category                      string                        `xml:"category,omitempty"`
	Priority                      int32                         `xml:"priority,omitempty"`
	FailoverGroupId               string                        `xml:"failoverGroupId,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderInfo", reflect.TypeOf((*VasaProviderInfo)(nil)).Elem())
}

type VasaProviderReconnectRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:VasaProviderReconnectRequestType", reflect.TypeOf((*VasaProviderReconnectRequestType)(nil)).Elem())
}

type VasaProviderReconnect_Task VasaProviderReconnectRequestType

func init() {
	types.Add("sms:VasaProviderReconnect_Task", reflect.TypeOf((*VasaProviderReconnect_Task)(nil)).Elem())
}

type VasaProviderReconnect_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VasaProviderRefreshCertificateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:VasaProviderRefreshCertificateRequestType", reflect.TypeOf((*VasaProviderRefreshCertificateRequestType)(nil)).Elem())
}

type VasaProviderRefreshCertificate_Task VasaProviderRefreshCertificateRequestType

func init() {
	types.Add("sms:VasaProviderRefreshCertificate_Task", reflect.TypeOf((*VasaProviderRefreshCertificate_Task)(nil)).Elem())
}

type VasaProviderRefreshCertificate_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VasaProviderRevokeCertificateRequestType struct {
	This types.ManagedObjectReference `xml:"_this"`
}

func init() {
	types.Add("sms:VasaProviderRevokeCertificateRequestType", reflect.TypeOf((*VasaProviderRevokeCertificateRequestType)(nil)).Elem())
}

type VasaProviderRevokeCertificate_Task VasaProviderRevokeCertificateRequestType

func init() {
	types.Add("sms:VasaProviderRevokeCertificate_Task", reflect.TypeOf((*VasaProviderRevokeCertificate_Task)(nil)).Elem())
}

type VasaProviderRevokeCertificate_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VasaProviderSpec struct {
	SmsProviderSpec

	Username    string `xml:"username"`
	Password    string `xml:"password"`
	Url         string `xml:"url"`
	Certificate string `xml:"certificate,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderSpec", reflect.TypeOf((*VasaProviderSpec)(nil)).Elem())
}

type VasaProviderSyncRequestType struct {
	This    types.ManagedObjectReference `xml:"_this"`
	ArrayId string                       `xml:"arrayId,omitempty"`
}

func init() {
	types.Add("sms:VasaProviderSyncRequestType", reflect.TypeOf((*VasaProviderSyncRequestType)(nil)).Elem())
}

type VasaProviderSync_Task VasaProviderSyncRequestType

func init() {
	types.Add("sms:VasaProviderSync_Task", reflect.TypeOf((*VasaProviderSync_Task)(nil)).Elem())
}

type VasaProviderSync_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type VasaVirtualDiskId struct {
	DeviceId

	DiskId string `xml:"diskId"`
}

func init() {
	types.Add("sms:VasaVirtualDiskId", reflect.TypeOf((*VasaVirtualDiskId)(nil)).Elem())
}

type VirtualDiskKey struct {
	DeviceId

	VmInstanceUUID string `xml:"vmInstanceUUID"`
	DeviceKey      int32  `xml:"deviceKey"`
}

func init() {
	types.Add("sms:VirtualDiskKey", reflect.TypeOf((*VirtualDiskKey)(nil)).Elem())
}

type VirtualDiskMoId struct {
	DeviceId

	VcUuid  string `xml:"vcUuid,omitempty"`
	VmMoid  string `xml:"vmMoid"`
	DiskKey string `xml:"diskKey"`
}

func init() {
	types.Add("sms:VirtualDiskMoId", reflect.TypeOf((*VirtualDiskMoId)(nil)).Elem())
}

type VirtualMachineFilePath struct {
	VirtualMachineId

	VcUuid  string `xml:"vcUuid,omitempty"`
	DsUrl   string `xml:"dsUrl"`
	VmxPath string `xml:"vmxPath"`
}

func init() {
	types.Add("sms:VirtualMachineFilePath", reflect.TypeOf((*VirtualMachineFilePath)(nil)).Elem())
}

type VirtualMachineId struct {
	DeviceId
}

func init() {
	types.Add("sms:VirtualMachineId", reflect.TypeOf((*VirtualMachineId)(nil)).Elem())
}

type VirtualMachineMoId struct {
	VirtualMachineId

	VcUuid string `xml:"vcUuid,omitempty"`
	VmMoid string `xml:"vmMoid"`
}

func init() {
	types.Add("sms:VirtualMachineMoId", reflect.TypeOf((*VirtualMachineMoId)(nil)).Elem())
}

type VirtualMachineUUID struct {
	VirtualMachineId

	VmInstanceUUID string `xml:"vmInstanceUUID"`
}

func init() {
	types.Add("sms:VirtualMachineUUID", reflect.TypeOf((*VirtualMachineUUID)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("sms:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
