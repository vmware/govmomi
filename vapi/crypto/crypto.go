/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package crypto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/crypto/internal"
	"github.com/vmware/govmomi/vapi/rest"
)

// Manager extends rest.Client, adding crypto related methods.
// Currently providing create, delete and export only.
// See crypto.ManagerKmip for getting provider details.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

type KmsProviderConstraints struct {
	TpmRequired bool `json:"tpm_required"`
}

type KmsProviderCreateSpec struct {
	Provider    string                 `json:"provider"`
	Constraints KmsProviderConstraints `json:"constraints"`
}

type KmsProviderExportSpec struct {
	Provider string `json:"provider"`
	Password string `json:"password,omitempty"`
}

type KmsProviderDownloadToken struct {
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}

type KmsProviderExportLocation struct {
	URL           string                   `json:"url"`
	DownloadToken KmsProviderDownloadToken `json:"download_token"`
}

type KmsProviderExport struct {
	Type     string                     `json:"type"`
	Location *KmsProviderExportLocation `json:"location,omitempty"`
}

func (c *Manager) KmsProviderCreate(ctx context.Context, spec KmsProviderCreateSpec) error {
	resource := c.Resource(internal.KmsProvidersPath)
	request := resource.Request(http.MethodPost, spec)
	return c.Do(ctx, request, nil)
}

func (c *Manager) KmsProviderDelete(ctx context.Context, provider string) error {
	resource := c.Resource(internal.KmsProvidersPath).WithSubpath(provider)
	request := resource.Request(http.MethodDelete)
	return c.Do(ctx, request, nil)
}

func (c *Manager) KmsProviderExport(ctx context.Context, spec KmsProviderExportSpec) (*KmsProviderExport, error) {
	resource := c.Resource(internal.KmsProvidersPath).WithParam("action", "export")
	request := resource.Request(http.MethodPost, spec)

	var res KmsProviderExport
	if err := c.Do(ctx, request, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Manager) KmsProviderExportRequest(ctx context.Context, export *KmsProviderExportLocation) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, export.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", export.DownloadToken.Token))

	return req, nil
}
