// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"fmt"
	"path"
	"strings"
)

// deriveImportTargets centralizes logic for choosing the OVF entity name (no slashes),
// upload folder, upload target, and final rename destination based on ImportParams
// and the local disk info. It mirrors the behavior used in Import.
func deriveImportTargets(p ImportParams, disk *Info) (entityName, target string, err error) {
	// original import name derived from local file (base name without .vmdk)
	origImport := disk.ImportName

	// If no remote path provided, fall back to original behavior where entityName
	// is derived from the local vmdk base name and upload folder == entityName.
	if strings.TrimSpace(p.Path) == "" {
		entityName = origImport
		target = fmt.Sprintf("%s/%s.vmdk", origImport, origImport)
		return
	}

	// remotePath is the provided target path (should include the vmdk filename)
	remotePath := strings.TrimSuffix(p.Path, "/")
	base := path.Base(remotePath)

	if path.Ext(base) != ".vmdk" {
		err = fmt.Errorf("remote target must end with .vmdk; the last element will be the vmdk name (e.g. windows/win2019/mydisk.vmdk)")
		return
	}

	// entityName is the vmdk file name without extension
	entityName = strings.TrimSuffix(base, ".vmdk")

	// dir is the folder portion of the provided path. If none, we implicitly create a folder named after the vmdk
	dir := path.Dir(remotePath)
	uploadFolder := dir
	if dir == "." || dir == "" {
		uploadFolder = entityName
	}

	// target is where ImportVApp will place the uploaded disk (uploadFolder/entityName.vmdk).
	target = fmt.Sprintf("%s/%s.vmdk", uploadFolder, entityName)

	return
}
