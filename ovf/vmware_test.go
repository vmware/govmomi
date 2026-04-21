// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnmarshalVmwareFixtures(t *testing.T) {
	dir := "fixtures/vmware"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Skipf("fixtures directory not found: %s", dir)
		return
	}

	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".ovf" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", dir, err)
	}

	for _, path := range files {
		t.Run(path, func(t *testing.T) {
			runFixtureTest(t, path, isNegativeFixture(path))
		})
	}
}
