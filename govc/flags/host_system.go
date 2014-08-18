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

type HostSystemFlag struct {
	*DatacenterFlag

	name string
	host *govmomi.HostSystem
	pool *govmomi.ResourcePool
}

func (f *HostSystemFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "host", "", "Host system")
}

func (f *HostSystemFlag) Process() error {
	return nil
}

func (f *HostSystemFlag) HostSystem() (*govmomi.HostSystem, error) {
	if f.name == "" {
		return nil, nil // optional
	}

	if f.host != nil {
		return f.host, nil
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	s := c.SearchIndex()

	if strings.Contains(f.name, "/") {
		// TODO: make use of the DatacenterFlag
		ref, err := s.FindByInventoryPath(f.name)

		if err != nil {
			return nil, err
		}
		if host, ok := ref.(*govmomi.HostSystem); ok {
			f.host = host
			return host, nil
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
	// TODO: find a lighter way
	for _, child := range cs {
		ref, err := s.FindChild(child, f.name)
		if err != nil {
			return nil, err
		}

		if host, ok := ref.(*govmomi.HostSystem); ok {
			f.host = host
			return host, nil
		}
	}

	return nil, fmt.Errorf("host system not found")
}

func (f *HostSystemFlag) HostResourcePool() (*govmomi.ResourcePool, error) {
	if f.pool != nil {
		return f.pool, nil
	}

	host, err := f.HostSystem()
	if err != nil {
		return nil, err
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	var h mo.HostSystem
	err = c.Properties(host.Reference(), []string{"parent"}, &h)
	if err != nil {
		return nil, err
	}

	var r mo.ComputeResource
	err = c.Properties(*h.Parent, []string{"resourcePool"}, &r)
	if err != nil {
		return nil, err
	}

	f.pool = &govmomi.ResourcePool{*r.ResourcePool}

	return f.pool, nil
}
