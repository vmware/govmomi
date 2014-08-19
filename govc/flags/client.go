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

const cDescr = "ESX or vCenter URL"

type ClientFlag struct {
	*DebugFlag

	register sync.Once
	url      *url.URL
	client   *govmomi.Client
}

func (c *ClientFlag) String() string {
	if c.url != nil {
		withoutCredentials := *c.url
		withoutCredentials.User = nil
		return withoutCredentials.String()
	}
	return ""
}

func (c *ClientFlag) Set(s string) error {
	var err error

	c.url, err = url.Parse(s)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientFlag) Register(f *flag.FlagSet) {
	c.register.Do(func() {
		c.Set(os.Getenv("GOVMOMI_URL"))
		f.Var(c, "u", cDescr)
	})
}

func (c *ClientFlag) Process() error {
	if c.url == nil {
		return errors.New("specify an " + cDescr)
	}

	return nil
}

func (c *ClientFlag) Client() (*govmomi.Client, error) {
	if c.client != nil {
		return c.client, nil
	}

	client, err := govmomi.NewClient(*c.url)
	if err != nil {
		return nil, err
	}

	c.client = client
	return c.client, nil
}
