/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package property_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

// Example to retrieve properties from a single object
func ExampleCollector_RetrieveOne() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		pc := property.DefaultCollector(c)

		obj, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		var vm mo.VirtualMachine
		err = pc.RetrieveOne(ctx, obj.Reference(), []string{"config.version"}, &vm)
		if err != nil {
			return err
		}

		fmt.Printf("hardware version %s", vm.Config.Version)
		return nil
	})
	// Output: hardware version vmx-13
}

func ExampleCollector_Retrieve() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		pc := property.DefaultCollector(c)

		obj, err := find.NewFinder(c).HostSystem(ctx, "DC0_H0")
		if err != nil {
			return err
		}

		var host mo.HostSystem
		err = pc.RetrieveOne(ctx, obj.Reference(), []string{"vm"}, &host)
		if err != nil {
			return err
		}

		var vms []mo.VirtualMachine
		err = pc.Retrieve(ctx, host.Vm, []string{"name"}, &vms)
		if err != nil {
			return err
		}

		fmt.Printf("host has %d vms:", len(vms))
		for i := range vms {
			fmt.Print(" ", vms[i].Name)
		}

		return nil
	})
	// Output: host has 2 vms: DC0_H0_VM0 DC0_H0_VM1
}
