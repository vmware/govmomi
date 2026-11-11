// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/types"
)

func TestErrorOutputTaskError(t *testing.T) {
	tests := []struct {
		name string
		err  task.Error
		want []string
	}{
		{
			name: "nil fault does not panic",
			err:  task.Error{LocalizedMethodFault: &types.LocalizedMethodFault{LocalizedMessage: "boom"}},
			want: []string{"boom"},
		},
		{
			name: "fault messages are printed once",
			err: task.Error{LocalizedMethodFault: &types.LocalizedMethodFault{
				LocalizedMessage: "boom",
				Fault: &types.InvalidArgument{RuntimeFault: types.RuntimeFault{MethodFault: types.MethodFault{
					FaultMessage: []types.LocalizableMessage{{Key: "k", Message: "detail"}},
				}}},
			}},
			want: []string{"detail (k)", "boom"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := (errorOutput{test.err}).Write(&buf); err != nil {
				t.Fatal(err)
			}

			var got []string
			for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
				// strip the "<argv0>: " prefix
				got = append(got, line[strings.Index(line, ": ")+2:])
			}

			if strings.Join(got, "|") != strings.Join(test.want, "|") {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
