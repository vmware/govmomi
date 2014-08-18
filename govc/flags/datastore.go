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
	"github.com/vmware/govmomi/vim25/mo"
)

type DatastoreFlag struct {
	*ClientFlag
	*DatacenterFlag

	register sync.Once
	name     string
	ds       *govmomi.Datastore
}

func (f *DatastoreFlag) Register(fs *flag.FlagSet) {
	f.register.Do(func() {
		f.name = os.Getenv("GOVMOMI_DATASTORE")
		fs.StringVar(&f.name, "ds", "", "Datastore")
	})
}

func (f *DatastoreFlag) Process() error {
	return nil
}

func (f *DatastoreFlag) Datastore() (*govmomi.Datastore, error) {
	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	folders, err := dc.Folders(c)
	if err != nil {
		return nil, err
	}
	df := folders.DatastoreFolder

	if f.name != "" {
		ref, err := c.SearchIndex().FindChild(df, f.name)
		if err == nil {
			return nil, err
		}
		f.ds = ref.(*govmomi.Datastore)
		return f.ds, nil
	}

	cs, err := df.Children(c)
	if err != nil {
		return nil, err
	}
	// Default to using the only datastore if there is only one.
	if len(cs) != 1 {
		return nil, fmt.Errorf("more than one datastore, please specify one")
	}

	f.ds = cs[0].(*govmomi.Datastore)
	return f.ds, nil
}

func (f *DatastoreFlag) Properties(p []string) (*mo.Datastore, error) {
	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	_, err = f.Datastore()
	if err != nil {
		return nil, err
	}

	var ds mo.Datastore
	if err := c.Properties(f.ds.Reference(), p, &ds); err != nil {
		return nil, err
	}
	return &ds, nil
}

func (f *DatastoreFlag) Name() (string, error) {
	if f.name == "" {
		ds, err := f.Properties([]string{"name"})
		if err != nil {
			return "", nil
		}
		f.name = ds.Name
	}
	return f.name, nil
}
