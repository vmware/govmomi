// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/vmware/govmomi/vim25/xml"
)

// maxToken limits the size of a decompressed SIGN token.
const maxToken = 1 << 20

// signParams parses the parameters of a "SIGN" Authorization header, as written by sts.Signer.SignRequest:
// key="value" pairs, separated by commas.
func signParams(auth string) map[string]string {
	params := make(map[string]string)
	for _, p := range strings.Split(auth, ",") {
		key, val, ok := strings.Cut(strings.TrimSpace(p), "=")
		if ok {
			params[key] = strings.Trim(val, `"`)
		}
	}
	return params
}

// signToken decodes the SAML token of a "SIGN" Authorization header: gzip compressed, then base64 encoded.
func signToken(params map[string]string) (string, error) {
	gz, err := base64.StdEncoding.DecodeString(params["token"])
	if err != nil {
		return "", err
	}
	z, err := gzip.NewReader(bytes.NewReader(gz))
	if err != nil {
		return "", err
	}
	token, err := io.ReadAll(io.LimitReader(z, maxToken))
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// nameID returns the Subject NameID of a SAML token, without validating it.
func nameID(token string) string {
	var subject struct {
		ID string `xml:"Subject>NameID"`
	}
	_ = xml.Unmarshal([]byte(token), &subject)
	return subject.ID
}

var errSignature = errors.New("invalid request signature")

// verifyRequest checks the signature of a request signed with a holder-of-key token, per sts.Signer.SignRequest,
// against the certificate that confirms the token.
func verifyRequest(r *http.Request, params map[string]string, cert *x509.Certificate) error {
	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok || params["signature_alg"] != "RSA-SHA256" {
		return errSignature
	}
	sig, err := base64.StdEncoding.DecodeString(params["signature"])
	if err != nil {
		return errSignature
	}
	nonce := params["nonce"]
	if nonce == "" {
		return errSignature
	}

	var body []byte
	if r.Body != nil {
		if body, err = io.ReadAll(r.Body); err != nil {
			return err
		}
		r.Body = io.NopCloser(bytes.NewReader(body))
	}
	bhash := sha256.Sum256(body)
	if want, err := base64.StdEncoding.DecodeString(params["bodyhash"]); err != nil || subtle.ConstantTimeCompare(want, bhash[:]) != 1 {
		return errSignature
	}

	// The client signs the host and port of its request URL, with port 80 when the URL has none.
	host, port, err := net.SplitHostPort(r.Host)
	if err != nil {
		host, port = strings.Trim(r.Host, "[]"), "80"
	}
	if strings.Contains(host, ":") {
		host = "[" + host + "]" // IPv6
	}

	var msg bytes.Buffer
	for _, s := range []string{nonce, r.Method, r.URL.Path, strings.ToLower(host), port} {
		msg.WriteString(s)
		msg.WriteByte('\n')
	}
	msg.Write(bhash[:])
	msg.WriteByte('\n')

	sum := sha256.Sum256(msg.Bytes())
	if rsa.VerifyPKCS1v15(key, crypto.SHA256, sum[:], sig) != nil {
		return errSignature
	}
	return nil
}

// tokenAuthorization authenticates a request by its "SIGN" Authorization header, returning the token's subject.
// When the simulator's SessionManager.ValidToken is set, the token must be valid, and a request with a
// holder-of-key token must be signed with the token's key. Otherwise any token is accepted.
func (s *handler) tokenAuthorization(r *http.Request, auth string) (string, bool) {
	params := signParams(auth)
	token, err := signToken(params)
	if err != nil {
		return "", false
	}

	valid := s.Map.SessionManager().ValidToken
	if valid == nil {
		return nameID(token), true
	}

	id, err := valid(token)
	if err != nil {
		return "", false
	}
	if id.Certificate != nil {
		if err := verifyRequest(r, params, id.Certificate); err != nil {
			return "", false
		}
	}
	return id.Subject, true
}
