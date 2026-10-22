// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package simulator

import "errors"

// buildToolboxArtifact is a no-op stub on non-Linux platforms.
// The toolbox binary build is only supported on Linux, where container-backed
// VM simulation runs.
func buildToolboxArtifact() (toolboxBinPath string, err error) {
	return "", errors.New("vmci-artifacts: only available on linux")
}

// VmciToolboxBinaryPath returns "" on non-Linux platforms where VMCI simulation
// is not supported.
func VmciToolboxBinaryPath() string { return "" }
