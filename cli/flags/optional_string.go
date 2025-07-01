// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
)

type optionalString struct {
	val **string
}

func (s *optionalString) Set(input string) error {
	*s.val = &input
	return nil
}

func (s *optionalString) Get() any {
	if *s.val == nil {
		return nil
	}
	return **s.val
}

func (s *optionalString) String() string {
	if s.val == nil || *s.val == nil {
		return "<nil>"
	}
	return **s.val
}

// NewOptionalString returns a flag.Value implementation where there is no default value.
// This avoids sending a default value over the wire as using flag.StringVar() would.
func NewOptionalString(v **string) flag.Value {
	return &optionalString{v}
}
