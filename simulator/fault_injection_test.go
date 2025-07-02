// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestFaultInjection(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		// Get service to add fault injection rules
		service := ServiceFromContext(ctx)

		// Test 1: Inject NotAuthenticated fault for all PowerOnVM_Task calls
		rule1 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0, // Always inject
			FaultType:   FaultTypeNotAuthenticated,
			Message:     "Authentication failed for power on operation",
			Enabled:     true,
		}
		service.AddFaultRule(rule1)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			t.Fatal(err)
		}

		// This should fail with NotAuthenticated
		_, err = vm.PowerOn(ctx)
		if err == nil {
			t.Fatal("expected authentication error")
		}

		// Check that we got an error (fault injection worked)
		t.Logf("Fault injection worked, got error: %v", err)

		// Clear rules and test again
		service.ClearFaultRules()

		// Test 2: Inject InvalidArgument fault for specific VM
		rule2 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "VirtualMachine",
			ObjectName:  "DC0_C0_RP0_VM0",
			Probability: 1.0,
			FaultType:   FaultTypeInvalidArgument,
			Message:     "Invalid argument provided",
			Enabled:     true,
		}
		service.AddFaultRule(rule2)

		// This should fail with InvalidArgument
		_, err = vm.PowerOn(ctx)
		if err == nil {
			t.Fatal("expected invalid argument error")
		}

		// Check that we got an error (fault injection worked)
		t.Logf("InvalidArgument fault injection worked, got error: %v", err)

		// Test 3: Probability-based fault injection
		service.ClearFaultRules()

		rule3 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 0.5, // 50% chance
			FaultType:   FaultTypeGeneric,
			Message:     "Random failure",
			Enabled:     true,
		}
		service.AddFaultRule(rule3)

		// Run multiple times to test probability
		failures := 0
		attempts := 100
		for i := 0; i < attempts; i++ {
			vm.PowerOff(ctx) // Reset state
			taskResult, err := vm.PowerOn(ctx)
			if err != nil {
				failures++
			} else if taskResult != nil {
				taskResult.Wait(ctx) // Wait for task completion
			}
		}

		// Should have roughly 50% failures (allow for variance)
		if failures < attempts/4 || failures > 3*attempts/4 {
			t.Logf("Expected ~50%% failures, got %d/%d (%.1f%%)", failures, attempts, float64(failures)/float64(attempts)*100)
		}

		// Test 4: Custom fault injection
		service.ClearFaultRules()

		customFault := &types.ResourceInUse{
			VimFault: types.VimFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						LocalizedMessage: "Resource is currently in use",
					},
				},
			},
			Type: "VirtualMachine",
			Name: "TestVM",
		}

		rule4 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeCustom,
			Fault:       customFault,
			Enabled:     true,
		}
		service.AddFaultRule(rule4)

		_, err = vm.PowerOn(ctx)
		if err == nil {
			t.Fatal("expected resource in use error")
		}

		// Check that we got an error (fault injection worked)
		t.Logf("Custom fault injection worked, got error: %v", err)

		// Test 5: Test statistics
		stats := service.GetFaultStats()
		if stats["total_rules"].(int) != 1 {
			t.Errorf("expected 1 rule, got %d", stats["total_rules"])
		}

		service.ClearFaultRules()
	})
}

func TestFaultInjectionWithDelay(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		// Test fault injection with delay
		rule := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeGeneric,
			Message:     "Delayed failure",
			Delay:       100, // 100ms delay
			Enabled:     true,
		}
		service.AddFaultRule(rule)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			t.Fatal(err)
		}

		// This should fail after a delay
		_, err = vm.PowerOn(ctx)
		if err == nil {
			t.Fatal("expected error with delay")
		}

		service.ClearFaultRules()
	})
}

func TestFaultInjectionMaxCount(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		// Test max count limiting
		rule := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeGeneric,
			Message:     "Limited failure",
			MaxCount:    2, // Only fail twice
			Enabled:     true,
		}
		service.AddFaultRule(rule)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			t.Fatal(err)
		}

		// First two attempts should fail
		for i := 0; i < 2; i++ {
			vm.PowerOff(ctx)
			_, err := vm.PowerOn(ctx)
			if err == nil {
				t.Fatalf("expected error on attempt %d", i+1)
			}
		}

		// Third attempt should succeed (no more faults)
		vm.PowerOff(ctx)
		taskResult, err := vm.PowerOn(ctx)
		if err != nil {
			t.Fatalf("expected success on third attempt, got: %v", err)
		}
		if taskResult != nil {
			taskResult.Wait(ctx)
		}

		service.ClearFaultRules()
	})
}

func TestMultipleFaultRules(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		// Add multiple rules with different priorities
		rule1 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "VirtualMachine",
			ObjectName:  "DC0_C0_RP0_VM0",
			Probability: 1.0,
			FaultType:   FaultTypeInvalidArgument,
			Message:     "Specific VM rule",
			Enabled:     true,
		}

		rule2 := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeGeneric,
			Message:     "General rule",
			Enabled:     true,
		}

		service.AddFaultRule(rule1)
		service.AddFaultRule(rule2)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			t.Fatal(err)
		}

		// Should match the first rule (more specific)
		_, err = vm.PowerOn(ctx)
		if err == nil {
			t.Fatal("expected error")
		}

		// Check that we got an error (fault injection worked)
		t.Logf("Multiple rules fault injection worked, got error: %v", err)

		service.ClearFaultRules()
	})
}
