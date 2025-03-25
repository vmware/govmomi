// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"net/http"
	"path"

	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/vapi/cluster/internal"
)

// Manager extends rest.Client, adding cluster related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// CreateModule creates a new module in a vCenter cluster.
func (c *Manager) CreateModule(ctx context.Context, ref mo.Reference) (string, error) {
	var s internal.CreateModule
	s.Spec.ID = ref.Reference().Value

	url := c.Resource(internal.ModulesPath)
	var res string
	return res, c.Do(ctx, url.Request(http.MethodPost, s), &res)
}

// DeleteModule deletes a specific module.
func (c *Manager) DeleteModule(ctx context.Context, id string) error {
	url := c.Resource(internal.ModulesPath + "/" + id)
	return c.Do(ctx, url.Request(http.MethodDelete), nil)
}

// ModuleSummary contains commonly used information about a module in a vCenter cluster.
type ModuleSummary struct {
	Cluster string `json:"cluster"`
	Module  string `json:"module"`
}

// ModuleSummaryList is used to JSON encode/decode a ModuleSummary.
type ModuleSummaryList struct {
	Summaries []ModuleSummary `json:"summaries"`
}

// ListModules returns information about the modules available in this vCenter server.
func (c *Manager) ListModules(ctx context.Context) ([]ModuleSummary, error) {
	var res ModuleSummaryList
	url := c.Resource(internal.ModulesPath)
	return res.Summaries, c.Do(ctx, url.Request(http.MethodGet), &res)
}

func memberPath(id string) string {
	return path.Join(internal.ModulesVMPath, id, "members")
}

// ListModuleMembers returns the virtual machines that are members of the module.
func (c *Manager) ListModuleMembers(ctx context.Context, id string) ([]types.ManagedObjectReference, error) {
	var m internal.ModuleMembers
	url := c.Resource(memberPath(id))
	err := c.Do(ctx, url.Request(http.MethodGet), &m)
	if err != nil {
		return nil, err
	}
	return m.AsReferences(), err
}

func (c *Manager) moduleMembers(ctx context.Context, action string, id string, vms ...mo.Reference) (bool, error) {
	url := c.Resource(memberPath(id)).WithParam("action", action)
	var m internal.ModuleMembers
	for i := range vms {
		m.VMs = append(m.VMs, vms[i].Reference().Value)
	}
	var res internal.Status
	return res.Success, c.Do(ctx, url.Request(http.MethodPost, m), &res)
}

// AddModuleMembers adds virtual machines to the module. These virtual machines are required to be in the same vCenter cluster.
// Returns true if all vms are added, false if a vm is already a member of the module or not within the module's cluster.
func (c *Manager) AddModuleMembers(ctx context.Context, id string, vms ...mo.Reference) (bool, error) {
	return c.moduleMembers(ctx, "add", id, vms...)
}

// RemoveModuleMembers removes virtual machines from the module.
// Returns true if all vms are removed, false if a vm is not a member of the module.
func (c *Manager) RemoveModuleMembers(ctx context.Context, id string, vms ...mo.Reference) (bool, error) {
	return c.moduleMembers(ctx, "remove", id, vms...)
}
