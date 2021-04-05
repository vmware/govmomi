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

package simulator

import (
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New())
		}
	})
}

func New() *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = vsan.Namespace
	r.Path = vsan.Path

	r.Put(&StretchedClusterSystem{
		ManagedObjectReference: vsan.VsanVcStretchedClusterSystem,
	})

	return r
}

type StretchedClusterSystem struct {
	vim.ManagedObjectReference
}

func (s *StretchedClusterSystem) VSANVcConvertToStretchedCluster(ctx *simulator.Context, req *types.VSANVcConvertToStretchedCluster) soap.HasFault {
	task := simulator.CreateTask(s, "convertToStretchedCluster", func(*simulator.Task) (vim.AnyType, vim.BaseMethodFault) {
		// TODO: validate req fields
		return nil, nil
	})

	return &methods.VSANVcConvertToStretchedClusterBody{
		Res: &types.VSANVcConvertToStretchedClusterResponse{
			Returnval: task.Run(ctx),
		},
	}
}
