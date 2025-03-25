// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
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
