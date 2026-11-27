// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package simulator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	vmciArtifactsOnce        sync.Once
	vmciArtifactsToolboxPath string
	vmciArtifactsErr         error
)

// artifactsBinDir returns the stable host directory used to hold the built
// toolbox binary across test invocations.
//
// The directory is simulator/bin/ relative to this source file. "go build -o"
// is invoked on every call (see buildToolboxArtifact); the directory only
// exists to give that output a stable, discoverable path, not to skip the
// build. Go's own content-addressed build cache (~/.cache/go-build) is what
// makes repeated invocations cheap when toolbox sources are unchanged — this
// directory deliberately carries no invalidation logic of its own, so a
// presence check can never serve a binary built from stale sources.
//
// Falls back to os.MkdirTemp when the source path is unavailable (e.g. binaries
// built with -trimpath or deployed outside the source tree).
func artifactsBinDir() (string, error) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return os.MkdirTemp("", "vcsim-vmci-artifacts-*")
	}
	dir := filepath.Join(filepath.Dir(thisFile), "bin")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("vmci-artifacts: mkdir %s: %w", dir, err)
	}
	return dir, nil
}

// buildToolboxArtifact returns the path to the govmomi/toolbox binary
// (injected as /usr/bin/vmware-rpctool in container-backed VMs with
// RUN.vmci=true), building it only if it is not already present in the
// cache directory.
//
// The function is idempotent and thread-safe (sync.Once per process).  It is
// called automatically inside simVM.start() when RUN.vmci=true.
//
// A build failure here is non-fatal to the caller: it is logged and the
// returned path is empty, so /usr/bin/vmware-rpctool auto-injection is
// skipped, but the container still starts. Callers that need to guarantee
// the binary is present should check VmciToolboxBinaryPath() for emptiness.
func buildToolboxArtifact() (toolboxBinPath string, err error) {
	vmciArtifactsOnce.Do(func() {
		dir, mkErr := artifactsBinDir()
		if mkErr != nil {
			vmciArtifactsErr = mkErr
			return
		}

		toolboxBin := filepath.Join(dir, "toolbox")
		// Always build: a presence check here previously let an ancient binary
		// (built before a toolbox wire-protocol change) go on serving requests
		// indefinitely, since nothing about this directory's content changes
		// when the sources it was built from do. "go build -o" always
		// overwrites toolboxBin with a binary matching the current sources;
		// Go's own build cache (keyed on source content, not a stat) is what
		// keeps the repeat-invocation cost low.
		toolboxBuild := exec.Command(
			"go", "build",
			"-o", toolboxBin,
			"github.com/vmware/govmomi/toolbox/toolbox",
		)
		toolboxBuild.Env = append(os.Environ(),
			"CGO_ENABLED=0",
			"GOOS=linux",
			"GOARCH=amd64",
		)
		if out, buildErr := toolboxBuild.CombinedOutput(); buildErr != nil {
			vmciArtifactsErr = fmt.Errorf("toolbox build: %w\n%s", buildErr, out)
			return
		}
		log.Printf("vmci-artifacts: built toolbox at %s", toolboxBin)
		vmciArtifactsToolboxPath = toolboxBin
	})
	return vmciArtifactsToolboxPath, vmciArtifactsErr
}

// VmciToolboxBinaryPath returns the host path of the govmomi/toolbox static
// binary built by buildToolboxArtifact.  It is auto-injected as
// /usr/bin/vmware-rpctool in every container-backed VM with RUN.vmci=true.
//
// Callers that also want to replace /usr/bin/vmtoolsd (e.g. integration tests
// whose init process needs to write guestinfo via the toolbox binary) can
// obtain the path here and add a RUN.volume.<label> ExtraConfig entry:
//
//	"RUN.volume.vmtoolsd": simulator.VmciToolboxBinaryPath() + ":/usr/bin/vmtoolsd:ro"
//
// Returns "" if the build failed (RUN.vmci containers still work via the
// auto-injected /usr/bin/vmware-rpctool, but /usr/bin/vmtoolsd callers will
// fall back to whatever is in the container image).
func VmciToolboxBinaryPath() string {
	path, _ := buildToolboxArtifact()
	return path
}
