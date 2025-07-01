// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package methods

import (
	"context"

	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/vim25/soap"
)

type AddIssueBody struct {
	Req    *types.AddIssue         `xml:"urn:eam AddIssue,omitempty"`
	Res    *types.AddIssueResponse `xml:"urn:eam AddIssueResponse,omitempty"`
	Fault_ *soap.Fault             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AddIssueBody) Fault() *soap.Fault { return b.Fault_ }

func AddIssue(ctx context.Context, r soap.RoundTripper, req *types.AddIssue) (*types.AddIssueResponse, error) {
	var reqBody, resBody AddIssueBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AgencyQueryRuntimeBody struct {
	Req    *types.AgencyQueryRuntime         `xml:"urn:eam AgencyQueryRuntime,omitempty"`
	Res    *types.AgencyQueryRuntimeResponse `xml:"urn:eam AgencyQueryRuntimeResponse,omitempty"`
	Fault_ *soap.Fault                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AgencyQueryRuntimeBody) Fault() *soap.Fault { return b.Fault_ }

func AgencyQueryRuntime(ctx context.Context, r soap.RoundTripper, req *types.AgencyQueryRuntime) (*types.AgencyQueryRuntimeResponse, error) {
	var reqBody, resBody AgencyQueryRuntimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type Agency_DisableBody struct {
	Req    *types.Agency_Disable         `xml:"urn:eam Agency_Disable,omitempty"`
	Res    *types.Agency_DisableResponse `xml:"urn:eam Agency_DisableResponse,omitempty"`
	Fault_ *soap.Fault                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Agency_DisableBody) Fault() *soap.Fault { return b.Fault_ }

func Agency_Disable(ctx context.Context, r soap.RoundTripper, req *types.Agency_Disable) (*types.Agency_DisableResponse, error) {
	var reqBody, resBody Agency_DisableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type Agency_EnableBody struct {
	Req    *types.Agency_Enable         `xml:"urn:eam Agency_Enable,omitempty"`
	Res    *types.Agency_EnableResponse `xml:"urn:eam Agency_EnableResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *Agency_EnableBody) Fault() *soap.Fault { return b.Fault_ }

func Agency_Enable(ctx context.Context, r soap.RoundTripper, req *types.Agency_Enable) (*types.Agency_EnableResponse, error) {
	var reqBody, resBody Agency_EnableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AgentQueryConfigBody struct {
	Req    *types.AgentQueryConfig         `xml:"urn:eam AgentQueryConfig,omitempty"`
	Res    *types.AgentQueryConfigResponse `xml:"urn:eam AgentQueryConfigResponse,omitempty"`
	Fault_ *soap.Fault                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AgentQueryConfigBody) Fault() *soap.Fault { return b.Fault_ }

func AgentQueryConfig(ctx context.Context, r soap.RoundTripper, req *types.AgentQueryConfig) (*types.AgentQueryConfigResponse, error) {
	var reqBody, resBody AgentQueryConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type AgentQueryRuntimeBody struct {
	Req    *types.AgentQueryRuntime         `xml:"urn:eam AgentQueryRuntime,omitempty"`
	Res    *types.AgentQueryRuntimeResponse `xml:"urn:eam AgentQueryRuntimeResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AgentQueryRuntimeBody) Fault() *soap.Fault { return b.Fault_ }

func AgentQueryRuntime(ctx context.Context, r soap.RoundTripper, req *types.AgentQueryRuntime) (*types.AgentQueryRuntimeResponse, error) {
	var reqBody, resBody AgentQueryRuntimeBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type CreateAgencyBody struct {
	Req    *types.CreateAgency         `xml:"urn:eam CreateAgency,omitempty"`
	Res    *types.CreateAgencyResponse `xml:"urn:eam CreateAgencyResponse,omitempty"`
	Fault_ *soap.Fault                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *CreateAgencyBody) Fault() *soap.Fault { return b.Fault_ }

func CreateAgency(ctx context.Context, r soap.RoundTripper, req *types.CreateAgency) (*types.CreateAgencyResponse, error) {
	var reqBody, resBody CreateAgencyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type DestroyAgencyBody struct {
	Req    *types.DestroyAgency         `xml:"urn:eam DestroyAgency,omitempty"`
	Res    *types.DestroyAgencyResponse `xml:"urn:eam DestroyAgencyResponse,omitempty"`
	Fault_ *soap.Fault                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *DestroyAgencyBody) Fault() *soap.Fault { return b.Fault_ }

func DestroyAgency(ctx context.Context, r soap.RoundTripper, req *types.DestroyAgency) (*types.DestroyAgencyResponse, error) {
	var reqBody, resBody DestroyAgencyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type GetMaintenanceModePolicyBody struct {
	Req    *types.GetMaintenanceModePolicy         `xml:"urn:eam GetMaintenanceModePolicy,omitempty"`
	Res    *types.GetMaintenanceModePolicyResponse `xml:"urn:eam GetMaintenanceModePolicyResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *GetMaintenanceModePolicyBody) Fault() *soap.Fault { return b.Fault_ }

func GetMaintenanceModePolicy(ctx context.Context, r soap.RoundTripper, req *types.GetMaintenanceModePolicy) (*types.GetMaintenanceModePolicyResponse, error) {
	var reqBody, resBody GetMaintenanceModePolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type MarkAsAvailableBody struct {
	Req    *types.MarkAsAvailable         `xml:"urn:eam MarkAsAvailable,omitempty"`
	Res    *types.MarkAsAvailableResponse `xml:"urn:eam MarkAsAvailableResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *MarkAsAvailableBody) Fault() *soap.Fault { return b.Fault_ }

func MarkAsAvailable(ctx context.Context, r soap.RoundTripper, req *types.MarkAsAvailable) (*types.MarkAsAvailableResponse, error) {
	var reqBody, resBody MarkAsAvailableBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAgencyBody struct {
	Req    *types.QueryAgency         `xml:"urn:eam QueryAgency,omitempty"`
	Res    *types.QueryAgencyResponse `xml:"urn:eam QueryAgencyResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAgencyBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAgency(ctx context.Context, r soap.RoundTripper, req *types.QueryAgency) (*types.QueryAgencyResponse, error) {
	var reqBody, resBody QueryAgencyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryAgentBody struct {
	Req    *types.QueryAgent         `xml:"urn:eam QueryAgent,omitempty"`
	Res    *types.QueryAgentResponse `xml:"urn:eam QueryAgentResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryAgentBody) Fault() *soap.Fault { return b.Fault_ }

func QueryAgent(ctx context.Context, r soap.RoundTripper, req *types.QueryAgent) (*types.QueryAgentResponse, error) {
	var reqBody, resBody QueryAgentBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryConfigBody struct {
	Req    *types.QueryConfig         `xml:"urn:eam QueryConfig,omitempty"`
	Res    *types.QueryConfigResponse `xml:"urn:eam QueryConfigResponse,omitempty"`
	Fault_ *soap.Fault                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryConfigBody) Fault() *soap.Fault { return b.Fault_ }

func QueryConfig(ctx context.Context, r soap.RoundTripper, req *types.QueryConfig) (*types.QueryConfigResponse, error) {
	var reqBody, resBody QueryConfigBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QueryIssueBody struct {
	Req    *types.QueryIssue         `xml:"urn:eam QueryIssue,omitempty"`
	Res    *types.QueryIssueResponse `xml:"urn:eam QueryIssueResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QueryIssueBody) Fault() *soap.Fault { return b.Fault_ }

func QueryIssue(ctx context.Context, r soap.RoundTripper, req *types.QueryIssue) (*types.QueryIssueResponse, error) {
	var reqBody, resBody QueryIssueBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type QuerySolutionIdBody struct {
	Req    *types.QuerySolutionId         `xml:"urn:eam QuerySolutionId,omitempty"`
	Res    *types.QuerySolutionIdResponse `xml:"urn:eam QuerySolutionIdResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *QuerySolutionIdBody) Fault() *soap.Fault { return b.Fault_ }

func QuerySolutionId(ctx context.Context, r soap.RoundTripper, req *types.QuerySolutionId) (*types.QuerySolutionIdResponse, error) {
	var reqBody, resBody QuerySolutionIdBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type RegisterAgentVmBody struct {
	Req    *types.RegisterAgentVm         `xml:"urn:eam RegisterAgentVm,omitempty"`
	Res    *types.RegisterAgentVmResponse `xml:"urn:eam RegisterAgentVmResponse,omitempty"`
	Fault_ *soap.Fault                    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RegisterAgentVmBody) Fault() *soap.Fault { return b.Fault_ }

func RegisterAgentVm(ctx context.Context, r soap.RoundTripper, req *types.RegisterAgentVm) (*types.RegisterAgentVmResponse, error) {
	var reqBody, resBody RegisterAgentVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveBody struct {
	Req    *types.Resolve         `xml:"urn:eam Resolve,omitempty"`
	Res    *types.ResolveResponse `xml:"urn:eam ResolveResponse,omitempty"`
	Fault_ *soap.Fault            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveBody) Fault() *soap.Fault { return b.Fault_ }

func Resolve(ctx context.Context, r soap.RoundTripper, req *types.Resolve) (*types.ResolveResponse, error) {
	var reqBody, resBody ResolveBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ResolveAllBody struct {
	Req    *types.ResolveAll         `xml:"urn:eam ResolveAll,omitempty"`
	Res    *types.ResolveAllResponse `xml:"urn:eam ResolveAllResponse,omitempty"`
	Fault_ *soap.Fault               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ResolveAllBody) Fault() *soap.Fault { return b.Fault_ }

func ResolveAll(ctx context.Context, r soap.RoundTripper, req *types.ResolveAll) (*types.ResolveAllResponse, error) {
	var reqBody, resBody ResolveAllBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type ScanForUnknownAgentVmBody struct {
	Req    *types.ScanForUnknownAgentVm         `xml:"urn:eam ScanForUnknownAgentVm,omitempty"`
	Res    *types.ScanForUnknownAgentVmResponse `xml:"urn:eam ScanForUnknownAgentVmResponse,omitempty"`
	Fault_ *soap.Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *ScanForUnknownAgentVmBody) Fault() *soap.Fault { return b.Fault_ }

func ScanForUnknownAgentVm(ctx context.Context, r soap.RoundTripper, req *types.ScanForUnknownAgentVm) (*types.ScanForUnknownAgentVmResponse, error) {
	var reqBody, resBody ScanForUnknownAgentVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type SetMaintenanceModePolicyBody struct {
	Req    *types.SetMaintenanceModePolicy         `xml:"urn:eam SetMaintenanceModePolicy,omitempty"`
	Res    *types.SetMaintenanceModePolicyResponse `xml:"urn:eam SetMaintenanceModePolicyResponse,omitempty"`
	Fault_ *soap.Fault                             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *SetMaintenanceModePolicyBody) Fault() *soap.Fault { return b.Fault_ }

func SetMaintenanceModePolicy(ctx context.Context, r soap.RoundTripper, req *types.SetMaintenanceModePolicy) (*types.SetMaintenanceModePolicyResponse, error) {
	var reqBody, resBody SetMaintenanceModePolicyBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UninstallBody struct {
	Req    *types.Uninstall         `xml:"urn:eam Uninstall,omitempty"`
	Res    *types.UninstallResponse `xml:"urn:eam UninstallResponse,omitempty"`
	Fault_ *soap.Fault              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UninstallBody) Fault() *soap.Fault { return b.Fault_ }

func Uninstall(ctx context.Context, r soap.RoundTripper, req *types.Uninstall) (*types.UninstallResponse, error) {
	var reqBody, resBody UninstallBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UnregisterAgentVmBody struct {
	Req    *types.UnregisterAgentVm         `xml:"urn:eam UnregisterAgentVm,omitempty"`
	Res    *types.UnregisterAgentVmResponse `xml:"urn:eam UnregisterAgentVmResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UnregisterAgentVmBody) Fault() *soap.Fault { return b.Fault_ }

func UnregisterAgentVm(ctx context.Context, r soap.RoundTripper, req *types.UnregisterAgentVm) (*types.UnregisterAgentVmResponse, error) {
	var reqBody, resBody UnregisterAgentVmBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}

type UpdateBody struct {
	Req    *types.Update         `xml:"urn:eam Update,omitempty"`
	Res    *types.UpdateResponse `xml:"urn:eam UpdateResponse,omitempty"`
	Fault_ *soap.Fault           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *UpdateBody) Fault() *soap.Fault { return b.Fault_ }

func Update(ctx context.Context, r soap.RoundTripper, req *types.Update) (*types.UpdateResponse, error) {
	var reqBody, resBody UpdateBody

	reqBody.Req = req

	if err := r.RoundTrip(ctx, &reqBody, &resBody); err != nil {
		return nil, err
	}

	return resBody.Res, nil
}
