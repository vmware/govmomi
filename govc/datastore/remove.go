/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

	"golang.org/x/net/context"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type remove struct {
	*flags.HostSystemFlag
	*flags.DatastoreFlag
}

func init() {
	cli.Register("datastore.remove", &remove{})
}

func (cmd *remove) Register(f *flag.FlagSet) {}

func (cmd *remove) Process() error { return nil }

func (cmd *remove) Usage() string {
	return "HOST..."
}

func (cmd *remove) Description() string {
	return `Remove datastore from HOST.
Example:
govc datastore.remove -ds nfsDatastore cluster1
`
}

func (cmd *remove) Run(f *flag.FlagSet) error {
	ctx := context.TODO()

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		hds, err := host.ConfigManager().DatastoreSystem(ctx)
		if err != nil {
			return err
		}

		err = hds.Remove(ctx, ds)
		if err != nil {
			return err
		}
	}

	return nil
}
