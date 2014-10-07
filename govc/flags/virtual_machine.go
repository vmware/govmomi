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

package flags

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
)

type VirtualMachineFlag struct {
	*ClientFlag
	*DatacenterFlag
	*SearchFlag
	*ListFlag

	register sync.Once
	name     string
	vm       *govmomi.VirtualMachine
}

func (flag *VirtualMachineFlag) Register(f *flag.FlagSet) {
	flag.SearchFlag = NewSearchFlag(SearchVirtualMachines)

	flag.register.Do(func() {
		env := "GOVC_VM"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Virtual machine [%s]", env)
		f.StringVar(&flag.name, "vm", value, usage)
	})
}

func (flag *VirtualMachineFlag) Process() error {
	return nil
}

func (flag *VirtualMachineFlag) findVirtualMachine(path string) ([]*govmomi.VirtualMachine, error) {
	c, err := flag.ClientFlag.Client()
	if err != nil {
		return nil, err
	}

	relativeFunc := func() (govmomi.Reference, error) {
		dc, err := flag.DatacenterFlag.Datacenter()
		if err != nil {
			return nil, err
		}

		f, err := dc.Folders(c)
		if err != nil {
			return nil, err
		}

		return f.VmFolder, nil
	}

	es, err := flag.ListFlag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var vms []*govmomi.VirtualMachine
	for _, e := range es {
		switch o := e.Object.(type) {
		case mo.VirtualMachine:
			vm := govmomi.NewVirtualMachine(c, o.Reference())
			vms = append(vms, vm)
		}
	}

	return vms, nil
}

func (flag *VirtualMachineFlag) findSpecifiedVirtualMachine(path string) (*govmomi.VirtualMachine, error) {
	vms, err := flag.findVirtualMachine(path)
	if err != nil {
		return nil, err
	}

	if len(vms) == 0 {
		return nil, errors.New("no such vm")
	}

	if len(vms) > 1 {
		return nil, errors.New("path resolves to multiple vms")
	}

	flag.vm = vms[0]
	return flag.vm, nil
}

func (flag *VirtualMachineFlag) VirtualMachine() (*govmomi.VirtualMachine, error) {
	if flag.vm != nil {
		return flag.vm, nil
	}

	// Use search flags if specified.
	if flag.SearchFlag.IsSet() {
		vm, err := flag.SearchFlag.VirtualMachine()
		if err != nil {
			return nil, err
		}

		flag.vm = vm
		return flag.vm, nil
	}

	// Never look for a default virtual machine.
	if flag.name == "" {
		return nil, nil
	}

	return flag.findSpecifiedVirtualMachine(flag.name)
}
