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
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
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

	if s != "" {
		c.url, err = url.Parse(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientFlag) Register(f *flag.FlagSet) {
	c.register.Do(func() {
		c.Set(os.Getenv("GOVC_URL"))
		f.Var(c, "u", cDescr)
	})
}

func (c *ClientFlag) Process() error {
	if c.url == nil {
		return errors.New("specify an " + cDescr)
	}

	return nil
}

func (c *ClientFlag) sessionFile() string {
	file := fmt.Sprintf("%s@%s", c.url.User.Username(), c.url.Host)
	return path.Join(os.Getenv("HOME"), ".govmomi", "sessions", file)
}

func (c *ClientFlag) loadClient() (*govmomi.Client, error) {
	var client govmomi.Client

	f, err := os.Open(c.sessionFile())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *ClientFlag) newClient() (*govmomi.Client, error) {
	client, err := govmomi.NewClient(*c.url)
	if err != nil {
		return nil, err
	}

	p := c.sessionFile()
	err = os.MkdirAll(path.Dir(p), 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// currentSessionValid returns whether or not the current session is valid. It
// is valid if the "currentSession" field of the SessionManager managed object
// can be retrieved. It is not valid it cannot be retrieved, but no error
// occurs. An error is returned otherwise.
func currentSessionValid(c *govmomi.Client) (bool, error) {
	var sm mo.SessionManager

	err := c.Properties(*c.ServiceContent.SessionManager, []string{"currentSession"}, &sm)
	if err != nil {
		return false, err
	}

	if sm.CurrentSession == nil {
		return false, nil
	}

	return true, nil
}

func (c *ClientFlag) Client() (*govmomi.Client, error) {
	if c.client != nil {
		return c.client, nil
	}

	var ok = false

	client, err := c.loadClient()
	if err != nil {
		return nil, err
	}

	if client != nil {
		ok, err = currentSessionValid(client)
		if err != nil {
			return nil, err
		}
	}

	if !ok {
		client, err = c.newClient()
		if err != nil {
			return nil, err
		}
	}

	c.client = client
	return c.client, nil
}
