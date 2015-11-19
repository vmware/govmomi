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

package nas

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type create struct {
	*flags.HostSystemFlag

	Force bool

	types.HostNasVolumeSpec
}

func init() {
	cli.Register("datastore.nas.create", &create{})
}

func (cmd *create) Register(f *flag.FlagSet) {
	fsTypes := []string{
		string(types.HostFileSystemVolumeFileSystemTypeNFS),
		string(types.HostFileSystemVolumeFileSystemTypeNFS41),
		string(types.HostFileSystemVolumeFileSystemTypeCIFS),
	}

	modes := []string{
		string(types.HostMountModeReadOnly),
		string(types.HostMountModeReadWrite),
	}

	f.BoolVar(&cmd.Force, "force", false, "Ignore DuplicateName error if datastore is already mounted on a host")

	f.StringVar(&cmd.LocalPath, "local-path", "", "Name of the NAS datastore")
	f.StringVar(&cmd.RemoteHost, "remote-host", "", "Remote hostname of the NAS datastore")
	f.StringVar(&cmd.RemotePath, "remote-path", "", "Remote path of the NFS mount point")
	f.StringVar(&cmd.AccessMode, "mode", modes[0],
		fmt.Sprintf("Access mode for the mount point (%s)", strings.Join(modes, "|")))
	f.StringVar(&cmd.Type, "type", fsTypes[0],
		fmt.Sprintf("Type of the the NAS volume (%s)", strings.Join(fsTypes, "|")))
	f.StringVar(&cmd.UserName, "username", "", "Username to use when connecting (CIFS only)")
	f.StringVar(&cmd.Password, "password", "", "Password to use when connecting (CIFS only)")
}

func (cmd *create) Process() error { return nil }

func (cmd *create) Usage() string {
	return "HOST..."
}

func (cmd *create) Description() string {
	return `Create datastore on HOST.
Example:
govc datastore.nas.create -local-path nfsDatastore -remote-host 10.143.2.232 -remote-path /share cluster1
`
}

func (cmd *create) Run(f *flag.FlagSet) error {
	var ctx = context.Background()

	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	object := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: fmt.Sprintf("%s:%s", cmd.RemoteHost, cmd.RemotePath),
	}

	for _, host := range hosts {
		ds, err := host.ConfigManager().DatastoreSystem(ctx)
		if err != nil {
			return err
		}

		_, err = ds.CreateNasDatastore(ctx, cmd.HostNasVolumeSpec)
		if err != nil {
			if soap.IsSoapFault(err) {
				switch fault := soap.ToSoapFault(err).VimFault().(type) {
				case types.PlatformConfigFault:
					if len(fault.FaultMessage) != 0 {
						return errors.New(fault.FaultMessage[0].Message)
					}
				case types.DuplicateName:
					if cmd.Force && fault.Object == object {
						fmt.Fprintf(os.Stderr, "%s: '%s' already mounted\n",
							host.InventoryPath, cmd.LocalPath)
						continue
					}
				}
			}

			return fmt.Errorf("%s: %s", host.InventoryPath, err)
		}
	}

	return nil
}
