/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
