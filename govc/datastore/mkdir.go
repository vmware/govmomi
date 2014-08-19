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

package datastore

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type mkdir struct {
	*flags.DatastorePathFlag

	createParents bool
}

func init() {
	cli.Register(&mkdir{})
}

func (c *mkdir) Register(f *flag.FlagSet) {
	f.BoolVar(&c.createParents, "p", false, "Create intermediate directories as needed")
}

func (c *mkdir) Process() error {
	return nil
}

func (c *mkdir) Run(f *flag.FlagSet) error {
	client, err := c.Client()
	if err != nil {
		return err
	}

	dc, err := c.Datacenter()
	if err != nil {
		return err
	}

	name, err := c.Name()
	if err != nil {
		return err
	}

	err = client.FileManager().MakeDirectory(name, dc, c.createParents)

	// ignore EEXIST if -p flag is given
	if err != nil && c.createParents {
		if merr, ok := err.(methods.Error); ok && merr.IsFault() {
			if _, ok := merr.Fault().(types.FileAlreadyExists); ok {
				return nil
			}
		}
	}

	return err
}
