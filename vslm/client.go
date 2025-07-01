// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	sc.Cookie = c.SessionCookie // vcSessionCookie soap.Header, value must be from vim25.Client

	req := types.RetrieveContent{
		This: ServiceInstance,
	}

	res, err := methods.RetrieveContent(ctx, sc, &req)
	if err != nil {
		return nil, err
	}

	return &Client{sc, res.Returnval}, nil
}
