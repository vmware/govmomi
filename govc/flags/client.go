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
	"crypto/sha1"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
)

const (
	envURL           = "GOVC_URL"
	envInsecure      = "GOVC_INSECURE"
	envMinAPIVersion = "GOVC_MIN_API_VERSION"
)

const cDescr = "ESX or vCenter URL"

type ClientFlag struct {
	*DebugFlag

	register sync.Once

	url           *url.URL
	insecure      bool
	minAPIVersion string

	client *govmomi.Client
}

func (flag *ClientFlag) URLWithoutPassword() *url.URL {
	if flag.url == nil {
		return nil
	}

	withoutCredentials := *flag.url
	withoutCredentials.User = url.User(flag.url.User.Username())
	return &withoutCredentials
}

func (flag *ClientFlag) String() string {
	url := flag.URLWithoutPassword()
	if url == nil {
		return ""
	}

	return url.String()
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

		if flag.url.User == nil {
			flag.url.User = url.UserPassword("", "")
		}
	}

	return nil
}

func (flag *ClientFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		{
			flag.Set(os.Getenv(envURL))
			usage := fmt.Sprintf("%s [%s]", cDescr, envURL)
			f.Var(flag, "u", usage)
		}

		{
			insecure := false
			switch env := strings.ToLower(os.Getenv(envInsecure)); env {
			case "1", "true":
				insecure = true
			}

			usage := fmt.Sprintf("Skip verification of server certificate [%s]", envInsecure)
			f.BoolVar(&flag.insecure, "k", insecure, usage)
		}

		{
			env := os.Getenv(envMinAPIVersion)
			if env == "" {
				env = "5.5"
			}

			flag.minAPIVersion = env
		}
	})
}

func (flag *ClientFlag) Process() error {
	if flag.url == nil {
		return errors.New("specify an " + cDescr)
	}

	return nil
}

func (flag *ClientFlag) sessionFile() string {
	url := flag.URLWithoutPassword()

	// Key session file off of full URI and insecure setting.
	// Hash key to get a predictable, canonical format.
	key := fmt.Sprintf("%s#insecure=%t", url.String(), flag.insecure)
	name := fmt.Sprintf("%040x", sha1.Sum([]byte(key)))
	return filepath.Join(os.Getenv("HOME"), ".govmomi", "sessions", name)
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
	err = os.MkdirAll(filepath.Dir(p), 0700)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0600)
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

// apiVersionValid returns whether or not the API version supported by the
// server the client is connected to is not recent enough.
func apiVersionValid(c *govmomi.Client, minVersionString string) error {
	realVersion, err := ParseVersion(c.ServiceContent.About.ApiVersion)
	if err != nil {
		return err
	}

	minVersion, err := ParseVersion(minVersionString)
	if err != nil {
		return err
	}

	if !minVersion.Lte(realVersion) {
		err = fmt.Errorf("Require API version %s, connected to API version %s (set %s to override)",
			minVersionString,
			c.ServiceContent.About.ApiVersion,
			envMinAPIVersion)
		return err
	}

	return nil
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

	// Check that the endpoint has the right API version
	err = apiVersionValid(c, flag.minAPIVersion)
	if err != nil {
		return nil, err
	}

	flag.client = c
	return flag.client, nil
}
