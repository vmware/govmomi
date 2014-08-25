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
	Client         *soap.Client
	ServiceContent types.ServiceContent
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
	}

	sc, err := serviceContent(c.Client)
	if err != nil {
		return nil, err
	}

	err = login(c.Client, u, sc)
	if err != nil {
		return nil, err
	}

	c.ServiceContent = sc

	return &c, nil
}

// RoundTrip dispatches to the client's SOAP client RoundTrip function.
func (c *Client) RoundTrip(req, res soap.HasFault) error {
	return c.Client.RoundTrip(req, res)
}

func (c *Client) UserSession() (*types.UserSession, error) {
	var sm mo.SessionManager

	err := mo.RetrieveProperties(c, c.ServiceContent.PropertyCollector, *c.ServiceContent.SessionManager, &sm)
	if err != nil {
		return nil, err
	}

	return sm.CurrentSession, nil
}

func (c *Client) Properties(obj types.ManagedObjectReference, p []string, dst interface{}) error {
	ospec := types.ObjectSpec{
		Obj:  obj,
		Skip: false,
	}

	pspec := types.PropertySpec{
		Type: obj.Type,
	}

	if p == nil {
		pspec.All = true
	} else {
		pspec.PathSet = p
	}

	req := types.RetrieveProperties{
		This: c.ServiceContent.PropertyCollector,
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{ospec},
				PropSet:   []types.PropertySpec{pspec},
			},
		},
	}

	return mo.RetrievePropertiesForRequest(c, req, dst)
}

func (c *Client) WaitForProperties(obj types.ManagedObjectReference, ps []string, f func([]types.PropertyChange) bool) error {
	p, err := c.NewPropertyCollector()
	if err != nil {
		return err
	}

	defer p.Destroy()

	req := types.CreateFilter{
		Spec: types.PropertyFilterSpec{
			ObjectSet: []types.ObjectSpec{
				{
					Obj: obj,
				},
			},
			PropSet: []types.PropertySpec{
				{
					PathSet: ps,
					Type:    obj.Type,
				},
			},
		},
	}

	err = p.CreateFilter(req)
	if err != nil {
		return err
	}

	for version := ""; ; {
		res, err := p.WaitForUpdates(version)
		if err != nil {
			return err
		}

		version = res.Version

		for _, fs := range res.FilterSet {
			for _, os := range fs.ObjectSet {
				if os.Obj == obj {
					if f(os.ChangeSet) {
						return nil
					}
				}
			}
		}
	}
}

// Ancestors returns the entire ancestry tree of a specified managed object.
// The return value includes the root node and the specified object itself.
func (c *Client) Ancestors(r Reference) ([]mo.ManagedEntity, error) {
	ospec := types.ObjectSpec{
		Obj: r.Reference(),
		SelectSet: []types.BaseSelectionSpec{
			&types.TraversalSpec{
				SelectionSpec: types.SelectionSpec{Name: "traverseParent"},
				Type:          "ManagedEntity",
				Path:          "parent",
				Skip:          false,
				SelectSet: []types.BaseSelectionSpec{
					&types.SelectionSpec{Name: "traverseParent"},
				},
			},
		},
		Skip: false,
	}

	pspec := types.PropertySpec{
		Type:    "ManagedEntity",
		PathSet: []string{"name", "parent"},
	}

	req := types.RetrieveProperties{
		This: c.ServiceContent.PropertyCollector,
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{ospec},
				PropSet:   []types.PropertySpec{pspec},
			},
		},
	}

	var ifaces []interface{}

	err := mo.RetrievePropertiesForRequest(c, req, &ifaces)
	if err != nil {
		return nil, err
	}

	var out []mo.ManagedEntity

	// Build ancestry tree by iteratively finding a new child.
	for len(out) < len(ifaces) {
		var find types.ManagedObjectReference

		if len(out) > 0 {
			find = out[len(out)-1].Self
		}

		// Find entity we're looking for given the last entity in the current tree.
		for _, iface := range ifaces {
			me := iface.(mo.IsManagedEntity).GetManagedEntity()
			if me.Parent == nil {
				out = append(out, me)
				break
			}

			if *me.Parent == find {
				out = append(out, me)
				break
			}
		}
	}

	return out, nil
}

func (c *Client) FileManager() FileManager {
	return FileManager{c}
}

func (c *Client) GuestOperationsManager() GuestOperationsManager {
	return GuestOperationsManager{c}
}

func (c *Client) OvfManager() OvfManager {
	return OvfManager{c}
}

func (c *Client) RootFolder() Folder {
	return Folder{c.ServiceContent.RootFolder}
}

func (c *Client) SearchIndex() SearchIndex {
	return SearchIndex{c}
}

func (c *Client) VirtualDiskManager() VirtualDiskManager {
	return VirtualDiskManager{c}
}

// NewPropertyCollector creates a new property collector based on the
// root property collector. It is the responsibility of the caller to
// clean up the property collector when done.
func (c *Client) NewPropertyCollector() (*PropertyCollector, error) {
	req := types.CreatePropertyCollector{
		This: c.ServiceContent.PropertyCollector,
	}

	res, err := methods.CreatePropertyCollector(c, &req)
	if err != nil {
		return nil, err
	}

	p := PropertyCollector{
		c: c,
		r: res.Returnval,
	}

	return &p, nil
}
