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

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type destroy struct {
	*flags.ClientFlag
	*flags.SearchFlag

	detach bool

	Client *govmomi.Client
}

func init() {
	cli.Register(&destroy{})
}

func (cmd *destroy) Register(f *flag.FlagSet) {
	cmd.SearchFlag = flags.NewSearchFlag(flags.SearchVirtualMachines)
	f.BoolVar(&cmd.detach, "detach", true, "Detach disks before destroying VM")
}

func (cmd *destroy) Process() error { return nil }

func (cmd *destroy) Run(f *flag.FlagSet) error {
	var err error

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	vms, err := cmd.SearchFlag.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		task, err := vm.PowerOff(cmd.Client)
		if err != nil {
			return err
		}

		// Ignore error since the VM may already been in powered off state.
		// vm.Destroy will fail if the VM is still powered on.
		_ = task.Wait()

		// Detach disks if necessary.
		if cmd.detach {
			err = cmd.DetachDisks(vm)
			if err != nil {
				return err
			}
		}

		task, err = vm.Destroy(cmd.Client)
		if err != nil {
			return err
		}

		err = task.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *destroy) DetachDisks(vm *govmomi.VirtualMachine) error {
	var mvm mo.VirtualMachine

	// TODO(PN): Use `config.hardware` here, see issue #44.
	err := cmd.Client.Properties(vm.Reference(), []string{"config"}, &mvm)
	if err != nil {
		return err
	}

	spec := new(configSpec)
	spec.RemoveDisks(&mvm)

	task, err := vm.Reconfigure(cmd.Client, spec.ToSpec())
	if err != nil {
		return err
	}

	err = task.Wait()
	if err != nil {
		return err
	}

	return nil
}
