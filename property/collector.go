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

package property

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type Collector struct {
	roundTripper soap.RoundTripper
	reference    types.ManagedObjectReference
}

// NewCollector creates a new property collector based on the root property
// collector. It is the responsibility of the caller to destroy it when done.
func NewCollector(ctx context.Context, rt soap.RoundTripper, sc types.ServiceContent) (*Collector, error) {
	req := types.CreatePropertyCollector{
		This: sc.PropertyCollector,
	}

	res, err := methods.CreatePropertyCollector(ctx, rt, &req)
	if err != nil {
		return nil, err
	}

	p := Collector{
		roundTripper: rt,
		reference:    res.Returnval,
	}

	return &p, nil
}

func (p Collector) Reference() types.ManagedObjectReference {
	return p.reference
}

func (p *Collector) Destroy(ctx context.Context) error {
	req := types.DestroyCollector{
		This: p.Reference(),
	}

	_, err := methods.DestroyCollector(ctx, p.roundTripper, &req)
	if err != nil {
		return err
	}

	p.reference = types.ManagedObjectReference{}
	return nil
}

func (p *Collector) CreateFilter(ctx context.Context, req types.CreateFilter) error {
	req.This = p.Reference()

	_, err := methods.CreateFilter(ctx, p.roundTripper, &req)
	if err != nil {
		return err
	}

	return nil
}

func (p *Collector) WaitForUpdates(ctx context.Context, v string) (*types.UpdateSet, error) {
	req := types.WaitForUpdatesEx{
		This:    p.Reference(),
		Version: v,
	}

	res, err := methods.WaitForUpdatesEx(ctx, p.roundTripper, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (p *Collector) Wait(ctx context.Context, obj types.ManagedObjectReference, ps []string, f func([]types.PropertyChange) bool) error {
	req := types.CreateFilter{
		Spec: types.PropertyFilterSpec{
			ObjectSet: []types.ObjectSpec{
				{
					Obj: obj,
				},
			},
			PropSet: []types.PropertySpec{
				{
					PathSet: ps,
					Type:    obj.Type,
				},
			},
		},
	}

	err := p.CreateFilter(ctx, req)
	if err != nil {
		return err
	}

	for version := ""; ; {
		res, err := p.WaitForUpdates(ctx, version)
		if err != nil {
			return err
		}

		version = res.Version

		for _, fs := range res.FilterSet {
			for _, os := range fs.ObjectSet {
				if os.Obj == obj {
					if f(os.ChangeSet) {
						return nil
					}
				}
			}
		}
	}
}
