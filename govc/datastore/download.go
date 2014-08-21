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
	*flags.DatastorePathFlag
}

func init() {
	cli.Register(&download{})
}

func (cmd *download) Register(f *flag.FlagSet) {
}

func (cmd *download) Process() error {
	return nil
}

func (cmd *download) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	u, err := cmd.URL()
	if err != nil {
		return err
	}

	return c.Client.DownloadFile(f.Arg(0), u)
}
