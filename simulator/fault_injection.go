// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// FaultFilterCallback evaluates whether a managed object is a "match", allowing arbitrarily complex
// filters by implementing this function directly.
// For convenience, use FaultFilterAll() or FaultFilterAny() to create filters from string expressions.
type FaultFilterCallback func(obj mo.Reference) bool

// FaultFilterAll (AND) takes a set of strings expressing propertyPath:value pairs using the property.Match format.
// Only equality matches are currently supported: "property.path == value"
// Examples:
//   - FaultFilterAll("Config.Name == DC0_C0_RP0_VM0")
//   - FaultFilterAll("Config.Name == DC0_C0*", "Runtime.PowerState == poweredOn")
//   - FaultFilterAll("Config.Name == *") // match all objects where the property exists
//
// Panics on code-level usage errors:
//   - Empty filter strings
//   - Invalid filter format (missing " == " separator)
//   - Empty property path
//   - Multiple filters targeting the same property path ( combine them into a single pattern)
func FaultFilterAll(filterStrs ...string) FaultFilterCallback {
	match := parsePropertyFilters(filterStrs)

	propertyPaths := make([]string, 0, len(match))
	for path := range match {
		propertyPaths = append(propertyPaths, path)
	}

	// Return a closure that captures the match and paths
	return func(obj mo.Reference) bool {
		if obj == nil {
			return false
		}

		props := getProperties(obj, propertyPaths)
		// AND match
		return match.List(props)
	}
}

// FaultFilterAny (OR) takes a set of strings expressing propertyPath:value pairs using the property.Match format.
// See FaultFilterAll for details.
func FaultFilterAny(filterStrs ...string) FaultFilterCallback {
	// For exclusion matches, we need to handle each filter separately
	// since multiple matches can target the same property path with different values
	matches := make([]property.Match, 0, len(filterStrs))
	allPaths := make(map[string]bool)
	for _, filterStr := range filterStrs {
		match := parsePropertyFilters([]string{filterStr})
		if len(match) > 0 {
			matches = append(matches, match)

			for path := range match {
				// deduplicate paths for collection
				allPaths[path] = true
			}
		}
	}

	paths := make([]string, 0, len(allPaths))
	for path := range allPaths {
		paths = append(paths, path)
	}

	// Return a closure that captures the matches and paths
	return func(obj mo.Reference) bool {
		if obj == nil {
			return false
		}

		props := getProperties(obj, paths)

		// OR match
		for _, match := range matches {
			if match.AnyList(props) {
				return true
			}
		}
		return false
	}
}

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

	// InclusionPropertyFilter is a callback function that evaluates whether the managed object matches the filter.
	// ObjectName above is a convenience as it's such a common filter, but can be expressed here.
	// Use FaultFilterAll (AND) or FaultFilterAny (OR) helper functions to create from string expressions for common cases.
	// Examples:
	//   - InclusionPropertyFilter: FaultFilterAll("Config.Name == DC0_C0_RP0_VM0")
	//   - InclusionPropertyFilter: FaultFilterAll("Config.Name == DC0_C0*", "Runtime.PowerState == poweredOn")
	//   - InclusionPropertyFilter: FaultFilterAll("Config.Name == *") // match all objects where the property exists
	//   - InclusionPropertyFilter: func(obj mo.Reference) bool { ... } // custom filter logic
	// Note: To express "property != value", use an exclusion filter: FaultFilterAny("property == value")
	// The order of evaluation is:
	// 1. ObjectName
	// 2. InclusionPropertyFilter (all must match)
	// 3. ExclusionPropertyFilter (none must match)
	InclusionPropertyFilter FaultFilterCallback

	// ExclusionPropertyFilter is a callback function that evaluates whether the managed object should be excluded.
	// This is evaluated after the name and inclusion filters.
	// If the function returns true, the object is excluded.
	// Use FaultFilterAll (AND) or FaultFilterAny (OR) helper functions to create from string expressions for common cases.
	// Examples:
	//   - ExclusionPropertyFilter: FaultFilterAny("Config.Name == test-vm") to exclude objects with name "test-vm"
	//   - ExclusionPropertyFilter: FaultFilterAny("Config.Name == ", "Runtime.PowerState == ") to exclude objects where Name or PowerState are empty
	//   - ExclusionPropertyFilter: func(obj mo.Reference) bool { ... } // custom exclusion logic
	ExclusionPropertyFilter FaultFilterCallback

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

		// Check inclusion property filter (order 2)
		if rule.InclusionPropertyFilter != nil {
			if !rule.InclusionPropertyFilter(handler) {
				continue
			}
		}

		// Check exclusion property filter (order 3)
		if rule.ExclusionPropertyFilter != nil {
			if rule.ExclusionPropertyFilter(handler) {
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

// parsePropertyFilters parses an array of string filter expressions and consolidates them
// into a single property.Match map. Only equality matches are supported: "property.path == value"
// Panics on code-level usage errors:
//   - Empty filter strings
//   - Invalid filter format (missing " == " separator)
//   - Empty property path or expected value
//   - Multiple filters targeting the same property path (user should combine them into a single pattern)
func parsePropertyFilters(filterStrs []string) property.Match {
	match := make(property.Match)

	for _, filterStr := range filterStrs {
		filterStr = strings.TrimSpace(filterStr)
		if filterStr == "" {
			panic("empty filter string in property filter - remove empty entries from filter array")
		}

		parts := strings.SplitN(filterStr, " == ", 2)
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid property filter format %q - expected format: \"property.path == value\"", filterStr))
		}

		propertyPath := strings.TrimSpace(parts[0])
		expectedValue := strings.TrimSpace(parts[1])
		if propertyPath == "" {
			panic(fmt.Sprintf("empty property path in filter %q - expected format: \"property.path == value\"", filterStr))
		}
		// nil expected values are ok - that's how we expect to match empty properties

		// Convert property path to camelCase (property collector format)
		// e.g., "Config.Name" -> "config.name"
		pathParts := strings.Split(propertyPath, ".")
		for i, part := range pathParts {
			pathParts[i] = lcFirst(part)
		}
		propertyPath = strings.Join(pathParts, ".")

		// Panic if this property path already exists (multiple filters for same property)
		if _, exists := match[propertyPath]; exists {
			panic(fmt.Sprintf("multiple inclusion filters specified for property %q - combine them into a single pattern", propertyPath))
		}

		match[propertyPath] = expectedValue
	}

	return match
}

// getProperties collects the specified properties from a managed object.
func getProperties(obj mo.Reference, paths []string) []types.DynamicProperty {
	if obj == nil || len(paths) == 0 {
		return nil
	}

	rval := getManagedObject(obj)
	props := make([]types.DynamicProperty, 0, len(paths))

	for _, propPath := range paths {
		propValue, err := fieldValue(rval, propPath)
		var val types.AnyType
		if err == nil {
			val = propValue
		}
		// Note: If property doesn't exist or is empty, val will be nil.
		// property.Match.Property() will return false for nil values.

		props = append(props, types.DynamicProperty{
			Name: propPath,
			Val:  val,
		})
	}

	return props
}
