// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

// A boxed array of `VslmDatastoreSyncStatus`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmDatastoreSyncStatus struct {
	VslmDatastoreSyncStatus []VslmDatastoreSyncStatus `xml:"VslmDatastoreSyncStatus,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmDatastoreSyncStatus", reflect.TypeOf((*ArrayOfVslmDatastoreSyncStatus)(nil)).Elem())
}

// A boxed array of `VslmQueryDatastoreInfoResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmQueryDatastoreInfoResult struct {
	VslmQueryDatastoreInfoResult []VslmQueryDatastoreInfoResult `xml:"VslmQueryDatastoreInfoResult,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmQueryDatastoreInfoResult", reflect.TypeOf((*ArrayOfVslmQueryDatastoreInfoResult)(nil)).Elem())
}

// A boxed array of `VslmVsoVStorageObjectAssociations`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmVsoVStorageObjectAssociations struct {
	VslmVsoVStorageObjectAssociations []VslmVsoVStorageObjectAssociations `xml:"VslmVsoVStorageObjectAssociations,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectAssociations", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectAssociations)(nil)).Elem())
}

// A boxed array of `VslmVsoVStorageObjectAssociationsVmDiskAssociation`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation struct {
	VslmVsoVStorageObjectAssociationsVmDiskAssociation []VslmVsoVStorageObjectAssociationsVmDiskAssociation `xml:"VslmVsoVStorageObjectAssociationsVmDiskAssociation,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectAssociationsVmDiskAssociation)(nil)).Elem())
}

// A boxed array of `VslmVsoVStorageObjectQuerySpec`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmVsoVStorageObjectQuerySpec struct {
	VslmVsoVStorageObjectQuerySpec []VslmVsoVStorageObjectQuerySpec `xml:"VslmVsoVStorageObjectQuerySpec,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectQuerySpec", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectQuerySpec)(nil)).Elem())
}

// A boxed array of `VslmVsoVStorageObjectResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmVsoVStorageObjectResult struct {
	VslmVsoVStorageObjectResult []VslmVsoVStorageObjectResult `xml:"VslmVsoVStorageObjectResult,omitempty" json:"_value"`
}

func init() {
	types.Add("vslm:ArrayOfVslmVsoVStorageObjectResult", reflect.TypeOf((*ArrayOfVslmVsoVStorageObjectResult)(nil)).Elem())
}

// A boxed array of `VslmVsoVStorageObjectSnapshotResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/vslm`.
type ArrayOfVslmVsoVStorageObjectSnapshotResult struct {
	VslmVsoVStorageObjectSnapshotResult []VslmVsoVStorageObjectSnapshotResult `xml:"VslmVsoVStorageObjectSnapshotResult,omitempty" json:"_value"`
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

// This data object type describes system information.
//
// This structure may be used only with operations rendered under `/vslm`.
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

// The parameters of `VslmVStorageObjectManager.VslmAttachDisk_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmAttachDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual disk to be operated. See
	// `ID`
	Id types.ID `xml:"id" json:"id"`
	// The virtual machine where the virtual disk is to be attached.
	//
	// Refers instance of `VirtualMachine`.
	Vm types.ManagedObjectReference `xml:"vm" json:"vm"`
	// Key of the controller the disk will connect to.
	// It can be unset if there is only one controller
	// (SCSI or SATA) with the available slot in the
	// virtual machine. If there are multiple SCSI or
	// SATA controllers available, user must specify
	// the controller; if there is no available
	// controllers, a `MissingController`
	// fault will be thrown.
	ControllerKey int32 `xml:"controllerKey,omitempty" json:"controllerKey,omitempty"`
	// The unit number of the attached disk on its controller.
	// If unset, the next available slot on the specified
	// controller or the only available controller will be
	// assigned to the attached disk.
	UnitNumber *int32 `xml:"unitNumber" json:"unitNumber,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmAttachTagToVStorageObject`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmAttachTagToVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The identifier(ID) of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The category to which the tag belongs.
	Category string `xml:"category" json:"category"`
	// The tag which has to be associated with the virtual storage
	// object.
	Tag string `xml:"tag" json:"tag"`
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

// The parameters of `VslmVStorageObjectManager.VslmClearVStorageObjectControlFlags`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmClearVStorageObjectControlFlagsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// control flags enum array to be cleared on the
	// VStorageObject. All control flags not included
	// in the array remain intact.
	ControlFlags []string `xml:"controlFlags,omitempty" json:"controlFlags,omitempty"`
}

func init() {
	types.Add("vslm:VslmClearVStorageObjectControlFlagsRequestType", reflect.TypeOf((*VslmClearVStorageObjectControlFlagsRequestType)(nil)).Elem())
}

type VslmClearVStorageObjectControlFlagsResponse struct {
}

// The parameters of `VslmVStorageObjectManager.VslmCloneVStorageObject_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmCloneVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The specification for cloning the virtual storage
	// object.
	Spec types.VslmCloneSpec `xml:"spec" json:"spec"`
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

// The parameters of `VslmVStorageObjectManager.VslmCreateDiskFromSnapshot_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmCreateDiskFromSnapshotRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of the virtual storage object.
	SnapshotId types.ID `xml:"snapshotId" json:"snapshotId"`
	// A user friendly name to be associated with the new disk.
	Name string `xml:"name" json:"name"`
	// SPBM Profile requirement on the new virtual storage object.
	// If not specified datastore default policy would be
	// assigned.
	Profile []types.VirtualMachineProfileSpec `xml:"profile,omitempty" json:"profile,omitempty"`
	// Crypto information of the new disk.
	Crypto *types.CryptoSpec `xml:"crypto,omitempty" json:"crypto,omitempty"`
	// Relative location in the specified datastore where disk needs
	// to be created. If not specified disk gets created at the
	// default VStorageObject location on the specified datastore.
	Path string `xml:"path,omitempty" json:"path,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmCreateDisk_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmCreateDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The specification of the virtual storage object
	// to be created.
	Spec types.VslmCreateSpec `xml:"spec" json:"spec"`
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

// The parameters of `VslmVStorageObjectManager.VslmCreateSnapshot_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmCreateSnapshotRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// A short description to be associated with the snapshot.
	Description string `xml:"description" json:"description"`
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

// DatastoreSyncStatus shows the catalog sync status of a datastore
// and is returned as a result of the VStorageObjectManager
// getGlobalCatalogSyncStatus API.
//
// When syncVClock == objectVClock the global catalog is in sync with the
// local catalog
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmDatastoreSyncStatus struct {
	types.DynamicData

	// The datastore URL as specified in `DatastoreInfo.url`
	DatastoreURL string `xml:"datastoreURL" json:"datastoreURL"`
	ObjectVClock int64  `xml:"objectVClock" json:"objectVClock"`
	SyncVClock   int64  `xml:"syncVClock" json:"syncVClock"`
	// The time representing the last successful sync of the datastore.
	SyncTime *time.Time `xml:"syncTime" json:"syncTime,omitempty"`
	// The number of retries for the Datastore synchronization in failure
	// cases.
	NumberOfRetries int32 `xml:"numberOfRetries,omitempty" json:"numberOfRetries,omitempty"`
	// The fault is set in case of error conditions.
	//
	// If the fault is set,
	// the objectVClock and syncVClock will be set to -1L.
	// Possible Faults:
	// SyncFault If specified datastoreURL failed to sync.
	Error *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("vslm:VslmDatastoreSyncStatus", reflect.TypeOf((*VslmDatastoreSyncStatus)(nil)).Elem())
}

// The parameters of `VslmVStorageObjectManager.VslmDeleteSnapshot_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmDeleteSnapshotRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of a virtual storage object.
	SnapshotId types.ID `xml:"snapshotId" json:"snapshotId"`
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

// The parameters of `VslmVStorageObjectManager.VslmDeleteVStorageObject_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmDeleteVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object to be deleted.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmDetachTagFromVStorageObject`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmDetachTagFromVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The identifier(ID) of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The category to which the tag belongs.
	Category string `xml:"category" json:"category"`
	// The tag which has to be disassociated with the virtual storage
	// object.
	Tag string `xml:"tag" json:"tag"`
}

func init() {
	types.Add("vslm:VslmDetachTagFromVStorageObjectRequestType", reflect.TypeOf((*VslmDetachTagFromVStorageObjectRequestType)(nil)).Elem())
}

type VslmDetachTagFromVStorageObjectResponse struct {
}

// The parameters of `VslmVStorageObjectManager.VslmExtendDisk_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmExtendDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual disk to be extended.
	Id types.ID `xml:"id" json:"id"`
	// The new capacity of the virtual disk in MB.
	NewCapacityInMB int64 `xml:"newCapacityInMB" json:"newCapacityInMB"`
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

// The super class for all VSLM Faults.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmFault struct {
	types.MethodFault

	// The fault message if available.
	Msg string `xml:"msg,omitempty" json:"msg,omitempty"`
}

func init() {
	types.Add("vslm:VslmFault", reflect.TypeOf((*VslmFault)(nil)).Elem())
}

type VslmFaultFault BaseVslmFault

func init() {
	types.Add("vslm:VslmFaultFault", reflect.TypeOf((*VslmFaultFault)(nil)).Elem())
}

// The parameters of `VslmVStorageObjectManager.VslmInflateDisk_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmInflateDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual disk to be inflated.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmListTagsAttachedToVStorageObject`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmListTagsAttachedToVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmListVStorageObjectForSpec`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmListVStorageObjectForSpecRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Query defined using array of
	// `VslmVsoVStorageObjectQuerySpec` objects.
	Query []VslmVsoVStorageObjectQuerySpec `xml:"query,omitempty" json:"query,omitempty"`
	// Maximum number of virtual storage object IDs to return.
	MaxResult int32 `xml:"maxResult" json:"maxResult"`
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

// The parameters of `VslmVStorageObjectManager.VslmListVStorageObjectsAttachedToTag`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmListVStorageObjectsAttachedToTagRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The category to which the tag belongs.
	Category string `xml:"category" json:"category"`
	// The tag to be queried.
	Tag string `xml:"tag" json:"tag"`
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

// The parameters of `VslmSessionManager.VslmLoginByToken`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmLoginByTokenRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The delegated token will be retrieved by the
	// client and delegated to VSLM. VSLM will use this token, on user's
	// behalf, to login to VC for authorization purposes. It is necessary
	// to convert the token to XML because the SAML token itself is
	// not a VMODL Data Object and cannot be used as a parameter.
	DelegatedTokenXml string `xml:"delegatedTokenXml" json:"delegatedTokenXml"`
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

// The parameters of `VslmVStorageObjectManager.VslmQueryChangedDiskAreas`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmQueryChangedDiskAreasRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of a virtual storage object for
	// which changes that have been made since "changeId"
	// should be computed.
	SnapshotId types.ID `xml:"snapshotId" json:"snapshotId"`
	// Start Offset in bytes at which to start computing
	// changes. Typically, callers will make multiple calls
	// to this function, starting with startOffset 0 and then
	// examine the "length" property in the returned
	// DiskChangeInfo structure, repeatedly calling
	// queryChangedDiskAreas until a map for the entire
	// virtual disk has been obtained.
	StartOffset int64 `xml:"startOffset" json:"startOffset"`
	// Identifier referring to a point in the past that should
	// be used as the point in time at which to begin including
	// changes to the disk in the result. A typical use case
	// would be a backup application obtaining a changeId from
	// a virtual disk's backing info when performing a backup.
	// When a subsequent incremental backup is to be performed,
	// this change Id can be used to obtain a list of changed
	// areas on disk.
	ChangeId string `xml:"changeId" json:"changeId"`
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

// The parameters of `VslmStorageLifecycleManager.VslmQueryDatastoreInfo`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmQueryDatastoreInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The datastore URL as specified in
	// `DatastoreInfo.url`
	DatastoreUrl string `xml:"datastoreUrl" json:"datastoreUrl"`
}

func init() {
	types.Add("vslm:VslmQueryDatastoreInfoRequestType", reflect.TypeOf((*VslmQueryDatastoreInfoRequestType)(nil)).Elem())
}

type VslmQueryDatastoreInfoResponse struct {
	Returnval []VslmQueryDatastoreInfoResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// The `VslmQueryDatastoreInfoResult` provides mapping information
// between `Datacenter` and `Datastore`.
//
// This API is returned as a result of
// `VslmStorageLifecycleManager.VslmQueryDatastoreInfo` invocation.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmQueryDatastoreInfoResult struct {
	types.DynamicData

	// Indicates the datacenter containing the
	// `VslmQueryDatastoreInfoResult.datastore`.
	//
	// Refers instance of `Datacenter`.
	Datacenter types.ManagedObjectReference `xml:"datacenter" json:"datacenter"`
	// Indicates the datastore which is contained within the
	// `VslmQueryDatastoreInfoResult.datacenter`.
	//
	// Refers instance of `Datastore`.
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
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

// The parameters of `VslmVStorageObjectManager.VslmQueryGlobalCatalogSyncStatusForDatastore`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmQueryGlobalCatalogSyncStatusForDatastoreRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// URL of the datastore to check synchronization status for
	DatastoreURL string `xml:"datastoreURL" json:"datastoreURL"`
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

// The parameters of `VslmVStorageObjectManager.VslmReconcileDatastoreInventory_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmReconcileDatastoreInventoryRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The datastore that needs to be reconciled.
	//
	// Refers instance of `Datastore`.
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

// The parameters of `VslmVStorageObjectManager.VslmRegisterDisk`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRegisterDiskRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// URL path to the virtual disk.
	Path string `xml:"path" json:"path"`
	// The descriptive name of the disk object. If
	// unset the name will be automatically determined
	// from the path. @see vim.vslm.BaseConfigInfo.name
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

func init() {
	types.Add("vslm:VslmRegisterDiskRequestType", reflect.TypeOf((*VslmRegisterDiskRequestType)(nil)).Elem())
}

type VslmRegisterDiskResponse struct {
	Returnval types.VStorageObject `xml:"returnval" json:"returnval"`
}

// The parameters of `VslmVStorageObjectManager.VslmRelocateVStorageObject_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRelocateVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The specification for relocation of the virtual
	// storage object.
	Spec types.VslmRelocateSpec `xml:"spec" json:"spec"`
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

// The parameters of `VslmVStorageObjectManager.VslmRenameVStorageObject`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRenameVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object to be renamed.
	Id types.ID `xml:"id" json:"id"`
	// The new name for the virtual storage object.
	Name string `xml:"name" json:"name"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveSnapshotDetails`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveSnapshotDetailsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of a virtual storage object.
	SnapshotId types.ID `xml:"snapshotId" json:"snapshotId"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveSnapshotInfo`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveSnapshotInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageInfrastructureObjectPolicy`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageInfrastructureObjectPolicyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Datastore on which policy needs to be retrieved.
	//
	// Refers instance of `Datastore`.
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObjectAssociations`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectAssociationsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The IDs of the virtual storage objects of the query.
	Ids []types.ID `xml:"ids,omitempty" json:"ids,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObjectMetadata`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectMetadataRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of virtual storage object.
	SnapshotId *types.ID `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	// The prefix of the metadata key that needs to be retrieved
	Prefix string `xml:"prefix,omitempty" json:"prefix,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObjectMetadataValue`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectMetadataValueRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of virtual storage object.
	SnapshotId *types.ID `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	// The key for the virtual storage object
	Key string `xml:"key" json:"key"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectMetadataValueRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectMetadataValueRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectMetadataValueResponse struct {
	Returnval string `xml:"returnval" json:"returnval"`
}

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObject`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object to be retrieved.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObjectState`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectStateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object the state to be retrieved.
	Id types.ID `xml:"id" json:"id"`
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

// The parameters of `VslmVStorageObjectManager.VslmRetrieveVStorageObjects`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRetrieveVStorageObjectsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The array of IDs of the virtual storage object to be
	// retrieved.
	Ids []types.ID `xml:"ids,omitempty" json:"ids,omitempty"`
}

func init() {
	types.Add("vslm:VslmRetrieveVStorageObjectsRequestType", reflect.TypeOf((*VslmRetrieveVStorageObjectsRequestType)(nil)).Elem())
}

type VslmRetrieveVStorageObjectsResponse struct {
	Returnval []VslmVsoVStorageObjectResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// The parameters of `VslmVStorageObjectManager.VslmRevertVStorageObject_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmRevertVStorageObjectRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// The ID of the snapshot of a virtual storage object.
	SnapshotId types.ID `xml:"snapshotId" json:"snapshotId"`
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

// The parameters of `VslmVStorageObjectManager.VslmScheduleReconcileDatastoreInventory`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmScheduleReconcileDatastoreInventoryRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The datastore that needs to be reconciled.
	//
	// Refers instance of `Datastore`.
	Datastore types.ManagedObjectReference `xml:"datastore" json:"datastore"`
}

func init() {
	types.Add("vslm:VslmScheduleReconcileDatastoreInventoryRequestType", reflect.TypeOf((*VslmScheduleReconcileDatastoreInventoryRequestType)(nil)).Elem())
}

type VslmScheduleReconcileDatastoreInventoryResponse struct {
}

// The `VslmServiceInstanceContent` data object defines properties for the
// `VslmServiceInstance` managed object.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmServiceInstanceContent struct {
	types.DynamicData

	// Contains information that identifies the Storage Lifecycle Management
	// service.
	AboutInfo VslmAboutInfo `xml:"aboutInfo" json:"aboutInfo"`
	// `VslmSessionManager` contains login APIs to connect to VSLM
	// service.
	//
	// Refers instance of `VslmSessionManager`.
	SessionManager types.ManagedObjectReference `xml:"sessionManager" json:"sessionManager"`
	// `VslmVStorageObjectManager` contains virtual storage object
	// APIs.
	//
	// Refers instance of `VslmVStorageObjectManager`.
	VStorageObjectManager types.ManagedObjectReference `xml:"vStorageObjectManager" json:"vStorageObjectManager"`
	// `VslmStorageLifecycleManager` contains callback APIs to VSLM
	// service.
	//
	// Refers instance of `VslmStorageLifecycleManager`.
	StorageLifecycleManager types.ManagedObjectReference `xml:"storageLifecycleManager" json:"storageLifecycleManager"`
}

func init() {
	types.Add("vslm:VslmServiceInstanceContent", reflect.TypeOf((*VslmServiceInstanceContent)(nil)).Elem())
}

type VslmSetVStorageObjectControlFlags VslmSetVStorageObjectControlFlagsRequestType

func init() {
	types.Add("vslm:VslmSetVStorageObjectControlFlags", reflect.TypeOf((*VslmSetVStorageObjectControlFlags)(nil)).Elem())
}

// The parameters of `VslmVStorageObjectManager.VslmSetVStorageObjectControlFlags`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmSetVStorageObjectControlFlagsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// control flags enum array to be set on the
	// VStorageObject. All control flags not included
	// in the array remain intact.
	ControlFlags []string `xml:"controlFlags,omitempty" json:"controlFlags,omitempty"`
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

// The parameters of `VslmStorageLifecycleManager.VslmSyncDatastore`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmSyncDatastoreRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The datastore URL as specified in
	// `DatastoreInfo.url`
	DatastoreUrl string `xml:"datastoreUrl" json:"datastoreUrl"`
	// If this is set to true, all information for this datastore
	// will be discarded from the catalog and reloaded from the
	// datastore's catalog
	FullSync bool `xml:"fullSync" json:"fullSync"`
	// If set, this call blocks until fcdId is persisted into db
	// if this fcdId is not found in queue, assume persisted and return
	FcdId *types.ID `xml:"fcdId,omitempty" json:"fcdId,omitempty"`
}

func init() {
	types.Add("vslm:VslmSyncDatastoreRequestType", reflect.TypeOf((*VslmSyncDatastoreRequestType)(nil)).Elem())
}

type VslmSyncDatastoreResponse struct {
}

// An SyncFault fault is thrown when there is a failure to synchronize
// the FCD global catalog information with the local catalog information.
//
// Pandora synchronizes the datastore periodically in the background, it
// recovers from any transient failures affecting the datastore or
// individual FCDs. In cases where the sync fault needs to be resolved
// immediately, explicitly triggering a
// `VslmStorageLifecycleManager.VslmSyncDatastore` should resolve the
// issue, unless there are underlying infrastructure issues affecting the
// datastore or FCD. If the fault is ignored there is
// a possibility that the FCD is unrecognized by Pandora or Pandora
// DB having stale information, consequently, affecting the return of
// `VslmVStorageObjectManager.VslmListVStorageObjectForSpec` and
// `VslmVStorageObjectManager.VslmRetrieveVStorageObjects` APIs.
// In cases where the `VslmSyncFault.id` is specified,
// the client can explicitly trigger
// `VslmStorageLifecycleManager.VslmSyncDatastore` to resolve
// the issue, else, could ignore the fault in anticipation of Pandora
// automatically resolving the error.
//
// This structure may be used only with operations rendered under `/vslm`.
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

// This data object type contains all information about a VSLM task.
//
// A task represents an operation performed by VirtualCenter or ESX.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskInfo struct {
	types.DynamicData

	// The unique key for the task.
	Key string `xml:"key" json:"key"`
	// The managed object that represents this task.
	//
	// Refers instance of `VslmTask`.
	Task types.ManagedObjectReference `xml:"task" json:"task"`
	// The description field of the task describes the current phase of
	// operation of the task.
	//
	// For a task that does a single monolithic
	// activity, this will be fixed and unchanging.
	// For tasks that have various substeps, this field will change
	// as the task progresses from one phase to another.
	Description *types.LocalizableMessage `xml:"description,omitempty" json:"description,omitempty"`
	// The name of the operation that created the task.
	//
	// This is not set
	// for internal tasks.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	// An identifier for this operation.
	//
	// This includes publicly visible
	// internal tasks and is a lookup in the TaskDescription methodInfo
	// data object.
	DescriptionId string `xml:"descriptionId" json:"descriptionId"`
	// Managed entity to which the operation applies.
	//
	// Refers instance of `ManagedEntity`.
	Entity *types.ManagedObjectReference `xml:"entity,omitempty" json:"entity,omitempty"`
	// The name of the managed entity, locale-specific, retained for the
	// history collector database.
	EntityName string `xml:"entityName,omitempty" json:"entityName,omitempty"`
	// If the state of the task is "running", then this property is a list of
	// managed entities that the operation has locked, with a shared lock.
	//
	// Refers instances of `ManagedEntity`.
	Locked []types.ManagedObjectReference `xml:"locked,omitempty" json:"locked,omitempty"`
	// Runtime status of the task.
	State VslmTaskInfoState `xml:"state" json:"state"`
	// Flag to indicate whether or not the client requested
	// cancellation of the task.
	Cancelled bool `xml:"cancelled" json:"cancelled"`
	// Flag to indicate whether or not the cancel task operation is supported.
	Cancelable bool `xml:"cancelable" json:"cancelable"`
	// If the task state is "error", then this property contains the fault code.
	Error *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
	// If the task state is "success", then this property may be used
	// to hold a return value.
	Result types.AnyType `xml:"result,omitempty,typeattr" json:"result,omitempty"`
	// If the task state is "running", then this property contains a
	// progress measurement, expressed as percentage completed, from 0 to 100.
	//
	// If this property is not set, then the command does not report progress.
	Progress int32 `xml:"progress,omitempty" json:"progress,omitempty"`
	// Kind of entity responsible for creating this task.
	Reason BaseVslmTaskReason `xml:"reason,typeattr" json:"reason"`
	// Time stamp when the task was created.
	QueueTime time.Time `xml:"queueTime" json:"queueTime"`
	// Time stamp when the task started running.
	StartTime *time.Time `xml:"startTime" json:"startTime,omitempty"`
	// Time stamp when the task was completed (whether success or failure).
	CompleteTime *time.Time `xml:"completeTime" json:"completeTime,omitempty"`
	// Event chain ID that leads to the corresponding events.
	EventChainId int32 `xml:"eventChainId" json:"eventChainId"`
	// The user entered tag to identify the operations and their side effects
	ChangeTag string `xml:"changeTag,omitempty" json:"changeTag,omitempty"`
	// Tasks can be created by another task.
	//
	// This shows `VslmTaskInfo.key` of the task spun off this task. This is to
	// track causality between tasks.
	ParentTaskKey string `xml:"parentTaskKey,omitempty" json:"parentTaskKey,omitempty"`
	// Tasks can be created by another task and such creation can go on for
	// multiple levels.
	//
	// This is the `VslmTaskInfo.key` of the task
	// that started the chain of tasks.
	RootTaskKey string `xml:"rootTaskKey,omitempty" json:"rootTaskKey,omitempty"`
	// The activation Id is a client-provided token to link an API call with a task.
	ActivationId string `xml:"activationId,omitempty" json:"activationId,omitempty"`
}

func init() {
	types.Add("vslm:VslmTaskInfo", reflect.TypeOf((*VslmTaskInfo)(nil)).Elem())
}

// Base type for all task reasons.
//
// Task reasons represent the kind of entity responsible for a task's creation.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskReason struct {
	types.DynamicData
}

func init() {
	types.Add("vslm:VslmTaskReason", reflect.TypeOf((*VslmTaskReason)(nil)).Elem())
}

// Indicates that the task was queued by an alarm.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskReasonAlarm struct {
	VslmTaskReason

	// The name of the alarm that queued the task, retained in the history
	// collector database.
	AlarmName string `xml:"alarmName" json:"alarmName"`
	// The alarm object that queued the task.
	//
	// Refers instance of `Alarm`.
	Alarm types.ManagedObjectReference `xml:"alarm" json:"alarm"`
	// The name of the managed entity on which the alarm is triggered,
	// retained in the history collector database.
	EntityName string `xml:"entityName" json:"entityName"`
	// The managed entity object on which the alarm is triggered.
	//
	// Refers instance of `ManagedEntity`.
	Entity types.ManagedObjectReference `xml:"entity" json:"entity"`
}

func init() {
	types.Add("vslm:VslmTaskReasonAlarm", reflect.TypeOf((*VslmTaskReasonAlarm)(nil)).Elem())
}

// Indicates that the task was queued by a scheduled task.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskReasonSchedule struct {
	VslmTaskReason

	// The name of the scheduled task that queued this task.
	Name string `xml:"name" json:"name"`
	// The scheduledTask object that queued this task.
	//
	// Refers instance of `ScheduledTask`.
	ScheduledTask types.ManagedObjectReference `xml:"scheduledTask" json:"scheduledTask"`
}

func init() {
	types.Add("vslm:VslmTaskReasonSchedule", reflect.TypeOf((*VslmTaskReasonSchedule)(nil)).Elem())
}

// Indicates that the task was started by the system (a default task).
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskReasonSystem struct {
	VslmTaskReason
}

func init() {
	types.Add("vslm:VslmTaskReasonSystem", reflect.TypeOf((*VslmTaskReasonSystem)(nil)).Elem())
}

// Indicates that the task was queued by a specific user.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmTaskReasonUser struct {
	VslmTaskReason

	// Name of the user that queued the task.
	UserName string `xml:"userName" json:"userName"`
}

func init() {
	types.Add("vslm:VslmTaskReasonUser", reflect.TypeOf((*VslmTaskReasonUser)(nil)).Elem())
}

// The parameters of `VslmVStorageObjectManager.VslmUpdateVStorageInfrastructureObjectPolicy_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmUpdateVStorageInfrastructureObjectPolicyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// specification to assign a SPBM policy to FCD infrastructure
	// object.
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

// The parameters of `VslmVStorageObjectManager.VslmUpdateVStorageObjectMetadata_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmUpdateVStorageObjectMetadataRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// array of key/value strings. (keys must be unique
	// within the list)
	Metadata []types.KeyValue `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// array of keys need to be deleted
	DeleteKeys []string `xml:"deleteKeys,omitempty" json:"deleteKeys,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmUpdateVstorageObjectCrypto_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmUpdateVstorageObjectCryptoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// New profile requirement on the virtual storage object.
	Profile []types.VirtualMachineProfileSpec `xml:"profile,omitempty" json:"profile,omitempty"`
	// The crypto information of each disk on the chain.
	DisksCrypto *types.DiskCryptoSpec `xml:"disksCrypto,omitempty" json:"disksCrypto,omitempty"`
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

// The parameters of `VslmVStorageObjectManager.VslmUpdateVstorageObjectPolicy_Task`.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmUpdateVstorageObjectPolicyRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// New profile requirement on the virtual storage object.
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

// This data object is a key-value pair whose key is the virtual storage
// object id, and value is the vm association information.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectAssociations struct {
	types.DynamicData

	// ID of this virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// Array of vm associations related to the virtual storage object.
	VmDiskAssociation []VslmVsoVStorageObjectAssociationsVmDiskAssociation `xml:"vmDiskAssociation,omitempty" json:"vmDiskAssociation,omitempty"`
	// Received error while generating associations.
	Fault *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectAssociations", reflect.TypeOf((*VslmVsoVStorageObjectAssociations)(nil)).Elem())
}

// This data object contains information of a VM Disk association.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectAssociationsVmDiskAssociation struct {
	types.DynamicData

	// ID of the virtual machine.
	VmId string `xml:"vmId" json:"vmId"`
	// Device key of the disk attached to the VM.
	DiskKey int32 `xml:"diskKey" json:"diskKey"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectAssociationsVmDiskAssociation", reflect.TypeOf((*VslmVsoVStorageObjectAssociationsVmDiskAssociation)(nil)).Elem())
}

// The `VslmVsoVStorageObjectQueryResult` contains the result of
// `VslmVStorageObjectManager.VslmListVStorageObjectForSpec` API.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectQueryResult struct {
	types.DynamicData

	// If set to false, more results were found than could be returned (either
	// limited by maxResult input argument in the
	// `VslmVStorageObjectManager.VslmListVStorageObjectForSpec` API or
	// truncated because the number of results exceeded the internal limit).
	AllRecordsReturned bool `xml:"allRecordsReturned" json:"allRecordsReturned"`
	// IDs of the VStorageObjects matching the query criteria
	// NOTE: This field will be removed once the dev/qe code is refactored.
	//
	// IDs will be returned in ascending order. If
	// `VslmVsoVStorageObjectQueryResult.allRecordsReturned` is set to false,
	// to get the additional results, repeat the query with ID &gt; last ID as
	// part of the query spec `VslmVsoVStorageObjectQuerySpec`.
	Id []types.ID `xml:"id,omitempty" json:"id,omitempty"`
	// Results of the query criteria.
	//
	// `IDs` will be returned in
	// ascending order. If `VslmVsoVStorageObjectQueryResult.allRecordsReturned`
	// is set to false,then, to get the additional results, repeat the query
	// with ID &gt; last ID as part of the query spec
	// `VslmVsoVStorageObjectQuerySpec`.
	QueryResults []VslmVsoVStorageObjectResult `xml:"queryResults,omitempty" json:"queryResults,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQueryResult", reflect.TypeOf((*VslmVsoVStorageObjectQueryResult)(nil)).Elem())
}

// The `VslmVsoVStorageObjectQuerySpec` describes the criteria to query
// VStorageObject from global catalog.
//
// `VslmVsoVStorageObjectQuerySpec` is sent as input to
// `VslmVStorageObjectManager.VslmListVStorageObjectForSpec` API.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectQuerySpec struct {
	types.DynamicData

	// This field specifies the searchable field.
	//
	// This can be one of the values from
	// `VslmVsoVStorageObjectQuerySpecQueryFieldEnum_enum`.
	QueryField string `xml:"queryField" json:"queryField"`
	// This field specifies the operator to compare the searchable field
	// `VslmVsoVStorageObjectQuerySpec.queryField` with the specified
	// value `VslmVsoVStorageObjectQuerySpec.queryValue`.
	//
	// This can be one of the values from `VslmVsoVStorageObjectQuerySpecQueryOperatorEnum_enum`.
	QueryOperator string `xml:"queryOperator" json:"queryOperator"`
	// This field specifies the value to be compared with the searchable field.
	QueryValue []string `xml:"queryValue,omitempty" json:"queryValue,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectQuerySpec", reflect.TypeOf((*VslmVsoVStorageObjectQuerySpec)(nil)).Elem())
}

// The `VslmVsoVStorageObjectResult` contains the result objects of
// `VslmVsoVStorageObjectQueryResult` which is returned as a result of
// `VslmVStorageObjectManager.VslmListVStorageObjectForSpec` and
// `VslmVStorageObjectManager.VslmRetrieveVStorageObjects` APIs.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectResult struct {
	types.DynamicData

	// The ID of the virtual storage object.
	Id types.ID `xml:"id" json:"id"`
	// Name of FCD.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	// The size in MB of this object.
	//
	// If the faults are set,
	// then the capacityInMB will be -1
	CapacityInMB int64 `xml:"capacityInMB" json:"capacityInMB"`
	// The create time information of the FCD.
	CreateTime *time.Time `xml:"createTime" json:"createTime,omitempty"`
	// The Datastore URL containing the FCD.
	DatastoreUrl string `xml:"datastoreUrl,omitempty" json:"datastoreUrl,omitempty"`
	// The disk path of the FCD.
	DiskPath string `xml:"diskPath,omitempty" json:"diskPath,omitempty"`
	// The rolled up used capacity of the FCD and it's snapshots.
	//
	// Returns -1L if the space information is currently unavailable.
	UsedCapacityInMB int64 `xml:"usedCapacityInMB,omitempty" json:"usedCapacityInMB,omitempty"`
	// The ID of the backing object of the virtual storage object.
	BackingObjectId *types.ID `xml:"backingObjectId,omitempty" json:"backingObjectId,omitempty"`
	// VStorageObjectSnapshotResult array containing information about all the
	// snapshots of the virtual storage object.
	SnapshotInfo []VslmVsoVStorageObjectSnapshotResult `xml:"snapshotInfo,omitempty" json:"snapshotInfo,omitempty"`
	// Metadata array of key/value strings.
	Metadata []types.KeyValue `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// The fault is set in case of error conditions and this property will
	// have the reason.
	//
	// Possible Faults:
	// NotFound If specified virtual storage object cannot be found.
	Error *types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectResult", reflect.TypeOf((*VslmVsoVStorageObjectResult)(nil)).Elem())
}

// The `VslmVsoVStorageObjectSnapshotResult` contains brief information about a
// snapshot of the object `VslmVsoVStorageObjectResult` which is returned as a
// result of `VslmVStorageObjectManager.VslmRetrieveVStorageObjects` API.
//
// This structure may be used only with operations rendered under `/vslm`.
type VslmVsoVStorageObjectSnapshotResult struct {
	types.DynamicData

	// The ID of the vsan object backing a snapshot of the virtual storage
	// object.
	BackingObjectId types.ID `xml:"backingObjectId" json:"backingObjectId"`
	// The description user passed in when creating this snapshot.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// The ID of this snapshot, created and used in fcd layer.
	SnapshotId *types.ID `xml:"snapshotId,omitempty" json:"snapshotId,omitempty"`
	// The file path of this snapshot.
	DiskPath string `xml:"diskPath,omitempty" json:"diskPath,omitempty"`
}

func init() {
	types.Add("vslm:VslmVsoVStorageObjectSnapshotResult", reflect.TypeOf((*VslmVsoVStorageObjectSnapshotResult)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("vslm:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
