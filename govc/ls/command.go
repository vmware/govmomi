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

package vm

import (
	"flag"
	"fmt"
	"io"
	"path"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.OutputFlag
}

func init() {
	cli.Register(&ls{})
}

func (l *ls) Register(f *flag.FlagSet) {}

func (l *ls) Process() error { return nil }

func (l *ls) Run(f *flag.FlagSet) error {
	client, err := l.Client()
	if err != nil {
		return err
	}

	var root types.ManagedObjectReference
	var res listResult

	arg := path.Clean(f.Arg(0))
	if len(arg) > 0 && arg[0] == '/' {
		root = client.ServiceContent.RootFolder
		arg = arg[1:]
	} else {
		dc, err := l.Datacenter()
		if err != nil {
			return err
		}

		root = dc.Reference()
		if arg == "." {
			arg = ""
		}
	}

	parts := strings.Split(arg, "/")
	if parts[0] == "" {
		parts = parts[1:]
	}

	for {
		full := l.JSON && len(parts) == 0

		switch root.Type {
		case "Folder":
			res, err = l.listFolder(root, full)
		case "Datacenter":
			res, err = l.listDatacenter(root, full)
		default:
			return fmt.Errorf("cannot traverse type " + root.Type)
		}

		if err != nil {
			return err
		}

		if len(parts) == 0 {
			break
		}

		if _, ok := res.byName[parts[0]]; !ok {
			return fmt.Errorf("%s not found", arg)
		}

		root = res.byName[parts[0]]
		parts = parts[1:]
	}

	return l.WriteResult(&res)
}

func (l *ls) listFolder(m types.ManagedObjectReference, full bool) (listResult, error) {
	var res = newListResult()
	var me mo.Folder

	c, err := l.Client()
	if err != nil {
		return res, err
	}

	err = c.Properties(m, []string{"name", "childType", "childEntity"}, &me)
	if err != nil {
		return res, err
	}

	req := types.RetrieveProperties{
		This: c.ServiceContent.PropertyCollector,
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{
					{
						Obj: m,
						SelectSet: []types.BaseSelectionSpec{
							&types.TraversalSpec{
								Path: "childEntity",
								Skip: false,
								Type: "Folder",
							},
						},
						Skip: true,
					},
				},
			},
		},
	}

	for _, t := range me.ChildType {
		// Retrieve only the managed entity's name
		pspec := types.PropertySpec{
			Type: t,
		}

		if full {
			pspec.All = true
		} else {
			pspec.PathSet = []string{"name"}
		}

		req.SpecSet[0].PropSet = append(req.SpecSet[0].PropSet, pspec)
	}

	var dst []interface{}

	err = mo.RetrievePropertiesForRequest(c, req, &dst)
	if err != nil {
		return res, err
	}

	for _, v := range dst {
		switch m := v.(type) {
		case mo.Folder:
			res.byName[m.Name] = m.Reference()
			res.Folders = append(res.Folders, m)
		case mo.Datacenter:
			res.byName[m.Name] = m.Reference()
			res.Datacenters = append(res.Datacenters, m)
		case mo.VirtualMachine:
			res.byName[m.Name] = m.Reference()
			res.VirtualMachines = append(res.VirtualMachines, m)
		case mo.Network:
			res.byName[m.Name] = m.Reference()
			res.Networks = append(res.Networks, m)
		case mo.ComputeResource:
			res.byName[m.Name] = m.Reference()
			res.ComputeResources = append(res.ComputeResources, m)
		case mo.Datastore:
			res.byName[m.Name] = m.Reference()
			res.Datastores = append(res.Datastores, m)
		default:
			panic("not implemented for type " + reflect.TypeOf(v).String())
		}
	}

	return res, nil
}

func (l *ls) listDatacenter(m types.ManagedObjectReference, all bool) (listResult, error) {
	var res = newListResult()

	c, err := l.Client()
	if err != nil {
		return res, err
	}

	pspec := types.PropertySpec{
		Type: "Folder",
	}

	if all {
		pspec.All = true
	} else {
		pspec.PathSet = []string{"name"}
	}

	req := types.RetrieveProperties{
		This: c.ServiceContent.PropertyCollector,
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{
					{
						Obj:  m,
						Skip: true,
					},
				},
				PropSet: []types.PropertySpec{
					pspec,
				},
			},
		},
	}

	// Include every datastore folder in the select set
	os := &req.SpecSet[0].ObjectSet[0]
	for _, f := range []string{"vmFolder", "hostFolder", "datastoreFolder", "networkFolder"} {
		s := types.TraversalSpec{
			Path: f,
			Skip: false,
			Type: "Datacenter",
		}

		os.SelectSet = append(os.SelectSet, &s)
	}

	var dst []interface{}

	err = mo.RetrievePropertiesForRequest(c, req, &dst)
	if err != nil {
		return res, err
	}

	for _, v := range dst {
		switch m := v.(type) {
		case mo.Folder:
			res.byName[m.Name] = m.Reference()
			res.Folders = append(res.Folders, m)
		default:
			panic("not implemented for type " + reflect.TypeOf(v).String())
		}
	}

	return res, nil
}

type listResult struct {
	byName map[string]types.ManagedObjectReference

	Folders          []mo.Folder          `json:",omitempty"`
	Datacenters      []mo.Datacenter      `json:",omitempty"`
	VirtualMachines  []mo.VirtualMachine  `json:",omitempty"`
	Networks         []mo.Network         `json:",omitempty"`
	ComputeResources []mo.ComputeResource `json:",omitempty"`
	Datastores       []mo.Datastore       `json:",omitempty"`
}

func newListResult() listResult {
	l := listResult{
		byName: make(map[string]types.ManagedObjectReference),
	}

	return l
}

func (l *listResult) WriteTo(w io.Writer) error {
	var err error

	for _, f := range l.Folders {
		if _, err = fmt.Fprintf(w, "%s/\n", f.Name); err != nil {
			return err
		}
	}

	for _, f := range l.Datacenters {
		if _, err = fmt.Fprintf(w, "%s (Datacenter)\n", f.Name); err != nil {
			return err
		}
	}

	for _, f := range l.VirtualMachines {
		if _, err = fmt.Fprintf(w, "%s (VirtualMachine)\n", f.Name); err != nil {
			return err
		}
	}

	for _, f := range l.Networks {
		if _, err = fmt.Fprintf(w, "%s (Network)\n", f.Name); err != nil {
			return err
		}
	}

	for _, f := range l.ComputeResources {
		if _, err = fmt.Fprintf(w, "%s (ComputeResource)\n", f.Name); err != nil {
			return err
		}
	}

	for _, f := range l.Datastores {
		if _, err = fmt.Fprintf(w, "%s (Datastore)\n", f.Name); err != nil {
			return err
		}
	}

	return nil
}
