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
	"errors"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type power struct {
	*flags.ClientFlag
	*flags.SearchFlag

	On    bool
	Off   bool
	Force bool
}

func init() {
	flag := power{
		SearchFlag: flags.NewSearchFlag(flags.SearchVirtualMachines),
	}

	cli.Register(&flag)
}

func (cmd *power) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.On, "on", false, "Power on")
	f.BoolVar(&cmd.Off, "off", false, "Power off")
	f.BoolVar(&cmd.Force, "force", false, "Force (ignore state error)")
}

func (cmd *power) Process() error {
	if !cmd.On && !cmd.Off || (cmd.On && cmd.Off) {
		return errors.New("specify -on OR -off")
	}
	return nil
}

func (cmd *power) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Arg(0))
	if err != nil {
		return err
	}

	for _, vm := range vms {
		switch {
		case cmd.On:
			err = vm.PowerOn(c)
		case cmd.Off:
			err = vm.PowerOff(c)
		}

		if !cmd.Force && err != nil {
			return err
		}
	}

	return nil
}
