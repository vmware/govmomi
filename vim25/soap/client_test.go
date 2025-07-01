// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

type mockRT struct{}

func (m mockRT) RoundTrip(req *http.Request) (*http.Response, error) {
	var res http.Response
	res.Header = req.Header.Clone()
	return &res, nil
}

func TestUserAgent(t *testing.T) {
	tests := []struct {
		name  string
		agent string
	}{
		{name: "default agent", agent: ""},
		{name: "custom agent", agent: "govmomi-test/0.0.0"},
	}

	const rawURL = "https://vcenter.local"
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("parse url: %v", err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClient(u, false)
			c.Transport = &mockRT{}

			req, err := http.NewRequest(http.MethodPost, rawURL, nil)
			if err != nil {
				t.Fatalf("create request: %v", err)
			}

			if test.agent != "" {
				c.UserAgent = test.agent
			}

			if err = c.Do(context.Background(), req, func(response *http.Response) error {
				got := response.Header.Get("user-agent")
				want := func() string {
					if test.agent == "" {
						return defaultUserAgent
					}
					return test.agent
				}

				if got != want() {
					return fmt.Errorf("user-agent header mismatch: got=%s want=%s", got, want())
				}
				return nil
			}); err != nil {
				t.Errorf("user-agent header validation failed: %v", err)
			}
		})
	}
}

func TestSplitHostPort(t *testing.T) {
	tests := []struct {
		url  string
		host string
		port string
	}{
		{"127.0.0.1", "127.0.0.1", ""},
		{"*:1234", "*", "1234"},
		{"127.0.0.1:80", "127.0.0.1", "80"},
		{"[::1]:6767", "[::1]", "6767"},
		{"[::1]", "[::1]", ""},
	}

	for _, test := range tests {
		host, port := splitHostPort(test.url)
		if host != test.host {
			t.Errorf("(%s) %s != %s", test.url, host, test.host)
		}
		if port != test.port {
			t.Errorf("(%s) %s != %s", test.url, port, test.port)
		}
	}
}

func TestMultipleCAPaths(t *testing.T) {
	err := setCAsOnClient("fixtures/invalid-cert.pem:fixtures/valid-cert.pem")

	certErr, ok := err.(errInvalidCACertificate)
	if !ok {
		t.Fatalf("Expected errInvalidCACertificate to occur")
	}
	if certErr.File != "fixtures/invalid-cert.pem" {
		t.Fatalf("Expected Err to show invalid file")
	}
}

func TestInvalidRootCAPath(t *testing.T) {
	err := setCAsOnClient("fixtures/there-is-no-such-file")

	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("os.PathError should have occurred: %#v", err)
	}
}

func TestValidRootCAs(t *testing.T) {
	err := setCAsOnClient("fixtures/valid-cert.pem")
	if err != nil {
		t.Fatalf("Err should not have occurred: %#v", err)
	}
}

func TestInvalidRootCAs(t *testing.T) {
	err := setCAsOnClient("fixtures/invalid-cert.pem")

	certErr, ok := err.(errInvalidCACertificate)
	if !ok {
		t.Fatalf("Expected errInvalidCACertificate to occur")
	}
	if certErr.File != "fixtures/invalid-cert.pem" {
		t.Fatalf("Expected Err to show invalid file")
	}
}

func setCAsOnClient(cas string) error {
	url := &url.URL{
		Scheme: "https",
		Host:   "some.host.tld:8080",
	}
	insecure := false

	client := NewClient(url, insecure)

	return client.SetRootCAs(cas)
}

func TestParseURL(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    *url.URL
		wantErr bool
	}{
		{
			name: "empty URL should return null",
			want: nil,
		},
		{
			name: "just endpoint should return full URL",
			s:    "some.vcenter.tld",
			want: &url.URL{
				Scheme: "https",
				Path:   "/sdk",
				Host:   "some.vcenter.tld",
				User:   url.UserPassword("", ""),
			},
		},
		{
			name: "URL with / on suffix should be trimmed",
			s:    "https://some.vcenter.tld/",
			want: &url.URL{
				Scheme: "https",
				Path:   "/sdk",
				Host:   "some.vcenter.tld",
				User:   url.UserPassword("", ""),
			},
		},
		{
			name: "URL with user and password should be used",
			s:    "https://user:password@some.vcenter.tld",
			want: &url.URL{
				Scheme: "https",
				Path:   "/sdk",
				Host:   "some.vcenter.tld",
				User:   url.UserPassword("user", "password"),
			},
		},
		{
			name: "existing path should be used",
			s:    "https://some.vcenter.tld/othersdk",
			want: &url.URL{
				Scheme: "https",
				Path:   "/othersdk",
				Host:   "some.vcenter.tld",
				User:   url.UserPassword("", ""),
			},
		},
		{
			name:    "Invalid URL should be rejected",
			s:       "https://user:password@some.vcenter.tld:xpto1234",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionCookie(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:1080",
		Path:   "sdk", // see comment in Client.SessionCookie
	}

	c := NewClient(u, true)

	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    "ANY",
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   false,
	}

	c.Jar.SetCookies(u, []*http.Cookie{cookie})

	val := c.SessionCookie()
	if val == nil {
		t.Fatal("no session cookie")
	}
}
