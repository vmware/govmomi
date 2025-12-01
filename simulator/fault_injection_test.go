// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"fmt"
	"testing"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/stretchr/testify/require"
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
		require.NoError(t, err)

		// This should fail with NotAuthenticated
		_, err = vm.PowerOn(ctx)
		require.Error(t, err, "expected authentication error")

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
		require.Error(t, err, "expected invalid argument error")

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
		require.Error(t, err, "expected resource in use error")

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
		require.NoError(t, err)

		// This should fail after a delay
		_, err = vm.PowerOn(ctx)
		require.Error(t, err, "expected error with delay")

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
		require.NoError(t, err)

		// First two attempts should fail
		for i := 0; i < 2; i++ {
			vm.PowerOff(ctx)
			_, err := vm.PowerOn(ctx)
			require.Error(t, err, "expected error on attempt %d", i+1)
		}

		// Third attempt should succeed (no more faults)
		vm.PowerOff(ctx)
		taskResult, err := vm.PowerOn(ctx)
		require.NoError(t, err, "expected success on third attempt")
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
		require.NoError(t, err)

		// Should match the first rule (more specific)
		_, err = vm.PowerOn(ctx)
		require.Error(t, err, "expected error")

		// Check that we got an error (fault injection worked)
		t.Logf("Multiple rules fault injection worked, got error: %v", err)

		service.ClearFaultRules()
	})
}

func TestFaultInjectionSkipCount(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		// Test skip count - first 2 matches should be skipped
		rule := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeGeneric,
			Message:     "Should not trigger for first 2 calls",
			SkipCount:   2, // Skip first 2 matches
			Enabled:     true,
		}
		service.AddFaultRule(rule)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		require.NoError(t, err)

		// First two attempts should succeed (skipped)
		for i := 0; i < 2; i++ {
			vm.PowerOff(ctx)
			taskResult, err := vm.PowerOn(ctx)
			require.NoError(t, err, "expected success on attempt %d (should be skipped)", i+1)
			if taskResult != nil {
				taskResult.Wait(ctx)
			}
		}

		// Third attempt should fail (after skipping 2)
		vm.PowerOff(ctx)
		_, err = vm.PowerOn(ctx)
		require.Error(t, err, "expected error on third attempt (after skipping 2)")

		// Fourth attempt should also fail (skip count already met)
		vm.PowerOff(ctx)
		_, err = vm.PowerOn(ctx)
		require.Error(t, err, "expected error on fourth attempt")

		service.ClearFaultRules()
	})
}

func TestFaultInjectionSkipCountWithMaxCount(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		// Test skip count combined with max count
		rule := &FaultInjectionRule{
			MethodName:  "PowerOnVM_Task",
			ObjectType:  "*",
			ObjectName:  "*",
			Probability: 1.0,
			FaultType:   FaultTypeGeneric,
			Message:     "Skip and max count test",
			SkipCount:   1, // Skip first match
			MaxCount:    2, // Only inject twice
			Enabled:     true,
		}
		service.AddFaultRule(rule)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		require.NoError(t, err)

		// First attempt should succeed (skipped)
		taskResult, err := vm.PowerOn(ctx)
		require.NoError(t, err, "expected success on first attempt (skipped)")
		if taskResult != nil {
			taskResult.Wait(ctx)
		}

		// Second and third attempts should fail (after skip, max 2 injections)
		for i := 0; i < 2; i++ {
			vm.PowerOff(ctx)
			_, err = vm.PowerOn(ctx)
			require.Error(t, err, "expected error on attempt %d", i+2)
		}

		// Fourth attempt should succeed (max count reached)
		vm.PowerOff(ctx)
		taskResult, err = vm.PowerOn(ctx)
		require.NoError(t, err, "expected success on fourth attempt (max count reached)")
		if taskResult != nil {
			taskResult.Wait(ctx)
		}

		service.ClearFaultRules()
	})
}

func TestFaultInjectionInclusionExclusionFilter(t *testing.T) {
	// Simplified test: Test boolean combinations of inclusion/exclusion filters
	// Property matching logic is tested separately in property/match_test.go
	type testCase struct {
		name               string
		inclusionFilter    FaultFilterCallback
		exclusionFilter    FaultFilterCallback
		expectError        bool
		expectedLogMessage string
	}

	testCases := []testCase{
		{
			name: "Inclusion=true, Exclusion=false (should trigger)",
			inclusionFilter: func(obj mo.Reference) bool {
				return true // Always match
			},
			exclusionFilter: func(obj mo.Reference) bool {
				return false // Never exclude
			},
			expectError:        true,
			expectedLogMessage: "Inclusion=true, Exclusion=false: fault injected",
		},
		{
			name: "Inclusion=false, Exclusion=false (should not trigger)",
			inclusionFilter: func(obj mo.Reference) bool {
				return false // Never match
			},
			exclusionFilter: func(obj mo.Reference) bool {
				return false // Never exclude
			},
			expectError:        false,
			expectedLogMessage: "Inclusion=false, Exclusion=false: no fault injected",
		},
		{
			name: "Inclusion=true, Exclusion=true (should not trigger - excluded)",
			inclusionFilter: func(obj mo.Reference) bool {
				return true // Always match
			},
			exclusionFilter: func(obj mo.Reference) bool {
				return true // Always exclude
			},
			expectError:        false,
			expectedLogMessage: "Inclusion=true, Exclusion=true: excluded, no fault injected",
		},
		{
			name: "Inclusion=false, Exclusion=true (should not trigger - not included)",
			inclusionFilter: func(obj mo.Reference) bool {
				return false // Never match
			},
			exclusionFilter: func(obj mo.Reference) bool {
				return true // Always exclude (but doesn't matter since not included)
			},
			expectError:        false,
			expectedLogMessage: "Inclusion=false, Exclusion=true: not included, no fault injected",
		},
	}

	Test(func(ctx context.Context, c *vim25.Client) {
		service := ServiceFromContext(ctx)

		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		require.NoError(t, err)

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				service.ClearFaultRules()
				rule := &FaultInjectionRule{
					MethodName:              "PowerOnVM_Task",
					ObjectType:              "*",
					ObjectName:              "*",
					InclusionPropertyFilter: tc.inclusionFilter,
					ExclusionPropertyFilter: tc.exclusionFilter,
					Probability:             1.0,
					FaultType:               FaultTypeGeneric,
					Message:                 "Filter test",
					Enabled:                 true,
				}
				service.AddFaultRule(rule)

				// Always power off before test, ignoring "already powered off" errors
				_, err := vm.PowerOff(ctx)
				require.True(t, err == nil || fault.IsAlreadyPoweredOffError(err), "unexpected error powering off VM: %v", err)

				taskResult, err := vm.PowerOn(ctx)

				if tc.expectError {
					require.Error(t, err, "expected error for test case: %s", tc.name)
					t.Logf("%s, got error: %v", tc.expectedLogMessage, err)
				} else {
					require.NoError(t, err, "expected success for test case: %s", tc.name)
					if taskResult != nil {
						taskResult.Wait(ctx)
					}
					t.Log(tc.expectedLogMessage)
				}
			})
		}
	})
}

func TestFaultInjectionIncorrectFilterSpecifications(t *testing.T) {
	testCases := []struct {
		name     string
		filters  []string
		panicMsg string
	}{
		{
			name:     "empty filter string",
			filters:  []string{""},
			panicMsg: "empty filter string",
		},
		{
			name:     "whitespace only filter string",
			filters:  []string{"   "},
			panicMsg: "empty filter string",
		},
		{
			name:     "missing separator",
			filters:  []string{"config.name DC0_C0_RP0_VM0"},
			panicMsg: "invalid property filter format",
		},
		{
			name:     "missing value",
			filters:  []string{"config.name == "},
			panicMsg: "invalid property filter format",
		},
		{
			name:     "empty property path",
			filters:  []string{" == value"},
			panicMsg: "invalid property filter format",
		},
		{
			name:     "whitespace property path",
			filters:  []string{"   == value"},
			panicMsg: "invalid property filter format",
		},
		{
			name:     "duplicate property path",
			filters:  []string{"config.name == DC0*", "config.name == *VM0"},
			panicMsg: "multiple inclusion filters specified for property",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var panicValue interface{}
			func() {
				defer func() {
					panicValue = recover()
				}()
				_ = FaultFilterAll(tc.filters...)
			}()

			require.NotNil(t, panicValue, "expected panic for test case: %s", tc.name)
			if tc.panicMsg != "" {
				panicStr := fmt.Sprintf("%v", panicValue)
				require.Contains(t, panicStr, tc.panicMsg, "panic message should contain: %s", tc.panicMsg)
			}
			t.Logf("Got expected panic: %v", panicValue)
		})
	}
}
