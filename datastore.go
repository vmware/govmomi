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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type Datastore struct {
	types.ManagedObjectReference
}

func (d Datastore) Reference() types.ManagedObjectReference {
	return d.ManagedObjectReference
}

// URL for datastore access over HTTP
func (d Datastore) URL(c *Client, dc *Datacenter, path string) (*url.URL, error) {
	var mdc mo.Datacenter
	if err := c.Properties(dc.Reference(), []string{"name"}, &mdc); err != nil {
		return nil, err
	}

	var mds mo.Datastore
	if err := c.Properties(d.Reference(), []string{"name"}, &mds); err != nil {
		return nil, err
	}

	u := c.Client.URL()

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

// UploadFile uploads the local file to the given datastore URL
func (d Datastore) UploadFile(c *Client, file string, u *url.URL) error {
	s, err := os.Stat(file)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest("PUT", u.String(), f)
	if err != nil {
		return err
	}

	req.ContentLength = s.Size()
	req.Header.Set("Content-Type", "application/octet-stream")

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		return nil

	}
	return errors.New(res.Status)
}

// DownloadFile downloads the given datastore URL to a local file
func (d Datastore) DownloadFile(c *Client, file string, u *url.URL) error {
	fh, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fh.Close()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	_, err = io.Copy(fh, res.Body)
	if err != nil {
		return err
	}

	return fh.Close()
}
