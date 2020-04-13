/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package cache

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Client interface to support client session caching
type Client interface {
	json.Marshaler
	json.Unmarshaler

	Valid() bool
	Path() string
}

// Session provides methods to cache authenticated vim25.Client and rest.Client sessions.
// Use of session cache avoids the expense of creating and deleting vSphere sessions.
// It also helps avoid the problem of "leaking sessions", as Session.Login will only
// create a new authenticated session if the cached session does not exist or is invalid.
// By default, username/password authentication is used to create new sessions.
// The Session.Login{SOAP,REST} fields can be set to use other methods,
// such as SAML token authentication (see govc session.login for example).
type Session struct {
	URL         *url.URL // URL of a vCenter or ESXi instance
	DirSOAP     string   // DirSOAP cache directory. Defaults to "$HOME/.govmomi/sessions"
	DirREST     string   // DirREST cache directory. Defaults to "$HOME/.govmomi/rest_sessions"
	Insecure    bool     // Insecure param for soap.NewClient (tls.Config.InsecureSkipVerify)
	Passthrough bool     // Passthrough disables caching when set to true

	LoginSOAP func(context.Context, *vim25.Client) error // LoginSOAP defaults to session.Manager.Login()
	LoginREST func(context.Context, *rest.Client) error  // LoginREST defaults to rest.Client.Login()
}

var (
	home = os.Getenv("GOVMOMI_HOME")
)

func init() {
	if home == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			dir = os.Getenv("HOME")
		}
		home = filepath.Join(dir, ".govmomi")
	}
}

// Endpoint returns a copy of the Session.URL with Password, Query and Fragment removed.
func (s *Session) Endpoint() *url.URL {
	if s.URL == nil {
		return nil
	}
	p := &url.URL{
		Scheme: s.URL.Scheme,
		Host:   s.URL.Host,
		Path:   s.URL.Path,
	}
	if u := s.URL.User; u != nil {
		p.User = url.User(u.Username()) // Remove password
	}
	return p
}

// key is a digest of the URL scheme + username + host + Client.Path()
func (s *Session) key(path string) string {
	p := s.Endpoint()
	p.Path = path

	// Key session file off of full URI and insecure setting.
	// Hash key to get a predictable, canonical format.
	key := fmt.Sprintf("%s#insecure=%t", p.String(), s.Insecure)
	return fmt.Sprintf("%040x", sha1.Sum([]byte(key)))
}

func (s *Session) file(p string) string {
	dir := ""

	switch p {
	case rest.Path:
		dir = s.DirREST
		if dir == "" {
			dir = filepath.Join(home, "rest_sessions")
		}
	default:
		dir = s.DirSOAP
		if dir == "" {
			dir = filepath.Join(home, "sessions")
		}
	}

	return filepath.Join(dir, s.key(p))
}

// Save a Client in the file cache.
// Session will not be saved if Session.Passthrough is true.
func (s *Session) Save(c Client) error {
	if s.Passthrough {
		return nil
	}

	p := s.file(c.Path())

	err := os.MkdirAll(filepath.Dir(p), 0700)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(c)
	if err != nil {
		_ = f.Close()
		return err
	}

	return f.Close()
}

func (s *Session) get(c Client) (bool, error) {
	f, err := os.Open(s.file(c.Path()))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(c)
	if err != nil {
		_ = f.Close()
		return false, err
	}

	return c.Valid(), f.Close()
}

func localTicket(ctx context.Context, m *session.Manager) (*url.Userinfo, error) {
	name := os.Getenv("USER")
	u, err := user.Current()
	if err == nil {
		name = u.Username
	}

	ticket, err := m.AcquireLocalTicket(ctx, name)
	if err != nil {
		return nil, err
	}

	password, err := ioutil.ReadFile(ticket.PasswordFilePath)
	if err != nil {
		return nil, err
	}

	return url.UserPassword(ticket.UserName, string(password)), nil
}

func (s *Session) loginSOAP(ctx context.Context, c *vim25.Client) error {
	m := session.NewManager(c)
	u := s.URL.User
	name := u.Username()

	if name == "" && !c.IsVC() {
		// If no username is provided, try to acquire a local ticket.
		// When invoked remotely, ESX returns an InvalidRequestFault.
		// So, rather than return an error here, fallthrough to Login() with the original User to
		// to avoid what would be a confusing error message.
		luser, lerr := localTicket(ctx, m)
		if lerr == nil {
			// We are running directly on an ESX or Workstation host and can use the ticket with Login()
			u = luser
			name = u.Username()
		}
	}
	if name == "" {
		// ServiceContent does not require authentication
		return nil
	}

	return m.Login(ctx, u)
}

func (s *Session) loginREST(ctx context.Context, c *rest.Client) error {
	return c.Login(ctx, s.URL.User)
}

func soapSessionValid(ctx context.Context, client *vim25.Client) (bool, error) {
	m := session.NewManager(client)
	u, err := m.UserSession(ctx)
	if err != nil {
		if soap.IsSoapFault(err) {
			fault := soap.ToSoapFault(err).VimFault()
			// If the PropertyCollector is not found, the saved session for this URL is not valid
			if _, ok := fault.(types.ManagedObjectNotFound); ok {
				return false, nil
			}
		}

		return false, err
	}

	return u != nil, nil
}

func restSessionValid(ctx context.Context, client *rest.Client) (bool, error) {
	s, err := client.Session(ctx)
	if err != nil {
		return false, err
	}
	return s != nil, nil
}

// Load a Client from the file cache.
// Returns false if no cache exists or is invalid.
// An error is returned if the file cannot be opened or is not json encoded.
// After loading the Client from the file:
// Returns true if the session is still valid, false otherwise indicating the client requires authentication.
// An error is returned if the session ID cannot be validated.
// Returns false if Session.Passthrough is true.
func (s *Session) Load(ctx context.Context, c Client, config func(*soap.Client) error) (bool, error) {
	if s.Passthrough {
		return false, nil
	}

	ok, err := s.get(c)
	if err != nil {
		return false, err

	}
	if !ok {
		return false, nil
	}

	switch client := c.(type) {
	case *vim25.Client:
		if config != nil {
			if err := config(client.Client); err != nil {
				return false, err
			}
		}
		return soapSessionValid(ctx, client)
	case *rest.Client:
		if config != nil {
			if err := config(client.Client); err != nil {
				return false, err
			}
		}
		return restSessionValid(ctx, client)
	default:
		panic(fmt.Sprintf("unsupported client type=%T", client))
	}
}

// Login returns a cached session via Load() if valid.
// Otherwise, creates a new authenticated session and saves to the cache.
// The config func can be used to apply soap.Client configuration, such as TLS settings.
// When Session.Passthrough is true, Login will always create a new session.
func (s *Session) Login(ctx context.Context, c Client, config func(*soap.Client) error) error {
	ok, err := s.Load(ctx, c, config)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	sc := soap.NewClient(s.URL, s.Insecure)

	if config != nil {
		err = config(sc)
		if err != nil {
			return err
		}
	}

	switch client := c.(type) {
	case *vim25.Client:
		vc, err := vim25.NewClient(ctx, sc)
		if err != nil {
			return err
		}

		login := s.loginSOAP
		if s.LoginSOAP != nil {
			login = s.LoginSOAP
		}
		if err = login(ctx, vc); err != nil {
			return err
		}

		*client = *vc
		c = client
	case *rest.Client:
		vc := &vim25.Client{Client: sc}
		rc := rest.NewClient(vc)

		login := s.loginREST
		if s.LoginREST != nil {
			login = s.LoginREST
		}
		if err = login(ctx, rc); err != nil {
			return err
		}

		*client = *rc
		c = client
	default:
		panic(fmt.Sprintf("unsupported client type=%T", client))
	}

	return s.Save(c)
}

// Login calls the Logout method for the given Client if Session.Passthrough is true.
// Otherwise returns nil.
func (s *Session) Logout(ctx context.Context, c Client) error {
	if s.Passthrough {
		switch client := c.(type) {
		case *vim25.Client:
			return session.NewManager(client).Logout(ctx)
		case *rest.Client:
			return client.Logout(ctx)
		default:
			panic(fmt.Sprintf("unsupported client type=%T", client))
		}
	}
	return nil
}
