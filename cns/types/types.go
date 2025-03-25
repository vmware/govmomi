// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
	vsanfstypes "github.com/vmware/govmomi/vsan/vsanfs/types"
)

type CnsCreateVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"-"`
	CreateSpecs []CnsVolumeCreateSpec        `xml:"createSpecs,omitempty" json:"createSpecs"`
}

func init() {
	types.Add("CnsCreateVolumeRequestType", reflect.TypeOf((*CnsCreateVolumeRequestType)(nil)).Elem())
}

type CnsCreateVolume CnsCreateVolumeRequestType

func init() {
	types.Add("CnsCreateVolume", reflect.TypeOf((*CnsCreateVolume)(nil)).Elem())
}

type CnsCreateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsEntityMetadata struct {
	types.DynamicData

	EntityName string           `xml:"entityName" json:"entityName"`
	Labels     []types.KeyValue `xml:"labels,omitempty" json:"labels"`
	Delete     bool             `xml:"delete,omitempty" json:"delete"`
	ClusterID  string           `xml:"clusterId,omitempty" json:"clusterID"`
}

func init() {
	types.Add("CnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

type CnsKubernetesEntityReference struct {
	EntityType string `xml:"entityType" json:"entityType"`
	EntityName string `xml:"entityName" json:"entityName"`
	Namespace  string `xml:"namespace,omitempty" json:"namespace"`
	ClusterID  string `xml:"clusterId,omitempty" json:"clusterID"`
}

type CnsKubernetesEntityMetadata struct {
	CnsEntityMetadata

	EntityType     string                         `xml:"entityType" json:"entityType"`
	Namespace      string                         `xml:"namespace,omitempty" json:"namespace"`
	ReferredEntity []CnsKubernetesEntityReference `xml:"referredEntity,omitempty" json:"referredEntity"`
}

func init() {
	types.Add("CnsKubernetesEntityMetadata", reflect.TypeOf((*CnsKubernetesEntityMetadata)(nil)).Elem())
}

type CnsVolumeMetadata struct {
	types.DynamicData

	ContainerCluster      CnsContainerCluster     `xml:"containerCluster" json:"containerCluster"`
	EntityMetadata        []BaseCnsEntityMetadata `xml:"entityMetadata,typeattr,omitempty" json:"entityMetadata"`
	ContainerClusterArray []CnsContainerCluster   `xml:"containerClusterArray,omitempty" json:"containerClusterArray"`
}

func init() {
	types.Add("CnsVolumeMetadata", reflect.TypeOf((*CnsVolumeMetadata)(nil)).Elem())
}

type CnsVolumeCreateSpec struct {
	types.DynamicData
	Name                 string                                `xml:"name" json:"name"`
	VolumeType           string                                `xml:"volumeType" json:"volumeType"`
	Datastores           []types.ManagedObjectReference        `xml:"datastores,omitempty" json:"datastores"`
	Metadata             CnsVolumeMetadata                     `xml:"metadata,omitempty" json:"metadata"`
	BackingObjectDetails BaseCnsBackingObjectDetails           `xml:"backingObjectDetails,typeattr" json:"backingObjectDetails"`
	Profile              []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr" json:"profile"`
	CreateSpec           BaseCnsBaseCreateSpec                 `xml:"createSpec,omitempty,typeattr" json:"createSpec"`
	VolumeSource         BaseCnsVolumeSource                   `xml:"volumeSource,omitempty,typeattr" json:"volumeSource"`
}

func init() {
	types.Add("CnsVolumeCreateSpec", reflect.TypeOf((*CnsVolumeCreateSpec)(nil)).Elem())
}

type CnsUpdateVolumeMetadataRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this" json:"-"`
	UpdateSpecs []CnsVolumeMetadataUpdateSpec `xml:"updateSpecs,omitempty" json:"updateSpecs"`
}

func init() {
	types.Add("CnsUpdateVolumeMetadataRequestType", reflect.TypeOf((*CnsUpdateVolumeMetadataRequestType)(nil)).Elem())
}

type CnsUpdateVolumeMetadata CnsUpdateVolumeMetadataRequestType

func init() {
	types.Add("CnsUpdateVolumeMetadata", reflect.TypeOf((*CnsUpdateVolumeMetadata)(nil)).Elem())
}

type CnsUpdateVolumeMetadataResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumeMetadataUpdateSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId       `xml:"volumeId" json:"volumeId"`
	Metadata CnsVolumeMetadata `xml:"metadata,omitempty" json:"metadata"`
}

func init() {
	types.Add("CnsVolumeMetadataUpdateSpec", reflect.TypeOf((*CnsVolumeMetadataUpdateSpec)(nil)).Elem())
}

type CnsDeleteVolumeRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"-"`
	VolumeIds  []CnsVolumeId                `xml:"volumeIds" json:"volumeIds"`
	DeleteDisk bool                         `xml:"deleteDisk" json:"deleteDisk"`
}

func init() {
	types.Add("CnsDeleteVolumeRequestType", reflect.TypeOf((*CnsDeleteVolumeRequestType)(nil)).Elem())
}

type CnsDeleteVolume CnsDeleteVolumeRequestType

func init() {
	types.Add("CnsDeleteVolume", reflect.TypeOf((*CnsDeleteVolume)(nil)).Elem())
}

type CnsDeleteVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsExtendVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"-"`
	ExtendSpecs []CnsVolumeExtendSpec        `xml:"extendSpecs,omitempty" json:"extendSpecs"`
}

func init() {
	types.Add("CnsExtendVolumeRequestType", reflect.TypeOf((*CnsExtendVolumeRequestType)(nil)).Elem())
}

type CnsExtendVolume CnsExtendVolumeRequestType

func init() {
	types.Add("CnsExtendVolume", reflect.TypeOf((*CnsExtendVolume)(nil)).Elem())
}

type CnsExtendVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumeExtendSpec struct {
	types.DynamicData

	VolumeId     CnsVolumeId `xml:"volumeId" json:"volumeId"`
	CapacityInMb int64       `xml:"capacityInMb" json:"capacityInMb"`
}

func init() {
	types.Add("CnsVolumeExtendSpec", reflect.TypeOf((*CnsVolumeExtendSpec)(nil)).Elem())
}

type CnsAttachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"-"`
	AttachSpecs []CnsVolumeAttachDetachSpec  `xml:"attachSpecs,omitempty" json:"attachSpecs"`
}

func init() {
	types.Add("CnsAttachVolumeRequestType", reflect.TypeOf((*CnsAttachVolumeRequestType)(nil)).Elem())
}

type CnsAttachVolume CnsAttachVolumeRequestType

func init() {
	types.Add("CnsAttachVolume", reflect.TypeOf((*CnsAttachVolume)(nil)).Elem())
}

type CnsAttachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsDetachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"-"`
	DetachSpecs []CnsVolumeAttachDetachSpec  `xml:"detachSpecs,omitempty" json:"detachSpecs"`
}

func init() {
	types.Add("CnsDetachVolumeRequestType", reflect.TypeOf((*CnsDetachVolumeRequestType)(nil)).Elem())
}

type CnsDetachVolume CnsDetachVolumeRequestType

func init() {
	types.Add("CnsDetachVolume", reflect.TypeOf((*CnsDetachVolume)(nil)).Elem())
}

type CnsDetachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumeAttachDetachSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId                  `xml:"volumeId" json:"volumeId"`
	Vm       types.ManagedObjectReference `xml:"vm" json:"vm"`
}

func init() {
	types.Add("CnsVolumeAttachDetachSpec", reflect.TypeOf((*CnsVolumeAttachDetachSpec)(nil)).Elem())
}

type CnsQueryVolume CnsQueryVolumeRequestType

func init() {
	types.Add("CnsQueryVolume", reflect.TypeOf((*CnsQueryVolume)(nil)).Elem())
}

type CnsQueryVolumeRequestType struct {
	This   types.ManagedObjectReference `xml:"_this" json:"-"`
	Filter CnsQueryFilter               `xml:"filter" json:"filter"`
}

func init() {
	types.Add("CnsQueryVolumeRequestType", reflect.TypeOf((*CnsQueryVolumeRequestType)(nil)).Elem())
}

type CnsQueryVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval" json:"returnval"`
}

type CnsQueryVolumeInfo CnsQueryVolumeInfoRequestType

func init() {
	types.Add("CnsQueryVolumeInfo", reflect.TypeOf((*CnsQueryVolumeInfo)(nil)).Elem())
}

type CnsQueryVolumeInfoRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"-"`
	VolumeIds []CnsVolumeId                `xml:"volumes" json:"volumeIds"`
}

type CnsQueryVolumeInfoResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsQueryAllVolume CnsQueryAllVolumeRequestType

func init() {
	types.Add("CnsQueryAllVolume", reflect.TypeOf((*CnsQueryAllVolume)(nil)).Elem())
}

type CnsQueryAllVolumeRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"-"`
	Filter    CnsQueryFilter               `xml:"filter" json:"filter"`
	Selection CnsQuerySelection            `xml:"selection" json:"selection"`
}

func init() {
	types.Add("CnsQueryAllVolumeRequestType", reflect.TypeOf((*CnsQueryVolumeRequestType)(nil)).Elem())
}

type CnsQueryAllVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval" json:"returnval"`
}

type CnsContainerCluster struct {
	types.DynamicData

	ClusterType         string `xml:"clusterType" json:"clusterType"`
	ClusterId           string `xml:"clusterId" json:"clusterId"`
	VSphereUser         string `xml:"vSphereUser" json:"vSphereUser"`
	ClusterFlavor       string `xml:"clusterFlavor,omitempty" json:"clusterFlavor"`
	ClusterDistribution string `xml:"clusterDistribution,omitempty" json:"clusterDistribution"`
	Delete              bool   `xml:"delete,omitempty" json:"delete"`
}

func init() {
	types.Add("CnsContainerCluster", reflect.TypeOf((*CnsContainerCluster)(nil)).Elem())
}

type CnsVolume struct {
	types.DynamicData

	VolumeId                     CnsVolumeId                 `xml:"volumeId" json:"volumeId"`
	DatastoreUrl                 string                      `xml:"datastoreUrl,omitempty" json:"datastoreUrl"`
	Name                         string                      `xml:"name,omitempty" json:"name"`
	VolumeType                   string                      `xml:"volumeType,omitempty" json:"volumeType"`
	StoragePolicyId              string                      `xml:"storagePolicyId,omitempty" json:"storagePolicyId"`
	Metadata                     CnsVolumeMetadata           `xml:"metadata,omitempty" json:"metadata"`
	BackingObjectDetails         BaseCnsBackingObjectDetails `xml:"backingObjectDetails,omitempty" json:"backingObjectDetails"`
	ComplianceStatus             string                      `xml:"complianceStatus,omitempty" json:"complianceStatus"`
	DatastoreAccessibilityStatus string                      `xml:"datastoreAccessibilityStatus,omitempty" json:"datastoreAccessibilityStatus"`
	HealthStatus                 string                      `xml:"healthStatus,omitempty" json:"healthStatus"`
}

func init() {
	types.Add("CnsVolume", reflect.TypeOf((*CnsVolume)(nil)).Elem())
}

type CnsVolumeOperationResult struct {
	types.DynamicData

	VolumeId CnsVolumeId                 `xml:"volumeId,omitempty" json:"volumeId"`
	Fault    *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault"`
}

func init() {
	types.Add("CnsVolumeOperationResult", reflect.TypeOf((*CnsVolumeOperationResult)(nil)).Elem())
}

type CnsVolumeOperationBatchResult struct {
	types.DynamicData

	VolumeResults []BaseCnsVolumeOperationResult `xml:"volumeResults,omitempty,typeattr" json:"volumeResults"`
}

func init() {
	types.Add("CnsVolumeOperationBatchResult", reflect.TypeOf((*CnsVolumeOperationBatchResult)(nil)).Elem())
}

type CnsPlacementResult struct {
	Datastore       types.ManagedObjectReference  `xml:"datastore,omitempty" json:"datastore"`
	PlacementFaults []*types.LocalizedMethodFault `xml:"placementFaults,omitempty" json:"placementFaults"`
}

func init() {
	types.Add("CnsPlacementResult", reflect.TypeOf((*CnsPlacementResult)(nil)).Elem())
}

type CnsVolumeCreateResult struct {
	CnsVolumeOperationResult
	Name             string               `xml:"name,omitempty" json:"name"`
	PlacementResults []CnsPlacementResult `xml:"placementResults,omitempty" json:"placementResults"`
}

func init() {
	types.Add("CnsVolumeCreateResult", reflect.TypeOf((*CnsVolumeCreateResult)(nil)).Elem())
}

type CnsVolumeAttachResult struct {
	CnsVolumeOperationResult

	DiskUUID string `xml:"diskUUID,omitempty" json:"diskUUID"`
}

func init() {
	types.Add("CnsVolumeAttachResult", reflect.TypeOf((*CnsVolumeAttachResult)(nil)).Elem())
}

type CnsVolumeId struct {
	types.DynamicData

	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("CnsVolumeId", reflect.TypeOf((*CnsVolumeId)(nil)).Elem())
}

type CnsBackingObjectDetails struct {
	types.DynamicData

	CapacityInMb int64 `xml:"capacityInMb,omitempty" json:"capacityInMb"`
}

func init() {
	types.Add("CnsBackingObjectDetails", reflect.TypeOf((*CnsBackingObjectDetails)(nil)).Elem())
}

type CnsBlockBackingDetails struct {
	CnsBackingObjectDetails

	BackingDiskId                  string `xml:"backingDiskId,omitempty" json:"backingDiskId"`
	BackingDiskUrlPath             string `xml:"backingDiskUrlPath,omitempty" json:"backingDiskUrlPath"`
	BackingDiskObjectId            string `xml:"backingDiskObjectId,omitempty" json:"backingDiskObjectId"`
	AggregatedSnapshotCapacityInMb int64  `xml:"aggregatedSnapshotCapacityInMb,omitempty" json:"aggregatedSnapshotCapacityInMb"`
	BackingDiskPath                string `xml:"backingDiskPath,omitempty" json:"backingDiskPath"`
}

func init() {
	types.Add("CnsBlockBackingDetails", reflect.TypeOf((*CnsBlockBackingDetails)(nil)).Elem())
}

type CnsFileBackingDetails struct {
	CnsBackingObjectDetails

	BackingFileId string `xml:"backingFileId,omitempty" json:"backingFileId"`
}

func init() {
	types.Add("CnsFileBackingDetails", reflect.TypeOf((*CnsFileBackingDetails)(nil)).Elem())
}

type CnsVsanFileShareBackingDetails struct {
	CnsFileBackingDetails

	Name         string           `xml:"name,omitempty" json:"name"`
	AccessPoints []types.KeyValue `xml:"accessPoints,omitempty" json:"accessPoints"`
}

func init() {
	types.Add("CnsVsanFileShareBackingDetails", reflect.TypeOf((*CnsVsanFileShareBackingDetails)(nil)).Elem())
}

type CnsBaseCreateSpec struct {
	types.DynamicData
}

func init() {
	types.Add("CnsBaseCreateSpec", reflect.TypeOf((*CnsBaseCreateSpec)(nil)).Elem())
}

type CnsFileCreateSpec struct {
	CnsBaseCreateSpec
}

func init() {
	types.Add("CnsFileCreateSpec", reflect.TypeOf((*CnsFileCreateSpec)(nil)).Elem())
}

type CnsVSANFileCreateSpec struct {
	CnsFileCreateSpec
	SoftQuotaInMb int64                                    `xml:"softQuotaInMb,omitempty" json:"softQuotaInMb"`
	Permission    []vsanfstypes.VsanFileShareNetPermission `xml:"permission,omitempty,typeattr" json:"permission"`
}

func init() {
	types.Add("CnsVSANFileCreateSpec", reflect.TypeOf((*CnsVSANFileCreateSpec)(nil)).Elem())
}

type CnsQueryFilter struct {
	types.DynamicData

	VolumeIds                    []CnsVolumeId                  `xml:"volumeIds,omitempty" json:"volumeIds"`
	Names                        []string                       `xml:"names,omitempty" json:"names"`
	ContainerClusterIds          []string                       `xml:"containerClusterIds,omitempty" json:"containerClusterIds"`
	StoragePolicyId              string                         `xml:"storagePolicyId,omitempty" json:"storagePolicyId"`
	Datastores                   []types.ManagedObjectReference `xml:"datastores,omitempty" json:"datastores"`
	Labels                       []types.KeyValue               `xml:"labels,omitempty" json:"labels"`
	ComplianceStatus             string                         `xml:"complianceStatus,omitempty" json:"complianceStatus"`
	DatastoreAccessibilityStatus string                         `xml:"datastoreAccessibilityStatus,omitempty" json:"datastoreAccessibilityStatus"`
	Cursor                       *CnsCursor                     `xml:"cursor,omitempty" json:"cursor"`
	HealthStatus                 string                         `xml:"healthStatus,omitempty" json:"healthStatus"`
}

func init() {
	types.Add("CnsQueryFilter", reflect.TypeOf((*CnsQueryFilter)(nil)).Elem())
}

type CnsQuerySelection struct {
	types.DynamicData

	Names []string `xml:"names,omitempty" json:"names"`
}

type CnsQueryResult struct {
	types.DynamicData

	Volumes []CnsVolume `xml:"volumes,omitempty" json:"volumes"`
	Cursor  CnsCursor   `xml:"cursor" json:"cursor"`
}

func init() {
	types.Add("CnsQueryResult", reflect.TypeOf((*CnsQueryResult)(nil)).Elem())
}

type CnsVolumeInfo struct {
	types.DynamicData
}

func init() {
	types.Add("CnsVolumeInfo", reflect.TypeOf((*CnsVolumeInfo)(nil)).Elem())
}

type CnsBlockVolumeInfo struct {
	CnsVolumeInfo

	VStorageObject types.VStorageObject `xml:"vStorageObject" json:"vStorageObject"`
}

func init() {
	types.Add("CnsBlockVolumeInfo", reflect.TypeOf((*CnsBlockVolumeInfo)(nil)).Elem())
}

type CnsQueryVolumeInfoResult struct {
	CnsVolumeOperationResult

	VolumeInfo BaseCnsVolumeInfo `xml:"volumeInfo,typeattr,omitempty" json:"volumeInfo"`
}

func init() {
	types.Add("CnsQueryVolumeInfoResult", reflect.TypeOf((*CnsQueryVolumeInfoResult)(nil)).Elem())
}

type CnsRelocateVolumeRequestType struct {
	This          types.ManagedObjectReference `xml:"_this" json:"-"`
	RelocateSpecs []BaseCnsVolumeRelocateSpec  `xml:"relocateSpecs,typeattr" json:"relocateSpecs"`
}

func init() {
	types.Add("CnsRelocateVolumeRequestType", reflect.TypeOf((*CnsRelocateVolumeRequestType)(nil)).Elem())
}

type CnsRelocateVolume CnsRelocateVolumeRequestType

func init() {
	types.Add("CnsRelocateVolume", reflect.TypeOf((*CnsRelocateVolume)(nil)).Elem())
}

type CnsRelocateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumeRelocateSpec struct {
	types.DynamicData

	VolumeId       CnsVolumeId                           `xml:"volumeId" json:"volumeId"`
	Datastore      types.ManagedObjectReference          `xml:"datastore" json:"datastore"`
	Profile        []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr" json:"profile"`
	ServiceLocator *types.ServiceLocator                 `xml:"serviceLocator,omitempty" json:"serviceLocator"`
}

func init() {
	types.Add("CnsVolumeRelocateSpec", reflect.TypeOf((*CnsVolumeRelocateSpec)(nil)).Elem())
}

type CnsBlockVolumeRelocateSpec struct {
	CnsVolumeRelocateSpec
}

func NewCnsBlockVolumeRelocateSpec(volumeId string, datastore types.ManagedObjectReference, profile ...types.BaseVirtualMachineProfileSpec) CnsBlockVolumeRelocateSpec {
	cnsVolumeID := CnsVolumeId{
		Id: volumeId,
	}
	volumeSpec := CnsVolumeRelocateSpec{
		VolumeId:  cnsVolumeID,
		Datastore: datastore,
		Profile:   profile,
	}
	blockVolSpec := CnsBlockVolumeRelocateSpec{
		CnsVolumeRelocateSpec: volumeSpec,
	}
	return blockVolSpec
}

func init() {
	types.Add("CnsBlockVolumeRelocateSpec", reflect.TypeOf((*CnsBlockVolumeRelocateSpec)(nil)).Elem())
}

type CnsCursor struct {
	types.DynamicData

	Offset       int64 `xml:"offset" json:"offset"`
	Limit        int64 `xml:"limit" json:"limit"`
	TotalRecords int64 `xml:"totalRecords,omitempty" json:"totalRecords"`
}

func init() {
	types.Add("CnsCursor", reflect.TypeOf((*CnsCursor)(nil)).Elem())
}

type CnsFault struct {
	types.BaseMethodFault `xml:"fault,typeattr"`

	Reason string `xml:"reason,omitempty" json:"reason"`
}

func init() {
	types.Add("CnsFault", reflect.TypeOf((*CnsFault)(nil)).Elem())
}

type CnsVolumeNotFoundFault struct {
	CnsFault

	VolumeId CnsVolumeId `xml:"volumeId" json:"volumeId"`
}

func init() {
	types.Add("CnsVolumeNotFoundFault", reflect.TypeOf((*CnsVolumeNotFoundFault)(nil)).Elem())
}

type CnsAlreadyRegisteredFault struct {
	CnsFault `xml:"fault,typeattr"`

	VolumeId CnsVolumeId `xml:"volumeId,omitempty" json:"volumeId"`
}

func init() {
	types.Add("CnsAlreadyRegisteredFault", reflect.TypeOf((*CnsAlreadyRegisteredFault)(nil)).Elem())
}

type CnsSnapshotNotFoundFault struct {
	CnsFault

	VolumeId   CnsVolumeId   `xml:"volumeId,omitempty" json:"volumeId"`
	SnapshotId CnsSnapshotId `xml:"SnapshotId" json:"snapshotId"`
}

func init() {
	types.Add("CnsSnapshotNotFoundFault", reflect.TypeOf((*CnsSnapshotNotFoundFault)(nil)).Elem())
}

type CnsSnapshotCreatedFault struct {
	CnsFault

	VolumeId   CnsVolumeId                  `xml:"volumeId" json:"volumeId"`
	SnapshotId CnsSnapshotId                `xml:"SnapshotId" json:"snapshotId"`
	Datastore  types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("CnsSnapshotCreatedFault", reflect.TypeOf((*CnsSnapshotCreatedFault)(nil)).Elem())
}

type CnsConfigureVolumeACLs CnsConfigureVolumeACLsRequestType

func init() {
	types.Add("vsan:CnsConfigureVolumeACLs", reflect.TypeOf((*CnsConfigureVolumeACLs)(nil)).Elem())
}

type CnsConfigureVolumeACLsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this" json:"-"`
	ACLConfigSpecs []CnsVolumeACLConfigureSpec  `xml:"ACLConfigSpecs" json:"aclConfigSpecs"`
}

func init() {
	types.Add("vsan:CnsConfigureVolumeACLsRequestType", reflect.TypeOf((*CnsConfigureVolumeACLsRequestType)(nil)).Elem())
}

type CnsConfigureVolumeACLsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumeACLConfigureSpec struct {
	types.DynamicData

	VolumeId              CnsVolumeId               `xml:"volumeId" json:"volumeId"`
	AccessControlSpecList []CnsNFSAccessControlSpec `xml:"accessControlSpecList,typeattr" json:"accessControlSpecList"`
}

type CnsNFSAccessControlSpec struct {
	types.DynamicData
	Permission []vsanfstypes.VsanFileShareNetPermission `xml:"netPermission,omitempty,typeattr" json:"permission"`
	Delete     bool                                     `xml:"delete,omitempty" json:"delete"`
}

func init() {
	types.Add("CnsNFSAccessControlSpec", reflect.TypeOf((*CnsNFSAccessControlSpec)(nil)).Elem())
}

type CnsQueryAsync CnsQueryAsyncRequestType

func init() {
	types.Add("CnsQueryAsync", reflect.TypeOf((*CnsQueryAsync)(nil)).Elem())
}

type CnsQueryAsyncRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"-"`
	Filter    CnsQueryFilter               `xml:"filter" json:"filter"`
	Selection *CnsQuerySelection           `xml:"selection,omitempty" json:"selection"`
}

func init() {
	types.Add("CnsQueryAsyncRequestType", reflect.TypeOf((*CnsQueryAsyncRequestType)(nil)).Elem())
}

type CnsQueryAsyncResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsAsyncQueryResult struct {
	CnsVolumeOperationResult

	QueryResult CnsQueryResult `xml:"queryResult,omitempty" json:"queryResult"`
}

func init() {
	types.Add("CnsAsyncQueryResult", reflect.TypeOf((*CnsAsyncQueryResult)(nil)).Elem())
}

// Cns Snapshot Types

type CnsCreateSnapshotsRequestType struct {
	This          types.ManagedObjectReference `xml:"_this" json:"-"`
	SnapshotSpecs []CnsSnapshotCreateSpec      `xml:"snapshotSpecs,omitempty" json:"snapshotSpecs"`
}

func init() {
	types.Add("CnsCreateSnapshotsRequestType", reflect.TypeOf((*CnsCreateSnapshotsRequestType)(nil)).Elem())
}

type CnsCreateSnapshots CnsCreateSnapshotsRequestType

func init() {
	types.Add("CnsCreateSnapshots", reflect.TypeOf((*CnsCreateSnapshots)(nil)).Elem())
}

type CnsCreateSnapshotsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsSnapshotCreateSpec struct {
	types.DynamicData

	VolumeId    CnsVolumeId `xml:"volumeId" json:"volumeId"`
	Description string      `xml:"description" json:"description"`
}

func init() {
	types.Add("CnsSnapshotCreateSpec", reflect.TypeOf((*CnsSnapshotCreateSpec)(nil)).Elem())
}

type CnsDeleteSnapshotsRequestType struct {
	This                types.ManagedObjectReference `xml:"_this" json:"-"`
	SnapshotDeleteSpecs []CnsSnapshotDeleteSpec      `xml:"snapshotDeleteSpecs,omitempty" json:"snapshotDeleteSpecs"`
}

func init() {
	types.Add("CnsDeleteSnapshotsRequestType", reflect.TypeOf((*CnsDeleteSnapshotsRequestType)(nil)).Elem())
}

type CnsDeleteSnapshots CnsDeleteSnapshotsRequestType

func init() {
	types.Add("CnsDeleteSnapshots", reflect.TypeOf((*CnsDeleteSnapshots)(nil)).Elem())
}

type CnsDeleteSnapshotsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsSnapshotId struct {
	types.DynamicData

	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("CnsSnapshotId", reflect.TypeOf((*CnsSnapshotId)(nil)).Elem())
}

type CnsSnapshotDeleteSpec struct {
	types.DynamicData

	VolumeId   CnsVolumeId   `xml:"volumeId" json:"volumeId"`
	SnapshotId CnsSnapshotId `xml:"snapshotId" json:"snapshotId"`
}

func init() {
	types.Add("CnsSnapshotDeleteSpec", reflect.TypeOf((*CnsSnapshotDeleteSpec)(nil)).Elem())
}

type CnsSnapshot struct {
	types.DynamicData

	SnapshotId  CnsSnapshotId `xml:"snapshotId" json:"snapshotId"`
	VolumeId    CnsVolumeId   `xml:"volumeId" json:"volumeId"`
	Description string        `xml:"description,omitempty" json:"description"`
	CreateTime  time.Time     `xml:"createTime" json:"createTime"`
}

func init() {
	types.Add("CnsSnapshot", reflect.TypeOf((*CnsSnapshot)(nil)).Elem())
}

type CnsSnapshotOperationResult struct {
	CnsVolumeOperationResult
}

func init() {
	types.Add("CnsSnapshotOperationResult", reflect.TypeOf((*CnsSnapshotOperationResult)(nil)).Elem())
}

type CnsSnapshotCreateResult struct {
	CnsSnapshotOperationResult
	Snapshot                       CnsSnapshot `xml:"snapshot,omitempty" json:"snapshot"`
	AggregatedSnapshotCapacityInMb int64       `xml:"aggregatedSnapshotCapacityInMb,omitempty" json:"aggregatedSnapshotCapacityInMb"`
}

func init() {
	types.Add("CnsSnapshotCreateResult", reflect.TypeOf((*CnsSnapshotCreateResult)(nil)).Elem())
}

type CnsSnapshotDeleteResult struct {
	CnsSnapshotOperationResult
	SnapshotId                     CnsSnapshotId `xml:"snapshotId,omitempty" json:"snapshotId"`
	AggregatedSnapshotCapacityInMb int64         `xml:"aggregatedSnapshotCapacityInMb,omitempty" json:"aggregatedSnapshotCapacityInMb"`
}

func init() {
	types.Add("CnsSnapshotDeleteResult", reflect.TypeOf((*CnsSnapshotDeleteResult)(nil)).Elem())
}

type CnsVolumeSource struct {
	types.DynamicData
}

func init() {
	types.Add("CnsVolumeSource", reflect.TypeOf((*CnsVolumeSource)(nil)).Elem())
}

type CnsSnapshotVolumeSource struct {
	CnsVolumeSource

	VolumeId   CnsVolumeId   `xml:"volumeId,omitempty" json:"volumeId"`
	SnapshotId CnsSnapshotId `xml:"snapshotId,omitempty" json:"snapshotId"`
}

func init() {
	types.Add("CnsSnapshotVolumeSource", reflect.TypeOf((*CnsSnapshotVolumeSource)(nil)).Elem())
}

// CNS QuerySnapshots related types

type CnsQuerySnapshotsRequestType struct {
	This                types.ManagedObjectReference `xml:"_this" json:"-"`
	SnapshotQueryFilter CnsSnapshotQueryFilter       `xml:"snapshotQueryFilter" json:"snapshotQueryFilter"`
}

func init() {
	types.Add("CnsQuerySnapshotsRequestType", reflect.TypeOf((*CnsQuerySnapshotsRequestType)(nil)).Elem())
}

type CnsQuerySnapshots CnsQuerySnapshotsRequestType

func init() {
	types.Add("CnsQuerySnapshots", reflect.TypeOf((*CnsQuerySnapshots)(nil)).Elem())
}

type CnsQuerySnapshotsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsSnapshotQueryResult struct {
	types.DynamicData

	Entries []CnsSnapshotQueryResultEntry `xml:"entries,omitempty" json:"entries"`
	Cursor  CnsCursor                     `xml:"cursor" json:"cursor"`
}

func init() {
	types.Add("CnsSnapshotQueryResult", reflect.TypeOf((*CnsSnapshotQueryResult)(nil)).Elem())
}

type CnsSnapshotQueryResultEntry struct {
	types.DynamicData

	Snapshot CnsSnapshot                 `xml:"snapshot,omitempty" json:"snapshot"`
	Error    *types.LocalizedMethodFault `xml:"error,omitempty" json:"error"`
}

func init() {
	types.Add("CnsSnapshotQueryResultEntry", reflect.TypeOf((*CnsSnapshotQueryResultEntry)(nil)).Elem())
}

type CnsSnapshotQueryFilter struct {
	types.DynamicData

	SnapshotQuerySpecs []CnsSnapshotQuerySpec `xml:"snapshotQuerySpecs,omitempty" json:"snapshotQuerySpecs"`
	Cursor             *CnsCursor             `xml:"cursor,omitempty" json:"cursor"`
}

func init() {
	types.Add("CnsSnapshotQueryFilter", reflect.TypeOf((*CnsSnapshotQueryFilter)(nil)).Elem())
}

type CnsSnapshotQuerySpec struct {
	types.DynamicData

	VolumeId   CnsVolumeId    `xml:"volumeId" json:"volumeId"`
	SnapshotId *CnsSnapshotId `xml:"snapshotId,omitempty" json:"snapshotId"`
}

func init() {
	types.Add("CnsSnapshotQuerySpec", reflect.TypeOf((*CnsSnapshotQuerySpec)(nil)).Elem())
}

type CnsReconfigVolumePolicy CnsReconfigVolumePolicyRequestType

func init() {
	types.Add("vsan:CnsReconfigVolumePolicy", reflect.TypeOf((*CnsReconfigVolumePolicy)(nil)).Elem())
}

type CnsReconfigVolumePolicyRequestType struct {
	This                      types.ManagedObjectReference  `xml:"_this" json:"-"`
	VolumePolicyReconfigSpecs []CnsVolumePolicyReconfigSpec `xml:"volumePolicyReconfigSpecs,omitempty" json:"volumePolicyReconfigSpecs"`
}

func init() {
	types.Add("vsan:CnsReconfigVolumePolicyRequestType", reflect.TypeOf((*CnsReconfigVolumePolicyRequestType)(nil)).Elem())
}

type CnsReconfigVolumePolicyResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type CnsVolumePolicyReconfigSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId                           `xml:"volumeId" json:"volumeId"`
	Profile  []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr" json:"profile"`
}

func init() {
	types.Add("vsan:CnsVolumePolicyReconfigSpec", reflect.TypeOf((*CnsVolumePolicyReconfigSpec)(nil)).Elem())
}

type CnsSyncDatastore CnsSyncDatastoreRequestType

func init() {
	types.Add("vsan:CnsSyncDatastore", reflect.TypeOf((*CnsSyncDatastore)(nil)).Elem())
}

type CnsSyncDatastoreRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"-"`
	DatastoreUrl string                       `xml:"datastoreUrl,omitempty" json:"datastoreUrl"`
	FullSync     *bool                        `xml:"fullSync" json:"fullSync"`
}

func init() {
	types.Add("vsan:CnsSyncDatastoreRequestType", reflect.TypeOf((*CnsSyncDatastoreRequestType)(nil)).Elem())
}

type CnsSyncDatastoreResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}
