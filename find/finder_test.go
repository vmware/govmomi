/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package find

import (
	"reflect"
	"testing"
)

func TestInsertFolder(t *testing.T) {
	tests := []struct {
		In   string
		Fldr string
		Out  string
	}{
		{
			In:   "*",
			Fldr: "fldr",
			Out:  "./fldr/*",
		},
		{
			In:   "/*",
			Fldr: "fldr",
			Out:  "fldr/*",
		},
		{
			In:   "/foo/*",
			Fldr: "fldr",
			Out:  "foo/fldr/*",
		},
		{
			In:   "/foo/../*",
			Fldr: "fldr",
			Out:  "fldr/*",
		},
		{
			In:   "/./foo/*",
			Fldr: "fldr",
			Out:  "foo/fldr/*",
		},
		{
			In:   "/../foo/*",
			Fldr: "fldr",
			Out:  "foo/fldr/*",
		},
		{
			In:   "/foo/bar/*",
			Fldr: "fldr",
			Out:  "foo/bar/fldr/*",
		},
		{
			In:   "/foo/bar/../*",
			Fldr: "fldr",
			Out:  "foo/fldr/*",
		},
		{
			In:   "./*",
			Fldr: "fldr",
			Out:  "./fldr/*",
		},
		{
			In:   "foo/*",
			Fldr: "fldr",
			Out:  "./foo/fldr/*",
		},
		{
			In:   "foo/../*",
			Fldr: "fldr",
			Out:  "./fldr/*",
		},
		{
			In:   "./foo/*",
			Fldr: "fldr",
			Out:  "./foo/fldr/*",
		},
		{
			In:   "../foo/*", // Special case...
			Fldr: "fldr",
			Out:  "../foo/fldr/*",
		},
		{
			In:   "foo/bar/../*",
			Fldr: "fldr",
			Out:  "./foo/fldr/*",
		},
	}

	for _, test := range tests {
		out := insertFolder(test.In, test.Fldr)
		if !reflect.DeepEqual(test.Out, out) {
			t.Errorf("Expected %s to return: %#v, actual: %#v", test.In, test.Out, out)
		}
	}
}
