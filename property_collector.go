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
	"github.com/vmware/govmomi/vim25/types"
)

type PropertyCollector struct {
	c *Client
	r types.ManagedObjectReference
}

func (p *PropertyCollector) Destroy() error {
	req := types.DestroyPropertyCollector{
		This: p.r,
	}

	_, err := methods.DestroyPropertyCollector(p.c, &req)
	if err != nil {
		return err
	}

	p.r = types.ManagedObjectReference{}
	return nil
}

func (p *PropertyCollector) CreateFilter(req types.CreateFilter) error {
	req.This = p.r

	_, err := methods.CreateFilter(p.c, &req)
	if err != nil {
		return err
	}

	return nil
}

func (p *PropertyCollector) WaitForUpdates(v string) (*types.UpdateSet, error) {
	req := types.WaitForUpdatesEx{
		This:    p.r,
		Version: v,
	}

	res, err := methods.WaitForUpdatesEx(p.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
