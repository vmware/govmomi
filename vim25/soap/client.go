/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package soap

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/vmware/govmomi/vim25/debug"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type RoundTripper interface {
	RoundTrip(req, res *Envelope) error
}

type Client struct {
	http.Client

	u url.URL
	t map[string]reflect.Type

	mu  sync.Mutex
	c   int            // Request counter
	log io.WriteCloser // Request log
}

func NewClient(u url.URL) *Client {
	c := Client{
		u: u,
	}

	if c.u.Scheme == "https" {
		c.Client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	c.Jar, _ = cookiejar.New(nil)
	c.u.User = nil
	c.t = types.TypeMap()

	if debug.Enabled() {
		c.log = debug.NewFile("client.log")
	}

	return &c
}

func (c *Client) RoundTrip(req, res *Envelope) error {
	var httpreq *http.Request
	var httpres *http.Response
	var err error

	c.mu.Lock()
	num := c.c
	c.c++
	c.mu.Unlock()

	b, err := xml.Marshal(req)
	if err != nil {
		panic(err)
	}

	xmlbody := io.MultiReader(strings.NewReader(xml.Header), bytes.NewReader(b))
	httpreq, err = http.NewRequest("POST", c.u.String(), xmlbody)
	if err != nil {
		panic(err)
	}

	httpreq.Header.Set(`Content-Type`, `text/xml; charset="utf-8"`)
	httpreq.Header.Set(`SOAPAction`, `urn:vim25/5.5`)

	if debug.Enabled() {
		b, _ := httputil.DumpRequest(httpreq, true)
		wc := debug.NewFile(fmt.Sprintf("%04d.req", num))
		wc.Write(b)
		wc.Close()
	}

	tstart := time.Now()
	httpres, err = c.Client.Do(httpreq)
	tstop := time.Now()

	if debug.Enabled() {
		now := time.Now().Format("2006-01-02T15-04-05.999999999")
		ms := tstop.Sub(tstart) / time.Millisecond
		fmt.Fprintf(c.log, "%s: %4d took %6dms\n", now, num, ms)
	}

	if err != nil {
		return err
	}

	if debug.Enabled() {
		b, _ := httputil.DumpResponse(httpres, true)
		wc := debug.NewFile(fmt.Sprintf("%04d.res", num))
		wc.Write(b)
		wc.Close()
	}

	dec := xml.NewDecoder(httpres.Body)
	dec.Types = c.t
	err = dec.Decode(res)
	if err != nil {
		return err
	}

	return err
}
