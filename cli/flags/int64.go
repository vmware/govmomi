// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
	"fmt"
	"strconv"
)

// This flag type is internal to stdlib:
// https://github.com/golang/go/blob/master/src/cmd/internal/obj/flag.go
type int64Value int64

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) Get() any {
	return int64(*i)
}

func (i *int64Value) String() string {
	return fmt.Sprintf("%v", *i)
}

// NewInt64 behaves as flag.IntVar, but using an int64 type.
func NewInt64(v *int64) flag.Value {
	return (*int64Value)(v)
}

type int64ptrValue struct {
	val **int64
}

func (i *int64ptrValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i.val = new(int64)
	**i.val = int64(v)
	return err
}

func (i *int64ptrValue) Get() any {
	if i.val == nil || *i.val == nil {
		return nil
	}
	return **i.val
}

func (i *int64ptrValue) String() string {
	return fmt.Sprintf("%v", i.Get())
}

func NewOptionalInt64(v **int64) flag.Value {
	return &int64ptrValue{val: v}
}
