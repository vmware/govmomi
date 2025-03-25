// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	_ "github.com/vmware/govmomi/vsan/types"
)

func RetrieveProperties(ctx context.Context, r soap.RoundTripper, req *types.RetrieveProperties) (*types.RetrievePropertiesResponse, error) {
	return methods.RetrieveProperties(ctx, r, req)
}

func RetrievePropertiesEx(ctx context.Context, r soap.RoundTripper, req *types.RetrievePropertiesEx) (*types.RetrievePropertiesExResponse, error) {
	return methods.RetrievePropertiesEx(ctx, r, req)
}
