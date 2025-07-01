//go:build !linux
// +build !linux

// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package hgfs

import (
	"os"
)

func (a *AttrV2) sysStat(info os.FileInfo) {
}
