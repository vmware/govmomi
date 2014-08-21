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
	"net/url"
	"os"
	"sync"

	"github.com/vmware/govmomi"
)

type DatastoreFlag struct {
	*DatacenterFlag

	register sync.Once
	name     string
	ds       *govmomi.Datastore
}

func (flag *DatastoreFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		flag.name = os.Getenv("GOVMOMI_DATASTORE")
		f.StringVar(&flag.name, "ds", "", "Datastore")
	})
}

func (flag *DatastoreFlag) Process() error {
	return nil
}

func (flag *DatastoreFlag) findDatastore(path string) ([]*govmomi.Datastore, error) {
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

		return f.DatastoreFolder, nil
	}

	es, err := flag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var dss []*govmomi.Datastore
	for _, e := range es {
		ref := e.Object.Reference()
		if ref.Type == "Datastore" {
			ds := govmomi.Datastore{
				ManagedObjectReference: ref,
				InventoryPath:          e.Path,
			}

			dss = append(dss, &ds)
		}
	}

	return dss, nil
}

func (flag *DatastoreFlag) findSpecifiedDatastore(path string) (*govmomi.Datastore, error) {
	dss, err := flag.findDatastore(path)
	if err != nil {
		return nil, err
	}

	if len(dss) == 0 {
		return nil, errors.New("no such datastore")
	}

	if len(dss) > 1 {
		return nil, errors.New("path resolves to multiple datastores")
	}

	flag.ds = dss[0]
	return flag.ds, nil
}

func (flag *DatastoreFlag) findDefaultDatastore() (*govmomi.Datastore, error) {
	dss, err := flag.findDatastore("*")
	if err != nil {
		return nil, err
	}

	if len(dss) == 0 {
		panic("no datastores") // Should never happen
	}

	if len(dss) > 1 {
		return nil, errors.New("please specify a datastore")
	}

	flag.ds = dss[0]
	return flag.ds, nil
}

func (flag *DatastoreFlag) Datastore() (*govmomi.Datastore, error) {
	if flag.ds != nil {
		return flag.ds, nil
	}

	if flag.name == "" {
		return flag.findDefaultDatastore()
	}

	return flag.findSpecifiedDatastore(flag.name)
}

func (flag *DatastoreFlag) DatastorePath(name string) (string, error) {
	ds, err := flag.Datastore()
	if err != nil {
		return "", err
	}

	return ds.Path(name), nil
}

func (flag *DatastoreFlag) DatastoreURL(path string) (*url.URL, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	dc, err := flag.Datacenter()
	if err != nil {
		return nil, err
	}

	ds, err := flag.Datastore()
	if err != nil {
		return nil, err
	}

	u, err := ds.URL(c, dc, path)
	if err != nil {
		return nil, err
	}

	return u, nil
}
