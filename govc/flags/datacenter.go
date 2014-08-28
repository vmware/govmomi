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
)

type DatacenterFlag struct {
	*ListFlag

	register sync.Once
	path     string
	dc       *govmomi.Datacenter
}

func (flag *DatacenterFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		flag.path = os.Getenv("GOVC_DATACENTER")
		f.StringVar(&flag.path, "dc", "", "Datacenter")
	})
}

func (flag *DatacenterFlag) Process() error {
	return nil
}

func (flag *DatacenterFlag) findDatacenter(path string) ([]*govmomi.Datacenter, error) {
	relativeFunc := func() (govmomi.Reference, error) {
		c, err := flag.Client()
		if err != nil {
			return nil, err
		}

		return c.RootFolder(), nil
	}

	es, err := flag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var dcs []*govmomi.Datacenter
	for _, e := range es {
		ref := e.Object.Reference()
		if ref.Type == "Datacenter" {
			dcs = append(dcs, govmomi.NewDatacenter(ref))
		}
	}

	return dcs, nil
}

func (flag *DatacenterFlag) findSpecifiedDatacenter(path string) (*govmomi.Datacenter, error) {
	dcs, err := flag.findDatacenter(path)
	if err != nil {
		return nil, err
	}

	if len(dcs) == 0 {
		return nil, errors.New("no such datacenter")
	}

	if len(dcs) > 1 {
		return nil, errors.New("path resolves to multiple datacenters")
	}

	flag.dc = dcs[0]
	return flag.dc, nil
}

func (flag *DatacenterFlag) findDefaultDatacenter() (*govmomi.Datacenter, error) {
	dcs, err := flag.findDatacenter("*")
	if err != nil {
		return nil, err
	}

	if len(dcs) == 0 {
		panic("no datacenters") // Should never happen
	}

	if len(dcs) > 1 {
		return nil, errors.New("please specify a datacenter")
	}

	flag.dc = dcs[0]
	return flag.dc, nil
}

func (flag *DatacenterFlag) Datacenter() (*govmomi.Datacenter, error) {
	if flag.dc != nil {
		return flag.dc, nil
	}

	if flag.path == "" {
		return flag.findDefaultDatacenter()
	}

	return flag.findSpecifiedDatacenter(flag.path)
}
