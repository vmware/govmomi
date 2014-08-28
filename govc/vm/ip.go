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
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type ip struct {
	*flags.OutputFlag
	*flags.SearchFlag
}

func init() {
	cli.Register("vm.ip", &ip{})
}

func (cmd *ip) Register(f *flag.FlagSet) {
	cmd.SearchFlag = flags.NewSearchFlag(flags.SearchVirtualMachines)
}

func (cmd *ip) Process() error { return nil }

func (cmd *ip) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		ip, err := vm.WaitForIP(c)
		if err != nil {
			return err
		}

		// TODO(PN): Display inventory path to VM
		fmt.Fprintf(cmd, "%s\n", ip)
	}

	return nil
}
