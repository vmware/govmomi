// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type TenantManager struct {
	mo.TenantTenantManager

	spEntities map[types.ManagedObjectReference]bool
}

func (t *TenantManager) init(r *Registry) {
	t.spEntities = make(map[types.ManagedObjectReference]bool)
}

func (t *TenantManager) markEntities(entities []types.ManagedObjectReference) {
	for _, e := range entities {
		t.spEntities[e] = true
	}
}

func (t *TenantManager) unmarkEntities(entities []types.ManagedObjectReference) {
	for _, e := range entities {
		_, ok := t.spEntities[e]
		if ok {
			delete(t.spEntities, e)
		}
	}
}

func (t *TenantManager) getEntities() []types.ManagedObjectReference {
	entities := []types.ManagedObjectReference{}
	for e := range t.spEntities {
		entities = append(entities, e)
	}
	return entities
}

func (t *TenantManager) MarkServiceProviderEntities(req *types.MarkServiceProviderEntities) soap.HasFault {
	body := new(methods.MarkServiceProviderEntitiesBody)
	t.markEntities(req.Entity)
	body.Res = &types.MarkServiceProviderEntitiesResponse{}
	return body
}

func (t *TenantManager) UnmarkServiceProviderEntities(req *types.UnmarkServiceProviderEntities) soap.HasFault {
	body := new(methods.UnmarkServiceProviderEntitiesBody)
	t.unmarkEntities(req.Entity)
	body.Res = &types.UnmarkServiceProviderEntitiesResponse{}
	return body
}

func (t *TenantManager) RetrieveServiceProviderEntities(req *types.RetrieveServiceProviderEntities) soap.HasFault {
	body := new(methods.RetrieveServiceProviderEntitiesBody)
	body.Res = &types.RetrieveServiceProviderEntitiesResponse{
		Returnval: t.getEntities(),
	}
	return body
}
