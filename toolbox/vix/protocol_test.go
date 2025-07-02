// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vix

import (
	"reflect"
	"testing"
)

func TestMarshalVixMsgStartProgramRequest(t *testing.T) {
	requests := []*StartProgramRequest{
		{},
		{
			ProgramPath: "/bin/date",
		},
		{
			ProgramPath: "/bin/date",
			Arguments:   "--date=@2147483647",
		},
		{
			ProgramPath: "/bin/date",
			WorkingDir:  "/tmp",
		},
		{
			ProgramPath: "/bin/date",
			WorkingDir:  "/tmp",
			EnvVars:     []string{"FOO=bar"},
		},
		{
			ProgramPath: "/bin/date",
			WorkingDir:  "/tmp",
			EnvVars:     []string{"FOO=bar", "BAR=foo"},
		},
	}

	for i, in := range requests {
		buf, err := in.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}

		out := new(StartProgramRequest)

		err = out.UnmarshalBinary(buf)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("%d marshal mismatch", i)
		}
	}
}
