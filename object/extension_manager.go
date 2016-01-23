/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// ExtensionManager represents an extension manager client
type ExtensionManager struct {
	Common
}

// GetExtensionManager wraps NewExtensionManager, returning ErrNotSupported
// when the client is not connected to a vCenter instance.
func GetExtensionManager(c *vim25.Client) (*ExtensionManager, error) {
	if c.ServiceContent.ExtensionManager == nil {
		return nil, ErrNotSupported
	}
	return NewExtensionManager(c), nil
}

// NewExtensionManager creates a new extension manager client
func NewExtensionManager(c *vim25.Client) *ExtensionManager {
	o := ExtensionManager{
		Common: NewCommon(c, *c.ServiceContent.ExtensionManager),
	}

	return &o
}

// List the known extensions
func (m ExtensionManager) List(ctx context.Context) ([]types.Extension, error) {
	var em mo.ExtensionManager

	err := m.Properties(ctx, m.Reference(), []string{"extensionList"}, &em)
	if err != nil {
		return nil, err
	}

	return em.ExtensionList, nil
}

// Find the extension by key
func (m ExtensionManager) Find(ctx context.Context, key string) (*types.Extension, error) {
	req := types.FindExtension{
		This:         m.Reference(),
		ExtensionKey: key,
	}

	res, err := methods.FindExtension(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// Register the extension
func (m ExtensionManager) Register(ctx context.Context, extension types.Extension) error {
	req := types.RegisterExtension{
		This:      m.Reference(),
		Extension: extension,
	}

	_, err := methods.RegisterExtension(ctx, m.c, &req)
	return err
}

// SetCertificate for the extension
func (m ExtensionManager) SetCertificate(ctx context.Context, key string, certificatePem string) error {
	req := types.SetExtensionCertificate{
		This:           m.Reference(),
		ExtensionKey:   key,
		CertificatePem: certificatePem,
	}

	_, err := methods.SetExtensionCertificate(ctx, m.c, &req)
	return err
}

// Unregister the extension by name
func (m ExtensionManager) Unregister(ctx context.Context, key string) error {
	req := types.UnregisterExtension{
		This:         m.Reference(),
		ExtensionKey: key,
	}

	_, err := methods.UnregisterExtension(ctx, m.c, &req)
	return err
}

// Update the extension
func (m ExtensionManager) Update(ctx context.Context, extension types.Extension) error {
	req := types.UpdateExtension{
		This:      m.Reference(),
		Extension: extension,
	}

	_, err := methods.UpdateExtension(ctx, m.c, &req)
	return err
}
