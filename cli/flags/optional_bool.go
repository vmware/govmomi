// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
	"fmt"
	"strconv"
)

type optionalBool struct {
	val **bool
}

func (b *optionalBool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b.val = &v
	return err
}

func (b *optionalBool) Get() any {
	if *b.val == nil {
		return nil
	}
	return **b.val
}

func (b *optionalBool) String() string {
	if b.val == nil || *b.val == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", **b.val)
}

func (b *optionalBool) IsBoolFlag() bool { return true }

// NewOptionalBool returns a flag.Value implementation where there is no default value.
// This avoids sending a default value over the wire as using flag.BoolVar() would.
func NewOptionalBool(v **bool) flag.Value {
	return &optionalBool{v}
}
