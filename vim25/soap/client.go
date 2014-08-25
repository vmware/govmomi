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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/debug"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type HasFault interface {
	Fault() *Fault
}

type RoundTripper interface {
	RoundTrip(req, res HasFault) error
}

var cn uint64 // Client counter

type Client struct {
	http.Client

	u url.URL

	cn  uint64         // Client counter
	rn  uint64         // Request counter
	log io.WriteCloser // Request log
}

func NewClient(u url.URL) *Client {

	c := Client{
		u: u,

		cn: atomic.AddUint64(&cn, 1),
		rn: 0,
	}

	if c.u.Scheme == "https" {
		c.Client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	c.Jar, _ = cookiejar.New(nil)
	c.u.User = nil

	if debug.Enabled() {
		c.log = debug.NewFile(fmt.Sprintf("%d-client.log", c.cn))
	}

	return &c
}

func (c *Client) URL() url.URL {
	return c.u
}

type marshaledClient struct {
	Cookies []*http.Cookie
	URL     *url.URL
}

func (c *Client) MarshalJSON() ([]byte, error) {
	m := marshaledClient{
		Cookies: c.Jar.Cookies(&c.u),
		URL:     &c.u,
	}

	return json.Marshal(m)
}

func (c *Client) UnmarshalJSON(b []byte) error {
	var m marshaledClient

	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	*c = *NewClient(*m.URL)
	c.Jar.SetCookies(m.URL, m.Cookies)

	return nil
}

func (c *Client) RoundTrip(reqBody, resBody HasFault) error {
	var httpreq *http.Request
	var httpres *http.Response
	var err error

	reqEnv := Envelope{Body: reqBody}
	resEnv := Envelope{Body: resBody}

	num := atomic.AddUint64(&c.rn, 1)

	b, err := xml.Marshal(reqEnv)
	if err != nil {
		panic(err)
	}

	rawreqbody := io.MultiReader(strings.NewReader(xml.Header), bytes.NewReader(b))
	if debug.Enabled() {
		f := debug.NewFile(fmt.Sprintf("%d-%04d.req.xml", c.cn, num))
		rawreqbody = io.TeeReader(rawreqbody, f)
	}

	httpreq, err = http.NewRequest("POST", c.u.String(), rawreqbody)
	if err != nil {
		panic(err)
	}

	httpreq.Header.Set(`Content-Type`, `text/xml; charset="utf-8"`)
	httpreq.Header.Set(`SOAPAction`, `urn:vim25/5.5`)

	if debug.Enabled() {
		b, _ := httputil.DumpRequest(httpreq, false)
		wc := debug.NewFile(fmt.Sprintf("%d-%04d.req.headers", c.cn, num))
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

	var rawresbody io.Reader = httpres.Body
	defer httpres.Body.Close()

	if debug.Enabled() {
		f := debug.NewFile(fmt.Sprintf("%d-%04d.res.xml", c.cn, num))
		rawresbody = io.TeeReader(rawresbody, f)
	}

	if debug.Enabled() {
		b, _ := httputil.DumpResponse(httpres, false)
		wc := debug.NewFile(fmt.Sprintf("%d-%04d.res.headers", c.cn, num))
		wc.Write(b)
		wc.Close()
	}

	dec := xml.NewDecoder(rawresbody)
	dec.TypeFunc = types.TypeFunc()
	err = dec.Decode(&resEnv)
	if err != nil {
		return err
	}

	if f := resBody.Fault(); f != nil {
		return WrapSoapFault(f)
	}

	return err
}

// ParseURL wraps url.Parse to rewrite the URL.Host field
// In the case of VM guest uploads or NFC lease URLs, a Host
// field with a value of "*" is rewritten to the Client's Host.
func (c *Client) ParseURL(urlStr string) (*url.URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	host := strings.Split(u.Host, ":")
	if host[0] == "*" {
		host[0] = strings.Split(c.URL().Host, ":")[0]
		u.Host = strings.Join(host, ":")
	}

	return u, nil
}

type Upload struct {
	Type       string
	Method     string
	ProgressCh chan<- vim25.Progress
}

var DefaultUpload = Upload{
	Type:   "application/octet-stream",
	Method: "PUT",
}

// UploadFile PUTs the local file to the given URL
func (c *Client) UploadFile(file string, u *url.URL, param *Upload) error {
	var err error

	if param == nil {
		param = &DefaultUpload
	}

	pr := progressReader{
		ch: param.ProgressCh,
	}

	// Mark progress reader as done when returning from this function.
	defer func() {
		pr.Done(err)
	}()

	s, err := os.Stat(file)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	pr.r = f
	pr.size = s.Size()
	req, err := http.NewRequest(param.Method, u.String(), &pr)
	if err != nil {
		return err
	}

	req.ContentLength = s.Size()
	req.Header.Set("Content-Type", param.Type)

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	default:
		err = errors.New(res.Status)
	}

	return err
}

type Download struct {
	Method     string
	ProgressCh chan<- vim25.Progress
}

var DefaultDownload = Download{
	Method: "GET",
}

// DownloadFile GETs the given URL to a local file
func (c *Client) DownloadFile(file string, u *url.URL, param *Download) error {
	var err error

	if param == nil {
		param = &DefaultDownload
	}

	pr := progressReader{
		ch: param.ProgressCh,
	}

	// Mark progress reader as done when returning from this function.
	defer func() {
		pr.Done(err)
	}()

	fh, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fh.Close()

	req, err := http.NewRequest(param.Method, u.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
	default:
		err = errors.New(res.Status)
	}

	if err != nil {
		return err
	}

	pr.r = res.Body
	pr.size = res.ContentLength
	_, err = io.Copy(fh, &pr)
	if err != nil {
		return err
	}

	// Assign error before returning so that it gets picked up by the deferred
	// function marking the progress reader as done.
	err = fh.Close()
	if err != nil {
		return err
	}

	return nil
}
