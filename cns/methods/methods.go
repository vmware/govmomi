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

package methods

import (
	"context"

	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type CnsCreateVolumeBody struct {
	Req    *types.CnsCreateVolume         `xml:"urn:vsan CnsCreateVolume,omitempty"`
	Res    *types.CnsCreateVolumeResponse `xml:"urn:vsan CnsCreateVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsCreateVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsCreateVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsCreateVolume) (*types.CnsCreateVolumeResponse, error) {
	var reqBody, resBody CnsCreateVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsUpdateVolumeBody struct {
	Req    *types.CnsUpdateVolumeMetadata         `xml:"urn:vsan CnsUpdateVolumeMetadata,omitempty"`
	Res    *types.CnsUpdateVolumeMetadataResponse `xml:"urn:vsan CnsUpdateVolumeMetadataResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsUpdateVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsUpdateVolumeMetadata(ctx context.Context, r soap.RoundTripper, req *types.CnsUpdateVolumeMetadata) (*types.CnsUpdateVolumeMetadataResponse, error) {
	var reqBody, resBody CnsUpdateVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsDeleteVolumeBody struct {
	Req    *types.CnsDeleteVolume         `xml:"urn:vsan CnsDeleteVolume,omitempty"`
	Res    *types.CnsDeleteVolumeResponse `xml:"urn:vsan CnsDeleteVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsDeleteVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsDeleteVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsDeleteVolume) (*types.CnsDeleteVolumeResponse, error) {
	var reqBody, resBody CnsDeleteVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsExtendVolumeBody struct {
	Req    *types.CnsExtendVolume         `xml:"urn:vsan CnsExtendVolume,omitempty"`
	Res    *types.CnsExtendVolumeResponse `xml:"urn:vsan CnsExtendVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsExtendVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsExtendVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsExtendVolume) (*types.CnsExtendVolumeResponse, error) {
	var reqBody, resBody CnsExtendVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsAttachVolumeBody struct {
	Req    *types.CnsAttachVolume         `xml:"urn:vsan CnsAttachVolume,omitempty"`
	Res    *types.CnsAttachVolumeResponse `xml:"urn:vsan CnsAttachVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsAttachVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsAttachVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsAttachVolume) (*types.CnsAttachVolumeResponse, error) {
	var reqBody, resBody CnsAttachVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsDetachVolumeBody struct {
	Req    *types.CnsDetachVolume         `xml:"urn:vsan CnsDetachVolume,omitempty"`
	Res    *types.CnsDetachVolumeResponse `xml:"urn:vsan CnsDetachVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsDetachVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsDetachVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsDetachVolume) (*types.CnsDetachVolumeResponse, error) {
	var reqBody, resBody CnsDetachVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryVolumeBody struct {
	Req    *types.CnsQueryVolume         `xml:"urn:vsan CnsQueryVolume,omitempty"`
	Res    *types.CnsQueryVolumeResponse `xml:"urn:vsan CnsQueryVolumeResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryVolume) (*types.CnsQueryVolumeResponse, error) {
	var reqBody, resBody CnsQueryVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryVolumeInfoBody struct {
	Req    *types.CnsQueryVolumeInfo         `xml:"urn:vsan CnsQueryVolumeInfo,omitempty"`
	Res    *types.CnsQueryVolumeInfoResponse `xml:"urn:vsan CnsQueryVolumeInfoResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryVolumeInfoBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryVolumeInfo(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryVolumeInfo) (*types.CnsQueryVolumeInfoResponse, error) {
	var reqBody, resBody CnsQueryVolumeInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryAllVolumeBody struct {
	Req    *types.CnsQueryAllVolume         `xml:"urn:vsan CnsQueryAllVolume,omitempty"`
	Res    *types.CnsQueryAllVolumeResponse `xml:"urn:vsan CnsQueryAllVolumeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryAllVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryAllVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryAllVolume) (*types.CnsQueryAllVolumeResponse, error) {
	var reqBody, resBody CnsQueryAllVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsRelocateVolumeBody struct {
	Req    *types.CnsRelocateVolume         `xml:"urn:vsan CnsRelocateVolume,omitempty"`
	Res    *types.CnsRelocateVolumeResponse `xml:"urn:vsan CnsRelocateVolumeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsRelocateVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func CnsRelocateVolume(ctx context.Context, r soap.RoundTripper, req *types.CnsRelocateVolume) (*types.CnsRelocateVolumeResponse, error) {
	var reqBody, resBody CnsRelocateVolumeBody
	reqBody.Req = req
	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsConfigureVolumeACLsBody struct {
	Req    *types.CnsConfigureVolumeACLs         `xml:"urn:vsan CnsConfigureVolumeACLs,omitempty"`
	Res    *types.CnsConfigureVolumeACLsResponse `xml:"urn:vsan CnsConfigureVolumeACLsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsConfigureVolumeACLsBody) Fault() *soap.Fault { return b.Fault_ }

func CnsConfigureVolumeACLs(ctx context.Context, r soap.RoundTripper, req *types.CnsConfigureVolumeACLs) (*types.CnsConfigureVolumeACLsResponse, error) {
	var reqBody, resBody CnsConfigureVolumeACLsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CnsQueryAsyncBody struct {
	Req    *types.CnsQueryAsync         `xml:"urn:vsan CnsQueryAsync,omitempty"`
	Res    *types.CnsQueryAsyncResponse `xml:"urn:vsan CnsQueryAsyncResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQueryAsyncBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQueryAsync(ctx context.Context, r soap.RoundTripper, req *types.CnsQueryAsync) (*types.CnsQueryAsyncResponse, error) {
	var reqBody, resBody CnsQueryAsyncBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

// CNS CreateSnapshots API

type CnsCreateSnapshotsBody struct {
	Req    *types.CnsCreateSnapshots         `xml:"urn:vsan CnsCreateSnapshots,omitempty"`
	Res    *types.CnsCreateSnapshotsResponse `xml:"urn:vsan CnsCreateSnapshotsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsCreateSnapshotsBody) Fault() *soap.Fault { return b.Fault_ }

func CnsCreateSnapshots(ctx context.Context, r soap.RoundTripper, req *types.CnsCreateSnapshots) (*types.CnsCreateSnapshotsResponse, error) {
	var reqBody, resBody CnsCreateSnapshotsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

// CNS DeleteSnapshot API

type CnsDeleteSnapshotBody struct {
	Req    *types.CnsDeleteSnapshots         `xml:"urn:vsan CnsDeleteSnapshots,omitempty"`
	Res    *types.CnsDeleteSnapshotsResponse `xml:"urn:vsan CnsDeleteSnapshotsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsDeleteSnapshotBody) Fault() *soap.Fault { return b.Fault_ }

func CnsDeleteSnapshots(ctx context.Context, r soap.RoundTripper, req *types.CnsDeleteSnapshots) (*types.CnsDeleteSnapshotsResponse, error) {
	var reqBody, resBody CnsDeleteSnapshotBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

// CNS QuerySnapshots API

type CnsQuerySnapshotsBody struct {
	Req    *types.CnsQuerySnapshots         `xml:"urn:vsan CnsQuerySnapshots,omitempty"`
	Res    *types.CnsQuerySnapshotsResponse `xml:"urn:vsan CnsQuerySnapshotsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsQuerySnapshotsBody) Fault() *soap.Fault { return b.Fault_ }

func CnsQuerySnapshots(ctx context.Context, r soap.RoundTripper, req *types.CnsQuerySnapshots) (*types.CnsQuerySnapshotsResponse, error) {
	var reqBody, resBody CnsQuerySnapshotsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
