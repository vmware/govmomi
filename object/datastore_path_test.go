// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import "testing"

func TestParseDatastorePath(t *testing.T) {
	tests := []struct {
		dsPath string
		dsFile string
		fail   bool
	}{
		{"", "", true},
		{"x", "", true},
		{"[", "", true},
		{"[nope", "", true},
		{"[te st]", "", false},
		{"[te st] foo", "foo", false},
		{"[te st] foo/foo.vmx", "foo/foo.vmx", false},
		{"[te st]foo bar/foo bar.vmx", "foo bar/foo bar.vmx", false},
		{" [te st]     bar/bar.vmx  ", "bar/bar.vmx", false},
	}

	for _, test := range tests {
		p := new(DatastorePath)
		ok := p.FromString(test.dsPath)

		if test.fail {
			if ok {
				t.Errorf("expected error for: %s", test.dsPath)
			}
		} else {
			if !ok {
				t.Errorf("failed to parse: %q", test.dsPath)
			} else {
				if test.dsFile != p.Path {
					t.Errorf("dsFile=%s", p.Path)
				}
				if p.Datastore != "te st" {
					t.Errorf("ds=%s", p.Datastore)
				}
			}
		}
	}

	s := "[datastore1] foo/bar.vmdk"
	p := new(DatastorePath)
	ok := p.FromString(s)
	if !ok {
		t.Fatal(s)
	}

	if p.String() != s {
		t.Fatal(p.String())
	}

	p.Path = ""

	if p.String() != "[datastore1]" {
		t.Fatal(p.String())
	}
}
