// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/simulator/internal"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type ServiceInstance struct {
	mo.ServiceInstance
}

func NewServiceInstance(ctx *Context, content types.ServiceContent, folder mo.Folder) (*Context, *ServiceInstance) {
	s := &ServiceInstance{}

	s.Self = vim25.ServiceInstance
	s.Content = content

	ctx.Map.Put(s)

	f := &Folder{Folder: folder}
	ctx.Map.Put(f)

	if content.About.ApiType == "HostAgent" {
		CreateDefaultESX(ctx, f)
	} else {
		content.About.InstanceUuid = uuid.New().String()
	}

	refs := mo.References(content)

	for i := range refs {
		if ctx.Map.Get(refs[i]) != nil {
			continue
		}
		content := types.ObjectContent{Obj: refs[i]}
		o, err := loadObject(ctx, content)
		if err != nil {
			panic(err)
		}
		ctx.Map.Put(o)
	}

	return ctx, s
}

func (s *ServiceInstance) ServiceContent() types.ServiceContent {
	return s.Content
}

func (s *ServiceInstance) RetrieveServiceContent(*types.RetrieveServiceContent) soap.HasFault {
	return &methods.RetrieveServiceContentBody{
		Res: &types.RetrieveServiceContentResponse{
			Returnval: s.Content,
		},
	}
}

func (*ServiceInstance) CurrentTime(*types.CurrentTime) soap.HasFault {
	return &methods.CurrentTimeBody{
		Res: &types.CurrentTimeResponse{
			Returnval: time.Now(),
		},
	}
}

func (s *ServiceInstance) RetrieveInternalContent(*internal.RetrieveInternalContent) soap.HasFault {
	return &internal.RetrieveInternalContentBody{
		Res: &internal.RetrieveInternalContentResponse{
			Returnval: internal.InternalServiceInstanceContent{
				NfcService: types.ManagedObjectReference{Type: "NfcService", Value: "NfcService"},
			},
		},
	}
}
