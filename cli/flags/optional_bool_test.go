// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
	"testing"
)

func TestOptionalBool(t *testing.T) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	var val *bool

	fs.Var(NewOptionalBool(&val), "obool", "optional bool")

	b := fs.Lookup("obool")

	if b.DefValue != "<nil>" {
		t.Fail()
	}

	if b.Value.String() != "<nil>" {
		t.Fail()
	}

	if b.Value.(flag.Getter).Get() != nil {
		t.Fail()
	}

	b.Value.Set("true")

	if b.Value.String() != "true" {
		t.Fail()
	}

	if b.Value.(flag.Getter).Get() != true {
		t.Fail()
	}

	if val == nil || *val != true {
		t.Fail()
	}

	b.Value.Set("false")

	if b.Value.String() != "false" {
		t.Fail()
	}

	if b.Value.(flag.Getter).Get() != false {
		t.Fail()
	}

	if val == nil || *val != false {
		t.Fail()
	}
}
