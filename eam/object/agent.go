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
