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

type HostSystemFlag struct {
	*ClientFlag
	*DatacenterFlag
	*SearchFlag
	*ListFlag

	register sync.Once
	name     string
	host     *govmomi.HostSystem
	pool     *govmomi.ResourcePool
}

func (flag *HostSystemFlag) Register(f *flag.FlagSet) {
	flag.SearchFlag = NewSearchFlag(SearchHosts)

	flag.register.Do(func() {
		f.StringVar(&flag.name, "host", os.Getenv("GOVC_HOST"), "Host system [GOVC_HOST]")
	})
}

func (flag *HostSystemFlag) Process() error {
	return nil
}

func (flag *HostSystemFlag) findHostSystem(path string) ([]*govmomi.HostSystem, error) {
	relativeFunc := func() (govmomi.Reference, error) {
		dc, err := flag.DatacenterFlag.Datacenter()
		if err != nil {
			return nil, err
		}

		c, err := flag.ClientFlag.Client()
		if err != nil {
			return nil, err
		}

		f, err := dc.Folders(c)
		if err != nil {
			return nil, err
		}

		return f.HostFolder, nil
	}

	es, err := flag.ListFlag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var hss []*govmomi.HostSystem
	for _, e := range es {
		switch o := e.Object.(type) {
		case mo.HostSystem:
			hs := govmomi.HostSystem{
				ManagedObjectReference: o.Reference(),
			}
			hss = append(hss, &hs)
		}
	}

	return hss, nil
}

func (flag *HostSystemFlag) findSpecifiedHostSystem(path string) (*govmomi.HostSystem, error) {
	hss, err := flag.findHostSystem(path)
	if err != nil {
		return nil, err
	}

	if len(hss) == 0 {
		return nil, errors.New("no such host")
	}

	if len(hss) > 1 {
		return nil, errors.New("path resolves to multiple hosts")
	}

	flag.host = hss[0]
	return flag.host, nil
}

func (flag *HostSystemFlag) HostSystem() (*govmomi.HostSystem, error) {
	if flag.host != nil {
		return flag.host, nil
	}

	// Use search flags if specified.
	if flag.SearchFlag.IsSet() {
		host, err := flag.SearchFlag.HostSystem()
		if err != nil {
			return nil, err
		}

		flag.host = host
		return flag.host, nil
	}

	// Never look for a default host system.
	// A host system parameter is optional for vm creation. It uses a mandatory
	// resource pool parameter to determine where the vm should be placed.
	if flag.name == "" {
		return nil, nil
	}

	return flag.findSpecifiedHostSystem(flag.name)
}

// ResourcePool returns the host system's resource pool, if the host system
// flag itself is specified and valid.
func (flag *HostSystemFlag) ResourcePool() (*govmomi.ResourcePool, error) {
	if flag.pool != nil {
		return flag.pool, nil
	}

	h, err := flag.HostSystem()
	if err != nil {
		return nil, err
	}

	c, err := flag.DatacenterFlag.Client()
	if err != nil {
		return nil, err
	}

	var mh mo.HostSystem
	err = c.Properties(h.Reference(), []string{"parent"}, &mh)
	if err != nil {
		return nil, err
	}

	var mcr mo.ComputeResource
	err = c.Properties(*mh.Parent, []string{"resourcePool"}, &mcr)
	if err != nil {
		return nil, err
	}

	flag.pool = govmomi.NewResourcePool(*mcr.ResourcePool)
	return flag.pool, nil
}
