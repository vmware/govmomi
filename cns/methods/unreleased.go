// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type CnsUpdateVolumeCryptoBody struct {
	Req    *types.CnsUpdateVolumeCrypto         `xml:"urn:vsan CnsUpdateVolumeCrypto,omitempty"`
	Res    *types.CnsUpdateVolumeCryptoResponse `xml:"urn:vsan CnsUpdateVolumeCryptoResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CnsUpdateVolumeCryptoBody) Fault() *soap.Fault { return b.Fault_ }

func CnsUpdateVolumeCrypto(ctx context.Context, r soap.RoundTripper, req *types.CnsUpdateVolumeCrypto) (*types.CnsUpdateVolumeCryptoResponse, error) {
	var reqBody, resBody CnsUpdateVolumeCryptoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
