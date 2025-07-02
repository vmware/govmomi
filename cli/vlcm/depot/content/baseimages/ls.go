// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package baseimages

import (
	"context"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/depots"
)

type lsResult []depots.BaseImagesSummary

func (r lsResult) Write(w io.Writer) error {
	return nil
}

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vlcm.depot.baseimages.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *ls) Usage() string {
	return "VLCM"
}

func (cmd *ls) Description() string {
	return `Displays the list of available ESXi base images.

Examples:
  govc vlcm.depot.baseimages.ls`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := depots.NewManager(rc)

	if res, err := dm.ListBaseImages(); err != nil {
		return err
	} else {
		if !cmd.All() {
			cmd.JSON = true
		}
		return cmd.WriteResult(lsResult(res))
	}
}
