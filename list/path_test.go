// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package list

import (
	"reflect"
	"testing"
)

func TestToParts(t *testing.T) {
	tests := []struct {
		In  string
		Out []string
	}{
		{
			In:  "/",
			Out: []string{},
		},
		{
			In:  "/foo",
			Out: []string{"foo"},
		},
		{
			In:  "/foo/..",
			Out: []string{},
		},
		{
			In:  "/./foo",
			Out: []string{"foo"},
		},
		{
			In:  "/../foo",
			Out: []string{"foo"},
		},
		{
			In:  "/foo/bar",
			Out: []string{"foo", "bar"},
		},
		{
			In:  "/foo/bar/..",
			Out: []string{"foo"},
		},
		{
			In:  "",
			Out: []string{"."},
		},
		{
			In:  ".",
			Out: []string{"."},
		},
		{
			In:  "foo",
			Out: []string{".", "foo"},
		},
		{
			In:  "foo/..",
			Out: []string{"."},
		},
		{
			In:  "./foo",
			Out: []string{".", "foo"},
		},
		{
			In:  "../foo", // Special case...
			Out: []string{"..", "foo"},
		},
		{
			In:  "foo/bar/..",
			Out: []string{".", "foo"},
		},
	}

	for _, test := range tests {
		out := ToParts(test.In)
		if !reflect.DeepEqual(test.Out, out) {
			t.Errorf("Expected %s to return: %#v, actual: %#v", test.In, test.Out, out)
		}
	}
}
