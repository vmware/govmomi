// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"flag"
	"testing"
)

func TestOptionalString(t *testing.T) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	var val *string

	fs.Var(NewOptionalString(&val), "ostring", "optional string")

	s := fs.Lookup("ostring")

	if s.DefValue != "<nil>" {
		t.Fail()
	}

	if s.Value.String() != "<nil>" {
		t.Fail()
	}

	if s.Value.(flag.Getter).Get() != nil {
		t.Fail()
	}

	s.Value.Set("test")

	if s.Value.String() != "test" {
		t.Fail()
	}

	if s.Value.(flag.Getter).Get() != "test" {
		t.Fail()
	}

	if val == nil || *val != "test" {
		t.Fail()
	}
}
