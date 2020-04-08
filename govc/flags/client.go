/*
Copyright (c) 2014-2018 VMware, Inc. All Rights Reserved.

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
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

const (
	envURL           = "GOVC_URL"
	envUsername      = "GOVC_USERNAME"
	envPassword      = "GOVC_PASSWORD"
	envCertificate   = "GOVC_CERTIFICATE"
	envPrivateKey    = "GOVC_PRIVATE_KEY"
	envInsecure      = "GOVC_INSECURE"
	envPersist       = "GOVC_PERSIST_SESSION"
	envMinAPIVersion = "GOVC_MIN_API_VERSION"
	envVimNamespace  = "GOVC_VIM_NAMESPACE"
	envVimVersion    = "GOVC_VIM_VERSION"
	envTLSCaCerts    = "GOVC_TLS_CA_CERTS"
	envTLSKnownHosts = "GOVC_TLS_KNOWN_HOSTS"

	defaultMinVimVersion = "5.5"
)

const cDescr = "ESX or vCenter URL"

type ClientFlag struct {
	common

	*DebugFlag

	username      string
	password      string
	cert          string
	key           string
	persist       bool
	minAPIVersion string
	vimNamespace  string
	vimVersion    string
	tlsCaCerts    string
	tlsKnownHosts string
	client        *vim25.Client
	restClient    *rest.Client
	Session       cache.Session
}

var (
	home          = os.Getenv("GOVMOMI_HOME")
	clientFlagKey = flagKey("client")
)

func init() {
	if home == "" {
		home = filepath.Join(os.Getenv("HOME"), ".govmomi")
	}
}

func NewClientFlag(ctx context.Context) (*ClientFlag, context.Context) {
	if v := ctx.Value(clientFlagKey); v != nil {
		return v.(*ClientFlag), ctx
	}

	v := &ClientFlag{}
	v.DebugFlag, ctx = NewDebugFlag(ctx)
	ctx = context.WithValue(ctx, clientFlagKey, v)
	return v, ctx
}

func (flag *ClientFlag) String() string {
	url := flag.Session.Endpoint()
	if url == nil {
		return ""
	}

	return url.String()
}

func (flag *ClientFlag) Set(s string) error {
	var err error

	flag.Session.URL, err = soap.ParseURL(s)

	return err
}

func (flag *ClientFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.DebugFlag.Register(ctx, f)

		{
			flag.Set(os.Getenv(envURL))
			usage := fmt.Sprintf("%s [%s]", cDescr, envURL)
			f.Var(flag, "u", usage)
		}

		{
			flag.username = os.Getenv(envUsername)
			flag.password = os.Getenv(envPassword)
		}

		{
			value := os.Getenv(envCertificate)
			usage := fmt.Sprintf("Certificate [%s]", envCertificate)
			f.StringVar(&flag.cert, "cert", value, usage)
		}

		{
			value := os.Getenv(envPrivateKey)
			usage := fmt.Sprintf("Private key [%s]", envPrivateKey)
			f.StringVar(&flag.key, "key", value, usage)
		}

		{
			insecure := false
			switch env := strings.ToLower(os.Getenv(envInsecure)); env {
			case "1", "true":
				insecure = true
			}

			usage := fmt.Sprintf("Skip verification of server certificate [%s]", envInsecure)
			f.BoolVar(&flag.Session.Insecure, "k", insecure, usage)
		}

		{
			persist := true
			switch env := strings.ToLower(os.Getenv(envPersist)); env {
			case "0", "false":
				persist = false
			}

			usage := fmt.Sprintf("Persist session to disk [%s]", envPersist)
			f.BoolVar(&flag.persist, "persist-session", persist, usage)
		}

		{
			env := os.Getenv(envMinAPIVersion)
			if env == "" {
				env = defaultMinVimVersion
			}

			flag.minAPIVersion = env
		}

		{
			value := os.Getenv(envVimNamespace)
			if value == "" {
				value = vim25.Namespace
			}
			usage := fmt.Sprintf("Vim namespace [%s]", envVimNamespace)
			f.StringVar(&flag.vimNamespace, "vim-namespace", value, usage)
		}

		{
			value := os.Getenv(envVimVersion)
			if value == "" {
				value = vim25.Version
			}
			usage := fmt.Sprintf("Vim version [%s]", envVimVersion)
			f.StringVar(&flag.vimVersion, "vim-version", value, usage)
		}

		{
			value := os.Getenv(envTLSCaCerts)
			usage := fmt.Sprintf("TLS CA certificates file [%s]", envTLSCaCerts)
			f.StringVar(&flag.tlsCaCerts, "tls-ca-certs", value, usage)
		}

		{
			value := os.Getenv(envTLSKnownHosts)
			usage := fmt.Sprintf("TLS known hosts file [%s]", envTLSKnownHosts)
			f.StringVar(&flag.tlsKnownHosts, "tls-known-hosts", value, usage)
		}
	})
}

func (flag *ClientFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		err := flag.DebugFlag.Process(ctx)
		if err != nil {
			return err
		}

		if flag.Session.URL == nil {
			return errors.New("specify an " + cDescr)
		}

		if !flag.persist {
			flag.Session.Passthrough = true
		}

		flag.username, err = session.Secret(flag.username)
		if err != nil {
			return err
		}
		flag.password, err = session.Secret(flag.password)
		if err != nil {
			return err
		}

		// Override username if set
		if flag.username != "" {
			var password string
			var ok bool

			if flag.Session.URL.User != nil {
				password, ok = flag.Session.URL.User.Password()
			}

			if ok {
				flag.Session.URL.User = url.UserPassword(flag.username, password)
			} else {
				flag.Session.URL.User = url.User(flag.username)
			}
		}

		// Override password if set
		if flag.password != "" {
			var username string

			if flag.Session.URL.User != nil {
				username = flag.Session.URL.User.Username()
			}

			flag.Session.URL.User = url.UserPassword(username, flag.password)
		}

		return nil
	})
}

func (flag *ClientFlag) ConfigureTLS(sc *soap.Client) error {
	if flag.cert != "" {
		cert, err := tls.LoadX509KeyPair(flag.cert, flag.key)
		if err != nil {
			return fmt.Errorf("%s=%q %s=%q: %s", envCertificate, flag.cert, envPrivateKey, flag.key, err)
		}

		sc.SetCertificate(cert)
	}

	// Set namespace and version
	sc.Namespace = "urn:" + flag.vimNamespace
	sc.Version = flag.vimVersion

	sc.UserAgent = fmt.Sprintf("govc/%s", Version)

	if err := flag.SetRootCAs(sc); err != nil {
		return err
	}

	if err := sc.LoadThumbprints(flag.tlsKnownHosts); err != nil {
		return err
	}

	t := sc.DefaultTransport()
	var err error

	value := os.Getenv("GOVC_TLS_HANDSHAKE_TIMEOUT")
	if value != "" {
		t.TLSHandshakeTimeout, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (flag *ClientFlag) SetRootCAs(c *soap.Client) error {
	if flag.tlsCaCerts != "" {
		return c.SetRootCAs(flag.tlsCaCerts)
	}
	return nil
}

func isDevelopmentVersion(apiVersion string) bool {
	// Skip version check for development builds which can be in the form of "r4A70F" or "6.5.x"
	return strings.Count(apiVersion, ".") == 0 || strings.HasSuffix(apiVersion, ".x")
}

// apiVersionValid returns whether or not the API version supported by the
// server the client is connected to is not recent enough.
func apiVersionValid(c *vim25.Client, minVersionString string) error {
	if minVersionString == "-" {
		// Disable version check
		return nil
	}

	apiVersion := c.ServiceContent.About.ApiVersion
	if isDevelopmentVersion(apiVersion) {
		return nil
	}

	realVersion, err := ParseVersion(apiVersion)
	if err != nil {
		return fmt.Errorf("error parsing API version %q: %s", apiVersion, err)
	}

	minVersion, err := ParseVersion(minVersionString)
	if err != nil {
		return fmt.Errorf("error parsing %s=%q: %s", envMinAPIVersion, minVersionString, err)
	}

	if !minVersion.Lte(realVersion) {
		err = fmt.Errorf("require API version %q, connected to API version %q (set %s to override)",
			minVersionString,
			c.ServiceContent.About.ApiVersion,
			envMinAPIVersion)
		return err
	}

	return nil
}

func (flag *ClientFlag) Client() (*vim25.Client, error) {
	if flag.client != nil {
		return flag.client, nil
	}

	c := new(vim25.Client)
	err := flag.Session.Login(context.Background(), c, flag.ConfigureTLS)
	if err != nil {
		return nil, err
	}

	// Check that the endpoint has the right API version
	err = apiVersionValid(c, flag.minAPIVersion)
	if err != nil {
		return nil, err
	}

	if flag.vimVersion == "" {
		err = c.UseServiceVersion()
		if err != nil {
			return nil, err
		}
	}

	// Retry twice when a temporary I/O error occurs.
	// This means a maximum of 3 attempts.
	c.RoundTripper = vim25.Retry(c.Client, vim25.TemporaryNetworkError(3))
	flag.client = c

	return flag.client, nil
}

func (flag *ClientFlag) RestClient() (*rest.Client, error) {
	if flag.restClient != nil {
		return flag.restClient, nil
	}

	c := new(rest.Client)

	err := flag.Session.Login(context.Background(), c, flag.ConfigureTLS)
	if err != nil {
		return nil, err
	}

	flag.restClient = c
	return flag.restClient, nil
}

func (flag *ClientFlag) Logout(ctx context.Context) error {
	if flag.client != nil {
		_ = flag.Session.Logout(ctx, flag.client)
	}

	if flag.restClient != nil {
		_ = flag.Session.Logout(ctx, flag.restClient)
	}

	return nil
}

// Environ returns the govc environment variables for this connection
func (flag *ClientFlag) Environ(extra bool) []string {
	var env []string
	add := func(k, v string) {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	u := *flag.Session.URL
	if u.User != nil {
		add(envUsername, u.User.Username())

		if p, ok := u.User.Password(); ok {
			add(envPassword, p)
		}

		u.User = nil
	}

	if u.Path == vim25.Path {
		u.Path = ""
	}
	u.Fragment = ""
	u.RawQuery = ""

	add(envURL, strings.TrimPrefix(u.String(), "https://"))

	keys := []string{
		envCertificate,
		envPrivateKey,
		envInsecure,
		envPersist,
		envMinAPIVersion,
		envVimNamespace,
		envVimVersion,
	}

	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			add(k, v)
		}
	}

	if extra {
		add("GOVC_URL_SCHEME", flag.Session.URL.Scheme)

		v := strings.SplitN(u.Host, ":", 2)
		add("GOVC_URL_HOST", v[0])
		if len(v) == 2 {
			add("GOVC_URL_PORT", v[1])
		}

		add("GOVC_URL_PATH", flag.Session.URL.Path)

		if f := flag.Session.URL.Fragment; f != "" {
			add("GOVC_URL_FRAGMENT", f)
		}

		if q := flag.Session.URL.RawQuery; q != "" {
			add("GOVC_URL_QUERY", q)
		}
	}

	return env
}

// WithCancel calls the given function, returning when complete or canceled via SIGINT.
func (flag *ClientFlag) WithCancel(ctx context.Context, f func(context.Context) error) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	wctx, cancel := context.WithCancel(ctx)
	defer cancel()

	done := make(chan bool)
	var werr error

	go func() {
		defer close(done)
		werr = f(wctx)
	}()

	select {
	case <-sig:
		cancel()
		<-done // Wait for f() to complete
	case <-done:
	}

	return werr
}
