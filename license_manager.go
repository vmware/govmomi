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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type LicenseManager struct {
	types.ManagedObjectReference

	c *Client
}

func NewLicenseManager(c *Client, ref types.ManagedObjectReference) LicenseManager {
	return LicenseManager{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (l LicenseManager) Reference() types.ManagedObjectReference {
	return l.ManagedObjectReference
}

func mapToKeyValueSlice(m map[string]string) []types.KeyValue {
	r := make([]types.KeyValue, len(m))
	for k, v := range m {
		r = append(r, types.KeyValue{Key: k, Value: v})
	}
	return r
}

func (l LicenseManager) AddLicense(key string, labels map[string]string) (types.LicenseManagerLicenseInfo, error) {
	req := types.AddLicense{
		This:       l.Reference(),
		LicenseKey: key,
		Labels:     mapToKeyValueSlice(labels),
	}

	res, err := methods.AddLicense(l.c, &req)
	if err != nil {
		return types.LicenseManagerLicenseInfo{}, err
	}

	return res.Returnval, nil
}

func (l LicenseManager) RemoveLicense(key string) error {
	req := types.RemoveLicense{
		This:       l.Reference(),
		LicenseKey: key,
	}

	_, err := methods.RemoveLicense(l.c, &req)
	return err
}

func (l LicenseManager) UpdateLicense(key string, labels map[string]string) (types.LicenseManagerLicenseInfo, error) {
	req := types.UpdateLicense{
		This:       l.Reference(),
		LicenseKey: key,
		Labels:     mapToKeyValueSlice(labels),
	}

	res, err := methods.UpdateLicense(l.c, &req)
	if err != nil {
		return types.LicenseManagerLicenseInfo{}, err
	}

	return res.Returnval, nil
}

func (l LicenseManager) ListLicenses() ([]types.LicenseManagerLicenseInfo, error) {
	var mlm mo.LicenseManager

	err := l.c.Properties(l.Reference(), []string{"licenses"}, &mlm)
	if err != nil {
		return nil, err
	}

	return mlm.Licenses, nil
}
