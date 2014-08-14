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
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

var debug = false

func init() {
	debug = (os.Getenv("DEBUG") != "")
}

type RoundTripper interface {
	RoundTrip(req, res *Envelope) error
}

type Client struct {
	http.Client

	u url.URL
	c *http.Cookie
	t map[string]reflect.Type
}

func NewClient(u url.URL) *Client {
	c := Client{
		u: u,
	}

	if c.u.Scheme == "https" {
		c.Client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	c.u.User = nil
	c.t = types.TypeMap()

	return &c
}

func (c *Client) RoundTrip(req, res *Envelope) error {
	var httpreq *http.Request
	var httpres *http.Response
	var err error

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

	if c.c != nil {
		httpreq.AddCookie(c.c)
	}

	if debug {
		b, _ := httputil.DumpRequest(httpreq, true)
		fmt.Printf("----------- request\n")
		fmt.Printf("%s\n", string(b))
	}

	httpres, err = c.Client.Do(httpreq)
	if err != nil {
		return err
	}

	if debug {
		b, _ := httputil.DumpResponse(httpres, true)
		fmt.Printf("----------- response\n")
		fmt.Printf("%s\n", string(b))
	}

	dec := xml.NewDecoder(httpres.Body)
	dec.Types = c.t
	err = dec.Decode(res)
	if err != nil {
		panic(err)
	}

	cookies := httpres.Cookies()
	for _, cookie := range cookies {
		if cookie.Name != "vmware_soap_session" {
			continue
		}

		c.c = cookie
		c.c.Secure = false // Don't care
	}

	return err
}
