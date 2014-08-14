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

package govmomi

import (
	"net/url"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var serviceInstance = types.ManagedObjectReference{
	Type:  "ServiceInstance",
	Value: "ServiceInstance",
}

type Client struct {
	*soap.Client

	types.ServiceContent

	u url.URL
}

func serviceContent(r soap.RoundTripper) (types.ServiceContent, error) {
	req := types.RetrieveServiceContent{
		This: serviceInstance,
	}

	res, err := methods.RetrieveServiceContent(r, &req)
	if err != nil {
		return types.ServiceContent{}, err
	}

	return res.Returnval, nil
}

func login(r soap.RoundTripper, u url.URL, sc types.ServiceContent) error {
	// Return if URL doesn't contain username/password information.
	if u.User == nil {
		return nil
	}

	req := types.Login{
		This: *sc.SessionManager,
	}

	req.UserName = u.User.Username()
	if pw, ok := u.User.Password(); ok {
		req.Password = pw
	}

	_, err := methods.Login(r, &req)
	if err != nil {
		return err
	}

	return nil
}

func NewClient(u url.URL) (*Client, error) {
	c := Client{
		Client: soap.NewClient(u),

		u: u,
	}

	sc, err := serviceContent(c.Client)
	if err != nil {
		return nil, err
	}

	err = login(c.Client, c.u, sc)
	if err != nil {
		return nil, err
	}

	c.ServiceContent = sc

	return &c, nil
}

func (c *Client) UserSession() (*types.UserSession, error) {
	var sm mo.SessionManager

	err := mo.RetrieveProperties(c, c.PropertyCollector, *c.SessionManager, &sm)
	if err != nil {
		return nil, err
	}

	return sm.CurrentSession, nil
}
