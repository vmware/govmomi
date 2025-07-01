// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importer

import (
	"runtime"
	"testing"
)

func TestImporter_manifestPath(t *testing.T) {
	// We can only test filepath operations on the target OS
	manifestTests := []struct {
		goos     string
		name     string
		path     string
		expected string
	}{
		{
			goos:     "linux",
			name:     "linux path",
			path:     "/home/user/foo/bar/qux.ovf",
			expected: "/home/user/foo/bar/qux.mf",
		},
		{
			goos:     "darwin",
			name:     "darwin path",
			path:     "/home/user/foo/bar/qux.ovf",
			expected: "/home/user/foo/bar/qux.mf",
		},
		{
			goos:     "windows",
			name:     "windows path",
			path:     "C:\\ProgramData\\Foo\\Bar\\Qux.ovf",
			expected: "C:\\ProgramData\\Foo\\Bar\\Qux.mf",
		},
	}

	imp := Importer{}

	for _, test := range manifestTests {
		if test.goos == runtime.GOOS {
			manifestPath := imp.manifestPath(test.path)
			if manifestPath != test.expected {
				t.Fatalf("'%s' failed: expected '%s', got '%s'", test.name, test.expected, manifestPath)
			}
		}
	}
}
