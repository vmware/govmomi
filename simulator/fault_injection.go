// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// FaultInjectionRule defines a rule for injecting faults into method responses.
type FaultInjectionRule struct {
	// MethodName is the SOAP method name to inject faults into
	// (e.g., "PowerOnVM_Task").
	// Use "*" to match all methods
	MethodName string

	// ObjectType is the managed object type to inject faults into
	// (e.g., "VirtualMachine").
	// Use "*" to match all object types
	ObjectType string

	// ObjectName is the name of the specific object to inject faults into
	// Use "*" to match all object names
	ObjectName string

	// InclusionPropertyFilter is an array of property filters to match the object.
	// ObjectName above is a convenience as it's such a common filter, but
	// could be expressed here. All filters in the array must match (AND).
	// Supported operators:
	//   - "==" for equality: "Config.Name == 'DC0_C0_RP0_VM0'"
	//   - "!=" for inequality: "Config.Name != 'other'"
	//   - "~=" for regexp match: "Config.Name ~= '^DC0.*'"
	// Examples:
	//   - []string{"Config.Name == 'DC0_C0_RP0_VM0'"}
	//   - []string{"Config.Name ~= '^DC0.*'", "Runtime.PowerState == 'poweredOn'"}
	// The order of evaluation is:
	// 1. ObjectName
	// 2. InclusionPropertyFilter (all must match)
	// 3. ExclusionPropertyFilter (none must match)
	InclusionPropertyFilter []string

	// ExclusionPropertyFilter is an array of property filters to exclude the object.
	// This is evaluated after the name and inclusion filters.
	// If any filter in the array matches, the object is excluded (OR).
	// Supports the same operators as InclusionPropertyFilter:
	//   - "==" for equality
	//   - "!=" for inequality
	//   - "~=" for regexp match
	// Example: []string{"Config.Name ~= '^test-.*'"} to exclude all objects with names starting with "test-"
	ExclusionPropertyFilter []string

	// Probability is the probability (0.0 to 1.0) that a fault will be
	// injected.
	Probability float64

	// FaultType is the type of fault to inject
	FaultType FaultType

	// Message is the error message for the fault
	Message string

	// Fault is the specific VMware fault to inject (optional, overrides
	// FaultType).
	Fault types.BaseMethodFault

	// Delay introduces a delay before returning the fault (in milliseconds)
	Delay int

	// MaxCount limits the number of times this rule can be triggered
	// (0 = unlimited)
	MaxCount int

	// SkipCount is the number of times to skip this rule before it is triggered
	SkipCount int

	// Enabled controls whether this rule is active
	Enabled bool

	// Internal fields
	count     int64
	skipCount int64
	rng       *rand.Rand
}

// FaultType represents the type of fault to inject
type FaultType int

const (
	// FaultTypeGeneric injects a generic fault
	FaultTypeGeneric FaultType = iota
	// FaultTypeNotAuthenticated injects authentication errors
	FaultTypeNotAuthenticated
	// FaultTypeNoPermission injects permission errors
	FaultTypeNoPermission
	// FaultTypeInvalidArgument injects invalid argument errors
	FaultTypeInvalidArgument
	// FaultTypeMethodNotFound injects method not found errors
	FaultTypeMethodNotFound
	// FaultTypeManagedObjectNotFound injects object not found errors
	FaultTypeManagedObjectNotFound
	// FaultTypeInvalidState injects invalid state errors
	FaultTypeInvalidState
	// FaultTypeResourceInUse injects resource in use errors
	FaultTypeResourceInUse
	// FaultTypeInsufficientResourcesFault injects insufficient resources errors
	FaultTypeInsufficientResourcesFault
	// FaultTypeCustom allows injection of custom faults specified in the Fault
	// field
	FaultTypeCustom
	// FaultTypeNetworkFailure simulates network-level failures
	FaultTypeNetworkFailure
	// FaultTypeTimeout simulates timeout errors
	FaultTypeTimeout
)

// FaultInjector manages fault injection rules and applies them to method calls
type FaultInjector struct {
	mu    sync.RWMutex
	rules []*FaultInjectionRule
	rng   *rand.Rand
}

// NewFaultInjector creates a new fault injector
func NewFaultInjector() *FaultInjector {
	return &FaultInjector{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// AddRule adds a fault injection rule
func (f *FaultInjector) AddRule(rule *FaultInjectionRule) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Initialize rule's RNG if not set
	if rule.rng == nil {
		rule.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	f.rules = append(f.rules, rule)
}

// RemoveRule removes a fault injection rule by index
func (f *FaultInjector) RemoveRule(index int) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if index < 0 || index >= len(f.rules) {
		return false
	}

	f.rules = append(f.rules[:index], f.rules[index+1:]...)
	return true
}

// ClearRules removes all fault injection rules
func (f *FaultInjector) ClearRules() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.rules = nil
}

// GetRules returns a copy of all fault injection rules
func (f *FaultInjector) GetRules() []*FaultInjectionRule {
	f.mu.RLock()
	defer f.mu.RUnlock()

	rules := make([]*FaultInjectionRule, len(f.rules))
	copy(rules, f.rules)
	return rules
}

// ShouldInjectFault determines if a fault should be injected for the given method call
func (f *FaultInjector) ShouldInjectFault(method *Method, objectName string, handler mo.Reference) *FaultInjectionRule {
	f.mu.RLock()
	defer f.mu.RUnlock()

	for _, rule := range f.rules {
		if !rule.Enabled {
			continue
		}

		// Check if we've exceeded the max count
		if rule.MaxCount > 0 && rule.count >= int64(rule.MaxCount) {
			continue
		}

		// Check method name match
		if rule.MethodName != "*" && rule.MethodName != method.Name {
			continue
		}

		// Check object type match
		if rule.ObjectType != "*" && rule.ObjectType != method.This.Type {
			continue
		}

		// Check object name match (order 1)
		if rule.ObjectName != "*" && rule.ObjectName != objectName {
			continue
		}

		// Check inclusion property filter (order 2) - all must match (AND logic)
		if len(rule.InclusionPropertyFilter) > 0 {
			allMatch := true
			for _, filter := range rule.InclusionPropertyFilter {
				if !evaluatePropertyFilter(handler, filter) {
					allMatch = false
					break
				}
			}
			if !allMatch {
				continue
			}
		}

		// Check exclusion property filter (order 3) - any match excludes (OR logic)
		if len(rule.ExclusionPropertyFilter) > 0 {
			anyMatch := false
			for _, filter := range rule.ExclusionPropertyFilter {
				if evaluatePropertyFilter(handler, filter) {
					anyMatch = true
					break
				}
			}
			if anyMatch {
				continue
			}
		}

		// Check probability
		if rule.rng.Float64() > rule.Probability {
			continue
		}

		// At this point, the rule matches. Check if we need to skip this match
		if rule.SkipCount > 0 && rule.skipCount < int64(rule.SkipCount) {
			rule.skipCount++
			continue
		}

		// This rule matches, increment count and return it
		rule.count++
		return rule
	}

	return nil
}

// CreateFault creates a SOAP fault based on the rule
func (f *FaultInjector) CreateFault(rule *FaultInjectionRule, method *Method) soap.HasFault {
	// Apply delay if specified
	if rule.Delay > 0 {
		time.Sleep(time.Duration(rule.Delay) * time.Millisecond)
	}

	var fault types.BaseMethodFault
	message := rule.Message
	if message == "" {
		message = "Fault injected by simulator"
	}

	// Use custom fault if specified
	if rule.FaultType == FaultTypeCustom && rule.Fault != nil {
		fault = rule.Fault
	} else {
		// Create fault based on type
		switch rule.FaultType {
		case FaultTypeNotAuthenticated:
			fault = &types.NotAuthenticated{
				NoPermission: types.NoPermission{
					Object:      &method.This,
					PrivilegeId: "System.View",
				},
			}
		case FaultTypeNoPermission:
			fault = &types.NoPermission{
				Object:      &method.This,
				PrivilegeId: "System.Modify",
			}
		case FaultTypeInvalidArgument:
			fault = &types.InvalidArgument{
				InvalidProperty: "injected_fault_property",
			}
		case FaultTypeMethodNotFound:
			fault = &types.MethodNotFound{
				Receiver: method.This,
				Method:   method.Name,
			}
		case FaultTypeManagedObjectNotFound:
			fault = &types.ManagedObjectNotFound{
				Obj: method.This,
			}
		case FaultTypeInvalidState:
			fault = &types.InvalidState{
				VimFault: types.VimFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: &types.SystemErrorFault{
								Reason: message,
							},
							LocalizedMessage: message,
						},
					},
				},
			}
		case FaultTypeResourceInUse:
			fault = &types.ResourceInUse{
				VimFault: types.VimFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							LocalizedMessage: message,
						},
					},
				},
				Type: method.This.Type,
				Name: "injected_resource",
			}
		case FaultTypeInsufficientResourcesFault:
			fault = &types.InsufficientResourcesFault{
				VimFault: types.VimFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							LocalizedMessage: message,
						},
					},
				},
			}
		case FaultTypeNetworkFailure:
			fault = &types.SystemErrorFault{
				Reason: "Network connection failed: " + message,
			}
		case FaultTypeTimeout:
			fault = &types.SystemErrorFault{
				Reason: "Operation timed out: " + message,
			}
		default:
			// FaultTypeGeneric
			fault = &types.SystemErrorFault{
				Reason: message,
			}
		}
	}

	return &serverFaultBody{Reason: Fault(message, fault)}
}

// GetStats returns statistics about fault injection
func (f *FaultInjector) GetStats() map[string]interface{} {
	f.mu.RLock()
	defer f.mu.RUnlock()

	stats := make(map[string]interface{})
	stats["total_rules"] = len(f.rules)

	enabledRules := 0
	totalInjections := int64(0)

	for i, rule := range f.rules {
		if rule.Enabled {
			enabledRules++
		}
		totalInjections += rule.count
		stats[fmt.Sprintf("rule_%d_count", i)] = rule.count
	}

	stats["enabled_rules"] = enabledRules
	stats["total_injections"] = totalInjections

	return stats
}

// ResetStats resets the injection counters for all rules
func (f *FaultInjector) ResetStats() {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, rule := range f.rules {
		rule.count = 0
		rule.skipCount = 0
	}
}

// evaluatePropertyFilter evaluates a property filter expression against a managed object.
// See struct for filter expression examples.
func evaluatePropertyFilter(obj mo.Reference, filter string) bool {
	if obj == nil || filter == "" {
		return false
	}

	// Parse the filter expression: "property.path == 'value'", "property.path != 'value'", or "property.path ~= 'regexp'"
	filter = strings.TrimSpace(filter)
	parts := strings.Fields(filter)
	if len(parts) != 3 {
		// If it doesn't match the expected format, try to treat it as a simple property path
		// and check if it exists and is truthy
		return evaluateSimplePropertyFilter(obj, filter)
	}

	propertyPath := parts[0]
	operator := parts[1]
	expectedValue := strings.Trim(parts[2], "'\"")

	// Get the managed object's reflect value
	rval := getManagedObject(obj)

	// Get the property value using fieldValue
	propValue, err := fieldValue(rval, propertyPath)
	if err != nil {
		return false
	}

	// Convert property value to string for comparison
	var propValueStr string
	switch v := propValue.(type) {
	case string:
		propValueStr = v
	case fmt.Stringer:
		propValueStr = v.String()
	default:
		// Try to get string representation
		propValueStr = fmt.Sprintf("%v", v)
	}

	// Evaluate the operator
	switch operator {
	case "==":
		return propValueStr == expectedValue
	case "!=":
		return propValueStr != expectedValue
	case "~=":
		re, err := regexp.Compile(expectedValue)
		if err != nil {
			return false
		}
		return re.MatchString(propValueStr)
	default:
		return false
	}
}

// evaluateSimplePropertyFilter evaluates a simple property path without an operator.
// Returns true if the property exists and has a truthy value.
func evaluateSimplePropertyFilter(obj mo.Reference, propertyPath string) bool {
	if obj == nil || propertyPath == "" {
		return false
	}

	rval := getManagedObject(obj)
	propValue, err := fieldValue(rval, propertyPath)
	if err != nil {
		return false
	}

	// Check if the value is truthy
	if propValue == nil {
		return false
	}

	// For boolean values, return the boolean value
	if b, ok := propValue.(bool); ok {
		return b
	}

	// For strings, return true if non-empty
	if s, ok := propValue.(string); ok {
		return s != ""
	}

	// For other types, return true if not nil
	return true
}
