package flags

import (
	"flag"
	"fmt"
	"strings"
)

// stringPtr is a two-state string value store that implements the flag
// package interfaces, and is designed for use with flags that can have zero
// values.
type stringPtr struct {
	value **string
}

// NewStringPtr returns a stringPtr as a flag.Value for the string pointer
// supplied by p.
func newStringPtr(p **string) flag.Value {
	return &stringPtr{value: p}
}

// Get implements Getter for stringValue.
func (v *stringPtr) Get() interface{} {
	return *v.value
}

// Set implements Value for stringValue.
func (v *stringPtr) Set(s string) error {
	**v.value = s
	return nil
}

// String implements Value for stringValue.
func (v *stringPtr) String() string {
	return fmt.Sprintf("%v", *v.value)
}

// stringSlice is a value store that implements the flag package
// interfaces, and is designed for use with flags that should have multiple
// values. The values are separated by commas, ie: -foobar="value1,value two".
// A nil slice here is a valid state and indicates the flag was not set.
type stringSlice struct {
	value *[]string
}

// NewStringSlice returns a stringSlice as a flag.Value for the
// string slice pointed at by p.
func NewStringSlice(p *[]string) flag.Value {
	return &stringSlice{value: p}
}

// Get implements Getter for stringSliceValue.
func (v *stringSlice) Get() interface{} {
	return *v.value
}

// Set implements Value for stringSliceValue.
func (v *stringSlice) Set(s string) error {
	*v.value = strings.Split(s, ",")
	return nil
}

// String implements Value for stringSliceValue.
//
// This only serves to satisfy the interface, it should not be used directly as
// it will not distinguish between nil and an empty slice.
func (v *stringSlice) String() string {
	var s string
	if v.value != nil && *v.value != nil {
		s = fmt.Sprintf("%s", *v.value)
	}
	return s
}
