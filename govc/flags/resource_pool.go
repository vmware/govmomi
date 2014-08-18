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

package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
)

type ResourcePoolFlag struct {
	*DatacenterFlag

	name string
	pool *govmomi.ResourcePool
}

func (f *ResourcePoolFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "pool", "", "Resource Pool")
}

func (f *ResourcePoolFlag) Process() error {
	return nil
}

func (f *ResourcePoolFlag) ResourcePool() (*govmomi.ResourcePool, error) {
	if f.pool != nil {
		return f.pool, nil
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	s := c.SearchIndex()

	if strings.Contains(f.name, "/") {
		// e.g. ha-datacenter/host/esxbox.localdomain/Resources
		// TODO: make use of the DatacenterFlag
		ref, err := s.FindByInventoryPath(f.name)
		if err != nil {
			return nil, err
		}

		if pool, ok := ref.(*govmomi.ResourcePool); ok {
			f.pool = pool
			return pool, nil
		}
	}

	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	folders, err := dc.Folders(c)
	if err != nil {
		return nil, err
	}

	cs, err := folders.HostFolder.Children(c)
	if err != nil {
		return nil, err
	}

	// TODO: defaulting to resourcePool of the first ComputeResource here,
	// should find a better default
	if f.name == "" {
		var cr mo.ComputeResource
		err = c.Properties(cs[0].Reference(), []string{"resourcePool"}, &cr)
		if err != nil {
			return nil, err
		}

		f.pool = &govmomi.ResourcePool{*cr.ResourcePool}
		return f.pool, nil
	}

	// TODO: find a lighter way
	for _, child := range cs {
		ref, err := s.FindChild(child, f.name)
		if err != nil {
			return nil, err
		}

		if pool, ok := ref.(*govmomi.ResourcePool); ok {
			f.pool = pool
			return pool, nil
		}
	}

	return nil, fmt.Errorf("resource pool not found")
}
