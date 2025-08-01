// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type AbandonHciWorkflowBody struct {
	Req    *types.AbandonHciWorkflow         `xml:"urn:vim25 AbandonHciWorkflow,omitempty"`
	Res    *types.AbandonHciWorkflowResponse `xml:"AbandonHciWorkflowResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AbandonHciWorkflowBody) Fault() *soap.Fault { return b.Fault_ }

func AbandonHciWorkflow(ctx context.Context, r soap.RoundTripper, req *types.AbandonHciWorkflow) (*types.AbandonHciWorkflowResponse, error) {
	var reqBody, resBody AbandonHciWorkflowBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AbdicateDomOwnershipBody struct {
	Req    *types.AbdicateDomOwnership         `xml:"urn:vim25 AbdicateDomOwnership,omitempty"`
	Res    *types.AbdicateDomOwnershipResponse `xml:"AbdicateDomOwnershipResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AbdicateDomOwnershipBody) Fault() *soap.Fault { return b.Fault_ }

func AbdicateDomOwnership(ctx context.Context, r soap.RoundTripper, req *types.AbdicateDomOwnership) (*types.AbdicateDomOwnershipResponse, error) {
	var reqBody, resBody AbdicateDomOwnershipBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AbortCustomization_TaskBody struct {
	Req    *types.AbortCustomization_Task         `xml:"urn:vim25 AbortCustomization_Task,omitempty"`
	Res    *types.AbortCustomization_TaskResponse `xml:"AbortCustomization_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AbortCustomization_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AbortCustomization_Task(ctx context.Context, r soap.RoundTripper, req *types.AbortCustomization_Task) (*types.AbortCustomization_TaskResponse, error) {
	var reqBody, resBody AbortCustomization_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcknowledgeAlarmBody struct {
	Req    *types.AcknowledgeAlarm         `xml:"urn:vim25 AcknowledgeAlarm,omitempty"`
	Res    *types.AcknowledgeAlarmResponse `xml:"AcknowledgeAlarmResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcknowledgeAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func AcknowledgeAlarm(ctx context.Context, r soap.RoundTripper, req *types.AcknowledgeAlarm) (*types.AcknowledgeAlarmResponse, error) {
	var reqBody, resBody AcknowledgeAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireCimServicesTicketBody struct {
	Req    *types.AcquireCimServicesTicket         `xml:"urn:vim25 AcquireCimServicesTicket,omitempty"`
	Res    *types.AcquireCimServicesTicketResponse `xml:"AcquireCimServicesTicketResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCimServicesTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireCimServicesTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireCimServicesTicket) (*types.AcquireCimServicesTicketResponse, error) {
	var reqBody, resBody AcquireCimServicesTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireCloneTicketBody struct {
	Req    *types.AcquireCloneTicket         `xml:"urn:vim25 AcquireCloneTicket,omitempty"`
	Res    *types.AcquireCloneTicketResponse `xml:"AcquireCloneTicketResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCloneTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireCloneTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireCloneTicket) (*types.AcquireCloneTicketResponse, error) {
	var reqBody, resBody AcquireCloneTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireCredentialsInGuestBody struct {
	Req    *types.AcquireCredentialsInGuest         `xml:"urn:vim25 AcquireCredentialsInGuest,omitempty"`
	Res    *types.AcquireCredentialsInGuestResponse `xml:"AcquireCredentialsInGuestResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCredentialsInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireCredentialsInGuest(ctx context.Context, r soap.RoundTripper, req *types.AcquireCredentialsInGuest) (*types.AcquireCredentialsInGuestResponse, error) {
	var reqBody, resBody AcquireCredentialsInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireGenericServiceTicketBody struct {
	Req    *types.AcquireGenericServiceTicket         `xml:"urn:vim25 AcquireGenericServiceTicket,omitempty"`
	Res    *types.AcquireGenericServiceTicketResponse `xml:"AcquireGenericServiceTicketResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireGenericServiceTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireGenericServiceTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireGenericServiceTicket) (*types.AcquireGenericServiceTicketResponse, error) {
	var reqBody, resBody AcquireGenericServiceTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireLocalTicketBody struct {
	Req    *types.AcquireLocalTicket         `xml:"urn:vim25 AcquireLocalTicket,omitempty"`
	Res    *types.AcquireLocalTicketResponse `xml:"AcquireLocalTicketResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireLocalTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireLocalTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireLocalTicket) (*types.AcquireLocalTicketResponse, error) {
	var reqBody, resBody AcquireLocalTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireMksTicketBody struct {
	Req    *types.AcquireMksTicket         `xml:"urn:vim25 AcquireMksTicket,omitempty"`
	Res    *types.AcquireMksTicketResponse `xml:"AcquireMksTicketResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireMksTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireMksTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireMksTicket) (*types.AcquireMksTicketResponse, error) {
	var reqBody, resBody AcquireMksTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AcquireTicketBody struct {
	Req    *types.AcquireTicket         `xml:"urn:vim25 AcquireTicket,omitempty"`
	Res    *types.AcquireTicketResponse `xml:"AcquireTicketResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireTicketBody) Fault() *soap.Fault { return b.Fault_ }

func AcquireTicket(ctx context.Context, r soap.RoundTripper, req *types.AcquireTicket) (*types.AcquireTicketResponse, error) {
	var reqBody, resBody AcquireTicketBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddAuthorizationRoleBody struct {
	Req    *types.AddAuthorizationRole         `xml:"urn:vim25 AddAuthorizationRole,omitempty"`
	Res    *types.AddAuthorizationRoleResponse `xml:"AddAuthorizationRoleResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddAuthorizationRoleBody) Fault() *soap.Fault { return b.Fault_ }

func AddAuthorizationRole(ctx context.Context, r soap.RoundTripper, req *types.AddAuthorizationRole) (*types.AddAuthorizationRoleResponse, error) {
	var reqBody, resBody AddAuthorizationRoleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddCustomFieldDefBody struct {
	Req    *types.AddCustomFieldDef         `xml:"urn:vim25 AddCustomFieldDef,omitempty"`
	Res    *types.AddCustomFieldDefResponse `xml:"AddCustomFieldDefResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddCustomFieldDefBody) Fault() *soap.Fault { return b.Fault_ }

func AddCustomFieldDef(ctx context.Context, r soap.RoundTripper, req *types.AddCustomFieldDef) (*types.AddCustomFieldDefResponse, error) {
	var reqBody, resBody AddCustomFieldDefBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddDVPortgroup_TaskBody struct {
	Req    *types.AddDVPortgroup_Task         `xml:"urn:vim25 AddDVPortgroup_Task,omitempty"`
	Res    *types.AddDVPortgroup_TaskResponse `xml:"AddDVPortgroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddDVPortgroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AddDVPortgroup_Task(ctx context.Context, r soap.RoundTripper, req *types.AddDVPortgroup_Task) (*types.AddDVPortgroup_TaskResponse, error) {
	var reqBody, resBody AddDVPortgroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddDisks_TaskBody struct {
	Req    *types.AddDisks_Task         `xml:"urn:vim25 AddDisks_Task,omitempty"`
	Res    *types.AddDisks_TaskResponse `xml:"AddDisks_TaskResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddDisks_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AddDisks_Task(ctx context.Context, r soap.RoundTripper, req *types.AddDisks_Task) (*types.AddDisks_TaskResponse, error) {
	var reqBody, resBody AddDisks_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddFilterBody struct {
	Req    *types.AddFilter         `xml:"urn:vim25 AddFilter,omitempty"`
	Res    *types.AddFilterResponse `xml:"AddFilterResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddFilterBody) Fault() *soap.Fault { return b.Fault_ }

func AddFilter(ctx context.Context, r soap.RoundTripper, req *types.AddFilter) (*types.AddFilterResponse, error) {
	var reqBody, resBody AddFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddFilterEntitiesBody struct {
	Req    *types.AddFilterEntities         `xml:"urn:vim25 AddFilterEntities,omitempty"`
	Res    *types.AddFilterEntitiesResponse `xml:"AddFilterEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddFilterEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func AddFilterEntities(ctx context.Context, r soap.RoundTripper, req *types.AddFilterEntities) (*types.AddFilterEntitiesResponse, error) {
	var reqBody, resBody AddFilterEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddGuestAliasBody struct {
	Req    *types.AddGuestAlias         `xml:"urn:vim25 AddGuestAlias,omitempty"`
	Res    *types.AddGuestAliasResponse `xml:"AddGuestAliasResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddGuestAliasBody) Fault() *soap.Fault { return b.Fault_ }

func AddGuestAlias(ctx context.Context, r soap.RoundTripper, req *types.AddGuestAlias) (*types.AddGuestAliasResponse, error) {
	var reqBody, resBody AddGuestAliasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddHost_TaskBody struct {
	Req    *types.AddHost_Task         `xml:"urn:vim25 AddHost_Task,omitempty"`
	Res    *types.AddHost_TaskResponse `xml:"AddHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AddHost_Task(ctx context.Context, r soap.RoundTripper, req *types.AddHost_Task) (*types.AddHost_TaskResponse, error) {
	var reqBody, resBody AddHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddInternetScsiSendTargetsBody struct {
	Req    *types.AddInternetScsiSendTargets         `xml:"urn:vim25 AddInternetScsiSendTargets,omitempty"`
	Res    *types.AddInternetScsiSendTargetsResponse `xml:"AddInternetScsiSendTargetsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddInternetScsiSendTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func AddInternetScsiSendTargets(ctx context.Context, r soap.RoundTripper, req *types.AddInternetScsiSendTargets) (*types.AddInternetScsiSendTargetsResponse, error) {
	var reqBody, resBody AddInternetScsiSendTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddInternetScsiStaticTargetsBody struct {
	Req    *types.AddInternetScsiStaticTargets         `xml:"urn:vim25 AddInternetScsiStaticTargets,omitempty"`
	Res    *types.AddInternetScsiStaticTargetsResponse `xml:"AddInternetScsiStaticTargetsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddInternetScsiStaticTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func AddInternetScsiStaticTargets(ctx context.Context, r soap.RoundTripper, req *types.AddInternetScsiStaticTargets) (*types.AddInternetScsiStaticTargetsResponse, error) {
	var reqBody, resBody AddInternetScsiStaticTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddKeyBody struct {
	Req    *types.AddKey         `xml:"urn:vim25 AddKey,omitempty"`
	Res    *types.AddKeyResponse `xml:"AddKeyResponse,omitempty"`
	Fault_ *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddKeyBody) Fault() *soap.Fault { return b.Fault_ }

func AddKey(ctx context.Context, r soap.RoundTripper, req *types.AddKey) (*types.AddKeyResponse, error) {
	var reqBody, resBody AddKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddKeysBody struct {
	Req    *types.AddKeys         `xml:"urn:vim25 AddKeys,omitempty"`
	Res    *types.AddKeysResponse `xml:"AddKeysResponse,omitempty"`
	Fault_ *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddKeysBody) Fault() *soap.Fault { return b.Fault_ }

func AddKeys(ctx context.Context, r soap.RoundTripper, req *types.AddKeys) (*types.AddKeysResponse, error) {
	var reqBody, resBody AddKeysBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddLicenseBody struct {
	Req    *types.AddLicense         `xml:"urn:vim25 AddLicense,omitempty"`
	Res    *types.AddLicenseResponse `xml:"AddLicenseResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func AddLicense(ctx context.Context, r soap.RoundTripper, req *types.AddLicense) (*types.AddLicenseResponse, error) {
	var reqBody, resBody AddLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddMonitoredEntitiesBody struct {
	Req    *types.AddMonitoredEntities         `xml:"urn:vim25 AddMonitoredEntities,omitempty"`
	Res    *types.AddMonitoredEntitiesResponse `xml:"AddMonitoredEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddMonitoredEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func AddMonitoredEntities(ctx context.Context, r soap.RoundTripper, req *types.AddMonitoredEntities) (*types.AddMonitoredEntitiesResponse, error) {
	var reqBody, resBody AddMonitoredEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddNetworkResourcePoolBody struct {
	Req    *types.AddNetworkResourcePool         `xml:"urn:vim25 AddNetworkResourcePool,omitempty"`
	Res    *types.AddNetworkResourcePoolResponse `xml:"AddNetworkResourcePoolResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddNetworkResourcePoolBody) Fault() *soap.Fault { return b.Fault_ }

func AddNetworkResourcePool(ctx context.Context, r soap.RoundTripper, req *types.AddNetworkResourcePool) (*types.AddNetworkResourcePoolResponse, error) {
	var reqBody, resBody AddNetworkResourcePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddPortGroupBody struct {
	Req    *types.AddPortGroup         `xml:"urn:vim25 AddPortGroup,omitempty"`
	Res    *types.AddPortGroupResponse `xml:"AddPortGroupResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddPortGroupBody) Fault() *soap.Fault { return b.Fault_ }

func AddPortGroup(ctx context.Context, r soap.RoundTripper, req *types.AddPortGroup) (*types.AddPortGroupResponse, error) {
	var reqBody, resBody AddPortGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddServiceConsoleVirtualNicBody struct {
	Req    *types.AddServiceConsoleVirtualNic         `xml:"urn:vim25 AddServiceConsoleVirtualNic,omitempty"`
	Res    *types.AddServiceConsoleVirtualNicResponse `xml:"AddServiceConsoleVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddServiceConsoleVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func AddServiceConsoleVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.AddServiceConsoleVirtualNic) (*types.AddServiceConsoleVirtualNicResponse, error) {
	var reqBody, resBody AddServiceConsoleVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddStandaloneHost_TaskBody struct {
	Req    *types.AddStandaloneHost_Task         `xml:"urn:vim25 AddStandaloneHost_Task,omitempty"`
	Res    *types.AddStandaloneHost_TaskResponse `xml:"AddStandaloneHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddStandaloneHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AddStandaloneHost_Task(ctx context.Context, r soap.RoundTripper, req *types.AddStandaloneHost_Task) (*types.AddStandaloneHost_TaskResponse, error) {
	var reqBody, resBody AddStandaloneHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddVirtualNicBody struct {
	Req    *types.AddVirtualNic         `xml:"urn:vim25 AddVirtualNic,omitempty"`
	Res    *types.AddVirtualNicResponse `xml:"AddVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func AddVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.AddVirtualNic) (*types.AddVirtualNicResponse, error) {
	var reqBody, resBody AddVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AddVirtualSwitchBody struct {
	Req    *types.AddVirtualSwitch         `xml:"urn:vim25 AddVirtualSwitch,omitempty"`
	Res    *types.AddVirtualSwitchResponse `xml:"AddVirtualSwitchResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddVirtualSwitchBody) Fault() *soap.Fault { return b.Fault_ }

func AddVirtualSwitch(ctx context.Context, r soap.RoundTripper, req *types.AddVirtualSwitch) (*types.AddVirtualSwitchResponse, error) {
	var reqBody, resBody AddVirtualSwitchBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AllocateIpv4AddressBody struct {
	Req    *types.AllocateIpv4Address         `xml:"urn:vim25 AllocateIpv4Address,omitempty"`
	Res    *types.AllocateIpv4AddressResponse `xml:"AllocateIpv4AddressResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AllocateIpv4AddressBody) Fault() *soap.Fault { return b.Fault_ }

func AllocateIpv4Address(ctx context.Context, r soap.RoundTripper, req *types.AllocateIpv4Address) (*types.AllocateIpv4AddressResponse, error) {
	var reqBody, resBody AllocateIpv4AddressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AllocateIpv6AddressBody struct {
	Req    *types.AllocateIpv6Address         `xml:"urn:vim25 AllocateIpv6Address,omitempty"`
	Res    *types.AllocateIpv6AddressResponse `xml:"AllocateIpv6AddressResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AllocateIpv6AddressBody) Fault() *soap.Fault { return b.Fault_ }

func AllocateIpv6Address(ctx context.Context, r soap.RoundTripper, req *types.AllocateIpv6Address) (*types.AllocateIpv6AddressResponse, error) {
	var reqBody, resBody AllocateIpv6AddressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AnswerVMBody struct {
	Req    *types.AnswerVM         `xml:"urn:vim25 AnswerVM,omitempty"`
	Res    *types.AnswerVMResponse `xml:"AnswerVMResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AnswerVMBody) Fault() *soap.Fault { return b.Fault_ }

func AnswerVM(ctx context.Context, r soap.RoundTripper, req *types.AnswerVM) (*types.AnswerVMResponse, error) {
	var reqBody, resBody AnswerVMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyEntitiesConfig_TaskBody struct {
	Req    *types.ApplyEntitiesConfig_Task         `xml:"urn:vim25 ApplyEntitiesConfig_Task,omitempty"`
	Res    *types.ApplyEntitiesConfig_TaskResponse `xml:"ApplyEntitiesConfig_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyEntitiesConfig_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyEntitiesConfig_Task(ctx context.Context, r soap.RoundTripper, req *types.ApplyEntitiesConfig_Task) (*types.ApplyEntitiesConfig_TaskResponse, error) {
	var reqBody, resBody ApplyEntitiesConfig_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyEvcModeVM_TaskBody struct {
	Req    *types.ApplyEvcModeVM_Task         `xml:"urn:vim25 ApplyEvcModeVM_Task,omitempty"`
	Res    *types.ApplyEvcModeVM_TaskResponse `xml:"ApplyEvcModeVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyEvcModeVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyEvcModeVM_Task(ctx context.Context, r soap.RoundTripper, req *types.ApplyEvcModeVM_Task) (*types.ApplyEvcModeVM_TaskResponse, error) {
	var reqBody, resBody ApplyEvcModeVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyHostConfig_TaskBody struct {
	Req    *types.ApplyHostConfig_Task         `xml:"urn:vim25 ApplyHostConfig_Task,omitempty"`
	Res    *types.ApplyHostConfig_TaskResponse `xml:"ApplyHostConfig_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyHostConfig_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyHostConfig_Task(ctx context.Context, r soap.RoundTripper, req *types.ApplyHostConfig_Task) (*types.ApplyHostConfig_TaskResponse, error) {
	var reqBody, resBody ApplyHostConfig_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyRecommendationBody struct {
	Req    *types.ApplyRecommendation         `xml:"urn:vim25 ApplyRecommendation,omitempty"`
	Res    *types.ApplyRecommendationResponse `xml:"ApplyRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyRecommendation(ctx context.Context, r soap.RoundTripper, req *types.ApplyRecommendation) (*types.ApplyRecommendationResponse, error) {
	var reqBody, resBody ApplyRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyStorageDrsRecommendationToPod_TaskBody struct {
	Req    *types.ApplyStorageDrsRecommendationToPod_Task         `xml:"urn:vim25 ApplyStorageDrsRecommendationToPod_Task,omitempty"`
	Res    *types.ApplyStorageDrsRecommendationToPod_TaskResponse `xml:"ApplyStorageDrsRecommendationToPod_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyStorageDrsRecommendationToPod_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyStorageDrsRecommendationToPod_Task(ctx context.Context, r soap.RoundTripper, req *types.ApplyStorageDrsRecommendationToPod_Task) (*types.ApplyStorageDrsRecommendationToPod_TaskResponse, error) {
	var reqBody, resBody ApplyStorageDrsRecommendationToPod_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ApplyStorageDrsRecommendation_TaskBody struct {
	Req    *types.ApplyStorageDrsRecommendation_Task         `xml:"urn:vim25 ApplyStorageDrsRecommendation_Task,omitempty"`
	Res    *types.ApplyStorageDrsRecommendation_TaskResponse `xml:"ApplyStorageDrsRecommendation_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyStorageDrsRecommendation_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ApplyStorageDrsRecommendation_Task(ctx context.Context, r soap.RoundTripper, req *types.ApplyStorageDrsRecommendation_Task) (*types.ApplyStorageDrsRecommendation_TaskResponse, error) {
	var reqBody, resBody ApplyStorageDrsRecommendation_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AreAlarmActionsEnabledBody struct {
	Req    *types.AreAlarmActionsEnabled         `xml:"urn:vim25 AreAlarmActionsEnabled,omitempty"`
	Res    *types.AreAlarmActionsEnabledResponse `xml:"AreAlarmActionsEnabledResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AreAlarmActionsEnabledBody) Fault() *soap.Fault { return b.Fault_ }

func AreAlarmActionsEnabled(ctx context.Context, r soap.RoundTripper, req *types.AreAlarmActionsEnabled) (*types.AreAlarmActionsEnabledResponse, error) {
	var reqBody, resBody AreAlarmActionsEnabledBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AssignUserToGroupBody struct {
	Req    *types.AssignUserToGroup         `xml:"urn:vim25 AssignUserToGroup,omitempty"`
	Res    *types.AssignUserToGroupResponse `xml:"AssignUserToGroupResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AssignUserToGroupBody) Fault() *soap.Fault { return b.Fault_ }

func AssignUserToGroup(ctx context.Context, r soap.RoundTripper, req *types.AssignUserToGroup) (*types.AssignUserToGroupResponse, error) {
	var reqBody, resBody AssignUserToGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AssociateProfileBody struct {
	Req    *types.AssociateProfile         `xml:"urn:vim25 AssociateProfile,omitempty"`
	Res    *types.AssociateProfileResponse `xml:"AssociateProfileResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AssociateProfileBody) Fault() *soap.Fault { return b.Fault_ }

func AssociateProfile(ctx context.Context, r soap.RoundTripper, req *types.AssociateProfile) (*types.AssociateProfileResponse, error) {
	var reqBody, resBody AssociateProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AttachDisk_TaskBody struct {
	Req    *types.AttachDisk_Task         `xml:"urn:vim25 AttachDisk_Task,omitempty"`
	Res    *types.AttachDisk_TaskResponse `xml:"AttachDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AttachDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.AttachDisk_Task) (*types.AttachDisk_TaskResponse, error) {
	var reqBody, resBody AttachDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AttachScsiLunBody struct {
	Req    *types.AttachScsiLun         `xml:"urn:vim25 AttachScsiLun,omitempty"`
	Res    *types.AttachScsiLunResponse `xml:"AttachScsiLunResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachScsiLunBody) Fault() *soap.Fault { return b.Fault_ }

func AttachScsiLun(ctx context.Context, r soap.RoundTripper, req *types.AttachScsiLun) (*types.AttachScsiLunResponse, error) {
	var reqBody, resBody AttachScsiLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AttachScsiLunEx_TaskBody struct {
	Req    *types.AttachScsiLunEx_Task         `xml:"urn:vim25 AttachScsiLunEx_Task,omitempty"`
	Res    *types.AttachScsiLunEx_TaskResponse `xml:"AttachScsiLunEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachScsiLunEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func AttachScsiLunEx_Task(ctx context.Context, r soap.RoundTripper, req *types.AttachScsiLunEx_Task) (*types.AttachScsiLunEx_TaskResponse, error) {
	var reqBody, resBody AttachScsiLunEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AttachTagToVStorageObjectBody struct {
	Req    *types.AttachTagToVStorageObject         `xml:"urn:vim25 AttachTagToVStorageObject,omitempty"`
	Res    *types.AttachTagToVStorageObjectResponse `xml:"AttachTagToVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachTagToVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func AttachTagToVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.AttachTagToVStorageObject) (*types.AttachTagToVStorageObjectResponse, error) {
	var reqBody, resBody AttachTagToVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AttachVmfsExtentBody struct {
	Req    *types.AttachVmfsExtent         `xml:"urn:vim25 AttachVmfsExtent,omitempty"`
	Res    *types.AttachVmfsExtentResponse `xml:"AttachVmfsExtentResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachVmfsExtentBody) Fault() *soap.Fault { return b.Fault_ }

func AttachVmfsExtent(ctx context.Context, r soap.RoundTripper, req *types.AttachVmfsExtent) (*types.AttachVmfsExtentResponse, error) {
	var reqBody, resBody AttachVmfsExtentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AutoStartPowerOffBody struct {
	Req    *types.AutoStartPowerOff         `xml:"urn:vim25 AutoStartPowerOff,omitempty"`
	Res    *types.AutoStartPowerOffResponse `xml:"AutoStartPowerOffResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AutoStartPowerOffBody) Fault() *soap.Fault { return b.Fault_ }

func AutoStartPowerOff(ctx context.Context, r soap.RoundTripper, req *types.AutoStartPowerOff) (*types.AutoStartPowerOffResponse, error) {
	var reqBody, resBody AutoStartPowerOffBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AutoStartPowerOnBody struct {
	Req    *types.AutoStartPowerOn         `xml:"urn:vim25 AutoStartPowerOn,omitempty"`
	Res    *types.AutoStartPowerOnResponse `xml:"AutoStartPowerOnResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AutoStartPowerOnBody) Fault() *soap.Fault { return b.Fault_ }

func AutoStartPowerOn(ctx context.Context, r soap.RoundTripper, req *types.AutoStartPowerOn) (*types.AutoStartPowerOnResponse, error) {
	var reqBody, resBody AutoStartPowerOnBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BackupFirmwareConfigurationBody struct {
	Req    *types.BackupFirmwareConfiguration         `xml:"urn:vim25 BackupFirmwareConfiguration,omitempty"`
	Res    *types.BackupFirmwareConfigurationResponse `xml:"BackupFirmwareConfigurationResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BackupFirmwareConfigurationBody) Fault() *soap.Fault { return b.Fault_ }

func BackupFirmwareConfiguration(ctx context.Context, r soap.RoundTripper, req *types.BackupFirmwareConfiguration) (*types.BackupFirmwareConfigurationResponse, error) {
	var reqBody, resBody BackupFirmwareConfigurationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BatchAddHostsToCluster_TaskBody struct {
	Req    *types.BatchAddHostsToCluster_Task         `xml:"urn:vim25 BatchAddHostsToCluster_Task,omitempty"`
	Res    *types.BatchAddHostsToCluster_TaskResponse `xml:"BatchAddHostsToCluster_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BatchAddHostsToCluster_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func BatchAddHostsToCluster_Task(ctx context.Context, r soap.RoundTripper, req *types.BatchAddHostsToCluster_Task) (*types.BatchAddHostsToCluster_TaskResponse, error) {
	var reqBody, resBody BatchAddHostsToCluster_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BatchAddStandaloneHosts_TaskBody struct {
	Req    *types.BatchAddStandaloneHosts_Task         `xml:"urn:vim25 BatchAddStandaloneHosts_Task,omitempty"`
	Res    *types.BatchAddStandaloneHosts_TaskResponse `xml:"BatchAddStandaloneHosts_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BatchAddStandaloneHosts_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func BatchAddStandaloneHosts_Task(ctx context.Context, r soap.RoundTripper, req *types.BatchAddStandaloneHosts_Task) (*types.BatchAddStandaloneHosts_TaskResponse, error) {
	var reqBody, resBody BatchAddStandaloneHosts_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BatchQueryConnectInfoBody struct {
	Req    *types.BatchQueryConnectInfo         `xml:"urn:vim25 BatchQueryConnectInfo,omitempty"`
	Res    *types.BatchQueryConnectInfoResponse `xml:"BatchQueryConnectInfoResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BatchQueryConnectInfoBody) Fault() *soap.Fault { return b.Fault_ }

func BatchQueryConnectInfo(ctx context.Context, r soap.RoundTripper, req *types.BatchQueryConnectInfo) (*types.BatchQueryConnectInfoResponse, error) {
	var reqBody, resBody BatchQueryConnectInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BindVnicBody struct {
	Req    *types.BindVnic         `xml:"urn:vim25 BindVnic,omitempty"`
	Res    *types.BindVnicResponse `xml:"BindVnicResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BindVnicBody) Fault() *soap.Fault { return b.Fault_ }

func BindVnic(ctx context.Context, r soap.RoundTripper, req *types.BindVnic) (*types.BindVnicResponse, error) {
	var reqBody, resBody BindVnicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type BrowseDiagnosticLogBody struct {
	Req    *types.BrowseDiagnosticLog         `xml:"urn:vim25 BrowseDiagnosticLog,omitempty"`
	Res    *types.BrowseDiagnosticLogResponse `xml:"BrowseDiagnosticLogResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BrowseDiagnosticLogBody) Fault() *soap.Fault { return b.Fault_ }

func BrowseDiagnosticLog(ctx context.Context, r soap.RoundTripper, req *types.BrowseDiagnosticLog) (*types.BrowseDiagnosticLogResponse, error) {
	var reqBody, resBody BrowseDiagnosticLogBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CanProvisionObjectsBody struct {
	Req    *types.CanProvisionObjects         `xml:"urn:vim25 CanProvisionObjects,omitempty"`
	Res    *types.CanProvisionObjectsResponse `xml:"CanProvisionObjectsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CanProvisionObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func CanProvisionObjects(ctx context.Context, r soap.RoundTripper, req *types.CanProvisionObjects) (*types.CanProvisionObjectsResponse, error) {
	var reqBody, resBody CanProvisionObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CancelRecommendationBody struct {
	Req    *types.CancelRecommendation         `xml:"urn:vim25 CancelRecommendation,omitempty"`
	Res    *types.CancelRecommendationResponse `xml:"CancelRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func CancelRecommendation(ctx context.Context, r soap.RoundTripper, req *types.CancelRecommendation) (*types.CancelRecommendationResponse, error) {
	var reqBody, resBody CancelRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CancelRetrievePropertiesExBody struct {
	Req    *types.CancelRetrievePropertiesEx         `xml:"urn:vim25 CancelRetrievePropertiesEx,omitempty"`
	Res    *types.CancelRetrievePropertiesExResponse `xml:"CancelRetrievePropertiesExResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelRetrievePropertiesExBody) Fault() *soap.Fault { return b.Fault_ }

func CancelRetrievePropertiesEx(ctx context.Context, r soap.RoundTripper, req *types.CancelRetrievePropertiesEx) (*types.CancelRetrievePropertiesExResponse, error) {
	var reqBody, resBody CancelRetrievePropertiesExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CancelStorageDrsRecommendationBody struct {
	Req    *types.CancelStorageDrsRecommendation         `xml:"urn:vim25 CancelStorageDrsRecommendation,omitempty"`
	Res    *types.CancelStorageDrsRecommendationResponse `xml:"CancelStorageDrsRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelStorageDrsRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func CancelStorageDrsRecommendation(ctx context.Context, r soap.RoundTripper, req *types.CancelStorageDrsRecommendation) (*types.CancelStorageDrsRecommendationResponse, error) {
	var reqBody, resBody CancelStorageDrsRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CancelTaskBody struct {
	Req    *types.CancelTask         `xml:"urn:vim25 CancelTask,omitempty"`
	Res    *types.CancelTaskResponse `xml:"CancelTaskResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelTaskBody) Fault() *soap.Fault { return b.Fault_ }

func CancelTask(ctx context.Context, r soap.RoundTripper, req *types.CancelTask) (*types.CancelTaskResponse, error) {
	var reqBody, resBody CancelTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CancelWaitForUpdatesBody struct {
	Req    *types.CancelWaitForUpdates         `xml:"urn:vim25 CancelWaitForUpdates,omitempty"`
	Res    *types.CancelWaitForUpdatesResponse `xml:"CancelWaitForUpdatesResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelWaitForUpdatesBody) Fault() *soap.Fault { return b.Fault_ }

func CancelWaitForUpdates(ctx context.Context, r soap.RoundTripper, req *types.CancelWaitForUpdates) (*types.CancelWaitForUpdatesResponse, error) {
	var reqBody, resBody CancelWaitForUpdatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CertMgrRefreshCACertificatesAndCRLs_TaskBody struct {
	Req    *types.CertMgrRefreshCACertificatesAndCRLs_Task         `xml:"urn:vim25 CertMgrRefreshCACertificatesAndCRLs_Task,omitempty"`
	Res    *types.CertMgrRefreshCACertificatesAndCRLs_TaskResponse `xml:"CertMgrRefreshCACertificatesAndCRLs_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CertMgrRefreshCACertificatesAndCRLs_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CertMgrRefreshCACertificatesAndCRLs_Task(ctx context.Context, r soap.RoundTripper, req *types.CertMgrRefreshCACertificatesAndCRLs_Task) (*types.CertMgrRefreshCACertificatesAndCRLs_TaskResponse, error) {
	var reqBody, resBody CertMgrRefreshCACertificatesAndCRLs_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CertMgrRefreshCertificates_TaskBody struct {
	Req    *types.CertMgrRefreshCertificates_Task         `xml:"urn:vim25 CertMgrRefreshCertificates_Task,omitempty"`
	Res    *types.CertMgrRefreshCertificates_TaskResponse `xml:"CertMgrRefreshCertificates_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CertMgrRefreshCertificates_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CertMgrRefreshCertificates_Task(ctx context.Context, r soap.RoundTripper, req *types.CertMgrRefreshCertificates_Task) (*types.CertMgrRefreshCertificates_TaskResponse, error) {
	var reqBody, resBody CertMgrRefreshCertificates_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CertMgrRevokeCertificates_TaskBody struct {
	Req    *types.CertMgrRevokeCertificates_Task         `xml:"urn:vim25 CertMgrRevokeCertificates_Task,omitempty"`
	Res    *types.CertMgrRevokeCertificates_TaskResponse `xml:"CertMgrRevokeCertificates_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CertMgrRevokeCertificates_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CertMgrRevokeCertificates_Task(ctx context.Context, r soap.RoundTripper, req *types.CertMgrRevokeCertificates_Task) (*types.CertMgrRevokeCertificates_TaskResponse, error) {
	var reqBody, resBody CertMgrRevokeCertificates_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeAccessModeBody struct {
	Req    *types.ChangeAccessMode         `xml:"urn:vim25 ChangeAccessMode,omitempty"`
	Res    *types.ChangeAccessModeResponse `xml:"ChangeAccessModeResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeAccessModeBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeAccessMode(ctx context.Context, r soap.RoundTripper, req *types.ChangeAccessMode) (*types.ChangeAccessModeResponse, error) {
	var reqBody, resBody ChangeAccessModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeFileAttributesInGuestBody struct {
	Req    *types.ChangeFileAttributesInGuest         `xml:"urn:vim25 ChangeFileAttributesInGuest,omitempty"`
	Res    *types.ChangeFileAttributesInGuestResponse `xml:"ChangeFileAttributesInGuestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeFileAttributesInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeFileAttributesInGuest(ctx context.Context, r soap.RoundTripper, req *types.ChangeFileAttributesInGuest) (*types.ChangeFileAttributesInGuestResponse, error) {
	var reqBody, resBody ChangeFileAttributesInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeKey_TaskBody struct {
	Req    *types.ChangeKey_Task         `xml:"urn:vim25 ChangeKey_Task,omitempty"`
	Res    *types.ChangeKey_TaskResponse `xml:"ChangeKey_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeKey_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeKey_Task(ctx context.Context, r soap.RoundTripper, req *types.ChangeKey_Task) (*types.ChangeKey_TaskResponse, error) {
	var reqBody, resBody ChangeKey_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeLockdownModeBody struct {
	Req    *types.ChangeLockdownMode         `xml:"urn:vim25 ChangeLockdownMode,omitempty"`
	Res    *types.ChangeLockdownModeResponse `xml:"ChangeLockdownModeResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeLockdownModeBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeLockdownMode(ctx context.Context, r soap.RoundTripper, req *types.ChangeLockdownMode) (*types.ChangeLockdownModeResponse, error) {
	var reqBody, resBody ChangeLockdownModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeNFSUserPasswordBody struct {
	Req    *types.ChangeNFSUserPassword         `xml:"urn:vim25 ChangeNFSUserPassword,omitempty"`
	Res    *types.ChangeNFSUserPasswordResponse `xml:"ChangeNFSUserPasswordResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeNFSUserPasswordBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeNFSUserPassword(ctx context.Context, r soap.RoundTripper, req *types.ChangeNFSUserPassword) (*types.ChangeNFSUserPasswordResponse, error) {
	var reqBody, resBody ChangeNFSUserPasswordBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangeOwnerBody struct {
	Req    *types.ChangeOwner         `xml:"urn:vim25 ChangeOwner,omitempty"`
	Res    *types.ChangeOwnerResponse `xml:"ChangeOwnerResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeOwnerBody) Fault() *soap.Fault { return b.Fault_ }

func ChangeOwner(ctx context.Context, r soap.RoundTripper, req *types.ChangeOwner) (*types.ChangeOwnerResponse, error) {
	var reqBody, resBody ChangeOwnerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ChangePasswordBody struct {
	Req    *types.ChangePassword         `xml:"urn:vim25 ChangePassword,omitempty"`
	Res    *types.ChangePasswordResponse `xml:"ChangePasswordResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangePasswordBody) Fault() *soap.Fault { return b.Fault_ }

func ChangePassword(ctx context.Context, r soap.RoundTripper, req *types.ChangePassword) (*types.ChangePasswordResponse, error) {
	var reqBody, resBody ChangePasswordBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckAddHostEvc_TaskBody struct {
	Req    *types.CheckAddHostEvc_Task         `xml:"urn:vim25 CheckAddHostEvc_Task,omitempty"`
	Res    *types.CheckAddHostEvc_TaskResponse `xml:"CheckAddHostEvc_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckAddHostEvc_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckAddHostEvc_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckAddHostEvc_Task) (*types.CheckAddHostEvc_TaskResponse, error) {
	var reqBody, resBody CheckAddHostEvc_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckAnswerFileStatus_TaskBody struct {
	Req    *types.CheckAnswerFileStatus_Task         `xml:"urn:vim25 CheckAnswerFileStatus_Task,omitempty"`
	Res    *types.CheckAnswerFileStatus_TaskResponse `xml:"CheckAnswerFileStatus_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckAnswerFileStatus_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckAnswerFileStatus_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckAnswerFileStatus_Task) (*types.CheckAnswerFileStatus_TaskResponse, error) {
	var reqBody, resBody CheckAnswerFileStatus_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckClone_TaskBody struct {
	Req    *types.CheckClone_Task         `xml:"urn:vim25 CheckClone_Task,omitempty"`
	Res    *types.CheckClone_TaskResponse `xml:"CheckClone_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckClone_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckClone_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckClone_Task) (*types.CheckClone_TaskResponse, error) {
	var reqBody, resBody CheckClone_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckCompatibility_TaskBody struct {
	Req    *types.CheckCompatibility_Task         `xml:"urn:vim25 CheckCompatibility_Task,omitempty"`
	Res    *types.CheckCompatibility_TaskResponse `xml:"CheckCompatibility_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCompatibility_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckCompatibility_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckCompatibility_Task) (*types.CheckCompatibility_TaskResponse, error) {
	var reqBody, resBody CheckCompatibility_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckCompliance_TaskBody struct {
	Req    *types.CheckCompliance_Task         `xml:"urn:vim25 CheckCompliance_Task,omitempty"`
	Res    *types.CheckCompliance_TaskResponse `xml:"CheckCompliance_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCompliance_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckCompliance_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckCompliance_Task) (*types.CheckCompliance_TaskResponse, error) {
	var reqBody, resBody CheckCompliance_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckConfigureEvcMode_TaskBody struct {
	Req    *types.CheckConfigureEvcMode_Task         `xml:"urn:vim25 CheckConfigureEvcMode_Task,omitempty"`
	Res    *types.CheckConfigureEvcMode_TaskResponse `xml:"CheckConfigureEvcMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckConfigureEvcMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckConfigureEvcMode_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckConfigureEvcMode_Task) (*types.CheckConfigureEvcMode_TaskResponse, error) {
	var reqBody, resBody CheckConfigureEvcMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckCustomizationResourcesBody struct {
	Req    *types.CheckCustomizationResources         `xml:"urn:vim25 CheckCustomizationResources,omitempty"`
	Res    *types.CheckCustomizationResourcesResponse `xml:"CheckCustomizationResourcesResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCustomizationResourcesBody) Fault() *soap.Fault { return b.Fault_ }

func CheckCustomizationResources(ctx context.Context, r soap.RoundTripper, req *types.CheckCustomizationResources) (*types.CheckCustomizationResourcesResponse, error) {
	var reqBody, resBody CheckCustomizationResourcesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckCustomizationSpecBody struct {
	Req    *types.CheckCustomizationSpec         `xml:"urn:vim25 CheckCustomizationSpec,omitempty"`
	Res    *types.CheckCustomizationSpecResponse `xml:"CheckCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func CheckCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.CheckCustomizationSpec) (*types.CheckCustomizationSpecResponse, error) {
	var reqBody, resBody CheckCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckForUpdatesBody struct {
	Req    *types.CheckForUpdates         `xml:"urn:vim25 CheckForUpdates,omitempty"`
	Res    *types.CheckForUpdatesResponse `xml:"CheckForUpdatesResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckForUpdatesBody) Fault() *soap.Fault { return b.Fault_ }

func CheckForUpdates(ctx context.Context, r soap.RoundTripper, req *types.CheckForUpdates) (*types.CheckForUpdatesResponse, error) {
	var reqBody, resBody CheckForUpdatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckHostPatch_TaskBody struct {
	Req    *types.CheckHostPatch_Task         `xml:"urn:vim25 CheckHostPatch_Task,omitempty"`
	Res    *types.CheckHostPatch_TaskResponse `xml:"CheckHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckHostPatch_Task) (*types.CheckHostPatch_TaskResponse, error) {
	var reqBody, resBody CheckHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckInstantClone_TaskBody struct {
	Req    *types.CheckInstantClone_Task         `xml:"urn:vim25 CheckInstantClone_Task,omitempty"`
	Res    *types.CheckInstantClone_TaskResponse `xml:"CheckInstantClone_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckInstantClone_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckInstantClone_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckInstantClone_Task) (*types.CheckInstantClone_TaskResponse, error) {
	var reqBody, resBody CheckInstantClone_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckLicenseFeatureBody struct {
	Req    *types.CheckLicenseFeature         `xml:"urn:vim25 CheckLicenseFeature,omitempty"`
	Res    *types.CheckLicenseFeatureResponse `xml:"CheckLicenseFeatureResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckLicenseFeatureBody) Fault() *soap.Fault { return b.Fault_ }

func CheckLicenseFeature(ctx context.Context, r soap.RoundTripper, req *types.CheckLicenseFeature) (*types.CheckLicenseFeatureResponse, error) {
	var reqBody, resBody CheckLicenseFeatureBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckMigrate_TaskBody struct {
	Req    *types.CheckMigrate_Task         `xml:"urn:vim25 CheckMigrate_Task,omitempty"`
	Res    *types.CheckMigrate_TaskResponse `xml:"CheckMigrate_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckMigrate_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckMigrate_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckMigrate_Task) (*types.CheckMigrate_TaskResponse, error) {
	var reqBody, resBody CheckMigrate_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckPowerOn_TaskBody struct {
	Req    *types.CheckPowerOn_Task         `xml:"urn:vim25 CheckPowerOn_Task,omitempty"`
	Res    *types.CheckPowerOn_TaskResponse `xml:"CheckPowerOn_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckPowerOn_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckPowerOn_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckPowerOn_Task) (*types.CheckPowerOn_TaskResponse, error) {
	var reqBody, resBody CheckPowerOn_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckProfileCompliance_TaskBody struct {
	Req    *types.CheckProfileCompliance_Task         `xml:"urn:vim25 CheckProfileCompliance_Task,omitempty"`
	Res    *types.CheckProfileCompliance_TaskResponse `xml:"CheckProfileCompliance_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckProfileCompliance_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckProfileCompliance_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckProfileCompliance_Task) (*types.CheckProfileCompliance_TaskResponse, error) {
	var reqBody, resBody CheckProfileCompliance_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckRelocate_TaskBody struct {
	Req    *types.CheckRelocate_Task         `xml:"urn:vim25 CheckRelocate_Task,omitempty"`
	Res    *types.CheckRelocate_TaskResponse `xml:"CheckRelocate_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckRelocate_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckRelocate_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckRelocate_Task) (*types.CheckRelocate_TaskResponse, error) {
	var reqBody, resBody CheckRelocate_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CheckVmConfig_TaskBody struct {
	Req    *types.CheckVmConfig_Task         `xml:"urn:vim25 CheckVmConfig_Task,omitempty"`
	Res    *types.CheckVmConfig_TaskResponse `xml:"CheckVmConfig_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckVmConfig_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CheckVmConfig_Task(ctx context.Context, r soap.RoundTripper, req *types.CheckVmConfig_Task) (*types.CheckVmConfig_TaskResponse, error) {
	var reqBody, resBody CheckVmConfig_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearComplianceStatusBody struct {
	Req    *types.ClearComplianceStatus         `xml:"urn:vim25 ClearComplianceStatus,omitempty"`
	Res    *types.ClearComplianceStatusResponse `xml:"ClearComplianceStatusResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearComplianceStatusBody) Fault() *soap.Fault { return b.Fault_ }

func ClearComplianceStatus(ctx context.Context, r soap.RoundTripper, req *types.ClearComplianceStatus) (*types.ClearComplianceStatusResponse, error) {
	var reqBody, resBody ClearComplianceStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearNFSUserBody struct {
	Req    *types.ClearNFSUser         `xml:"urn:vim25 ClearNFSUser,omitempty"`
	Res    *types.ClearNFSUserResponse `xml:"ClearNFSUserResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearNFSUserBody) Fault() *soap.Fault { return b.Fault_ }

func ClearNFSUser(ctx context.Context, r soap.RoundTripper, req *types.ClearNFSUser) (*types.ClearNFSUserResponse, error) {
	var reqBody, resBody ClearNFSUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearSystemEventLogBody struct {
	Req    *types.ClearSystemEventLog         `xml:"urn:vim25 ClearSystemEventLog,omitempty"`
	Res    *types.ClearSystemEventLogResponse `xml:"ClearSystemEventLogResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearSystemEventLogBody) Fault() *soap.Fault { return b.Fault_ }

func ClearSystemEventLog(ctx context.Context, r soap.RoundTripper, req *types.ClearSystemEventLog) (*types.ClearSystemEventLogResponse, error) {
	var reqBody, resBody ClearSystemEventLogBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearTriggeredAlarmsBody struct {
	Req    *types.ClearTriggeredAlarms         `xml:"urn:vim25 ClearTriggeredAlarms,omitempty"`
	Res    *types.ClearTriggeredAlarmsResponse `xml:"ClearTriggeredAlarmsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearTriggeredAlarmsBody) Fault() *soap.Fault { return b.Fault_ }

func ClearTriggeredAlarms(ctx context.Context, r soap.RoundTripper, req *types.ClearTriggeredAlarms) (*types.ClearTriggeredAlarmsResponse, error) {
	var reqBody, resBody ClearTriggeredAlarmsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClearVStorageObjectControlFlagsBody struct {
	Req    *types.ClearVStorageObjectControlFlags         `xml:"urn:vim25 ClearVStorageObjectControlFlags,omitempty"`
	Res    *types.ClearVStorageObjectControlFlagsResponse `xml:"ClearVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func ClearVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.ClearVStorageObjectControlFlags) (*types.ClearVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody ClearVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CloneSessionBody struct {
	Req    *types.CloneSession         `xml:"urn:vim25 CloneSession,omitempty"`
	Res    *types.CloneSessionResponse `xml:"CloneSessionResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneSessionBody) Fault() *soap.Fault { return b.Fault_ }

func CloneSession(ctx context.Context, r soap.RoundTripper, req *types.CloneSession) (*types.CloneSessionResponse, error) {
	var reqBody, resBody CloneSessionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CloneVApp_TaskBody struct {
	Req    *types.CloneVApp_Task         `xml:"urn:vim25 CloneVApp_Task,omitempty"`
	Res    *types.CloneVApp_TaskResponse `xml:"CloneVApp_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneVApp_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CloneVApp_Task(ctx context.Context, r soap.RoundTripper, req *types.CloneVApp_Task) (*types.CloneVApp_TaskResponse, error) {
	var reqBody, resBody CloneVApp_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CloneVM_TaskBody struct {
	Req    *types.CloneVM_Task         `xml:"urn:vim25 CloneVM_Task,omitempty"`
	Res    *types.CloneVM_TaskResponse `xml:"CloneVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CloneVM_Task(ctx context.Context, r soap.RoundTripper, req *types.CloneVM_Task) (*types.CloneVM_TaskResponse, error) {
	var reqBody, resBody CloneVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CloneVStorageObject_TaskBody struct {
	Req    *types.CloneVStorageObject_Task         `xml:"urn:vim25 CloneVStorageObject_Task,omitempty"`
	Res    *types.CloneVStorageObject_TaskResponse `xml:"CloneVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CloneVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.CloneVStorageObject_Task) (*types.CloneVStorageObject_TaskResponse, error) {
	var reqBody, resBody CloneVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CloseInventoryViewFolderBody struct {
	Req    *types.CloseInventoryViewFolder         `xml:"urn:vim25 CloseInventoryViewFolder,omitempty"`
	Res    *types.CloseInventoryViewFolderResponse `xml:"CloseInventoryViewFolderResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloseInventoryViewFolderBody) Fault() *soap.Fault { return b.Fault_ }

func CloseInventoryViewFolder(ctx context.Context, r soap.RoundTripper, req *types.CloseInventoryViewFolder) (*types.CloseInventoryViewFolderResponse, error) {
	var reqBody, resBody CloseInventoryViewFolderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ClusterEnterMaintenanceModeBody struct {
	Req    *types.ClusterEnterMaintenanceMode         `xml:"urn:vim25 ClusterEnterMaintenanceMode,omitempty"`
	Res    *types.ClusterEnterMaintenanceModeResponse `xml:"ClusterEnterMaintenanceModeResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClusterEnterMaintenanceModeBody) Fault() *soap.Fault { return b.Fault_ }

func ClusterEnterMaintenanceMode(ctx context.Context, r soap.RoundTripper, req *types.ClusterEnterMaintenanceMode) (*types.ClusterEnterMaintenanceModeResponse, error) {
	var reqBody, resBody ClusterEnterMaintenanceModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CompositeHostProfile_TaskBody struct {
	Req    *types.CompositeHostProfile_Task         `xml:"urn:vim25 CompositeHostProfile_Task,omitempty"`
	Res    *types.CompositeHostProfile_TaskResponse `xml:"CompositeHostProfile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CompositeHostProfile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CompositeHostProfile_Task(ctx context.Context, r soap.RoundTripper, req *types.CompositeHostProfile_Task) (*types.CompositeHostProfile_TaskResponse, error) {
	var reqBody, resBody CompositeHostProfile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ComputeDiskPartitionInfoBody struct {
	Req    *types.ComputeDiskPartitionInfo         `xml:"urn:vim25 ComputeDiskPartitionInfo,omitempty"`
	Res    *types.ComputeDiskPartitionInfoResponse `xml:"ComputeDiskPartitionInfoResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ComputeDiskPartitionInfoBody) Fault() *soap.Fault { return b.Fault_ }

func ComputeDiskPartitionInfo(ctx context.Context, r soap.RoundTripper, req *types.ComputeDiskPartitionInfo) (*types.ComputeDiskPartitionInfoResponse, error) {
	var reqBody, resBody ComputeDiskPartitionInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ComputeDiskPartitionInfoForResizeBody struct {
	Req    *types.ComputeDiskPartitionInfoForResize         `xml:"urn:vim25 ComputeDiskPartitionInfoForResize,omitempty"`
	Res    *types.ComputeDiskPartitionInfoForResizeResponse `xml:"ComputeDiskPartitionInfoForResizeResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ComputeDiskPartitionInfoForResizeBody) Fault() *soap.Fault { return b.Fault_ }

func ComputeDiskPartitionInfoForResize(ctx context.Context, r soap.RoundTripper, req *types.ComputeDiskPartitionInfoForResize) (*types.ComputeDiskPartitionInfoForResizeResponse, error) {
	var reqBody, resBody ComputeDiskPartitionInfoForResizeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureCryptoKeyBody struct {
	Req    *types.ConfigureCryptoKey         `xml:"urn:vim25 ConfigureCryptoKey,omitempty"`
	Res    *types.ConfigureCryptoKeyResponse `xml:"ConfigureCryptoKeyResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureCryptoKeyBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureCryptoKey(ctx context.Context, r soap.RoundTripper, req *types.ConfigureCryptoKey) (*types.ConfigureCryptoKeyResponse, error) {
	var reqBody, resBody ConfigureCryptoKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureDatastoreIORM_TaskBody struct {
	Req    *types.ConfigureDatastoreIORM_Task         `xml:"urn:vim25 ConfigureDatastoreIORM_Task,omitempty"`
	Res    *types.ConfigureDatastoreIORM_TaskResponse `xml:"ConfigureDatastoreIORM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureDatastoreIORM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureDatastoreIORM_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureDatastoreIORM_Task) (*types.ConfigureDatastoreIORM_TaskResponse, error) {
	var reqBody, resBody ConfigureDatastoreIORM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureDatastorePrincipalBody struct {
	Req    *types.ConfigureDatastorePrincipal         `xml:"urn:vim25 ConfigureDatastorePrincipal,omitempty"`
	Res    *types.ConfigureDatastorePrincipalResponse `xml:"ConfigureDatastorePrincipalResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureDatastorePrincipalBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureDatastorePrincipal(ctx context.Context, r soap.RoundTripper, req *types.ConfigureDatastorePrincipal) (*types.ConfigureDatastorePrincipalResponse, error) {
	var reqBody, resBody ConfigureDatastorePrincipalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureEvcMode_TaskBody struct {
	Req    *types.ConfigureEvcMode_Task         `xml:"urn:vim25 ConfigureEvcMode_Task,omitempty"`
	Res    *types.ConfigureEvcMode_TaskResponse `xml:"ConfigureEvcMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureEvcMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureEvcMode_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureEvcMode_Task) (*types.ConfigureEvcMode_TaskResponse, error) {
	var reqBody, resBody ConfigureEvcMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureHCI_TaskBody struct {
	Req    *types.ConfigureHCI_Task         `xml:"urn:vim25 ConfigureHCI_Task,omitempty"`
	Res    *types.ConfigureHCI_TaskResponse `xml:"ConfigureHCI_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureHCI_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureHCI_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureHCI_Task) (*types.ConfigureHCI_TaskResponse, error) {
	var reqBody, resBody ConfigureHCI_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureHostCache_TaskBody struct {
	Req    *types.ConfigureHostCache_Task         `xml:"urn:vim25 ConfigureHostCache_Task,omitempty"`
	Res    *types.ConfigureHostCache_TaskResponse `xml:"ConfigureHostCache_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureHostCache_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureHostCache_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureHostCache_Task) (*types.ConfigureHostCache_TaskResponse, error) {
	var reqBody, resBody ConfigureHostCache_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureLicenseSourceBody struct {
	Req    *types.ConfigureLicenseSource         `xml:"urn:vim25 ConfigureLicenseSource,omitempty"`
	Res    *types.ConfigureLicenseSourceResponse `xml:"ConfigureLicenseSourceResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureLicenseSourceBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureLicenseSource(ctx context.Context, r soap.RoundTripper, req *types.ConfigureLicenseSource) (*types.ConfigureLicenseSourceResponse, error) {
	var reqBody, resBody ConfigureLicenseSourceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigurePowerPolicyBody struct {
	Req    *types.ConfigurePowerPolicy         `xml:"urn:vim25 ConfigurePowerPolicy,omitempty"`
	Res    *types.ConfigurePowerPolicyResponse `xml:"ConfigurePowerPolicyResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigurePowerPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigurePowerPolicy(ctx context.Context, r soap.RoundTripper, req *types.ConfigurePowerPolicy) (*types.ConfigurePowerPolicyResponse, error) {
	var reqBody, resBody ConfigurePowerPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureStorageDrsForPod_TaskBody struct {
	Req    *types.ConfigureStorageDrsForPod_Task         `xml:"urn:vim25 ConfigureStorageDrsForPod_Task,omitempty"`
	Res    *types.ConfigureStorageDrsForPod_TaskResponse `xml:"ConfigureStorageDrsForPod_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureStorageDrsForPod_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureStorageDrsForPod_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureStorageDrsForPod_Task) (*types.ConfigureStorageDrsForPod_TaskResponse, error) {
	var reqBody, resBody ConfigureStorageDrsForPod_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureVFlashResourceEx_TaskBody struct {
	Req    *types.ConfigureVFlashResourceEx_Task         `xml:"urn:vim25 ConfigureVFlashResourceEx_Task,omitempty"`
	Res    *types.ConfigureVFlashResourceEx_TaskResponse `xml:"ConfigureVFlashResourceEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureVFlashResourceEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureVFlashResourceEx_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureVFlashResourceEx_Task) (*types.ConfigureVFlashResourceEx_TaskResponse, error) {
	var reqBody, resBody ConfigureVFlashResourceEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConnectNvmeControllerBody struct {
	Req    *types.ConnectNvmeController         `xml:"urn:vim25 ConnectNvmeController,omitempty"`
	Res    *types.ConnectNvmeControllerResponse `xml:"ConnectNvmeControllerResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConnectNvmeControllerBody) Fault() *soap.Fault { return b.Fault_ }

func ConnectNvmeController(ctx context.Context, r soap.RoundTripper, req *types.ConnectNvmeController) (*types.ConnectNvmeControllerResponse, error) {
	var reqBody, resBody ConnectNvmeControllerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConnectNvmeControllerEx_TaskBody struct {
	Req    *types.ConnectNvmeControllerEx_Task         `xml:"urn:vim25 ConnectNvmeControllerEx_Task,omitempty"`
	Res    *types.ConnectNvmeControllerEx_TaskResponse `xml:"ConnectNvmeControllerEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConnectNvmeControllerEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConnectNvmeControllerEx_Task(ctx context.Context, r soap.RoundTripper, req *types.ConnectNvmeControllerEx_Task) (*types.ConnectNvmeControllerEx_TaskResponse, error) {
	var reqBody, resBody ConnectNvmeControllerEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConsolidateVMDisks_TaskBody struct {
	Req    *types.ConsolidateVMDisks_Task         `xml:"urn:vim25 ConsolidateVMDisks_Task,omitempty"`
	Res    *types.ConsolidateVMDisks_TaskResponse `xml:"ConsolidateVMDisks_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConsolidateVMDisks_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConsolidateVMDisks_Task(ctx context.Context, r soap.RoundTripper, req *types.ConsolidateVMDisks_Task) (*types.ConsolidateVMDisks_TaskResponse, error) {
	var reqBody, resBody ConsolidateVMDisks_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ContinueRetrievePropertiesExBody struct {
	Req    *types.ContinueRetrievePropertiesEx         `xml:"urn:vim25 ContinueRetrievePropertiesEx,omitempty"`
	Res    *types.ContinueRetrievePropertiesExResponse `xml:"ContinueRetrievePropertiesExResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ContinueRetrievePropertiesExBody) Fault() *soap.Fault { return b.Fault_ }

func ContinueRetrievePropertiesEx(ctx context.Context, r soap.RoundTripper, req *types.ContinueRetrievePropertiesEx) (*types.ContinueRetrievePropertiesExResponse, error) {
	var reqBody, resBody ContinueRetrievePropertiesExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConvertNamespacePathToUuidPathBody struct {
	Req    *types.ConvertNamespacePathToUuidPath         `xml:"urn:vim25 ConvertNamespacePathToUuidPath,omitempty"`
	Res    *types.ConvertNamespacePathToUuidPathResponse `xml:"ConvertNamespacePathToUuidPathResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConvertNamespacePathToUuidPathBody) Fault() *soap.Fault { return b.Fault_ }

func ConvertNamespacePathToUuidPath(ctx context.Context, r soap.RoundTripper, req *types.ConvertNamespacePathToUuidPath) (*types.ConvertNamespacePathToUuidPathResponse, error) {
	var reqBody, resBody ConvertNamespacePathToUuidPathBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CopyDatastoreFile_TaskBody struct {
	Req    *types.CopyDatastoreFile_Task         `xml:"urn:vim25 CopyDatastoreFile_Task,omitempty"`
	Res    *types.CopyDatastoreFile_TaskResponse `xml:"CopyDatastoreFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CopyDatastoreFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CopyDatastoreFile_Task(ctx context.Context, r soap.RoundTripper, req *types.CopyDatastoreFile_Task) (*types.CopyDatastoreFile_TaskResponse, error) {
	var reqBody, resBody CopyDatastoreFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CopyVirtualDisk_TaskBody struct {
	Req    *types.CopyVirtualDisk_Task         `xml:"urn:vim25 CopyVirtualDisk_Task,omitempty"`
	Res    *types.CopyVirtualDisk_TaskResponse `xml:"CopyVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CopyVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CopyVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.CopyVirtualDisk_Task) (*types.CopyVirtualDisk_TaskResponse, error) {
	var reqBody, resBody CopyVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateAlarmBody struct {
	Req    *types.CreateAlarm         `xml:"urn:vim25 CreateAlarm,omitempty"`
	Res    *types.CreateAlarmResponse `xml:"CreateAlarmResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func CreateAlarm(ctx context.Context, r soap.RoundTripper, req *types.CreateAlarm) (*types.CreateAlarmResponse, error) {
	var reqBody, resBody CreateAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateChildVM_TaskBody struct {
	Req    *types.CreateChildVM_Task         `xml:"urn:vim25 CreateChildVM_Task,omitempty"`
	Res    *types.CreateChildVM_TaskResponse `xml:"CreateChildVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateChildVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateChildVM_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateChildVM_Task) (*types.CreateChildVM_TaskResponse, error) {
	var reqBody, resBody CreateChildVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateClusterBody struct {
	Req    *types.CreateCluster         `xml:"urn:vim25 CreateCluster,omitempty"`
	Res    *types.CreateClusterResponse `xml:"CreateClusterResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateClusterBody) Fault() *soap.Fault { return b.Fault_ }

func CreateCluster(ctx context.Context, r soap.RoundTripper, req *types.CreateCluster) (*types.CreateClusterResponse, error) {
	var reqBody, resBody CreateClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateClusterExBody struct {
	Req    *types.CreateClusterEx         `xml:"urn:vim25 CreateClusterEx,omitempty"`
	Res    *types.CreateClusterExResponse `xml:"CreateClusterExResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateClusterExBody) Fault() *soap.Fault { return b.Fault_ }

func CreateClusterEx(ctx context.Context, r soap.RoundTripper, req *types.CreateClusterEx) (*types.CreateClusterExResponse, error) {
	var reqBody, resBody CreateClusterExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateCollectorForEventsBody struct {
	Req    *types.CreateCollectorForEvents         `xml:"urn:vim25 CreateCollectorForEvents,omitempty"`
	Res    *types.CreateCollectorForEventsResponse `xml:"CreateCollectorForEventsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCollectorForEventsBody) Fault() *soap.Fault { return b.Fault_ }

func CreateCollectorForEvents(ctx context.Context, r soap.RoundTripper, req *types.CreateCollectorForEvents) (*types.CreateCollectorForEventsResponse, error) {
	var reqBody, resBody CreateCollectorForEventsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateCollectorForTasksBody struct {
	Req    *types.CreateCollectorForTasks         `xml:"urn:vim25 CreateCollectorForTasks,omitempty"`
	Res    *types.CreateCollectorForTasksResponse `xml:"CreateCollectorForTasksResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCollectorForTasksBody) Fault() *soap.Fault { return b.Fault_ }

func CreateCollectorForTasks(ctx context.Context, r soap.RoundTripper, req *types.CreateCollectorForTasks) (*types.CreateCollectorForTasksResponse, error) {
	var reqBody, resBody CreateCollectorForTasksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateCollectorWithInfoFilterForTasksBody struct {
	Req    *types.CreateCollectorWithInfoFilterForTasks         `xml:"urn:vim25 CreateCollectorWithInfoFilterForTasks,omitempty"`
	Res    *types.CreateCollectorWithInfoFilterForTasksResponse `xml:"CreateCollectorWithInfoFilterForTasksResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCollectorWithInfoFilterForTasksBody) Fault() *soap.Fault { return b.Fault_ }

func CreateCollectorWithInfoFilterForTasks(ctx context.Context, r soap.RoundTripper, req *types.CreateCollectorWithInfoFilterForTasks) (*types.CreateCollectorWithInfoFilterForTasksResponse, error) {
	var reqBody, resBody CreateCollectorWithInfoFilterForTasksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateContainerViewBody struct {
	Req    *types.CreateContainerView         `xml:"urn:vim25 CreateContainerView,omitempty"`
	Res    *types.CreateContainerViewResponse `xml:"CreateContainerViewResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateContainerViewBody) Fault() *soap.Fault { return b.Fault_ }

func CreateContainerView(ctx context.Context, r soap.RoundTripper, req *types.CreateContainerView) (*types.CreateContainerViewResponse, error) {
	var reqBody, resBody CreateContainerViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateCustomizationSpecBody struct {
	Req    *types.CreateCustomizationSpec         `xml:"urn:vim25 CreateCustomizationSpec,omitempty"`
	Res    *types.CreateCustomizationSpecResponse `xml:"CreateCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func CreateCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.CreateCustomizationSpec) (*types.CreateCustomizationSpecResponse, error) {
	var reqBody, resBody CreateCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDVPortgroup_TaskBody struct {
	Req    *types.CreateDVPortgroup_Task         `xml:"urn:vim25 CreateDVPortgroup_Task,omitempty"`
	Res    *types.CreateDVPortgroup_TaskResponse `xml:"CreateDVPortgroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDVPortgroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDVPortgroup_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateDVPortgroup_Task) (*types.CreateDVPortgroup_TaskResponse, error) {
	var reqBody, resBody CreateDVPortgroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDVS_TaskBody struct {
	Req    *types.CreateDVS_Task         `xml:"urn:vim25 CreateDVS_Task,omitempty"`
	Res    *types.CreateDVS_TaskResponse `xml:"CreateDVS_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDVS_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDVS_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateDVS_Task) (*types.CreateDVS_TaskResponse, error) {
	var reqBody, resBody CreateDVS_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDatacenterBody struct {
	Req    *types.CreateDatacenter         `xml:"urn:vim25 CreateDatacenter,omitempty"`
	Res    *types.CreateDatacenterResponse `xml:"CreateDatacenterResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDatacenterBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDatacenter(ctx context.Context, r soap.RoundTripper, req *types.CreateDatacenter) (*types.CreateDatacenterResponse, error) {
	var reqBody, resBody CreateDatacenterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDefaultProfileBody struct {
	Req    *types.CreateDefaultProfile         `xml:"urn:vim25 CreateDefaultProfile,omitempty"`
	Res    *types.CreateDefaultProfileResponse `xml:"CreateDefaultProfileResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDefaultProfileBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDefaultProfile(ctx context.Context, r soap.RoundTripper, req *types.CreateDefaultProfile) (*types.CreateDefaultProfileResponse, error) {
	var reqBody, resBody CreateDefaultProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDescriptorBody struct {
	Req    *types.CreateDescriptor         `xml:"urn:vim25 CreateDescriptor,omitempty"`
	Res    *types.CreateDescriptorResponse `xml:"CreateDescriptorResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDescriptorBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDescriptor(ctx context.Context, r soap.RoundTripper, req *types.CreateDescriptor) (*types.CreateDescriptorResponse, error) {
	var reqBody, resBody CreateDescriptorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDiagnosticPartitionBody struct {
	Req    *types.CreateDiagnosticPartition         `xml:"urn:vim25 CreateDiagnosticPartition,omitempty"`
	Res    *types.CreateDiagnosticPartitionResponse `xml:"CreateDiagnosticPartitionResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDiagnosticPartitionBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDiagnosticPartition(ctx context.Context, r soap.RoundTripper, req *types.CreateDiagnosticPartition) (*types.CreateDiagnosticPartitionResponse, error) {
	var reqBody, resBody CreateDiagnosticPartitionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDirectoryBody struct {
	Req    *types.CreateDirectory         `xml:"urn:vim25 CreateDirectory,omitempty"`
	Res    *types.CreateDirectoryResponse `xml:"CreateDirectoryResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDirectoryBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDirectory(ctx context.Context, r soap.RoundTripper, req *types.CreateDirectory) (*types.CreateDirectoryResponse, error) {
	var reqBody, resBody CreateDirectoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDiskFromSnapshot_TaskBody struct {
	Req    *types.CreateDiskFromSnapshot_Task         `xml:"urn:vim25 CreateDiskFromSnapshot_Task,omitempty"`
	Res    *types.CreateDiskFromSnapshot_TaskResponse `xml:"CreateDiskFromSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDiskFromSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDiskFromSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateDiskFromSnapshot_Task) (*types.CreateDiskFromSnapshot_TaskResponse, error) {
	var reqBody, resBody CreateDiskFromSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateDisk_TaskBody struct {
	Req    *types.CreateDisk_Task         `xml:"urn:vim25 CreateDisk_Task,omitempty"`
	Res    *types.CreateDisk_TaskResponse `xml:"CreateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateDisk_Task) (*types.CreateDisk_TaskResponse, error) {
	var reqBody, resBody CreateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateFilterBody struct {
	Req    *types.CreateFilter         `xml:"urn:vim25 CreateFilter,omitempty"`
	Res    *types.CreateFilterResponse `xml:"CreateFilterResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateFilterBody) Fault() *soap.Fault { return b.Fault_ }

func CreateFilter(ctx context.Context, r soap.RoundTripper, req *types.CreateFilter) (*types.CreateFilterResponse, error) {
	var reqBody, resBody CreateFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateFolderBody struct {
	Req    *types.CreateFolder         `xml:"urn:vim25 CreateFolder,omitempty"`
	Res    *types.CreateFolderResponse `xml:"CreateFolderResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateFolderBody) Fault() *soap.Fault { return b.Fault_ }

func CreateFolder(ctx context.Context, r soap.RoundTripper, req *types.CreateFolder) (*types.CreateFolderResponse, error) {
	var reqBody, resBody CreateFolderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateGroupBody struct {
	Req    *types.CreateGroup         `xml:"urn:vim25 CreateGroup,omitempty"`
	Res    *types.CreateGroupResponse `xml:"CreateGroupResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateGroupBody) Fault() *soap.Fault { return b.Fault_ }

func CreateGroup(ctx context.Context, r soap.RoundTripper, req *types.CreateGroup) (*types.CreateGroupResponse, error) {
	var reqBody, resBody CreateGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateImportSpecBody struct {
	Req    *types.CreateImportSpec         `xml:"urn:vim25 CreateImportSpec,omitempty"`
	Res    *types.CreateImportSpecResponse `xml:"CreateImportSpecResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateImportSpecBody) Fault() *soap.Fault { return b.Fault_ }

func CreateImportSpec(ctx context.Context, r soap.RoundTripper, req *types.CreateImportSpec) (*types.CreateImportSpecResponse, error) {
	var reqBody, resBody CreateImportSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateInventoryViewBody struct {
	Req    *types.CreateInventoryView         `xml:"urn:vim25 CreateInventoryView,omitempty"`
	Res    *types.CreateInventoryViewResponse `xml:"CreateInventoryViewResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateInventoryViewBody) Fault() *soap.Fault { return b.Fault_ }

func CreateInventoryView(ctx context.Context, r soap.RoundTripper, req *types.CreateInventoryView) (*types.CreateInventoryViewResponse, error) {
	var reqBody, resBody CreateInventoryViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateIpPoolBody struct {
	Req    *types.CreateIpPool         `xml:"urn:vim25 CreateIpPool,omitempty"`
	Res    *types.CreateIpPoolResponse `xml:"CreateIpPoolResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateIpPoolBody) Fault() *soap.Fault { return b.Fault_ }

func CreateIpPool(ctx context.Context, r soap.RoundTripper, req *types.CreateIpPool) (*types.CreateIpPoolResponse, error) {
	var reqBody, resBody CreateIpPoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateListViewBody struct {
	Req    *types.CreateListView         `xml:"urn:vim25 CreateListView,omitempty"`
	Res    *types.CreateListViewResponse `xml:"CreateListViewResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateListViewBody) Fault() *soap.Fault { return b.Fault_ }

func CreateListView(ctx context.Context, r soap.RoundTripper, req *types.CreateListView) (*types.CreateListViewResponse, error) {
	var reqBody, resBody CreateListViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateListViewFromViewBody struct {
	Req    *types.CreateListViewFromView         `xml:"urn:vim25 CreateListViewFromView,omitempty"`
	Res    *types.CreateListViewFromViewResponse `xml:"CreateListViewFromViewResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateListViewFromViewBody) Fault() *soap.Fault { return b.Fault_ }

func CreateListViewFromView(ctx context.Context, r soap.RoundTripper, req *types.CreateListViewFromView) (*types.CreateListViewFromViewResponse, error) {
	var reqBody, resBody CreateListViewFromViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateLocalDatastoreBody struct {
	Req    *types.CreateLocalDatastore         `xml:"urn:vim25 CreateLocalDatastore,omitempty"`
	Res    *types.CreateLocalDatastoreResponse `xml:"CreateLocalDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateLocalDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func CreateLocalDatastore(ctx context.Context, r soap.RoundTripper, req *types.CreateLocalDatastore) (*types.CreateLocalDatastoreResponse, error) {
	var reqBody, resBody CreateLocalDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateNasDatastoreBody struct {
	Req    *types.CreateNasDatastore         `xml:"urn:vim25 CreateNasDatastore,omitempty"`
	Res    *types.CreateNasDatastoreResponse `xml:"CreateNasDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateNasDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func CreateNasDatastore(ctx context.Context, r soap.RoundTripper, req *types.CreateNasDatastore) (*types.CreateNasDatastoreResponse, error) {
	var reqBody, resBody CreateNasDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateNvdimmNamespace_TaskBody struct {
	Req    *types.CreateNvdimmNamespace_Task         `xml:"urn:vim25 CreateNvdimmNamespace_Task,omitempty"`
	Res    *types.CreateNvdimmNamespace_TaskResponse `xml:"CreateNvdimmNamespace_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateNvdimmNamespace_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateNvdimmNamespace_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateNvdimmNamespace_Task) (*types.CreateNvdimmNamespace_TaskResponse, error) {
	var reqBody, resBody CreateNvdimmNamespace_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateNvdimmPMemNamespace_TaskBody struct {
	Req    *types.CreateNvdimmPMemNamespace_Task         `xml:"urn:vim25 CreateNvdimmPMemNamespace_Task,omitempty"`
	Res    *types.CreateNvdimmPMemNamespace_TaskResponse `xml:"CreateNvdimmPMemNamespace_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateNvdimmPMemNamespace_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateNvdimmPMemNamespace_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateNvdimmPMemNamespace_Task) (*types.CreateNvdimmPMemNamespace_TaskResponse, error) {
	var reqBody, resBody CreateNvdimmPMemNamespace_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateNvmeOverRdmaAdapterBody struct {
	Req    *types.CreateNvmeOverRdmaAdapter         `xml:"urn:vim25 CreateNvmeOverRdmaAdapter,omitempty"`
	Res    *types.CreateNvmeOverRdmaAdapterResponse `xml:"CreateNvmeOverRdmaAdapterResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateNvmeOverRdmaAdapterBody) Fault() *soap.Fault { return b.Fault_ }

func CreateNvmeOverRdmaAdapter(ctx context.Context, r soap.RoundTripper, req *types.CreateNvmeOverRdmaAdapter) (*types.CreateNvmeOverRdmaAdapterResponse, error) {
	var reqBody, resBody CreateNvmeOverRdmaAdapterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateObjectScheduledTaskBody struct {
	Req    *types.CreateObjectScheduledTask         `xml:"urn:vim25 CreateObjectScheduledTask,omitempty"`
	Res    *types.CreateObjectScheduledTaskResponse `xml:"CreateObjectScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateObjectScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateObjectScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.CreateObjectScheduledTask) (*types.CreateObjectScheduledTaskResponse, error) {
	var reqBody, resBody CreateObjectScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreatePerfIntervalBody struct {
	Req    *types.CreatePerfInterval         `xml:"urn:vim25 CreatePerfInterval,omitempty"`
	Res    *types.CreatePerfIntervalResponse `xml:"CreatePerfIntervalResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreatePerfIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func CreatePerfInterval(ctx context.Context, r soap.RoundTripper, req *types.CreatePerfInterval) (*types.CreatePerfIntervalResponse, error) {
	var reqBody, resBody CreatePerfIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateProfileBody struct {
	Req    *types.CreateProfile         `xml:"urn:vim25 CreateProfile,omitempty"`
	Res    *types.CreateProfileResponse `xml:"CreateProfileResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateProfileBody) Fault() *soap.Fault { return b.Fault_ }

func CreateProfile(ctx context.Context, r soap.RoundTripper, req *types.CreateProfile) (*types.CreateProfileResponse, error) {
	var reqBody, resBody CreateProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreatePropertyCollectorBody struct {
	Req    *types.CreatePropertyCollector         `xml:"urn:vim25 CreatePropertyCollector,omitempty"`
	Res    *types.CreatePropertyCollectorResponse `xml:"CreatePropertyCollectorResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreatePropertyCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func CreatePropertyCollector(ctx context.Context, r soap.RoundTripper, req *types.CreatePropertyCollector) (*types.CreatePropertyCollectorResponse, error) {
	var reqBody, resBody CreatePropertyCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateRegistryKeyInGuestBody struct {
	Req    *types.CreateRegistryKeyInGuest         `xml:"urn:vim25 CreateRegistryKeyInGuest,omitempty"`
	Res    *types.CreateRegistryKeyInGuestResponse `xml:"CreateRegistryKeyInGuestResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateRegistryKeyInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func CreateRegistryKeyInGuest(ctx context.Context, r soap.RoundTripper, req *types.CreateRegistryKeyInGuest) (*types.CreateRegistryKeyInGuestResponse, error) {
	var reqBody, resBody CreateRegistryKeyInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateResourcePoolBody struct {
	Req    *types.CreateResourcePool         `xml:"urn:vim25 CreateResourcePool,omitempty"`
	Res    *types.CreateResourcePoolResponse `xml:"CreateResourcePoolResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateResourcePoolBody) Fault() *soap.Fault { return b.Fault_ }

func CreateResourcePool(ctx context.Context, r soap.RoundTripper, req *types.CreateResourcePool) (*types.CreateResourcePoolResponse, error) {
	var reqBody, resBody CreateResourcePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateScheduledTaskBody struct {
	Req    *types.CreateScheduledTask         `xml:"urn:vim25 CreateScheduledTask,omitempty"`
	Res    *types.CreateScheduledTaskResponse `xml:"CreateScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.CreateScheduledTask) (*types.CreateScheduledTaskResponse, error) {
	var reqBody, resBody CreateScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateScreenshot_TaskBody struct {
	Req    *types.CreateScreenshot_Task         `xml:"urn:vim25 CreateScreenshot_Task,omitempty"`
	Res    *types.CreateScreenshot_TaskResponse `xml:"CreateScreenshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateScreenshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateScreenshot_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateScreenshot_Task) (*types.CreateScreenshot_TaskResponse, error) {
	var reqBody, resBody CreateScreenshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateSecondaryVMEx_TaskBody struct {
	Req    *types.CreateSecondaryVMEx_Task         `xml:"urn:vim25 CreateSecondaryVMEx_Task,omitempty"`
	Res    *types.CreateSecondaryVMEx_TaskResponse `xml:"CreateSecondaryVMEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSecondaryVMEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateSecondaryVMEx_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateSecondaryVMEx_Task) (*types.CreateSecondaryVMEx_TaskResponse, error) {
	var reqBody, resBody CreateSecondaryVMEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateSecondaryVM_TaskBody struct {
	Req    *types.CreateSecondaryVM_Task         `xml:"urn:vim25 CreateSecondaryVM_Task,omitempty"`
	Res    *types.CreateSecondaryVM_TaskResponse `xml:"CreateSecondaryVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSecondaryVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateSecondaryVM_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateSecondaryVM_Task) (*types.CreateSecondaryVM_TaskResponse, error) {
	var reqBody, resBody CreateSecondaryVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateSnapshotEx_TaskBody struct {
	Req    *types.CreateSnapshotEx_Task         `xml:"urn:vim25 CreateSnapshotEx_Task,omitempty"`
	Res    *types.CreateSnapshotEx_TaskResponse `xml:"CreateSnapshotEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSnapshotEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateSnapshotEx_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateSnapshotEx_Task) (*types.CreateSnapshotEx_TaskResponse, error) {
	var reqBody, resBody CreateSnapshotEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateSnapshot_TaskBody struct {
	Req    *types.CreateSnapshot_Task         `xml:"urn:vim25 CreateSnapshot_Task,omitempty"`
	Res    *types.CreateSnapshot_TaskResponse `xml:"CreateSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateSnapshot_Task) (*types.CreateSnapshot_TaskResponse, error) {
	var reqBody, resBody CreateSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateSoftwareAdapterBody struct {
	Req    *types.CreateSoftwareAdapter         `xml:"urn:vim25 CreateSoftwareAdapter,omitempty"`
	Res    *types.CreateSoftwareAdapterResponse `xml:"CreateSoftwareAdapterResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSoftwareAdapterBody) Fault() *soap.Fault { return b.Fault_ }

func CreateSoftwareAdapter(ctx context.Context, r soap.RoundTripper, req *types.CreateSoftwareAdapter) (*types.CreateSoftwareAdapterResponse, error) {
	var reqBody, resBody CreateSoftwareAdapterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateStoragePodBody struct {
	Req    *types.CreateStoragePod         `xml:"urn:vim25 CreateStoragePod,omitempty"`
	Res    *types.CreateStoragePodResponse `xml:"CreateStoragePodResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateStoragePodBody) Fault() *soap.Fault { return b.Fault_ }

func CreateStoragePod(ctx context.Context, r soap.RoundTripper, req *types.CreateStoragePod) (*types.CreateStoragePodResponse, error) {
	var reqBody, resBody CreateStoragePodBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateTaskBody struct {
	Req    *types.CreateTask         `xml:"urn:vim25 CreateTask,omitempty"`
	Res    *types.CreateTaskResponse `xml:"CreateTaskResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateTask(ctx context.Context, r soap.RoundTripper, req *types.CreateTask) (*types.CreateTaskResponse, error) {
	var reqBody, resBody CreateTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateTemporaryDirectoryInGuestBody struct {
	Req    *types.CreateTemporaryDirectoryInGuest         `xml:"urn:vim25 CreateTemporaryDirectoryInGuest,omitempty"`
	Res    *types.CreateTemporaryDirectoryInGuestResponse `xml:"CreateTemporaryDirectoryInGuestResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTemporaryDirectoryInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func CreateTemporaryDirectoryInGuest(ctx context.Context, r soap.RoundTripper, req *types.CreateTemporaryDirectoryInGuest) (*types.CreateTemporaryDirectoryInGuestResponse, error) {
	var reqBody, resBody CreateTemporaryDirectoryInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateTemporaryFileInGuestBody struct {
	Req    *types.CreateTemporaryFileInGuest         `xml:"urn:vim25 CreateTemporaryFileInGuest,omitempty"`
	Res    *types.CreateTemporaryFileInGuestResponse `xml:"CreateTemporaryFileInGuestResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTemporaryFileInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func CreateTemporaryFileInGuest(ctx context.Context, r soap.RoundTripper, req *types.CreateTemporaryFileInGuest) (*types.CreateTemporaryFileInGuestResponse, error) {
	var reqBody, resBody CreateTemporaryFileInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateUserBody struct {
	Req    *types.CreateUser         `xml:"urn:vim25 CreateUser,omitempty"`
	Res    *types.CreateUserResponse `xml:"CreateUserResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateUserBody) Fault() *soap.Fault { return b.Fault_ }

func CreateUser(ctx context.Context, r soap.RoundTripper, req *types.CreateUser) (*types.CreateUserResponse, error) {
	var reqBody, resBody CreateUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateVAppBody struct {
	Req    *types.CreateVApp         `xml:"urn:vim25 CreateVApp,omitempty"`
	Res    *types.CreateVAppResponse `xml:"CreateVAppResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVAppBody) Fault() *soap.Fault { return b.Fault_ }

func CreateVApp(ctx context.Context, r soap.RoundTripper, req *types.CreateVApp) (*types.CreateVAppResponse, error) {
	var reqBody, resBody CreateVAppBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateVM_TaskBody struct {
	Req    *types.CreateVM_Task         `xml:"urn:vim25 CreateVM_Task,omitempty"`
	Res    *types.CreateVM_TaskResponse `xml:"CreateVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateVM_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateVM_Task) (*types.CreateVM_TaskResponse, error) {
	var reqBody, resBody CreateVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateVirtualDisk_TaskBody struct {
	Req    *types.CreateVirtualDisk_Task         `xml:"urn:vim25 CreateVirtualDisk_Task,omitempty"`
	Res    *types.CreateVirtualDisk_TaskResponse `xml:"CreateVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateVirtualDisk_Task) (*types.CreateVirtualDisk_TaskResponse, error) {
	var reqBody, resBody CreateVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateVmfsDatastoreBody struct {
	Req    *types.CreateVmfsDatastore         `xml:"urn:vim25 CreateVmfsDatastore,omitempty"`
	Res    *types.CreateVmfsDatastoreResponse `xml:"CreateVmfsDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVmfsDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func CreateVmfsDatastore(ctx context.Context, r soap.RoundTripper, req *types.CreateVmfsDatastore) (*types.CreateVmfsDatastoreResponse, error) {
	var reqBody, resBody CreateVmfsDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateVvolDatastoreBody struct {
	Req    *types.CreateVvolDatastore         `xml:"urn:vim25 CreateVvolDatastore,omitempty"`
	Res    *types.CreateVvolDatastoreResponse `xml:"CreateVvolDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVvolDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func CreateVvolDatastore(ctx context.Context, r soap.RoundTripper, req *types.CreateVvolDatastore) (*types.CreateVvolDatastoreResponse, error) {
	var reqBody, resBody CreateVvolDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CryptoManagerHostDisableBody struct {
	Req    *types.CryptoManagerHostDisable         `xml:"urn:vim25 CryptoManagerHostDisable,omitempty"`
	Res    *types.CryptoManagerHostDisableResponse `xml:"CryptoManagerHostDisableResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CryptoManagerHostDisableBody) Fault() *soap.Fault { return b.Fault_ }

func CryptoManagerHostDisable(ctx context.Context, r soap.RoundTripper, req *types.CryptoManagerHostDisable) (*types.CryptoManagerHostDisableResponse, error) {
	var reqBody, resBody CryptoManagerHostDisableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CryptoManagerHostEnableBody struct {
	Req    *types.CryptoManagerHostEnable         `xml:"urn:vim25 CryptoManagerHostEnable,omitempty"`
	Res    *types.CryptoManagerHostEnableResponse `xml:"CryptoManagerHostEnableResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CryptoManagerHostEnableBody) Fault() *soap.Fault { return b.Fault_ }

func CryptoManagerHostEnable(ctx context.Context, r soap.RoundTripper, req *types.CryptoManagerHostEnable) (*types.CryptoManagerHostEnableResponse, error) {
	var reqBody, resBody CryptoManagerHostEnableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CryptoManagerHostPrepareBody struct {
	Req    *types.CryptoManagerHostPrepare         `xml:"urn:vim25 CryptoManagerHostPrepare,omitempty"`
	Res    *types.CryptoManagerHostPrepareResponse `xml:"CryptoManagerHostPrepareResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CryptoManagerHostPrepareBody) Fault() *soap.Fault { return b.Fault_ }

func CryptoManagerHostPrepare(ctx context.Context, r soap.RoundTripper, req *types.CryptoManagerHostPrepare) (*types.CryptoManagerHostPrepareResponse, error) {
	var reqBody, resBody CryptoManagerHostPrepareBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CryptoUnlock_TaskBody struct {
	Req    *types.CryptoUnlock_Task         `xml:"urn:vim25 CryptoUnlock_Task,omitempty"`
	Res    *types.CryptoUnlock_TaskResponse `xml:"CryptoUnlock_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CryptoUnlock_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CryptoUnlock_Task(ctx context.Context, r soap.RoundTripper, req *types.CryptoUnlock_Task) (*types.CryptoUnlock_TaskResponse, error) {
	var reqBody, resBody CryptoUnlock_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CurrentTimeBody struct {
	Req    *types.CurrentTime         `xml:"urn:vim25 CurrentTime,omitempty"`
	Res    *types.CurrentTimeResponse `xml:"CurrentTimeResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CurrentTimeBody) Fault() *soap.Fault { return b.Fault_ }

func CurrentTime(ctx context.Context, r soap.RoundTripper, req *types.CurrentTime) (*types.CurrentTimeResponse, error) {
	var reqBody, resBody CurrentTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CustomizationSpecItemToXmlBody struct {
	Req    *types.CustomizationSpecItemToXml         `xml:"urn:vim25 CustomizationSpecItemToXml,omitempty"`
	Res    *types.CustomizationSpecItemToXmlResponse `xml:"CustomizationSpecItemToXmlResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CustomizationSpecItemToXmlBody) Fault() *soap.Fault { return b.Fault_ }

func CustomizationSpecItemToXml(ctx context.Context, r soap.RoundTripper, req *types.CustomizationSpecItemToXml) (*types.CustomizationSpecItemToXmlResponse, error) {
	var reqBody, resBody CustomizationSpecItemToXmlBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CustomizeGuest_TaskBody struct {
	Req    *types.CustomizeGuest_Task         `xml:"urn:vim25 CustomizeGuest_Task,omitempty"`
	Res    *types.CustomizeGuest_TaskResponse `xml:"CustomizeGuest_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CustomizeGuest_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CustomizeGuest_Task(ctx context.Context, r soap.RoundTripper, req *types.CustomizeGuest_Task) (*types.CustomizeGuest_TaskResponse, error) {
	var reqBody, resBody CustomizeGuest_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CustomizeVM_TaskBody struct {
	Req    *types.CustomizeVM_Task         `xml:"urn:vim25 CustomizeVM_Task,omitempty"`
	Res    *types.CustomizeVM_TaskResponse `xml:"CustomizeVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CustomizeVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CustomizeVM_Task(ctx context.Context, r soap.RoundTripper, req *types.CustomizeVM_Task) (*types.CustomizeVM_TaskResponse, error) {
	var reqBody, resBody CustomizeVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DVPortgroupRollback_TaskBody struct {
	Req    *types.DVPortgroupRollback_Task         `xml:"urn:vim25 DVPortgroupRollback_Task,omitempty"`
	Res    *types.DVPortgroupRollback_TaskResponse `xml:"DVPortgroupRollback_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVPortgroupRollback_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DVPortgroupRollback_Task(ctx context.Context, r soap.RoundTripper, req *types.DVPortgroupRollback_Task) (*types.DVPortgroupRollback_TaskResponse, error) {
	var reqBody, resBody DVPortgroupRollback_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DVSManagerExportEntity_TaskBody struct {
	Req    *types.DVSManagerExportEntity_Task         `xml:"urn:vim25 DVSManagerExportEntity_Task,omitempty"`
	Res    *types.DVSManagerExportEntity_TaskResponse `xml:"DVSManagerExportEntity_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerExportEntity_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DVSManagerExportEntity_Task(ctx context.Context, r soap.RoundTripper, req *types.DVSManagerExportEntity_Task) (*types.DVSManagerExportEntity_TaskResponse, error) {
	var reqBody, resBody DVSManagerExportEntity_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DVSManagerImportEntity_TaskBody struct {
	Req    *types.DVSManagerImportEntity_Task         `xml:"urn:vim25 DVSManagerImportEntity_Task,omitempty"`
	Res    *types.DVSManagerImportEntity_TaskResponse `xml:"DVSManagerImportEntity_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerImportEntity_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DVSManagerImportEntity_Task(ctx context.Context, r soap.RoundTripper, req *types.DVSManagerImportEntity_Task) (*types.DVSManagerImportEntity_TaskResponse, error) {
	var reqBody, resBody DVSManagerImportEntity_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DVSManagerLookupDvPortGroupBody struct {
	Req    *types.DVSManagerLookupDvPortGroup         `xml:"urn:vim25 DVSManagerLookupDvPortGroup,omitempty"`
	Res    *types.DVSManagerLookupDvPortGroupResponse `xml:"DVSManagerLookupDvPortGroupResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerLookupDvPortGroupBody) Fault() *soap.Fault { return b.Fault_ }

func DVSManagerLookupDvPortGroup(ctx context.Context, r soap.RoundTripper, req *types.DVSManagerLookupDvPortGroup) (*types.DVSManagerLookupDvPortGroupResponse, error) {
	var reqBody, resBody DVSManagerLookupDvPortGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DVSRollback_TaskBody struct {
	Req    *types.DVSRollback_Task         `xml:"urn:vim25 DVSRollback_Task,omitempty"`
	Res    *types.DVSRollback_TaskResponse `xml:"DVSRollback_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSRollback_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DVSRollback_Task(ctx context.Context, r soap.RoundTripper, req *types.DVSRollback_Task) (*types.DVSRollback_TaskResponse, error) {
	var reqBody, resBody DVSRollback_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DatastoreEnterMaintenanceModeBody struct {
	Req    *types.DatastoreEnterMaintenanceMode         `xml:"urn:vim25 DatastoreEnterMaintenanceMode,omitempty"`
	Res    *types.DatastoreEnterMaintenanceModeResponse `xml:"DatastoreEnterMaintenanceModeResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DatastoreEnterMaintenanceModeBody) Fault() *soap.Fault { return b.Fault_ }

func DatastoreEnterMaintenanceMode(ctx context.Context, r soap.RoundTripper, req *types.DatastoreEnterMaintenanceMode) (*types.DatastoreEnterMaintenanceModeResponse, error) {
	var reqBody, resBody DatastoreEnterMaintenanceModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DatastoreExitMaintenanceMode_TaskBody struct {
	Req    *types.DatastoreExitMaintenanceMode_Task         `xml:"urn:vim25 DatastoreExitMaintenanceMode_Task,omitempty"`
	Res    *types.DatastoreExitMaintenanceMode_TaskResponse `xml:"DatastoreExitMaintenanceMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DatastoreExitMaintenanceMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DatastoreExitMaintenanceMode_Task(ctx context.Context, r soap.RoundTripper, req *types.DatastoreExitMaintenanceMode_Task) (*types.DatastoreExitMaintenanceMode_TaskResponse, error) {
	var reqBody, resBody DatastoreExitMaintenanceMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DecodeLicenseBody struct {
	Req    *types.DecodeLicense         `xml:"urn:vim25 DecodeLicense,omitempty"`
	Res    *types.DecodeLicenseResponse `xml:"DecodeLicenseResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DecodeLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func DecodeLicense(ctx context.Context, r soap.RoundTripper, req *types.DecodeLicense) (*types.DecodeLicenseResponse, error) {
	var reqBody, resBody DecodeLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DefragmentAllDisksBody struct {
	Req    *types.DefragmentAllDisks         `xml:"urn:vim25 DefragmentAllDisks,omitempty"`
	Res    *types.DefragmentAllDisksResponse `xml:"DefragmentAllDisksResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DefragmentAllDisksBody) Fault() *soap.Fault { return b.Fault_ }

func DefragmentAllDisks(ctx context.Context, r soap.RoundTripper, req *types.DefragmentAllDisks) (*types.DefragmentAllDisksResponse, error) {
	var reqBody, resBody DefragmentAllDisksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DefragmentVirtualDisk_TaskBody struct {
	Req    *types.DefragmentVirtualDisk_Task         `xml:"urn:vim25 DefragmentVirtualDisk_Task,omitempty"`
	Res    *types.DefragmentVirtualDisk_TaskResponse `xml:"DefragmentVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DefragmentVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DefragmentVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.DefragmentVirtualDisk_Task) (*types.DefragmentVirtualDisk_TaskResponse, error) {
	var reqBody, resBody DefragmentVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteCustomizationSpecBody struct {
	Req    *types.DeleteCustomizationSpec         `xml:"urn:vim25 DeleteCustomizationSpec,omitempty"`
	Res    *types.DeleteCustomizationSpecResponse `xml:"DeleteCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.DeleteCustomizationSpec) (*types.DeleteCustomizationSpecResponse, error) {
	var reqBody, resBody DeleteCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteDatastoreFile_TaskBody struct {
	Req    *types.DeleteDatastoreFile_Task         `xml:"urn:vim25 DeleteDatastoreFile_Task,omitempty"`
	Res    *types.DeleteDatastoreFile_TaskResponse `xml:"DeleteDatastoreFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDatastoreFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteDatastoreFile_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteDatastoreFile_Task) (*types.DeleteDatastoreFile_TaskResponse, error) {
	var reqBody, resBody DeleteDatastoreFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteDirectoryBody struct {
	Req    *types.DeleteDirectory         `xml:"urn:vim25 DeleteDirectory,omitempty"`
	Res    *types.DeleteDirectoryResponse `xml:"DeleteDirectoryResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDirectoryBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteDirectory(ctx context.Context, r soap.RoundTripper, req *types.DeleteDirectory) (*types.DeleteDirectoryResponse, error) {
	var reqBody, resBody DeleteDirectoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteDirectoryInGuestBody struct {
	Req    *types.DeleteDirectoryInGuest         `xml:"urn:vim25 DeleteDirectoryInGuest,omitempty"`
	Res    *types.DeleteDirectoryInGuestResponse `xml:"DeleteDirectoryInGuestResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDirectoryInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteDirectoryInGuest(ctx context.Context, r soap.RoundTripper, req *types.DeleteDirectoryInGuest) (*types.DeleteDirectoryInGuestResponse, error) {
	var reqBody, resBody DeleteDirectoryInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteFileBody struct {
	Req    *types.DeleteFile         `xml:"urn:vim25 DeleteFile,omitempty"`
	Res    *types.DeleteFileResponse `xml:"DeleteFileResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteFileBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteFile(ctx context.Context, r soap.RoundTripper, req *types.DeleteFile) (*types.DeleteFileResponse, error) {
	var reqBody, resBody DeleteFileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteFileInGuestBody struct {
	Req    *types.DeleteFileInGuest         `xml:"urn:vim25 DeleteFileInGuest,omitempty"`
	Res    *types.DeleteFileInGuestResponse `xml:"DeleteFileInGuestResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteFileInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteFileInGuest(ctx context.Context, r soap.RoundTripper, req *types.DeleteFileInGuest) (*types.DeleteFileInGuestResponse, error) {
	var reqBody, resBody DeleteFileInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteHostSpecificationBody struct {
	Req    *types.DeleteHostSpecification         `xml:"urn:vim25 DeleteHostSpecification,omitempty"`
	Res    *types.DeleteHostSpecificationResponse `xml:"DeleteHostSpecificationResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteHostSpecificationBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteHostSpecification(ctx context.Context, r soap.RoundTripper, req *types.DeleteHostSpecification) (*types.DeleteHostSpecificationResponse, error) {
	var reqBody, resBody DeleteHostSpecificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteHostSubSpecificationBody struct {
	Req    *types.DeleteHostSubSpecification         `xml:"urn:vim25 DeleteHostSubSpecification,omitempty"`
	Res    *types.DeleteHostSubSpecificationResponse `xml:"DeleteHostSubSpecificationResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteHostSubSpecificationBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteHostSubSpecification(ctx context.Context, r soap.RoundTripper, req *types.DeleteHostSubSpecification) (*types.DeleteHostSubSpecificationResponse, error) {
	var reqBody, resBody DeleteHostSubSpecificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteNvdimmBlockNamespaces_TaskBody struct {
	Req    *types.DeleteNvdimmBlockNamespaces_Task         `xml:"urn:vim25 DeleteNvdimmBlockNamespaces_Task,omitempty"`
	Res    *types.DeleteNvdimmBlockNamespaces_TaskResponse `xml:"DeleteNvdimmBlockNamespaces_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteNvdimmBlockNamespaces_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteNvdimmBlockNamespaces_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteNvdimmBlockNamespaces_Task) (*types.DeleteNvdimmBlockNamespaces_TaskResponse, error) {
	var reqBody, resBody DeleteNvdimmBlockNamespaces_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteNvdimmNamespace_TaskBody struct {
	Req    *types.DeleteNvdimmNamespace_Task         `xml:"urn:vim25 DeleteNvdimmNamespace_Task,omitempty"`
	Res    *types.DeleteNvdimmNamespace_TaskResponse `xml:"DeleteNvdimmNamespace_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteNvdimmNamespace_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteNvdimmNamespace_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteNvdimmNamespace_Task) (*types.DeleteNvdimmNamespace_TaskResponse, error) {
	var reqBody, resBody DeleteNvdimmNamespace_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteRegistryKeyInGuestBody struct {
	Req    *types.DeleteRegistryKeyInGuest         `xml:"urn:vim25 DeleteRegistryKeyInGuest,omitempty"`
	Res    *types.DeleteRegistryKeyInGuestResponse `xml:"DeleteRegistryKeyInGuestResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteRegistryKeyInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteRegistryKeyInGuest(ctx context.Context, r soap.RoundTripper, req *types.DeleteRegistryKeyInGuest) (*types.DeleteRegistryKeyInGuestResponse, error) {
	var reqBody, resBody DeleteRegistryKeyInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteRegistryValueInGuestBody struct {
	Req    *types.DeleteRegistryValueInGuest         `xml:"urn:vim25 DeleteRegistryValueInGuest,omitempty"`
	Res    *types.DeleteRegistryValueInGuestResponse `xml:"DeleteRegistryValueInGuestResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteRegistryValueInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteRegistryValueInGuest(ctx context.Context, r soap.RoundTripper, req *types.DeleteRegistryValueInGuest) (*types.DeleteRegistryValueInGuestResponse, error) {
	var reqBody, resBody DeleteRegistryValueInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteScsiLunStateBody struct {
	Req    *types.DeleteScsiLunState         `xml:"urn:vim25 DeleteScsiLunState,omitempty"`
	Res    *types.DeleteScsiLunStateResponse `xml:"DeleteScsiLunStateResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteScsiLunStateBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteScsiLunState(ctx context.Context, r soap.RoundTripper, req *types.DeleteScsiLunState) (*types.DeleteScsiLunStateResponse, error) {
	var reqBody, resBody DeleteScsiLunStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteSnapshot_TaskBody struct {
	Req    *types.DeleteSnapshot_Task         `xml:"urn:vim25 DeleteSnapshot_Task,omitempty"`
	Res    *types.DeleteSnapshot_TaskResponse `xml:"DeleteSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteSnapshot_Task) (*types.DeleteSnapshot_TaskResponse, error) {
	var reqBody, resBody DeleteSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVStorageObjectEx_TaskBody struct {
	Req    *types.DeleteVStorageObjectEx_Task         `xml:"urn:vim25 DeleteVStorageObjectEx_Task,omitempty"`
	Res    *types.DeleteVStorageObjectEx_TaskResponse `xml:"DeleteVStorageObjectEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVStorageObjectEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVStorageObjectEx_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteVStorageObjectEx_Task) (*types.DeleteVStorageObjectEx_TaskResponse, error) {
	var reqBody, resBody DeleteVStorageObjectEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVStorageObject_TaskBody struct {
	Req    *types.DeleteVStorageObject_Task         `xml:"urn:vim25 DeleteVStorageObject_Task,omitempty"`
	Res    *types.DeleteVStorageObject_TaskResponse `xml:"DeleteVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteVStorageObject_Task) (*types.DeleteVStorageObject_TaskResponse, error) {
	var reqBody, resBody DeleteVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVffsVolumeStateBody struct {
	Req    *types.DeleteVffsVolumeState         `xml:"urn:vim25 DeleteVffsVolumeState,omitempty"`
	Res    *types.DeleteVffsVolumeStateResponse `xml:"DeleteVffsVolumeStateResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVffsVolumeStateBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVffsVolumeState(ctx context.Context, r soap.RoundTripper, req *types.DeleteVffsVolumeState) (*types.DeleteVffsVolumeStateResponse, error) {
	var reqBody, resBody DeleteVffsVolumeStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVirtualDisk_TaskBody struct {
	Req    *types.DeleteVirtualDisk_Task         `xml:"urn:vim25 DeleteVirtualDisk_Task,omitempty"`
	Res    *types.DeleteVirtualDisk_TaskResponse `xml:"DeleteVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.DeleteVirtualDisk_Task) (*types.DeleteVirtualDisk_TaskResponse, error) {
	var reqBody, resBody DeleteVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVmfsVolumeStateBody struct {
	Req    *types.DeleteVmfsVolumeState         `xml:"urn:vim25 DeleteVmfsVolumeState,omitempty"`
	Res    *types.DeleteVmfsVolumeStateResponse `xml:"DeleteVmfsVolumeStateResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVmfsVolumeStateBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVmfsVolumeState(ctx context.Context, r soap.RoundTripper, req *types.DeleteVmfsVolumeState) (*types.DeleteVmfsVolumeStateResponse, error) {
	var reqBody, resBody DeleteVmfsVolumeStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeleteVsanObjectsBody struct {
	Req    *types.DeleteVsanObjects         `xml:"urn:vim25 DeleteVsanObjects,omitempty"`
	Res    *types.DeleteVsanObjectsResponse `xml:"DeleteVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func DeleteVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.DeleteVsanObjects) (*types.DeleteVsanObjectsResponse, error) {
	var reqBody, resBody DeleteVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeselectVnicBody struct {
	Req    *types.DeselectVnic         `xml:"urn:vim25 DeselectVnic,omitempty"`
	Res    *types.DeselectVnicResponse `xml:"DeselectVnicResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeselectVnicBody) Fault() *soap.Fault { return b.Fault_ }

func DeselectVnic(ctx context.Context, r soap.RoundTripper, req *types.DeselectVnic) (*types.DeselectVnicResponse, error) {
	var reqBody, resBody DeselectVnicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeselectVnicForNicTypeBody struct {
	Req    *types.DeselectVnicForNicType         `xml:"urn:vim25 DeselectVnicForNicType,omitempty"`
	Res    *types.DeselectVnicForNicTypeResponse `xml:"DeselectVnicForNicTypeResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeselectVnicForNicTypeBody) Fault() *soap.Fault { return b.Fault_ }

func DeselectVnicForNicType(ctx context.Context, r soap.RoundTripper, req *types.DeselectVnicForNicType) (*types.DeselectVnicForNicTypeResponse, error) {
	var reqBody, resBody DeselectVnicForNicTypeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyChildrenBody struct {
	Req    *types.DestroyChildren         `xml:"urn:vim25 DestroyChildren,omitempty"`
	Res    *types.DestroyChildrenResponse `xml:"DestroyChildrenResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyChildrenBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyChildren(ctx context.Context, r soap.RoundTripper, req *types.DestroyChildren) (*types.DestroyChildrenResponse, error) {
	var reqBody, resBody DestroyChildrenBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyCollectorBody struct {
	Req    *types.DestroyCollector         `xml:"urn:vim25 DestroyCollector,omitempty"`
	Res    *types.DestroyCollectorResponse `xml:"DestroyCollectorResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyCollector(ctx context.Context, r soap.RoundTripper, req *types.DestroyCollector) (*types.DestroyCollectorResponse, error) {
	var reqBody, resBody DestroyCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyDatastoreBody struct {
	Req    *types.DestroyDatastore         `xml:"urn:vim25 DestroyDatastore,omitempty"`
	Res    *types.DestroyDatastoreResponse `xml:"DestroyDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyDatastore(ctx context.Context, r soap.RoundTripper, req *types.DestroyDatastore) (*types.DestroyDatastoreResponse, error) {
	var reqBody, resBody DestroyDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyIpPoolBody struct {
	Req    *types.DestroyIpPool         `xml:"urn:vim25 DestroyIpPool,omitempty"`
	Res    *types.DestroyIpPoolResponse `xml:"DestroyIpPoolResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyIpPoolBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyIpPool(ctx context.Context, r soap.RoundTripper, req *types.DestroyIpPool) (*types.DestroyIpPoolResponse, error) {
	var reqBody, resBody DestroyIpPoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyNetworkBody struct {
	Req    *types.DestroyNetwork         `xml:"urn:vim25 DestroyNetwork,omitempty"`
	Res    *types.DestroyNetworkResponse `xml:"DestroyNetworkResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyNetworkBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyNetwork(ctx context.Context, r soap.RoundTripper, req *types.DestroyNetwork) (*types.DestroyNetworkResponse, error) {
	var reqBody, resBody DestroyNetworkBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyProfileBody struct {
	Req    *types.DestroyProfile         `xml:"urn:vim25 DestroyProfile,omitempty"`
	Res    *types.DestroyProfileResponse `xml:"DestroyProfileResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyProfileBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyProfile(ctx context.Context, r soap.RoundTripper, req *types.DestroyProfile) (*types.DestroyProfileResponse, error) {
	var reqBody, resBody DestroyProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyPropertyCollectorBody struct {
	Req    *types.DestroyPropertyCollector         `xml:"urn:vim25 DestroyPropertyCollector,omitempty"`
	Res    *types.DestroyPropertyCollectorResponse `xml:"DestroyPropertyCollectorResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyPropertyCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyPropertyCollector(ctx context.Context, r soap.RoundTripper, req *types.DestroyPropertyCollector) (*types.DestroyPropertyCollectorResponse, error) {
	var reqBody, resBody DestroyPropertyCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyPropertyFilterBody struct {
	Req    *types.DestroyPropertyFilter         `xml:"urn:vim25 DestroyPropertyFilter,omitempty"`
	Res    *types.DestroyPropertyFilterResponse `xml:"DestroyPropertyFilterResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyPropertyFilterBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyPropertyFilter(ctx context.Context, r soap.RoundTripper, req *types.DestroyPropertyFilter) (*types.DestroyPropertyFilterResponse, error) {
	var reqBody, resBody DestroyPropertyFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyVffsBody struct {
	Req    *types.DestroyVffs         `xml:"urn:vim25 DestroyVffs,omitempty"`
	Res    *types.DestroyVffsResponse `xml:"DestroyVffsResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyVffsBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyVffs(ctx context.Context, r soap.RoundTripper, req *types.DestroyVffs) (*types.DestroyVffsResponse, error) {
	var reqBody, resBody DestroyVffsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyViewBody struct {
	Req    *types.DestroyView         `xml:"urn:vim25 DestroyView,omitempty"`
	Res    *types.DestroyViewResponse `xml:"DestroyViewResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyViewBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyView(ctx context.Context, r soap.RoundTripper, req *types.DestroyView) (*types.DestroyViewResponse, error) {
	var reqBody, resBody DestroyViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type Destroy_TaskBody struct {
	Req    *types.Destroy_Task         `xml:"urn:vim25 Destroy_Task,omitempty"`
	Res    *types.Destroy_TaskResponse `xml:"Destroy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Destroy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func Destroy_Task(ctx context.Context, r soap.RoundTripper, req *types.Destroy_Task) (*types.Destroy_TaskResponse, error) {
	var reqBody, resBody Destroy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DetachDisk_TaskBody struct {
	Req    *types.DetachDisk_Task         `xml:"urn:vim25 DetachDisk_Task,omitempty"`
	Res    *types.DetachDisk_TaskResponse `xml:"DetachDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DetachDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DetachDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.DetachDisk_Task) (*types.DetachDisk_TaskResponse, error) {
	var reqBody, resBody DetachDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DetachScsiLunBody struct {
	Req    *types.DetachScsiLun         `xml:"urn:vim25 DetachScsiLun,omitempty"`
	Res    *types.DetachScsiLunResponse `xml:"DetachScsiLunResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DetachScsiLunBody) Fault() *soap.Fault { return b.Fault_ }

func DetachScsiLun(ctx context.Context, r soap.RoundTripper, req *types.DetachScsiLun) (*types.DetachScsiLunResponse, error) {
	var reqBody, resBody DetachScsiLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DetachScsiLunEx_TaskBody struct {
	Req    *types.DetachScsiLunEx_Task         `xml:"urn:vim25 DetachScsiLunEx_Task,omitempty"`
	Res    *types.DetachScsiLunEx_TaskResponse `xml:"DetachScsiLunEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DetachScsiLunEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DetachScsiLunEx_Task(ctx context.Context, r soap.RoundTripper, req *types.DetachScsiLunEx_Task) (*types.DetachScsiLunEx_TaskResponse, error) {
	var reqBody, resBody DetachScsiLunEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DetachTagFromVStorageObjectBody struct {
	Req    *types.DetachTagFromVStorageObject         `xml:"urn:vim25 DetachTagFromVStorageObject,omitempty"`
	Res    *types.DetachTagFromVStorageObjectResponse `xml:"DetachTagFromVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DetachTagFromVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func DetachTagFromVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.DetachTagFromVStorageObject) (*types.DetachTagFromVStorageObjectResponse, error) {
	var reqBody, resBody DetachTagFromVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DirectPathProfileManagerCreateBody struct {
	Req    *types.DirectPathProfileManagerCreate         `xml:"urn:vim25 DirectPathProfileManagerCreate,omitempty"`
	Res    *types.DirectPathProfileManagerCreateResponse `xml:"DirectPathProfileManagerCreateResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DirectPathProfileManagerCreateBody) Fault() *soap.Fault { return b.Fault_ }

func DirectPathProfileManagerCreate(ctx context.Context, r soap.RoundTripper, req *types.DirectPathProfileManagerCreate) (*types.DirectPathProfileManagerCreateResponse, error) {
	var reqBody, resBody DirectPathProfileManagerCreateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DirectPathProfileManagerDeleteBody struct {
	Req    *types.DirectPathProfileManagerDelete         `xml:"urn:vim25 DirectPathProfileManagerDelete,omitempty"`
	Res    *types.DirectPathProfileManagerDeleteResponse `xml:"DirectPathProfileManagerDeleteResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DirectPathProfileManagerDeleteBody) Fault() *soap.Fault { return b.Fault_ }

func DirectPathProfileManagerDelete(ctx context.Context, r soap.RoundTripper, req *types.DirectPathProfileManagerDelete) (*types.DirectPathProfileManagerDeleteResponse, error) {
	var reqBody, resBody DirectPathProfileManagerDeleteBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DirectPathProfileManagerListBody struct {
	Req    *types.DirectPathProfileManagerList         `xml:"urn:vim25 DirectPathProfileManagerList,omitempty"`
	Res    *types.DirectPathProfileManagerListResponse `xml:"DirectPathProfileManagerListResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DirectPathProfileManagerListBody) Fault() *soap.Fault { return b.Fault_ }

func DirectPathProfileManagerList(ctx context.Context, r soap.RoundTripper, req *types.DirectPathProfileManagerList) (*types.DirectPathProfileManagerListResponse, error) {
	var reqBody, resBody DirectPathProfileManagerListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DirectPathProfileManagerQueryCapacityBody struct {
	Req    *types.DirectPathProfileManagerQueryCapacity         `xml:"urn:vim25 DirectPathProfileManagerQueryCapacity,omitempty"`
	Res    *types.DirectPathProfileManagerQueryCapacityResponse `xml:"DirectPathProfileManagerQueryCapacityResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DirectPathProfileManagerQueryCapacityBody) Fault() *soap.Fault { return b.Fault_ }

func DirectPathProfileManagerQueryCapacity(ctx context.Context, r soap.RoundTripper, req *types.DirectPathProfileManagerQueryCapacity) (*types.DirectPathProfileManagerQueryCapacityResponse, error) {
	var reqBody, resBody DirectPathProfileManagerQueryCapacityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DirectPathProfileManagerUpdateBody struct {
	Req    *types.DirectPathProfileManagerUpdate         `xml:"urn:vim25 DirectPathProfileManagerUpdate,omitempty"`
	Res    *types.DirectPathProfileManagerUpdateResponse `xml:"DirectPathProfileManagerUpdateResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DirectPathProfileManagerUpdateBody) Fault() *soap.Fault { return b.Fault_ }

func DirectPathProfileManagerUpdate(ctx context.Context, r soap.RoundTripper, req *types.DirectPathProfileManagerUpdate) (*types.DirectPathProfileManagerUpdateResponse, error) {
	var reqBody, resBody DirectPathProfileManagerUpdateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableAlarmBody struct {
	Req    *types.DisableAlarm         `xml:"urn:vim25 DisableAlarm,omitempty"`
	Res    *types.DisableAlarmResponse `xml:"DisableAlarmResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func DisableAlarm(ctx context.Context, r soap.RoundTripper, req *types.DisableAlarm) (*types.DisableAlarmResponse, error) {
	var reqBody, resBody DisableAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableClusteredVmdkSupportBody struct {
	Req    *types.DisableClusteredVmdkSupport         `xml:"urn:vim25 DisableClusteredVmdkSupport,omitempty"`
	Res    *types.DisableClusteredVmdkSupportResponse `xml:"DisableClusteredVmdkSupportResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableClusteredVmdkSupportBody) Fault() *soap.Fault { return b.Fault_ }

func DisableClusteredVmdkSupport(ctx context.Context, r soap.RoundTripper, req *types.DisableClusteredVmdkSupport) (*types.DisableClusteredVmdkSupportResponse, error) {
	var reqBody, resBody DisableClusteredVmdkSupportBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableEvcMode_TaskBody struct {
	Req    *types.DisableEvcMode_Task         `xml:"urn:vim25 DisableEvcMode_Task,omitempty"`
	Res    *types.DisableEvcMode_TaskResponse `xml:"DisableEvcMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableEvcMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DisableEvcMode_Task(ctx context.Context, r soap.RoundTripper, req *types.DisableEvcMode_Task) (*types.DisableEvcMode_TaskResponse, error) {
	var reqBody, resBody DisableEvcMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableFeatureBody struct {
	Req    *types.DisableFeature         `xml:"urn:vim25 DisableFeature,omitempty"`
	Res    *types.DisableFeatureResponse `xml:"DisableFeatureResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableFeatureBody) Fault() *soap.Fault { return b.Fault_ }

func DisableFeature(ctx context.Context, r soap.RoundTripper, req *types.DisableFeature) (*types.DisableFeatureResponse, error) {
	var reqBody, resBody DisableFeatureBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableHyperThreadingBody struct {
	Req    *types.DisableHyperThreading         `xml:"urn:vim25 DisableHyperThreading,omitempty"`
	Res    *types.DisableHyperThreadingResponse `xml:"DisableHyperThreadingResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableHyperThreadingBody) Fault() *soap.Fault { return b.Fault_ }

func DisableHyperThreading(ctx context.Context, r soap.RoundTripper, req *types.DisableHyperThreading) (*types.DisableHyperThreadingResponse, error) {
	var reqBody, resBody DisableHyperThreadingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableMultipathPathBody struct {
	Req    *types.DisableMultipathPath         `xml:"urn:vim25 DisableMultipathPath,omitempty"`
	Res    *types.DisableMultipathPathResponse `xml:"DisableMultipathPathResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableMultipathPathBody) Fault() *soap.Fault { return b.Fault_ }

func DisableMultipathPath(ctx context.Context, r soap.RoundTripper, req *types.DisableMultipathPath) (*types.DisableMultipathPathResponse, error) {
	var reqBody, resBody DisableMultipathPathBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableNetworkBoot_TaskBody struct {
	Req    *types.DisableNetworkBoot_Task         `xml:"urn:vim25 DisableNetworkBoot_Task,omitempty"`
	Res    *types.DisableNetworkBoot_TaskResponse `xml:"DisableNetworkBoot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableNetworkBoot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DisableNetworkBoot_Task(ctx context.Context, r soap.RoundTripper, req *types.DisableNetworkBoot_Task) (*types.DisableNetworkBoot_TaskResponse, error) {
	var reqBody, resBody DisableNetworkBoot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableRulesetBody struct {
	Req    *types.DisableRuleset         `xml:"urn:vim25 DisableRuleset,omitempty"`
	Res    *types.DisableRulesetResponse `xml:"DisableRulesetResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableRulesetBody) Fault() *soap.Fault { return b.Fault_ }

func DisableRuleset(ctx context.Context, r soap.RoundTripper, req *types.DisableRuleset) (*types.DisableRulesetResponse, error) {
	var reqBody, resBody DisableRulesetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableSecondaryVM_TaskBody struct {
	Req    *types.DisableSecondaryVM_Task         `xml:"urn:vim25 DisableSecondaryVM_Task,omitempty"`
	Res    *types.DisableSecondaryVM_TaskResponse `xml:"DisableSecondaryVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableSecondaryVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DisableSecondaryVM_Task(ctx context.Context, r soap.RoundTripper, req *types.DisableSecondaryVM_Task) (*types.DisableSecondaryVM_TaskResponse, error) {
	var reqBody, resBody DisableSecondaryVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisableSmartCardAuthenticationBody struct {
	Req    *types.DisableSmartCardAuthentication         `xml:"urn:vim25 DisableSmartCardAuthentication,omitempty"`
	Res    *types.DisableSmartCardAuthenticationResponse `xml:"DisableSmartCardAuthenticationResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableSmartCardAuthenticationBody) Fault() *soap.Fault { return b.Fault_ }

func DisableSmartCardAuthentication(ctx context.Context, r soap.RoundTripper, req *types.DisableSmartCardAuthentication) (*types.DisableSmartCardAuthenticationResponse, error) {
	var reqBody, resBody DisableSmartCardAuthenticationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisconnectHost_TaskBody struct {
	Req    *types.DisconnectHost_Task         `xml:"urn:vim25 DisconnectHost_Task,omitempty"`
	Res    *types.DisconnectHost_TaskResponse `xml:"DisconnectHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisconnectHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DisconnectHost_Task(ctx context.Context, r soap.RoundTripper, req *types.DisconnectHost_Task) (*types.DisconnectHost_TaskResponse, error) {
	var reqBody, resBody DisconnectHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisconnectNvmeControllerBody struct {
	Req    *types.DisconnectNvmeController         `xml:"urn:vim25 DisconnectNvmeController,omitempty"`
	Res    *types.DisconnectNvmeControllerResponse `xml:"DisconnectNvmeControllerResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisconnectNvmeControllerBody) Fault() *soap.Fault { return b.Fault_ }

func DisconnectNvmeController(ctx context.Context, r soap.RoundTripper, req *types.DisconnectNvmeController) (*types.DisconnectNvmeControllerResponse, error) {
	var reqBody, resBody DisconnectNvmeControllerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DisconnectNvmeControllerEx_TaskBody struct {
	Req    *types.DisconnectNvmeControllerEx_Task         `xml:"urn:vim25 DisconnectNvmeControllerEx_Task,omitempty"`
	Res    *types.DisconnectNvmeControllerEx_TaskResponse `xml:"DisconnectNvmeControllerEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisconnectNvmeControllerEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DisconnectNvmeControllerEx_Task(ctx context.Context, r soap.RoundTripper, req *types.DisconnectNvmeControllerEx_Task) (*types.DisconnectNvmeControllerEx_TaskResponse, error) {
	var reqBody, resBody DisconnectNvmeControllerEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DiscoverFcoeHbasBody struct {
	Req    *types.DiscoverFcoeHbas         `xml:"urn:vim25 DiscoverFcoeHbas,omitempty"`
	Res    *types.DiscoverFcoeHbasResponse `xml:"DiscoverFcoeHbasResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DiscoverFcoeHbasBody) Fault() *soap.Fault { return b.Fault_ }

func DiscoverFcoeHbas(ctx context.Context, r soap.RoundTripper, req *types.DiscoverFcoeHbas) (*types.DiscoverFcoeHbasResponse, error) {
	var reqBody, resBody DiscoverFcoeHbasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DiscoverNvmeControllersBody struct {
	Req    *types.DiscoverNvmeControllers         `xml:"urn:vim25 DiscoverNvmeControllers,omitempty"`
	Res    *types.DiscoverNvmeControllersResponse `xml:"DiscoverNvmeControllersResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DiscoverNvmeControllersBody) Fault() *soap.Fault { return b.Fault_ }

func DiscoverNvmeControllers(ctx context.Context, r soap.RoundTripper, req *types.DiscoverNvmeControllers) (*types.DiscoverNvmeControllersResponse, error) {
	var reqBody, resBody DiscoverNvmeControllersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DissociateProfileBody struct {
	Req    *types.DissociateProfile         `xml:"urn:vim25 DissociateProfile,omitempty"`
	Res    *types.DissociateProfileResponse `xml:"DissociateProfileResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DissociateProfileBody) Fault() *soap.Fault { return b.Fault_ }

func DissociateProfile(ctx context.Context, r soap.RoundTripper, req *types.DissociateProfile) (*types.DissociateProfileResponse, error) {
	var reqBody, resBody DissociateProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DoesCustomizationSpecExistBody struct {
	Req    *types.DoesCustomizationSpecExist         `xml:"urn:vim25 DoesCustomizationSpecExist,omitempty"`
	Res    *types.DoesCustomizationSpecExistResponse `xml:"DoesCustomizationSpecExistResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DoesCustomizationSpecExistBody) Fault() *soap.Fault { return b.Fault_ }

func DoesCustomizationSpecExist(ctx context.Context, r soap.RoundTripper, req *types.DoesCustomizationSpecExist) (*types.DoesCustomizationSpecExistResponse, error) {
	var reqBody, resBody DoesCustomizationSpecExistBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DownloadDescriptionTreeBody struct {
	Req    *types.DownloadDescriptionTree         `xml:"urn:vim25 DownloadDescriptionTree,omitempty"`
	Res    *types.DownloadDescriptionTreeResponse `xml:"DownloadDescriptionTreeResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DownloadDescriptionTreeBody) Fault() *soap.Fault { return b.Fault_ }

func DownloadDescriptionTree(ctx context.Context, r soap.RoundTripper, req *types.DownloadDescriptionTree) (*types.DownloadDescriptionTreeResponse, error) {
	var reqBody, resBody DownloadDescriptionTreeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DropConnectionsBody struct {
	Req    *types.DropConnections         `xml:"urn:vim25 DropConnections,omitempty"`
	Res    *types.DropConnectionsResponse `xml:"DropConnectionsResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DropConnectionsBody) Fault() *soap.Fault { return b.Fault_ }

func DropConnections(ctx context.Context, r soap.RoundTripper, req *types.DropConnections) (*types.DropConnectionsResponse, error) {
	var reqBody, resBody DropConnectionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DuplicateCustomizationSpecBody struct {
	Req    *types.DuplicateCustomizationSpec         `xml:"urn:vim25 DuplicateCustomizationSpec,omitempty"`
	Res    *types.DuplicateCustomizationSpecResponse `xml:"DuplicateCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DuplicateCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func DuplicateCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.DuplicateCustomizationSpec) (*types.DuplicateCustomizationSpecResponse, error) {
	var reqBody, resBody DuplicateCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DvsReconfigureVmVnicNetworkResourcePool_TaskBody struct {
	Req    *types.DvsReconfigureVmVnicNetworkResourcePool_Task         `xml:"urn:vim25 DvsReconfigureVmVnicNetworkResourcePool_Task,omitempty"`
	Res    *types.DvsReconfigureVmVnicNetworkResourcePool_TaskResponse `xml:"DvsReconfigureVmVnicNetworkResourcePool_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DvsReconfigureVmVnicNetworkResourcePool_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DvsReconfigureVmVnicNetworkResourcePool_Task(ctx context.Context, r soap.RoundTripper, req *types.DvsReconfigureVmVnicNetworkResourcePool_Task) (*types.DvsReconfigureVmVnicNetworkResourcePool_TaskResponse, error) {
	var reqBody, resBody DvsReconfigureVmVnicNetworkResourcePool_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EagerZeroVirtualDisk_TaskBody struct {
	Req    *types.EagerZeroVirtualDisk_Task         `xml:"urn:vim25 EagerZeroVirtualDisk_Task,omitempty"`
	Res    *types.EagerZeroVirtualDisk_TaskResponse `xml:"EagerZeroVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EagerZeroVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EagerZeroVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.EagerZeroVirtualDisk_Task) (*types.EagerZeroVirtualDisk_TaskResponse, error) {
	var reqBody, resBody EagerZeroVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EmitSyslogMarkBody struct {
	Req    *types.EmitSyslogMark         `xml:"urn:vim25 EmitSyslogMark,omitempty"`
	Res    *types.EmitSyslogMarkResponse `xml:"EmitSyslogMarkResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EmitSyslogMarkBody) Fault() *soap.Fault { return b.Fault_ }

func EmitSyslogMark(ctx context.Context, r soap.RoundTripper, req *types.EmitSyslogMark) (*types.EmitSyslogMarkResponse, error) {
	var reqBody, resBody EmitSyslogMarkBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableAlarmBody struct {
	Req    *types.EnableAlarm         `xml:"urn:vim25 EnableAlarm,omitempty"`
	Res    *types.EnableAlarmResponse `xml:"EnableAlarmResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func EnableAlarm(ctx context.Context, r soap.RoundTripper, req *types.EnableAlarm) (*types.EnableAlarmResponse, error) {
	var reqBody, resBody EnableAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableAlarmActionsBody struct {
	Req    *types.EnableAlarmActions         `xml:"urn:vim25 EnableAlarmActions,omitempty"`
	Res    *types.EnableAlarmActionsResponse `xml:"EnableAlarmActionsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableAlarmActionsBody) Fault() *soap.Fault { return b.Fault_ }

func EnableAlarmActions(ctx context.Context, r soap.RoundTripper, req *types.EnableAlarmActions) (*types.EnableAlarmActionsResponse, error) {
	var reqBody, resBody EnableAlarmActionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableClusteredVmdkSupportBody struct {
	Req    *types.EnableClusteredVmdkSupport         `xml:"urn:vim25 EnableClusteredVmdkSupport,omitempty"`
	Res    *types.EnableClusteredVmdkSupportResponse `xml:"EnableClusteredVmdkSupportResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableClusteredVmdkSupportBody) Fault() *soap.Fault { return b.Fault_ }

func EnableClusteredVmdkSupport(ctx context.Context, r soap.RoundTripper, req *types.EnableClusteredVmdkSupport) (*types.EnableClusteredVmdkSupportResponse, error) {
	var reqBody, resBody EnableClusteredVmdkSupportBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableCryptoBody struct {
	Req    *types.EnableCrypto         `xml:"urn:vim25 EnableCrypto,omitempty"`
	Res    *types.EnableCryptoResponse `xml:"EnableCryptoResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableCryptoBody) Fault() *soap.Fault { return b.Fault_ }

func EnableCrypto(ctx context.Context, r soap.RoundTripper, req *types.EnableCrypto) (*types.EnableCryptoResponse, error) {
	var reqBody, resBody EnableCryptoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableFeatureBody struct {
	Req    *types.EnableFeature         `xml:"urn:vim25 EnableFeature,omitempty"`
	Res    *types.EnableFeatureResponse `xml:"EnableFeatureResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableFeatureBody) Fault() *soap.Fault { return b.Fault_ }

func EnableFeature(ctx context.Context, r soap.RoundTripper, req *types.EnableFeature) (*types.EnableFeatureResponse, error) {
	var reqBody, resBody EnableFeatureBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableHyperThreadingBody struct {
	Req    *types.EnableHyperThreading         `xml:"urn:vim25 EnableHyperThreading,omitempty"`
	Res    *types.EnableHyperThreadingResponse `xml:"EnableHyperThreadingResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableHyperThreadingBody) Fault() *soap.Fault { return b.Fault_ }

func EnableHyperThreading(ctx context.Context, r soap.RoundTripper, req *types.EnableHyperThreading) (*types.EnableHyperThreadingResponse, error) {
	var reqBody, resBody EnableHyperThreadingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableMultipathPathBody struct {
	Req    *types.EnableMultipathPath         `xml:"urn:vim25 EnableMultipathPath,omitempty"`
	Res    *types.EnableMultipathPathResponse `xml:"EnableMultipathPathResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableMultipathPathBody) Fault() *soap.Fault { return b.Fault_ }

func EnableMultipathPath(ctx context.Context, r soap.RoundTripper, req *types.EnableMultipathPath) (*types.EnableMultipathPathResponse, error) {
	var reqBody, resBody EnableMultipathPathBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableNetworkBoot_TaskBody struct {
	Req    *types.EnableNetworkBoot_Task         `xml:"urn:vim25 EnableNetworkBoot_Task,omitempty"`
	Res    *types.EnableNetworkBoot_TaskResponse `xml:"EnableNetworkBoot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableNetworkBoot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EnableNetworkBoot_Task(ctx context.Context, r soap.RoundTripper, req *types.EnableNetworkBoot_Task) (*types.EnableNetworkBoot_TaskResponse, error) {
	var reqBody, resBody EnableNetworkBoot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableNetworkResourceManagementBody struct {
	Req    *types.EnableNetworkResourceManagement         `xml:"urn:vim25 EnableNetworkResourceManagement,omitempty"`
	Res    *types.EnableNetworkResourceManagementResponse `xml:"EnableNetworkResourceManagementResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableNetworkResourceManagementBody) Fault() *soap.Fault { return b.Fault_ }

func EnableNetworkResourceManagement(ctx context.Context, r soap.RoundTripper, req *types.EnableNetworkResourceManagement) (*types.EnableNetworkResourceManagementResponse, error) {
	var reqBody, resBody EnableNetworkResourceManagementBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableRulesetBody struct {
	Req    *types.EnableRuleset         `xml:"urn:vim25 EnableRuleset,omitempty"`
	Res    *types.EnableRulesetResponse `xml:"EnableRulesetResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableRulesetBody) Fault() *soap.Fault { return b.Fault_ }

func EnableRuleset(ctx context.Context, r soap.RoundTripper, req *types.EnableRuleset) (*types.EnableRulesetResponse, error) {
	var reqBody, resBody EnableRulesetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableSecondaryVM_TaskBody struct {
	Req    *types.EnableSecondaryVM_Task         `xml:"urn:vim25 EnableSecondaryVM_Task,omitempty"`
	Res    *types.EnableSecondaryVM_TaskResponse `xml:"EnableSecondaryVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableSecondaryVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EnableSecondaryVM_Task(ctx context.Context, r soap.RoundTripper, req *types.EnableSecondaryVM_Task) (*types.EnableSecondaryVM_TaskResponse, error) {
	var reqBody, resBody EnableSecondaryVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnableSmartCardAuthenticationBody struct {
	Req    *types.EnableSmartCardAuthentication         `xml:"urn:vim25 EnableSmartCardAuthentication,omitempty"`
	Res    *types.EnableSmartCardAuthenticationResponse `xml:"EnableSmartCardAuthenticationResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableSmartCardAuthenticationBody) Fault() *soap.Fault { return b.Fault_ }

func EnableSmartCardAuthentication(ctx context.Context, r soap.RoundTripper, req *types.EnableSmartCardAuthentication) (*types.EnableSmartCardAuthenticationResponse, error) {
	var reqBody, resBody EnableSmartCardAuthenticationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnterLockdownModeBody struct {
	Req    *types.EnterLockdownMode         `xml:"urn:vim25 EnterLockdownMode,omitempty"`
	Res    *types.EnterLockdownModeResponse `xml:"EnterLockdownModeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnterLockdownModeBody) Fault() *soap.Fault { return b.Fault_ }

func EnterLockdownMode(ctx context.Context, r soap.RoundTripper, req *types.EnterLockdownMode) (*types.EnterLockdownModeResponse, error) {
	var reqBody, resBody EnterLockdownModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EnterMaintenanceMode_TaskBody struct {
	Req    *types.EnterMaintenanceMode_Task         `xml:"urn:vim25 EnterMaintenanceMode_Task,omitempty"`
	Res    *types.EnterMaintenanceMode_TaskResponse `xml:"EnterMaintenanceMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnterMaintenanceMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EnterMaintenanceMode_Task(ctx context.Context, r soap.RoundTripper, req *types.EnterMaintenanceMode_Task) (*types.EnterMaintenanceMode_TaskResponse, error) {
	var reqBody, resBody EnterMaintenanceMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EstimateDatabaseSizeBody struct {
	Req    *types.EstimateDatabaseSize         `xml:"urn:vim25 EstimateDatabaseSize,omitempty"`
	Res    *types.EstimateDatabaseSizeResponse `xml:"EstimateDatabaseSizeResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EstimateDatabaseSizeBody) Fault() *soap.Fault { return b.Fault_ }

func EstimateDatabaseSize(ctx context.Context, r soap.RoundTripper, req *types.EstimateDatabaseSize) (*types.EstimateDatabaseSizeResponse, error) {
	var reqBody, resBody EstimateDatabaseSizeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EstimateStorageForConsolidateSnapshots_TaskBody struct {
	Req    *types.EstimateStorageForConsolidateSnapshots_Task         `xml:"urn:vim25 EstimateStorageForConsolidateSnapshots_Task,omitempty"`
	Res    *types.EstimateStorageForConsolidateSnapshots_TaskResponse `xml:"EstimateStorageForConsolidateSnapshots_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EstimateStorageForConsolidateSnapshots_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EstimateStorageForConsolidateSnapshots_Task(ctx context.Context, r soap.RoundTripper, req *types.EstimateStorageForConsolidateSnapshots_Task) (*types.EstimateStorageForConsolidateSnapshots_TaskResponse, error) {
	var reqBody, resBody EstimateStorageForConsolidateSnapshots_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EsxAgentHostManagerUpdateConfigBody struct {
	Req    *types.EsxAgentHostManagerUpdateConfig         `xml:"urn:vim25 EsxAgentHostManagerUpdateConfig,omitempty"`
	Res    *types.EsxAgentHostManagerUpdateConfigResponse `xml:"EsxAgentHostManagerUpdateConfigResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EsxAgentHostManagerUpdateConfigBody) Fault() *soap.Fault { return b.Fault_ }

func EsxAgentHostManagerUpdateConfig(ctx context.Context, r soap.RoundTripper, req *types.EsxAgentHostManagerUpdateConfig) (*types.EsxAgentHostManagerUpdateConfigResponse, error) {
	var reqBody, resBody EsxAgentHostManagerUpdateConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EvacuateVsanNode_TaskBody struct {
	Req    *types.EvacuateVsanNode_Task         `xml:"urn:vim25 EvacuateVsanNode_Task,omitempty"`
	Res    *types.EvacuateVsanNode_TaskResponse `xml:"EvacuateVsanNode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EvacuateVsanNode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func EvacuateVsanNode_Task(ctx context.Context, r soap.RoundTripper, req *types.EvacuateVsanNode_Task) (*types.EvacuateVsanNode_TaskResponse, error) {
	var reqBody, resBody EvacuateVsanNode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type EvcManagerBody struct {
	Req    *types.EvcManager         `xml:"urn:vim25 EvcManager,omitempty"`
	Res    *types.EvcManagerResponse `xml:"EvcManagerResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EvcManagerBody) Fault() *soap.Fault { return b.Fault_ }

func EvcManager(ctx context.Context, r soap.RoundTripper, req *types.EvcManager) (*types.EvcManagerResponse, error) {
	var reqBody, resBody EvcManagerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExecuteHostProfileBody struct {
	Req    *types.ExecuteHostProfile         `xml:"urn:vim25 ExecuteHostProfile,omitempty"`
	Res    *types.ExecuteHostProfileResponse `xml:"ExecuteHostProfileResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExecuteHostProfileBody) Fault() *soap.Fault { return b.Fault_ }

func ExecuteHostProfile(ctx context.Context, r soap.RoundTripper, req *types.ExecuteHostProfile) (*types.ExecuteHostProfileResponse, error) {
	var reqBody, resBody ExecuteHostProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExecuteSimpleCommandBody struct {
	Req    *types.ExecuteSimpleCommand         `xml:"urn:vim25 ExecuteSimpleCommand,omitempty"`
	Res    *types.ExecuteSimpleCommandResponse `xml:"ExecuteSimpleCommandResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExecuteSimpleCommandBody) Fault() *soap.Fault { return b.Fault_ }

func ExecuteSimpleCommand(ctx context.Context, r soap.RoundTripper, req *types.ExecuteSimpleCommand) (*types.ExecuteSimpleCommandResponse, error) {
	var reqBody, resBody ExecuteSimpleCommandBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExitLockdownModeBody struct {
	Req    *types.ExitLockdownMode         `xml:"urn:vim25 ExitLockdownMode,omitempty"`
	Res    *types.ExitLockdownModeResponse `xml:"ExitLockdownModeResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExitLockdownModeBody) Fault() *soap.Fault { return b.Fault_ }

func ExitLockdownMode(ctx context.Context, r soap.RoundTripper, req *types.ExitLockdownMode) (*types.ExitLockdownModeResponse, error) {
	var reqBody, resBody ExitLockdownModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExitMaintenanceMode_TaskBody struct {
	Req    *types.ExitMaintenanceMode_Task         `xml:"urn:vim25 ExitMaintenanceMode_Task,omitempty"`
	Res    *types.ExitMaintenanceMode_TaskResponse `xml:"ExitMaintenanceMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExitMaintenanceMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ExitMaintenanceMode_Task(ctx context.Context, r soap.RoundTripper, req *types.ExitMaintenanceMode_Task) (*types.ExitMaintenanceMode_TaskResponse, error) {
	var reqBody, resBody ExitMaintenanceMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExpandVmfsDatastoreBody struct {
	Req    *types.ExpandVmfsDatastore         `xml:"urn:vim25 ExpandVmfsDatastore,omitempty"`
	Res    *types.ExpandVmfsDatastoreResponse `xml:"ExpandVmfsDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExpandVmfsDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func ExpandVmfsDatastore(ctx context.Context, r soap.RoundTripper, req *types.ExpandVmfsDatastore) (*types.ExpandVmfsDatastoreResponse, error) {
	var reqBody, resBody ExpandVmfsDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExpandVmfsExtentBody struct {
	Req    *types.ExpandVmfsExtent         `xml:"urn:vim25 ExpandVmfsExtent,omitempty"`
	Res    *types.ExpandVmfsExtentResponse `xml:"ExpandVmfsExtentResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExpandVmfsExtentBody) Fault() *soap.Fault { return b.Fault_ }

func ExpandVmfsExtent(ctx context.Context, r soap.RoundTripper, req *types.ExpandVmfsExtent) (*types.ExpandVmfsExtentResponse, error) {
	var reqBody, resBody ExpandVmfsExtentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExportAnswerFile_TaskBody struct {
	Req    *types.ExportAnswerFile_Task         `xml:"urn:vim25 ExportAnswerFile_Task,omitempty"`
	Res    *types.ExportAnswerFile_TaskResponse `xml:"ExportAnswerFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportAnswerFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ExportAnswerFile_Task(ctx context.Context, r soap.RoundTripper, req *types.ExportAnswerFile_Task) (*types.ExportAnswerFile_TaskResponse, error) {
	var reqBody, resBody ExportAnswerFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExportProfileBody struct {
	Req    *types.ExportProfile         `xml:"urn:vim25 ExportProfile,omitempty"`
	Res    *types.ExportProfileResponse `xml:"ExportProfileResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportProfileBody) Fault() *soap.Fault { return b.Fault_ }

func ExportProfile(ctx context.Context, r soap.RoundTripper, req *types.ExportProfile) (*types.ExportProfileResponse, error) {
	var reqBody, resBody ExportProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExportSnapshotBody struct {
	Req    *types.ExportSnapshot         `xml:"urn:vim25 ExportSnapshot,omitempty"`
	Res    *types.ExportSnapshotResponse `xml:"ExportSnapshotResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportSnapshotBody) Fault() *soap.Fault { return b.Fault_ }

func ExportSnapshot(ctx context.Context, r soap.RoundTripper, req *types.ExportSnapshot) (*types.ExportSnapshotResponse, error) {
	var reqBody, resBody ExportSnapshotBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExportVAppBody struct {
	Req    *types.ExportVApp         `xml:"urn:vim25 ExportVApp,omitempty"`
	Res    *types.ExportVAppResponse `xml:"ExportVAppResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportVAppBody) Fault() *soap.Fault { return b.Fault_ }

func ExportVApp(ctx context.Context, r soap.RoundTripper, req *types.ExportVApp) (*types.ExportVAppResponse, error) {
	var reqBody, resBody ExportVAppBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExportVmBody struct {
	Req    *types.ExportVm         `xml:"urn:vim25 ExportVm,omitempty"`
	Res    *types.ExportVmResponse `xml:"ExportVmResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportVmBody) Fault() *soap.Fault { return b.Fault_ }

func ExportVm(ctx context.Context, r soap.RoundTripper, req *types.ExportVm) (*types.ExportVmResponse, error) {
	var reqBody, resBody ExportVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtendDisk_TaskBody struct {
	Req    *types.ExtendDisk_Task         `xml:"urn:vim25 ExtendDisk_Task,omitempty"`
	Res    *types.ExtendDisk_TaskResponse `xml:"ExtendDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ExtendDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.ExtendDisk_Task) (*types.ExtendDisk_TaskResponse, error) {
	var reqBody, resBody ExtendDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtendHCI_TaskBody struct {
	Req    *types.ExtendHCI_Task         `xml:"urn:vim25 ExtendHCI_Task,omitempty"`
	Res    *types.ExtendHCI_TaskResponse `xml:"ExtendHCI_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendHCI_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ExtendHCI_Task(ctx context.Context, r soap.RoundTripper, req *types.ExtendHCI_Task) (*types.ExtendHCI_TaskResponse, error) {
	var reqBody, resBody ExtendHCI_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtendVffsBody struct {
	Req    *types.ExtendVffs         `xml:"urn:vim25 ExtendVffs,omitempty"`
	Res    *types.ExtendVffsResponse `xml:"ExtendVffsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVffsBody) Fault() *soap.Fault { return b.Fault_ }

func ExtendVffs(ctx context.Context, r soap.RoundTripper, req *types.ExtendVffs) (*types.ExtendVffsResponse, error) {
	var reqBody, resBody ExtendVffsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtendVirtualDisk_TaskBody struct {
	Req    *types.ExtendVirtualDisk_Task         `xml:"urn:vim25 ExtendVirtualDisk_Task,omitempty"`
	Res    *types.ExtendVirtualDisk_TaskResponse `xml:"ExtendVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ExtendVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.ExtendVirtualDisk_Task) (*types.ExtendVirtualDisk_TaskResponse, error) {
	var reqBody, resBody ExtendVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtendVmfsDatastoreBody struct {
	Req    *types.ExtendVmfsDatastore         `xml:"urn:vim25 ExtendVmfsDatastore,omitempty"`
	Res    *types.ExtendVmfsDatastoreResponse `xml:"ExtendVmfsDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVmfsDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func ExtendVmfsDatastore(ctx context.Context, r soap.RoundTripper, req *types.ExtendVmfsDatastore) (*types.ExtendVmfsDatastoreResponse, error) {
	var reqBody, resBody ExtendVmfsDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ExtractOvfEnvironmentBody struct {
	Req    *types.ExtractOvfEnvironment         `xml:"urn:vim25 ExtractOvfEnvironment,omitempty"`
	Res    *types.ExtractOvfEnvironmentResponse `xml:"ExtractOvfEnvironmentResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtractOvfEnvironmentBody) Fault() *soap.Fault { return b.Fault_ }

func ExtractOvfEnvironment(ctx context.Context, r soap.RoundTripper, req *types.ExtractOvfEnvironment) (*types.ExtractOvfEnvironmentResponse, error) {
	var reqBody, resBody ExtractOvfEnvironmentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchAuditRecordsBody struct {
	Req    *types.FetchAuditRecords         `xml:"urn:vim25 FetchAuditRecords,omitempty"`
	Res    *types.FetchAuditRecordsResponse `xml:"FetchAuditRecordsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchAuditRecordsBody) Fault() *soap.Fault { return b.Fault_ }

func FetchAuditRecords(ctx context.Context, r soap.RoundTripper, req *types.FetchAuditRecords) (*types.FetchAuditRecordsResponse, error) {
	var reqBody, resBody FetchAuditRecordsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchDVPortKeysBody struct {
	Req    *types.FetchDVPortKeys         `xml:"urn:vim25 FetchDVPortKeys,omitempty"`
	Res    *types.FetchDVPortKeysResponse `xml:"FetchDVPortKeysResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchDVPortKeysBody) Fault() *soap.Fault { return b.Fault_ }

func FetchDVPortKeys(ctx context.Context, r soap.RoundTripper, req *types.FetchDVPortKeys) (*types.FetchDVPortKeysResponse, error) {
	var reqBody, resBody FetchDVPortKeysBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchDVPortsBody struct {
	Req    *types.FetchDVPorts         `xml:"urn:vim25 FetchDVPorts,omitempty"`
	Res    *types.FetchDVPortsResponse `xml:"FetchDVPortsResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchDVPortsBody) Fault() *soap.Fault { return b.Fault_ }

func FetchDVPorts(ctx context.Context, r soap.RoundTripper, req *types.FetchDVPorts) (*types.FetchDVPortsResponse, error) {
	var reqBody, resBody FetchDVPortsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchSystemEventLogBody struct {
	Req    *types.FetchSystemEventLog         `xml:"urn:vim25 FetchSystemEventLog,omitempty"`
	Res    *types.FetchSystemEventLogResponse `xml:"FetchSystemEventLogResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchSystemEventLogBody) Fault() *soap.Fault { return b.Fault_ }

func FetchSystemEventLog(ctx context.Context, r soap.RoundTripper, req *types.FetchSystemEventLog) (*types.FetchSystemEventLogResponse, error) {
	var reqBody, resBody FetchSystemEventLogBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchUserPrivilegeOnEntitiesBody struct {
	Req    *types.FetchUserPrivilegeOnEntities         `xml:"urn:vim25 FetchUserPrivilegeOnEntities,omitempty"`
	Res    *types.FetchUserPrivilegeOnEntitiesResponse `xml:"FetchUserPrivilegeOnEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchUserPrivilegeOnEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func FetchUserPrivilegeOnEntities(ctx context.Context, r soap.RoundTripper, req *types.FetchUserPrivilegeOnEntities) (*types.FetchUserPrivilegeOnEntitiesResponse, error) {
	var reqBody, resBody FetchUserPrivilegeOnEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindAllByDnsNameBody struct {
	Req    *types.FindAllByDnsName         `xml:"urn:vim25 FindAllByDnsName,omitempty"`
	Res    *types.FindAllByDnsNameResponse `xml:"FindAllByDnsNameResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByDnsNameBody) Fault() *soap.Fault { return b.Fault_ }

func FindAllByDnsName(ctx context.Context, r soap.RoundTripper, req *types.FindAllByDnsName) (*types.FindAllByDnsNameResponse, error) {
	var reqBody, resBody FindAllByDnsNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindAllByIpBody struct {
	Req    *types.FindAllByIp         `xml:"urn:vim25 FindAllByIp,omitempty"`
	Res    *types.FindAllByIpResponse `xml:"FindAllByIpResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByIpBody) Fault() *soap.Fault { return b.Fault_ }

func FindAllByIp(ctx context.Context, r soap.RoundTripper, req *types.FindAllByIp) (*types.FindAllByIpResponse, error) {
	var reqBody, resBody FindAllByIpBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindAllByUuidBody struct {
	Req    *types.FindAllByUuid         `xml:"urn:vim25 FindAllByUuid,omitempty"`
	Res    *types.FindAllByUuidResponse `xml:"FindAllByUuidResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByUuidBody) Fault() *soap.Fault { return b.Fault_ }

func FindAllByUuid(ctx context.Context, r soap.RoundTripper, req *types.FindAllByUuid) (*types.FindAllByUuidResponse, error) {
	var reqBody, resBody FindAllByUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindAssociatedProfileBody struct {
	Req    *types.FindAssociatedProfile         `xml:"urn:vim25 FindAssociatedProfile,omitempty"`
	Res    *types.FindAssociatedProfileResponse `xml:"FindAssociatedProfileResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAssociatedProfileBody) Fault() *soap.Fault { return b.Fault_ }

func FindAssociatedProfile(ctx context.Context, r soap.RoundTripper, req *types.FindAssociatedProfile) (*types.FindAssociatedProfileResponse, error) {
	var reqBody, resBody FindAssociatedProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindByDatastorePathBody struct {
	Req    *types.FindByDatastorePath         `xml:"urn:vim25 FindByDatastorePath,omitempty"`
	Res    *types.FindByDatastorePathResponse `xml:"FindByDatastorePathResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByDatastorePathBody) Fault() *soap.Fault { return b.Fault_ }

func FindByDatastorePath(ctx context.Context, r soap.RoundTripper, req *types.FindByDatastorePath) (*types.FindByDatastorePathResponse, error) {
	var reqBody, resBody FindByDatastorePathBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindByDnsNameBody struct {
	Req    *types.FindByDnsName         `xml:"urn:vim25 FindByDnsName,omitempty"`
	Res    *types.FindByDnsNameResponse `xml:"FindByDnsNameResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByDnsNameBody) Fault() *soap.Fault { return b.Fault_ }

func FindByDnsName(ctx context.Context, r soap.RoundTripper, req *types.FindByDnsName) (*types.FindByDnsNameResponse, error) {
	var reqBody, resBody FindByDnsNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindByInventoryPathBody struct {
	Req    *types.FindByInventoryPath         `xml:"urn:vim25 FindByInventoryPath,omitempty"`
	Res    *types.FindByInventoryPathResponse `xml:"FindByInventoryPathResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByInventoryPathBody) Fault() *soap.Fault { return b.Fault_ }

func FindByInventoryPath(ctx context.Context, r soap.RoundTripper, req *types.FindByInventoryPath) (*types.FindByInventoryPathResponse, error) {
	var reqBody, resBody FindByInventoryPathBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindByIpBody struct {
	Req    *types.FindByIp         `xml:"urn:vim25 FindByIp,omitempty"`
	Res    *types.FindByIpResponse `xml:"FindByIpResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByIpBody) Fault() *soap.Fault { return b.Fault_ }

func FindByIp(ctx context.Context, r soap.RoundTripper, req *types.FindByIp) (*types.FindByIpResponse, error) {
	var reqBody, resBody FindByIpBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindByUuidBody struct {
	Req    *types.FindByUuid         `xml:"urn:vim25 FindByUuid,omitempty"`
	Res    *types.FindByUuidResponse `xml:"FindByUuidResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByUuidBody) Fault() *soap.Fault { return b.Fault_ }

func FindByUuid(ctx context.Context, r soap.RoundTripper, req *types.FindByUuid) (*types.FindByUuidResponse, error) {
	var reqBody, resBody FindByUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindChildBody struct {
	Req    *types.FindChild         `xml:"urn:vim25 FindChild,omitempty"`
	Res    *types.FindChildResponse `xml:"FindChildResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindChildBody) Fault() *soap.Fault { return b.Fault_ }

func FindChild(ctx context.Context, r soap.RoundTripper, req *types.FindChild) (*types.FindChildResponse, error) {
	var reqBody, resBody FindChildBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindExtensionBody struct {
	Req    *types.FindExtension         `xml:"urn:vim25 FindExtension,omitempty"`
	Res    *types.FindExtensionResponse `xml:"FindExtensionResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindExtensionBody) Fault() *soap.Fault { return b.Fault_ }

func FindExtension(ctx context.Context, r soap.RoundTripper, req *types.FindExtension) (*types.FindExtensionResponse, error) {
	var reqBody, resBody FindExtensionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FindRulesForVmBody struct {
	Req    *types.FindRulesForVm         `xml:"urn:vim25 FindRulesForVm,omitempty"`
	Res    *types.FindRulesForVmResponse `xml:"FindRulesForVmResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindRulesForVmBody) Fault() *soap.Fault { return b.Fault_ }

func FindRulesForVm(ctx context.Context, r soap.RoundTripper, req *types.FindRulesForVm) (*types.FindRulesForVmResponse, error) {
	var reqBody, resBody FindRulesForVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FormatVffsBody struct {
	Req    *types.FormatVffs         `xml:"urn:vim25 FormatVffs,omitempty"`
	Res    *types.FormatVffsResponse `xml:"FormatVffsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FormatVffsBody) Fault() *soap.Fault { return b.Fault_ }

func FormatVffs(ctx context.Context, r soap.RoundTripper, req *types.FormatVffs) (*types.FormatVffsResponse, error) {
	var reqBody, resBody FormatVffsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FormatVmfsBody struct {
	Req    *types.FormatVmfs         `xml:"urn:vim25 FormatVmfs,omitempty"`
	Res    *types.FormatVmfsResponse `xml:"FormatVmfsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FormatVmfsBody) Fault() *soap.Fault { return b.Fault_ }

func FormatVmfs(ctx context.Context, r soap.RoundTripper, req *types.FormatVmfs) (*types.FormatVmfsResponse, error) {
	var reqBody, resBody FormatVmfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateCertificateSigningRequestBody struct {
	Req    *types.GenerateCertificateSigningRequest         `xml:"urn:vim25 GenerateCertificateSigningRequest,omitempty"`
	Res    *types.GenerateCertificateSigningRequestResponse `xml:"GenerateCertificateSigningRequestResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateCertificateSigningRequestBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateCertificateSigningRequest(ctx context.Context, r soap.RoundTripper, req *types.GenerateCertificateSigningRequest) (*types.GenerateCertificateSigningRequestResponse, error) {
	var reqBody, resBody GenerateCertificateSigningRequestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateCertificateSigningRequestByDnBody struct {
	Req    *types.GenerateCertificateSigningRequestByDn         `xml:"urn:vim25 GenerateCertificateSigningRequestByDn,omitempty"`
	Res    *types.GenerateCertificateSigningRequestByDnResponse `xml:"GenerateCertificateSigningRequestByDnResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateCertificateSigningRequestByDnBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateCertificateSigningRequestByDn(ctx context.Context, r soap.RoundTripper, req *types.GenerateCertificateSigningRequestByDn) (*types.GenerateCertificateSigningRequestByDnResponse, error) {
	var reqBody, resBody GenerateCertificateSigningRequestByDnBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateClientCsrBody struct {
	Req    *types.GenerateClientCsr         `xml:"urn:vim25 GenerateClientCsr,omitempty"`
	Res    *types.GenerateClientCsrResponse `xml:"GenerateClientCsrResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateClientCsrBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateClientCsr(ctx context.Context, r soap.RoundTripper, req *types.GenerateClientCsr) (*types.GenerateClientCsrResponse, error) {
	var reqBody, resBody GenerateClientCsrBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateConfigTaskListBody struct {
	Req    *types.GenerateConfigTaskList         `xml:"urn:vim25 GenerateConfigTaskList,omitempty"`
	Res    *types.GenerateConfigTaskListResponse `xml:"GenerateConfigTaskListResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateConfigTaskListBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateConfigTaskList(ctx context.Context, r soap.RoundTripper, req *types.GenerateConfigTaskList) (*types.GenerateConfigTaskListResponse, error) {
	var reqBody, resBody GenerateConfigTaskListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateHostConfigTaskSpec_TaskBody struct {
	Req    *types.GenerateHostConfigTaskSpec_Task         `xml:"urn:vim25 GenerateHostConfigTaskSpec_Task,omitempty"`
	Res    *types.GenerateHostConfigTaskSpec_TaskResponse `xml:"GenerateHostConfigTaskSpec_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateHostConfigTaskSpec_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateHostConfigTaskSpec_Task(ctx context.Context, r soap.RoundTripper, req *types.GenerateHostConfigTaskSpec_Task) (*types.GenerateHostConfigTaskSpec_TaskResponse, error) {
	var reqBody, resBody GenerateHostConfigTaskSpec_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateHostProfileTaskList_TaskBody struct {
	Req    *types.GenerateHostProfileTaskList_Task         `xml:"urn:vim25 GenerateHostProfileTaskList_Task,omitempty"`
	Res    *types.GenerateHostProfileTaskList_TaskResponse `xml:"GenerateHostProfileTaskList_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateHostProfileTaskList_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateHostProfileTaskList_Task(ctx context.Context, r soap.RoundTripper, req *types.GenerateHostProfileTaskList_Task) (*types.GenerateHostProfileTaskList_TaskResponse, error) {
	var reqBody, resBody GenerateHostProfileTaskList_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateKeyBody struct {
	Req    *types.GenerateKey         `xml:"urn:vim25 GenerateKey,omitempty"`
	Res    *types.GenerateKeyResponse `xml:"GenerateKeyResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateKeyBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateKey(ctx context.Context, r soap.RoundTripper, req *types.GenerateKey) (*types.GenerateKeyResponse, error) {
	var reqBody, resBody GenerateKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateLogBundles_TaskBody struct {
	Req    *types.GenerateLogBundles_Task         `xml:"urn:vim25 GenerateLogBundles_Task,omitempty"`
	Res    *types.GenerateLogBundles_TaskResponse `xml:"GenerateLogBundles_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateLogBundles_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateLogBundles_Task(ctx context.Context, r soap.RoundTripper, req *types.GenerateLogBundles_Task) (*types.GenerateLogBundles_TaskResponse, error) {
	var reqBody, resBody GenerateLogBundles_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GenerateSelfSignedClientCertBody struct {
	Req    *types.GenerateSelfSignedClientCert         `xml:"urn:vim25 GenerateSelfSignedClientCert,omitempty"`
	Res    *types.GenerateSelfSignedClientCertResponse `xml:"GenerateSelfSignedClientCertResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateSelfSignedClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func GenerateSelfSignedClientCert(ctx context.Context, r soap.RoundTripper, req *types.GenerateSelfSignedClientCert) (*types.GenerateSelfSignedClientCertResponse, error) {
	var reqBody, resBody GenerateSelfSignedClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetAlarmBody struct {
	Req    *types.GetAlarm         `xml:"urn:vim25 GetAlarm,omitempty"`
	Res    *types.GetAlarmResponse `xml:"GetAlarmResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func GetAlarm(ctx context.Context, r soap.RoundTripper, req *types.GetAlarm) (*types.GetAlarmResponse, error) {
	var reqBody, resBody GetAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetAlarmStateBody struct {
	Req    *types.GetAlarmState         `xml:"urn:vim25 GetAlarmState,omitempty"`
	Res    *types.GetAlarmStateResponse `xml:"GetAlarmStateResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetAlarmStateBody) Fault() *soap.Fault { return b.Fault_ }

func GetAlarmState(ctx context.Context, r soap.RoundTripper, req *types.GetAlarmState) (*types.GetAlarmStateResponse, error) {
	var reqBody, resBody GetAlarmStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetCryptoKeyStatusBody struct {
	Req    *types.GetCryptoKeyStatus         `xml:"urn:vim25 GetCryptoKeyStatus,omitempty"`
	Res    *types.GetCryptoKeyStatusResponse `xml:"GetCryptoKeyStatusResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetCryptoKeyStatusBody) Fault() *soap.Fault { return b.Fault_ }

func GetCryptoKeyStatus(ctx context.Context, r soap.RoundTripper, req *types.GetCryptoKeyStatus) (*types.GetCryptoKeyStatusResponse, error) {
	var reqBody, resBody GetCryptoKeyStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetCustomizationSpecBody struct {
	Req    *types.GetCustomizationSpec         `xml:"urn:vim25 GetCustomizationSpec,omitempty"`
	Res    *types.GetCustomizationSpecResponse `xml:"GetCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func GetCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.GetCustomizationSpec) (*types.GetCustomizationSpecResponse, error) {
	var reqBody, resBody GetCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetDefaultKmsClusterBody struct {
	Req    *types.GetDefaultKmsCluster         `xml:"urn:vim25 GetDefaultKmsCluster,omitempty"`
	Res    *types.GetDefaultKmsClusterResponse `xml:"GetDefaultKmsClusterResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetDefaultKmsClusterBody) Fault() *soap.Fault { return b.Fault_ }

func GetDefaultKmsCluster(ctx context.Context, r soap.RoundTripper, req *types.GetDefaultKmsCluster) (*types.GetDefaultKmsClusterResponse, error) {
	var reqBody, resBody GetDefaultKmsClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetPublicKeyBody struct {
	Req    *types.GetPublicKey         `xml:"urn:vim25 GetPublicKey,omitempty"`
	Res    *types.GetPublicKeyResponse `xml:"GetPublicKeyResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetPublicKeyBody) Fault() *soap.Fault { return b.Fault_ }

func GetPublicKey(ctx context.Context, r soap.RoundTripper, req *types.GetPublicKey) (*types.GetPublicKeyResponse, error) {
	var reqBody, resBody GetPublicKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetResourceUsageBody struct {
	Req    *types.GetResourceUsage         `xml:"urn:vim25 GetResourceUsage,omitempty"`
	Res    *types.GetResourceUsageResponse `xml:"GetResourceUsageResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetResourceUsageBody) Fault() *soap.Fault { return b.Fault_ }

func GetResourceUsage(ctx context.Context, r soap.RoundTripper, req *types.GetResourceUsage) (*types.GetResourceUsageResponse, error) {
	var reqBody, resBody GetResourceUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetSiteInfoBody struct {
	Req    *types.GetSiteInfo         `xml:"urn:vim25 GetSiteInfo,omitempty"`
	Res    *types.GetSiteInfoResponse `xml:"GetSiteInfoResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetSiteInfoBody) Fault() *soap.Fault { return b.Fault_ }

func GetSiteInfo(ctx context.Context, r soap.RoundTripper, req *types.GetSiteInfo) (*types.GetSiteInfoResponse, error) {
	var reqBody, resBody GetSiteInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetSystemVMsRestrictedDatastoresBody struct {
	Req    *types.GetSystemVMsRestrictedDatastores         `xml:"urn:vim25 GetSystemVMsRestrictedDatastores,omitempty"`
	Res    *types.GetSystemVMsRestrictedDatastoresResponse `xml:"GetSystemVMsRestrictedDatastoresResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetSystemVMsRestrictedDatastoresBody) Fault() *soap.Fault { return b.Fault_ }

func GetSystemVMsRestrictedDatastores(ctx context.Context, r soap.RoundTripper, req *types.GetSystemVMsRestrictedDatastores) (*types.GetSystemVMsRestrictedDatastoresResponse, error) {
	var reqBody, resBody GetSystemVMsRestrictedDatastoresBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVchaClusterHealthBody struct {
	Req    *types.GetVchaClusterHealth         `xml:"urn:vim25 GetVchaClusterHealth,omitempty"`
	Res    *types.GetVchaClusterHealthResponse `xml:"GetVchaClusterHealthResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVchaClusterHealthBody) Fault() *soap.Fault { return b.Fault_ }

func GetVchaClusterHealth(ctx context.Context, r soap.RoundTripper, req *types.GetVchaClusterHealth) (*types.GetVchaClusterHealthResponse, error) {
	var reqBody, resBody GetVchaClusterHealthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVsanObjExtAttrsBody struct {
	Req    *types.GetVsanObjExtAttrs         `xml:"urn:vim25 GetVsanObjExtAttrs,omitempty"`
	Res    *types.GetVsanObjExtAttrsResponse `xml:"GetVsanObjExtAttrsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVsanObjExtAttrsBody) Fault() *soap.Fault { return b.Fault_ }

func GetVsanObjExtAttrs(ctx context.Context, r soap.RoundTripper, req *types.GetVsanObjExtAttrs) (*types.GetVsanObjExtAttrsResponse, error) {
	var reqBody, resBody GetVsanObjExtAttrsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HasMonitoredEntityBody struct {
	Req    *types.HasMonitoredEntity         `xml:"urn:vim25 HasMonitoredEntity,omitempty"`
	Res    *types.HasMonitoredEntityResponse `xml:"HasMonitoredEntityResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasMonitoredEntityBody) Fault() *soap.Fault { return b.Fault_ }

func HasMonitoredEntity(ctx context.Context, r soap.RoundTripper, req *types.HasMonitoredEntity) (*types.HasMonitoredEntityResponse, error) {
	var reqBody, resBody HasMonitoredEntityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HasPrivilegeOnEntitiesBody struct {
	Req    *types.HasPrivilegeOnEntities         `xml:"urn:vim25 HasPrivilegeOnEntities,omitempty"`
	Res    *types.HasPrivilegeOnEntitiesResponse `xml:"HasPrivilegeOnEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasPrivilegeOnEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func HasPrivilegeOnEntities(ctx context.Context, r soap.RoundTripper, req *types.HasPrivilegeOnEntities) (*types.HasPrivilegeOnEntitiesResponse, error) {
	var reqBody, resBody HasPrivilegeOnEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HasPrivilegeOnEntityBody struct {
	Req    *types.HasPrivilegeOnEntity         `xml:"urn:vim25 HasPrivilegeOnEntity,omitempty"`
	Res    *types.HasPrivilegeOnEntityResponse `xml:"HasPrivilegeOnEntityResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasPrivilegeOnEntityBody) Fault() *soap.Fault { return b.Fault_ }

func HasPrivilegeOnEntity(ctx context.Context, r soap.RoundTripper, req *types.HasPrivilegeOnEntity) (*types.HasPrivilegeOnEntityResponse, error) {
	var reqBody, resBody HasPrivilegeOnEntityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HasProviderBody struct {
	Req    *types.HasProvider         `xml:"urn:vim25 HasProvider,omitempty"`
	Res    *types.HasProviderResponse `xml:"HasProviderResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasProviderBody) Fault() *soap.Fault { return b.Fault_ }

func HasProvider(ctx context.Context, r soap.RoundTripper, req *types.HasProvider) (*types.HasProviderResponse, error) {
	var reqBody, resBody HasProviderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HasUserPrivilegeOnEntitiesBody struct {
	Req    *types.HasUserPrivilegeOnEntities         `xml:"urn:vim25 HasUserPrivilegeOnEntities,omitempty"`
	Res    *types.HasUserPrivilegeOnEntitiesResponse `xml:"HasUserPrivilegeOnEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasUserPrivilegeOnEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func HasUserPrivilegeOnEntities(ctx context.Context, r soap.RoundTripper, req *types.HasUserPrivilegeOnEntities) (*types.HasUserPrivilegeOnEntitiesResponse, error) {
	var reqBody, resBody HasUserPrivilegeOnEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostClearVStorageObjectControlFlagsBody struct {
	Req    *types.HostClearVStorageObjectControlFlags         `xml:"urn:vim25 HostClearVStorageObjectControlFlags,omitempty"`
	Res    *types.HostClearVStorageObjectControlFlagsResponse `xml:"HostClearVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostClearVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func HostClearVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.HostClearVStorageObjectControlFlags) (*types.HostClearVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody HostClearVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostCloneVStorageObject_TaskBody struct {
	Req    *types.HostCloneVStorageObject_Task         `xml:"urn:vim25 HostCloneVStorageObject_Task,omitempty"`
	Res    *types.HostCloneVStorageObject_TaskResponse `xml:"HostCloneVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostCloneVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostCloneVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.HostCloneVStorageObject_Task) (*types.HostCloneVStorageObject_TaskResponse, error) {
	var reqBody, resBody HostCloneVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostConfigVFlashCacheBody struct {
	Req    *types.HostConfigVFlashCache         `xml:"urn:vim25 HostConfigVFlashCache,omitempty"`
	Res    *types.HostConfigVFlashCacheResponse `xml:"HostConfigVFlashCacheResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostConfigVFlashCacheBody) Fault() *soap.Fault { return b.Fault_ }

func HostConfigVFlashCache(ctx context.Context, r soap.RoundTripper, req *types.HostConfigVFlashCache) (*types.HostConfigVFlashCacheResponse, error) {
	var reqBody, resBody HostConfigVFlashCacheBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostConfigureVFlashResourceBody struct {
	Req    *types.HostConfigureVFlashResource         `xml:"urn:vim25 HostConfigureVFlashResource,omitempty"`
	Res    *types.HostConfigureVFlashResourceResponse `xml:"HostConfigureVFlashResourceResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostConfigureVFlashResourceBody) Fault() *soap.Fault { return b.Fault_ }

func HostConfigureVFlashResource(ctx context.Context, r soap.RoundTripper, req *types.HostConfigureVFlashResource) (*types.HostConfigureVFlashResourceResponse, error) {
	var reqBody, resBody HostConfigureVFlashResourceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostCreateDisk_TaskBody struct {
	Req    *types.HostCreateDisk_Task         `xml:"urn:vim25 HostCreateDisk_Task,omitempty"`
	Res    *types.HostCreateDisk_TaskResponse `xml:"HostCreateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostCreateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostCreateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.HostCreateDisk_Task) (*types.HostCreateDisk_TaskResponse, error) {
	var reqBody, resBody HostCreateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostDeleteVStorageObjectEx_TaskBody struct {
	Req    *types.HostDeleteVStorageObjectEx_Task         `xml:"urn:vim25 HostDeleteVStorageObjectEx_Task,omitempty"`
	Res    *types.HostDeleteVStorageObjectEx_TaskResponse `xml:"HostDeleteVStorageObjectEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostDeleteVStorageObjectEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostDeleteVStorageObjectEx_Task(ctx context.Context, r soap.RoundTripper, req *types.HostDeleteVStorageObjectEx_Task) (*types.HostDeleteVStorageObjectEx_TaskResponse, error) {
	var reqBody, resBody HostDeleteVStorageObjectEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostDeleteVStorageObject_TaskBody struct {
	Req    *types.HostDeleteVStorageObject_Task         `xml:"urn:vim25 HostDeleteVStorageObject_Task,omitempty"`
	Res    *types.HostDeleteVStorageObject_TaskResponse `xml:"HostDeleteVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostDeleteVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostDeleteVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.HostDeleteVStorageObject_Task) (*types.HostDeleteVStorageObject_TaskResponse, error) {
	var reqBody, resBody HostDeleteVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostExtendDisk_TaskBody struct {
	Req    *types.HostExtendDisk_Task         `xml:"urn:vim25 HostExtendDisk_Task,omitempty"`
	Res    *types.HostExtendDisk_TaskResponse `xml:"HostExtendDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostExtendDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostExtendDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.HostExtendDisk_Task) (*types.HostExtendDisk_TaskResponse, error) {
	var reqBody, resBody HostExtendDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostGetVFlashModuleDefaultConfigBody struct {
	Req    *types.HostGetVFlashModuleDefaultConfig         `xml:"urn:vim25 HostGetVFlashModuleDefaultConfig,omitempty"`
	Res    *types.HostGetVFlashModuleDefaultConfigResponse `xml:"HostGetVFlashModuleDefaultConfigResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostGetVFlashModuleDefaultConfigBody) Fault() *soap.Fault { return b.Fault_ }

func HostGetVFlashModuleDefaultConfig(ctx context.Context, r soap.RoundTripper, req *types.HostGetVFlashModuleDefaultConfig) (*types.HostGetVFlashModuleDefaultConfigResponse, error) {
	var reqBody, resBody HostGetVFlashModuleDefaultConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostImageConfigGetAcceptanceBody struct {
	Req    *types.HostImageConfigGetAcceptance         `xml:"urn:vim25 HostImageConfigGetAcceptance,omitempty"`
	Res    *types.HostImageConfigGetAcceptanceResponse `xml:"HostImageConfigGetAcceptanceResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostImageConfigGetAcceptanceBody) Fault() *soap.Fault { return b.Fault_ }

func HostImageConfigGetAcceptance(ctx context.Context, r soap.RoundTripper, req *types.HostImageConfigGetAcceptance) (*types.HostImageConfigGetAcceptanceResponse, error) {
	var reqBody, resBody HostImageConfigGetAcceptanceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostImageConfigGetProfileBody struct {
	Req    *types.HostImageConfigGetProfile         `xml:"urn:vim25 HostImageConfigGetProfile,omitempty"`
	Res    *types.HostImageConfigGetProfileResponse `xml:"HostImageConfigGetProfileResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostImageConfigGetProfileBody) Fault() *soap.Fault { return b.Fault_ }

func HostImageConfigGetProfile(ctx context.Context, r soap.RoundTripper, req *types.HostImageConfigGetProfile) (*types.HostImageConfigGetProfileResponse, error) {
	var reqBody, resBody HostImageConfigGetProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostInflateDisk_TaskBody struct {
	Req    *types.HostInflateDisk_Task         `xml:"urn:vim25 HostInflateDisk_Task,omitempty"`
	Res    *types.HostInflateDisk_TaskResponse `xml:"HostInflateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostInflateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostInflateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.HostInflateDisk_Task) (*types.HostInflateDisk_TaskResponse, error) {
	var reqBody, resBody HostInflateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostListVStorageObjectBody struct {
	Req    *types.HostListVStorageObject         `xml:"urn:vim25 HostListVStorageObject,omitempty"`
	Res    *types.HostListVStorageObjectResponse `xml:"HostListVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostListVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func HostListVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.HostListVStorageObject) (*types.HostListVStorageObjectResponse, error) {
	var reqBody, resBody HostListVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostProfileResetValidationStateBody struct {
	Req    *types.HostProfileResetValidationState         `xml:"urn:vim25 HostProfileResetValidationState,omitempty"`
	Res    *types.HostProfileResetValidationStateResponse `xml:"HostProfileResetValidationStateResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostProfileResetValidationStateBody) Fault() *soap.Fault { return b.Fault_ }

func HostProfileResetValidationState(ctx context.Context, r soap.RoundTripper, req *types.HostProfileResetValidationState) (*types.HostProfileResetValidationStateResponse, error) {
	var reqBody, resBody HostProfileResetValidationStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostQueryVirtualDiskUuidBody struct {
	Req    *types.HostQueryVirtualDiskUuid         `xml:"urn:vim25 HostQueryVirtualDiskUuid,omitempty"`
	Res    *types.HostQueryVirtualDiskUuidResponse `xml:"HostQueryVirtualDiskUuidResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostQueryVirtualDiskUuidBody) Fault() *soap.Fault { return b.Fault_ }

func HostQueryVirtualDiskUuid(ctx context.Context, r soap.RoundTripper, req *types.HostQueryVirtualDiskUuid) (*types.HostQueryVirtualDiskUuidResponse, error) {
	var reqBody, resBody HostQueryVirtualDiskUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostReconcileDatastoreInventory_TaskBody struct {
	Req    *types.HostReconcileDatastoreInventory_Task         `xml:"urn:vim25 HostReconcileDatastoreInventory_Task,omitempty"`
	Res    *types.HostReconcileDatastoreInventory_TaskResponse `xml:"HostReconcileDatastoreInventory_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostReconcileDatastoreInventory_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostReconcileDatastoreInventory_Task(ctx context.Context, r soap.RoundTripper, req *types.HostReconcileDatastoreInventory_Task) (*types.HostReconcileDatastoreInventory_TaskResponse, error) {
	var reqBody, resBody HostReconcileDatastoreInventory_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRegisterDiskBody struct {
	Req    *types.HostRegisterDisk         `xml:"urn:vim25 HostRegisterDisk,omitempty"`
	Res    *types.HostRegisterDiskResponse `xml:"HostRegisterDiskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRegisterDiskBody) Fault() *soap.Fault { return b.Fault_ }

func HostRegisterDisk(ctx context.Context, r soap.RoundTripper, req *types.HostRegisterDisk) (*types.HostRegisterDiskResponse, error) {
	var reqBody, resBody HostRegisterDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRelocateVStorageObject_TaskBody struct {
	Req    *types.HostRelocateVStorageObject_Task         `xml:"urn:vim25 HostRelocateVStorageObject_Task,omitempty"`
	Res    *types.HostRelocateVStorageObject_TaskResponse `xml:"HostRelocateVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRelocateVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostRelocateVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.HostRelocateVStorageObject_Task) (*types.HostRelocateVStorageObject_TaskResponse, error) {
	var reqBody, resBody HostRelocateVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRemoveVFlashResourceBody struct {
	Req    *types.HostRemoveVFlashResource         `xml:"urn:vim25 HostRemoveVFlashResource,omitempty"`
	Res    *types.HostRemoveVFlashResourceResponse `xml:"HostRemoveVFlashResourceResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRemoveVFlashResourceBody) Fault() *soap.Fault { return b.Fault_ }

func HostRemoveVFlashResource(ctx context.Context, r soap.RoundTripper, req *types.HostRemoveVFlashResource) (*types.HostRemoveVFlashResourceResponse, error) {
	var reqBody, resBody HostRemoveVFlashResourceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRenameVStorageObjectBody struct {
	Req    *types.HostRenameVStorageObject         `xml:"urn:vim25 HostRenameVStorageObject,omitempty"`
	Res    *types.HostRenameVStorageObjectResponse `xml:"HostRenameVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRenameVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func HostRenameVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.HostRenameVStorageObject) (*types.HostRenameVStorageObjectResponse, error) {
	var reqBody, resBody HostRenameVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRetrieveVStorageInfrastructureObjectPolicyBody struct {
	Req    *types.HostRetrieveVStorageInfrastructureObjectPolicy         `xml:"urn:vim25 HostRetrieveVStorageInfrastructureObjectPolicy,omitempty"`
	Res    *types.HostRetrieveVStorageInfrastructureObjectPolicyResponse `xml:"HostRetrieveVStorageInfrastructureObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRetrieveVStorageInfrastructureObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func HostRetrieveVStorageInfrastructureObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.HostRetrieveVStorageInfrastructureObjectPolicy) (*types.HostRetrieveVStorageInfrastructureObjectPolicyResponse, error) {
	var reqBody, resBody HostRetrieveVStorageInfrastructureObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRetrieveVStorageObjectBody struct {
	Req    *types.HostRetrieveVStorageObject         `xml:"urn:vim25 HostRetrieveVStorageObject,omitempty"`
	Res    *types.HostRetrieveVStorageObjectResponse `xml:"HostRetrieveVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRetrieveVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func HostRetrieveVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.HostRetrieveVStorageObject) (*types.HostRetrieveVStorageObjectResponse, error) {
	var reqBody, resBody HostRetrieveVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRetrieveVStorageObjectMetadataBody struct {
	Req    *types.HostRetrieveVStorageObjectMetadata         `xml:"urn:vim25 HostRetrieveVStorageObjectMetadata,omitempty"`
	Res    *types.HostRetrieveVStorageObjectMetadataResponse `xml:"HostRetrieveVStorageObjectMetadataResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRetrieveVStorageObjectMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func HostRetrieveVStorageObjectMetadata(ctx context.Context, r soap.RoundTripper, req *types.HostRetrieveVStorageObjectMetadata) (*types.HostRetrieveVStorageObjectMetadataResponse, error) {
	var reqBody, resBody HostRetrieveVStorageObjectMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRetrieveVStorageObjectMetadataValueBody struct {
	Req    *types.HostRetrieveVStorageObjectMetadataValue         `xml:"urn:vim25 HostRetrieveVStorageObjectMetadataValue,omitempty"`
	Res    *types.HostRetrieveVStorageObjectMetadataValueResponse `xml:"HostRetrieveVStorageObjectMetadataValueResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRetrieveVStorageObjectMetadataValueBody) Fault() *soap.Fault { return b.Fault_ }

func HostRetrieveVStorageObjectMetadataValue(ctx context.Context, r soap.RoundTripper, req *types.HostRetrieveVStorageObjectMetadataValue) (*types.HostRetrieveVStorageObjectMetadataValueResponse, error) {
	var reqBody, resBody HostRetrieveVStorageObjectMetadataValueBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostRetrieveVStorageObjectStateBody struct {
	Req    *types.HostRetrieveVStorageObjectState         `xml:"urn:vim25 HostRetrieveVStorageObjectState,omitempty"`
	Res    *types.HostRetrieveVStorageObjectStateResponse `xml:"HostRetrieveVStorageObjectStateResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRetrieveVStorageObjectStateBody) Fault() *soap.Fault { return b.Fault_ }

func HostRetrieveVStorageObjectState(ctx context.Context, r soap.RoundTripper, req *types.HostRetrieveVStorageObjectState) (*types.HostRetrieveVStorageObjectStateResponse, error) {
	var reqBody, resBody HostRetrieveVStorageObjectStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostScheduleReconcileDatastoreInventoryBody struct {
	Req    *types.HostScheduleReconcileDatastoreInventory         `xml:"urn:vim25 HostScheduleReconcileDatastoreInventory,omitempty"`
	Res    *types.HostScheduleReconcileDatastoreInventoryResponse `xml:"HostScheduleReconcileDatastoreInventoryResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostScheduleReconcileDatastoreInventoryBody) Fault() *soap.Fault { return b.Fault_ }

func HostScheduleReconcileDatastoreInventory(ctx context.Context, r soap.RoundTripper, req *types.HostScheduleReconcileDatastoreInventory) (*types.HostScheduleReconcileDatastoreInventoryResponse, error) {
	var reqBody, resBody HostScheduleReconcileDatastoreInventoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSetVStorageObjectControlFlagsBody struct {
	Req    *types.HostSetVStorageObjectControlFlags         `xml:"urn:vim25 HostSetVStorageObjectControlFlags,omitempty"`
	Res    *types.HostSetVStorageObjectControlFlagsResponse `xml:"HostSetVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSetVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func HostSetVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.HostSetVStorageObjectControlFlags) (*types.HostSetVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody HostSetVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSetVirtualDiskUuid_TaskBody struct {
	Req    *types.HostSetVirtualDiskUuid_Task         `xml:"urn:vim25 HostSetVirtualDiskUuid_Task,omitempty"`
	Res    *types.HostSetVirtualDiskUuid_TaskResponse `xml:"HostSetVirtualDiskUuid_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSetVirtualDiskUuid_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostSetVirtualDiskUuid_Task(ctx context.Context, r soap.RoundTripper, req *types.HostSetVirtualDiskUuid_Task) (*types.HostSetVirtualDiskUuid_TaskResponse, error) {
	var reqBody, resBody HostSetVirtualDiskUuid_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostSpecGetUpdatedHostsBody struct {
	Req    *types.HostSpecGetUpdatedHosts         `xml:"urn:vim25 HostSpecGetUpdatedHosts,omitempty"`
	Res    *types.HostSpecGetUpdatedHostsResponse `xml:"HostSpecGetUpdatedHostsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostSpecGetUpdatedHostsBody) Fault() *soap.Fault { return b.Fault_ }

func HostSpecGetUpdatedHosts(ctx context.Context, r soap.RoundTripper, req *types.HostSpecGetUpdatedHosts) (*types.HostSpecGetUpdatedHostsResponse, error) {
	var reqBody, resBody HostSpecGetUpdatedHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostUpdateVStorageObjectMetadataEx_TaskBody struct {
	Req    *types.HostUpdateVStorageObjectMetadataEx_Task         `xml:"urn:vim25 HostUpdateVStorageObjectMetadataEx_Task,omitempty"`
	Res    *types.HostUpdateVStorageObjectMetadataEx_TaskResponse `xml:"HostUpdateVStorageObjectMetadataEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostUpdateVStorageObjectMetadataEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostUpdateVStorageObjectMetadataEx_Task(ctx context.Context, r soap.RoundTripper, req *types.HostUpdateVStorageObjectMetadataEx_Task) (*types.HostUpdateVStorageObjectMetadataEx_TaskResponse, error) {
	var reqBody, resBody HostUpdateVStorageObjectMetadataEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostUpdateVStorageObjectMetadata_TaskBody struct {
	Req    *types.HostUpdateVStorageObjectMetadata_Task         `xml:"urn:vim25 HostUpdateVStorageObjectMetadata_Task,omitempty"`
	Res    *types.HostUpdateVStorageObjectMetadata_TaskResponse `xml:"HostUpdateVStorageObjectMetadata_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostUpdateVStorageObjectMetadata_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostUpdateVStorageObjectMetadata_Task(ctx context.Context, r soap.RoundTripper, req *types.HostUpdateVStorageObjectMetadata_Task) (*types.HostUpdateVStorageObjectMetadata_TaskResponse, error) {
	var reqBody, resBody HostUpdateVStorageObjectMetadata_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostVStorageObjectCreateDiskFromSnapshot_TaskBody struct {
	Req    *types.HostVStorageObjectCreateDiskFromSnapshot_Task         `xml:"urn:vim25 HostVStorageObjectCreateDiskFromSnapshot_Task,omitempty"`
	Res    *types.HostVStorageObjectCreateDiskFromSnapshot_TaskResponse `xml:"HostVStorageObjectCreateDiskFromSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostVStorageObjectCreateDiskFromSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostVStorageObjectCreateDiskFromSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.HostVStorageObjectCreateDiskFromSnapshot_Task) (*types.HostVStorageObjectCreateDiskFromSnapshot_TaskResponse, error) {
	var reqBody, resBody HostVStorageObjectCreateDiskFromSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostVStorageObjectCreateSnapshot_TaskBody struct {
	Req    *types.HostVStorageObjectCreateSnapshot_Task         `xml:"urn:vim25 HostVStorageObjectCreateSnapshot_Task,omitempty"`
	Res    *types.HostVStorageObjectCreateSnapshot_TaskResponse `xml:"HostVStorageObjectCreateSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostVStorageObjectCreateSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostVStorageObjectCreateSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.HostVStorageObjectCreateSnapshot_Task) (*types.HostVStorageObjectCreateSnapshot_TaskResponse, error) {
	var reqBody, resBody HostVStorageObjectCreateSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostVStorageObjectDeleteSnapshot_TaskBody struct {
	Req    *types.HostVStorageObjectDeleteSnapshot_Task         `xml:"urn:vim25 HostVStorageObjectDeleteSnapshot_Task,omitempty"`
	Res    *types.HostVStorageObjectDeleteSnapshot_TaskResponse `xml:"HostVStorageObjectDeleteSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostVStorageObjectDeleteSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostVStorageObjectDeleteSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.HostVStorageObjectDeleteSnapshot_Task) (*types.HostVStorageObjectDeleteSnapshot_TaskResponse, error) {
	var reqBody, resBody HostVStorageObjectDeleteSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostVStorageObjectRetrieveSnapshotInfoBody struct {
	Req    *types.HostVStorageObjectRetrieveSnapshotInfo         `xml:"urn:vim25 HostVStorageObjectRetrieveSnapshotInfo,omitempty"`
	Res    *types.HostVStorageObjectRetrieveSnapshotInfoResponse `xml:"HostVStorageObjectRetrieveSnapshotInfoResponse,omitempty"`
	Fault_ *soap.Fault                                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostVStorageObjectRetrieveSnapshotInfoBody) Fault() *soap.Fault { return b.Fault_ }

func HostVStorageObjectRetrieveSnapshotInfo(ctx context.Context, r soap.RoundTripper, req *types.HostVStorageObjectRetrieveSnapshotInfo) (*types.HostVStorageObjectRetrieveSnapshotInfoResponse, error) {
	var reqBody, resBody HostVStorageObjectRetrieveSnapshotInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HostVStorageObjectRevert_TaskBody struct {
	Req    *types.HostVStorageObjectRevert_Task         `xml:"urn:vim25 HostVStorageObjectRevert_Task,omitempty"`
	Res    *types.HostVStorageObjectRevert_TaskResponse `xml:"HostVStorageObjectRevert_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostVStorageObjectRevert_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HostVStorageObjectRevert_Task(ctx context.Context, r soap.RoundTripper, req *types.HostVStorageObjectRevert_Task) (*types.HostVStorageObjectRevert_TaskResponse, error) {
	var reqBody, resBody HostVStorageObjectRevert_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseAbortBody struct {
	Req    *types.HttpNfcLeaseAbort         `xml:"urn:vim25 HttpNfcLeaseAbort,omitempty"`
	Res    *types.HttpNfcLeaseAbortResponse `xml:"HttpNfcLeaseAbortResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseAbortBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseAbort(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseAbort) (*types.HttpNfcLeaseAbortResponse, error) {
	var reqBody, resBody HttpNfcLeaseAbortBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseCompleteBody struct {
	Req    *types.HttpNfcLeaseComplete         `xml:"urn:vim25 HttpNfcLeaseComplete,omitempty"`
	Res    *types.HttpNfcLeaseCompleteResponse `xml:"HttpNfcLeaseCompleteResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseCompleteBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseComplete(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseComplete) (*types.HttpNfcLeaseCompleteResponse, error) {
	var reqBody, resBody HttpNfcLeaseCompleteBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseGetManifestBody struct {
	Req    *types.HttpNfcLeaseGetManifest         `xml:"urn:vim25 HttpNfcLeaseGetManifest,omitempty"`
	Res    *types.HttpNfcLeaseGetManifestResponse `xml:"HttpNfcLeaseGetManifestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseGetManifestBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseGetManifest(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseGetManifest) (*types.HttpNfcLeaseGetManifestResponse, error) {
	var reqBody, resBody HttpNfcLeaseGetManifestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseProbeUrlsBody struct {
	Req    *types.HttpNfcLeaseProbeUrls         `xml:"urn:vim25 HttpNfcLeaseProbeUrls,omitempty"`
	Res    *types.HttpNfcLeaseProbeUrlsResponse `xml:"HttpNfcLeaseProbeUrlsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseProbeUrlsBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseProbeUrls(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseProbeUrls) (*types.HttpNfcLeaseProbeUrlsResponse, error) {
	var reqBody, resBody HttpNfcLeaseProbeUrlsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseProgressBody struct {
	Req    *types.HttpNfcLeaseProgress         `xml:"urn:vim25 HttpNfcLeaseProgress,omitempty"`
	Res    *types.HttpNfcLeaseProgressResponse `xml:"HttpNfcLeaseProgressResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseProgressBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseProgress(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseProgress) (*types.HttpNfcLeaseProgressResponse, error) {
	var reqBody, resBody HttpNfcLeaseProgressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeasePullFromUrls_TaskBody struct {
	Req    *types.HttpNfcLeasePullFromUrls_Task         `xml:"urn:vim25 HttpNfcLeasePullFromUrls_Task,omitempty"`
	Res    *types.HttpNfcLeasePullFromUrls_TaskResponse `xml:"HttpNfcLeasePullFromUrls_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeasePullFromUrls_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeasePullFromUrls_Task(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeasePullFromUrls_Task) (*types.HttpNfcLeasePullFromUrls_TaskResponse, error) {
	var reqBody, resBody HttpNfcLeasePullFromUrls_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type HttpNfcLeaseSetManifestChecksumTypeBody struct {
	Req    *types.HttpNfcLeaseSetManifestChecksumType         `xml:"urn:vim25 HttpNfcLeaseSetManifestChecksumType,omitempty"`
	Res    *types.HttpNfcLeaseSetManifestChecksumTypeResponse `xml:"HttpNfcLeaseSetManifestChecksumTypeResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseSetManifestChecksumTypeBody) Fault() *soap.Fault { return b.Fault_ }

func HttpNfcLeaseSetManifestChecksumType(ctx context.Context, r soap.RoundTripper, req *types.HttpNfcLeaseSetManifestChecksumType) (*types.HttpNfcLeaseSetManifestChecksumTypeResponse, error) {
	var reqBody, resBody HttpNfcLeaseSetManifestChecksumTypeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ImpersonateUserBody struct {
	Req    *types.ImpersonateUser         `xml:"urn:vim25 ImpersonateUser,omitempty"`
	Res    *types.ImpersonateUserResponse `xml:"ImpersonateUserResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImpersonateUserBody) Fault() *soap.Fault { return b.Fault_ }

func ImpersonateUser(ctx context.Context, r soap.RoundTripper, req *types.ImpersonateUser) (*types.ImpersonateUserResponse, error) {
	var reqBody, resBody ImpersonateUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ImportCertificateForCAM_TaskBody struct {
	Req    *types.ImportCertificateForCAM_Task         `xml:"urn:vim25 ImportCertificateForCAM_Task,omitempty"`
	Res    *types.ImportCertificateForCAM_TaskResponse `xml:"ImportCertificateForCAM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImportCertificateForCAM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ImportCertificateForCAM_Task(ctx context.Context, r soap.RoundTripper, req *types.ImportCertificateForCAM_Task) (*types.ImportCertificateForCAM_TaskResponse, error) {
	var reqBody, resBody ImportCertificateForCAM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ImportUnmanagedSnapshotBody struct {
	Req    *types.ImportUnmanagedSnapshot         `xml:"urn:vim25 ImportUnmanagedSnapshot,omitempty"`
	Res    *types.ImportUnmanagedSnapshotResponse `xml:"ImportUnmanagedSnapshotResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImportUnmanagedSnapshotBody) Fault() *soap.Fault { return b.Fault_ }

func ImportUnmanagedSnapshot(ctx context.Context, r soap.RoundTripper, req *types.ImportUnmanagedSnapshot) (*types.ImportUnmanagedSnapshotResponse, error) {
	var reqBody, resBody ImportUnmanagedSnapshotBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ImportVAppBody struct {
	Req    *types.ImportVApp         `xml:"urn:vim25 ImportVApp,omitempty"`
	Res    *types.ImportVAppResponse `xml:"ImportVAppResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImportVAppBody) Fault() *soap.Fault { return b.Fault_ }

func ImportVApp(ctx context.Context, r soap.RoundTripper, req *types.ImportVApp) (*types.ImportVAppResponse, error) {
	var reqBody, resBody ImportVAppBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type IncreaseDirectorySizeBody struct {
	Req    *types.IncreaseDirectorySize         `xml:"urn:vim25 IncreaseDirectorySize,omitempty"`
	Res    *types.IncreaseDirectorySizeResponse `xml:"IncreaseDirectorySizeResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IncreaseDirectorySizeBody) Fault() *soap.Fault { return b.Fault_ }

func IncreaseDirectorySize(ctx context.Context, r soap.RoundTripper, req *types.IncreaseDirectorySize) (*types.IncreaseDirectorySizeResponse, error) {
	var reqBody, resBody IncreaseDirectorySizeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InflateDisk_TaskBody struct {
	Req    *types.InflateDisk_Task         `xml:"urn:vim25 InflateDisk_Task,omitempty"`
	Res    *types.InflateDisk_TaskResponse `xml:"InflateDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InflateDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InflateDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.InflateDisk_Task) (*types.InflateDisk_TaskResponse, error) {
	var reqBody, resBody InflateDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InflateVirtualDisk_TaskBody struct {
	Req    *types.InflateVirtualDisk_Task         `xml:"urn:vim25 InflateVirtualDisk_Task,omitempty"`
	Res    *types.InflateVirtualDisk_TaskResponse `xml:"InflateVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InflateVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InflateVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.InflateVirtualDisk_Task) (*types.InflateVirtualDisk_TaskResponse, error) {
	var reqBody, resBody InflateVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitializeDisks_TaskBody struct {
	Req    *types.InitializeDisks_Task         `xml:"urn:vim25 InitializeDisks_Task,omitempty"`
	Res    *types.InitializeDisks_TaskResponse `xml:"InitializeDisks_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitializeDisks_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InitializeDisks_Task(ctx context.Context, r soap.RoundTripper, req *types.InitializeDisks_Task) (*types.InitializeDisks_TaskResponse, error) {
	var reqBody, resBody InitializeDisks_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitiateFileTransferFromGuestBody struct {
	Req    *types.InitiateFileTransferFromGuest         `xml:"urn:vim25 InitiateFileTransferFromGuest,omitempty"`
	Res    *types.InitiateFileTransferFromGuestResponse `xml:"InitiateFileTransferFromGuestResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateFileTransferFromGuestBody) Fault() *soap.Fault { return b.Fault_ }

func InitiateFileTransferFromGuest(ctx context.Context, r soap.RoundTripper, req *types.InitiateFileTransferFromGuest) (*types.InitiateFileTransferFromGuestResponse, error) {
	var reqBody, resBody InitiateFileTransferFromGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitiateFileTransferToGuestBody struct {
	Req    *types.InitiateFileTransferToGuest         `xml:"urn:vim25 InitiateFileTransferToGuest,omitempty"`
	Res    *types.InitiateFileTransferToGuestResponse `xml:"InitiateFileTransferToGuestResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateFileTransferToGuestBody) Fault() *soap.Fault { return b.Fault_ }

func InitiateFileTransferToGuest(ctx context.Context, r soap.RoundTripper, req *types.InitiateFileTransferToGuest) (*types.InitiateFileTransferToGuestResponse, error) {
	var reqBody, resBody InitiateFileTransferToGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitiateTransitionToVLCM_TaskBody struct {
	Req    *types.InitiateTransitionToVLCM_Task         `xml:"urn:vim25 InitiateTransitionToVLCM_Task,omitempty"`
	Res    *types.InitiateTransitionToVLCM_TaskResponse `xml:"InitiateTransitionToVLCM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateTransitionToVLCM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InitiateTransitionToVLCM_Task(ctx context.Context, r soap.RoundTripper, req *types.InitiateTransitionToVLCM_Task) (*types.InitiateTransitionToVLCM_TaskResponse, error) {
	var reqBody, resBody InitiateTransitionToVLCM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallHostPatchV2_TaskBody struct {
	Req    *types.InstallHostPatchV2_Task         `xml:"urn:vim25 InstallHostPatchV2_Task,omitempty"`
	Res    *types.InstallHostPatchV2_TaskResponse `xml:"InstallHostPatchV2_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallHostPatchV2_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InstallHostPatchV2_Task(ctx context.Context, r soap.RoundTripper, req *types.InstallHostPatchV2_Task) (*types.InstallHostPatchV2_TaskResponse, error) {
	var reqBody, resBody InstallHostPatchV2_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallHostPatch_TaskBody struct {
	Req    *types.InstallHostPatch_Task         `xml:"urn:vim25 InstallHostPatch_Task,omitempty"`
	Res    *types.InstallHostPatch_TaskResponse `xml:"InstallHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InstallHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.InstallHostPatch_Task) (*types.InstallHostPatch_TaskResponse, error) {
	var reqBody, resBody InstallHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallIoFilter_TaskBody struct {
	Req    *types.InstallIoFilter_Task         `xml:"urn:vim25 InstallIoFilter_Task,omitempty"`
	Res    *types.InstallIoFilter_TaskResponse `xml:"InstallIoFilter_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallIoFilter_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InstallIoFilter_Task(ctx context.Context, r soap.RoundTripper, req *types.InstallIoFilter_Task) (*types.InstallIoFilter_TaskResponse, error) {
	var reqBody, resBody InstallIoFilter_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallServerCertificateBody struct {
	Req    *types.InstallServerCertificate         `xml:"urn:vim25 InstallServerCertificate,omitempty"`
	Res    *types.InstallServerCertificateResponse `xml:"InstallServerCertificateResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallServerCertificateBody) Fault() *soap.Fault { return b.Fault_ }

func InstallServerCertificate(ctx context.Context, r soap.RoundTripper, req *types.InstallServerCertificate) (*types.InstallServerCertificateResponse, error) {
	var reqBody, resBody InstallServerCertificateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallSmartCardTrustAnchorBody struct {
	Req    *types.InstallSmartCardTrustAnchor         `xml:"urn:vim25 InstallSmartCardTrustAnchor,omitempty"`
	Res    *types.InstallSmartCardTrustAnchorResponse `xml:"InstallSmartCardTrustAnchorResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallSmartCardTrustAnchorBody) Fault() *soap.Fault { return b.Fault_ }

func InstallSmartCardTrustAnchor(ctx context.Context, r soap.RoundTripper, req *types.InstallSmartCardTrustAnchor) (*types.InstallSmartCardTrustAnchorResponse, error) {
	var reqBody, resBody InstallSmartCardTrustAnchorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstantClone_TaskBody struct {
	Req    *types.InstantClone_Task         `xml:"urn:vim25 InstantClone_Task,omitempty"`
	Res    *types.InstantClone_TaskResponse `xml:"InstantClone_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstantClone_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InstantClone_Task(ctx context.Context, r soap.RoundTripper, req *types.InstantClone_Task) (*types.InstantClone_TaskResponse, error) {
	var reqBody, resBody InstantClone_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type IsClusteredVmdkEnabledBody struct {
	Req    *types.IsClusteredVmdkEnabled         `xml:"urn:vim25 IsClusteredVmdkEnabled,omitempty"`
	Res    *types.IsClusteredVmdkEnabledResponse `xml:"IsClusteredVmdkEnabledResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IsClusteredVmdkEnabledBody) Fault() *soap.Fault { return b.Fault_ }

func IsClusteredVmdkEnabled(ctx context.Context, r soap.RoundTripper, req *types.IsClusteredVmdkEnabled) (*types.IsClusteredVmdkEnabledResponse, error) {
	var reqBody, resBody IsClusteredVmdkEnabledBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type IsGuestOsCustomizableBody struct {
	Req    *types.IsGuestOsCustomizable         `xml:"urn:vim25 IsGuestOsCustomizable,omitempty"`
	Res    *types.IsGuestOsCustomizableResponse `xml:"IsGuestOsCustomizableResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IsGuestOsCustomizableBody) Fault() *soap.Fault { return b.Fault_ }

func IsGuestOsCustomizable(ctx context.Context, r soap.RoundTripper, req *types.IsGuestOsCustomizable) (*types.IsGuestOsCustomizableResponse, error) {
	var reqBody, resBody IsGuestOsCustomizableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type IsKmsClusterActiveBody struct {
	Req    *types.IsKmsClusterActive         `xml:"urn:vim25 IsKmsClusterActive,omitempty"`
	Res    *types.IsKmsClusterActiveResponse `xml:"IsKmsClusterActiveResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IsKmsClusterActiveBody) Fault() *soap.Fault { return b.Fault_ }

func IsKmsClusterActive(ctx context.Context, r soap.RoundTripper, req *types.IsKmsClusterActive) (*types.IsKmsClusterActiveResponse, error) {
	var reqBody, resBody IsKmsClusterActiveBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type IsSharedGraphicsActiveBody struct {
	Req    *types.IsSharedGraphicsActive         `xml:"urn:vim25 IsSharedGraphicsActive,omitempty"`
	Res    *types.IsSharedGraphicsActiveResponse `xml:"IsSharedGraphicsActiveResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IsSharedGraphicsActiveBody) Fault() *soap.Fault { return b.Fault_ }

func IsSharedGraphicsActive(ctx context.Context, r soap.RoundTripper, req *types.IsSharedGraphicsActive) (*types.IsSharedGraphicsActiveResponse, error) {
	var reqBody, resBody IsSharedGraphicsActiveBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type JoinDomainWithCAM_TaskBody struct {
	Req    *types.JoinDomainWithCAM_Task         `xml:"urn:vim25 JoinDomainWithCAM_Task,omitempty"`
	Res    *types.JoinDomainWithCAM_TaskResponse `xml:"JoinDomainWithCAM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *JoinDomainWithCAM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func JoinDomainWithCAM_Task(ctx context.Context, r soap.RoundTripper, req *types.JoinDomainWithCAM_Task) (*types.JoinDomainWithCAM_TaskResponse, error) {
	var reqBody, resBody JoinDomainWithCAM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type JoinDomain_TaskBody struct {
	Req    *types.JoinDomain_Task         `xml:"urn:vim25 JoinDomain_Task,omitempty"`
	Res    *types.JoinDomain_TaskResponse `xml:"JoinDomain_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *JoinDomain_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func JoinDomain_Task(ctx context.Context, r soap.RoundTripper, req *types.JoinDomain_Task) (*types.JoinDomain_TaskResponse, error) {
	var reqBody, resBody JoinDomain_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LeaveCurrentDomain_TaskBody struct {
	Req    *types.LeaveCurrentDomain_Task         `xml:"urn:vim25 LeaveCurrentDomain_Task,omitempty"`
	Res    *types.LeaveCurrentDomain_TaskResponse `xml:"LeaveCurrentDomain_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LeaveCurrentDomain_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func LeaveCurrentDomain_Task(ctx context.Context, r soap.RoundTripper, req *types.LeaveCurrentDomain_Task) (*types.LeaveCurrentDomain_TaskResponse, error) {
	var reqBody, resBody LeaveCurrentDomain_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListCACertificateRevocationListsBody struct {
	Req    *types.ListCACertificateRevocationLists         `xml:"urn:vim25 ListCACertificateRevocationLists,omitempty"`
	Res    *types.ListCACertificateRevocationListsResponse `xml:"ListCACertificateRevocationListsResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListCACertificateRevocationListsBody) Fault() *soap.Fault { return b.Fault_ }

func ListCACertificateRevocationLists(ctx context.Context, r soap.RoundTripper, req *types.ListCACertificateRevocationLists) (*types.ListCACertificateRevocationListsResponse, error) {
	var reqBody, resBody ListCACertificateRevocationListsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListCACertificatesBody struct {
	Req    *types.ListCACertificates         `xml:"urn:vim25 ListCACertificates,omitempty"`
	Res    *types.ListCACertificatesResponse `xml:"ListCACertificatesResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListCACertificatesBody) Fault() *soap.Fault { return b.Fault_ }

func ListCACertificates(ctx context.Context, r soap.RoundTripper, req *types.ListCACertificates) (*types.ListCACertificatesResponse, error) {
	var reqBody, resBody ListCACertificatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListFilesInGuestBody struct {
	Req    *types.ListFilesInGuest         `xml:"urn:vim25 ListFilesInGuest,omitempty"`
	Res    *types.ListFilesInGuestResponse `xml:"ListFilesInGuestResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListFilesInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ListFilesInGuest(ctx context.Context, r soap.RoundTripper, req *types.ListFilesInGuest) (*types.ListFilesInGuestResponse, error) {
	var reqBody, resBody ListFilesInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListGuestAliasesBody struct {
	Req    *types.ListGuestAliases         `xml:"urn:vim25 ListGuestAliases,omitempty"`
	Res    *types.ListGuestAliasesResponse `xml:"ListGuestAliasesResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListGuestAliasesBody) Fault() *soap.Fault { return b.Fault_ }

func ListGuestAliases(ctx context.Context, r soap.RoundTripper, req *types.ListGuestAliases) (*types.ListGuestAliasesResponse, error) {
	var reqBody, resBody ListGuestAliasesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListGuestMappedAliasesBody struct {
	Req    *types.ListGuestMappedAliases         `xml:"urn:vim25 ListGuestMappedAliases,omitempty"`
	Res    *types.ListGuestMappedAliasesResponse `xml:"ListGuestMappedAliasesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListGuestMappedAliasesBody) Fault() *soap.Fault { return b.Fault_ }

func ListGuestMappedAliases(ctx context.Context, r soap.RoundTripper, req *types.ListGuestMappedAliases) (*types.ListGuestMappedAliasesResponse, error) {
	var reqBody, resBody ListGuestMappedAliasesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListKeysBody struct {
	Req    *types.ListKeys         `xml:"urn:vim25 ListKeys,omitempty"`
	Res    *types.ListKeysResponse `xml:"ListKeysResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListKeysBody) Fault() *soap.Fault { return b.Fault_ }

func ListKeys(ctx context.Context, r soap.RoundTripper, req *types.ListKeys) (*types.ListKeysResponse, error) {
	var reqBody, resBody ListKeysBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListKmipServersBody struct {
	Req    *types.ListKmipServers         `xml:"urn:vim25 ListKmipServers,omitempty"`
	Res    *types.ListKmipServersResponse `xml:"ListKmipServersResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListKmipServersBody) Fault() *soap.Fault { return b.Fault_ }

func ListKmipServers(ctx context.Context, r soap.RoundTripper, req *types.ListKmipServers) (*types.ListKmipServersResponse, error) {
	var reqBody, resBody ListKmipServersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListKmsClustersBody struct {
	Req    *types.ListKmsClusters         `xml:"urn:vim25 ListKmsClusters,omitempty"`
	Res    *types.ListKmsClustersResponse `xml:"ListKmsClustersResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListKmsClustersBody) Fault() *soap.Fault { return b.Fault_ }

func ListKmsClusters(ctx context.Context, r soap.RoundTripper, req *types.ListKmsClusters) (*types.ListKmsClustersResponse, error) {
	var reqBody, resBody ListKmsClustersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListProcessesInGuestBody struct {
	Req    *types.ListProcessesInGuest         `xml:"urn:vim25 ListProcessesInGuest,omitempty"`
	Res    *types.ListProcessesInGuestResponse `xml:"ListProcessesInGuestResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListProcessesInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ListProcessesInGuest(ctx context.Context, r soap.RoundTripper, req *types.ListProcessesInGuest) (*types.ListProcessesInGuestResponse, error) {
	var reqBody, resBody ListProcessesInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListRegistryKeysInGuestBody struct {
	Req    *types.ListRegistryKeysInGuest         `xml:"urn:vim25 ListRegistryKeysInGuest,omitempty"`
	Res    *types.ListRegistryKeysInGuestResponse `xml:"ListRegistryKeysInGuestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListRegistryKeysInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ListRegistryKeysInGuest(ctx context.Context, r soap.RoundTripper, req *types.ListRegistryKeysInGuest) (*types.ListRegistryKeysInGuestResponse, error) {
	var reqBody, resBody ListRegistryKeysInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListRegistryValuesInGuestBody struct {
	Req    *types.ListRegistryValuesInGuest         `xml:"urn:vim25 ListRegistryValuesInGuest,omitempty"`
	Res    *types.ListRegistryValuesInGuestResponse `xml:"ListRegistryValuesInGuestResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListRegistryValuesInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ListRegistryValuesInGuest(ctx context.Context, r soap.RoundTripper, req *types.ListRegistryValuesInGuest) (*types.ListRegistryValuesInGuestResponse, error) {
	var reqBody, resBody ListRegistryValuesInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListSmartCardTrustAnchorsBody struct {
	Req    *types.ListSmartCardTrustAnchors         `xml:"urn:vim25 ListSmartCardTrustAnchors,omitempty"`
	Res    *types.ListSmartCardTrustAnchorsResponse `xml:"ListSmartCardTrustAnchorsResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListSmartCardTrustAnchorsBody) Fault() *soap.Fault { return b.Fault_ }

func ListSmartCardTrustAnchors(ctx context.Context, r soap.RoundTripper, req *types.ListSmartCardTrustAnchors) (*types.ListSmartCardTrustAnchorsResponse, error) {
	var reqBody, resBody ListSmartCardTrustAnchorsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListTagsAttachedToVStorageObjectBody struct {
	Req    *types.ListTagsAttachedToVStorageObject         `xml:"urn:vim25 ListTagsAttachedToVStorageObject,omitempty"`
	Res    *types.ListTagsAttachedToVStorageObjectResponse `xml:"ListTagsAttachedToVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListTagsAttachedToVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func ListTagsAttachedToVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.ListTagsAttachedToVStorageObject) (*types.ListTagsAttachedToVStorageObjectResponse, error) {
	var reqBody, resBody ListTagsAttachedToVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListVStorageObjectBody struct {
	Req    *types.ListVStorageObject         `xml:"urn:vim25 ListVStorageObject,omitempty"`
	Res    *types.ListVStorageObjectResponse `xml:"ListVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func ListVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.ListVStorageObject) (*types.ListVStorageObjectResponse, error) {
	var reqBody, resBody ListVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ListVStorageObjectsAttachedToTagBody struct {
	Req    *types.ListVStorageObjectsAttachedToTag         `xml:"urn:vim25 ListVStorageObjectsAttachedToTag,omitempty"`
	Res    *types.ListVStorageObjectsAttachedToTagResponse `xml:"ListVStorageObjectsAttachedToTagResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListVStorageObjectsAttachedToTagBody) Fault() *soap.Fault { return b.Fault_ }

func ListVStorageObjectsAttachedToTag(ctx context.Context, r soap.RoundTripper, req *types.ListVStorageObjectsAttachedToTag) (*types.ListVStorageObjectsAttachedToTagResponse, error) {
	var reqBody, resBody ListVStorageObjectsAttachedToTagBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LogUserEventBody struct {
	Req    *types.LogUserEvent         `xml:"urn:vim25 LogUserEvent,omitempty"`
	Res    *types.LogUserEventResponse `xml:"LogUserEventResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LogUserEventBody) Fault() *soap.Fault { return b.Fault_ }

func LogUserEvent(ctx context.Context, r soap.RoundTripper, req *types.LogUserEvent) (*types.LogUserEventResponse, error) {
	var reqBody, resBody LogUserEventBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LoginBody struct {
	Req    *types.Login         `xml:"urn:vim25 Login,omitempty"`
	Res    *types.LoginResponse `xml:"LoginResponse,omitempty"`
	Fault_ *soap.Fault          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginBody) Fault() *soap.Fault { return b.Fault_ }

func Login(ctx context.Context, r soap.RoundTripper, req *types.Login) (*types.LoginResponse, error) {
	var reqBody, resBody LoginBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LoginBySSPIBody struct {
	Req    *types.LoginBySSPI         `xml:"urn:vim25 LoginBySSPI,omitempty"`
	Res    *types.LoginBySSPIResponse `xml:"LoginBySSPIResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginBySSPIBody) Fault() *soap.Fault { return b.Fault_ }

func LoginBySSPI(ctx context.Context, r soap.RoundTripper, req *types.LoginBySSPI) (*types.LoginBySSPIResponse, error) {
	var reqBody, resBody LoginBySSPIBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LoginByTokenBody struct {
	Req    *types.LoginByToken         `xml:"urn:vim25 LoginByToken,omitempty"`
	Res    *types.LoginByTokenResponse `xml:"LoginByTokenResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginByTokenBody) Fault() *soap.Fault { return b.Fault_ }

func LoginByToken(ctx context.Context, r soap.RoundTripper, req *types.LoginByToken) (*types.LoginByTokenResponse, error) {
	var reqBody, resBody LoginByTokenBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LoginExtensionByCertificateBody struct {
	Req    *types.LoginExtensionByCertificate         `xml:"urn:vim25 LoginExtensionByCertificate,omitempty"`
	Res    *types.LoginExtensionByCertificateResponse `xml:"LoginExtensionByCertificateResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginExtensionByCertificateBody) Fault() *soap.Fault { return b.Fault_ }

func LoginExtensionByCertificate(ctx context.Context, r soap.RoundTripper, req *types.LoginExtensionByCertificate) (*types.LoginExtensionByCertificateResponse, error) {
	var reqBody, resBody LoginExtensionByCertificateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LoginExtensionBySubjectNameBody struct {
	Req    *types.LoginExtensionBySubjectName         `xml:"urn:vim25 LoginExtensionBySubjectName,omitempty"`
	Res    *types.LoginExtensionBySubjectNameResponse `xml:"LoginExtensionBySubjectNameResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginExtensionBySubjectNameBody) Fault() *soap.Fault { return b.Fault_ }

func LoginExtensionBySubjectName(ctx context.Context, r soap.RoundTripper, req *types.LoginExtensionBySubjectName) (*types.LoginExtensionBySubjectNameResponse, error) {
	var reqBody, resBody LoginExtensionBySubjectNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LogoutBody struct {
	Req    *types.Logout         `xml:"urn:vim25 Logout,omitempty"`
	Res    *types.LogoutResponse `xml:"LogoutResponse,omitempty"`
	Fault_ *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LogoutBody) Fault() *soap.Fault { return b.Fault_ }

func Logout(ctx context.Context, r soap.RoundTripper, req *types.Logout) (*types.LogoutResponse, error) {
	var reqBody, resBody LogoutBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LookupDvPortGroupBody struct {
	Req    *types.LookupDvPortGroup         `xml:"urn:vim25 LookupDvPortGroup,omitempty"`
	Res    *types.LookupDvPortGroupResponse `xml:"LookupDvPortGroupResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LookupDvPortGroupBody) Fault() *soap.Fault { return b.Fault_ }

func LookupDvPortGroup(ctx context.Context, r soap.RoundTripper, req *types.LookupDvPortGroup) (*types.LookupDvPortGroupResponse, error) {
	var reqBody, resBody LookupDvPortGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type LookupVmOverheadMemoryBody struct {
	Req    *types.LookupVmOverheadMemory         `xml:"urn:vim25 LookupVmOverheadMemory,omitempty"`
	Res    *types.LookupVmOverheadMemoryResponse `xml:"LookupVmOverheadMemoryResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LookupVmOverheadMemoryBody) Fault() *soap.Fault { return b.Fault_ }

func LookupVmOverheadMemory(ctx context.Context, r soap.RoundTripper, req *types.LookupVmOverheadMemory) (*types.LookupVmOverheadMemoryResponse, error) {
	var reqBody, resBody LookupVmOverheadMemoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MakeDirectoryBody struct {
	Req    *types.MakeDirectory         `xml:"urn:vim25 MakeDirectory,omitempty"`
	Res    *types.MakeDirectoryResponse `xml:"MakeDirectoryResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakeDirectoryBody) Fault() *soap.Fault { return b.Fault_ }

func MakeDirectory(ctx context.Context, r soap.RoundTripper, req *types.MakeDirectory) (*types.MakeDirectoryResponse, error) {
	var reqBody, resBody MakeDirectoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MakeDirectoryInGuestBody struct {
	Req    *types.MakeDirectoryInGuest         `xml:"urn:vim25 MakeDirectoryInGuest,omitempty"`
	Res    *types.MakeDirectoryInGuestResponse `xml:"MakeDirectoryInGuestResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakeDirectoryInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func MakeDirectoryInGuest(ctx context.Context, r soap.RoundTripper, req *types.MakeDirectoryInGuest) (*types.MakeDirectoryInGuestResponse, error) {
	var reqBody, resBody MakeDirectoryInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MakePrimaryVM_TaskBody struct {
	Req    *types.MakePrimaryVM_Task         `xml:"urn:vim25 MakePrimaryVM_Task,omitempty"`
	Res    *types.MakePrimaryVM_TaskResponse `xml:"MakePrimaryVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakePrimaryVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MakePrimaryVM_Task(ctx context.Context, r soap.RoundTripper, req *types.MakePrimaryVM_Task) (*types.MakePrimaryVM_TaskResponse, error) {
	var reqBody, resBody MakePrimaryVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsLocal_TaskBody struct {
	Req    *types.MarkAsLocal_Task         `xml:"urn:vim25 MarkAsLocal_Task,omitempty"`
	Res    *types.MarkAsLocal_TaskResponse `xml:"MarkAsLocal_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsLocal_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsLocal_Task(ctx context.Context, r soap.RoundTripper, req *types.MarkAsLocal_Task) (*types.MarkAsLocal_TaskResponse, error) {
	var reqBody, resBody MarkAsLocal_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsNonLocal_TaskBody struct {
	Req    *types.MarkAsNonLocal_Task         `xml:"urn:vim25 MarkAsNonLocal_Task,omitempty"`
	Res    *types.MarkAsNonLocal_TaskResponse `xml:"MarkAsNonLocal_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsNonLocal_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsNonLocal_Task(ctx context.Context, r soap.RoundTripper, req *types.MarkAsNonLocal_Task) (*types.MarkAsNonLocal_TaskResponse, error) {
	var reqBody, resBody MarkAsNonLocal_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsNonSsd_TaskBody struct {
	Req    *types.MarkAsNonSsd_Task         `xml:"urn:vim25 MarkAsNonSsd_Task,omitempty"`
	Res    *types.MarkAsNonSsd_TaskResponse `xml:"MarkAsNonSsd_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsNonSsd_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsNonSsd_Task(ctx context.Context, r soap.RoundTripper, req *types.MarkAsNonSsd_Task) (*types.MarkAsNonSsd_TaskResponse, error) {
	var reqBody, resBody MarkAsNonSsd_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsSsd_TaskBody struct {
	Req    *types.MarkAsSsd_Task         `xml:"urn:vim25 MarkAsSsd_Task,omitempty"`
	Res    *types.MarkAsSsd_TaskResponse `xml:"MarkAsSsd_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsSsd_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsSsd_Task(ctx context.Context, r soap.RoundTripper, req *types.MarkAsSsd_Task) (*types.MarkAsSsd_TaskResponse, error) {
	var reqBody, resBody MarkAsSsd_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsTemplateBody struct {
	Req    *types.MarkAsTemplate         `xml:"urn:vim25 MarkAsTemplate,omitempty"`
	Res    *types.MarkAsTemplateResponse `xml:"MarkAsTemplateResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsTemplateBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsTemplate(ctx context.Context, r soap.RoundTripper, req *types.MarkAsTemplate) (*types.MarkAsTemplateResponse, error) {
	var reqBody, resBody MarkAsTemplateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsVirtualMachineBody struct {
	Req    *types.MarkAsVirtualMachine         `xml:"urn:vim25 MarkAsVirtualMachine,omitempty"`
	Res    *types.MarkAsVirtualMachineResponse `xml:"MarkAsVirtualMachineResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsVirtualMachineBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsVirtualMachine(ctx context.Context, r soap.RoundTripper, req *types.MarkAsVirtualMachine) (*types.MarkAsVirtualMachineResponse, error) {
	var reqBody, resBody MarkAsVirtualMachineBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkDefaultBody struct {
	Req    *types.MarkDefault         `xml:"urn:vim25 MarkDefault,omitempty"`
	Res    *types.MarkDefaultResponse `xml:"MarkDefaultResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkDefaultBody) Fault() *soap.Fault { return b.Fault_ }

func MarkDefault(ctx context.Context, r soap.RoundTripper, req *types.MarkDefault) (*types.MarkDefaultResponse, error) {
	var reqBody, resBody MarkDefaultBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkForRemovalBody struct {
	Req    *types.MarkForRemoval         `xml:"urn:vim25 MarkForRemoval,omitempty"`
	Res    *types.MarkForRemovalResponse `xml:"MarkForRemovalResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkForRemovalBody) Fault() *soap.Fault { return b.Fault_ }

func MarkForRemoval(ctx context.Context, r soap.RoundTripper, req *types.MarkForRemoval) (*types.MarkForRemovalResponse, error) {
	var reqBody, resBody MarkForRemovalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkPerenniallyReservedBody struct {
	Req    *types.MarkPerenniallyReserved         `xml:"urn:vim25 MarkPerenniallyReserved,omitempty"`
	Res    *types.MarkPerenniallyReservedResponse `xml:"MarkPerenniallyReservedResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkPerenniallyReservedBody) Fault() *soap.Fault { return b.Fault_ }

func MarkPerenniallyReserved(ctx context.Context, r soap.RoundTripper, req *types.MarkPerenniallyReserved) (*types.MarkPerenniallyReservedResponse, error) {
	var reqBody, resBody MarkPerenniallyReservedBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkPerenniallyReservedEx_TaskBody struct {
	Req    *types.MarkPerenniallyReservedEx_Task         `xml:"urn:vim25 MarkPerenniallyReservedEx_Task,omitempty"`
	Res    *types.MarkPerenniallyReservedEx_TaskResponse `xml:"MarkPerenniallyReservedEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkPerenniallyReservedEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MarkPerenniallyReservedEx_Task(ctx context.Context, r soap.RoundTripper, req *types.MarkPerenniallyReservedEx_Task) (*types.MarkPerenniallyReservedEx_TaskResponse, error) {
	var reqBody, resBody MarkPerenniallyReservedEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkServiceProviderEntitiesBody struct {
	Req    *types.MarkServiceProviderEntities         `xml:"urn:vim25 MarkServiceProviderEntities,omitempty"`
	Res    *types.MarkServiceProviderEntitiesResponse `xml:"MarkServiceProviderEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkServiceProviderEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func MarkServiceProviderEntities(ctx context.Context, r soap.RoundTripper, req *types.MarkServiceProviderEntities) (*types.MarkServiceProviderEntitiesResponse, error) {
	var reqBody, resBody MarkServiceProviderEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MergeDvs_TaskBody struct {
	Req    *types.MergeDvs_Task         `xml:"urn:vim25 MergeDvs_Task,omitempty"`
	Res    *types.MergeDvs_TaskResponse `xml:"MergeDvs_TaskResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MergeDvs_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MergeDvs_Task(ctx context.Context, r soap.RoundTripper, req *types.MergeDvs_Task) (*types.MergeDvs_TaskResponse, error) {
	var reqBody, resBody MergeDvs_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MergePermissionsBody struct {
	Req    *types.MergePermissions         `xml:"urn:vim25 MergePermissions,omitempty"`
	Res    *types.MergePermissionsResponse `xml:"MergePermissionsResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MergePermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func MergePermissions(ctx context.Context, r soap.RoundTripper, req *types.MergePermissions) (*types.MergePermissionsResponse, error) {
	var reqBody, resBody MergePermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MigrateVM_TaskBody struct {
	Req    *types.MigrateVM_Task         `xml:"urn:vim25 MigrateVM_Task,omitempty"`
	Res    *types.MigrateVM_TaskResponse `xml:"MigrateVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MigrateVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MigrateVM_Task(ctx context.Context, r soap.RoundTripper, req *types.MigrateVM_Task) (*types.MigrateVM_TaskResponse, error) {
	var reqBody, resBody MigrateVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ModifyListViewBody struct {
	Req    *types.ModifyListView         `xml:"urn:vim25 ModifyListView,omitempty"`
	Res    *types.ModifyListViewResponse `xml:"ModifyListViewResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ModifyListViewBody) Fault() *soap.Fault { return b.Fault_ }

func ModifyListView(ctx context.Context, r soap.RoundTripper, req *types.ModifyListView) (*types.ModifyListViewResponse, error) {
	var reqBody, resBody ModifyListViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MountToolsInstallerBody struct {
	Req    *types.MountToolsInstaller         `xml:"urn:vim25 MountToolsInstaller,omitempty"`
	Res    *types.MountToolsInstallerResponse `xml:"MountToolsInstallerResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountToolsInstallerBody) Fault() *soap.Fault { return b.Fault_ }

func MountToolsInstaller(ctx context.Context, r soap.RoundTripper, req *types.MountToolsInstaller) (*types.MountToolsInstallerResponse, error) {
	var reqBody, resBody MountToolsInstallerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MountVffsVolumeBody struct {
	Req    *types.MountVffsVolume         `xml:"urn:vim25 MountVffsVolume,omitempty"`
	Res    *types.MountVffsVolumeResponse `xml:"MountVffsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountVffsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func MountVffsVolume(ctx context.Context, r soap.RoundTripper, req *types.MountVffsVolume) (*types.MountVffsVolumeResponse, error) {
	var reqBody, resBody MountVffsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MountVmfsVolumeBody struct {
	Req    *types.MountVmfsVolume         `xml:"urn:vim25 MountVmfsVolume,omitempty"`
	Res    *types.MountVmfsVolumeResponse `xml:"MountVmfsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountVmfsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func MountVmfsVolume(ctx context.Context, r soap.RoundTripper, req *types.MountVmfsVolume) (*types.MountVmfsVolumeResponse, error) {
	var reqBody, resBody MountVmfsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MountVmfsVolumeEx_TaskBody struct {
	Req    *types.MountVmfsVolumeEx_Task         `xml:"urn:vim25 MountVmfsVolumeEx_Task,omitempty"`
	Res    *types.MountVmfsVolumeEx_TaskResponse `xml:"MountVmfsVolumeEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountVmfsVolumeEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MountVmfsVolumeEx_Task(ctx context.Context, r soap.RoundTripper, req *types.MountVmfsVolumeEx_Task) (*types.MountVmfsVolumeEx_TaskResponse, error) {
	var reqBody, resBody MountVmfsVolumeEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveDVPort_TaskBody struct {
	Req    *types.MoveDVPort_Task         `xml:"urn:vim25 MoveDVPort_Task,omitempty"`
	Res    *types.MoveDVPort_TaskResponse `xml:"MoveDVPort_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDVPort_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveDVPort_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveDVPort_Task) (*types.MoveDVPort_TaskResponse, error) {
	var reqBody, resBody MoveDVPort_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveDatastoreFile_TaskBody struct {
	Req    *types.MoveDatastoreFile_Task         `xml:"urn:vim25 MoveDatastoreFile_Task,omitempty"`
	Res    *types.MoveDatastoreFile_TaskResponse `xml:"MoveDatastoreFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDatastoreFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveDatastoreFile_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveDatastoreFile_Task) (*types.MoveDatastoreFile_TaskResponse, error) {
	var reqBody, resBody MoveDatastoreFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveDirectoryInGuestBody struct {
	Req    *types.MoveDirectoryInGuest         `xml:"urn:vim25 MoveDirectoryInGuest,omitempty"`
	Res    *types.MoveDirectoryInGuestResponse `xml:"MoveDirectoryInGuestResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDirectoryInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func MoveDirectoryInGuest(ctx context.Context, r soap.RoundTripper, req *types.MoveDirectoryInGuest) (*types.MoveDirectoryInGuestResponse, error) {
	var reqBody, resBody MoveDirectoryInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveFileInGuestBody struct {
	Req    *types.MoveFileInGuest         `xml:"urn:vim25 MoveFileInGuest,omitempty"`
	Res    *types.MoveFileInGuestResponse `xml:"MoveFileInGuestResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveFileInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func MoveFileInGuest(ctx context.Context, r soap.RoundTripper, req *types.MoveFileInGuest) (*types.MoveFileInGuestResponse, error) {
	var reqBody, resBody MoveFileInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveHostInto_TaskBody struct {
	Req    *types.MoveHostInto_Task         `xml:"urn:vim25 MoveHostInto_Task,omitempty"`
	Res    *types.MoveHostInto_TaskResponse `xml:"MoveHostInto_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveHostInto_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveHostInto_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveHostInto_Task) (*types.MoveHostInto_TaskResponse, error) {
	var reqBody, resBody MoveHostInto_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveIntoFolder_TaskBody struct {
	Req    *types.MoveIntoFolder_Task         `xml:"urn:vim25 MoveIntoFolder_Task,omitempty"`
	Res    *types.MoveIntoFolder_TaskResponse `xml:"MoveIntoFolder_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveIntoFolder_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveIntoFolder_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveIntoFolder_Task) (*types.MoveIntoFolder_TaskResponse, error) {
	var reqBody, resBody MoveIntoFolder_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveIntoResourcePoolBody struct {
	Req    *types.MoveIntoResourcePool         `xml:"urn:vim25 MoveIntoResourcePool,omitempty"`
	Res    *types.MoveIntoResourcePoolResponse `xml:"MoveIntoResourcePoolResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveIntoResourcePoolBody) Fault() *soap.Fault { return b.Fault_ }

func MoveIntoResourcePool(ctx context.Context, r soap.RoundTripper, req *types.MoveIntoResourcePool) (*types.MoveIntoResourcePoolResponse, error) {
	var reqBody, resBody MoveIntoResourcePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveInto_TaskBody struct {
	Req    *types.MoveInto_Task         `xml:"urn:vim25 MoveInto_Task,omitempty"`
	Res    *types.MoveInto_TaskResponse `xml:"MoveInto_TaskResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveInto_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveInto_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveInto_Task) (*types.MoveInto_TaskResponse, error) {
	var reqBody, resBody MoveInto_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MoveVirtualDisk_TaskBody struct {
	Req    *types.MoveVirtualDisk_Task         `xml:"urn:vim25 MoveVirtualDisk_Task,omitempty"`
	Res    *types.MoveVirtualDisk_TaskResponse `xml:"MoveVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func MoveVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.MoveVirtualDisk_Task) (*types.MoveVirtualDisk_TaskResponse, error) {
	var reqBody, resBody MoveVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type NotifyAffectedServicesBody struct {
	Req    *types.NotifyAffectedServices         `xml:"urn:vim25 NotifyAffectedServices,omitempty"`
	Res    *types.NotifyAffectedServicesResponse `xml:"NotifyAffectedServicesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *NotifyAffectedServicesBody) Fault() *soap.Fault { return b.Fault_ }

func NotifyAffectedServices(ctx context.Context, r soap.RoundTripper, req *types.NotifyAffectedServices) (*types.NotifyAffectedServicesResponse, error) {
	var reqBody, resBody NotifyAffectedServicesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type OpenInventoryViewFolderBody struct {
	Req    *types.OpenInventoryViewFolder         `xml:"urn:vim25 OpenInventoryViewFolder,omitempty"`
	Res    *types.OpenInventoryViewFolderResponse `xml:"OpenInventoryViewFolderResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *OpenInventoryViewFolderBody) Fault() *soap.Fault { return b.Fault_ }

func OpenInventoryViewFolder(ctx context.Context, r soap.RoundTripper, req *types.OpenInventoryViewFolder) (*types.OpenInventoryViewFolderResponse, error) {
	var reqBody, resBody OpenInventoryViewFolderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type OverwriteCustomizationSpecBody struct {
	Req    *types.OverwriteCustomizationSpec         `xml:"urn:vim25 OverwriteCustomizationSpec,omitempty"`
	Res    *types.OverwriteCustomizationSpecResponse `xml:"OverwriteCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *OverwriteCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func OverwriteCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.OverwriteCustomizationSpec) (*types.OverwriteCustomizationSpecResponse, error) {
	var reqBody, resBody OverwriteCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ParseDescriptorBody struct {
	Req    *types.ParseDescriptor         `xml:"urn:vim25 ParseDescriptor,omitempty"`
	Res    *types.ParseDescriptorResponse `xml:"ParseDescriptorResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ParseDescriptorBody) Fault() *soap.Fault { return b.Fault_ }

func ParseDescriptor(ctx context.Context, r soap.RoundTripper, req *types.ParseDescriptor) (*types.ParseDescriptorResponse, error) {
	var reqBody, resBody ParseDescriptorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformDvsProductSpecOperation_TaskBody struct {
	Req    *types.PerformDvsProductSpecOperation_Task         `xml:"urn:vim25 PerformDvsProductSpecOperation_Task,omitempty"`
	Res    *types.PerformDvsProductSpecOperation_TaskResponse `xml:"PerformDvsProductSpecOperation_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformDvsProductSpecOperation_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PerformDvsProductSpecOperation_Task(ctx context.Context, r soap.RoundTripper, req *types.PerformDvsProductSpecOperation_Task) (*types.PerformDvsProductSpecOperation_TaskResponse, error) {
	var reqBody, resBody PerformDvsProductSpecOperation_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgradePreflightCheckBody struct {
	Req    *types.PerformVsanUpgradePreflightCheck         `xml:"urn:vim25 PerformVsanUpgradePreflightCheck,omitempty"`
	Res    *types.PerformVsanUpgradePreflightCheckResponse `xml:"PerformVsanUpgradePreflightCheckResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgradePreflightCheckBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgradePreflightCheck(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgradePreflightCheck) (*types.PerformVsanUpgradePreflightCheckResponse, error) {
	var reqBody, resBody PerformVsanUpgradePreflightCheckBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PerformVsanUpgrade_TaskBody struct {
	Req    *types.PerformVsanUpgrade_Task         `xml:"urn:vim25 PerformVsanUpgrade_Task,omitempty"`
	Res    *types.PerformVsanUpgrade_TaskResponse `xml:"PerformVsanUpgrade_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformVsanUpgrade_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PerformVsanUpgrade_Task(ctx context.Context, r soap.RoundTripper, req *types.PerformVsanUpgrade_Task) (*types.PerformVsanUpgrade_TaskResponse, error) {
	var reqBody, resBody PerformVsanUpgrade_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PlaceVmBody struct {
	Req    *types.PlaceVm         `xml:"urn:vim25 PlaceVm,omitempty"`
	Res    *types.PlaceVmResponse `xml:"PlaceVmResponse,omitempty"`
	Fault_ *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PlaceVmBody) Fault() *soap.Fault { return b.Fault_ }

func PlaceVm(ctx context.Context, r soap.RoundTripper, req *types.PlaceVm) (*types.PlaceVmResponse, error) {
	var reqBody, resBody PlaceVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PostEventBody struct {
	Req    *types.PostEvent         `xml:"urn:vim25 PostEvent,omitempty"`
	Res    *types.PostEventResponse `xml:"PostEventResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PostEventBody) Fault() *soap.Fault { return b.Fault_ }

func PostEvent(ctx context.Context, r soap.RoundTripper, req *types.PostEvent) (*types.PostEventResponse, error) {
	var reqBody, resBody PostEventBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PostHealthUpdatesBody struct {
	Req    *types.PostHealthUpdates         `xml:"urn:vim25 PostHealthUpdates,omitempty"`
	Res    *types.PostHealthUpdatesResponse `xml:"PostHealthUpdatesResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PostHealthUpdatesBody) Fault() *soap.Fault { return b.Fault_ }

func PostHealthUpdates(ctx context.Context, r soap.RoundTripper, req *types.PostHealthUpdates) (*types.PostHealthUpdatesResponse, error) {
	var reqBody, resBody PostHealthUpdatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerDownHostToStandBy_TaskBody struct {
	Req    *types.PowerDownHostToStandBy_Task         `xml:"urn:vim25 PowerDownHostToStandBy_Task,omitempty"`
	Res    *types.PowerDownHostToStandBy_TaskResponse `xml:"PowerDownHostToStandBy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerDownHostToStandBy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerDownHostToStandBy_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerDownHostToStandBy_Task) (*types.PowerDownHostToStandBy_TaskResponse, error) {
	var reqBody, resBody PowerDownHostToStandBy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerOffVApp_TaskBody struct {
	Req    *types.PowerOffVApp_Task         `xml:"urn:vim25 PowerOffVApp_Task,omitempty"`
	Res    *types.PowerOffVApp_TaskResponse `xml:"PowerOffVApp_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOffVApp_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerOffVApp_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerOffVApp_Task) (*types.PowerOffVApp_TaskResponse, error) {
	var reqBody, resBody PowerOffVApp_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerOffVM_TaskBody struct {
	Req    *types.PowerOffVM_Task         `xml:"urn:vim25 PowerOffVM_Task,omitempty"`
	Res    *types.PowerOffVM_TaskResponse `xml:"PowerOffVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOffVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerOffVM_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerOffVM_Task) (*types.PowerOffVM_TaskResponse, error) {
	var reqBody, resBody PowerOffVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerOnMultiVM_TaskBody struct {
	Req    *types.PowerOnMultiVM_Task         `xml:"urn:vim25 PowerOnMultiVM_Task,omitempty"`
	Res    *types.PowerOnMultiVM_TaskResponse `xml:"PowerOnMultiVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnMultiVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerOnMultiVM_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerOnMultiVM_Task) (*types.PowerOnMultiVM_TaskResponse, error) {
	var reqBody, resBody PowerOnMultiVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerOnVApp_TaskBody struct {
	Req    *types.PowerOnVApp_Task         `xml:"urn:vim25 PowerOnVApp_Task,omitempty"`
	Res    *types.PowerOnVApp_TaskResponse `xml:"PowerOnVApp_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnVApp_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerOnVApp_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerOnVApp_Task) (*types.PowerOnVApp_TaskResponse, error) {
	var reqBody, resBody PowerOnVApp_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerOnVM_TaskBody struct {
	Req    *types.PowerOnVM_Task         `xml:"urn:vim25 PowerOnVM_Task,omitempty"`
	Res    *types.PowerOnVM_TaskResponse `xml:"PowerOnVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerOnVM_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerOnVM_Task) (*types.PowerOnVM_TaskResponse, error) {
	var reqBody, resBody PowerOnVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PowerUpHostFromStandBy_TaskBody struct {
	Req    *types.PowerUpHostFromStandBy_Task         `xml:"urn:vim25 PowerUpHostFromStandBy_Task,omitempty"`
	Res    *types.PowerUpHostFromStandBy_TaskResponse `xml:"PowerUpHostFromStandBy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerUpHostFromStandBy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PowerUpHostFromStandBy_Task(ctx context.Context, r soap.RoundTripper, req *types.PowerUpHostFromStandBy_Task) (*types.PowerUpHostFromStandBy_TaskResponse, error) {
	var reqBody, resBody PowerUpHostFromStandBy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PrepareCryptoBody struct {
	Req    *types.PrepareCrypto         `xml:"urn:vim25 PrepareCrypto,omitempty"`
	Res    *types.PrepareCryptoResponse `xml:"PrepareCryptoResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PrepareCryptoBody) Fault() *soap.Fault { return b.Fault_ }

func PrepareCrypto(ctx context.Context, r soap.RoundTripper, req *types.PrepareCrypto) (*types.PrepareCryptoResponse, error) {
	var reqBody, resBody PrepareCryptoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PromoteDisks_TaskBody struct {
	Req    *types.PromoteDisks_Task         `xml:"urn:vim25 PromoteDisks_Task,omitempty"`
	Res    *types.PromoteDisks_TaskResponse `xml:"PromoteDisks_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PromoteDisks_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PromoteDisks_Task(ctx context.Context, r soap.RoundTripper, req *types.PromoteDisks_Task) (*types.PromoteDisks_TaskResponse, error) {
	var reqBody, resBody PromoteDisks_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ProvisionServerPrivateKeyBody struct {
	Req    *types.ProvisionServerPrivateKey         `xml:"urn:vim25 ProvisionServerPrivateKey,omitempty"`
	Res    *types.ProvisionServerPrivateKeyResponse `xml:"ProvisionServerPrivateKeyResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ProvisionServerPrivateKeyBody) Fault() *soap.Fault { return b.Fault_ }

func ProvisionServerPrivateKey(ctx context.Context, r soap.RoundTripper, req *types.ProvisionServerPrivateKey) (*types.ProvisionServerPrivateKeyResponse, error) {
	var reqBody, resBody ProvisionServerPrivateKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PutUsbScanCodesBody struct {
	Req    *types.PutUsbScanCodes         `xml:"urn:vim25 PutUsbScanCodes,omitempty"`
	Res    *types.PutUsbScanCodesResponse `xml:"PutUsbScanCodesResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PutUsbScanCodesBody) Fault() *soap.Fault { return b.Fault_ }

func PutUsbScanCodes(ctx context.Context, r soap.RoundTripper, req *types.PutUsbScanCodes) (*types.PutUsbScanCodesResponse, error) {
	var reqBody, resBody PutUsbScanCodesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAnswerFileStatusBody struct {
	Req    *types.QueryAnswerFileStatus         `xml:"urn:vim25 QueryAnswerFileStatus,omitempty"`
	Res    *types.QueryAnswerFileStatusResponse `xml:"QueryAnswerFileStatusResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAnswerFileStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAnswerFileStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryAnswerFileStatus) (*types.QueryAnswerFileStatusResponse, error) {
	var reqBody, resBody QueryAnswerFileStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAssignedLicensesBody struct {
	Req    *types.QueryAssignedLicenses         `xml:"urn:vim25 QueryAssignedLicenses,omitempty"`
	Res    *types.QueryAssignedLicensesResponse `xml:"QueryAssignedLicensesResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAssignedLicensesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAssignedLicenses(ctx context.Context, r soap.RoundTripper, req *types.QueryAssignedLicenses) (*types.QueryAssignedLicensesResponse, error) {
	var reqBody, resBody QueryAssignedLicensesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailableDisksForVmfsBody struct {
	Req    *types.QueryAvailableDisksForVmfs         `xml:"urn:vim25 QueryAvailableDisksForVmfs,omitempty"`
	Res    *types.QueryAvailableDisksForVmfsResponse `xml:"QueryAvailableDisksForVmfsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableDisksForVmfsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailableDisksForVmfs(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailableDisksForVmfs) (*types.QueryAvailableDisksForVmfsResponse, error) {
	var reqBody, resBody QueryAvailableDisksForVmfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailableDvsSpecBody struct {
	Req    *types.QueryAvailableDvsSpec         `xml:"urn:vim25 QueryAvailableDvsSpec,omitempty"`
	Res    *types.QueryAvailableDvsSpecResponse `xml:"QueryAvailableDvsSpecResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableDvsSpecBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailableDvsSpec(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailableDvsSpec) (*types.QueryAvailableDvsSpecResponse, error) {
	var reqBody, resBody QueryAvailableDvsSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailablePartitionBody struct {
	Req    *types.QueryAvailablePartition         `xml:"urn:vim25 QueryAvailablePartition,omitempty"`
	Res    *types.QueryAvailablePartitionResponse `xml:"QueryAvailablePartitionResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailablePartitionBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailablePartition(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailablePartition) (*types.QueryAvailablePartitionResponse, error) {
	var reqBody, resBody QueryAvailablePartitionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailablePerfMetricBody struct {
	Req    *types.QueryAvailablePerfMetric         `xml:"urn:vim25 QueryAvailablePerfMetric,omitempty"`
	Res    *types.QueryAvailablePerfMetricResponse `xml:"QueryAvailablePerfMetricResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailablePerfMetricBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailablePerfMetric(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailablePerfMetric) (*types.QueryAvailablePerfMetricResponse, error) {
	var reqBody, resBody QueryAvailablePerfMetricBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailableSsdsBody struct {
	Req    *types.QueryAvailableSsds         `xml:"urn:vim25 QueryAvailableSsds,omitempty"`
	Res    *types.QueryAvailableSsdsResponse `xml:"QueryAvailableSsdsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableSsdsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailableSsds(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailableSsds) (*types.QueryAvailableSsdsResponse, error) {
	var reqBody, resBody QueryAvailableSsdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAvailableTimeZonesBody struct {
	Req    *types.QueryAvailableTimeZones         `xml:"urn:vim25 QueryAvailableTimeZones,omitempty"`
	Res    *types.QueryAvailableTimeZonesResponse `xml:"QueryAvailableTimeZonesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableTimeZonesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAvailableTimeZones(ctx context.Context, r soap.RoundTripper, req *types.QueryAvailableTimeZones) (*types.QueryAvailableTimeZonesResponse, error) {
	var reqBody, resBody QueryAvailableTimeZonesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryBootDevicesBody struct {
	Req    *types.QueryBootDevices         `xml:"urn:vim25 QueryBootDevices,omitempty"`
	Res    *types.QueryBootDevicesResponse `xml:"QueryBootDevicesResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryBootDevicesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryBootDevices(ctx context.Context, r soap.RoundTripper, req *types.QueryBootDevices) (*types.QueryBootDevicesResponse, error) {
	var reqBody, resBody QueryBootDevicesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryBoundVnicsBody struct {
	Req    *types.QueryBoundVnics         `xml:"urn:vim25 QueryBoundVnics,omitempty"`
	Res    *types.QueryBoundVnicsResponse `xml:"QueryBoundVnicsResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryBoundVnicsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryBoundVnics(ctx context.Context, r soap.RoundTripper, req *types.QueryBoundVnics) (*types.QueryBoundVnicsResponse, error) {
	var reqBody, resBody QueryBoundVnicsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCandidateNicsBody struct {
	Req    *types.QueryCandidateNics         `xml:"urn:vim25 QueryCandidateNics,omitempty"`
	Res    *types.QueryCandidateNicsResponse `xml:"QueryCandidateNicsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCandidateNicsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCandidateNics(ctx context.Context, r soap.RoundTripper, req *types.QueryCandidateNics) (*types.QueryCandidateNicsResponse, error) {
	var reqBody, resBody QueryCandidateNicsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryChangedDiskAreasBody struct {
	Req    *types.QueryChangedDiskAreas         `xml:"urn:vim25 QueryChangedDiskAreas,omitempty"`
	Res    *types.QueryChangedDiskAreasResponse `xml:"QueryChangedDiskAreasResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryChangedDiskAreasBody) Fault() *soap.Fault { return b.Fault_ }

func QueryChangedDiskAreas(ctx context.Context, r soap.RoundTripper, req *types.QueryChangedDiskAreas) (*types.QueryChangedDiskAreasResponse, error) {
	var reqBody, resBody QueryChangedDiskAreasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCmmdsBody struct {
	Req    *types.QueryCmmds         `xml:"urn:vim25 QueryCmmds,omitempty"`
	Res    *types.QueryCmmdsResponse `xml:"QueryCmmdsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCmmdsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCmmds(ctx context.Context, r soap.RoundTripper, req *types.QueryCmmds) (*types.QueryCmmdsResponse, error) {
	var reqBody, resBody QueryCmmdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCompatibleHostForExistingDvsBody struct {
	Req    *types.QueryCompatibleHostForExistingDvs         `xml:"urn:vim25 QueryCompatibleHostForExistingDvs,omitempty"`
	Res    *types.QueryCompatibleHostForExistingDvsResponse `xml:"QueryCompatibleHostForExistingDvsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCompatibleHostForExistingDvsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCompatibleHostForExistingDvs(ctx context.Context, r soap.RoundTripper, req *types.QueryCompatibleHostForExistingDvs) (*types.QueryCompatibleHostForExistingDvsResponse, error) {
	var reqBody, resBody QueryCompatibleHostForExistingDvsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCompatibleHostForNewDvsBody struct {
	Req    *types.QueryCompatibleHostForNewDvs         `xml:"urn:vim25 QueryCompatibleHostForNewDvs,omitempty"`
	Res    *types.QueryCompatibleHostForNewDvsResponse `xml:"QueryCompatibleHostForNewDvsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCompatibleHostForNewDvsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCompatibleHostForNewDvs(ctx context.Context, r soap.RoundTripper, req *types.QueryCompatibleHostForNewDvs) (*types.QueryCompatibleHostForNewDvsResponse, error) {
	var reqBody, resBody QueryCompatibleHostForNewDvsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCompatibleVmnicsFromHostsBody struct {
	Req    *types.QueryCompatibleVmnicsFromHosts         `xml:"urn:vim25 QueryCompatibleVmnicsFromHosts,omitempty"`
	Res    *types.QueryCompatibleVmnicsFromHostsResponse `xml:"QueryCompatibleVmnicsFromHostsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCompatibleVmnicsFromHostsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCompatibleVmnicsFromHosts(ctx context.Context, r soap.RoundTripper, req *types.QueryCompatibleVmnicsFromHosts) (*types.QueryCompatibleVmnicsFromHostsResponse, error) {
	var reqBody, resBody QueryCompatibleVmnicsFromHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryComplianceStatusBody struct {
	Req    *types.QueryComplianceStatus         `xml:"urn:vim25 QueryComplianceStatus,omitempty"`
	Res    *types.QueryComplianceStatusResponse `xml:"QueryComplianceStatusResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryComplianceStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryComplianceStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryComplianceStatus) (*types.QueryComplianceStatusResponse, error) {
	var reqBody, resBody QueryComplianceStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfigOptionBody struct {
	Req    *types.QueryConfigOption         `xml:"urn:vim25 QueryConfigOption,omitempty"`
	Res    *types.QueryConfigOptionResponse `xml:"QueryConfigOptionResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigOptionBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfigOption(ctx context.Context, r soap.RoundTripper, req *types.QueryConfigOption) (*types.QueryConfigOptionResponse, error) {
	var reqBody, resBody QueryConfigOptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfigOptionDescriptorBody struct {
	Req    *types.QueryConfigOptionDescriptor         `xml:"urn:vim25 QueryConfigOptionDescriptor,omitempty"`
	Res    *types.QueryConfigOptionDescriptorResponse `xml:"QueryConfigOptionDescriptorResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigOptionDescriptorBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfigOptionDescriptor(ctx context.Context, r soap.RoundTripper, req *types.QueryConfigOptionDescriptor) (*types.QueryConfigOptionDescriptorResponse, error) {
	var reqBody, resBody QueryConfigOptionDescriptorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfigOptionExBody struct {
	Req    *types.QueryConfigOptionEx         `xml:"urn:vim25 QueryConfigOptionEx,omitempty"`
	Res    *types.QueryConfigOptionExResponse `xml:"QueryConfigOptionExResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigOptionExBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfigOptionEx(ctx context.Context, r soap.RoundTripper, req *types.QueryConfigOptionEx) (*types.QueryConfigOptionExResponse, error) {
	var reqBody, resBody QueryConfigOptionExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfigTargetBody struct {
	Req    *types.QueryConfigTarget         `xml:"urn:vim25 QueryConfigTarget,omitempty"`
	Res    *types.QueryConfigTargetResponse `xml:"QueryConfigTargetResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigTargetBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfigTarget(ctx context.Context, r soap.RoundTripper, req *types.QueryConfigTarget) (*types.QueryConfigTargetResponse, error) {
	var reqBody, resBody QueryConfigTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfiguredModuleOptionStringBody struct {
	Req    *types.QueryConfiguredModuleOptionString         `xml:"urn:vim25 QueryConfiguredModuleOptionString,omitempty"`
	Res    *types.QueryConfiguredModuleOptionStringResponse `xml:"QueryConfiguredModuleOptionStringResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfiguredModuleOptionStringBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfiguredModuleOptionString(ctx context.Context, r soap.RoundTripper, req *types.QueryConfiguredModuleOptionString) (*types.QueryConfiguredModuleOptionStringResponse, error) {
	var reqBody, resBody QueryConfiguredModuleOptionStringBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConnectionInfoBody struct {
	Req    *types.QueryConnectionInfo         `xml:"urn:vim25 QueryConnectionInfo,omitempty"`
	Res    *types.QueryConnectionInfoResponse `xml:"QueryConnectionInfoResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConnectionInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConnectionInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryConnectionInfo) (*types.QueryConnectionInfoResponse, error) {
	var reqBody, resBody QueryConnectionInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConnectionInfoViaSpecBody struct {
	Req    *types.QueryConnectionInfoViaSpec         `xml:"urn:vim25 QueryConnectionInfoViaSpec,omitempty"`
	Res    *types.QueryConnectionInfoViaSpecResponse `xml:"QueryConnectionInfoViaSpecResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConnectionInfoViaSpecBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConnectionInfoViaSpec(ctx context.Context, r soap.RoundTripper, req *types.QueryConnectionInfoViaSpec) (*types.QueryConnectionInfoViaSpecResponse, error) {
	var reqBody, resBody QueryConnectionInfoViaSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConnectionsBody struct {
	Req    *types.QueryConnections         `xml:"urn:vim25 QueryConnections,omitempty"`
	Res    *types.QueryConnectionsResponse `xml:"QueryConnectionsResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConnectionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConnections(ctx context.Context, r soap.RoundTripper, req *types.QueryConnections) (*types.QueryConnectionsResponse, error) {
	var reqBody, resBody QueryConnectionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryCryptoKeyStatusBody struct {
	Req    *types.QueryCryptoKeyStatus         `xml:"urn:vim25 QueryCryptoKeyStatus,omitempty"`
	Res    *types.QueryCryptoKeyStatusResponse `xml:"QueryCryptoKeyStatusResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCryptoKeyStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryCryptoKeyStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryCryptoKeyStatus) (*types.QueryCryptoKeyStatusResponse, error) {
	var reqBody, resBody QueryCryptoKeyStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDatastorePerformanceSummaryBody struct {
	Req    *types.QueryDatastorePerformanceSummary         `xml:"urn:vim25 QueryDatastorePerformanceSummary,omitempty"`
	Res    *types.QueryDatastorePerformanceSummaryResponse `xml:"QueryDatastorePerformanceSummaryResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDatastorePerformanceSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDatastorePerformanceSummary(ctx context.Context, r soap.RoundTripper, req *types.QueryDatastorePerformanceSummary) (*types.QueryDatastorePerformanceSummaryResponse, error) {
	var reqBody, resBody QueryDatastorePerformanceSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDateTimeBody struct {
	Req    *types.QueryDateTime         `xml:"urn:vim25 QueryDateTime,omitempty"`
	Res    *types.QueryDateTimeResponse `xml:"QueryDateTimeResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDateTimeBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDateTime(ctx context.Context, r soap.RoundTripper, req *types.QueryDateTime) (*types.QueryDateTimeResponse, error) {
	var reqBody, resBody QueryDateTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDescriptionsBody struct {
	Req    *types.QueryDescriptions         `xml:"urn:vim25 QueryDescriptions,omitempty"`
	Res    *types.QueryDescriptionsResponse `xml:"QueryDescriptionsResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDescriptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDescriptions(ctx context.Context, r soap.RoundTripper, req *types.QueryDescriptions) (*types.QueryDescriptionsResponse, error) {
	var reqBody, resBody QueryDescriptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDirectoryInfoBody struct {
	Req    *types.QueryDirectoryInfo         `xml:"urn:vim25 QueryDirectoryInfo,omitempty"`
	Res    *types.QueryDirectoryInfoResponse `xml:"QueryDirectoryInfoResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDirectoryInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDirectoryInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryDirectoryInfo) (*types.QueryDirectoryInfoResponse, error) {
	var reqBody, resBody QueryDirectoryInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDisksForVsanBody struct {
	Req    *types.QueryDisksForVsan         `xml:"urn:vim25 QueryDisksForVsan,omitempty"`
	Res    *types.QueryDisksForVsanResponse `xml:"QueryDisksForVsanResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDisksForVsanBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDisksForVsan(ctx context.Context, r soap.RoundTripper, req *types.QueryDisksForVsan) (*types.QueryDisksForVsanResponse, error) {
	var reqBody, resBody QueryDisksForVsanBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDisksUsingFilterBody struct {
	Req    *types.QueryDisksUsingFilter         `xml:"urn:vim25 QueryDisksUsingFilter,omitempty"`
	Res    *types.QueryDisksUsingFilterResponse `xml:"QueryDisksUsingFilterResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDisksUsingFilterBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDisksUsingFilter(ctx context.Context, r soap.RoundTripper, req *types.QueryDisksUsingFilter) (*types.QueryDisksUsingFilterResponse, error) {
	var reqBody, resBody QueryDisksUsingFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDvsByUuidBody struct {
	Req    *types.QueryDvsByUuid         `xml:"urn:vim25 QueryDvsByUuid,omitempty"`
	Res    *types.QueryDvsByUuidResponse `xml:"QueryDvsByUuidResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsByUuidBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDvsByUuid(ctx context.Context, r soap.RoundTripper, req *types.QueryDvsByUuid) (*types.QueryDvsByUuidResponse, error) {
	var reqBody, resBody QueryDvsByUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDvsCheckCompatibilityBody struct {
	Req    *types.QueryDvsCheckCompatibility         `xml:"urn:vim25 QueryDvsCheckCompatibility,omitempty"`
	Res    *types.QueryDvsCheckCompatibilityResponse `xml:"QueryDvsCheckCompatibilityResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsCheckCompatibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDvsCheckCompatibility(ctx context.Context, r soap.RoundTripper, req *types.QueryDvsCheckCompatibility) (*types.QueryDvsCheckCompatibilityResponse, error) {
	var reqBody, resBody QueryDvsCheckCompatibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDvsCompatibleHostSpecBody struct {
	Req    *types.QueryDvsCompatibleHostSpec         `xml:"urn:vim25 QueryDvsCompatibleHostSpec,omitempty"`
	Res    *types.QueryDvsCompatibleHostSpecResponse `xml:"QueryDvsCompatibleHostSpecResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsCompatibleHostSpecBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDvsCompatibleHostSpec(ctx context.Context, r soap.RoundTripper, req *types.QueryDvsCompatibleHostSpec) (*types.QueryDvsCompatibleHostSpecResponse, error) {
	var reqBody, resBody QueryDvsCompatibleHostSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDvsConfigTargetBody struct {
	Req    *types.QueryDvsConfigTarget         `xml:"urn:vim25 QueryDvsConfigTarget,omitempty"`
	Res    *types.QueryDvsConfigTargetResponse `xml:"QueryDvsConfigTargetResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsConfigTargetBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDvsConfigTarget(ctx context.Context, r soap.RoundTripper, req *types.QueryDvsConfigTarget) (*types.QueryDvsConfigTargetResponse, error) {
	var reqBody, resBody QueryDvsConfigTargetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDvsFeatureCapabilityBody struct {
	Req    *types.QueryDvsFeatureCapability         `xml:"urn:vim25 QueryDvsFeatureCapability,omitempty"`
	Res    *types.QueryDvsFeatureCapabilityResponse `xml:"QueryDvsFeatureCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsFeatureCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDvsFeatureCapability(ctx context.Context, r soap.RoundTripper, req *types.QueryDvsFeatureCapability) (*types.QueryDvsFeatureCapabilityResponse, error) {
	var reqBody, resBody QueryDvsFeatureCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryEventsBody struct {
	Req    *types.QueryEvents         `xml:"urn:vim25 QueryEvents,omitempty"`
	Res    *types.QueryEventsResponse `xml:"QueryEventsResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryEventsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryEvents(ctx context.Context, r soap.RoundTripper, req *types.QueryEvents) (*types.QueryEventsResponse, error) {
	var reqBody, resBody QueryEventsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryExpressionMetadataBody struct {
	Req    *types.QueryExpressionMetadata         `xml:"urn:vim25 QueryExpressionMetadata,omitempty"`
	Res    *types.QueryExpressionMetadataResponse `xml:"QueryExpressionMetadataResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryExpressionMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryExpressionMetadata(ctx context.Context, r soap.RoundTripper, req *types.QueryExpressionMetadata) (*types.QueryExpressionMetadataResponse, error) {
	var reqBody, resBody QueryExpressionMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryExtensionIpAllocationUsageBody struct {
	Req    *types.QueryExtensionIpAllocationUsage         `xml:"urn:vim25 QueryExtensionIpAllocationUsage,omitempty"`
	Res    *types.QueryExtensionIpAllocationUsageResponse `xml:"QueryExtensionIpAllocationUsageResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryExtensionIpAllocationUsageBody) Fault() *soap.Fault { return b.Fault_ }

func QueryExtensionIpAllocationUsage(ctx context.Context, r soap.RoundTripper, req *types.QueryExtensionIpAllocationUsage) (*types.QueryExtensionIpAllocationUsageResponse, error) {
	var reqBody, resBody QueryExtensionIpAllocationUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFaultToleranceCompatibilityBody struct {
	Req    *types.QueryFaultToleranceCompatibility         `xml:"urn:vim25 QueryFaultToleranceCompatibility,omitempty"`
	Res    *types.QueryFaultToleranceCompatibilityResponse `xml:"QueryFaultToleranceCompatibilityResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFaultToleranceCompatibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFaultToleranceCompatibility(ctx context.Context, r soap.RoundTripper, req *types.QueryFaultToleranceCompatibility) (*types.QueryFaultToleranceCompatibilityResponse, error) {
	var reqBody, resBody QueryFaultToleranceCompatibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFaultToleranceCompatibilityExBody struct {
	Req    *types.QueryFaultToleranceCompatibilityEx         `xml:"urn:vim25 QueryFaultToleranceCompatibilityEx,omitempty"`
	Res    *types.QueryFaultToleranceCompatibilityExResponse `xml:"QueryFaultToleranceCompatibilityExResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFaultToleranceCompatibilityExBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFaultToleranceCompatibilityEx(ctx context.Context, r soap.RoundTripper, req *types.QueryFaultToleranceCompatibilityEx) (*types.QueryFaultToleranceCompatibilityExResponse, error) {
	var reqBody, resBody QueryFaultToleranceCompatibilityExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFileLockInfoBody struct {
	Req    *types.QueryFileLockInfo         `xml:"urn:vim25 QueryFileLockInfo,omitempty"`
	Res    *types.QueryFileLockInfoResponse `xml:"QueryFileLockInfoResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFileLockInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFileLockInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryFileLockInfo) (*types.QueryFileLockInfoResponse, error) {
	var reqBody, resBody QueryFileLockInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFilterEntitiesBody struct {
	Req    *types.QueryFilterEntities         `xml:"urn:vim25 QueryFilterEntities,omitempty"`
	Res    *types.QueryFilterEntitiesResponse `xml:"QueryFilterEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFilterEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFilterEntities(ctx context.Context, r soap.RoundTripper, req *types.QueryFilterEntities) (*types.QueryFilterEntitiesResponse, error) {
	var reqBody, resBody QueryFilterEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFilterInfoIdsBody struct {
	Req    *types.QueryFilterInfoIds         `xml:"urn:vim25 QueryFilterInfoIds,omitempty"`
	Res    *types.QueryFilterInfoIdsResponse `xml:"QueryFilterInfoIdsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFilterInfoIdsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFilterInfoIds(ctx context.Context, r soap.RoundTripper, req *types.QueryFilterInfoIds) (*types.QueryFilterInfoIdsResponse, error) {
	var reqBody, resBody QueryFilterInfoIdsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFilterListBody struct {
	Req    *types.QueryFilterList         `xml:"urn:vim25 QueryFilterList,omitempty"`
	Res    *types.QueryFilterListResponse `xml:"QueryFilterListResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFilterListBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFilterList(ctx context.Context, r soap.RoundTripper, req *types.QueryFilterList) (*types.QueryFilterListResponse, error) {
	var reqBody, resBody QueryFilterListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFilterNameBody struct {
	Req    *types.QueryFilterName         `xml:"urn:vim25 QueryFilterName,omitempty"`
	Res    *types.QueryFilterNameResponse `xml:"QueryFilterNameResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFilterNameBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFilterName(ctx context.Context, r soap.RoundTripper, req *types.QueryFilterName) (*types.QueryFilterNameResponse, error) {
	var reqBody, resBody QueryFilterNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryFirmwareConfigUploadURLBody struct {
	Req    *types.QueryFirmwareConfigUploadURL         `xml:"urn:vim25 QueryFirmwareConfigUploadURL,omitempty"`
	Res    *types.QueryFirmwareConfigUploadURLResponse `xml:"QueryFirmwareConfigUploadURLResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFirmwareConfigUploadURLBody) Fault() *soap.Fault { return b.Fault_ }

func QueryFirmwareConfigUploadURL(ctx context.Context, r soap.RoundTripper, req *types.QueryFirmwareConfigUploadURL) (*types.QueryFirmwareConfigUploadURLResponse, error) {
	var reqBody, resBody QueryFirmwareConfigUploadURLBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHealthUpdateInfosBody struct {
	Req    *types.QueryHealthUpdateInfos         `xml:"urn:vim25 QueryHealthUpdateInfos,omitempty"`
	Res    *types.QueryHealthUpdateInfosResponse `xml:"QueryHealthUpdateInfosResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHealthUpdateInfosBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHealthUpdateInfos(ctx context.Context, r soap.RoundTripper, req *types.QueryHealthUpdateInfos) (*types.QueryHealthUpdateInfosResponse, error) {
	var reqBody, resBody QueryHealthUpdateInfosBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHealthUpdatesBody struct {
	Req    *types.QueryHealthUpdates         `xml:"urn:vim25 QueryHealthUpdates,omitempty"`
	Res    *types.QueryHealthUpdatesResponse `xml:"QueryHealthUpdatesResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHealthUpdatesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHealthUpdates(ctx context.Context, r soap.RoundTripper, req *types.QueryHealthUpdates) (*types.QueryHealthUpdatesResponse, error) {
	var reqBody, resBody QueryHealthUpdatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostConnectionInfoBody struct {
	Req    *types.QueryHostConnectionInfo         `xml:"urn:vim25 QueryHostConnectionInfo,omitempty"`
	Res    *types.QueryHostConnectionInfoResponse `xml:"QueryHostConnectionInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostConnectionInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostConnectionInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryHostConnectionInfo) (*types.QueryHostConnectionInfoResponse, error) {
	var reqBody, resBody QueryHostConnectionInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostPatch_TaskBody struct {
	Req    *types.QueryHostPatch_Task         `xml:"urn:vim25 QueryHostPatch_Task,omitempty"`
	Res    *types.QueryHostPatch_TaskResponse `xml:"QueryHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.QueryHostPatch_Task) (*types.QueryHostPatch_TaskResponse, error) {
	var reqBody, resBody QueryHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostProfileMetadataBody struct {
	Req    *types.QueryHostProfileMetadata         `xml:"urn:vim25 QueryHostProfileMetadata,omitempty"`
	Res    *types.QueryHostProfileMetadataResponse `xml:"QueryHostProfileMetadataResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostProfileMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostProfileMetadata(ctx context.Context, r soap.RoundTripper, req *types.QueryHostProfileMetadata) (*types.QueryHostProfileMetadataResponse, error) {
	var reqBody, resBody QueryHostProfileMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostStatusBody struct {
	Req    *types.QueryHostStatus         `xml:"urn:vim25 QueryHostStatus,omitempty"`
	Res    *types.QueryHostStatusResponse `xml:"QueryHostStatusResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryHostStatus) (*types.QueryHostStatusResponse, error) {
	var reqBody, resBody QueryHostStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryHostsWithAttachedLunBody struct {
	Req    *types.QueryHostsWithAttachedLun         `xml:"urn:vim25 QueryHostsWithAttachedLun,omitempty"`
	Res    *types.QueryHostsWithAttachedLunResponse `xml:"QueryHostsWithAttachedLunResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostsWithAttachedLunBody) Fault() *soap.Fault { return b.Fault_ }

func QueryHostsWithAttachedLun(ctx context.Context, r soap.RoundTripper, req *types.QueryHostsWithAttachedLun) (*types.QueryHostsWithAttachedLunResponse, error) {
	var reqBody, resBody QueryHostsWithAttachedLunBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIORMConfigOptionBody struct {
	Req    *types.QueryIORMConfigOption         `xml:"urn:vim25 QueryIORMConfigOption,omitempty"`
	Res    *types.QueryIORMConfigOptionResponse `xml:"QueryIORMConfigOptionResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIORMConfigOptionBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIORMConfigOption(ctx context.Context, r soap.RoundTripper, req *types.QueryIORMConfigOption) (*types.QueryIORMConfigOptionResponse, error) {
	var reqBody, resBody QueryIORMConfigOptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIPAllocationsBody struct {
	Req    *types.QueryIPAllocations         `xml:"urn:vim25 QueryIPAllocations,omitempty"`
	Res    *types.QueryIPAllocationsResponse `xml:"QueryIPAllocationsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIPAllocationsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIPAllocations(ctx context.Context, r soap.RoundTripper, req *types.QueryIPAllocations) (*types.QueryIPAllocationsResponse, error) {
	var reqBody, resBody QueryIPAllocationsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIoFilterInfoBody struct {
	Req    *types.QueryIoFilterInfo         `xml:"urn:vim25 QueryIoFilterInfo,omitempty"`
	Res    *types.QueryIoFilterInfoResponse `xml:"QueryIoFilterInfoResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIoFilterInfoBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIoFilterInfo(ctx context.Context, r soap.RoundTripper, req *types.QueryIoFilterInfo) (*types.QueryIoFilterInfoResponse, error) {
	var reqBody, resBody QueryIoFilterInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIoFilterIssuesBody struct {
	Req    *types.QueryIoFilterIssues         `xml:"urn:vim25 QueryIoFilterIssues,omitempty"`
	Res    *types.QueryIoFilterIssuesResponse `xml:"QueryIoFilterIssuesResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIoFilterIssuesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIoFilterIssues(ctx context.Context, r soap.RoundTripper, req *types.QueryIoFilterIssues) (*types.QueryIoFilterIssuesResponse, error) {
	var reqBody, resBody QueryIoFilterIssuesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIpPoolsBody struct {
	Req    *types.QueryIpPools         `xml:"urn:vim25 QueryIpPools,omitempty"`
	Res    *types.QueryIpPoolsResponse `xml:"QueryIpPoolsResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIpPoolsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIpPools(ctx context.Context, r soap.RoundTripper, req *types.QueryIpPools) (*types.QueryIpPoolsResponse, error) {
	var reqBody, resBody QueryIpPoolsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryLicenseSourceAvailabilityBody struct {
	Req    *types.QueryLicenseSourceAvailability         `xml:"urn:vim25 QueryLicenseSourceAvailability,omitempty"`
	Res    *types.QueryLicenseSourceAvailabilityResponse `xml:"QueryLicenseSourceAvailabilityResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLicenseSourceAvailabilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryLicenseSourceAvailability(ctx context.Context, r soap.RoundTripper, req *types.QueryLicenseSourceAvailability) (*types.QueryLicenseSourceAvailabilityResponse, error) {
	var reqBody, resBody QueryLicenseSourceAvailabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryLicenseUsageBody struct {
	Req    *types.QueryLicenseUsage         `xml:"urn:vim25 QueryLicenseUsage,omitempty"`
	Res    *types.QueryLicenseUsageResponse `xml:"QueryLicenseUsageResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLicenseUsageBody) Fault() *soap.Fault { return b.Fault_ }

func QueryLicenseUsage(ctx context.Context, r soap.RoundTripper, req *types.QueryLicenseUsage) (*types.QueryLicenseUsageResponse, error) {
	var reqBody, resBody QueryLicenseUsageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryLockdownExceptionsBody struct {
	Req    *types.QueryLockdownExceptions         `xml:"urn:vim25 QueryLockdownExceptions,omitempty"`
	Res    *types.QueryLockdownExceptionsResponse `xml:"QueryLockdownExceptionsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLockdownExceptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryLockdownExceptions(ctx context.Context, r soap.RoundTripper, req *types.QueryLockdownExceptions) (*types.QueryLockdownExceptionsResponse, error) {
	var reqBody, resBody QueryLockdownExceptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryManagedByBody struct {
	Req    *types.QueryManagedBy         `xml:"urn:vim25 QueryManagedBy,omitempty"`
	Res    *types.QueryManagedByResponse `xml:"QueryManagedByResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryManagedByBody) Fault() *soap.Fault { return b.Fault_ }

func QueryManagedBy(ctx context.Context, r soap.RoundTripper, req *types.QueryManagedBy) (*types.QueryManagedByResponse, error) {
	var reqBody, resBody QueryManagedByBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryMaxQueueDepthBody struct {
	Req    *types.QueryMaxQueueDepth         `xml:"urn:vim25 QueryMaxQueueDepth,omitempty"`
	Res    *types.QueryMaxQueueDepthResponse `xml:"QueryMaxQueueDepthResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMaxQueueDepthBody) Fault() *soap.Fault { return b.Fault_ }

func QueryMaxQueueDepth(ctx context.Context, r soap.RoundTripper, req *types.QueryMaxQueueDepth) (*types.QueryMaxQueueDepthResponse, error) {
	var reqBody, resBody QueryMaxQueueDepthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryMemoryOverheadBody struct {
	Req    *types.QueryMemoryOverhead         `xml:"urn:vim25 QueryMemoryOverhead,omitempty"`
	Res    *types.QueryMemoryOverheadResponse `xml:"QueryMemoryOverheadResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMemoryOverheadBody) Fault() *soap.Fault { return b.Fault_ }

func QueryMemoryOverhead(ctx context.Context, r soap.RoundTripper, req *types.QueryMemoryOverhead) (*types.QueryMemoryOverheadResponse, error) {
	var reqBody, resBody QueryMemoryOverheadBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryMemoryOverheadExBody struct {
	Req    *types.QueryMemoryOverheadEx         `xml:"urn:vim25 QueryMemoryOverheadEx,omitempty"`
	Res    *types.QueryMemoryOverheadExResponse `xml:"QueryMemoryOverheadExResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMemoryOverheadExBody) Fault() *soap.Fault { return b.Fault_ }

func QueryMemoryOverheadEx(ctx context.Context, r soap.RoundTripper, req *types.QueryMemoryOverheadEx) (*types.QueryMemoryOverheadExResponse, error) {
	var reqBody, resBody QueryMemoryOverheadExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryMigrationDependenciesBody struct {
	Req    *types.QueryMigrationDependencies         `xml:"urn:vim25 QueryMigrationDependencies,omitempty"`
	Res    *types.QueryMigrationDependenciesResponse `xml:"QueryMigrationDependenciesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMigrationDependenciesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryMigrationDependencies(ctx context.Context, r soap.RoundTripper, req *types.QueryMigrationDependencies) (*types.QueryMigrationDependenciesResponse, error) {
	var reqBody, resBody QueryMigrationDependenciesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryModulesBody struct {
	Req    *types.QueryModules         `xml:"urn:vim25 QueryModules,omitempty"`
	Res    *types.QueryModulesResponse `xml:"QueryModulesResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryModulesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryModules(ctx context.Context, r soap.RoundTripper, req *types.QueryModules) (*types.QueryModulesResponse, error) {
	var reqBody, resBody QueryModulesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryMonitoredEntitiesBody struct {
	Req    *types.QueryMonitoredEntities         `xml:"urn:vim25 QueryMonitoredEntities,omitempty"`
	Res    *types.QueryMonitoredEntitiesResponse `xml:"QueryMonitoredEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMonitoredEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryMonitoredEntities(ctx context.Context, r soap.RoundTripper, req *types.QueryMonitoredEntities) (*types.QueryMonitoredEntitiesResponse, error) {
	var reqBody, resBody QueryMonitoredEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryNFSUserBody struct {
	Req    *types.QueryNFSUser         `xml:"urn:vim25 QueryNFSUser,omitempty"`
	Res    *types.QueryNFSUserResponse `xml:"QueryNFSUserResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNFSUserBody) Fault() *soap.Fault { return b.Fault_ }

func QueryNFSUser(ctx context.Context, r soap.RoundTripper, req *types.QueryNFSUser) (*types.QueryNFSUserResponse, error) {
	var reqBody, resBody QueryNFSUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryNetConfigBody struct {
	Req    *types.QueryNetConfig         `xml:"urn:vim25 QueryNetConfig,omitempty"`
	Res    *types.QueryNetConfigResponse `xml:"QueryNetConfigResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNetConfigBody) Fault() *soap.Fault { return b.Fault_ }

func QueryNetConfig(ctx context.Context, r soap.RoundTripper, req *types.QueryNetConfig) (*types.QueryNetConfigResponse, error) {
	var reqBody, resBody QueryNetConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryNetworkHintBody struct {
	Req    *types.QueryNetworkHint         `xml:"urn:vim25 QueryNetworkHint,omitempty"`
	Res    *types.QueryNetworkHintResponse `xml:"QueryNetworkHintResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNetworkHintBody) Fault() *soap.Fault { return b.Fault_ }

func QueryNetworkHint(ctx context.Context, r soap.RoundTripper, req *types.QueryNetworkHint) (*types.QueryNetworkHintResponse, error) {
	var reqBody, resBody QueryNetworkHintBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryObjectsOnPhysicalVsanDiskBody struct {
	Req    *types.QueryObjectsOnPhysicalVsanDisk         `xml:"urn:vim25 QueryObjectsOnPhysicalVsanDisk,omitempty"`
	Res    *types.QueryObjectsOnPhysicalVsanDiskResponse `xml:"QueryObjectsOnPhysicalVsanDiskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryObjectsOnPhysicalVsanDiskBody) Fault() *soap.Fault { return b.Fault_ }

func QueryObjectsOnPhysicalVsanDisk(ctx context.Context, r soap.RoundTripper, req *types.QueryObjectsOnPhysicalVsanDisk) (*types.QueryObjectsOnPhysicalVsanDiskResponse, error) {
	var reqBody, resBody QueryObjectsOnPhysicalVsanDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryOptionsBody struct {
	Req    *types.QueryOptions         `xml:"urn:vim25 QueryOptions,omitempty"`
	Res    *types.QueryOptionsResponse `xml:"QueryOptionsResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryOptions) (*types.QueryOptionsResponse, error) {
	var reqBody, resBody QueryOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPartitionCreateDescBody struct {
	Req    *types.QueryPartitionCreateDesc         `xml:"urn:vim25 QueryPartitionCreateDesc,omitempty"`
	Res    *types.QueryPartitionCreateDescResponse `xml:"QueryPartitionCreateDescResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPartitionCreateDescBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPartitionCreateDesc(ctx context.Context, r soap.RoundTripper, req *types.QueryPartitionCreateDesc) (*types.QueryPartitionCreateDescResponse, error) {
	var reqBody, resBody QueryPartitionCreateDescBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPartitionCreateOptionsBody struct {
	Req    *types.QueryPartitionCreateOptions         `xml:"urn:vim25 QueryPartitionCreateOptions,omitempty"`
	Res    *types.QueryPartitionCreateOptionsResponse `xml:"QueryPartitionCreateOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPartitionCreateOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPartitionCreateOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryPartitionCreateOptions) (*types.QueryPartitionCreateOptionsResponse, error) {
	var reqBody, resBody QueryPartitionCreateOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPathSelectionPolicyOptionsBody struct {
	Req    *types.QueryPathSelectionPolicyOptions         `xml:"urn:vim25 QueryPathSelectionPolicyOptions,omitempty"`
	Res    *types.QueryPathSelectionPolicyOptionsResponse `xml:"QueryPathSelectionPolicyOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPathSelectionPolicyOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPathSelectionPolicyOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryPathSelectionPolicyOptions) (*types.QueryPathSelectionPolicyOptionsResponse, error) {
	var reqBody, resBody QueryPathSelectionPolicyOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPerfBody struct {
	Req    *types.QueryPerf         `xml:"urn:vim25 QueryPerf,omitempty"`
	Res    *types.QueryPerfResponse `xml:"QueryPerfResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPerf(ctx context.Context, r soap.RoundTripper, req *types.QueryPerf) (*types.QueryPerfResponse, error) {
	var reqBody, resBody QueryPerfBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPerfCompositeBody struct {
	Req    *types.QueryPerfComposite         `xml:"urn:vim25 QueryPerfComposite,omitempty"`
	Res    *types.QueryPerfCompositeResponse `xml:"QueryPerfCompositeResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCompositeBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPerfComposite(ctx context.Context, r soap.RoundTripper, req *types.QueryPerfComposite) (*types.QueryPerfCompositeResponse, error) {
	var reqBody, resBody QueryPerfCompositeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPerfCounterBody struct {
	Req    *types.QueryPerfCounter         `xml:"urn:vim25 QueryPerfCounter,omitempty"`
	Res    *types.QueryPerfCounterResponse `xml:"QueryPerfCounterResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCounterBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPerfCounter(ctx context.Context, r soap.RoundTripper, req *types.QueryPerfCounter) (*types.QueryPerfCounterResponse, error) {
	var reqBody, resBody QueryPerfCounterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPerfCounterByLevelBody struct {
	Req    *types.QueryPerfCounterByLevel         `xml:"urn:vim25 QueryPerfCounterByLevel,omitempty"`
	Res    *types.QueryPerfCounterByLevelResponse `xml:"QueryPerfCounterByLevelResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCounterByLevelBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPerfCounterByLevel(ctx context.Context, r soap.RoundTripper, req *types.QueryPerfCounterByLevel) (*types.QueryPerfCounterByLevelResponse, error) {
	var reqBody, resBody QueryPerfCounterByLevelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPerfProviderSummaryBody struct {
	Req    *types.QueryPerfProviderSummary         `xml:"urn:vim25 QueryPerfProviderSummary,omitempty"`
	Res    *types.QueryPerfProviderSummaryResponse `xml:"QueryPerfProviderSummaryResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfProviderSummaryBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPerfProviderSummary(ctx context.Context, r soap.RoundTripper, req *types.QueryPerfProviderSummary) (*types.QueryPerfProviderSummaryResponse, error) {
	var reqBody, resBody QueryPerfProviderSummaryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPhysicalVsanDisksBody struct {
	Req    *types.QueryPhysicalVsanDisks         `xml:"urn:vim25 QueryPhysicalVsanDisks,omitempty"`
	Res    *types.QueryPhysicalVsanDisksResponse `xml:"QueryPhysicalVsanDisksResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPhysicalVsanDisksBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPhysicalVsanDisks(ctx context.Context, r soap.RoundTripper, req *types.QueryPhysicalVsanDisks) (*types.QueryPhysicalVsanDisksResponse, error) {
	var reqBody, resBody QueryPhysicalVsanDisksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPnicStatusBody struct {
	Req    *types.QueryPnicStatus         `xml:"urn:vim25 QueryPnicStatus,omitempty"`
	Res    *types.QueryPnicStatusResponse `xml:"QueryPnicStatusResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPnicStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPnicStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryPnicStatus) (*types.QueryPnicStatusResponse, error) {
	var reqBody, resBody QueryPnicStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryPolicyMetadataBody struct {
	Req    *types.QueryPolicyMetadata         `xml:"urn:vim25 QueryPolicyMetadata,omitempty"`
	Res    *types.QueryPolicyMetadataResponse `xml:"QueryPolicyMetadataResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPolicyMetadataBody) Fault() *soap.Fault { return b.Fault_ }

func QueryPolicyMetadata(ctx context.Context, r soap.RoundTripper, req *types.QueryPolicyMetadata) (*types.QueryPolicyMetadataResponse, error) {
	var reqBody, resBody QueryPolicyMetadataBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProductLockerLocationBody struct {
	Req    *types.QueryProductLockerLocation         `xml:"urn:vim25 QueryProductLockerLocation,omitempty"`
	Res    *types.QueryProductLockerLocationResponse `xml:"QueryProductLockerLocationResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProductLockerLocationBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProductLockerLocation(ctx context.Context, r soap.RoundTripper, req *types.QueryProductLockerLocation) (*types.QueryProductLockerLocationResponse, error) {
	var reqBody, resBody QueryProductLockerLocationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProfileStructureBody struct {
	Req    *types.QueryProfileStructure         `xml:"urn:vim25 QueryProfileStructure,omitempty"`
	Res    *types.QueryProfileStructureResponse `xml:"QueryProfileStructureResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProfileStructureBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProfileStructure(ctx context.Context, r soap.RoundTripper, req *types.QueryProfileStructure) (*types.QueryProfileStructureResponse, error) {
	var reqBody, resBody QueryProfileStructureBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProviderListBody struct {
	Req    *types.QueryProviderList         `xml:"urn:vim25 QueryProviderList,omitempty"`
	Res    *types.QueryProviderListResponse `xml:"QueryProviderListResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProviderListBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProviderList(ctx context.Context, r soap.RoundTripper, req *types.QueryProviderList) (*types.QueryProviderListResponse, error) {
	var reqBody, resBody QueryProviderListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryProviderNameBody struct {
	Req    *types.QueryProviderName         `xml:"urn:vim25 QueryProviderName,omitempty"`
	Res    *types.QueryProviderNameResponse `xml:"QueryProviderNameResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProviderNameBody) Fault() *soap.Fault { return b.Fault_ }

func QueryProviderName(ctx context.Context, r soap.RoundTripper, req *types.QueryProviderName) (*types.QueryProviderNameResponse, error) {
	var reqBody, resBody QueryProviderNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryResourceConfigOptionBody struct {
	Req    *types.QueryResourceConfigOption         `xml:"urn:vim25 QueryResourceConfigOption,omitempty"`
	Res    *types.QueryResourceConfigOptionResponse `xml:"QueryResourceConfigOptionResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryResourceConfigOptionBody) Fault() *soap.Fault { return b.Fault_ }

func QueryResourceConfigOption(ctx context.Context, r soap.RoundTripper, req *types.QueryResourceConfigOption) (*types.QueryResourceConfigOptionResponse, error) {
	var reqBody, resBody QueryResourceConfigOptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryServiceListBody struct {
	Req    *types.QueryServiceList         `xml:"urn:vim25 QueryServiceList,omitempty"`
	Res    *types.QueryServiceListResponse `xml:"QueryServiceListResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryServiceListBody) Fault() *soap.Fault { return b.Fault_ }

func QueryServiceList(ctx context.Context, r soap.RoundTripper, req *types.QueryServiceList) (*types.QueryServiceListResponse, error) {
	var reqBody, resBody QueryServiceListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryStorageArrayTypePolicyOptionsBody struct {
	Req    *types.QueryStorageArrayTypePolicyOptions         `xml:"urn:vim25 QueryStorageArrayTypePolicyOptions,omitempty"`
	Res    *types.QueryStorageArrayTypePolicyOptionsResponse `xml:"QueryStorageArrayTypePolicyOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryStorageArrayTypePolicyOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryStorageArrayTypePolicyOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryStorageArrayTypePolicyOptions) (*types.QueryStorageArrayTypePolicyOptionsResponse, error) {
	var reqBody, resBody QueryStorageArrayTypePolicyOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySupportedFeaturesBody struct {
	Req    *types.QuerySupportedFeatures         `xml:"urn:vim25 QuerySupportedFeatures,omitempty"`
	Res    *types.QuerySupportedFeaturesResponse `xml:"QuerySupportedFeaturesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySupportedFeaturesBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySupportedFeatures(ctx context.Context, r soap.RoundTripper, req *types.QuerySupportedFeatures) (*types.QuerySupportedFeaturesResponse, error) {
	var reqBody, resBody QuerySupportedFeaturesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySupportedNetworkOffloadSpecBody struct {
	Req    *types.QuerySupportedNetworkOffloadSpec         `xml:"urn:vim25 QuerySupportedNetworkOffloadSpec,omitempty"`
	Res    *types.QuerySupportedNetworkOffloadSpecResponse `xml:"QuerySupportedNetworkOffloadSpecResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySupportedNetworkOffloadSpecBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySupportedNetworkOffloadSpec(ctx context.Context, r soap.RoundTripper, req *types.QuerySupportedNetworkOffloadSpec) (*types.QuerySupportedNetworkOffloadSpecResponse, error) {
	var reqBody, resBody QuerySupportedNetworkOffloadSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySyncingVsanObjectsBody struct {
	Req    *types.QuerySyncingVsanObjects         `xml:"urn:vim25 QuerySyncingVsanObjects,omitempty"`
	Res    *types.QuerySyncingVsanObjectsResponse `xml:"QuerySyncingVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySyncingVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySyncingVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.QuerySyncingVsanObjects) (*types.QuerySyncingVsanObjectsResponse, error) {
	var reqBody, resBody QuerySyncingVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySystemUsersBody struct {
	Req    *types.QuerySystemUsers         `xml:"urn:vim25 QuerySystemUsers,omitempty"`
	Res    *types.QuerySystemUsersResponse `xml:"QuerySystemUsersResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySystemUsersBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySystemUsers(ctx context.Context, r soap.RoundTripper, req *types.QuerySystemUsers) (*types.QuerySystemUsersResponse, error) {
	var reqBody, resBody QuerySystemUsersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryTargetCapabilitiesBody struct {
	Req    *types.QueryTargetCapabilities         `xml:"urn:vim25 QueryTargetCapabilities,omitempty"`
	Res    *types.QueryTargetCapabilitiesResponse `xml:"QueryTargetCapabilitiesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryTargetCapabilitiesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryTargetCapabilities(ctx context.Context, r soap.RoundTripper, req *types.QueryTargetCapabilities) (*types.QueryTargetCapabilitiesResponse, error) {
	var reqBody, resBody QueryTargetCapabilitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryTpmAttestationReportBody struct {
	Req    *types.QueryTpmAttestationReport         `xml:"urn:vim25 QueryTpmAttestationReport,omitempty"`
	Res    *types.QueryTpmAttestationReportResponse `xml:"QueryTpmAttestationReportResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryTpmAttestationReportBody) Fault() *soap.Fault { return b.Fault_ }

func QueryTpmAttestationReport(ctx context.Context, r soap.RoundTripper, req *types.QueryTpmAttestationReport) (*types.QueryTpmAttestationReportResponse, error) {
	var reqBody, resBody QueryTpmAttestationReportBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryUnmonitoredHostsBody struct {
	Req    *types.QueryUnmonitoredHosts         `xml:"urn:vim25 QueryUnmonitoredHosts,omitempty"`
	Res    *types.QueryUnmonitoredHostsResponse `xml:"QueryUnmonitoredHostsResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnmonitoredHostsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryUnmonitoredHosts(ctx context.Context, r soap.RoundTripper, req *types.QueryUnmonitoredHosts) (*types.QueryUnmonitoredHostsResponse, error) {
	var reqBody, resBody QueryUnmonitoredHostsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryUnownedFilesBody struct {
	Req    *types.QueryUnownedFiles         `xml:"urn:vim25 QueryUnownedFiles,omitempty"`
	Res    *types.QueryUnownedFilesResponse `xml:"QueryUnownedFilesResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnownedFilesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryUnownedFiles(ctx context.Context, r soap.RoundTripper, req *types.QueryUnownedFiles) (*types.QueryUnownedFilesResponse, error) {
	var reqBody, resBody QueryUnownedFilesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryUnresolvedVmfsVolumeBody struct {
	Req    *types.QueryUnresolvedVmfsVolume         `xml:"urn:vim25 QueryUnresolvedVmfsVolume,omitempty"`
	Res    *types.QueryUnresolvedVmfsVolumeResponse `xml:"QueryUnresolvedVmfsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnresolvedVmfsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func QueryUnresolvedVmfsVolume(ctx context.Context, r soap.RoundTripper, req *types.QueryUnresolvedVmfsVolume) (*types.QueryUnresolvedVmfsVolumeResponse, error) {
	var reqBody, resBody QueryUnresolvedVmfsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryUnresolvedVmfsVolumesBody struct {
	Req    *types.QueryUnresolvedVmfsVolumes         `xml:"urn:vim25 QueryUnresolvedVmfsVolumes,omitempty"`
	Res    *types.QueryUnresolvedVmfsVolumesResponse `xml:"QueryUnresolvedVmfsVolumesResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnresolvedVmfsVolumesBody) Fault() *soap.Fault { return b.Fault_ }

func QueryUnresolvedVmfsVolumes(ctx context.Context, r soap.RoundTripper, req *types.QueryUnresolvedVmfsVolumes) (*types.QueryUnresolvedVmfsVolumesResponse, error) {
	var reqBody, resBody QueryUnresolvedVmfsVolumesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryUsedVlanIdInDvsBody struct {
	Req    *types.QueryUsedVlanIdInDvs         `xml:"urn:vim25 QueryUsedVlanIdInDvs,omitempty"`
	Res    *types.QueryUsedVlanIdInDvsResponse `xml:"QueryUsedVlanIdInDvsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUsedVlanIdInDvsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryUsedVlanIdInDvs(ctx context.Context, r soap.RoundTripper, req *types.QueryUsedVlanIdInDvs) (*types.QueryUsedVlanIdInDvsResponse, error) {
	var reqBody, resBody QueryUsedVlanIdInDvsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVMotionCompatibilityBody struct {
	Req    *types.QueryVMotionCompatibility         `xml:"urn:vim25 QueryVMotionCompatibility,omitempty"`
	Res    *types.QueryVMotionCompatibilityResponse `xml:"QueryVMotionCompatibilityResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVMotionCompatibilityBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVMotionCompatibility(ctx context.Context, r soap.RoundTripper, req *types.QueryVMotionCompatibility) (*types.QueryVMotionCompatibilityResponse, error) {
	var reqBody, resBody QueryVMotionCompatibilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVMotionCompatibilityEx_TaskBody struct {
	Req    *types.QueryVMotionCompatibilityEx_Task         `xml:"urn:vim25 QueryVMotionCompatibilityEx_Task,omitempty"`
	Res    *types.QueryVMotionCompatibilityEx_TaskResponse `xml:"QueryVMotionCompatibilityEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVMotionCompatibilityEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVMotionCompatibilityEx_Task(ctx context.Context, r soap.RoundTripper, req *types.QueryVMotionCompatibilityEx_Task) (*types.QueryVMotionCompatibilityEx_TaskResponse, error) {
	var reqBody, resBody QueryVMotionCompatibilityEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVirtualDiskFragmentationBody struct {
	Req    *types.QueryVirtualDiskFragmentation         `xml:"urn:vim25 QueryVirtualDiskFragmentation,omitempty"`
	Res    *types.QueryVirtualDiskFragmentationResponse `xml:"QueryVirtualDiskFragmentationResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskFragmentationBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVirtualDiskFragmentation(ctx context.Context, r soap.RoundTripper, req *types.QueryVirtualDiskFragmentation) (*types.QueryVirtualDiskFragmentationResponse, error) {
	var reqBody, resBody QueryVirtualDiskFragmentationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVirtualDiskGeometryBody struct {
	Req    *types.QueryVirtualDiskGeometry         `xml:"urn:vim25 QueryVirtualDiskGeometry,omitempty"`
	Res    *types.QueryVirtualDiskGeometryResponse `xml:"QueryVirtualDiskGeometryResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskGeometryBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVirtualDiskGeometry(ctx context.Context, r soap.RoundTripper, req *types.QueryVirtualDiskGeometry) (*types.QueryVirtualDiskGeometryResponse, error) {
	var reqBody, resBody QueryVirtualDiskGeometryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVirtualDiskUuidBody struct {
	Req    *types.QueryVirtualDiskUuid         `xml:"urn:vim25 QueryVirtualDiskUuid,omitempty"`
	Res    *types.QueryVirtualDiskUuidResponse `xml:"QueryVirtualDiskUuidResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskUuidBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVirtualDiskUuid(ctx context.Context, r soap.RoundTripper, req *types.QueryVirtualDiskUuid) (*types.QueryVirtualDiskUuidResponse, error) {
	var reqBody, resBody QueryVirtualDiskUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVirtualDiskUuidExBody struct {
	Req    *types.QueryVirtualDiskUuidEx         `xml:"urn:vim25 QueryVirtualDiskUuidEx,omitempty"`
	Res    *types.QueryVirtualDiskUuidExResponse `xml:"QueryVirtualDiskUuidExResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskUuidExBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVirtualDiskUuidEx(ctx context.Context, r soap.RoundTripper, req *types.QueryVirtualDiskUuidEx) (*types.QueryVirtualDiskUuidExResponse, error) {
	var reqBody, resBody QueryVirtualDiskUuidExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmfsConfigOptionBody struct {
	Req    *types.QueryVmfsConfigOption         `xml:"urn:vim25 QueryVmfsConfigOption,omitempty"`
	Res    *types.QueryVmfsConfigOptionResponse `xml:"QueryVmfsConfigOptionResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsConfigOptionBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmfsConfigOption(ctx context.Context, r soap.RoundTripper, req *types.QueryVmfsConfigOption) (*types.QueryVmfsConfigOptionResponse, error) {
	var reqBody, resBody QueryVmfsConfigOptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmfsDatastoreCreateOptionsBody struct {
	Req    *types.QueryVmfsDatastoreCreateOptions         `xml:"urn:vim25 QueryVmfsDatastoreCreateOptions,omitempty"`
	Res    *types.QueryVmfsDatastoreCreateOptionsResponse `xml:"QueryVmfsDatastoreCreateOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreCreateOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmfsDatastoreCreateOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryVmfsDatastoreCreateOptions) (*types.QueryVmfsDatastoreCreateOptionsResponse, error) {
	var reqBody, resBody QueryVmfsDatastoreCreateOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmfsDatastoreExpandOptionsBody struct {
	Req    *types.QueryVmfsDatastoreExpandOptions         `xml:"urn:vim25 QueryVmfsDatastoreExpandOptions,omitempty"`
	Res    *types.QueryVmfsDatastoreExpandOptionsResponse `xml:"QueryVmfsDatastoreExpandOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreExpandOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmfsDatastoreExpandOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryVmfsDatastoreExpandOptions) (*types.QueryVmfsDatastoreExpandOptionsResponse, error) {
	var reqBody, resBody QueryVmfsDatastoreExpandOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVmfsDatastoreExtendOptionsBody struct {
	Req    *types.QueryVmfsDatastoreExtendOptions         `xml:"urn:vim25 QueryVmfsDatastoreExtendOptions,omitempty"`
	Res    *types.QueryVmfsDatastoreExtendOptionsResponse `xml:"QueryVmfsDatastoreExtendOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreExtendOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVmfsDatastoreExtendOptions(ctx context.Context, r soap.RoundTripper, req *types.QueryVmfsDatastoreExtendOptions) (*types.QueryVmfsDatastoreExtendOptionsResponse, error) {
	var reqBody, resBody QueryVmfsDatastoreExtendOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVnicStatusBody struct {
	Req    *types.QueryVnicStatus         `xml:"urn:vim25 QueryVnicStatus,omitempty"`
	Res    *types.QueryVnicStatusResponse `xml:"QueryVnicStatusResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVnicStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVnicStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryVnicStatus) (*types.QueryVnicStatusResponse, error) {
	var reqBody, resBody QueryVnicStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanObjectUuidsByFilterBody struct {
	Req    *types.QueryVsanObjectUuidsByFilter         `xml:"urn:vim25 QueryVsanObjectUuidsByFilter,omitempty"`
	Res    *types.QueryVsanObjectUuidsByFilterResponse `xml:"QueryVsanObjectUuidsByFilterResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanObjectUuidsByFilterBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanObjectUuidsByFilter(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanObjectUuidsByFilter) (*types.QueryVsanObjectUuidsByFilterResponse, error) {
	var reqBody, resBody QueryVsanObjectUuidsByFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanObjectsBody struct {
	Req    *types.QueryVsanObjects         `xml:"urn:vim25 QueryVsanObjects,omitempty"`
	Res    *types.QueryVsanObjectsResponse `xml:"QueryVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanObjects) (*types.QueryVsanObjectsResponse, error) {
	var reqBody, resBody QueryVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanStatisticsBody struct {
	Req    *types.QueryVsanStatistics         `xml:"urn:vim25 QueryVsanStatistics,omitempty"`
	Res    *types.QueryVsanStatisticsResponse `xml:"QueryVsanStatisticsResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanStatisticsBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanStatistics(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanStatistics) (*types.QueryVsanStatisticsResponse, error) {
	var reqBody, resBody QueryVsanStatisticsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryVsanUpgradeStatusBody struct {
	Req    *types.QueryVsanUpgradeStatus         `xml:"urn:vim25 QueryVsanUpgradeStatus,omitempty"`
	Res    *types.QueryVsanUpgradeStatusResponse `xml:"QueryVsanUpgradeStatusResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanUpgradeStatusBody) Fault() *soap.Fault { return b.Fault_ }

func QueryVsanUpgradeStatus(ctx context.Context, r soap.RoundTripper, req *types.QueryVsanUpgradeStatus) (*types.QueryVsanUpgradeStatusResponse, error) {
	var reqBody, resBody QueryVsanUpgradeStatusBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadEnvironmentVariableInGuestBody struct {
	Req    *types.ReadEnvironmentVariableInGuest         `xml:"urn:vim25 ReadEnvironmentVariableInGuest,omitempty"`
	Res    *types.ReadEnvironmentVariableInGuestResponse `xml:"ReadEnvironmentVariableInGuestResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadEnvironmentVariableInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ReadEnvironmentVariableInGuest(ctx context.Context, r soap.RoundTripper, req *types.ReadEnvironmentVariableInGuest) (*types.ReadEnvironmentVariableInGuestResponse, error) {
	var reqBody, resBody ReadEnvironmentVariableInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadNextEventsBody struct {
	Req    *types.ReadNextEvents         `xml:"urn:vim25 ReadNextEvents,omitempty"`
	Res    *types.ReadNextEventsResponse `xml:"ReadNextEventsResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadNextEventsBody) Fault() *soap.Fault { return b.Fault_ }

func ReadNextEvents(ctx context.Context, r soap.RoundTripper, req *types.ReadNextEvents) (*types.ReadNextEventsResponse, error) {
	var reqBody, resBody ReadNextEventsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadNextTasksBody struct {
	Req    *types.ReadNextTasks         `xml:"urn:vim25 ReadNextTasks,omitempty"`
	Res    *types.ReadNextTasksResponse `xml:"ReadNextTasksResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadNextTasksBody) Fault() *soap.Fault { return b.Fault_ }

func ReadNextTasks(ctx context.Context, r soap.RoundTripper, req *types.ReadNextTasks) (*types.ReadNextTasksResponse, error) {
	var reqBody, resBody ReadNextTasksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadNextTasksByViewSpecBody struct {
	Req    *types.ReadNextTasksByViewSpec         `xml:"urn:vim25 ReadNextTasksByViewSpec,omitempty"`
	Res    *types.ReadNextTasksByViewSpecResponse `xml:"ReadNextTasksByViewSpecResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadNextTasksByViewSpecBody) Fault() *soap.Fault { return b.Fault_ }

func ReadNextTasksByViewSpec(ctx context.Context, r soap.RoundTripper, req *types.ReadNextTasksByViewSpec) (*types.ReadNextTasksByViewSpecResponse, error) {
	var reqBody, resBody ReadNextTasksByViewSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadPreviousEventsBody struct {
	Req    *types.ReadPreviousEvents         `xml:"urn:vim25 ReadPreviousEvents,omitempty"`
	Res    *types.ReadPreviousEventsResponse `xml:"ReadPreviousEventsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadPreviousEventsBody) Fault() *soap.Fault { return b.Fault_ }

func ReadPreviousEvents(ctx context.Context, r soap.RoundTripper, req *types.ReadPreviousEvents) (*types.ReadPreviousEventsResponse, error) {
	var reqBody, resBody ReadPreviousEventsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReadPreviousTasksBody struct {
	Req    *types.ReadPreviousTasks         `xml:"urn:vim25 ReadPreviousTasks,omitempty"`
	Res    *types.ReadPreviousTasksResponse `xml:"ReadPreviousTasksResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadPreviousTasksBody) Fault() *soap.Fault { return b.Fault_ }

func ReadPreviousTasks(ctx context.Context, r soap.RoundTripper, req *types.ReadPreviousTasks) (*types.ReadPreviousTasksResponse, error) {
	var reqBody, resBody ReadPreviousTasksBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RebootGuestBody struct {
	Req    *types.RebootGuest         `xml:"urn:vim25 RebootGuest,omitempty"`
	Res    *types.RebootGuestResponse `xml:"RebootGuestResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebootGuestBody) Fault() *soap.Fault { return b.Fault_ }

func RebootGuest(ctx context.Context, r soap.RoundTripper, req *types.RebootGuest) (*types.RebootGuestResponse, error) {
	var reqBody, resBody RebootGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RebootHost_TaskBody struct {
	Req    *types.RebootHost_Task         `xml:"urn:vim25 RebootHost_Task,omitempty"`
	Res    *types.RebootHost_TaskResponse `xml:"RebootHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebootHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RebootHost_Task(ctx context.Context, r soap.RoundTripper, req *types.RebootHost_Task) (*types.RebootHost_TaskResponse, error) {
	var reqBody, resBody RebootHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RecommendDatastoresBody struct {
	Req    *types.RecommendDatastores         `xml:"urn:vim25 RecommendDatastores,omitempty"`
	Res    *types.RecommendDatastoresResponse `xml:"RecommendDatastoresResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RecommendDatastoresBody) Fault() *soap.Fault { return b.Fault_ }

func RecommendDatastores(ctx context.Context, r soap.RoundTripper, req *types.RecommendDatastores) (*types.RecommendDatastoresResponse, error) {
	var reqBody, resBody RecommendDatastoresBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RecommendHostsForVmBody struct {
	Req    *types.RecommendHostsForVm         `xml:"urn:vim25 RecommendHostsForVm,omitempty"`
	Res    *types.RecommendHostsForVmResponse `xml:"RecommendHostsForVmResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RecommendHostsForVmBody) Fault() *soap.Fault { return b.Fault_ }

func RecommendHostsForVm(ctx context.Context, r soap.RoundTripper, req *types.RecommendHostsForVm) (*types.RecommendHostsForVmResponse, error) {
	var reqBody, resBody RecommendHostsForVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RecommissionVsanNode_TaskBody struct {
	Req    *types.RecommissionVsanNode_Task         `xml:"urn:vim25 RecommissionVsanNode_Task,omitempty"`
	Res    *types.RecommissionVsanNode_TaskResponse `xml:"RecommissionVsanNode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RecommissionVsanNode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RecommissionVsanNode_Task(ctx context.Context, r soap.RoundTripper, req *types.RecommissionVsanNode_Task) (*types.RecommissionVsanNode_TaskResponse, error) {
	var reqBody, resBody RecommissionVsanNode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconcileDatastoreInventoryEx_TaskBody struct {
	Req    *types.ReconcileDatastoreInventoryEx_Task         `xml:"urn:vim25 ReconcileDatastoreInventoryEx_Task,omitempty"`
	Res    *types.ReconcileDatastoreInventoryEx_TaskResponse `xml:"ReconcileDatastoreInventoryEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconcileDatastoreInventoryEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconcileDatastoreInventoryEx_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconcileDatastoreInventoryEx_Task) (*types.ReconcileDatastoreInventoryEx_TaskResponse, error) {
	var reqBody, resBody ReconcileDatastoreInventoryEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconcileDatastoreInventory_TaskBody struct {
	Req    *types.ReconcileDatastoreInventory_Task         `xml:"urn:vim25 ReconcileDatastoreInventory_Task,omitempty"`
	Res    *types.ReconcileDatastoreInventory_TaskResponse `xml:"ReconcileDatastoreInventory_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconcileDatastoreInventory_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconcileDatastoreInventory_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconcileDatastoreInventory_Task) (*types.ReconcileDatastoreInventory_TaskResponse, error) {
	var reqBody, resBody ReconcileDatastoreInventory_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigVM_TaskBody struct {
	Req    *types.ReconfigVM_Task         `xml:"urn:vim25 ReconfigVM_Task,omitempty"`
	Res    *types.ReconfigVM_TaskResponse `xml:"ReconfigVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigVM_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigVM_Task) (*types.ReconfigVM_TaskResponse, error) {
	var reqBody, resBody ReconfigVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigurationSatisfiableBody struct {
	Req    *types.ReconfigurationSatisfiable         `xml:"urn:vim25 ReconfigurationSatisfiable,omitempty"`
	Res    *types.ReconfigurationSatisfiableResponse `xml:"ReconfigurationSatisfiableResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigurationSatisfiableBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigurationSatisfiable(ctx context.Context, r soap.RoundTripper, req *types.ReconfigurationSatisfiable) (*types.ReconfigurationSatisfiableResponse, error) {
	var reqBody, resBody ReconfigurationSatisfiableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureAlarmBody struct {
	Req    *types.ReconfigureAlarm         `xml:"urn:vim25 ReconfigureAlarm,omitempty"`
	Res    *types.ReconfigureAlarmResponse `xml:"ReconfigureAlarmResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureAlarm(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureAlarm) (*types.ReconfigureAlarmResponse, error) {
	var reqBody, resBody ReconfigureAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureAutostartBody struct {
	Req    *types.ReconfigureAutostart         `xml:"urn:vim25 ReconfigureAutostart,omitempty"`
	Res    *types.ReconfigureAutostartResponse `xml:"ReconfigureAutostartResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureAutostartBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureAutostart(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureAutostart) (*types.ReconfigureAutostartResponse, error) {
	var reqBody, resBody ReconfigureAutostartBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureCluster_TaskBody struct {
	Req    *types.ReconfigureCluster_Task         `xml:"urn:vim25 ReconfigureCluster_Task,omitempty"`
	Res    *types.ReconfigureCluster_TaskResponse `xml:"ReconfigureCluster_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureCluster_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureCluster_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureCluster_Task) (*types.ReconfigureCluster_TaskResponse, error) {
	var reqBody, resBody ReconfigureCluster_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureComputeResource_TaskBody struct {
	Req    *types.ReconfigureComputeResource_Task         `xml:"urn:vim25 ReconfigureComputeResource_Task,omitempty"`
	Res    *types.ReconfigureComputeResource_TaskResponse `xml:"ReconfigureComputeResource_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureComputeResource_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureComputeResource_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureComputeResource_Task) (*types.ReconfigureComputeResource_TaskResponse, error) {
	var reqBody, resBody ReconfigureComputeResource_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureDVPort_TaskBody struct {
	Req    *types.ReconfigureDVPort_Task         `xml:"urn:vim25 ReconfigureDVPort_Task,omitempty"`
	Res    *types.ReconfigureDVPort_TaskResponse `xml:"ReconfigureDVPort_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDVPort_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureDVPort_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureDVPort_Task) (*types.ReconfigureDVPort_TaskResponse, error) {
	var reqBody, resBody ReconfigureDVPort_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureDVPortgroup_TaskBody struct {
	Req    *types.ReconfigureDVPortgroup_Task         `xml:"urn:vim25 ReconfigureDVPortgroup_Task,omitempty"`
	Res    *types.ReconfigureDVPortgroup_TaskResponse `xml:"ReconfigureDVPortgroup_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDVPortgroup_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureDVPortgroup_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureDVPortgroup_Task) (*types.ReconfigureDVPortgroup_TaskResponse, error) {
	var reqBody, resBody ReconfigureDVPortgroup_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureDatacenter_TaskBody struct {
	Req    *types.ReconfigureDatacenter_Task         `xml:"urn:vim25 ReconfigureDatacenter_Task,omitempty"`
	Res    *types.ReconfigureDatacenter_TaskResponse `xml:"ReconfigureDatacenter_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDatacenter_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureDatacenter_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureDatacenter_Task) (*types.ReconfigureDatacenter_TaskResponse, error) {
	var reqBody, resBody ReconfigureDatacenter_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureDomObjectBody struct {
	Req    *types.ReconfigureDomObject         `xml:"urn:vim25 ReconfigureDomObject,omitempty"`
	Res    *types.ReconfigureDomObjectResponse `xml:"ReconfigureDomObjectResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDomObjectBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureDomObject(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureDomObject) (*types.ReconfigureDomObjectResponse, error) {
	var reqBody, resBody ReconfigureDomObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureDvs_TaskBody struct {
	Req    *types.ReconfigureDvs_Task         `xml:"urn:vim25 ReconfigureDvs_Task,omitempty"`
	Res    *types.ReconfigureDvs_TaskResponse `xml:"ReconfigureDvs_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDvs_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureDvs_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureDvs_Task) (*types.ReconfigureDvs_TaskResponse, error) {
	var reqBody, resBody ReconfigureDvs_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureHostForDAS_TaskBody struct {
	Req    *types.ReconfigureHostForDAS_Task         `xml:"urn:vim25 ReconfigureHostForDAS_Task,omitempty"`
	Res    *types.ReconfigureHostForDAS_TaskResponse `xml:"ReconfigureHostForDAS_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureHostForDAS_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureHostForDAS_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureHostForDAS_Task) (*types.ReconfigureHostForDAS_TaskResponse, error) {
	var reqBody, resBody ReconfigureHostForDAS_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureScheduledTaskBody struct {
	Req    *types.ReconfigureScheduledTask         `xml:"urn:vim25 ReconfigureScheduledTask,omitempty"`
	Res    *types.ReconfigureScheduledTaskResponse `xml:"ReconfigureScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureScheduledTask) (*types.ReconfigureScheduledTaskResponse, error) {
	var reqBody, resBody ReconfigureScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureServiceConsoleReservationBody struct {
	Req    *types.ReconfigureServiceConsoleReservation         `xml:"urn:vim25 ReconfigureServiceConsoleReservation,omitempty"`
	Res    *types.ReconfigureServiceConsoleReservationResponse `xml:"ReconfigureServiceConsoleReservationResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureServiceConsoleReservationBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureServiceConsoleReservation(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureServiceConsoleReservation) (*types.ReconfigureServiceConsoleReservationResponse, error) {
	var reqBody, resBody ReconfigureServiceConsoleReservationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureSnmpAgentBody struct {
	Req    *types.ReconfigureSnmpAgent         `xml:"urn:vim25 ReconfigureSnmpAgent,omitempty"`
	Res    *types.ReconfigureSnmpAgentResponse `xml:"ReconfigureSnmpAgentResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureSnmpAgentBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureSnmpAgent(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureSnmpAgent) (*types.ReconfigureSnmpAgentResponse, error) {
	var reqBody, resBody ReconfigureSnmpAgentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconfigureVirtualMachineReservationBody struct {
	Req    *types.ReconfigureVirtualMachineReservation         `xml:"urn:vim25 ReconfigureVirtualMachineReservation,omitempty"`
	Res    *types.ReconfigureVirtualMachineReservationResponse `xml:"ReconfigureVirtualMachineReservationResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureVirtualMachineReservationBody) Fault() *soap.Fault { return b.Fault_ }

func ReconfigureVirtualMachineReservation(ctx context.Context, r soap.RoundTripper, req *types.ReconfigureVirtualMachineReservation) (*types.ReconfigureVirtualMachineReservationResponse, error) {
	var reqBody, resBody ReconfigureVirtualMachineReservationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReconnectHost_TaskBody struct {
	Req    *types.ReconnectHost_Task         `xml:"urn:vim25 ReconnectHost_Task,omitempty"`
	Res    *types.ReconnectHost_TaskResponse `xml:"ReconnectHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconnectHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReconnectHost_Task(ctx context.Context, r soap.RoundTripper, req *types.ReconnectHost_Task) (*types.ReconnectHost_TaskResponse, error) {
	var reqBody, resBody ReconnectHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RectifyDvsHost_TaskBody struct {
	Req    *types.RectifyDvsHost_Task         `xml:"urn:vim25 RectifyDvsHost_Task,omitempty"`
	Res    *types.RectifyDvsHost_TaskResponse `xml:"RectifyDvsHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RectifyDvsHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RectifyDvsHost_Task(ctx context.Context, r soap.RoundTripper, req *types.RectifyDvsHost_Task) (*types.RectifyDvsHost_TaskResponse, error) {
	var reqBody, resBody RectifyDvsHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RectifyDvsOnHost_TaskBody struct {
	Req    *types.RectifyDvsOnHost_Task         `xml:"urn:vim25 RectifyDvsOnHost_Task,omitempty"`
	Res    *types.RectifyDvsOnHost_TaskResponse `xml:"RectifyDvsOnHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RectifyDvsOnHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RectifyDvsOnHost_Task(ctx context.Context, r soap.RoundTripper, req *types.RectifyDvsOnHost_Task) (*types.RectifyDvsOnHost_TaskResponse, error) {
	var reqBody, resBody RectifyDvsOnHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshBody struct {
	Req    *types.Refresh         `xml:"urn:vim25 Refresh,omitempty"`
	Res    *types.RefreshResponse `xml:"RefreshResponse,omitempty"`
	Fault_ *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshBody) Fault() *soap.Fault { return b.Fault_ }

func Refresh(ctx context.Context, r soap.RoundTripper, req *types.Refresh) (*types.RefreshResponse, error) {
	var reqBody, resBody RefreshBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshDVPortStateBody struct {
	Req    *types.RefreshDVPortState         `xml:"urn:vim25 RefreshDVPortState,omitempty"`
	Res    *types.RefreshDVPortStateResponse `xml:"RefreshDVPortStateResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDVPortStateBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshDVPortState(ctx context.Context, r soap.RoundTripper, req *types.RefreshDVPortState) (*types.RefreshDVPortStateResponse, error) {
	var reqBody, resBody RefreshDVPortStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshDatastoreBody struct {
	Req    *types.RefreshDatastore         `xml:"urn:vim25 RefreshDatastore,omitempty"`
	Res    *types.RefreshDatastoreResponse `xml:"RefreshDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshDatastore(ctx context.Context, r soap.RoundTripper, req *types.RefreshDatastore) (*types.RefreshDatastoreResponse, error) {
	var reqBody, resBody RefreshDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshDatastoreStorageInfoBody struct {
	Req    *types.RefreshDatastoreStorageInfo         `xml:"urn:vim25 RefreshDatastoreStorageInfo,omitempty"`
	Res    *types.RefreshDatastoreStorageInfoResponse `xml:"RefreshDatastoreStorageInfoResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDatastoreStorageInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshDatastoreStorageInfo(ctx context.Context, r soap.RoundTripper, req *types.RefreshDatastoreStorageInfo) (*types.RefreshDatastoreStorageInfoResponse, error) {
	var reqBody, resBody RefreshDatastoreStorageInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshDateTimeSystemBody struct {
	Req    *types.RefreshDateTimeSystem         `xml:"urn:vim25 RefreshDateTimeSystem,omitempty"`
	Res    *types.RefreshDateTimeSystemResponse `xml:"RefreshDateTimeSystemResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDateTimeSystemBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshDateTimeSystem(ctx context.Context, r soap.RoundTripper, req *types.RefreshDateTimeSystem) (*types.RefreshDateTimeSystemResponse, error) {
	var reqBody, resBody RefreshDateTimeSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshFirewallBody struct {
	Req    *types.RefreshFirewall         `xml:"urn:vim25 RefreshFirewall,omitempty"`
	Res    *types.RefreshFirewallResponse `xml:"RefreshFirewallResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshFirewallBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshFirewall(ctx context.Context, r soap.RoundTripper, req *types.RefreshFirewall) (*types.RefreshFirewallResponse, error) {
	var reqBody, resBody RefreshFirewallBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshGraphicsManagerBody struct {
	Req    *types.RefreshGraphicsManager         `xml:"urn:vim25 RefreshGraphicsManager,omitempty"`
	Res    *types.RefreshGraphicsManagerResponse `xml:"RefreshGraphicsManagerResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshGraphicsManagerBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshGraphicsManager(ctx context.Context, r soap.RoundTripper, req *types.RefreshGraphicsManager) (*types.RefreshGraphicsManagerResponse, error) {
	var reqBody, resBody RefreshGraphicsManagerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshHealthStatusSystemBody struct {
	Req    *types.RefreshHealthStatusSystem         `xml:"urn:vim25 RefreshHealthStatusSystem,omitempty"`
	Res    *types.RefreshHealthStatusSystemResponse `xml:"RefreshHealthStatusSystemResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshHealthStatusSystemBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshHealthStatusSystem(ctx context.Context, r soap.RoundTripper, req *types.RefreshHealthStatusSystem) (*types.RefreshHealthStatusSystemResponse, error) {
	var reqBody, resBody RefreshHealthStatusSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshNetworkSystemBody struct {
	Req    *types.RefreshNetworkSystem         `xml:"urn:vim25 RefreshNetworkSystem,omitempty"`
	Res    *types.RefreshNetworkSystemResponse `xml:"RefreshNetworkSystemResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshNetworkSystemBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshNetworkSystem(ctx context.Context, r soap.RoundTripper, req *types.RefreshNetworkSystem) (*types.RefreshNetworkSystemResponse, error) {
	var reqBody, resBody RefreshNetworkSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshRecommendationBody struct {
	Req    *types.RefreshRecommendation         `xml:"urn:vim25 RefreshRecommendation,omitempty"`
	Res    *types.RefreshRecommendationResponse `xml:"RefreshRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshRecommendation(ctx context.Context, r soap.RoundTripper, req *types.RefreshRecommendation) (*types.RefreshRecommendationResponse, error) {
	var reqBody, resBody RefreshRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshRuntimeBody struct {
	Req    *types.RefreshRuntime         `xml:"urn:vim25 RefreshRuntime,omitempty"`
	Res    *types.RefreshRuntimeResponse `xml:"RefreshRuntimeResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshRuntimeBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshRuntime(ctx context.Context, r soap.RoundTripper, req *types.RefreshRuntime) (*types.RefreshRuntimeResponse, error) {
	var reqBody, resBody RefreshRuntimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshServicesBody struct {
	Req    *types.RefreshServices         `xml:"urn:vim25 RefreshServices,omitempty"`
	Res    *types.RefreshServicesResponse `xml:"RefreshServicesResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshServicesBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshServices(ctx context.Context, r soap.RoundTripper, req *types.RefreshServices) (*types.RefreshServicesResponse, error) {
	var reqBody, resBody RefreshServicesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshStorageDrsRecommendationBody struct {
	Req    *types.RefreshStorageDrsRecommendation         `xml:"urn:vim25 RefreshStorageDrsRecommendation,omitempty"`
	Res    *types.RefreshStorageDrsRecommendationResponse `xml:"RefreshStorageDrsRecommendationResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageDrsRecommendationBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshStorageDrsRecommendation(ctx context.Context, r soap.RoundTripper, req *types.RefreshStorageDrsRecommendation) (*types.RefreshStorageDrsRecommendationResponse, error) {
	var reqBody, resBody RefreshStorageDrsRecommendationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshStorageDrsRecommendationsForPod_TaskBody struct {
	Req    *types.RefreshStorageDrsRecommendationsForPod_Task         `xml:"urn:vim25 RefreshStorageDrsRecommendationsForPod_Task,omitempty"`
	Res    *types.RefreshStorageDrsRecommendationsForPod_TaskResponse `xml:"RefreshStorageDrsRecommendationsForPod_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageDrsRecommendationsForPod_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshStorageDrsRecommendationsForPod_Task(ctx context.Context, r soap.RoundTripper, req *types.RefreshStorageDrsRecommendationsForPod_Task) (*types.RefreshStorageDrsRecommendationsForPod_TaskResponse, error) {
	var reqBody, resBody RefreshStorageDrsRecommendationsForPod_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshStorageInfoBody struct {
	Req    *types.RefreshStorageInfo         `xml:"urn:vim25 RefreshStorageInfo,omitempty"`
	Res    *types.RefreshStorageInfoResponse `xml:"RefreshStorageInfoResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshStorageInfo(ctx context.Context, r soap.RoundTripper, req *types.RefreshStorageInfo) (*types.RefreshStorageInfoResponse, error) {
	var reqBody, resBody RefreshStorageInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RefreshStorageSystemBody struct {
	Req    *types.RefreshStorageSystem         `xml:"urn:vim25 RefreshStorageSystem,omitempty"`
	Res    *types.RefreshStorageSystemResponse `xml:"RefreshStorageSystemResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageSystemBody) Fault() *soap.Fault { return b.Fault_ }

func RefreshStorageSystem(ctx context.Context, r soap.RoundTripper, req *types.RefreshStorageSystem) (*types.RefreshStorageSystemResponse, error) {
	var reqBody, resBody RefreshStorageSystemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterChildVM_TaskBody struct {
	Req    *types.RegisterChildVM_Task         `xml:"urn:vim25 RegisterChildVM_Task,omitempty"`
	Res    *types.RegisterChildVM_TaskResponse `xml:"RegisterChildVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterChildVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterChildVM_Task(ctx context.Context, r soap.RoundTripper, req *types.RegisterChildVM_Task) (*types.RegisterChildVM_TaskResponse, error) {
	var reqBody, resBody RegisterChildVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterDiskBody struct {
	Req    *types.RegisterDisk         `xml:"urn:vim25 RegisterDisk,omitempty"`
	Res    *types.RegisterDiskResponse `xml:"RegisterDiskResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterDiskBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterDisk(ctx context.Context, r soap.RoundTripper, req *types.RegisterDisk) (*types.RegisterDiskResponse, error) {
	var reqBody, resBody RegisterDiskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterExtensionBody struct {
	Req    *types.RegisterExtension         `xml:"urn:vim25 RegisterExtension,omitempty"`
	Res    *types.RegisterExtensionResponse `xml:"RegisterExtensionResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterExtensionBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterExtension(ctx context.Context, r soap.RoundTripper, req *types.RegisterExtension) (*types.RegisterExtensionResponse, error) {
	var reqBody, resBody RegisterExtensionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterHealthUpdateProviderBody struct {
	Req    *types.RegisterHealthUpdateProvider         `xml:"urn:vim25 RegisterHealthUpdateProvider,omitempty"`
	Res    *types.RegisterHealthUpdateProviderResponse `xml:"RegisterHealthUpdateProviderResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterHealthUpdateProviderBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterHealthUpdateProvider(ctx context.Context, r soap.RoundTripper, req *types.RegisterHealthUpdateProvider) (*types.RegisterHealthUpdateProviderResponse, error) {
	var reqBody, resBody RegisterHealthUpdateProviderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterKmipServerBody struct {
	Req    *types.RegisterKmipServer         `xml:"urn:vim25 RegisterKmipServer,omitempty"`
	Res    *types.RegisterKmipServerResponse `xml:"RegisterKmipServerResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterKmipServerBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterKmipServer(ctx context.Context, r soap.RoundTripper, req *types.RegisterKmipServer) (*types.RegisterKmipServerResponse, error) {
	var reqBody, resBody RegisterKmipServerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterKmsClusterBody struct {
	Req    *types.RegisterKmsCluster         `xml:"urn:vim25 RegisterKmsCluster,omitempty"`
	Res    *types.RegisterKmsClusterResponse `xml:"RegisterKmsClusterResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterKmsClusterBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterKmsCluster(ctx context.Context, r soap.RoundTripper, req *types.RegisterKmsCluster) (*types.RegisterKmsClusterResponse, error) {
	var reqBody, resBody RegisterKmsClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterVM_TaskBody struct {
	Req    *types.RegisterVM_Task         `xml:"urn:vim25 RegisterVM_Task,omitempty"`
	Res    *types.RegisterVM_TaskResponse `xml:"RegisterVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterVM_Task(ctx context.Context, r soap.RoundTripper, req *types.RegisterVM_Task) (*types.RegisterVM_TaskResponse, error) {
	var reqBody, resBody RegisterVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReleaseCredentialsInGuestBody struct {
	Req    *types.ReleaseCredentialsInGuest         `xml:"urn:vim25 ReleaseCredentialsInGuest,omitempty"`
	Res    *types.ReleaseCredentialsInGuestResponse `xml:"ReleaseCredentialsInGuestResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReleaseCredentialsInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ReleaseCredentialsInGuest(ctx context.Context, r soap.RoundTripper, req *types.ReleaseCredentialsInGuest) (*types.ReleaseCredentialsInGuestResponse, error) {
	var reqBody, resBody ReleaseCredentialsInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReleaseIpAllocationBody struct {
	Req    *types.ReleaseIpAllocation         `xml:"urn:vim25 ReleaseIpAllocation,omitempty"`
	Res    *types.ReleaseIpAllocationResponse `xml:"ReleaseIpAllocationResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReleaseIpAllocationBody) Fault() *soap.Fault { return b.Fault_ }

func ReleaseIpAllocation(ctx context.Context, r soap.RoundTripper, req *types.ReleaseIpAllocation) (*types.ReleaseIpAllocationResponse, error) {
	var reqBody, resBody ReleaseIpAllocationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReleaseManagedSnapshotBody struct {
	Req    *types.ReleaseManagedSnapshot         `xml:"urn:vim25 ReleaseManagedSnapshot,omitempty"`
	Res    *types.ReleaseManagedSnapshotResponse `xml:"ReleaseManagedSnapshotResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReleaseManagedSnapshotBody) Fault() *soap.Fault { return b.Fault_ }

func ReleaseManagedSnapshot(ctx context.Context, r soap.RoundTripper, req *types.ReleaseManagedSnapshot) (*types.ReleaseManagedSnapshotResponse, error) {
	var reqBody, resBody ReleaseManagedSnapshotBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReloadBody struct {
	Req    *types.Reload         `xml:"urn:vim25 Reload,omitempty"`
	Res    *types.ReloadResponse `xml:"ReloadResponse,omitempty"`
	Fault_ *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReloadBody) Fault() *soap.Fault { return b.Fault_ }

func Reload(ctx context.Context, r soap.RoundTripper, req *types.Reload) (*types.ReloadResponse, error) {
	var reqBody, resBody ReloadBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RelocateVM_TaskBody struct {
	Req    *types.RelocateVM_Task         `xml:"urn:vim25 RelocateVM_Task,omitempty"`
	Res    *types.RelocateVM_TaskResponse `xml:"RelocateVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RelocateVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RelocateVM_Task(ctx context.Context, r soap.RoundTripper, req *types.RelocateVM_Task) (*types.RelocateVM_TaskResponse, error) {
	var reqBody, resBody RelocateVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RelocateVStorageObject_TaskBody struct {
	Req    *types.RelocateVStorageObject_Task         `xml:"urn:vim25 RelocateVStorageObject_Task,omitempty"`
	Res    *types.RelocateVStorageObject_TaskResponse `xml:"RelocateVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RelocateVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RelocateVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.RelocateVStorageObject_Task) (*types.RelocateVStorageObject_TaskResponse, error) {
	var reqBody, resBody RelocateVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveAlarmBody struct {
	Req    *types.RemoveAlarm         `xml:"urn:vim25 RemoveAlarm,omitempty"`
	Res    *types.RemoveAlarmResponse `xml:"RemoveAlarmResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAlarmBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveAlarm(ctx context.Context, r soap.RoundTripper, req *types.RemoveAlarm) (*types.RemoveAlarmResponse, error) {
	var reqBody, resBody RemoveAlarmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveAllSnapshots_TaskBody struct {
	Req    *types.RemoveAllSnapshots_Task         `xml:"urn:vim25 RemoveAllSnapshots_Task,omitempty"`
	Res    *types.RemoveAllSnapshots_TaskResponse `xml:"RemoveAllSnapshots_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAllSnapshots_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveAllSnapshots_Task(ctx context.Context, r soap.RoundTripper, req *types.RemoveAllSnapshots_Task) (*types.RemoveAllSnapshots_TaskResponse, error) {
	var reqBody, resBody RemoveAllSnapshots_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveAssignedLicenseBody struct {
	Req    *types.RemoveAssignedLicense         `xml:"urn:vim25 RemoveAssignedLicense,omitempty"`
	Res    *types.RemoveAssignedLicenseResponse `xml:"RemoveAssignedLicenseResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAssignedLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveAssignedLicense(ctx context.Context, r soap.RoundTripper, req *types.RemoveAssignedLicense) (*types.RemoveAssignedLicenseResponse, error) {
	var reqBody, resBody RemoveAssignedLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveAuthorizationRoleBody struct {
	Req    *types.RemoveAuthorizationRole         `xml:"urn:vim25 RemoveAuthorizationRole,omitempty"`
	Res    *types.RemoveAuthorizationRoleResponse `xml:"RemoveAuthorizationRoleResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAuthorizationRoleBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveAuthorizationRole(ctx context.Context, r soap.RoundTripper, req *types.RemoveAuthorizationRole) (*types.RemoveAuthorizationRoleResponse, error) {
	var reqBody, resBody RemoveAuthorizationRoleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveCustomFieldDefBody struct {
	Req    *types.RemoveCustomFieldDef         `xml:"urn:vim25 RemoveCustomFieldDef,omitempty"`
	Res    *types.RemoveCustomFieldDefResponse `xml:"RemoveCustomFieldDefResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveCustomFieldDefBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveCustomFieldDef(ctx context.Context, r soap.RoundTripper, req *types.RemoveCustomFieldDef) (*types.RemoveCustomFieldDefResponse, error) {
	var reqBody, resBody RemoveCustomFieldDefBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDatastoreBody struct {
	Req    *types.RemoveDatastore         `xml:"urn:vim25 RemoveDatastore,omitempty"`
	Res    *types.RemoveDatastoreResponse `xml:"RemoveDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDatastore(ctx context.Context, r soap.RoundTripper, req *types.RemoveDatastore) (*types.RemoveDatastoreResponse, error) {
	var reqBody, resBody RemoveDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDatastoreEx_TaskBody struct {
	Req    *types.RemoveDatastoreEx_Task         `xml:"urn:vim25 RemoveDatastoreEx_Task,omitempty"`
	Res    *types.RemoveDatastoreEx_TaskResponse `xml:"RemoveDatastoreEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDatastoreEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDatastoreEx_Task(ctx context.Context, r soap.RoundTripper, req *types.RemoveDatastoreEx_Task) (*types.RemoveDatastoreEx_TaskResponse, error) {
	var reqBody, resBody RemoveDatastoreEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDiskMapping_TaskBody struct {
	Req    *types.RemoveDiskMapping_Task         `xml:"urn:vim25 RemoveDiskMapping_Task,omitempty"`
	Res    *types.RemoveDiskMapping_TaskResponse `xml:"RemoveDiskMapping_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDiskMapping_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDiskMapping_Task(ctx context.Context, r soap.RoundTripper, req *types.RemoveDiskMapping_Task) (*types.RemoveDiskMapping_TaskResponse, error) {
	var reqBody, resBody RemoveDiskMapping_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveDisk_TaskBody struct {
	Req    *types.RemoveDisk_Task         `xml:"urn:vim25 RemoveDisk_Task,omitempty"`
	Res    *types.RemoveDisk_TaskResponse `xml:"RemoveDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.RemoveDisk_Task) (*types.RemoveDisk_TaskResponse, error) {
	var reqBody, resBody RemoveDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveEntityPermissionBody struct {
	Req    *types.RemoveEntityPermission         `xml:"urn:vim25 RemoveEntityPermission,omitempty"`
	Res    *types.RemoveEntityPermissionResponse `xml:"RemoveEntityPermissionResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveEntityPermissionBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveEntityPermission(ctx context.Context, r soap.RoundTripper, req *types.RemoveEntityPermission) (*types.RemoveEntityPermissionResponse, error) {
	var reqBody, resBody RemoveEntityPermissionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveFilterBody struct {
	Req    *types.RemoveFilter         `xml:"urn:vim25 RemoveFilter,omitempty"`
	Res    *types.RemoveFilterResponse `xml:"RemoveFilterResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveFilterBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveFilter(ctx context.Context, r soap.RoundTripper, req *types.RemoveFilter) (*types.RemoveFilterResponse, error) {
	var reqBody, resBody RemoveFilterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveFilterEntitiesBody struct {
	Req    *types.RemoveFilterEntities         `xml:"urn:vim25 RemoveFilterEntities,omitempty"`
	Res    *types.RemoveFilterEntitiesResponse `xml:"RemoveFilterEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveFilterEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveFilterEntities(ctx context.Context, r soap.RoundTripper, req *types.RemoveFilterEntities) (*types.RemoveFilterEntitiesResponse, error) {
	var reqBody, resBody RemoveFilterEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveGroupBody struct {
	Req    *types.RemoveGroup         `xml:"urn:vim25 RemoveGroup,omitempty"`
	Res    *types.RemoveGroupResponse `xml:"RemoveGroupResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveGroupBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveGroup(ctx context.Context, r soap.RoundTripper, req *types.RemoveGroup) (*types.RemoveGroupResponse, error) {
	var reqBody, resBody RemoveGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveGuestAliasBody struct {
	Req    *types.RemoveGuestAlias         `xml:"urn:vim25 RemoveGuestAlias,omitempty"`
	Res    *types.RemoveGuestAliasResponse `xml:"RemoveGuestAliasResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveGuestAliasBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveGuestAlias(ctx context.Context, r soap.RoundTripper, req *types.RemoveGuestAlias) (*types.RemoveGuestAliasResponse, error) {
	var reqBody, resBody RemoveGuestAliasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveGuestAliasByCertBody struct {
	Req    *types.RemoveGuestAliasByCert         `xml:"urn:vim25 RemoveGuestAliasByCert,omitempty"`
	Res    *types.RemoveGuestAliasByCertResponse `xml:"RemoveGuestAliasByCertResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveGuestAliasByCertBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveGuestAliasByCert(ctx context.Context, r soap.RoundTripper, req *types.RemoveGuestAliasByCert) (*types.RemoveGuestAliasByCertResponse, error) {
	var reqBody, resBody RemoveGuestAliasByCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveInternetScsiSendTargetsBody struct {
	Req    *types.RemoveInternetScsiSendTargets         `xml:"urn:vim25 RemoveInternetScsiSendTargets,omitempty"`
	Res    *types.RemoveInternetScsiSendTargetsResponse `xml:"RemoveInternetScsiSendTargetsResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveInternetScsiSendTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveInternetScsiSendTargets(ctx context.Context, r soap.RoundTripper, req *types.RemoveInternetScsiSendTargets) (*types.RemoveInternetScsiSendTargetsResponse, error) {
	var reqBody, resBody RemoveInternetScsiSendTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveInternetScsiStaticTargetsBody struct {
	Req    *types.RemoveInternetScsiStaticTargets         `xml:"urn:vim25 RemoveInternetScsiStaticTargets,omitempty"`
	Res    *types.RemoveInternetScsiStaticTargetsResponse `xml:"RemoveInternetScsiStaticTargetsResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveInternetScsiStaticTargetsBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveInternetScsiStaticTargets(ctx context.Context, r soap.RoundTripper, req *types.RemoveInternetScsiStaticTargets) (*types.RemoveInternetScsiStaticTargetsResponse, error) {
	var reqBody, resBody RemoveInternetScsiStaticTargetsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveKeyBody struct {
	Req    *types.RemoveKey         `xml:"urn:vim25 RemoveKey,omitempty"`
	Res    *types.RemoveKeyResponse `xml:"RemoveKeyResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveKeyBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveKey(ctx context.Context, r soap.RoundTripper, req *types.RemoveKey) (*types.RemoveKeyResponse, error) {
	var reqBody, resBody RemoveKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveKeysBody struct {
	Req    *types.RemoveKeys         `xml:"urn:vim25 RemoveKeys,omitempty"`
	Res    *types.RemoveKeysResponse `xml:"RemoveKeysResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveKeysBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveKeys(ctx context.Context, r soap.RoundTripper, req *types.RemoveKeys) (*types.RemoveKeysResponse, error) {
	var reqBody, resBody RemoveKeysBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveKmipServerBody struct {
	Req    *types.RemoveKmipServer         `xml:"urn:vim25 RemoveKmipServer,omitempty"`
	Res    *types.RemoveKmipServerResponse `xml:"RemoveKmipServerResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveKmipServerBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveKmipServer(ctx context.Context, r soap.RoundTripper, req *types.RemoveKmipServer) (*types.RemoveKmipServerResponse, error) {
	var reqBody, resBody RemoveKmipServerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveLicenseBody struct {
	Req    *types.RemoveLicense         `xml:"urn:vim25 RemoveLicense,omitempty"`
	Res    *types.RemoveLicenseResponse `xml:"RemoveLicenseResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveLicense(ctx context.Context, r soap.RoundTripper, req *types.RemoveLicense) (*types.RemoveLicenseResponse, error) {
	var reqBody, resBody RemoveLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveLicenseLabelBody struct {
	Req    *types.RemoveLicenseLabel         `xml:"urn:vim25 RemoveLicenseLabel,omitempty"`
	Res    *types.RemoveLicenseLabelResponse `xml:"RemoveLicenseLabelResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveLicenseLabelBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveLicenseLabel(ctx context.Context, r soap.RoundTripper, req *types.RemoveLicenseLabel) (*types.RemoveLicenseLabelResponse, error) {
	var reqBody, resBody RemoveLicenseLabelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveMonitoredEntitiesBody struct {
	Req    *types.RemoveMonitoredEntities         `xml:"urn:vim25 RemoveMonitoredEntities,omitempty"`
	Res    *types.RemoveMonitoredEntitiesResponse `xml:"RemoveMonitoredEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveMonitoredEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveMonitoredEntities(ctx context.Context, r soap.RoundTripper, req *types.RemoveMonitoredEntities) (*types.RemoveMonitoredEntitiesResponse, error) {
	var reqBody, resBody RemoveMonitoredEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveNetworkResourcePoolBody struct {
	Req    *types.RemoveNetworkResourcePool         `xml:"urn:vim25 RemoveNetworkResourcePool,omitempty"`
	Res    *types.RemoveNetworkResourcePoolResponse `xml:"RemoveNetworkResourcePoolResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveNetworkResourcePoolBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveNetworkResourcePool(ctx context.Context, r soap.RoundTripper, req *types.RemoveNetworkResourcePool) (*types.RemoveNetworkResourcePoolResponse, error) {
	var reqBody, resBody RemoveNetworkResourcePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveNvmeOverRdmaAdapterBody struct {
	Req    *types.RemoveNvmeOverRdmaAdapter         `xml:"urn:vim25 RemoveNvmeOverRdmaAdapter,omitempty"`
	Res    *types.RemoveNvmeOverRdmaAdapterResponse `xml:"RemoveNvmeOverRdmaAdapterResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveNvmeOverRdmaAdapterBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveNvmeOverRdmaAdapter(ctx context.Context, r soap.RoundTripper, req *types.RemoveNvmeOverRdmaAdapter) (*types.RemoveNvmeOverRdmaAdapterResponse, error) {
	var reqBody, resBody RemoveNvmeOverRdmaAdapterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemovePerfIntervalBody struct {
	Req    *types.RemovePerfInterval         `xml:"urn:vim25 RemovePerfInterval,omitempty"`
	Res    *types.RemovePerfIntervalResponse `xml:"RemovePerfIntervalResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemovePerfIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func RemovePerfInterval(ctx context.Context, r soap.RoundTripper, req *types.RemovePerfInterval) (*types.RemovePerfIntervalResponse, error) {
	var reqBody, resBody RemovePerfIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemovePortGroupBody struct {
	Req    *types.RemovePortGroup         `xml:"urn:vim25 RemovePortGroup,omitempty"`
	Res    *types.RemovePortGroupResponse `xml:"RemovePortGroupResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemovePortGroupBody) Fault() *soap.Fault { return b.Fault_ }

func RemovePortGroup(ctx context.Context, r soap.RoundTripper, req *types.RemovePortGroup) (*types.RemovePortGroupResponse, error) {
	var reqBody, resBody RemovePortGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveScheduledTaskBody struct {
	Req    *types.RemoveScheduledTask         `xml:"urn:vim25 RemoveScheduledTask,omitempty"`
	Res    *types.RemoveScheduledTaskResponse `xml:"RemoveScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.RemoveScheduledTask) (*types.RemoveScheduledTaskResponse, error) {
	var reqBody, resBody RemoveScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveServiceConsoleVirtualNicBody struct {
	Req    *types.RemoveServiceConsoleVirtualNic         `xml:"urn:vim25 RemoveServiceConsoleVirtualNic,omitempty"`
	Res    *types.RemoveServiceConsoleVirtualNicResponse `xml:"RemoveServiceConsoleVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveServiceConsoleVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveServiceConsoleVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.RemoveServiceConsoleVirtualNic) (*types.RemoveServiceConsoleVirtualNicResponse, error) {
	var reqBody, resBody RemoveServiceConsoleVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveSmartCardTrustAnchorBody struct {
	Req    *types.RemoveSmartCardTrustAnchor         `xml:"urn:vim25 RemoveSmartCardTrustAnchor,omitempty"`
	Res    *types.RemoveSmartCardTrustAnchorResponse `xml:"RemoveSmartCardTrustAnchorResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSmartCardTrustAnchorBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveSmartCardTrustAnchor(ctx context.Context, r soap.RoundTripper, req *types.RemoveSmartCardTrustAnchor) (*types.RemoveSmartCardTrustAnchorResponse, error) {
	var reqBody, resBody RemoveSmartCardTrustAnchorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveSmartCardTrustAnchorByFingerprintBody struct {
	Req    *types.RemoveSmartCardTrustAnchorByFingerprint         `xml:"urn:vim25 RemoveSmartCardTrustAnchorByFingerprint,omitempty"`
	Res    *types.RemoveSmartCardTrustAnchorByFingerprintResponse `xml:"RemoveSmartCardTrustAnchorByFingerprintResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSmartCardTrustAnchorByFingerprintBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveSmartCardTrustAnchorByFingerprint(ctx context.Context, r soap.RoundTripper, req *types.RemoveSmartCardTrustAnchorByFingerprint) (*types.RemoveSmartCardTrustAnchorByFingerprintResponse, error) {
	var reqBody, resBody RemoveSmartCardTrustAnchorByFingerprintBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveSmartCardTrustAnchorCertificateBody struct {
	Req    *types.RemoveSmartCardTrustAnchorCertificate         `xml:"urn:vim25 RemoveSmartCardTrustAnchorCertificate,omitempty"`
	Res    *types.RemoveSmartCardTrustAnchorCertificateResponse `xml:"RemoveSmartCardTrustAnchorCertificateResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSmartCardTrustAnchorCertificateBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveSmartCardTrustAnchorCertificate(ctx context.Context, r soap.RoundTripper, req *types.RemoveSmartCardTrustAnchorCertificate) (*types.RemoveSmartCardTrustAnchorCertificateResponse, error) {
	var reqBody, resBody RemoveSmartCardTrustAnchorCertificateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveSnapshot_TaskBody struct {
	Req    *types.RemoveSnapshot_Task         `xml:"urn:vim25 RemoveSnapshot_Task,omitempty"`
	Res    *types.RemoveSnapshot_TaskResponse `xml:"RemoveSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.RemoveSnapshot_Task) (*types.RemoveSnapshot_TaskResponse, error) {
	var reqBody, resBody RemoveSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveSoftwareAdapterBody struct {
	Req    *types.RemoveSoftwareAdapter         `xml:"urn:vim25 RemoveSoftwareAdapter,omitempty"`
	Res    *types.RemoveSoftwareAdapterResponse `xml:"RemoveSoftwareAdapterResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSoftwareAdapterBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveSoftwareAdapter(ctx context.Context, r soap.RoundTripper, req *types.RemoveSoftwareAdapter) (*types.RemoveSoftwareAdapterResponse, error) {
	var reqBody, resBody RemoveSoftwareAdapterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveUserBody struct {
	Req    *types.RemoveUser         `xml:"urn:vim25 RemoveUser,omitempty"`
	Res    *types.RemoveUserResponse `xml:"RemoveUserResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveUserBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveUser(ctx context.Context, r soap.RoundTripper, req *types.RemoveUser) (*types.RemoveUserResponse, error) {
	var reqBody, resBody RemoveUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveVirtualNicBody struct {
	Req    *types.RemoveVirtualNic         `xml:"urn:vim25 RemoveVirtualNic,omitempty"`
	Res    *types.RemoveVirtualNicResponse `xml:"RemoveVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.RemoveVirtualNic) (*types.RemoveVirtualNicResponse, error) {
	var reqBody, resBody RemoveVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RemoveVirtualSwitchBody struct {
	Req    *types.RemoveVirtualSwitch         `xml:"urn:vim25 RemoveVirtualSwitch,omitempty"`
	Res    *types.RemoveVirtualSwitchResponse `xml:"RemoveVirtualSwitchResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveVirtualSwitchBody) Fault() *soap.Fault { return b.Fault_ }

func RemoveVirtualSwitch(ctx context.Context, r soap.RoundTripper, req *types.RemoveVirtualSwitch) (*types.RemoveVirtualSwitchResponse, error) {
	var reqBody, resBody RemoveVirtualSwitchBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameCustomFieldDefBody struct {
	Req    *types.RenameCustomFieldDef         `xml:"urn:vim25 RenameCustomFieldDef,omitempty"`
	Res    *types.RenameCustomFieldDefResponse `xml:"RenameCustomFieldDefResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameCustomFieldDefBody) Fault() *soap.Fault { return b.Fault_ }

func RenameCustomFieldDef(ctx context.Context, r soap.RoundTripper, req *types.RenameCustomFieldDef) (*types.RenameCustomFieldDefResponse, error) {
	var reqBody, resBody RenameCustomFieldDefBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameCustomizationSpecBody struct {
	Req    *types.RenameCustomizationSpec         `xml:"urn:vim25 RenameCustomizationSpec,omitempty"`
	Res    *types.RenameCustomizationSpecResponse `xml:"RenameCustomizationSpecResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameCustomizationSpecBody) Fault() *soap.Fault { return b.Fault_ }

func RenameCustomizationSpec(ctx context.Context, r soap.RoundTripper, req *types.RenameCustomizationSpec) (*types.RenameCustomizationSpecResponse, error) {
	var reqBody, resBody RenameCustomizationSpecBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameDatastoreBody struct {
	Req    *types.RenameDatastore         `xml:"urn:vim25 RenameDatastore,omitempty"`
	Res    *types.RenameDatastoreResponse `xml:"RenameDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func RenameDatastore(ctx context.Context, r soap.RoundTripper, req *types.RenameDatastore) (*types.RenameDatastoreResponse, error) {
	var reqBody, resBody RenameDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameSnapshotBody struct {
	Req    *types.RenameSnapshot         `xml:"urn:vim25 RenameSnapshot,omitempty"`
	Res    *types.RenameSnapshotResponse `xml:"RenameSnapshotResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameSnapshotBody) Fault() *soap.Fault { return b.Fault_ }

func RenameSnapshot(ctx context.Context, r soap.RoundTripper, req *types.RenameSnapshot) (*types.RenameSnapshotResponse, error) {
	var reqBody, resBody RenameSnapshotBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameVStorageObjectBody struct {
	Req    *types.RenameVStorageObject         `xml:"urn:vim25 RenameVStorageObject,omitempty"`
	Res    *types.RenameVStorageObjectResponse `xml:"RenameVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func RenameVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.RenameVStorageObject) (*types.RenameVStorageObjectResponse, error) {
	var reqBody, resBody RenameVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RenameVStorageObjectExBody struct {
	Req    *types.RenameVStorageObjectEx         `xml:"urn:vim25 RenameVStorageObjectEx,omitempty"`
	Res    *types.RenameVStorageObjectExResponse `xml:"RenameVStorageObjectExResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameVStorageObjectExBody) Fault() *soap.Fault { return b.Fault_ }

func RenameVStorageObjectEx(ctx context.Context, r soap.RoundTripper, req *types.RenameVStorageObjectEx) (*types.RenameVStorageObjectExResponse, error) {
	var reqBody, resBody RenameVStorageObjectExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type Rename_TaskBody struct {
	Req    *types.Rename_Task         `xml:"urn:vim25 Rename_Task,omitempty"`
	Res    *types.Rename_TaskResponse `xml:"Rename_TaskResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Rename_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func Rename_Task(ctx context.Context, r soap.RoundTripper, req *types.Rename_Task) (*types.Rename_TaskResponse, error) {
	var reqBody, resBody Rename_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReplaceCACertificatesAndCRLsBody struct {
	Req    *types.ReplaceCACertificatesAndCRLs         `xml:"urn:vim25 ReplaceCACertificatesAndCRLs,omitempty"`
	Res    *types.ReplaceCACertificatesAndCRLsResponse `xml:"ReplaceCACertificatesAndCRLsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReplaceCACertificatesAndCRLsBody) Fault() *soap.Fault { return b.Fault_ }

func ReplaceCACertificatesAndCRLs(ctx context.Context, r soap.RoundTripper, req *types.ReplaceCACertificatesAndCRLs) (*types.ReplaceCACertificatesAndCRLsResponse, error) {
	var reqBody, resBody ReplaceCACertificatesAndCRLsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReplaceSmartCardTrustAnchorsBody struct {
	Req    *types.ReplaceSmartCardTrustAnchors         `xml:"urn:vim25 ReplaceSmartCardTrustAnchors,omitempty"`
	Res    *types.ReplaceSmartCardTrustAnchorsResponse `xml:"ReplaceSmartCardTrustAnchorsResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReplaceSmartCardTrustAnchorsBody) Fault() *soap.Fault { return b.Fault_ }

func ReplaceSmartCardTrustAnchors(ctx context.Context, r soap.RoundTripper, req *types.ReplaceSmartCardTrustAnchors) (*types.ReplaceSmartCardTrustAnchorsResponse, error) {
	var reqBody, resBody ReplaceSmartCardTrustAnchorsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RescanAllHbaBody struct {
	Req    *types.RescanAllHba         `xml:"urn:vim25 RescanAllHba,omitempty"`
	Res    *types.RescanAllHbaResponse `xml:"RescanAllHbaResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanAllHbaBody) Fault() *soap.Fault { return b.Fault_ }

func RescanAllHba(ctx context.Context, r soap.RoundTripper, req *types.RescanAllHba) (*types.RescanAllHbaResponse, error) {
	var reqBody, resBody RescanAllHbaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RescanHbaBody struct {
	Req    *types.RescanHba         `xml:"urn:vim25 RescanHba,omitempty"`
	Res    *types.RescanHbaResponse `xml:"RescanHbaResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanHbaBody) Fault() *soap.Fault { return b.Fault_ }

func RescanHba(ctx context.Context, r soap.RoundTripper, req *types.RescanHba) (*types.RescanHbaResponse, error) {
	var reqBody, resBody RescanHbaBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RescanVffsBody struct {
	Req    *types.RescanVffs         `xml:"urn:vim25 RescanVffs,omitempty"`
	Res    *types.RescanVffsResponse `xml:"RescanVffsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanVffsBody) Fault() *soap.Fault { return b.Fault_ }

func RescanVffs(ctx context.Context, r soap.RoundTripper, req *types.RescanVffs) (*types.RescanVffsResponse, error) {
	var reqBody, resBody RescanVffsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RescanVmfsBody struct {
	Req    *types.RescanVmfs         `xml:"urn:vim25 RescanVmfs,omitempty"`
	Res    *types.RescanVmfsResponse `xml:"RescanVmfsResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanVmfsBody) Fault() *soap.Fault { return b.Fault_ }

func RescanVmfs(ctx context.Context, r soap.RoundTripper, req *types.RescanVmfs) (*types.RescanVmfsResponse, error) {
	var reqBody, resBody RescanVmfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetCollectorBody struct {
	Req    *types.ResetCollector         `xml:"urn:vim25 ResetCollector,omitempty"`
	Res    *types.ResetCollectorResponse `xml:"ResetCollectorResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func ResetCollector(ctx context.Context, r soap.RoundTripper, req *types.ResetCollector) (*types.ResetCollectorResponse, error) {
	var reqBody, resBody ResetCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetCounterLevelMappingBody struct {
	Req    *types.ResetCounterLevelMapping         `xml:"urn:vim25 ResetCounterLevelMapping,omitempty"`
	Res    *types.ResetCounterLevelMappingResponse `xml:"ResetCounterLevelMappingResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetCounterLevelMappingBody) Fault() *soap.Fault { return b.Fault_ }

func ResetCounterLevelMapping(ctx context.Context, r soap.RoundTripper, req *types.ResetCounterLevelMapping) (*types.ResetCounterLevelMappingResponse, error) {
	var reqBody, resBody ResetCounterLevelMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetEntityPermissionsBody struct {
	Req    *types.ResetEntityPermissions         `xml:"urn:vim25 ResetEntityPermissions,omitempty"`
	Res    *types.ResetEntityPermissionsResponse `xml:"ResetEntityPermissionsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetEntityPermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func ResetEntityPermissions(ctx context.Context, r soap.RoundTripper, req *types.ResetEntityPermissions) (*types.ResetEntityPermissionsResponse, error) {
	var reqBody, resBody ResetEntityPermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetFirmwareToFactoryDefaultsBody struct {
	Req    *types.ResetFirmwareToFactoryDefaults         `xml:"urn:vim25 ResetFirmwareToFactoryDefaults,omitempty"`
	Res    *types.ResetFirmwareToFactoryDefaultsResponse `xml:"ResetFirmwareToFactoryDefaultsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetFirmwareToFactoryDefaultsBody) Fault() *soap.Fault { return b.Fault_ }

func ResetFirmwareToFactoryDefaults(ctx context.Context, r soap.RoundTripper, req *types.ResetFirmwareToFactoryDefaults) (*types.ResetFirmwareToFactoryDefaultsResponse, error) {
	var reqBody, resBody ResetFirmwareToFactoryDefaultsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetGuestInformationBody struct {
	Req    *types.ResetGuestInformation         `xml:"urn:vim25 ResetGuestInformation,omitempty"`
	Res    *types.ResetGuestInformationResponse `xml:"ResetGuestInformationResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetGuestInformationBody) Fault() *soap.Fault { return b.Fault_ }

func ResetGuestInformation(ctx context.Context, r soap.RoundTripper, req *types.ResetGuestInformation) (*types.ResetGuestInformationResponse, error) {
	var reqBody, resBody ResetGuestInformationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetListViewBody struct {
	Req    *types.ResetListView         `xml:"urn:vim25 ResetListView,omitempty"`
	Res    *types.ResetListViewResponse `xml:"ResetListViewResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetListViewBody) Fault() *soap.Fault { return b.Fault_ }

func ResetListView(ctx context.Context, r soap.RoundTripper, req *types.ResetListView) (*types.ResetListViewResponse, error) {
	var reqBody, resBody ResetListViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetListViewFromViewBody struct {
	Req    *types.ResetListViewFromView         `xml:"urn:vim25 ResetListViewFromView,omitempty"`
	Res    *types.ResetListViewFromViewResponse `xml:"ResetListViewFromViewResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetListViewFromViewBody) Fault() *soap.Fault { return b.Fault_ }

func ResetListViewFromView(ctx context.Context, r soap.RoundTripper, req *types.ResetListViewFromView) (*types.ResetListViewFromViewResponse, error) {
	var reqBody, resBody ResetListViewFromViewBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetSystemHealthInfoBody struct {
	Req    *types.ResetSystemHealthInfo         `xml:"urn:vim25 ResetSystemHealthInfo,omitempty"`
	Res    *types.ResetSystemHealthInfoResponse `xml:"ResetSystemHealthInfoResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetSystemHealthInfoBody) Fault() *soap.Fault { return b.Fault_ }

func ResetSystemHealthInfo(ctx context.Context, r soap.RoundTripper, req *types.ResetSystemHealthInfo) (*types.ResetSystemHealthInfoResponse, error) {
	var reqBody, resBody ResetSystemHealthInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResetVM_TaskBody struct {
	Req    *types.ResetVM_Task         `xml:"urn:vim25 ResetVM_Task,omitempty"`
	Res    *types.ResetVM_TaskResponse `xml:"ResetVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ResetVM_Task(ctx context.Context, r soap.RoundTripper, req *types.ResetVM_Task) (*types.ResetVM_TaskResponse, error) {
	var reqBody, resBody ResetVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResignatureUnresolvedVmfsVolume_TaskBody struct {
	Req    *types.ResignatureUnresolvedVmfsVolume_Task         `xml:"urn:vim25 ResignatureUnresolvedVmfsVolume_Task,omitempty"`
	Res    *types.ResignatureUnresolvedVmfsVolume_TaskResponse `xml:"ResignatureUnresolvedVmfsVolume_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResignatureUnresolvedVmfsVolume_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ResignatureUnresolvedVmfsVolume_Task(ctx context.Context, r soap.RoundTripper, req *types.ResignatureUnresolvedVmfsVolume_Task) (*types.ResignatureUnresolvedVmfsVolume_TaskResponse, error) {
	var reqBody, resBody ResignatureUnresolvedVmfsVolume_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveInstallationErrorsOnCluster_TaskBody struct {
	Req    *types.ResolveInstallationErrorsOnCluster_Task         `xml:"urn:vim25 ResolveInstallationErrorsOnCluster_Task,omitempty"`
	Res    *types.ResolveInstallationErrorsOnCluster_TaskResponse `xml:"ResolveInstallationErrorsOnCluster_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveInstallationErrorsOnCluster_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ResolveInstallationErrorsOnCluster_Task(ctx context.Context, r soap.RoundTripper, req *types.ResolveInstallationErrorsOnCluster_Task) (*types.ResolveInstallationErrorsOnCluster_TaskResponse, error) {
	var reqBody, resBody ResolveInstallationErrorsOnCluster_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveInstallationErrorsOnHost_TaskBody struct {
	Req    *types.ResolveInstallationErrorsOnHost_Task         `xml:"urn:vim25 ResolveInstallationErrorsOnHost_Task,omitempty"`
	Res    *types.ResolveInstallationErrorsOnHost_TaskResponse `xml:"ResolveInstallationErrorsOnHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveInstallationErrorsOnHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ResolveInstallationErrorsOnHost_Task(ctx context.Context, r soap.RoundTripper, req *types.ResolveInstallationErrorsOnHost_Task) (*types.ResolveInstallationErrorsOnHost_TaskResponse, error) {
	var reqBody, resBody ResolveInstallationErrorsOnHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveMultipleUnresolvedVmfsVolumesBody struct {
	Req    *types.ResolveMultipleUnresolvedVmfsVolumes         `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumes,omitempty"`
	Res    *types.ResolveMultipleUnresolvedVmfsVolumesResponse `xml:"ResolveMultipleUnresolvedVmfsVolumesResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveMultipleUnresolvedVmfsVolumesBody) Fault() *soap.Fault { return b.Fault_ }

func ResolveMultipleUnresolvedVmfsVolumes(ctx context.Context, r soap.RoundTripper, req *types.ResolveMultipleUnresolvedVmfsVolumes) (*types.ResolveMultipleUnresolvedVmfsVolumesResponse, error) {
	var reqBody, resBody ResolveMultipleUnresolvedVmfsVolumesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody struct {
	Req    *types.ResolveMultipleUnresolvedVmfsVolumesEx_Task         `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumesEx_Task,omitempty"`
	Res    *types.ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse `xml:"ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ResolveMultipleUnresolvedVmfsVolumesEx_Task(ctx context.Context, r soap.RoundTripper, req *types.ResolveMultipleUnresolvedVmfsVolumesEx_Task) (*types.ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse, error) {
	var reqBody, resBody ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RestartServiceBody struct {
	Req    *types.RestartService         `xml:"urn:vim25 RestartService,omitempty"`
	Res    *types.RestartServiceResponse `xml:"RestartServiceResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestartServiceBody) Fault() *soap.Fault { return b.Fault_ }

func RestartService(ctx context.Context, r soap.RoundTripper, req *types.RestartService) (*types.RestartServiceResponse, error) {
	var reqBody, resBody RestartServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RestartServiceConsoleVirtualNicBody struct {
	Req    *types.RestartServiceConsoleVirtualNic         `xml:"urn:vim25 RestartServiceConsoleVirtualNic,omitempty"`
	Res    *types.RestartServiceConsoleVirtualNicResponse `xml:"RestartServiceConsoleVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestartServiceConsoleVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func RestartServiceConsoleVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.RestartServiceConsoleVirtualNic) (*types.RestartServiceConsoleVirtualNicResponse, error) {
	var reqBody, resBody RestartServiceConsoleVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RestoreFirmwareConfigurationBody struct {
	Req    *types.RestoreFirmwareConfiguration         `xml:"urn:vim25 RestoreFirmwareConfiguration,omitempty"`
	Res    *types.RestoreFirmwareConfigurationResponse `xml:"RestoreFirmwareConfigurationResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestoreFirmwareConfigurationBody) Fault() *soap.Fault { return b.Fault_ }

func RestoreFirmwareConfiguration(ctx context.Context, r soap.RoundTripper, req *types.RestoreFirmwareConfiguration) (*types.RestoreFirmwareConfigurationResponse, error) {
	var reqBody, resBody RestoreFirmwareConfigurationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAllPermissionsBody struct {
	Req    *types.RetrieveAllPermissions         `xml:"urn:vim25 RetrieveAllPermissions,omitempty"`
	Res    *types.RetrieveAllPermissionsResponse `xml:"RetrieveAllPermissionsResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAllPermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAllPermissions(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAllPermissions) (*types.RetrieveAllPermissionsResponse, error) {
	var reqBody, resBody RetrieveAllPermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAnswerFileBody struct {
	Req    *types.RetrieveAnswerFile         `xml:"urn:vim25 RetrieveAnswerFile,omitempty"`
	Res    *types.RetrieveAnswerFileResponse `xml:"RetrieveAnswerFileResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAnswerFileBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAnswerFile(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAnswerFile) (*types.RetrieveAnswerFileResponse, error) {
	var reqBody, resBody RetrieveAnswerFileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveAnswerFileForProfileBody struct {
	Req    *types.RetrieveAnswerFileForProfile         `xml:"urn:vim25 RetrieveAnswerFileForProfile,omitempty"`
	Res    *types.RetrieveAnswerFileForProfileResponse `xml:"RetrieveAnswerFileForProfileResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAnswerFileForProfileBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveAnswerFileForProfile(ctx context.Context, r soap.RoundTripper, req *types.RetrieveAnswerFileForProfile) (*types.RetrieveAnswerFileForProfileResponse, error) {
	var reqBody, resBody RetrieveAnswerFileForProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveArgumentDescriptionBody struct {
	Req    *types.RetrieveArgumentDescription         `xml:"urn:vim25 RetrieveArgumentDescription,omitempty"`
	Res    *types.RetrieveArgumentDescriptionResponse `xml:"RetrieveArgumentDescriptionResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveArgumentDescriptionBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveArgumentDescription(ctx context.Context, r soap.RoundTripper, req *types.RetrieveArgumentDescription) (*types.RetrieveArgumentDescriptionResponse, error) {
	var reqBody, resBody RetrieveArgumentDescriptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveCertificateInfoListBody struct {
	Req    *types.RetrieveCertificateInfoList         `xml:"urn:vim25 RetrieveCertificateInfoList,omitempty"`
	Res    *types.RetrieveCertificateInfoListResponse `xml:"RetrieveCertificateInfoListResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveCertificateInfoListBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveCertificateInfoList(ctx context.Context, r soap.RoundTripper, req *types.RetrieveCertificateInfoList) (*types.RetrieveCertificateInfoListResponse, error) {
	var reqBody, resBody RetrieveCertificateInfoListBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveClientCertBody struct {
	Req    *types.RetrieveClientCert         `xml:"urn:vim25 RetrieveClientCert,omitempty"`
	Res    *types.RetrieveClientCertResponse `xml:"RetrieveClientCertResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveClientCert(ctx context.Context, r soap.RoundTripper, req *types.RetrieveClientCert) (*types.RetrieveClientCertResponse, error) {
	var reqBody, resBody RetrieveClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveClientCsrBody struct {
	Req    *types.RetrieveClientCsr         `xml:"urn:vim25 RetrieveClientCsr,omitempty"`
	Res    *types.RetrieveClientCsrResponse `xml:"RetrieveClientCsrResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveClientCsrBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveClientCsr(ctx context.Context, r soap.RoundTripper, req *types.RetrieveClientCsr) (*types.RetrieveClientCsrResponse, error) {
	var reqBody, resBody RetrieveClientCsrBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveDasAdvancedRuntimeInfoBody struct {
	Req    *types.RetrieveDasAdvancedRuntimeInfo         `xml:"urn:vim25 RetrieveDasAdvancedRuntimeInfo,omitempty"`
	Res    *types.RetrieveDasAdvancedRuntimeInfoResponse `xml:"RetrieveDasAdvancedRuntimeInfoResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDasAdvancedRuntimeInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveDasAdvancedRuntimeInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveDasAdvancedRuntimeInfo) (*types.RetrieveDasAdvancedRuntimeInfoResponse, error) {
	var reqBody, resBody RetrieveDasAdvancedRuntimeInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveDescriptionBody struct {
	Req    *types.RetrieveDescription         `xml:"urn:vim25 RetrieveDescription,omitempty"`
	Res    *types.RetrieveDescriptionResponse `xml:"RetrieveDescriptionResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDescriptionBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveDescription(ctx context.Context, r soap.RoundTripper, req *types.RetrieveDescription) (*types.RetrieveDescriptionResponse, error) {
	var reqBody, resBody RetrieveDescriptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveDiskPartitionInfoBody struct {
	Req    *types.RetrieveDiskPartitionInfo         `xml:"urn:vim25 RetrieveDiskPartitionInfo,omitempty"`
	Res    *types.RetrieveDiskPartitionInfoResponse `xml:"RetrieveDiskPartitionInfoResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDiskPartitionInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveDiskPartitionInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveDiskPartitionInfo) (*types.RetrieveDiskPartitionInfoResponse, error) {
	var reqBody, resBody RetrieveDiskPartitionInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveDynamicPassthroughInfoBody struct {
	Req    *types.RetrieveDynamicPassthroughInfo         `xml:"urn:vim25 RetrieveDynamicPassthroughInfo,omitempty"`
	Res    *types.RetrieveDynamicPassthroughInfoResponse `xml:"RetrieveDynamicPassthroughInfoResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDynamicPassthroughInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveDynamicPassthroughInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveDynamicPassthroughInfo) (*types.RetrieveDynamicPassthroughInfoResponse, error) {
	var reqBody, resBody RetrieveDynamicPassthroughInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveEntityPermissionsBody struct {
	Req    *types.RetrieveEntityPermissions         `xml:"urn:vim25 RetrieveEntityPermissions,omitempty"`
	Res    *types.RetrieveEntityPermissionsResponse `xml:"RetrieveEntityPermissionsResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveEntityPermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveEntityPermissions(ctx context.Context, r soap.RoundTripper, req *types.RetrieveEntityPermissions) (*types.RetrieveEntityPermissionsResponse, error) {
	var reqBody, resBody RetrieveEntityPermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveEntityScheduledTaskBody struct {
	Req    *types.RetrieveEntityScheduledTask         `xml:"urn:vim25 RetrieveEntityScheduledTask,omitempty"`
	Res    *types.RetrieveEntityScheduledTaskResponse `xml:"RetrieveEntityScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveEntityScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveEntityScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.RetrieveEntityScheduledTask) (*types.RetrieveEntityScheduledTaskResponse, error) {
	var reqBody, resBody RetrieveEntityScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveFreeEpcMemoryBody struct {
	Req    *types.RetrieveFreeEpcMemory         `xml:"urn:vim25 RetrieveFreeEpcMemory,omitempty"`
	Res    *types.RetrieveFreeEpcMemoryResponse `xml:"RetrieveFreeEpcMemoryResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveFreeEpcMemoryBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveFreeEpcMemory(ctx context.Context, r soap.RoundTripper, req *types.RetrieveFreeEpcMemory) (*types.RetrieveFreeEpcMemoryResponse, error) {
	var reqBody, resBody RetrieveFreeEpcMemoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveHardwareUptimeBody struct {
	Req    *types.RetrieveHardwareUptime         `xml:"urn:vim25 RetrieveHardwareUptime,omitempty"`
	Res    *types.RetrieveHardwareUptimeResponse `xml:"RetrieveHardwareUptimeResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHardwareUptimeBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveHardwareUptime(ctx context.Context, r soap.RoundTripper, req *types.RetrieveHardwareUptime) (*types.RetrieveHardwareUptimeResponse, error) {
	var reqBody, resBody RetrieveHardwareUptimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveHostAccessControlEntriesBody struct {
	Req    *types.RetrieveHostAccessControlEntries         `xml:"urn:vim25 RetrieveHostAccessControlEntries,omitempty"`
	Res    *types.RetrieveHostAccessControlEntriesResponse `xml:"RetrieveHostAccessControlEntriesResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHostAccessControlEntriesBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveHostAccessControlEntries(ctx context.Context, r soap.RoundTripper, req *types.RetrieveHostAccessControlEntries) (*types.RetrieveHostAccessControlEntriesResponse, error) {
	var reqBody, resBody RetrieveHostAccessControlEntriesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveHostCustomizationsBody struct {
	Req    *types.RetrieveHostCustomizations         `xml:"urn:vim25 RetrieveHostCustomizations,omitempty"`
	Res    *types.RetrieveHostCustomizationsResponse `xml:"RetrieveHostCustomizationsResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHostCustomizationsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveHostCustomizations(ctx context.Context, r soap.RoundTripper, req *types.RetrieveHostCustomizations) (*types.RetrieveHostCustomizationsResponse, error) {
	var reqBody, resBody RetrieveHostCustomizationsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveHostCustomizationsForProfileBody struct {
	Req    *types.RetrieveHostCustomizationsForProfile         `xml:"urn:vim25 RetrieveHostCustomizationsForProfile,omitempty"`
	Res    *types.RetrieveHostCustomizationsForProfileResponse `xml:"RetrieveHostCustomizationsForProfileResponse,omitempty"`
	Fault_ *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHostCustomizationsForProfileBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveHostCustomizationsForProfile(ctx context.Context, r soap.RoundTripper, req *types.RetrieveHostCustomizationsForProfile) (*types.RetrieveHostCustomizationsForProfileResponse, error) {
	var reqBody, resBody RetrieveHostCustomizationsForProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveHostSpecificationBody struct {
	Req    *types.RetrieveHostSpecification         `xml:"urn:vim25 RetrieveHostSpecification,omitempty"`
	Res    *types.RetrieveHostSpecificationResponse `xml:"RetrieveHostSpecificationResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHostSpecificationBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveHostSpecification(ctx context.Context, r soap.RoundTripper, req *types.RetrieveHostSpecification) (*types.RetrieveHostSpecificationResponse, error) {
	var reqBody, resBody RetrieveHostSpecificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveKmipServerCertBody struct {
	Req    *types.RetrieveKmipServerCert         `xml:"urn:vim25 RetrieveKmipServerCert,omitempty"`
	Res    *types.RetrieveKmipServerCertResponse `xml:"RetrieveKmipServerCertResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveKmipServerCertBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveKmipServerCert(ctx context.Context, r soap.RoundTripper, req *types.RetrieveKmipServerCert) (*types.RetrieveKmipServerCertResponse, error) {
	var reqBody, resBody RetrieveKmipServerCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveKmipServersStatus_TaskBody struct {
	Req    *types.RetrieveKmipServersStatus_Task         `xml:"urn:vim25 RetrieveKmipServersStatus_Task,omitempty"`
	Res    *types.RetrieveKmipServersStatus_TaskResponse `xml:"RetrieveKmipServersStatus_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveKmipServersStatus_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveKmipServersStatus_Task(ctx context.Context, r soap.RoundTripper, req *types.RetrieveKmipServersStatus_Task) (*types.RetrieveKmipServersStatus_TaskResponse, error) {
	var reqBody, resBody RetrieveKmipServersStatus_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveObjectScheduledTaskBody struct {
	Req    *types.RetrieveObjectScheduledTask         `xml:"urn:vim25 RetrieveObjectScheduledTask,omitempty"`
	Res    *types.RetrieveObjectScheduledTaskResponse `xml:"RetrieveObjectScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveObjectScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveObjectScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.RetrieveObjectScheduledTask) (*types.RetrieveObjectScheduledTaskResponse, error) {
	var reqBody, resBody RetrieveObjectScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveProductComponentsBody struct {
	Req    *types.RetrieveProductComponents         `xml:"urn:vim25 RetrieveProductComponents,omitempty"`
	Res    *types.RetrieveProductComponentsResponse `xml:"RetrieveProductComponentsResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveProductComponentsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveProductComponents(ctx context.Context, r soap.RoundTripper, req *types.RetrieveProductComponents) (*types.RetrieveProductComponentsResponse, error) {
	var reqBody, resBody RetrieveProductComponentsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrievePropertiesBody struct {
	Req    *types.RetrieveProperties         `xml:"urn:vim25 RetrieveProperties,omitempty"`
	Res    *types.RetrievePropertiesResponse `xml:"RetrievePropertiesResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrievePropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveProperties(ctx context.Context, r soap.RoundTripper, req *types.RetrieveProperties) (*types.RetrievePropertiesResponse, error) {
	var reqBody, resBody RetrievePropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrievePropertiesExBody struct {
	Req    *types.RetrievePropertiesEx         `xml:"urn:vim25 RetrievePropertiesEx,omitempty"`
	Res    *types.RetrievePropertiesExResponse `xml:"RetrievePropertiesExResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrievePropertiesExBody) Fault() *soap.Fault { return b.Fault_ }

func RetrievePropertiesEx(ctx context.Context, r soap.RoundTripper, req *types.RetrievePropertiesEx) (*types.RetrievePropertiesExResponse, error) {
	var reqBody, resBody RetrievePropertiesExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveRolePermissionsBody struct {
	Req    *types.RetrieveRolePermissions         `xml:"urn:vim25 RetrieveRolePermissions,omitempty"`
	Res    *types.RetrieveRolePermissionsResponse `xml:"RetrieveRolePermissionsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveRolePermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveRolePermissions(ctx context.Context, r soap.RoundTripper, req *types.RetrieveRolePermissions) (*types.RetrieveRolePermissionsResponse, error) {
	var reqBody, resBody RetrieveRolePermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveSelfSignedClientCertBody struct {
	Req    *types.RetrieveSelfSignedClientCert         `xml:"urn:vim25 RetrieveSelfSignedClientCert,omitempty"`
	Res    *types.RetrieveSelfSignedClientCertResponse `xml:"RetrieveSelfSignedClientCertResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveSelfSignedClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveSelfSignedClientCert(ctx context.Context, r soap.RoundTripper, req *types.RetrieveSelfSignedClientCert) (*types.RetrieveSelfSignedClientCertResponse, error) {
	var reqBody, resBody RetrieveSelfSignedClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveServiceContentBody struct {
	Req    *types.RetrieveServiceContent         `xml:"urn:vim25 RetrieveServiceContent,omitempty"`
	Res    *types.RetrieveServiceContentResponse `xml:"RetrieveServiceContentResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveServiceContentBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveServiceContent(ctx context.Context, r soap.RoundTripper, req *types.RetrieveServiceContent) (*types.RetrieveServiceContentResponse, error) {
	var reqBody, resBody RetrieveServiceContentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveServiceProviderEntitiesBody struct {
	Req    *types.RetrieveServiceProviderEntities         `xml:"urn:vim25 RetrieveServiceProviderEntities,omitempty"`
	Res    *types.RetrieveServiceProviderEntitiesResponse `xml:"RetrieveServiceProviderEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveServiceProviderEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveServiceProviderEntities(ctx context.Context, r soap.RoundTripper, req *types.RetrieveServiceProviderEntities) (*types.RetrieveServiceProviderEntitiesResponse, error) {
	var reqBody, resBody RetrieveServiceProviderEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveSnapshotDetailsBody struct {
	Req    *types.RetrieveSnapshotDetails         `xml:"urn:vim25 RetrieveSnapshotDetails,omitempty"`
	Res    *types.RetrieveSnapshotDetailsResponse `xml:"RetrieveSnapshotDetailsResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveSnapshotDetailsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveSnapshotDetails(ctx context.Context, r soap.RoundTripper, req *types.RetrieveSnapshotDetails) (*types.RetrieveSnapshotDetailsResponse, error) {
	var reqBody, resBody RetrieveSnapshotDetailsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveSnapshotInfoBody struct {
	Req    *types.RetrieveSnapshotInfo         `xml:"urn:vim25 RetrieveSnapshotInfo,omitempty"`
	Res    *types.RetrieveSnapshotInfoResponse `xml:"RetrieveSnapshotInfoResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveSnapshotInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveSnapshotInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveSnapshotInfo) (*types.RetrieveSnapshotInfoResponse, error) {
	var reqBody, resBody RetrieveSnapshotInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveUserGroupsBody struct {
	Req    *types.RetrieveUserGroups         `xml:"urn:vim25 RetrieveUserGroups,omitempty"`
	Res    *types.RetrieveUserGroupsResponse `xml:"RetrieveUserGroupsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveUserGroupsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveUserGroups(ctx context.Context, r soap.RoundTripper, req *types.RetrieveUserGroups) (*types.RetrieveUserGroupsResponse, error) {
	var reqBody, resBody RetrieveUserGroupsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVStorageInfrastructureObjectPolicyBody struct {
	Req    *types.RetrieveVStorageInfrastructureObjectPolicy         `xml:"urn:vim25 RetrieveVStorageInfrastructureObjectPolicy,omitempty"`
	Res    *types.RetrieveVStorageInfrastructureObjectPolicyResponse `xml:"RetrieveVStorageInfrastructureObjectPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVStorageInfrastructureObjectPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVStorageInfrastructureObjectPolicy(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVStorageInfrastructureObjectPolicy) (*types.RetrieveVStorageInfrastructureObjectPolicyResponse, error) {
	var reqBody, resBody RetrieveVStorageInfrastructureObjectPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVStorageObjectBody struct {
	Req    *types.RetrieveVStorageObject         `xml:"urn:vim25 RetrieveVStorageObject,omitempty"`
	Res    *types.RetrieveVStorageObjectResponse `xml:"RetrieveVStorageObjectResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVStorageObjectBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVStorageObject(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVStorageObject) (*types.RetrieveVStorageObjectResponse, error) {
	var reqBody, resBody RetrieveVStorageObjectBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVStorageObjectAssociationsBody struct {
	Req    *types.RetrieveVStorageObjectAssociations         `xml:"urn:vim25 RetrieveVStorageObjectAssociations,omitempty"`
	Res    *types.RetrieveVStorageObjectAssociationsResponse `xml:"RetrieveVStorageObjectAssociationsResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVStorageObjectAssociationsBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVStorageObjectAssociations(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVStorageObjectAssociations) (*types.RetrieveVStorageObjectAssociationsResponse, error) {
	var reqBody, resBody RetrieveVStorageObjectAssociationsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVStorageObjectStateBody struct {
	Req    *types.RetrieveVStorageObjectState         `xml:"urn:vim25 RetrieveVStorageObjectState,omitempty"`
	Res    *types.RetrieveVStorageObjectStateResponse `xml:"RetrieveVStorageObjectStateResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVStorageObjectStateBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVStorageObjectState(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVStorageObjectState) (*types.RetrieveVStorageObjectStateResponse, error) {
	var reqBody, resBody RetrieveVStorageObjectStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVendorDeviceGroupInfoBody struct {
	Req    *types.RetrieveVendorDeviceGroupInfo         `xml:"urn:vim25 RetrieveVendorDeviceGroupInfo,omitempty"`
	Res    *types.RetrieveVendorDeviceGroupInfoResponse `xml:"RetrieveVendorDeviceGroupInfoResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVendorDeviceGroupInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVendorDeviceGroupInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVendorDeviceGroupInfo) (*types.RetrieveVendorDeviceGroupInfoResponse, error) {
	var reqBody, resBody RetrieveVendorDeviceGroupInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVgpuDeviceInfoBody struct {
	Req    *types.RetrieveVgpuDeviceInfo         `xml:"urn:vim25 RetrieveVgpuDeviceInfo,omitempty"`
	Res    *types.RetrieveVgpuDeviceInfoResponse `xml:"RetrieveVgpuDeviceInfoResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVgpuDeviceInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVgpuDeviceInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVgpuDeviceInfo) (*types.RetrieveVgpuDeviceInfoResponse, error) {
	var reqBody, resBody RetrieveVgpuDeviceInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RetrieveVgpuProfileInfoBody struct {
	Req    *types.RetrieveVgpuProfileInfo         `xml:"urn:vim25 RetrieveVgpuProfileInfo,omitempty"`
	Res    *types.RetrieveVgpuProfileInfoResponse `xml:"RetrieveVgpuProfileInfoResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveVgpuProfileInfoBody) Fault() *soap.Fault { return b.Fault_ }

func RetrieveVgpuProfileInfo(ctx context.Context, r soap.RoundTripper, req *types.RetrieveVgpuProfileInfo) (*types.RetrieveVgpuProfileInfoResponse, error) {
	var reqBody, resBody RetrieveVgpuProfileInfoBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RevertToCurrentSnapshot_TaskBody struct {
	Req    *types.RevertToCurrentSnapshot_Task         `xml:"urn:vim25 RevertToCurrentSnapshot_Task,omitempty"`
	Res    *types.RevertToCurrentSnapshot_TaskResponse `xml:"RevertToCurrentSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertToCurrentSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RevertToCurrentSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.RevertToCurrentSnapshot_Task) (*types.RevertToCurrentSnapshot_TaskResponse, error) {
	var reqBody, resBody RevertToCurrentSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RevertToSnapshot_TaskBody struct {
	Req    *types.RevertToSnapshot_Task         `xml:"urn:vim25 RevertToSnapshot_Task,omitempty"`
	Res    *types.RevertToSnapshot_TaskResponse `xml:"RevertToSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertToSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RevertToSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.RevertToSnapshot_Task) (*types.RevertToSnapshot_TaskResponse, error) {
	var reqBody, resBody RevertToSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RevertVStorageObjectEx_TaskBody struct {
	Req    *types.RevertVStorageObjectEx_Task         `xml:"urn:vim25 RevertVStorageObjectEx_Task,omitempty"`
	Res    *types.RevertVStorageObjectEx_TaskResponse `xml:"RevertVStorageObjectEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertVStorageObjectEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RevertVStorageObjectEx_Task(ctx context.Context, r soap.RoundTripper, req *types.RevertVStorageObjectEx_Task) (*types.RevertVStorageObjectEx_TaskResponse, error) {
	var reqBody, resBody RevertVStorageObjectEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RevertVStorageObject_TaskBody struct {
	Req    *types.RevertVStorageObject_Task         `xml:"urn:vim25 RevertVStorageObject_Task,omitempty"`
	Res    *types.RevertVStorageObject_TaskResponse `xml:"RevertVStorageObject_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertVStorageObject_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func RevertVStorageObject_Task(ctx context.Context, r soap.RoundTripper, req *types.RevertVStorageObject_Task) (*types.RevertVStorageObject_TaskResponse, error) {
	var reqBody, resBody RevertVStorageObject_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RewindCollectorBody struct {
	Req    *types.RewindCollector         `xml:"urn:vim25 RewindCollector,omitempty"`
	Res    *types.RewindCollectorResponse `xml:"RewindCollectorResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RewindCollectorBody) Fault() *soap.Fault { return b.Fault_ }

func RewindCollector(ctx context.Context, r soap.RoundTripper, req *types.RewindCollector) (*types.RewindCollectorResponse, error) {
	var reqBody, resBody RewindCollectorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RunScheduledTaskBody struct {
	Req    *types.RunScheduledTask         `xml:"urn:vim25 RunScheduledTask,omitempty"`
	Res    *types.RunScheduledTaskResponse `xml:"RunScheduledTaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RunScheduledTaskBody) Fault() *soap.Fault { return b.Fault_ }

func RunScheduledTask(ctx context.Context, r soap.RoundTripper, req *types.RunScheduledTask) (*types.RunScheduledTaskResponse, error) {
	var reqBody, resBody RunScheduledTaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RunVsanPhysicalDiskDiagnosticsBody struct {
	Req    *types.RunVsanPhysicalDiskDiagnostics         `xml:"urn:vim25 RunVsanPhysicalDiskDiagnostics,omitempty"`
	Res    *types.RunVsanPhysicalDiskDiagnosticsResponse `xml:"RunVsanPhysicalDiskDiagnosticsResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RunVsanPhysicalDiskDiagnosticsBody) Fault() *soap.Fault { return b.Fault_ }

func RunVsanPhysicalDiskDiagnostics(ctx context.Context, r soap.RoundTripper, req *types.RunVsanPhysicalDiskDiagnostics) (*types.RunVsanPhysicalDiskDiagnosticsResponse, error) {
	var reqBody, resBody RunVsanPhysicalDiskDiagnosticsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ScanHostPatchV2_TaskBody struct {
	Req    *types.ScanHostPatchV2_Task         `xml:"urn:vim25 ScanHostPatchV2_Task,omitempty"`
	Res    *types.ScanHostPatchV2_TaskResponse `xml:"ScanHostPatchV2_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScanHostPatchV2_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ScanHostPatchV2_Task(ctx context.Context, r soap.RoundTripper, req *types.ScanHostPatchV2_Task) (*types.ScanHostPatchV2_TaskResponse, error) {
	var reqBody, resBody ScanHostPatchV2_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ScanHostPatch_TaskBody struct {
	Req    *types.ScanHostPatch_Task         `xml:"urn:vim25 ScanHostPatch_Task,omitempty"`
	Res    *types.ScanHostPatch_TaskResponse `xml:"ScanHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScanHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ScanHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.ScanHostPatch_Task) (*types.ScanHostPatch_TaskResponse, error) {
	var reqBody, resBody ScanHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ScheduleReconcileDatastoreInventoryBody struct {
	Req    *types.ScheduleReconcileDatastoreInventory         `xml:"urn:vim25 ScheduleReconcileDatastoreInventory,omitempty"`
	Res    *types.ScheduleReconcileDatastoreInventoryResponse `xml:"ScheduleReconcileDatastoreInventoryResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScheduleReconcileDatastoreInventoryBody) Fault() *soap.Fault { return b.Fault_ }

func ScheduleReconcileDatastoreInventory(ctx context.Context, r soap.RoundTripper, req *types.ScheduleReconcileDatastoreInventory) (*types.ScheduleReconcileDatastoreInventoryResponse, error) {
	var reqBody, resBody ScheduleReconcileDatastoreInventoryBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SearchDatastoreSubFolders_TaskBody struct {
	Req    *types.SearchDatastoreSubFolders_Task         `xml:"urn:vim25 SearchDatastoreSubFolders_Task,omitempty"`
	Res    *types.SearchDatastoreSubFolders_TaskResponse `xml:"SearchDatastoreSubFolders_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SearchDatastoreSubFolders_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SearchDatastoreSubFolders_Task(ctx context.Context, r soap.RoundTripper, req *types.SearchDatastoreSubFolders_Task) (*types.SearchDatastoreSubFolders_TaskResponse, error) {
	var reqBody, resBody SearchDatastoreSubFolders_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SearchDatastore_TaskBody struct {
	Req    *types.SearchDatastore_Task         `xml:"urn:vim25 SearchDatastore_Task,omitempty"`
	Res    *types.SearchDatastore_TaskResponse `xml:"SearchDatastore_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SearchDatastore_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SearchDatastore_Task(ctx context.Context, r soap.RoundTripper, req *types.SearchDatastore_Task) (*types.SearchDatastore_TaskResponse, error) {
	var reqBody, resBody SearchDatastore_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SelectActivePartitionBody struct {
	Req    *types.SelectActivePartition         `xml:"urn:vim25 SelectActivePartition,omitempty"`
	Res    *types.SelectActivePartitionResponse `xml:"SelectActivePartitionResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectActivePartitionBody) Fault() *soap.Fault { return b.Fault_ }

func SelectActivePartition(ctx context.Context, r soap.RoundTripper, req *types.SelectActivePartition) (*types.SelectActivePartitionResponse, error) {
	var reqBody, resBody SelectActivePartitionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SelectVnicBody struct {
	Req    *types.SelectVnic         `xml:"urn:vim25 SelectVnic,omitempty"`
	Res    *types.SelectVnicResponse `xml:"SelectVnicResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectVnicBody) Fault() *soap.Fault { return b.Fault_ }

func SelectVnic(ctx context.Context, r soap.RoundTripper, req *types.SelectVnic) (*types.SelectVnicResponse, error) {
	var reqBody, resBody SelectVnicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SelectVnicForNicTypeBody struct {
	Req    *types.SelectVnicForNicType         `xml:"urn:vim25 SelectVnicForNicType,omitempty"`
	Res    *types.SelectVnicForNicTypeResponse `xml:"SelectVnicForNicTypeResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectVnicForNicTypeBody) Fault() *soap.Fault { return b.Fault_ }

func SelectVnicForNicType(ctx context.Context, r soap.RoundTripper, req *types.SelectVnicForNicType) (*types.SelectVnicForNicTypeResponse, error) {
	var reqBody, resBody SelectVnicForNicTypeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SendNMIBody struct {
	Req    *types.SendNMI         `xml:"urn:vim25 SendNMI,omitempty"`
	Res    *types.SendNMIResponse `xml:"SendNMIResponse,omitempty"`
	Fault_ *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SendNMIBody) Fault() *soap.Fault { return b.Fault_ }

func SendNMI(ctx context.Context, r soap.RoundTripper, req *types.SendNMI) (*types.SendNMIResponse, error) {
	var reqBody, resBody SendNMIBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SendTestNotificationBody struct {
	Req    *types.SendTestNotification         `xml:"urn:vim25 SendTestNotification,omitempty"`
	Res    *types.SendTestNotificationResponse `xml:"SendTestNotificationResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SendTestNotificationBody) Fault() *soap.Fault { return b.Fault_ }

func SendTestNotification(ctx context.Context, r soap.RoundTripper, req *types.SendTestNotification) (*types.SendTestNotificationResponse, error) {
	var reqBody, resBody SendTestNotificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SessionIsActiveBody struct {
	Req    *types.SessionIsActive         `xml:"urn:vim25 SessionIsActive,omitempty"`
	Res    *types.SessionIsActiveResponse `xml:"SessionIsActiveResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SessionIsActiveBody) Fault() *soap.Fault { return b.Fault_ }

func SessionIsActive(ctx context.Context, r soap.RoundTripper, req *types.SessionIsActive) (*types.SessionIsActiveResponse, error) {
	var reqBody, resBody SessionIsActiveBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetCollectorPageSizeBody struct {
	Req    *types.SetCollectorPageSize         `xml:"urn:vim25 SetCollectorPageSize,omitempty"`
	Res    *types.SetCollectorPageSizeResponse `xml:"SetCollectorPageSizeResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetCollectorPageSizeBody) Fault() *soap.Fault { return b.Fault_ }

func SetCollectorPageSize(ctx context.Context, r soap.RoundTripper, req *types.SetCollectorPageSize) (*types.SetCollectorPageSizeResponse, error) {
	var reqBody, resBody SetCollectorPageSizeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetCryptoModeBody struct {
	Req    *types.SetCryptoMode         `xml:"urn:vim25 SetCryptoMode,omitempty"`
	Res    *types.SetCryptoModeResponse `xml:"SetCryptoModeResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetCryptoModeBody) Fault() *soap.Fault { return b.Fault_ }

func SetCryptoMode(ctx context.Context, r soap.RoundTripper, req *types.SetCryptoMode) (*types.SetCryptoModeResponse, error) {
	var reqBody, resBody SetCryptoModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetDefaultKmsClusterBody struct {
	Req    *types.SetDefaultKmsCluster         `xml:"urn:vim25 SetDefaultKmsCluster,omitempty"`
	Res    *types.SetDefaultKmsClusterResponse `xml:"SetDefaultKmsClusterResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetDefaultKmsClusterBody) Fault() *soap.Fault { return b.Fault_ }

func SetDefaultKmsCluster(ctx context.Context, r soap.RoundTripper, req *types.SetDefaultKmsCluster) (*types.SetDefaultKmsClusterResponse, error) {
	var reqBody, resBody SetDefaultKmsClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetDisplayTopologyBody struct {
	Req    *types.SetDisplayTopology         `xml:"urn:vim25 SetDisplayTopology,omitempty"`
	Res    *types.SetDisplayTopologyResponse `xml:"SetDisplayTopologyResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetDisplayTopologyBody) Fault() *soap.Fault { return b.Fault_ }

func SetDisplayTopology(ctx context.Context, r soap.RoundTripper, req *types.SetDisplayTopology) (*types.SetDisplayTopologyResponse, error) {
	var reqBody, resBody SetDisplayTopologyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetEntityPermissionsBody struct {
	Req    *types.SetEntityPermissions         `xml:"urn:vim25 SetEntityPermissions,omitempty"`
	Res    *types.SetEntityPermissionsResponse `xml:"SetEntityPermissionsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetEntityPermissionsBody) Fault() *soap.Fault { return b.Fault_ }

func SetEntityPermissions(ctx context.Context, r soap.RoundTripper, req *types.SetEntityPermissions) (*types.SetEntityPermissionsResponse, error) {
	var reqBody, resBody SetEntityPermissionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetExtensionCertificateBody struct {
	Req    *types.SetExtensionCertificate         `xml:"urn:vim25 SetExtensionCertificate,omitempty"`
	Res    *types.SetExtensionCertificateResponse `xml:"SetExtensionCertificateResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetExtensionCertificateBody) Fault() *soap.Fault { return b.Fault_ }

func SetExtensionCertificate(ctx context.Context, r soap.RoundTripper, req *types.SetExtensionCertificate) (*types.SetExtensionCertificateResponse, error) {
	var reqBody, resBody SetExtensionCertificateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetFieldBody struct {
	Req    *types.SetField         `xml:"urn:vim25 SetField,omitempty"`
	Res    *types.SetFieldResponse `xml:"SetFieldResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetFieldBody) Fault() *soap.Fault { return b.Fault_ }

func SetField(ctx context.Context, r soap.RoundTripper, req *types.SetField) (*types.SetFieldResponse, error) {
	var reqBody, resBody SetFieldBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetKeyCustomAttributesBody struct {
	Req    *types.SetKeyCustomAttributes         `xml:"urn:vim25 SetKeyCustomAttributes,omitempty"`
	Res    *types.SetKeyCustomAttributesResponse `xml:"SetKeyCustomAttributesResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetKeyCustomAttributesBody) Fault() *soap.Fault { return b.Fault_ }

func SetKeyCustomAttributes(ctx context.Context, r soap.RoundTripper, req *types.SetKeyCustomAttributes) (*types.SetKeyCustomAttributesResponse, error) {
	var reqBody, resBody SetKeyCustomAttributesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetLicenseEditionBody struct {
	Req    *types.SetLicenseEdition         `xml:"urn:vim25 SetLicenseEdition,omitempty"`
	Res    *types.SetLicenseEditionResponse `xml:"SetLicenseEditionResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetLicenseEditionBody) Fault() *soap.Fault { return b.Fault_ }

func SetLicenseEdition(ctx context.Context, r soap.RoundTripper, req *types.SetLicenseEdition) (*types.SetLicenseEditionResponse, error) {
	var reqBody, resBody SetLicenseEditionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetLocaleBody struct {
	Req    *types.SetLocale         `xml:"urn:vim25 SetLocale,omitempty"`
	Res    *types.SetLocaleResponse `xml:"SetLocaleResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetLocaleBody) Fault() *soap.Fault { return b.Fault_ }

func SetLocale(ctx context.Context, r soap.RoundTripper, req *types.SetLocale) (*types.SetLocaleResponse, error) {
	var reqBody, resBody SetLocaleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetMaxQueueDepthBody struct {
	Req    *types.SetMaxQueueDepth         `xml:"urn:vim25 SetMaxQueueDepth,omitempty"`
	Res    *types.SetMaxQueueDepthResponse `xml:"SetMaxQueueDepthResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetMaxQueueDepthBody) Fault() *soap.Fault { return b.Fault_ }

func SetMaxQueueDepth(ctx context.Context, r soap.RoundTripper, req *types.SetMaxQueueDepth) (*types.SetMaxQueueDepthResponse, error) {
	var reqBody, resBody SetMaxQueueDepthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetMultipathLunPolicyBody struct {
	Req    *types.SetMultipathLunPolicy         `xml:"urn:vim25 SetMultipathLunPolicy,omitempty"`
	Res    *types.SetMultipathLunPolicyResponse `xml:"SetMultipathLunPolicyResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetMultipathLunPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func SetMultipathLunPolicy(ctx context.Context, r soap.RoundTripper, req *types.SetMultipathLunPolicy) (*types.SetMultipathLunPolicyResponse, error) {
	var reqBody, resBody SetMultipathLunPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetNFSUserBody struct {
	Req    *types.SetNFSUser         `xml:"urn:vim25 SetNFSUser,omitempty"`
	Res    *types.SetNFSUserResponse `xml:"SetNFSUserResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetNFSUserBody) Fault() *soap.Fault { return b.Fault_ }

func SetNFSUser(ctx context.Context, r soap.RoundTripper, req *types.SetNFSUser) (*types.SetNFSUserResponse, error) {
	var reqBody, resBody SetNFSUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetPublicKeyBody struct {
	Req    *types.SetPublicKey         `xml:"urn:vim25 SetPublicKey,omitempty"`
	Res    *types.SetPublicKeyResponse `xml:"SetPublicKeyResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetPublicKeyBody) Fault() *soap.Fault { return b.Fault_ }

func SetPublicKey(ctx context.Context, r soap.RoundTripper, req *types.SetPublicKey) (*types.SetPublicKeyResponse, error) {
	var reqBody, resBody SetPublicKeyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetRegistryValueInGuestBody struct {
	Req    *types.SetRegistryValueInGuest         `xml:"urn:vim25 SetRegistryValueInGuest,omitempty"`
	Res    *types.SetRegistryValueInGuestResponse `xml:"SetRegistryValueInGuestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetRegistryValueInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func SetRegistryValueInGuest(ctx context.Context, r soap.RoundTripper, req *types.SetRegistryValueInGuest) (*types.SetRegistryValueInGuestResponse, error) {
	var reqBody, resBody SetRegistryValueInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetScreenResolutionBody struct {
	Req    *types.SetScreenResolution         `xml:"urn:vim25 SetScreenResolution,omitempty"`
	Res    *types.SetScreenResolutionResponse `xml:"SetScreenResolutionResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetScreenResolutionBody) Fault() *soap.Fault { return b.Fault_ }

func SetScreenResolution(ctx context.Context, r soap.RoundTripper, req *types.SetScreenResolution) (*types.SetScreenResolutionResponse, error) {
	var reqBody, resBody SetScreenResolutionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetServiceAccountBody struct {
	Req    *types.SetServiceAccount         `xml:"urn:vim25 SetServiceAccount,omitempty"`
	Res    *types.SetServiceAccountResponse `xml:"SetServiceAccountResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetServiceAccountBody) Fault() *soap.Fault { return b.Fault_ }

func SetServiceAccount(ctx context.Context, r soap.RoundTripper, req *types.SetServiceAccount) (*types.SetServiceAccountResponse, error) {
	var reqBody, resBody SetServiceAccountBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetTaskDescriptionBody struct {
	Req    *types.SetTaskDescription         `xml:"urn:vim25 SetTaskDescription,omitempty"`
	Res    *types.SetTaskDescriptionResponse `xml:"SetTaskDescriptionResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetTaskDescriptionBody) Fault() *soap.Fault { return b.Fault_ }

func SetTaskDescription(ctx context.Context, r soap.RoundTripper, req *types.SetTaskDescription) (*types.SetTaskDescriptionResponse, error) {
	var reqBody, resBody SetTaskDescriptionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetTaskStateBody struct {
	Req    *types.SetTaskState         `xml:"urn:vim25 SetTaskState,omitempty"`
	Res    *types.SetTaskStateResponse `xml:"SetTaskStateResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetTaskStateBody) Fault() *soap.Fault { return b.Fault_ }

func SetTaskState(ctx context.Context, r soap.RoundTripper, req *types.SetTaskState) (*types.SetTaskStateResponse, error) {
	var reqBody, resBody SetTaskStateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetVStorageObjectControlFlagsBody struct {
	Req    *types.SetVStorageObjectControlFlags         `xml:"urn:vim25 SetVStorageObjectControlFlags,omitempty"`
	Res    *types.SetVStorageObjectControlFlagsResponse `xml:"SetVStorageObjectControlFlagsResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetVStorageObjectControlFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func SetVStorageObjectControlFlags(ctx context.Context, r soap.RoundTripper, req *types.SetVStorageObjectControlFlags) (*types.SetVStorageObjectControlFlagsResponse, error) {
	var reqBody, resBody SetVStorageObjectControlFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetVirtualDiskUuidBody struct {
	Req    *types.SetVirtualDiskUuid         `xml:"urn:vim25 SetVirtualDiskUuid,omitempty"`
	Res    *types.SetVirtualDiskUuidResponse `xml:"SetVirtualDiskUuidResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetVirtualDiskUuidBody) Fault() *soap.Fault { return b.Fault_ }

func SetVirtualDiskUuid(ctx context.Context, r soap.RoundTripper, req *types.SetVirtualDiskUuid) (*types.SetVirtualDiskUuidResponse, error) {
	var reqBody, resBody SetVirtualDiskUuidBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetVirtualDiskUuidEx_TaskBody struct {
	Req    *types.SetVirtualDiskUuidEx_Task         `xml:"urn:vim25 SetVirtualDiskUuidEx_Task,omitempty"`
	Res    *types.SetVirtualDiskUuidEx_TaskResponse `xml:"SetVirtualDiskUuidEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetVirtualDiskUuidEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SetVirtualDiskUuidEx_Task(ctx context.Context, r soap.RoundTripper, req *types.SetVirtualDiskUuidEx_Task) (*types.SetVirtualDiskUuidEx_TaskResponse, error) {
	var reqBody, resBody SetVirtualDiskUuidEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ShrinkVirtualDisk_TaskBody struct {
	Req    *types.ShrinkVirtualDisk_Task         `xml:"urn:vim25 ShrinkVirtualDisk_Task,omitempty"`
	Res    *types.ShrinkVirtualDisk_TaskResponse `xml:"ShrinkVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShrinkVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ShrinkVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.ShrinkVirtualDisk_Task) (*types.ShrinkVirtualDisk_TaskResponse, error) {
	var reqBody, resBody ShrinkVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ShutdownGuestBody struct {
	Req    *types.ShutdownGuest         `xml:"urn:vim25 ShutdownGuest,omitempty"`
	Res    *types.ShutdownGuestResponse `xml:"ShutdownGuestResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShutdownGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ShutdownGuest(ctx context.Context, r soap.RoundTripper, req *types.ShutdownGuest) (*types.ShutdownGuestResponse, error) {
	var reqBody, resBody ShutdownGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ShutdownHost_TaskBody struct {
	Req    *types.ShutdownHost_Task         `xml:"urn:vim25 ShutdownHost_Task,omitempty"`
	Res    *types.ShutdownHost_TaskResponse `xml:"ShutdownHost_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShutdownHost_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ShutdownHost_Task(ctx context.Context, r soap.RoundTripper, req *types.ShutdownHost_Task) (*types.ShutdownHost_TaskResponse, error) {
	var reqBody, resBody ShutdownHost_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StageHostPatch_TaskBody struct {
	Req    *types.StageHostPatch_Task         `xml:"urn:vim25 StageHostPatch_Task,omitempty"`
	Res    *types.StageHostPatch_TaskResponse `xml:"StageHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StageHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StageHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.StageHostPatch_Task) (*types.StageHostPatch_TaskResponse, error) {
	var reqBody, resBody StageHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StampAllRulesWithUuid_TaskBody struct {
	Req    *types.StampAllRulesWithUuid_Task         `xml:"urn:vim25 StampAllRulesWithUuid_Task,omitempty"`
	Res    *types.StampAllRulesWithUuid_TaskResponse `xml:"StampAllRulesWithUuid_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StampAllRulesWithUuid_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StampAllRulesWithUuid_Task(ctx context.Context, r soap.RoundTripper, req *types.StampAllRulesWithUuid_Task) (*types.StampAllRulesWithUuid_TaskResponse, error) {
	var reqBody, resBody StampAllRulesWithUuid_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StandbyGuestBody struct {
	Req    *types.StandbyGuest         `xml:"urn:vim25 StandbyGuest,omitempty"`
	Res    *types.StandbyGuestResponse `xml:"StandbyGuestResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StandbyGuestBody) Fault() *soap.Fault { return b.Fault_ }

func StandbyGuest(ctx context.Context, r soap.RoundTripper, req *types.StandbyGuest) (*types.StandbyGuestResponse, error) {
	var reqBody, resBody StandbyGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartGuestNetwork_TaskBody struct {
	Req    *types.StartGuestNetwork_Task         `xml:"urn:vim25 StartGuestNetwork_Task,omitempty"`
	Res    *types.StartGuestNetwork_TaskResponse `xml:"StartGuestNetwork_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartGuestNetwork_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StartGuestNetwork_Task(ctx context.Context, r soap.RoundTripper, req *types.StartGuestNetwork_Task) (*types.StartGuestNetwork_TaskResponse, error) {
	var reqBody, resBody StartGuestNetwork_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartProgramInGuestBody struct {
	Req    *types.StartProgramInGuest         `xml:"urn:vim25 StartProgramInGuest,omitempty"`
	Res    *types.StartProgramInGuestResponse `xml:"StartProgramInGuestResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartProgramInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func StartProgramInGuest(ctx context.Context, r soap.RoundTripper, req *types.StartProgramInGuest) (*types.StartProgramInGuestResponse, error) {
	var reqBody, resBody StartProgramInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartRecording_TaskBody struct {
	Req    *types.StartRecording_Task         `xml:"urn:vim25 StartRecording_Task,omitempty"`
	Res    *types.StartRecording_TaskResponse `xml:"StartRecording_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartRecording_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StartRecording_Task(ctx context.Context, r soap.RoundTripper, req *types.StartRecording_Task) (*types.StartRecording_TaskResponse, error) {
	var reqBody, resBody StartRecording_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartReplaying_TaskBody struct {
	Req    *types.StartReplaying_Task         `xml:"urn:vim25 StartReplaying_Task,omitempty"`
	Res    *types.StartReplaying_TaskResponse `xml:"StartReplaying_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartReplaying_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StartReplaying_Task(ctx context.Context, r soap.RoundTripper, req *types.StartReplaying_Task) (*types.StartReplaying_TaskResponse, error) {
	var reqBody, resBody StartReplaying_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartServiceBody struct {
	Req    *types.StartService         `xml:"urn:vim25 StartService,omitempty"`
	Res    *types.StartServiceResponse `xml:"StartServiceResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartServiceBody) Fault() *soap.Fault { return b.Fault_ }

func StartService(ctx context.Context, r soap.RoundTripper, req *types.StartService) (*types.StartServiceResponse, error) {
	var reqBody, resBody StartServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StopRecording_TaskBody struct {
	Req    *types.StopRecording_Task         `xml:"urn:vim25 StopRecording_Task,omitempty"`
	Res    *types.StopRecording_TaskResponse `xml:"StopRecording_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopRecording_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StopRecording_Task(ctx context.Context, r soap.RoundTripper, req *types.StopRecording_Task) (*types.StopRecording_TaskResponse, error) {
	var reqBody, resBody StopRecording_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StopReplaying_TaskBody struct {
	Req    *types.StopReplaying_Task         `xml:"urn:vim25 StopReplaying_Task,omitempty"`
	Res    *types.StopReplaying_TaskResponse `xml:"StopReplaying_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopReplaying_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func StopReplaying_Task(ctx context.Context, r soap.RoundTripper, req *types.StopReplaying_Task) (*types.StopReplaying_TaskResponse, error) {
	var reqBody, resBody StopReplaying_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StopServiceBody struct {
	Req    *types.StopService         `xml:"urn:vim25 StopService,omitempty"`
	Res    *types.StopServiceResponse `xml:"StopServiceResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopServiceBody) Fault() *soap.Fault { return b.Fault_ }

func StopService(ctx context.Context, r soap.RoundTripper, req *types.StopService) (*types.StopServiceResponse, error) {
	var reqBody, resBody StopServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SuspendVApp_TaskBody struct {
	Req    *types.SuspendVApp_Task         `xml:"urn:vim25 SuspendVApp_Task,omitempty"`
	Res    *types.SuspendVApp_TaskResponse `xml:"SuspendVApp_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SuspendVApp_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SuspendVApp_Task(ctx context.Context, r soap.RoundTripper, req *types.SuspendVApp_Task) (*types.SuspendVApp_TaskResponse, error) {
	var reqBody, resBody SuspendVApp_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SuspendVM_TaskBody struct {
	Req    *types.SuspendVM_Task         `xml:"urn:vim25 SuspendVM_Task,omitempty"`
	Res    *types.SuspendVM_TaskResponse `xml:"SuspendVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SuspendVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SuspendVM_Task(ctx context.Context, r soap.RoundTripper, req *types.SuspendVM_Task) (*types.SuspendVM_TaskResponse, error) {
	var reqBody, resBody SuspendVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TerminateFaultTolerantVM_TaskBody struct {
	Req    *types.TerminateFaultTolerantVM_Task         `xml:"urn:vim25 TerminateFaultTolerantVM_Task,omitempty"`
	Res    *types.TerminateFaultTolerantVM_TaskResponse `xml:"TerminateFaultTolerantVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateFaultTolerantVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TerminateFaultTolerantVM_Task(ctx context.Context, r soap.RoundTripper, req *types.TerminateFaultTolerantVM_Task) (*types.TerminateFaultTolerantVM_TaskResponse, error) {
	var reqBody, resBody TerminateFaultTolerantVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TerminateProcessInGuestBody struct {
	Req    *types.TerminateProcessInGuest         `xml:"urn:vim25 TerminateProcessInGuest,omitempty"`
	Res    *types.TerminateProcessInGuestResponse `xml:"TerminateProcessInGuestResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateProcessInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func TerminateProcessInGuest(ctx context.Context, r soap.RoundTripper, req *types.TerminateProcessInGuest) (*types.TerminateProcessInGuestResponse, error) {
	var reqBody, resBody TerminateProcessInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TerminateSessionBody struct {
	Req    *types.TerminateSession         `xml:"urn:vim25 TerminateSession,omitempty"`
	Res    *types.TerminateSessionResponse `xml:"TerminateSessionResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateSessionBody) Fault() *soap.Fault { return b.Fault_ }

func TerminateSession(ctx context.Context, r soap.RoundTripper, req *types.TerminateSession) (*types.TerminateSessionResponse, error) {
	var reqBody, resBody TerminateSessionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TerminateVMBody struct {
	Req    *types.TerminateVM         `xml:"urn:vim25 TerminateVM,omitempty"`
	Res    *types.TerminateVMResponse `xml:"TerminateVMResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateVMBody) Fault() *soap.Fault { return b.Fault_ }

func TerminateVM(ctx context.Context, r soap.RoundTripper, req *types.TerminateVM) (*types.TerminateVMResponse, error) {
	var reqBody, resBody TerminateVMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TestTimeServiceBody struct {
	Req    *types.TestTimeService         `xml:"urn:vim25 TestTimeService,omitempty"`
	Res    *types.TestTimeServiceResponse `xml:"TestTimeServiceResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TestTimeServiceBody) Fault() *soap.Fault { return b.Fault_ }

func TestTimeService(ctx context.Context, r soap.RoundTripper, req *types.TestTimeService) (*types.TestTimeServiceResponse, error) {
	var reqBody, resBody TestTimeServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TurnDiskLocatorLedOff_TaskBody struct {
	Req    *types.TurnDiskLocatorLedOff_Task         `xml:"urn:vim25 TurnDiskLocatorLedOff_Task,omitempty"`
	Res    *types.TurnDiskLocatorLedOff_TaskResponse `xml:"TurnDiskLocatorLedOff_TaskResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TurnDiskLocatorLedOff_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TurnDiskLocatorLedOff_Task(ctx context.Context, r soap.RoundTripper, req *types.TurnDiskLocatorLedOff_Task) (*types.TurnDiskLocatorLedOff_TaskResponse, error) {
	var reqBody, resBody TurnDiskLocatorLedOff_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TurnDiskLocatorLedOn_TaskBody struct {
	Req    *types.TurnDiskLocatorLedOn_Task         `xml:"urn:vim25 TurnDiskLocatorLedOn_Task,omitempty"`
	Res    *types.TurnDiskLocatorLedOn_TaskResponse `xml:"TurnDiskLocatorLedOn_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TurnDiskLocatorLedOn_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TurnDiskLocatorLedOn_Task(ctx context.Context, r soap.RoundTripper, req *types.TurnDiskLocatorLedOn_Task) (*types.TurnDiskLocatorLedOn_TaskResponse, error) {
	var reqBody, resBody TurnDiskLocatorLedOn_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type TurnOffFaultToleranceForVM_TaskBody struct {
	Req    *types.TurnOffFaultToleranceForVM_Task         `xml:"urn:vim25 TurnOffFaultToleranceForVM_Task,omitempty"`
	Res    *types.TurnOffFaultToleranceForVM_TaskResponse `xml:"TurnOffFaultToleranceForVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TurnOffFaultToleranceForVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func TurnOffFaultToleranceForVM_Task(ctx context.Context, r soap.RoundTripper, req *types.TurnOffFaultToleranceForVM_Task) (*types.TurnOffFaultToleranceForVM_TaskResponse, error) {
	var reqBody, resBody TurnOffFaultToleranceForVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnassignUserFromGroupBody struct {
	Req    *types.UnassignUserFromGroup         `xml:"urn:vim25 UnassignUserFromGroup,omitempty"`
	Res    *types.UnassignUserFromGroupResponse `xml:"UnassignUserFromGroupResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnassignUserFromGroupBody) Fault() *soap.Fault { return b.Fault_ }

func UnassignUserFromGroup(ctx context.Context, r soap.RoundTripper, req *types.UnassignUserFromGroup) (*types.UnassignUserFromGroupResponse, error) {
	var reqBody, resBody UnassignUserFromGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnbindVnicBody struct {
	Req    *types.UnbindVnic         `xml:"urn:vim25 UnbindVnic,omitempty"`
	Res    *types.UnbindVnicResponse `xml:"UnbindVnicResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnbindVnicBody) Fault() *soap.Fault { return b.Fault_ }

func UnbindVnic(ctx context.Context, r soap.RoundTripper, req *types.UnbindVnic) (*types.UnbindVnicResponse, error) {
	var reqBody, resBody UnbindVnicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UninstallHostPatch_TaskBody struct {
	Req    *types.UninstallHostPatch_Task         `xml:"urn:vim25 UninstallHostPatch_Task,omitempty"`
	Res    *types.UninstallHostPatch_TaskResponse `xml:"UninstallHostPatch_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallHostPatch_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UninstallHostPatch_Task(ctx context.Context, r soap.RoundTripper, req *types.UninstallHostPatch_Task) (*types.UninstallHostPatch_TaskResponse, error) {
	var reqBody, resBody UninstallHostPatch_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UninstallIoFilter_TaskBody struct {
	Req    *types.UninstallIoFilter_Task         `xml:"urn:vim25 UninstallIoFilter_Task,omitempty"`
	Res    *types.UninstallIoFilter_TaskResponse `xml:"UninstallIoFilter_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallIoFilter_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UninstallIoFilter_Task(ctx context.Context, r soap.RoundTripper, req *types.UninstallIoFilter_Task) (*types.UninstallIoFilter_TaskResponse, error) {
	var reqBody, resBody UninstallIoFilter_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UninstallServiceBody struct {
	Req    *types.UninstallService         `xml:"urn:vim25 UninstallService,omitempty"`
	Res    *types.UninstallServiceResponse `xml:"UninstallServiceResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallServiceBody) Fault() *soap.Fault { return b.Fault_ }

func UninstallService(ctx context.Context, r soap.RoundTripper, req *types.UninstallService) (*types.UninstallServiceResponse, error) {
	var reqBody, resBody UninstallServiceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmapVmfsVolumeEx_TaskBody struct {
	Req    *types.UnmapVmfsVolumeEx_Task         `xml:"urn:vim25 UnmapVmfsVolumeEx_Task,omitempty"`
	Res    *types.UnmapVmfsVolumeEx_TaskResponse `xml:"UnmapVmfsVolumeEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmapVmfsVolumeEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnmapVmfsVolumeEx_Task(ctx context.Context, r soap.RoundTripper, req *types.UnmapVmfsVolumeEx_Task) (*types.UnmapVmfsVolumeEx_TaskResponse, error) {
	var reqBody, resBody UnmapVmfsVolumeEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmarkServiceProviderEntitiesBody struct {
	Req    *types.UnmarkServiceProviderEntities         `xml:"urn:vim25 UnmarkServiceProviderEntities,omitempty"`
	Res    *types.UnmarkServiceProviderEntitiesResponse `xml:"UnmarkServiceProviderEntitiesResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmarkServiceProviderEntitiesBody) Fault() *soap.Fault { return b.Fault_ }

func UnmarkServiceProviderEntities(ctx context.Context, r soap.RoundTripper, req *types.UnmarkServiceProviderEntities) (*types.UnmarkServiceProviderEntitiesResponse, error) {
	var reqBody, resBody UnmarkServiceProviderEntitiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountDiskMapping_TaskBody struct {
	Req    *types.UnmountDiskMapping_Task         `xml:"urn:vim25 UnmountDiskMapping_Task,omitempty"`
	Res    *types.UnmountDiskMapping_TaskResponse `xml:"UnmountDiskMapping_TaskResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountDiskMapping_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountDiskMapping_Task(ctx context.Context, r soap.RoundTripper, req *types.UnmountDiskMapping_Task) (*types.UnmountDiskMapping_TaskResponse, error) {
	var reqBody, resBody UnmountDiskMapping_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountForceMountedVmfsVolumeBody struct {
	Req    *types.UnmountForceMountedVmfsVolume         `xml:"urn:vim25 UnmountForceMountedVmfsVolume,omitempty"`
	Res    *types.UnmountForceMountedVmfsVolumeResponse `xml:"UnmountForceMountedVmfsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountForceMountedVmfsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountForceMountedVmfsVolume(ctx context.Context, r soap.RoundTripper, req *types.UnmountForceMountedVmfsVolume) (*types.UnmountForceMountedVmfsVolumeResponse, error) {
	var reqBody, resBody UnmountForceMountedVmfsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountToolsInstallerBody struct {
	Req    *types.UnmountToolsInstaller         `xml:"urn:vim25 UnmountToolsInstaller,omitempty"`
	Res    *types.UnmountToolsInstallerResponse `xml:"UnmountToolsInstallerResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountToolsInstallerBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountToolsInstaller(ctx context.Context, r soap.RoundTripper, req *types.UnmountToolsInstaller) (*types.UnmountToolsInstallerResponse, error) {
	var reqBody, resBody UnmountToolsInstallerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountVffsVolumeBody struct {
	Req    *types.UnmountVffsVolume         `xml:"urn:vim25 UnmountVffsVolume,omitempty"`
	Res    *types.UnmountVffsVolumeResponse `xml:"UnmountVffsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountVffsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountVffsVolume(ctx context.Context, r soap.RoundTripper, req *types.UnmountVffsVolume) (*types.UnmountVffsVolumeResponse, error) {
	var reqBody, resBody UnmountVffsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountVmfsVolumeBody struct {
	Req    *types.UnmountVmfsVolume         `xml:"urn:vim25 UnmountVmfsVolume,omitempty"`
	Res    *types.UnmountVmfsVolumeResponse `xml:"UnmountVmfsVolumeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountVmfsVolumeBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountVmfsVolume(ctx context.Context, r soap.RoundTripper, req *types.UnmountVmfsVolume) (*types.UnmountVmfsVolumeResponse, error) {
	var reqBody, resBody UnmountVmfsVolumeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnmountVmfsVolumeEx_TaskBody struct {
	Req    *types.UnmountVmfsVolumeEx_Task         `xml:"urn:vim25 UnmountVmfsVolumeEx_Task,omitempty"`
	Res    *types.UnmountVmfsVolumeEx_TaskResponse `xml:"UnmountVmfsVolumeEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountVmfsVolumeEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnmountVmfsVolumeEx_Task(ctx context.Context, r soap.RoundTripper, req *types.UnmountVmfsVolumeEx_Task) (*types.UnmountVmfsVolumeEx_TaskResponse, error) {
	var reqBody, resBody UnmountVmfsVolumeEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterAndDestroy_TaskBody struct {
	Req    *types.UnregisterAndDestroy_Task         `xml:"urn:vim25 UnregisterAndDestroy_Task,omitempty"`
	Res    *types.UnregisterAndDestroy_TaskResponse `xml:"UnregisterAndDestroy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterAndDestroy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterAndDestroy_Task(ctx context.Context, r soap.RoundTripper, req *types.UnregisterAndDestroy_Task) (*types.UnregisterAndDestroy_TaskResponse, error) {
	var reqBody, resBody UnregisterAndDestroy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterExtensionBody struct {
	Req    *types.UnregisterExtension         `xml:"urn:vim25 UnregisterExtension,omitempty"`
	Res    *types.UnregisterExtensionResponse `xml:"UnregisterExtensionResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterExtensionBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterExtension(ctx context.Context, r soap.RoundTripper, req *types.UnregisterExtension) (*types.UnregisterExtensionResponse, error) {
	var reqBody, resBody UnregisterExtensionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterHealthUpdateProviderBody struct {
	Req    *types.UnregisterHealthUpdateProvider         `xml:"urn:vim25 UnregisterHealthUpdateProvider,omitempty"`
	Res    *types.UnregisterHealthUpdateProviderResponse `xml:"UnregisterHealthUpdateProviderResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterHealthUpdateProviderBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterHealthUpdateProvider(ctx context.Context, r soap.RoundTripper, req *types.UnregisterHealthUpdateProvider) (*types.UnregisterHealthUpdateProviderResponse, error) {
	var reqBody, resBody UnregisterHealthUpdateProviderBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterKmsClusterBody struct {
	Req    *types.UnregisterKmsCluster         `xml:"urn:vim25 UnregisterKmsCluster,omitempty"`
	Res    *types.UnregisterKmsClusterResponse `xml:"UnregisterKmsClusterResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterKmsClusterBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterKmsCluster(ctx context.Context, r soap.RoundTripper, req *types.UnregisterKmsCluster) (*types.UnregisterKmsClusterResponse, error) {
	var reqBody, resBody UnregisterKmsClusterBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterVMBody struct {
	Req    *types.UnregisterVM         `xml:"urn:vim25 UnregisterVM,omitempty"`
	Res    *types.UnregisterVMResponse `xml:"UnregisterVMResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterVMBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterVM(ctx context.Context, r soap.RoundTripper, req *types.UnregisterVM) (*types.UnregisterVMResponse, error) {
	var reqBody, resBody UnregisterVMBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateAnswerFile_TaskBody struct {
	Req    *types.UpdateAnswerFile_Task         `xml:"urn:vim25 UpdateAnswerFile_Task,omitempty"`
	Res    *types.UpdateAnswerFile_TaskResponse `xml:"UpdateAnswerFile_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAnswerFile_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateAnswerFile_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateAnswerFile_Task) (*types.UpdateAnswerFile_TaskResponse, error) {
	var reqBody, resBody UpdateAnswerFile_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateAssignableHardwareConfigBody struct {
	Req    *types.UpdateAssignableHardwareConfig         `xml:"urn:vim25 UpdateAssignableHardwareConfig,omitempty"`
	Res    *types.UpdateAssignableHardwareConfigResponse `xml:"UpdateAssignableHardwareConfigResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAssignableHardwareConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateAssignableHardwareConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateAssignableHardwareConfig) (*types.UpdateAssignableHardwareConfigResponse, error) {
	var reqBody, resBody UpdateAssignableHardwareConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateAssignedLicenseBody struct {
	Req    *types.UpdateAssignedLicense         `xml:"urn:vim25 UpdateAssignedLicense,omitempty"`
	Res    *types.UpdateAssignedLicenseResponse `xml:"UpdateAssignedLicenseResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAssignedLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateAssignedLicense(ctx context.Context, r soap.RoundTripper, req *types.UpdateAssignedLicense) (*types.UpdateAssignedLicenseResponse, error) {
	var reqBody, resBody UpdateAssignedLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateAuthorizationRoleBody struct {
	Req    *types.UpdateAuthorizationRole         `xml:"urn:vim25 UpdateAuthorizationRole,omitempty"`
	Res    *types.UpdateAuthorizationRoleResponse `xml:"UpdateAuthorizationRoleResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAuthorizationRoleBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateAuthorizationRole(ctx context.Context, r soap.RoundTripper, req *types.UpdateAuthorizationRole) (*types.UpdateAuthorizationRoleResponse, error) {
	var reqBody, resBody UpdateAuthorizationRoleBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateBootDeviceBody struct {
	Req    *types.UpdateBootDevice         `xml:"urn:vim25 UpdateBootDevice,omitempty"`
	Res    *types.UpdateBootDeviceResponse `xml:"UpdateBootDeviceResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateBootDeviceBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateBootDevice(ctx context.Context, r soap.RoundTripper, req *types.UpdateBootDevice) (*types.UpdateBootDeviceResponse, error) {
	var reqBody, resBody UpdateBootDeviceBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateChildResourceConfigurationBody struct {
	Req    *types.UpdateChildResourceConfiguration         `xml:"urn:vim25 UpdateChildResourceConfiguration,omitempty"`
	Res    *types.UpdateChildResourceConfigurationResponse `xml:"UpdateChildResourceConfigurationResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateChildResourceConfigurationBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateChildResourceConfiguration(ctx context.Context, r soap.RoundTripper, req *types.UpdateChildResourceConfiguration) (*types.UpdateChildResourceConfigurationResponse, error) {
	var reqBody, resBody UpdateChildResourceConfigurationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateClusterProfileBody struct {
	Req    *types.UpdateClusterProfile         `xml:"urn:vim25 UpdateClusterProfile,omitempty"`
	Res    *types.UpdateClusterProfileResponse `xml:"UpdateClusterProfileResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateClusterProfileBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateClusterProfile(ctx context.Context, r soap.RoundTripper, req *types.UpdateClusterProfile) (*types.UpdateClusterProfileResponse, error) {
	var reqBody, resBody UpdateClusterProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateConfigBody struct {
	Req    *types.UpdateConfig         `xml:"urn:vim25 UpdateConfig,omitempty"`
	Res    *types.UpdateConfigResponse `xml:"UpdateConfigResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateConfig) (*types.UpdateConfigResponse, error) {
	var reqBody, resBody UpdateConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateConsoleIpRouteConfigBody struct {
	Req    *types.UpdateConsoleIpRouteConfig         `xml:"urn:vim25 UpdateConsoleIpRouteConfig,omitempty"`
	Res    *types.UpdateConsoleIpRouteConfigResponse `xml:"UpdateConsoleIpRouteConfigResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateConsoleIpRouteConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateConsoleIpRouteConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateConsoleIpRouteConfig) (*types.UpdateConsoleIpRouteConfigResponse, error) {
	var reqBody, resBody UpdateConsoleIpRouteConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateCounterLevelMappingBody struct {
	Req    *types.UpdateCounterLevelMapping         `xml:"urn:vim25 UpdateCounterLevelMapping,omitempty"`
	Res    *types.UpdateCounterLevelMappingResponse `xml:"UpdateCounterLevelMappingResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateCounterLevelMappingBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateCounterLevelMapping(ctx context.Context, r soap.RoundTripper, req *types.UpdateCounterLevelMapping) (*types.UpdateCounterLevelMappingResponse, error) {
	var reqBody, resBody UpdateCounterLevelMappingBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDVSHealthCheckConfig_TaskBody struct {
	Req    *types.UpdateDVSHealthCheckConfig_Task         `xml:"urn:vim25 UpdateDVSHealthCheckConfig_Task,omitempty"`
	Res    *types.UpdateDVSHealthCheckConfig_TaskResponse `xml:"UpdateDVSHealthCheckConfig_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDVSHealthCheckConfig_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDVSHealthCheckConfig_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateDVSHealthCheckConfig_Task) (*types.UpdateDVSHealthCheckConfig_TaskResponse, error) {
	var reqBody, resBody UpdateDVSHealthCheckConfig_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDVSLacpGroupConfig_TaskBody struct {
	Req    *types.UpdateDVSLacpGroupConfig_Task         `xml:"urn:vim25 UpdateDVSLacpGroupConfig_Task,omitempty"`
	Res    *types.UpdateDVSLacpGroupConfig_TaskResponse `xml:"UpdateDVSLacpGroupConfig_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDVSLacpGroupConfig_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDVSLacpGroupConfig_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateDVSLacpGroupConfig_Task) (*types.UpdateDVSLacpGroupConfig_TaskResponse, error) {
	var reqBody, resBody UpdateDVSLacpGroupConfig_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDateTimeBody struct {
	Req    *types.UpdateDateTime         `xml:"urn:vim25 UpdateDateTime,omitempty"`
	Res    *types.UpdateDateTimeResponse `xml:"UpdateDateTimeResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDateTimeBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDateTime(ctx context.Context, r soap.RoundTripper, req *types.UpdateDateTime) (*types.UpdateDateTimeResponse, error) {
	var reqBody, resBody UpdateDateTimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDateTimeConfigBody struct {
	Req    *types.UpdateDateTimeConfig         `xml:"urn:vim25 UpdateDateTimeConfig,omitempty"`
	Res    *types.UpdateDateTimeConfigResponse `xml:"UpdateDateTimeConfigResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDateTimeConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDateTimeConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateDateTimeConfig) (*types.UpdateDateTimeConfigResponse, error) {
	var reqBody, resBody UpdateDateTimeConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDefaultPolicyBody struct {
	Req    *types.UpdateDefaultPolicy         `xml:"urn:vim25 UpdateDefaultPolicy,omitempty"`
	Res    *types.UpdateDefaultPolicyResponse `xml:"UpdateDefaultPolicyResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDefaultPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDefaultPolicy(ctx context.Context, r soap.RoundTripper, req *types.UpdateDefaultPolicy) (*types.UpdateDefaultPolicyResponse, error) {
	var reqBody, resBody UpdateDefaultPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDiskPartitionsBody struct {
	Req    *types.UpdateDiskPartitions         `xml:"urn:vim25 UpdateDiskPartitions,omitempty"`
	Res    *types.UpdateDiskPartitionsResponse `xml:"UpdateDiskPartitionsResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDiskPartitionsBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDiskPartitions(ctx context.Context, r soap.RoundTripper, req *types.UpdateDiskPartitions) (*types.UpdateDiskPartitionsResponse, error) {
	var reqBody, resBody UpdateDiskPartitionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDnsConfigBody struct {
	Req    *types.UpdateDnsConfig         `xml:"urn:vim25 UpdateDnsConfig,omitempty"`
	Res    *types.UpdateDnsConfigResponse `xml:"UpdateDnsConfigResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDnsConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDnsConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateDnsConfig) (*types.UpdateDnsConfigResponse, error) {
	var reqBody, resBody UpdateDnsConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateDvsCapabilityBody struct {
	Req    *types.UpdateDvsCapability         `xml:"urn:vim25 UpdateDvsCapability,omitempty"`
	Res    *types.UpdateDvsCapabilityResponse `xml:"UpdateDvsCapabilityResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDvsCapabilityBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateDvsCapability(ctx context.Context, r soap.RoundTripper, req *types.UpdateDvsCapability) (*types.UpdateDvsCapabilityResponse, error) {
	var reqBody, resBody UpdateDvsCapabilityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateExtensionBody struct {
	Req    *types.UpdateExtension         `xml:"urn:vim25 UpdateExtension,omitempty"`
	Res    *types.UpdateExtensionResponse `xml:"UpdateExtensionResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateExtensionBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateExtension(ctx context.Context, r soap.RoundTripper, req *types.UpdateExtension) (*types.UpdateExtensionResponse, error) {
	var reqBody, resBody UpdateExtensionBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateFlagsBody struct {
	Req    *types.UpdateFlags         `xml:"urn:vim25 UpdateFlags,omitempty"`
	Res    *types.UpdateFlagsResponse `xml:"UpdateFlagsResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateFlagsBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateFlags(ctx context.Context, r soap.RoundTripper, req *types.UpdateFlags) (*types.UpdateFlagsResponse, error) {
	var reqBody, resBody UpdateFlagsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateGraphicsConfigBody struct {
	Req    *types.UpdateGraphicsConfig         `xml:"urn:vim25 UpdateGraphicsConfig,omitempty"`
	Res    *types.UpdateGraphicsConfigResponse `xml:"UpdateGraphicsConfigResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateGraphicsConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateGraphicsConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateGraphicsConfig) (*types.UpdateGraphicsConfigResponse, error) {
	var reqBody, resBody UpdateGraphicsConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateHostImageAcceptanceLevelBody struct {
	Req    *types.UpdateHostImageAcceptanceLevel         `xml:"urn:vim25 UpdateHostImageAcceptanceLevel,omitempty"`
	Res    *types.UpdateHostImageAcceptanceLevelResponse `xml:"UpdateHostImageAcceptanceLevelResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostImageAcceptanceLevelBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateHostImageAcceptanceLevel(ctx context.Context, r soap.RoundTripper, req *types.UpdateHostImageAcceptanceLevel) (*types.UpdateHostImageAcceptanceLevelResponse, error) {
	var reqBody, resBody UpdateHostImageAcceptanceLevelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateHostProfileBody struct {
	Req    *types.UpdateHostProfile         `xml:"urn:vim25 UpdateHostProfile,omitempty"`
	Res    *types.UpdateHostProfileResponse `xml:"UpdateHostProfileResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostProfileBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateHostProfile(ctx context.Context, r soap.RoundTripper, req *types.UpdateHostProfile) (*types.UpdateHostProfileResponse, error) {
	var reqBody, resBody UpdateHostProfileBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateHostSpecificationBody struct {
	Req    *types.UpdateHostSpecification         `xml:"urn:vim25 UpdateHostSpecification,omitempty"`
	Res    *types.UpdateHostSpecificationResponse `xml:"UpdateHostSpecificationResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostSpecificationBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateHostSpecification(ctx context.Context, r soap.RoundTripper, req *types.UpdateHostSpecification) (*types.UpdateHostSpecificationResponse, error) {
	var reqBody, resBody UpdateHostSpecificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateHostSubSpecificationBody struct {
	Req    *types.UpdateHostSubSpecification         `xml:"urn:vim25 UpdateHostSubSpecification,omitempty"`
	Res    *types.UpdateHostSubSpecificationResponse `xml:"UpdateHostSubSpecificationResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostSubSpecificationBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateHostSubSpecification(ctx context.Context, r soap.RoundTripper, req *types.UpdateHostSubSpecification) (*types.UpdateHostSubSpecificationResponse, error) {
	var reqBody, resBody UpdateHostSubSpecificationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateHppMultipathLunPolicyBody struct {
	Req    *types.UpdateHppMultipathLunPolicy         `xml:"urn:vim25 UpdateHppMultipathLunPolicy,omitempty"`
	Res    *types.UpdateHppMultipathLunPolicyResponse `xml:"UpdateHppMultipathLunPolicyResponse,omitempty"`
	Fault_ *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHppMultipathLunPolicyBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateHppMultipathLunPolicy(ctx context.Context, r soap.RoundTripper, req *types.UpdateHppMultipathLunPolicy) (*types.UpdateHppMultipathLunPolicyResponse, error) {
	var reqBody, resBody UpdateHppMultipathLunPolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiAdvancedOptionsBody struct {
	Req    *types.UpdateInternetScsiAdvancedOptions         `xml:"urn:vim25 UpdateInternetScsiAdvancedOptions,omitempty"`
	Res    *types.UpdateInternetScsiAdvancedOptionsResponse `xml:"UpdateInternetScsiAdvancedOptionsResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAdvancedOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiAdvancedOptions(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiAdvancedOptions) (*types.UpdateInternetScsiAdvancedOptionsResponse, error) {
	var reqBody, resBody UpdateInternetScsiAdvancedOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiAliasBody struct {
	Req    *types.UpdateInternetScsiAlias         `xml:"urn:vim25 UpdateInternetScsiAlias,omitempty"`
	Res    *types.UpdateInternetScsiAliasResponse `xml:"UpdateInternetScsiAliasResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAliasBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiAlias(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiAlias) (*types.UpdateInternetScsiAliasResponse, error) {
	var reqBody, resBody UpdateInternetScsiAliasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiAuthenticationPropertiesBody struct {
	Req    *types.UpdateInternetScsiAuthenticationProperties         `xml:"urn:vim25 UpdateInternetScsiAuthenticationProperties,omitempty"`
	Res    *types.UpdateInternetScsiAuthenticationPropertiesResponse `xml:"UpdateInternetScsiAuthenticationPropertiesResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAuthenticationPropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiAuthenticationProperties(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiAuthenticationProperties) (*types.UpdateInternetScsiAuthenticationPropertiesResponse, error) {
	var reqBody, resBody UpdateInternetScsiAuthenticationPropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiDigestPropertiesBody struct {
	Req    *types.UpdateInternetScsiDigestProperties         `xml:"urn:vim25 UpdateInternetScsiDigestProperties,omitempty"`
	Res    *types.UpdateInternetScsiDigestPropertiesResponse `xml:"UpdateInternetScsiDigestPropertiesResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiDigestPropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiDigestProperties(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiDigestProperties) (*types.UpdateInternetScsiDigestPropertiesResponse, error) {
	var reqBody, resBody UpdateInternetScsiDigestPropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiDiscoveryPropertiesBody struct {
	Req    *types.UpdateInternetScsiDiscoveryProperties         `xml:"urn:vim25 UpdateInternetScsiDiscoveryProperties,omitempty"`
	Res    *types.UpdateInternetScsiDiscoveryPropertiesResponse `xml:"UpdateInternetScsiDiscoveryPropertiesResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiDiscoveryPropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiDiscoveryProperties(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiDiscoveryProperties) (*types.UpdateInternetScsiDiscoveryPropertiesResponse, error) {
	var reqBody, resBody UpdateInternetScsiDiscoveryPropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiIPPropertiesBody struct {
	Req    *types.UpdateInternetScsiIPProperties         `xml:"urn:vim25 UpdateInternetScsiIPProperties,omitempty"`
	Res    *types.UpdateInternetScsiIPPropertiesResponse `xml:"UpdateInternetScsiIPPropertiesResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiIPPropertiesBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiIPProperties(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiIPProperties) (*types.UpdateInternetScsiIPPropertiesResponse, error) {
	var reqBody, resBody UpdateInternetScsiIPPropertiesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateInternetScsiNameBody struct {
	Req    *types.UpdateInternetScsiName         `xml:"urn:vim25 UpdateInternetScsiName,omitempty"`
	Res    *types.UpdateInternetScsiNameResponse `xml:"UpdateInternetScsiNameResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiNameBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateInternetScsiName(ctx context.Context, r soap.RoundTripper, req *types.UpdateInternetScsiName) (*types.UpdateInternetScsiNameResponse, error) {
	var reqBody, resBody UpdateInternetScsiNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateIpConfigBody struct {
	Req    *types.UpdateIpConfig         `xml:"urn:vim25 UpdateIpConfig,omitempty"`
	Res    *types.UpdateIpConfigResponse `xml:"UpdateIpConfigResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateIpConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateIpConfig) (*types.UpdateIpConfigResponse, error) {
	var reqBody, resBody UpdateIpConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateIpPoolBody struct {
	Req    *types.UpdateIpPool         `xml:"urn:vim25 UpdateIpPool,omitempty"`
	Res    *types.UpdateIpPoolResponse `xml:"UpdateIpPoolResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpPoolBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateIpPool(ctx context.Context, r soap.RoundTripper, req *types.UpdateIpPool) (*types.UpdateIpPoolResponse, error) {
	var reqBody, resBody UpdateIpPoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateIpRouteConfigBody struct {
	Req    *types.UpdateIpRouteConfig         `xml:"urn:vim25 UpdateIpRouteConfig,omitempty"`
	Res    *types.UpdateIpRouteConfigResponse `xml:"UpdateIpRouteConfigResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpRouteConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateIpRouteConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateIpRouteConfig) (*types.UpdateIpRouteConfigResponse, error) {
	var reqBody, resBody UpdateIpRouteConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateIpRouteTableConfigBody struct {
	Req    *types.UpdateIpRouteTableConfig         `xml:"urn:vim25 UpdateIpRouteTableConfig,omitempty"`
	Res    *types.UpdateIpRouteTableConfigResponse `xml:"UpdateIpRouteTableConfigResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpRouteTableConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateIpRouteTableConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateIpRouteTableConfig) (*types.UpdateIpRouteTableConfigResponse, error) {
	var reqBody, resBody UpdateIpRouteTableConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateIpmiBody struct {
	Req    *types.UpdateIpmi         `xml:"urn:vim25 UpdateIpmi,omitempty"`
	Res    *types.UpdateIpmiResponse `xml:"UpdateIpmiResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpmiBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateIpmi(ctx context.Context, r soap.RoundTripper, req *types.UpdateIpmi) (*types.UpdateIpmiResponse, error) {
	var reqBody, resBody UpdateIpmiBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateKmipServerBody struct {
	Req    *types.UpdateKmipServer         `xml:"urn:vim25 UpdateKmipServer,omitempty"`
	Res    *types.UpdateKmipServerResponse `xml:"UpdateKmipServerResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateKmipServerBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateKmipServer(ctx context.Context, r soap.RoundTripper, req *types.UpdateKmipServer) (*types.UpdateKmipServerResponse, error) {
	var reqBody, resBody UpdateKmipServerBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateKmsSignedCsrClientCertBody struct {
	Req    *types.UpdateKmsSignedCsrClientCert         `xml:"urn:vim25 UpdateKmsSignedCsrClientCert,omitempty"`
	Res    *types.UpdateKmsSignedCsrClientCertResponse `xml:"UpdateKmsSignedCsrClientCertResponse,omitempty"`
	Fault_ *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateKmsSignedCsrClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateKmsSignedCsrClientCert(ctx context.Context, r soap.RoundTripper, req *types.UpdateKmsSignedCsrClientCert) (*types.UpdateKmsSignedCsrClientCertResponse, error) {
	var reqBody, resBody UpdateKmsSignedCsrClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateLicenseBody struct {
	Req    *types.UpdateLicense         `xml:"urn:vim25 UpdateLicense,omitempty"`
	Res    *types.UpdateLicenseResponse `xml:"UpdateLicenseResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLicenseBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateLicense(ctx context.Context, r soap.RoundTripper, req *types.UpdateLicense) (*types.UpdateLicenseResponse, error) {
	var reqBody, resBody UpdateLicenseBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateLicenseLabelBody struct {
	Req    *types.UpdateLicenseLabel         `xml:"urn:vim25 UpdateLicenseLabel,omitempty"`
	Res    *types.UpdateLicenseLabelResponse `xml:"UpdateLicenseLabelResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLicenseLabelBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateLicenseLabel(ctx context.Context, r soap.RoundTripper, req *types.UpdateLicenseLabel) (*types.UpdateLicenseLabelResponse, error) {
	var reqBody, resBody UpdateLicenseLabelBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateLinkedChildrenBody struct {
	Req    *types.UpdateLinkedChildren         `xml:"urn:vim25 UpdateLinkedChildren,omitempty"`
	Res    *types.UpdateLinkedChildrenResponse `xml:"UpdateLinkedChildrenResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLinkedChildrenBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateLinkedChildren(ctx context.Context, r soap.RoundTripper, req *types.UpdateLinkedChildren) (*types.UpdateLinkedChildrenResponse, error) {
	var reqBody, resBody UpdateLinkedChildrenBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateLocalSwapDatastoreBody struct {
	Req    *types.UpdateLocalSwapDatastore         `xml:"urn:vim25 UpdateLocalSwapDatastore,omitempty"`
	Res    *types.UpdateLocalSwapDatastoreResponse `xml:"UpdateLocalSwapDatastoreResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLocalSwapDatastoreBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateLocalSwapDatastore(ctx context.Context, r soap.RoundTripper, req *types.UpdateLocalSwapDatastore) (*types.UpdateLocalSwapDatastoreResponse, error) {
	var reqBody, resBody UpdateLocalSwapDatastoreBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateLockdownExceptionsBody struct {
	Req    *types.UpdateLockdownExceptions         `xml:"urn:vim25 UpdateLockdownExceptions,omitempty"`
	Res    *types.UpdateLockdownExceptionsResponse `xml:"UpdateLockdownExceptionsResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLockdownExceptionsBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateLockdownExceptions(ctx context.Context, r soap.RoundTripper, req *types.UpdateLockdownExceptions) (*types.UpdateLockdownExceptionsResponse, error) {
	var reqBody, resBody UpdateLockdownExceptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateModuleOptionStringBody struct {
	Req    *types.UpdateModuleOptionString         `xml:"urn:vim25 UpdateModuleOptionString,omitempty"`
	Res    *types.UpdateModuleOptionStringResponse `xml:"UpdateModuleOptionStringResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateModuleOptionStringBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateModuleOptionString(ctx context.Context, r soap.RoundTripper, req *types.UpdateModuleOptionString) (*types.UpdateModuleOptionStringResponse, error) {
	var reqBody, resBody UpdateModuleOptionStringBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateNetworkConfigBody struct {
	Req    *types.UpdateNetworkConfig         `xml:"urn:vim25 UpdateNetworkConfig,omitempty"`
	Res    *types.UpdateNetworkConfigResponse `xml:"UpdateNetworkConfigResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateNetworkConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateNetworkConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateNetworkConfig) (*types.UpdateNetworkConfigResponse, error) {
	var reqBody, resBody UpdateNetworkConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateNetworkResourcePoolBody struct {
	Req    *types.UpdateNetworkResourcePool         `xml:"urn:vim25 UpdateNetworkResourcePool,omitempty"`
	Res    *types.UpdateNetworkResourcePoolResponse `xml:"UpdateNetworkResourcePoolResponse,omitempty"`
	Fault_ *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateNetworkResourcePoolBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateNetworkResourcePool(ctx context.Context, r soap.RoundTripper, req *types.UpdateNetworkResourcePool) (*types.UpdateNetworkResourcePoolResponse, error) {
	var reqBody, resBody UpdateNetworkResourcePoolBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateOptionsBody struct {
	Req    *types.UpdateOptions         `xml:"urn:vim25 UpdateOptions,omitempty"`
	Res    *types.UpdateOptionsResponse `xml:"UpdateOptionsResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateOptionsBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateOptions(ctx context.Context, r soap.RoundTripper, req *types.UpdateOptions) (*types.UpdateOptionsResponse, error) {
	var reqBody, resBody UpdateOptionsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdatePassthruConfigBody struct {
	Req    *types.UpdatePassthruConfig         `xml:"urn:vim25 UpdatePassthruConfig,omitempty"`
	Res    *types.UpdatePassthruConfigResponse `xml:"UpdatePassthruConfigResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePassthruConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdatePassthruConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdatePassthruConfig) (*types.UpdatePassthruConfigResponse, error) {
	var reqBody, resBody UpdatePassthruConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdatePerfIntervalBody struct {
	Req    *types.UpdatePerfInterval         `xml:"urn:vim25 UpdatePerfInterval,omitempty"`
	Res    *types.UpdatePerfIntervalResponse `xml:"UpdatePerfIntervalResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePerfIntervalBody) Fault() *soap.Fault { return b.Fault_ }

func UpdatePerfInterval(ctx context.Context, r soap.RoundTripper, req *types.UpdatePerfInterval) (*types.UpdatePerfIntervalResponse, error) {
	var reqBody, resBody UpdatePerfIntervalBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdatePhysicalNicLinkSpeedBody struct {
	Req    *types.UpdatePhysicalNicLinkSpeed         `xml:"urn:vim25 UpdatePhysicalNicLinkSpeed,omitempty"`
	Res    *types.UpdatePhysicalNicLinkSpeedResponse `xml:"UpdatePhysicalNicLinkSpeedResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePhysicalNicLinkSpeedBody) Fault() *soap.Fault { return b.Fault_ }

func UpdatePhysicalNicLinkSpeed(ctx context.Context, r soap.RoundTripper, req *types.UpdatePhysicalNicLinkSpeed) (*types.UpdatePhysicalNicLinkSpeedResponse, error) {
	var reqBody, resBody UpdatePhysicalNicLinkSpeedBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdatePortGroupBody struct {
	Req    *types.UpdatePortGroup         `xml:"urn:vim25 UpdatePortGroup,omitempty"`
	Res    *types.UpdatePortGroupResponse `xml:"UpdatePortGroupResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePortGroupBody) Fault() *soap.Fault { return b.Fault_ }

func UpdatePortGroup(ctx context.Context, r soap.RoundTripper, req *types.UpdatePortGroup) (*types.UpdatePortGroupResponse, error) {
	var reqBody, resBody UpdatePortGroupBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateProductLockerLocation_TaskBody struct {
	Req    *types.UpdateProductLockerLocation_Task         `xml:"urn:vim25 UpdateProductLockerLocation_Task,omitempty"`
	Res    *types.UpdateProductLockerLocation_TaskResponse `xml:"UpdateProductLockerLocation_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateProductLockerLocation_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateProductLockerLocation_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateProductLockerLocation_Task) (*types.UpdateProductLockerLocation_TaskResponse, error) {
	var reqBody, resBody UpdateProductLockerLocation_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateProgressBody struct {
	Req    *types.UpdateProgress         `xml:"urn:vim25 UpdateProgress,omitempty"`
	Res    *types.UpdateProgressResponse `xml:"UpdateProgressResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateProgressBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateProgress(ctx context.Context, r soap.RoundTripper, req *types.UpdateProgress) (*types.UpdateProgressResponse, error) {
	var reqBody, resBody UpdateProgressBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateReferenceHostBody struct {
	Req    *types.UpdateReferenceHost         `xml:"urn:vim25 UpdateReferenceHost,omitempty"`
	Res    *types.UpdateReferenceHostResponse `xml:"UpdateReferenceHostResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateReferenceHostBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateReferenceHost(ctx context.Context, r soap.RoundTripper, req *types.UpdateReferenceHost) (*types.UpdateReferenceHostResponse, error) {
	var reqBody, resBody UpdateReferenceHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateRulesetBody struct {
	Req    *types.UpdateRuleset         `xml:"urn:vim25 UpdateRuleset,omitempty"`
	Res    *types.UpdateRulesetResponse `xml:"UpdateRulesetResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateRulesetBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateRuleset(ctx context.Context, r soap.RoundTripper, req *types.UpdateRuleset) (*types.UpdateRulesetResponse, error) {
	var reqBody, resBody UpdateRulesetBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateScsiLunDisplayNameBody struct {
	Req    *types.UpdateScsiLunDisplayName         `xml:"urn:vim25 UpdateScsiLunDisplayName,omitempty"`
	Res    *types.UpdateScsiLunDisplayNameResponse `xml:"UpdateScsiLunDisplayNameResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateScsiLunDisplayNameBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateScsiLunDisplayName(ctx context.Context, r soap.RoundTripper, req *types.UpdateScsiLunDisplayName) (*types.UpdateScsiLunDisplayNameResponse, error) {
	var reqBody, resBody UpdateScsiLunDisplayNameBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateSelfSignedClientCertBody struct {
	Req    *types.UpdateSelfSignedClientCert         `xml:"urn:vim25 UpdateSelfSignedClientCert,omitempty"`
	Res    *types.UpdateSelfSignedClientCertResponse `xml:"UpdateSelfSignedClientCertResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSelfSignedClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateSelfSignedClientCert(ctx context.Context, r soap.RoundTripper, req *types.UpdateSelfSignedClientCert) (*types.UpdateSelfSignedClientCertResponse, error) {
	var reqBody, resBody UpdateSelfSignedClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateServiceConsoleVirtualNicBody struct {
	Req    *types.UpdateServiceConsoleVirtualNic         `xml:"urn:vim25 UpdateServiceConsoleVirtualNic,omitempty"`
	Res    *types.UpdateServiceConsoleVirtualNicResponse `xml:"UpdateServiceConsoleVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServiceConsoleVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateServiceConsoleVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.UpdateServiceConsoleVirtualNic) (*types.UpdateServiceConsoleVirtualNicResponse, error) {
	var reqBody, resBody UpdateServiceConsoleVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateServiceMessageBody struct {
	Req    *types.UpdateServiceMessage         `xml:"urn:vim25 UpdateServiceMessage,omitempty"`
	Res    *types.UpdateServiceMessageResponse `xml:"UpdateServiceMessageResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServiceMessageBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateServiceMessage(ctx context.Context, r soap.RoundTripper, req *types.UpdateServiceMessage) (*types.UpdateServiceMessageResponse, error) {
	var reqBody, resBody UpdateServiceMessageBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateServicePolicyBody struct {
	Req    *types.UpdateServicePolicy         `xml:"urn:vim25 UpdateServicePolicy,omitempty"`
	Res    *types.UpdateServicePolicyResponse `xml:"UpdateServicePolicyResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServicePolicyBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateServicePolicy(ctx context.Context, r soap.RoundTripper, req *types.UpdateServicePolicy) (*types.UpdateServicePolicyResponse, error) {
	var reqBody, resBody UpdateServicePolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateSoftwareInternetScsiEnabledBody struct {
	Req    *types.UpdateSoftwareInternetScsiEnabled         `xml:"urn:vim25 UpdateSoftwareInternetScsiEnabled,omitempty"`
	Res    *types.UpdateSoftwareInternetScsiEnabledResponse `xml:"UpdateSoftwareInternetScsiEnabledResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSoftwareInternetScsiEnabledBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateSoftwareInternetScsiEnabled(ctx context.Context, r soap.RoundTripper, req *types.UpdateSoftwareInternetScsiEnabled) (*types.UpdateSoftwareInternetScsiEnabledResponse, error) {
	var reqBody, resBody UpdateSoftwareInternetScsiEnabledBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateSystemResourcesBody struct {
	Req    *types.UpdateSystemResources         `xml:"urn:vim25 UpdateSystemResources,omitempty"`
	Res    *types.UpdateSystemResourcesResponse `xml:"UpdateSystemResourcesResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSystemResourcesBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateSystemResources(ctx context.Context, r soap.RoundTripper, req *types.UpdateSystemResources) (*types.UpdateSystemResourcesResponse, error) {
	var reqBody, resBody UpdateSystemResourcesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateSystemSwapConfigurationBody struct {
	Req    *types.UpdateSystemSwapConfiguration         `xml:"urn:vim25 UpdateSystemSwapConfiguration,omitempty"`
	Res    *types.UpdateSystemSwapConfigurationResponse `xml:"UpdateSystemSwapConfigurationResponse,omitempty"`
	Fault_ *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSystemSwapConfigurationBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateSystemSwapConfiguration(ctx context.Context, r soap.RoundTripper, req *types.UpdateSystemSwapConfiguration) (*types.UpdateSystemSwapConfigurationResponse, error) {
	var reqBody, resBody UpdateSystemSwapConfigurationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateSystemUsersBody struct {
	Req    *types.UpdateSystemUsers         `xml:"urn:vim25 UpdateSystemUsers,omitempty"`
	Res    *types.UpdateSystemUsersResponse `xml:"UpdateSystemUsersResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSystemUsersBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateSystemUsers(ctx context.Context, r soap.RoundTripper, req *types.UpdateSystemUsers) (*types.UpdateSystemUsersResponse, error) {
	var reqBody, resBody UpdateSystemUsersBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateUserBody struct {
	Req    *types.UpdateUser         `xml:"urn:vim25 UpdateUser,omitempty"`
	Res    *types.UpdateUserResponse `xml:"UpdateUserResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateUserBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateUser(ctx context.Context, r soap.RoundTripper, req *types.UpdateUser) (*types.UpdateUserResponse, error) {
	var reqBody, resBody UpdateUserBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVAppConfigBody struct {
	Req    *types.UpdateVAppConfig         `xml:"urn:vim25 UpdateVAppConfig,omitempty"`
	Res    *types.UpdateVAppConfigResponse `xml:"UpdateVAppConfigResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVAppConfigBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVAppConfig(ctx context.Context, r soap.RoundTripper, req *types.UpdateVAppConfig) (*types.UpdateVAppConfigResponse, error) {
	var reqBody, resBody UpdateVAppConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVStorageInfrastructureObjectPolicy_TaskBody struct {
	Req    *types.UpdateVStorageInfrastructureObjectPolicy_Task         `xml:"urn:vim25 UpdateVStorageInfrastructureObjectPolicy_Task,omitempty"`
	Res    *types.UpdateVStorageInfrastructureObjectPolicy_TaskResponse `xml:"UpdateVStorageInfrastructureObjectPolicy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVStorageInfrastructureObjectPolicy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVStorageInfrastructureObjectPolicy_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVStorageInfrastructureObjectPolicy_Task) (*types.UpdateVStorageInfrastructureObjectPolicy_TaskResponse, error) {
	var reqBody, resBody UpdateVStorageInfrastructureObjectPolicy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVStorageObjectCrypto_TaskBody struct {
	Req    *types.UpdateVStorageObjectCrypto_Task         `xml:"urn:vim25 UpdateVStorageObjectCrypto_Task,omitempty"`
	Res    *types.UpdateVStorageObjectCrypto_TaskResponse `xml:"UpdateVStorageObjectCrypto_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVStorageObjectCrypto_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVStorageObjectCrypto_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVStorageObjectCrypto_Task) (*types.UpdateVStorageObjectCrypto_TaskResponse, error) {
	var reqBody, resBody UpdateVStorageObjectCrypto_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVStorageObjectPolicy_TaskBody struct {
	Req    *types.UpdateVStorageObjectPolicy_Task         `xml:"urn:vim25 UpdateVStorageObjectPolicy_Task,omitempty"`
	Res    *types.UpdateVStorageObjectPolicy_TaskResponse `xml:"UpdateVStorageObjectPolicy_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVStorageObjectPolicy_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVStorageObjectPolicy_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVStorageObjectPolicy_Task) (*types.UpdateVStorageObjectPolicy_TaskResponse, error) {
	var reqBody, resBody UpdateVStorageObjectPolicy_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVVolVirtualMachineFiles_TaskBody struct {
	Req    *types.UpdateVVolVirtualMachineFiles_Task         `xml:"urn:vim25 UpdateVVolVirtualMachineFiles_Task,omitempty"`
	Res    *types.UpdateVVolVirtualMachineFiles_TaskResponse `xml:"UpdateVVolVirtualMachineFiles_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVVolVirtualMachineFiles_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVVolVirtualMachineFiles_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVVolVirtualMachineFiles_Task) (*types.UpdateVVolVirtualMachineFiles_TaskResponse, error) {
	var reqBody, resBody UpdateVVolVirtualMachineFiles_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVirtualMachineFiles_TaskBody struct {
	Req    *types.UpdateVirtualMachineFiles_Task         `xml:"urn:vim25 UpdateVirtualMachineFiles_Task,omitempty"`
	Res    *types.UpdateVirtualMachineFiles_TaskResponse `xml:"UpdateVirtualMachineFiles_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualMachineFiles_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVirtualMachineFiles_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVirtualMachineFiles_Task) (*types.UpdateVirtualMachineFiles_TaskResponse, error) {
	var reqBody, resBody UpdateVirtualMachineFiles_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVirtualNicBody struct {
	Req    *types.UpdateVirtualNic         `xml:"urn:vim25 UpdateVirtualNic,omitempty"`
	Res    *types.UpdateVirtualNicResponse `xml:"UpdateVirtualNicResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualNicBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVirtualNic(ctx context.Context, r soap.RoundTripper, req *types.UpdateVirtualNic) (*types.UpdateVirtualNicResponse, error) {
	var reqBody, resBody UpdateVirtualNicBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVirtualSwitchBody struct {
	Req    *types.UpdateVirtualSwitch         `xml:"urn:vim25 UpdateVirtualSwitch,omitempty"`
	Res    *types.UpdateVirtualSwitchResponse `xml:"UpdateVirtualSwitchResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualSwitchBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVirtualSwitch(ctx context.Context, r soap.RoundTripper, req *types.UpdateVirtualSwitch) (*types.UpdateVirtualSwitchResponse, error) {
	var reqBody, resBody UpdateVirtualSwitchBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVmfsUnmapBandwidthBody struct {
	Req    *types.UpdateVmfsUnmapBandwidth         `xml:"urn:vim25 UpdateVmfsUnmapBandwidth,omitempty"`
	Res    *types.UpdateVmfsUnmapBandwidthResponse `xml:"UpdateVmfsUnmapBandwidthResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVmfsUnmapBandwidthBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVmfsUnmapBandwidth(ctx context.Context, r soap.RoundTripper, req *types.UpdateVmfsUnmapBandwidth) (*types.UpdateVmfsUnmapBandwidthResponse, error) {
	var reqBody, resBody UpdateVmfsUnmapBandwidthBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVmfsUnmapPriorityBody struct {
	Req    *types.UpdateVmfsUnmapPriority         `xml:"urn:vim25 UpdateVmfsUnmapPriority,omitempty"`
	Res    *types.UpdateVmfsUnmapPriorityResponse `xml:"UpdateVmfsUnmapPriorityResponse,omitempty"`
	Fault_ *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVmfsUnmapPriorityBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVmfsUnmapPriority(ctx context.Context, r soap.RoundTripper, req *types.UpdateVmfsUnmapPriority) (*types.UpdateVmfsUnmapPriorityResponse, error) {
	var reqBody, resBody UpdateVmfsUnmapPriorityBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateVsan_TaskBody struct {
	Req    *types.UpdateVsan_Task         `xml:"urn:vim25 UpdateVsan_Task,omitempty"`
	Res    *types.UpdateVsan_TaskResponse `xml:"UpdateVsan_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVsan_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpdateVsan_Task(ctx context.Context, r soap.RoundTripper, req *types.UpdateVsan_Task) (*types.UpdateVsan_TaskResponse, error) {
	var reqBody, resBody UpdateVsan_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeIoFilter_TaskBody struct {
	Req    *types.UpgradeIoFilter_Task         `xml:"urn:vim25 UpgradeIoFilter_Task,omitempty"`
	Res    *types.UpgradeIoFilter_TaskResponse `xml:"UpgradeIoFilter_TaskResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeIoFilter_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeIoFilter_Task(ctx context.Context, r soap.RoundTripper, req *types.UpgradeIoFilter_Task) (*types.UpgradeIoFilter_TaskResponse, error) {
	var reqBody, resBody UpgradeIoFilter_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeTools_TaskBody struct {
	Req    *types.UpgradeTools_Task         `xml:"urn:vim25 UpgradeTools_Task,omitempty"`
	Res    *types.UpgradeTools_TaskResponse `xml:"UpgradeTools_TaskResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeTools_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeTools_Task(ctx context.Context, r soap.RoundTripper, req *types.UpgradeTools_Task) (*types.UpgradeTools_TaskResponse, error) {
	var reqBody, resBody UpgradeTools_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeVM_TaskBody struct {
	Req    *types.UpgradeVM_Task         `xml:"urn:vim25 UpgradeVM_Task,omitempty"`
	Res    *types.UpgradeVM_TaskResponse `xml:"UpgradeVM_TaskResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVM_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeVM_Task(ctx context.Context, r soap.RoundTripper, req *types.UpgradeVM_Task) (*types.UpgradeVM_TaskResponse, error) {
	var reqBody, resBody UpgradeVM_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeVmLayoutBody struct {
	Req    *types.UpgradeVmLayout         `xml:"urn:vim25 UpgradeVmLayout,omitempty"`
	Res    *types.UpgradeVmLayoutResponse `xml:"UpgradeVmLayoutResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVmLayoutBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeVmLayout(ctx context.Context, r soap.RoundTripper, req *types.UpgradeVmLayout) (*types.UpgradeVmLayoutResponse, error) {
	var reqBody, resBody UpgradeVmLayoutBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeVmfsBody struct {
	Req    *types.UpgradeVmfs         `xml:"urn:vim25 UpgradeVmfs,omitempty"`
	Res    *types.UpgradeVmfsResponse `xml:"UpgradeVmfsResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVmfsBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeVmfs(ctx context.Context, r soap.RoundTripper, req *types.UpgradeVmfs) (*types.UpgradeVmfsResponse, error) {
	var reqBody, resBody UpgradeVmfsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpgradeVsanObjectsBody struct {
	Req    *types.UpgradeVsanObjects         `xml:"urn:vim25 UpgradeVsanObjects,omitempty"`
	Res    *types.UpgradeVsanObjectsResponse `xml:"UpgradeVsanObjectsResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVsanObjectsBody) Fault() *soap.Fault { return b.Fault_ }

func UpgradeVsanObjects(ctx context.Context, r soap.RoundTripper, req *types.UpgradeVsanObjects) (*types.UpgradeVsanObjectsResponse, error) {
	var reqBody, resBody UpgradeVsanObjectsBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UploadClientCertBody struct {
	Req    *types.UploadClientCert         `xml:"urn:vim25 UploadClientCert,omitempty"`
	Res    *types.UploadClientCertResponse `xml:"UploadClientCertResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UploadClientCertBody) Fault() *soap.Fault { return b.Fault_ }

func UploadClientCert(ctx context.Context, r soap.RoundTripper, req *types.UploadClientCert) (*types.UploadClientCertResponse, error) {
	var reqBody, resBody UploadClientCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UploadKmipServerCertBody struct {
	Req    *types.UploadKmipServerCert         `xml:"urn:vim25 UploadKmipServerCert,omitempty"`
	Res    *types.UploadKmipServerCertResponse `xml:"UploadKmipServerCertResponse,omitempty"`
	Fault_ *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UploadKmipServerCertBody) Fault() *soap.Fault { return b.Fault_ }

func UploadKmipServerCert(ctx context.Context, r soap.RoundTripper, req *types.UploadKmipServerCert) (*types.UploadKmipServerCertResponse, error) {
	var reqBody, resBody UploadKmipServerCertBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VCenterUpdateVStorageObjectMetadataEx_TaskBody struct {
	Req    *types.VCenterUpdateVStorageObjectMetadataEx_Task         `xml:"urn:vim25 VCenterUpdateVStorageObjectMetadataEx_Task,omitempty"`
	Res    *types.VCenterUpdateVStorageObjectMetadataEx_TaskResponse `xml:"VCenterUpdateVStorageObjectMetadataEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VCenterUpdateVStorageObjectMetadataEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VCenterUpdateVStorageObjectMetadataEx_Task(ctx context.Context, r soap.RoundTripper, req *types.VCenterUpdateVStorageObjectMetadataEx_Task) (*types.VCenterUpdateVStorageObjectMetadataEx_TaskResponse, error) {
	var reqBody, resBody VCenterUpdateVStorageObjectMetadataEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VStorageObjectCreateSnapshotEx_TaskBody struct {
	Req    *types.VStorageObjectCreateSnapshotEx_Task         `xml:"urn:vim25 VStorageObjectCreateSnapshotEx_Task,omitempty"`
	Res    *types.VStorageObjectCreateSnapshotEx_TaskResponse `xml:"VStorageObjectCreateSnapshotEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VStorageObjectCreateSnapshotEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VStorageObjectCreateSnapshotEx_Task(ctx context.Context, r soap.RoundTripper, req *types.VStorageObjectCreateSnapshotEx_Task) (*types.VStorageObjectCreateSnapshotEx_TaskResponse, error) {
	var reqBody, resBody VStorageObjectCreateSnapshotEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VStorageObjectCreateSnapshot_TaskBody struct {
	Req    *types.VStorageObjectCreateSnapshot_Task         `xml:"urn:vim25 VStorageObjectCreateSnapshot_Task,omitempty"`
	Res    *types.VStorageObjectCreateSnapshot_TaskResponse `xml:"VStorageObjectCreateSnapshot_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VStorageObjectCreateSnapshot_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VStorageObjectCreateSnapshot_Task(ctx context.Context, r soap.RoundTripper, req *types.VStorageObjectCreateSnapshot_Task) (*types.VStorageObjectCreateSnapshot_TaskResponse, error) {
	var reqBody, resBody VStorageObjectCreateSnapshot_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VStorageObjectDeleteSnapshotEx_TaskBody struct {
	Req    *types.VStorageObjectDeleteSnapshotEx_Task         `xml:"urn:vim25 VStorageObjectDeleteSnapshotEx_Task,omitempty"`
	Res    *types.VStorageObjectDeleteSnapshotEx_TaskResponse `xml:"VStorageObjectDeleteSnapshotEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VStorageObjectDeleteSnapshotEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VStorageObjectDeleteSnapshotEx_Task(ctx context.Context, r soap.RoundTripper, req *types.VStorageObjectDeleteSnapshotEx_Task) (*types.VStorageObjectDeleteSnapshotEx_TaskResponse, error) {
	var reqBody, resBody VStorageObjectDeleteSnapshotEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VStorageObjectExtendDiskEx_TaskBody struct {
	Req    *types.VStorageObjectExtendDiskEx_Task         `xml:"urn:vim25 VStorageObjectExtendDiskEx_Task,omitempty"`
	Res    *types.VStorageObjectExtendDiskEx_TaskResponse `xml:"VStorageObjectExtendDiskEx_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VStorageObjectExtendDiskEx_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func VStorageObjectExtendDiskEx_Task(ctx context.Context, r soap.RoundTripper, req *types.VStorageObjectExtendDiskEx_Task) (*types.VStorageObjectExtendDiskEx_TaskResponse, error) {
	var reqBody, resBody VStorageObjectExtendDiskEx_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateCredentialsInGuestBody struct {
	Req    *types.ValidateCredentialsInGuest         `xml:"urn:vim25 ValidateCredentialsInGuest,omitempty"`
	Res    *types.ValidateCredentialsInGuestResponse `xml:"ValidateCredentialsInGuestResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateCredentialsInGuestBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateCredentialsInGuest(ctx context.Context, r soap.RoundTripper, req *types.ValidateCredentialsInGuest) (*types.ValidateCredentialsInGuestResponse, error) {
	var reqBody, resBody ValidateCredentialsInGuestBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateHCIConfigurationBody struct {
	Req    *types.ValidateHCIConfiguration         `xml:"urn:vim25 ValidateHCIConfiguration,omitempty"`
	Res    *types.ValidateHCIConfigurationResponse `xml:"ValidateHCIConfigurationResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateHCIConfigurationBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateHCIConfiguration(ctx context.Context, r soap.RoundTripper, req *types.ValidateHCIConfiguration) (*types.ValidateHCIConfigurationResponse, error) {
	var reqBody, resBody ValidateHCIConfigurationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateHostBody struct {
	Req    *types.ValidateHost         `xml:"urn:vim25 ValidateHost,omitempty"`
	Res    *types.ValidateHostResponse `xml:"ValidateHostResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateHostBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateHost(ctx context.Context, r soap.RoundTripper, req *types.ValidateHost) (*types.ValidateHostResponse, error) {
	var reqBody, resBody ValidateHostBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateHostProfileComposition_TaskBody struct {
	Req    *types.ValidateHostProfileComposition_Task         `xml:"urn:vim25 ValidateHostProfileComposition_Task,omitempty"`
	Res    *types.ValidateHostProfileComposition_TaskResponse `xml:"ValidateHostProfileComposition_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateHostProfileComposition_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateHostProfileComposition_Task(ctx context.Context, r soap.RoundTripper, req *types.ValidateHostProfileComposition_Task) (*types.ValidateHostProfileComposition_TaskResponse, error) {
	var reqBody, resBody ValidateHostProfileComposition_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateMigrationBody struct {
	Req    *types.ValidateMigration         `xml:"urn:vim25 ValidateMigration,omitempty"`
	Res    *types.ValidateMigrationResponse `xml:"ValidateMigrationResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateMigrationBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateMigration(ctx context.Context, r soap.RoundTripper, req *types.ValidateMigration) (*types.ValidateMigrationResponse, error) {
	var reqBody, resBody ValidateMigrationBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ValidateStoragePodConfigBody struct {
	Req    *types.ValidateStoragePodConfig         `xml:"urn:vim25 ValidateStoragePodConfig,omitempty"`
	Res    *types.ValidateStoragePodConfigResponse `xml:"ValidateStoragePodConfigResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateStoragePodConfigBody) Fault() *soap.Fault { return b.Fault_ }

func ValidateStoragePodConfig(ctx context.Context, r soap.RoundTripper, req *types.ValidateStoragePodConfig) (*types.ValidateStoragePodConfigResponse, error) {
	var reqBody, resBody ValidateStoragePodConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type VstorageObjectVCenterQueryChangedDiskAreasBody struct {
	Req    *types.VstorageObjectVCenterQueryChangedDiskAreas         `xml:"urn:vim25 VstorageObjectVCenterQueryChangedDiskAreas,omitempty"`
	Res    *types.VstorageObjectVCenterQueryChangedDiskAreasResponse `xml:"VstorageObjectVCenterQueryChangedDiskAreasResponse,omitempty"`
	Fault_ *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *VstorageObjectVCenterQueryChangedDiskAreasBody) Fault() *soap.Fault { return b.Fault_ }

func VstorageObjectVCenterQueryChangedDiskAreas(ctx context.Context, r soap.RoundTripper, req *types.VstorageObjectVCenterQueryChangedDiskAreas) (*types.VstorageObjectVCenterQueryChangedDiskAreasResponse, error) {
	var reqBody, resBody VstorageObjectVCenterQueryChangedDiskAreasBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type WaitForUpdatesBody struct {
	Req    *types.WaitForUpdates         `xml:"urn:vim25 WaitForUpdates,omitempty"`
	Res    *types.WaitForUpdatesResponse `xml:"WaitForUpdatesResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *WaitForUpdatesBody) Fault() *soap.Fault { return b.Fault_ }

func WaitForUpdates(ctx context.Context, r soap.RoundTripper, req *types.WaitForUpdates) (*types.WaitForUpdatesResponse, error) {
	var reqBody, resBody WaitForUpdatesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type WaitForUpdatesExBody struct {
	Req    *types.WaitForUpdatesEx         `xml:"urn:vim25 WaitForUpdatesEx,omitempty"`
	Res    *types.WaitForUpdatesExResponse `xml:"WaitForUpdatesExResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *WaitForUpdatesExBody) Fault() *soap.Fault { return b.Fault_ }

func WaitForUpdatesEx(ctx context.Context, r soap.RoundTripper, req *types.WaitForUpdatesEx) (*types.WaitForUpdatesExResponse, error) {
	var reqBody, resBody WaitForUpdatesExBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type XmlToCustomizationSpecItemBody struct {
	Req    *types.XmlToCustomizationSpecItem         `xml:"urn:vim25 XmlToCustomizationSpecItem,omitempty"`
	Res    *types.XmlToCustomizationSpecItemResponse `xml:"XmlToCustomizationSpecItemResponse,omitempty"`
	Fault_ *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *XmlToCustomizationSpecItemBody) Fault() *soap.Fault { return b.Fault_ }

func XmlToCustomizationSpecItem(ctx context.Context, r soap.RoundTripper, req *types.XmlToCustomizationSpecItem) (*types.XmlToCustomizationSpecItemResponse, error) {
	var reqBody, resBody XmlToCustomizationSpecItemBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ZeroFillVirtualDisk_TaskBody struct {
	Req    *types.ZeroFillVirtualDisk_Task         `xml:"urn:vim25 ZeroFillVirtualDisk_Task,omitempty"`
	Res    *types.ZeroFillVirtualDisk_TaskResponse `xml:"ZeroFillVirtualDisk_TaskResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ZeroFillVirtualDisk_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ZeroFillVirtualDisk_Task(ctx context.Context, r soap.RoundTripper, req *types.ZeroFillVirtualDisk_Task) (*types.ZeroFillVirtualDisk_TaskResponse, error) {
	var reqBody, resBody ZeroFillVirtualDisk_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ConfigureVcha_TaskBody struct {
	Req    *types.ConfigureVcha_Task         `xml:"urn:vim25 configureVcha_Task,omitempty"`
	Res    *types.ConfigureVcha_TaskResponse `xml:"configureVcha_TaskResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureVcha_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ConfigureVcha_Task(ctx context.Context, r soap.RoundTripper, req *types.ConfigureVcha_Task) (*types.ConfigureVcha_TaskResponse, error) {
	var reqBody, resBody ConfigureVcha_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreatePassiveNode_TaskBody struct {
	Req    *types.CreatePassiveNode_Task         `xml:"urn:vim25 createPassiveNode_Task,omitempty"`
	Res    *types.CreatePassiveNode_TaskResponse `xml:"createPassiveNode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreatePassiveNode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreatePassiveNode_Task(ctx context.Context, r soap.RoundTripper, req *types.CreatePassiveNode_Task) (*types.CreatePassiveNode_TaskResponse, error) {
	var reqBody, resBody CreatePassiveNode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateWitnessNode_TaskBody struct {
	Req    *types.CreateWitnessNode_Task         `xml:"urn:vim25 createWitnessNode_Task,omitempty"`
	Res    *types.CreateWitnessNode_TaskResponse `xml:"createWitnessNode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateWitnessNode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func CreateWitnessNode_Task(ctx context.Context, r soap.RoundTripper, req *types.CreateWitnessNode_Task) (*types.CreateWitnessNode_TaskResponse, error) {
	var reqBody, resBody CreateWitnessNode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DeployVcha_TaskBody struct {
	Req    *types.DeployVcha_Task         `xml:"urn:vim25 deployVcha_Task,omitempty"`
	Res    *types.DeployVcha_TaskResponse `xml:"deployVcha_TaskResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeployVcha_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DeployVcha_Task(ctx context.Context, r soap.RoundTripper, req *types.DeployVcha_Task) (*types.DeployVcha_TaskResponse, error) {
	var reqBody, resBody DeployVcha_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyVcha_TaskBody struct {
	Req    *types.DestroyVcha_Task         `xml:"urn:vim25 destroyVcha_Task,omitempty"`
	Res    *types.DestroyVcha_TaskResponse `xml:"destroyVcha_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyVcha_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyVcha_Task(ctx context.Context, r soap.RoundTripper, req *types.DestroyVcha_Task) (*types.DestroyVcha_TaskResponse, error) {
	var reqBody, resBody DestroyVcha_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type FetchSoftwarePackagesBody struct {
	Req    *types.FetchSoftwarePackages         `xml:"urn:vim25 fetchSoftwarePackages,omitempty"`
	Res    *types.FetchSoftwarePackagesResponse `xml:"fetchSoftwarePackagesResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchSoftwarePackagesBody) Fault() *soap.Fault { return b.Fault_ }

func FetchSoftwarePackages(ctx context.Context, r soap.RoundTripper, req *types.FetchSoftwarePackages) (*types.FetchSoftwarePackagesResponse, error) {
	var reqBody, resBody FetchSoftwarePackagesBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetClusterModeBody struct {
	Req    *types.GetClusterMode         `xml:"urn:vim25 getClusterMode,omitempty"`
	Res    *types.GetClusterModeResponse `xml:"getClusterModeResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetClusterModeBody) Fault() *soap.Fault { return b.Fault_ }

func GetClusterMode(ctx context.Context, r soap.RoundTripper, req *types.GetClusterMode) (*types.GetClusterModeResponse, error) {
	var reqBody, resBody GetClusterModeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetVchaConfigBody struct {
	Req    *types.GetVchaConfig         `xml:"urn:vim25 getVchaConfig,omitempty"`
	Res    *types.GetVchaConfigResponse `xml:"getVchaConfigResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetVchaConfigBody) Fault() *soap.Fault { return b.Fault_ }

func GetVchaConfig(ctx context.Context, r soap.RoundTripper, req *types.GetVchaConfig) (*types.GetVchaConfigResponse, error) {
	var reqBody, resBody GetVchaConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InitiateFailover_TaskBody struct {
	Req    *types.InitiateFailover_Task         `xml:"urn:vim25 initiateFailover_Task,omitempty"`
	Res    *types.InitiateFailover_TaskResponse `xml:"initiateFailover_TaskResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateFailover_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func InitiateFailover_Task(ctx context.Context, r soap.RoundTripper, req *types.InitiateFailover_Task) (*types.InitiateFailover_TaskResponse, error) {
	var reqBody, resBody InitiateFailover_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type InstallDateBody struct {
	Req    *types.InstallDate         `xml:"urn:vim25 installDate,omitempty"`
	Res    *types.InstallDateResponse `xml:"installDateResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallDateBody) Fault() *soap.Fault { return b.Fault_ }

func InstallDate(ctx context.Context, r soap.RoundTripper, req *types.InstallDate) (*types.InstallDateResponse, error) {
	var reqBody, resBody InstallDateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type PrepareVcha_TaskBody struct {
	Req    *types.PrepareVcha_Task         `xml:"urn:vim25 prepareVcha_Task,omitempty"`
	Res    *types.PrepareVcha_TaskResponse `xml:"prepareVcha_TaskResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PrepareVcha_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func PrepareVcha_Task(ctx context.Context, r soap.RoundTripper, req *types.PrepareVcha_Task) (*types.PrepareVcha_TaskResponse, error) {
	var reqBody, resBody PrepareVcha_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryDatacenterConfigOptionDescriptorBody struct {
	Req    *types.QueryDatacenterConfigOptionDescriptor         `xml:"urn:vim25 queryDatacenterConfigOptionDescriptor,omitempty"`
	Res    *types.QueryDatacenterConfigOptionDescriptorResponse `xml:"queryDatacenterConfigOptionDescriptorResponse,omitempty"`
	Fault_ *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDatacenterConfigOptionDescriptorBody) Fault() *soap.Fault { return b.Fault_ }

func QueryDatacenterConfigOptionDescriptor(ctx context.Context, r soap.RoundTripper, req *types.QueryDatacenterConfigOptionDescriptor) (*types.QueryDatacenterConfigOptionDescriptorResponse, error) {
	var reqBody, resBody QueryDatacenterConfigOptionDescriptorBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ReloadVirtualMachineFromPath_TaskBody struct {
	Req    *types.ReloadVirtualMachineFromPath_Task         `xml:"urn:vim25 reloadVirtualMachineFromPath_Task,omitempty"`
	Res    *types.ReloadVirtualMachineFromPath_TaskResponse `xml:"reloadVirtualMachineFromPath_TaskResponse,omitempty"`
	Fault_ *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReloadVirtualMachineFromPath_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func ReloadVirtualMachineFromPath_Task(ctx context.Context, r soap.RoundTripper, req *types.ReloadVirtualMachineFromPath_Task) (*types.ReloadVirtualMachineFromPath_TaskResponse, error) {
	var reqBody, resBody ReloadVirtualMachineFromPath_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetClusterMode_TaskBody struct {
	Req    *types.SetClusterMode_Task         `xml:"urn:vim25 setClusterMode_Task,omitempty"`
	Res    *types.SetClusterMode_TaskResponse `xml:"setClusterMode_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetClusterMode_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func SetClusterMode_Task(ctx context.Context, r soap.RoundTripper, req *types.SetClusterMode_Task) (*types.SetClusterMode_TaskResponse, error) {
	var reqBody, resBody SetClusterMode_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetCustomValueBody struct {
	Req    *types.SetCustomValue         `xml:"urn:vim25 setCustomValue,omitempty"`
	Res    *types.SetCustomValueResponse `xml:"setCustomValueResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetCustomValueBody) Fault() *soap.Fault { return b.Fault_ }

func SetCustomValue(ctx context.Context, r soap.RoundTripper, req *types.SetCustomValue) (*types.SetCustomValueResponse, error) {
	var reqBody, resBody SetCustomValueBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type StartDpuFailoverBody struct {
	Req    *types.StartDpuFailover         `xml:"urn:vim25 startDpuFailover,omitempty"`
	Res    *types.StartDpuFailoverResponse `xml:"startDpuFailoverResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartDpuFailoverBody) Fault() *soap.Fault { return b.Fault_ }

func StartDpuFailover(ctx context.Context, r soap.RoundTripper, req *types.StartDpuFailover) (*types.StartDpuFailoverResponse, error) {
	var reqBody, resBody StartDpuFailoverBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterVApp_TaskBody struct {
	Req    *types.UnregisterVApp_Task         `xml:"urn:vim25 unregisterVApp_Task,omitempty"`
	Res    *types.UnregisterVApp_TaskResponse `xml:"unregisterVApp_TaskResponse,omitempty"`
	Fault_ *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterVApp_TaskBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterVApp_Task(ctx context.Context, r soap.RoundTripper, req *types.UnregisterVApp_Task) (*types.UnregisterVApp_TaskResponse, error) {
	var reqBody, resBody UnregisterVApp_TaskBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
