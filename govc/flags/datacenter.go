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
	"os"
	"sync"

	"github.com/vmware/govmomi"
)

type DatacenterFlag struct {
	*ClientFlag

	register sync.Once
	name     string
	dc       *govmomi.Datacenter
}

func (f *DatacenterFlag) Register(fs *flag.FlagSet) {
	f.register.Do(func() {
		f.name = os.Getenv("GOVMOMI_DATACENTER")
		fs.StringVar(&f.name, "dc", "", "Datacenter")
	})
}

func (f *DatacenterFlag) Process() error {
	return nil
}

func (f *DatacenterFlag) Datacenter() (*govmomi.Datacenter, error) {
	if f.dc != nil {
		return f.dc, nil
	}

	if f.name != "" {
		dc := govmomi.NewDatacenter(f.name)
		f.dc = &dc
		return f.dc, nil
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	// Default to using the only datacenter if there is only one.
	cs, err := govmomi.Folder{c.ServiceContent.RootFolder}.Children(c)
	if err != nil {
		return nil, err
	}

	if len(cs) != 1 {
		return nil, fmt.Errorf("more than one datacenter, please specify one")
	}

	f.dc = cs[0].(*govmomi.Datacenter)
	return f.dc, nil
}
