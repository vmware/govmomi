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
	"errors"
	"flag"
	"os"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
)

type ResourcePoolFlag struct {
	*DatacenterFlag

	register sync.Once
	name     string
	pool     *govmomi.ResourcePool
}

func (flag *ResourcePoolFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		f.StringVar(&flag.name, "pool", os.Getenv("GOVC_RESOURCE_POOL"), "Resource Pool")
	})
}

func (flag *ResourcePoolFlag) Process() error {
	return nil
}

func (flag *ResourcePoolFlag) findResourcePool(path string) ([]*govmomi.ResourcePool, error) {
	relativeFunc := func() (govmomi.Reference, error) {
		dc, err := flag.Datacenter()
		if err != nil {
			return nil, err
		}

		c, err := flag.Client()
		if err != nil {
			return nil, err
		}

		f, err := dc.Folders(c)
		if err != nil {
			return nil, err
		}

		return f.HostFolder, nil
	}

	es, err := flag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var rps []*govmomi.ResourcePool
	for _, e := range es {
		switch o := e.Object.(type) {
		case mo.ComputeResource:
			// Use a compute resouce's root resource pool.
			n := govmomi.ResourcePool{
				ManagedObjectReference: *o.ResourcePool,
			}
			rps = append(rps, &n)
		case mo.ResourcePool:
			n := govmomi.ResourcePool{
				ManagedObjectReference: o.Reference(),
			}
			rps = append(rps, &n)
		}
	}

	return rps, nil
}

func (flag *ResourcePoolFlag) findSpecifiedResourcePool(path string) (*govmomi.ResourcePool, error) {
	rps, err := flag.findResourcePool(path)
	if err != nil {
		return nil, err
	}

	if len(rps) == 0 {
		return nil, errors.New("no such resource pool")
	}

	if len(rps) > 1 {
		return nil, errors.New("path resolves to multiple resource pools")
	}

	flag.pool = rps[0]
	return flag.pool, nil
}

func (flag *ResourcePoolFlag) findDefaultResourcePool() (*govmomi.ResourcePool, error) {
	rps, err := flag.findResourcePool("*/Resources")
	if err != nil {
		return nil, err
	}

	if len(rps) == 0 {
		panic("no resource pools") // Should never happen
	}

	if len(rps) > 1 {
		return nil, errors.New("please specify a resource pool")
	}

	flag.pool = rps[0]
	return flag.pool, nil
}

func (flag *ResourcePoolFlag) ResourcePool() (*govmomi.ResourcePool, error) {
	if flag.pool != nil {
		return flag.pool, nil
	}

	if flag.name == "" {
		return flag.findDefaultResourcePool()
	}

	return flag.findSpecifiedResourcePool(flag.name)
}
