// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pbm

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/pbm/methods"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

const (
	Namespace = "pbm"
	Path      = "/pbm"
)

var (
	ServiceInstance = vim.ManagedObjectReference{
		Type:  "PbmServiceInstance",
		Value: "ServiceInstance",
	}
)

type Client struct {
	*soap.Client

	ServiceContent types.PbmServiceInstanceContent

	RoundTripper soap.RoundTripper
}

func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient(Path, Namespace)
	sc.Cookie = c.SessionCookie // vcSessionCookie soap.Header, value must be from vim25.Client

	req := types.PbmRetrieveServiceContent{
		This: ServiceInstance,
	}

	res, err := methods.PbmRetrieveServiceContent(ctx, sc, &req)
	if err != nil {
		return nil, err
	}

	return &Client{sc, res.Returnval, sc}, nil
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	return c.RoundTripper.RoundTrip(ctx, req, res)
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

func (c *Client) QueryProfileDetails(ctx context.Context, category string, fetchAllFields bool) (*types.PbmQueryProfileDetailsResponse, error) {
	req := types.PbmQueryProfileDetails{
		This:            c.ServiceContent.ProfileManager,
		ProfileCategory: category,
		FetchAllFields:  fetchAllFields,
	}

	res, err := methods.PbmQueryProfileDetails(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
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

type PlacementCompatibilityResult []types.PbmPlacementCompatibilityResult

func (c *Client) CheckRequirements(ctx context.Context, hubs []types.PbmPlacementHub, ref *types.PbmServerObjectRef, preq []types.BasePbmPlacementRequirement) (PlacementCompatibilityResult, error) {
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

func (l PlacementCompatibilityResult) CompatibleDatastores() []types.PbmPlacementHub {
	var compatibleDatastores []types.PbmPlacementHub

	for _, res := range l {
		if len(res.Error) == 0 {
			compatibleDatastores = append(compatibleDatastores, res.Hub)
		}
	}
	return compatibleDatastores
}

func (l PlacementCompatibilityResult) NonCompatibleDatastores() []types.PbmPlacementHub {
	var nonCompatibleDatastores []types.PbmPlacementHub

	for _, res := range l {
		if len(res.Error) > 0 {
			nonCompatibleDatastores = append(nonCompatibleDatastores, res.Hub)
		}
	}
	return nonCompatibleDatastores
}

func (c *Client) CreateProfile(ctx context.Context, capabilityProfileCreateSpec types.PbmCapabilityProfileCreateSpec) (*types.PbmProfileId, error) {
	req := types.PbmCreate{
		This:       c.ServiceContent.ProfileManager,
		CreateSpec: capabilityProfileCreateSpec,
	}

	res, err := methods.PbmCreate(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (c *Client) UpdateProfile(ctx context.Context, id types.PbmProfileId, updateSpec types.PbmCapabilityProfileUpdateSpec) error {
	req := types.PbmUpdate{
		This:       c.ServiceContent.ProfileManager,
		ProfileId:  id,
		UpdateSpec: updateSpec,
	}

	_, err := methods.PbmUpdate(ctx, c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteProfile(ctx context.Context, ids []types.PbmProfileId) ([]types.PbmProfileOperationOutcome, error) {
	req := types.PbmDelete{
		This:      c.ServiceContent.ProfileManager,
		ProfileId: ids,
	}

	res, err := methods.PbmDelete(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) QueryAssociatedEntity(ctx context.Context, id types.PbmProfileId, entityType string) ([]types.PbmServerObjectRef, error) {
	req := types.PbmQueryAssociatedEntity{
		This:       c.ServiceContent.ProfileManager,
		Profile:    id,
		EntityType: entityType,
	}

	res, err := methods.PbmQueryAssociatedEntity(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) QueryAssociatedEntities(ctx context.Context, ids []types.PbmProfileId) ([]types.PbmQueryProfileResult, error) {
	req := types.PbmQueryAssociatedEntities{
		This:     c.ServiceContent.ProfileManager,
		Profiles: ids,
	}

	res, err := methods.PbmQueryAssociatedEntities(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) ProfileIDByName(ctx context.Context, profileName string) (string, error) {
	resourceType := types.PbmProfileResourceType{
		ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
	}
	category := types.PbmProfileCategoryEnumREQUIREMENT
	ids, err := c.QueryProfile(ctx, resourceType, string(category))
	if err != nil {
		return "", err
	}

	profiles, err := c.RetrieveContent(ctx, ids)
	if err != nil {
		return "", err
	}

	for i := range profiles {
		profile := profiles[i].GetPbmProfile()
		if profile.Name == profileName {
			return profile.ProfileId.UniqueId, nil
		}
	}
	return "", fmt.Errorf("no pbm profile found with name: %q", profileName)
}

func (c *Client) FetchCapabilityMetadata(ctx context.Context, rtype *types.PbmProfileResourceType, vendorUuid string) ([]types.PbmCapabilityMetadataPerCategory, error) {
	req := types.PbmFetchCapabilityMetadata{
		This:         c.ServiceContent.ProfileManager,
		ResourceType: rtype,
		VendorUuid:   vendorUuid,
	}

	res, err := methods.PbmFetchCapabilityMetadata(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) FetchComplianceResult(ctx context.Context, entities []types.PbmServerObjectRef) ([]types.PbmComplianceResult, error) {
	req := types.PbmFetchComplianceResult{
		This:     c.ServiceContent.ComplianceManager,
		Entities: entities,
	}

	res, err := methods.PbmFetchComplianceResult(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// GetProfileNameByID gets storage profile name by ID
func (c *Client) GetProfileNameByID(ctx context.Context, profileID string) (string, error) {
	resourceType := types.PbmProfileResourceType{
		ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
	}
	category := types.PbmProfileCategoryEnumREQUIREMENT
	ids, err := c.QueryProfile(ctx, resourceType, string(category))
	if err != nil {
		return "", err
	}

	profiles, err := c.RetrieveContent(ctx, ids)
	if err != nil {
		return "", err
	}

	for i := range profiles {
		profile := profiles[i].GetPbmProfile()
		if profile.ProfileId.UniqueId == profileID {
			return profile.Name, nil
		}
	}
	return "", fmt.Errorf("no pbm profile found with id: %q", profileID)
}

func (c *Client) QueryAssociatedProfile(ctx context.Context, entity types.PbmServerObjectRef) ([]types.PbmProfileId, error) {
	req := types.PbmQueryAssociatedProfile{
		This:   c.ServiceContent.ProfileManager,
		Entity: entity,
	}

	res, err := methods.PbmQueryAssociatedProfile(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) QueryAssociatedProfiles(ctx context.Context, entities []types.PbmServerObjectRef) ([]types.PbmQueryProfileResult, error) {
	req := types.PbmQueryAssociatedProfiles{
		This:     c.ServiceContent.ProfileManager,
		Entities: entities,
	}

	res, err := methods.PbmQueryAssociatedProfiles(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (c *Client) QueryIOFiltersFromProfileId(
	ctx context.Context, profileID string) ([]types.PbmProfileToIofilterMap, error) {

	req := types.PbmQueryIOFiltersFromProfileId{
		This:       c.ServiceContent.ProfileManager,
		ProfileIds: []types.PbmProfileId{{UniqueId: profileID}},
	}
	res, err := methods.PbmQueryIOFiltersFromProfileId(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

func (c *Client) SupportsEncryption(
	ctx context.Context, profileID string) (bool, error) {

	list, err := c.QueryIOFiltersFromProfileId(ctx, profileID)
	if err != nil {
		return false, err
	}
	for i := range list {
		for j := range list[i].Iofilters {
			f := list[i].Iofilters[j]
			if f.FilterType == string(types.PbmIofilterInfoFilterTypeENCRYPTION) {
				return true, nil
			}
		}
	}
	return false, nil
}
func (c *Client) ResolveK8sCompliantNames(ctx context.Context) error {
	req := types.PbmResolveK8sCompliantNames{
		This: c.ServiceContent.ProfileManager,
	}

	_, err := methods.PbmResolveK8sCompliantNames(ctx, c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateK8sCompliantNames(ctx context.Context, profileID string,
	k8sCompliantName string, otherK8sCompliantNames []string) error {
	req := types.PbmUpdateK8sCompliantNames{
		This: c.ServiceContent.ProfileManager,
		K8sCompliantNameSpec: types.PbmProfileK8sCompliantNameSpec{
			ProfileId:              profileID,
			K8sCompliantName:       k8sCompliantName,
			OtherK8sCompliantNames: otherK8sCompliantNames,
		},
	}

	_, err := methods.PbmUpdateK8sCompliantNames(ctx, c, &req)
	if err != nil {
		return err
	}

	return nil
}
