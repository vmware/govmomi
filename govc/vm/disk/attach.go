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

package disk

import (
	"flag"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type attach struct {
	*flags.DatastoreFlag
	*flags.DiskFlag
	*flags.VirtualMachineFlag

	persist bool
	link    bool

	Client         *govmomi.Client
	Datastore      *govmomi.Datastore
	VirtualMachine *govmomi.VirtualMachine
}

func init() {
	cli.Register("vm.disk.attach", &attach{})
}

func (cmd *attach) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.persist, "persist", true, "Persist attached disk")
	f.BoolVar(&cmd.link, "link", true, "Link specified disk")
}

func (cmd *attach) Process() error { return nil }

func (cmd *attach) Run(f *flag.FlagSet) error {
	var err error

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	cmd.VirtualMachine, err = cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}
	if cmd.VirtualMachine == nil {
		return flag.ErrHelp
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return err
	}
	if cmd.Datastore == nil {
		return flag.ErrHelp
	}

	var mvm mo.VirtualMachine

	err = cmd.Client.Properties(cmd.VirtualMachine.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return err
	}

	disk, err := cmd.DiskFlag.Disk()
	if err != nil {
		return err
	}
	if disk == nil {
		return flag.ErrHelp
	}

	controllerKey, err := FindController(mvm)
	if err != nil {
		return err
	}

	err = AttachDisk(disk, cmd.VirtualMachine, cmd.Datastore, controllerKey, cmd.link, cmd.persist)
	if err != nil {
		return err
	}

	return nil
}
