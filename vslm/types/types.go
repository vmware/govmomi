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
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type ArrayOfVslmDatastoreSyncStatus struct {
	VslmDatastoreSyncStatus []VslmDatastoreSyncStatus `xml:"VslmDatastoreSyncStatus,omitempty" json:"VslmDatastoreSyncStatus,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmDatastoreSyncStatus", reflect.TypeOf((*ArrayOfVslmDatastoreSyncStatus)(nil)).Elem())
}

type ArrayOfVslmQueryDatastoreInfoResult struct {
	VslmQueryDatastoreInfoResult []VslmQueryDatastoreInfoResult `xml:"VslmQueryDatastoreInfoResult,omitempty" json:"VslmQueryDatastoreInfoResult,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmQueryDatastoreInfoResult", reflect.TypeOf((*ArrayOfVslmQueryDatastoreInfoResult)(nil)).Elem())
}

type ArrayOfVslmVsoVStorageObjectAssociations struct {
	VslmVsoVStorageObjectAssociations []VslmVsoVStorageObjectAssociations `xml:"VslmVsoVStorageObjectAssociations,omitempty" json:"VslmVsoVStorageObjectAssociations,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectAssociations", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectAssociations)(nil)).Elem())
}

type ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation struct {
	VslmVsoVStorageObjectAssociationsVmDiskAssociation []VslmVsoVStorageObjectAssociationsVmDiskAssociation `xml:"VslmVsoVStorageObjectAssociationsVmDiskAssociation,omitempty" json:"VslmVsoVStorageObjectAssociationsVmDiskAssociation,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation)(nil)).Elem())
}

type ArrayOfVslmVsoVStorageObjectQuerySpec struct {
	VslmVsoVStorageObjectQuerySpec []VslmVsoVStorageObjectQuerySpec `xml:"VslmVsoVStorageObjectQuerySpec,omitempty" json:"VslmVsoVStorageObjectQuerySpec,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectQuerySpec", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectQuerySpec)(nil)).Elem())
}

type ArrayOfVslmVsoVStorageObjectResult struct {
	VslmVsoVStorageObjectResult []VslmVsoVStorageObjectResult `xml:"VslmVsoVStorageObjectResult,omitempty" json:"VslmVsoVStorageObjectResult,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectResult", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectResult)(nil)).Elem())
}

type ArrayOfVslmVsoVStorageObjectSnapshotResult struct {
	VslmVsoVStorageObjectSnapshotResult []VslmVsoVStorageObjectSnapshotResult `xml:"VslmVsoVStorageObjectSnapshotResult,omitempty" json:"VslmVsoVStorageObjectSnapshotResult,omitempty"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectSnapshotResult", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectSnapshotResult)(nil)).Elem())
}

type RetrieveContent RetrieveContentRequestType

func init() {
	types.Add("vslm:RetrieveContent", reflect.TypeOf((*RetrieveContent)(nil)).Elem())
}

type RetrieveContentRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:RetrieveContentRequestType", reflect.TypeOf((*RetrieveContentRequestType)(nil)).Elem())
}

type RetrieveContentResponse struct {
	Returnval VslmServiceInstanceContent `xml:"returnval" json:"returnval"`
}

type VslmAboutInfo struct {
	types.DynamicData

	Name         string `xml:"name" json:"name"`
	FullName     string `xml:"fullName" json:"fullName"`
	Vendor       string `xml:"vendor" json:"vendor"`
	ApiVersion   string `xml:"apiVersion" json:"apiVersion"`
	InstanceUuid string `xml:"instanceUuid" json:"instanceUuid"`
}

func init() {
	types.Add("vslm:VslmAboutInfo", reflect.TypeOf((*VslmAboutInfo)(nil)).Elem())
}

type VslmAttachDiskRequestType struct {
	This          types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id            types.ID                     `xml:"id" json:"id"`
	Vm            types.ManagedObjectReference `xml:"vm" json:"vm"`
	ControllerKey int32                        `xml:"controllerKey,omitempty" json:"controllerKey,omitempty"`
	UnitNumber    *int32                       `xml:"unitNumber" json:"unitNumber,omitempty"`
}

func init() {
	types.Add("vslm:VslmAttachDiskRequestType", reflect.TypeOf((*VslmAttachDiskRequestType)(nil)).Elem())
}

type VslmAttachDisk_Task VslmAttachDiskRequestType

func init() {
	types.Add("vslm:VslmAttachDisk_Task", reflect.TypeOf((*VslmAttachDisk_Task)(nil)).Elem())
}

type VslmAttachDisk_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmAttachTagToVStorageObject VslmAttachTagToVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmAttachTagToVStorageObject", reflect.TypeOf((*VslmAttachTagToVStorageObject)(nil)).Elem())
}

type VslmAttachTagToVStorageObjectRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id       types.ID                     `xml:"id" json:"id"`
	Category string                       `xml:"category" json:"category"`
	Tag      string                       `xml:"tag" json:"tag"`
}

func init() {
	types.Add("vslm:VslmAttachTagToVStorageObjectRequestType", reflect.TypeOf((*VslmAttachTagToVStorageObjectRequestType)(nil)).Elem())
}

type VslmAttachTagToVStorageObjectResponse struct {
}

type VslmCancelTask VslmCancelTaskRequestType

func init() {
	types.Add("vslm:VslmCancelTask", reflect.TypeOf((*VslmCancelTask)(nil)).Elem())
}

type VslmCancelTaskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:VslmCancelTaskRequestType", reflect.TypeOf((*VslmCancelTaskRequestType)(nil)).Elem())
}

type VslmCancelTaskResponse struct {
}

type VslmClearVStorageObjectControlFlags VslmClearVStorageObjectControlFlagsRequestType

func init() {
	types.Add("vslm:VslmClearVStorageObjectControlFlags", reflect.TypeOf((*VslmClearVStorageObjectControlFlags)(nil)).Elem())
}

type VslmClearVStorageObjectControlFlagsRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id           types.ID                     `xml:"id" json:"id"`
	ControlFlags []string                     `xml:"controlFlags,omitempty" json:"controlFlags,omitempty"`
}

func init() {
	types.Add("vslm:VslmClearVStorageObjectControlFlagsRequestType", reflect.TypeOf((*VslmClearVStorageObjectControlFlagsRequestType)(nil)).Elem())
}

type VslmClearVStorageObjectControlFlagsResponse struct {
}

type VslmCloneVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
	Spec types.VslmCloneSpec          `xml:"spec" json:"spec"`
}

func init() {
	types.Add("vslm:VslmCloneVStorageObjectRequestType", reflect.TypeOf((*VslmCloneVStorageObjectRequestType)(nil)).Elem())
}

type VslmCloneVStorageObject_Task VslmCloneVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmCloneVStorageObject_Task", reflect.TypeOf((*VslmCloneVStorageObject_Task)(nil)).Elem())
}

type VslmCloneVStorageObject_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmCreateDiskFromSnapshotRequestType struct {
	This       types.ManagedObjectReference      `xml:"_this" json:"_this"`
	Id         types.ID                          `xml:"id" json:"id"`
	SnapshotId types.ID                          `xml:"snapshotId" json:"snapshotId"`
	Name       string                            `xml:"name" json:"name"`
	Profile    []types.VirtualMachineProfileSpec `xml:"profile,omitempty" json:"profile,omitempty"`
	Crypto     *types.CryptoSpec                 `xml:"crypto,omitempty" json:"crypto,omitempty"`
	Path       string                            `xml:"path,omitempty" json:"path,omitempty"`
}

func init() {
	types.Add("vslm:VslmCreateDiskFromSnapshotRequestType", reflect.TypeOf((*VslmCreateDiskFromSnapshotRequestType)(nil)).Elem())
}

type VslmCreateDiskFromSnapshot_Task VslmCreateDiskFromSnapshotRequestType

func init() {
	types.Add("vslm:VslmCreateDiskFromSnapshot_Task", reflect.TypeOf((*VslmCreateDiskFromSnapshot_Task)(nil)).Elem())
}

type VslmCreateDiskFromSnapshot_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmCreateDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Spec types.VslmCreateSpec         `xml:"spec" json:"spec"`
}

func init() {
	types.Add("vslm:VslmCreateDiskRequestType", reflect.TypeOf((*VslmCreateDiskRequestType)(nil)).Elem())
}

type VslmCreateDisk_Task VslmCreateDiskRequestType

func init() {
	types.Add("vslm:VslmCreateDisk_Task", reflect.TypeOf((*VslmCreateDisk_Task)(nil)).Elem())
}

type VslmCreateDisk_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmCreateSnapshotRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id          types.ID                     `xml:"id" json:"id"`
	Description string                       `xml:"description" json:"description"`
}

func init() {
	types.Add("vslm:VslmCreateSnapshotRequestType", reflect.TypeOf((*VslmCreateSnapshotRequestType)(nil)).Elem())
}

type VslmCreateSnapshot_Task VslmCreateSnapshotRequestType

func init() {
	types.Add("vslm:VslmCreateSnapshot_Task", reflect.TypeOf((*VslmCreateSnapshot_Task)(nil)).Elem())
}

type VslmCreateSnapshot_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmDatastoreSyncStatus struct {
	types.DynamicData

	DatastoreURL    string                      `xml:"datastoreURL" json:"datastoreURL"`
	ObjectVClock    int64                       `xml:"objectVClock" json:"objectVClock"`
	SyncVClock      int64                       `xml:"syncVClock" json:"syncVClock"`
	SyncTime        *time.Time                  `xml:"syncTime" json:"syncTime,omitempty"`
	NumberOfRetries int32                       `xml:"numberOfRetries,omitempty" json:"numberOfRetries,omitempty"`
	Error           *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("vslm:VslmDatastoreSyncStatus", reflect.TypeOf((*VslmDatastoreSyncStatus)(nil)).Elem())
}

type VslmDeleteSnapshotRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	SnapshotId types.ID                     `xml:"snapshotId" json:"snapshotId"`
}

func init() {
	types.Add("vslm:VslmDeleteSnapshotRequestType", reflect.TypeOf((*VslmDeleteSnapshotRequestType)(nil)).Elem())
}

type VslmDeleteSnapshot_Task VslmDeleteSnapshotRequestType

func init() {
	types.Add("vslm:VslmDeleteSnapshot_Task", reflect.TypeOf((*VslmDeleteSnapshot_Task)(nil)).Elem())
}

type VslmDeleteSnapshot_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmDeleteVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmDeleteVStorageObjectRequestType", reflect.TypeOf((*VslmDeleteVStorageObjectRequestType)(nil)).Elem())
}

type VslmDeleteVStorageObject_Task VslmDeleteVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmDeleteVStorageObject_Task", reflect.TypeOf((*VslmDeleteVStorageObject_Task)(nil)).Elem())
}

type VslmDeleteVStorageObject_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmDetachTagFromVStorageObject VslmDetachTagFromVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmDetachTagFromVStorageObject", reflect.TypeOf((*VslmDetachTagFromVStorageObject)(nil)).Elem())
}

type VslmDetachTagFromVStorageObjectRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id       types.ID                     `xml:"id" json:"id"`
	Category string                       `xml:"category" json:"category"`
	Tag      string                       `xml:"tag" json:"tag"`
}

func init() {
	types.Add("vslm:VslmDetachTagFromVStorageObjectRequestType", reflect.TypeOf((*VslmDetachTagFromVStorageObjectRequestType)(nil)).Elem())
}

type VslmDetachTagFromVStorageObjectResponse struct {
}

type VslmExtendDiskRequestType struct {
	This            types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id              types.ID                     `xml:"id" json:"id"`
	NewCapacityInMB int64                        `xml:"newCapacityInMB" json:"newCapacityInMB"`
}

func init() {
	types.Add("vslm:VslmExtendDiskRequestType", reflect.TypeOf((*VslmExtendDiskRequestType)(nil)).Elem())
}

type VslmExtendDisk_Task VslmExtendDiskRequestType

func init() {
	types.Add("vslm:VslmExtendDisk_Task", reflect.TypeOf((*VslmExtendDisk_Task)(nil)).Elem())
}

type VslmExtendDisk_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmFault struct {
	types.MethodFault

	Msg string `xml:"msg,omitempty" json:"msg,omitempty"`
}

func init() {
	types.Add("vslm:VslmFault", reflect.TypeOf((*VslmFault)(nil)).Elem())
}

type VslmFaultFault BaseVslmFault

func init() {
	types.Add("vslm:VslmFaultFault", reflect.TypeOf((*VslmFaultFault)(nil)).Elem())
}

type VslmInflateDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmInflateDiskRequestType", reflect.TypeOf((*VslmInflateDiskRequestType)(nil)).Elem())
}

type VslmInflateDisk_Task VslmInflateDiskRequestType

func init() {
	types.Add("vslm:VslmInflateDisk_Task", reflect.TypeOf((*VslmInflateDisk_Task)(nil)).Elem())
}

type VslmInflateDisk_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmListTagsAttachedToVStorageObject VslmListTagsAttachedToVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmListTagsAttachedToVStorageObject", reflect.TypeOf((*VslmListTagsAttachedToVStorageObject)(nil)).Elem())
}

type VslmListTagsAttachedToVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmListTagsAttachedToVStorageObjectRequestType", reflect.TypeOf((*VslmListTagsAttachedToVStorageObjectRequestType)(nil)).Elem())
}

type VslmListTagsAttachedToVStorageObjectResponse struct {
	Returnval []types.VslmTagEntry `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmListVStorageObjectForSpec VslmListVStorageObjectForSpecRequestType

func init() {
	types.Add("vslm:VslmListVStorageObjectForSpec", reflect.TypeOf((*VslmListVStorageObjectForSpec)(nil)).Elem())
}

type VslmListVStorageObjectForSpecRequestType struct {
	This      types.ManagedObjectReference     `xml:"_this" json:"_this"`
	Query     []VslmVsoVStorageObjectQuerySpec `xml:"query,omitempty" json:"query,omitempty"`
	MaxResult int32                            `xml:"maxResult" json:"maxResult"`
}

func init() {
	types.Add("vslm:VslmListVStorageObjectForSpecRequestType", reflect.TypeOf((*VslmListVStorageObjectForSpecRequestType)(nil)).Elem())
}

type VslmListVStorageObjectForSpecResponse struct {
	Returnval *VslmVsoVStorageObjectQueryResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmListVStorageObjectsAttachedToTag VslmListVStorageObjectsAttachedToTagRequestType

func init() {
	types.Add("vslm:VslmListVStorageObjectsAttachedToTag", reflect.TypeOf((*VslmListVStorageObjectsAttachedToTag)(nil)).Elem())
}

type VslmListVStorageObjectsAttachedToTagRequestType struct {
	This     types.ManagedObjectReference `xml:"_this" json:"_this"`
	Category string                       `xml:"category" json:"category"`
	Tag      string                       `xml:"tag" json:"tag"`
}

func init() {
	types.Add("vslm:VslmListVStorageObjectsAttachedToTagRequestType", reflect.TypeOf((*VslmListVStorageObjectsAttachedToTagRequestType)(nil)).Elem())
}

type VslmListVStorageObjectsAttachedToTagResponse struct {
	Returnval []types.ID `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmLoginByToken VslmLoginByTokenRequestType

func init() {
	types.Add("vslm:VslmLoginByToken", reflect.TypeOf((*VslmLoginByToken)(nil)).Elem())
}

type VslmLoginByTokenRequestType struct {
	This              types.ManagedObjectReference `xml:"_this" json:"_this"`
	DelegatedTokenXml string                       `xml:"delegatedTokenXml" json:"delegatedTokenXml"`
}

func init() {
	types.Add("vslm:VslmLoginByTokenRequestType", reflect.TypeOf((*VslmLoginByTokenRequestType)(nil)).Elem())
}

type VslmLoginByTokenResponse struct {
}

type VslmLogout VslmLogoutRequestType

func init() {
	types.Add("vslm:VslmLogout", reflect.TypeOf((*VslmLogout)(nil)).Elem())
}

type VslmLogoutRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:VslmLogoutRequestType", reflect.TypeOf((*VslmLogoutRequestType)(nil)).Elem())
}

type VslmLogoutResponse struct {
}

type VslmQueryChangedDiskAreas VslmQueryChangedDiskAreasRequestType

func init() {
	types.Add("vslm:VslmQueryChangedDiskAreas", reflect.TypeOf((*VslmQueryChangedDiskAreas)(nil)).Elem())
}

type VslmQueryChangedDiskAreasRequestType struct {
	This        types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id          types.ID                     `xml:"id" json:"id"`
	SnapshotId  types.ID                     `xml:"snapshotId" json:"snapshotId"`
	StartOffset int64                        `xml:"startOffset" json:"startOffset"`
	ChangeId    string                       `xml:"changeId" json:"changeId"`
}

func init() {
	types.Add("vslm:VslmQueryChangedDiskAreasRequestType", reflect.TypeOf((*VslmQueryChangedDiskAreasRequestType)(nil)).Elem())
}

type VslmQueryChangedDiskAreasResponse struct {
	Returnval types.DiskChangeInfo `xml:"returnval" json:"returnval"`
}

type VslmQueryDatastoreInfo VslmQueryDatastoreInfoRequestType

func init() {
	types.Add("vslm:VslmQueryDatastoreInfo", reflect.TypeOf((*VslmQueryDatastoreInfo)(nil)).Elem())
}

type VslmQueryDatastoreInfoRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"_this"`
	DatastoreUrl string                       `xml:"datastoreUrl" json:"datastoreUrl"`
}

func init() {
	types.Add("vslm:VslmQueryDatastoreInfoRequestType", reflect.TypeOf((*VslmQueryDatastoreInfoRequestType)(nil)).Elem())
}

type VslmQueryDatastoreInfoResponse struct {
	Returnval []VslmQueryDatastoreInfoResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmQueryDatastoreInfoResult struct {
	types.DynamicData

	Datacenter types.ManagedObjectReference `xml:"datacenter" json:"datacenter"`
	Datastore  types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("vslm:VslmQueryDatastoreInfoResult", reflect.TypeOf((*VslmQueryDatastoreInfoResult)(nil)).Elem())
}

type VslmQueryGlobalCatalogSyncStatus VslmQueryGlobalCatalogSyncStatusRequestType

func init() {
	types.Add("vslm:VslmQueryGlobalCatalogSyncStatus", reflect.TypeOf((*VslmQueryGlobalCatalogSyncStatus)(nil)).Elem())
}

type VslmQueryGlobalCatalogSyncStatusForDatastore VslmQueryGlobalCatalogSyncStatusForDatastoreRequestType

func init() {
	types.Add("vslm:VslmQueryGlobalCatalogSyncStatusForDatastore", reflect.TypeOf((*VslmQueryGlobalCatalogSyncStatusForDatastore)(nil)).Elem())
}

type VslmQueryGlobalCatalogSyncStatusForDatastoreRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"_this"`
	DatastoreURL string                       `xml:"datastoreURL" json:"datastoreURL"`
}

func init() {
	types.Add("vslm:VslmQueryGlobalCatalogSyncStatusForDatastoreRequestType", reflect.TypeOf((*VslmQueryGlobalCatalogSyncStatusForDatastoreRequestType)(nil)).Elem())
}

type VslmQueryGlobalCatalogSyncStatusForDatastoreResponse struct {
	Returnval *VslmDatastoreSyncStatus `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmQueryGlobalCatalogSyncStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:VslmQueryGlobalCatalogSyncStatusRequestType", reflect.TypeOf((*VslmQueryGlobalCatalogSyncStatusRequestType)(nil)).Elem())
}

type VslmQueryGlobalCatalogSyncStatusResponse struct {
	Returnval []VslmDatastoreSyncStatus `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmQueryInfo VslmQueryInfoRequestType

func init() {
	types.Add("vslm:VslmQueryInfo", reflect.TypeOf((*VslmQueryInfo)(nil)).Elem())
}

type VslmQueryInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:VslmQueryInfoRequestType", reflect.TypeOf((*VslmQueryInfoRequestType)(nil)).Elem())
}

type VslmQueryInfoResponse struct {
	Returnval VslmTaskInfo `xml:"returnval" json:"returnval"`
}

type VslmQueryTaskResult VslmQueryTaskResultRequestType

func init() {
	types.Add("vslm:VslmQueryTaskResult", reflect.TypeOf((*VslmQueryTaskResult)(nil)).Elem())
}

type VslmQueryTaskResultRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("vslm:VslmQueryTaskResultRequestType", reflect.TypeOf((*VslmQueryTaskResultRequestType)(nil)).Elem())
}

type VslmQueryTaskResultResponse struct {
	Returnval types.AnyType `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

type VslmReconcileDatastoreInventoryRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"_this"`
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("vslm:VslmReconcileDatastoreInventoryRequestType", reflect.TypeOf((*VslmReconcileDatastoreInventoryRequestType)(nil)).Elem())
}

type VslmReconcileDatastoreInventory_Task VslmReconcileDatastoreInventoryRequestType

func init() {
	types.Add("vslm:VslmReconcileDatastoreInventory_Task", reflect.TypeOf((*VslmReconcileDatastoreInventory_Task)(nil)).Elem())
}

type VslmReconcileDatastoreInventory_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmRegisterDisk VslmRegisterDiskRequestType

func init() {
	types.Add("vslm:VslmRegisterDisk", reflect.TypeOf((*VslmRegisterDisk)(nil)).Elem())
}

type VslmRegisterDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Path string                       `xml:"path" json:"path"`
	Name string                       `xml:"name,omitempty" json:"name,omitempty"`
}

func init() {
	types.Add("vslm:VslmRegisterDiskRequestType", reflect.TypeOf((*VslmRegisterDiskRequestType)(nil)).Elem())
}

type VslmRegisterDiskResponse struct {
	Returnval types.VStorageObject `xml:"returnval" json:"returnval"`
}

type VslmRelocateVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
	Spec types.VslmRelocateSpec       `xml:"spec" json:"spec"`
}

func init() {
	types.Add("vslm:VslmRelocateVStorageObjectRequestType", reflect.TypeOf((*VslmRelocateVStorageObjectRequestType)(nil)).Elem())
}

type VslmRelocateVStorageObject_Task VslmRelocateVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmRelocateVStorageObject_Task", reflect.TypeOf((*VslmRelocateVStorageObject_Task)(nil)).Elem())
}

type VslmRelocateVStorageObject_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmRenameVStorageObject VslmRenameVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmRenameVStorageObject", reflect.TypeOf((*VslmRenameVStorageObject)(nil)).Elem())
}

type VslmRenameVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
	Name string                       `xml:"name" json:"name"`
}

func init() {
	types.Add("vslm:VslmRenameVStorageObjectRequestType", reflect.TypeOf((*VslmRenameVStorageObjectRequestType)(nil)).Elem())
}

type VslmRenameVStorageObjectResponse struct {
}

type VslmRetrieveSnapshotDetails VslmRetrieveSnapshotDetailsRequestType

func init() {
	types.Add("vslm:VslmRetrieveSnapshotDetails", reflect.TypeOf((*VslmRetrieveSnapshotDetails)(nil)).Elem())
}

type VslmRetrieveSnapshotDetailsRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	SnapshotId types.ID                     `xml:"snapshotId" json:"snapshotId"`
}

func init() {
	types.Add("vslm:VslmRetrieveSnapshotDetailsRequestType", reflect.TypeOf((*VslmRetrieveSnapshotDetailsRequestType)(nil)).Elem())
}

type VslmRetrieveSnapshotDetailsResponse struct {
	Returnval types.VStorageObjectSnapshotDetails `xml:"returnval" json:"returnval"`
}

type VslmRetrieveSnapshotInfo VslmRetrieveSnapshotInfoRequestType

func init() {
	types.Add("vslm:VslmRetrieveSnapshotInfo", reflect.TypeOf((*VslmRetrieveSnapshotInfo)(nil)).Elem())
}

type VslmRetrieveSnapshotInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmRetrieveSnapshotInfoRequestType", reflect.TypeOf((*VslmRetrieveSnapshotInfoRequestType)(nil)).Elem())
}

type VslmRetrieveSnapshotInfoResponse struct {
	Returnval types.VStorageObjectSnapshotInfo `xml:"returnval" json:"returnval"`
}

type VslmRetrieveVStorageInfrastructureObjectPolicy VslmRetrieveVStorageInfrastructureObjectPolicyRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageInfrastructureObjectPolicy", reflect.TypeOf((*VslmRetrieveVStorageInfrastructureObjectPolicy)(nil)).Elem())
}

type VslmRetrieveVStorageInfrastructureObjectPolicyRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"_this"`
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageInfrastructureObjectPolicyRequestType", reflect.TypeOf((*VslmRetrieveVStorageInfrastructureObjectPolicyRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageInfrastructureObjectPolicyResponse struct {
	Returnval []types.VslmInfrastructureObjectPolicy `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmRetrieveVStorageObject VslmRetrieveVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObject", reflect.TypeOf((*VslmRetrieveVStorageObject)(nil)).Elem())
}

type VslmRetrieveVStorageObjectAssociations VslmRetrieveVStorageObjectAssociationsRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectAssociations", reflect.TypeOf((*VslmRetrieveVStorageObjectAssociations)(nil)).Elem())
}

type VslmRetrieveVStorageObjectAssociationsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Ids  []types.ID                   `xml:"ids,omitempty" json:"ids,omitempty"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectAssociationsRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectAssociationsRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectAssociationsResponse struct {
	Returnval []VslmVsoVStorageObjectAssociations `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmRetrieveVStorageObjectMetadata VslmRetrieveVStorageObjectMetadataRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectMetadata", reflect.TypeOf((*VslmRetrieveVStorageObjectMetadata)(nil)).Elem())
}

type VslmRetrieveVStorageObjectMetadataRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	SnapshotId *types.ID                    `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	Prefix     string                       `xml:"prefix,omitempty" json:"prefix,omitempty"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectMetadataRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectMetadataRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectMetadataResponse struct {
	Returnval []types.KeyValue `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmRetrieveVStorageObjectMetadataValue VslmRetrieveVStorageObjectMetadataValueRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectMetadataValue", reflect.TypeOf((*VslmRetrieveVStorageObjectMetadataValue)(nil)).Elem())
}

type VslmRetrieveVStorageObjectMetadataValueRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	SnapshotId *types.ID                    `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	Key        string                       `xml:"key" json:"key"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectMetadataValueRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectMetadataValueRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectMetadataValueResponse struct {
	Returnval string `xml:"returnval" json:"returnval"`
}

type VslmRetrieveVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectResponse struct {
	Returnval types.VStorageObject `xml:"returnval" json:"returnval"`
}

type VslmRetrieveVStorageObjectState VslmRetrieveVStorageObjectStateRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectState", reflect.TypeOf((*VslmRetrieveVStorageObjectState)(nil)).Elem())
}

type VslmRetrieveVStorageObjectStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id   types.ID                     `xml:"id" json:"id"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectStateRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectStateRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectStateResponse struct {
	Returnval types.VStorageObjectStateInfo `xml:"returnval" json:"returnval"`
}

type VslmRetrieveVStorageObjects VslmRetrieveVStorageObjectsRequestType

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjects", reflect.TypeOf((*VslmRetrieveVStorageObjects)(nil)).Elem())
}

type VslmRetrieveVStorageObjectsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	Ids  []types.ID                   `xml:"ids,omitempty" json:"ids,omitempty"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectsRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectsRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectsResponse struct {
	Returnval []VslmVsoVStorageObjectResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type VslmRevertVStorageObjectRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	SnapshotId types.ID                     `xml:"snapshotId" json:"snapshotId"`
}

func init() {
	types.Add("vslm:VslmRevertVStorageObjectRequestType", reflect.TypeOf((*VslmRevertVStorageObjectRequestType)(nil)).Elem())
}

type VslmRevertVStorageObject_Task VslmRevertVStorageObjectRequestType

func init() {
	types.Add("vslm:VslmRevertVStorageObject_Task", reflect.TypeOf((*VslmRevertVStorageObject_Task)(nil)).Elem())
}

type VslmRevertVStorageObject_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmScheduleReconcileDatastoreInventory VslmScheduleReconcileDatastoreInventoryRequestType

func init() {
	types.Add("vslm:VslmScheduleReconcileDatastoreInventory", reflect.TypeOf((*VslmScheduleReconcileDatastoreInventory)(nil)).Elem())
}

type VslmScheduleReconcileDatastoreInventoryRequestType struct {
	This      types.ManagedObjectReference `xml:"_this" json:"_this"`
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("vslm:VslmScheduleReconcileDatastoreInventoryRequestType", reflect.TypeOf((*VslmScheduleReconcileDatastoreInventoryRequestType)(nil)).Elem())
}

type VslmScheduleReconcileDatastoreInventoryResponse struct {
}

type VslmServiceInstanceContent struct {
	types.DynamicData

	AboutInfo               VslmAboutInfo                `xml:"aboutInfo" json:"aboutInfo"`
	SessionManager          types.ManagedObjectReference `xml:"sessionManager" json:"sessionManager"`
	VStorageObjectManager   types.ManagedObjectReference `xml:"vStorageObjectManager" json:"vStorageObjectManager"`
	StorageLifecycleManager types.ManagedObjectReference `xml:"storageLifecycleManager" json:"storageLifecycleManager"`
}

func init() {
	types.Add("vslm:VslmServiceInstanceContent", reflect.TypeOf((*VslmServiceInstanceContent)(nil)).Elem())
}

type VslmSetVStorageObjectControlFlags VslmSetVStorageObjectControlFlagsRequestType

func init() {
	types.Add("vslm:VslmSetVStorageObjectControlFlags", reflect.TypeOf((*VslmSetVStorageObjectControlFlags)(nil)).Elem())
}

type VslmSetVStorageObjectControlFlagsRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id           types.ID                     `xml:"id" json:"id"`
	ControlFlags []string                     `xml:"controlFlags,omitempty" json:"controlFlags,omitempty"`
}

func init() {
	types.Add("vslm:VslmSetVStorageObjectControlFlagsRequestType", reflect.TypeOf((*VslmSetVStorageObjectControlFlagsRequestType)(nil)).Elem())
}

type VslmSetVStorageObjectControlFlagsResponse struct {
}

type VslmSyncDatastore VslmSyncDatastoreRequestType

func init() {
	types.Add("vslm:VslmSyncDatastore", reflect.TypeOf((*VslmSyncDatastore)(nil)).Elem())
}

type VslmSyncDatastoreRequestType struct {
	This         types.ManagedObjectReference `xml:"_this" json:"_this"`
	DatastoreUrl string                       `xml:"datastoreUrl" json:"datastoreUrl"`
	FullSync     bool                         `xml:"fullSync" json:"fullSync"`
	FcdId        *types.ID                    `xml:"fcdId,omitempty" json:"fcdId,omitempty"`
}

func init() {
	types.Add("vslm:VslmSyncDatastoreRequestType", reflect.TypeOf((*VslmSyncDatastoreRequestType)(nil)).Elem())
}

type VslmSyncDatastoreResponse struct {
}

type VslmSyncFault struct {
	VslmFault

	Id *types.ID `xml:"id,omitempty" json:"id,omitempty"`
}

func init() {
	types.Add("vslm:VslmSyncFault", reflect.TypeOf((*VslmSyncFault)(nil)).Elem())
}

type VslmSyncFaultFault VslmSyncFault

func init() {
	types.Add("vslm:VslmSyncFaultFault", reflect.TypeOf((*VslmSyncFaultFault)(nil)).Elem())
}

type VslmTaskInfo struct {
	types.DynamicData

	Key           string                         `xml:"key" json:"key"`
	Task          types.ManagedObjectReference   `xml:"task" json:"task"`
	Description   *types.LocalizableMessage      `xml:"description,omitempty" json:"description,omitempty"`
	Name          string                         `xml:"name,omitempty" json:"name,omitempty"`
	DescriptionId string                         `xml:"descriptionId" json:"descriptionId"`
	Entity        *types.ManagedObjectReference  `xml:"entity,omitempty" json:"entity,omitempty"`
	EntityName    string                         `xml:"entityName,omitempty" json:"entityName,omitempty"`
	Locked        []types.ManagedObjectReference `xml:"locked,omitempty" json:"locked,omitempty"`
	State         VslmTaskInfoState              `xml:"state" json:"state"`
	Cancelled     bool                           `xml:"cancelled" json:"cancelled"`
	Cancelable    bool                           `xml:"cancelable" json:"cancelable"`
	Error         *types.LocalizedMethodFault    `xml:"error,omitempty" json:"error,omitempty"`
	Result        types.AnyType                  `xml:"result,omitempty,typeattr" json:"result,omitempty"`
	Progress      int32                          `xml:"progress,omitempty" json:"progress,omitempty"`
	Reason        BaseVslmTaskReason             `xml:"reason,typeattr" json:"reason"`
	QueueTime     time.Time                      `xml:"queueTime" json:"queueTime"`
	StartTime     *time.Time                     `xml:"startTime" json:"startTime,omitempty"`
	CompleteTime  *time.Time                     `xml:"completeTime" json:"completeTime,omitempty"`
	EventChainId  int32                          `xml:"eventChainId" json:"eventChainId"`
	ChangeTag     string                         `xml:"changeTag,omitempty" json:"changeTag,omitempty"`
	ParentTaskKey string                         `xml:"parentTaskKey,omitempty" json:"parentTaskKey,omitempty"`
	RootTaskKey   string                         `xml:"rootTaskKey,omitempty" json:"rootTaskKey,omitempty"`
	ActivationId  string                         `xml:"activationId,omitempty" json:"activationId,omitempty"`
}

func init() {
	types.Add("vslm:VslmTaskInfo", reflect.TypeOf((*VslmTaskInfo)(nil)).Elem())
}

type VslmTaskReason struct {
	types.DynamicData
}

func init() {
	types.Add("vslm:VslmTaskReason", reflect.TypeOf((*VslmTaskReason)(nil)).Elem())
}

type VslmTaskReasonAlarm struct {
	VslmTaskReason

	AlarmName  string                       `xml:"alarmName" json:"alarmName"`
	Alarm      types.ManagedObjectReference `xml:"alarm" json:"alarm"`
	EntityName string                       `xml:"entityName" json:"entityName"`
	Entity     types.ManagedObjectReference `xml:"entity" json:"entity"`
}

func init() {
	types.Add("vslm:VslmTaskReasonAlarm", reflect.TypeOf((*VslmTaskReasonAlarm)(nil)).Elem())
}

type VslmTaskReasonSchedule struct {
	VslmTaskReason

	Name          string                       `xml:"name" json:"name"`
	ScheduledTask types.ManagedObjectReference `xml:"scheduledTask" json:"scheduledTask"`
}

func init() {
	types.Add("vslm:VslmTaskReasonSchedule", reflect.TypeOf((*VslmTaskReasonSchedule)(nil)).Elem())
}

type VslmTaskReasonSystem struct {
	VslmTaskReason
}

func init() {
	types.Add("vslm:VslmTaskReasonSystem", reflect.TypeOf((*VslmTaskReasonSystem)(nil)).Elem())
}

type VslmTaskReasonUser struct {
	VslmTaskReason

	UserName string `xml:"userName" json:"userName"`
}

func init() {
	types.Add("vslm:VslmTaskReasonUser", reflect.TypeOf((*VslmTaskReasonUser)(nil)).Elem())
}

type VslmUpdateVStorageInfrastructureObjectPolicyRequestType struct {
	This types.ManagedObjectReference             `xml:"_this" json:"_this"`
	Spec types.VslmInfrastructureObjectPolicySpec `xml:"spec" json:"spec"`
}

func init() {
	types.Add("vslm:VslmUpdateVStorageInfrastructureObjectPolicyRequestType", reflect.TypeOf((*VslmUpdateVStorageInfrastructureObjectPolicyRequestType)(nil)).Elem())
}

type VslmUpdateVStorageInfrastructureObjectPolicy_Task VslmUpdateVStorageInfrastructureObjectPolicyRequestType

func init() {
	types.Add("vslm:VslmUpdateVStorageInfrastructureObjectPolicy_Task", reflect.TypeOf((*VslmUpdateVStorageInfrastructureObjectPolicy_Task)(nil)).Elem())
}

type VslmUpdateVStorageInfrastructureObjectPolicy_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmUpdateVStorageObjectMetadataRequestType struct {
	This       types.ManagedObjectReference `xml:"_this" json:"_this"`
	Id         types.ID                     `xml:"id" json:"id"`
	Metadata   []types.KeyValue             `xml:"metadata,omitempty" json:"metadata,omitempty"`
	DeleteKeys []string                     `xml:"deleteKeys,omitempty" json:"deleteKeys,omitempty"`
}

func init() {
	types.Add("vslm:VslmUpdateVStorageObjectMetadataRequestType", reflect.TypeOf((*VslmUpdateVStorageObjectMetadataRequestType)(nil)).Elem())
}

type VslmUpdateVStorageObjectMetadata_Task VslmUpdateVStorageObjectMetadataRequestType

func init() {
	types.Add("vslm:VslmUpdateVStorageObjectMetadata_Task", reflect.TypeOf((*VslmUpdateVStorageObjectMetadata_Task)(nil)).Elem())
}

type VslmUpdateVStorageObjectMetadata_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmUpdateVstorageObjectCryptoRequestType struct {
	This        types.ManagedObjectReference      `xml:"_this" json:"_this"`
	Id          types.ID                          `xml:"id" json:"id"`
	Profile     []types.VirtualMachineProfileSpec `xml:"profile,omitempty" json:"profile,omitempty"`
	DisksCrypto *types.DiskCryptoSpec             `xml:"disksCrypto,omitempty" json:"disksCrypto,omitempty"`
}

func init() {
	types.Add("vslm:VslmUpdateVstorageObjectCryptoRequestType", reflect.TypeOf((*VslmUpdateVstorageObjectCryptoRequestType)(nil)).Elem())
}

type VslmUpdateVstorageObjectCrypto_Task VslmUpdateVstorageObjectCryptoRequestType

func init() {
	types.Add("vslm:VslmUpdateVstorageObjectCrypto_Task", reflect.TypeOf((*VslmUpdateVstorageObjectCrypto_Task)(nil)).Elem())
}

type VslmUpdateVstorageObjectCrypto_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmUpdateVstorageObjectPolicyRequestType struct {
	This    types.ManagedObjectReference      `xml:"_this" json:"_this"`
	Id      types.ID                          `xml:"id" json:"id"`
	Profile []types.VirtualMachineProfileSpec `xml:"profile,omitempty" json:"profile,omitempty"`
}

func init() {
	types.Add("vslm:VslmUpdateVstorageObjectPolicyRequestType", reflect.TypeOf((*VslmUpdateVstorageObjectPolicyRequestType)(nil)).Elem())
}

type VslmUpdateVstorageObjectPolicy_Task VslmUpdateVstorageObjectPolicyRequestType

func init() {
	types.Add("vslm:VslmUpdateVstorageObjectPolicy_Task", reflect.TypeOf((*VslmUpdateVstorageObjectPolicy_Task)(nil)).Elem())
}

type VslmUpdateVstorageObjectPolicy_TaskResponse struct {
	Returnval types.ManagedObjectReference `xml:"returnval" json:"returnval"`
}

type VslmVsoVStorageObjectAssociations struct {
	types.DynamicData

	Id                types.ID                                             `xml:"id" json:"id"`
	VmDiskAssociation []VslmVsoVStorageObjectAssociationsVmDiskAssociation `xml:"vmDiskAssociation,omitempty" json:"vmDiskAssociation,omitempty"`
	Fault             *types.LocalizedMethodFault                          `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectAssociations", reflect.TypeOf((*VslmVsoVStorageObjectAssociations)(nil)).Elem())
}

type VslmVsoVStorageObjectAssociationsVmDiskAssociation struct {
	types.DynamicData

	VmId    string `xml:"vmId" json:"vmId"`
	DiskKey int32  `xml:"diskKey" json:"diskKey"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectAssociationsVmDiskAssociation", reflect.TypeOf((*VslmVsoVStorageObjectAssociationsVmDiskAssociation)(nil)).Elem())
}

type VslmVsoVStorageObjectQueryResult struct {
	types.DynamicData

	AllRecordsReturned bool                          `xml:"allRecordsReturned" json:"allRecordsReturned"`
	Id                 []types.ID                    `xml:"id,omitempty" json:"id,omitempty"`
	QueryResults       []VslmVsoVStorageObjectResult `xml:"queryResults,omitempty" json:"queryResults,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQueryResult", reflect.TypeOf((*VslmVsoVStorageObjectQueryResult)(nil)).Elem())
}

type VslmVsoVStorageObjectQuerySpec struct {
	types.DynamicData

	QueryField    string   `xml:"queryField" json:"queryField"`
	QueryOperator string   `xml:"queryOperator" json:"queryOperator"`
	QueryValue    []string `xml:"queryValue,omitempty" json:"queryValue,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpec", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpec)(nil)).Elem())
}

type VslmVsoVStorageObjectResult struct {
	types.DynamicData

	Id               types.ID                              `xml:"id" json:"id"`
	Name             string                                `xml:"name,omitempty" json:"name,omitempty"`
	CapacityInMB     int64                                 `xml:"capacityInMB" json:"capacityInMB"`
	CreateTime       *time.Time                            `xml:"createTime" json:"createTime,omitempty"`
	DatastoreUrl     string                                `xml:"datastoreUrl,omitempty" json:"datastoreUrl,omitempty"`
	DiskPath         string                                `xml:"diskPath,omitempty" json:"diskPath,omitempty"`
	UsedCapacityInMB int64                                 `xml:"usedCapacityInMB,omitempty" json:"usedCapacityInMB,omitempty"`
	BackingObjectId  *types.ID                             `xml:"backingObjectId,omitempty" json:"backingObjectId,omitempty"`
	SnapshotInfo     []VslmVsoVStorageObjectSnapshotResult `xml:"snapshotInfo,omitempty" json:"snapshotInfo,omitempty"`
	Metadata         []types.KeyValue                      `xml:"metadata,omitempty" json:"metadata,omitempty"`
	Error            *types.LocalizedMethodFault           `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectResult", reflect.TypeOf((*VslmVsoVStorageObjectResult)(nil)).Elem())
}

type VslmVsoVStorageObjectSnapshotResult struct {
	types.DynamicData

	BackingObjectId types.ID  `xml:"backingObjectId" json:"backingObjectId"`
	Description     string    `xml:"description,omitempty" json:"description,omitempty"`
	SnapshotId      *types.ID `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	DiskPath        string    `xml:"diskPath,omitempty" json:"diskPath,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectSnapshotResult", reflect.TypeOf((*VslmVsoVStorageObjectSnapshotResult)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("vslm:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
