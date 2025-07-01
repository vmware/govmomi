// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type PbmQueryIOFiltersFromProfileIdBody struct {
	Req    *types.PbmQueryIOFiltersFromProfileId         `xml:"urn:internalpbm PbmQueryIOFiltersFromProfileId,omitempty"`
	Res    *types.PbmQueryIOFiltersFromProfileIdResponse `xml:"urn:internalpbm PbmQueryIOFiltersFromProfileIdResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PbmQueryIOFiltersFromProfileIdBody) Fault() *soap.Fault { return b.Fault_ }

func PbmQueryIOFiltersFromProfileId(ctx context.Context, r soap.RoundTripper, req *types.PbmQueryIOFiltersFromProfileId) (*types.PbmQueryIOFiltersFromProfileIdResponse, error) {
	var reqBody, resBody PbmQueryIOFiltersFromProfileIdBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PbmQueryProfileDetailsBody struct {
	Req    *types.PbmQueryProfileDetails         `xml:"urn:internalpbm PbmQueryProfileDetails,omitempty"`
	Res    *types.PbmQueryProfileDetailsResponse `xml:"urn:internalpbm PbmQueryProfileDetailsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PbmQueryProfileDetailsBody) Fault() *soap.Fault { return b.Fault_ }

func PbmQueryProfileDetails(ctx context.Context, r soap.RoundTripper, req *types.PbmQueryProfileDetails) (*types.PbmQueryProfileDetailsResponse, error) {
	var reqBody, resBody PbmQueryProfileDetailsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PbmResolveK8sCompliantNamesBody struct {
	Req    *types.PbmResolveK8sCompliantNames         `xml:"urn:internalpbm PbmResolveK8sCompliantNames,omitempty"`
	Res    *types.PbmResolveK8sCompliantNamesResponse `xml:"urn:internalpbm PbmResolveK8sCompliantNamesResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PbmResolveK8sCompliantNamesBody) Fault() *soap.Fault { return b.Fault_ }

func PbmResolveK8sCompliantNames(ctx context.Context, r soap.RoundTripper, req *types.PbmResolveK8sCompliantNames) (*types.PbmResolveK8sCompliantNamesResponse, error) {
	var reqBody, resBody PbmResolveK8sCompliantNamesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PbmUpdateK8sCompliantNamesBody struct {
	Req    *types.PbmUpdateK8sCompliantNames         `xml:"urn:internalpbm PbmUpdateK8sCompliantNames,omitempty"`
	Res    *types.PbmUpdateK8sCompliantNamesResponse `xml:"urn:internalpbm PbmUpdateK8sCompliantNamesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PbmUpdateK8sCompliantNamesBody) Fault() *soap.Fault { return b.Fault_ }

func PbmUpdateK8sCompliantNames(ctx context.Context, r soap.RoundTripper, req *types.PbmUpdateK8sCompliantNames) (*types.PbmUpdateK8sCompliantNamesResponse, error) {
	var reqBody, resBody PbmUpdateK8sCompliantNamesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PbmValidateProfileK8sCompliantNameBody struct {
	Req    *types.PbmValidateProfileK8sCompliantName         `xml:"urn:internalpbm PbmValidateProfileK8sCompliantName,omitempty"`
	Res    *types.PbmValidateProfileK8sCompliantNameResponse `xml:"urn:internalpbm PbmValidateProfileK8sCompliantNameResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PbmValidateProfileK8sCompliantNameBody) Fault() *soap.Fault { return b.Fault_ }

func PbmValidateProfileK8sCompliantName(ctx context.Context, r soap.RoundTripper, req *types.PbmValidateProfileK8sCompliantName) (*types.PbmValidateProfileK8sCompliantNameResponse, error) {
	var reqBody, resBody PbmValidateProfileK8sCompliantNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
