// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	// ModulesPath is rest endpoint for the Cluster Modules API
	ModulesPath = "/vcenter/cluster/modules"
	// ModulesVMPath is rest endpoint for the Cluster Modules Members API
	ModulesVMPath = "/vcenter/cluster/modules/vm"
)

// Status is used for JSON encode/decode
type Status struct {
	Success bool `json:"success"`
}

// CreateModule is used for JSON encode/decode
type CreateModule struct {
	Spec struct {
		ID string `json:"cluster"`
	} `json:"spec"`
}

// ModuleMembers is used for JSON encode/decode
type ModuleMembers struct {
	VMs []string `json:"vms"`
}

// AsReferences converts the ModuleMembers.VM field to morefs
func (m *ModuleMembers) AsReferences() []types.ManagedObjectReference {
	refs := make([]types.ManagedObjectReference, 0, len(m.VMs))
	for _, id := range m.VMs {
		refs = append(refs, types.ManagedObjectReference{
			Type:  "VirtualMachine",
			Value: id,
		})
	}
	return refs
}

// ClusterVM returns all VM references in the given cluster
func ClusterVM(c *vim25.Client, cluster mo.Reference) ([]mo.Reference, error) {
	ctx := context.Background()
	kind := []string{"VirtualMachine"}

	m := view.NewManager(c)
	v, err := m.CreateContainerView(ctx, cluster.Reference(), kind, true)
	if err != nil {
		return nil, err
	}
	defer func() { _ = v.Destroy(ctx) }()

	refs, err := v.Find(ctx, kind, nil)
	if err != nil {
		return nil, err
	}

	vms := make([]mo.Reference, 0, len(refs))
	for i := range refs {
		vms = append(vms, refs[i])
	}

	return vms, nil
}
