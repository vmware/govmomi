// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
	"fmt"
	"strconv"
)

// This flag type is internal to stdlib:
// https://github.com/golang/go/blob/master/src/cmd/internal/obj/flag.go
type int32Value int32

func (i *int32Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 32)
	*i = int32Value(v)
	return err
}

func (i *int32Value) Get() any {
	return int32(*i)
}

func (i *int32Value) String() string {
	return fmt.Sprintf("%v", *i)
}

// NewInt32 behaves as flag.IntVar, but using an int32 type.
func NewInt32(v *int32) flag.Value {
	return (*int32Value)(v)
}

type int32ptrValue struct {
	val **int32
}

func (i *int32ptrValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 32)
	*i.val = new(int32)
	**i.val = int32(v)
	return err
}

func (i *int32ptrValue) Get() any {
	if i.val == nil || *i.val == nil {
		return nil
	}
	return *i.val
}

func (i *int32ptrValue) String() string {
	return fmt.Sprintf("%v", i.Get())
}

func NewOptionalInt32(v **int32) flag.Value {
	return &int32ptrValue{val: v}
}
