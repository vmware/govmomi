// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vapi/authentication"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
)

type login struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clone  bool
	issue  bool
	jwt    string
	renew  bool
	long   bool
	vapi   bool
	ticket string
	life   time.Duration
	cookie string
	token  string
	ext    string
	as     string
	method string
}

func init() {
	cli.Register("session.login", &login{})
}

func (cmd *login) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.clone, "clone", false, "Acquire clone ticket")
	f.BoolVar(&cmd.issue, "issue", false, "Issue SAML token")
	f.StringVar(&cmd.jwt, "jwt", "", "Exchange SAML token for JWT audience")
	f.BoolVar(&cmd.renew, "renew", false, "Renew SAML token")
	f.BoolVar(&cmd.vapi, "r", false, "REST login")
	f.DurationVar(&cmd.life, "lifetime", time.Minute*10, "SAML token lifetime")
	f.BoolVar(&cmd.long, "l", false, "Output session cookie")
	f.StringVar(&cmd.ticket, "ticket", "", "Use clone ticket for login")
	f.StringVar(&cmd.cookie, "cookie", "", "Set HTTP cookie for an existing session")
	f.StringVar(&cmd.token, "token", "", "Use SAML token for login or as issue identity")
	f.StringVar(&cmd.ext, "extension", "", "Extension name")
	f.StringVar(&cmd.as, "as", "", "Impersonate user")
	f.StringVar(&cmd.method, "X", "", "HTTP method")
}

func (cmd *login) Process(ctx context.Context) error {
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *login) Usage() string {
	return "[PATH]"
}

func (cmd *login) Description() string {
	return `Session login.

The session.login command is optional, all other govc commands will auto login when given credentials.
The session.login command can be used to:
- Persist a session without writing to disk via the '-cookie' flag
- Acquire a clone ticket
- Login using a clone ticket
- Login using a vCenter Extension certificate
- Issue a SAML token
- Renew a SAML token
- Exchange a SAML token for a JSON Web Token (JWT)
- Login using a SAML token
- Impersonate a user
- Avoid passing credentials to other govc commands
- Send an authenticated raw HTTP request

The session.login command can be used for authenticated curl-style HTTP requests when a PATH arg is given.
PATH may also contain a query string. The '-u' flag (GOVC_URL) is used for the URL scheme, host and port.
The request method (-X) defaults to GET. When set to POST, PUT or PATCH, a request body must be provided via stdin.

Examples:
  govc session.login -u root:password@host # Creates a cached session in ~/.govmomi/sessions
  govc session.ls -u root@host # Use the cached session with another command
  ticket=$(govc session.login -u root@host -clone)
  govc session.login -u root@host -ticket $ticket
  govc session.login -u Administrator@vsphere.local:password@host -as other@vsphere.local
  govc session.login -u host -extension com.vmware.vsan.health -cert rui.crt -key rui.key
  token=$(govc session.login -u host -cert user.crt -key user.key -issue) # HoK token
  bearer=$(govc session.login -u user:pass@host -issue) # Bearer token
  token=$(govc session.login -u host -cert user.crt -key user.key -issue -token "$bearer")
  govc session.login -u host -cert user.crt -key user.key -token "$token"
  token=$(govc session.login -u host -cert user.crt -key user.key -renew -lifetime 24h -token "$token")
  govc session.login -jwt vmware-tes:vc:nsxd-v2:nsx -token "$token"
  # HTTP requests
  govc session.login -r -X GET /api/vcenter/namespace-management/clusters | jq .
  govc session.login -r -X POST /rest/vcenter/cluster/modules <<<'{"spec": {"cluster": "domain-c9"}}'`
}

type ticketResult struct {
	cmd    *login
	Ticket string `json:",omitempty"`
	Token  string `json:",omitempty"`
	Cookie string `json:",omitempty"`
}

func (r *ticketResult) Write(w io.Writer) error {
	var output []string

	for _, val := range []string{r.Ticket, r.Token, r.Cookie} {
		if val != "" {
			output = append(output, val)
		}
	}

	if len(output) == 0 {
		return nil
	}

	fmt.Fprintln(w, strings.Join(output, " "))

	return nil
}

// Logout is called by cli.Run()
// We override ClientFlag's Logout impl to avoid ending a session when -persist-session=false,
// otherwise Logout would invalidate the cookie and/or ticket.
func (cmd *login) Logout(ctx context.Context) error {
	if cmd.long || cmd.clone || cmd.issue {
		return nil
	}
	return cmd.ClientFlag.Logout(ctx)
}

func (cmd *login) cloneSession(ctx context.Context, c *vim25.Client) error {
	return session.NewManager(c).CloneSession(ctx, cmd.ticket)
}

func (cmd *login) issueToken(ctx context.Context, vc *vim25.Client) (string, error) {
	c, err := sts.NewClient(ctx, vc)
	if err != nil {
		return "", err
	}
	c.RoundTripper = cmd.RoundTripper(c.Client)

	req := sts.TokenRequest{
		Certificate: c.Certificate(),
		Userinfo:    cmd.Session.URL.User,
		Renewable:   true,
		Delegatable: true,
		ActAs:       cmd.token != "",
		Token:       cmd.token,
		Lifetime:    cmd.life,
	}

	issue := c.Issue
	if cmd.renew {
		issue = c.Renew
	}

	s, err := issue(ctx, req)
	if err != nil {
		return "", err
	}

	if req.Token != "" {
		duration := s.Lifetime.Expires.Sub(s.Lifetime.Created)
		if duration < req.Lifetime {
			// The granted lifetime is that of the bearer token, which is 5min max.
			// Extend the lifetime via Renew.
			req.Token = s.Token
			if s, err = c.Renew(ctx, req); err != nil {
				return "", err
			}
		}
	}

	return s.Token, nil
}

func (cmd *login) exchangeTokenJWT(ctx context.Context, c *rest.Client) (string, error) {
	spec := authentication.TokenIssueSpec{
		Audience:           cmd.jwt,
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		RequestedTokenType: "urn:ietf:params:oauth:token-type:id_token",
		SubjectToken:       base64.StdEncoding.EncodeToString([]byte(cmd.token)),
		SubjectTokenType:   "urn:ietf:params:oauth:token-type:saml2",
	}

	info, err := authentication.NewManager(c).Issue(ctx, spec)
	if err != nil {
		return "", err
	}
	return info.AccessToken, nil
}

func (cmd *login) loginByToken(ctx context.Context, c *vim25.Client) error {
	header := soap.Header{
		Security: &sts.Signer{
			Certificate: c.Certificate(),
			Token:       cmd.token,
		},
	}

	// something behind the LoginByToken scene requires a version from /sdk/vimServiceVersions.xml
	// in the SOAPAction header. For example, if vim25.Version is "7.0" but the service version is "6.3",
	// LoginByToken fails with: 'VersionMismatchFaultCode: Unsupported version URI "urn:vim25/7.0"'
	if c.Version == vim25.Version {
		_ = c.UseServiceVersion()
	}

	return session.NewManager(c).LoginByToken(c.WithHeader(ctx, header))
}

func (cmd *login) loginRestByToken(ctx context.Context, c *rest.Client) error {
	signer := &sts.Signer{
		Certificate: c.Certificate(),
		Token:       cmd.token,
	}

	return c.LoginByToken(c.WithSigner(ctx, signer))
}

func (cmd *login) loginByExtension(ctx context.Context, c *vim25.Client) error {
	return session.NewManager(c).LoginExtensionByCertificate(ctx, cmd.ext)
}

func (cmd *login) impersonateUser(ctx context.Context, c *vim25.Client) error {
	m := session.NewManager(c)
	if err := m.Login(ctx, cmd.Session.URL.User); err != nil {
		return err
	}
	return m.ImpersonateUser(ctx, cmd.as)
}

func (cmd *login) setCookie(ctx context.Context, c *vim25.Client) error {
	url := c.URL()
	jar := c.Client.Jar
	cookies := jar.Cookies(url)
	add := true

	cookie := &http.Cookie{
		Name: soap.SessionCookieName,
	}

	for _, e := range cookies {
		if e.Name == cookie.Name {
			add = false
			cookie = e
			break
		}
	}

	if cmd.cookie == "" {
		// This is the cookie from Set-Cookie after a Login or CloneSession
		cmd.cookie = cookie.Value
	} else {
		// The cookie flag is set, set the HTTP header and skip Login()
		cookie.Value = cmd.cookie
		if add {
			cookies = append(cookies, cookie)
		}
		jar.SetCookies(url, cookies)

		// Check the session is still valid
		_, err := methods.GetCurrentTime(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *login) setRestCookie(ctx context.Context, c *rest.Client) error {
	if cmd.cookie == "" {
		cmd.cookie = c.SessionID()
	} else {
		c.SessionID(cmd.cookie)

		// Check the session is still valid
		s, err := c.Session(ctx)
		if err != nil {
			return err
		}
		if s == nil {
			return errors.New(http.StatusText(http.StatusUnauthorized))
		}
	}

	return nil
}

func nologinSOAP(_ context.Context, _ *vim25.Client) error {
	return nil
}

func nologinREST(_ context.Context, _ *rest.Client) error {
	return nil
}

func (cmd *login) Run(ctx context.Context, f *flag.FlagSet) error {
	if cmd.renew {
		cmd.issue = true
	}
	switch {
	case cmd.ticket != "":
		cmd.Session.LoginSOAP = cmd.cloneSession
	case cmd.cookie != "":
		if cmd.vapi {
			cmd.Session.LoginSOAP = nologinSOAP
			cmd.Session.LoginREST = cmd.setRestCookie
		} else {
			cmd.Session.LoginSOAP = cmd.setCookie
			cmd.Session.LoginREST = nologinREST
		}
	case cmd.token != "":
		cmd.Session.LoginSOAP = cmd.loginByToken
		cmd.Session.LoginREST = cmd.loginRestByToken
	case cmd.ext != "":
		cmd.Session.LoginSOAP = cmd.loginByExtension
	case cmd.as != "":
		cmd.Session.LoginSOAP = cmd.impersonateUser
	case cmd.issue:
		cmd.Session.LoginSOAP = nologinSOAP
		cmd.Session.LoginREST = nologinREST
	case cmd.jwt != "":
		cmd.Session.LoginSOAP = nologinSOAP
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	r := &ticketResult{cmd: cmd}

	var rc *rest.Client
	if cmd.vapi || cmd.jwt != "" {
		rc, err = cmd.RestClient()
		if err != nil {
			return err
		}
	}

	switch {
	case cmd.clone:
		m := session.NewManager(c)
		r.Ticket, err = m.AcquireCloneTicket(ctx)
		if err != nil {
			return err
		}
	case cmd.issue:
		r.Token, err = cmd.issueToken(ctx, c)
		if err != nil {
			return err
		}
		return cmd.WriteResult(r)
	case cmd.jwt != "":
		r.Token, err = cmd.exchangeTokenJWT(ctx, rc)
	}

	if f.NArg() == 1 {
		u, err := url.Parse(f.Arg(0))
		if err != nil {
			return err
		}
		vc := c.URL()
		u.Scheme = vc.Scheme
		u.Host = vc.Host

		var body io.Reader

		switch cmd.method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			// strings.Reader here as /api wants a Content-Length header
			b, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			body = bytes.NewReader(b)
		default:
			body = strings.NewReader("")
		}

		req, err := http.NewRequest(cmd.method, u.String(), body)
		if err != nil {
			return err
		}

		if cmd.vapi {
			return rc.Do(ctx, req, cmd.Out)
		}

		return c.Do(ctx, req, func(res *http.Response) error {
			if res.StatusCode != http.StatusOK {
				return errors.New(res.Status)
			}
			_, err := io.Copy(cmd.Out, res.Body)
			return err
		})
	}

	if cmd.cookie == "" {
		if cmd.vapi {
			_ = cmd.setRestCookie(ctx, rc)
		} else {
			_ = cmd.setCookie(ctx, c)
		}
		if cmd.cookie == "" {
			return flag.ErrHelp
		}
	}

	if cmd.long {
		r.Cookie = cmd.cookie
	}

	return cmd.WriteResult(r)
}
