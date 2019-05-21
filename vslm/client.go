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

package vslm

import (
	"context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm/methods"
	"github.com/vmware/govmomi/vslm/types"
)

const (
	Namespace = "vslm"
	Path      = "/vslm/sdk"
)

var (
	ServiceInstance = vim.ManagedObjectReference{
		Type:  "VslmServiceInstance",
		Value: "ServiceInstance",
	}
)

type Client struct {
	*soap.Client

	ServiceContent types.VslmServiceInstanceContent
}

func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient(Path, Namespace)
	req := types.RetrieveContent{
		This: ServiceInstance,
	}

	res, err := methods.RetrieveContent(ctx, sc, &req)
	if err != nil {
		return nil, err
	}

	return &Client{sc, res.Returnval}, nil
}
