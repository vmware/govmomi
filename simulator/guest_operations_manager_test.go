/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func TestFileInfo(t *testing.T) {
	switch runtime.GOOS {
	case "linux", "darwin":
	default:
		// listFiles() returns a `find` command to run inside a linux docker container.
		// The `find` command also works on darwin, skip otherwise.
		t.Skipf("GOOS=%s", runtime.GOOS)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	req := &types.ListFilesInGuest{
		FilePath:     pwd,
		MatchPattern: "*_test.go",
	}

	args := listFiles(req)
	path, err := exec.LookPath(args[0])
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Cmd{Path: path, Args: args}
	res, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	for _, info := range toFileInfo(string(res)) {
		if info.Path == "" {
			t.Fail()
		}
		if !strings.HasSuffix(info.Path, "_test.go") {
			t.Fail()
		}
		if info.Type == "" {
			t.Fail()
		}
		if info.Size == 0 {
			t.Fail()
		}
		attr, ok := info.Attributes.(*types.GuestPosixFileAttributes)
		if !ok {
			t.Fail()
		}
		if attr.ModificationTime == nil {
			t.Fail()
		}
		if attr.AccessTime == nil {
			t.Fail()
		}
		if attr.Permissions == 0 {
			t.Fail()
		}
	}
}
