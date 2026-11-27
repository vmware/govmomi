// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/sts/internal"
)

const (
	samlNS = "urn:oasis:names:tc:SAML:2.0:assertion"
	xsNS   = "http://www.w3.org/2001/XMLSchema"
	excC14 = "http://www.w3.org/2001/10/xml-exc-c14n#"

	confirmBearer = "urn:oasis:names:tc:SAML:2.0:cm:bearer"
	confirmHoK    = "urn:oasis:names:tc:SAML:2.0:cm:holder-of-key"

	groupsAttr   = "http://rsa.com/schemas/attr-names/2009/01/GroupIdentity"
	solutionAttr = "http://vmware.com/schemas/attr-names/2011/07/isSolution"

	// clockSkew is the tolerance applied to a token's validity period.
	clockSkew = time.Minute
)

// token is the content of a SAML assertion that the simulator issues and validates.
type token struct {
	Subject      string
	Groups       []string // Groups are "<domain>\<name>", as in the GroupIdentity attribute
	Solution     bool     // Solution is true when Subject is a solution user
	Certificate  *x509.Certificate
	NotBefore    time.Time
	NotOnOrAfter time.Time
}

// escape escapes s as canonical XML character data, or as an attribute value when attr is true.
func escape(s string, attr bool) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c == '&':
			b.WriteString("&amp;")
		case c == '<':
			b.WriteString("&lt;")
		case c == '>' && !attr:
			b.WriteString("&gt;")
		case c == '"' && attr:
			b.WriteString("&quot;")
		case c == '\t' && attr:
			b.WriteString("&#x9;")
		case c == '\n' && attr:
			b.WriteString("&#xA;")
		case c == '\r':
			b.WriteString("&#xD;")
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

// assertion returns t as a saml2:Assertion, less its signature, and the index where the signature is placed.
// The assertion is written in exclusive canonical form (xml-exc-c14n, with the xs and xsi prefixes inclusive), so
// that its bytes are the input to the signature's digest, and the signature can be checked without an XML
// canonicalizer.
func (t *token) assertion(id, issuer string, issued time.Time) (string, int) {
	now := issued.UTC().Format(internal.Time)
	var b strings.Builder

	fmt.Fprintf(&b, `<saml2:Assertion xmlns:saml2="%s" xmlns:xs="%s" xmlns:xsi="%s" ID="%s" IssueInstant="%s" Version="2.0">`,
		samlNS, xsNS, internal.XSI, id, now)
	fmt.Fprintf(&b, `<saml2:Issuer Format="urn:oasis:names:tc:SAML:2.0:nameid-format:entity">%s</saml2:Issuer>`,
		escape(issuer, false))
	pos := b.Len()

	b.WriteString(`<saml2:Subject>`)
	fmt.Fprintf(&b, `<saml2:NameID Format="http://schemas.xmlsoap.org/claims/UPN">%s</saml2:NameID>`,
		escape(t.Subject, false))
	if t.Certificate != nil {
		fmt.Fprintf(&b, `<saml2:SubjectConfirmation Method="%s">`, confirmHoK)
		b.WriteString(`<saml2:SubjectConfirmationData xsi:type="saml2:KeyInfoConfirmationDataType">`)
		fmt.Fprintf(&b, `<ds:KeyInfo xmlns:ds="%s"><ds:X509Data><ds:X509Certificate>%s</ds:X509Certificate></ds:X509Data></ds:KeyInfo>`,
			internal.DSIG, base64.StdEncoding.EncodeToString(t.Certificate.Raw))
	} else {
		fmt.Fprintf(&b, `<saml2:SubjectConfirmation Method="%s">`, confirmBearer)
		fmt.Fprintf(&b, `<saml2:SubjectConfirmationData NotOnOrAfter="%s">`, t.NotOnOrAfter.UTC().Format(internal.Time))
	}
	b.WriteString(`</saml2:SubjectConfirmationData></saml2:SubjectConfirmation></saml2:Subject>`)

	fmt.Fprintf(&b, `<saml2:Conditions NotBefore="%s" NotOnOrAfter="%s"></saml2:Conditions>`,
		t.NotBefore.UTC().Format(internal.Time), t.NotOnOrAfter.UTC().Format(internal.Time))
	fmt.Fprintf(&b, `<saml2:AuthnStatement AuthnInstant="%s"><saml2:AuthnContext>`, now)
	b.WriteString(`<saml2:AuthnContextClassRef>urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport</saml2:AuthnContextClassRef>`)
	b.WriteString(`</saml2:AuthnContext></saml2:AuthnStatement>`)

	b.WriteString(`<saml2:AttributeStatement>`)
	attr := func(friendly, name string, values ...string) {
		fmt.Fprintf(&b, `<saml2:Attribute FriendlyName="%s" Name="%s" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:uri">`,
			friendly, name)
		for _, v := range values {
			fmt.Fprintf(&b, `<saml2:AttributeValue xsi:type="xs:string">%s</saml2:AttributeValue>`, escape(v, false))
		}
		b.WriteString(`</saml2:Attribute>`)
	}
	attr("Groups", groupsAttr, t.Groups...)
	attr("Subject Type", solutionAttr, fmt.Sprint(t.Solution))
	b.WriteString(`</saml2:AttributeStatement></saml2:Assertion>`)

	return b.String(), pos
}

// signedInfo returns the canonical ds:SignedInfo for an assertion with the given id and digest.
// When ns is true the ds namespace is declared, as it is when SignedInfo is canonicalized to be signed.
func signedInfo(id string, digest []byte, ns bool) string {
	decl := ""
	if ns {
		decl = fmt.Sprintf(` xmlns:ds="%s"`, internal.DSIG)
	}
	return fmt.Sprintf(`<ds:SignedInfo%s>`, decl) +
		fmt.Sprintf(`<ds:CanonicalizationMethod Algorithm="%s"></ds:CanonicalizationMethod>`, excC14) +
		fmt.Sprintf(`<ds:SignatureMethod Algorithm="%s"></ds:SignatureMethod>`, internal.SHA256) +
		fmt.Sprintf(`<ds:Reference URI="#%s"><ds:Transforms>`, id) +
		`<ds:Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"></ds:Transform>` +
		fmt.Sprintf(`<ds:Transform Algorithm="%s"><ec:InclusiveNamespaces xmlns:ec="%s" PrefixList="xs xsi"></ec:InclusiveNamespaces></ds:Transform>`, excC14, excC14) +
		`</ds:Transforms><ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"></ds:DigestMethod>` +
		fmt.Sprintf(`<ds:DigestValue>%s</ds:DigestValue></ds:Reference></ds:SignedInfo>`, base64.StdEncoding.EncodeToString(digest))
}

// sign returns t as a SAML assertion with an enveloped signature by key, whose certificate is included.
func (t *token) sign(issuer string, key *rsa.PrivateKey, cert *x509.Certificate) (string, error) {
	id := "_" + uuid.NewString()
	unsigned, pos := t.assertion(id, issuer, time.Now())

	digest := sha256.Sum256([]byte(unsigned))
	sum := sha256.Sum256([]byte(signedInfo(id, digest[:], true)))
	sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, sum[:])
	if err != nil {
		return "", err
	}

	signature := fmt.Sprintf(`<ds:Signature xmlns:ds="%s">`, internal.DSIG) +
		signedInfo(id, digest[:], false) +
		fmt.Sprintf(`<ds:SignatureValue>%s</ds:SignatureValue>`, base64.StdEncoding.EncodeToString(sig)) +
		fmt.Sprintf(`<ds:KeyInfo><ds:X509Data><ds:X509Certificate>%s</ds:X509Certificate></ds:X509Data></ds:KeyInfo>`,
			base64.StdEncoding.EncodeToString(cert.Raw)) +
		`</ds:Signature>`

	return unsigned[:pos] + signature + unsigned[pos:], nil
}

// between returns the text of s between the first open and the next close that follows it.
func between(s, open, close string) (string, bool) {
	_, after, ok := strings.Cut(s, open)
	if !ok {
		return "", false
	}
	inner, _, ok := strings.Cut(after, close)
	return inner, ok
}

var errInvalidToken = errors.New("sts: invalid token")

// verify validates that assertion was signed by key, and is valid at now, returning its content.
// Only the form written by token.sign is accepted: this is not a general XML signature validator,
// and any change to the bytes of a signed assertion invalidates it.
func verify(assertion string, key *rsa.PublicKey, now time.Time) (*token, error) {
	const open, close = `<ds:Signature xmlns:ds="` + internal.DSIG + `">`, `</ds:Signature>`

	start := strings.Index(assertion, open)
	if start < 0 {
		return nil, errInvalidToken
	}
	end := strings.Index(assertion[start:], close)
	if end < 0 {
		return nil, errInvalidToken
	}
	end += start + len(close)
	signature := assertion[start:end]
	unsigned := assertion[:start] + assertion[end:]
	if strings.Contains(unsigned, "<ds:Signature") {
		return nil, errInvalidToken // only one signature
	}

	var parsed struct {
		XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:assertion Assertion"`
		ID      string   `xml:"ID,attr"`
		Subject struct {
			NameID       string `xml:"NameID"`
			Confirmation struct {
				Method      string `xml:"Method,attr"`
				Certificate string `xml:"SubjectConfirmationData>KeyInfo>X509Data>X509Certificate"`
			} `xml:"SubjectConfirmation"`
		} `xml:"Subject"`
		Conditions struct {
			NotBefore    string `xml:"NotBefore,attr"`
			NotOnOrAfter string `xml:"NotOnOrAfter,attr"`
		} `xml:"Conditions"`
		Attributes []struct {
			Name   string   `xml:"Name,attr"`
			Values []string `xml:"AttributeValue"`
		} `xml:"AttributeStatement>Attribute"`
	}
	// Decode only what the digest covers, so no content can be added outside it, in the signature.
	if err := xml.Unmarshal([]byte(unsigned), &parsed); err != nil || parsed.ID == "" {
		return nil, errInvalidToken
	}

	digest, ok := between(signature, "<ds:DigestValue>", "</ds:DigestValue>")
	if !ok {
		return nil, errInvalidToken
	}
	value, ok := between(signature, "<ds:SignatureValue>", "</ds:SignatureValue>")
	if !ok {
		return nil, errInvalidToken
	}
	want, err := base64.StdEncoding.DecodeString(digest)
	if err != nil {
		return nil, errInvalidToken
	}
	sig, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, errInvalidToken
	}

	// The signed SignedInfo is rebuilt rather than taken from the signature,
	// so the only algorithms and reference accepted are those token.sign writes.
	info := signedInfo(parsed.ID, want, false)
	if !strings.HasPrefix(signature[len(open):], info) {
		return nil, errInvalidToken
	}
	got := sha256.Sum256([]byte(unsigned))
	if subtle.ConstantTimeCompare(got[:], want) != 1 {
		return nil, errInvalidToken
	}
	sum := sha256.Sum256([]byte(signedInfo(parsed.ID, want, true)))
	if rsa.VerifyPKCS1v15(key, crypto.SHA256, sum[:], sig) != nil {
		return nil, errInvalidToken
	}

	t := &token{Subject: parsed.Subject.NameID}
	if t.Subject == "" {
		return nil, errInvalidToken
	}
	if t.NotBefore, err = time.Parse(internal.Time, parsed.Conditions.NotBefore); err != nil {
		return nil, errInvalidToken
	}
	if t.NotOnOrAfter, err = time.Parse(internal.Time, parsed.Conditions.NotOnOrAfter); err != nil {
		return nil, errInvalidToken
	}
	if now.Add(clockSkew).Before(t.NotBefore) || !now.Add(-clockSkew).Before(t.NotOnOrAfter) {
		return nil, errors.New("sts: token is expired or not yet valid")
	}

	switch parsed.Subject.Confirmation.Method {
	case confirmBearer:
	case confirmHoK:
		der, err := base64.StdEncoding.DecodeString(parsed.Subject.Confirmation.Certificate)
		if err != nil {
			return nil, errInvalidToken
		}
		if t.Certificate, err = x509.ParseCertificate(der); err != nil {
			return nil, errInvalidToken
		}
	default:
		return nil, errInvalidToken
	}

	for _, a := range parsed.Attributes {
		switch a.Name {
		case groupsAttr:
			t.Groups = a.Values
		case solutionAttr:
			t.Solution = len(a.Values) == 1 && a.Values[0] == "true"
		}
	}

	return t, nil
}
