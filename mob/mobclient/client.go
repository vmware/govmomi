/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package mobclient

import (
	"bytes"
	"context"
	"fmt"
	"github.com/vmware/govmomi/mob/internal"
	"github.com/vmware/govmomi/vim25/soap"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

type Client struct {
	mu sync.Mutex
	*soap.Client
	Cookie *http.Cookie
}

// Session information
type Session struct {
	User         string    `json:"user"`
	Created      time.Time `json:"created_time"`
	LastAccessed time.Time `json:"last_accessed_time"`
}

func NewClient(c *soap.Client) *Client {
	return &Client{Client: c}
}

type statusError struct {
	res *http.Response
}

func (e *statusError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.res.Request.Method, e.res.Request.URL, e.res.Status)
}

// Resource helper for the given path.
func (c *Client) Resource(path string) *Resource {
	r := &Resource{u: c.URL()}
	r.u.Path = path
	return r
}

// Do sends the http.Request, decoding resBody if provided.
func (c *Client) Do(ctx context.Context, req *http.Request, resBody *string) (string, error) {
	switch req.Method {
	case http.MethodPost, http.MethodPatch:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if c.Cookie != nil {
		req.AddCookie(c.Cookie)
	}
	err := c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			// TODO: structured error types
			detail, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &statusError{res}
		}

		for _, cookie := range res.Cookies() {
			if cookie.Name == internal.SessionCookieName {
				c.Cookie = cookie
			}
		}

		if resBody == nil {
			return nil
		}

		output, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		*resBody = string(output)
		return nil
	})
	return c.extractXsrfToken(*resBody), err

}

// Login creates a new session via Basic Authentication with the given url.Userinfo.
func (c *Client) Login(ctx context.Context, user *url.Userinfo) error {
	sessionCallRes := c.Resource(internal.SessionPath)
	req := sessionCallRes.Request(http.MethodGet)
	if user != nil {
		if password, ok := user.Password(); ok {
			req.SetBasicAuth(user.Username(), password)
		}
	}

	responseBody := ""
	_, err := c.Do(ctx, req, &responseBody)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) extractXsrfToken(out string) string {
	regex, _ := regexp.Compile("<input name=\"vmware-session-nonce\" type=\"hidden\" value=\"(.+?)\">")
	loc := regex.FindStringSubmatch(out)
	if len(loc) < 2 {
		return ""
	}
	return loc[1]
}
