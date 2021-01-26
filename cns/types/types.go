/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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
	vsanfstypes "github.com/vmware/govmomi/vsan/vsanfs/types"
)

type CnsCreateVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	CreateSpecs []CnsVolumeCreateSpec        `xml:"createSpecs,omitempty"`
}

func init() {
	types.Add("CnsCreateVolumeRequestType", reflect.TypeOf((*CnsCreateVolumeRequestType)(nil)).Elem())
}

type CnsCreateVolume CnsCreateVolumeRequestType

func init() {
	types.Add("CnsCreateVolume", reflect.TypeOf((*CnsCreateVolume)(nil)).Elem())
}

type CnsCreateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsEntityMetadata struct {
	types.DynamicData

	EntityName string           `xml:"entityName"`
	Labels     []types.KeyValue `xml:"labels,omitempty"`
	Delete     bool             `xml:"delete,omitempty"`
	ClusterID  string           `xml:"clusterId,omitempty"`
}

func init() {
	types.Add("CnsEntityMetadata", reflect.TypeOf((*CnsEntityMetadata)(nil)).Elem())
}

type CnsKubernetesEntityReference struct {
	EntityType string `xml:"entityType"`
	EntityName string `xml:"entityName"`
	Namespace  string `xml:"namespace,omitempty"`
	ClusterID  string `xml:"clusterId,omitempty"`
}

type CnsKubernetesEntityMetadata struct {
	CnsEntityMetadata

	EntityType     string                         `xml:"entityType"`
	Namespace      string                         `xml:"namespace,omitempty"`
	ReferredEntity []CnsKubernetesEntityReference `xml:"referredEntity,omitempty"`
}

func init() {
	types.Add("CnsKubernetesEntityMetadata", reflect.TypeOf((*CnsKubernetesEntityMetadata)(nil)).Elem())
}

type CnsVolumeMetadata struct {
	types.DynamicData

	ContainerCluster      CnsContainerCluster     `xml:"containerCluster"`
	EntityMetadata        []BaseCnsEntityMetadata `xml:"entityMetadata,typeattr,omitempty"`
	ContainerClusterArray []CnsContainerCluster   `xml:"containerClusterArray,omitempty"`
}

func init() {
	types.Add("CnsVolumeMetadata", reflect.TypeOf((*CnsVolumeMetadata)(nil)).Elem())
}

type CnsVolumeCreateSpec struct {
	types.DynamicData
	Name                 string                                `xml:"name"`
	VolumeType           string                                `xml:"volumeType"`
	Datastores           []types.ManagedObjectReference        `xml:"datastores,omitempty"`
	Metadata             CnsVolumeMetadata                     `xml:"metadata,omitempty"`
	BackingObjectDetails BaseCnsBackingObjectDetails           `xml:"backingObjectDetails,typeattr"`
	Profile              []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
	CreateSpec           BaseCnsBaseCreateSpec                 `xml:"createSpec,omitempty,typeattr"`
}

func init() {
	types.Add("CnsVolumeCreateSpec", reflect.TypeOf((*CnsVolumeCreateSpec)(nil)).Elem())
}

type CnsUpdateVolumeMetadataRequestType struct {
	This        types.ManagedObjectReference  `xml:"_this"`
	UpdateSpecs []CnsVolumeMetadataUpdateSpec `xml:"updateSpecs,omitempty"`
}

func init() {
	types.Add("CnsUpdateVolumeMetadataRequestType", reflect.TypeOf((*CnsUpdateVolumeMetadataRequestType)(nil)).Elem())
}

type CnsUpdateVolumeMetadata CnsUpdateVolumeMetadataRequestType

func init() {
	types.Add("CnsUpdateVolumeMetadata", reflect.TypeOf((*CnsUpdateVolumeMetadata)(nil)).Elem())
}

type CnsUpdateVolumeMetadataResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVolumeMetadataUpdateSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId       `xml:"volumeId"`
	Metadata CnsVolumeMetadata `xml:"metadata,omitempty"`
}

func init() {
	types.Add("CnsVolumeMetadataUpdateSpec", reflect.TypeOf((*CnsVolumeMetadataUpdateSpec)(nil)).Elem())
}

type CnsDeleteVolumeRequestType struct {
	This       types.ManagedObjectReference `xml:"_this"`
	VolumeIds  []CnsVolumeId                `xml:"volumeIds"`
	DeleteDisk bool                         `xml:"deleteDisk"`
}

func init() {
	types.Add("CnsDeleteVolumeRequestType", reflect.TypeOf((*CnsDeleteVolumeRequestType)(nil)).Elem())
}

type CnsDeleteVolume CnsDeleteVolumeRequestType

func init() {
	types.Add("CnsDeleteVolume", reflect.TypeOf((*CnsDeleteVolume)(nil)).Elem())
}

type CnsDeleteVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsExtendVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	ExtendSpecs []CnsVolumeExtendSpec        `xml:"extendSpecs,omitempty"`
}

func init() {
	types.Add("CnsExtendVolumeRequestType", reflect.TypeOf((*CnsExtendVolumeRequestType)(nil)).Elem())
}

type CnsExtendVolume CnsExtendVolumeRequestType

func init() {
	types.Add("CnsExtendVolume", reflect.TypeOf((*CnsExtendVolume)(nil)).Elem())
}

type CnsExtendVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVolumeExtendSpec struct {
	types.DynamicData

	VolumeId     CnsVolumeId `xml:"volumeId"`
	CapacityInMb int64       `xml:"capacityInMb"`
}

func init() {
	types.Add("CnsVolumeExtendSpec", reflect.TypeOf((*CnsVolumeExtendSpec)(nil)).Elem())
}

type CnsAttachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	AttachSpecs []CnsVolumeAttachDetachSpec  `xml:"attachSpecs,omitempty"`
}

func init() {
	types.Add("CnsAttachVolumeRequestType", reflect.TypeOf((*CnsAttachVolumeRequestType)(nil)).Elem())
}

type CnsAttachVolume CnsAttachVolumeRequestType

func init() {
	types.Add("CnsAttachVolume", reflect.TypeOf((*CnsAttachVolume)(nil)).Elem())
}

type CnsAttachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsDetachVolumeRequestType struct {
	This        types.ManagedObjectReference `xml:"_this"`
	DetachSpecs []CnsVolumeAttachDetachSpec  `xml:"detachSpecs,omitempty"`
}

func init() {
	types.Add("CnsDetachVolumeRequestType", reflect.TypeOf((*CnsDetachVolumeRequestType)(nil)).Elem())
}

type CnsDetachVolume CnsDetachVolumeRequestType

func init() {
	types.Add("CnsDetachVolume", reflect.TypeOf((*CnsDetachVolume)(nil)).Elem())
}

type CnsDetachVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVolumeAttachDetachSpec struct {
	types.DynamicData

	VolumeId CnsVolumeId                  `xml:"volumeId"`
	Vm       types.ManagedObjectReference `xml:"vm"`
}

func init() {
	types.Add("CnsVolumeAttachDetachSpec", reflect.TypeOf((*CnsVolumeAttachDetachSpec)(nil)).Elem())
}

type CnsQueryVolume CnsQueryVolumeRequestType

func init() {
	types.Add("CnsQueryVolume", reflect.TypeOf((*CnsQueryVolume)(nil)).Elem())
}

type CnsQueryVolumeRequestType struct {
	This   types.ManagedObjectReference `xml:"_this"`
	Filter CnsQueryFilter               `xml:"filter"`
}

func init() {
	types.Add("CnsQueryVolumeRequestType", reflect.TypeOf((*CnsQueryVolumeRequestType)(nil)).Elem())
}

type CnsQueryVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval"`
}

type CnsQueryVolumeInfo CnsQueryVolumeInfoRequestType

func init() {
	types.Add("CnsQueryVolumeInfo", reflect.TypeOf((*CnsQueryVolumeInfo)(nil)).Elem())
}

type CnsQueryVolumeInfoRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	VolumeIds []CnsVolumeId                `xml:"volumes"`
}

type CnsQueryVolumeInfoResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsQueryAllVolume CnsQueryAllVolumeRequestType

func init() {
	types.Add("CnsQueryAllVolume", reflect.TypeOf((*CnsQueryAllVolume)(nil)).Elem())
}

type CnsQueryAllVolumeRequestType struct {
	This      types.ManagedObjectReference `xml:"_this"`
	Filter    CnsQueryFilter               `xml:"filter"`
	Selection CnsQuerySelection            `xml:"selection"`
}

func init() {
	types.Add("CnsQueryAllVolumeRequestType", reflect.TypeOf((*CnsQueryVolumeRequestType)(nil)).Elem())
}

type CnsQueryAllVolumeResponse struct {
	Returnval CnsQueryResult `xml:"returnval"`
}

type CnsContainerCluster struct {
	types.DynamicData

	ClusterType         string `xml:"clusterType"`
	ClusterId           string `xml:"clusterId"`
	VSphereUser         string `xml:"vSphereUser"`
	ClusterFlavor       string `xml:"clusterFlavor,omitempty"`
	ClusterDistribution string `xml:"clusterDistribution,omitempty"`
}

func init() {
	types.Add("CnsContainerCluster", reflect.TypeOf((*CnsContainerCluster)(nil)).Elem())
}

type CnsVolume struct {
	types.DynamicData

	VolumeId                     CnsVolumeId                 `xml:"volumeId"`
	DatastoreUrl                 string                      `xml:"datastoreUrl,omitempty"`
	Name                         string                      `xml:"name,omitempty"`
	VolumeType                   string                      `xml:"volumeType,omitempty"`
	StoragePolicyId              string                      `xml:"storagePolicyId,omitempty"`
	Metadata                     CnsVolumeMetadata           `xml:"metadata,omitempty"`
	BackingObjectDetails         BaseCnsBackingObjectDetails `xml:"backingObjectDetails,omitempty"`
	ComplianceStatus             string                      `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                      `xml:"datastoreAccessibilityStatus,omitempty"`
	HealthStatus                 string                      `xml:"healthStatus,omitempty"`
}

func init() {
	types.Add("CnsVolume", reflect.TypeOf((*CnsVolume)(nil)).Elem())
}

type CnsVolumeOperationResult struct {
	types.DynamicData

	VolumeId CnsVolumeId                 `xml:"volumeId,omitempty"`
	Fault    *types.LocalizedMethodFault `xml:"fault,omitempty"`
}

func init() {
	types.Add("CnsVolumeOperationResult", reflect.TypeOf((*CnsVolumeOperationResult)(nil)).Elem())
}

type CnsVolumeOperationBatchResult struct {
	types.DynamicData

	VolumeResults []BaseCnsVolumeOperationResult `xml:"volumeResults,omitempty,typeattr"`
}

func init() {
	types.Add("CnsVolumeOperationBatchResult", reflect.TypeOf((*CnsVolumeOperationBatchResult)(nil)).Elem())
}

type CnsPlacementResult struct {
	Datastore       types.ManagedObjectReference  `xml:"datastore,omitempty"`
	PlacementFaults []*types.LocalizedMethodFault `xml:"placementFaults,omitempty"`
}

func init() {
	types.Add("CnsPlacementResult", reflect.TypeOf((*CnsPlacementResult)(nil)).Elem())
}

type CnsVolumeCreateResult struct {
	CnsVolumeOperationResult
	Name             string               `xml:"name,omitempty"`
	PlacementResults []CnsPlacementResult `xml:"placementResults,omitempty"`
}

func init() {
	types.Add("CnsVolumeCreateResult", reflect.TypeOf((*CnsVolumeCreateResult)(nil)).Elem())
}

type CnsVolumeAttachResult struct {
	CnsVolumeOperationResult

	DiskUUID string `xml:"diskUUID,omitempty"`
}

func init() {
	types.Add("CnsVolumeAttachResult", reflect.TypeOf((*CnsVolumeAttachResult)(nil)).Elem())
}

type CnsVolumeId struct {
	types.DynamicData

	Id string `xml:"id"`
}

func init() {
	types.Add("CnsVolumeId", reflect.TypeOf((*CnsVolumeId)(nil)).Elem())
}

type CnsBackingObjectDetails struct {
	types.DynamicData

	CapacityInMb int64 `xml:"capacityInMb,omitempty"`
}

func init() {
	types.Add("CnsBackingObjectDetails", reflect.TypeOf((*CnsBackingObjectDetails)(nil)).Elem())
}

type CnsBlockBackingDetails struct {
	CnsBackingObjectDetails

	BackingDiskId      string `xml:"backingDiskId,omitempty"`
	BackingDiskUrlPath string `xml:"backingDiskUrlPath,omitempty"`
}

func init() {
	types.Add("CnsBlockBackingDetails", reflect.TypeOf((*CnsBlockBackingDetails)(nil)).Elem())
}

type CnsFileBackingDetails struct {
	CnsBackingObjectDetails

	BackingFileId string `xml:"backingFileId,omitempty"`
}

func init() {
	types.Add("CnsFileBackingDetails", reflect.TypeOf((*CnsFileBackingDetails)(nil)).Elem())
}

type CnsVsanFileShareBackingDetails struct {
	CnsFileBackingDetails

	Name         string           `xml:"name,omitempty"`
	AccessPoints []types.KeyValue `xml:"accessPoints,omitempty"`
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
	SoftQuotaInMb int64                                    `xml:"softQuotaInMb,omitempty"`
	Permission    []vsanfstypes.VsanFileShareNetPermission `xml:"permission,omitempty,typeattr"`
}

func init() {
	types.Add("CnsVSANFileCreateSpec", reflect.TypeOf((*CnsVSANFileCreateSpec)(nil)).Elem())
}

type CnsQueryFilter struct {
	types.DynamicData

	VolumeIds                    []CnsVolumeId                  `xml:"volumeIds,omitempty"`
	Names                        []string                       `xml:"names,omitempty"`
	ContainerClusterIds          []string                       `xml:"containerClusterIds,omitempty"`
	StoragePolicyId              string                         `xml:"storagePolicyId,omitempty"`
	Datastores                   []types.ManagedObjectReference `xml:"datastores,omitempty"`
	Labels                       []types.KeyValue               `xml:"labels,omitempty"`
	ComplianceStatus             string                         `xml:"complianceStatus,omitempty"`
	DatastoreAccessibilityStatus string                         `xml:"datastoreAccessibilityStatus,omitempty"`
	Cursor                       *CnsCursor                     `xml:"cursor,omitempty"`
	healthStatus                 string                         `xml:"healthStatus,omitempty"`
}

func init() {
	types.Add("CnsQueryFilter", reflect.TypeOf((*CnsQueryFilter)(nil)).Elem())
}

type CnsQuerySelection struct {
	types.DynamicData

	Names []string `xml:"names,omitempty"`
}

type CnsQueryResult struct {
	types.DynamicData

	Volumes []CnsVolume `xml:"volumes,omitempty"`
	Cursor  CnsCursor   `xml:"cursor"`
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

	VStorageObject types.VStorageObject `xml:"vStorageObject"`
}

func init() {
	types.Add("CnsBlockVolumeInfo", reflect.TypeOf((*CnsBlockVolumeInfo)(nil)).Elem())
}

type CnsQueryVolumeInfoResult struct {
	CnsVolumeOperationResult

	VolumeInfo BaseCnsVolumeInfo `xml:"volumeInfo,typeattr,omitempty"`
}

func init() {
	types.Add("CnsQueryVolumeInfoResult", reflect.TypeOf((*CnsQueryVolumeInfoResult)(nil)).Elem())
}

type CnsRelocateVolumeRequestType struct {
	This          types.ManagedObjectReference `xml:"_this"`
	RelocateSpecs []BaseCnsVolumeRelocateSpec  `xml:"relocateSpecs,typeattr"`
}

func init() {
	types.Add("CnsRelocateVolumeRequestType", reflect.TypeOf((*CnsRelocateVolumeRequestType)(nil)).Elem())
}

type CnsRelocateVolume CnsRelocateVolumeRequestType

func init() {
	types.Add("CnsRelocateVolume", reflect.TypeOf((*CnsRelocateVolume)(nil)).Elem())
}

type CnsRelocateVolumeResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVolumeRelocateSpec struct {
	types.DynamicData

	VolumeId  CnsVolumeId                           `xml:"volumeId"`
	Datastore types.ManagedObjectReference          `xml:"datastore"`
	Profile   []types.BaseVirtualMachineProfileSpec `xml:"profile,omitempty,typeattr"`
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

	Offset       int64 `xml:"offset"`
	Limit        int64 `xml:"limit"`
	TotalRecords int64 `xml:"totalRecords,omitempty"`
}

func init() {
	types.Add("CnsCursor", reflect.TypeOf((*CnsCursor)(nil)).Elem())
}

type CnsFault struct {
	types.BaseMethodFault `xml:"fault,typeattr"`

	Reason string `xml:"reason,omitempty"`
}

func init() {
	types.Add("CnsFault", reflect.TypeOf((*CnsFault)(nil)).Elem())
}

type CnsAlreadyRegisteredFault struct {
	CnsFault `xml:"fault,typeattr"`

	VolumeId CnsVolumeId `xml:"volumeId,omitempty"`
}

func init() {
	types.Add("CnsAlreadyRegisteredFault", reflect.TypeOf((*CnsAlreadyRegisteredFault)(nil)).Elem())
}

type CnsConfigureVolumeACLs CnsConfigureVolumeACLsRequestType

func init() {
	types.Add("vsan:CnsConfigureVolumeACLs", reflect.TypeOf((*CnsConfigureVolumeACLs)(nil)).Elem())
}

type CnsConfigureVolumeACLsRequestType struct {
	This           types.ManagedObjectReference `xml:"_this"`
	ACLConfigSpecs []CnsVolumeACLConfigureSpec  `xml:"ACLConfigSpecs"`
}

func init() {
	types.Add("vsan:CnsConfigureVolumeACLsRequestType", reflect.TypeOf((*CnsConfigureVolumeACLsRequestType)(nil)).Elem())
}

type CnsConfigureVolumeACLsResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval"`
}

type CnsVolumeACLConfigureSpec struct {
	types.DynamicData

	VolumeId              CnsVolumeId               `xml:"volumeId"`
	AccessControlSpecList []CnsNFSAccessControlSpec `xml:"accessControlSpecList,typeattr"`
}

type CnsNFSAccessControlSpec struct {
	types.DynamicData
	Permission []vsanfstypes.VsanFileShareNetPermission `xml:"netPermission,omitempty,typeattr"`
	Delete     bool                                     `xml:"delete,omitempty"`
}

func init() {
	types.Add("CnsNFSAccessControlSpec", reflect.TypeOf((*CnsNFSAccessControlSpec)(nil)).Elem())
}
