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
)

type download struct {
	*flags.ClientFlag
	*flags.DatastorePath
}

func init() {
	cli.Register(&download{})
}

func (c *download) Register(f *flag.FlagSet) {
}

func (c *download) Process() error {
	return nil
}

func (c *download) Run(f *flag.FlagSet) error {
	client, err := c.Client()
	if err != nil {
		return err
	}

	ds, err := c.Datastore()
	if err != nil {
		return err
	}

	u, err := c.URL()
	if err != nil {
		return err
	}

	return ds.UploadFile(client, f.Arg(0), u)
}
