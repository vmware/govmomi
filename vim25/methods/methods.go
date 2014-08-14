/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type AcknowledgeAlarmBody struct {
	Req   *types.AcknowledgeAlarm         `xml:"urn:vim25 AcknowledgeAlarm,omitempty"`
	Res   *types.AcknowledgeAlarmResponse `xml:"urn:vim25 AcknowledgeAlarmResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcknowledgeAlarmBody) fault() *soap.Fault { return b.Fault }

func AcknowledgeAlarm(r soap.RoundTripper, req *types.AcknowledgeAlarm) (*types.AcknowledgeAlarmResponse, Error) {
	var body = AcknowledgeAlarmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireCimServicesTicketBody struct {
	Req   *types.AcquireCimServicesTicket         `xml:"urn:vim25 AcquireCimServicesTicket,omitempty"`
	Res   *types.AcquireCimServicesTicketResponse `xml:"urn:vim25 AcquireCimServicesTicketResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCimServicesTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireCimServicesTicket(r soap.RoundTripper, req *types.AcquireCimServicesTicket) (*types.AcquireCimServicesTicketResponse, Error) {
	var body = AcquireCimServicesTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireCloneTicketBody struct {
	Req   *types.AcquireCloneTicket         `xml:"urn:vim25 AcquireCloneTicket,omitempty"`
	Res   *types.AcquireCloneTicketResponse `xml:"urn:vim25 AcquireCloneTicketResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCloneTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireCloneTicket(r soap.RoundTripper, req *types.AcquireCloneTicket) (*types.AcquireCloneTicketResponse, Error) {
	var body = AcquireCloneTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireCredentialsInGuestBody struct {
	Req   *types.AcquireCredentialsInGuest         `xml:"urn:vim25 AcquireCredentialsInGuest,omitempty"`
	Res   *types.AcquireCredentialsInGuestResponse `xml:"urn:vim25 AcquireCredentialsInGuestResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireCredentialsInGuestBody) fault() *soap.Fault { return b.Fault }

func AcquireCredentialsInGuest(r soap.RoundTripper, req *types.AcquireCredentialsInGuest) (*types.AcquireCredentialsInGuestResponse, Error) {
	var body = AcquireCredentialsInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireGenericServiceTicketBody struct {
	Req   *types.AcquireGenericServiceTicket         `xml:"urn:vim25 AcquireGenericServiceTicket,omitempty"`
	Res   *types.AcquireGenericServiceTicketResponse `xml:"urn:vim25 AcquireGenericServiceTicketResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireGenericServiceTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireGenericServiceTicket(r soap.RoundTripper, req *types.AcquireGenericServiceTicket) (*types.AcquireGenericServiceTicketResponse, Error) {
	var body = AcquireGenericServiceTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireLocalTicketBody struct {
	Req   *types.AcquireLocalTicket         `xml:"urn:vim25 AcquireLocalTicket,omitempty"`
	Res   *types.AcquireLocalTicketResponse `xml:"urn:vim25 AcquireLocalTicketResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireLocalTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireLocalTicket(r soap.RoundTripper, req *types.AcquireLocalTicket) (*types.AcquireLocalTicketResponse, Error) {
	var body = AcquireLocalTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireMksTicketBody struct {
	Req   *types.AcquireMksTicket         `xml:"urn:vim25 AcquireMksTicket,omitempty"`
	Res   *types.AcquireMksTicketResponse `xml:"urn:vim25 AcquireMksTicketResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireMksTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireMksTicket(r soap.RoundTripper, req *types.AcquireMksTicket) (*types.AcquireMksTicketResponse, Error) {
	var body = AcquireMksTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AcquireTicketBody struct {
	Req   *types.AcquireTicket         `xml:"urn:vim25 AcquireTicket,omitempty"`
	Res   *types.AcquireTicketResponse `xml:"urn:vim25 AcquireTicketResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireTicketBody) fault() *soap.Fault { return b.Fault }

func AcquireTicket(r soap.RoundTripper, req *types.AcquireTicket) (*types.AcquireTicketResponse, Error) {
	var body = AcquireTicketBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddAuthorizationRoleBody struct {
	Req   *types.AddAuthorizationRole         `xml:"urn:vim25 AddAuthorizationRole,omitempty"`
	Res   *types.AddAuthorizationRoleResponse `xml:"urn:vim25 AddAuthorizationRoleResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddAuthorizationRoleBody) fault() *soap.Fault { return b.Fault }

func AddAuthorizationRole(r soap.RoundTripper, req *types.AddAuthorizationRole) (*types.AddAuthorizationRoleResponse, Error) {
	var body = AddAuthorizationRoleBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddCustomFieldDefBody struct {
	Req   *types.AddCustomFieldDef         `xml:"urn:vim25 AddCustomFieldDef,omitempty"`
	Res   *types.AddCustomFieldDefResponse `xml:"urn:vim25 AddCustomFieldDefResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddCustomFieldDefBody) fault() *soap.Fault { return b.Fault }

func AddCustomFieldDef(r soap.RoundTripper, req *types.AddCustomFieldDef) (*types.AddCustomFieldDefResponse, Error) {
	var body = AddCustomFieldDefBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddDVPortgroup_TaskBody struct {
	Req   *types.AddDVPortgroup_Task         `xml:"urn:vim25 AddDVPortgroup_Task,omitempty"`
	Res   *types.AddDVPortgroup_TaskResponse `xml:"urn:vim25 AddDVPortgroup_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddDVPortgroup_TaskBody) fault() *soap.Fault { return b.Fault }

func AddDVPortgroup_Task(r soap.RoundTripper, req *types.AddDVPortgroup_Task) (*types.AddDVPortgroup_TaskResponse, Error) {
	var body = AddDVPortgroup_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddDisks_TaskBody struct {
	Req   *types.AddDisks_Task         `xml:"urn:vim25 AddDisks_Task,omitempty"`
	Res   *types.AddDisks_TaskResponse `xml:"urn:vim25 AddDisks_TaskResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddDisks_TaskBody) fault() *soap.Fault { return b.Fault }

func AddDisks_Task(r soap.RoundTripper, req *types.AddDisks_Task) (*types.AddDisks_TaskResponse, Error) {
	var body = AddDisks_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddHost_TaskBody struct {
	Req   *types.AddHost_Task         `xml:"urn:vim25 AddHost_Task,omitempty"`
	Res   *types.AddHost_TaskResponse `xml:"urn:vim25 AddHost_TaskResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddHost_TaskBody) fault() *soap.Fault { return b.Fault }

func AddHost_Task(r soap.RoundTripper, req *types.AddHost_Task) (*types.AddHost_TaskResponse, Error) {
	var body = AddHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddInternetScsiSendTargetsBody struct {
	Req   *types.AddInternetScsiSendTargets         `xml:"urn:vim25 AddInternetScsiSendTargets,omitempty"`
	Res   *types.AddInternetScsiSendTargetsResponse `xml:"urn:vim25 AddInternetScsiSendTargetsResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddInternetScsiSendTargetsBody) fault() *soap.Fault { return b.Fault }

func AddInternetScsiSendTargets(r soap.RoundTripper, req *types.AddInternetScsiSendTargets) (*types.AddInternetScsiSendTargetsResponse, Error) {
	var body = AddInternetScsiSendTargetsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddInternetScsiStaticTargetsBody struct {
	Req   *types.AddInternetScsiStaticTargets         `xml:"urn:vim25 AddInternetScsiStaticTargets,omitempty"`
	Res   *types.AddInternetScsiStaticTargetsResponse `xml:"urn:vim25 AddInternetScsiStaticTargetsResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddInternetScsiStaticTargetsBody) fault() *soap.Fault { return b.Fault }

func AddInternetScsiStaticTargets(r soap.RoundTripper, req *types.AddInternetScsiStaticTargets) (*types.AddInternetScsiStaticTargetsResponse, Error) {
	var body = AddInternetScsiStaticTargetsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddLicenseBody struct {
	Req   *types.AddLicense         `xml:"urn:vim25 AddLicense,omitempty"`
	Res   *types.AddLicenseResponse `xml:"urn:vim25 AddLicenseResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddLicenseBody) fault() *soap.Fault { return b.Fault }

func AddLicense(r soap.RoundTripper, req *types.AddLicense) (*types.AddLicenseResponse, Error) {
	var body = AddLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddNetworkResourcePoolBody struct {
	Req   *types.AddNetworkResourcePool         `xml:"urn:vim25 AddNetworkResourcePool,omitempty"`
	Res   *types.AddNetworkResourcePoolResponse `xml:"urn:vim25 AddNetworkResourcePoolResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddNetworkResourcePoolBody) fault() *soap.Fault { return b.Fault }

func AddNetworkResourcePool(r soap.RoundTripper, req *types.AddNetworkResourcePool) (*types.AddNetworkResourcePoolResponse, Error) {
	var body = AddNetworkResourcePoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddPortGroupBody struct {
	Req   *types.AddPortGroup         `xml:"urn:vim25 AddPortGroup,omitempty"`
	Res   *types.AddPortGroupResponse `xml:"urn:vim25 AddPortGroupResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddPortGroupBody) fault() *soap.Fault { return b.Fault }

func AddPortGroup(r soap.RoundTripper, req *types.AddPortGroup) (*types.AddPortGroupResponse, Error) {
	var body = AddPortGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddServiceConsoleVirtualNicBody struct {
	Req   *types.AddServiceConsoleVirtualNic         `xml:"urn:vim25 AddServiceConsoleVirtualNic,omitempty"`
	Res   *types.AddServiceConsoleVirtualNicResponse `xml:"urn:vim25 AddServiceConsoleVirtualNicResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddServiceConsoleVirtualNicBody) fault() *soap.Fault { return b.Fault }

func AddServiceConsoleVirtualNic(r soap.RoundTripper, req *types.AddServiceConsoleVirtualNic) (*types.AddServiceConsoleVirtualNicResponse, Error) {
	var body = AddServiceConsoleVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddStandaloneHost_TaskBody struct {
	Req   *types.AddStandaloneHost_Task         `xml:"urn:vim25 AddStandaloneHost_Task,omitempty"`
	Res   *types.AddStandaloneHost_TaskResponse `xml:"urn:vim25 AddStandaloneHost_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddStandaloneHost_TaskBody) fault() *soap.Fault { return b.Fault }

func AddStandaloneHost_Task(r soap.RoundTripper, req *types.AddStandaloneHost_Task) (*types.AddStandaloneHost_TaskResponse, Error) {
	var body = AddStandaloneHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddVirtualNicBody struct {
	Req   *types.AddVirtualNic         `xml:"urn:vim25 AddVirtualNic,omitempty"`
	Res   *types.AddVirtualNicResponse `xml:"urn:vim25 AddVirtualNicResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddVirtualNicBody) fault() *soap.Fault { return b.Fault }

func AddVirtualNic(r soap.RoundTripper, req *types.AddVirtualNic) (*types.AddVirtualNicResponse, Error) {
	var body = AddVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AddVirtualSwitchBody struct {
	Req   *types.AddVirtualSwitch         `xml:"urn:vim25 AddVirtualSwitch,omitempty"`
	Res   *types.AddVirtualSwitchResponse `xml:"urn:vim25 AddVirtualSwitchResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddVirtualSwitchBody) fault() *soap.Fault { return b.Fault }

func AddVirtualSwitch(r soap.RoundTripper, req *types.AddVirtualSwitch) (*types.AddVirtualSwitchResponse, Error) {
	var body = AddVirtualSwitchBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AllocateIpv4AddressBody struct {
	Req   *types.AllocateIpv4Address         `xml:"urn:vim25 AllocateIpv4Address,omitempty"`
	Res   *types.AllocateIpv4AddressResponse `xml:"urn:vim25 AllocateIpv4AddressResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AllocateIpv4AddressBody) fault() *soap.Fault { return b.Fault }

func AllocateIpv4Address(r soap.RoundTripper, req *types.AllocateIpv4Address) (*types.AllocateIpv4AddressResponse, Error) {
	var body = AllocateIpv4AddressBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AllocateIpv6AddressBody struct {
	Req   *types.AllocateIpv6Address         `xml:"urn:vim25 AllocateIpv6Address,omitempty"`
	Res   *types.AllocateIpv6AddressResponse `xml:"urn:vim25 AllocateIpv6AddressResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AllocateIpv6AddressBody) fault() *soap.Fault { return b.Fault }

func AllocateIpv6Address(r soap.RoundTripper, req *types.AllocateIpv6Address) (*types.AllocateIpv6AddressResponse, Error) {
	var body = AllocateIpv6AddressBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AnswerVMBody struct {
	Req   *types.AnswerVM         `xml:"urn:vim25 AnswerVM,omitempty"`
	Res   *types.AnswerVMResponse `xml:"urn:vim25 AnswerVMResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AnswerVMBody) fault() *soap.Fault { return b.Fault }

func AnswerVM(r soap.RoundTripper, req *types.AnswerVM) (*types.AnswerVMResponse, Error) {
	var body = AnswerVMBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ApplyHostConfig_TaskBody struct {
	Req   *types.ApplyHostConfig_Task         `xml:"urn:vim25 ApplyHostConfig_Task,omitempty"`
	Res   *types.ApplyHostConfig_TaskResponse `xml:"urn:vim25 ApplyHostConfig_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyHostConfig_TaskBody) fault() *soap.Fault { return b.Fault }

func ApplyHostConfig_Task(r soap.RoundTripper, req *types.ApplyHostConfig_Task) (*types.ApplyHostConfig_TaskResponse, Error) {
	var body = ApplyHostConfig_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ApplyRecommendationBody struct {
	Req   *types.ApplyRecommendation         `xml:"urn:vim25 ApplyRecommendation,omitempty"`
	Res   *types.ApplyRecommendationResponse `xml:"urn:vim25 ApplyRecommendationResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyRecommendationBody) fault() *soap.Fault { return b.Fault }

func ApplyRecommendation(r soap.RoundTripper, req *types.ApplyRecommendation) (*types.ApplyRecommendationResponse, Error) {
	var body = ApplyRecommendationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ApplyStorageDrsRecommendationToPod_TaskBody struct {
	Req   *types.ApplyStorageDrsRecommendationToPod_Task         `xml:"urn:vim25 ApplyStorageDrsRecommendationToPod_Task,omitempty"`
	Res   *types.ApplyStorageDrsRecommendationToPod_TaskResponse `xml:"urn:vim25 ApplyStorageDrsRecommendationToPod_TaskResponse,omitempty"`
	Fault *soap.Fault                                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyStorageDrsRecommendationToPod_TaskBody) fault() *soap.Fault { return b.Fault }

func ApplyStorageDrsRecommendationToPod_Task(r soap.RoundTripper, req *types.ApplyStorageDrsRecommendationToPod_Task) (*types.ApplyStorageDrsRecommendationToPod_TaskResponse, Error) {
	var body = ApplyStorageDrsRecommendationToPod_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ApplyStorageDrsRecommendation_TaskBody struct {
	Req   *types.ApplyStorageDrsRecommendation_Task         `xml:"urn:vim25 ApplyStorageDrsRecommendation_Task,omitempty"`
	Res   *types.ApplyStorageDrsRecommendation_TaskResponse `xml:"urn:vim25 ApplyStorageDrsRecommendation_TaskResponse,omitempty"`
	Fault *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ApplyStorageDrsRecommendation_TaskBody) fault() *soap.Fault { return b.Fault }

func ApplyStorageDrsRecommendation_Task(r soap.RoundTripper, req *types.ApplyStorageDrsRecommendation_Task) (*types.ApplyStorageDrsRecommendation_TaskResponse, Error) {
	var body = ApplyStorageDrsRecommendation_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AreAlarmActionsEnabledBody struct {
	Req   *types.AreAlarmActionsEnabled         `xml:"urn:vim25 AreAlarmActionsEnabled,omitempty"`
	Res   *types.AreAlarmActionsEnabledResponse `xml:"urn:vim25 AreAlarmActionsEnabledResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AreAlarmActionsEnabledBody) fault() *soap.Fault { return b.Fault }

func AreAlarmActionsEnabled(r soap.RoundTripper, req *types.AreAlarmActionsEnabled) (*types.AreAlarmActionsEnabledResponse, Error) {
	var body = AreAlarmActionsEnabledBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AssignUserToGroupBody struct {
	Req   *types.AssignUserToGroup         `xml:"urn:vim25 AssignUserToGroup,omitempty"`
	Res   *types.AssignUserToGroupResponse `xml:"urn:vim25 AssignUserToGroupResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AssignUserToGroupBody) fault() *soap.Fault { return b.Fault }

func AssignUserToGroup(r soap.RoundTripper, req *types.AssignUserToGroup) (*types.AssignUserToGroupResponse, Error) {
	var body = AssignUserToGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AssociateProfileBody struct {
	Req   *types.AssociateProfile         `xml:"urn:vim25 AssociateProfile,omitempty"`
	Res   *types.AssociateProfileResponse `xml:"urn:vim25 AssociateProfileResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AssociateProfileBody) fault() *soap.Fault { return b.Fault }

func AssociateProfile(r soap.RoundTripper, req *types.AssociateProfile) (*types.AssociateProfileResponse, Error) {
	var body = AssociateProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AttachScsiLunBody struct {
	Req   *types.AttachScsiLun         `xml:"urn:vim25 AttachScsiLun,omitempty"`
	Res   *types.AttachScsiLunResponse `xml:"urn:vim25 AttachScsiLunResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachScsiLunBody) fault() *soap.Fault { return b.Fault }

func AttachScsiLun(r soap.RoundTripper, req *types.AttachScsiLun) (*types.AttachScsiLunResponse, Error) {
	var body = AttachScsiLunBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AttachVmfsExtentBody struct {
	Req   *types.AttachVmfsExtent         `xml:"urn:vim25 AttachVmfsExtent,omitempty"`
	Res   *types.AttachVmfsExtentResponse `xml:"urn:vim25 AttachVmfsExtentResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AttachVmfsExtentBody) fault() *soap.Fault { return b.Fault }

func AttachVmfsExtent(r soap.RoundTripper, req *types.AttachVmfsExtent) (*types.AttachVmfsExtentResponse, Error) {
	var body = AttachVmfsExtentBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AutoStartPowerOffBody struct {
	Req   *types.AutoStartPowerOff         `xml:"urn:vim25 AutoStartPowerOff,omitempty"`
	Res   *types.AutoStartPowerOffResponse `xml:"urn:vim25 AutoStartPowerOffResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AutoStartPowerOffBody) fault() *soap.Fault { return b.Fault }

func AutoStartPowerOff(r soap.RoundTripper, req *types.AutoStartPowerOff) (*types.AutoStartPowerOffResponse, Error) {
	var body = AutoStartPowerOffBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type AutoStartPowerOnBody struct {
	Req   *types.AutoStartPowerOn         `xml:"urn:vim25 AutoStartPowerOn,omitempty"`
	Res   *types.AutoStartPowerOnResponse `xml:"urn:vim25 AutoStartPowerOnResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AutoStartPowerOnBody) fault() *soap.Fault { return b.Fault }

func AutoStartPowerOn(r soap.RoundTripper, req *types.AutoStartPowerOn) (*types.AutoStartPowerOnResponse, Error) {
	var body = AutoStartPowerOnBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type BackupFirmwareConfigurationBody struct {
	Req   *types.BackupFirmwareConfiguration         `xml:"urn:vim25 BackupFirmwareConfiguration,omitempty"`
	Res   *types.BackupFirmwareConfigurationResponse `xml:"urn:vim25 BackupFirmwareConfigurationResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BackupFirmwareConfigurationBody) fault() *soap.Fault { return b.Fault }

func BackupFirmwareConfiguration(r soap.RoundTripper, req *types.BackupFirmwareConfiguration) (*types.BackupFirmwareConfigurationResponse, Error) {
	var body = BackupFirmwareConfigurationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type BindVnicBody struct {
	Req   *types.BindVnic         `xml:"urn:vim25 BindVnic,omitempty"`
	Res   *types.BindVnicResponse `xml:"urn:vim25 BindVnicResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BindVnicBody) fault() *soap.Fault { return b.Fault }

func BindVnic(r soap.RoundTripper, req *types.BindVnic) (*types.BindVnicResponse, Error) {
	var body = BindVnicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type BrowseDiagnosticLogBody struct {
	Req   *types.BrowseDiagnosticLog         `xml:"urn:vim25 BrowseDiagnosticLog,omitempty"`
	Res   *types.BrowseDiagnosticLogResponse `xml:"urn:vim25 BrowseDiagnosticLogResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *BrowseDiagnosticLogBody) fault() *soap.Fault { return b.Fault }

func BrowseDiagnosticLog(r soap.RoundTripper, req *types.BrowseDiagnosticLog) (*types.BrowseDiagnosticLogResponse, Error) {
	var body = BrowseDiagnosticLogBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CancelRecommendationBody struct {
	Req   *types.CancelRecommendation         `xml:"urn:vim25 CancelRecommendation,omitempty"`
	Res   *types.CancelRecommendationResponse `xml:"urn:vim25 CancelRecommendationResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelRecommendationBody) fault() *soap.Fault { return b.Fault }

func CancelRecommendation(r soap.RoundTripper, req *types.CancelRecommendation) (*types.CancelRecommendationResponse, Error) {
	var body = CancelRecommendationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CancelRetrievePropertiesExBody struct {
	Req   *types.CancelRetrievePropertiesEx         `xml:"urn:vim25 CancelRetrievePropertiesEx,omitempty"`
	Res   *types.CancelRetrievePropertiesExResponse `xml:"urn:vim25 CancelRetrievePropertiesExResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelRetrievePropertiesExBody) fault() *soap.Fault { return b.Fault }

func CancelRetrievePropertiesEx(r soap.RoundTripper, req *types.CancelRetrievePropertiesEx) (*types.CancelRetrievePropertiesExResponse, Error) {
	var body = CancelRetrievePropertiesExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CancelStorageDrsRecommendationBody struct {
	Req   *types.CancelStorageDrsRecommendation         `xml:"urn:vim25 CancelStorageDrsRecommendation,omitempty"`
	Res   *types.CancelStorageDrsRecommendationResponse `xml:"urn:vim25 CancelStorageDrsRecommendationResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelStorageDrsRecommendationBody) fault() *soap.Fault { return b.Fault }

func CancelStorageDrsRecommendation(r soap.RoundTripper, req *types.CancelStorageDrsRecommendation) (*types.CancelStorageDrsRecommendationResponse, Error) {
	var body = CancelStorageDrsRecommendationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CancelTaskBody struct {
	Req   *types.CancelTask         `xml:"urn:vim25 CancelTask,omitempty"`
	Res   *types.CancelTaskResponse `xml:"urn:vim25 CancelTaskResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelTaskBody) fault() *soap.Fault { return b.Fault }

func CancelTask(r soap.RoundTripper, req *types.CancelTask) (*types.CancelTaskResponse, Error) {
	var body = CancelTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CancelWaitForUpdatesBody struct {
	Req   *types.CancelWaitForUpdates         `xml:"urn:vim25 CancelWaitForUpdates,omitempty"`
	Res   *types.CancelWaitForUpdatesResponse `xml:"urn:vim25 CancelWaitForUpdatesResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CancelWaitForUpdatesBody) fault() *soap.Fault { return b.Fault }

func CancelWaitForUpdates(r soap.RoundTripper, req *types.CancelWaitForUpdates) (*types.CancelWaitForUpdatesResponse, Error) {
	var body = CancelWaitForUpdatesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ChangeFileAttributesInGuestBody struct {
	Req   *types.ChangeFileAttributesInGuest         `xml:"urn:vim25 ChangeFileAttributesInGuest,omitempty"`
	Res   *types.ChangeFileAttributesInGuestResponse `xml:"urn:vim25 ChangeFileAttributesInGuestResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeFileAttributesInGuestBody) fault() *soap.Fault { return b.Fault }

func ChangeFileAttributesInGuest(r soap.RoundTripper, req *types.ChangeFileAttributesInGuest) (*types.ChangeFileAttributesInGuestResponse, Error) {
	var body = ChangeFileAttributesInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ChangeOwnerBody struct {
	Req   *types.ChangeOwner         `xml:"urn:vim25 ChangeOwner,omitempty"`
	Res   *types.ChangeOwnerResponse `xml:"urn:vim25 ChangeOwnerResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ChangeOwnerBody) fault() *soap.Fault { return b.Fault }

func ChangeOwner(r soap.RoundTripper, req *types.ChangeOwner) (*types.ChangeOwnerResponse, Error) {
	var body = ChangeOwnerBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckAnswerFileStatus_TaskBody struct {
	Req   *types.CheckAnswerFileStatus_Task         `xml:"urn:vim25 CheckAnswerFileStatus_Task,omitempty"`
	Res   *types.CheckAnswerFileStatus_TaskResponse `xml:"urn:vim25 CheckAnswerFileStatus_TaskResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckAnswerFileStatus_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckAnswerFileStatus_Task(r soap.RoundTripper, req *types.CheckAnswerFileStatus_Task) (*types.CheckAnswerFileStatus_TaskResponse, Error) {
	var body = CheckAnswerFileStatus_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckCompatibility_TaskBody struct {
	Req   *types.CheckCompatibility_Task         `xml:"urn:vim25 CheckCompatibility_Task,omitempty"`
	Res   *types.CheckCompatibility_TaskResponse `xml:"urn:vim25 CheckCompatibility_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCompatibility_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckCompatibility_Task(r soap.RoundTripper, req *types.CheckCompatibility_Task) (*types.CheckCompatibility_TaskResponse, Error) {
	var body = CheckCompatibility_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckCompliance_TaskBody struct {
	Req   *types.CheckCompliance_Task         `xml:"urn:vim25 CheckCompliance_Task,omitempty"`
	Res   *types.CheckCompliance_TaskResponse `xml:"urn:vim25 CheckCompliance_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCompliance_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckCompliance_Task(r soap.RoundTripper, req *types.CheckCompliance_Task) (*types.CheckCompliance_TaskResponse, Error) {
	var body = CheckCompliance_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckCustomizationResourcesBody struct {
	Req   *types.CheckCustomizationResources         `xml:"urn:vim25 CheckCustomizationResources,omitempty"`
	Res   *types.CheckCustomizationResourcesResponse `xml:"urn:vim25 CheckCustomizationResourcesResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCustomizationResourcesBody) fault() *soap.Fault { return b.Fault }

func CheckCustomizationResources(r soap.RoundTripper, req *types.CheckCustomizationResources) (*types.CheckCustomizationResourcesResponse, Error) {
	var body = CheckCustomizationResourcesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckCustomizationSpecBody struct {
	Req   *types.CheckCustomizationSpec         `xml:"urn:vim25 CheckCustomizationSpec,omitempty"`
	Res   *types.CheckCustomizationSpecResponse `xml:"urn:vim25 CheckCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func CheckCustomizationSpec(r soap.RoundTripper, req *types.CheckCustomizationSpec) (*types.CheckCustomizationSpecResponse, Error) {
	var body = CheckCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckForUpdatesBody struct {
	Req   *types.CheckForUpdates         `xml:"urn:vim25 CheckForUpdates,omitempty"`
	Res   *types.CheckForUpdatesResponse `xml:"urn:vim25 CheckForUpdatesResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckForUpdatesBody) fault() *soap.Fault { return b.Fault }

func CheckForUpdates(r soap.RoundTripper, req *types.CheckForUpdates) (*types.CheckForUpdatesResponse, Error) {
	var body = CheckForUpdatesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckHostPatch_TaskBody struct {
	Req   *types.CheckHostPatch_Task         `xml:"urn:vim25 CheckHostPatch_Task,omitempty"`
	Res   *types.CheckHostPatch_TaskResponse `xml:"urn:vim25 CheckHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckHostPatch_Task(r soap.RoundTripper, req *types.CheckHostPatch_Task) (*types.CheckHostPatch_TaskResponse, Error) {
	var body = CheckHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckLicenseFeatureBody struct {
	Req   *types.CheckLicenseFeature         `xml:"urn:vim25 CheckLicenseFeature,omitempty"`
	Res   *types.CheckLicenseFeatureResponse `xml:"urn:vim25 CheckLicenseFeatureResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckLicenseFeatureBody) fault() *soap.Fault { return b.Fault }

func CheckLicenseFeature(r soap.RoundTripper, req *types.CheckLicenseFeature) (*types.CheckLicenseFeatureResponse, Error) {
	var body = CheckLicenseFeatureBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckMigrate_TaskBody struct {
	Req   *types.CheckMigrate_Task         `xml:"urn:vim25 CheckMigrate_Task,omitempty"`
	Res   *types.CheckMigrate_TaskResponse `xml:"urn:vim25 CheckMigrate_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckMigrate_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckMigrate_Task(r soap.RoundTripper, req *types.CheckMigrate_Task) (*types.CheckMigrate_TaskResponse, Error) {
	var body = CheckMigrate_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckProfileCompliance_TaskBody struct {
	Req   *types.CheckProfileCompliance_Task         `xml:"urn:vim25 CheckProfileCompliance_Task,omitempty"`
	Res   *types.CheckProfileCompliance_TaskResponse `xml:"urn:vim25 CheckProfileCompliance_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckProfileCompliance_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckProfileCompliance_Task(r soap.RoundTripper, req *types.CheckProfileCompliance_Task) (*types.CheckProfileCompliance_TaskResponse, Error) {
	var body = CheckProfileCompliance_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CheckRelocate_TaskBody struct {
	Req   *types.CheckRelocate_Task         `xml:"urn:vim25 CheckRelocate_Task,omitempty"`
	Res   *types.CheckRelocate_TaskResponse `xml:"urn:vim25 CheckRelocate_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CheckRelocate_TaskBody) fault() *soap.Fault { return b.Fault }

func CheckRelocate_Task(r soap.RoundTripper, req *types.CheckRelocate_Task) (*types.CheckRelocate_TaskResponse, Error) {
	var body = CheckRelocate_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ClearComplianceStatusBody struct {
	Req   *types.ClearComplianceStatus         `xml:"urn:vim25 ClearComplianceStatus,omitempty"`
	Res   *types.ClearComplianceStatusResponse `xml:"urn:vim25 ClearComplianceStatusResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClearComplianceStatusBody) fault() *soap.Fault { return b.Fault }

func ClearComplianceStatus(r soap.RoundTripper, req *types.ClearComplianceStatus) (*types.ClearComplianceStatusResponse, Error) {
	var body = ClearComplianceStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CloneSessionBody struct {
	Req   *types.CloneSession         `xml:"urn:vim25 CloneSession,omitempty"`
	Res   *types.CloneSessionResponse `xml:"urn:vim25 CloneSessionResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneSessionBody) fault() *soap.Fault { return b.Fault }

func CloneSession(r soap.RoundTripper, req *types.CloneSession) (*types.CloneSessionResponse, Error) {
	var body = CloneSessionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CloneVApp_TaskBody struct {
	Req   *types.CloneVApp_Task         `xml:"urn:vim25 CloneVApp_Task,omitempty"`
	Res   *types.CloneVApp_TaskResponse `xml:"urn:vim25 CloneVApp_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneVApp_TaskBody) fault() *soap.Fault { return b.Fault }

func CloneVApp_Task(r soap.RoundTripper, req *types.CloneVApp_Task) (*types.CloneVApp_TaskResponse, Error) {
	var body = CloneVApp_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CloneVM_TaskBody struct {
	Req   *types.CloneVM_Task         `xml:"urn:vim25 CloneVM_Task,omitempty"`
	Res   *types.CloneVM_TaskResponse `xml:"urn:vim25 CloneVM_TaskResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloneVM_TaskBody) fault() *soap.Fault { return b.Fault }

func CloneVM_Task(r soap.RoundTripper, req *types.CloneVM_Task) (*types.CloneVM_TaskResponse, Error) {
	var body = CloneVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CloseInventoryViewFolderBody struct {
	Req   *types.CloseInventoryViewFolder         `xml:"urn:vim25 CloseInventoryViewFolder,omitempty"`
	Res   *types.CloseInventoryViewFolderResponse `xml:"urn:vim25 CloseInventoryViewFolderResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CloseInventoryViewFolderBody) fault() *soap.Fault { return b.Fault }

func CloseInventoryViewFolder(r soap.RoundTripper, req *types.CloseInventoryViewFolder) (*types.CloseInventoryViewFolderResponse, Error) {
	var body = CloseInventoryViewFolderBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ClusterEnterMaintenanceModeBody struct {
	Req   *types.ClusterEnterMaintenanceMode         `xml:"urn:vim25 ClusterEnterMaintenanceMode,omitempty"`
	Res   *types.ClusterEnterMaintenanceModeResponse `xml:"urn:vim25 ClusterEnterMaintenanceModeResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ClusterEnterMaintenanceModeBody) fault() *soap.Fault { return b.Fault }

func ClusterEnterMaintenanceMode(r soap.RoundTripper, req *types.ClusterEnterMaintenanceMode) (*types.ClusterEnterMaintenanceModeResponse, Error) {
	var body = ClusterEnterMaintenanceModeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ComputeDiskPartitionInfoBody struct {
	Req   *types.ComputeDiskPartitionInfo         `xml:"urn:vim25 ComputeDiskPartitionInfo,omitempty"`
	Res   *types.ComputeDiskPartitionInfoResponse `xml:"urn:vim25 ComputeDiskPartitionInfoResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ComputeDiskPartitionInfoBody) fault() *soap.Fault { return b.Fault }

func ComputeDiskPartitionInfo(r soap.RoundTripper, req *types.ComputeDiskPartitionInfo) (*types.ComputeDiskPartitionInfoResponse, Error) {
	var body = ComputeDiskPartitionInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ComputeDiskPartitionInfoForResizeBody struct {
	Req   *types.ComputeDiskPartitionInfoForResize         `xml:"urn:vim25 ComputeDiskPartitionInfoForResize,omitempty"`
	Res   *types.ComputeDiskPartitionInfoForResizeResponse `xml:"urn:vim25 ComputeDiskPartitionInfoForResizeResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ComputeDiskPartitionInfoForResizeBody) fault() *soap.Fault { return b.Fault }

func ComputeDiskPartitionInfoForResize(r soap.RoundTripper, req *types.ComputeDiskPartitionInfoForResize) (*types.ComputeDiskPartitionInfoForResizeResponse, Error) {
	var body = ComputeDiskPartitionInfoForResizeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureDatastoreIORM_TaskBody struct {
	Req   *types.ConfigureDatastoreIORM_Task         `xml:"urn:vim25 ConfigureDatastoreIORM_Task,omitempty"`
	Res   *types.ConfigureDatastoreIORM_TaskResponse `xml:"urn:vim25 ConfigureDatastoreIORM_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureDatastoreIORM_TaskBody) fault() *soap.Fault { return b.Fault }

func ConfigureDatastoreIORM_Task(r soap.RoundTripper, req *types.ConfigureDatastoreIORM_Task) (*types.ConfigureDatastoreIORM_TaskResponse, Error) {
	var body = ConfigureDatastoreIORM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureDatastorePrincipalBody struct {
	Req   *types.ConfigureDatastorePrincipal         `xml:"urn:vim25 ConfigureDatastorePrincipal,omitempty"`
	Res   *types.ConfigureDatastorePrincipalResponse `xml:"urn:vim25 ConfigureDatastorePrincipalResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureDatastorePrincipalBody) fault() *soap.Fault { return b.Fault }

func ConfigureDatastorePrincipal(r soap.RoundTripper, req *types.ConfigureDatastorePrincipal) (*types.ConfigureDatastorePrincipalResponse, Error) {
	var body = ConfigureDatastorePrincipalBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureHostCache_TaskBody struct {
	Req   *types.ConfigureHostCache_Task         `xml:"urn:vim25 ConfigureHostCache_Task,omitempty"`
	Res   *types.ConfigureHostCache_TaskResponse `xml:"urn:vim25 ConfigureHostCache_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureHostCache_TaskBody) fault() *soap.Fault { return b.Fault }

func ConfigureHostCache_Task(r soap.RoundTripper, req *types.ConfigureHostCache_Task) (*types.ConfigureHostCache_TaskResponse, Error) {
	var body = ConfigureHostCache_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureLicenseSourceBody struct {
	Req   *types.ConfigureLicenseSource         `xml:"urn:vim25 ConfigureLicenseSource,omitempty"`
	Res   *types.ConfigureLicenseSourceResponse `xml:"urn:vim25 ConfigureLicenseSourceResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureLicenseSourceBody) fault() *soap.Fault { return b.Fault }

func ConfigureLicenseSource(r soap.RoundTripper, req *types.ConfigureLicenseSource) (*types.ConfigureLicenseSourceResponse, Error) {
	var body = ConfigureLicenseSourceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigurePowerPolicyBody struct {
	Req   *types.ConfigurePowerPolicy         `xml:"urn:vim25 ConfigurePowerPolicy,omitempty"`
	Res   *types.ConfigurePowerPolicyResponse `xml:"urn:vim25 ConfigurePowerPolicyResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigurePowerPolicyBody) fault() *soap.Fault { return b.Fault }

func ConfigurePowerPolicy(r soap.RoundTripper, req *types.ConfigurePowerPolicy) (*types.ConfigurePowerPolicyResponse, Error) {
	var body = ConfigurePowerPolicyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureStorageDrsForPod_TaskBody struct {
	Req   *types.ConfigureStorageDrsForPod_Task         `xml:"urn:vim25 ConfigureStorageDrsForPod_Task,omitempty"`
	Res   *types.ConfigureStorageDrsForPod_TaskResponse `xml:"urn:vim25 ConfigureStorageDrsForPod_TaskResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureStorageDrsForPod_TaskBody) fault() *soap.Fault { return b.Fault }

func ConfigureStorageDrsForPod_Task(r soap.RoundTripper, req *types.ConfigureStorageDrsForPod_Task) (*types.ConfigureStorageDrsForPod_TaskResponse, Error) {
	var body = ConfigureStorageDrsForPod_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConfigureVFlashResourceEx_TaskBody struct {
	Req   *types.ConfigureVFlashResourceEx_Task         `xml:"urn:vim25 ConfigureVFlashResourceEx_Task,omitempty"`
	Res   *types.ConfigureVFlashResourceEx_TaskResponse `xml:"urn:vim25 ConfigureVFlashResourceEx_TaskResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConfigureVFlashResourceEx_TaskBody) fault() *soap.Fault { return b.Fault }

func ConfigureVFlashResourceEx_Task(r soap.RoundTripper, req *types.ConfigureVFlashResourceEx_Task) (*types.ConfigureVFlashResourceEx_TaskResponse, Error) {
	var body = ConfigureVFlashResourceEx_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ConsolidateVMDisks_TaskBody struct {
	Req   *types.ConsolidateVMDisks_Task         `xml:"urn:vim25 ConsolidateVMDisks_Task,omitempty"`
	Res   *types.ConsolidateVMDisks_TaskResponse `xml:"urn:vim25 ConsolidateVMDisks_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ConsolidateVMDisks_TaskBody) fault() *soap.Fault { return b.Fault }

func ConsolidateVMDisks_Task(r soap.RoundTripper, req *types.ConsolidateVMDisks_Task) (*types.ConsolidateVMDisks_TaskResponse, Error) {
	var body = ConsolidateVMDisks_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ContinueRetrievePropertiesExBody struct {
	Req   *types.ContinueRetrievePropertiesEx         `xml:"urn:vim25 ContinueRetrievePropertiesEx,omitempty"`
	Res   *types.ContinueRetrievePropertiesExResponse `xml:"urn:vim25 ContinueRetrievePropertiesExResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ContinueRetrievePropertiesExBody) fault() *soap.Fault { return b.Fault }

func ContinueRetrievePropertiesEx(r soap.RoundTripper, req *types.ContinueRetrievePropertiesEx) (*types.ContinueRetrievePropertiesExResponse, Error) {
	var body = ContinueRetrievePropertiesExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CopyDatastoreFile_TaskBody struct {
	Req   *types.CopyDatastoreFile_Task         `xml:"urn:vim25 CopyDatastoreFile_Task,omitempty"`
	Res   *types.CopyDatastoreFile_TaskResponse `xml:"urn:vim25 CopyDatastoreFile_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CopyDatastoreFile_TaskBody) fault() *soap.Fault { return b.Fault }

func CopyDatastoreFile_Task(r soap.RoundTripper, req *types.CopyDatastoreFile_Task) (*types.CopyDatastoreFile_TaskResponse, Error) {
	var body = CopyDatastoreFile_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CopyVirtualDisk_TaskBody struct {
	Req   *types.CopyVirtualDisk_Task         `xml:"urn:vim25 CopyVirtualDisk_Task,omitempty"`
	Res   *types.CopyVirtualDisk_TaskResponse `xml:"urn:vim25 CopyVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CopyVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func CopyVirtualDisk_Task(r soap.RoundTripper, req *types.CopyVirtualDisk_Task) (*types.CopyVirtualDisk_TaskResponse, Error) {
	var body = CopyVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateAlarmBody struct {
	Req   *types.CreateAlarm         `xml:"urn:vim25 CreateAlarm,omitempty"`
	Res   *types.CreateAlarmResponse `xml:"urn:vim25 CreateAlarmResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateAlarmBody) fault() *soap.Fault { return b.Fault }

func CreateAlarm(r soap.RoundTripper, req *types.CreateAlarm) (*types.CreateAlarmResponse, Error) {
	var body = CreateAlarmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateChildVM_TaskBody struct {
	Req   *types.CreateChildVM_Task         `xml:"urn:vim25 CreateChildVM_Task,omitempty"`
	Res   *types.CreateChildVM_TaskResponse `xml:"urn:vim25 CreateChildVM_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateChildVM_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateChildVM_Task(r soap.RoundTripper, req *types.CreateChildVM_Task) (*types.CreateChildVM_TaskResponse, Error) {
	var body = CreateChildVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateClusterBody struct {
	Req   *types.CreateCluster         `xml:"urn:vim25 CreateCluster,omitempty"`
	Res   *types.CreateClusterResponse `xml:"urn:vim25 CreateClusterResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateClusterBody) fault() *soap.Fault { return b.Fault }

func CreateCluster(r soap.RoundTripper, req *types.CreateCluster) (*types.CreateClusterResponse, Error) {
	var body = CreateClusterBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateClusterExBody struct {
	Req   *types.CreateClusterEx         `xml:"urn:vim25 CreateClusterEx,omitempty"`
	Res   *types.CreateClusterExResponse `xml:"urn:vim25 CreateClusterExResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateClusterExBody) fault() *soap.Fault { return b.Fault }

func CreateClusterEx(r soap.RoundTripper, req *types.CreateClusterEx) (*types.CreateClusterExResponse, Error) {
	var body = CreateClusterExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateCollectorForEventsBody struct {
	Req   *types.CreateCollectorForEvents         `xml:"urn:vim25 CreateCollectorForEvents,omitempty"`
	Res   *types.CreateCollectorForEventsResponse `xml:"urn:vim25 CreateCollectorForEventsResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCollectorForEventsBody) fault() *soap.Fault { return b.Fault }

func CreateCollectorForEvents(r soap.RoundTripper, req *types.CreateCollectorForEvents) (*types.CreateCollectorForEventsResponse, Error) {
	var body = CreateCollectorForEventsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateCollectorForTasksBody struct {
	Req   *types.CreateCollectorForTasks         `xml:"urn:vim25 CreateCollectorForTasks,omitempty"`
	Res   *types.CreateCollectorForTasksResponse `xml:"urn:vim25 CreateCollectorForTasksResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCollectorForTasksBody) fault() *soap.Fault { return b.Fault }

func CreateCollectorForTasks(r soap.RoundTripper, req *types.CreateCollectorForTasks) (*types.CreateCollectorForTasksResponse, Error) {
	var body = CreateCollectorForTasksBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateContainerViewBody struct {
	Req   *types.CreateContainerView         `xml:"urn:vim25 CreateContainerView,omitempty"`
	Res   *types.CreateContainerViewResponse `xml:"urn:vim25 CreateContainerViewResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateContainerViewBody) fault() *soap.Fault { return b.Fault }

func CreateContainerView(r soap.RoundTripper, req *types.CreateContainerView) (*types.CreateContainerViewResponse, Error) {
	var body = CreateContainerViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateCustomizationSpecBody struct {
	Req   *types.CreateCustomizationSpec         `xml:"urn:vim25 CreateCustomizationSpec,omitempty"`
	Res   *types.CreateCustomizationSpecResponse `xml:"urn:vim25 CreateCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func CreateCustomizationSpec(r soap.RoundTripper, req *types.CreateCustomizationSpec) (*types.CreateCustomizationSpecResponse, Error) {
	var body = CreateCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDVPortgroup_TaskBody struct {
	Req   *types.CreateDVPortgroup_Task         `xml:"urn:vim25 CreateDVPortgroup_Task,omitempty"`
	Res   *types.CreateDVPortgroup_TaskResponse `xml:"urn:vim25 CreateDVPortgroup_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDVPortgroup_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateDVPortgroup_Task(r soap.RoundTripper, req *types.CreateDVPortgroup_Task) (*types.CreateDVPortgroup_TaskResponse, Error) {
	var body = CreateDVPortgroup_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDVS_TaskBody struct {
	Req   *types.CreateDVS_Task         `xml:"urn:vim25 CreateDVS_Task,omitempty"`
	Res   *types.CreateDVS_TaskResponse `xml:"urn:vim25 CreateDVS_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDVS_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateDVS_Task(r soap.RoundTripper, req *types.CreateDVS_Task) (*types.CreateDVS_TaskResponse, Error) {
	var body = CreateDVS_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDatacenterBody struct {
	Req   *types.CreateDatacenter         `xml:"urn:vim25 CreateDatacenter,omitempty"`
	Res   *types.CreateDatacenterResponse `xml:"urn:vim25 CreateDatacenterResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDatacenterBody) fault() *soap.Fault { return b.Fault }

func CreateDatacenter(r soap.RoundTripper, req *types.CreateDatacenter) (*types.CreateDatacenterResponse, Error) {
	var body = CreateDatacenterBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDefaultProfileBody struct {
	Req   *types.CreateDefaultProfile         `xml:"urn:vim25 CreateDefaultProfile,omitempty"`
	Res   *types.CreateDefaultProfileResponse `xml:"urn:vim25 CreateDefaultProfileResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDefaultProfileBody) fault() *soap.Fault { return b.Fault }

func CreateDefaultProfile(r soap.RoundTripper, req *types.CreateDefaultProfile) (*types.CreateDefaultProfileResponse, Error) {
	var body = CreateDefaultProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDescriptorBody struct {
	Req   *types.CreateDescriptor         `xml:"urn:vim25 CreateDescriptor,omitempty"`
	Res   *types.CreateDescriptorResponse `xml:"urn:vim25 CreateDescriptorResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDescriptorBody) fault() *soap.Fault { return b.Fault }

func CreateDescriptor(r soap.RoundTripper, req *types.CreateDescriptor) (*types.CreateDescriptorResponse, Error) {
	var body = CreateDescriptorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDiagnosticPartitionBody struct {
	Req   *types.CreateDiagnosticPartition         `xml:"urn:vim25 CreateDiagnosticPartition,omitempty"`
	Res   *types.CreateDiagnosticPartitionResponse `xml:"urn:vim25 CreateDiagnosticPartitionResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDiagnosticPartitionBody) fault() *soap.Fault { return b.Fault }

func CreateDiagnosticPartition(r soap.RoundTripper, req *types.CreateDiagnosticPartition) (*types.CreateDiagnosticPartitionResponse, Error) {
	var body = CreateDiagnosticPartitionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateDirectoryBody struct {
	Req   *types.CreateDirectory         `xml:"urn:vim25 CreateDirectory,omitempty"`
	Res   *types.CreateDirectoryResponse `xml:"urn:vim25 CreateDirectoryResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateDirectoryBody) fault() *soap.Fault { return b.Fault }

func CreateDirectory(r soap.RoundTripper, req *types.CreateDirectory) (*types.CreateDirectoryResponse, Error) {
	var body = CreateDirectoryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateFilterBody struct {
	Req   *types.CreateFilter         `xml:"urn:vim25 CreateFilter,omitempty"`
	Res   *types.CreateFilterResponse `xml:"urn:vim25 CreateFilterResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateFilterBody) fault() *soap.Fault { return b.Fault }

func CreateFilter(r soap.RoundTripper, req *types.CreateFilter) (*types.CreateFilterResponse, Error) {
	var body = CreateFilterBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateFolderBody struct {
	Req   *types.CreateFolder         `xml:"urn:vim25 CreateFolder,omitempty"`
	Res   *types.CreateFolderResponse `xml:"urn:vim25 CreateFolderResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateFolderBody) fault() *soap.Fault { return b.Fault }

func CreateFolder(r soap.RoundTripper, req *types.CreateFolder) (*types.CreateFolderResponse, Error) {
	var body = CreateFolderBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateGroupBody struct {
	Req   *types.CreateGroup         `xml:"urn:vim25 CreateGroup,omitempty"`
	Res   *types.CreateGroupResponse `xml:"urn:vim25 CreateGroupResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateGroupBody) fault() *soap.Fault { return b.Fault }

func CreateGroup(r soap.RoundTripper, req *types.CreateGroup) (*types.CreateGroupResponse, Error) {
	var body = CreateGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateImportSpecBody struct {
	Req   *types.CreateImportSpec         `xml:"urn:vim25 CreateImportSpec,omitempty"`
	Res   *types.CreateImportSpecResponse `xml:"urn:vim25 CreateImportSpecResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateImportSpecBody) fault() *soap.Fault { return b.Fault }

func CreateImportSpec(r soap.RoundTripper, req *types.CreateImportSpec) (*types.CreateImportSpecResponse, Error) {
	var body = CreateImportSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateInventoryViewBody struct {
	Req   *types.CreateInventoryView         `xml:"urn:vim25 CreateInventoryView,omitempty"`
	Res   *types.CreateInventoryViewResponse `xml:"urn:vim25 CreateInventoryViewResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateInventoryViewBody) fault() *soap.Fault { return b.Fault }

func CreateInventoryView(r soap.RoundTripper, req *types.CreateInventoryView) (*types.CreateInventoryViewResponse, Error) {
	var body = CreateInventoryViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateIpPoolBody struct {
	Req   *types.CreateIpPool         `xml:"urn:vim25 CreateIpPool,omitempty"`
	Res   *types.CreateIpPoolResponse `xml:"urn:vim25 CreateIpPoolResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateIpPoolBody) fault() *soap.Fault { return b.Fault }

func CreateIpPool(r soap.RoundTripper, req *types.CreateIpPool) (*types.CreateIpPoolResponse, Error) {
	var body = CreateIpPoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateListViewBody struct {
	Req   *types.CreateListView         `xml:"urn:vim25 CreateListView,omitempty"`
	Res   *types.CreateListViewResponse `xml:"urn:vim25 CreateListViewResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateListViewBody) fault() *soap.Fault { return b.Fault }

func CreateListView(r soap.RoundTripper, req *types.CreateListView) (*types.CreateListViewResponse, Error) {
	var body = CreateListViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateListViewFromViewBody struct {
	Req   *types.CreateListViewFromView         `xml:"urn:vim25 CreateListViewFromView,omitempty"`
	Res   *types.CreateListViewFromViewResponse `xml:"urn:vim25 CreateListViewFromViewResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateListViewFromViewBody) fault() *soap.Fault { return b.Fault }

func CreateListViewFromView(r soap.RoundTripper, req *types.CreateListViewFromView) (*types.CreateListViewFromViewResponse, Error) {
	var body = CreateListViewFromViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateLocalDatastoreBody struct {
	Req   *types.CreateLocalDatastore         `xml:"urn:vim25 CreateLocalDatastore,omitempty"`
	Res   *types.CreateLocalDatastoreResponse `xml:"urn:vim25 CreateLocalDatastoreResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateLocalDatastoreBody) fault() *soap.Fault { return b.Fault }

func CreateLocalDatastore(r soap.RoundTripper, req *types.CreateLocalDatastore) (*types.CreateLocalDatastoreResponse, Error) {
	var body = CreateLocalDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateNasDatastoreBody struct {
	Req   *types.CreateNasDatastore         `xml:"urn:vim25 CreateNasDatastore,omitempty"`
	Res   *types.CreateNasDatastoreResponse `xml:"urn:vim25 CreateNasDatastoreResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateNasDatastoreBody) fault() *soap.Fault { return b.Fault }

func CreateNasDatastore(r soap.RoundTripper, req *types.CreateNasDatastore) (*types.CreateNasDatastoreResponse, Error) {
	var body = CreateNasDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateObjectScheduledTaskBody struct {
	Req   *types.CreateObjectScheduledTask         `xml:"urn:vim25 CreateObjectScheduledTask,omitempty"`
	Res   *types.CreateObjectScheduledTaskResponse `xml:"urn:vim25 CreateObjectScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateObjectScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func CreateObjectScheduledTask(r soap.RoundTripper, req *types.CreateObjectScheduledTask) (*types.CreateObjectScheduledTaskResponse, Error) {
	var body = CreateObjectScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreatePerfIntervalBody struct {
	Req   *types.CreatePerfInterval         `xml:"urn:vim25 CreatePerfInterval,omitempty"`
	Res   *types.CreatePerfIntervalResponse `xml:"urn:vim25 CreatePerfIntervalResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreatePerfIntervalBody) fault() *soap.Fault { return b.Fault }

func CreatePerfInterval(r soap.RoundTripper, req *types.CreatePerfInterval) (*types.CreatePerfIntervalResponse, Error) {
	var body = CreatePerfIntervalBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateProfileBody struct {
	Req   *types.CreateProfile         `xml:"urn:vim25 CreateProfile,omitempty"`
	Res   *types.CreateProfileResponse `xml:"urn:vim25 CreateProfileResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateProfileBody) fault() *soap.Fault { return b.Fault }

func CreateProfile(r soap.RoundTripper, req *types.CreateProfile) (*types.CreateProfileResponse, Error) {
	var body = CreateProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreatePropertyCollectorBody struct {
	Req   *types.CreatePropertyCollector         `xml:"urn:vim25 CreatePropertyCollector,omitempty"`
	Res   *types.CreatePropertyCollectorResponse `xml:"urn:vim25 CreatePropertyCollectorResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreatePropertyCollectorBody) fault() *soap.Fault { return b.Fault }

func CreatePropertyCollector(r soap.RoundTripper, req *types.CreatePropertyCollector) (*types.CreatePropertyCollectorResponse, Error) {
	var body = CreatePropertyCollectorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateResourcePoolBody struct {
	Req   *types.CreateResourcePool         `xml:"urn:vim25 CreateResourcePool,omitempty"`
	Res   *types.CreateResourcePoolResponse `xml:"urn:vim25 CreateResourcePoolResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateResourcePoolBody) fault() *soap.Fault { return b.Fault }

func CreateResourcePool(r soap.RoundTripper, req *types.CreateResourcePool) (*types.CreateResourcePoolResponse, Error) {
	var body = CreateResourcePoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateScheduledTaskBody struct {
	Req   *types.CreateScheduledTask         `xml:"urn:vim25 CreateScheduledTask,omitempty"`
	Res   *types.CreateScheduledTaskResponse `xml:"urn:vim25 CreateScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func CreateScheduledTask(r soap.RoundTripper, req *types.CreateScheduledTask) (*types.CreateScheduledTaskResponse, Error) {
	var body = CreateScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateScreenshot_TaskBody struct {
	Req   *types.CreateScreenshot_Task         `xml:"urn:vim25 CreateScreenshot_Task,omitempty"`
	Res   *types.CreateScreenshot_TaskResponse `xml:"urn:vim25 CreateScreenshot_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateScreenshot_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateScreenshot_Task(r soap.RoundTripper, req *types.CreateScreenshot_Task) (*types.CreateScreenshot_TaskResponse, Error) {
	var body = CreateScreenshot_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateSecondaryVM_TaskBody struct {
	Req   *types.CreateSecondaryVM_Task         `xml:"urn:vim25 CreateSecondaryVM_Task,omitempty"`
	Res   *types.CreateSecondaryVM_TaskResponse `xml:"urn:vim25 CreateSecondaryVM_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSecondaryVM_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateSecondaryVM_Task(r soap.RoundTripper, req *types.CreateSecondaryVM_Task) (*types.CreateSecondaryVM_TaskResponse, Error) {
	var body = CreateSecondaryVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateSnapshot_TaskBody struct {
	Req   *types.CreateSnapshot_Task         `xml:"urn:vim25 CreateSnapshot_Task,omitempty"`
	Res   *types.CreateSnapshot_TaskResponse `xml:"urn:vim25 CreateSnapshot_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateSnapshot_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateSnapshot_Task(r soap.RoundTripper, req *types.CreateSnapshot_Task) (*types.CreateSnapshot_TaskResponse, Error) {
	var body = CreateSnapshot_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateStoragePodBody struct {
	Req   *types.CreateStoragePod         `xml:"urn:vim25 CreateStoragePod,omitempty"`
	Res   *types.CreateStoragePodResponse `xml:"urn:vim25 CreateStoragePodResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateStoragePodBody) fault() *soap.Fault { return b.Fault }

func CreateStoragePod(r soap.RoundTripper, req *types.CreateStoragePod) (*types.CreateStoragePodResponse, Error) {
	var body = CreateStoragePodBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateTaskBody struct {
	Req   *types.CreateTask         `xml:"urn:vim25 CreateTask,omitempty"`
	Res   *types.CreateTaskResponse `xml:"urn:vim25 CreateTaskResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTaskBody) fault() *soap.Fault { return b.Fault }

func CreateTask(r soap.RoundTripper, req *types.CreateTask) (*types.CreateTaskResponse, Error) {
	var body = CreateTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateTemporaryDirectoryInGuestBody struct {
	Req   *types.CreateTemporaryDirectoryInGuest         `xml:"urn:vim25 CreateTemporaryDirectoryInGuest,omitempty"`
	Res   *types.CreateTemporaryDirectoryInGuestResponse `xml:"urn:vim25 CreateTemporaryDirectoryInGuestResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTemporaryDirectoryInGuestBody) fault() *soap.Fault { return b.Fault }

func CreateTemporaryDirectoryInGuest(r soap.RoundTripper, req *types.CreateTemporaryDirectoryInGuest) (*types.CreateTemporaryDirectoryInGuestResponse, Error) {
	var body = CreateTemporaryDirectoryInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateTemporaryFileInGuestBody struct {
	Req   *types.CreateTemporaryFileInGuest         `xml:"urn:vim25 CreateTemporaryFileInGuest,omitempty"`
	Res   *types.CreateTemporaryFileInGuestResponse `xml:"urn:vim25 CreateTemporaryFileInGuestResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateTemporaryFileInGuestBody) fault() *soap.Fault { return b.Fault }

func CreateTemporaryFileInGuest(r soap.RoundTripper, req *types.CreateTemporaryFileInGuest) (*types.CreateTemporaryFileInGuestResponse, Error) {
	var body = CreateTemporaryFileInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateUserBody struct {
	Req   *types.CreateUser         `xml:"urn:vim25 CreateUser,omitempty"`
	Res   *types.CreateUserResponse `xml:"urn:vim25 CreateUserResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateUserBody) fault() *soap.Fault { return b.Fault }

func CreateUser(r soap.RoundTripper, req *types.CreateUser) (*types.CreateUserResponse, Error) {
	var body = CreateUserBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateVAppBody struct {
	Req   *types.CreateVApp         `xml:"urn:vim25 CreateVApp,omitempty"`
	Res   *types.CreateVAppResponse `xml:"urn:vim25 CreateVAppResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVAppBody) fault() *soap.Fault { return b.Fault }

func CreateVApp(r soap.RoundTripper, req *types.CreateVApp) (*types.CreateVAppResponse, Error) {
	var body = CreateVAppBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateVM_TaskBody struct {
	Req   *types.CreateVM_Task         `xml:"urn:vim25 CreateVM_Task,omitempty"`
	Res   *types.CreateVM_TaskResponse `xml:"urn:vim25 CreateVM_TaskResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVM_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateVM_Task(r soap.RoundTripper, req *types.CreateVM_Task) (*types.CreateVM_TaskResponse, Error) {
	var body = CreateVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateVirtualDisk_TaskBody struct {
	Req   *types.CreateVirtualDisk_Task         `xml:"urn:vim25 CreateVirtualDisk_Task,omitempty"`
	Res   *types.CreateVirtualDisk_TaskResponse `xml:"urn:vim25 CreateVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func CreateVirtualDisk_Task(r soap.RoundTripper, req *types.CreateVirtualDisk_Task) (*types.CreateVirtualDisk_TaskResponse, Error) {
	var body = CreateVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CreateVmfsDatastoreBody struct {
	Req   *types.CreateVmfsDatastore         `xml:"urn:vim25 CreateVmfsDatastore,omitempty"`
	Res   *types.CreateVmfsDatastoreResponse `xml:"urn:vim25 CreateVmfsDatastoreResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateVmfsDatastoreBody) fault() *soap.Fault { return b.Fault }

func CreateVmfsDatastore(r soap.RoundTripper, req *types.CreateVmfsDatastore) (*types.CreateVmfsDatastoreResponse, Error) {
	var body = CreateVmfsDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CurrentTimeBody struct {
	Req   *types.CurrentTime         `xml:"urn:vim25 CurrentTime,omitempty"`
	Res   *types.CurrentTimeResponse `xml:"urn:vim25 CurrentTimeResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CurrentTimeBody) fault() *soap.Fault { return b.Fault }

func CurrentTime(r soap.RoundTripper, req *types.CurrentTime) (*types.CurrentTimeResponse, Error) {
	var body = CurrentTimeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CustomizationSpecItemToXmlBody struct {
	Req   *types.CustomizationSpecItemToXml         `xml:"urn:vim25 CustomizationSpecItemToXml,omitempty"`
	Res   *types.CustomizationSpecItemToXmlResponse `xml:"urn:vim25 CustomizationSpecItemToXmlResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CustomizationSpecItemToXmlBody) fault() *soap.Fault { return b.Fault }

func CustomizationSpecItemToXml(r soap.RoundTripper, req *types.CustomizationSpecItemToXml) (*types.CustomizationSpecItemToXmlResponse, Error) {
	var body = CustomizationSpecItemToXmlBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type CustomizeVM_TaskBody struct {
	Req   *types.CustomizeVM_Task         `xml:"urn:vim25 CustomizeVM_Task,omitempty"`
	Res   *types.CustomizeVM_TaskResponse `xml:"urn:vim25 CustomizeVM_TaskResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CustomizeVM_TaskBody) fault() *soap.Fault { return b.Fault }

func CustomizeVM_Task(r soap.RoundTripper, req *types.CustomizeVM_Task) (*types.CustomizeVM_TaskResponse, Error) {
	var body = CustomizeVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DVPortgroupRollback_TaskBody struct {
	Req   *types.DVPortgroupRollback_Task         `xml:"urn:vim25 DVPortgroupRollback_Task,omitempty"`
	Res   *types.DVPortgroupRollback_TaskResponse `xml:"urn:vim25 DVPortgroupRollback_TaskResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVPortgroupRollback_TaskBody) fault() *soap.Fault { return b.Fault }

func DVPortgroupRollback_Task(r soap.RoundTripper, req *types.DVPortgroupRollback_Task) (*types.DVPortgroupRollback_TaskResponse, Error) {
	var body = DVPortgroupRollback_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DVSManagerExportEntity_TaskBody struct {
	Req   *types.DVSManagerExportEntity_Task         `xml:"urn:vim25 DVSManagerExportEntity_Task,omitempty"`
	Res   *types.DVSManagerExportEntity_TaskResponse `xml:"urn:vim25 DVSManagerExportEntity_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerExportEntity_TaskBody) fault() *soap.Fault { return b.Fault }

func DVSManagerExportEntity_Task(r soap.RoundTripper, req *types.DVSManagerExportEntity_Task) (*types.DVSManagerExportEntity_TaskResponse, Error) {
	var body = DVSManagerExportEntity_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DVSManagerImportEntity_TaskBody struct {
	Req   *types.DVSManagerImportEntity_Task         `xml:"urn:vim25 DVSManagerImportEntity_Task,omitempty"`
	Res   *types.DVSManagerImportEntity_TaskResponse `xml:"urn:vim25 DVSManagerImportEntity_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerImportEntity_TaskBody) fault() *soap.Fault { return b.Fault }

func DVSManagerImportEntity_Task(r soap.RoundTripper, req *types.DVSManagerImportEntity_Task) (*types.DVSManagerImportEntity_TaskResponse, Error) {
	var body = DVSManagerImportEntity_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DVSManagerLookupDvPortGroupBody struct {
	Req   *types.DVSManagerLookupDvPortGroup         `xml:"urn:vim25 DVSManagerLookupDvPortGroup,omitempty"`
	Res   *types.DVSManagerLookupDvPortGroupResponse `xml:"urn:vim25 DVSManagerLookupDvPortGroupResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSManagerLookupDvPortGroupBody) fault() *soap.Fault { return b.Fault }

func DVSManagerLookupDvPortGroup(r soap.RoundTripper, req *types.DVSManagerLookupDvPortGroup) (*types.DVSManagerLookupDvPortGroupResponse, Error) {
	var body = DVSManagerLookupDvPortGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DVSRollback_TaskBody struct {
	Req   *types.DVSRollback_Task         `xml:"urn:vim25 DVSRollback_Task,omitempty"`
	Res   *types.DVSRollback_TaskResponse `xml:"urn:vim25 DVSRollback_TaskResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DVSRollback_TaskBody) fault() *soap.Fault { return b.Fault }

func DVSRollback_Task(r soap.RoundTripper, req *types.DVSRollback_Task) (*types.DVSRollback_TaskResponse, Error) {
	var body = DVSRollback_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DatastoreEnterMaintenanceModeBody struct {
	Req   *types.DatastoreEnterMaintenanceMode         `xml:"urn:vim25 DatastoreEnterMaintenanceMode,omitempty"`
	Res   *types.DatastoreEnterMaintenanceModeResponse `xml:"urn:vim25 DatastoreEnterMaintenanceModeResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DatastoreEnterMaintenanceModeBody) fault() *soap.Fault { return b.Fault }

func DatastoreEnterMaintenanceMode(r soap.RoundTripper, req *types.DatastoreEnterMaintenanceMode) (*types.DatastoreEnterMaintenanceModeResponse, Error) {
	var body = DatastoreEnterMaintenanceModeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DatastoreExitMaintenanceMode_TaskBody struct {
	Req   *types.DatastoreExitMaintenanceMode_Task         `xml:"urn:vim25 DatastoreExitMaintenanceMode_Task,omitempty"`
	Res   *types.DatastoreExitMaintenanceMode_TaskResponse `xml:"urn:vim25 DatastoreExitMaintenanceMode_TaskResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DatastoreExitMaintenanceMode_TaskBody) fault() *soap.Fault { return b.Fault }

func DatastoreExitMaintenanceMode_Task(r soap.RoundTripper, req *types.DatastoreExitMaintenanceMode_Task) (*types.DatastoreExitMaintenanceMode_TaskResponse, Error) {
	var body = DatastoreExitMaintenanceMode_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DecodeLicenseBody struct {
	Req   *types.DecodeLicense         `xml:"urn:vim25 DecodeLicense,omitempty"`
	Res   *types.DecodeLicenseResponse `xml:"urn:vim25 DecodeLicenseResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DecodeLicenseBody) fault() *soap.Fault { return b.Fault }

func DecodeLicense(r soap.RoundTripper, req *types.DecodeLicense) (*types.DecodeLicenseResponse, Error) {
	var body = DecodeLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DefragmentAllDisksBody struct {
	Req   *types.DefragmentAllDisks         `xml:"urn:vim25 DefragmentAllDisks,omitempty"`
	Res   *types.DefragmentAllDisksResponse `xml:"urn:vim25 DefragmentAllDisksResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DefragmentAllDisksBody) fault() *soap.Fault { return b.Fault }

func DefragmentAllDisks(r soap.RoundTripper, req *types.DefragmentAllDisks) (*types.DefragmentAllDisksResponse, Error) {
	var body = DefragmentAllDisksBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DefragmentVirtualDisk_TaskBody struct {
	Req   *types.DefragmentVirtualDisk_Task         `xml:"urn:vim25 DefragmentVirtualDisk_Task,omitempty"`
	Res   *types.DefragmentVirtualDisk_TaskResponse `xml:"urn:vim25 DefragmentVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DefragmentVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func DefragmentVirtualDisk_Task(r soap.RoundTripper, req *types.DefragmentVirtualDisk_Task) (*types.DefragmentVirtualDisk_TaskResponse, Error) {
	var body = DefragmentVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteCustomizationSpecBody struct {
	Req   *types.DeleteCustomizationSpec         `xml:"urn:vim25 DeleteCustomizationSpec,omitempty"`
	Res   *types.DeleteCustomizationSpecResponse `xml:"urn:vim25 DeleteCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func DeleteCustomizationSpec(r soap.RoundTripper, req *types.DeleteCustomizationSpec) (*types.DeleteCustomizationSpecResponse, Error) {
	var body = DeleteCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteDatastoreFile_TaskBody struct {
	Req   *types.DeleteDatastoreFile_Task         `xml:"urn:vim25 DeleteDatastoreFile_Task,omitempty"`
	Res   *types.DeleteDatastoreFile_TaskResponse `xml:"urn:vim25 DeleteDatastoreFile_TaskResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDatastoreFile_TaskBody) fault() *soap.Fault { return b.Fault }

func DeleteDatastoreFile_Task(r soap.RoundTripper, req *types.DeleteDatastoreFile_Task) (*types.DeleteDatastoreFile_TaskResponse, Error) {
	var body = DeleteDatastoreFile_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteDirectoryBody struct {
	Req   *types.DeleteDirectory         `xml:"urn:vim25 DeleteDirectory,omitempty"`
	Res   *types.DeleteDirectoryResponse `xml:"urn:vim25 DeleteDirectoryResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDirectoryBody) fault() *soap.Fault { return b.Fault }

func DeleteDirectory(r soap.RoundTripper, req *types.DeleteDirectory) (*types.DeleteDirectoryResponse, Error) {
	var body = DeleteDirectoryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteDirectoryInGuestBody struct {
	Req   *types.DeleteDirectoryInGuest         `xml:"urn:vim25 DeleteDirectoryInGuest,omitempty"`
	Res   *types.DeleteDirectoryInGuestResponse `xml:"urn:vim25 DeleteDirectoryInGuestResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteDirectoryInGuestBody) fault() *soap.Fault { return b.Fault }

func DeleteDirectoryInGuest(r soap.RoundTripper, req *types.DeleteDirectoryInGuest) (*types.DeleteDirectoryInGuestResponse, Error) {
	var body = DeleteDirectoryInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteFileBody struct {
	Req   *types.DeleteFile         `xml:"urn:vim25 DeleteFile,omitempty"`
	Res   *types.DeleteFileResponse `xml:"urn:vim25 DeleteFileResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteFileBody) fault() *soap.Fault { return b.Fault }

func DeleteFile(r soap.RoundTripper, req *types.DeleteFile) (*types.DeleteFileResponse, Error) {
	var body = DeleteFileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteFileInGuestBody struct {
	Req   *types.DeleteFileInGuest         `xml:"urn:vim25 DeleteFileInGuest,omitempty"`
	Res   *types.DeleteFileInGuestResponse `xml:"urn:vim25 DeleteFileInGuestResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteFileInGuestBody) fault() *soap.Fault { return b.Fault }

func DeleteFileInGuest(r soap.RoundTripper, req *types.DeleteFileInGuest) (*types.DeleteFileInGuestResponse, Error) {
	var body = DeleteFileInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteScsiLunStateBody struct {
	Req   *types.DeleteScsiLunState         `xml:"urn:vim25 DeleteScsiLunState,omitempty"`
	Res   *types.DeleteScsiLunStateResponse `xml:"urn:vim25 DeleteScsiLunStateResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteScsiLunStateBody) fault() *soap.Fault { return b.Fault }

func DeleteScsiLunState(r soap.RoundTripper, req *types.DeleteScsiLunState) (*types.DeleteScsiLunStateResponse, Error) {
	var body = DeleteScsiLunStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteVffsVolumeStateBody struct {
	Req   *types.DeleteVffsVolumeState         `xml:"urn:vim25 DeleteVffsVolumeState,omitempty"`
	Res   *types.DeleteVffsVolumeStateResponse `xml:"urn:vim25 DeleteVffsVolumeStateResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVffsVolumeStateBody) fault() *soap.Fault { return b.Fault }

func DeleteVffsVolumeState(r soap.RoundTripper, req *types.DeleteVffsVolumeState) (*types.DeleteVffsVolumeStateResponse, Error) {
	var body = DeleteVffsVolumeStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteVirtualDisk_TaskBody struct {
	Req   *types.DeleteVirtualDisk_Task         `xml:"urn:vim25 DeleteVirtualDisk_Task,omitempty"`
	Res   *types.DeleteVirtualDisk_TaskResponse `xml:"urn:vim25 DeleteVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func DeleteVirtualDisk_Task(r soap.RoundTripper, req *types.DeleteVirtualDisk_Task) (*types.DeleteVirtualDisk_TaskResponse, Error) {
	var body = DeleteVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeleteVmfsVolumeStateBody struct {
	Req   *types.DeleteVmfsVolumeState         `xml:"urn:vim25 DeleteVmfsVolumeState,omitempty"`
	Res   *types.DeleteVmfsVolumeStateResponse `xml:"urn:vim25 DeleteVmfsVolumeStateResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeleteVmfsVolumeStateBody) fault() *soap.Fault { return b.Fault }

func DeleteVmfsVolumeState(r soap.RoundTripper, req *types.DeleteVmfsVolumeState) (*types.DeleteVmfsVolumeStateResponse, Error) {
	var body = DeleteVmfsVolumeStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeselectVnicBody struct {
	Req   *types.DeselectVnic         `xml:"urn:vim25 DeselectVnic,omitempty"`
	Res   *types.DeselectVnicResponse `xml:"urn:vim25 DeselectVnicResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeselectVnicBody) fault() *soap.Fault { return b.Fault }

func DeselectVnic(r soap.RoundTripper, req *types.DeselectVnic) (*types.DeselectVnicResponse, Error) {
	var body = DeselectVnicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DeselectVnicForNicTypeBody struct {
	Req   *types.DeselectVnicForNicType         `xml:"urn:vim25 DeselectVnicForNicType,omitempty"`
	Res   *types.DeselectVnicForNicTypeResponse `xml:"urn:vim25 DeselectVnicForNicTypeResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DeselectVnicForNicTypeBody) fault() *soap.Fault { return b.Fault }

func DeselectVnicForNicType(r soap.RoundTripper, req *types.DeselectVnicForNicType) (*types.DeselectVnicForNicTypeResponse, Error) {
	var body = DeselectVnicForNicTypeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyChildrenBody struct {
	Req   *types.DestroyChildren         `xml:"urn:vim25 DestroyChildren,omitempty"`
	Res   *types.DestroyChildrenResponse `xml:"urn:vim25 DestroyChildrenResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyChildrenBody) fault() *soap.Fault { return b.Fault }

func DestroyChildren(r soap.RoundTripper, req *types.DestroyChildren) (*types.DestroyChildrenResponse, Error) {
	var body = DestroyChildrenBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyCollectorBody struct {
	Req   *types.DestroyCollector         `xml:"urn:vim25 DestroyCollector,omitempty"`
	Res   *types.DestroyCollectorResponse `xml:"urn:vim25 DestroyCollectorResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyCollectorBody) fault() *soap.Fault { return b.Fault }

func DestroyCollector(r soap.RoundTripper, req *types.DestroyCollector) (*types.DestroyCollectorResponse, Error) {
	var body = DestroyCollectorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyDatastoreBody struct {
	Req   *types.DestroyDatastore         `xml:"urn:vim25 DestroyDatastore,omitempty"`
	Res   *types.DestroyDatastoreResponse `xml:"urn:vim25 DestroyDatastoreResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyDatastoreBody) fault() *soap.Fault { return b.Fault }

func DestroyDatastore(r soap.RoundTripper, req *types.DestroyDatastore) (*types.DestroyDatastoreResponse, Error) {
	var body = DestroyDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyIpPoolBody struct {
	Req   *types.DestroyIpPool         `xml:"urn:vim25 DestroyIpPool,omitempty"`
	Res   *types.DestroyIpPoolResponse `xml:"urn:vim25 DestroyIpPoolResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyIpPoolBody) fault() *soap.Fault { return b.Fault }

func DestroyIpPool(r soap.RoundTripper, req *types.DestroyIpPool) (*types.DestroyIpPoolResponse, Error) {
	var body = DestroyIpPoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyNetworkBody struct {
	Req   *types.DestroyNetwork         `xml:"urn:vim25 DestroyNetwork,omitempty"`
	Res   *types.DestroyNetworkResponse `xml:"urn:vim25 DestroyNetworkResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyNetworkBody) fault() *soap.Fault { return b.Fault }

func DestroyNetwork(r soap.RoundTripper, req *types.DestroyNetwork) (*types.DestroyNetworkResponse, Error) {
	var body = DestroyNetworkBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyProfileBody struct {
	Req   *types.DestroyProfile         `xml:"urn:vim25 DestroyProfile,omitempty"`
	Res   *types.DestroyProfileResponse `xml:"urn:vim25 DestroyProfileResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyProfileBody) fault() *soap.Fault { return b.Fault }

func DestroyProfile(r soap.RoundTripper, req *types.DestroyProfile) (*types.DestroyProfileResponse, Error) {
	var body = DestroyProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyPropertyCollectorBody struct {
	Req   *types.DestroyPropertyCollector         `xml:"urn:vim25 DestroyPropertyCollector,omitempty"`
	Res   *types.DestroyPropertyCollectorResponse `xml:"urn:vim25 DestroyPropertyCollectorResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyPropertyCollectorBody) fault() *soap.Fault { return b.Fault }

func DestroyPropertyCollector(r soap.RoundTripper, req *types.DestroyPropertyCollector) (*types.DestroyPropertyCollectorResponse, Error) {
	var body = DestroyPropertyCollectorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyPropertyFilterBody struct {
	Req   *types.DestroyPropertyFilter         `xml:"urn:vim25 DestroyPropertyFilter,omitempty"`
	Res   *types.DestroyPropertyFilterResponse `xml:"urn:vim25 DestroyPropertyFilterResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyPropertyFilterBody) fault() *soap.Fault { return b.Fault }

func DestroyPropertyFilter(r soap.RoundTripper, req *types.DestroyPropertyFilter) (*types.DestroyPropertyFilterResponse, Error) {
	var body = DestroyPropertyFilterBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyVffsBody struct {
	Req   *types.DestroyVffs         `xml:"urn:vim25 DestroyVffs,omitempty"`
	Res   *types.DestroyVffsResponse `xml:"urn:vim25 DestroyVffsResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyVffsBody) fault() *soap.Fault { return b.Fault }

func DestroyVffs(r soap.RoundTripper, req *types.DestroyVffs) (*types.DestroyVffsResponse, Error) {
	var body = DestroyVffsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DestroyViewBody struct {
	Req   *types.DestroyView         `xml:"urn:vim25 DestroyView,omitempty"`
	Res   *types.DestroyViewResponse `xml:"urn:vim25 DestroyViewResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyViewBody) fault() *soap.Fault { return b.Fault }

func DestroyView(r soap.RoundTripper, req *types.DestroyView) (*types.DestroyViewResponse, Error) {
	var body = DestroyViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type Destroy_TaskBody struct {
	Req   *types.Destroy_Task         `xml:"urn:vim25 Destroy_Task,omitempty"`
	Res   *types.Destroy_TaskResponse `xml:"urn:vim25 Destroy_TaskResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Destroy_TaskBody) fault() *soap.Fault { return b.Fault }

func Destroy_Task(r soap.RoundTripper, req *types.Destroy_Task) (*types.Destroy_TaskResponse, Error) {
	var body = Destroy_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DetachScsiLunBody struct {
	Req   *types.DetachScsiLun         `xml:"urn:vim25 DetachScsiLun,omitempty"`
	Res   *types.DetachScsiLunResponse `xml:"urn:vim25 DetachScsiLunResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DetachScsiLunBody) fault() *soap.Fault { return b.Fault }

func DetachScsiLun(r soap.RoundTripper, req *types.DetachScsiLun) (*types.DetachScsiLunResponse, Error) {
	var body = DetachScsiLunBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisableFeatureBody struct {
	Req   *types.DisableFeature         `xml:"urn:vim25 DisableFeature,omitempty"`
	Res   *types.DisableFeatureResponse `xml:"urn:vim25 DisableFeatureResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableFeatureBody) fault() *soap.Fault { return b.Fault }

func DisableFeature(r soap.RoundTripper, req *types.DisableFeature) (*types.DisableFeatureResponse, Error) {
	var body = DisableFeatureBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisableHyperThreadingBody struct {
	Req   *types.DisableHyperThreading         `xml:"urn:vim25 DisableHyperThreading,omitempty"`
	Res   *types.DisableHyperThreadingResponse `xml:"urn:vim25 DisableHyperThreadingResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableHyperThreadingBody) fault() *soap.Fault { return b.Fault }

func DisableHyperThreading(r soap.RoundTripper, req *types.DisableHyperThreading) (*types.DisableHyperThreadingResponse, Error) {
	var body = DisableHyperThreadingBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisableMultipathPathBody struct {
	Req   *types.DisableMultipathPath         `xml:"urn:vim25 DisableMultipathPath,omitempty"`
	Res   *types.DisableMultipathPathResponse `xml:"urn:vim25 DisableMultipathPathResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableMultipathPathBody) fault() *soap.Fault { return b.Fault }

func DisableMultipathPath(r soap.RoundTripper, req *types.DisableMultipathPath) (*types.DisableMultipathPathResponse, Error) {
	var body = DisableMultipathPathBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisableRulesetBody struct {
	Req   *types.DisableRuleset         `xml:"urn:vim25 DisableRuleset,omitempty"`
	Res   *types.DisableRulesetResponse `xml:"urn:vim25 DisableRulesetResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableRulesetBody) fault() *soap.Fault { return b.Fault }

func DisableRuleset(r soap.RoundTripper, req *types.DisableRuleset) (*types.DisableRulesetResponse, Error) {
	var body = DisableRulesetBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisableSecondaryVM_TaskBody struct {
	Req   *types.DisableSecondaryVM_Task         `xml:"urn:vim25 DisableSecondaryVM_Task,omitempty"`
	Res   *types.DisableSecondaryVM_TaskResponse `xml:"urn:vim25 DisableSecondaryVM_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisableSecondaryVM_TaskBody) fault() *soap.Fault { return b.Fault }

func DisableSecondaryVM_Task(r soap.RoundTripper, req *types.DisableSecondaryVM_Task) (*types.DisableSecondaryVM_TaskResponse, Error) {
	var body = DisableSecondaryVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DisconnectHost_TaskBody struct {
	Req   *types.DisconnectHost_Task         `xml:"urn:vim25 DisconnectHost_Task,omitempty"`
	Res   *types.DisconnectHost_TaskResponse `xml:"urn:vim25 DisconnectHost_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DisconnectHost_TaskBody) fault() *soap.Fault { return b.Fault }

func DisconnectHost_Task(r soap.RoundTripper, req *types.DisconnectHost_Task) (*types.DisconnectHost_TaskResponse, Error) {
	var body = DisconnectHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DiscoverFcoeHbasBody struct {
	Req   *types.DiscoverFcoeHbas         `xml:"urn:vim25 DiscoverFcoeHbas,omitempty"`
	Res   *types.DiscoverFcoeHbasResponse `xml:"urn:vim25 DiscoverFcoeHbasResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DiscoverFcoeHbasBody) fault() *soap.Fault { return b.Fault }

func DiscoverFcoeHbas(r soap.RoundTripper, req *types.DiscoverFcoeHbas) (*types.DiscoverFcoeHbasResponse, Error) {
	var body = DiscoverFcoeHbasBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DissociateProfileBody struct {
	Req   *types.DissociateProfile         `xml:"urn:vim25 DissociateProfile,omitempty"`
	Res   *types.DissociateProfileResponse `xml:"urn:vim25 DissociateProfileResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DissociateProfileBody) fault() *soap.Fault { return b.Fault }

func DissociateProfile(r soap.RoundTripper, req *types.DissociateProfile) (*types.DissociateProfileResponse, Error) {
	var body = DissociateProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DoesCustomizationSpecExistBody struct {
	Req   *types.DoesCustomizationSpecExist         `xml:"urn:vim25 DoesCustomizationSpecExist,omitempty"`
	Res   *types.DoesCustomizationSpecExistResponse `xml:"urn:vim25 DoesCustomizationSpecExistResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DoesCustomizationSpecExistBody) fault() *soap.Fault { return b.Fault }

func DoesCustomizationSpecExist(r soap.RoundTripper, req *types.DoesCustomizationSpecExist) (*types.DoesCustomizationSpecExistResponse, Error) {
	var body = DoesCustomizationSpecExistBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type DuplicateCustomizationSpecBody struct {
	Req   *types.DuplicateCustomizationSpec         `xml:"urn:vim25 DuplicateCustomizationSpec,omitempty"`
	Res   *types.DuplicateCustomizationSpecResponse `xml:"urn:vim25 DuplicateCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DuplicateCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func DuplicateCustomizationSpec(r soap.RoundTripper, req *types.DuplicateCustomizationSpec) (*types.DuplicateCustomizationSpecResponse, Error) {
	var body = DuplicateCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EagerZeroVirtualDisk_TaskBody struct {
	Req   *types.EagerZeroVirtualDisk_Task         `xml:"urn:vim25 EagerZeroVirtualDisk_Task,omitempty"`
	Res   *types.EagerZeroVirtualDisk_TaskResponse `xml:"urn:vim25 EagerZeroVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EagerZeroVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func EagerZeroVirtualDisk_Task(r soap.RoundTripper, req *types.EagerZeroVirtualDisk_Task) (*types.EagerZeroVirtualDisk_TaskResponse, Error) {
	var body = EagerZeroVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableAlarmActionsBody struct {
	Req   *types.EnableAlarmActions         `xml:"urn:vim25 EnableAlarmActions,omitempty"`
	Res   *types.EnableAlarmActionsResponse `xml:"urn:vim25 EnableAlarmActionsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableAlarmActionsBody) fault() *soap.Fault { return b.Fault }

func EnableAlarmActions(r soap.RoundTripper, req *types.EnableAlarmActions) (*types.EnableAlarmActionsResponse, Error) {
	var body = EnableAlarmActionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableFeatureBody struct {
	Req   *types.EnableFeature         `xml:"urn:vim25 EnableFeature,omitempty"`
	Res   *types.EnableFeatureResponse `xml:"urn:vim25 EnableFeatureResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableFeatureBody) fault() *soap.Fault { return b.Fault }

func EnableFeature(r soap.RoundTripper, req *types.EnableFeature) (*types.EnableFeatureResponse, Error) {
	var body = EnableFeatureBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableHyperThreadingBody struct {
	Req   *types.EnableHyperThreading         `xml:"urn:vim25 EnableHyperThreading,omitempty"`
	Res   *types.EnableHyperThreadingResponse `xml:"urn:vim25 EnableHyperThreadingResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableHyperThreadingBody) fault() *soap.Fault { return b.Fault }

func EnableHyperThreading(r soap.RoundTripper, req *types.EnableHyperThreading) (*types.EnableHyperThreadingResponse, Error) {
	var body = EnableHyperThreadingBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableMultipathPathBody struct {
	Req   *types.EnableMultipathPath         `xml:"urn:vim25 EnableMultipathPath,omitempty"`
	Res   *types.EnableMultipathPathResponse `xml:"urn:vim25 EnableMultipathPathResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableMultipathPathBody) fault() *soap.Fault { return b.Fault }

func EnableMultipathPath(r soap.RoundTripper, req *types.EnableMultipathPath) (*types.EnableMultipathPathResponse, Error) {
	var body = EnableMultipathPathBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableNetworkResourceManagementBody struct {
	Req   *types.EnableNetworkResourceManagement         `xml:"urn:vim25 EnableNetworkResourceManagement,omitempty"`
	Res   *types.EnableNetworkResourceManagementResponse `xml:"urn:vim25 EnableNetworkResourceManagementResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableNetworkResourceManagementBody) fault() *soap.Fault { return b.Fault }

func EnableNetworkResourceManagement(r soap.RoundTripper, req *types.EnableNetworkResourceManagement) (*types.EnableNetworkResourceManagementResponse, Error) {
	var body = EnableNetworkResourceManagementBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableRulesetBody struct {
	Req   *types.EnableRuleset         `xml:"urn:vim25 EnableRuleset,omitempty"`
	Res   *types.EnableRulesetResponse `xml:"urn:vim25 EnableRulesetResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableRulesetBody) fault() *soap.Fault { return b.Fault }

func EnableRuleset(r soap.RoundTripper, req *types.EnableRuleset) (*types.EnableRulesetResponse, Error) {
	var body = EnableRulesetBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnableSecondaryVM_TaskBody struct {
	Req   *types.EnableSecondaryVM_Task         `xml:"urn:vim25 EnableSecondaryVM_Task,omitempty"`
	Res   *types.EnableSecondaryVM_TaskResponse `xml:"urn:vim25 EnableSecondaryVM_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnableSecondaryVM_TaskBody) fault() *soap.Fault { return b.Fault }

func EnableSecondaryVM_Task(r soap.RoundTripper, req *types.EnableSecondaryVM_Task) (*types.EnableSecondaryVM_TaskResponse, Error) {
	var body = EnableSecondaryVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnterLockdownModeBody struct {
	Req   *types.EnterLockdownMode         `xml:"urn:vim25 EnterLockdownMode,omitempty"`
	Res   *types.EnterLockdownModeResponse `xml:"urn:vim25 EnterLockdownModeResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnterLockdownModeBody) fault() *soap.Fault { return b.Fault }

func EnterLockdownMode(r soap.RoundTripper, req *types.EnterLockdownMode) (*types.EnterLockdownModeResponse, Error) {
	var body = EnterLockdownModeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EnterMaintenanceMode_TaskBody struct {
	Req   *types.EnterMaintenanceMode_Task         `xml:"urn:vim25 EnterMaintenanceMode_Task,omitempty"`
	Res   *types.EnterMaintenanceMode_TaskResponse `xml:"urn:vim25 EnterMaintenanceMode_TaskResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EnterMaintenanceMode_TaskBody) fault() *soap.Fault { return b.Fault }

func EnterMaintenanceMode_Task(r soap.RoundTripper, req *types.EnterMaintenanceMode_Task) (*types.EnterMaintenanceMode_TaskResponse, Error) {
	var body = EnterMaintenanceMode_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EstimateDatabaseSizeBody struct {
	Req   *types.EstimateDatabaseSize         `xml:"urn:vim25 EstimateDatabaseSize,omitempty"`
	Res   *types.EstimateDatabaseSizeResponse `xml:"urn:vim25 EstimateDatabaseSizeResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EstimateDatabaseSizeBody) fault() *soap.Fault { return b.Fault }

func EstimateDatabaseSize(r soap.RoundTripper, req *types.EstimateDatabaseSize) (*types.EstimateDatabaseSizeResponse, Error) {
	var body = EstimateDatabaseSizeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EstimateStorageForConsolidateSnapshots_TaskBody struct {
	Req   *types.EstimateStorageForConsolidateSnapshots_Task         `xml:"urn:vim25 EstimateStorageForConsolidateSnapshots_Task,omitempty"`
	Res   *types.EstimateStorageForConsolidateSnapshots_TaskResponse `xml:"urn:vim25 EstimateStorageForConsolidateSnapshots_TaskResponse,omitempty"`
	Fault *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EstimateStorageForConsolidateSnapshots_TaskBody) fault() *soap.Fault { return b.Fault }

func EstimateStorageForConsolidateSnapshots_Task(r soap.RoundTripper, req *types.EstimateStorageForConsolidateSnapshots_Task) (*types.EstimateStorageForConsolidateSnapshots_TaskResponse, Error) {
	var body = EstimateStorageForConsolidateSnapshots_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type EsxAgentHostManagerUpdateConfigBody struct {
	Req   *types.EsxAgentHostManagerUpdateConfig         `xml:"urn:vim25 EsxAgentHostManagerUpdateConfig,omitempty"`
	Res   *types.EsxAgentHostManagerUpdateConfigResponse `xml:"urn:vim25 EsxAgentHostManagerUpdateConfigResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *EsxAgentHostManagerUpdateConfigBody) fault() *soap.Fault { return b.Fault }

func EsxAgentHostManagerUpdateConfig(r soap.RoundTripper, req *types.EsxAgentHostManagerUpdateConfig) (*types.EsxAgentHostManagerUpdateConfigResponse, Error) {
	var body = EsxAgentHostManagerUpdateConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExecuteHostProfileBody struct {
	Req   *types.ExecuteHostProfile         `xml:"urn:vim25 ExecuteHostProfile,omitempty"`
	Res   *types.ExecuteHostProfileResponse `xml:"urn:vim25 ExecuteHostProfileResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExecuteHostProfileBody) fault() *soap.Fault { return b.Fault }

func ExecuteHostProfile(r soap.RoundTripper, req *types.ExecuteHostProfile) (*types.ExecuteHostProfileResponse, Error) {
	var body = ExecuteHostProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExecuteSimpleCommandBody struct {
	Req   *types.ExecuteSimpleCommand         `xml:"urn:vim25 ExecuteSimpleCommand,omitempty"`
	Res   *types.ExecuteSimpleCommandResponse `xml:"urn:vim25 ExecuteSimpleCommandResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExecuteSimpleCommandBody) fault() *soap.Fault { return b.Fault }

func ExecuteSimpleCommand(r soap.RoundTripper, req *types.ExecuteSimpleCommand) (*types.ExecuteSimpleCommandResponse, Error) {
	var body = ExecuteSimpleCommandBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExitLockdownModeBody struct {
	Req   *types.ExitLockdownMode         `xml:"urn:vim25 ExitLockdownMode,omitempty"`
	Res   *types.ExitLockdownModeResponse `xml:"urn:vim25 ExitLockdownModeResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExitLockdownModeBody) fault() *soap.Fault { return b.Fault }

func ExitLockdownMode(r soap.RoundTripper, req *types.ExitLockdownMode) (*types.ExitLockdownModeResponse, Error) {
	var body = ExitLockdownModeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExitMaintenanceMode_TaskBody struct {
	Req   *types.ExitMaintenanceMode_Task         `xml:"urn:vim25 ExitMaintenanceMode_Task,omitempty"`
	Res   *types.ExitMaintenanceMode_TaskResponse `xml:"urn:vim25 ExitMaintenanceMode_TaskResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExitMaintenanceMode_TaskBody) fault() *soap.Fault { return b.Fault }

func ExitMaintenanceMode_Task(r soap.RoundTripper, req *types.ExitMaintenanceMode_Task) (*types.ExitMaintenanceMode_TaskResponse, Error) {
	var body = ExitMaintenanceMode_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExpandVmfsDatastoreBody struct {
	Req   *types.ExpandVmfsDatastore         `xml:"urn:vim25 ExpandVmfsDatastore,omitempty"`
	Res   *types.ExpandVmfsDatastoreResponse `xml:"urn:vim25 ExpandVmfsDatastoreResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExpandVmfsDatastoreBody) fault() *soap.Fault { return b.Fault }

func ExpandVmfsDatastore(r soap.RoundTripper, req *types.ExpandVmfsDatastore) (*types.ExpandVmfsDatastoreResponse, Error) {
	var body = ExpandVmfsDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExpandVmfsExtentBody struct {
	Req   *types.ExpandVmfsExtent         `xml:"urn:vim25 ExpandVmfsExtent,omitempty"`
	Res   *types.ExpandVmfsExtentResponse `xml:"urn:vim25 ExpandVmfsExtentResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExpandVmfsExtentBody) fault() *soap.Fault { return b.Fault }

func ExpandVmfsExtent(r soap.RoundTripper, req *types.ExpandVmfsExtent) (*types.ExpandVmfsExtentResponse, Error) {
	var body = ExpandVmfsExtentBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExportAnswerFile_TaskBody struct {
	Req   *types.ExportAnswerFile_Task         `xml:"urn:vim25 ExportAnswerFile_Task,omitempty"`
	Res   *types.ExportAnswerFile_TaskResponse `xml:"urn:vim25 ExportAnswerFile_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportAnswerFile_TaskBody) fault() *soap.Fault { return b.Fault }

func ExportAnswerFile_Task(r soap.RoundTripper, req *types.ExportAnswerFile_Task) (*types.ExportAnswerFile_TaskResponse, Error) {
	var body = ExportAnswerFile_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExportProfileBody struct {
	Req   *types.ExportProfile         `xml:"urn:vim25 ExportProfile,omitempty"`
	Res   *types.ExportProfileResponse `xml:"urn:vim25 ExportProfileResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportProfileBody) fault() *soap.Fault { return b.Fault }

func ExportProfile(r soap.RoundTripper, req *types.ExportProfile) (*types.ExportProfileResponse, Error) {
	var body = ExportProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExportSnapshotBody struct {
	Req   *types.ExportSnapshot         `xml:"urn:vim25 ExportSnapshot,omitempty"`
	Res   *types.ExportSnapshotResponse `xml:"urn:vim25 ExportSnapshotResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportSnapshotBody) fault() *soap.Fault { return b.Fault }

func ExportSnapshot(r soap.RoundTripper, req *types.ExportSnapshot) (*types.ExportSnapshotResponse, Error) {
	var body = ExportSnapshotBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExportVAppBody struct {
	Req   *types.ExportVApp         `xml:"urn:vim25 ExportVApp,omitempty"`
	Res   *types.ExportVAppResponse `xml:"urn:vim25 ExportVAppResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportVAppBody) fault() *soap.Fault { return b.Fault }

func ExportVApp(r soap.RoundTripper, req *types.ExportVApp) (*types.ExportVAppResponse, Error) {
	var body = ExportVAppBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExportVmBody struct {
	Req   *types.ExportVm         `xml:"urn:vim25 ExportVm,omitempty"`
	Res   *types.ExportVmResponse `xml:"urn:vim25 ExportVmResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExportVmBody) fault() *soap.Fault { return b.Fault }

func ExportVm(r soap.RoundTripper, req *types.ExportVm) (*types.ExportVmResponse, Error) {
	var body = ExportVmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExtendVffsBody struct {
	Req   *types.ExtendVffs         `xml:"urn:vim25 ExtendVffs,omitempty"`
	Res   *types.ExtendVffsResponse `xml:"urn:vim25 ExtendVffsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVffsBody) fault() *soap.Fault { return b.Fault }

func ExtendVffs(r soap.RoundTripper, req *types.ExtendVffs) (*types.ExtendVffsResponse, Error) {
	var body = ExtendVffsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExtendVirtualDisk_TaskBody struct {
	Req   *types.ExtendVirtualDisk_Task         `xml:"urn:vim25 ExtendVirtualDisk_Task,omitempty"`
	Res   *types.ExtendVirtualDisk_TaskResponse `xml:"urn:vim25 ExtendVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func ExtendVirtualDisk_Task(r soap.RoundTripper, req *types.ExtendVirtualDisk_Task) (*types.ExtendVirtualDisk_TaskResponse, Error) {
	var body = ExtendVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExtendVmfsDatastoreBody struct {
	Req   *types.ExtendVmfsDatastore         `xml:"urn:vim25 ExtendVmfsDatastore,omitempty"`
	Res   *types.ExtendVmfsDatastoreResponse `xml:"urn:vim25 ExtendVmfsDatastoreResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtendVmfsDatastoreBody) fault() *soap.Fault { return b.Fault }

func ExtendVmfsDatastore(r soap.RoundTripper, req *types.ExtendVmfsDatastore) (*types.ExtendVmfsDatastoreResponse, Error) {
	var body = ExtendVmfsDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ExtractOvfEnvironmentBody struct {
	Req   *types.ExtractOvfEnvironment         `xml:"urn:vim25 ExtractOvfEnvironment,omitempty"`
	Res   *types.ExtractOvfEnvironmentResponse `xml:"urn:vim25 ExtractOvfEnvironmentResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ExtractOvfEnvironmentBody) fault() *soap.Fault { return b.Fault }

func ExtractOvfEnvironment(r soap.RoundTripper, req *types.ExtractOvfEnvironment) (*types.ExtractOvfEnvironmentResponse, Error) {
	var body = ExtractOvfEnvironmentBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FetchDVPortKeysBody struct {
	Req   *types.FetchDVPortKeys         `xml:"urn:vim25 FetchDVPortKeys,omitempty"`
	Res   *types.FetchDVPortKeysResponse `xml:"urn:vim25 FetchDVPortKeysResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchDVPortKeysBody) fault() *soap.Fault { return b.Fault }

func FetchDVPortKeys(r soap.RoundTripper, req *types.FetchDVPortKeys) (*types.FetchDVPortKeysResponse, Error) {
	var body = FetchDVPortKeysBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FetchDVPortsBody struct {
	Req   *types.FetchDVPorts         `xml:"urn:vim25 FetchDVPorts,omitempty"`
	Res   *types.FetchDVPortsResponse `xml:"urn:vim25 FetchDVPortsResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchDVPortsBody) fault() *soap.Fault { return b.Fault }

func FetchDVPorts(r soap.RoundTripper, req *types.FetchDVPorts) (*types.FetchDVPortsResponse, Error) {
	var body = FetchDVPortsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindAllByDnsNameBody struct {
	Req   *types.FindAllByDnsName         `xml:"urn:vim25 FindAllByDnsName,omitempty"`
	Res   *types.FindAllByDnsNameResponse `xml:"urn:vim25 FindAllByDnsNameResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByDnsNameBody) fault() *soap.Fault { return b.Fault }

func FindAllByDnsName(r soap.RoundTripper, req *types.FindAllByDnsName) (*types.FindAllByDnsNameResponse, Error) {
	var body = FindAllByDnsNameBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindAllByIpBody struct {
	Req   *types.FindAllByIp         `xml:"urn:vim25 FindAllByIp,omitempty"`
	Res   *types.FindAllByIpResponse `xml:"urn:vim25 FindAllByIpResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByIpBody) fault() *soap.Fault { return b.Fault }

func FindAllByIp(r soap.RoundTripper, req *types.FindAllByIp) (*types.FindAllByIpResponse, Error) {
	var body = FindAllByIpBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindAllByUuidBody struct {
	Req   *types.FindAllByUuid         `xml:"urn:vim25 FindAllByUuid,omitempty"`
	Res   *types.FindAllByUuidResponse `xml:"urn:vim25 FindAllByUuidResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAllByUuidBody) fault() *soap.Fault { return b.Fault }

func FindAllByUuid(r soap.RoundTripper, req *types.FindAllByUuid) (*types.FindAllByUuidResponse, Error) {
	var body = FindAllByUuidBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindAssociatedProfileBody struct {
	Req   *types.FindAssociatedProfile         `xml:"urn:vim25 FindAssociatedProfile,omitempty"`
	Res   *types.FindAssociatedProfileResponse `xml:"urn:vim25 FindAssociatedProfileResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindAssociatedProfileBody) fault() *soap.Fault { return b.Fault }

func FindAssociatedProfile(r soap.RoundTripper, req *types.FindAssociatedProfile) (*types.FindAssociatedProfileResponse, Error) {
	var body = FindAssociatedProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindByDatastorePathBody struct {
	Req   *types.FindByDatastorePath         `xml:"urn:vim25 FindByDatastorePath,omitempty"`
	Res   *types.FindByDatastorePathResponse `xml:"urn:vim25 FindByDatastorePathResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByDatastorePathBody) fault() *soap.Fault { return b.Fault }

func FindByDatastorePath(r soap.RoundTripper, req *types.FindByDatastorePath) (*types.FindByDatastorePathResponse, Error) {
	var body = FindByDatastorePathBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindByDnsNameBody struct {
	Req   *types.FindByDnsName         `xml:"urn:vim25 FindByDnsName,omitempty"`
	Res   *types.FindByDnsNameResponse `xml:"urn:vim25 FindByDnsNameResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByDnsNameBody) fault() *soap.Fault { return b.Fault }

func FindByDnsName(r soap.RoundTripper, req *types.FindByDnsName) (*types.FindByDnsNameResponse, Error) {
	var body = FindByDnsNameBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindByInventoryPathBody struct {
	Req   *types.FindByInventoryPath         `xml:"urn:vim25 FindByInventoryPath,omitempty"`
	Res   *types.FindByInventoryPathResponse `xml:"urn:vim25 FindByInventoryPathResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByInventoryPathBody) fault() *soap.Fault { return b.Fault }

func FindByInventoryPath(r soap.RoundTripper, req *types.FindByInventoryPath) (*types.FindByInventoryPathResponse, Error) {
	var body = FindByInventoryPathBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindByIpBody struct {
	Req   *types.FindByIp         `xml:"urn:vim25 FindByIp,omitempty"`
	Res   *types.FindByIpResponse `xml:"urn:vim25 FindByIpResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByIpBody) fault() *soap.Fault { return b.Fault }

func FindByIp(r soap.RoundTripper, req *types.FindByIp) (*types.FindByIpResponse, Error) {
	var body = FindByIpBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindByUuidBody struct {
	Req   *types.FindByUuid         `xml:"urn:vim25 FindByUuid,omitempty"`
	Res   *types.FindByUuidResponse `xml:"urn:vim25 FindByUuidResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindByUuidBody) fault() *soap.Fault { return b.Fault }

func FindByUuid(r soap.RoundTripper, req *types.FindByUuid) (*types.FindByUuidResponse, Error) {
	var body = FindByUuidBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindChildBody struct {
	Req   *types.FindChild         `xml:"urn:vim25 FindChild,omitempty"`
	Res   *types.FindChildResponse `xml:"urn:vim25 FindChildResponse,omitempty"`
	Fault *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindChildBody) fault() *soap.Fault { return b.Fault }

func FindChild(r soap.RoundTripper, req *types.FindChild) (*types.FindChildResponse, Error) {
	var body = FindChildBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FindExtensionBody struct {
	Req   *types.FindExtension         `xml:"urn:vim25 FindExtension,omitempty"`
	Res   *types.FindExtensionResponse `xml:"urn:vim25 FindExtensionResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FindExtensionBody) fault() *soap.Fault { return b.Fault }

func FindExtension(r soap.RoundTripper, req *types.FindExtension) (*types.FindExtensionResponse, Error) {
	var body = FindExtensionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FormatVffsBody struct {
	Req   *types.FormatVffs         `xml:"urn:vim25 FormatVffs,omitempty"`
	Res   *types.FormatVffsResponse `xml:"urn:vim25 FormatVffsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FormatVffsBody) fault() *soap.Fault { return b.Fault }

func FormatVffs(r soap.RoundTripper, req *types.FormatVffs) (*types.FormatVffsResponse, Error) {
	var body = FormatVffsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type FormatVmfsBody struct {
	Req   *types.FormatVmfs         `xml:"urn:vim25 FormatVmfs,omitempty"`
	Res   *types.FormatVmfsResponse `xml:"urn:vim25 FormatVmfsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FormatVmfsBody) fault() *soap.Fault { return b.Fault }

func FormatVmfs(r soap.RoundTripper, req *types.FormatVmfs) (*types.FormatVmfsResponse, Error) {
	var body = FormatVmfsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GenerateConfigTaskListBody struct {
	Req   *types.GenerateConfigTaskList         `xml:"urn:vim25 GenerateConfigTaskList,omitempty"`
	Res   *types.GenerateConfigTaskListResponse `xml:"urn:vim25 GenerateConfigTaskListResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateConfigTaskListBody) fault() *soap.Fault { return b.Fault }

func GenerateConfigTaskList(r soap.RoundTripper, req *types.GenerateConfigTaskList) (*types.GenerateConfigTaskListResponse, Error) {
	var body = GenerateConfigTaskListBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GenerateHostProfileTaskList_TaskBody struct {
	Req   *types.GenerateHostProfileTaskList_Task         `xml:"urn:vim25 GenerateHostProfileTaskList_Task,omitempty"`
	Res   *types.GenerateHostProfileTaskList_TaskResponse `xml:"urn:vim25 GenerateHostProfileTaskList_TaskResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateHostProfileTaskList_TaskBody) fault() *soap.Fault { return b.Fault }

func GenerateHostProfileTaskList_Task(r soap.RoundTripper, req *types.GenerateHostProfileTaskList_Task) (*types.GenerateHostProfileTaskList_TaskResponse, Error) {
	var body = GenerateHostProfileTaskList_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GenerateLogBundles_TaskBody struct {
	Req   *types.GenerateLogBundles_Task         `xml:"urn:vim25 GenerateLogBundles_Task,omitempty"`
	Res   *types.GenerateLogBundles_TaskResponse `xml:"urn:vim25 GenerateLogBundles_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GenerateLogBundles_TaskBody) fault() *soap.Fault { return b.Fault }

func GenerateLogBundles_Task(r soap.RoundTripper, req *types.GenerateLogBundles_Task) (*types.GenerateLogBundles_TaskResponse, Error) {
	var body = GenerateLogBundles_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GetAlarmBody struct {
	Req   *types.GetAlarm         `xml:"urn:vim25 GetAlarm,omitempty"`
	Res   *types.GetAlarmResponse `xml:"urn:vim25 GetAlarmResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetAlarmBody) fault() *soap.Fault { return b.Fault }

func GetAlarm(r soap.RoundTripper, req *types.GetAlarm) (*types.GetAlarmResponse, Error) {
	var body = GetAlarmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GetAlarmStateBody struct {
	Req   *types.GetAlarmState         `xml:"urn:vim25 GetAlarmState,omitempty"`
	Res   *types.GetAlarmStateResponse `xml:"urn:vim25 GetAlarmStateResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetAlarmStateBody) fault() *soap.Fault { return b.Fault }

func GetAlarmState(r soap.RoundTripper, req *types.GetAlarmState) (*types.GetAlarmStateResponse, Error) {
	var body = GetAlarmStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GetCustomizationSpecBody struct {
	Req   *types.GetCustomizationSpec         `xml:"urn:vim25 GetCustomizationSpec,omitempty"`
	Res   *types.GetCustomizationSpecResponse `xml:"urn:vim25 GetCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func GetCustomizationSpec(r soap.RoundTripper, req *types.GetCustomizationSpec) (*types.GetCustomizationSpecResponse, Error) {
	var body = GetCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type GetPublicKeyBody struct {
	Req   *types.GetPublicKey         `xml:"urn:vim25 GetPublicKey,omitempty"`
	Res   *types.GetPublicKeyResponse `xml:"urn:vim25 GetPublicKeyResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetPublicKeyBody) fault() *soap.Fault { return b.Fault }

func GetPublicKey(r soap.RoundTripper, req *types.GetPublicKey) (*types.GetPublicKeyResponse, Error) {
	var body = GetPublicKeyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HasPrivilegeOnEntitiesBody struct {
	Req   *types.HasPrivilegeOnEntities         `xml:"urn:vim25 HasPrivilegeOnEntities,omitempty"`
	Res   *types.HasPrivilegeOnEntitiesResponse `xml:"urn:vim25 HasPrivilegeOnEntitiesResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasPrivilegeOnEntitiesBody) fault() *soap.Fault { return b.Fault }

func HasPrivilegeOnEntities(r soap.RoundTripper, req *types.HasPrivilegeOnEntities) (*types.HasPrivilegeOnEntitiesResponse, Error) {
	var body = HasPrivilegeOnEntitiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HasPrivilegeOnEntityBody struct {
	Req   *types.HasPrivilegeOnEntity         `xml:"urn:vim25 HasPrivilegeOnEntity,omitempty"`
	Res   *types.HasPrivilegeOnEntityResponse `xml:"urn:vim25 HasPrivilegeOnEntityResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HasPrivilegeOnEntityBody) fault() *soap.Fault { return b.Fault }

func HasPrivilegeOnEntity(r soap.RoundTripper, req *types.HasPrivilegeOnEntity) (*types.HasPrivilegeOnEntityResponse, Error) {
	var body = HasPrivilegeOnEntityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostConfigVFlashCacheBody struct {
	Req   *types.HostConfigVFlashCache         `xml:"urn:vim25 HostConfigVFlashCache,omitempty"`
	Res   *types.HostConfigVFlashCacheResponse `xml:"urn:vim25 HostConfigVFlashCacheResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostConfigVFlashCacheBody) fault() *soap.Fault { return b.Fault }

func HostConfigVFlashCache(r soap.RoundTripper, req *types.HostConfigVFlashCache) (*types.HostConfigVFlashCacheResponse, Error) {
	var body = HostConfigVFlashCacheBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostConfigureVFlashResourceBody struct {
	Req   *types.HostConfigureVFlashResource         `xml:"urn:vim25 HostConfigureVFlashResource,omitempty"`
	Res   *types.HostConfigureVFlashResourceResponse `xml:"urn:vim25 HostConfigureVFlashResourceResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostConfigureVFlashResourceBody) fault() *soap.Fault { return b.Fault }

func HostConfigureVFlashResource(r soap.RoundTripper, req *types.HostConfigureVFlashResource) (*types.HostConfigureVFlashResourceResponse, Error) {
	var body = HostConfigureVFlashResourceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostGetVFlashModuleDefaultConfigBody struct {
	Req   *types.HostGetVFlashModuleDefaultConfig         `xml:"urn:vim25 HostGetVFlashModuleDefaultConfig,omitempty"`
	Res   *types.HostGetVFlashModuleDefaultConfigResponse `xml:"urn:vim25 HostGetVFlashModuleDefaultConfigResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostGetVFlashModuleDefaultConfigBody) fault() *soap.Fault { return b.Fault }

func HostGetVFlashModuleDefaultConfig(r soap.RoundTripper, req *types.HostGetVFlashModuleDefaultConfig) (*types.HostGetVFlashModuleDefaultConfigResponse, Error) {
	var body = HostGetVFlashModuleDefaultConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostImageConfigGetAcceptanceBody struct {
	Req   *types.HostImageConfigGetAcceptance         `xml:"urn:vim25 HostImageConfigGetAcceptance,omitempty"`
	Res   *types.HostImageConfigGetAcceptanceResponse `xml:"urn:vim25 HostImageConfigGetAcceptanceResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostImageConfigGetAcceptanceBody) fault() *soap.Fault { return b.Fault }

func HostImageConfigGetAcceptance(r soap.RoundTripper, req *types.HostImageConfigGetAcceptance) (*types.HostImageConfigGetAcceptanceResponse, Error) {
	var body = HostImageConfigGetAcceptanceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostImageConfigGetProfileBody struct {
	Req   *types.HostImageConfigGetProfile         `xml:"urn:vim25 HostImageConfigGetProfile,omitempty"`
	Res   *types.HostImageConfigGetProfileResponse `xml:"urn:vim25 HostImageConfigGetProfileResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostImageConfigGetProfileBody) fault() *soap.Fault { return b.Fault }

func HostImageConfigGetProfile(r soap.RoundTripper, req *types.HostImageConfigGetProfile) (*types.HostImageConfigGetProfileResponse, Error) {
	var body = HostImageConfigGetProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HostRemoveVFlashResourceBody struct {
	Req   *types.HostRemoveVFlashResource         `xml:"urn:vim25 HostRemoveVFlashResource,omitempty"`
	Res   *types.HostRemoveVFlashResourceResponse `xml:"urn:vim25 HostRemoveVFlashResourceResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HostRemoveVFlashResourceBody) fault() *soap.Fault { return b.Fault }

func HostRemoveVFlashResource(r soap.RoundTripper, req *types.HostRemoveVFlashResource) (*types.HostRemoveVFlashResourceResponse, Error) {
	var body = HostRemoveVFlashResourceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HttpNfcLeaseAbortBody struct {
	Req   *types.HttpNfcLeaseAbort         `xml:"urn:vim25 HttpNfcLeaseAbort,omitempty"`
	Res   *types.HttpNfcLeaseAbortResponse `xml:"urn:vim25 HttpNfcLeaseAbortResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseAbortBody) fault() *soap.Fault { return b.Fault }

func HttpNfcLeaseAbort(r soap.RoundTripper, req *types.HttpNfcLeaseAbort) (*types.HttpNfcLeaseAbortResponse, Error) {
	var body = HttpNfcLeaseAbortBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HttpNfcLeaseCompleteBody struct {
	Req   *types.HttpNfcLeaseComplete         `xml:"urn:vim25 HttpNfcLeaseComplete,omitempty"`
	Res   *types.HttpNfcLeaseCompleteResponse `xml:"urn:vim25 HttpNfcLeaseCompleteResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseCompleteBody) fault() *soap.Fault { return b.Fault }

func HttpNfcLeaseComplete(r soap.RoundTripper, req *types.HttpNfcLeaseComplete) (*types.HttpNfcLeaseCompleteResponse, Error) {
	var body = HttpNfcLeaseCompleteBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HttpNfcLeaseGetManifestBody struct {
	Req   *types.HttpNfcLeaseGetManifest         `xml:"urn:vim25 HttpNfcLeaseGetManifest,omitempty"`
	Res   *types.HttpNfcLeaseGetManifestResponse `xml:"urn:vim25 HttpNfcLeaseGetManifestResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseGetManifestBody) fault() *soap.Fault { return b.Fault }

func HttpNfcLeaseGetManifest(r soap.RoundTripper, req *types.HttpNfcLeaseGetManifest) (*types.HttpNfcLeaseGetManifestResponse, Error) {
	var body = HttpNfcLeaseGetManifestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type HttpNfcLeaseProgressBody struct {
	Req   *types.HttpNfcLeaseProgress         `xml:"urn:vim25 HttpNfcLeaseProgress,omitempty"`
	Res   *types.HttpNfcLeaseProgressResponse `xml:"urn:vim25 HttpNfcLeaseProgressResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *HttpNfcLeaseProgressBody) fault() *soap.Fault { return b.Fault }

func HttpNfcLeaseProgress(r soap.RoundTripper, req *types.HttpNfcLeaseProgress) (*types.HttpNfcLeaseProgressResponse, Error) {
	var body = HttpNfcLeaseProgressBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ImpersonateUserBody struct {
	Req   *types.ImpersonateUser         `xml:"urn:vim25 ImpersonateUser,omitempty"`
	Res   *types.ImpersonateUserResponse `xml:"urn:vim25 ImpersonateUserResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImpersonateUserBody) fault() *soap.Fault { return b.Fault }

func ImpersonateUser(r soap.RoundTripper, req *types.ImpersonateUser) (*types.ImpersonateUserResponse, Error) {
	var body = ImpersonateUserBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ImportCertificateForCAM_TaskBody struct {
	Req   *types.ImportCertificateForCAM_Task         `xml:"urn:vim25 ImportCertificateForCAM_Task,omitempty"`
	Res   *types.ImportCertificateForCAM_TaskResponse `xml:"urn:vim25 ImportCertificateForCAM_TaskResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImportCertificateForCAM_TaskBody) fault() *soap.Fault { return b.Fault }

func ImportCertificateForCAM_Task(r soap.RoundTripper, req *types.ImportCertificateForCAM_Task) (*types.ImportCertificateForCAM_TaskResponse, Error) {
	var body = ImportCertificateForCAM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ImportVAppBody struct {
	Req   *types.ImportVApp         `xml:"urn:vim25 ImportVApp,omitempty"`
	Res   *types.ImportVAppResponse `xml:"urn:vim25 ImportVAppResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ImportVAppBody) fault() *soap.Fault { return b.Fault }

func ImportVApp(r soap.RoundTripper, req *types.ImportVApp) (*types.ImportVAppResponse, Error) {
	var body = ImportVAppBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InflateVirtualDisk_TaskBody struct {
	Req   *types.InflateVirtualDisk_Task         `xml:"urn:vim25 InflateVirtualDisk_Task,omitempty"`
	Res   *types.InflateVirtualDisk_TaskResponse `xml:"urn:vim25 InflateVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InflateVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func InflateVirtualDisk_Task(r soap.RoundTripper, req *types.InflateVirtualDisk_Task) (*types.InflateVirtualDisk_TaskResponse, Error) {
	var body = InflateVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InitializeDisks_TaskBody struct {
	Req   *types.InitializeDisks_Task         `xml:"urn:vim25 InitializeDisks_Task,omitempty"`
	Res   *types.InitializeDisks_TaskResponse `xml:"urn:vim25 InitializeDisks_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitializeDisks_TaskBody) fault() *soap.Fault { return b.Fault }

func InitializeDisks_Task(r soap.RoundTripper, req *types.InitializeDisks_Task) (*types.InitializeDisks_TaskResponse, Error) {
	var body = InitializeDisks_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InitiateFileTransferFromGuestBody struct {
	Req   *types.InitiateFileTransferFromGuest         `xml:"urn:vim25 InitiateFileTransferFromGuest,omitempty"`
	Res   *types.InitiateFileTransferFromGuestResponse `xml:"urn:vim25 InitiateFileTransferFromGuestResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateFileTransferFromGuestBody) fault() *soap.Fault { return b.Fault }

func InitiateFileTransferFromGuest(r soap.RoundTripper, req *types.InitiateFileTransferFromGuest) (*types.InitiateFileTransferFromGuestResponse, Error) {
	var body = InitiateFileTransferFromGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InitiateFileTransferToGuestBody struct {
	Req   *types.InitiateFileTransferToGuest         `xml:"urn:vim25 InitiateFileTransferToGuest,omitempty"`
	Res   *types.InitiateFileTransferToGuestResponse `xml:"urn:vim25 InitiateFileTransferToGuestResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InitiateFileTransferToGuestBody) fault() *soap.Fault { return b.Fault }

func InitiateFileTransferToGuest(r soap.RoundTripper, req *types.InitiateFileTransferToGuest) (*types.InitiateFileTransferToGuestResponse, Error) {
	var body = InitiateFileTransferToGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InstallHostPatchV2_TaskBody struct {
	Req   *types.InstallHostPatchV2_Task         `xml:"urn:vim25 InstallHostPatchV2_Task,omitempty"`
	Res   *types.InstallHostPatchV2_TaskResponse `xml:"urn:vim25 InstallHostPatchV2_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallHostPatchV2_TaskBody) fault() *soap.Fault { return b.Fault }

func InstallHostPatchV2_Task(r soap.RoundTripper, req *types.InstallHostPatchV2_Task) (*types.InstallHostPatchV2_TaskResponse, Error) {
	var body = InstallHostPatchV2_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type InstallHostPatch_TaskBody struct {
	Req   *types.InstallHostPatch_Task         `xml:"urn:vim25 InstallHostPatch_Task,omitempty"`
	Res   *types.InstallHostPatch_TaskResponse `xml:"urn:vim25 InstallHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *InstallHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func InstallHostPatch_Task(r soap.RoundTripper, req *types.InstallHostPatch_Task) (*types.InstallHostPatch_TaskResponse, Error) {
	var body = InstallHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type IsSharedGraphicsActiveBody struct {
	Req   *types.IsSharedGraphicsActive         `xml:"urn:vim25 IsSharedGraphicsActive,omitempty"`
	Res   *types.IsSharedGraphicsActiveResponse `xml:"urn:vim25 IsSharedGraphicsActiveResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *IsSharedGraphicsActiveBody) fault() *soap.Fault { return b.Fault }

func IsSharedGraphicsActive(r soap.RoundTripper, req *types.IsSharedGraphicsActive) (*types.IsSharedGraphicsActiveResponse, Error) {
	var body = IsSharedGraphicsActiveBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type JoinDomainWithCAM_TaskBody struct {
	Req   *types.JoinDomainWithCAM_Task         `xml:"urn:vim25 JoinDomainWithCAM_Task,omitempty"`
	Res   *types.JoinDomainWithCAM_TaskResponse `xml:"urn:vim25 JoinDomainWithCAM_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *JoinDomainWithCAM_TaskBody) fault() *soap.Fault { return b.Fault }

func JoinDomainWithCAM_Task(r soap.RoundTripper, req *types.JoinDomainWithCAM_Task) (*types.JoinDomainWithCAM_TaskResponse, Error) {
	var body = JoinDomainWithCAM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type JoinDomain_TaskBody struct {
	Req   *types.JoinDomain_Task         `xml:"urn:vim25 JoinDomain_Task,omitempty"`
	Res   *types.JoinDomain_TaskResponse `xml:"urn:vim25 JoinDomain_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *JoinDomain_TaskBody) fault() *soap.Fault { return b.Fault }

func JoinDomain_Task(r soap.RoundTripper, req *types.JoinDomain_Task) (*types.JoinDomain_TaskResponse, Error) {
	var body = JoinDomain_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LeaveCurrentDomain_TaskBody struct {
	Req   *types.LeaveCurrentDomain_Task         `xml:"urn:vim25 LeaveCurrentDomain_Task,omitempty"`
	Res   *types.LeaveCurrentDomain_TaskResponse `xml:"urn:vim25 LeaveCurrentDomain_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LeaveCurrentDomain_TaskBody) fault() *soap.Fault { return b.Fault }

func LeaveCurrentDomain_Task(r soap.RoundTripper, req *types.LeaveCurrentDomain_Task) (*types.LeaveCurrentDomain_TaskResponse, Error) {
	var body = LeaveCurrentDomain_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ListFilesInGuestBody struct {
	Req   *types.ListFilesInGuest         `xml:"urn:vim25 ListFilesInGuest,omitempty"`
	Res   *types.ListFilesInGuestResponse `xml:"urn:vim25 ListFilesInGuestResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListFilesInGuestBody) fault() *soap.Fault { return b.Fault }

func ListFilesInGuest(r soap.RoundTripper, req *types.ListFilesInGuest) (*types.ListFilesInGuestResponse, Error) {
	var body = ListFilesInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ListProcessesInGuestBody struct {
	Req   *types.ListProcessesInGuest         `xml:"urn:vim25 ListProcessesInGuest,omitempty"`
	Res   *types.ListProcessesInGuestResponse `xml:"urn:vim25 ListProcessesInGuestResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ListProcessesInGuestBody) fault() *soap.Fault { return b.Fault }

func ListProcessesInGuest(r soap.RoundTripper, req *types.ListProcessesInGuest) (*types.ListProcessesInGuestResponse, Error) {
	var body = ListProcessesInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LogUserEventBody struct {
	Req   *types.LogUserEvent         `xml:"urn:vim25 LogUserEvent,omitempty"`
	Res   *types.LogUserEventResponse `xml:"urn:vim25 LogUserEventResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LogUserEventBody) fault() *soap.Fault { return b.Fault }

func LogUserEvent(r soap.RoundTripper, req *types.LogUserEvent) (*types.LogUserEventResponse, Error) {
	var body = LogUserEventBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LoginBody struct {
	Req   *types.Login         `xml:"urn:vim25 Login,omitempty"`
	Res   *types.LoginResponse `xml:"urn:vim25 LoginResponse,omitempty"`
	Fault *soap.Fault          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginBody) fault() *soap.Fault { return b.Fault }

func Login(r soap.RoundTripper, req *types.Login) (*types.LoginResponse, Error) {
	var body = LoginBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LoginBySSPIBody struct {
	Req   *types.LoginBySSPI         `xml:"urn:vim25 LoginBySSPI,omitempty"`
	Res   *types.LoginBySSPIResponse `xml:"urn:vim25 LoginBySSPIResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginBySSPIBody) fault() *soap.Fault { return b.Fault }

func LoginBySSPI(r soap.RoundTripper, req *types.LoginBySSPI) (*types.LoginBySSPIResponse, Error) {
	var body = LoginBySSPIBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LoginByTokenBody struct {
	Req   *types.LoginByToken         `xml:"urn:vim25 LoginByToken,omitempty"`
	Res   *types.LoginByTokenResponse `xml:"urn:vim25 LoginByTokenResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginByTokenBody) fault() *soap.Fault { return b.Fault }

func LoginByToken(r soap.RoundTripper, req *types.LoginByToken) (*types.LoginByTokenResponse, Error) {
	var body = LoginByTokenBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LoginExtensionByCertificateBody struct {
	Req   *types.LoginExtensionByCertificate         `xml:"urn:vim25 LoginExtensionByCertificate,omitempty"`
	Res   *types.LoginExtensionByCertificateResponse `xml:"urn:vim25 LoginExtensionByCertificateResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginExtensionByCertificateBody) fault() *soap.Fault { return b.Fault }

func LoginExtensionByCertificate(r soap.RoundTripper, req *types.LoginExtensionByCertificate) (*types.LoginExtensionByCertificateResponse, Error) {
	var body = LoginExtensionByCertificateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LoginExtensionBySubjectNameBody struct {
	Req   *types.LoginExtensionBySubjectName         `xml:"urn:vim25 LoginExtensionBySubjectName,omitempty"`
	Res   *types.LoginExtensionBySubjectNameResponse `xml:"urn:vim25 LoginExtensionBySubjectNameResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LoginExtensionBySubjectNameBody) fault() *soap.Fault { return b.Fault }

func LoginExtensionBySubjectName(r soap.RoundTripper, req *types.LoginExtensionBySubjectName) (*types.LoginExtensionBySubjectNameResponse, Error) {
	var body = LoginExtensionBySubjectNameBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LogoutBody struct {
	Req   *types.Logout         `xml:"urn:vim25 Logout,omitempty"`
	Res   *types.LogoutResponse `xml:"urn:vim25 LogoutResponse,omitempty"`
	Fault *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LogoutBody) fault() *soap.Fault { return b.Fault }

func Logout(r soap.RoundTripper, req *types.Logout) (*types.LogoutResponse, Error) {
	var body = LogoutBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type LookupDvPortGroupBody struct {
	Req   *types.LookupDvPortGroup         `xml:"urn:vim25 LookupDvPortGroup,omitempty"`
	Res   *types.LookupDvPortGroupResponse `xml:"urn:vim25 LookupDvPortGroupResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *LookupDvPortGroupBody) fault() *soap.Fault { return b.Fault }

func LookupDvPortGroup(r soap.RoundTripper, req *types.LookupDvPortGroup) (*types.LookupDvPortGroupResponse, Error) {
	var body = LookupDvPortGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MakeDirectoryBody struct {
	Req   *types.MakeDirectory         `xml:"urn:vim25 MakeDirectory,omitempty"`
	Res   *types.MakeDirectoryResponse `xml:"urn:vim25 MakeDirectoryResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakeDirectoryBody) fault() *soap.Fault { return b.Fault }

func MakeDirectory(r soap.RoundTripper, req *types.MakeDirectory) (*types.MakeDirectoryResponse, Error) {
	var body = MakeDirectoryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MakeDirectoryInGuestBody struct {
	Req   *types.MakeDirectoryInGuest         `xml:"urn:vim25 MakeDirectoryInGuest,omitempty"`
	Res   *types.MakeDirectoryInGuestResponse `xml:"urn:vim25 MakeDirectoryInGuestResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakeDirectoryInGuestBody) fault() *soap.Fault { return b.Fault }

func MakeDirectoryInGuest(r soap.RoundTripper, req *types.MakeDirectoryInGuest) (*types.MakeDirectoryInGuestResponse, Error) {
	var body = MakeDirectoryInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MakePrimaryVM_TaskBody struct {
	Req   *types.MakePrimaryVM_Task         `xml:"urn:vim25 MakePrimaryVM_Task,omitempty"`
	Res   *types.MakePrimaryVM_TaskResponse `xml:"urn:vim25 MakePrimaryVM_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MakePrimaryVM_TaskBody) fault() *soap.Fault { return b.Fault }

func MakePrimaryVM_Task(r soap.RoundTripper, req *types.MakePrimaryVM_Task) (*types.MakePrimaryVM_TaskResponse, Error) {
	var body = MakePrimaryVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MarkAsTemplateBody struct {
	Req   *types.MarkAsTemplate         `xml:"urn:vim25 MarkAsTemplate,omitempty"`
	Res   *types.MarkAsTemplateResponse `xml:"urn:vim25 MarkAsTemplateResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsTemplateBody) fault() *soap.Fault { return b.Fault }

func MarkAsTemplate(r soap.RoundTripper, req *types.MarkAsTemplate) (*types.MarkAsTemplateResponse, Error) {
	var body = MarkAsTemplateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MarkAsVirtualMachineBody struct {
	Req   *types.MarkAsVirtualMachine         `xml:"urn:vim25 MarkAsVirtualMachine,omitempty"`
	Res   *types.MarkAsVirtualMachineResponse `xml:"urn:vim25 MarkAsVirtualMachineResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsVirtualMachineBody) fault() *soap.Fault { return b.Fault }

func MarkAsVirtualMachine(r soap.RoundTripper, req *types.MarkAsVirtualMachine) (*types.MarkAsVirtualMachineResponse, Error) {
	var body = MarkAsVirtualMachineBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MarkForRemovalBody struct {
	Req   *types.MarkForRemoval         `xml:"urn:vim25 MarkForRemoval,omitempty"`
	Res   *types.MarkForRemovalResponse `xml:"urn:vim25 MarkForRemovalResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkForRemovalBody) fault() *soap.Fault { return b.Fault }

func MarkForRemoval(r soap.RoundTripper, req *types.MarkForRemoval) (*types.MarkForRemovalResponse, Error) {
	var body = MarkForRemovalBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MergeDvs_TaskBody struct {
	Req   *types.MergeDvs_Task         `xml:"urn:vim25 MergeDvs_Task,omitempty"`
	Res   *types.MergeDvs_TaskResponse `xml:"urn:vim25 MergeDvs_TaskResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MergeDvs_TaskBody) fault() *soap.Fault { return b.Fault }

func MergeDvs_Task(r soap.RoundTripper, req *types.MergeDvs_Task) (*types.MergeDvs_TaskResponse, Error) {
	var body = MergeDvs_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MergePermissionsBody struct {
	Req   *types.MergePermissions         `xml:"urn:vim25 MergePermissions,omitempty"`
	Res   *types.MergePermissionsResponse `xml:"urn:vim25 MergePermissionsResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MergePermissionsBody) fault() *soap.Fault { return b.Fault }

func MergePermissions(r soap.RoundTripper, req *types.MergePermissions) (*types.MergePermissionsResponse, Error) {
	var body = MergePermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MigrateVM_TaskBody struct {
	Req   *types.MigrateVM_Task         `xml:"urn:vim25 MigrateVM_Task,omitempty"`
	Res   *types.MigrateVM_TaskResponse `xml:"urn:vim25 MigrateVM_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MigrateVM_TaskBody) fault() *soap.Fault { return b.Fault }

func MigrateVM_Task(r soap.RoundTripper, req *types.MigrateVM_Task) (*types.MigrateVM_TaskResponse, Error) {
	var body = MigrateVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ModifyListViewBody struct {
	Req   *types.ModifyListView         `xml:"urn:vim25 ModifyListView,omitempty"`
	Res   *types.ModifyListViewResponse `xml:"urn:vim25 ModifyListViewResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ModifyListViewBody) fault() *soap.Fault { return b.Fault }

func ModifyListView(r soap.RoundTripper, req *types.ModifyListView) (*types.ModifyListViewResponse, Error) {
	var body = ModifyListViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MountToolsInstallerBody struct {
	Req   *types.MountToolsInstaller         `xml:"urn:vim25 MountToolsInstaller,omitempty"`
	Res   *types.MountToolsInstallerResponse `xml:"urn:vim25 MountToolsInstallerResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountToolsInstallerBody) fault() *soap.Fault { return b.Fault }

func MountToolsInstaller(r soap.RoundTripper, req *types.MountToolsInstaller) (*types.MountToolsInstallerResponse, Error) {
	var body = MountToolsInstallerBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MountVffsVolumeBody struct {
	Req   *types.MountVffsVolume         `xml:"urn:vim25 MountVffsVolume,omitempty"`
	Res   *types.MountVffsVolumeResponse `xml:"urn:vim25 MountVffsVolumeResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountVffsVolumeBody) fault() *soap.Fault { return b.Fault }

func MountVffsVolume(r soap.RoundTripper, req *types.MountVffsVolume) (*types.MountVffsVolumeResponse, Error) {
	var body = MountVffsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MountVmfsVolumeBody struct {
	Req   *types.MountVmfsVolume         `xml:"urn:vim25 MountVmfsVolume,omitempty"`
	Res   *types.MountVmfsVolumeResponse `xml:"urn:vim25 MountVmfsVolumeResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MountVmfsVolumeBody) fault() *soap.Fault { return b.Fault }

func MountVmfsVolume(r soap.RoundTripper, req *types.MountVmfsVolume) (*types.MountVmfsVolumeResponse, Error) {
	var body = MountVmfsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveDVPort_TaskBody struct {
	Req   *types.MoveDVPort_Task         `xml:"urn:vim25 MoveDVPort_Task,omitempty"`
	Res   *types.MoveDVPort_TaskResponse `xml:"urn:vim25 MoveDVPort_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDVPort_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveDVPort_Task(r soap.RoundTripper, req *types.MoveDVPort_Task) (*types.MoveDVPort_TaskResponse, Error) {
	var body = MoveDVPort_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveDatastoreFile_TaskBody struct {
	Req   *types.MoveDatastoreFile_Task         `xml:"urn:vim25 MoveDatastoreFile_Task,omitempty"`
	Res   *types.MoveDatastoreFile_TaskResponse `xml:"urn:vim25 MoveDatastoreFile_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDatastoreFile_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveDatastoreFile_Task(r soap.RoundTripper, req *types.MoveDatastoreFile_Task) (*types.MoveDatastoreFile_TaskResponse, Error) {
	var body = MoveDatastoreFile_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveDirectoryInGuestBody struct {
	Req   *types.MoveDirectoryInGuest         `xml:"urn:vim25 MoveDirectoryInGuest,omitempty"`
	Res   *types.MoveDirectoryInGuestResponse `xml:"urn:vim25 MoveDirectoryInGuestResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveDirectoryInGuestBody) fault() *soap.Fault { return b.Fault }

func MoveDirectoryInGuest(r soap.RoundTripper, req *types.MoveDirectoryInGuest) (*types.MoveDirectoryInGuestResponse, Error) {
	var body = MoveDirectoryInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveFileInGuestBody struct {
	Req   *types.MoveFileInGuest         `xml:"urn:vim25 MoveFileInGuest,omitempty"`
	Res   *types.MoveFileInGuestResponse `xml:"urn:vim25 MoveFileInGuestResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveFileInGuestBody) fault() *soap.Fault { return b.Fault }

func MoveFileInGuest(r soap.RoundTripper, req *types.MoveFileInGuest) (*types.MoveFileInGuestResponse, Error) {
	var body = MoveFileInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveHostInto_TaskBody struct {
	Req   *types.MoveHostInto_Task         `xml:"urn:vim25 MoveHostInto_Task,omitempty"`
	Res   *types.MoveHostInto_TaskResponse `xml:"urn:vim25 MoveHostInto_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveHostInto_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveHostInto_Task(r soap.RoundTripper, req *types.MoveHostInto_Task) (*types.MoveHostInto_TaskResponse, Error) {
	var body = MoveHostInto_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveIntoFolder_TaskBody struct {
	Req   *types.MoveIntoFolder_Task         `xml:"urn:vim25 MoveIntoFolder_Task,omitempty"`
	Res   *types.MoveIntoFolder_TaskResponse `xml:"urn:vim25 MoveIntoFolder_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveIntoFolder_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveIntoFolder_Task(r soap.RoundTripper, req *types.MoveIntoFolder_Task) (*types.MoveIntoFolder_TaskResponse, Error) {
	var body = MoveIntoFolder_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveIntoResourcePoolBody struct {
	Req   *types.MoveIntoResourcePool         `xml:"urn:vim25 MoveIntoResourcePool,omitempty"`
	Res   *types.MoveIntoResourcePoolResponse `xml:"urn:vim25 MoveIntoResourcePoolResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveIntoResourcePoolBody) fault() *soap.Fault { return b.Fault }

func MoveIntoResourcePool(r soap.RoundTripper, req *types.MoveIntoResourcePool) (*types.MoveIntoResourcePoolResponse, Error) {
	var body = MoveIntoResourcePoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveInto_TaskBody struct {
	Req   *types.MoveInto_Task         `xml:"urn:vim25 MoveInto_Task,omitempty"`
	Res   *types.MoveInto_TaskResponse `xml:"urn:vim25 MoveInto_TaskResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveInto_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveInto_Task(r soap.RoundTripper, req *types.MoveInto_Task) (*types.MoveInto_TaskResponse, Error) {
	var body = MoveInto_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type MoveVirtualDisk_TaskBody struct {
	Req   *types.MoveVirtualDisk_Task         `xml:"urn:vim25 MoveVirtualDisk_Task,omitempty"`
	Res   *types.MoveVirtualDisk_TaskResponse `xml:"urn:vim25 MoveVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MoveVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func MoveVirtualDisk_Task(r soap.RoundTripper, req *types.MoveVirtualDisk_Task) (*types.MoveVirtualDisk_TaskResponse, Error) {
	var body = MoveVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type OpenInventoryViewFolderBody struct {
	Req   *types.OpenInventoryViewFolder         `xml:"urn:vim25 OpenInventoryViewFolder,omitempty"`
	Res   *types.OpenInventoryViewFolderResponse `xml:"urn:vim25 OpenInventoryViewFolderResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *OpenInventoryViewFolderBody) fault() *soap.Fault { return b.Fault }

func OpenInventoryViewFolder(r soap.RoundTripper, req *types.OpenInventoryViewFolder) (*types.OpenInventoryViewFolderResponse, Error) {
	var body = OpenInventoryViewFolderBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type OverwriteCustomizationSpecBody struct {
	Req   *types.OverwriteCustomizationSpec         `xml:"urn:vim25 OverwriteCustomizationSpec,omitempty"`
	Res   *types.OverwriteCustomizationSpecResponse `xml:"urn:vim25 OverwriteCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *OverwriteCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func OverwriteCustomizationSpec(r soap.RoundTripper, req *types.OverwriteCustomizationSpec) (*types.OverwriteCustomizationSpecResponse, Error) {
	var body = OverwriteCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ParseDescriptorBody struct {
	Req   *types.ParseDescriptor         `xml:"urn:vim25 ParseDescriptor,omitempty"`
	Res   *types.ParseDescriptorResponse `xml:"urn:vim25 ParseDescriptorResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ParseDescriptorBody) fault() *soap.Fault { return b.Fault }

func ParseDescriptor(r soap.RoundTripper, req *types.ParseDescriptor) (*types.ParseDescriptorResponse, Error) {
	var body = ParseDescriptorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PerformDvsProductSpecOperation_TaskBody struct {
	Req   *types.PerformDvsProductSpecOperation_Task         `xml:"urn:vim25 PerformDvsProductSpecOperation_Task,omitempty"`
	Res   *types.PerformDvsProductSpecOperation_TaskResponse `xml:"urn:vim25 PerformDvsProductSpecOperation_TaskResponse,omitempty"`
	Fault *soap.Fault                                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PerformDvsProductSpecOperation_TaskBody) fault() *soap.Fault { return b.Fault }

func PerformDvsProductSpecOperation_Task(r soap.RoundTripper, req *types.PerformDvsProductSpecOperation_Task) (*types.PerformDvsProductSpecOperation_TaskResponse, Error) {
	var body = PerformDvsProductSpecOperation_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PostEventBody struct {
	Req   *types.PostEvent         `xml:"urn:vim25 PostEvent,omitempty"`
	Res   *types.PostEventResponse `xml:"urn:vim25 PostEventResponse,omitempty"`
	Fault *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PostEventBody) fault() *soap.Fault { return b.Fault }

func PostEvent(r soap.RoundTripper, req *types.PostEvent) (*types.PostEventResponse, Error) {
	var body = PostEventBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerDownHostToStandBy_TaskBody struct {
	Req   *types.PowerDownHostToStandBy_Task         `xml:"urn:vim25 PowerDownHostToStandBy_Task,omitempty"`
	Res   *types.PowerDownHostToStandBy_TaskResponse `xml:"urn:vim25 PowerDownHostToStandBy_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerDownHostToStandBy_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerDownHostToStandBy_Task(r soap.RoundTripper, req *types.PowerDownHostToStandBy_Task) (*types.PowerDownHostToStandBy_TaskResponse, Error) {
	var body = PowerDownHostToStandBy_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerOffVApp_TaskBody struct {
	Req   *types.PowerOffVApp_Task         `xml:"urn:vim25 PowerOffVApp_Task,omitempty"`
	Res   *types.PowerOffVApp_TaskResponse `xml:"urn:vim25 PowerOffVApp_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOffVApp_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerOffVApp_Task(r soap.RoundTripper, req *types.PowerOffVApp_Task) (*types.PowerOffVApp_TaskResponse, Error) {
	var body = PowerOffVApp_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerOffVM_TaskBody struct {
	Req   *types.PowerOffVM_Task         `xml:"urn:vim25 PowerOffVM_Task,omitempty"`
	Res   *types.PowerOffVM_TaskResponse `xml:"urn:vim25 PowerOffVM_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOffVM_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerOffVM_Task(r soap.RoundTripper, req *types.PowerOffVM_Task) (*types.PowerOffVM_TaskResponse, Error) {
	var body = PowerOffVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerOnMultiVM_TaskBody struct {
	Req   *types.PowerOnMultiVM_Task         `xml:"urn:vim25 PowerOnMultiVM_Task,omitempty"`
	Res   *types.PowerOnMultiVM_TaskResponse `xml:"urn:vim25 PowerOnMultiVM_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnMultiVM_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerOnMultiVM_Task(r soap.RoundTripper, req *types.PowerOnMultiVM_Task) (*types.PowerOnMultiVM_TaskResponse, Error) {
	var body = PowerOnMultiVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerOnVApp_TaskBody struct {
	Req   *types.PowerOnVApp_Task         `xml:"urn:vim25 PowerOnVApp_Task,omitempty"`
	Res   *types.PowerOnVApp_TaskResponse `xml:"urn:vim25 PowerOnVApp_TaskResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnVApp_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerOnVApp_Task(r soap.RoundTripper, req *types.PowerOnVApp_Task) (*types.PowerOnVApp_TaskResponse, Error) {
	var body = PowerOnVApp_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerOnVM_TaskBody struct {
	Req   *types.PowerOnVM_Task         `xml:"urn:vim25 PowerOnVM_Task,omitempty"`
	Res   *types.PowerOnVM_TaskResponse `xml:"urn:vim25 PowerOnVM_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerOnVM_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerOnVM_Task(r soap.RoundTripper, req *types.PowerOnVM_Task) (*types.PowerOnVM_TaskResponse, Error) {
	var body = PowerOnVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PowerUpHostFromStandBy_TaskBody struct {
	Req   *types.PowerUpHostFromStandBy_Task         `xml:"urn:vim25 PowerUpHostFromStandBy_Task,omitempty"`
	Res   *types.PowerUpHostFromStandBy_TaskResponse `xml:"urn:vim25 PowerUpHostFromStandBy_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PowerUpHostFromStandBy_TaskBody) fault() *soap.Fault { return b.Fault }

func PowerUpHostFromStandBy_Task(r soap.RoundTripper, req *types.PowerUpHostFromStandBy_Task) (*types.PowerUpHostFromStandBy_TaskResponse, Error) {
	var body = PowerUpHostFromStandBy_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type PromoteDisks_TaskBody struct {
	Req   *types.PromoteDisks_Task         `xml:"urn:vim25 PromoteDisks_Task,omitempty"`
	Res   *types.PromoteDisks_TaskResponse `xml:"urn:vim25 PromoteDisks_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *PromoteDisks_TaskBody) fault() *soap.Fault { return b.Fault }

func PromoteDisks_Task(r soap.RoundTripper, req *types.PromoteDisks_Task) (*types.PromoteDisks_TaskResponse, Error) {
	var body = PromoteDisks_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAnswerFileStatusBody struct {
	Req   *types.QueryAnswerFileStatus         `xml:"urn:vim25 QueryAnswerFileStatus,omitempty"`
	Res   *types.QueryAnswerFileStatusResponse `xml:"urn:vim25 QueryAnswerFileStatusResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAnswerFileStatusBody) fault() *soap.Fault { return b.Fault }

func QueryAnswerFileStatus(r soap.RoundTripper, req *types.QueryAnswerFileStatus) (*types.QueryAnswerFileStatusResponse, Error) {
	var body = QueryAnswerFileStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAssignedLicensesBody struct {
	Req   *types.QueryAssignedLicenses         `xml:"urn:vim25 QueryAssignedLicenses,omitempty"`
	Res   *types.QueryAssignedLicensesResponse `xml:"urn:vim25 QueryAssignedLicensesResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAssignedLicensesBody) fault() *soap.Fault { return b.Fault }

func QueryAssignedLicenses(r soap.RoundTripper, req *types.QueryAssignedLicenses) (*types.QueryAssignedLicensesResponse, Error) {
	var body = QueryAssignedLicensesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailableDisksForVmfsBody struct {
	Req   *types.QueryAvailableDisksForVmfs         `xml:"urn:vim25 QueryAvailableDisksForVmfs,omitempty"`
	Res   *types.QueryAvailableDisksForVmfsResponse `xml:"urn:vim25 QueryAvailableDisksForVmfsResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableDisksForVmfsBody) fault() *soap.Fault { return b.Fault }

func QueryAvailableDisksForVmfs(r soap.RoundTripper, req *types.QueryAvailableDisksForVmfs) (*types.QueryAvailableDisksForVmfsResponse, Error) {
	var body = QueryAvailableDisksForVmfsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailableDvsSpecBody struct {
	Req   *types.QueryAvailableDvsSpec         `xml:"urn:vim25 QueryAvailableDvsSpec,omitempty"`
	Res   *types.QueryAvailableDvsSpecResponse `xml:"urn:vim25 QueryAvailableDvsSpecResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableDvsSpecBody) fault() *soap.Fault { return b.Fault }

func QueryAvailableDvsSpec(r soap.RoundTripper, req *types.QueryAvailableDvsSpec) (*types.QueryAvailableDvsSpecResponse, Error) {
	var body = QueryAvailableDvsSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailablePartitionBody struct {
	Req   *types.QueryAvailablePartition         `xml:"urn:vim25 QueryAvailablePartition,omitempty"`
	Res   *types.QueryAvailablePartitionResponse `xml:"urn:vim25 QueryAvailablePartitionResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailablePartitionBody) fault() *soap.Fault { return b.Fault }

func QueryAvailablePartition(r soap.RoundTripper, req *types.QueryAvailablePartition) (*types.QueryAvailablePartitionResponse, Error) {
	var body = QueryAvailablePartitionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailablePerfMetricBody struct {
	Req   *types.QueryAvailablePerfMetric         `xml:"urn:vim25 QueryAvailablePerfMetric,omitempty"`
	Res   *types.QueryAvailablePerfMetricResponse `xml:"urn:vim25 QueryAvailablePerfMetricResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailablePerfMetricBody) fault() *soap.Fault { return b.Fault }

func QueryAvailablePerfMetric(r soap.RoundTripper, req *types.QueryAvailablePerfMetric) (*types.QueryAvailablePerfMetricResponse, Error) {
	var body = QueryAvailablePerfMetricBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailableSsdsBody struct {
	Req   *types.QueryAvailableSsds         `xml:"urn:vim25 QueryAvailableSsds,omitempty"`
	Res   *types.QueryAvailableSsdsResponse `xml:"urn:vim25 QueryAvailableSsdsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableSsdsBody) fault() *soap.Fault { return b.Fault }

func QueryAvailableSsds(r soap.RoundTripper, req *types.QueryAvailableSsds) (*types.QueryAvailableSsdsResponse, Error) {
	var body = QueryAvailableSsdsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryAvailableTimeZonesBody struct {
	Req   *types.QueryAvailableTimeZones         `xml:"urn:vim25 QueryAvailableTimeZones,omitempty"`
	Res   *types.QueryAvailableTimeZonesResponse `xml:"urn:vim25 QueryAvailableTimeZonesResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAvailableTimeZonesBody) fault() *soap.Fault { return b.Fault }

func QueryAvailableTimeZones(r soap.RoundTripper, req *types.QueryAvailableTimeZones) (*types.QueryAvailableTimeZonesResponse, Error) {
	var body = QueryAvailableTimeZonesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryBootDevicesBody struct {
	Req   *types.QueryBootDevices         `xml:"urn:vim25 QueryBootDevices,omitempty"`
	Res   *types.QueryBootDevicesResponse `xml:"urn:vim25 QueryBootDevicesResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryBootDevicesBody) fault() *soap.Fault { return b.Fault }

func QueryBootDevices(r soap.RoundTripper, req *types.QueryBootDevices) (*types.QueryBootDevicesResponse, Error) {
	var body = QueryBootDevicesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryBoundVnicsBody struct {
	Req   *types.QueryBoundVnics         `xml:"urn:vim25 QueryBoundVnics,omitempty"`
	Res   *types.QueryBoundVnicsResponse `xml:"urn:vim25 QueryBoundVnicsResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryBoundVnicsBody) fault() *soap.Fault { return b.Fault }

func QueryBoundVnics(r soap.RoundTripper, req *types.QueryBoundVnics) (*types.QueryBoundVnicsResponse, Error) {
	var body = QueryBoundVnicsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryCandidateNicsBody struct {
	Req   *types.QueryCandidateNics         `xml:"urn:vim25 QueryCandidateNics,omitempty"`
	Res   *types.QueryCandidateNicsResponse `xml:"urn:vim25 QueryCandidateNicsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCandidateNicsBody) fault() *soap.Fault { return b.Fault }

func QueryCandidateNics(r soap.RoundTripper, req *types.QueryCandidateNics) (*types.QueryCandidateNicsResponse, Error) {
	var body = QueryCandidateNicsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryChangedDiskAreasBody struct {
	Req   *types.QueryChangedDiskAreas         `xml:"urn:vim25 QueryChangedDiskAreas,omitempty"`
	Res   *types.QueryChangedDiskAreasResponse `xml:"urn:vim25 QueryChangedDiskAreasResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryChangedDiskAreasBody) fault() *soap.Fault { return b.Fault }

func QueryChangedDiskAreas(r soap.RoundTripper, req *types.QueryChangedDiskAreas) (*types.QueryChangedDiskAreasResponse, Error) {
	var body = QueryChangedDiskAreasBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryCmmdsBody struct {
	Req   *types.QueryCmmds         `xml:"urn:vim25 QueryCmmds,omitempty"`
	Res   *types.QueryCmmdsResponse `xml:"urn:vim25 QueryCmmdsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCmmdsBody) fault() *soap.Fault { return b.Fault }

func QueryCmmds(r soap.RoundTripper, req *types.QueryCmmds) (*types.QueryCmmdsResponse, Error) {
	var body = QueryCmmdsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryCompatibleHostForExistingDvsBody struct {
	Req   *types.QueryCompatibleHostForExistingDvs         `xml:"urn:vim25 QueryCompatibleHostForExistingDvs,omitempty"`
	Res   *types.QueryCompatibleHostForExistingDvsResponse `xml:"urn:vim25 QueryCompatibleHostForExistingDvsResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCompatibleHostForExistingDvsBody) fault() *soap.Fault { return b.Fault }

func QueryCompatibleHostForExistingDvs(r soap.RoundTripper, req *types.QueryCompatibleHostForExistingDvs) (*types.QueryCompatibleHostForExistingDvsResponse, Error) {
	var body = QueryCompatibleHostForExistingDvsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryCompatibleHostForNewDvsBody struct {
	Req   *types.QueryCompatibleHostForNewDvs         `xml:"urn:vim25 QueryCompatibleHostForNewDvs,omitempty"`
	Res   *types.QueryCompatibleHostForNewDvsResponse `xml:"urn:vim25 QueryCompatibleHostForNewDvsResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryCompatibleHostForNewDvsBody) fault() *soap.Fault { return b.Fault }

func QueryCompatibleHostForNewDvs(r soap.RoundTripper, req *types.QueryCompatibleHostForNewDvs) (*types.QueryCompatibleHostForNewDvsResponse, Error) {
	var body = QueryCompatibleHostForNewDvsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryComplianceStatusBody struct {
	Req   *types.QueryComplianceStatus         `xml:"urn:vim25 QueryComplianceStatus,omitempty"`
	Res   *types.QueryComplianceStatusResponse `xml:"urn:vim25 QueryComplianceStatusResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryComplianceStatusBody) fault() *soap.Fault { return b.Fault }

func QueryComplianceStatus(r soap.RoundTripper, req *types.QueryComplianceStatus) (*types.QueryComplianceStatusResponse, Error) {
	var body = QueryComplianceStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryConfigOptionBody struct {
	Req   *types.QueryConfigOption         `xml:"urn:vim25 QueryConfigOption,omitempty"`
	Res   *types.QueryConfigOptionResponse `xml:"urn:vim25 QueryConfigOptionResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigOptionBody) fault() *soap.Fault { return b.Fault }

func QueryConfigOption(r soap.RoundTripper, req *types.QueryConfigOption) (*types.QueryConfigOptionResponse, Error) {
	var body = QueryConfigOptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryConfigOptionDescriptorBody struct {
	Req   *types.QueryConfigOptionDescriptor         `xml:"urn:vim25 QueryConfigOptionDescriptor,omitempty"`
	Res   *types.QueryConfigOptionDescriptorResponse `xml:"urn:vim25 QueryConfigOptionDescriptorResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigOptionDescriptorBody) fault() *soap.Fault { return b.Fault }

func QueryConfigOptionDescriptor(r soap.RoundTripper, req *types.QueryConfigOptionDescriptor) (*types.QueryConfigOptionDescriptorResponse, Error) {
	var body = QueryConfigOptionDescriptorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryConfigTargetBody struct {
	Req   *types.QueryConfigTarget         `xml:"urn:vim25 QueryConfigTarget,omitempty"`
	Res   *types.QueryConfigTargetResponse `xml:"urn:vim25 QueryConfigTargetResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigTargetBody) fault() *soap.Fault { return b.Fault }

func QueryConfigTarget(r soap.RoundTripper, req *types.QueryConfigTarget) (*types.QueryConfigTargetResponse, Error) {
	var body = QueryConfigTargetBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryConfiguredModuleOptionStringBody struct {
	Req   *types.QueryConfiguredModuleOptionString         `xml:"urn:vim25 QueryConfiguredModuleOptionString,omitempty"`
	Res   *types.QueryConfiguredModuleOptionStringResponse `xml:"urn:vim25 QueryConfiguredModuleOptionStringResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfiguredModuleOptionStringBody) fault() *soap.Fault { return b.Fault }

func QueryConfiguredModuleOptionString(r soap.RoundTripper, req *types.QueryConfiguredModuleOptionString) (*types.QueryConfiguredModuleOptionStringResponse, Error) {
	var body = QueryConfiguredModuleOptionStringBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryConnectionInfoBody struct {
	Req   *types.QueryConnectionInfo         `xml:"urn:vim25 QueryConnectionInfo,omitempty"`
	Res   *types.QueryConnectionInfoResponse `xml:"urn:vim25 QueryConnectionInfoResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConnectionInfoBody) fault() *soap.Fault { return b.Fault }

func QueryConnectionInfo(r soap.RoundTripper, req *types.QueryConnectionInfo) (*types.QueryConnectionInfoResponse, Error) {
	var body = QueryConnectionInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDatastorePerformanceSummaryBody struct {
	Req   *types.QueryDatastorePerformanceSummary         `xml:"urn:vim25 QueryDatastorePerformanceSummary,omitempty"`
	Res   *types.QueryDatastorePerformanceSummaryResponse `xml:"urn:vim25 QueryDatastorePerformanceSummaryResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDatastorePerformanceSummaryBody) fault() *soap.Fault { return b.Fault }

func QueryDatastorePerformanceSummary(r soap.RoundTripper, req *types.QueryDatastorePerformanceSummary) (*types.QueryDatastorePerformanceSummaryResponse, Error) {
	var body = QueryDatastorePerformanceSummaryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDateTimeBody struct {
	Req   *types.QueryDateTime         `xml:"urn:vim25 QueryDateTime,omitempty"`
	Res   *types.QueryDateTimeResponse `xml:"urn:vim25 QueryDateTimeResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDateTimeBody) fault() *soap.Fault { return b.Fault }

func QueryDateTime(r soap.RoundTripper, req *types.QueryDateTime) (*types.QueryDateTimeResponse, Error) {
	var body = QueryDateTimeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDescriptionsBody struct {
	Req   *types.QueryDescriptions         `xml:"urn:vim25 QueryDescriptions,omitempty"`
	Res   *types.QueryDescriptionsResponse `xml:"urn:vim25 QueryDescriptionsResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDescriptionsBody) fault() *soap.Fault { return b.Fault }

func QueryDescriptions(r soap.RoundTripper, req *types.QueryDescriptions) (*types.QueryDescriptionsResponse, Error) {
	var body = QueryDescriptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDisksForVsanBody struct {
	Req   *types.QueryDisksForVsan         `xml:"urn:vim25 QueryDisksForVsan,omitempty"`
	Res   *types.QueryDisksForVsanResponse `xml:"urn:vim25 QueryDisksForVsanResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDisksForVsanBody) fault() *soap.Fault { return b.Fault }

func QueryDisksForVsan(r soap.RoundTripper, req *types.QueryDisksForVsan) (*types.QueryDisksForVsanResponse, Error) {
	var body = QueryDisksForVsanBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDvsByUuidBody struct {
	Req   *types.QueryDvsByUuid         `xml:"urn:vim25 QueryDvsByUuid,omitempty"`
	Res   *types.QueryDvsByUuidResponse `xml:"urn:vim25 QueryDvsByUuidResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsByUuidBody) fault() *soap.Fault { return b.Fault }

func QueryDvsByUuid(r soap.RoundTripper, req *types.QueryDvsByUuid) (*types.QueryDvsByUuidResponse, Error) {
	var body = QueryDvsByUuidBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDvsCheckCompatibilityBody struct {
	Req   *types.QueryDvsCheckCompatibility         `xml:"urn:vim25 QueryDvsCheckCompatibility,omitempty"`
	Res   *types.QueryDvsCheckCompatibilityResponse `xml:"urn:vim25 QueryDvsCheckCompatibilityResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsCheckCompatibilityBody) fault() *soap.Fault { return b.Fault }

func QueryDvsCheckCompatibility(r soap.RoundTripper, req *types.QueryDvsCheckCompatibility) (*types.QueryDvsCheckCompatibilityResponse, Error) {
	var body = QueryDvsCheckCompatibilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDvsCompatibleHostSpecBody struct {
	Req   *types.QueryDvsCompatibleHostSpec         `xml:"urn:vim25 QueryDvsCompatibleHostSpec,omitempty"`
	Res   *types.QueryDvsCompatibleHostSpecResponse `xml:"urn:vim25 QueryDvsCompatibleHostSpecResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsCompatibleHostSpecBody) fault() *soap.Fault { return b.Fault }

func QueryDvsCompatibleHostSpec(r soap.RoundTripper, req *types.QueryDvsCompatibleHostSpec) (*types.QueryDvsCompatibleHostSpecResponse, Error) {
	var body = QueryDvsCompatibleHostSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDvsConfigTargetBody struct {
	Req   *types.QueryDvsConfigTarget         `xml:"urn:vim25 QueryDvsConfigTarget,omitempty"`
	Res   *types.QueryDvsConfigTargetResponse `xml:"urn:vim25 QueryDvsConfigTargetResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsConfigTargetBody) fault() *soap.Fault { return b.Fault }

func QueryDvsConfigTarget(r soap.RoundTripper, req *types.QueryDvsConfigTarget) (*types.QueryDvsConfigTargetResponse, Error) {
	var body = QueryDvsConfigTargetBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryDvsFeatureCapabilityBody struct {
	Req   *types.QueryDvsFeatureCapability         `xml:"urn:vim25 QueryDvsFeatureCapability,omitempty"`
	Res   *types.QueryDvsFeatureCapabilityResponse `xml:"urn:vim25 QueryDvsFeatureCapabilityResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryDvsFeatureCapabilityBody) fault() *soap.Fault { return b.Fault }

func QueryDvsFeatureCapability(r soap.RoundTripper, req *types.QueryDvsFeatureCapability) (*types.QueryDvsFeatureCapabilityResponse, Error) {
	var body = QueryDvsFeatureCapabilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryEventsBody struct {
	Req   *types.QueryEvents         `xml:"urn:vim25 QueryEvents,omitempty"`
	Res   *types.QueryEventsResponse `xml:"urn:vim25 QueryEventsResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryEventsBody) fault() *soap.Fault { return b.Fault }

func QueryEvents(r soap.RoundTripper, req *types.QueryEvents) (*types.QueryEventsResponse, Error) {
	var body = QueryEventsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryExpressionMetadataBody struct {
	Req   *types.QueryExpressionMetadata         `xml:"urn:vim25 QueryExpressionMetadata,omitempty"`
	Res   *types.QueryExpressionMetadataResponse `xml:"urn:vim25 QueryExpressionMetadataResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryExpressionMetadataBody) fault() *soap.Fault { return b.Fault }

func QueryExpressionMetadata(r soap.RoundTripper, req *types.QueryExpressionMetadata) (*types.QueryExpressionMetadataResponse, Error) {
	var body = QueryExpressionMetadataBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryExtensionIpAllocationUsageBody struct {
	Req   *types.QueryExtensionIpAllocationUsage         `xml:"urn:vim25 QueryExtensionIpAllocationUsage,omitempty"`
	Res   *types.QueryExtensionIpAllocationUsageResponse `xml:"urn:vim25 QueryExtensionIpAllocationUsageResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryExtensionIpAllocationUsageBody) fault() *soap.Fault { return b.Fault }

func QueryExtensionIpAllocationUsage(r soap.RoundTripper, req *types.QueryExtensionIpAllocationUsage) (*types.QueryExtensionIpAllocationUsageResponse, Error) {
	var body = QueryExtensionIpAllocationUsageBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryFaultToleranceCompatibilityBody struct {
	Req   *types.QueryFaultToleranceCompatibility         `xml:"urn:vim25 QueryFaultToleranceCompatibility,omitempty"`
	Res   *types.QueryFaultToleranceCompatibilityResponse `xml:"urn:vim25 QueryFaultToleranceCompatibilityResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFaultToleranceCompatibilityBody) fault() *soap.Fault { return b.Fault }

func QueryFaultToleranceCompatibility(r soap.RoundTripper, req *types.QueryFaultToleranceCompatibility) (*types.QueryFaultToleranceCompatibilityResponse, Error) {
	var body = QueryFaultToleranceCompatibilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryFirmwareConfigUploadURLBody struct {
	Req   *types.QueryFirmwareConfigUploadURL         `xml:"urn:vim25 QueryFirmwareConfigUploadURL,omitempty"`
	Res   *types.QueryFirmwareConfigUploadURLResponse `xml:"urn:vim25 QueryFirmwareConfigUploadURLResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryFirmwareConfigUploadURLBody) fault() *soap.Fault { return b.Fault }

func QueryFirmwareConfigUploadURL(r soap.RoundTripper, req *types.QueryFirmwareConfigUploadURL) (*types.QueryFirmwareConfigUploadURLResponse, Error) {
	var body = QueryFirmwareConfigUploadURLBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryHostConnectionInfoBody struct {
	Req   *types.QueryHostConnectionInfo         `xml:"urn:vim25 QueryHostConnectionInfo,omitempty"`
	Res   *types.QueryHostConnectionInfoResponse `xml:"urn:vim25 QueryHostConnectionInfoResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostConnectionInfoBody) fault() *soap.Fault { return b.Fault }

func QueryHostConnectionInfo(r soap.RoundTripper, req *types.QueryHostConnectionInfo) (*types.QueryHostConnectionInfoResponse, Error) {
	var body = QueryHostConnectionInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryHostPatch_TaskBody struct {
	Req   *types.QueryHostPatch_Task         `xml:"urn:vim25 QueryHostPatch_Task,omitempty"`
	Res   *types.QueryHostPatch_TaskResponse `xml:"urn:vim25 QueryHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func QueryHostPatch_Task(r soap.RoundTripper, req *types.QueryHostPatch_Task) (*types.QueryHostPatch_TaskResponse, Error) {
	var body = QueryHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryHostProfileMetadataBody struct {
	Req   *types.QueryHostProfileMetadata         `xml:"urn:vim25 QueryHostProfileMetadata,omitempty"`
	Res   *types.QueryHostProfileMetadataResponse `xml:"urn:vim25 QueryHostProfileMetadataResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostProfileMetadataBody) fault() *soap.Fault { return b.Fault }

func QueryHostProfileMetadata(r soap.RoundTripper, req *types.QueryHostProfileMetadata) (*types.QueryHostProfileMetadataResponse, Error) {
	var body = QueryHostProfileMetadataBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryHostStatusBody struct {
	Req   *types.QueryHostStatus         `xml:"urn:vim25 QueryHostStatus,omitempty"`
	Res   *types.QueryHostStatusResponse `xml:"urn:vim25 QueryHostStatusResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryHostStatusBody) fault() *soap.Fault { return b.Fault }

func QueryHostStatus(r soap.RoundTripper, req *types.QueryHostStatus) (*types.QueryHostStatusResponse, Error) {
	var body = QueryHostStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryIORMConfigOptionBody struct {
	Req   *types.QueryIORMConfigOption         `xml:"urn:vim25 QueryIORMConfigOption,omitempty"`
	Res   *types.QueryIORMConfigOptionResponse `xml:"urn:vim25 QueryIORMConfigOptionResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIORMConfigOptionBody) fault() *soap.Fault { return b.Fault }

func QueryIORMConfigOption(r soap.RoundTripper, req *types.QueryIORMConfigOption) (*types.QueryIORMConfigOptionResponse, Error) {
	var body = QueryIORMConfigOptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryIPAllocationsBody struct {
	Req   *types.QueryIPAllocations         `xml:"urn:vim25 QueryIPAllocations,omitempty"`
	Res   *types.QueryIPAllocationsResponse `xml:"urn:vim25 QueryIPAllocationsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIPAllocationsBody) fault() *soap.Fault { return b.Fault }

func QueryIPAllocations(r soap.RoundTripper, req *types.QueryIPAllocations) (*types.QueryIPAllocationsResponse, Error) {
	var body = QueryIPAllocationsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryIpPoolsBody struct {
	Req   *types.QueryIpPools         `xml:"urn:vim25 QueryIpPools,omitempty"`
	Res   *types.QueryIpPoolsResponse `xml:"urn:vim25 QueryIpPoolsResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIpPoolsBody) fault() *soap.Fault { return b.Fault }

func QueryIpPools(r soap.RoundTripper, req *types.QueryIpPools) (*types.QueryIpPoolsResponse, Error) {
	var body = QueryIpPoolsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryLicenseSourceAvailabilityBody struct {
	Req   *types.QueryLicenseSourceAvailability         `xml:"urn:vim25 QueryLicenseSourceAvailability,omitempty"`
	Res   *types.QueryLicenseSourceAvailabilityResponse `xml:"urn:vim25 QueryLicenseSourceAvailabilityResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLicenseSourceAvailabilityBody) fault() *soap.Fault { return b.Fault }

func QueryLicenseSourceAvailability(r soap.RoundTripper, req *types.QueryLicenseSourceAvailability) (*types.QueryLicenseSourceAvailabilityResponse, Error) {
	var body = QueryLicenseSourceAvailabilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryLicenseUsageBody struct {
	Req   *types.QueryLicenseUsage         `xml:"urn:vim25 QueryLicenseUsage,omitempty"`
	Res   *types.QueryLicenseUsageResponse `xml:"urn:vim25 QueryLicenseUsageResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryLicenseUsageBody) fault() *soap.Fault { return b.Fault }

func QueryLicenseUsage(r soap.RoundTripper, req *types.QueryLicenseUsage) (*types.QueryLicenseUsageResponse, Error) {
	var body = QueryLicenseUsageBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryManagedByBody struct {
	Req   *types.QueryManagedBy         `xml:"urn:vim25 QueryManagedBy,omitempty"`
	Res   *types.QueryManagedByResponse `xml:"urn:vim25 QueryManagedByResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryManagedByBody) fault() *soap.Fault { return b.Fault }

func QueryManagedBy(r soap.RoundTripper, req *types.QueryManagedBy) (*types.QueryManagedByResponse, Error) {
	var body = QueryManagedByBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryMemoryOverheadBody struct {
	Req   *types.QueryMemoryOverhead         `xml:"urn:vim25 QueryMemoryOverhead,omitempty"`
	Res   *types.QueryMemoryOverheadResponse `xml:"urn:vim25 QueryMemoryOverheadResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMemoryOverheadBody) fault() *soap.Fault { return b.Fault }

func QueryMemoryOverhead(r soap.RoundTripper, req *types.QueryMemoryOverhead) (*types.QueryMemoryOverheadResponse, Error) {
	var body = QueryMemoryOverheadBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryMemoryOverheadExBody struct {
	Req   *types.QueryMemoryOverheadEx         `xml:"urn:vim25 QueryMemoryOverheadEx,omitempty"`
	Res   *types.QueryMemoryOverheadExResponse `xml:"urn:vim25 QueryMemoryOverheadExResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMemoryOverheadExBody) fault() *soap.Fault { return b.Fault }

func QueryMemoryOverheadEx(r soap.RoundTripper, req *types.QueryMemoryOverheadEx) (*types.QueryMemoryOverheadExResponse, Error) {
	var body = QueryMemoryOverheadExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryMigrationDependenciesBody struct {
	Req   *types.QueryMigrationDependencies         `xml:"urn:vim25 QueryMigrationDependencies,omitempty"`
	Res   *types.QueryMigrationDependenciesResponse `xml:"urn:vim25 QueryMigrationDependenciesResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryMigrationDependenciesBody) fault() *soap.Fault { return b.Fault }

func QueryMigrationDependencies(r soap.RoundTripper, req *types.QueryMigrationDependencies) (*types.QueryMigrationDependenciesResponse, Error) {
	var body = QueryMigrationDependenciesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryModulesBody struct {
	Req   *types.QueryModules         `xml:"urn:vim25 QueryModules,omitempty"`
	Res   *types.QueryModulesResponse `xml:"urn:vim25 QueryModulesResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryModulesBody) fault() *soap.Fault { return b.Fault }

func QueryModules(r soap.RoundTripper, req *types.QueryModules) (*types.QueryModulesResponse, Error) {
	var body = QueryModulesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryNetConfigBody struct {
	Req   *types.QueryNetConfig         `xml:"urn:vim25 QueryNetConfig,omitempty"`
	Res   *types.QueryNetConfigResponse `xml:"urn:vim25 QueryNetConfigResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNetConfigBody) fault() *soap.Fault { return b.Fault }

func QueryNetConfig(r soap.RoundTripper, req *types.QueryNetConfig) (*types.QueryNetConfigResponse, Error) {
	var body = QueryNetConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryNetworkHintBody struct {
	Req   *types.QueryNetworkHint         `xml:"urn:vim25 QueryNetworkHint,omitempty"`
	Res   *types.QueryNetworkHintResponse `xml:"urn:vim25 QueryNetworkHintResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryNetworkHintBody) fault() *soap.Fault { return b.Fault }

func QueryNetworkHint(r soap.RoundTripper, req *types.QueryNetworkHint) (*types.QueryNetworkHintResponse, Error) {
	var body = QueryNetworkHintBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryObjectsOnPhysicalVsanDiskBody struct {
	Req   *types.QueryObjectsOnPhysicalVsanDisk         `xml:"urn:vim25 QueryObjectsOnPhysicalVsanDisk,omitempty"`
	Res   *types.QueryObjectsOnPhysicalVsanDiskResponse `xml:"urn:vim25 QueryObjectsOnPhysicalVsanDiskResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryObjectsOnPhysicalVsanDiskBody) fault() *soap.Fault { return b.Fault }

func QueryObjectsOnPhysicalVsanDisk(r soap.RoundTripper, req *types.QueryObjectsOnPhysicalVsanDisk) (*types.QueryObjectsOnPhysicalVsanDiskResponse, Error) {
	var body = QueryObjectsOnPhysicalVsanDiskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryOptionsBody struct {
	Req   *types.QueryOptions         `xml:"urn:vim25 QueryOptions,omitempty"`
	Res   *types.QueryOptionsResponse `xml:"urn:vim25 QueryOptionsResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryOptions(r soap.RoundTripper, req *types.QueryOptions) (*types.QueryOptionsResponse, Error) {
	var body = QueryOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPartitionCreateDescBody struct {
	Req   *types.QueryPartitionCreateDesc         `xml:"urn:vim25 QueryPartitionCreateDesc,omitempty"`
	Res   *types.QueryPartitionCreateDescResponse `xml:"urn:vim25 QueryPartitionCreateDescResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPartitionCreateDescBody) fault() *soap.Fault { return b.Fault }

func QueryPartitionCreateDesc(r soap.RoundTripper, req *types.QueryPartitionCreateDesc) (*types.QueryPartitionCreateDescResponse, Error) {
	var body = QueryPartitionCreateDescBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPartitionCreateOptionsBody struct {
	Req   *types.QueryPartitionCreateOptions         `xml:"urn:vim25 QueryPartitionCreateOptions,omitempty"`
	Res   *types.QueryPartitionCreateOptionsResponse `xml:"urn:vim25 QueryPartitionCreateOptionsResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPartitionCreateOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryPartitionCreateOptions(r soap.RoundTripper, req *types.QueryPartitionCreateOptions) (*types.QueryPartitionCreateOptionsResponse, Error) {
	var body = QueryPartitionCreateOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPathSelectionPolicyOptionsBody struct {
	Req   *types.QueryPathSelectionPolicyOptions         `xml:"urn:vim25 QueryPathSelectionPolicyOptions,omitempty"`
	Res   *types.QueryPathSelectionPolicyOptionsResponse `xml:"urn:vim25 QueryPathSelectionPolicyOptionsResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPathSelectionPolicyOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryPathSelectionPolicyOptions(r soap.RoundTripper, req *types.QueryPathSelectionPolicyOptions) (*types.QueryPathSelectionPolicyOptionsResponse, Error) {
	var body = QueryPathSelectionPolicyOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPerfBody struct {
	Req   *types.QueryPerf         `xml:"urn:vim25 QueryPerf,omitempty"`
	Res   *types.QueryPerfResponse `xml:"urn:vim25 QueryPerfResponse,omitempty"`
	Fault *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfBody) fault() *soap.Fault { return b.Fault }

func QueryPerf(r soap.RoundTripper, req *types.QueryPerf) (*types.QueryPerfResponse, Error) {
	var body = QueryPerfBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPerfCompositeBody struct {
	Req   *types.QueryPerfComposite         `xml:"urn:vim25 QueryPerfComposite,omitempty"`
	Res   *types.QueryPerfCompositeResponse `xml:"urn:vim25 QueryPerfCompositeResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCompositeBody) fault() *soap.Fault { return b.Fault }

func QueryPerfComposite(r soap.RoundTripper, req *types.QueryPerfComposite) (*types.QueryPerfCompositeResponse, Error) {
	var body = QueryPerfCompositeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPerfCounterBody struct {
	Req   *types.QueryPerfCounter         `xml:"urn:vim25 QueryPerfCounter,omitempty"`
	Res   *types.QueryPerfCounterResponse `xml:"urn:vim25 QueryPerfCounterResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCounterBody) fault() *soap.Fault { return b.Fault }

func QueryPerfCounter(r soap.RoundTripper, req *types.QueryPerfCounter) (*types.QueryPerfCounterResponse, Error) {
	var body = QueryPerfCounterBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPerfCounterByLevelBody struct {
	Req   *types.QueryPerfCounterByLevel         `xml:"urn:vim25 QueryPerfCounterByLevel,omitempty"`
	Res   *types.QueryPerfCounterByLevelResponse `xml:"urn:vim25 QueryPerfCounterByLevelResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfCounterByLevelBody) fault() *soap.Fault { return b.Fault }

func QueryPerfCounterByLevel(r soap.RoundTripper, req *types.QueryPerfCounterByLevel) (*types.QueryPerfCounterByLevelResponse, Error) {
	var body = QueryPerfCounterByLevelBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPerfProviderSummaryBody struct {
	Req   *types.QueryPerfProviderSummary         `xml:"urn:vim25 QueryPerfProviderSummary,omitempty"`
	Res   *types.QueryPerfProviderSummaryResponse `xml:"urn:vim25 QueryPerfProviderSummaryResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPerfProviderSummaryBody) fault() *soap.Fault { return b.Fault }

func QueryPerfProviderSummary(r soap.RoundTripper, req *types.QueryPerfProviderSummary) (*types.QueryPerfProviderSummaryResponse, Error) {
	var body = QueryPerfProviderSummaryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPhysicalVsanDisksBody struct {
	Req   *types.QueryPhysicalVsanDisks         `xml:"urn:vim25 QueryPhysicalVsanDisks,omitempty"`
	Res   *types.QueryPhysicalVsanDisksResponse `xml:"urn:vim25 QueryPhysicalVsanDisksResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPhysicalVsanDisksBody) fault() *soap.Fault { return b.Fault }

func QueryPhysicalVsanDisks(r soap.RoundTripper, req *types.QueryPhysicalVsanDisks) (*types.QueryPhysicalVsanDisksResponse, Error) {
	var body = QueryPhysicalVsanDisksBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPnicStatusBody struct {
	Req   *types.QueryPnicStatus         `xml:"urn:vim25 QueryPnicStatus,omitempty"`
	Res   *types.QueryPnicStatusResponse `xml:"urn:vim25 QueryPnicStatusResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPnicStatusBody) fault() *soap.Fault { return b.Fault }

func QueryPnicStatus(r soap.RoundTripper, req *types.QueryPnicStatus) (*types.QueryPnicStatusResponse, Error) {
	var body = QueryPnicStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryPolicyMetadataBody struct {
	Req   *types.QueryPolicyMetadata         `xml:"urn:vim25 QueryPolicyMetadata,omitempty"`
	Res   *types.QueryPolicyMetadataResponse `xml:"urn:vim25 QueryPolicyMetadataResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryPolicyMetadataBody) fault() *soap.Fault { return b.Fault }

func QueryPolicyMetadata(r soap.RoundTripper, req *types.QueryPolicyMetadata) (*types.QueryPolicyMetadataResponse, Error) {
	var body = QueryPolicyMetadataBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryProfileStructureBody struct {
	Req   *types.QueryProfileStructure         `xml:"urn:vim25 QueryProfileStructure,omitempty"`
	Res   *types.QueryProfileStructureResponse `xml:"urn:vim25 QueryProfileStructureResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryProfileStructureBody) fault() *soap.Fault { return b.Fault }

func QueryProfileStructure(r soap.RoundTripper, req *types.QueryProfileStructure) (*types.QueryProfileStructureResponse, Error) {
	var body = QueryProfileStructureBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryResourceConfigOptionBody struct {
	Req   *types.QueryResourceConfigOption         `xml:"urn:vim25 QueryResourceConfigOption,omitempty"`
	Res   *types.QueryResourceConfigOptionResponse `xml:"urn:vim25 QueryResourceConfigOptionResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryResourceConfigOptionBody) fault() *soap.Fault { return b.Fault }

func QueryResourceConfigOption(r soap.RoundTripper, req *types.QueryResourceConfigOption) (*types.QueryResourceConfigOptionResponse, Error) {
	var body = QueryResourceConfigOptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryServiceListBody struct {
	Req   *types.QueryServiceList         `xml:"urn:vim25 QueryServiceList,omitempty"`
	Res   *types.QueryServiceListResponse `xml:"urn:vim25 QueryServiceListResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryServiceListBody) fault() *soap.Fault { return b.Fault }

func QueryServiceList(r soap.RoundTripper, req *types.QueryServiceList) (*types.QueryServiceListResponse, Error) {
	var body = QueryServiceListBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryStorageArrayTypePolicyOptionsBody struct {
	Req   *types.QueryStorageArrayTypePolicyOptions         `xml:"urn:vim25 QueryStorageArrayTypePolicyOptions,omitempty"`
	Res   *types.QueryStorageArrayTypePolicyOptionsResponse `xml:"urn:vim25 QueryStorageArrayTypePolicyOptionsResponse,omitempty"`
	Fault *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryStorageArrayTypePolicyOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryStorageArrayTypePolicyOptions(r soap.RoundTripper, req *types.QueryStorageArrayTypePolicyOptions) (*types.QueryStorageArrayTypePolicyOptionsResponse, Error) {
	var body = QueryStorageArrayTypePolicyOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QuerySupportedFeaturesBody struct {
	Req   *types.QuerySupportedFeatures         `xml:"urn:vim25 QuerySupportedFeatures,omitempty"`
	Res   *types.QuerySupportedFeaturesResponse `xml:"urn:vim25 QuerySupportedFeaturesResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySupportedFeaturesBody) fault() *soap.Fault { return b.Fault }

func QuerySupportedFeatures(r soap.RoundTripper, req *types.QuerySupportedFeatures) (*types.QuerySupportedFeaturesResponse, Error) {
	var body = QuerySupportedFeaturesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryTargetCapabilitiesBody struct {
	Req   *types.QueryTargetCapabilities         `xml:"urn:vim25 QueryTargetCapabilities,omitempty"`
	Res   *types.QueryTargetCapabilitiesResponse `xml:"urn:vim25 QueryTargetCapabilitiesResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryTargetCapabilitiesBody) fault() *soap.Fault { return b.Fault }

func QueryTargetCapabilities(r soap.RoundTripper, req *types.QueryTargetCapabilities) (*types.QueryTargetCapabilitiesResponse, Error) {
	var body = QueryTargetCapabilitiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryTpmAttestationReportBody struct {
	Req   *types.QueryTpmAttestationReport         `xml:"urn:vim25 QueryTpmAttestationReport,omitempty"`
	Res   *types.QueryTpmAttestationReportResponse `xml:"urn:vim25 QueryTpmAttestationReportResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryTpmAttestationReportBody) fault() *soap.Fault { return b.Fault }

func QueryTpmAttestationReport(r soap.RoundTripper, req *types.QueryTpmAttestationReport) (*types.QueryTpmAttestationReportResponse, Error) {
	var body = QueryTpmAttestationReportBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryUnownedFilesBody struct {
	Req   *types.QueryUnownedFiles         `xml:"urn:vim25 QueryUnownedFiles,omitempty"`
	Res   *types.QueryUnownedFilesResponse `xml:"urn:vim25 QueryUnownedFilesResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnownedFilesBody) fault() *soap.Fault { return b.Fault }

func QueryUnownedFiles(r soap.RoundTripper, req *types.QueryUnownedFiles) (*types.QueryUnownedFilesResponse, Error) {
	var body = QueryUnownedFilesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryUnresolvedVmfsVolumeBody struct {
	Req   *types.QueryUnresolvedVmfsVolume         `xml:"urn:vim25 QueryUnresolvedVmfsVolume,omitempty"`
	Res   *types.QueryUnresolvedVmfsVolumeResponse `xml:"urn:vim25 QueryUnresolvedVmfsVolumeResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnresolvedVmfsVolumeBody) fault() *soap.Fault { return b.Fault }

func QueryUnresolvedVmfsVolume(r soap.RoundTripper, req *types.QueryUnresolvedVmfsVolume) (*types.QueryUnresolvedVmfsVolumeResponse, Error) {
	var body = QueryUnresolvedVmfsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryUnresolvedVmfsVolumesBody struct {
	Req   *types.QueryUnresolvedVmfsVolumes         `xml:"urn:vim25 QueryUnresolvedVmfsVolumes,omitempty"`
	Res   *types.QueryUnresolvedVmfsVolumesResponse `xml:"urn:vim25 QueryUnresolvedVmfsVolumesResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUnresolvedVmfsVolumesBody) fault() *soap.Fault { return b.Fault }

func QueryUnresolvedVmfsVolumes(r soap.RoundTripper, req *types.QueryUnresolvedVmfsVolumes) (*types.QueryUnresolvedVmfsVolumesResponse, Error) {
	var body = QueryUnresolvedVmfsVolumesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryUsedVlanIdInDvsBody struct {
	Req   *types.QueryUsedVlanIdInDvs         `xml:"urn:vim25 QueryUsedVlanIdInDvs,omitempty"`
	Res   *types.QueryUsedVlanIdInDvsResponse `xml:"urn:vim25 QueryUsedVlanIdInDvsResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryUsedVlanIdInDvsBody) fault() *soap.Fault { return b.Fault }

func QueryUsedVlanIdInDvs(r soap.RoundTripper, req *types.QueryUsedVlanIdInDvs) (*types.QueryUsedVlanIdInDvsResponse, Error) {
	var body = QueryUsedVlanIdInDvsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVMotionCompatibilityBody struct {
	Req   *types.QueryVMotionCompatibility         `xml:"urn:vim25 QueryVMotionCompatibility,omitempty"`
	Res   *types.QueryVMotionCompatibilityResponse `xml:"urn:vim25 QueryVMotionCompatibilityResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVMotionCompatibilityBody) fault() *soap.Fault { return b.Fault }

func QueryVMotionCompatibility(r soap.RoundTripper, req *types.QueryVMotionCompatibility) (*types.QueryVMotionCompatibilityResponse, Error) {
	var body = QueryVMotionCompatibilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVMotionCompatibilityEx_TaskBody struct {
	Req   *types.QueryVMotionCompatibilityEx_Task         `xml:"urn:vim25 QueryVMotionCompatibilityEx_Task,omitempty"`
	Res   *types.QueryVMotionCompatibilityEx_TaskResponse `xml:"urn:vim25 QueryVMotionCompatibilityEx_TaskResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVMotionCompatibilityEx_TaskBody) fault() *soap.Fault { return b.Fault }

func QueryVMotionCompatibilityEx_Task(r soap.RoundTripper, req *types.QueryVMotionCompatibilityEx_Task) (*types.QueryVMotionCompatibilityEx_TaskResponse, Error) {
	var body = QueryVMotionCompatibilityEx_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVirtualDiskFragmentationBody struct {
	Req   *types.QueryVirtualDiskFragmentation         `xml:"urn:vim25 QueryVirtualDiskFragmentation,omitempty"`
	Res   *types.QueryVirtualDiskFragmentationResponse `xml:"urn:vim25 QueryVirtualDiskFragmentationResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskFragmentationBody) fault() *soap.Fault { return b.Fault }

func QueryVirtualDiskFragmentation(r soap.RoundTripper, req *types.QueryVirtualDiskFragmentation) (*types.QueryVirtualDiskFragmentationResponse, Error) {
	var body = QueryVirtualDiskFragmentationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVirtualDiskGeometryBody struct {
	Req   *types.QueryVirtualDiskGeometry         `xml:"urn:vim25 QueryVirtualDiskGeometry,omitempty"`
	Res   *types.QueryVirtualDiskGeometryResponse `xml:"urn:vim25 QueryVirtualDiskGeometryResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskGeometryBody) fault() *soap.Fault { return b.Fault }

func QueryVirtualDiskGeometry(r soap.RoundTripper, req *types.QueryVirtualDiskGeometry) (*types.QueryVirtualDiskGeometryResponse, Error) {
	var body = QueryVirtualDiskGeometryBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVirtualDiskUuidBody struct {
	Req   *types.QueryVirtualDiskUuid         `xml:"urn:vim25 QueryVirtualDiskUuid,omitempty"`
	Res   *types.QueryVirtualDiskUuidResponse `xml:"urn:vim25 QueryVirtualDiskUuidResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVirtualDiskUuidBody) fault() *soap.Fault { return b.Fault }

func QueryVirtualDiskUuid(r soap.RoundTripper, req *types.QueryVirtualDiskUuid) (*types.QueryVirtualDiskUuidResponse, Error) {
	var body = QueryVirtualDiskUuidBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVmfsDatastoreCreateOptionsBody struct {
	Req   *types.QueryVmfsDatastoreCreateOptions         `xml:"urn:vim25 QueryVmfsDatastoreCreateOptions,omitempty"`
	Res   *types.QueryVmfsDatastoreCreateOptionsResponse `xml:"urn:vim25 QueryVmfsDatastoreCreateOptionsResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreCreateOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryVmfsDatastoreCreateOptions(r soap.RoundTripper, req *types.QueryVmfsDatastoreCreateOptions) (*types.QueryVmfsDatastoreCreateOptionsResponse, Error) {
	var body = QueryVmfsDatastoreCreateOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVmfsDatastoreExpandOptionsBody struct {
	Req   *types.QueryVmfsDatastoreExpandOptions         `xml:"urn:vim25 QueryVmfsDatastoreExpandOptions,omitempty"`
	Res   *types.QueryVmfsDatastoreExpandOptionsResponse `xml:"urn:vim25 QueryVmfsDatastoreExpandOptionsResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreExpandOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryVmfsDatastoreExpandOptions(r soap.RoundTripper, req *types.QueryVmfsDatastoreExpandOptions) (*types.QueryVmfsDatastoreExpandOptionsResponse, Error) {
	var body = QueryVmfsDatastoreExpandOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVmfsDatastoreExtendOptionsBody struct {
	Req   *types.QueryVmfsDatastoreExtendOptions         `xml:"urn:vim25 QueryVmfsDatastoreExtendOptions,omitempty"`
	Res   *types.QueryVmfsDatastoreExtendOptionsResponse `xml:"urn:vim25 QueryVmfsDatastoreExtendOptionsResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVmfsDatastoreExtendOptionsBody) fault() *soap.Fault { return b.Fault }

func QueryVmfsDatastoreExtendOptions(r soap.RoundTripper, req *types.QueryVmfsDatastoreExtendOptions) (*types.QueryVmfsDatastoreExtendOptionsResponse, Error) {
	var body = QueryVmfsDatastoreExtendOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVnicStatusBody struct {
	Req   *types.QueryVnicStatus         `xml:"urn:vim25 QueryVnicStatus,omitempty"`
	Res   *types.QueryVnicStatusResponse `xml:"urn:vim25 QueryVnicStatusResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVnicStatusBody) fault() *soap.Fault { return b.Fault }

func QueryVnicStatus(r soap.RoundTripper, req *types.QueryVnicStatus) (*types.QueryVnicStatusResponse, Error) {
	var body = QueryVnicStatusBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type QueryVsanObjectsBody struct {
	Req   *types.QueryVsanObjects         `xml:"urn:vim25 QueryVsanObjects,omitempty"`
	Res   *types.QueryVsanObjectsResponse `xml:"urn:vim25 QueryVsanObjectsResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryVsanObjectsBody) fault() *soap.Fault { return b.Fault }

func QueryVsanObjects(r soap.RoundTripper, req *types.QueryVsanObjects) (*types.QueryVsanObjectsResponse, Error) {
	var body = QueryVsanObjectsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReadEnvironmentVariableInGuestBody struct {
	Req   *types.ReadEnvironmentVariableInGuest         `xml:"urn:vim25 ReadEnvironmentVariableInGuest,omitempty"`
	Res   *types.ReadEnvironmentVariableInGuestResponse `xml:"urn:vim25 ReadEnvironmentVariableInGuestResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadEnvironmentVariableInGuestBody) fault() *soap.Fault { return b.Fault }

func ReadEnvironmentVariableInGuest(r soap.RoundTripper, req *types.ReadEnvironmentVariableInGuest) (*types.ReadEnvironmentVariableInGuestResponse, Error) {
	var body = ReadEnvironmentVariableInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReadNextEventsBody struct {
	Req   *types.ReadNextEvents         `xml:"urn:vim25 ReadNextEvents,omitempty"`
	Res   *types.ReadNextEventsResponse `xml:"urn:vim25 ReadNextEventsResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadNextEventsBody) fault() *soap.Fault { return b.Fault }

func ReadNextEvents(r soap.RoundTripper, req *types.ReadNextEvents) (*types.ReadNextEventsResponse, Error) {
	var body = ReadNextEventsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReadNextTasksBody struct {
	Req   *types.ReadNextTasks         `xml:"urn:vim25 ReadNextTasks,omitempty"`
	Res   *types.ReadNextTasksResponse `xml:"urn:vim25 ReadNextTasksResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadNextTasksBody) fault() *soap.Fault { return b.Fault }

func ReadNextTasks(r soap.RoundTripper, req *types.ReadNextTasks) (*types.ReadNextTasksResponse, Error) {
	var body = ReadNextTasksBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReadPreviousEventsBody struct {
	Req   *types.ReadPreviousEvents         `xml:"urn:vim25 ReadPreviousEvents,omitempty"`
	Res   *types.ReadPreviousEventsResponse `xml:"urn:vim25 ReadPreviousEventsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadPreviousEventsBody) fault() *soap.Fault { return b.Fault }

func ReadPreviousEvents(r soap.RoundTripper, req *types.ReadPreviousEvents) (*types.ReadPreviousEventsResponse, Error) {
	var body = ReadPreviousEventsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReadPreviousTasksBody struct {
	Req   *types.ReadPreviousTasks         `xml:"urn:vim25 ReadPreviousTasks,omitempty"`
	Res   *types.ReadPreviousTasksResponse `xml:"urn:vim25 ReadPreviousTasksResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReadPreviousTasksBody) fault() *soap.Fault { return b.Fault }

func ReadPreviousTasks(r soap.RoundTripper, req *types.ReadPreviousTasks) (*types.ReadPreviousTasksResponse, Error) {
	var body = ReadPreviousTasksBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RebootGuestBody struct {
	Req   *types.RebootGuest         `xml:"urn:vim25 RebootGuest,omitempty"`
	Res   *types.RebootGuestResponse `xml:"urn:vim25 RebootGuestResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebootGuestBody) fault() *soap.Fault { return b.Fault }

func RebootGuest(r soap.RoundTripper, req *types.RebootGuest) (*types.RebootGuestResponse, Error) {
	var body = RebootGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RebootHost_TaskBody struct {
	Req   *types.RebootHost_Task         `xml:"urn:vim25 RebootHost_Task,omitempty"`
	Res   *types.RebootHost_TaskResponse `xml:"urn:vim25 RebootHost_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RebootHost_TaskBody) fault() *soap.Fault { return b.Fault }

func RebootHost_Task(r soap.RoundTripper, req *types.RebootHost_Task) (*types.RebootHost_TaskResponse, Error) {
	var body = RebootHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RecommendDatastoresBody struct {
	Req   *types.RecommendDatastores         `xml:"urn:vim25 RecommendDatastores,omitempty"`
	Res   *types.RecommendDatastoresResponse `xml:"urn:vim25 RecommendDatastoresResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RecommendDatastoresBody) fault() *soap.Fault { return b.Fault }

func RecommendDatastores(r soap.RoundTripper, req *types.RecommendDatastores) (*types.RecommendDatastoresResponse, Error) {
	var body = RecommendDatastoresBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RecommendHostsForVmBody struct {
	Req   *types.RecommendHostsForVm         `xml:"urn:vim25 RecommendHostsForVm,omitempty"`
	Res   *types.RecommendHostsForVmResponse `xml:"urn:vim25 RecommendHostsForVmResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RecommendHostsForVmBody) fault() *soap.Fault { return b.Fault }

func RecommendHostsForVm(r soap.RoundTripper, req *types.RecommendHostsForVm) (*types.RecommendHostsForVmResponse, Error) {
	var body = RecommendHostsForVmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigVM_TaskBody struct {
	Req   *types.ReconfigVM_Task         `xml:"urn:vim25 ReconfigVM_Task,omitempty"`
	Res   *types.ReconfigVM_TaskResponse `xml:"urn:vim25 ReconfigVM_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigVM_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigVM_Task(r soap.RoundTripper, req *types.ReconfigVM_Task) (*types.ReconfigVM_TaskResponse, Error) {
	var body = ReconfigVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureAlarmBody struct {
	Req   *types.ReconfigureAlarm         `xml:"urn:vim25 ReconfigureAlarm,omitempty"`
	Res   *types.ReconfigureAlarmResponse `xml:"urn:vim25 ReconfigureAlarmResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureAlarmBody) fault() *soap.Fault { return b.Fault }

func ReconfigureAlarm(r soap.RoundTripper, req *types.ReconfigureAlarm) (*types.ReconfigureAlarmResponse, Error) {
	var body = ReconfigureAlarmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureAutostartBody struct {
	Req   *types.ReconfigureAutostart         `xml:"urn:vim25 ReconfigureAutostart,omitempty"`
	Res   *types.ReconfigureAutostartResponse `xml:"urn:vim25 ReconfigureAutostartResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureAutostartBody) fault() *soap.Fault { return b.Fault }

func ReconfigureAutostart(r soap.RoundTripper, req *types.ReconfigureAutostart) (*types.ReconfigureAutostartResponse, Error) {
	var body = ReconfigureAutostartBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureCluster_TaskBody struct {
	Req   *types.ReconfigureCluster_Task         `xml:"urn:vim25 ReconfigureCluster_Task,omitempty"`
	Res   *types.ReconfigureCluster_TaskResponse `xml:"urn:vim25 ReconfigureCluster_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureCluster_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureCluster_Task(r soap.RoundTripper, req *types.ReconfigureCluster_Task) (*types.ReconfigureCluster_TaskResponse, Error) {
	var body = ReconfigureCluster_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureComputeResource_TaskBody struct {
	Req   *types.ReconfigureComputeResource_Task         `xml:"urn:vim25 ReconfigureComputeResource_Task,omitempty"`
	Res   *types.ReconfigureComputeResource_TaskResponse `xml:"urn:vim25 ReconfigureComputeResource_TaskResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureComputeResource_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureComputeResource_Task(r soap.RoundTripper, req *types.ReconfigureComputeResource_Task) (*types.ReconfigureComputeResource_TaskResponse, Error) {
	var body = ReconfigureComputeResource_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureDVPort_TaskBody struct {
	Req   *types.ReconfigureDVPort_Task         `xml:"urn:vim25 ReconfigureDVPort_Task,omitempty"`
	Res   *types.ReconfigureDVPort_TaskResponse `xml:"urn:vim25 ReconfigureDVPort_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDVPort_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureDVPort_Task(r soap.RoundTripper, req *types.ReconfigureDVPort_Task) (*types.ReconfigureDVPort_TaskResponse, Error) {
	var body = ReconfigureDVPort_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureDVPortgroup_TaskBody struct {
	Req   *types.ReconfigureDVPortgroup_Task         `xml:"urn:vim25 ReconfigureDVPortgroup_Task,omitempty"`
	Res   *types.ReconfigureDVPortgroup_TaskResponse `xml:"urn:vim25 ReconfigureDVPortgroup_TaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDVPortgroup_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureDVPortgroup_Task(r soap.RoundTripper, req *types.ReconfigureDVPortgroup_Task) (*types.ReconfigureDVPortgroup_TaskResponse, Error) {
	var body = ReconfigureDVPortgroup_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureDatacenter_TaskBody struct {
	Req   *types.ReconfigureDatacenter_Task         `xml:"urn:vim25 ReconfigureDatacenter_Task,omitempty"`
	Res   *types.ReconfigureDatacenter_TaskResponse `xml:"urn:vim25 ReconfigureDatacenter_TaskResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDatacenter_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureDatacenter_Task(r soap.RoundTripper, req *types.ReconfigureDatacenter_Task) (*types.ReconfigureDatacenter_TaskResponse, Error) {
	var body = ReconfigureDatacenter_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureDvs_TaskBody struct {
	Req   *types.ReconfigureDvs_Task         `xml:"urn:vim25 ReconfigureDvs_Task,omitempty"`
	Res   *types.ReconfigureDvs_TaskResponse `xml:"urn:vim25 ReconfigureDvs_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureDvs_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureDvs_Task(r soap.RoundTripper, req *types.ReconfigureDvs_Task) (*types.ReconfigureDvs_TaskResponse, Error) {
	var body = ReconfigureDvs_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureHostForDAS_TaskBody struct {
	Req   *types.ReconfigureHostForDAS_Task         `xml:"urn:vim25 ReconfigureHostForDAS_Task,omitempty"`
	Res   *types.ReconfigureHostForDAS_TaskResponse `xml:"urn:vim25 ReconfigureHostForDAS_TaskResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureHostForDAS_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureHostForDAS_Task(r soap.RoundTripper, req *types.ReconfigureHostForDAS_Task) (*types.ReconfigureHostForDAS_TaskResponse, Error) {
	var body = ReconfigureHostForDAS_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureScheduledTaskBody struct {
	Req   *types.ReconfigureScheduledTask         `xml:"urn:vim25 ReconfigureScheduledTask,omitempty"`
	Res   *types.ReconfigureScheduledTaskResponse `xml:"urn:vim25 ReconfigureScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func ReconfigureScheduledTask(r soap.RoundTripper, req *types.ReconfigureScheduledTask) (*types.ReconfigureScheduledTaskResponse, Error) {
	var body = ReconfigureScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureServiceConsoleReservationBody struct {
	Req   *types.ReconfigureServiceConsoleReservation         `xml:"urn:vim25 ReconfigureServiceConsoleReservation,omitempty"`
	Res   *types.ReconfigureServiceConsoleReservationResponse `xml:"urn:vim25 ReconfigureServiceConsoleReservationResponse,omitempty"`
	Fault *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureServiceConsoleReservationBody) fault() *soap.Fault { return b.Fault }

func ReconfigureServiceConsoleReservation(r soap.RoundTripper, req *types.ReconfigureServiceConsoleReservation) (*types.ReconfigureServiceConsoleReservationResponse, Error) {
	var body = ReconfigureServiceConsoleReservationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureSnmpAgentBody struct {
	Req   *types.ReconfigureSnmpAgent         `xml:"urn:vim25 ReconfigureSnmpAgent,omitempty"`
	Res   *types.ReconfigureSnmpAgentResponse `xml:"urn:vim25 ReconfigureSnmpAgentResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureSnmpAgentBody) fault() *soap.Fault { return b.Fault }

func ReconfigureSnmpAgent(r soap.RoundTripper, req *types.ReconfigureSnmpAgent) (*types.ReconfigureSnmpAgentResponse, Error) {
	var body = ReconfigureSnmpAgentBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconfigureVirtualMachineReservationBody struct {
	Req   *types.ReconfigureVirtualMachineReservation         `xml:"urn:vim25 ReconfigureVirtualMachineReservation,omitempty"`
	Res   *types.ReconfigureVirtualMachineReservationResponse `xml:"urn:vim25 ReconfigureVirtualMachineReservationResponse,omitempty"`
	Fault *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconfigureVirtualMachineReservationBody) fault() *soap.Fault { return b.Fault }

func ReconfigureVirtualMachineReservation(r soap.RoundTripper, req *types.ReconfigureVirtualMachineReservation) (*types.ReconfigureVirtualMachineReservationResponse, Error) {
	var body = ReconfigureVirtualMachineReservationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReconnectHost_TaskBody struct {
	Req   *types.ReconnectHost_Task         `xml:"urn:vim25 ReconnectHost_Task,omitempty"`
	Res   *types.ReconnectHost_TaskResponse `xml:"urn:vim25 ReconnectHost_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReconnectHost_TaskBody) fault() *soap.Fault { return b.Fault }

func ReconnectHost_Task(r soap.RoundTripper, req *types.ReconnectHost_Task) (*types.ReconnectHost_TaskResponse, Error) {
	var body = ReconnectHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RectifyDvsHost_TaskBody struct {
	Req   *types.RectifyDvsHost_Task         `xml:"urn:vim25 RectifyDvsHost_Task,omitempty"`
	Res   *types.RectifyDvsHost_TaskResponse `xml:"urn:vim25 RectifyDvsHost_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RectifyDvsHost_TaskBody) fault() *soap.Fault { return b.Fault }

func RectifyDvsHost_Task(r soap.RoundTripper, req *types.RectifyDvsHost_Task) (*types.RectifyDvsHost_TaskResponse, Error) {
	var body = RectifyDvsHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RectifyDvsOnHost_TaskBody struct {
	Req   *types.RectifyDvsOnHost_Task         `xml:"urn:vim25 RectifyDvsOnHost_Task,omitempty"`
	Res   *types.RectifyDvsOnHost_TaskResponse `xml:"urn:vim25 RectifyDvsOnHost_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RectifyDvsOnHost_TaskBody) fault() *soap.Fault { return b.Fault }

func RectifyDvsOnHost_Task(r soap.RoundTripper, req *types.RectifyDvsOnHost_Task) (*types.RectifyDvsOnHost_TaskResponse, Error) {
	var body = RectifyDvsOnHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshBody struct {
	Req   *types.Refresh         `xml:"urn:vim25 Refresh,omitempty"`
	Res   *types.RefreshResponse `xml:"urn:vim25 RefreshResponse,omitempty"`
	Fault *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshBody) fault() *soap.Fault { return b.Fault }

func Refresh(r soap.RoundTripper, req *types.Refresh) (*types.RefreshResponse, Error) {
	var body = RefreshBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshDVPortStateBody struct {
	Req   *types.RefreshDVPortState         `xml:"urn:vim25 RefreshDVPortState,omitempty"`
	Res   *types.RefreshDVPortStateResponse `xml:"urn:vim25 RefreshDVPortStateResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDVPortStateBody) fault() *soap.Fault { return b.Fault }

func RefreshDVPortState(r soap.RoundTripper, req *types.RefreshDVPortState) (*types.RefreshDVPortStateResponse, Error) {
	var body = RefreshDVPortStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshDatastoreBody struct {
	Req   *types.RefreshDatastore         `xml:"urn:vim25 RefreshDatastore,omitempty"`
	Res   *types.RefreshDatastoreResponse `xml:"urn:vim25 RefreshDatastoreResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDatastoreBody) fault() *soap.Fault { return b.Fault }

func RefreshDatastore(r soap.RoundTripper, req *types.RefreshDatastore) (*types.RefreshDatastoreResponse, Error) {
	var body = RefreshDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshDatastoreStorageInfoBody struct {
	Req   *types.RefreshDatastoreStorageInfo         `xml:"urn:vim25 RefreshDatastoreStorageInfo,omitempty"`
	Res   *types.RefreshDatastoreStorageInfoResponse `xml:"urn:vim25 RefreshDatastoreStorageInfoResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDatastoreStorageInfoBody) fault() *soap.Fault { return b.Fault }

func RefreshDatastoreStorageInfo(r soap.RoundTripper, req *types.RefreshDatastoreStorageInfo) (*types.RefreshDatastoreStorageInfoResponse, Error) {
	var body = RefreshDatastoreStorageInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshDateTimeSystemBody struct {
	Req   *types.RefreshDateTimeSystem         `xml:"urn:vim25 RefreshDateTimeSystem,omitempty"`
	Res   *types.RefreshDateTimeSystemResponse `xml:"urn:vim25 RefreshDateTimeSystemResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshDateTimeSystemBody) fault() *soap.Fault { return b.Fault }

func RefreshDateTimeSystem(r soap.RoundTripper, req *types.RefreshDateTimeSystem) (*types.RefreshDateTimeSystemResponse, Error) {
	var body = RefreshDateTimeSystemBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshFirewallBody struct {
	Req   *types.RefreshFirewall         `xml:"urn:vim25 RefreshFirewall,omitempty"`
	Res   *types.RefreshFirewallResponse `xml:"urn:vim25 RefreshFirewallResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshFirewallBody) fault() *soap.Fault { return b.Fault }

func RefreshFirewall(r soap.RoundTripper, req *types.RefreshFirewall) (*types.RefreshFirewallResponse, Error) {
	var body = RefreshFirewallBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshGraphicsManagerBody struct {
	Req   *types.RefreshGraphicsManager         `xml:"urn:vim25 RefreshGraphicsManager,omitempty"`
	Res   *types.RefreshGraphicsManagerResponse `xml:"urn:vim25 RefreshGraphicsManagerResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshGraphicsManagerBody) fault() *soap.Fault { return b.Fault }

func RefreshGraphicsManager(r soap.RoundTripper, req *types.RefreshGraphicsManager) (*types.RefreshGraphicsManagerResponse, Error) {
	var body = RefreshGraphicsManagerBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshHealthStatusSystemBody struct {
	Req   *types.RefreshHealthStatusSystem         `xml:"urn:vim25 RefreshHealthStatusSystem,omitempty"`
	Res   *types.RefreshHealthStatusSystemResponse `xml:"urn:vim25 RefreshHealthStatusSystemResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshHealthStatusSystemBody) fault() *soap.Fault { return b.Fault }

func RefreshHealthStatusSystem(r soap.RoundTripper, req *types.RefreshHealthStatusSystem) (*types.RefreshHealthStatusSystemResponse, Error) {
	var body = RefreshHealthStatusSystemBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshNetworkSystemBody struct {
	Req   *types.RefreshNetworkSystem         `xml:"urn:vim25 RefreshNetworkSystem,omitempty"`
	Res   *types.RefreshNetworkSystemResponse `xml:"urn:vim25 RefreshNetworkSystemResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshNetworkSystemBody) fault() *soap.Fault { return b.Fault }

func RefreshNetworkSystem(r soap.RoundTripper, req *types.RefreshNetworkSystem) (*types.RefreshNetworkSystemResponse, Error) {
	var body = RefreshNetworkSystemBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshRecommendationBody struct {
	Req   *types.RefreshRecommendation         `xml:"urn:vim25 RefreshRecommendation,omitempty"`
	Res   *types.RefreshRecommendationResponse `xml:"urn:vim25 RefreshRecommendationResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshRecommendationBody) fault() *soap.Fault { return b.Fault }

func RefreshRecommendation(r soap.RoundTripper, req *types.RefreshRecommendation) (*types.RefreshRecommendationResponse, Error) {
	var body = RefreshRecommendationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshRuntimeBody struct {
	Req   *types.RefreshRuntime         `xml:"urn:vim25 RefreshRuntime,omitempty"`
	Res   *types.RefreshRuntimeResponse `xml:"urn:vim25 RefreshRuntimeResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshRuntimeBody) fault() *soap.Fault { return b.Fault }

func RefreshRuntime(r soap.RoundTripper, req *types.RefreshRuntime) (*types.RefreshRuntimeResponse, Error) {
	var body = RefreshRuntimeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshServicesBody struct {
	Req   *types.RefreshServices         `xml:"urn:vim25 RefreshServices,omitempty"`
	Res   *types.RefreshServicesResponse `xml:"urn:vim25 RefreshServicesResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshServicesBody) fault() *soap.Fault { return b.Fault }

func RefreshServices(r soap.RoundTripper, req *types.RefreshServices) (*types.RefreshServicesResponse, Error) {
	var body = RefreshServicesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshStorageDrsRecommendationBody struct {
	Req   *types.RefreshStorageDrsRecommendation         `xml:"urn:vim25 RefreshStorageDrsRecommendation,omitempty"`
	Res   *types.RefreshStorageDrsRecommendationResponse `xml:"urn:vim25 RefreshStorageDrsRecommendationResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageDrsRecommendationBody) fault() *soap.Fault { return b.Fault }

func RefreshStorageDrsRecommendation(r soap.RoundTripper, req *types.RefreshStorageDrsRecommendation) (*types.RefreshStorageDrsRecommendationResponse, Error) {
	var body = RefreshStorageDrsRecommendationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshStorageInfoBody struct {
	Req   *types.RefreshStorageInfo         `xml:"urn:vim25 RefreshStorageInfo,omitempty"`
	Res   *types.RefreshStorageInfoResponse `xml:"urn:vim25 RefreshStorageInfoResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageInfoBody) fault() *soap.Fault { return b.Fault }

func RefreshStorageInfo(r soap.RoundTripper, req *types.RefreshStorageInfo) (*types.RefreshStorageInfoResponse, Error) {
	var body = RefreshStorageInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RefreshStorageSystemBody struct {
	Req   *types.RefreshStorageSystem         `xml:"urn:vim25 RefreshStorageSystem,omitempty"`
	Res   *types.RefreshStorageSystemResponse `xml:"urn:vim25 RefreshStorageSystemResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RefreshStorageSystemBody) fault() *soap.Fault { return b.Fault }

func RefreshStorageSystem(r soap.RoundTripper, req *types.RefreshStorageSystem) (*types.RefreshStorageSystemResponse, Error) {
	var body = RefreshStorageSystemBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RegisterChildVM_TaskBody struct {
	Req   *types.RegisterChildVM_Task         `xml:"urn:vim25 RegisterChildVM_Task,omitempty"`
	Res   *types.RegisterChildVM_TaskResponse `xml:"urn:vim25 RegisterChildVM_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterChildVM_TaskBody) fault() *soap.Fault { return b.Fault }

func RegisterChildVM_Task(r soap.RoundTripper, req *types.RegisterChildVM_Task) (*types.RegisterChildVM_TaskResponse, Error) {
	var body = RegisterChildVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RegisterExtensionBody struct {
	Req   *types.RegisterExtension         `xml:"urn:vim25 RegisterExtension,omitempty"`
	Res   *types.RegisterExtensionResponse `xml:"urn:vim25 RegisterExtensionResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterExtensionBody) fault() *soap.Fault { return b.Fault }

func RegisterExtension(r soap.RoundTripper, req *types.RegisterExtension) (*types.RegisterExtensionResponse, Error) {
	var body = RegisterExtensionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RegisterVM_TaskBody struct {
	Req   *types.RegisterVM_Task         `xml:"urn:vim25 RegisterVM_Task,omitempty"`
	Res   *types.RegisterVM_TaskResponse `xml:"urn:vim25 RegisterVM_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterVM_TaskBody) fault() *soap.Fault { return b.Fault }

func RegisterVM_Task(r soap.RoundTripper, req *types.RegisterVM_Task) (*types.RegisterVM_TaskResponse, Error) {
	var body = RegisterVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReleaseCredentialsInGuestBody struct {
	Req   *types.ReleaseCredentialsInGuest         `xml:"urn:vim25 ReleaseCredentialsInGuest,omitempty"`
	Res   *types.ReleaseCredentialsInGuestResponse `xml:"urn:vim25 ReleaseCredentialsInGuestResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReleaseCredentialsInGuestBody) fault() *soap.Fault { return b.Fault }

func ReleaseCredentialsInGuest(r soap.RoundTripper, req *types.ReleaseCredentialsInGuest) (*types.ReleaseCredentialsInGuestResponse, Error) {
	var body = ReleaseCredentialsInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReleaseIpAllocationBody struct {
	Req   *types.ReleaseIpAllocation         `xml:"urn:vim25 ReleaseIpAllocation,omitempty"`
	Res   *types.ReleaseIpAllocationResponse `xml:"urn:vim25 ReleaseIpAllocationResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReleaseIpAllocationBody) fault() *soap.Fault { return b.Fault }

func ReleaseIpAllocation(r soap.RoundTripper, req *types.ReleaseIpAllocation) (*types.ReleaseIpAllocationResponse, Error) {
	var body = ReleaseIpAllocationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ReloadBody struct {
	Req   *types.Reload         `xml:"urn:vim25 Reload,omitempty"`
	Res   *types.ReloadResponse `xml:"urn:vim25 ReloadResponse,omitempty"`
	Fault *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ReloadBody) fault() *soap.Fault { return b.Fault }

func Reload(r soap.RoundTripper, req *types.Reload) (*types.ReloadResponse, Error) {
	var body = ReloadBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RelocateVM_TaskBody struct {
	Req   *types.RelocateVM_Task         `xml:"urn:vim25 RelocateVM_Task,omitempty"`
	Res   *types.RelocateVM_TaskResponse `xml:"urn:vim25 RelocateVM_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RelocateVM_TaskBody) fault() *soap.Fault { return b.Fault }

func RelocateVM_Task(r soap.RoundTripper, req *types.RelocateVM_Task) (*types.RelocateVM_TaskResponse, Error) {
	var body = RelocateVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveAlarmBody struct {
	Req   *types.RemoveAlarm         `xml:"urn:vim25 RemoveAlarm,omitempty"`
	Res   *types.RemoveAlarmResponse `xml:"urn:vim25 RemoveAlarmResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAlarmBody) fault() *soap.Fault { return b.Fault }

func RemoveAlarm(r soap.RoundTripper, req *types.RemoveAlarm) (*types.RemoveAlarmResponse, Error) {
	var body = RemoveAlarmBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveAllSnapshots_TaskBody struct {
	Req   *types.RemoveAllSnapshots_Task         `xml:"urn:vim25 RemoveAllSnapshots_Task,omitempty"`
	Res   *types.RemoveAllSnapshots_TaskResponse `xml:"urn:vim25 RemoveAllSnapshots_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAllSnapshots_TaskBody) fault() *soap.Fault { return b.Fault }

func RemoveAllSnapshots_Task(r soap.RoundTripper, req *types.RemoveAllSnapshots_Task) (*types.RemoveAllSnapshots_TaskResponse, Error) {
	var body = RemoveAllSnapshots_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveAssignedLicenseBody struct {
	Req   *types.RemoveAssignedLicense         `xml:"urn:vim25 RemoveAssignedLicense,omitempty"`
	Res   *types.RemoveAssignedLicenseResponse `xml:"urn:vim25 RemoveAssignedLicenseResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAssignedLicenseBody) fault() *soap.Fault { return b.Fault }

func RemoveAssignedLicense(r soap.RoundTripper, req *types.RemoveAssignedLicense) (*types.RemoveAssignedLicenseResponse, Error) {
	var body = RemoveAssignedLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveAuthorizationRoleBody struct {
	Req   *types.RemoveAuthorizationRole         `xml:"urn:vim25 RemoveAuthorizationRole,omitempty"`
	Res   *types.RemoveAuthorizationRoleResponse `xml:"urn:vim25 RemoveAuthorizationRoleResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveAuthorizationRoleBody) fault() *soap.Fault { return b.Fault }

func RemoveAuthorizationRole(r soap.RoundTripper, req *types.RemoveAuthorizationRole) (*types.RemoveAuthorizationRoleResponse, Error) {
	var body = RemoveAuthorizationRoleBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveCustomFieldDefBody struct {
	Req   *types.RemoveCustomFieldDef         `xml:"urn:vim25 RemoveCustomFieldDef,omitempty"`
	Res   *types.RemoveCustomFieldDefResponse `xml:"urn:vim25 RemoveCustomFieldDefResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveCustomFieldDefBody) fault() *soap.Fault { return b.Fault }

func RemoveCustomFieldDef(r soap.RoundTripper, req *types.RemoveCustomFieldDef) (*types.RemoveCustomFieldDefResponse, Error) {
	var body = RemoveCustomFieldDefBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveDatastoreBody struct {
	Req   *types.RemoveDatastore         `xml:"urn:vim25 RemoveDatastore,omitempty"`
	Res   *types.RemoveDatastoreResponse `xml:"urn:vim25 RemoveDatastoreResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDatastoreBody) fault() *soap.Fault { return b.Fault }

func RemoveDatastore(r soap.RoundTripper, req *types.RemoveDatastore) (*types.RemoveDatastoreResponse, Error) {
	var body = RemoveDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveDiskMapping_TaskBody struct {
	Req   *types.RemoveDiskMapping_Task         `xml:"urn:vim25 RemoveDiskMapping_Task,omitempty"`
	Res   *types.RemoveDiskMapping_TaskResponse `xml:"urn:vim25 RemoveDiskMapping_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDiskMapping_TaskBody) fault() *soap.Fault { return b.Fault }

func RemoveDiskMapping_Task(r soap.RoundTripper, req *types.RemoveDiskMapping_Task) (*types.RemoveDiskMapping_TaskResponse, Error) {
	var body = RemoveDiskMapping_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveDisk_TaskBody struct {
	Req   *types.RemoveDisk_Task         `xml:"urn:vim25 RemoveDisk_Task,omitempty"`
	Res   *types.RemoveDisk_TaskResponse `xml:"urn:vim25 RemoveDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func RemoveDisk_Task(r soap.RoundTripper, req *types.RemoveDisk_Task) (*types.RemoveDisk_TaskResponse, Error) {
	var body = RemoveDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveEntityPermissionBody struct {
	Req   *types.RemoveEntityPermission         `xml:"urn:vim25 RemoveEntityPermission,omitempty"`
	Res   *types.RemoveEntityPermissionResponse `xml:"urn:vim25 RemoveEntityPermissionResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveEntityPermissionBody) fault() *soap.Fault { return b.Fault }

func RemoveEntityPermission(r soap.RoundTripper, req *types.RemoveEntityPermission) (*types.RemoveEntityPermissionResponse, Error) {
	var body = RemoveEntityPermissionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveGroupBody struct {
	Req   *types.RemoveGroup         `xml:"urn:vim25 RemoveGroup,omitempty"`
	Res   *types.RemoveGroupResponse `xml:"urn:vim25 RemoveGroupResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveGroupBody) fault() *soap.Fault { return b.Fault }

func RemoveGroup(r soap.RoundTripper, req *types.RemoveGroup) (*types.RemoveGroupResponse, Error) {
	var body = RemoveGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveInternetScsiSendTargetsBody struct {
	Req   *types.RemoveInternetScsiSendTargets         `xml:"urn:vim25 RemoveInternetScsiSendTargets,omitempty"`
	Res   *types.RemoveInternetScsiSendTargetsResponse `xml:"urn:vim25 RemoveInternetScsiSendTargetsResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveInternetScsiSendTargetsBody) fault() *soap.Fault { return b.Fault }

func RemoveInternetScsiSendTargets(r soap.RoundTripper, req *types.RemoveInternetScsiSendTargets) (*types.RemoveInternetScsiSendTargetsResponse, Error) {
	var body = RemoveInternetScsiSendTargetsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveInternetScsiStaticTargetsBody struct {
	Req   *types.RemoveInternetScsiStaticTargets         `xml:"urn:vim25 RemoveInternetScsiStaticTargets,omitempty"`
	Res   *types.RemoveInternetScsiStaticTargetsResponse `xml:"urn:vim25 RemoveInternetScsiStaticTargetsResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveInternetScsiStaticTargetsBody) fault() *soap.Fault { return b.Fault }

func RemoveInternetScsiStaticTargets(r soap.RoundTripper, req *types.RemoveInternetScsiStaticTargets) (*types.RemoveInternetScsiStaticTargetsResponse, Error) {
	var body = RemoveInternetScsiStaticTargetsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveLicenseBody struct {
	Req   *types.RemoveLicense         `xml:"urn:vim25 RemoveLicense,omitempty"`
	Res   *types.RemoveLicenseResponse `xml:"urn:vim25 RemoveLicenseResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveLicenseBody) fault() *soap.Fault { return b.Fault }

func RemoveLicense(r soap.RoundTripper, req *types.RemoveLicense) (*types.RemoveLicenseResponse, Error) {
	var body = RemoveLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveLicenseLabelBody struct {
	Req   *types.RemoveLicenseLabel         `xml:"urn:vim25 RemoveLicenseLabel,omitempty"`
	Res   *types.RemoveLicenseLabelResponse `xml:"urn:vim25 RemoveLicenseLabelResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveLicenseLabelBody) fault() *soap.Fault { return b.Fault }

func RemoveLicenseLabel(r soap.RoundTripper, req *types.RemoveLicenseLabel) (*types.RemoveLicenseLabelResponse, Error) {
	var body = RemoveLicenseLabelBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveNetworkResourcePoolBody struct {
	Req   *types.RemoveNetworkResourcePool         `xml:"urn:vim25 RemoveNetworkResourcePool,omitempty"`
	Res   *types.RemoveNetworkResourcePoolResponse `xml:"urn:vim25 RemoveNetworkResourcePoolResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveNetworkResourcePoolBody) fault() *soap.Fault { return b.Fault }

func RemoveNetworkResourcePool(r soap.RoundTripper, req *types.RemoveNetworkResourcePool) (*types.RemoveNetworkResourcePoolResponse, Error) {
	var body = RemoveNetworkResourcePoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemovePerfIntervalBody struct {
	Req   *types.RemovePerfInterval         `xml:"urn:vim25 RemovePerfInterval,omitempty"`
	Res   *types.RemovePerfIntervalResponse `xml:"urn:vim25 RemovePerfIntervalResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemovePerfIntervalBody) fault() *soap.Fault { return b.Fault }

func RemovePerfInterval(r soap.RoundTripper, req *types.RemovePerfInterval) (*types.RemovePerfIntervalResponse, Error) {
	var body = RemovePerfIntervalBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemovePortGroupBody struct {
	Req   *types.RemovePortGroup         `xml:"urn:vim25 RemovePortGroup,omitempty"`
	Res   *types.RemovePortGroupResponse `xml:"urn:vim25 RemovePortGroupResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemovePortGroupBody) fault() *soap.Fault { return b.Fault }

func RemovePortGroup(r soap.RoundTripper, req *types.RemovePortGroup) (*types.RemovePortGroupResponse, Error) {
	var body = RemovePortGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveScheduledTaskBody struct {
	Req   *types.RemoveScheduledTask         `xml:"urn:vim25 RemoveScheduledTask,omitempty"`
	Res   *types.RemoveScheduledTaskResponse `xml:"urn:vim25 RemoveScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func RemoveScheduledTask(r soap.RoundTripper, req *types.RemoveScheduledTask) (*types.RemoveScheduledTaskResponse, Error) {
	var body = RemoveScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveServiceConsoleVirtualNicBody struct {
	Req   *types.RemoveServiceConsoleVirtualNic         `xml:"urn:vim25 RemoveServiceConsoleVirtualNic,omitempty"`
	Res   *types.RemoveServiceConsoleVirtualNicResponse `xml:"urn:vim25 RemoveServiceConsoleVirtualNicResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveServiceConsoleVirtualNicBody) fault() *soap.Fault { return b.Fault }

func RemoveServiceConsoleVirtualNic(r soap.RoundTripper, req *types.RemoveServiceConsoleVirtualNic) (*types.RemoveServiceConsoleVirtualNicResponse, Error) {
	var body = RemoveServiceConsoleVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveSnapshot_TaskBody struct {
	Req   *types.RemoveSnapshot_Task         `xml:"urn:vim25 RemoveSnapshot_Task,omitempty"`
	Res   *types.RemoveSnapshot_TaskResponse `xml:"urn:vim25 RemoveSnapshot_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveSnapshot_TaskBody) fault() *soap.Fault { return b.Fault }

func RemoveSnapshot_Task(r soap.RoundTripper, req *types.RemoveSnapshot_Task) (*types.RemoveSnapshot_TaskResponse, Error) {
	var body = RemoveSnapshot_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveUserBody struct {
	Req   *types.RemoveUser         `xml:"urn:vim25 RemoveUser,omitempty"`
	Res   *types.RemoveUserResponse `xml:"urn:vim25 RemoveUserResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveUserBody) fault() *soap.Fault { return b.Fault }

func RemoveUser(r soap.RoundTripper, req *types.RemoveUser) (*types.RemoveUserResponse, Error) {
	var body = RemoveUserBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveVirtualNicBody struct {
	Req   *types.RemoveVirtualNic         `xml:"urn:vim25 RemoveVirtualNic,omitempty"`
	Res   *types.RemoveVirtualNicResponse `xml:"urn:vim25 RemoveVirtualNicResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveVirtualNicBody) fault() *soap.Fault { return b.Fault }

func RemoveVirtualNic(r soap.RoundTripper, req *types.RemoveVirtualNic) (*types.RemoveVirtualNicResponse, Error) {
	var body = RemoveVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RemoveVirtualSwitchBody struct {
	Req   *types.RemoveVirtualSwitch         `xml:"urn:vim25 RemoveVirtualSwitch,omitempty"`
	Res   *types.RemoveVirtualSwitchResponse `xml:"urn:vim25 RemoveVirtualSwitchResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RemoveVirtualSwitchBody) fault() *soap.Fault { return b.Fault }

func RemoveVirtualSwitch(r soap.RoundTripper, req *types.RemoveVirtualSwitch) (*types.RemoveVirtualSwitchResponse, Error) {
	var body = RemoveVirtualSwitchBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RenameCustomFieldDefBody struct {
	Req   *types.RenameCustomFieldDef         `xml:"urn:vim25 RenameCustomFieldDef,omitempty"`
	Res   *types.RenameCustomFieldDefResponse `xml:"urn:vim25 RenameCustomFieldDefResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameCustomFieldDefBody) fault() *soap.Fault { return b.Fault }

func RenameCustomFieldDef(r soap.RoundTripper, req *types.RenameCustomFieldDef) (*types.RenameCustomFieldDefResponse, Error) {
	var body = RenameCustomFieldDefBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RenameCustomizationSpecBody struct {
	Req   *types.RenameCustomizationSpec         `xml:"urn:vim25 RenameCustomizationSpec,omitempty"`
	Res   *types.RenameCustomizationSpecResponse `xml:"urn:vim25 RenameCustomizationSpecResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameCustomizationSpecBody) fault() *soap.Fault { return b.Fault }

func RenameCustomizationSpec(r soap.RoundTripper, req *types.RenameCustomizationSpec) (*types.RenameCustomizationSpecResponse, Error) {
	var body = RenameCustomizationSpecBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RenameDatastoreBody struct {
	Req   *types.RenameDatastore         `xml:"urn:vim25 RenameDatastore,omitempty"`
	Res   *types.RenameDatastoreResponse `xml:"urn:vim25 RenameDatastoreResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameDatastoreBody) fault() *soap.Fault { return b.Fault }

func RenameDatastore(r soap.RoundTripper, req *types.RenameDatastore) (*types.RenameDatastoreResponse, Error) {
	var body = RenameDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RenameSnapshotBody struct {
	Req   *types.RenameSnapshot         `xml:"urn:vim25 RenameSnapshot,omitempty"`
	Res   *types.RenameSnapshotResponse `xml:"urn:vim25 RenameSnapshotResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RenameSnapshotBody) fault() *soap.Fault { return b.Fault }

func RenameSnapshot(r soap.RoundTripper, req *types.RenameSnapshot) (*types.RenameSnapshotResponse, Error) {
	var body = RenameSnapshotBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type Rename_TaskBody struct {
	Req   *types.Rename_Task         `xml:"urn:vim25 Rename_Task,omitempty"`
	Res   *types.Rename_TaskResponse `xml:"urn:vim25 Rename_TaskResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Rename_TaskBody) fault() *soap.Fault { return b.Fault }

func Rename_Task(r soap.RoundTripper, req *types.Rename_Task) (*types.Rename_TaskResponse, Error) {
	var body = Rename_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RescanAllHbaBody struct {
	Req   *types.RescanAllHba         `xml:"urn:vim25 RescanAllHba,omitempty"`
	Res   *types.RescanAllHbaResponse `xml:"urn:vim25 RescanAllHbaResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanAllHbaBody) fault() *soap.Fault { return b.Fault }

func RescanAllHba(r soap.RoundTripper, req *types.RescanAllHba) (*types.RescanAllHbaResponse, Error) {
	var body = RescanAllHbaBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RescanHbaBody struct {
	Req   *types.RescanHba         `xml:"urn:vim25 RescanHba,omitempty"`
	Res   *types.RescanHbaResponse `xml:"urn:vim25 RescanHbaResponse,omitempty"`
	Fault *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanHbaBody) fault() *soap.Fault { return b.Fault }

func RescanHba(r soap.RoundTripper, req *types.RescanHba) (*types.RescanHbaResponse, Error) {
	var body = RescanHbaBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RescanVffsBody struct {
	Req   *types.RescanVffs         `xml:"urn:vim25 RescanVffs,omitempty"`
	Res   *types.RescanVffsResponse `xml:"urn:vim25 RescanVffsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanVffsBody) fault() *soap.Fault { return b.Fault }

func RescanVffs(r soap.RoundTripper, req *types.RescanVffs) (*types.RescanVffsResponse, Error) {
	var body = RescanVffsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RescanVmfsBody struct {
	Req   *types.RescanVmfs         `xml:"urn:vim25 RescanVmfs,omitempty"`
	Res   *types.RescanVmfsResponse `xml:"urn:vim25 RescanVmfsResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RescanVmfsBody) fault() *soap.Fault { return b.Fault }

func RescanVmfs(r soap.RoundTripper, req *types.RescanVmfs) (*types.RescanVmfsResponse, Error) {
	var body = RescanVmfsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetCollectorBody struct {
	Req   *types.ResetCollector         `xml:"urn:vim25 ResetCollector,omitempty"`
	Res   *types.ResetCollectorResponse `xml:"urn:vim25 ResetCollectorResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetCollectorBody) fault() *soap.Fault { return b.Fault }

func ResetCollector(r soap.RoundTripper, req *types.ResetCollector) (*types.ResetCollectorResponse, Error) {
	var body = ResetCollectorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetCounterLevelMappingBody struct {
	Req   *types.ResetCounterLevelMapping         `xml:"urn:vim25 ResetCounterLevelMapping,omitempty"`
	Res   *types.ResetCounterLevelMappingResponse `xml:"urn:vim25 ResetCounterLevelMappingResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetCounterLevelMappingBody) fault() *soap.Fault { return b.Fault }

func ResetCounterLevelMapping(r soap.RoundTripper, req *types.ResetCounterLevelMapping) (*types.ResetCounterLevelMappingResponse, Error) {
	var body = ResetCounterLevelMappingBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetEntityPermissionsBody struct {
	Req   *types.ResetEntityPermissions         `xml:"urn:vim25 ResetEntityPermissions,omitempty"`
	Res   *types.ResetEntityPermissionsResponse `xml:"urn:vim25 ResetEntityPermissionsResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetEntityPermissionsBody) fault() *soap.Fault { return b.Fault }

func ResetEntityPermissions(r soap.RoundTripper, req *types.ResetEntityPermissions) (*types.ResetEntityPermissionsResponse, Error) {
	var body = ResetEntityPermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetFirmwareToFactoryDefaultsBody struct {
	Req   *types.ResetFirmwareToFactoryDefaults         `xml:"urn:vim25 ResetFirmwareToFactoryDefaults,omitempty"`
	Res   *types.ResetFirmwareToFactoryDefaultsResponse `xml:"urn:vim25 ResetFirmwareToFactoryDefaultsResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetFirmwareToFactoryDefaultsBody) fault() *soap.Fault { return b.Fault }

func ResetFirmwareToFactoryDefaults(r soap.RoundTripper, req *types.ResetFirmwareToFactoryDefaults) (*types.ResetFirmwareToFactoryDefaultsResponse, Error) {
	var body = ResetFirmwareToFactoryDefaultsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetGuestInformationBody struct {
	Req   *types.ResetGuestInformation         `xml:"urn:vim25 ResetGuestInformation,omitempty"`
	Res   *types.ResetGuestInformationResponse `xml:"urn:vim25 ResetGuestInformationResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetGuestInformationBody) fault() *soap.Fault { return b.Fault }

func ResetGuestInformation(r soap.RoundTripper, req *types.ResetGuestInformation) (*types.ResetGuestInformationResponse, Error) {
	var body = ResetGuestInformationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetListViewBody struct {
	Req   *types.ResetListView         `xml:"urn:vim25 ResetListView,omitempty"`
	Res   *types.ResetListViewResponse `xml:"urn:vim25 ResetListViewResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetListViewBody) fault() *soap.Fault { return b.Fault }

func ResetListView(r soap.RoundTripper, req *types.ResetListView) (*types.ResetListViewResponse, Error) {
	var body = ResetListViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetListViewFromViewBody struct {
	Req   *types.ResetListViewFromView         `xml:"urn:vim25 ResetListViewFromView,omitempty"`
	Res   *types.ResetListViewFromViewResponse `xml:"urn:vim25 ResetListViewFromViewResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetListViewFromViewBody) fault() *soap.Fault { return b.Fault }

func ResetListViewFromView(r soap.RoundTripper, req *types.ResetListViewFromView) (*types.ResetListViewFromViewResponse, Error) {
	var body = ResetListViewFromViewBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetSystemHealthInfoBody struct {
	Req   *types.ResetSystemHealthInfo         `xml:"urn:vim25 ResetSystemHealthInfo,omitempty"`
	Res   *types.ResetSystemHealthInfoResponse `xml:"urn:vim25 ResetSystemHealthInfoResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetSystemHealthInfoBody) fault() *soap.Fault { return b.Fault }

func ResetSystemHealthInfo(r soap.RoundTripper, req *types.ResetSystemHealthInfo) (*types.ResetSystemHealthInfoResponse, Error) {
	var body = ResetSystemHealthInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResetVM_TaskBody struct {
	Req   *types.ResetVM_Task         `xml:"urn:vim25 ResetVM_Task,omitempty"`
	Res   *types.ResetVM_TaskResponse `xml:"urn:vim25 ResetVM_TaskResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResetVM_TaskBody) fault() *soap.Fault { return b.Fault }

func ResetVM_Task(r soap.RoundTripper, req *types.ResetVM_Task) (*types.ResetVM_TaskResponse, Error) {
	var body = ResetVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResignatureUnresolvedVmfsVolume_TaskBody struct {
	Req   *types.ResignatureUnresolvedVmfsVolume_Task         `xml:"urn:vim25 ResignatureUnresolvedVmfsVolume_Task,omitempty"`
	Res   *types.ResignatureUnresolvedVmfsVolume_TaskResponse `xml:"urn:vim25 ResignatureUnresolvedVmfsVolume_TaskResponse,omitempty"`
	Fault *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResignatureUnresolvedVmfsVolume_TaskBody) fault() *soap.Fault { return b.Fault }

func ResignatureUnresolvedVmfsVolume_Task(r soap.RoundTripper, req *types.ResignatureUnresolvedVmfsVolume_Task) (*types.ResignatureUnresolvedVmfsVolume_TaskResponse, Error) {
	var body = ResignatureUnresolvedVmfsVolume_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResolveMultipleUnresolvedVmfsVolumesBody struct {
	Req   *types.ResolveMultipleUnresolvedVmfsVolumes         `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumes,omitempty"`
	Res   *types.ResolveMultipleUnresolvedVmfsVolumesResponse `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumesResponse,omitempty"`
	Fault *soap.Fault                                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveMultipleUnresolvedVmfsVolumesBody) fault() *soap.Fault { return b.Fault }

func ResolveMultipleUnresolvedVmfsVolumes(r soap.RoundTripper, req *types.ResolveMultipleUnresolvedVmfsVolumes) (*types.ResolveMultipleUnresolvedVmfsVolumesResponse, Error) {
	var body = ResolveMultipleUnresolvedVmfsVolumesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody struct {
	Req   *types.ResolveMultipleUnresolvedVmfsVolumesEx_Task         `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumesEx_Task,omitempty"`
	Res   *types.ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse `xml:"urn:vim25 ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse,omitempty"`
	Fault *soap.Fault                                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody) fault() *soap.Fault { return b.Fault }

func ResolveMultipleUnresolvedVmfsVolumesEx_Task(r soap.RoundTripper, req *types.ResolveMultipleUnresolvedVmfsVolumesEx_Task) (*types.ResolveMultipleUnresolvedVmfsVolumesEx_TaskResponse, Error) {
	var body = ResolveMultipleUnresolvedVmfsVolumesEx_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RestartServiceBody struct {
	Req   *types.RestartService         `xml:"urn:vim25 RestartService,omitempty"`
	Res   *types.RestartServiceResponse `xml:"urn:vim25 RestartServiceResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestartServiceBody) fault() *soap.Fault { return b.Fault }

func RestartService(r soap.RoundTripper, req *types.RestartService) (*types.RestartServiceResponse, Error) {
	var body = RestartServiceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RestartServiceConsoleVirtualNicBody struct {
	Req   *types.RestartServiceConsoleVirtualNic         `xml:"urn:vim25 RestartServiceConsoleVirtualNic,omitempty"`
	Res   *types.RestartServiceConsoleVirtualNicResponse `xml:"urn:vim25 RestartServiceConsoleVirtualNicResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestartServiceConsoleVirtualNicBody) fault() *soap.Fault { return b.Fault }

func RestartServiceConsoleVirtualNic(r soap.RoundTripper, req *types.RestartServiceConsoleVirtualNic) (*types.RestartServiceConsoleVirtualNicResponse, Error) {
	var body = RestartServiceConsoleVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RestoreFirmwareConfigurationBody struct {
	Req   *types.RestoreFirmwareConfiguration         `xml:"urn:vim25 RestoreFirmwareConfiguration,omitempty"`
	Res   *types.RestoreFirmwareConfigurationResponse `xml:"urn:vim25 RestoreFirmwareConfigurationResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RestoreFirmwareConfigurationBody) fault() *soap.Fault { return b.Fault }

func RestoreFirmwareConfiguration(r soap.RoundTripper, req *types.RestoreFirmwareConfiguration) (*types.RestoreFirmwareConfigurationResponse, Error) {
	var body = RestoreFirmwareConfigurationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveAllPermissionsBody struct {
	Req   *types.RetrieveAllPermissions         `xml:"urn:vim25 RetrieveAllPermissions,omitempty"`
	Res   *types.RetrieveAllPermissionsResponse `xml:"urn:vim25 RetrieveAllPermissionsResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAllPermissionsBody) fault() *soap.Fault { return b.Fault }

func RetrieveAllPermissions(r soap.RoundTripper, req *types.RetrieveAllPermissions) (*types.RetrieveAllPermissionsResponse, Error) {
	var body = RetrieveAllPermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveAnswerFileBody struct {
	Req   *types.RetrieveAnswerFile         `xml:"urn:vim25 RetrieveAnswerFile,omitempty"`
	Res   *types.RetrieveAnswerFileResponse `xml:"urn:vim25 RetrieveAnswerFileResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAnswerFileBody) fault() *soap.Fault { return b.Fault }

func RetrieveAnswerFile(r soap.RoundTripper, req *types.RetrieveAnswerFile) (*types.RetrieveAnswerFileResponse, Error) {
	var body = RetrieveAnswerFileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveAnswerFileForProfileBody struct {
	Req   *types.RetrieveAnswerFileForProfile         `xml:"urn:vim25 RetrieveAnswerFileForProfile,omitempty"`
	Res   *types.RetrieveAnswerFileForProfileResponse `xml:"urn:vim25 RetrieveAnswerFileForProfileResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveAnswerFileForProfileBody) fault() *soap.Fault { return b.Fault }

func RetrieveAnswerFileForProfile(r soap.RoundTripper, req *types.RetrieveAnswerFileForProfile) (*types.RetrieveAnswerFileForProfileResponse, Error) {
	var body = RetrieveAnswerFileForProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveArgumentDescriptionBody struct {
	Req   *types.RetrieveArgumentDescription         `xml:"urn:vim25 RetrieveArgumentDescription,omitempty"`
	Res   *types.RetrieveArgumentDescriptionResponse `xml:"urn:vim25 RetrieveArgumentDescriptionResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveArgumentDescriptionBody) fault() *soap.Fault { return b.Fault }

func RetrieveArgumentDescription(r soap.RoundTripper, req *types.RetrieveArgumentDescription) (*types.RetrieveArgumentDescriptionResponse, Error) {
	var body = RetrieveArgumentDescriptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveDasAdvancedRuntimeInfoBody struct {
	Req   *types.RetrieveDasAdvancedRuntimeInfo         `xml:"urn:vim25 RetrieveDasAdvancedRuntimeInfo,omitempty"`
	Res   *types.RetrieveDasAdvancedRuntimeInfoResponse `xml:"urn:vim25 RetrieveDasAdvancedRuntimeInfoResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDasAdvancedRuntimeInfoBody) fault() *soap.Fault { return b.Fault }

func RetrieveDasAdvancedRuntimeInfo(r soap.RoundTripper, req *types.RetrieveDasAdvancedRuntimeInfo) (*types.RetrieveDasAdvancedRuntimeInfoResponse, Error) {
	var body = RetrieveDasAdvancedRuntimeInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveDescriptionBody struct {
	Req   *types.RetrieveDescription         `xml:"urn:vim25 RetrieveDescription,omitempty"`
	Res   *types.RetrieveDescriptionResponse `xml:"urn:vim25 RetrieveDescriptionResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDescriptionBody) fault() *soap.Fault { return b.Fault }

func RetrieveDescription(r soap.RoundTripper, req *types.RetrieveDescription) (*types.RetrieveDescriptionResponse, Error) {
	var body = RetrieveDescriptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveDiskPartitionInfoBody struct {
	Req   *types.RetrieveDiskPartitionInfo         `xml:"urn:vim25 RetrieveDiskPartitionInfo,omitempty"`
	Res   *types.RetrieveDiskPartitionInfoResponse `xml:"urn:vim25 RetrieveDiskPartitionInfoResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveDiskPartitionInfoBody) fault() *soap.Fault { return b.Fault }

func RetrieveDiskPartitionInfo(r soap.RoundTripper, req *types.RetrieveDiskPartitionInfo) (*types.RetrieveDiskPartitionInfoResponse, Error) {
	var body = RetrieveDiskPartitionInfoBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveEntityPermissionsBody struct {
	Req   *types.RetrieveEntityPermissions         `xml:"urn:vim25 RetrieveEntityPermissions,omitempty"`
	Res   *types.RetrieveEntityPermissionsResponse `xml:"urn:vim25 RetrieveEntityPermissionsResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveEntityPermissionsBody) fault() *soap.Fault { return b.Fault }

func RetrieveEntityPermissions(r soap.RoundTripper, req *types.RetrieveEntityPermissions) (*types.RetrieveEntityPermissionsResponse, Error) {
	var body = RetrieveEntityPermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveEntityScheduledTaskBody struct {
	Req   *types.RetrieveEntityScheduledTask         `xml:"urn:vim25 RetrieveEntityScheduledTask,omitempty"`
	Res   *types.RetrieveEntityScheduledTaskResponse `xml:"urn:vim25 RetrieveEntityScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveEntityScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func RetrieveEntityScheduledTask(r soap.RoundTripper, req *types.RetrieveEntityScheduledTask) (*types.RetrieveEntityScheduledTaskResponse, Error) {
	var body = RetrieveEntityScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveHardwareUptimeBody struct {
	Req   *types.RetrieveHardwareUptime         `xml:"urn:vim25 RetrieveHardwareUptime,omitempty"`
	Res   *types.RetrieveHardwareUptimeResponse `xml:"urn:vim25 RetrieveHardwareUptimeResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveHardwareUptimeBody) fault() *soap.Fault { return b.Fault }

func RetrieveHardwareUptime(r soap.RoundTripper, req *types.RetrieveHardwareUptime) (*types.RetrieveHardwareUptimeResponse, Error) {
	var body = RetrieveHardwareUptimeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveObjectScheduledTaskBody struct {
	Req   *types.RetrieveObjectScheduledTask         `xml:"urn:vim25 RetrieveObjectScheduledTask,omitempty"`
	Res   *types.RetrieveObjectScheduledTaskResponse `xml:"urn:vim25 RetrieveObjectScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveObjectScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func RetrieveObjectScheduledTask(r soap.RoundTripper, req *types.RetrieveObjectScheduledTask) (*types.RetrieveObjectScheduledTaskResponse, Error) {
	var body = RetrieveObjectScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveProductComponentsBody struct {
	Req   *types.RetrieveProductComponents         `xml:"urn:vim25 RetrieveProductComponents,omitempty"`
	Res   *types.RetrieveProductComponentsResponse `xml:"urn:vim25 RetrieveProductComponentsResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveProductComponentsBody) fault() *soap.Fault { return b.Fault }

func RetrieveProductComponents(r soap.RoundTripper, req *types.RetrieveProductComponents) (*types.RetrieveProductComponentsResponse, Error) {
	var body = RetrieveProductComponentsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrievePropertiesBody struct {
	Req   *types.RetrieveProperties         `xml:"urn:vim25 RetrieveProperties,omitempty"`
	Res   *types.RetrievePropertiesResponse `xml:"urn:vim25 RetrievePropertiesResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrievePropertiesBody) fault() *soap.Fault { return b.Fault }

func RetrieveProperties(r soap.RoundTripper, req *types.RetrieveProperties) (*types.RetrievePropertiesResponse, Error) {
	var body = RetrievePropertiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrievePropertiesExBody struct {
	Req   *types.RetrievePropertiesEx         `xml:"urn:vim25 RetrievePropertiesEx,omitempty"`
	Res   *types.RetrievePropertiesExResponse `xml:"urn:vim25 RetrievePropertiesExResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrievePropertiesExBody) fault() *soap.Fault { return b.Fault }

func RetrievePropertiesEx(r soap.RoundTripper, req *types.RetrievePropertiesEx) (*types.RetrievePropertiesExResponse, Error) {
	var body = RetrievePropertiesExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveRolePermissionsBody struct {
	Req   *types.RetrieveRolePermissions         `xml:"urn:vim25 RetrieveRolePermissions,omitempty"`
	Res   *types.RetrieveRolePermissionsResponse `xml:"urn:vim25 RetrieveRolePermissionsResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveRolePermissionsBody) fault() *soap.Fault { return b.Fault }

func RetrieveRolePermissions(r soap.RoundTripper, req *types.RetrieveRolePermissions) (*types.RetrieveRolePermissionsResponse, Error) {
	var body = RetrieveRolePermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveServiceContentBody struct {
	Req   *types.RetrieveServiceContent         `xml:"urn:vim25 RetrieveServiceContent,omitempty"`
	Res   *types.RetrieveServiceContentResponse `xml:"urn:vim25 RetrieveServiceContentResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveServiceContentBody) fault() *soap.Fault { return b.Fault }

func RetrieveServiceContent(r soap.RoundTripper, req *types.RetrieveServiceContent) (*types.RetrieveServiceContentResponse, Error) {
	var body = RetrieveServiceContentBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RetrieveUserGroupsBody struct {
	Req   *types.RetrieveUserGroups         `xml:"urn:vim25 RetrieveUserGroups,omitempty"`
	Res   *types.RetrieveUserGroupsResponse `xml:"urn:vim25 RetrieveUserGroupsResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveUserGroupsBody) fault() *soap.Fault { return b.Fault }

func RetrieveUserGroups(r soap.RoundTripper, req *types.RetrieveUserGroups) (*types.RetrieveUserGroupsResponse, Error) {
	var body = RetrieveUserGroupsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RevertToCurrentSnapshot_TaskBody struct {
	Req   *types.RevertToCurrentSnapshot_Task         `xml:"urn:vim25 RevertToCurrentSnapshot_Task,omitempty"`
	Res   *types.RevertToCurrentSnapshot_TaskResponse `xml:"urn:vim25 RevertToCurrentSnapshot_TaskResponse,omitempty"`
	Fault *soap.Fault                                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertToCurrentSnapshot_TaskBody) fault() *soap.Fault { return b.Fault }

func RevertToCurrentSnapshot_Task(r soap.RoundTripper, req *types.RevertToCurrentSnapshot_Task) (*types.RevertToCurrentSnapshot_TaskResponse, Error) {
	var body = RevertToCurrentSnapshot_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RevertToSnapshot_TaskBody struct {
	Req   *types.RevertToSnapshot_Task         `xml:"urn:vim25 RevertToSnapshot_Task,omitempty"`
	Res   *types.RevertToSnapshot_TaskResponse `xml:"urn:vim25 RevertToSnapshot_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RevertToSnapshot_TaskBody) fault() *soap.Fault { return b.Fault }

func RevertToSnapshot_Task(r soap.RoundTripper, req *types.RevertToSnapshot_Task) (*types.RevertToSnapshot_TaskResponse, Error) {
	var body = RevertToSnapshot_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RewindCollectorBody struct {
	Req   *types.RewindCollector         `xml:"urn:vim25 RewindCollector,omitempty"`
	Res   *types.RewindCollectorResponse `xml:"urn:vim25 RewindCollectorResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RewindCollectorBody) fault() *soap.Fault { return b.Fault }

func RewindCollector(r soap.RoundTripper, req *types.RewindCollector) (*types.RewindCollectorResponse, Error) {
	var body = RewindCollectorBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type RunScheduledTaskBody struct {
	Req   *types.RunScheduledTask         `xml:"urn:vim25 RunScheduledTask,omitempty"`
	Res   *types.RunScheduledTaskResponse `xml:"urn:vim25 RunScheduledTaskResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RunScheduledTaskBody) fault() *soap.Fault { return b.Fault }

func RunScheduledTask(r soap.RoundTripper, req *types.RunScheduledTask) (*types.RunScheduledTaskResponse, Error) {
	var body = RunScheduledTaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ScanHostPatchV2_TaskBody struct {
	Req   *types.ScanHostPatchV2_Task         `xml:"urn:vim25 ScanHostPatchV2_Task,omitempty"`
	Res   *types.ScanHostPatchV2_TaskResponse `xml:"urn:vim25 ScanHostPatchV2_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScanHostPatchV2_TaskBody) fault() *soap.Fault { return b.Fault }

func ScanHostPatchV2_Task(r soap.RoundTripper, req *types.ScanHostPatchV2_Task) (*types.ScanHostPatchV2_TaskResponse, Error) {
	var body = ScanHostPatchV2_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ScanHostPatch_TaskBody struct {
	Req   *types.ScanHostPatch_Task         `xml:"urn:vim25 ScanHostPatch_Task,omitempty"`
	Res   *types.ScanHostPatch_TaskResponse `xml:"urn:vim25 ScanHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScanHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func ScanHostPatch_Task(r soap.RoundTripper, req *types.ScanHostPatch_Task) (*types.ScanHostPatch_TaskResponse, Error) {
	var body = ScanHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SearchDatastoreSubFolders_TaskBody struct {
	Req   *types.SearchDatastoreSubFolders_Task         `xml:"urn:vim25 SearchDatastoreSubFolders_Task,omitempty"`
	Res   *types.SearchDatastoreSubFolders_TaskResponse `xml:"urn:vim25 SearchDatastoreSubFolders_TaskResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SearchDatastoreSubFolders_TaskBody) fault() *soap.Fault { return b.Fault }

func SearchDatastoreSubFolders_Task(r soap.RoundTripper, req *types.SearchDatastoreSubFolders_Task) (*types.SearchDatastoreSubFolders_TaskResponse, Error) {
	var body = SearchDatastoreSubFolders_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SearchDatastore_TaskBody struct {
	Req   *types.SearchDatastore_Task         `xml:"urn:vim25 SearchDatastore_Task,omitempty"`
	Res   *types.SearchDatastore_TaskResponse `xml:"urn:vim25 SearchDatastore_TaskResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SearchDatastore_TaskBody) fault() *soap.Fault { return b.Fault }

func SearchDatastore_Task(r soap.RoundTripper, req *types.SearchDatastore_Task) (*types.SearchDatastore_TaskResponse, Error) {
	var body = SearchDatastore_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SelectActivePartitionBody struct {
	Req   *types.SelectActivePartition         `xml:"urn:vim25 SelectActivePartition,omitempty"`
	Res   *types.SelectActivePartitionResponse `xml:"urn:vim25 SelectActivePartitionResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectActivePartitionBody) fault() *soap.Fault { return b.Fault }

func SelectActivePartition(r soap.RoundTripper, req *types.SelectActivePartition) (*types.SelectActivePartitionResponse, Error) {
	var body = SelectActivePartitionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SelectVnicBody struct {
	Req   *types.SelectVnic         `xml:"urn:vim25 SelectVnic,omitempty"`
	Res   *types.SelectVnicResponse `xml:"urn:vim25 SelectVnicResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectVnicBody) fault() *soap.Fault { return b.Fault }

func SelectVnic(r soap.RoundTripper, req *types.SelectVnic) (*types.SelectVnicResponse, Error) {
	var body = SelectVnicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SelectVnicForNicTypeBody struct {
	Req   *types.SelectVnicForNicType         `xml:"urn:vim25 SelectVnicForNicType,omitempty"`
	Res   *types.SelectVnicForNicTypeResponse `xml:"urn:vim25 SelectVnicForNicTypeResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SelectVnicForNicTypeBody) fault() *soap.Fault { return b.Fault }

func SelectVnicForNicType(r soap.RoundTripper, req *types.SelectVnicForNicType) (*types.SelectVnicForNicTypeResponse, Error) {
	var body = SelectVnicForNicTypeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SendTestNotificationBody struct {
	Req   *types.SendTestNotification         `xml:"urn:vim25 SendTestNotification,omitempty"`
	Res   *types.SendTestNotificationResponse `xml:"urn:vim25 SendTestNotificationResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SendTestNotificationBody) fault() *soap.Fault { return b.Fault }

func SendTestNotification(r soap.RoundTripper, req *types.SendTestNotification) (*types.SendTestNotificationResponse, Error) {
	var body = SendTestNotificationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SessionIsActiveBody struct {
	Req   *types.SessionIsActive         `xml:"urn:vim25 SessionIsActive,omitempty"`
	Res   *types.SessionIsActiveResponse `xml:"urn:vim25 SessionIsActiveResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SessionIsActiveBody) fault() *soap.Fault { return b.Fault }

func SessionIsActive(r soap.RoundTripper, req *types.SessionIsActive) (*types.SessionIsActiveResponse, Error) {
	var body = SessionIsActiveBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetCollectorPageSizeBody struct {
	Req   *types.SetCollectorPageSize         `xml:"urn:vim25 SetCollectorPageSize,omitempty"`
	Res   *types.SetCollectorPageSizeResponse `xml:"urn:vim25 SetCollectorPageSizeResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetCollectorPageSizeBody) fault() *soap.Fault { return b.Fault }

func SetCollectorPageSize(r soap.RoundTripper, req *types.SetCollectorPageSize) (*types.SetCollectorPageSizeResponse, Error) {
	var body = SetCollectorPageSizeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetDisplayTopologyBody struct {
	Req   *types.SetDisplayTopology         `xml:"urn:vim25 SetDisplayTopology,omitempty"`
	Res   *types.SetDisplayTopologyResponse `xml:"urn:vim25 SetDisplayTopologyResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetDisplayTopologyBody) fault() *soap.Fault { return b.Fault }

func SetDisplayTopology(r soap.RoundTripper, req *types.SetDisplayTopology) (*types.SetDisplayTopologyResponse, Error) {
	var body = SetDisplayTopologyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetEntityPermissionsBody struct {
	Req   *types.SetEntityPermissions         `xml:"urn:vim25 SetEntityPermissions,omitempty"`
	Res   *types.SetEntityPermissionsResponse `xml:"urn:vim25 SetEntityPermissionsResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetEntityPermissionsBody) fault() *soap.Fault { return b.Fault }

func SetEntityPermissions(r soap.RoundTripper, req *types.SetEntityPermissions) (*types.SetEntityPermissionsResponse, Error) {
	var body = SetEntityPermissionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetExtensionCertificateBody struct {
	Req   *types.SetExtensionCertificate         `xml:"urn:vim25 SetExtensionCertificate,omitempty"`
	Res   *types.SetExtensionCertificateResponse `xml:"urn:vim25 SetExtensionCertificateResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetExtensionCertificateBody) fault() *soap.Fault { return b.Fault }

func SetExtensionCertificate(r soap.RoundTripper, req *types.SetExtensionCertificate) (*types.SetExtensionCertificateResponse, Error) {
	var body = SetExtensionCertificateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetFieldBody struct {
	Req   *types.SetField         `xml:"urn:vim25 SetField,omitempty"`
	Res   *types.SetFieldResponse `xml:"urn:vim25 SetFieldResponse,omitempty"`
	Fault *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetFieldBody) fault() *soap.Fault { return b.Fault }

func SetField(r soap.RoundTripper, req *types.SetField) (*types.SetFieldResponse, Error) {
	var body = SetFieldBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetLicenseEditionBody struct {
	Req   *types.SetLicenseEdition         `xml:"urn:vim25 SetLicenseEdition,omitempty"`
	Res   *types.SetLicenseEditionResponse `xml:"urn:vim25 SetLicenseEditionResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetLicenseEditionBody) fault() *soap.Fault { return b.Fault }

func SetLicenseEdition(r soap.RoundTripper, req *types.SetLicenseEdition) (*types.SetLicenseEditionResponse, Error) {
	var body = SetLicenseEditionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetLocaleBody struct {
	Req   *types.SetLocale         `xml:"urn:vim25 SetLocale,omitempty"`
	Res   *types.SetLocaleResponse `xml:"urn:vim25 SetLocaleResponse,omitempty"`
	Fault *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetLocaleBody) fault() *soap.Fault { return b.Fault }

func SetLocale(r soap.RoundTripper, req *types.SetLocale) (*types.SetLocaleResponse, Error) {
	var body = SetLocaleBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetMultipathLunPolicyBody struct {
	Req   *types.SetMultipathLunPolicy         `xml:"urn:vim25 SetMultipathLunPolicy,omitempty"`
	Res   *types.SetMultipathLunPolicyResponse `xml:"urn:vim25 SetMultipathLunPolicyResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetMultipathLunPolicyBody) fault() *soap.Fault { return b.Fault }

func SetMultipathLunPolicy(r soap.RoundTripper, req *types.SetMultipathLunPolicy) (*types.SetMultipathLunPolicyResponse, Error) {
	var body = SetMultipathLunPolicyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetPublicKeyBody struct {
	Req   *types.SetPublicKey         `xml:"urn:vim25 SetPublicKey,omitempty"`
	Res   *types.SetPublicKeyResponse `xml:"urn:vim25 SetPublicKeyResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetPublicKeyBody) fault() *soap.Fault { return b.Fault }

func SetPublicKey(r soap.RoundTripper, req *types.SetPublicKey) (*types.SetPublicKeyResponse, Error) {
	var body = SetPublicKeyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetScreenResolutionBody struct {
	Req   *types.SetScreenResolution         `xml:"urn:vim25 SetScreenResolution,omitempty"`
	Res   *types.SetScreenResolutionResponse `xml:"urn:vim25 SetScreenResolutionResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetScreenResolutionBody) fault() *soap.Fault { return b.Fault }

func SetScreenResolution(r soap.RoundTripper, req *types.SetScreenResolution) (*types.SetScreenResolutionResponse, Error) {
	var body = SetScreenResolutionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetTaskDescriptionBody struct {
	Req   *types.SetTaskDescription         `xml:"urn:vim25 SetTaskDescription,omitempty"`
	Res   *types.SetTaskDescriptionResponse `xml:"urn:vim25 SetTaskDescriptionResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetTaskDescriptionBody) fault() *soap.Fault { return b.Fault }

func SetTaskDescription(r soap.RoundTripper, req *types.SetTaskDescription) (*types.SetTaskDescriptionResponse, Error) {
	var body = SetTaskDescriptionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetTaskStateBody struct {
	Req   *types.SetTaskState         `xml:"urn:vim25 SetTaskState,omitempty"`
	Res   *types.SetTaskStateResponse `xml:"urn:vim25 SetTaskStateResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetTaskStateBody) fault() *soap.Fault { return b.Fault }

func SetTaskState(r soap.RoundTripper, req *types.SetTaskState) (*types.SetTaskStateResponse, Error) {
	var body = SetTaskStateBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SetVirtualDiskUuidBody struct {
	Req   *types.SetVirtualDiskUuid         `xml:"urn:vim25 SetVirtualDiskUuid,omitempty"`
	Res   *types.SetVirtualDiskUuidResponse `xml:"urn:vim25 SetVirtualDiskUuidResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetVirtualDiskUuidBody) fault() *soap.Fault { return b.Fault }

func SetVirtualDiskUuid(r soap.RoundTripper, req *types.SetVirtualDiskUuid) (*types.SetVirtualDiskUuidResponse, Error) {
	var body = SetVirtualDiskUuidBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ShrinkVirtualDisk_TaskBody struct {
	Req   *types.ShrinkVirtualDisk_Task         `xml:"urn:vim25 ShrinkVirtualDisk_Task,omitempty"`
	Res   *types.ShrinkVirtualDisk_TaskResponse `xml:"urn:vim25 ShrinkVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShrinkVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func ShrinkVirtualDisk_Task(r soap.RoundTripper, req *types.ShrinkVirtualDisk_Task) (*types.ShrinkVirtualDisk_TaskResponse, Error) {
	var body = ShrinkVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ShutdownGuestBody struct {
	Req   *types.ShutdownGuest         `xml:"urn:vim25 ShutdownGuest,omitempty"`
	Res   *types.ShutdownGuestResponse `xml:"urn:vim25 ShutdownGuestResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShutdownGuestBody) fault() *soap.Fault { return b.Fault }

func ShutdownGuest(r soap.RoundTripper, req *types.ShutdownGuest) (*types.ShutdownGuestResponse, Error) {
	var body = ShutdownGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ShutdownHost_TaskBody struct {
	Req   *types.ShutdownHost_Task         `xml:"urn:vim25 ShutdownHost_Task,omitempty"`
	Res   *types.ShutdownHost_TaskResponse `xml:"urn:vim25 ShutdownHost_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ShutdownHost_TaskBody) fault() *soap.Fault { return b.Fault }

func ShutdownHost_Task(r soap.RoundTripper, req *types.ShutdownHost_Task) (*types.ShutdownHost_TaskResponse, Error) {
	var body = ShutdownHost_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StageHostPatch_TaskBody struct {
	Req   *types.StageHostPatch_Task         `xml:"urn:vim25 StageHostPatch_Task,omitempty"`
	Res   *types.StageHostPatch_TaskResponse `xml:"urn:vim25 StageHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StageHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func StageHostPatch_Task(r soap.RoundTripper, req *types.StageHostPatch_Task) (*types.StageHostPatch_TaskResponse, Error) {
	var body = StageHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StandbyGuestBody struct {
	Req   *types.StandbyGuest         `xml:"urn:vim25 StandbyGuest,omitempty"`
	Res   *types.StandbyGuestResponse `xml:"urn:vim25 StandbyGuestResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StandbyGuestBody) fault() *soap.Fault { return b.Fault }

func StandbyGuest(r soap.RoundTripper, req *types.StandbyGuest) (*types.StandbyGuestResponse, Error) {
	var body = StandbyGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StartProgramInGuestBody struct {
	Req   *types.StartProgramInGuest         `xml:"urn:vim25 StartProgramInGuest,omitempty"`
	Res   *types.StartProgramInGuestResponse `xml:"urn:vim25 StartProgramInGuestResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartProgramInGuestBody) fault() *soap.Fault { return b.Fault }

func StartProgramInGuest(r soap.RoundTripper, req *types.StartProgramInGuest) (*types.StartProgramInGuestResponse, Error) {
	var body = StartProgramInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StartRecording_TaskBody struct {
	Req   *types.StartRecording_Task         `xml:"urn:vim25 StartRecording_Task,omitempty"`
	Res   *types.StartRecording_TaskResponse `xml:"urn:vim25 StartRecording_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartRecording_TaskBody) fault() *soap.Fault { return b.Fault }

func StartRecording_Task(r soap.RoundTripper, req *types.StartRecording_Task) (*types.StartRecording_TaskResponse, Error) {
	var body = StartRecording_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StartReplaying_TaskBody struct {
	Req   *types.StartReplaying_Task         `xml:"urn:vim25 StartReplaying_Task,omitempty"`
	Res   *types.StartReplaying_TaskResponse `xml:"urn:vim25 StartReplaying_TaskResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartReplaying_TaskBody) fault() *soap.Fault { return b.Fault }

func StartReplaying_Task(r soap.RoundTripper, req *types.StartReplaying_Task) (*types.StartReplaying_TaskResponse, Error) {
	var body = StartReplaying_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StartServiceBody struct {
	Req   *types.StartService         `xml:"urn:vim25 StartService,omitempty"`
	Res   *types.StartServiceResponse `xml:"urn:vim25 StartServiceResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StartServiceBody) fault() *soap.Fault { return b.Fault }

func StartService(r soap.RoundTripper, req *types.StartService) (*types.StartServiceResponse, Error) {
	var body = StartServiceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StopRecording_TaskBody struct {
	Req   *types.StopRecording_Task         `xml:"urn:vim25 StopRecording_Task,omitempty"`
	Res   *types.StopRecording_TaskResponse `xml:"urn:vim25 StopRecording_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopRecording_TaskBody) fault() *soap.Fault { return b.Fault }

func StopRecording_Task(r soap.RoundTripper, req *types.StopRecording_Task) (*types.StopRecording_TaskResponse, Error) {
	var body = StopRecording_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StopReplaying_TaskBody struct {
	Req   *types.StopReplaying_Task         `xml:"urn:vim25 StopReplaying_Task,omitempty"`
	Res   *types.StopReplaying_TaskResponse `xml:"urn:vim25 StopReplaying_TaskResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopReplaying_TaskBody) fault() *soap.Fault { return b.Fault }

func StopReplaying_Task(r soap.RoundTripper, req *types.StopReplaying_Task) (*types.StopReplaying_TaskResponse, Error) {
	var body = StopReplaying_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type StopServiceBody struct {
	Req   *types.StopService         `xml:"urn:vim25 StopService,omitempty"`
	Res   *types.StopServiceResponse `xml:"urn:vim25 StopServiceResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *StopServiceBody) fault() *soap.Fault { return b.Fault }

func StopService(r soap.RoundTripper, req *types.StopService) (*types.StopServiceResponse, Error) {
	var body = StopServiceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SuspendVApp_TaskBody struct {
	Req   *types.SuspendVApp_Task         `xml:"urn:vim25 SuspendVApp_Task,omitempty"`
	Res   *types.SuspendVApp_TaskResponse `xml:"urn:vim25 SuspendVApp_TaskResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SuspendVApp_TaskBody) fault() *soap.Fault { return b.Fault }

func SuspendVApp_Task(r soap.RoundTripper, req *types.SuspendVApp_Task) (*types.SuspendVApp_TaskResponse, Error) {
	var body = SuspendVApp_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type SuspendVM_TaskBody struct {
	Req   *types.SuspendVM_Task         `xml:"urn:vim25 SuspendVM_Task,omitempty"`
	Res   *types.SuspendVM_TaskResponse `xml:"urn:vim25 SuspendVM_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SuspendVM_TaskBody) fault() *soap.Fault { return b.Fault }

func SuspendVM_Task(r soap.RoundTripper, req *types.SuspendVM_Task) (*types.SuspendVM_TaskResponse, Error) {
	var body = SuspendVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type TerminateFaultTolerantVM_TaskBody struct {
	Req   *types.TerminateFaultTolerantVM_Task         `xml:"urn:vim25 TerminateFaultTolerantVM_Task,omitempty"`
	Res   *types.TerminateFaultTolerantVM_TaskResponse `xml:"urn:vim25 TerminateFaultTolerantVM_TaskResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateFaultTolerantVM_TaskBody) fault() *soap.Fault { return b.Fault }

func TerminateFaultTolerantVM_Task(r soap.RoundTripper, req *types.TerminateFaultTolerantVM_Task) (*types.TerminateFaultTolerantVM_TaskResponse, Error) {
	var body = TerminateFaultTolerantVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type TerminateProcessInGuestBody struct {
	Req   *types.TerminateProcessInGuest         `xml:"urn:vim25 TerminateProcessInGuest,omitempty"`
	Res   *types.TerminateProcessInGuestResponse `xml:"urn:vim25 TerminateProcessInGuestResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateProcessInGuestBody) fault() *soap.Fault { return b.Fault }

func TerminateProcessInGuest(r soap.RoundTripper, req *types.TerminateProcessInGuest) (*types.TerminateProcessInGuestResponse, Error) {
	var body = TerminateProcessInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type TerminateSessionBody struct {
	Req   *types.TerminateSession         `xml:"urn:vim25 TerminateSession,omitempty"`
	Res   *types.TerminateSessionResponse `xml:"urn:vim25 TerminateSessionResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateSessionBody) fault() *soap.Fault { return b.Fault }

func TerminateSession(r soap.RoundTripper, req *types.TerminateSession) (*types.TerminateSessionResponse, Error) {
	var body = TerminateSessionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type TerminateVMBody struct {
	Req   *types.TerminateVM         `xml:"urn:vim25 TerminateVM,omitempty"`
	Res   *types.TerminateVMResponse `xml:"urn:vim25 TerminateVMResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TerminateVMBody) fault() *soap.Fault { return b.Fault }

func TerminateVM(r soap.RoundTripper, req *types.TerminateVM) (*types.TerminateVMResponse, Error) {
	var body = TerminateVMBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type TurnOffFaultToleranceForVM_TaskBody struct {
	Req   *types.TurnOffFaultToleranceForVM_Task         `xml:"urn:vim25 TurnOffFaultToleranceForVM_Task,omitempty"`
	Res   *types.TurnOffFaultToleranceForVM_TaskResponse `xml:"urn:vim25 TurnOffFaultToleranceForVM_TaskResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *TurnOffFaultToleranceForVM_TaskBody) fault() *soap.Fault { return b.Fault }

func TurnOffFaultToleranceForVM_Task(r soap.RoundTripper, req *types.TurnOffFaultToleranceForVM_Task) (*types.TurnOffFaultToleranceForVM_TaskResponse, Error) {
	var body = TurnOffFaultToleranceForVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnassignUserFromGroupBody struct {
	Req   *types.UnassignUserFromGroup         `xml:"urn:vim25 UnassignUserFromGroup,omitempty"`
	Res   *types.UnassignUserFromGroupResponse `xml:"urn:vim25 UnassignUserFromGroupResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnassignUserFromGroupBody) fault() *soap.Fault { return b.Fault }

func UnassignUserFromGroup(r soap.RoundTripper, req *types.UnassignUserFromGroup) (*types.UnassignUserFromGroupResponse, Error) {
	var body = UnassignUserFromGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnbindVnicBody struct {
	Req   *types.UnbindVnic         `xml:"urn:vim25 UnbindVnic,omitempty"`
	Res   *types.UnbindVnicResponse `xml:"urn:vim25 UnbindVnicResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnbindVnicBody) fault() *soap.Fault { return b.Fault }

func UnbindVnic(r soap.RoundTripper, req *types.UnbindVnic) (*types.UnbindVnicResponse, Error) {
	var body = UnbindVnicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UninstallHostPatch_TaskBody struct {
	Req   *types.UninstallHostPatch_Task         `xml:"urn:vim25 UninstallHostPatch_Task,omitempty"`
	Res   *types.UninstallHostPatch_TaskResponse `xml:"urn:vim25 UninstallHostPatch_TaskResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallHostPatch_TaskBody) fault() *soap.Fault { return b.Fault }

func UninstallHostPatch_Task(r soap.RoundTripper, req *types.UninstallHostPatch_Task) (*types.UninstallHostPatch_TaskResponse, Error) {
	var body = UninstallHostPatch_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UninstallServiceBody struct {
	Req   *types.UninstallService         `xml:"urn:vim25 UninstallService,omitempty"`
	Res   *types.UninstallServiceResponse `xml:"urn:vim25 UninstallServiceResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallServiceBody) fault() *soap.Fault { return b.Fault }

func UninstallService(r soap.RoundTripper, req *types.UninstallService) (*types.UninstallServiceResponse, Error) {
	var body = UninstallServiceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnmountForceMountedVmfsVolumeBody struct {
	Req   *types.UnmountForceMountedVmfsVolume         `xml:"urn:vim25 UnmountForceMountedVmfsVolume,omitempty"`
	Res   *types.UnmountForceMountedVmfsVolumeResponse `xml:"urn:vim25 UnmountForceMountedVmfsVolumeResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountForceMountedVmfsVolumeBody) fault() *soap.Fault { return b.Fault }

func UnmountForceMountedVmfsVolume(r soap.RoundTripper, req *types.UnmountForceMountedVmfsVolume) (*types.UnmountForceMountedVmfsVolumeResponse, Error) {
	var body = UnmountForceMountedVmfsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnmountToolsInstallerBody struct {
	Req   *types.UnmountToolsInstaller         `xml:"urn:vim25 UnmountToolsInstaller,omitempty"`
	Res   *types.UnmountToolsInstallerResponse `xml:"urn:vim25 UnmountToolsInstallerResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountToolsInstallerBody) fault() *soap.Fault { return b.Fault }

func UnmountToolsInstaller(r soap.RoundTripper, req *types.UnmountToolsInstaller) (*types.UnmountToolsInstallerResponse, Error) {
	var body = UnmountToolsInstallerBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnmountVffsVolumeBody struct {
	Req   *types.UnmountVffsVolume         `xml:"urn:vim25 UnmountVffsVolume,omitempty"`
	Res   *types.UnmountVffsVolumeResponse `xml:"urn:vim25 UnmountVffsVolumeResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountVffsVolumeBody) fault() *soap.Fault { return b.Fault }

func UnmountVffsVolume(r soap.RoundTripper, req *types.UnmountVffsVolume) (*types.UnmountVffsVolumeResponse, Error) {
	var body = UnmountVffsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnmountVmfsVolumeBody struct {
	Req   *types.UnmountVmfsVolume         `xml:"urn:vim25 UnmountVmfsVolume,omitempty"`
	Res   *types.UnmountVmfsVolumeResponse `xml:"urn:vim25 UnmountVmfsVolumeResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnmountVmfsVolumeBody) fault() *soap.Fault { return b.Fault }

func UnmountVmfsVolume(r soap.RoundTripper, req *types.UnmountVmfsVolume) (*types.UnmountVmfsVolumeResponse, Error) {
	var body = UnmountVmfsVolumeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnregisterAndDestroy_TaskBody struct {
	Req   *types.UnregisterAndDestroy_Task         `xml:"urn:vim25 UnregisterAndDestroy_Task,omitempty"`
	Res   *types.UnregisterAndDestroy_TaskResponse `xml:"urn:vim25 UnregisterAndDestroy_TaskResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterAndDestroy_TaskBody) fault() *soap.Fault { return b.Fault }

func UnregisterAndDestroy_Task(r soap.RoundTripper, req *types.UnregisterAndDestroy_Task) (*types.UnregisterAndDestroy_TaskResponse, Error) {
	var body = UnregisterAndDestroy_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnregisterExtensionBody struct {
	Req   *types.UnregisterExtension         `xml:"urn:vim25 UnregisterExtension,omitempty"`
	Res   *types.UnregisterExtensionResponse `xml:"urn:vim25 UnregisterExtensionResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterExtensionBody) fault() *soap.Fault { return b.Fault }

func UnregisterExtension(r soap.RoundTripper, req *types.UnregisterExtension) (*types.UnregisterExtensionResponse, Error) {
	var body = UnregisterExtensionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UnregisterVMBody struct {
	Req   *types.UnregisterVM         `xml:"urn:vim25 UnregisterVM,omitempty"`
	Res   *types.UnregisterVMResponse `xml:"urn:vim25 UnregisterVMResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterVMBody) fault() *soap.Fault { return b.Fault }

func UnregisterVM(r soap.RoundTripper, req *types.UnregisterVM) (*types.UnregisterVMResponse, Error) {
	var body = UnregisterVMBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateAnswerFile_TaskBody struct {
	Req   *types.UpdateAnswerFile_Task         `xml:"urn:vim25 UpdateAnswerFile_Task,omitempty"`
	Res   *types.UpdateAnswerFile_TaskResponse `xml:"urn:vim25 UpdateAnswerFile_TaskResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAnswerFile_TaskBody) fault() *soap.Fault { return b.Fault }

func UpdateAnswerFile_Task(r soap.RoundTripper, req *types.UpdateAnswerFile_Task) (*types.UpdateAnswerFile_TaskResponse, Error) {
	var body = UpdateAnswerFile_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateAssignedLicenseBody struct {
	Req   *types.UpdateAssignedLicense         `xml:"urn:vim25 UpdateAssignedLicense,omitempty"`
	Res   *types.UpdateAssignedLicenseResponse `xml:"urn:vim25 UpdateAssignedLicenseResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAssignedLicenseBody) fault() *soap.Fault { return b.Fault }

func UpdateAssignedLicense(r soap.RoundTripper, req *types.UpdateAssignedLicense) (*types.UpdateAssignedLicenseResponse, Error) {
	var body = UpdateAssignedLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateAuthorizationRoleBody struct {
	Req   *types.UpdateAuthorizationRole         `xml:"urn:vim25 UpdateAuthorizationRole,omitempty"`
	Res   *types.UpdateAuthorizationRoleResponse `xml:"urn:vim25 UpdateAuthorizationRoleResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateAuthorizationRoleBody) fault() *soap.Fault { return b.Fault }

func UpdateAuthorizationRole(r soap.RoundTripper, req *types.UpdateAuthorizationRole) (*types.UpdateAuthorizationRoleResponse, Error) {
	var body = UpdateAuthorizationRoleBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateBootDeviceBody struct {
	Req   *types.UpdateBootDevice         `xml:"urn:vim25 UpdateBootDevice,omitempty"`
	Res   *types.UpdateBootDeviceResponse `xml:"urn:vim25 UpdateBootDeviceResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateBootDeviceBody) fault() *soap.Fault { return b.Fault }

func UpdateBootDevice(r soap.RoundTripper, req *types.UpdateBootDevice) (*types.UpdateBootDeviceResponse, Error) {
	var body = UpdateBootDeviceBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateChildResourceConfigurationBody struct {
	Req   *types.UpdateChildResourceConfiguration         `xml:"urn:vim25 UpdateChildResourceConfiguration,omitempty"`
	Res   *types.UpdateChildResourceConfigurationResponse `xml:"urn:vim25 UpdateChildResourceConfigurationResponse,omitempty"`
	Fault *soap.Fault                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateChildResourceConfigurationBody) fault() *soap.Fault { return b.Fault }

func UpdateChildResourceConfiguration(r soap.RoundTripper, req *types.UpdateChildResourceConfiguration) (*types.UpdateChildResourceConfigurationResponse, Error) {
	var body = UpdateChildResourceConfigurationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateClusterProfileBody struct {
	Req   *types.UpdateClusterProfile         `xml:"urn:vim25 UpdateClusterProfile,omitempty"`
	Res   *types.UpdateClusterProfileResponse `xml:"urn:vim25 UpdateClusterProfileResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateClusterProfileBody) fault() *soap.Fault { return b.Fault }

func UpdateClusterProfile(r soap.RoundTripper, req *types.UpdateClusterProfile) (*types.UpdateClusterProfileResponse, Error) {
	var body = UpdateClusterProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateConfigBody struct {
	Req   *types.UpdateConfig         `xml:"urn:vim25 UpdateConfig,omitempty"`
	Res   *types.UpdateConfigResponse `xml:"urn:vim25 UpdateConfigResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateConfig(r soap.RoundTripper, req *types.UpdateConfig) (*types.UpdateConfigResponse, Error) {
	var body = UpdateConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateConsoleIpRouteConfigBody struct {
	Req   *types.UpdateConsoleIpRouteConfig         `xml:"urn:vim25 UpdateConsoleIpRouteConfig,omitempty"`
	Res   *types.UpdateConsoleIpRouteConfigResponse `xml:"urn:vim25 UpdateConsoleIpRouteConfigResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateConsoleIpRouteConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateConsoleIpRouteConfig(r soap.RoundTripper, req *types.UpdateConsoleIpRouteConfig) (*types.UpdateConsoleIpRouteConfigResponse, Error) {
	var body = UpdateConsoleIpRouteConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateCounterLevelMappingBody struct {
	Req   *types.UpdateCounterLevelMapping         `xml:"urn:vim25 UpdateCounterLevelMapping,omitempty"`
	Res   *types.UpdateCounterLevelMappingResponse `xml:"urn:vim25 UpdateCounterLevelMappingResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateCounterLevelMappingBody) fault() *soap.Fault { return b.Fault }

func UpdateCounterLevelMapping(r soap.RoundTripper, req *types.UpdateCounterLevelMapping) (*types.UpdateCounterLevelMappingResponse, Error) {
	var body = UpdateCounterLevelMappingBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDVSHealthCheckConfig_TaskBody struct {
	Req   *types.UpdateDVSHealthCheckConfig_Task         `xml:"urn:vim25 UpdateDVSHealthCheckConfig_Task,omitempty"`
	Res   *types.UpdateDVSHealthCheckConfig_TaskResponse `xml:"urn:vim25 UpdateDVSHealthCheckConfig_TaskResponse,omitempty"`
	Fault *soap.Fault                                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDVSHealthCheckConfig_TaskBody) fault() *soap.Fault { return b.Fault }

func UpdateDVSHealthCheckConfig_Task(r soap.RoundTripper, req *types.UpdateDVSHealthCheckConfig_Task) (*types.UpdateDVSHealthCheckConfig_TaskResponse, Error) {
	var body = UpdateDVSHealthCheckConfig_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDVSLacpGroupConfig_TaskBody struct {
	Req   *types.UpdateDVSLacpGroupConfig_Task         `xml:"urn:vim25 UpdateDVSLacpGroupConfig_Task,omitempty"`
	Res   *types.UpdateDVSLacpGroupConfig_TaskResponse `xml:"urn:vim25 UpdateDVSLacpGroupConfig_TaskResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDVSLacpGroupConfig_TaskBody) fault() *soap.Fault { return b.Fault }

func UpdateDVSLacpGroupConfig_Task(r soap.RoundTripper, req *types.UpdateDVSLacpGroupConfig_Task) (*types.UpdateDVSLacpGroupConfig_TaskResponse, Error) {
	var body = UpdateDVSLacpGroupConfig_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDateTimeBody struct {
	Req   *types.UpdateDateTime         `xml:"urn:vim25 UpdateDateTime,omitempty"`
	Res   *types.UpdateDateTimeResponse `xml:"urn:vim25 UpdateDateTimeResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDateTimeBody) fault() *soap.Fault { return b.Fault }

func UpdateDateTime(r soap.RoundTripper, req *types.UpdateDateTime) (*types.UpdateDateTimeResponse, Error) {
	var body = UpdateDateTimeBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDateTimeConfigBody struct {
	Req   *types.UpdateDateTimeConfig         `xml:"urn:vim25 UpdateDateTimeConfig,omitempty"`
	Res   *types.UpdateDateTimeConfigResponse `xml:"urn:vim25 UpdateDateTimeConfigResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDateTimeConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateDateTimeConfig(r soap.RoundTripper, req *types.UpdateDateTimeConfig) (*types.UpdateDateTimeConfigResponse, Error) {
	var body = UpdateDateTimeConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDefaultPolicyBody struct {
	Req   *types.UpdateDefaultPolicy         `xml:"urn:vim25 UpdateDefaultPolicy,omitempty"`
	Res   *types.UpdateDefaultPolicyResponse `xml:"urn:vim25 UpdateDefaultPolicyResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDefaultPolicyBody) fault() *soap.Fault { return b.Fault }

func UpdateDefaultPolicy(r soap.RoundTripper, req *types.UpdateDefaultPolicy) (*types.UpdateDefaultPolicyResponse, Error) {
	var body = UpdateDefaultPolicyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDiskPartitionsBody struct {
	Req   *types.UpdateDiskPartitions         `xml:"urn:vim25 UpdateDiskPartitions,omitempty"`
	Res   *types.UpdateDiskPartitionsResponse `xml:"urn:vim25 UpdateDiskPartitionsResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDiskPartitionsBody) fault() *soap.Fault { return b.Fault }

func UpdateDiskPartitions(r soap.RoundTripper, req *types.UpdateDiskPartitions) (*types.UpdateDiskPartitionsResponse, Error) {
	var body = UpdateDiskPartitionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDnsConfigBody struct {
	Req   *types.UpdateDnsConfig         `xml:"urn:vim25 UpdateDnsConfig,omitempty"`
	Res   *types.UpdateDnsConfigResponse `xml:"urn:vim25 UpdateDnsConfigResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDnsConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateDnsConfig(r soap.RoundTripper, req *types.UpdateDnsConfig) (*types.UpdateDnsConfigResponse, Error) {
	var body = UpdateDnsConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateDvsCapabilityBody struct {
	Req   *types.UpdateDvsCapability         `xml:"urn:vim25 UpdateDvsCapability,omitempty"`
	Res   *types.UpdateDvsCapabilityResponse `xml:"urn:vim25 UpdateDvsCapabilityResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateDvsCapabilityBody) fault() *soap.Fault { return b.Fault }

func UpdateDvsCapability(r soap.RoundTripper, req *types.UpdateDvsCapability) (*types.UpdateDvsCapabilityResponse, Error) {
	var body = UpdateDvsCapabilityBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateExtensionBody struct {
	Req   *types.UpdateExtension         `xml:"urn:vim25 UpdateExtension,omitempty"`
	Res   *types.UpdateExtensionResponse `xml:"urn:vim25 UpdateExtensionResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateExtensionBody) fault() *soap.Fault { return b.Fault }

func UpdateExtension(r soap.RoundTripper, req *types.UpdateExtension) (*types.UpdateExtensionResponse, Error) {
	var body = UpdateExtensionBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateFlagsBody struct {
	Req   *types.UpdateFlags         `xml:"urn:vim25 UpdateFlags,omitempty"`
	Res   *types.UpdateFlagsResponse `xml:"urn:vim25 UpdateFlagsResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateFlagsBody) fault() *soap.Fault { return b.Fault }

func UpdateFlags(r soap.RoundTripper, req *types.UpdateFlags) (*types.UpdateFlagsResponse, Error) {
	var body = UpdateFlagsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateHostImageAcceptanceLevelBody struct {
	Req   *types.UpdateHostImageAcceptanceLevel         `xml:"urn:vim25 UpdateHostImageAcceptanceLevel,omitempty"`
	Res   *types.UpdateHostImageAcceptanceLevelResponse `xml:"urn:vim25 UpdateHostImageAcceptanceLevelResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostImageAcceptanceLevelBody) fault() *soap.Fault { return b.Fault }

func UpdateHostImageAcceptanceLevel(r soap.RoundTripper, req *types.UpdateHostImageAcceptanceLevel) (*types.UpdateHostImageAcceptanceLevelResponse, Error) {
	var body = UpdateHostImageAcceptanceLevelBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateHostProfileBody struct {
	Req   *types.UpdateHostProfile         `xml:"urn:vim25 UpdateHostProfile,omitempty"`
	Res   *types.UpdateHostProfileResponse `xml:"urn:vim25 UpdateHostProfileResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateHostProfileBody) fault() *soap.Fault { return b.Fault }

func UpdateHostProfile(r soap.RoundTripper, req *types.UpdateHostProfile) (*types.UpdateHostProfileResponse, Error) {
	var body = UpdateHostProfileBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiAdvancedOptionsBody struct {
	Req   *types.UpdateInternetScsiAdvancedOptions         `xml:"urn:vim25 UpdateInternetScsiAdvancedOptions,omitempty"`
	Res   *types.UpdateInternetScsiAdvancedOptionsResponse `xml:"urn:vim25 UpdateInternetScsiAdvancedOptionsResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAdvancedOptionsBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiAdvancedOptions(r soap.RoundTripper, req *types.UpdateInternetScsiAdvancedOptions) (*types.UpdateInternetScsiAdvancedOptionsResponse, Error) {
	var body = UpdateInternetScsiAdvancedOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiAliasBody struct {
	Req   *types.UpdateInternetScsiAlias         `xml:"urn:vim25 UpdateInternetScsiAlias,omitempty"`
	Res   *types.UpdateInternetScsiAliasResponse `xml:"urn:vim25 UpdateInternetScsiAliasResponse,omitempty"`
	Fault *soap.Fault                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAliasBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiAlias(r soap.RoundTripper, req *types.UpdateInternetScsiAlias) (*types.UpdateInternetScsiAliasResponse, Error) {
	var body = UpdateInternetScsiAliasBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiAuthenticationPropertiesBody struct {
	Req   *types.UpdateInternetScsiAuthenticationProperties         `xml:"urn:vim25 UpdateInternetScsiAuthenticationProperties,omitempty"`
	Res   *types.UpdateInternetScsiAuthenticationPropertiesResponse `xml:"urn:vim25 UpdateInternetScsiAuthenticationPropertiesResponse,omitempty"`
	Fault *soap.Fault                                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiAuthenticationPropertiesBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiAuthenticationProperties(r soap.RoundTripper, req *types.UpdateInternetScsiAuthenticationProperties) (*types.UpdateInternetScsiAuthenticationPropertiesResponse, Error) {
	var body = UpdateInternetScsiAuthenticationPropertiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiDigestPropertiesBody struct {
	Req   *types.UpdateInternetScsiDigestProperties         `xml:"urn:vim25 UpdateInternetScsiDigestProperties,omitempty"`
	Res   *types.UpdateInternetScsiDigestPropertiesResponse `xml:"urn:vim25 UpdateInternetScsiDigestPropertiesResponse,omitempty"`
	Fault *soap.Fault                                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiDigestPropertiesBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiDigestProperties(r soap.RoundTripper, req *types.UpdateInternetScsiDigestProperties) (*types.UpdateInternetScsiDigestPropertiesResponse, Error) {
	var body = UpdateInternetScsiDigestPropertiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiDiscoveryPropertiesBody struct {
	Req   *types.UpdateInternetScsiDiscoveryProperties         `xml:"urn:vim25 UpdateInternetScsiDiscoveryProperties,omitempty"`
	Res   *types.UpdateInternetScsiDiscoveryPropertiesResponse `xml:"urn:vim25 UpdateInternetScsiDiscoveryPropertiesResponse,omitempty"`
	Fault *soap.Fault                                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiDiscoveryPropertiesBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiDiscoveryProperties(r soap.RoundTripper, req *types.UpdateInternetScsiDiscoveryProperties) (*types.UpdateInternetScsiDiscoveryPropertiesResponse, Error) {
	var body = UpdateInternetScsiDiscoveryPropertiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiIPPropertiesBody struct {
	Req   *types.UpdateInternetScsiIPProperties         `xml:"urn:vim25 UpdateInternetScsiIPProperties,omitempty"`
	Res   *types.UpdateInternetScsiIPPropertiesResponse `xml:"urn:vim25 UpdateInternetScsiIPPropertiesResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiIPPropertiesBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiIPProperties(r soap.RoundTripper, req *types.UpdateInternetScsiIPProperties) (*types.UpdateInternetScsiIPPropertiesResponse, Error) {
	var body = UpdateInternetScsiIPPropertiesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateInternetScsiNameBody struct {
	Req   *types.UpdateInternetScsiName         `xml:"urn:vim25 UpdateInternetScsiName,omitempty"`
	Res   *types.UpdateInternetScsiNameResponse `xml:"urn:vim25 UpdateInternetScsiNameResponse,omitempty"`
	Fault *soap.Fault                           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateInternetScsiNameBody) fault() *soap.Fault { return b.Fault }

func UpdateInternetScsiName(r soap.RoundTripper, req *types.UpdateInternetScsiName) (*types.UpdateInternetScsiNameResponse, Error) {
	var body = UpdateInternetScsiNameBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateIpConfigBody struct {
	Req   *types.UpdateIpConfig         `xml:"urn:vim25 UpdateIpConfig,omitempty"`
	Res   *types.UpdateIpConfigResponse `xml:"urn:vim25 UpdateIpConfigResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateIpConfig(r soap.RoundTripper, req *types.UpdateIpConfig) (*types.UpdateIpConfigResponse, Error) {
	var body = UpdateIpConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateIpPoolBody struct {
	Req   *types.UpdateIpPool         `xml:"urn:vim25 UpdateIpPool,omitempty"`
	Res   *types.UpdateIpPoolResponse `xml:"urn:vim25 UpdateIpPoolResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpPoolBody) fault() *soap.Fault { return b.Fault }

func UpdateIpPool(r soap.RoundTripper, req *types.UpdateIpPool) (*types.UpdateIpPoolResponse, Error) {
	var body = UpdateIpPoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateIpRouteConfigBody struct {
	Req   *types.UpdateIpRouteConfig         `xml:"urn:vim25 UpdateIpRouteConfig,omitempty"`
	Res   *types.UpdateIpRouteConfigResponse `xml:"urn:vim25 UpdateIpRouteConfigResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpRouteConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateIpRouteConfig(r soap.RoundTripper, req *types.UpdateIpRouteConfig) (*types.UpdateIpRouteConfigResponse, Error) {
	var body = UpdateIpRouteConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateIpRouteTableConfigBody struct {
	Req   *types.UpdateIpRouteTableConfig         `xml:"urn:vim25 UpdateIpRouteTableConfig,omitempty"`
	Res   *types.UpdateIpRouteTableConfigResponse `xml:"urn:vim25 UpdateIpRouteTableConfigResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpRouteTableConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateIpRouteTableConfig(r soap.RoundTripper, req *types.UpdateIpRouteTableConfig) (*types.UpdateIpRouteTableConfigResponse, Error) {
	var body = UpdateIpRouteTableConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateIpmiBody struct {
	Req   *types.UpdateIpmi         `xml:"urn:vim25 UpdateIpmi,omitempty"`
	Res   *types.UpdateIpmiResponse `xml:"urn:vim25 UpdateIpmiResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateIpmiBody) fault() *soap.Fault { return b.Fault }

func UpdateIpmi(r soap.RoundTripper, req *types.UpdateIpmi) (*types.UpdateIpmiResponse, Error) {
	var body = UpdateIpmiBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateLicenseBody struct {
	Req   *types.UpdateLicense         `xml:"urn:vim25 UpdateLicense,omitempty"`
	Res   *types.UpdateLicenseResponse `xml:"urn:vim25 UpdateLicenseResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLicenseBody) fault() *soap.Fault { return b.Fault }

func UpdateLicense(r soap.RoundTripper, req *types.UpdateLicense) (*types.UpdateLicenseResponse, Error) {
	var body = UpdateLicenseBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateLicenseLabelBody struct {
	Req   *types.UpdateLicenseLabel         `xml:"urn:vim25 UpdateLicenseLabel,omitempty"`
	Res   *types.UpdateLicenseLabelResponse `xml:"urn:vim25 UpdateLicenseLabelResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLicenseLabelBody) fault() *soap.Fault { return b.Fault }

func UpdateLicenseLabel(r soap.RoundTripper, req *types.UpdateLicenseLabel) (*types.UpdateLicenseLabelResponse, Error) {
	var body = UpdateLicenseLabelBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateLinkedChildrenBody struct {
	Req   *types.UpdateLinkedChildren         `xml:"urn:vim25 UpdateLinkedChildren,omitempty"`
	Res   *types.UpdateLinkedChildrenResponse `xml:"urn:vim25 UpdateLinkedChildrenResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLinkedChildrenBody) fault() *soap.Fault { return b.Fault }

func UpdateLinkedChildren(r soap.RoundTripper, req *types.UpdateLinkedChildren) (*types.UpdateLinkedChildrenResponse, Error) {
	var body = UpdateLinkedChildrenBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateLocalSwapDatastoreBody struct {
	Req   *types.UpdateLocalSwapDatastore         `xml:"urn:vim25 UpdateLocalSwapDatastore,omitempty"`
	Res   *types.UpdateLocalSwapDatastoreResponse `xml:"urn:vim25 UpdateLocalSwapDatastoreResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateLocalSwapDatastoreBody) fault() *soap.Fault { return b.Fault }

func UpdateLocalSwapDatastore(r soap.RoundTripper, req *types.UpdateLocalSwapDatastore) (*types.UpdateLocalSwapDatastoreResponse, Error) {
	var body = UpdateLocalSwapDatastoreBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateModuleOptionStringBody struct {
	Req   *types.UpdateModuleOptionString         `xml:"urn:vim25 UpdateModuleOptionString,omitempty"`
	Res   *types.UpdateModuleOptionStringResponse `xml:"urn:vim25 UpdateModuleOptionStringResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateModuleOptionStringBody) fault() *soap.Fault { return b.Fault }

func UpdateModuleOptionString(r soap.RoundTripper, req *types.UpdateModuleOptionString) (*types.UpdateModuleOptionStringResponse, Error) {
	var body = UpdateModuleOptionStringBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateNetworkConfigBody struct {
	Req   *types.UpdateNetworkConfig         `xml:"urn:vim25 UpdateNetworkConfig,omitempty"`
	Res   *types.UpdateNetworkConfigResponse `xml:"urn:vim25 UpdateNetworkConfigResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateNetworkConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateNetworkConfig(r soap.RoundTripper, req *types.UpdateNetworkConfig) (*types.UpdateNetworkConfigResponse, Error) {
	var body = UpdateNetworkConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateNetworkResourcePoolBody struct {
	Req   *types.UpdateNetworkResourcePool         `xml:"urn:vim25 UpdateNetworkResourcePool,omitempty"`
	Res   *types.UpdateNetworkResourcePoolResponse `xml:"urn:vim25 UpdateNetworkResourcePoolResponse,omitempty"`
	Fault *soap.Fault                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateNetworkResourcePoolBody) fault() *soap.Fault { return b.Fault }

func UpdateNetworkResourcePool(r soap.RoundTripper, req *types.UpdateNetworkResourcePool) (*types.UpdateNetworkResourcePoolResponse, Error) {
	var body = UpdateNetworkResourcePoolBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateOptionsBody struct {
	Req   *types.UpdateOptions         `xml:"urn:vim25 UpdateOptions,omitempty"`
	Res   *types.UpdateOptionsResponse `xml:"urn:vim25 UpdateOptionsResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateOptionsBody) fault() *soap.Fault { return b.Fault }

func UpdateOptions(r soap.RoundTripper, req *types.UpdateOptions) (*types.UpdateOptionsResponse, Error) {
	var body = UpdateOptionsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdatePassthruConfigBody struct {
	Req   *types.UpdatePassthruConfig         `xml:"urn:vim25 UpdatePassthruConfig,omitempty"`
	Res   *types.UpdatePassthruConfigResponse `xml:"urn:vim25 UpdatePassthruConfigResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePassthruConfigBody) fault() *soap.Fault { return b.Fault }

func UpdatePassthruConfig(r soap.RoundTripper, req *types.UpdatePassthruConfig) (*types.UpdatePassthruConfigResponse, Error) {
	var body = UpdatePassthruConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdatePerfIntervalBody struct {
	Req   *types.UpdatePerfInterval         `xml:"urn:vim25 UpdatePerfInterval,omitempty"`
	Res   *types.UpdatePerfIntervalResponse `xml:"urn:vim25 UpdatePerfIntervalResponse,omitempty"`
	Fault *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePerfIntervalBody) fault() *soap.Fault { return b.Fault }

func UpdatePerfInterval(r soap.RoundTripper, req *types.UpdatePerfInterval) (*types.UpdatePerfIntervalResponse, Error) {
	var body = UpdatePerfIntervalBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdatePhysicalNicLinkSpeedBody struct {
	Req   *types.UpdatePhysicalNicLinkSpeed         `xml:"urn:vim25 UpdatePhysicalNicLinkSpeed,omitempty"`
	Res   *types.UpdatePhysicalNicLinkSpeedResponse `xml:"urn:vim25 UpdatePhysicalNicLinkSpeedResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePhysicalNicLinkSpeedBody) fault() *soap.Fault { return b.Fault }

func UpdatePhysicalNicLinkSpeed(r soap.RoundTripper, req *types.UpdatePhysicalNicLinkSpeed) (*types.UpdatePhysicalNicLinkSpeedResponse, Error) {
	var body = UpdatePhysicalNicLinkSpeedBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdatePortGroupBody struct {
	Req   *types.UpdatePortGroup         `xml:"urn:vim25 UpdatePortGroup,omitempty"`
	Res   *types.UpdatePortGroupResponse `xml:"urn:vim25 UpdatePortGroupResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdatePortGroupBody) fault() *soap.Fault { return b.Fault }

func UpdatePortGroup(r soap.RoundTripper, req *types.UpdatePortGroup) (*types.UpdatePortGroupResponse, Error) {
	var body = UpdatePortGroupBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateProgressBody struct {
	Req   *types.UpdateProgress         `xml:"urn:vim25 UpdateProgress,omitempty"`
	Res   *types.UpdateProgressResponse `xml:"urn:vim25 UpdateProgressResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateProgressBody) fault() *soap.Fault { return b.Fault }

func UpdateProgress(r soap.RoundTripper, req *types.UpdateProgress) (*types.UpdateProgressResponse, Error) {
	var body = UpdateProgressBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateReferenceHostBody struct {
	Req   *types.UpdateReferenceHost         `xml:"urn:vim25 UpdateReferenceHost,omitempty"`
	Res   *types.UpdateReferenceHostResponse `xml:"urn:vim25 UpdateReferenceHostResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateReferenceHostBody) fault() *soap.Fault { return b.Fault }

func UpdateReferenceHost(r soap.RoundTripper, req *types.UpdateReferenceHost) (*types.UpdateReferenceHostResponse, Error) {
	var body = UpdateReferenceHostBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateRulesetBody struct {
	Req   *types.UpdateRuleset         `xml:"urn:vim25 UpdateRuleset,omitempty"`
	Res   *types.UpdateRulesetResponse `xml:"urn:vim25 UpdateRulesetResponse,omitempty"`
	Fault *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateRulesetBody) fault() *soap.Fault { return b.Fault }

func UpdateRuleset(r soap.RoundTripper, req *types.UpdateRuleset) (*types.UpdateRulesetResponse, Error) {
	var body = UpdateRulesetBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateScsiLunDisplayNameBody struct {
	Req   *types.UpdateScsiLunDisplayName         `xml:"urn:vim25 UpdateScsiLunDisplayName,omitempty"`
	Res   *types.UpdateScsiLunDisplayNameResponse `xml:"urn:vim25 UpdateScsiLunDisplayNameResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateScsiLunDisplayNameBody) fault() *soap.Fault { return b.Fault }

func UpdateScsiLunDisplayName(r soap.RoundTripper, req *types.UpdateScsiLunDisplayName) (*types.UpdateScsiLunDisplayNameResponse, Error) {
	var body = UpdateScsiLunDisplayNameBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateServiceConsoleVirtualNicBody struct {
	Req   *types.UpdateServiceConsoleVirtualNic         `xml:"urn:vim25 UpdateServiceConsoleVirtualNic,omitempty"`
	Res   *types.UpdateServiceConsoleVirtualNicResponse `xml:"urn:vim25 UpdateServiceConsoleVirtualNicResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServiceConsoleVirtualNicBody) fault() *soap.Fault { return b.Fault }

func UpdateServiceConsoleVirtualNic(r soap.RoundTripper, req *types.UpdateServiceConsoleVirtualNic) (*types.UpdateServiceConsoleVirtualNicResponse, Error) {
	var body = UpdateServiceConsoleVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateServiceMessageBody struct {
	Req   *types.UpdateServiceMessage         `xml:"urn:vim25 UpdateServiceMessage,omitempty"`
	Res   *types.UpdateServiceMessageResponse `xml:"urn:vim25 UpdateServiceMessageResponse,omitempty"`
	Fault *soap.Fault                         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServiceMessageBody) fault() *soap.Fault { return b.Fault }

func UpdateServiceMessage(r soap.RoundTripper, req *types.UpdateServiceMessage) (*types.UpdateServiceMessageResponse, Error) {
	var body = UpdateServiceMessageBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateServicePolicyBody struct {
	Req   *types.UpdateServicePolicy         `xml:"urn:vim25 UpdateServicePolicy,omitempty"`
	Res   *types.UpdateServicePolicyResponse `xml:"urn:vim25 UpdateServicePolicyResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateServicePolicyBody) fault() *soap.Fault { return b.Fault }

func UpdateServicePolicy(r soap.RoundTripper, req *types.UpdateServicePolicy) (*types.UpdateServicePolicyResponse, Error) {
	var body = UpdateServicePolicyBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateSoftwareInternetScsiEnabledBody struct {
	Req   *types.UpdateSoftwareInternetScsiEnabled         `xml:"urn:vim25 UpdateSoftwareInternetScsiEnabled,omitempty"`
	Res   *types.UpdateSoftwareInternetScsiEnabledResponse `xml:"urn:vim25 UpdateSoftwareInternetScsiEnabledResponse,omitempty"`
	Fault *soap.Fault                                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSoftwareInternetScsiEnabledBody) fault() *soap.Fault { return b.Fault }

func UpdateSoftwareInternetScsiEnabled(r soap.RoundTripper, req *types.UpdateSoftwareInternetScsiEnabled) (*types.UpdateSoftwareInternetScsiEnabledResponse, Error) {
	var body = UpdateSoftwareInternetScsiEnabledBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateSystemResourcesBody struct {
	Req   *types.UpdateSystemResources         `xml:"urn:vim25 UpdateSystemResources,omitempty"`
	Res   *types.UpdateSystemResourcesResponse `xml:"urn:vim25 UpdateSystemResourcesResponse,omitempty"`
	Fault *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSystemResourcesBody) fault() *soap.Fault { return b.Fault }

func UpdateSystemResources(r soap.RoundTripper, req *types.UpdateSystemResources) (*types.UpdateSystemResourcesResponse, Error) {
	var body = UpdateSystemResourcesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateSystemSwapConfigurationBody struct {
	Req   *types.UpdateSystemSwapConfiguration         `xml:"urn:vim25 UpdateSystemSwapConfiguration,omitempty"`
	Res   *types.UpdateSystemSwapConfigurationResponse `xml:"urn:vim25 UpdateSystemSwapConfigurationResponse,omitempty"`
	Fault *soap.Fault                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateSystemSwapConfigurationBody) fault() *soap.Fault { return b.Fault }

func UpdateSystemSwapConfiguration(r soap.RoundTripper, req *types.UpdateSystemSwapConfiguration) (*types.UpdateSystemSwapConfigurationResponse, Error) {
	var body = UpdateSystemSwapConfigurationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateUserBody struct {
	Req   *types.UpdateUser         `xml:"urn:vim25 UpdateUser,omitempty"`
	Res   *types.UpdateUserResponse `xml:"urn:vim25 UpdateUserResponse,omitempty"`
	Fault *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateUserBody) fault() *soap.Fault { return b.Fault }

func UpdateUser(r soap.RoundTripper, req *types.UpdateUser) (*types.UpdateUserResponse, Error) {
	var body = UpdateUserBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateVAppConfigBody struct {
	Req   *types.UpdateVAppConfig         `xml:"urn:vim25 UpdateVAppConfig,omitempty"`
	Res   *types.UpdateVAppConfigResponse `xml:"urn:vim25 UpdateVAppConfigResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVAppConfigBody) fault() *soap.Fault { return b.Fault }

func UpdateVAppConfig(r soap.RoundTripper, req *types.UpdateVAppConfig) (*types.UpdateVAppConfigResponse, Error) {
	var body = UpdateVAppConfigBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateVirtualMachineFiles_TaskBody struct {
	Req   *types.UpdateVirtualMachineFiles_Task         `xml:"urn:vim25 UpdateVirtualMachineFiles_Task,omitempty"`
	Res   *types.UpdateVirtualMachineFiles_TaskResponse `xml:"urn:vim25 UpdateVirtualMachineFiles_TaskResponse,omitempty"`
	Fault *soap.Fault                                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualMachineFiles_TaskBody) fault() *soap.Fault { return b.Fault }

func UpdateVirtualMachineFiles_Task(r soap.RoundTripper, req *types.UpdateVirtualMachineFiles_Task) (*types.UpdateVirtualMachineFiles_TaskResponse, Error) {
	var body = UpdateVirtualMachineFiles_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateVirtualNicBody struct {
	Req   *types.UpdateVirtualNic         `xml:"urn:vim25 UpdateVirtualNic,omitempty"`
	Res   *types.UpdateVirtualNicResponse `xml:"urn:vim25 UpdateVirtualNicResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualNicBody) fault() *soap.Fault { return b.Fault }

func UpdateVirtualNic(r soap.RoundTripper, req *types.UpdateVirtualNic) (*types.UpdateVirtualNicResponse, Error) {
	var body = UpdateVirtualNicBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateVirtualSwitchBody struct {
	Req   *types.UpdateVirtualSwitch         `xml:"urn:vim25 UpdateVirtualSwitch,omitempty"`
	Res   *types.UpdateVirtualSwitchResponse `xml:"urn:vim25 UpdateVirtualSwitchResponse,omitempty"`
	Fault *soap.Fault                        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVirtualSwitchBody) fault() *soap.Fault { return b.Fault }

func UpdateVirtualSwitch(r soap.RoundTripper, req *types.UpdateVirtualSwitch) (*types.UpdateVirtualSwitchResponse, Error) {
	var body = UpdateVirtualSwitchBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpdateVsan_TaskBody struct {
	Req   *types.UpdateVsan_Task         `xml:"urn:vim25 UpdateVsan_Task,omitempty"`
	Res   *types.UpdateVsan_TaskResponse `xml:"urn:vim25 UpdateVsan_TaskResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateVsan_TaskBody) fault() *soap.Fault { return b.Fault }

func UpdateVsan_Task(r soap.RoundTripper, req *types.UpdateVsan_Task) (*types.UpdateVsan_TaskResponse, Error) {
	var body = UpdateVsan_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpgradeTools_TaskBody struct {
	Req   *types.UpgradeTools_Task         `xml:"urn:vim25 UpgradeTools_Task,omitempty"`
	Res   *types.UpgradeTools_TaskResponse `xml:"urn:vim25 UpgradeTools_TaskResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeTools_TaskBody) fault() *soap.Fault { return b.Fault }

func UpgradeTools_Task(r soap.RoundTripper, req *types.UpgradeTools_Task) (*types.UpgradeTools_TaskResponse, Error) {
	var body = UpgradeTools_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpgradeVM_TaskBody struct {
	Req   *types.UpgradeVM_Task         `xml:"urn:vim25 UpgradeVM_Task,omitempty"`
	Res   *types.UpgradeVM_TaskResponse `xml:"urn:vim25 UpgradeVM_TaskResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVM_TaskBody) fault() *soap.Fault { return b.Fault }

func UpgradeVM_Task(r soap.RoundTripper, req *types.UpgradeVM_Task) (*types.UpgradeVM_TaskResponse, Error) {
	var body = UpgradeVM_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpgradeVmLayoutBody struct {
	Req   *types.UpgradeVmLayout         `xml:"urn:vim25 UpgradeVmLayout,omitempty"`
	Res   *types.UpgradeVmLayoutResponse `xml:"urn:vim25 UpgradeVmLayoutResponse,omitempty"`
	Fault *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVmLayoutBody) fault() *soap.Fault { return b.Fault }

func UpgradeVmLayout(r soap.RoundTripper, req *types.UpgradeVmLayout) (*types.UpgradeVmLayoutResponse, Error) {
	var body = UpgradeVmLayoutBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type UpgradeVmfsBody struct {
	Req   *types.UpgradeVmfs         `xml:"urn:vim25 UpgradeVmfs,omitempty"`
	Res   *types.UpgradeVmfsResponse `xml:"urn:vim25 UpgradeVmfsResponse,omitempty"`
	Fault *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpgradeVmfsBody) fault() *soap.Fault { return b.Fault }

func UpgradeVmfs(r soap.RoundTripper, req *types.UpgradeVmfs) (*types.UpgradeVmfsResponse, Error) {
	var body = UpgradeVmfsBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ValidateCredentialsInGuestBody struct {
	Req   *types.ValidateCredentialsInGuest         `xml:"urn:vim25 ValidateCredentialsInGuest,omitempty"`
	Res   *types.ValidateCredentialsInGuestResponse `xml:"urn:vim25 ValidateCredentialsInGuestResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateCredentialsInGuestBody) fault() *soap.Fault { return b.Fault }

func ValidateCredentialsInGuest(r soap.RoundTripper, req *types.ValidateCredentialsInGuest) (*types.ValidateCredentialsInGuestResponse, Error) {
	var body = ValidateCredentialsInGuestBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ValidateHostBody struct {
	Req   *types.ValidateHost         `xml:"urn:vim25 ValidateHost,omitempty"`
	Res   *types.ValidateHostResponse `xml:"urn:vim25 ValidateHostResponse,omitempty"`
	Fault *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateHostBody) fault() *soap.Fault { return b.Fault }

func ValidateHost(r soap.RoundTripper, req *types.ValidateHost) (*types.ValidateHostResponse, Error) {
	var body = ValidateHostBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ValidateMigrationBody struct {
	Req   *types.ValidateMigration         `xml:"urn:vim25 ValidateMigration,omitempty"`
	Res   *types.ValidateMigrationResponse `xml:"urn:vim25 ValidateMigrationResponse,omitempty"`
	Fault *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ValidateMigrationBody) fault() *soap.Fault { return b.Fault }

func ValidateMigration(r soap.RoundTripper, req *types.ValidateMigration) (*types.ValidateMigrationResponse, Error) {
	var body = ValidateMigrationBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type WaitForUpdatesBody struct {
	Req   *types.WaitForUpdates         `xml:"urn:vim25 WaitForUpdates,omitempty"`
	Res   *types.WaitForUpdatesResponse `xml:"urn:vim25 WaitForUpdatesResponse,omitempty"`
	Fault *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *WaitForUpdatesBody) fault() *soap.Fault { return b.Fault }

func WaitForUpdates(r soap.RoundTripper, req *types.WaitForUpdates) (*types.WaitForUpdatesResponse, Error) {
	var body = WaitForUpdatesBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type WaitForUpdatesExBody struct {
	Req   *types.WaitForUpdatesEx         `xml:"urn:vim25 WaitForUpdatesEx,omitempty"`
	Res   *types.WaitForUpdatesExResponse `xml:"urn:vim25 WaitForUpdatesExResponse,omitempty"`
	Fault *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *WaitForUpdatesExBody) fault() *soap.Fault { return b.Fault }

func WaitForUpdatesEx(r soap.RoundTripper, req *types.WaitForUpdatesEx) (*types.WaitForUpdatesExResponse, Error) {
	var body = WaitForUpdatesExBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type XmlToCustomizationSpecItemBody struct {
	Req   *types.XmlToCustomizationSpecItem         `xml:"urn:vim25 XmlToCustomizationSpecItem,omitempty"`
	Res   *types.XmlToCustomizationSpecItemResponse `xml:"urn:vim25 XmlToCustomizationSpecItemResponse,omitempty"`
	Fault *soap.Fault                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *XmlToCustomizationSpecItemBody) fault() *soap.Fault { return b.Fault }

func XmlToCustomizationSpecItem(r soap.RoundTripper, req *types.XmlToCustomizationSpecItem) (*types.XmlToCustomizationSpecItemResponse, Error) {
	var body = XmlToCustomizationSpecItemBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}

type ZeroFillVirtualDisk_TaskBody struct {
	Req   *types.ZeroFillVirtualDisk_Task         `xml:"urn:vim25 ZeroFillVirtualDisk_Task,omitempty"`
	Res   *types.ZeroFillVirtualDisk_TaskResponse `xml:"urn:vim25 ZeroFillVirtualDisk_TaskResponse,omitempty"`
	Fault *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ZeroFillVirtualDisk_TaskBody) fault() *soap.Fault { return b.Fault }

func ZeroFillVirtualDisk_Task(r soap.RoundTripper, req *types.ZeroFillVirtualDisk_Task) (*types.ZeroFillVirtualDisk_TaskResponse, Error) {
	var body = ZeroFillVirtualDisk_TaskBody{
		Req: req,
	}

	if err := roundTrip(r, &body); err != nil {
		return nil, err
	}

	return body.Res, nil
}
