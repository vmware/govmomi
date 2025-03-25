// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type Fault struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (f Fault) Error() string {
	return f.Message
}

func (f Fault) MessageDetail() string {
	if f.Detail != "" {
		return fmt.Sprintf("%s %s", f.Message, f.Detail)
	}

	return f.Message
}

type Executor struct {
	c    *vim25.Client
	host mo.Reference
	mme  *internal.ReflectManagedMethodExecuter
	dtm  *internal.InternalDynamicTypeManager
	info map[string]*CommandInfo

	Trace func(*internal.ExecuteSoapRequest, *internal.ExecuteSoapResponse)
}

func NewExecutor(ctx context.Context, c *vim25.Client, host mo.Reference) (*Executor, error) {
	e := &Executor{
		c:    c,
		host: host,
		info: make(map[string]*CommandInfo),
	}

	{
		req := internal.RetrieveManagedMethodExecuterRequest{
			This: host.Reference(),
		}

		res, err := internal.RetrieveManagedMethodExecuter(ctx, c, &req)
		if err != nil {
			return nil, err
		}

		e.mme = res.Returnval
	}

	{
		req := internal.RetrieveDynamicTypeManagerRequest{
			This: host.Reference(),
		}

		res, err := internal.RetrieveDynamicTypeManager(ctx, c, &req)
		if err != nil {
			return nil, err
		}

		e.dtm = res.Returnval
	}

	return e, nil
}

func (e *Executor) Client() *vim25.Client {
	return e.c
}

func (e *Executor) DynamicTypeManager() types.ManagedObjectReference {
	return e.dtm.ManagedObjectReference
}

func (e *Executor) CommandInfo(ctx context.Context, ns string) (*CommandInfo, error) {
	info, ok := e.info[ns]
	if ok {
		return info, nil
	}

	req := internal.ExecuteSoapRequest{
		Moid:   "ha-dynamic-type-manager-local-cli-cliinfo",
		Method: "vim.CLIInfo.FetchCLIInfo",
		Argument: []internal.ReflectManagedMethodExecuterSoapArgument{
			NewCommand(nil).Argument("typeName", "vim.EsxCLI."+ns),
		},
	}

	info = new(CommandInfo)
	if err := e.Execute(ctx, &req, info); err != nil {
		return nil, err
	}

	e.info[ns] = info

	return info, nil
}

func (e *Executor) CommandInfoMethod(ctx context.Context, c *Command) (*CommandInfoMethod, error) {
	ns := c.Namespace()

	info, err := e.CommandInfo(ctx, ns)
	if err != nil {
		return nil, err
	}

	name := c.Name()
	for _, method := range info.Method {
		if method.Name == name {
			return &method, nil
		}
	}

	return nil, fmt.Errorf("method '%s' not found in name space '%s'", name, c.Namespace())
}

func (e *Executor) NewRequest(ctx context.Context, args []string) (*internal.ExecuteSoapRequest, *CommandInfoMethod, error) {
	c := NewCommand(args)

	info, err := e.CommandInfoMethod(ctx, c)
	if err != nil {
		return nil, nil, err
	}

	sargs, err := c.Parse(info.Param)
	if err != nil {
		return nil, nil, err
	}

	sreq := internal.ExecuteSoapRequest{
		Moid:     c.Moid(),
		Method:   c.Method(),
		Argument: sargs,
	}

	return &sreq, info, nil
}

func (e *Executor) Execute(ctx context.Context, req *internal.ExecuteSoapRequest, res any) error {
	req.This = e.mme.ManagedObjectReference
	req.Version = "urn:vim25/5.0"

	x, err := internal.ExecuteSoap(ctx, e.c, req)
	if err != nil {
		return err
	}

	if e.Trace != nil {
		e.Trace(req, x)
	}

	if x.Returnval != nil {
		if x.Returnval.Fault != nil {
			return &Fault{
				x.Returnval.Fault.FaultMsg,
				x.Returnval.Fault.FaultDetail,
			}
		}

		if err := xml.Unmarshal([]byte(x.Returnval.Response), res); err != nil {
			return err
		}
	}

	return nil
}

func (e *Executor) Run(ctx context.Context, args []string) (*Response, error) {
	req, info, err := e.NewRequest(ctx, args)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Info: info,
	}

	if err := e.Execute(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
