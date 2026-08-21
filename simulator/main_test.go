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
	// Remove leftover vcsim GuestRPC socket directories from crashed test runs.
	//
	// The socket lives at <TempDir>/vcsim-rpc-<uid>/rpc.sock, so the pattern
	// must match the DIRECTORY: filepath.Match does not let * cross a path
	// separator, so a "vcsim-rpc-*.sock" pattern matched nothing and the
	// directories accumulated under TempDir.
	for _, pattern := range []string{
		filepath.Join(os.TempDir(), "vcsim-rpc-*"),
	} {
		if matches, err := filepath.Glob(pattern); err == nil {
			for _, f := range matches {
				_ = os.RemoveAll(f)
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
