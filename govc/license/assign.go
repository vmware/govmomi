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

package license

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/license"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type assign struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.HostSystemFlag

	name   string
	remove bool
}

func init() {
	cli.Register("license.assign", &assign{})
}

func (cmd *assign) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.name, "name", "", "Display name")
	f.BoolVar(&cmd.remove, "remove", false, "Remove assignment")
}

func (cmd *assign) Process() error { return nil }

func (cmd *assign) Usage() string {
	return "KEY"
}

func (cmd *assign) Run(f *flag.FlagSet) error {
	ctx := context.TODO()

	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	key := f.Arg(0)

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := license.NewManager(client).AssignmentManager(ctx)
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	var id string

	if host == nil {
		// Default to vCenter UUID
		id = client.ServiceContent.About.InstanceUuid
	} else {
		id = host.Reference().Value
	}

	if cmd.remove {
		return m.Remove(ctx, id)
	}

	info, err := m.Update(ctx, id, key, cmd.name)
	if err != nil {
		return err
	}

	return cmd.WriteResult(licenseOutput([]types.LicenseManagerLicenseInfo{*info}))
}
