// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// ExampleFaultInjector demonstrates how to use the fault injection system.
func ExampleFaultInjector() {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		// Get the simulator service to configure fault injection
		service := simulator.ServiceFromContext(ctx)

		// Example 1: Always fail PowerOn operations with authentication error
		authRule := &simulator.FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*", // Match any object type
			ObjectName:  "*", // Match any object name
			Probability: 1.0, // Always inject (100% probability)
			FaultType:   simulator.FaultTypeNotAuthenticated,
			Message:     "Simulated authentication failure",
			Enabled:     true,
		}
		service.AddFaultRule(authRule)

		finder := find.NewFinder(c)
		vm, _ := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")

		// This will fail with NotAuthenticated
		_, err := vm.PowerOn(ctx)
		if err != nil {
			fmt.Printf("PowerOn failed as expected: %v\n", err)
		}

		// Example 2: Probabilistic failures for specific VM
		service.ClearFaultRules()

		probabilisticRule := &simulator.FaultInjectionRule{
			MethodName:  "*",              // Match any method
			ObjectType:  "VirtualMachine", // Only VMs
			ObjectName:  "DC0_C0_RP0_VM0", // Specific VM
			Probability: 0.3,              // 30% chance of failure
			FaultType:   simulator.FaultTypeInvalidState,
			Message:     "VM is in an invalid state",
			Enabled:     true,
		}
		service.AddFaultRule(probabilisticRule)

		// Example 3: Custom fault with specific VMware error
		customFault := &types.ResourceInUse{
			VimFault: types.VimFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						LocalizedMessage: "CPU resource is currently allocated to another VM",
					},
				},
			},
			Type: "VirtualMachine",
			Name: "DC0_C0_RP0_VM0",
		}

		customRule := &simulator.FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "VirtualMachine",
			ObjectName:  "DC0_C0_RP0_VM0",
			Probability: 1.0,
			FaultType:   simulator.FaultTypeCustom,
			Fault:       customFault, // Use custom VMware fault
			Delay:       500,         // Add 500ms delay before fault
			MaxCount:    3,           // Only inject 3 times
			Enabled:     true,
		}
		service.AddFaultRule(customRule)

		// Example 4: Multiple rules for different scenarios
		service.ClearFaultRules()

		// Rule for all VM power operations
		powerRule := &simulator.FaultInjectionRule{
			MethodName:  "Power*", // Wildcard matching would need to be implemented
			ObjectType:  "VirtualMachine",
			ObjectName:  "*",
			Probability: 0.1,
			FaultType:   simulator.FaultTypeGeneric,
			Message:     "Power operation failed randomly",
			Enabled:     true,
		}

		// Rule for network configuration
		networkRule := &simulator.FaultInjectionRule{
			MethodName:  "*",
			ObjectType:  "Network",
			ObjectName:  "*",
			Probability: 0.05,
			FaultType:   simulator.FaultTypeNoPermission,
			Message:     "Insufficient permissions for network operation",
			Enabled:     true,
		}

		service.AddFaultRule(powerRule)
		service.AddFaultRule(networkRule)

		// Get and display statistics
		stats := service.GetFaultStats()
		fmt.Printf("Total rules: %d\n", stats["total_rules"])
		fmt.Printf("Enabled rules: %d\n", stats["enabled_rules"])

		// Clean up
		service.ClearFaultRules()

		fmt.Println("Fault injection example completed")
	})

	// Output:
	// PowerOn failed as expected: ServerFaultCode: Simulated authentication failure
	// Total rules: 2
	// Enabled rules: 2
	// Fault injection example completed
}
