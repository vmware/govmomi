// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	Path = "/rest"
)

// Resource wraps url.URL with helpers
type Resource struct {
	u *url.URL
}

func (r *Resource) String() string {
	return r.u.String()
}

// WithSubpath appends the provided subpath to the URL.Path
func (r *Resource) WithSubpath(subpath string) *Resource {
	r.u.Path += "/" + subpath
	return r
}

// WithID appends id to the URL.Path
func (r *Resource) WithID(id string) *Resource {
	r.u.Path += "/id:" + id
	return r
}

// WithAction sets adds action to the URL.RawQuery
func (r *Resource) WithAction(action string) *Resource {
	return r.WithParam("~action", action)
}

// WithParam adds one parameter on the URL.RawQuery
func (r *Resource) WithParam(name string, value string) *Resource {
	// ParseQuery handles empty case, and we control access to query string so shouldn't encounter an error case
	params, err := url.ParseQuery(r.u.RawQuery)
	if err != nil {
		panic(err)
	}
	params[name] = append(params[name], value)
	r.u.RawQuery = params.Encode()
	return r
}

// WithPathEncodedParam appends a parameter on the URL.RawQuery,
// For special cases where URL Path-style encoding is needed
func (r *Resource) WithPathEncodedParam(name string, value string) *Resource {
	t := &url.URL{Path: value}
	encodedValue := t.String()
	t = &url.URL{Path: name}
	encodedName := t.String()
	// ParseQuery handles empty case, and we control access to query string so shouldn't encounter an error case
	params, err := url.ParseQuery(r.u.RawQuery)
	if err != nil {
		panic(err)
	}
	// Values.Encode() doesn't escape exactly how we want, so we need to build the query string ourselves
	if len(params) >= 1 {
		r.u.RawQuery = r.u.RawQuery + "&" + encodedName + "=" + encodedValue
	} else {
		r.u.RawQuery = r.u.RawQuery + encodedName + "=" + encodedValue
	}
	return r
}

// Request returns a new http.Request for the given method.
// An optional body can be provided for POST and PATCH methods.
func (r *Resource) Request(method string, body ...any) *http.Request {
	rdr := io.MultiReader() // empty body by default
	if len(body) != 0 {
		rdr = encode(body[0])
	}
	req, err := http.NewRequest(method, r.u.String(), rdr)
	if err != nil {
		panic(err)
	}
	return req
}

type errorReader struct {
	e error
}

func (e errorReader) Read([]byte) (int, error) {
	return -1, e.e
}

// encode body as JSON, deferring any errors until io.Reader is used.
func encode(body any) io.Reader {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(body)
	if err != nil {
		return errorReader{err}
	}
	return &b
}
