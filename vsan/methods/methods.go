// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vsan/types"
)

type VsanPerfDiagnoseBody struct {
	Req    *types.VsanPerfDiagnose         `xml:"urn:vsan VsanPerfDiagnose,omitempty"`
	Res    *types.VsanPerfDiagnoseResponse `xml:"urn:vsan VsanPerfDiagnoseResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDiagnoseBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDiagnose(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDiagnose) (*types.VsanPerfDiagnoseResponse, error) {
	var reqBody, resBody VsanPerfDiagnoseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryClusterHealthBody struct {
	Req    *types.VsanPerfQueryClusterHealth         `xml:"urn:vsan VsanPerfQueryClusterHealth,omitempty"`
	Res    *types.VsanPerfQueryClusterHealthResponse `xml:"urn:vsan VsanPerfQueryClusterHealthResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryClusterHealthBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryClusterHealth(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryClusterHealth) (*types.VsanPerfQueryClusterHealthResponse, error) {
	var reqBody, resBody VsanPerfQueryClusterHealthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanPerfDiagnosisResultBody struct {
	Req    *types.GetVsanPerfDiagnosisResult         `xml:"urn:vsan GetVsanPerfDiagnosisResult,omitempty"`
	Res    *types.GetVsanPerfDiagnosisResultResponse `xml:"urn:vsan GetVsanPerfDiagnosisResultResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanPerfDiagnosisResultBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanPerfDiagnosisResult(ctx context.Context, r soap.RoundTripper, req *types.GetVsanPerfDiagnosisResult) (*types.GetVsanPerfDiagnosisResultResponse, error) {
	var reqBody, resBody GetVsanPerfDiagnosisResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfToggleVerboseModeBody struct {
	Req    *types.VsanPerfToggleVerboseMode         `xml:"urn:vsan VsanPerfToggleVerboseMode,omitempty"`
	Res    *types.VsanPerfToggleVerboseModeResponse `xml:"urn:vsan VsanPerfToggleVerboseModeResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfToggleVerboseModeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfToggleVerboseMode(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfToggleVerboseMode) (*types.VsanPerfToggleVerboseModeResponse, error) {
	var reqBody, resBody VsanPerfToggleVerboseModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDiagnoseTaskBody struct {
	Req    *types.VsanPerfDiagnoseTask         `xml:"urn:vsan VsanPerfDiagnoseTask,omitempty"`
	Res    *types.VsanPerfDiagnoseTaskResponse `xml:"urn:vsan VsanPerfDiagnoseTaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDiagnoseTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDiagnoseTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDiagnoseTask) (*types.VsanPerfDiagnoseTaskResponse, error) {
	var reqBody, resBody VsanPerfDiagnoseTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryRemoteServerClustersBody struct {
	Req    *types.QueryRemoteServerClusters         `xml:"urn:vsan QueryRemoteServerClusters,omitempty"`
	Res    *types.QueryRemoteServerClustersResponse `xml:"urn:vsan QueryRemoteServerClustersResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryRemoteServerClustersBody) Fault() *soap.Fault { return b.Fault_ }

func QueryRemoteServerClusters(ctx context.Context, r soap.RoundTripper, req *types.QueryRemoteServerClusters) (*types.QueryRemoteServerClustersResponse, error) {
	var reqBody, resBody QueryRemoteServerClustersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryTimeRangesBody struct {
	Req    *types.VsanPerfQueryTimeRanges         `xml:"urn:vsan VsanPerfQueryTimeRanges,omitempty"`
	Res    *types.VsanPerfQueryTimeRangesResponse `xml:"urn:vsan VsanPerfQueryTimeRangesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryTimeRangesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryTimeRanges(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryTimeRanges) (*types.VsanPerfQueryTimeRangesResponse, error) {
	var reqBody, resBody VsanPerfQueryTimeRangesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfSetStatsObjectPolicyBody struct {
	Req    *types.VsanPerfSetStatsObjectPolicy         `xml:"urn:vsan VsanPerfSetStatsObjectPolicy,omitempty"`
	Res    *types.VsanPerfSetStatsObjectPolicyResponse `xml:"urn:vsan VsanPerfSetStatsObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfSetStatsObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfSetStatsObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfSetStatsObjectPolicy) (*types.VsanPerfSetStatsObjectPolicyResponse, error) {
	var reqBody, resBody VsanPerfSetStatsObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfCreateStatsObjectBody struct {
	Req    *types.VsanPerfCreateStatsObject         `xml:"urn:vsan VsanPerfCreateStatsObject,omitempty"`
	Res    *types.VsanPerfCreateStatsObjectResponse `xml:"urn:vsan VsanPerfCreateStatsObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfCreateStatsObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfCreateStatsObject(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfCreateStatsObject) (*types.VsanPerfCreateStatsObjectResponse, error) {
	var reqBody, resBody VsanPerfCreateStatsObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryPerfBody struct {
	Req    *types.VsanPerfQueryPerf         `xml:"urn:vsan VsanPerfQueryPerf,omitempty"`
	Res    *types.VsanPerfQueryPerfResponse `xml:"urn:vsan VsanPerfQueryPerfResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryPerfBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryPerf(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryPerf) (*types.VsanPerfQueryPerfResponse, error) {
	var reqBody, resBody VsanPerfQueryPerfBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfCreateStatsObjectTaskBody struct {
	Req    *types.VsanPerfCreateStatsObjectTask         `xml:"urn:vsan VsanPerfCreateStatsObjectTask,omitempty"`
	Res    *types.VsanPerfCreateStatsObjectTaskResponse `xml:"urn:vsan VsanPerfCreateStatsObjectTaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfCreateStatsObjectTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfCreateStatsObjectTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfCreateStatsObjectTask) (*types.VsanPerfCreateStatsObjectTaskResponse, error) {
	var reqBody, resBody VsanPerfCreateStatsObjectTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteStatsObjectTaskBody struct {
	Req    *types.VsanPerfDeleteStatsObjectTask         `xml:"urn:vsan VsanPerfDeleteStatsObjectTask,omitempty"`
	Res    *types.VsanPerfDeleteStatsObjectTaskResponse `xml:"urn:vsan VsanPerfDeleteStatsObjectTaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteStatsObjectTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteStatsObjectTask(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteStatsObjectTask) (*types.VsanPerfDeleteStatsObjectTaskResponse, error) {
	var reqBody, resBody VsanPerfDeleteStatsObjectTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetAggregatedEntityTypesBody struct {
	Req    *types.VsanPerfGetAggregatedEntityTypes         `xml:"urn:vsan VsanPerfGetAggregatedEntityTypes,omitempty"`
	Res    *types.VsanPerfGetAggregatedEntityTypesResponse `xml:"urn:vsan VsanPerfGetAggregatedEntityTypesResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetAggregatedEntityTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetAggregatedEntityTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetAggregatedEntityTypes) (*types.VsanPerfGetAggregatedEntityTypesResponse, error) {
	var reqBody, resBody VsanPerfGetAggregatedEntityTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteTimeRangeBody struct {
	Req    *types.VsanPerfDeleteTimeRange         `xml:"urn:vsan VsanPerfDeleteTimeRange,omitempty"`
	Res    *types.VsanPerfDeleteTimeRangeResponse `xml:"urn:vsan VsanPerfDeleteTimeRangeResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteTimeRangeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteTimeRange(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteTimeRange) (*types.VsanPerfDeleteTimeRangeResponse, error) {
	var reqBody, resBody VsanPerfDeleteTimeRangeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryStatsObjectInformationBody struct {
	Req    *types.VsanPerfQueryStatsObjectInformation         `xml:"urn:vsan VsanPerfQueryStatsObjectInformation,omitempty"`
	Res    *types.VsanPerfQueryStatsObjectInformationResponse `xml:"urn:vsan VsanPerfQueryStatsObjectInformationResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryStatsObjectInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryStatsObjectInformation(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryStatsObjectInformation) (*types.VsanPerfQueryStatsObjectInformationResponse, error) {
	var reqBody, resBody VsanPerfQueryStatsObjectInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfDeleteStatsObjectBody struct {
	Req    *types.VsanPerfDeleteStatsObject         `xml:"urn:vsan VsanPerfDeleteStatsObject,omitempty"`
	Res    *types.VsanPerfDeleteStatsObjectResponse `xml:"urn:vsan VsanPerfDeleteStatsObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfDeleteStatsObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfDeleteStatsObject(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfDeleteStatsObject) (*types.VsanPerfDeleteStatsObjectResponse, error) {
	var reqBody, resBody VsanPerfDeleteStatsObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfSaveTimeRangesBody struct {
	Req    *types.VsanPerfSaveTimeRanges         `xml:"urn:vsan VsanPerfSaveTimeRanges,omitempty"`
	Res    *types.VsanPerfSaveTimeRangesResponse `xml:"urn:vsan VsanPerfSaveTimeRangesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfSaveTimeRangesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfSaveTimeRanges(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfSaveTimeRanges) (*types.VsanPerfSaveTimeRangesResponse, error) {
	var reqBody, resBody VsanPerfSaveTimeRangesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfQueryNodeInformationBody struct {
	Req    *types.VsanPerfQueryNodeInformation         `xml:"urn:vsan VsanPerfQueryNodeInformation,omitempty"`
	Res    *types.VsanPerfQueryNodeInformationResponse `xml:"urn:vsan VsanPerfQueryNodeInformationResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfQueryNodeInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfQueryNodeInformation(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfQueryNodeInformation) (*types.VsanPerfQueryNodeInformationResponse, error) {
	var reqBody, resBody VsanPerfQueryNodeInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetSupportedDiagnosticExceptionsBody struct {
	Req    *types.VsanPerfGetSupportedDiagnosticExceptions         `xml:"urn:vsan VsanPerfGetSupportedDiagnosticExceptions,omitempty"`
	Res    *types.VsanPerfGetSupportedDiagnosticExceptionsResponse `xml:"urn:vsan VsanPerfGetSupportedDiagnosticExceptionsResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetSupportedDiagnosticExceptionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetSupportedDiagnosticExceptions(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetSupportedDiagnosticExceptions) (*types.VsanPerfGetSupportedDiagnosticExceptionsResponse, error) {
	var reqBody, resBody VsanPerfGetSupportedDiagnosticExceptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerfGetSupportedEntityTypesBody struct {
	Req    *types.VsanPerfGetSupportedEntityTypes         `xml:"urn:vsan VsanPerfGetSupportedEntityTypes,omitempty"`
	Res    *types.VsanPerfGetSupportedEntityTypesResponse `xml:"urn:vsan VsanPerfGetSupportedEntityTypesResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerfGetSupportedEntityTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerfGetSupportedEntityTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanPerfGetSupportedEntityTypes) (*types.VsanPerfGetSupportedEntityTypesResponse, error) {
	var reqBody, resBody VsanPerfGetSupportedEntityTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountDiskMappingExBody struct {
	Req    *types.UnmountDiskMappingEx         `xml:"urn:vsan UnmountDiskMappingEx,omitempty"`
	Res    *types.UnmountDiskMappingExResponse `xml:"urn:vsan UnmountDiskMappingExResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountDiskMappingExBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountDiskMappingEx(ctx context.Context, r soap.RoundTripper, req *types.UnmountDiskMappingEx) (*types.UnmountDiskMappingExResponse, error) {
	var reqBody, resBody UnmountDiskMappingExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitializeDiskMappingsBody struct {
	Req    *types.InitializeDiskMappings         `xml:"urn:vsan InitializeDiskMappings,omitempty"`
	Res    *types.InitializeDiskMappingsResponse `xml:"urn:vsan InitializeDiskMappingsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitializeDiskMappingsBody) Fault() *soap.Fault { return b.Fault_ }

func InitializeDiskMappings(ctx context.Context, r soap.RoundTripper, req *types.InitializeDiskMappings) (*types.InitializeDiskMappingsResponse, error) {
	var reqBody, resBody InitializeDiskMappingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDiskMappingsBody struct {
	Req    *types.QueryDiskMappings         `xml:"urn:vsan QueryDiskMappings,omitempty"`
	Res    *types.QueryDiskMappingsResponse `xml:"urn:vsan QueryDiskMappingsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDiskMappingsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDiskMappings(ctx context.Context, r soap.RoundTripper, req *types.QueryDiskMappings) (*types.QueryDiskMappingsResponse, error) {
	var reqBody, resBody QueryDiskMappingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDiskExBody struct {
	Req    *types.RemoveDiskEx         `xml:"urn:vsan RemoveDiskEx,omitempty"`
	Res    *types.RemoveDiskExResponse `xml:"urn:vsan RemoveDiskExResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDiskExBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDiskEx(ctx context.Context, r soap.RoundTripper, req *types.RemoveDiskEx) (*types.RemoveDiskExResponse, error) {
	var reqBody, resBody RemoveDiskExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryClusterDataEfficiencyCapacityStateBody struct {
	Req    *types.QueryClusterDataEfficiencyCapacityState         `xml:"urn:vsan QueryClusterDataEfficiencyCapacityState,omitempty"`
	Res    *types.QueryClusterDataEfficiencyCapacityStateResponse `xml:"urn:vsan QueryClusterDataEfficiencyCapacityStateResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryClusterDataEfficiencyCapacityStateBody) Fault() *soap.Fault { return b.Fault_ }

func QueryClusterDataEfficiencyCapacityState(ctx context.Context, r soap.RoundTripper, req *types.QueryClusterDataEfficiencyCapacityState) (*types.QueryClusterDataEfficiencyCapacityStateResponse, error) {
	var reqBody, resBody QueryClusterDataEfficiencyCapacityStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAllFlashCapabilitiesBody struct {
	Req    *types.RetrieveAllFlashCapabilities         `xml:"urn:vsan RetrieveAllFlashCapabilities,omitempty"`
	Res    *types.RetrieveAllFlashCapabilitiesResponse `xml:"urn:vsan RetrieveAllFlashCapabilitiesResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAllFlashCapabilitiesBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAllFlashCapabilities(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAllFlashCapabilities) (*types.RetrieveAllFlashCapabilitiesResponse, error) {
	var reqBody, resBody RetrieveAllFlashCapabilitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanManagedDisksBody struct {
	Req    *types.QueryVsanManagedDisks         `xml:"urn:vsan QueryVsanManagedDisks,omitempty"`
	Res    *types.QueryVsanManagedDisksResponse `xml:"urn:vsan QueryVsanManagedDisksResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanManagedDisksBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanManagedDisks(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanManagedDisks) (*types.QueryVsanManagedDisksResponse, error) {
	var reqBody, resBody QueryVsanManagedDisksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RebuildDiskMappingBody struct {
	Req    *types.RebuildDiskMapping         `xml:"urn:vsan RebuildDiskMapping,omitempty"`
	Res    *types.RebuildDiskMappingResponse `xml:"urn:vsan RebuildDiskMappingResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebuildDiskMappingBody) Fault() *soap.Fault { return b.Fault_ }

func RebuildDiskMapping(ctx context.Context, r soap.RoundTripper, req *types.RebuildDiskMapping) (*types.RebuildDiskMappingResponse, error) {
	var reqBody, resBody RebuildDiskMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDiskMappingExBody struct {
	Req    *types.RemoveDiskMappingEx         `xml:"urn:vsan RemoveDiskMappingEx,omitempty"`
	Res    *types.RemoveDiskMappingExResponse `xml:"urn:vsan RemoveDiskMappingExResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDiskMappingExBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDiskMappingEx(ctx context.Context, r soap.RoundTripper, req *types.RemoveDiskMappingEx) (*types.RemoveDiskMappingExResponse, error) {
	var reqBody, resBody RemoveDiskMappingExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCompleteMigrateVmsToVdsBody struct {
	Req    *types.VsanCompleteMigrateVmsToVds         `xml:"urn:vsan VsanCompleteMigrateVmsToVds,omitempty"`
	Res    *types.VsanCompleteMigrateVmsToVdsResponse `xml:"urn:vsan VsanCompleteMigrateVmsToVdsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCompleteMigrateVmsToVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCompleteMigrateVmsToVds(ctx context.Context, r soap.RoundTripper, req *types.VsanCompleteMigrateVmsToVds) (*types.VsanCompleteMigrateVmsToVdsResponse, error) {
	var reqBody, resBody VsanCompleteMigrateVmsToVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanMigrateVmsToVdsBody struct {
	Req    *types.VsanMigrateVmsToVds         `xml:"urn:vsan VsanMigrateVmsToVds,omitempty"`
	Res    *types.VsanMigrateVmsToVdsResponse `xml:"urn:vsan VsanMigrateVmsToVdsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanMigrateVmsToVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanMigrateVmsToVds(ctx context.Context, r soap.RoundTripper, req *types.VsanMigrateVmsToVds) (*types.VsanMigrateVmsToVdsResponse, error) {
	var reqBody, resBody VsanMigrateVmsToVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanCloudHealthStatusBody struct {
	Req    *types.QueryVsanCloudHealthStatus         `xml:"urn:vsan QueryVsanCloudHealthStatus,omitempty"`
	Res    *types.QueryVsanCloudHealthStatusResponse `xml:"urn:vsan QueryVsanCloudHealthStatusResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanCloudHealthStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanCloudHealthStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanCloudHealthStatus) (*types.QueryVsanCloudHealthStatusResponse, error) {
	var reqBody, resBody QueryVsanCloudHealthStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformOnlineHealthCheckBody struct {
	Req    *types.VsanPerformOnlineHealthCheck         `xml:"urn:vsan VsanPerformOnlineHealthCheck,omitempty"`
	Res    *types.VsanPerformOnlineHealthCheckResponse `xml:"urn:vsan VsanPerformOnlineHealthCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformOnlineHealthCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformOnlineHealthCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformOnlineHealthCheck) (*types.VsanPerformOnlineHealthCheckResponse, error) {
	var reqBody, resBody VsanPerformOnlineHealthCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorsFromTargetBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorsFromTarget         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromTarget,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorsFromTargetResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromTargetResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorsFromTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorsFromTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorsFromTarget) (*types.VsanVitRemoveIscsiInitiatorsFromTargetResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorsFromTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorsFromGroupBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorsFromGroup         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorsFromGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorsFromGroupResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorsFromGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorsFromGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorsFromGroup) (*types.VsanVitRemoveIscsiInitiatorsFromGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorsFromGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEditIscsiLUNBody struct {
	Req    *types.VsanVitEditIscsiLUN         `xml:"urn:vsan VsanVitEditIscsiLUN,omitempty"`
	Res    *types.VsanVitEditIscsiLUNResponse `xml:"urn:vsan VsanVitEditIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEditIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEditIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEditIscsiLUN) (*types.VsanVitEditIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitEditIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiLUNBody struct {
	Req    *types.VsanVitGetIscsiLUN         `xml:"urn:vsan VsanVitGetIscsiLUN,omitempty"`
	Res    *types.VsanVitGetIscsiLUNResponse `xml:"urn:vsan VsanVitGetIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiLUN) (*types.VsanVitGetIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitGetIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitEditIscsiTargetBody struct {
	Req    *types.VsanVitEditIscsiTarget         `xml:"urn:vsan VsanVitEditIscsiTarget,omitempty"`
	Res    *types.VsanVitEditIscsiTargetResponse `xml:"urn:vsan VsanVitEditIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitEditIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitEditIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitEditIscsiTarget) (*types.VsanVitEditIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitEditIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorsToGroupBody struct {
	Req    *types.VsanVitAddIscsiInitiatorsToGroup         `xml:"urn:vsan VsanVitAddIscsiInitiatorsToGroup,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorsToGroupResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorsToGroupResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorsToGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorsToGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorsToGroup) (*types.VsanVitAddIscsiInitiatorsToGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorsToGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorsToTargetBody struct {
	Req    *types.VsanVitAddIscsiInitiatorsToTarget         `xml:"urn:vsan VsanVitAddIscsiInitiatorsToTarget,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorsToTargetResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorsToTargetResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorsToTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorsToTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorsToTarget) (*types.VsanVitAddIscsiInitiatorsToTargetResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorsToTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitQueryIscsiTargetServiceVersionBody struct {
	Req    *types.VsanVitQueryIscsiTargetServiceVersion         `xml:"urn:vsan VsanVitQueryIscsiTargetServiceVersion,omitempty"`
	Res    *types.VsanVitQueryIscsiTargetServiceVersionResponse `xml:"urn:vsan VsanVitQueryIscsiTargetServiceVersionResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitQueryIscsiTargetServiceVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitQueryIscsiTargetServiceVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanVitQueryIscsiTargetServiceVersion) (*types.VsanVitQueryIscsiTargetServiceVersionResponse, error) {
	var reqBody, resBody VsanVitQueryIscsiTargetServiceVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiTargetToGroupBody struct {
	Req    *types.VsanVitAddIscsiTargetToGroup         `xml:"urn:vsan VsanVitAddIscsiTargetToGroup,omitempty"`
	Res    *types.VsanVitAddIscsiTargetToGroupResponse `xml:"urn:vsan VsanVitAddIscsiTargetToGroupResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiTargetToGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiTargetToGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiTargetToGroup) (*types.VsanVitAddIscsiTargetToGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiTargetToGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiTargetFromGroupBody struct {
	Req    *types.VsanVitRemoveIscsiTargetFromGroup         `xml:"urn:vsan VsanVitRemoveIscsiTargetFromGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiTargetFromGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiTargetFromGroupResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiTargetFromGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiTargetFromGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiTargetFromGroup) (*types.VsanVitRemoveIscsiTargetFromGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiTargetFromGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiLUNsBody struct {
	Req    *types.VsanVitGetIscsiLUNs         `xml:"urn:vsan VsanVitGetIscsiLUNs,omitempty"`
	Res    *types.VsanVitGetIscsiLUNsResponse `xml:"urn:vsan VsanVitGetIscsiLUNsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiLUNsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiLUNs(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiLUNs) (*types.VsanVitGetIscsiLUNsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiLUNsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiLUNBody struct {
	Req    *types.VsanVitRemoveIscsiLUN         `xml:"urn:vsan VsanVitRemoveIscsiLUN,omitempty"`
	Res    *types.VsanVitRemoveIscsiLUNResponse `xml:"urn:vsan VsanVitRemoveIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiLUN) (*types.VsanVitRemoveIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitGetIscsiInitiatorGroup         `xml:"urn:vsan VsanVitGetIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitGetIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitGetIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiInitiatorGroup) (*types.VsanVitGetIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitGetIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitRemoveIscsiInitiatorGroup         `xml:"urn:vsan VsanVitRemoveIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitRemoveIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitRemoveIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiInitiatorGroup) (*types.VsanVitRemoveIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetHomeObjectBody struct {
	Req    *types.VsanVitGetHomeObject         `xml:"urn:vsan VsanVitGetHomeObject,omitempty"`
	Res    *types.VsanVitGetHomeObjectResponse `xml:"urn:vsan VsanVitGetHomeObjectResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetHomeObjectBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetHomeObject(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetHomeObject) (*types.VsanVitGetHomeObjectResponse, error) {
	var reqBody, resBody VsanVitGetHomeObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetBody struct {
	Req    *types.VsanVitGetIscsiTarget         `xml:"urn:vsan VsanVitGetIscsiTarget,omitempty"`
	Res    *types.VsanVitGetIscsiTargetResponse `xml:"urn:vsan VsanVitGetIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTarget) (*types.VsanVitGetIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitRemoveIscsiTargetBody struct {
	Req    *types.VsanVitRemoveIscsiTarget         `xml:"urn:vsan VsanVitRemoveIscsiTarget,omitempty"`
	Res    *types.VsanVitRemoveIscsiTargetResponse `xml:"urn:vsan VsanVitRemoveIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitRemoveIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitRemoveIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitRemoveIscsiTarget) (*types.VsanVitRemoveIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitRemoveIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiLUNBody struct {
	Req    *types.VsanVitAddIscsiLUN         `xml:"urn:vsan VsanVitAddIscsiLUN,omitempty"`
	Res    *types.VsanVitAddIscsiLUNResponse `xml:"urn:vsan VsanVitAddIscsiLUNResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiLUNBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiLUN(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiLUN) (*types.VsanVitAddIscsiLUNResponse, error) {
	var reqBody, resBody VsanVitAddIscsiLUNBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiInitiatorGroupsBody struct {
	Req    *types.VsanVitGetIscsiInitiatorGroups         `xml:"urn:vsan VsanVitGetIscsiInitiatorGroups,omitempty"`
	Res    *types.VsanVitGetIscsiInitiatorGroupsResponse `xml:"urn:vsan VsanVitGetIscsiInitiatorGroupsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiInitiatorGroupsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiInitiatorGroups(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiInitiatorGroups) (*types.VsanVitGetIscsiInitiatorGroupsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiInitiatorGroupsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitGetIscsiTargetsBody struct {
	Req    *types.VsanVitGetIscsiTargets         `xml:"urn:vsan VsanVitGetIscsiTargets,omitempty"`
	Res    *types.VsanVitGetIscsiTargetsResponse `xml:"urn:vsan VsanVitGetIscsiTargetsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitGetIscsiTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitGetIscsiTargets(ctx context.Context, r soap.RoundTripper, req *types.VsanVitGetIscsiTargets) (*types.VsanVitGetIscsiTargetsResponse, error) {
	var reqBody, resBody VsanVitGetIscsiTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiTargetBody struct {
	Req    *types.VsanVitAddIscsiTarget         `xml:"urn:vsan VsanVitAddIscsiTarget,omitempty"`
	Res    *types.VsanVitAddIscsiTargetResponse `xml:"urn:vsan VsanVitAddIscsiTargetResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiTargetBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiTarget(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiTarget) (*types.VsanVitAddIscsiTargetResponse, error) {
	var reqBody, resBody VsanVitAddIscsiTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVitAddIscsiInitiatorGroupBody struct {
	Req    *types.VsanVitAddIscsiInitiatorGroup         `xml:"urn:vsan VsanVitAddIscsiInitiatorGroup,omitempty"`
	Res    *types.VsanVitAddIscsiInitiatorGroupResponse `xml:"urn:vsan VsanVitAddIscsiInitiatorGroupResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVitAddIscsiInitiatorGroupBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVitAddIscsiInitiatorGroup(ctx context.Context, r soap.RoundTripper, req *types.VsanVitAddIscsiInitiatorGroup) (*types.VsanVitAddIscsiInitiatorGroupResponse, error) {
	var reqBody, resBody VsanVitAddIscsiInitiatorGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateVsanClusterBody struct {
	Req    *types.VsanRemediateVsanCluster         `xml:"urn:vsan VsanRemediateVsanCluster,omitempty"`
	Res    *types.VsanRemediateVsanClusterResponse `xml:"urn:vsan VsanRemediateVsanClusterResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateVsanClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateVsanCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateVsanCluster) (*types.VsanRemediateVsanClusterResponse, error) {
	var reqBody, resBody VsanRemediateVsanClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRemediateVsanHostBody struct {
	Req    *types.VsanRemediateVsanHost         `xml:"urn:vsan VsanRemediateVsanHost,omitempty"`
	Res    *types.VsanRemediateVsanHostResponse `xml:"urn:vsan VsanRemediateVsanHostResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRemediateVsanHostBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRemediateVsanHost(ctx context.Context, r soap.RoundTripper, req *types.VsanRemediateVsanHost) (*types.VsanRemediateVsanHostResponse, error) {
	var reqBody, resBody VsanRemediateVsanHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetCapabilitiesBody struct {
	Req    *types.VsanGetCapabilities         `xml:"urn:vsan VsanGetCapabilities,omitempty"`
	Res    *types.VsanGetCapabilitiesResponse `xml:"urn:vsan VsanGetCapabilitiesResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetCapabilitiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetCapabilities(ctx context.Context, r soap.RoundTripper, req *types.VsanGetCapabilities) (*types.VsanGetCapabilitiesResponse, error) {
	var reqBody, resBody VsanGetCapabilitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostUpdateFirmwareBody struct {
	Req    *types.VsanHostUpdateFirmware         `xml:"urn:vsan VsanHostUpdateFirmware,omitempty"`
	Res    *types.VsanHostUpdateFirmwareResponse `xml:"urn:vsan VsanHostUpdateFirmwareResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostUpdateFirmwareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostUpdateFirmware(ctx context.Context, r soap.RoundTripper, req *types.VsanHostUpdateFirmware) (*types.VsanHostUpdateFirmwareResponse, error) {
	var reqBody, resBody VsanHostUpdateFirmwareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchIsoDepotCookieBody struct {
	Req    *types.FetchIsoDepotCookie         `xml:"urn:vsan FetchIsoDepotCookie,omitempty"`
	Res    *types.FetchIsoDepotCookieResponse `xml:"urn:vsan FetchIsoDepotCookieResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchIsoDepotCookieBody) Fault() *soap.Fault { return b.Fault_ }

func FetchIsoDepotCookie(ctx context.Context, r soap.RoundTripper, req *types.FetchIsoDepotCookie) (*types.FetchIsoDepotCookieResponse, error) {
	var reqBody, resBody FetchIsoDepotCookieBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanVumConfigBody struct {
	Req    *types.GetVsanVumConfig         `xml:"urn:vsan GetVsanVumConfig,omitempty"`
	Res    *types.GetVsanVumConfigResponse `xml:"urn:vsan GetVsanVumConfigResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanVumConfigBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanVumConfig(ctx context.Context, r soap.RoundTripper, req *types.GetVsanVumConfig) (*types.GetVsanVumConfigResponse, error) {
	var reqBody, resBody GetVsanVumConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUploadReleaseDbBody struct {
	Req    *types.VsanVcUploadReleaseDb         `xml:"urn:vsan VsanVcUploadReleaseDb,omitempty"`
	Res    *types.VsanVcUploadReleaseDbResponse `xml:"urn:vsan VsanVcUploadReleaseDbResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUploadReleaseDbBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUploadReleaseDb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUploadReleaseDb) (*types.VsanVcUploadReleaseDbResponse, error) {
	var reqBody, resBody VsanVcUploadReleaseDbBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetResourceCheckStatusBody struct {
	Req    *types.VsanGetResourceCheckStatus         `xml:"urn:vsan VsanGetResourceCheckStatus,omitempty"`
	Res    *types.VsanGetResourceCheckStatusResponse `xml:"urn:vsan VsanGetResourceCheckStatusResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetResourceCheckStatusBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetResourceCheckStatus(ctx context.Context, r soap.RoundTripper, req *types.VsanGetResourceCheckStatus) (*types.VsanGetResourceCheckStatusResponse, error) {
	var reqBody, resBody VsanGetResourceCheckStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformResourceCheckBody struct {
	Req    *types.VsanPerformResourceCheck         `xml:"urn:vsan VsanPerformResourceCheck,omitempty"`
	Res    *types.VsanPerformResourceCheckResponse `xml:"urn:vsan VsanPerformResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformResourceCheck) (*types.VsanPerformResourceCheckResponse, error) {
	var reqBody, resBody VsanPerformResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCancelResourceCheckBody struct {
	Req    *types.VsanHostCancelResourceCheck         `xml:"urn:vsan VsanHostCancelResourceCheck,omitempty"`
	Res    *types.VsanHostCancelResourceCheckResponse `xml:"urn:vsan VsanHostCancelResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCancelResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCancelResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCancelResourceCheck) (*types.VsanHostCancelResourceCheckResponse, error) {
	var reqBody, resBody VsanHostCancelResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostPerformResourceCheckBody struct {
	Req    *types.VsanHostPerformResourceCheck         `xml:"urn:vsan VsanHostPerformResourceCheck,omitempty"`
	Res    *types.VsanHostPerformResourceCheckResponse `xml:"urn:vsan VsanHostPerformResourceCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostPerformResourceCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostPerformResourceCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanHostPerformResourceCheck) (*types.VsanHostPerformResourceCheckResponse, error) {
	var reqBody, resBody VsanHostPerformResourceCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VosSetVsanObjectPolicyBody struct {
	Req    *types.VosSetVsanObjectPolicy         `xml:"urn:vsan VosSetVsanObjectPolicy,omitempty"`
	Res    *types.VosSetVsanObjectPolicyResponse `xml:"urn:vsan VosSetVsanObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VosSetVsanObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func VosSetVsanObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.VosSetVsanObjectPolicy) (*types.VosSetVsanObjectPolicyResponse, error) {
	var reqBody, resBody VosSetVsanObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDeleteObjects_TaskBody struct {
	Req    *types.VsanDeleteObjects_Task         `xml:"urn:vsan VsanDeleteObjects_Task,omitempty"`
	Res    *types.VsanDeleteObjects_TaskResponse `xml:"urn:vsan VsanDeleteObjects_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDeleteObjects_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDeleteObjects_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDeleteObjects_Task) (*types.VsanDeleteObjects_TaskResponse, error) {
	var reqBody, resBody VsanDeleteObjects_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VosQueryVsanObjectInformationBody struct {
	Req    *types.VosQueryVsanObjectInformation         `xml:"urn:vsan VosQueryVsanObjectInformation,omitempty"`
	Res    *types.VosQueryVsanObjectInformationResponse `xml:"urn:vsan VosQueryVsanObjectInformationResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VosQueryVsanObjectInformationBody) Fault() *soap.Fault { return b.Fault_ }

func VosQueryVsanObjectInformation(ctx context.Context, r soap.RoundTripper, req *types.VosQueryVsanObjectInformation) (*types.VosQueryVsanObjectInformationResponse, error) {
	var reqBody, resBody VosQueryVsanObjectInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RelayoutObjectsBody struct {
	Req    *types.RelayoutObjects         `xml:"urn:vsan RelayoutObjects,omitempty"`
	Res    *types.RelayoutObjectsResponse `xml:"urn:vsan RelayoutObjectsResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RelayoutObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func RelayoutObjects(ctx context.Context, r soap.RoundTripper, req *types.RelayoutObjects) (*types.RelayoutObjectsResponse, error) {
	var reqBody, resBody RelayoutObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryInaccessibleVmSwapObjectsBody struct {
	Req    *types.VsanQueryInaccessibleVmSwapObjects         `xml:"urn:vsan VsanQueryInaccessibleVmSwapObjects,omitempty"`
	Res    *types.VsanQueryInaccessibleVmSwapObjectsResponse `xml:"urn:vsan VsanQueryInaccessibleVmSwapObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryInaccessibleVmSwapObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryInaccessibleVmSwapObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryInaccessibleVmSwapObjects) (*types.VsanQueryInaccessibleVmSwapObjectsResponse, error) {
	var reqBody, resBody VsanQueryInaccessibleVmSwapObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySyncingVsanObjectsSummaryBody struct {
	Req    *types.QuerySyncingVsanObjectsSummary         `xml:"urn:vsan QuerySyncingVsanObjectsSummary,omitempty"`
	Res    *types.QuerySyncingVsanObjectsSummaryResponse `xml:"urn:vsan QuerySyncingVsanObjectsSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySyncingVsanObjectsSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySyncingVsanObjectsSummary(ctx context.Context, r soap.RoundTripper, req *types.QuerySyncingVsanObjectsSummary) (*types.QuerySyncingVsanObjectsSummaryResponse, error) {
	var reqBody, resBody QuerySyncingVsanObjectsSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryObjectIdentitiesBody struct {
	Req    *types.VsanQueryObjectIdentities         `xml:"urn:vsan VsanQueryObjectIdentities,omitempty"`
	Res    *types.VsanQueryObjectIdentitiesResponse `xml:"urn:vsan VsanQueryObjectIdentitiesResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryObjectIdentitiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryObjectIdentities(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryObjectIdentities) (*types.VsanQueryObjectIdentitiesResponse, error) {
	var reqBody, resBody VsanQueryObjectIdentitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanValidateConfigSpecBody struct {
	Req    *types.VsanValidateConfigSpec         `xml:"urn:vsan VsanValidateConfigSpec,omitempty"`
	Res    *types.VsanValidateConfigSpecResponse `xml:"urn:vsan VsanValidateConfigSpecResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanValidateConfigSpecBody) Fault() *soap.Fault { return b.Fault_ }

func VsanValidateConfigSpec(ctx context.Context, r soap.RoundTripper, req *types.VsanValidateConfigSpec) (*types.VsanValidateConfigSpecResponse, error) {
	var reqBody, resBody VsanValidateConfigSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterReconfigBody struct {
	Req    *types.VsanClusterReconfig         `xml:"urn:vsan VsanClusterReconfig,omitempty"`
	Res    *types.VsanClusterReconfigResponse `xml:"urn:vsan VsanClusterReconfigResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterReconfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterReconfig(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterReconfig) (*types.VsanClusterReconfigResponse, error) {
	var reqBody, resBody VsanClusterReconfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetRuntimeStatsBody struct {
	Req    *types.VsanClusterGetRuntimeStats         `xml:"urn:vsan VsanClusterGetRuntimeStats,omitempty"`
	Res    *types.VsanClusterGetRuntimeStatsResponse `xml:"urn:vsan VsanClusterGetRuntimeStatsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetRuntimeStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetRuntimeStats(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetRuntimeStats) (*types.VsanClusterGetRuntimeStatsResponse, error) {
	var reqBody, resBody VsanClusterGetRuntimeStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterDrsStatsBody struct {
	Req    *types.VsanQueryClusterDrsStats         `xml:"urn:vsan VsanQueryClusterDrsStats,omitempty"`
	Res    *types.VsanQueryClusterDrsStatsResponse `xml:"urn:vsan VsanQueryClusterDrsStatsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterDrsStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterDrsStats(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterDrsStats) (*types.VsanQueryClusterDrsStatsResponse, error) {
	var reqBody, resBody VsanQueryClusterDrsStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetConfigBody struct {
	Req    *types.VsanClusterGetConfig         `xml:"urn:vsan VsanClusterGetConfig,omitempty"`
	Res    *types.VsanClusterGetConfigResponse `xml:"urn:vsan VsanClusterGetConfigResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetConfig) (*types.VsanClusterGetConfigResponse, error) {
	var reqBody, resBody VsanClusterGetConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanEncryptedClusterRekey_TaskBody struct {
	Req    *types.VsanEncryptedClusterRekey_Task         `xml:"urn:vsan VsanEncryptedClusterRekey_Task,omitempty"`
	Res    *types.VsanEncryptedClusterRekey_TaskResponse `xml:"urn:vsan VsanEncryptedClusterRekey_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanEncryptedClusterRekey_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanEncryptedClusterRekey_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanEncryptedClusterRekey_Task) (*types.VsanEncryptedClusterRekey_TaskResponse, error) {
	var reqBody, resBody VsanEncryptedClusterRekey_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPurgeHclFilesBody struct {
	Req    *types.VsanPurgeHclFiles         `xml:"urn:vsan VsanPurgeHclFiles,omitempty"`
	Res    *types.VsanPurgeHclFilesResponse `xml:"urn:vsan VsanPurgeHclFilesResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPurgeHclFilesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPurgeHclFiles(ctx context.Context, r soap.RoundTripper, req *types.VsanPurgeHclFiles) (*types.VsanPurgeHclFilesResponse, error) {
	var reqBody, resBody VsanPurgeHclFilesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterCreateVmHealthHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterCreateVmHealthHistoryTest         `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterCreateVmHealthHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterCreateVmHealthHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterCreateVmHealthHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterCreateVmHealthHistoryTest) (*types.VsanQueryVcClusterCreateVmHealthHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterCreateVmHealthHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterObjExtAttrsBody struct {
	Req    *types.VsanQueryVcClusterObjExtAttrs         `xml:"urn:vsan VsanQueryVcClusterObjExtAttrs,omitempty"`
	Res    *types.VsanQueryVcClusterObjExtAttrsResponse `xml:"urn:vsan VsanQueryVcClusterObjExtAttrsResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterObjExtAttrsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterObjExtAttrs(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterObjExtAttrs) (*types.VsanQueryVcClusterObjExtAttrsResponse, error) {
	var reqBody, resBody VsanQueryVcClusterObjExtAttrsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetLogLevelBody struct {
	Req    *types.VsanHealthSetLogLevel         `xml:"urn:vsan VsanHealthSetLogLevel,omitempty"`
	Res    *types.VsanHealthSetLogLevelResponse `xml:"urn:vsan VsanHealthSetLogLevelResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetLogLevelBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetLogLevel(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetLogLevel) (*types.VsanHealthSetLogLevelResponse, error) {
	var reqBody, resBody VsanHealthSetLogLevelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthTestVsanClusterTelemetryProxyBody struct {
	Req    *types.VsanHealthTestVsanClusterTelemetryProxy         `xml:"urn:vsan VsanHealthTestVsanClusterTelemetryProxy,omitempty"`
	Res    *types.VsanHealthTestVsanClusterTelemetryProxyResponse `xml:"urn:vsan VsanHealthTestVsanClusterTelemetryProxyResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthTestVsanClusterTelemetryProxyBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthTestVsanClusterTelemetryProxy(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthTestVsanClusterTelemetryProxy) (*types.VsanHealthTestVsanClusterTelemetryProxyResponse, error) {
	var reqBody, resBody VsanHealthTestVsanClusterTelemetryProxyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUploadHclDbBody struct {
	Req    *types.VsanVcUploadHclDb         `xml:"urn:vsan VsanVcUploadHclDb,omitempty"`
	Res    *types.VsanVcUploadHclDbResponse `xml:"urn:vsan VsanVcUploadHclDbResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUploadHclDbBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUploadHclDb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUploadHclDb) (*types.VsanVcUploadHclDbResponse, error) {
	var reqBody, resBody VsanVcUploadHclDbBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterSmartStatsSummaryBody struct {
	Req    *types.VsanQueryVcClusterSmartStatsSummary         `xml:"urn:vsan VsanQueryVcClusterSmartStatsSummary,omitempty"`
	Res    *types.VsanQueryVcClusterSmartStatsSummaryResponse `xml:"urn:vsan VsanQueryVcClusterSmartStatsSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterSmartStatsSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterSmartStatsSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterSmartStatsSummary) (*types.VsanQueryVcClusterSmartStatsSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterSmartStatsSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcUpdateHclDbFromWebBody struct {
	Req    *types.VsanVcUpdateHclDbFromWeb         `xml:"urn:vsan VsanVcUpdateHclDbFromWeb,omitempty"`
	Res    *types.VsanVcUpdateHclDbFromWebResponse `xml:"urn:vsan VsanVcUpdateHclDbFromWebResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcUpdateHclDbFromWebBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcUpdateHclDbFromWeb(ctx context.Context, r soap.RoundTripper, req *types.VsanVcUpdateHclDbFromWeb) (*types.VsanVcUpdateHclDbFromWebResponse, error) {
	var reqBody, resBody VsanVcUpdateHclDbFromWebBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStopRebalanceClusterBody struct {
	Req    *types.VsanStopRebalanceCluster         `xml:"urn:vsan VsanStopRebalanceCluster,omitempty"`
	Res    *types.VsanStopRebalanceClusterResponse `xml:"urn:vsan VsanStopRebalanceClusterResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStopRebalanceClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStopRebalanceCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanStopRebalanceCluster) (*types.VsanStopRebalanceClusterResponse, error) {
	var reqBody, resBody VsanStopRebalanceClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthGetVsanClusterSilentChecksBody struct {
	Req    *types.VsanHealthGetVsanClusterSilentChecks         `xml:"urn:vsan VsanHealthGetVsanClusterSilentChecks,omitempty"`
	Res    *types.VsanHealthGetVsanClusterSilentChecksResponse `xml:"urn:vsan VsanHealthGetVsanClusterSilentChecksResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthGetVsanClusterSilentChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthGetVsanClusterSilentChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthGetVsanClusterSilentChecks) (*types.VsanHealthGetVsanClusterSilentChecksResponse, error) {
	var reqBody, resBody VsanHealthGetVsanClusterSilentChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFileServiceHealthSummaryBody struct {
	Req    *types.VsanClusterQueryFileServiceHealthSummary         `xml:"urn:vsan VsanClusterQueryFileServiceHealthSummary,omitempty"`
	Res    *types.VsanClusterQueryFileServiceHealthSummaryResponse `xml:"urn:vsan VsanClusterQueryFileServiceHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFileServiceHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFileServiceHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFileServiceHealthSummary) (*types.VsanClusterQueryFileServiceHealthSummaryResponse, error) {
	var reqBody, resBody VsanClusterQueryFileServiceHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthRepairClusterObjectsImmediateBody struct {
	Req    *types.VsanHealthRepairClusterObjectsImmediate         `xml:"urn:vsan VsanHealthRepairClusterObjectsImmediate,omitempty"`
	Res    *types.VsanHealthRepairClusterObjectsImmediateResponse `xml:"urn:vsan VsanHealthRepairClusterObjectsImmediateResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthRepairClusterObjectsImmediateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthRepairClusterObjectsImmediate(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthRepairClusterObjectsImmediate) (*types.VsanHealthRepairClusterObjectsImmediateResponse, error) {
	var reqBody, resBody VsanHealthRepairClusterObjectsImmediateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterNetworkPerfTestBody struct {
	Req    *types.VsanQueryVcClusterNetworkPerfTest         `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTest,omitempty"`
	Res    *types.VsanQueryVcClusterNetworkPerfTestResponse `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTestResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterNetworkPerfTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterNetworkPerfTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterNetworkPerfTest) (*types.VsanQueryVcClusterNetworkPerfTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterNetworkPerfTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterVmdkLoadHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterVmdkLoadHistoryTest         `xml:"urn:vsan VsanQueryVcClusterVmdkLoadHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterVmdkLoadHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterVmdkLoadHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterVmdkLoadHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterVmdkLoadHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterVmdkLoadHistoryTest) (*types.VsanQueryVcClusterVmdkLoadHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterVmdkLoadHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthIsRebalanceRunningBody struct {
	Req    *types.VsanHealthIsRebalanceRunning         `xml:"urn:vsan VsanHealthIsRebalanceRunning,omitempty"`
	Res    *types.VsanHealthIsRebalanceRunningResponse `xml:"urn:vsan VsanHealthIsRebalanceRunningResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthIsRebalanceRunningBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthIsRebalanceRunning(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthIsRebalanceRunning) (*types.VsanHealthIsRebalanceRunningResponse, error) {
	var reqBody, resBody VsanHealthIsRebalanceRunningBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterCreateVmHealthTestBody struct {
	Req    *types.VsanQueryVcClusterCreateVmHealthTest         `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthTest,omitempty"`
	Res    *types.VsanQueryVcClusterCreateVmHealthTestResponse `xml:"urn:vsan VsanQueryVcClusterCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterCreateVmHealthTest) (*types.VsanQueryVcClusterCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanProxyConfigBody struct {
	Req    *types.VsanHealthQueryVsanProxyConfig         `xml:"urn:vsan VsanHealthQueryVsanProxyConfig,omitempty"`
	Res    *types.VsanHealthQueryVsanProxyConfigResponse `xml:"urn:vsan VsanHealthQueryVsanProxyConfigResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanProxyConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanProxyConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanProxyConfig) (*types.VsanHealthQueryVsanProxyConfigResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanProxyConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanClusterHealthCheckIntervalBody struct {
	Req    *types.VsanHealthQueryVsanClusterHealthCheckInterval         `xml:"urn:vsan VsanHealthQueryVsanClusterHealthCheckInterval,omitempty"`
	Res    *types.VsanHealthQueryVsanClusterHealthCheckIntervalResponse `xml:"urn:vsan VsanHealthQueryVsanClusterHealthCheckIntervalResponse,omitempty"`
	Fault_ *soap.Fault                                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanClusterHealthCheckIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanClusterHealthCheckInterval(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanClusterHealthCheckInterval) (*types.VsanHealthQueryVsanClusterHealthCheckIntervalResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanClusterHealthCheckIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryAllSupportedHealthChecksBody struct {
	Req    *types.VsanQueryAllSupportedHealthChecks         `xml:"urn:vsan VsanQueryAllSupportedHealthChecks,omitempty"`
	Res    *types.VsanQueryAllSupportedHealthChecksResponse `xml:"urn:vsan VsanQueryAllSupportedHealthChecksResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryAllSupportedHealthChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryAllSupportedHealthChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryAllSupportedHealthChecks) (*types.VsanQueryAllSupportedHealthChecksResponse, error) {
	var reqBody, resBody VsanQueryAllSupportedHealthChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterGetHclInfoBody struct {
	Req    *types.VsanVcClusterGetHclInfo         `xml:"urn:vsan VsanVcClusterGetHclInfo,omitempty"`
	Res    *types.VsanVcClusterGetHclInfoResponse `xml:"urn:vsan VsanVcClusterGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterGetHclInfo) (*types.VsanVcClusterGetHclInfoResponse, error) {
	var reqBody, resBody VsanVcClusterGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryAttachToSrHistoryBody struct {
	Req    *types.VsanQueryAttachToSrHistory         `xml:"urn:vsan VsanQueryAttachToSrHistory,omitempty"`
	Res    *types.VsanQueryAttachToSrHistoryResponse `xml:"urn:vsan VsanQueryAttachToSrHistoryResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryAttachToSrHistoryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryAttachToSrHistory(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryAttachToSrHistory) (*types.VsanQueryAttachToSrHistoryResponse, error) {
	var reqBody, resBody VsanQueryAttachToSrHistoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetReleaseRecommendationBody struct {
	Req    *types.VsanGetReleaseRecommendation         `xml:"urn:vsan VsanGetReleaseRecommendation,omitempty"`
	Res    *types.VsanGetReleaseRecommendationResponse `xml:"urn:vsan VsanGetReleaseRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetReleaseRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetReleaseRecommendation(ctx context.Context, r soap.RoundTripper, req *types.VsanGetReleaseRecommendation) (*types.VsanGetReleaseRecommendationResponse, error) {
	var reqBody, resBody VsanGetReleaseRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetHclConstraintsBody struct {
	Req    *types.VsanGetHclConstraints         `xml:"urn:vsan VsanGetHclConstraints,omitempty"`
	Res    *types.VsanGetHclConstraintsResponse `xml:"urn:vsan VsanGetHclConstraintsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetHclConstraintsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetHclConstraints(ctx context.Context, r soap.RoundTripper, req *types.VsanGetHclConstraints) (*types.VsanGetHclConstraintsResponse, error) {
	var reqBody, resBody VsanGetHclConstraintsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRebalanceClusterBody struct {
	Req    *types.VsanRebalanceCluster         `xml:"urn:vsan VsanRebalanceCluster,omitempty"`
	Res    *types.VsanRebalanceClusterResponse `xml:"urn:vsan VsanRebalanceClusterResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRebalanceClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRebalanceCluster(ctx context.Context, r soap.RoundTripper, req *types.VsanRebalanceCluster) (*types.VsanRebalanceClusterResponse, error) {
	var reqBody, resBody VsanRebalanceClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterRunVmdkLoadTestBody struct {
	Req    *types.VsanVcClusterRunVmdkLoadTest         `xml:"urn:vsan VsanVcClusterRunVmdkLoadTest,omitempty"`
	Res    *types.VsanVcClusterRunVmdkLoadTestResponse `xml:"urn:vsan VsanVcClusterRunVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterRunVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterRunVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterRunVmdkLoadTest) (*types.VsanVcClusterRunVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanVcClusterRunVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSendVsanTelemetryBody struct {
	Req    *types.VsanHealthSendVsanTelemetry         `xml:"urn:vsan VsanHealthSendVsanTelemetry,omitempty"`
	Res    *types.VsanHealthSendVsanTelemetryResponse `xml:"urn:vsan VsanHealthSendVsanTelemetryResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSendVsanTelemetryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSendVsanTelemetry(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSendVsanTelemetry) (*types.VsanHealthSendVsanTelemetryResponse, error) {
	var reqBody, resBody VsanHealthSendVsanTelemetryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterNetworkPerfHistoryTestBody struct {
	Req    *types.VsanQueryVcClusterNetworkPerfHistoryTest         `xml:"urn:vsan VsanQueryVcClusterNetworkPerfHistoryTest,omitempty"`
	Res    *types.VsanQueryVcClusterNetworkPerfHistoryTestResponse `xml:"urn:vsan VsanQueryVcClusterNetworkPerfHistoryTestResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterNetworkPerfHistoryTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterNetworkPerfHistoryTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterNetworkPerfHistoryTest) (*types.VsanQueryVcClusterNetworkPerfHistoryTestResponse, error) {
	var reqBody, resBody VsanQueryVcClusterNetworkPerfHistoryTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterHealthSummaryBody struct {
	Req    *types.VsanQueryVcClusterHealthSummary         `xml:"urn:vsan VsanQueryVcClusterHealthSummary,omitempty"`
	Res    *types.VsanQueryVcClusterHealthSummaryResponse `xml:"urn:vsan VsanQueryVcClusterHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterHealthSummary) (*types.VsanQueryVcClusterHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryVcClusterHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterHealthSummaryTaskBody struct {
	Req    *types.VsanQueryVcClusterHealthSummaryTask         `xml:"urn:vsan VsanQueryVcClusterHealthSummaryTask,omitempty"`
	Res    *types.VsanQueryVcClusterHealthSummaryTaskResponse `xml:"urn:vsan VsanQueryVcClusterHealthSummaryTaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterHealthSummaryTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterHealthSummaryTask(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterHealthSummaryTask) (*types.VsanQueryVcClusterHealthSummaryTaskResponse, error) {
	var reqBody, resBody VsanQueryVcClusterHealthSummaryTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterNetworkPerfTaskBody struct {
	Req    *types.VsanQueryVcClusterNetworkPerfTask         `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTask,omitempty"`
	Res    *types.VsanQueryVcClusterNetworkPerfTaskResponse `xml:"urn:vsan VsanQueryVcClusterNetworkPerfTaskResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterNetworkPerfTaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterNetworkPerfTask(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterNetworkPerfTask) (*types.VsanQueryVcClusterNetworkPerfTaskResponse, error) {
	var reqBody, resBody VsanQueryVcClusterNetworkPerfTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthQueryVsanClusterHealthConfigBody struct {
	Req    *types.VsanHealthQueryVsanClusterHealthConfig         `xml:"urn:vsan VsanHealthQueryVsanClusterHealthConfig,omitempty"`
	Res    *types.VsanHealthQueryVsanClusterHealthConfigResponse `xml:"urn:vsan VsanHealthQueryVsanClusterHealthConfigResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthQueryVsanClusterHealthConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthQueryVsanClusterHealthConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthQueryVsanClusterHealthConfig) (*types.VsanHealthQueryVsanClusterHealthConfigResponse, error) {
	var reqBody, resBody VsanHealthQueryVsanClusterHealthConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanAttachVsanSupportBundleToSrBody struct {
	Req    *types.VsanAttachVsanSupportBundleToSr         `xml:"urn:vsan VsanAttachVsanSupportBundleToSr,omitempty"`
	Res    *types.VsanAttachVsanSupportBundleToSrResponse `xml:"urn:vsan VsanAttachVsanSupportBundleToSrResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanAttachVsanSupportBundleToSrBody) Fault() *soap.Fault { return b.Fault_ }

func VsanAttachVsanSupportBundleToSr(ctx context.Context, r soap.RoundTripper, req *types.VsanAttachVsanSupportBundleToSr) (*types.VsanAttachVsanSupportBundleToSrResponse, error) {
	var reqBody, resBody VsanAttachVsanSupportBundleToSrBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadHclFile_TaskBody struct {
	Req    *types.VsanDownloadHclFile_Task         `xml:"urn:vsan VsanDownloadHclFile_Task,omitempty"`
	Res    *types.VsanDownloadHclFile_TaskResponse `xml:"urn:vsan VsanDownloadHclFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadHclFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadHclFile_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadHclFile_Task) (*types.VsanDownloadHclFile_TaskResponse, error) {
	var reqBody, resBody VsanDownloadHclFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVcClusterVmdkWorkloadTypesBody struct {
	Req    *types.VsanQueryVcClusterVmdkWorkloadTypes         `xml:"urn:vsan VsanQueryVcClusterVmdkWorkloadTypes,omitempty"`
	Res    *types.VsanQueryVcClusterVmdkWorkloadTypesResponse `xml:"urn:vsan VsanQueryVcClusterVmdkWorkloadTypesResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVcClusterVmdkWorkloadTypesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVcClusterVmdkWorkloadTypes(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVcClusterVmdkWorkloadTypes) (*types.VsanQueryVcClusterVmdkWorkloadTypesResponse, error) {
	var reqBody, resBody VsanQueryVcClusterVmdkWorkloadTypesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterSilentChecksBody struct {
	Req    *types.VsanHealthSetVsanClusterSilentChecks         `xml:"urn:vsan VsanHealthSetVsanClusterSilentChecks,omitempty"`
	Res    *types.VsanHealthSetVsanClusterSilentChecksResponse `xml:"urn:vsan VsanHealthSetVsanClusterSilentChecksResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterSilentChecksBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterSilentChecks(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterSilentChecks) (*types.VsanHealthSetVsanClusterSilentChecksResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterSilentChecksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcClusterQueryVerifyHealthSystemVersionsBody struct {
	Req    *types.VsanVcClusterQueryVerifyHealthSystemVersions         `xml:"urn:vsan VsanVcClusterQueryVerifyHealthSystemVersions,omitempty"`
	Res    *types.VsanVcClusterQueryVerifyHealthSystemVersionsResponse `xml:"urn:vsan VsanVcClusterQueryVerifyHealthSystemVersionsResponse,omitempty"`
	Fault_ *soap.Fault                                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcClusterQueryVerifyHealthSystemVersionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcClusterQueryVerifyHealthSystemVersions(ctx context.Context, r soap.RoundTripper, req *types.VsanVcClusterQueryVerifyHealthSystemVersions) (*types.VsanVcClusterQueryVerifyHealthSystemVersionsResponse, error) {
	var reqBody, resBody VsanVcClusterQueryVerifyHealthSystemVersionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterTelemetryConfigBody struct {
	Req    *types.VsanHealthSetVsanClusterTelemetryConfig         `xml:"urn:vsan VsanHealthSetVsanClusterTelemetryConfig,omitempty"`
	Res    *types.VsanHealthSetVsanClusterTelemetryConfigResponse `xml:"urn:vsan VsanHealthSetVsanClusterTelemetryConfigResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterTelemetryConfigBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterTelemetryConfig(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterTelemetryConfig) (*types.VsanHealthSetVsanClusterTelemetryConfigResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterTelemetryConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadAndInstallVendorTool_TaskBody struct {
	Req    *types.VsanDownloadAndInstallVendorTool_Task         `xml:"urn:vsan VsanDownloadAndInstallVendorTool_Task,omitempty"`
	Res    *types.VsanDownloadAndInstallVendorTool_TaskResponse `xml:"urn:vsan VsanDownloadAndInstallVendorTool_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadAndInstallVendorTool_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadAndInstallVendorTool_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadAndInstallVendorTool_Task) (*types.VsanDownloadAndInstallVendorTool_TaskResponse, error) {
	var reqBody, resBody VsanDownloadAndInstallVendorTool_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHealthSetVsanClusterHealthCheckIntervalBody struct {
	Req    *types.VsanHealthSetVsanClusterHealthCheckInterval         `xml:"urn:vsan VsanHealthSetVsanClusterHealthCheckInterval,omitempty"`
	Res    *types.VsanHealthSetVsanClusterHealthCheckIntervalResponse `xml:"urn:vsan VsanHealthSetVsanClusterHealthCheckIntervalResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHealthSetVsanClusterHealthCheckIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHealthSetVsanClusterHealthCheckInterval(ctx context.Context, r soap.RoundTripper, req *types.VsanHealthSetVsanClusterHealthCheckInterval) (*types.VsanHealthSetVsanClusterHealthCheckIntervalResponse, error) {
	var reqBody, resBody VsanHealthSetVsanClusterHealthCheckIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPostConfigForVcsaBody struct {
	Req    *types.VsanPostConfigForVcsa         `xml:"urn:vsan VsanPostConfigForVcsa,omitempty"`
	Res    *types.VsanPostConfigForVcsaResponse `xml:"urn:vsan VsanPostConfigForVcsaResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPostConfigForVcsaBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPostConfigForVcsa(ctx context.Context, r soap.RoundTripper, req *types.VsanPostConfigForVcsa) (*types.VsanPostConfigForVcsaResponse, error) {
	var reqBody, resBody VsanPostConfigForVcsaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcsaGetBootstrapProgressBody struct {
	Req    *types.VsanVcsaGetBootstrapProgress         `xml:"urn:vsan VsanVcsaGetBootstrapProgress,omitempty"`
	Res    *types.VsanVcsaGetBootstrapProgressResponse `xml:"urn:vsan VsanVcsaGetBootstrapProgressResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcsaGetBootstrapProgressBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcsaGetBootstrapProgress(ctx context.Context, r soap.RoundTripper, req *types.VsanVcsaGetBootstrapProgress) (*types.VsanVcsaGetBootstrapProgressResponse, error) {
	var reqBody, resBody VsanVcsaGetBootstrapProgressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPrepareVsanForVcsaBody struct {
	Req    *types.VsanPrepareVsanForVcsa         `xml:"urn:vsan VsanPrepareVsanForVcsa,omitempty"`
	Res    *types.VsanPrepareVsanForVcsaResponse `xml:"urn:vsan VsanPrepareVsanForVcsaResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPrepareVsanForVcsaBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPrepareVsanForVcsa(ctx context.Context, r soap.RoundTripper, req *types.VsanPrepareVsanForVcsa) (*types.VsanPrepareVsanForVcsaResponse, error) {
	var reqBody, resBody VsanPrepareVsanForVcsaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVdsMigrateVssBody struct {
	Req    *types.VsanVdsMigrateVss         `xml:"urn:vsan VsanVdsMigrateVss,omitempty"`
	Res    *types.VsanVdsMigrateVssResponse `xml:"urn:vsan VsanVdsMigrateVssResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVdsMigrateVssBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVdsMigrateVss(ctx context.Context, r soap.RoundTripper, req *types.VsanVdsMigrateVss) (*types.VsanVdsMigrateVssResponse, error) {
	var reqBody, resBody VsanVdsMigrateVssBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVdsGetMigrationPlanBody struct {
	Req    *types.VsanVdsGetMigrationPlan         `xml:"urn:vsan VsanVdsGetMigrationPlan,omitempty"`
	Res    *types.VsanVdsGetMigrationPlanResponse `xml:"urn:vsan VsanVdsGetMigrationPlanResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVdsGetMigrationPlanBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVdsGetMigrationPlan(ctx context.Context, r soap.RoundTripper, req *types.VsanVdsGetMigrationPlan) (*types.VsanVdsGetMigrationPlanResponse, error) {
	var reqBody, resBody VsanVdsGetMigrationPlanBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVssMigrateVdsBody struct {
	Req    *types.VsanVssMigrateVds         `xml:"urn:vsan VsanVssMigrateVds,omitempty"`
	Res    *types.VsanVssMigrateVdsResponse `xml:"urn:vsan VsanVssMigrateVdsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVssMigrateVdsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVssMigrateVds(ctx context.Context, r soap.RoundTripper, req *types.VsanVssMigrateVds) (*types.VsanVssMigrateVdsResponse, error) {
	var reqBody, resBody VsanVssMigrateVdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRollbackVdsToVssBody struct {
	Req    *types.VsanRollbackVdsToVss         `xml:"urn:vsan VsanRollbackVdsToVss,omitempty"`
	Res    *types.VsanRollbackVdsToVssResponse `xml:"urn:vsan VsanRollbackVdsToVssResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRollbackVdsToVssBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRollbackVdsToVss(ctx context.Context, r soap.RoundTripper, req *types.VsanRollbackVdsToVss) (*types.VsanRollbackVdsToVssResponse, error) {
	var reqBody, resBody VsanRollbackVdsToVssBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MountPrecheckBody struct {
	Req    *types.MountPrecheck         `xml:"urn:vsan MountPrecheck,omitempty"`
	Res    *types.MountPrecheckResponse `xml:"urn:vsan MountPrecheckResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountPrecheckBody) Fault() *soap.Fault { return b.Fault_ }

func MountPrecheck(ctx context.Context, r soap.RoundTripper, req *types.MountPrecheck) (*types.MountPrecheckResponse, error) {
	var reqBody, resBody MountPrecheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradeExBody struct {
	Req    *types.PerformVsanUpgradeEx         `xml:"urn:vsan PerformVsanUpgradeEx,omitempty"`
	Res    *types.PerformVsanUpgradeExResponse `xml:"urn:vsan PerformVsanUpgradeExResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradeExBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradeEx(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradeEx) (*types.PerformVsanUpgradeExResponse, error) {
	var reqBody, resBody PerformVsanUpgradeExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryUpgradeStatusExBody struct {
	Req    *types.VsanQueryUpgradeStatusEx         `xml:"urn:vsan VsanQueryUpgradeStatusEx,omitempty"`
	Res    *types.VsanQueryUpgradeStatusExResponse `xml:"urn:vsan VsanQueryUpgradeStatusExResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryUpgradeStatusExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryUpgradeStatusEx(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryUpgradeStatusEx) (*types.VsanQueryUpgradeStatusExResponse, error) {
	var reqBody, resBody VsanQueryUpgradeStatusExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveSupportedVsanFormatVersionBody struct {
	Req    *types.RetrieveSupportedVsanFormatVersion         `xml:"urn:vsan RetrieveSupportedVsanFormatVersion,omitempty"`
	Res    *types.RetrieveSupportedVsanFormatVersionResponse `xml:"urn:vsan RetrieveSupportedVsanFormatVersionResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveSupportedVsanFormatVersionBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveSupportedVsanFormatVersion(ctx context.Context, r soap.RoundTripper, req *types.RetrieveSupportedVsanFormatVersion) (*types.RetrieveSupportedVsanFormatVersionResponse, error) {
	var reqBody, resBody RetrieveSupportedVsanFormatVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradePreflightCheckExBody struct {
	Req    *types.PerformVsanUpgradePreflightCheckEx         `xml:"urn:vsan PerformVsanUpgradePreflightCheckEx,omitempty"`
	Res    *types.PerformVsanUpgradePreflightCheckExResponse `xml:"urn:vsan PerformVsanUpgradePreflightCheckExResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradePreflightCheckExBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradePreflightCheckEx(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradePreflightCheckEx) (*types.PerformVsanUpgradePreflightCheckExResponse, error) {
	var reqBody, resBody PerformVsanUpgradePreflightCheckExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradePreflightAsyncCheck_TaskBody struct {
	Req    *types.PerformVsanUpgradePreflightAsyncCheck_Task         `xml:"urn:vsan PerformVsanUpgradePreflightAsyncCheck_Task,omitempty"`
	Res    *types.PerformVsanUpgradePreflightAsyncCheck_TaskResponse `xml:"urn:vsan PerformVsanUpgradePreflightAsyncCheck_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradePreflightAsyncCheck_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradePreflightAsyncCheck_Task(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradePreflightAsyncCheck_Task) (*types.PerformVsanUpgradePreflightAsyncCheck_TaskResponse, error) {
	var reqBody, resBody PerformVsanUpgradePreflightAsyncCheck_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQuerySpaceUsageBody struct {
	Req    *types.VsanQuerySpaceUsage         `xml:"urn:vsan VsanQuerySpaceUsage,omitempty"`
	Res    *types.VsanQuerySpaceUsageResponse `xml:"urn:vsan VsanQuerySpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQuerySpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQuerySpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQuerySpaceUsage) (*types.VsanQuerySpaceUsageResponse, error) {
	var reqBody, resBody VsanQuerySpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryEntitySpaceUsageBody struct {
	Req    *types.VsanQueryEntitySpaceUsage         `xml:"urn:vsan VsanQueryEntitySpaceUsage,omitempty"`
	Res    *types.VsanQueryEntitySpaceUsageResponse `xml:"urn:vsan VsanQueryEntitySpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryEntitySpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryEntitySpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryEntitySpaceUsage) (*types.VsanQueryEntitySpaceUsageResponse, error) {
	var reqBody, resBody VsanQueryEntitySpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanManagedStorageSpaceUsageBody struct {
	Req    *types.QueryVsanManagedStorageSpaceUsage         `xml:"urn:vsan QueryVsanManagedStorageSpaceUsage,omitempty"`
	Res    *types.QueryVsanManagedStorageSpaceUsageResponse `xml:"urn:vsan QueryVsanManagedStorageSpaceUsageResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanManagedStorageSpaceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanManagedStorageSpaceUsage(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanManagedStorageSpaceUsage) (*types.QueryVsanManagedStorageSpaceUsageResponse, error) {
	var reqBody, resBody QueryVsanManagedStorageSpaceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartIoInsightBody struct {
	Req    *types.StartIoInsight         `xml:"urn:vsan StartIoInsight,omitempty"`
	Res    *types.StartIoInsightResponse `xml:"urn:vsan StartIoInsightResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartIoInsightBody) Fault() *soap.Fault { return b.Fault_ }

func StartIoInsight(ctx context.Context, r soap.RoundTripper, req *types.StartIoInsight) (*types.StartIoInsightResponse, error) {
	var reqBody, resBody StartIoInsightBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIoInsightInstancesBody struct {
	Req    *types.QueryIoInsightInstances         `xml:"urn:vsan QueryIoInsightInstances,omitempty"`
	Res    *types.QueryIoInsightInstancesResponse `xml:"urn:vsan QueryIoInsightInstancesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIoInsightInstancesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIoInsightInstances(ctx context.Context, r soap.RoundTripper, req *types.QueryIoInsightInstances) (*types.QueryIoInsightInstancesResponse, error) {
	var reqBody, resBody QueryIoInsightInstancesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameIoInsightInstanceBody struct {
	Req    *types.RenameIoInsightInstance         `xml:"urn:vsan RenameIoInsightInstance,omitempty"`
	Res    *types.RenameIoInsightInstanceResponse `xml:"urn:vsan RenameIoInsightInstanceResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameIoInsightInstanceBody) Fault() *soap.Fault { return b.Fault_ }

func RenameIoInsightInstance(ctx context.Context, r soap.RoundTripper, req *types.RenameIoInsightInstance) (*types.RenameIoInsightInstanceResponse, error) {
	var reqBody, resBody RenameIoInsightInstanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StopIoInsightBody struct {
	Req    *types.StopIoInsight         `xml:"urn:vsan StopIoInsight,omitempty"`
	Res    *types.StopIoInsightResponse `xml:"urn:vsan StopIoInsightResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopIoInsightBody) Fault() *soap.Fault { return b.Fault_ }

func StopIoInsight(ctx context.Context, r soap.RoundTripper, req *types.StopIoInsight) (*types.StopIoInsightResponse, error) {
	var reqBody, resBody StopIoInsightBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteIoInsightInstanceBody struct {
	Req    *types.DeleteIoInsightInstance         `xml:"urn:vsan DeleteIoInsightInstance,omitempty"`
	Res    *types.DeleteIoInsightInstanceResponse `xml:"urn:vsan DeleteIoInsightInstanceResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteIoInsightInstanceBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteIoInsightInstance(ctx context.Context, r soap.RoundTripper, req *types.DeleteIoInsightInstance) (*types.DeleteIoInsightInstanceResponse, error) {
	var reqBody, resBody DeleteIoInsightInstanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibInstall_TaskBody struct {
	Req    *types.VsanVibInstall_Task         `xml:"urn:vsan VsanVibInstall_Task,omitempty"`
	Res    *types.VsanVibInstall_TaskResponse `xml:"urn:vsan VsanVibInstall_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibInstall_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibInstall_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanVibInstall_Task) (*types.VsanVibInstall_TaskResponse, error) {
	var reqBody, resBody VsanVibInstall_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibInstallPreflightCheckBody struct {
	Req    *types.VsanVibInstallPreflightCheck         `xml:"urn:vsan VsanVibInstallPreflightCheck,omitempty"`
	Res    *types.VsanVibInstallPreflightCheckResponse `xml:"urn:vsan VsanVibInstallPreflightCheckResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibInstallPreflightCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibInstallPreflightCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanVibInstallPreflightCheck) (*types.VsanVibInstallPreflightCheckResponse, error) {
	var reqBody, resBody VsanVibInstallPreflightCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVibScanBody struct {
	Req    *types.VsanVibScan         `xml:"urn:vsan VsanVibScan,omitempty"`
	Res    *types.VsanVibScanResponse `xml:"urn:vsan VsanVibScanResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVibScanBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVibScan(ctx context.Context, r soap.RoundTripper, req *types.VsanVibScan) (*types.VsanVibScanResponse, error) {
	var reqBody, resBody VsanVibScanBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanUnmountDiskMappingExBody struct {
	Req    *types.VsanUnmountDiskMappingEx         `xml:"urn:vsan VsanUnmountDiskMappingEx,omitempty"`
	Res    *types.VsanUnmountDiskMappingExResponse `xml:"urn:vsan VsanUnmountDiskMappingExResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanUnmountDiskMappingExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanUnmountDiskMappingEx(ctx context.Context, r soap.RoundTripper, req *types.VsanUnmountDiskMappingEx) (*types.VsanUnmountDiskMappingExResponse, error) {
	var reqBody, resBody VsanUnmountDiskMappingExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQuerySyncingVsanObjectsBody struct {
	Req    *types.VsanQuerySyncingVsanObjects         `xml:"urn:vsan VsanQuerySyncingVsanObjects,omitempty"`
	Res    *types.VsanQuerySyncingVsanObjectsResponse `xml:"urn:vsan VsanQuerySyncingVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQuerySyncingVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQuerySyncingVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanQuerySyncingVsanObjects) (*types.VsanQuerySyncingVsanObjectsResponse, error) {
	var reqBody, resBody VsanQuerySyncingVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryWipeDiskBody struct {
	Req    *types.VsanHostQueryWipeDisk         `xml:"urn:vsan VsanHostQueryWipeDisk,omitempty"`
	Res    *types.VsanHostQueryWipeDiskResponse `xml:"urn:vsan VsanHostQueryWipeDiskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryWipeDiskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryWipeDisk(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryWipeDisk) (*types.VsanHostQueryWipeDiskResponse, error) {
	var reqBody, resBody VsanHostQueryWipeDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostStatusExBody struct {
	Req    *types.VsanQueryHostStatusEx         `xml:"urn:vsan VsanQueryHostStatusEx,omitempty"`
	Res    *types.VsanQueryHostStatusExResponse `xml:"urn:vsan VsanQueryHostStatusExResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostStatusExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostStatusEx(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostStatusEx) (*types.VsanQueryHostStatusExResponse, error) {
	var reqBody, resBody VsanQueryHostStatusExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostDrsStatsBody struct {
	Req    *types.VsanQueryHostDrsStats         `xml:"urn:vsan VsanQueryHostDrsStats,omitempty"`
	Res    *types.VsanQueryHostDrsStatsResponse `xml:"urn:vsan VsanQueryHostDrsStatsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostDrsStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostDrsStats(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostDrsStats) (*types.VsanQueryHostDrsStatsResponse, error) {
	var reqBody, resBody VsanQueryHostDrsStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryWhatIfEvacuationResultBody struct {
	Req    *types.VsanQueryWhatIfEvacuationResult         `xml:"urn:vsan VsanQueryWhatIfEvacuationResult,omitempty"`
	Res    *types.VsanQueryWhatIfEvacuationResultResponse `xml:"urn:vsan VsanQueryWhatIfEvacuationResultResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryWhatIfEvacuationResultBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryWhatIfEvacuationResult(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryWhatIfEvacuationResult) (*types.VsanQueryWhatIfEvacuationResultResponse, error) {
	var reqBody, resBody VsanQueryWhatIfEvacuationResultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostGetRuntimeStatsBody struct {
	Req    *types.VsanHostGetRuntimeStats         `xml:"urn:vsan VsanHostGetRuntimeStats,omitempty"`
	Res    *types.VsanHostGetRuntimeStatsResponse `xml:"urn:vsan VsanHostGetRuntimeStatsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostGetRuntimeStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostGetRuntimeStats(ctx context.Context, r soap.RoundTripper, req *types.VsanHostGetRuntimeStats) (*types.VsanHostGetRuntimeStatsResponse, error) {
	var reqBody, resBody VsanHostGetRuntimeStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostAbortWipeDiskBody struct {
	Req    *types.VsanHostAbortWipeDisk         `xml:"urn:vsan VsanHostAbortWipeDisk,omitempty"`
	Res    *types.VsanHostAbortWipeDiskResponse `xml:"urn:vsan VsanHostAbortWipeDiskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostAbortWipeDiskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostAbortWipeDisk(ctx context.Context, r soap.RoundTripper, req *types.VsanHostAbortWipeDisk) (*types.VsanHostAbortWipeDiskResponse, error) {
	var reqBody, resBody VsanHostAbortWipeDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetAboutInfoExBody struct {
	Req    *types.VsanGetAboutInfoEx         `xml:"urn:vsan VsanGetAboutInfoEx,omitempty"`
	Res    *types.VsanGetAboutInfoExResponse `xml:"urn:vsan VsanGetAboutInfoExResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetAboutInfoExBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetAboutInfoEx(ctx context.Context, r soap.RoundTripper, req *types.VsanGetAboutInfoEx) (*types.VsanGetAboutInfoExResponse, error) {
	var reqBody, resBody VsanGetAboutInfoExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostWipeDiskBody struct {
	Req    *types.VsanHostWipeDisk         `xml:"urn:vsan VsanHostWipeDisk,omitempty"`
	Res    *types.VsanHostWipeDiskResponse `xml:"urn:vsan VsanHostWipeDiskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostWipeDiskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostWipeDisk(ctx context.Context, r soap.RoundTripper, req *types.VsanHostWipeDisk) (*types.VsanHostWipeDiskResponse, error) {
	var reqBody, resBody VsanHostWipeDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterCreateFsDomainBody struct {
	Req    *types.VsanClusterCreateFsDomain         `xml:"urn:vsan VsanClusterCreateFsDomain,omitempty"`
	Res    *types.VsanClusterCreateFsDomainResponse `xml:"urn:vsan VsanClusterCreateFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterCreateFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterCreateFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterCreateFsDomain) (*types.VsanClusterCreateFsDomainResponse, error) {
	var reqBody, resBody VsanClusterCreateFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryFileServiceOvfsBody struct {
	Req    *types.VsanQueryFileServiceOvfs         `xml:"urn:vsan VsanQueryFileServiceOvfs,omitempty"`
	Res    *types.VsanQueryFileServiceOvfsResponse `xml:"urn:vsan VsanQueryFileServiceOvfsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryFileServiceOvfsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryFileServiceOvfs(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryFileServiceOvfs) (*types.VsanQueryFileServiceOvfsResponse, error) {
	var reqBody, resBody VsanQueryFileServiceOvfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRebalanceFileServiceBody struct {
	Req    *types.VsanRebalanceFileService         `xml:"urn:vsan VsanRebalanceFileService,omitempty"`
	Res    *types.VsanRebalanceFileServiceResponse `xml:"urn:vsan VsanRebalanceFileServiceResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRebalanceFileServiceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRebalanceFileService(ctx context.Context, r soap.RoundTripper, req *types.VsanRebalanceFileService) (*types.VsanRebalanceFileServiceResponse, error) {
	var reqBody, resBody VsanRebalanceFileServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterRemoveFsDomainBody struct {
	Req    *types.VsanClusterRemoveFsDomain         `xml:"urn:vsan VsanClusterRemoveFsDomain,omitempty"`
	Res    *types.VsanClusterRemoveFsDomainResponse `xml:"urn:vsan VsanClusterRemoveFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterRemoveFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterRemoveFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterRemoveFsDomain) (*types.VsanClusterRemoveFsDomainResponse, error) {
	var reqBody, resBody VsanClusterRemoveFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanPerformFileServiceEnablePreflightCheckBody struct {
	Req    *types.VsanPerformFileServiceEnablePreflightCheck         `xml:"urn:vsan VsanPerformFileServiceEnablePreflightCheck,omitempty"`
	Res    *types.VsanPerformFileServiceEnablePreflightCheckResponse `xml:"urn:vsan VsanPerformFileServiceEnablePreflightCheckResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanPerformFileServiceEnablePreflightCheckBody) Fault() *soap.Fault { return b.Fault_ }

func VsanPerformFileServiceEnablePreflightCheck(ctx context.Context, r soap.RoundTripper, req *types.VsanPerformFileServiceEnablePreflightCheck) (*types.VsanPerformFileServiceEnablePreflightCheckResponse, error) {
	var reqBody, resBody VsanPerformFileServiceEnablePreflightCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFindOvfDownloadUrlBody struct {
	Req    *types.VsanFindOvfDownloadUrl         `xml:"urn:vsan VsanFindOvfDownloadUrl,omitempty"`
	Res    *types.VsanFindOvfDownloadUrlResponse `xml:"urn:vsan VsanFindOvfDownloadUrlResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFindOvfDownloadUrlBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFindOvfDownloadUrl(ctx context.Context, r soap.RoundTripper, req *types.VsanFindOvfDownloadUrl) (*types.VsanFindOvfDownloadUrlResponse, error) {
	var reqBody, resBody VsanFindOvfDownloadUrlBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterReconfigureFsDomainBody struct {
	Req    *types.VsanClusterReconfigureFsDomain         `xml:"urn:vsan VsanClusterReconfigureFsDomain,omitempty"`
	Res    *types.VsanClusterReconfigureFsDomainResponse `xml:"urn:vsan VsanClusterReconfigureFsDomainResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterReconfigureFsDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterReconfigureFsDomain(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterReconfigureFsDomain) (*types.VsanClusterReconfigureFsDomainResponse, error) {
	var reqBody, resBody VsanClusterReconfigureFsDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFsDomainsBody struct {
	Req    *types.VsanClusterQueryFsDomains         `xml:"urn:vsan VsanClusterQueryFsDomains,omitempty"`
	Res    *types.VsanClusterQueryFsDomainsResponse `xml:"urn:vsan VsanClusterQueryFsDomainsResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFsDomainsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFsDomains(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFsDomains) (*types.VsanClusterQueryFsDomainsResponse, error) {
	var reqBody, resBody VsanClusterQueryFsDomainsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanUpgradeFsvmBody struct {
	Req    *types.VsanUpgradeFsvm         `xml:"urn:vsan VsanUpgradeFsvm,omitempty"`
	Res    *types.VsanUpgradeFsvmResponse `xml:"urn:vsan VsanUpgradeFsvmResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanUpgradeFsvmBody) Fault() *soap.Fault { return b.Fault_ }

func VsanUpgradeFsvm(ctx context.Context, r soap.RoundTripper, req *types.VsanUpgradeFsvm) (*types.VsanUpgradeFsvmResponse, error) {
	var reqBody, resBody VsanUpgradeFsvmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterRemoveShareBody struct {
	Req    *types.VsanClusterRemoveShare         `xml:"urn:vsan VsanClusterRemoveShare,omitempty"`
	Res    *types.VsanClusterRemoveShareResponse `xml:"urn:vsan VsanClusterRemoveShareResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterRemoveShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterRemoveShare(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterRemoveShare) (*types.VsanClusterRemoveShareResponse, error) {
	var reqBody, resBody VsanClusterRemoveShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCreateFileShareBody struct {
	Req    *types.VsanCreateFileShare         `xml:"urn:vsan VsanCreateFileShare,omitempty"`
	Res    *types.VsanCreateFileShareResponse `xml:"urn:vsan VsanCreateFileShareResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCreateFileShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCreateFileShare(ctx context.Context, r soap.RoundTripper, req *types.VsanCreateFileShare) (*types.VsanCreateFileShareResponse, error) {
	var reqBody, resBody VsanCreateFileShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanDownloadFileServiceOvfBody struct {
	Req    *types.VsanDownloadFileServiceOvf         `xml:"urn:vsan VsanDownloadFileServiceOvf,omitempty"`
	Res    *types.VsanDownloadFileServiceOvfResponse `xml:"urn:vsan VsanDownloadFileServiceOvfResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanDownloadFileServiceOvfBody) Fault() *soap.Fault { return b.Fault_ }

func VsanDownloadFileServiceOvf(ctx context.Context, r soap.RoundTripper, req *types.VsanDownloadFileServiceOvf) (*types.VsanDownloadFileServiceOvfResponse, error) {
	var reqBody, resBody VsanDownloadFileServiceOvfBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanReconfigureFileShareBody struct {
	Req    *types.VsanReconfigureFileShare         `xml:"urn:vsan VsanReconfigureFileShare,omitempty"`
	Res    *types.VsanReconfigureFileShareResponse `xml:"urn:vsan VsanReconfigureFileShareResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanReconfigureFileShareBody) Fault() *soap.Fault { return b.Fault_ }

func VsanReconfigureFileShare(ctx context.Context, r soap.RoundTripper, req *types.VsanReconfigureFileShare) (*types.VsanReconfigureFileShareResponse, error) {
	var reqBody, resBody VsanReconfigureFileShareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterQueryFileSharesBody struct {
	Req    *types.VsanClusterQueryFileShares         `xml:"urn:vsan VsanClusterQueryFileShares,omitempty"`
	Res    *types.VsanClusterQueryFileSharesResponse `xml:"urn:vsan VsanClusterQueryFileSharesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterQueryFileSharesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterQueryFileShares(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterQueryFileShares) (*types.VsanClusterQueryFileSharesResponse, error) {
	var reqBody, resBody VsanClusterQueryFileSharesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryAdvCfgBody struct {
	Req    *types.VsanHostQueryAdvCfg         `xml:"urn:vsan VsanHostQueryAdvCfg,omitempty"`
	Res    *types.VsanHostQueryAdvCfgResponse `xml:"urn:vsan VsanHostQueryAdvCfgResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryAdvCfgBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryAdvCfg(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryAdvCfg) (*types.VsanHostQueryAdvCfgResponse, error) {
	var reqBody, resBody VsanHostQueryAdvCfgBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryRunIperfClientBody struct {
	Req    *types.VsanHostQueryRunIperfClient         `xml:"urn:vsan VsanHostQueryRunIperfClient,omitempty"`
	Res    *types.VsanHostQueryRunIperfClientResponse `xml:"urn:vsan VsanHostQueryRunIperfClientResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryRunIperfClientBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryRunIperfClient(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryRunIperfClient) (*types.VsanHostQueryRunIperfClientResponse, error) {
	var reqBody, resBody VsanHostQueryRunIperfClientBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryObjectHealthSummaryBody struct {
	Req    *types.VsanHostQueryObjectHealthSummary         `xml:"urn:vsan VsanHostQueryObjectHealthSummary,omitempty"`
	Res    *types.VsanHostQueryObjectHealthSummaryResponse `xml:"urn:vsan VsanHostQueryObjectHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryObjectHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryObjectHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryObjectHealthSummary) (*types.VsanHostQueryObjectHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryObjectHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStopProactiveRebalanceBody struct {
	Req    *types.VsanStopProactiveRebalance         `xml:"urn:vsan VsanStopProactiveRebalance,omitempty"`
	Res    *types.VsanStopProactiveRebalanceResponse `xml:"urn:vsan VsanStopProactiveRebalanceResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStopProactiveRebalanceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStopProactiveRebalance(ctx context.Context, r soap.RoundTripper, req *types.VsanStopProactiveRebalance) (*types.VsanStopProactiveRebalanceResponse, error) {
	var reqBody, resBody VsanStopProactiveRebalanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryFileServiceHealthSummaryBody struct {
	Req    *types.VsanHostQueryFileServiceHealthSummary         `xml:"urn:vsan VsanHostQueryFileServiceHealthSummary,omitempty"`
	Res    *types.VsanHostQueryFileServiceHealthSummaryResponse `xml:"urn:vsan VsanHostQueryFileServiceHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryFileServiceHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryFileServiceHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryFileServiceHealthSummary) (*types.VsanHostQueryFileServiceHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryFileServiceHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostClomdLivenessBody struct {
	Req    *types.VsanHostClomdLiveness         `xml:"urn:vsan VsanHostClomdLiveness,omitempty"`
	Res    *types.VsanHostClomdLivenessResponse `xml:"urn:vsan VsanHostClomdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostClomdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostClomdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanHostClomdLiveness) (*types.VsanHostClomdLivenessResponse, error) {
	var reqBody, resBody VsanHostClomdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostRepairImmediateObjectsBody struct {
	Req    *types.VsanHostRepairImmediateObjects         `xml:"urn:vsan VsanHostRepairImmediateObjects,omitempty"`
	Res    *types.VsanHostRepairImmediateObjectsResponse `xml:"urn:vsan VsanHostRepairImmediateObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostRepairImmediateObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostRepairImmediateObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanHostRepairImmediateObjects) (*types.VsanHostRepairImmediateObjectsResponse, error) {
	var reqBody, resBody VsanHostRepairImmediateObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryVerifyNetworkSettingsBody struct {
	Req    *types.VsanHostQueryVerifyNetworkSettings         `xml:"urn:vsan VsanHostQueryVerifyNetworkSettings,omitempty"`
	Res    *types.VsanHostQueryVerifyNetworkSettingsResponse `xml:"urn:vsan VsanHostQueryVerifyNetworkSettingsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryVerifyNetworkSettingsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryVerifyNetworkSettings(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryVerifyNetworkSettings) (*types.VsanHostQueryVerifyNetworkSettingsResponse, error) {
	var reqBody, resBody VsanHostQueryVerifyNetworkSettingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCleanupVmdkLoadTestBody struct {
	Req    *types.VsanHostCleanupVmdkLoadTest         `xml:"urn:vsan VsanHostCleanupVmdkLoadTest,omitempty"`
	Res    *types.VsanHostCleanupVmdkLoadTestResponse `xml:"urn:vsan VsanHostCleanupVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCleanupVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCleanupVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCleanupVmdkLoadTest) (*types.VsanHostCleanupVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostCleanupVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanStartProactiveRebalanceBody struct {
	Req    *types.VsanStartProactiveRebalance         `xml:"urn:vsan VsanStartProactiveRebalance,omitempty"`
	Res    *types.VsanStartProactiveRebalanceResponse `xml:"urn:vsan VsanStartProactiveRebalanceResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanStartProactiveRebalanceBody) Fault() *soap.Fault { return b.Fault_ }

func VsanStartProactiveRebalance(ctx context.Context, r soap.RoundTripper, req *types.VsanStartProactiveRebalance) (*types.VsanStartProactiveRebalanceResponse, error) {
	var reqBody, resBody VsanStartProactiveRebalanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryEncryptionHealthSummaryBody struct {
	Req    *types.VsanHostQueryEncryptionHealthSummary         `xml:"urn:vsan VsanHostQueryEncryptionHealthSummary,omitempty"`
	Res    *types.VsanHostQueryEncryptionHealthSummaryResponse `xml:"urn:vsan VsanHostQueryEncryptionHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryEncryptionHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryEncryptionHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryEncryptionHealthSummary) (*types.VsanHostQueryEncryptionHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryEncryptionHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanFlashScsiControllerFirmware_TaskBody struct {
	Req    *types.VsanFlashScsiControllerFirmware_Task         `xml:"urn:vsan VsanFlashScsiControllerFirmware_Task,omitempty"`
	Res    *types.VsanFlashScsiControllerFirmware_TaskResponse `xml:"urn:vsan VsanFlashScsiControllerFirmware_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanFlashScsiControllerFirmware_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VsanFlashScsiControllerFirmware_Task(ctx context.Context, r soap.RoundTripper, req *types.VsanFlashScsiControllerFirmware_Task) (*types.VsanFlashScsiControllerFirmware_TaskResponse, error) {
	var reqBody, resBody VsanFlashScsiControllerFirmware_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryHostEMMStateBody struct {
	Req    *types.VsanQueryHostEMMState         `xml:"urn:vsan VsanQueryHostEMMState,omitempty"`
	Res    *types.VsanQueryHostEMMStateResponse `xml:"urn:vsan VsanQueryHostEMMStateResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryHostEMMStateBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryHostEMMState(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryHostEMMState) (*types.VsanQueryHostEMMStateResponse, error) {
	var reqBody, resBody VsanQueryHostEMMStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanWaitForVsanHealthGenerationIdChangeBody struct {
	Req    *types.VsanWaitForVsanHealthGenerationIdChange         `xml:"urn:vsan VsanWaitForVsanHealthGenerationIdChange,omitempty"`
	Res    *types.VsanWaitForVsanHealthGenerationIdChangeResponse `xml:"urn:vsan VsanWaitForVsanHealthGenerationIdChangeResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanWaitForVsanHealthGenerationIdChangeBody) Fault() *soap.Fault { return b.Fault_ }

func VsanWaitForVsanHealthGenerationIdChange(ctx context.Context, r soap.RoundTripper, req *types.VsanWaitForVsanHealthGenerationIdChange) (*types.VsanWaitForVsanHealthGenerationIdChangeResponse, error) {
	var reqBody, resBody VsanWaitForVsanHealthGenerationIdChangeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryHealthSystemVersionBody struct {
	Req    *types.VsanHostQueryHealthSystemVersion         `xml:"urn:vsan VsanHostQueryHealthSystemVersion,omitempty"`
	Res    *types.VsanHostQueryHealthSystemVersionResponse `xml:"urn:vsan VsanHostQueryHealthSystemVersionResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryHealthSystemVersionBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryHealthSystemVersion(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryHealthSystemVersion) (*types.VsanHostQueryHealthSystemVersionResponse, error) {
	var reqBody, resBody VsanHostQueryHealthSystemVersionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetHclInfoBody struct {
	Req    *types.VsanGetHclInfo         `xml:"urn:vsan VsanGetHclInfo,omitempty"`
	Res    *types.VsanGetHclInfoResponse `xml:"urn:vsan VsanGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanGetHclInfo) (*types.VsanGetHclInfoResponse, error) {
	var reqBody, resBody VsanGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostRunVmdkLoadTestBody struct {
	Req    *types.VsanHostRunVmdkLoadTest         `xml:"urn:vsan VsanHostRunVmdkLoadTest,omitempty"`
	Res    *types.VsanHostRunVmdkLoadTestResponse `xml:"urn:vsan VsanHostRunVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostRunVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostRunVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostRunVmdkLoadTest) (*types.VsanHostRunVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostRunVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQuerySmartStatsBody struct {
	Req    *types.VsanHostQuerySmartStats         `xml:"urn:vsan VsanHostQuerySmartStats,omitempty"`
	Res    *types.VsanHostQuerySmartStatsResponse `xml:"urn:vsan VsanHostQuerySmartStatsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQuerySmartStatsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQuerySmartStats(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQuerySmartStats) (*types.VsanHostQuerySmartStatsResponse, error) {
	var reqBody, resBody VsanHostQuerySmartStatsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostPrepareVmdkLoadTestBody struct {
	Req    *types.VsanHostPrepareVmdkLoadTest         `xml:"urn:vsan VsanHostPrepareVmdkLoadTest,omitempty"`
	Res    *types.VsanHostPrepareVmdkLoadTestResponse `xml:"urn:vsan VsanHostPrepareVmdkLoadTestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostPrepareVmdkLoadTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostPrepareVmdkLoadTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostPrepareVmdkLoadTest) (*types.VsanHostPrepareVmdkLoadTestResponse, error) {
	var reqBody, resBody VsanHostPrepareVmdkLoadTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryRunIperfServerBody struct {
	Req    *types.VsanHostQueryRunIperfServer         `xml:"urn:vsan VsanHostQueryRunIperfServer,omitempty"`
	Res    *types.VsanHostQueryRunIperfServerResponse `xml:"urn:vsan VsanHostQueryRunIperfServerResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryRunIperfServerBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryRunIperfServer(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryRunIperfServer) (*types.VsanHostQueryRunIperfServerResponse, error) {
	var reqBody, resBody VsanHostQueryRunIperfServerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanGetProactiveRebalanceInfoBody struct {
	Req    *types.VsanGetProactiveRebalanceInfo         `xml:"urn:vsan VsanGetProactiveRebalanceInfo,omitempty"`
	Res    *types.VsanGetProactiveRebalanceInfoResponse `xml:"urn:vsan VsanGetProactiveRebalanceInfoResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanGetProactiveRebalanceInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanGetProactiveRebalanceInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanGetProactiveRebalanceInfo) (*types.VsanGetProactiveRebalanceInfoResponse, error) {
	var reqBody, resBody VsanGetProactiveRebalanceInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryPhysicalDiskHealthSummaryBody struct {
	Req    *types.VsanHostQueryPhysicalDiskHealthSummary         `xml:"urn:vsan VsanHostQueryPhysicalDiskHealthSummary,omitempty"`
	Res    *types.VsanHostQueryPhysicalDiskHealthSummaryResponse `xml:"urn:vsan VsanHostQueryPhysicalDiskHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryPhysicalDiskHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryPhysicalDiskHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryPhysicalDiskHealthSummary) (*types.VsanHostQueryPhysicalDiskHealthSummaryResponse, error) {
	var reqBody, resBody VsanHostQueryPhysicalDiskHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryHostInfoByUuidsBody struct {
	Req    *types.VsanHostQueryHostInfoByUuids         `xml:"urn:vsan VsanHostQueryHostInfoByUuids,omitempty"`
	Res    *types.VsanHostQueryHostInfoByUuidsResponse `xml:"urn:vsan VsanHostQueryHostInfoByUuidsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryHostInfoByUuidsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryHostInfoByUuids(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryHostInfoByUuids) (*types.VsanHostQueryHostInfoByUuidsResponse, error) {
	var reqBody, resBody VsanHostQueryHostInfoByUuidsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostCreateVmHealthTestBody struct {
	Req    *types.VsanHostCreateVmHealthTest         `xml:"urn:vsan VsanHostCreateVmHealthTest,omitempty"`
	Res    *types.VsanHostCreateVmHealthTestResponse `xml:"urn:vsan VsanHostCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanHostCreateVmHealthTest) (*types.VsanHostCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanHostCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanHostQueryCheckLimitsBody struct {
	Req    *types.VsanHostQueryCheckLimits         `xml:"urn:vsan VsanHostQueryCheckLimits,omitempty"`
	Res    *types.VsanHostQueryCheckLimitsResponse `xml:"urn:vsan VsanHostQueryCheckLimitsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanHostQueryCheckLimitsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanHostQueryCheckLimits(ctx context.Context, r soap.RoundTripper, req *types.VsanHostQueryCheckLimits) (*types.VsanHostQueryCheckLimitsResponse, error) {
	var reqBody, resBody VsanHostQueryCheckLimitsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcIsWitnessHostBody struct {
	Req    *types.VSANVcIsWitnessHost         `xml:"urn:vsan VSANVcIsWitnessHost,omitempty"`
	Res    *types.VSANVcIsWitnessHostResponse `xml:"urn:vsan VSANVcIsWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcIsWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcIsWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcIsWitnessHost) (*types.VSANVcIsWitnessHostResponse, error) {
	var reqBody, resBody VSANVcIsWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcReplaceWitnessHostForClustersBody struct {
	Req    *types.VsanVcReplaceWitnessHostForClusters         `xml:"urn:vsan VsanVcReplaceWitnessHostForClusters,omitempty"`
	Res    *types.VsanVcReplaceWitnessHostForClustersResponse `xml:"urn:vsan VsanVcReplaceWitnessHostForClustersResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcReplaceWitnessHostForClustersBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcReplaceWitnessHostForClusters(ctx context.Context, r soap.RoundTripper, req *types.VsanVcReplaceWitnessHostForClusters) (*types.VsanVcReplaceWitnessHostForClustersResponse, error) {
	var reqBody, resBody VsanVcReplaceWitnessHostForClustersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanVcAddWitnessHostForClustersBody struct {
	Req    *types.VsanVcAddWitnessHostForClusters         `xml:"urn:vsan VsanVcAddWitnessHostForClusters,omitempty"`
	Res    *types.VsanVcAddWitnessHostForClustersResponse `xml:"urn:vsan VsanVcAddWitnessHostForClustersResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanVcAddWitnessHostForClustersBody) Fault() *soap.Fault { return b.Fault_ }

func VsanVcAddWitnessHostForClusters(ctx context.Context, r soap.RoundTripper, req *types.VsanVcAddWitnessHostForClusters) (*types.VsanVcAddWitnessHostForClustersResponse, error) {
	var reqBody, resBody VsanVcAddWitnessHostForClustersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcSetPreferredFaultDomainBody struct {
	Req    *types.VSANVcSetPreferredFaultDomain         `xml:"urn:vsan VSANVcSetPreferredFaultDomain,omitempty"`
	Res    *types.VSANVcSetPreferredFaultDomainResponse `xml:"urn:vsan VSANVcSetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcSetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcSetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANVcSetPreferredFaultDomain) (*types.VSANVcSetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANVcSetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySharedWitnessClusterInfoBody struct {
	Req    *types.QuerySharedWitnessClusterInfo         `xml:"urn:vsan QuerySharedWitnessClusterInfo,omitempty"`
	Res    *types.QuerySharedWitnessClusterInfoResponse `xml:"urn:vsan QuerySharedWitnessClusterInfoResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySharedWitnessClusterInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySharedWitnessClusterInfo(ctx context.Context, r soap.RoundTripper, req *types.QuerySharedWitnessClusterInfo) (*types.QuerySharedWitnessClusterInfoResponse, error) {
	var reqBody, resBody QuerySharedWitnessClusterInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetPreferredFaultDomainBody struct {
	Req    *types.VSANVcGetPreferredFaultDomain         `xml:"urn:vsan VSANVcGetPreferredFaultDomain,omitempty"`
	Res    *types.VSANVcGetPreferredFaultDomainResponse `xml:"urn:vsan VSANVcGetPreferredFaultDomainResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetPreferredFaultDomainBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetPreferredFaultDomain(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetPreferredFaultDomain) (*types.VSANVcGetPreferredFaultDomainResponse, error) {
	var reqBody, resBody VSANVcGetPreferredFaultDomainBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANIsWitnessVirtualApplianceBody struct {
	Req    *types.VSANIsWitnessVirtualAppliance         `xml:"urn:vsan VSANIsWitnessVirtualAppliance,omitempty"`
	Res    *types.VSANIsWitnessVirtualApplianceResponse `xml:"urn:vsan VSANIsWitnessVirtualApplianceResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANIsWitnessVirtualApplianceBody) Fault() *soap.Fault { return b.Fault_ }

func VSANIsWitnessVirtualAppliance(ctx context.Context, r soap.RoundTripper, req *types.VSANIsWitnessVirtualAppliance) (*types.VSANIsWitnessVirtualApplianceResponse, error) {
	var reqBody, resBody VSANIsWitnessVirtualApplianceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcAddWitnessHostBody struct {
	Req    *types.VSANVcAddWitnessHost         `xml:"urn:vsan VSANVcAddWitnessHost,omitempty"`
	Res    *types.VSANVcAddWitnessHostResponse `xml:"urn:vsan VSANVcAddWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcAddWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcAddWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcAddWitnessHost) (*types.VSANVcAddWitnessHostResponse, error) {
	var reqBody, resBody VSANVcAddWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcGetWitnessHostsBody struct {
	Req    *types.VSANVcGetWitnessHosts         `xml:"urn:vsan VSANVcGetWitnessHosts,omitempty"`
	Res    *types.VSANVcGetWitnessHostsResponse `xml:"urn:vsan VSANVcGetWitnessHostsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcGetWitnessHostsBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcGetWitnessHosts(ctx context.Context, r soap.RoundTripper, req *types.VSANVcGetWitnessHosts) (*types.VSANVcGetWitnessHostsResponse, error) {
	var reqBody, resBody VSANVcGetWitnessHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcRetrieveStretchedClusterVcCapabilityBody struct {
	Req    *types.VSANVcRetrieveStretchedClusterVcCapability         `xml:"urn:vsan VSANVcRetrieveStretchedClusterVcCapability,omitempty"`
	Res    *types.VSANVcRetrieveStretchedClusterVcCapabilityResponse `xml:"urn:vsan VSANVcRetrieveStretchedClusterVcCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcRetrieveStretchedClusterVcCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcRetrieveStretchedClusterVcCapability(ctx context.Context, r soap.RoundTripper, req *types.VSANVcRetrieveStretchedClusterVcCapability) (*types.VSANVcRetrieveStretchedClusterVcCapabilityResponse, error) {
	var reqBody, resBody VSANVcRetrieveStretchedClusterVcCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcConvertToStretchedClusterBody struct {
	Req    *types.VSANVcConvertToStretchedCluster         `xml:"urn:vsan VSANVcConvertToStretchedCluster,omitempty"`
	Res    *types.VSANVcConvertToStretchedClusterResponse `xml:"urn:vsan VSANVcConvertToStretchedClusterResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcConvertToStretchedClusterBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcConvertToStretchedCluster(ctx context.Context, r soap.RoundTripper, req *types.VSANVcConvertToStretchedCluster) (*types.VSANVcConvertToStretchedClusterResponse, error) {
	var reqBody, resBody VSANVcConvertToStretchedClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VSANVcRemoveWitnessHostBody struct {
	Req    *types.VSANVcRemoveWitnessHost         `xml:"urn:vsan VSANVcRemoveWitnessHost,omitempty"`
	Res    *types.VSANVcRemoveWitnessHostResponse `xml:"urn:vsan VSANVcRemoveWitnessHostResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VSANVcRemoveWitnessHostBody) Fault() *soap.Fault { return b.Fault_ }

func VSANVcRemoveWitnessHost(ctx context.Context, r soap.RoundTripper, req *types.VSANVcRemoveWitnessHost) (*types.VSANVcRemoveWitnessHostResponse, error) {
	var reqBody, resBody VSANVcRemoveWitnessHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySharedWitnessCompatibilityBody struct {
	Req    *types.QuerySharedWitnessCompatibility         `xml:"urn:vsan QuerySharedWitnessCompatibility,omitempty"`
	Res    *types.QuerySharedWitnessCompatibilityResponse `xml:"urn:vsan QuerySharedWitnessCompatibilityResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySharedWitnessCompatibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySharedWitnessCompatibility(ctx context.Context, r soap.RoundTripper, req *types.QuerySharedWitnessCompatibility) (*types.QuerySharedWitnessCompatibilityResponse, error) {
	var reqBody, resBody QuerySharedWitnessCompatibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterPhysicalDiskHealthSummaryBody struct {
	Req    *types.VsanQueryClusterPhysicalDiskHealthSummary         `xml:"urn:vsan VsanQueryClusterPhysicalDiskHealthSummary,omitempty"`
	Res    *types.VsanQueryClusterPhysicalDiskHealthSummaryResponse `xml:"urn:vsan VsanQueryClusterPhysicalDiskHealthSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterPhysicalDiskHealthSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterPhysicalDiskHealthSummary(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterPhysicalDiskHealthSummary) (*types.VsanQueryClusterPhysicalDiskHealthSummaryResponse, error) {
	var reqBody, resBody VsanQueryClusterPhysicalDiskHealthSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterNetworkPerfTestBody struct {
	Req    *types.VsanQueryClusterNetworkPerfTest         `xml:"urn:vsan VsanQueryClusterNetworkPerfTest,omitempty"`
	Res    *types.VsanQueryClusterNetworkPerfTestResponse `xml:"urn:vsan VsanQueryClusterNetworkPerfTestResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterNetworkPerfTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterNetworkPerfTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterNetworkPerfTest) (*types.VsanQueryClusterNetworkPerfTestResponse, error) {
	var reqBody, resBody VsanQueryClusterNetworkPerfTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterAdvCfgSyncBody struct {
	Req    *types.VsanQueryClusterAdvCfgSync         `xml:"urn:vsan VsanQueryClusterAdvCfgSync,omitempty"`
	Res    *types.VsanQueryClusterAdvCfgSyncResponse `xml:"urn:vsan VsanQueryClusterAdvCfgSyncResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterAdvCfgSyncBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterAdvCfgSync(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterAdvCfgSync) (*types.VsanQueryClusterAdvCfgSyncResponse, error) {
	var reqBody, resBody VsanQueryClusterAdvCfgSyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRepairClusterImmediateObjectsBody struct {
	Req    *types.VsanRepairClusterImmediateObjects         `xml:"urn:vsan VsanRepairClusterImmediateObjects,omitempty"`
	Res    *types.VsanRepairClusterImmediateObjectsResponse `xml:"urn:vsan VsanRepairClusterImmediateObjectsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRepairClusterImmediateObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRepairClusterImmediateObjects(ctx context.Context, r soap.RoundTripper, req *types.VsanRepairClusterImmediateObjects) (*types.VsanRepairClusterImmediateObjectsResponse, error) {
	var reqBody, resBody VsanRepairClusterImmediateObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryVerifyClusterNetworkSettingsBody struct {
	Req    *types.VsanQueryVerifyClusterNetworkSettings         `xml:"urn:vsan VsanQueryVerifyClusterNetworkSettings,omitempty"`
	Res    *types.VsanQueryVerifyClusterNetworkSettingsResponse `xml:"urn:vsan VsanQueryVerifyClusterNetworkSettingsResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryVerifyClusterNetworkSettingsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryVerifyClusterNetworkSettings(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryVerifyClusterNetworkSettings) (*types.VsanQueryVerifyClusterNetworkSettingsResponse, error) {
	var reqBody, resBody VsanQueryVerifyClusterNetworkSettingsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCreateVmHealthTestBody struct {
	Req    *types.VsanQueryClusterCreateVmHealthTest         `xml:"urn:vsan VsanQueryClusterCreateVmHealthTest,omitempty"`
	Res    *types.VsanQueryClusterCreateVmHealthTestResponse `xml:"urn:vsan VsanQueryClusterCreateVmHealthTestResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCreateVmHealthTestBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCreateVmHealthTest(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCreateVmHealthTest) (*types.VsanQueryClusterCreateVmHealthTestResponse, error) {
	var reqBody, resBody VsanQueryClusterCreateVmHealthTestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterHealthSystemVersionsBody struct {
	Req    *types.VsanQueryClusterHealthSystemVersions         `xml:"urn:vsan VsanQueryClusterHealthSystemVersions,omitempty"`
	Res    *types.VsanQueryClusterHealthSystemVersionsResponse `xml:"urn:vsan VsanQueryClusterHealthSystemVersionsResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterHealthSystemVersionsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterHealthSystemVersions(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterHealthSystemVersions) (*types.VsanQueryClusterHealthSystemVersionsResponse, error) {
	var reqBody, resBody VsanQueryClusterHealthSystemVersionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanClusterGetHclInfoBody struct {
	Req    *types.VsanClusterGetHclInfo         `xml:"urn:vsan VsanClusterGetHclInfo,omitempty"`
	Res    *types.VsanClusterGetHclInfoResponse `xml:"urn:vsan VsanClusterGetHclInfoResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanClusterGetHclInfoBody) Fault() *soap.Fault { return b.Fault_ }

func VsanClusterGetHclInfo(ctx context.Context, r soap.RoundTripper, req *types.VsanClusterGetHclInfo) (*types.VsanClusterGetHclInfoResponse, error) {
	var reqBody, resBody VsanClusterGetHclInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCheckLimitsBody struct {
	Req    *types.VsanQueryClusterCheckLimits         `xml:"urn:vsan VsanQueryClusterCheckLimits,omitempty"`
	Res    *types.VsanQueryClusterCheckLimitsResponse `xml:"urn:vsan VsanQueryClusterCheckLimitsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCheckLimitsBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCheckLimits(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCheckLimits) (*types.VsanQueryClusterCheckLimitsResponse, error) {
	var reqBody, resBody VsanQueryClusterCheckLimitsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanQueryClusterCaptureVsanPcapBody struct {
	Req    *types.VsanQueryClusterCaptureVsanPcap         `xml:"urn:vsan VsanQueryClusterCaptureVsanPcap,omitempty"`
	Res    *types.VsanQueryClusterCaptureVsanPcapResponse `xml:"urn:vsan VsanQueryClusterCaptureVsanPcapResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanQueryClusterCaptureVsanPcapBody) Fault() *soap.Fault { return b.Fault_ }

func VsanQueryClusterCaptureVsanPcap(ctx context.Context, r soap.RoundTripper, req *types.VsanQueryClusterCaptureVsanPcap) (*types.VsanQueryClusterCaptureVsanPcapResponse, error) {
	var reqBody, resBody VsanQueryClusterCaptureVsanPcapBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanCheckClusterClomdLivenessBody struct {
	Req    *types.VsanCheckClusterClomdLiveness         `xml:"urn:vsan VsanCheckClusterClomdLiveness,omitempty"`
	Res    *types.VsanCheckClusterClomdLivenessResponse `xml:"urn:vsan VsanCheckClusterClomdLivenessResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanCheckClusterClomdLivenessBody) Fault() *soap.Fault { return b.Fault_ }

func VsanCheckClusterClomdLiveness(ctx context.Context, r soap.RoundTripper, req *types.VsanCheckClusterClomdLiveness) (*types.VsanCheckClusterClomdLivenessResponse, error) {
	var reqBody, resBody VsanCheckClusterClomdLivenessBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VsanRetrievePropertiesBody struct {
	Req    *types.VsanRetrieveProperties         `xml:"urn:vsan VsanRetrieveProperties,omitempty"`
	Res    *types.VsanRetrievePropertiesResponse `xml:"urn:vsan VsanRetrievePropertiesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VsanRetrievePropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func VsanRetrieveProperties(ctx context.Context, r soap.RoundTripper, req *types.VsanRetrieveProperties) (*types.VsanRetrievePropertiesResponse, error) {
	var reqBody, resBody VsanRetrievePropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
