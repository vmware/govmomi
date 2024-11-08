/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
