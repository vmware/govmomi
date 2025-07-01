// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"fmt"
	"os"
)

func fileExtendedInfoFormat(dir string, info os.FileInfo) string {
	const format = "<fxi>" +
		"<Name>%s</Name>" +
		"<ft>%d</ft>" +
		"<fs>%d</fs>" +
		"<mt>%d</mt>" +
		"<ct>%d</ct>" +
		"<at>%d</at>" +
		"</fxi>"

	props := 0
	size := info.Size()
	mtime := info.ModTime().Unix()
	ctime := 0
	atime := 0

	return fmt.Sprintf(format, info.Name(), props, size, mtime, ctime, atime)
}
