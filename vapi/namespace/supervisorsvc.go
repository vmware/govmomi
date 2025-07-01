// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	Name                         string `json:"display_name"`
	State                        string `json:"state"`
	Description                  string `json:"description"`
	MustBeInstalled              bool   `json:"must_be_installed"`
	HasDefaultVersionsRegistered bool   `json:"has_default_versions_registered"`
}

// SupervisorServiceVersionSummary describes a vSphere Supervisor Service version.
type SupervisorServiceVersionSummary struct {
	SupervisorServiceInfo
	Version string `json:"version"`
}

// SupervisorServiceVersionInfo details a vSphere Supervisor Service version.
type SupervisorServiceVersionInfo struct {
	SupervisorServiceInfo
	Eula                string `json:"EULA"`
	Content             string `json:"content"`
	ContentType         string `json:"content_type"`
	TrustVerified       bool   `json:"trust_verified"`
	RegisteredByDefault bool   `json:"registered_by_default"`
}

// SupervisorService defines a new SupervisorService specification
type SupervisorService struct {
	// The specification required to create a Supervisor Service with a version from inline content that is based on the vSphere application service format.
	VsphereService *SupervisorServicesVSphereSpec `json:"vsphere_spec,omitempty"`
	// The specification required to create a Supervisor Service with a version from inline content that is based on the Carvel application package format.
	CarvelService *SupervisorServicesCarvelSpec `json:"carvel_spec,omitempty"`
}

// SupervisorServiceVersion defines a new SupervisorService version specification
type SupervisorServiceVersion struct {
	// The specification required to create a Supervisor Service with a version from inline content that is based on the vSphere application service format.
	VsphereService *SupervisorServicesVSphereVersionCreateSpec `json:"vsphere_spec,omitempty"`
	// The specification required to create a Supervisor Service with a version from inline content that is based on the Carvel application package format.
	CarvelService *CarvelVersionCreateSpec `json:"carvel_spec,omitempty"`
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

// The “SupervisorServicesCarvelSpec“ class provides a specification required to create a Supervisor Service with a version from Carvel application package format (Package and PackageMetadata resources should be declared).
type SupervisorServicesCarvelSpec struct {
	// Supervisor service version specification that provides the service definitions for one Supervisor Service version.
	VersionSpec CarvelVersionCreateSpec `json:"version_spec"`
}

// The “CarvelVersionCreateSpec“ class provides a specification required to create a Supervisor Service version from Carvel application package format (Package and PackageMetadata resources should be declared).
type CarvelVersionCreateSpec struct {
	// Inline content that contains all service definition of the version in Carvel application package format, which shall be base64 encoded.
	Content string `json:"content"`
}

// CreateSupervisorService creates a new Supervisor Service on vCenter.
func (c *Manager) CreateSupervisorService(ctx context.Context, service *SupervisorService) error {
	url := c.Resource(internal.SupervisorServicesPath)
	return c.Do(ctx, url.Request(http.MethodPost, service), nil)
}

// ListSupervisorServices returns a summary of registered Supervisor Services.
func (c *Manager) ListSupervisorServices(ctx context.Context) ([]SupervisorServiceSummary, error) {
	var res []SupervisorServiceSummary
	url := c.Resource(internal.SupervisorServicesPath)
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// GetSupervisorService gets the information of a specific Supervisor Service.
func (c *Manager) GetSupervisorService(ctx context.Context, id string) (SupervisorServiceInfo, error) {
	var res SupervisorServiceInfo
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id))
	return res, c.Do(ctx, url.Request(http.MethodGet, nil), &res)
}

// ActivateSupervisorServices activates a previously registered Supervisor Service.
func (c *Manager) ActivateSupervisorServices(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id)).WithParam("action", "activate")
	return c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
}

// DeactivateSupervisorServices deactivates a previously registered Supervisor Service.
func (c *Manager) DeactivateSupervisorServices(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id)).WithParam("action", "deactivate")
	return c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
}

// RemoveSupervisorService removes a previously deactivated Supervisor Service.
func (c *Manager) RemoveSupervisorService(ctx context.Context, id string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id))
	err := c.Do(ctx, url.Request(http.MethodDelete, nil), nil)
	return err
}

// ListSupervisorServiceVersions lists all versions of the given Supervisor Service.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/get/
func (c *Manager) ListSupervisorServiceVersions(ctx context.Context, id string) ([]SupervisorServiceVersionSummary, error) {
	var res []SupervisorServiceVersionSummary
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath))
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// RemoveSupervisorServiceVersion removes a previously deactivated Supervisor Service version.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/version/delete/
func (c *Manager) RemoveSupervisorServiceVersion(ctx context.Context, id string, version string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath, version))
	err := c.Do(ctx, url.Request(http.MethodDelete, nil), nil)
	return err
}

// CreateSupervisorServiceVersion creates a new version for an existing Supervisor Service.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/post/
func (c *Manager) CreateSupervisorServiceVersion(ctx context.Context, id string, service *SupervisorServiceVersion) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath))
	return c.Do(ctx, url.Request(http.MethodPost, service), nil)
}

// DeactivateSupervisorServiceVersion deactivates a version of an existing Supervisor Service.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/versionactiondeactivate/patch/
func (c *Manager) DeactivateSupervisorServiceVersion(ctx context.Context, id string, version string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath, version)).WithParam("action", "deactivate")
	return c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
}

// ActivateSupervisorServiceVersion activates a version of an existing Supervisor Service.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/versionactiondeactivate/patch/
func (c *Manager) ActivateSupervisorServiceVersion(ctx context.Context, id string, version string) error {
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath, version)).WithParam("action", "activate")
	return c.Do(ctx, url.Request(http.MethodPatch, nil), nil)
}

// GetSupervisorServiceVersion gets the information of a specific Supervisor Service version.
// https://developer.broadcom.com/xapis/vsphere-automation-api/8.0.3/vcenter/api/vcenter/namespace-management/supervisor-services/supervisor_service/versions/version/get/
func (c *Manager) GetSupervisorServiceVersion(ctx context.Context, id string, version string) (SupervisorServiceVersionInfo, error) {
	var res SupervisorServiceVersionInfo
	url := c.Resource(path.Join(internal.SupervisorServicesPath, id, internal.SupervisorServicesVersionsPath, version))
	return res, c.Do(ctx, url.Request(http.MethodGet, nil), &res)
}
