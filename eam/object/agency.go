/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package object

import (
	"context"

	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type Agency struct {
	EamObject
}

// NewAgency returns a wrapper for an Agency managed object.
func NewAgency(c *eam.Client, ref vim.ManagedObjectReference) Agency {
	return Agency{
		EamObject: EamObject{
			c: c,
			r: ref,
		},
	}
}

func (m Agency) Agents(ctx context.Context) ([]Agent, error) {
	resp, err := methods.QueryAgent(ctx, m.c, &types.QueryAgent{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	objs := make([]Agent, len(resp.Returnval))
	for i := range resp.Returnval {
		objs[i].c = m.c
		objs[i].r = resp.Returnval[i]
	}
	return objs, nil
}

func (m Agency) Config(ctx context.Context) (*types.AgencyConfigInfo, error) {
	resp, err := methods.QueryConfig(ctx, m.c, &types.QueryConfig{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	return resp.Returnval.GetAgencyConfigInfo(), nil
}

func (m Agency) Runtime(ctx context.Context) (*types.EamObjectRuntimeInfo, error) {
	resp, err := methods.AgencyQueryRuntime(ctx, m.c, &types.AgencyQueryRuntime{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	return resp.Returnval.GetEamObjectRuntimeInfo(), nil
}

func (m Agency) SolutionId(ctx context.Context) (string, error) {
	resp, err := methods.QuerySolutionId(ctx, m.c, &types.QuerySolutionId{
		This: m.r,
	})
	if err != nil {
		return "", err
	}
	return resp.Returnval, nil
}

func (m Agency) Destroy(ctx context.Context) error {
	_, err := methods.DestroyAgency(ctx, m.c, &types.DestroyAgency{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Agency) Disable(ctx context.Context) error {
	_, err := methods.Disable(ctx, m.c, &types.Disable{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Agency) Enable(ctx context.Context) error {
	_, err := methods.Enable(ctx, m.c, &types.Enable{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Agency) RegisterAgentVm(
	ctx context.Context,
	agentVmMoRef vim.ManagedObjectReference) (*Agent, error) {

	resp, err := methods.RegisterAgentVm(ctx, m.c, &types.RegisterAgentVm{
		This:    m.r,
		AgentVm: agentVmMoRef,
	})
	if err != nil {
		return nil, err
	}
	return NewAgent(m.c, resp.Returnval), nil
}

func (m Agency) Uninstall(ctx context.Context) error {
	_, err := methods.Uninstall(ctx, m.c, &types.Uninstall{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Agency) UnregisterAgentVm(
	ctx context.Context,
	agentVmMoRef vim.ManagedObjectReference) error {

	_, err := methods.UnregisterAgentVm(ctx, m.c, &types.UnregisterAgentVm{
		This:    m.r,
		AgentVm: agentVmMoRef,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Agency) Update(
	ctx context.Context,
	config types.AgencyConfigInfo) error {

	_, err := methods.Update(ctx, m.c, &types.Update{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}
