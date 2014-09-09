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
	"regexp"
	"strings"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
)

const cDescr = "ESX or vCenter URL"

type ClientFlag struct {
	*DebugFlag

	register sync.Once

	url      *url.URL
	insecure bool
	client   *govmomi.Client
}

func (flag *ClientFlag) String() string {
	if flag.url != nil {
		withoutCredentials := *flag.url
		withoutCredentials.User = nil
		return withoutCredentials.String()
	}
	return ""
}

var schemeMatch = regexp.MustCompile(`^\w+://`)

func (flag *ClientFlag) Set(s string) error {
	var err error

	if s != "" {
		// Default the scheme to https
		if !schemeMatch.MatchString(s) {
			s = "https://" + s
		}

		flag.url, err = url.Parse(s)
		if err != nil {
			return err
		}

		// Default the path to /sdk
		if flag.url.Path == "" {
			flag.url.Path = "/sdk"
		}
	}

	return nil
}

func (flag *ClientFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		flag.Set(os.Getenv("GOVC_URL"))
		f.Var(flag, "u", cDescr+" [GOVC_URL]")

		insecure := false
		switch env := strings.ToLower(os.Getenv("GOVC_INSECURE")); env {
		case "1", "true":
			insecure = true
		}

		f.BoolVar(&flag.insecure, "k", insecure, "Skip verification of server certificate [GOVC_INSECURE]")
	})
}

func (flag *ClientFlag) Process() error {
	if flag.url == nil {
		return errors.New("specify an " + cDescr)
	}

	return nil
}

func (flag *ClientFlag) sessionFile() string {
	file := fmt.Sprintf("%s@%s?insecure=%t", flag.url.User.Username(), flag.url.Host, flag.insecure)
	return path.Join(os.Getenv("HOME"), ".govmomi", "sessions", file)
}

func (flag *ClientFlag) loadClient() (*govmomi.Client, error) {
	var c govmomi.Client

	f, err := os.Open(flag.sessionFile())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (flag *ClientFlag) newClient() (*govmomi.Client, error) {
	c, err := govmomi.NewClient(*flag.url, flag.insecure)
	if err != nil {
		return nil, err
	}

	p := flag.sessionFile()
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
	err = enc.Encode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
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

func (flag *ClientFlag) Client() (*govmomi.Client, error) {
	if flag.client != nil {
		return flag.client, nil
	}

	var ok = false

	c, err := flag.loadClient()
	if err != nil {
		return nil, err
	}

	if c != nil {
		ok, err = currentSessionValid(c)
		if err != nil {
			return nil, err
		}
	}

	if !ok {
		c, err = flag.newClient()
		if err != nil {
			return nil, err
		}
	}

	flag.client = c
	return flag.client, nil
}
