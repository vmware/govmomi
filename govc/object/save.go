/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package object

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type save struct {
	*flags.FolderFlag

	n       int
	dir     string
	force   bool
	verbose bool
	recurse bool
	one     bool
	kind    kinds
	summary map[string]int
}

func init() {
	cli.Register("object.save", &save{})
}

func (cmd *save) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	f.BoolVar(&cmd.one, "1", false, "Save ROOT only, without its children")
	f.StringVar(&cmd.dir, "d", "", "Save objects in directory")
	f.BoolVar(&cmd.force, "f", false, "Remove existing object directory")
	f.BoolVar(&cmd.recurse, "r", true, "Include children of the container view root")
	f.Var(&cmd.kind, "type", "Resource types to save.  Defaults to all types")
	f.BoolVar(&cmd.verbose, "v", false, "Verbose output")
}

func (cmd *save) Usage() string {
	return "[PATH]"
}

func (cmd *save) Description() string {
	return `Save managed objects.

By default, the object tree and all properties are saved, starting at PATH.
PATH defaults to ServiceContent, but can be specified to save a subset of objects.
The primary use case for this command is to save inventory from a live vCenter and
load it into a vcsim instance.

Examples:
  govc object.save -d my-vcenter
  vcsim -load my-vcenter`
}

// write encodes data to file name
func (cmd *save) write(name string, data interface{}) error {
	f, err := os.Create(filepath.Join(cmd.dir, name) + ".xml")
	if err != nil {
		return err
	}
	e := xml.NewEncoder(f)
	e.Indent("", "  ")
	if err = e.Encode(data); err != nil {
		_ = f.Close()
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

type saveMethod struct {
	Name string
	Data interface{}
}

func saveDVS(ctx context.Context, c *vim25.Client, ref types.ManagedObjectReference) ([]saveMethod, error) {
	res, err := methods.FetchDVPorts(ctx, c, &types.FetchDVPorts{This: ref})
	if err != nil {
		return nil, err
	}
	return []saveMethod{{"FetchDVPorts", res}}, nil
}

// saveObjects maps object types to functions that can save data that isn't available via the PropertyCollector
var saveObjects = map[string]func(context.Context, *vim25.Client, types.ManagedObjectReference) ([]saveMethod, error){
	"VmwareDistributedVirtualSwitch": saveDVS,
}

func (cmd *save) save(content []types.ObjectContent) error {
	for _, x := range content {
		x.MissingSet = nil // drop any NoPermission faults
		cmd.summary[x.Obj.Type]++
		if cmd.verbose {
			fmt.Printf("Saving %s...", x.Obj)
		}
		ref := x.Obj.Encode()
		name := fmt.Sprintf("%04d-%s", cmd.n, ref)
		cmd.n++
		if err := cmd.write(name, x); err != nil {
			return err
		}
		if cmd.verbose {
			fmt.Println("ok")
		}

		c, _ := cmd.Client()
		if method, ok := saveObjects[x.Obj.Type]; ok {
			objs, err := method(context.Background(), c, x.Obj)
			if err != nil {
				return err
			}
			dir := filepath.Join(cmd.dir, ref)
			if err = os.Mkdir(dir, 0755); err != nil {
				return err
			}
			for _, obj := range objs {
				err = cmd.write(filepath.Join(ref, obj.Name), obj.Data)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (cmd *save) Run(ctx context.Context, f *flag.FlagSet) error {
	cmd.summary = make(map[string]int)
	c, err := cmd.Client()
	if err != nil {
		return err
	}
	if cmd.dir == "" {
		u := c.URL()
		name := u.Fragment
		if name == "" {
			name = u.Hostname()
		}
		cmd.dir = "vcsim-" + name
	}
	mkdir := os.Mkdir
	if cmd.force {
		mkdir = os.MkdirAll
	}
	if err := mkdir(cmd.dir, 0755); err != nil {
		return err
	}

	var content []types.ObjectContent
	pc := property.DefaultCollector(c)
	root := vim25.ServiceInstance
	if f.NArg() == 1 {
		root, err = cmd.ManagedObject(ctx, f.Arg(0))
		if err != nil {
			if !root.FromString(f.Arg(0)) {
				return err
			}
		}
		if cmd.one {
			err = pc.RetrieveOne(ctx, root, nil, &content)
			if err != nil {
				return nil
			}
			if err = cmd.save(content); err != nil {
				return err
			}
			return nil
		}
	}

	req := types.RetrievePropertiesEx{
		This:    pc.Reference(),
		Options: types.RetrieveOptions{MaxObjects: 10},
	}

	if root == vim25.ServiceInstance {
		err := pc.RetrieveOne(ctx, root, []string{"content"}, &content)
		if err != nil {
			return nil
		}
		if err = cmd.save(content); err != nil {
			return err
		}
		if cmd.one {
			return nil
		}

		root = c.ServiceContent.RootFolder

		for _, p := range content[0].PropSet {
			if c, ok := p.Val.(types.ServiceContent); ok {
				for _, ref := range mo.References(c) {
					all := types.NewBool(true)
					switch ref.Type {
					case "LicenseManager", "ServiceManager":
						all = nil
					}
					req.SpecSet = append(req.SpecSet, types.PropertyFilterSpec{
						ObjectSet: []types.ObjectSpec{{
							Obj: ref,
						}},
						PropSet: []types.PropertySpec{{
							Type:    ref.Type,
							All:     all,
							PathSet: nil,
						}},
					})
				}
				break
			}
		}
	}

	m := view.NewManager(c)
	v, err := m.CreateContainerView(ctx, root, cmd.kind, cmd.recurse)
	if err != nil {
		return err
	}

	defer func() {
		_ = v.Destroy(ctx)
	}()

	all := types.NewBool(true)
	req.SpecSet = append(req.SpecSet, types.PropertyFilterSpec{
		ObjectSet: []types.ObjectSpec{{
			Obj:  v.Reference(),
			Skip: types.NewBool(false),
			SelectSet: []types.BaseSelectionSpec{
				&types.TraversalSpec{
					Type: v.Reference().Type,
					Path: "view",
					SelectSet: []types.BaseSelectionSpec{
						&types.SelectionSpec{
							Name: "computeTraversalSpec",
						},
						&types.SelectionSpec{
							Name: "datastoreTraversalSpec",
						},
						&types.SelectionSpec{
							Name: "hostDatastoreSystemTraversalSpec",
						},
						&types.SelectionSpec{
							Name: "hostNetworkSystemTraversalSpec",
						},
						&types.SelectionSpec{
							Name: "hostVirtualNicManagerTraversalSpec",
						},
						&types.SelectionSpec{
							Name: "entityTraversalSpec",
						},
					},
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "computeTraversalSpec",
					},
					Type: "ComputeResource",
					Path: "environmentBrowser",
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "datastoreTraversalSpec",
					},
					Type: "Datastore",
					Path: "browser",
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "hostNetworkSystemTraversalSpec",
					},
					Type: "HostSystem",
					Path: "configManager.networkSystem",
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "hostVirtualNicManagerTraversalSpec",
					},
					Type: "HostSystem",
					Path: "configManager.virtualNicManager",
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "hostDatastoreSystemTraversalSpec",
					},
					Type: "HostSystem",
					Path: "configManager.datastoreSystem",
				},
				&types.TraversalSpec{
					SelectionSpec: types.SelectionSpec{
						Name: "entityTraversalSpec",
					},
					Type: "ManagedEntity",
					Path: "recentTask",
				},
			},
		}},
		PropSet: []types.PropertySpec{
			{Type: "EnvironmentBrowser", All: all},
			{Type: "HostDatastoreBrowser", All: all},
			{Type: "HostDatastoreSystem", All: all},
			{Type: "HostNetworkSystem", All: all},
			{Type: "HostVirtualNicManager", All: all},
			{Type: "ManagedEntity", All: all},
			{Type: "Task", All: all},
		},
	})

	res, err := methods.RetrievePropertiesEx(ctx, c, &req)
	if err != nil {
		return err
	}
	if err = cmd.save(res.Returnval.Objects); err != nil {
		return err
	}

	token := res.Returnval.Token
	for token != "" {
		cres, err := methods.ContinueRetrievePropertiesEx(ctx, c, &types.ContinueRetrievePropertiesEx{
			This:  req.This,
			Token: token,
		})
		if err != nil {
			return err
		}
		token = cres.Returnval.Token
		if err = cmd.save(cres.Returnval.Objects); err != nil {
			return err
		}
	}

	var summary []string
	for k, v := range cmd.summary {
		if v == 1 && !cmd.verbose {
			continue
		}
		summary = append(summary, fmt.Sprintf("%s: %d", k, v))
	}
	sort.Strings(summary)

	s := ", including"
	if cmd.verbose {
		s = ""
	}
	fmt.Printf("Saved %d total objects to %q%s:\n", cmd.n, cmd.dir, s)
	for i := range summary {
		fmt.Println(summary[i])
	}

	return nil
}
