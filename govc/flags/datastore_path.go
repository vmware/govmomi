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
)

type DatastorePathFlag struct {
	*DatastoreFlag

	name string
}

func (f *DatastorePathFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "n", "", "Datastore path name")
}

func (f *DatastorePathFlag) Process() error {
	if f.name == "" {
		return errors.New("-n flag is required")
	}
	return nil
}

func (f *DatastorePathFlag) Path() (string, error) {
	return f.DatastorePath(f.name)
}

func (f *DatastorePathFlag) URL() (*url.URL, error) {
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

	return ds.URL(client, dc, f.name)
}
