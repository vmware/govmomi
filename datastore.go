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

package govmomi

import (
	"fmt"
	"path"

	"net/url"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Datastore struct {
	types.ManagedObjectReference

	InventoryPath string

	c *Client
}

func NewDatastore(c *Client, ref types.ManagedObjectReference) *Datastore {
	return &Datastore{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (d Datastore) Reference() types.ManagedObjectReference {
	return d.ManagedObjectReference
}

func (d Datastore) Name() string {
	return path.Base(d.InventoryPath)
}

func (d Datastore) Path(path string) string {
	name := d.Name()
	if name == "" {
		panic("expected non-empty name")
	}

	return fmt.Sprintf("[%s] %s", name, path)
}

// URL for datastore access over HTTP
func (d Datastore) URL(dc *Datacenter, path string) (*url.URL, error) {
	var mdc mo.Datacenter
	if err := d.c.Properties(dc.Reference(), []string{"name"}, &mdc); err != nil {
		return nil, err
	}

	var mds mo.Datastore
	if err := d.c.Properties(d.Reference(), []string{"name"}, &mds); err != nil {
		return nil, err
	}

	u := d.c.Client.URL()

	return &url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   fmt.Sprintf("/folder/%s", path),
		RawQuery: url.Values{
			"dcPath": []string{mdc.Name},
			"dsName": []string{mds.Name},
		}.Encode(),
	}, nil
}

func (d Datastore) Browser() (*HostDatastoreBrowser, error) {
	var do mo.Datastore

	err := d.c.Properties(d.Reference(), []string{"browser"}, &do)
	if err != nil {
		return nil, err
	}

	return NewHostDatastoreBrowser(d.c, do.Browser), nil
}
