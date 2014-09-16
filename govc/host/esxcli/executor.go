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

package esxcli

import (
	"fmt"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type Executor struct {
	c    *govmomi.Client
	host *govmomi.HostSystem
	mme  *types.ReflectManagedMethodExecuter
	dtm  *types.InternalDynamicTypeManager
}

func NewExecutor(c *govmomi.Client, host *govmomi.HostSystem) (*Executor, error) {
	e := &Executor{
		c:    c,
		host: host,
	}

	{
		req := types.RetrieveManagedMethodExecuter{
			This: host.Reference(),
		}

		res, err := methods.RetrieveManagedMethodExecuter(c, &req)
		if err != nil {
			return nil, err
		}

		e.mme = res.Returnval
	}

	{
		req := types.RetrieveDynamicTypeManager{
			This: host.Reference(),
		}

		res, err := methods.RetrieveDynamicTypeManager(c, &req)
		if err != nil {
			return nil, err
		}

		e.dtm = res.Returnval
	}

	return e, nil
}

func (e *Executor) ParamTypeInfo(c *Command) ([]types.DynamicTypeMgrParamTypeInfo, error) {
	req := types.DynamicTypeMgrQueryTypeInfo{
		This: e.dtm.ManagedObjectReference,
		FilterSpec: &types.DynamicTypeMgrTypeFilterSpec{
			TypeSubstr: c.Namespace(),
		},
	}

	res, err := methods.DynamicTypeMgrQueryTypeInfo(e.c, &req)
	if err != nil {
		return nil, err
	}

	name := c.Name()
	for _, info := range res.Returnval.ManagedTypeInfo {
		for _, method := range info.Method {
			if method.Name == name {
				return method.ParamTypeInfo, nil
			}
		}
	}

	return nil, fmt.Errorf("method '%s' not found in name space '%s'", name, c.Namespace())
}

func (e *Executor) NewRequest(args []string) (*types.ExecuteSoap, error) {
	c := NewCommand(args)

	params, err := e.ParamTypeInfo(c)
	if err != nil {
		return nil, err
	}

	sargs, err := c.Parse(params)
	if err != nil {
		return nil, err
	}

	sreq := types.ExecuteSoap{
		This:     e.mme.ManagedObjectReference,
		Moid:     c.Moid(),
		Method:   c.Method(),
		Version:  "urn:vim25/5.0",
		Argument: sargs,
	}

	return &sreq, nil
}

func (e *Executor) Run(args []string) (*Response, error) {
	req, err := e.NewRequest(args)
	if err != nil {
		return nil, err
	}

	x, err := methods.ExecuteSoap(e.c, req)
	if err != nil {
		return nil, err
	}

	res := &Response{}

	if x.Returnval != nil {
		if err := xml.Unmarshal([]byte(x.Returnval.Response), res); err != nil {
			return nil, err
		}
	}

	return res, nil
}
