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

package event_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleManager_Events() {
	simulator.Example(func(ctx context.Context, c *vim25.Client) error {
		m := event.NewManager(c)

		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		objs := []types.ManagedObjectReference{vm.Reference()}

		return m.Events(ctx, objs, 10, false, false, func(ref types.ManagedObjectReference, events []types.BaseEvent) error {
			event.Sort(events)
			for _, event := range events {
				fmt.Printf("%T\n", event)
			}
			return nil
		})
	})
	// Output:
	// *types.VmBeingCreatedEvent
	// *types.VmInstanceUuidAssignedEvent
	// *types.VmUuidAssignedEvent
	// *types.VmCreatedEvent
	// *types.VmStartingEvent
	// *types.VmPoweredOnEvent
}
