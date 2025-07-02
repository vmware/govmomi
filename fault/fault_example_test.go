// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fault_test

import (
	"fmt"
	"reflect"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleAs_faultByAddress() {
	var (
		err    any
		target *types.InvalidPowerState
	)

	err = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.InvalidPowerState{
				ExistingState:  types.VirtualMachinePowerStatePoweredOn,
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
			},
			LocalizedMessage: "vm must be powered off to encrypt",
		},
	}

	localizedMessage, ok := fault.As(err, &target)

	fmt.Printf("result              = %v\n", ok)
	fmt.Printf("localizedMessage    = %v\n", localizedMessage)
	fmt.Printf("existingPowerState  = %v\n", target.ExistingState)
	fmt.Printf("requestedPowerState = %v\n", target.RequestedState)

	// Output:
	// result              = true
	// localizedMessage    = vm must be powered off to encrypt
	// existingPowerState  = poweredOn
	// requestedPowerState = poweredOff
}

type valueFault uint8

func (f valueFault) GetMethodFault() *types.MethodFault {
	return nil
}

func ExampleAs_faultByValue() {
	var (
		err    any
		target valueFault
	)

	err = &types.SystemError{
		RuntimeFault: types.RuntimeFault{
			MethodFault: types.MethodFault{
				FaultCause: &types.LocalizedMethodFault{
					Fault:            valueFault(1),
					LocalizedMessage: "fault by value",
				},
			},
		},
	}

	localizedMessage, ok := fault.As(err, &target)

	fmt.Printf("result              = %v\n", ok)
	fmt.Printf("localizedMessage    = %v\n", localizedMessage)
	fmt.Printf("value               = %d\n", target)

	// Output:
	// result              = true
	// localizedMessage    = fault by value
	// value               = 1
}

func ExampleIs_baseMethodFault() {
	var (
		err    any
		target types.BaseMethodFault
	)

	err = &types.SystemError{}
	target = &types.SystemError{}

	fmt.Printf("result = %v\n", fault.Is(err, target))

	// Output:
	// result = true
}

func ExampleIs_nestedFault() {
	var (
		err    any
		target types.BaseMethodFault
	)

	err = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault: &types.SystemError{},
					},
				},
			},
		},
	}
	target = &types.SystemError{}

	fmt.Printf("result = %v\n", fault.Is(err, target))

	// Output:
	// result = true
}

func ExampleIs_soapFault() {
	var (
		err    any
		target types.BaseMethodFault
	)

	err = soap.WrapSoapFault(&soap.Fault{
		Detail: struct {
			Fault types.AnyType "xml:\",any,typeattr\""
		}{
			Fault: &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault: &types.SystemError{},
					},
				},
			},
		},
	})
	target = &types.SystemError{}

	fmt.Printf("result = %v\n", fault.Is(err, target))

	// Output:
	// result = true
}

func ExampleIs_vimFault() {
	var (
		err    any
		target types.BaseMethodFault
	)

	err = soap.WrapVimFault(&types.RuntimeFault{
		MethodFault: types.MethodFault{
			FaultCause: &types.LocalizedMethodFault{
				Fault: &types.SystemError{},
			},
		},
	})
	target = &types.SystemError{}

	fmt.Printf("result = %v\n", fault.Is(err, target))

	// Output:
	// result = true
}

func ExampleIn_printAllTypeNamesAndMessages() {
	var (
		err     any
		onFault fault.OnFaultFn
	)

	err = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault:            &types.SystemError{},
						LocalizedMessage: "inner message",
					},
				},
			},
			LocalizedMessage: "outer message",
		},
	}

	onFault = func(
		fault types.BaseMethodFault,
		localizedMessage string,
		localizableMessages []types.LocalizableMessage) bool {

		fmt.Printf("type             = %s\n", reflect.ValueOf(fault).Elem().Type().Name())
		fmt.Printf("localizedMessage = %s\n", localizedMessage)

		// Return false to continue discovering faults.
		return false
	}

	fault.In(err, onFault)

	// Output:
	// type             = RuntimeFault
	// localizedMessage = outer message
	// type             = SystemError
	// localizedMessage = inner message
}

func ExampleIn_printFirstDiscoveredTypeNameAndMessage() {
	var (
		err     any
		onFault fault.OnFaultFn
	)

	err = task.Error{
		LocalizedMethodFault: &types.LocalizedMethodFault{
			Fault: &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault:            &types.SystemError{},
						LocalizedMessage: "inner message",
					},
				},
			},
			LocalizedMessage: "outer message",
		},
	}

	onFault = func(
		fault types.BaseMethodFault,
		localizedMessage string,
		localizableMessages []types.LocalizableMessage) bool {

		fmt.Printf("type             = %s\n", reflect.ValueOf(fault).Elem().Type().Name())
		fmt.Printf("localizedMessage = %s\n", localizedMessage)

		// Return true to force discovery to halt.
		return true
	}

	fault.In(err, onFault)

	// Output:
	// type             = RuntimeFault
	// localizedMessage = outer message
}
