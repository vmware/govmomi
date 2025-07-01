// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

/*
Package simulator provides fault injection capabilities to simulate various
error conditions in vSphere operations for testing and development purposes.

# Fault Injection Overview

The fault injection system allows you to inject various types of faults into
SOAP method calls made to the simulator. This is useful for:

- Testing error handling in client applications
- Simulating network and resource failures
- Validating retry and recovery logic
- Chaos engineering and resilience testing

# Basic Usage

To use fault injection, access the Service from the simulator context and add
fault rules:

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		service := simulator.ServiceFromContext(ctx)

		// Add a rule to inject authentication errors
		rule := &simulator.FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "VirtualMachine",
			ObjectName:  "*",
			Probability: 0.1, // 10% of calls will fail
			FaultType:   simulator.FaultTypeNotAuthenticated,
			Message:     "Authentication failed",
			Enabled:     true,
		}
		service.AddFaultRule(rule)

		// Now VM power on operations will randomly fail
		finder := find.NewFinder(c)
		vm, _ := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		_, err := vm.PowerOn(ctx) // May fail with NotAuthenticated
	})

# Fault Types

The following built-in fault types are available:

- FaultTypeGeneric: Generic system errors
- FaultTypeNotAuthenticated: Authentication failures
- FaultTypeNoPermission: Permission denied errors
- FaultTypeInvalidArgument: Invalid parameter errors
- FaultTypeMethodNotFound: Method not found errors
- FaultTypeManagedObjectNotFound: Object not found errors
- FaultTypeInvalidState: Invalid object state errors
- FaultTypeResourceInUse: Resource busy/in use errors
- FaultTypeInsufficientResourcesFault: Resource exhaustion errors
- FaultTypeNetworkFailure: Network connectivity issues
- FaultTypeTimeout: Operation timeout errors
- FaultTypeCustom: Custom VMware fault (specify in Fault field)

# Rule Matching

Fault injection rules are matched based on:

1. Method name (e.g., "PowerOnVM_Task") - use "*" for all methods
2. Object type (e.g., "VirtualMachine") - use "*" for all types
3. Object name (e.g., "DC0_C0_RP0_VM0") - use "*" for all objects
4. Probability (0.0 to 1.0) - percentage chance the fault will be injected

Rules are evaluated in the order they were added. The first matching rule is
used.

# Advanced Features

## Probabilistic Injection

Control how often faults are injected:

	rule := &simulator.FaultInjectionRule{
		MethodName:  "*",
		ObjectType:  "*",
		ObjectName:  "*",
		Probability: 0.05, // 5% of all operations fail
		FaultType:   simulator.FaultTypeNetworkFailure,
		Enabled:     true,
	}

## Limited Injection Count

Limit how many times a rule triggers:

	rule := &simulator.FaultInjectionRule{
		MethodName:  "PowerOnVM_Task",
		ObjectType:  "*",
		ObjectName:  "*",
		Probability: 1.0,
		FaultType:   simulator.FaultTypeResourceInUse,
		MaxCount:    3, // Only fail first 3 attempts
		Enabled:     true,
	}

## Delays

Add delays before returning faults to simulate slow responses:

	rule := &simulator.FaultInjectionRule{
		MethodName:  "*",
		ObjectType:  "*",
		ObjectName:  "*",
		Probability: 0.02,
		FaultType:   simulator.FaultTypeTimeout,
		Delay:       5000, // 5 second delay
		Enabled:     true,
	}

## Custom Faults

Inject specific VMware fault types:

	customFault := &types.ResourceInUse{
		VimFault: types.VimFault{
			MethodFault: types.MethodFault{
				FaultCause: &types.LocalizedMethodFault{
					LocalizedMessage: "CPU resource pool exhausted",
				},
			},
		},
		Type: "ResourcePool",
		Name: "Production-RP",
	}

	rule := &simulator.FaultInjectionRule{
		MethodName:  "PowerOnVM_Task",
		ObjectType:  "VirtualMachine",
		ObjectName:  "*",
		Probability: 1.0,
		FaultType:   simulator.FaultTypeCustom,
		Fault:       customFault,
		Enabled:     true,
	}

# Management Methods

The Service provides several methods for managing fault injection:

	service := simulator.ServiceFromContext(ctx)

	// Add rules
	service.AddFaultRule(rule)

	// Remove specific rule by index
	service.RemoveFaultRule(0)

	// Get all rules
	rules := service.GetFaultRules()

	// Clear all rules
	service.ClearFaultRules()

	// Get statistics
	stats := service.GetFaultStats()

	// Reset counters
	service.ResetFaultStats()

# Example Scenarios

## Simulating Authentication Service Outage

	authOutageRule := &simulator.FaultInjectionRule{
		MethodName:  "Login",
		ObjectType:  "*",
		ObjectName:  "*",
		Probability: 1.0,
		FaultType:   simulator.FaultTypeNotAuthenticated,
		Message:     "Authentication service unavailable",
		MaxCount:    10, // Outage affects first 10 login attempts
		Enabled:     true,
	}

## Simulating Intermittent Network Issues

	networkRule := &simulator.FaultInjectionRule{
		MethodName:  "*",
		ObjectType:  "*",
		ObjectName:  "*",
		Probability: 0.02, // 2% of operations affected
		FaultType:   simulator.FaultTypeNetworkFailure,
		Delay:       3000, // 3 second timeout
		Enabled:     true,
	}

## Simulating Resource Exhaustion

	resourceRule := &simulator.FaultInjectionRule{
		MethodName:  "PowerOnVM_Task",
		ObjectType:  "VirtualMachine",
		ObjectName:  "*",
		Probability: 0.8, // 80% of power-ons fail
		FaultType:   simulator.FaultTypeInsufficientResourcesFault,
		Message:     "Insufficient memory resources",
		Enabled:     true,
	}

# Thread Safety

The fault injection system is thread-safe and can be used from multiple
goroutines simultaneously. All operations on fault rules are protected by
internal mutexes.

# Performance Impact

Fault injection adds minimal overhead to normal operations. Rules are evaluated
efficiently, and random number generation is optimized for concurrent access.
When no rules are enabled, the performance impact is negligible.

See also: https://github.com/vmware/govmomi/blob/main/vcsim/README.md
*/
package simulator
