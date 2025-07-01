// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package event_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// ensure event.Manager implements the mo.Reference interface
var _ mo.Reference = new(event.Manager)

func ExampleManager_Events() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
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
