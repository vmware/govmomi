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
	"strings"

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

func (e *Executor) Run(args []string) (*Response, error) {
	req := Request{}
	req.ParseArgs(args)

	ns := req.Args[:len(req.Args)-1]

	sreq := types.ExecuteSoap{
		This:     e.mme.ManagedObjectReference,
		Moid:     "ha-cli-handler-" + strings.Join(ns, "-"),
		Method:   "vim.EsxCLI." + strings.Join(req.Args, "."),
		Version:  "urn:vim25/5.0",
		Argument: []types.ReflectManagedMethodExecuterSoapArgument{},
	}

	for key, val := range req.Flags {
		sreq.Argument = append(sreq.Argument, types.ReflectManagedMethodExecuterSoapArgument{
			Name: key,
			Val:  fmt.Sprintf("<%s>%s</%s>", key, val, key),
		})
	}

	sres, err := methods.ExecuteSoap(e.c, &sreq)
	if err != nil {
		return nil, err
	}

	res := &Response{}

	if sres.Returnval != nil {
		if err := xml.Unmarshal([]byte(sres.Returnval.Response), res); err != nil {
			return nil, err
		}
	}

	return res, nil
}
