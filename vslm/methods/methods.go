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

package methods

import (
	"context"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vslm/types"
)

type RetrieveContentBody struct {
	Req    *types.RetrieveContent         `xml:"urn:vslm RetrieveContent,omitempty"`
	Res    *types.RetrieveContentResponse `xml:"urn:vslm RetrieveContentResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveContentBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveContent(ctx context.Context, r soap.RoundTripper, req *types.RetrieveContent) (*types.RetrieveContentResponse, error) {
	var reqBody, resBody RetrieveContentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmAttachDisk_TaskBody struct {
	Req    *types.VslmAttachDisk_Task         `xml:"urn:vslm VslmAttachDisk_Task,omitempty"`
	Res    *types.VslmAttachDisk_TaskResponse `xml:"urn:vslm VslmAttachDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmAttachDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmAttachDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmAttachDisk_Task) (*types.VslmAttachDisk_TaskResponse, error) {
	var reqBody, resBody VslmAttachDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmAttachTagToVStorageObjectBody struct {
	Req    *types.VslmAttachTagToVStorageObject         `xml:"urn:vslm VslmAttachTagToVStorageObject,omitempty"`
	Res    *types.VslmAttachTagToVStorageObjectResponse `xml:"urn:vslm VslmAttachTagToVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmAttachTagToVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VslmAttachTagToVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.VslmAttachTagToVStorageObject) (*types.VslmAttachTagToVStorageObjectResponse, error) {
	var reqBody, resBody VslmAttachTagToVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmCancelTaskBody struct {
	Req    *types.VslmCancelTask         `xml:"urn:vslm VslmCancelTask,omitempty"`
	Res    *types.VslmCancelTaskResponse `xml:"urn:vslm VslmCancelTaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmCancelTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmCancelTask(ctx context.Context, r soap.RoundTripper, req *types.VslmCancelTask) (*types.VslmCancelTaskResponse, error) {
	var reqBody, resBody VslmCancelTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmClearVStorageObjectControlFlagsBody struct {
	Req    *types.VslmClearVStorageObjectControlFlags         `xml:"urn:vslm VslmClearVStorageObjectControlFlags,omitempty"`
	Res    *types.VslmClearVStorageObjectControlFlagsResponse `xml:"urn:vslm VslmClearVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmClearVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func VslmClearVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.VslmClearVStorageObjectControlFlags) (*types.VslmClearVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody VslmClearVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmCloneVStorageObject_TaskBody struct {
	Req    *types.VslmCloneVStorageObject_Task         `xml:"urn:vslm VslmCloneVStorageObject_Task,omitempty"`
	Res    *types.VslmCloneVStorageObject_TaskResponse `xml:"urn:vslm VslmCloneVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmCloneVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmCloneVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmCloneVStorageObject_Task) (*types.VslmCloneVStorageObject_TaskResponse, error) {
	var reqBody, resBody VslmCloneVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmCreateDiskFromSnapshot_TaskBody struct {
	Req    *types.VslmCreateDiskFromSnapshot_Task         `xml:"urn:vslm VslmCreateDiskFromSnapshot_Task,omitempty"`
	Res    *types.VslmCreateDiskFromSnapshot_TaskResponse `xml:"urn:vslm VslmCreateDiskFromSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmCreateDiskFromSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmCreateDiskFromSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmCreateDiskFromSnapshot_Task) (*types.VslmCreateDiskFromSnapshot_TaskResponse, error) {
	var reqBody, resBody VslmCreateDiskFromSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmCreateDisk_TaskBody struct {
	Req    *types.VslmCreateDisk_Task         `xml:"urn:vslm VslmCreateDisk_Task,omitempty"`
	Res    *types.VslmCreateDisk_TaskResponse `xml:"urn:vslm VslmCreateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmCreateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmCreateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmCreateDisk_Task) (*types.VslmCreateDisk_TaskResponse, error) {
	var reqBody, resBody VslmCreateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmCreateSnapshot_TaskBody struct {
	Req    *types.VslmCreateSnapshot_Task         `xml:"urn:vslm VslmCreateSnapshot_Task,omitempty"`
	Res    *types.VslmCreateSnapshot_TaskResponse `xml:"urn:vslm VslmCreateSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmCreateSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmCreateSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmCreateSnapshot_Task) (*types.VslmCreateSnapshot_TaskResponse, error) {
	var reqBody, resBody VslmCreateSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmDeleteSnapshot_TaskBody struct {
	Req    *types.VslmDeleteSnapshot_Task         `xml:"urn:vslm VslmDeleteSnapshot_Task,omitempty"`
	Res    *types.VslmDeleteSnapshot_TaskResponse `xml:"urn:vslm VslmDeleteSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmDeleteSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmDeleteSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmDeleteSnapshot_Task) (*types.VslmDeleteSnapshot_TaskResponse, error) {
	var reqBody, resBody VslmDeleteSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmDeleteVStorageObject_TaskBody struct {
	Req    *types.VslmDeleteVStorageObject_Task         `xml:"urn:vslm VslmDeleteVStorageObject_Task,omitempty"`
	Res    *types.VslmDeleteVStorageObject_TaskResponse `xml:"urn:vslm VslmDeleteVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmDeleteVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmDeleteVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmDeleteVStorageObject_Task) (*types.VslmDeleteVStorageObject_TaskResponse, error) {
	var reqBody, resBody VslmDeleteVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmDetachTagFromVStorageObjectBody struct {
	Req    *types.VslmDetachTagFromVStorageObject         `xml:"urn:vslm VslmDetachTagFromVStorageObject,omitempty"`
	Res    *types.VslmDetachTagFromVStorageObjectResponse `xml:"urn:vslm VslmDetachTagFromVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmDetachTagFromVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VslmDetachTagFromVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.VslmDetachTagFromVStorageObject) (*types.VslmDetachTagFromVStorageObjectResponse, error) {
	var reqBody, resBody VslmDetachTagFromVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmExtendDisk_TaskBody struct {
	Req    *types.VslmExtendDisk_Task         `xml:"urn:vslm VslmExtendDisk_Task,omitempty"`
	Res    *types.VslmExtendDisk_TaskResponse `xml:"urn:vslm VslmExtendDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmExtendDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmExtendDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmExtendDisk_Task) (*types.VslmExtendDisk_TaskResponse, error) {
	var reqBody, resBody VslmExtendDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmInflateDisk_TaskBody struct {
	Req    *types.VslmInflateDisk_Task         `xml:"urn:vslm VslmInflateDisk_Task,omitempty"`
	Res    *types.VslmInflateDisk_TaskResponse `xml:"urn:vslm VslmInflateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmInflateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmInflateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmInflateDisk_Task) (*types.VslmInflateDisk_TaskResponse, error) {
	var reqBody, resBody VslmInflateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmListTagsAttachedToVStorageObjectBody struct {
	Req    *types.VslmListTagsAttachedToVStorageObject         `xml:"urn:vslm VslmListTagsAttachedToVStorageObject,omitempty"`
	Res    *types.VslmListTagsAttachedToVStorageObjectResponse `xml:"urn:vslm VslmListTagsAttachedToVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmListTagsAttachedToVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VslmListTagsAttachedToVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.VslmListTagsAttachedToVStorageObject) (*types.VslmListTagsAttachedToVStorageObjectResponse, error) {
	var reqBody, resBody VslmListTagsAttachedToVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmListVStorageObjectForSpecBody struct {
	Req    *types.VslmListVStorageObjectForSpec         `xml:"urn:vslm VslmListVStorageObjectForSpec,omitempty"`
	Res    *types.VslmListVStorageObjectForSpecResponse `xml:"urn:vslm VslmListVStorageObjectForSpecResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmListVStorageObjectForSpecBody) Fault() *soap.Fault { return b.Fault_ }

func VslmListVStorageObjectForSpec(ctx context.Context, r soap.RoundTripper, req *types.VslmListVStorageObjectForSpec) (*types.VslmListVStorageObjectForSpecResponse, error) {
	var reqBody, resBody VslmListVStorageObjectForSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmListVStorageObjectsAttachedToTagBody struct {
	Req    *types.VslmListVStorageObjectsAttachedToTag         `xml:"urn:vslm VslmListVStorageObjectsAttachedToTag,omitempty"`
	Res    *types.VslmListVStorageObjectsAttachedToTagResponse `xml:"urn:vslm VslmListVStorageObjectsAttachedToTagResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmListVStorageObjectsAttachedToTagBody) Fault() *soap.Fault { return b.Fault_ }

func VslmListVStorageObjectsAttachedToTag(ctx context.Context, r soap.RoundTripper, req *types.VslmListVStorageObjectsAttachedToTag) (*types.VslmListVStorageObjectsAttachedToTagResponse, error) {
	var reqBody, resBody VslmListVStorageObjectsAttachedToTagBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmLoginByTokenBody struct {
	Req    *types.VslmLoginByToken         `xml:"urn:vslm VslmLoginByToken,omitempty"`
	Res    *types.VslmLoginByTokenResponse `xml:"urn:vslm VslmLoginByTokenResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmLoginByTokenBody) Fault() *soap.Fault { return b.Fault_ }

func VslmLoginByToken(ctx context.Context, r soap.RoundTripper, req *types.VslmLoginByToken) (*types.VslmLoginByTokenResponse, error) {
	var reqBody, resBody VslmLoginByTokenBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmLogoutBody struct {
	Req    *types.VslmLogout         `xml:"urn:vslm VslmLogout,omitempty"`
	Res    *types.VslmLogoutResponse `xml:"urn:vslm VslmLogoutResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmLogoutBody) Fault() *soap.Fault { return b.Fault_ }

func VslmLogout(ctx context.Context, r soap.RoundTripper, req *types.VslmLogout) (*types.VslmLogoutResponse, error) {
	var reqBody, resBody VslmLogoutBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryChangedDiskAreasBody struct {
	Req    *types.VslmQueryChangedDiskAreas         `xml:"urn:vslm VslmQueryChangedDiskAreas,omitempty"`
	Res    *types.VslmQueryChangedDiskAreasResponse `xml:"urn:vslm VslmQueryChangedDiskAreasResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryChangedDiskAreasBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryChangedDiskAreas(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryChangedDiskAreas) (*types.VslmQueryChangedDiskAreasResponse, error) {
	var reqBody, resBody VslmQueryChangedDiskAreasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryDatastoreInfoBody struct {
	Req    *types.VslmQueryDatastoreInfo         `xml:"urn:vslm VslmQueryDatastoreInfo,omitempty"`
	Res    *types.VslmQueryDatastoreInfoResponse `xml:"urn:vslm VslmQueryDatastoreInfoResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryDatastoreInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryDatastoreInfo(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryDatastoreInfo) (*types.VslmQueryDatastoreInfoResponse, error) {
	var reqBody, resBody VslmQueryDatastoreInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryGlobalCatalogSyncStatusBody struct {
	Req    *types.VslmQueryGlobalCatalogSyncStatus         `xml:"urn:vslm VslmQueryGlobalCatalogSyncStatus,omitempty"`
	Res    *types.VslmQueryGlobalCatalogSyncStatusResponse `xml:"urn:vslm VslmQueryGlobalCatalogSyncStatusResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryGlobalCatalogSyncStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryGlobalCatalogSyncStatus(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryGlobalCatalogSyncStatus) (*types.VslmQueryGlobalCatalogSyncStatusResponse, error) {
	var reqBody, resBody VslmQueryGlobalCatalogSyncStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryGlobalCatalogSyncStatusForDatastoreBody struct {
	Req    *types.VslmQueryGlobalCatalogSyncStatusForDatastore         `xml:"urn:vslm VslmQueryGlobalCatalogSyncStatusForDatastore,omitempty"`
	Res    *types.VslmQueryGlobalCatalogSyncStatusForDatastoreResponse `xml:"urn:vslm VslmQueryGlobalCatalogSyncStatusForDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryGlobalCatalogSyncStatusForDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryGlobalCatalogSyncStatusForDatastore(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryGlobalCatalogSyncStatusForDatastore) (*types.VslmQueryGlobalCatalogSyncStatusForDatastoreResponse, error) {
	var reqBody, resBody VslmQueryGlobalCatalogSyncStatusForDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryInfoBody struct {
	Req    *types.VslmQueryInfo         `xml:"urn:vslm VslmQueryInfo,omitempty"`
	Res    *types.VslmQueryInfoResponse `xml:"urn:vslm VslmQueryInfoResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryInfo(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryInfo) (*types.VslmQueryInfoResponse, error) {
	var reqBody, resBody VslmQueryInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmQueryTaskResultBody struct {
	Req    *types.VslmQueryTaskResult         `xml:"urn:vslm VslmQueryTaskResult,omitempty"`
	Res    *types.VslmQueryTaskResultResponse `xml:"urn:vslm VslmQueryTaskResultResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmQueryTaskResultBody) Fault() *soap.Fault { return b.Fault_ }

func VslmQueryTaskResult(ctx context.Context, r soap.RoundTripper, req *types.VslmQueryTaskResult) (*types.VslmQueryTaskResultResponse, error) {
	var reqBody, resBody VslmQueryTaskResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmReconcileDatastoreInventory_TaskBody struct {
	Req    *types.VslmReconcileDatastoreInventory_Task         `xml:"urn:vslm VslmReconcileDatastoreInventory_Task,omitempty"`
	Res    *types.VslmReconcileDatastoreInventory_TaskResponse `xml:"urn:vslm VslmReconcileDatastoreInventory_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmReconcileDatastoreInventory_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmReconcileDatastoreInventory_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmReconcileDatastoreInventory_Task) (*types.VslmReconcileDatastoreInventory_TaskResponse, error) {
	var reqBody, resBody VslmReconcileDatastoreInventory_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRegisterDiskBody struct {
	Req    *types.VslmRegisterDisk         `xml:"urn:vslm VslmRegisterDisk,omitempty"`
	Res    *types.VslmRegisterDiskResponse `xml:"urn:vslm VslmRegisterDiskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRegisterDiskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRegisterDisk(ctx context.Context, r soap.RoundTripper, req *types.VslmRegisterDisk) (*types.VslmRegisterDiskResponse, error) {
	var reqBody, resBody VslmRegisterDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRelocateVStorageObject_TaskBody struct {
	Req    *types.VslmRelocateVStorageObject_Task         `xml:"urn:vslm VslmRelocateVStorageObject_Task,omitempty"`
	Res    *types.VslmRelocateVStorageObject_TaskResponse `xml:"urn:vslm VslmRelocateVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRelocateVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRelocateVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmRelocateVStorageObject_Task) (*types.VslmRelocateVStorageObject_TaskResponse, error) {
	var reqBody, resBody VslmRelocateVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRenameVStorageObjectBody struct {
	Req    *types.VslmRenameVStorageObject         `xml:"urn:vslm VslmRenameVStorageObject,omitempty"`
	Res    *types.VslmRenameVStorageObjectResponse `xml:"urn:vslm VslmRenameVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRenameVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRenameVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.VslmRenameVStorageObject) (*types.VslmRenameVStorageObjectResponse, error) {
	var reqBody, resBody VslmRenameVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveSnapshotInfoBody struct {
	Req    *types.VslmRetrieveSnapshotInfo         `xml:"urn:vslm VslmRetrieveSnapshotInfo,omitempty"`
	Res    *types.VslmRetrieveSnapshotInfoResponse `xml:"urn:vslm VslmRetrieveSnapshotInfoResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveSnapshotInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveSnapshotInfo(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveSnapshotInfo) (*types.VslmRetrieveSnapshotInfoResponse, error) {
	var reqBody, resBody VslmRetrieveSnapshotInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageInfrastructureObjectPolicyBody struct {
	Req    *types.VslmRetrieveVStorageInfrastructureObjectPolicy         `xml:"urn:vslm VslmRetrieveVStorageInfrastructureObjectPolicy,omitempty"`
	Res    *types.VslmRetrieveVStorageInfrastructureObjectPolicyResponse `xml:"urn:vslm VslmRetrieveVStorageInfrastructureObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageInfrastructureObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageInfrastructureObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageInfrastructureObjectPolicy) (*types.VslmRetrieveVStorageInfrastructureObjectPolicyResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageInfrastructureObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectBody struct {
	Req    *types.VslmRetrieveVStorageObject         `xml:"urn:vslm VslmRetrieveVStorageObject,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectResponse `xml:"urn:vslm VslmRetrieveVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObject) (*types.VslmRetrieveVStorageObjectResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectAssociationsBody struct {
	Req    *types.VslmRetrieveVStorageObjectAssociations         `xml:"urn:vslm VslmRetrieveVStorageObjectAssociations,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectAssociationsResponse `xml:"urn:vslm VslmRetrieveVStorageObjectAssociationsResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectAssociationsBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObjectAssociations(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObjectAssociations) (*types.VslmRetrieveVStorageObjectAssociationsResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectAssociationsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectMetadataBody struct {
	Req    *types.VslmRetrieveVStorageObjectMetadata         `xml:"urn:vslm VslmRetrieveVStorageObjectMetadata,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectMetadataResponse `xml:"urn:vslm VslmRetrieveVStorageObjectMetadataResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObjectMetadata(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObjectMetadata) (*types.VslmRetrieveVStorageObjectMetadataResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectMetadataValueBody struct {
	Req    *types.VslmRetrieveVStorageObjectMetadataValue         `xml:"urn:vslm VslmRetrieveVStorageObjectMetadataValue,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectMetadataValueResponse `xml:"urn:vslm VslmRetrieveVStorageObjectMetadataValueResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectMetadataValueBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObjectMetadataValue(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObjectMetadataValue) (*types.VslmRetrieveVStorageObjectMetadataValueResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectMetadataValueBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectStateBody struct {
	Req    *types.VslmRetrieveVStorageObjectState         `xml:"urn:vslm VslmRetrieveVStorageObjectState,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectStateResponse `xml:"urn:vslm VslmRetrieveVStorageObjectStateResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectStateBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObjectState(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObjectState) (*types.VslmRetrieveVStorageObjectStateResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRetrieveVStorageObjectsBody struct {
	Req    *types.VslmRetrieveVStorageObjects         `xml:"urn:vslm VslmRetrieveVStorageObjects,omitempty"`
	Res    *types.VslmRetrieveVStorageObjectsResponse `xml:"urn:vslm VslmRetrieveVStorageObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRetrieveVStorageObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRetrieveVStorageObjects(ctx context.Context, r soap.RoundTripper, req *types.VslmRetrieveVStorageObjects) (*types.VslmRetrieveVStorageObjectsResponse, error) {
	var reqBody, resBody VslmRetrieveVStorageObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmRevertVStorageObject_TaskBody struct {
	Req    *types.VslmRevertVStorageObject_Task         `xml:"urn:vslm VslmRevertVStorageObject_Task,omitempty"`
	Res    *types.VslmRevertVStorageObject_TaskResponse `xml:"urn:vslm VslmRevertVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmRevertVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmRevertVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmRevertVStorageObject_Task) (*types.VslmRevertVStorageObject_TaskResponse, error) {
	var reqBody, resBody VslmRevertVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmScheduleReconcileDatastoreInventoryBody struct {
	Req    *types.VslmScheduleReconcileDatastoreInventory         `xml:"urn:vslm VslmScheduleReconcileDatastoreInventory,omitempty"`
	Res    *types.VslmScheduleReconcileDatastoreInventoryResponse `xml:"urn:vslm VslmScheduleReconcileDatastoreInventoryResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmScheduleReconcileDatastoreInventoryBody) Fault() *soap.Fault { return b.Fault_ }

func VslmScheduleReconcileDatastoreInventory(ctx context.Context, r soap.RoundTripper, req *types.VslmScheduleReconcileDatastoreInventory) (*types.VslmScheduleReconcileDatastoreInventoryResponse, error) {
	var reqBody, resBody VslmScheduleReconcileDatastoreInventoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmSetVStorageObjectControlFlagsBody struct {
	Req    *types.VslmSetVStorageObjectControlFlags         `xml:"urn:vslm VslmSetVStorageObjectControlFlags,omitempty"`
	Res    *types.VslmSetVStorageObjectControlFlagsResponse `xml:"urn:vslm VslmSetVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmSetVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func VslmSetVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.VslmSetVStorageObjectControlFlags) (*types.VslmSetVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody VslmSetVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmSyncDatastoreBody struct {
	Req    *types.VslmSyncDatastore         `xml:"urn:vslm VslmSyncDatastore,omitempty"`
	Res    *types.VslmSyncDatastoreResponse `xml:"urn:vslm VslmSyncDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmSyncDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func VslmSyncDatastore(ctx context.Context, r soap.RoundTripper, req *types.VslmSyncDatastore) (*types.VslmSyncDatastoreResponse, error) {
	var reqBody, resBody VslmSyncDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmUpdateVStorageInfrastructureObjectPolicy_TaskBody struct {
	Req    *types.VslmUpdateVStorageInfrastructureObjectPolicy_Task         `xml:"urn:vslm VslmUpdateVStorageInfrastructureObjectPolicy_Task,omitempty"`
	Res    *types.VslmUpdateVStorageInfrastructureObjectPolicy_TaskResponse `xml:"urn:vslm VslmUpdateVStorageInfrastructureObjectPolicy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmUpdateVStorageInfrastructureObjectPolicy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmUpdateVStorageInfrastructureObjectPolicy_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmUpdateVStorageInfrastructureObjectPolicy_Task) (*types.VslmUpdateVStorageInfrastructureObjectPolicy_TaskResponse, error) {
	var reqBody, resBody VslmUpdateVStorageInfrastructureObjectPolicy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmUpdateVStorageObjectMetadata_TaskBody struct {
	Req    *types.VslmUpdateVStorageObjectMetadata_Task         `xml:"urn:vslm VslmUpdateVStorageObjectMetadata_Task,omitempty"`
	Res    *types.VslmUpdateVStorageObjectMetadata_TaskResponse `xml:"urn:vslm VslmUpdateVStorageObjectMetadata_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmUpdateVStorageObjectMetadata_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmUpdateVStorageObjectMetadata_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmUpdateVStorageObjectMetadata_Task) (*types.VslmUpdateVStorageObjectMetadata_TaskResponse, error) {
	var reqBody, resBody VslmUpdateVStorageObjectMetadata_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VslmUpdateVstorageObjectPolicy_TaskBody struct {
	Req    *types.VslmUpdateVstorageObjectPolicy_Task         `xml:"urn:vslm VslmUpdateVstorageObjectPolicy_Task,omitempty"`
	Res    *types.VslmUpdateVstorageObjectPolicy_TaskResponse `xml:"urn:vslm VslmUpdateVstorageObjectPolicy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VslmUpdateVstorageObjectPolicy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VslmUpdateVstorageObjectPolicy_Task(ctx context.Context, r soap.RoundTripper, req *types.VslmUpdateVstorageObjectPolicy_Task) (*types.VslmUpdateVstorageObjectPolicy_TaskResponse, error) {
	var reqBody, resBody VslmUpdateVstorageObjectPolicy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
