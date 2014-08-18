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
	"net/url"
)

type DatastorePath struct {
	*DatastoreFlag

	path string
	name string
}

func (f *DatastorePath) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.path, "n", "", "Datastore path name")
}

func (f *DatastorePath) Process() error {
	return nil
}

func (f *DatastorePath) Name() (string, error) {
	if f.name == "" {
		ds, err := f.DatastoreFlag.Name()
		if err != nil {
			return "", err
		}
		f.name = fmt.Sprintf("[%s] %s", ds, f.path)
	}

	return f.name, nil
}

func (f *DatastorePath) URL() (*url.URL, error) {
	client, err := f.Client()
	if err != nil {
		return nil, err
	}

	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	ds, err := f.Datastore()
	if err != nil {
		return nil, err
	}

	return ds.URL(client, dc, f.path)
}
