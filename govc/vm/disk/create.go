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
	"errors"
	"flag"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type create struct {
	*flags.VirtualMachineFlag

	Name  string
	Bytes ByteValue

	Client         *govmomi.Client
	VirtualMachine *govmomi.VirtualMachine
}

func init() {
	cli.Register("vm.disk.create", &create{})
}

func (cmd *create) Register(f *flag.FlagSet) {
	err := (&cmd.Bytes).Set("10G")
	if err != nil {
		panic(err)
	}

	f.StringVar(&cmd.Name, "name", "", "Name for new disk")
	f.Var(&cmd.Bytes, "size", "Size of new disk")
}

func (cmd *create) Process() error { return nil }

func (cmd *create) Run(f *flag.FlagSet) error {
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
		return errors.New("please specify a vm")
	}

	var mvm mo.VirtualMachine

	err = cmd.Client.Properties(cmd.VirtualMachine.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return err
	}

	dev, err := FindDisk(cmd.Name, mvm)
	if err != nil {
		return err
	}

	if dev == nil {
		cmd.Log("Creating disk\n")

		controllerKey, err := FindController(mvm)
		if err != nil {
			return err
		}

		err = CreateDisk(cmd.Name, cmd.Bytes, controllerKey, cmd.VirtualMachine)
		if err != nil {
			return err
		}
	} else {
		cmd.Log("Disk already present\n")
	}

	return nil
}
