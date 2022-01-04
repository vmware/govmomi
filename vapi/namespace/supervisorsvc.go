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

package namespace

import (
	"context"
	"net/http"
	"path"

	"github.com/vmware/govmomi/vapi/namespace/internal"
)

// SupervisorServiceSummary for a supervisor service existent in vSphere.
type SupervisorServiceSummary struct {
	ID    string `json:"supervisor_service"`
	Name  string `json:"display_name"`
	State string `json:"state"`
}

// SupervisorServiceInfo for a supervisor service existent in vSphere.
type SupervisorServiceInfo struct {
	Name        string `json:"display_name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

// SupervisorService defines a new SupervisorService specification
type SupervisorService struct {
	VsphereService SupervisorServicesVSphereSpec `json:"vsphere_spec,omitempty"`
}

// SupervisorServicesVSphereSpec defines a new SupervisorService specification of vSphere type
type SupervisorServicesVSphereSpec struct {
	VersionSpec SupervisorServicesVSphereVersionCreateSpec `json:"version_spec"`
}

// SupervisorServicesVSphereVersionCreateSpec defines a new SupervisorService specification for vSphere
type SupervisorServicesVSphereVersionCreateSpec struct {
	Content         string `json:"content"`
	TrustedProvider bool   `json:"trusted_provider,omitempty"`
	AcceptEula      bool   `json:"accept_EULA,omitempty"`
}

// CreateSupervisorService creates a new Supervisor Service on vSphere Namespaces endpoint.
func (c *Manager) CreateSupervisorService(ctx context.Context, service *SupervisorService) error {
	url := c.Resource(internal.SupervisorServicesPath)
	return c.Do(ctx, url.Request(http.MethodPost, service), nil)
}

// ListSupervisorServices returns a summary of all clusters with vSphere Namespaces enabled.
func (c *Manager) ListSupervisorServices(ctx context.Context) ([]SupervisorServiceSummary, error) {
	var res []SupervisorServiceSummary
	url := c.Resource(internal.SupervisorServicesPath)
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetSupervisorService gets the information of a specific supervisor service.
func (c *Manager) GetSupervisorService(ctx context.Context, id string) (SupervisorServiceInfo, error) {
	var res SupervisorServiceInfo
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id))
	return res, c.Do(ctx, url.Request(http.MethodGet, nil), &res)
}

// ActivateSupervisorServices activates a previously registered Supervisor Service.
func (c *Manager) ActivateSupervisorServices(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id)).WithParam("action", "activate")
	err := c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
	return err
}

// DeactivateSupervisorServices deactivates a previously registered Supervisor Service.
func (c *Manager) DeactivateSupervisorServices(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id)).WithParam("action", "deactivate")
	err := c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
	return err
}

// RemoveSupervisorService removes a previously deactivated supervisor service.
func (c *Manager) RemoveSupervisorService(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id))
	err := c.Do(ctx, url.Request(http.MethodDelete, nil), nil)
	return err
}
