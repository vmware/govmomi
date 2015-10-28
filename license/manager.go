/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package license

import (
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type Manager struct {
	object.Common
}

func NewManager(c *vim25.Client) *Manager {
	m := Manager{
		object.NewCommon(c, *c.ServiceContent.LicenseManager),
	}

	return &m
}

func mapToKeyValueSlice(m map[string]string) []types.KeyValue {
	r := make([]types.KeyValue, len(m))
	for k, v := range m {
		r = append(r, types.KeyValue{Key: k, Value: v})
	}
	return r
}

func (m Manager) Add(ctx context.Context, key string, labels map[string]string) (types.LicenseManagerLicenseInfo, error) {
	req := types.AddLicense{
		This:       m.Reference(),
		LicenseKey: key,
		Labels:     mapToKeyValueSlice(labels),
	}

	res, err := methods.AddLicense(ctx, m.Client(), &req)
	if err != nil {
		return types.LicenseManagerLicenseInfo{}, err
	}

	return res.Returnval, nil
}

func (m Manager) Decode(ctx context.Context, key string) (types.LicenseManagerLicenseInfo, error) {
	req := types.DecodeLicense{
		This:       m.Reference(),
		LicenseKey: key,
	}

	res, err := methods.DecodeLicense(ctx, m.Client(), &req)
	if err != nil {
		return types.LicenseManagerLicenseInfo{}, err
	}

	return res.Returnval, nil
}

func (m Manager) Remove(ctx context.Context, key string) error {
	req := types.RemoveLicense{
		This:       m.Reference(),
		LicenseKey: key,
	}

	_, err := methods.RemoveLicense(ctx, m.Client(), &req)
	return err
}

func (m Manager) Update(ctx context.Context, key string, labels map[string]string) (types.LicenseManagerLicenseInfo, error) {
	req := types.UpdateLicense{
		This:       m.Reference(),
		LicenseKey: key,
		Labels:     mapToKeyValueSlice(labels),
	}

	res, err := methods.UpdateLicense(ctx, m.Client(), &req)
	if err != nil {
		return types.LicenseManagerLicenseInfo{}, err
	}

	return res.Returnval, nil
}

func (m Manager) List(ctx context.Context) ([]types.LicenseManagerLicenseInfo, error) {
	var mlm mo.LicenseManager

	err := m.Properties(ctx, m.Reference(), []string{"licenses"}, &mlm)
	if err != nil {
		return nil, err
	}

	return mlm.Licenses, nil
}

func (m Manager) AssignmentManager(ctx context.Context) (*AssignmentManager, error) {
	var mlm mo.LicenseManager

	err := m.Properties(ctx, m.Reference(), []string{"licenseAssignmentManager"}, &mlm)
	if err != nil {
		return nil, err
	}

	if mlm.LicenseAssignmentManager == nil {
		return nil, object.ErrNotSupported
	}

	am := AssignmentManager{
		object.NewCommon(m.Client(), *mlm.LicenseAssignmentManager),
	}

	return &am, nil
}
