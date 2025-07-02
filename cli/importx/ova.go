// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importx

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/ovf/importer"
)

type ova struct {
	*ovfx
}

func init() {
	cli.Register("import.ova", &ova{&ovfx{}})
}

func (cmd *ova) Usage() string {
	return "PATH_TO_OVA"
}

func (cmd *ova) Run(ctx context.Context, f *flag.FlagSet) error {
	fpath, err := cmd.Prepare(f)
	if err != nil {
		return err
	}

	archive := &importer.TapeArchive{Path: fpath}
	archive.Client = cmd.Importer.Client

	cmd.Importer.Archive = archive
	fpath = "*.ovf"

	return cmd.Import(ctx, fpath)
}
