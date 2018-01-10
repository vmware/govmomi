package flags

import (
	"flag"
	"fmt"
	"strings"
)

// enum is a value store that validates against a set of strings. The validation is performed at set-time.
type enum struct {
	value   *string
	allowed []string
}

// NewEnum returns a enum as a flag.Value for the string pointed at by p, and the validations supplied in allowed.
func NewEnum(p *string, allowed []string) flag.Value {
	return &enum{value: p, allowed: allowed}
}

// Get implements Getter for stringSliceValue.
func (e *enum) Get() interface{} {
	return *e.value
}

// Set implements Value for stringSliceValue.
func (e *enum) Set(s string) error {
	if err := validateEnum(s, e.allowed); err != nil {
		return err
	}
	*e.value = s
	return nil
}

// String implements Value for enum.
func (e *enum) String() string {
	var s string
	if e.value != nil {
		s = fmt.Sprintf("%s", *e.value)
	}
	return s
}

// enumSlice is a value store that validates against a set of strings. The validation is performed at set-time.
type enumSlice struct {
	value   *[]string
	allowed []string
}

// NewEnumSlice returns a enumSlice as a flag.Value for the
// string pointed at by p, validated by the values in allowed.
func NewEnumSlice(p *[]string, allowed []string) flag.Value {
	return &enumSlice{value: p, allowed: allowed}
}

// Get implements Getter for stringSliceValue.
func (e *enumSlice) Get() interface{} {
	return *e.value
}

// Set implements Value for stringSliceValue.
func (e *enumSlice) Set(s string) error {
	vs := strings.Split(s, ",")
	for _, v := range vs {
		if err := validateEnum(v, e.allowed); err != nil {
			return err
		}
	}
	*e.value = vs
	return nil
}

// String implements Value for enumSlice.
func (e *enumSlice) String() string {
	var s string
	if e.value != nil && *e.value != nil {
		s = fmt.Sprintf("%s", *e.value)
	}
	return s
}

func validateEnum(s string, allowed []string) error {
	for _, a := range allowed {
		if a == s {
			return nil
		}
	}
	return fmt.Errorf("valid values are %s", strings.Join(allowed, ","))
}
