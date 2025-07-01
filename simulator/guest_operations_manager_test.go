// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
