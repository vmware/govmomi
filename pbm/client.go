/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package pbm

import (
	"context"

	"github.com/vmware/govmomi/pbm/methods"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

type Client struct {
	*soap.Client

	ServiceContent types.PbmServiceInstanceContent
}

func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient("/pbm/sdk", "urn:pbm")

	req := types.PbmRetrieveServiceContent{
		This: vim.ManagedObjectReference{
			Type:  "PbmServiceInstance",
			Value: "ServiceInstance",
		},
	}

	res, err := methods.PbmRetrieveServiceContent(ctx, sc, &req)
	if err != nil {
		return nil, err
	}

	return &Client{sc, res.Returnval}, nil
}

func (c *Client) QueryProfile(ctx context.Context, rtype types.PbmProfileResourceType, category string) ([]types.PbmProfileId, error) {
	req := types.PbmQueryProfile{
		This:            c.ServiceContent.ProfileManager,
		ResourceType:    rtype,
		ProfileCategory: category,
	}

	res, err := methods.PbmQueryProfile(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) RetrieveContent(ctx context.Context, ids []types.PbmProfileId) ([]types.BasePbmProfile, error) {
	req := types.PbmRetrieveContent{
		This:       c.ServiceContent.ProfileManager,
		ProfileIds: ids,
	}

	res, err := methods.PbmRetrieveContent(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) CheckRequirements(ctx context.Context, hubs []types.PbmPlacementHub, ref *types.PbmServerObjectRef, preq []types.BasePbmPlacementRequirement) ([]types.PbmPlacementCompatibilityResult, error) {
	req := types.PbmCheckRequirements{
		This:                        c.ServiceContent.PlacementSolver,
		HubsToSearch:                hubs,
		PlacementSubjectRef:         ref,
		PlacementSubjectRequirement: preq,
	}

	res, err := methods.PbmCheckRequirements(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
