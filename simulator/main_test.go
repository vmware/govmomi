// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Remove any leftover vcsim GuestRPC sockets from previously crashed test
	// runs.  A fresh Start() would also remove them, but sweeping here catches
	// all patterns before any test even begins.
	for _, pattern := range []string{
		filepath.Join(os.TempDir(), "vcsim-rpc-*.sock"),
	} {
		if matches, err := filepath.Glob(pattern); err == nil {
			for _, f := range matches {
				_ = os.Remove(f)
			}
		}
	}

	// Force-remove any stale vcsim containers left by previously crashed tests.
	// Without this, docker create spends the full grace period (up to 60 s+)
	// waiting for the old container to stop before it can remove it, which
	// easily exhausts the per-test timeout.
	if out, err := exec.Command("docker", "ps", "-aq", "--filter", "name=vcsim").Output(); err == nil {
		for _, id := range strings.Fields(string(out)) {
			_ = exec.Command("docker", "rm", "-f", id).Run()
		}
	}

	os.Exit(m.Run())
}
