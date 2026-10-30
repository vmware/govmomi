// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/sts/internal"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			path, handler := New(s.Listen, r.OptionManager().Setting)
			if handler == nil {
				return
			}
			s.Handle(path, handler)
			s.Handle(sts.SystemPath, handler)
			r.SessionManager().ValidToken = handler.(*Handler).validToken
		}
	})
}

// Handler is the STS simulator. It issues SAML tokens signed by its own key, and validates them.
// Neither requests nor responses are logged, as they carry credentials.
type Handler struct {
	// URL is the scheme and host of the SSO server, as clients reach it.
	URL url.URL
	// Domain is the SSO domain, for example "vsphere.local".
	Domain string

	mu     sync.Mutex
	key    *rsa.PrivateKey
	cert   *x509.Certificate
	minted []string
}

// New creates an STS simulator and configures the simulator endpoint in the given settings.
// The path returned is that of the settings "config.vpxd.sso.sts.uri" property.
// The http.Handler returned is a *Handler.
func New(u *url.URL, settings []vim.BaseOptionValue) (string, http.Handler) {
	for i := range settings {
		setting := settings[i].GetOptionValue()
		if setting.Key == "config.vpxd.sso.sts.uri" {
			endpoint, _ := url.Parse(setting.Value.(string))
			h := &Handler{
				URL:    url.URL{Scheme: u.Scheme, Host: u.Host},
				Domain: path.Base(endpoint.Path),
			}
			return endpoint.Path, h
		}
	}
	return "", nil
}

// Lookup returns the STS simulator registered with s, or nil if there is none.
func Lookup(s *simulator.Service) *Handler {
	h, _ := s.ServeMux.Handler(&http.Request{URL: &url.URL{Path: sts.SystemPath}})
	sim, _ := h.(*Handler)
	return sim
}

// Minted returns every credential the simulator has issued, in the form it was issued.
// Tests use them as canaries: none of them should appear in a log.
func (s *Handler) Minted() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.minted...)
}

func (s *Handler) mint(credential string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.minted = append(s.minted, credential)
}

// signer returns the key and certificate the simulator signs tokens with, generating them on first use.
func (s *Handler) signer() (*rsa.PrivateKey, *x509.Certificate, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.key != nil {
		return s.key, s.cert, nil
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	now := time.Now()
	tmpl := x509.Certificate{
		SerialNumber: big.NewInt(now.UnixNano()),
		Subject:      pkix.Name{CommonName: "ssoserverSign"},
		NotBefore:    now.Add(-time.Hour),
		NotAfter:     now.AddDate(10, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	if err != nil {
		return nil, nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, nil, err
	}
	s.key, s.cert = key, cert
	return key, cert, nil
}

// verify validates a SAML assertion issued by the simulator.
func (s *Handler) verify(assertion string) (*token, error) {
	s.mu.Lock()
	key := s.key
	s.mu.Unlock()
	if key == nil {
		return nil, errInvalidToken // no token has been issued
	}
	return verify(assertion, &key.PublicKey, time.Now())
}

// validToken implements simulator.SessionManager.ValidToken.
func (s *Handler) validToken(assertion string) (*simulator.TokenIdentity, error) {
	t, err := s.verify(assertion)
	if err != nil {
		return nil, err
	}
	return &simulator.TokenIdentity{Subject: t.Subject, Certificate: t.Certificate}, nil
}

// request is the part of a WS-Trust request that the simulator uses.
type request struct {
	Header struct {
		Security struct {
			BinarySecurityToken string `xml:"BinarySecurityToken"`
			Username            string `xml:"UsernameToken>Username"`
			Content             string `xml:",innerxml"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		RST struct {
			Created     string `xml:"Lifetime>Created"`
			Expires     string `xml:"Lifetime>Expires"`
			KeyType     string `xml:"KeyType"`
			ActAs       *inner `xml:"ActAs"`
			RenewTarget *inner `xml:"RenewTarget"`
		} `xml:"RequestSecurityToken"`
	} `xml:"Body"`
}

type inner struct {
	Content string `xml:",innerxml"`
}

// rawElement returns the first element named local in data, as the bytes that encode it, or "" if there is none.
func rawElement(data, local string) string {
	dec := xml.NewDecoder(strings.NewReader(data))
	for {
		start := dec.InputOffset()
		tok, err := dec.Token()
		if err != nil {
			return ""
		}
		if e, ok := tok.(xml.StartElement); ok && e.Name.Local == local {
			if err := dec.Skip(); err != nil {
				return ""
			}
			return data[start:dec.InputOffset()]
		}
	}
}

// fault is a WS-Trust fault, returned to the client as a SOAP fault.
type fault struct {
	code, message string
}

func (f *fault) Error() string {
	return f.message
}

var (
	errAuthentication = &fault{"wst:FailedAuthentication", "Authentication failed"}
	errRequest        = &fault{"wst:InvalidRequest", "Invalid request"}
)

// defaultGroups are those of the default Administrator user.
var defaultGroups = []string{
	"Users", "Administrators", "CAAdmins", "ComponentManager.Administrators",
	"SystemConfiguration.BashShellAdministrators", "SystemConfiguration.Administrators",
	"LicenseService.Administrators", "ActAsUsers", "Everyone",
}

// principal returns the identity a request authenticates, and the holder-of-key certificate, if any, it asks to
// confirm the token with. A token presented to renew, to act as, or to authenticate the request must be one the
// simulator issued. Passwords are not checked.
func (s *Handler) principal(req *request) (*token, error) {
	security := &req.Header.Security
	rst := &req.Body.RST

	var cert *x509.Certificate
	if security.BinarySecurityToken != "" {
		der, err := base64.StdEncoding.DecodeString(strings.TrimSpace(security.BinarySecurityToken))
		if err != nil {
			return nil, errRequest
		}
		if cert, err = x509.ParseCertificate(der); err != nil {
			return nil, errRequest
		}
	}

	var header *token // a token that authenticates the request, for a token holder without a certificate
	if assertion := rawElement(security.Content, "Assertion"); assertion != "" {
		t, err := s.verify(assertion)
		if err != nil {
			return nil, errAuthentication
		}
		header = t
		if cert == nil {
			cert = t.Certificate
		}
	}

	var t *token
	switch {
	case rst.RenewTarget != nil:
		renew, err := s.verify(rawElement(rst.RenewTarget.Content, "Assertion"))
		if err != nil {
			return nil, errAuthentication
		}
		return renew, nil // renewal keeps the token's identity and confirmation
	case rst.ActAs != nil:
		actas, err := s.verify(rawElement(rst.ActAs.Content, "Assertion"))
		if err != nil {
			return nil, errAuthentication
		}
		t = &token{Subject: actas.Subject, Groups: actas.Groups, Solution: actas.Solution}
	case security.Username != "":
		t = &token{Subject: security.Username}
	case security.BinarySecurityToken != "":
		// A solution user authenticates by certificate: its name is the certificate's CN.
		name := cert.Subject.CommonName
		if name == "" {
			return nil, errAuthentication
		}
		t = &token{Subject: name, Solution: true}
	case header != nil:
		t = &token{Subject: header.Subject, Groups: header.Groups, Solution: header.Solution}
	default:
		return nil, errRequest
	}

	if !strings.Contains(t.Subject, "@") {
		t.Subject += "@" + s.Domain
	}
	if t.Groups == nil {
		groups := []string{"Users", "Everyone"}
		switch {
		case strings.EqualFold(t.Subject, "Administrator@"+s.Domain):
			groups = defaultGroups
		case t.Solution:
			groups = []string{"Users", "SolutionUsers", "Everyone"}
		}
		for _, g := range groups {
			t.Groups = append(t.Groups, s.Domain+`\`+g)
		}
	}
	if rst.KeyType != "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Bearer" {
		t.Certificate = cert
	}

	return t, nil
}

// issue returns a signed token for the principal req authenticates, with the lifetime it requests.
func (s *Handler) issue(req *request) (string, *internal.Lifetime, error) {
	t, err := s.principal(req)
	if err != nil {
		return "", nil, err
	}

	lifetime := 5 * time.Minute
	created, cerr := time.Parse(internal.Time, req.Body.RST.Created)
	expires, eerr := time.Parse(internal.Time, req.Body.RST.Expires)
	if cerr == nil && eerr == nil && expires.After(created) {
		lifetime = expires.Sub(created)
	}
	now := time.Now().UTC()
	t.NotBefore, t.NotOnOrAfter = now, now.Add(lifetime)

	key, cert, err := s.signer()
	if err != nil {
		return "", nil, err
	}
	issuer := s.URL.JoinPath("websso", "SAML2", "Metadata", s.Domain).String()
	assertion, err := t.sign(issuer, key, cert)
	if err != nil {
		return "", nil, err
	}
	s.mint(assertion)

	return assertion, &internal.Lifetime{
		Created: t.NotBefore.Format(internal.Time),
		Expires: t.NotOnOrAfter.Format(internal.Time),
	}, nil
}

// ServeHTTP handles STS requests.
func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("SOAPAction")
	action = strings.TrimSuffix(action, `"`) // PowerCLI sts client quotes the header value

	env := soap.Envelope{}
	kind := path.Base(action)
	if kind != "Issue" && kind != "Renew" {
		log.Printf("sts: unsupported action=%s", action)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var req request
	if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
		s.fault(w, errRequest)
		return
	}

	assertion, lifetime, err := s.issue(&req)
	if err != nil {
		s.fault(w, err)
		return
	}

	res := internal.RequestSecurityTokenResponse{
		RequestedSecurityToken: internal.RequestedSecurityToken{
			Assertion: assertion,
		},
		Lifetime: lifetime,
	}

	switch kind {
	case "Issue":
		env.Body = internal.RequestSecurityTokenBody{
			Res: &internal.RequestSecurityTokenResponseCollection{
				RequestSecurityTokenResponse: res,
			},
		}
	case "Renew":
		env.Body = internal.RenewSecurityTokenBody{
			Res: &res,
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, internal.Marshal(env))
}

// fault responds with a SOAP fault for err. The fault never includes the request.
func (s *Handler) fault(w http.ResponseWriter, err error) {
	var f *fault
	if !errors.As(err, &f) {
		f = &fault{"wst:RequestFailed", "The request failed"}
	}
	env := soap.Envelope{
		Body: internal.RequestSecurityTokenBody{
			Fault_: &soap.Fault{Code: f.code, String: f.message},
		},
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, internal.Marshal(env))
}
