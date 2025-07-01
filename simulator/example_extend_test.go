// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"fmt"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// BusyVM changes the behavior of simulator.VirtualMachine
type BusyVM struct {
	*simulator.VirtualMachine
}

// Override simulator.VirtualMachine.PowerOffVMTask to inject faults
func (vm *BusyVM) PowerOffVMTask(ctx *simulator.Context, req *types.PowerOffVM_Task) soap.HasFault {
	task := simulator.CreateTask(req.This, "powerOff", func(*simulator.Task) (types.AnyType, types.BaseMethodFault) {
		return nil, &types.TaskInProgress{}
	})

	return &methods.PowerOffVM_TaskBody{
		Res: &types.PowerOffVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

// Add AcquireTicket method, not implemented by simulator.VirtualMachine
func (vm *BusyVM) AcquireTicket(req *types.AcquireTicket) soap.HasFault {
	body := &methods.AcquireTicketBody{}

	if req.TicketType != "mks" {
		body.Fault_ = simulator.Fault("", &types.InvalidArgument{})
	}

	body.Res = &types.AcquireTicketResponse{
		Returnval: types.VirtualMachineTicket{
			Ticket: "welcome 2 the machine",
		},
	}

	return body
}

// Example of extending the simulator to inject faults.
func Example() {
	ctx := context.Background()
	model := simulator.VPX()

	defer model.Remove()
	_ = model.Create()

	s := model.Service.NewServer()
	defer s.Close()

	// NewClient connects to s.URL over https and invokes 2 SOAP methods (RetrieveServiceContent + Login)
	c, _ := govmomi.NewClient(ctx, s.URL, true)

	// Shortcut to choose any VM, rather than using the more verbose Finder or ContainerView.
	obj := model.Map().Any("VirtualMachine").(*simulator.VirtualMachine)
	// Validate VM is powered on
	if obj.Runtime.PowerState != "poweredOn" {
		log.Fatal(obj.Runtime.PowerState)
	}

	// Wrap the existing vm object, using the same vm.Self (ManagedObjectReference) value as the Map key.
	model.Map().Put(&BusyVM{obj})

	vm := object.NewVirtualMachine(c.Client, obj.Reference())

	// Start a PowerOff task using the SOAP client.
	task, _ := vm.PowerOff(ctx)

	// Wait for task completion, expecting failure.
	err := task.Wait(ctx)
	if err == nil {
		log.Fatal("expected error")
	}

	// Invalid ticket type, expecting failure.
	_, err = vm.AcquireTicket(ctx, "pks")
	if err == nil {
		log.Fatal("expected error")
	}

	mks, _ := vm.AcquireTicket(ctx, "mks")

	fmt.Println(mks.Ticket, obj.Runtime.PowerState)
	// Output: welcome 2 the machine poweredOn
}
