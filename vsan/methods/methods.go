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
	"github.com/vmware/govmomi/vsan/types"
)

// Supported Entities
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

//Object Identities

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

// Health summary
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

// Space usage
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

// Vsan performance
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

// Resyncing summary
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

type QueryVsanObjectsBody struct {
	Req    *types.QueryVsanObjects         `xml:"urn:vsan QueryVsanObjects,omitempty"`
	Res    *types.QueryVsanObjectsResponse `xml:"urn:vsan QueryVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanObjects) (*types.QueryVsanObjectsResponse, error) {
	var reqBody, resBody QueryVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
