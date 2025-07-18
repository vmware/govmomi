// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// HostCertificateInfo provides helpers for types.HostCertificateManagerCertificateInfo
type HostCertificateInfo struct {
	types.HostCertificateManagerCertificateInfo

	ThumbprintSHA1   string `json:"thumbprintSHA1"`
	ThumbprintSHA256 string `json:"thumbprintSHA256"`

	Err         error             `json:"err"`
	Certificate *x509.Certificate `json:"-"`

	subjectName *pkix.Name
	issuerName  *pkix.Name
}

// FromCertificate converts x509.Certificate to HostCertificateInfo
func (info *HostCertificateInfo) FromCertificate(cert *x509.Certificate) *HostCertificateInfo {
	info.Certificate = cert
	info.subjectName = &cert.Subject
	info.issuerName = &cert.Issuer

	info.Issuer = info.fromName(info.issuerName)
	info.NotBefore = &cert.NotBefore
	info.NotAfter = &cert.NotAfter
	info.Subject = info.fromName(info.subjectName)

	info.ThumbprintSHA1 = soap.ThumbprintSHA1(cert)
	info.ThumbprintSHA256 = soap.ThumbprintSHA256(cert)

	if info.Status == "" {
		info.Status = string(types.HostCertificateManagerCertificateInfoCertificateStatusUnknown)
	}

	return info
}

func (info *HostCertificateInfo) FromPEM(cert []byte) (*HostCertificateInfo, error) {
	block, _ := pem.Decode(cert)
	if block == nil {
		return nil, errors.New("failed to pem.Decode cert")
	}
	x, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return info.FromCertificate(x), nil
}

// FromURL connects to the given URL.Host via tls.Dial with the given tls.Config and populates the HostCertificateInfo
// via tls.ConnectionState.  If the certificate was verified with the given tls.Config, the Err field will be nil.
// Otherwise, Err will be set to the x509.UnknownAuthorityError or x509.HostnameError.
// If tls.Dial returns an error of any other type, that error is returned.
func (info *HostCertificateInfo) FromURL(u *url.URL, config *tls.Config) error {
	addr := u.Host
	if !(strings.LastIndex(addr, ":") > strings.LastIndex(addr, "]")) {
		addr += ":443"
	}

	conn, err := tls.Dial("tcp", addr, config)
	if err != nil {
		if !soap.IsCertificateUntrusted(err) {
			return err
		}

		info.Err = err

		conn, err = tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return err
		}
	} else {
		info.Status = string(types.HostCertificateManagerCertificateInfoCertificateStatusGood)
	}

	state := conn.ConnectionState()
	_ = conn.Close()
	info.FromCertificate(state.PeerCertificates[0])

	return nil
}

var emailAddressOID = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

func (info *HostCertificateInfo) fromName(name *pkix.Name) string {
	var attrs []string

	oids := map[string]string{
		emailAddressOID.String(): "emailAddress",
	}

	for _, attr := range name.Names {
		if key, ok := oids[attr.Type.String()]; ok {
			attrs = append(attrs, fmt.Sprintf("%s=%s", key, attr.Value))
		}
	}

	attrs = append(attrs, fmt.Sprintf("CN=%s", name.CommonName))

	add := func(key string, vals []string) {
		for _, val := range vals {
			attrs = append(attrs, fmt.Sprintf("%s=%s", key, val))
		}
	}

	elts := []struct {
		key string
		val []string
	}{
		{"OU", name.OrganizationalUnit},
		{"O", name.Organization},
		{"L", name.Locality},
		{"ST", name.Province},
		{"C", name.Country},
	}

	for _, elt := range elts {
		add(elt.key, elt.val)
	}

	return strings.Join(attrs, ",")
}

func (info *HostCertificateInfo) toName(s string) *pkix.Name {
	var name pkix.Name

	for _, pair := range strings.Split(s, ",") {
		attr := strings.SplitN(pair, "=", 2)
		if len(attr) != 2 {
			continue
		}

		v := attr[1]

		switch strings.ToLower(attr[0]) {
		case "cn":
			name.CommonName = v
		case "ou":
			name.OrganizationalUnit = append(name.OrganizationalUnit, v)
		case "o":
			name.Organization = append(name.Organization, v)
		case "l":
			name.Locality = append(name.Locality, v)
		case "st":
			name.Province = append(name.Province, v)
		case "c":
			name.Country = append(name.Country, v)
		case "emailaddress":
			name.Names = append(name.Names, pkix.AttributeTypeAndValue{Type: emailAddressOID, Value: v})
		}
	}

	return &name
}

// SubjectName parses Subject into a pkix.Name
func (info *HostCertificateInfo) SubjectName() *pkix.Name {
	if info.subjectName != nil {
		return info.subjectName
	}

	return info.toName(info.Subject)
}

// IssuerName parses Issuer into a pkix.Name
func (info *HostCertificateInfo) IssuerName() *pkix.Name {
	if info.issuerName != nil {
		return info.issuerName
	}

	return info.toName(info.Issuer)
}

// Write outputs info similar to the Chrome Certificate Viewer.
func (info *HostCertificateInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	s := func(val string) string {
		if val != "" {
			return val
		}
		return "<Not Part Of Certificate>"
	}

	ss := func(val []string) string {
		return s(strings.Join(val, ","))
	}

	name := func(n *pkix.Name) {
		fmt.Fprintf(tw, "  Common Name (CN):\t%s\n", s(n.CommonName))
		fmt.Fprintf(tw, "  Organization (O):\t%s\n", ss(n.Organization))
		fmt.Fprintf(tw, "  Organizational Unit (OU):\t%s\n", ss(n.OrganizationalUnit))
	}

	status := info.Status
	if info.Err != nil {
		status = fmt.Sprintf("ERROR %s", info.Err)
	}
	fmt.Fprintf(tw, "Certificate Status:\t%s\n", status)

	fmt.Fprintln(tw, "Issued To:\t")
	name(info.SubjectName())

	fmt.Fprintln(tw, "Issued By:\t")
	name(info.IssuerName())

	fmt.Fprintln(tw, "Validity Period:\t")
	fmt.Fprintf(tw, "  Issued On:\t%s\n", info.NotBefore)
	fmt.Fprintf(tw, "  Expires On:\t%s\n", info.NotAfter)

	if info.ThumbprintSHA1 != "" {
		fmt.Fprintln(tw, "Thumbprints:\t")
		if info.ThumbprintSHA256 != "" {
			fmt.Fprintf(tw, "  SHA-256 Thumbprint:\t%s\n", info.ThumbprintSHA256)
		}
		fmt.Fprintf(tw, "  SHA-1 Thumbprint:\t%s\n", info.ThumbprintSHA1)
	}

	return tw.Flush()
}
