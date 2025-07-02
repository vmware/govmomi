// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
