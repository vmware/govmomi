// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importer

import (
	"fmt"
	"path"
)

type importable struct {
	localPath  string
	remotePath string
}

func (i importable) Ext() string {
	return path.Ext(i.localPath)
}

func (i importable) Base() string {
	return path.Base(i.localPath)
}

func (i importable) BaseClean() string {
	b := i.Base()
	e := i.Ext()
	return b[:len(b)-len(e)]
}

func (i importable) RemoteSrcVMDK() string {
	file := fmt.Sprintf("%s-src.vmdk", i.BaseClean())
	return i.toRemotePath(file)
}

func (i importable) RemoteDstVMDK() string {
	file := fmt.Sprintf("%s.vmdk", i.BaseClean())
	return i.toRemotePath(file)
}

func (i importable) toRemotePath(p string) string {
	if i.remotePath == "" {
		return p
	}

	return path.Join(i.remotePath, p)
}
