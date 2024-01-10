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
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
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

func ExampleWait() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		pc := property.DefaultCollector(c)

		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		// power off VM after some time
		go func() {
			time.Sleep(time.Millisecond * 100)
			_, err := vm.PowerOff(ctx)
			if err != nil {
				panic(err)
			}
		}()

		return property.Wait(ctx, pc, vm.Reference(), []string{"runtime.powerState"}, func(changes []types.PropertyChange) bool {
			for _, change := range changes {
				state := change.Val.(types.VirtualMachinePowerState)
				fmt.Println(state)
				if state == types.VirtualMachinePowerStatePoweredOff {
					return true
				}
			}

			// continue polling
			return false
		})
	})
	// Output:
	// poweredOn
	// poweredOff
}

func ExampleCollector_WaitForUpdatesEx_addingRemovingPropertyFilters() {
	model := simulator.VPX()
	model.Datacenter = 1
	model.Cluster = 0
	model.Pool = 0
	model.Machine = 1
	model.Autostart = false

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		// Set up the finder and get a VM.
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			return fmt.Errorf("default datacenter not found: %w", err)
		}
		finder.SetDatacenter(datacenter)
		vmList, err := finder.VirtualMachineList(ctx, "*")
		if len(vmList) == 0 {
			return fmt.Errorf("vmList == 0")
		}
		vm := vmList[0]

		pc, err := property.DefaultCollector(c).Create(ctx)
		if err != nil {
			return fmt.Errorf("failed to create new property collector: %w", err)
		}

		// Start a goroutine to wait for power state changes to the VM. They
		// should not be triggered as there is no property filter yet defined.
		chanResult := make(chan any)
		cancelCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			if err := pc.WaitForUpdatesEx(
				cancelCtx,
				property.WaitOptions{},
				func(updates []types.ObjectUpdate) bool {
					return waitForPowerStateChanges(
						cancelCtx,
						vm,
						chanResult,
						updates,
						types.VirtualMachinePowerStatePoweredOff)
				}); err != nil {

				chanResult <- err
				return
			}
		}()

		// Power on the VM to cause a property change.
		if _, err := vm.PowerOn(ctx); err != nil {
			return fmt.Errorf("error while powering on vm: %w", err)
		}

		// The power change should be ignored.
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("poweredOn event not received")
		case result := <-chanResult:
			switch tResult := result.(type) {
			case types.VirtualMachinePowerState:
				return fmt.Errorf("update should not have been received without a property filter")
			case error:
				return fmt.Errorf("error while waiting for updates: %v", tResult)
			}
		}

		// Now create a property filter that will catch the update.
		pf, err := pc.CreateFilter(ctx, getDatacenterToVMFolderFilter(datacenter))
		if err != nil {
			return fmt.Errorf("failed to create dc2vm property filter: %w", err)
		}

		// Power off the VM to cause a property change.
		if _, err := vm.PowerOff(ctx); err != nil {
			return fmt.Errorf("error while powering off vm: %w", err)
		}

		// The power change should now be noticed.
		select {
		case <-time.After(3 * time.Second):
			return fmt.Errorf("timed out while waiting for property update")
		case result := <-chanResult:
			switch tResult := result.(type) {
			case types.VirtualMachinePowerState:
				if tResult != types.VirtualMachinePowerStatePoweredOff {
					return fmt.Errorf("unexpected power state: %v", tResult)
				}
				fmt.Println("poweredOff event received")
			case error:
				return fmt.Errorf("error while waiting for updates: %w", tResult)
			}
		}

		// Destroy the property filter and repeat, and the power change should
		// once again be ignored.
		if err := pf.Destroy(ctx); err != nil {
			return fmt.Errorf("failed to destroy property filter: %w", err)
		}

		// Power on the VM to cause a property change.
		if _, err := vm.PowerOn(ctx); err != nil {
			return fmt.Errorf("error while powering on vm: %w", err)
		}

		// The power change should be ignored.
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("poweredOn event not received")
		case result := <-chanResult:
			switch tResult := result.(type) {
			case types.VirtualMachinePowerState:
				return fmt.Errorf("update should not have been received after property filter was destroyed")
			case error:
				return fmt.Errorf("error while waiting for updates: %v", tResult)
			}
		}

		return nil
	}, model)

	// Output:
	// poweredOn event not received
	// poweredOff event received
	// poweredOn event not received
}

func ExampleCollector_WaitForUpdatesEx_errConcurrentCollector() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		pc := property.DefaultCollector(c)

		waitOptions := property.WaitOptions{
			Options: &types.WaitOptions{
				MaxWaitSeconds: addrOf(int32(1)),
			},
		}

		onUpdatesFn := func(_ []types.ObjectUpdate) bool {
			return false
		}

		waitForChanges := func(chanErr chan error) {
			defer close(chanErr)
			chanErr <- pc.WaitForUpdatesEx(ctx, waitOptions, onUpdatesFn)
		}

		// Start two goroutines that wait for changes, but only one will begin
		// waiting -- the other will return property.ErrConcurrentCollector.
		chanErr1, chanErr2 := make(chan error), make(chan error)
		go waitForChanges(chanErr1)
		go waitForChanges(chanErr2)

		err1 := <-chanErr1
		err2 := <-chanErr2

		if err1 == nil && err2 == nil {
			return fmt.Errorf(
				"one of the WaitForUpdate calls should have returned %s",
				property.ErrConcurrentCollector)
		}

		if err1 == property.ErrConcurrentCollector &&
			err2 == property.ErrConcurrentCollector {

			return fmt.Errorf(
				"both of the WaitForUpdate calls returned %s",
				property.ErrConcurrentCollector)
		}

		fmt.Println("WaitForUpdatesEx call succeeded")
		fmt.Println("WaitForUpdatesEx call returned ErrConcurrentCollector")

		// The third WaitForUpdatesEx call should be able to successfully obtain
		// the lock since the other two calls are completed.
		if err := pc.WaitForUpdatesEx(ctx, waitOptions, onUpdatesFn); err != nil {
			return fmt.Errorf(
				"unexpected error from third call to WaitForUpdatesEx: %s", err)
		}

		fmt.Println("WaitForUpdatesEx call succeeded")

		return nil
	})

	// Output:
	// WaitForUpdatesEx call succeeded
	// WaitForUpdatesEx call returned ErrConcurrentCollector
	// WaitForUpdatesEx call succeeded
}
