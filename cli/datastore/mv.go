// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
)

type mv struct {
	target
}

func init() {
	cli.Register("datastore.mv", &mv{})
}

func (cmd *mv) Usage() string {
	return "SRC DST"
}

func (cmd *mv) Description() string {
	return `Move SRC to DST on DATASTORE.

Examples:
  govc datastore.mv foo/foo.vmx foo/foo.vmx.old
  govc datastore.mv -f my.vmx foo/foo.vmx`
}

func (cmd *mv) Run(ctx context.Context, f *flag.FlagSet) error {
	args := f.Args()
	if len(args) != 2 {
		return flag.ErrHelp
	}

	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	src, err := cmd.DatastorePath(args[0])
	if err != nil {
		return err
	}

	dst, err := cmd.target.ds.DatastorePath(args[1])
	if err != nil {
		return err
	}

	mv := m.MoveFile
	if cmd.kind {
		mv = m.Move
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Moving %s to %s...", src, dst))
	defer logger.Wait()

	return mv(m.WithProgress(ctx, logger), src, dst)
}
