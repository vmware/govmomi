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
	"reflect"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type list struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.OutputFlag
}

func init() {
	cli.Register(&list{})
}

func (c *list) Register(f *flag.FlagSet) {}

func (c *list) Process() error { return nil }

func (c *list) Run(f *flag.FlagSet) error {
	client, err := c.Client()
	if err != nil {
		return err
	}

	dc, err := c.Datacenter()
	if err != nil {
		return err
	}

	folders, err := dc.Folders(client)
	if err != nil {
		return err
	}

	es, err := folders.VmFolder.Children(client)
	if err != nil {
		return err
	}

	result := listResult{}
	for _, e := range es {
		switch t := e.(type) {
		case *govmomi.Folder:
			var m mo.Folder

			err = client.Properties(e.Reference(), nil, &m)
			if err != nil {
				return err
			}

			result.Folders = append(result.Folders, m)
		case *govmomi.VirtualMachine:
			var m mo.VirtualMachine

			err = client.Properties(e.Reference(), nil, &m)
			if err != nil {
				return err
			}

			result.VirtualMachines = append(result.VirtualMachines, m)
		default:
			panic("unexpected type: " + reflect.TypeOf(t).String())
		}
	}

	return c.WriteResult(&result)
}

type listResult struct {
	Folders         []mo.Folder
	VirtualMachines []mo.VirtualMachine
}

func (l *listResult) WriteTo(w io.Writer) error {
	var err error

	if len(l.Folders) > 0 {
		if _, err = fmt.Fprintf(w, "## Folders:\n"); err != nil {
			return err
		}
		for _, f := range l.Folders {
			if _, err = fmt.Fprintf(w, "%s/", f.Name); err != nil {
				return err
			}
		}
	}

	if len(l.VirtualMachines) > 0 {
		if _, err = fmt.Fprintf(w, "## Virtual Machines:\n"); err != nil {
			return err
		}
		for _, m := range l.VirtualMachines {
			if _, err = fmt.Fprintf(w, "%s %s\n", m.Name, m.Runtime.PowerState); err != nil {
				return err
			}
		}
	}

	return nil
}
