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
	"github.com/vmware/govmomi/eam"
	"github.com/vmware/govmomi/eam/types"
	"github.com/vmware/govmomi/simulator"
	vimmo "github.com/vmware/govmomi/vim25/mo"
	vim "github.com/vmware/govmomi/vim25/types"
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
	r.Namespace = eam.Namespace
	r.Path = eam.Path
	r.Handler = invalidLoginForNilSessionFn

	r.Put(&EsxAgentManager{
		EamObject: EamObject{
			Self: eam.EsxAgentManager,
		},
	})

	return r
}

// invalidLoginForNilSessionFn returns EamInvalidLogin if the provided
// ctx.Session is nil. This is for validating all calls to EAM methods
// have a valid credential.
func invalidLoginForNilSessionFn(
	ctx *simulator.Context,
	_ *simulator.Method) (vimmo.Reference, vim.BaseMethodFault) {

	if ctx.Session == nil {
		return nil, new(types.EamInvalidLogin)
	}
	return nil, nil
}
