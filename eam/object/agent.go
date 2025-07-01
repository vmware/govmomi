// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"

	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/methods"
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type Agent struct {
	EamObject
}

// NewAgent returns a wrapper for an Agent managed object.
func NewAgent(c *eam.Client, ref vim.ManagedObjectReference) *Agent {
	return &Agent{
		EamObject: EamObject{
			c: c,
			r: ref,
		},
	}
}

func (m Agent) Config(ctx context.Context) (*types.AgentConfigInfo, error) {
	resp, err := methods.AgentQueryConfig(ctx, m.c, &types.AgentQueryConfig{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Returnval, nil
}

func (m Agent) Runtime(ctx context.Context) (*types.AgentRuntimeInfo, error) {
	resp, err := methods.AgentQueryRuntime(ctx, m.c, &types.AgentQueryRuntime{
		This: m.r,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Returnval, nil
}

func (m Agent) MarkAsAvailable(ctx context.Context) error {
	_, err := methods.MarkAsAvailable(ctx, m.c, &types.MarkAsAvailable{
		This: m.r,
	})
	if err != nil {
		return err
	}
	return nil
}
