// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package sts

import (
	"crypto/tls"
	"net/http"
	"strings"
	"testing"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

var _ soap.Signer = new(Signer)

// LocalhostCert copied from simulator/internal/testcert.go
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDOjCCAiKgAwIBAgIRAK6d/JpGL75P2wOSYc6WalEwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBAK9lInM4OAKI4z+wDNkvcZC6S6exFSOysp7NPJyaEAhW93kPY7gO
f6H5aP3V3YU0vYpCSnz/UhyDD+/knBof1J3Do7FVwYtC293vrtXffNtAvygfJodW
1dPllp17ZJbq76ei9oWq1Y5Hox/sVYmNVNztvBfK1mtpS8z8Qrk1LWCyLiDHkvDA
hCy2OjuaopxC6qQejdWT1PxwbqptuLVakQmecpiFrupy8DTG0x0rxxdMdAATywhY
Gm49A/FroagZ6HMz3bm39we/w6VIx3pX1lbUUyrfjvBgfUlRwxyZABBj2STGsOQJ
a451eEcESXcSEWzjGjUQ1Wf+zzxr2GAHmI8CAwEAAaOBiDCBhTAOBgNVHQ8BAf8E
BAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAdBgNV
HQ4EFgQUjtR2VSuxchTxe0UNDVqWDNMR37AwLgYDVR0RBCcwJYILZXhhbXBsZS5j
b22HBH8AAAGHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZIhvcNAQELBQADggEBACK7
1L15IeYPiQHFDml3EgDoRnd/tAaZP9nUZIIPLRgUnQITNAtjFgBvqneQZwlQtco2
s8YXSivBQiATlBEtGBzVxv36R4tTXldIgJxaCUxxZtqALLwyGqSaI/cwE0pWa6Z0
Op2wkzUmoQ5rRrJfRM+C6HR/+lWkNtHRzaUFOSlxNJbPo53K53OoDriwEc1PvEYP
wFeUXwTzCZ68pAlWUmDKCyp+lPhjIt2Gznig+BSPCNJqmwKM76oFyywi3HIP56rD
/cwUtoplF68uVuD8HXb1ggGsqtGiAT4GLT8tU5w+BtK8ZIs/LK7mdi7W8aIOhUUH
l1lgeV3oEQue3A7SogA=
-----END CERTIFICATE-----`)

// LocalhostKey copied from simulator/internal/testcert.go
var LocalhostKey = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCvZSJzODgCiOM/
sAzZL3GQukunsRUjsrKezTycmhAIVvd5D2O4Dn+h+Wj91d2FNL2KQkp8/1Icgw/v
5JwaH9Sdw6OxVcGLQtvd767V33zbQL8oHyaHVtXT5Zade2SW6u+novaFqtWOR6Mf
7FWJjVTc7bwXytZraUvM/EK5NS1gsi4gx5LwwIQstjo7mqKcQuqkHo3Vk9T8cG6q
bbi1WpEJnnKYha7qcvA0xtMdK8cXTHQAE8sIWBpuPQPxa6GoGehzM925t/cHv8Ol
SMd6V9ZW1FMq347wYH1JUcMcmQAQY9kkxrDkCWuOdXhHBEl3EhFs4xo1ENVn/s88
a9hgB5iPAgMBAAECggEAEpeS3knQThx6kk60HfWUgTXuPRldV0pi+shgq2z9VBT7
6J5EAMewqdfJVFbuQ2eCy/wY70UVTCZscw51qaNEI3EQkgS4Hm345n64tr0Y/BjR
6ovaxq/ivLJyk8D3ubOvscJphWPFfW6EkSa5LnqHy197972tmvcvbMw0unMzmzM7
DenXdoIJQu1SqLiLUiDXEfkvCReekqhe1jATCwTzIBTCnTWxgI4Ox2qsBaxuwrnl
D1GpWy4sh8NpDB0EBwdrjAOmLDOyvsy2X65DIlHS/k7901tvzyNjRrsr2Ig0sAz4
w0ke6CUKQ2B+Pqn3p6bvxRYMP08+ZjlQpPuU4RrxGQKBgQDd3HCrZCgUJAGcfzYX
ZzSmSoxB9+sEuVUZGU+ICMPlG0Dd8aEkXIyDjGMWsLFMIvzNBf4wB1FwdLaCJu6y
0PbX3KVfg/Yc4lvYUuQ+1nD/3gm2hE46lZuSfbmExH5SQVLSbSQf9S/5BTHAWQO9
PNie71AZ8fO5YDBM18tq2V7dBQKBgQDKYk1+Zup5p0BeRyCNnVYnpngO+pAZZmld
gYiRn8095QJ/Y+yV0kYc2+0+5f6St462E+PVH3tX1l9OG3ujVQnWfu2iSsRJRMUf
0blxqTWvqdcGi8SLpVjkrHn30scFNWmojhJv3k42H3nUMC1+WU3rp2f7+W58afyd
NY9x4sqzgwKBgQCoeMq9+3JLyQPIOPl0UBS06gsT1RUMI0gxpPy1yiInich6QRAi
snypMCPWiRo5PKBHd/OLuSLoiFhHARVliDTJum2B2I09Zc5kuJ1F8kUgpxUtGc7l
wdG/LeWAok1iXORtkh9KfT+Ok5kx/OZP/zJnjkZ/TTHMZPSIhZ2cZ7AXmQKBgHMP
HjWNtyKApsSytVwtpgyWxMznQMNgCOkjOoxoCJx2tUvNeHTY/glsM14+DdRFzTnQ
5weEhXAzrS1PzKPYNeafdOR+k0eAdH2Zk09+PspmyZusHIqz72zabeEqEQHyEubE
FtFI1rhIfs/WsBaUGQuvuhtz/I95BiguiiXaJRmXAoGADwcO6YXoWXga07gGRwZP
LYKwt5wBh13LAGbSsUyCSK5FG6ZrTmzaFdAGV1U4wc/wgiIgv33m8BG4Ikxvpa0r
Wg3dbhBx9Oya8QWIMBPk72KKEzsSDfi+Cn52ZmxTkWbBDCnkRhG77Ooi8vJ3dhq4
fHeAu1F9OwF83SBi1oNySd8=
-----END PRIVATE KEY-----`)

func TestSigner(t *testing.T) {
	cert, err := tls.X509KeyPair(LocalhostCert, LocalhostKey)
	if err != nil {
		t.Fatal(err)
	}

	s := &Signer{}

	env := soap.Envelope{
		Header: new(soap.Header),
		Body: &methods.LoginByTokenBody{
			Req: &types.LoginByToken{
				This: types.ManagedObjectReference{
					Type:  "SessionManager",
					Value: "SessionManager",
				},
				Locale: session.Locale,
			},
		},
	}

	_, err = s.Sign(env)
	if err != nil {
		t.Error(err)
	}

	s.Certificate = &cert

	_, err = s.Sign(env)
	if err == nil {
		t.Error("expected error")
	}

	// Sign() just needs to parse the Assertion.ID
	s.Token = `<saml2:Assertion ID="tokenID"></saml2:Assertion>`

	b, err := s.Sign(env)
	if err != nil {
		t.Error(err)
	}

	// TODO: new(internal.Security) would require more use of xml.Name instead of field tags
	security := struct {
		InnerXML string `xml:",innerxml"`
	}{}

	senv := soap.Envelope{
		Header: &soap.Header{
			Security: &security,
		},
	}
	err = xml.Unmarshal(b, &senv)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(security.InnerXML, "<wsu:Timestamp ") {
		t.Error("missing timestamp")
	}

	if !strings.HasSuffix(security.InnerXML, "</ds:Signature>") {
		t.Error("missing signature")
	}

	req, _ := http.NewRequest(http.MethodPost, "https://127.0.0.1", nil)
	err = s.SignRequest(req)
	if err != nil {
		t.Fatal(err)
	}

	req, _ = http.NewRequest(http.MethodPost, "https://[0:0:0:0:0:0:0:1]", nil)
	err = s.SignRequest(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		ip string
		v6 bool
	}{
		{"0:0:0:0:0:0:0:1", true},
		{"10.0.0.42", false},
	}

	for _, test := range tests {
		v6 := isIPv6(test.ip)
		if v6 != test.v6 {
			t.Errorf("%s: expected %t, got %t", test.ip, test.v6, v6)
		}
	}
}
