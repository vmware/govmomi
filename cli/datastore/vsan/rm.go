// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vsan

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type rm struct {
	*flags.DatastoreFlag

	force   bool
	verbose bool
}

func init() {
	cli.Register("datastore.vsan.dom.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.force, "f", false, "Force delete")
	f.BoolVar(&cmd.verbose, "v", false, "Print deleted UUIDs to stdout, failed to stderr")
}

func (cmd *rm) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rm) Usage() string {
	return "UUID..."
}

func (cmd *rm) Description() string {
	return `Remove vSAN DOM objects in DS.

Examples:
  govc datastore.vsan.dom.rm d85aa758-63f5-500a-3150-0200308e589c
  govc datastore.vsan.dom.rm -f d85aa758-63f5-500a-3150-0200308e589c
  govc datastore.vsan.dom.ls -o | xargs govc datastore.vsan.dom.rm`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	hosts, err := ds.AttachedHosts(ctx)
	if err != nil {
		return err
	}

	if len(hosts) == 0 {
		return flag.ErrHelp
	}

	m, err := hosts[0].ConfigManager().VsanInternalSystem(ctx)
	if err != nil {
		return err
	}

	res, err := m.DeleteVsanObjects(ctx, f.Args(), &cmd.force)
	if err != nil {
		return err
	}

	if cmd.verbose {
		for _, r := range res {
			if r.Success {
				fmt.Fprintln(cmd.Out, r.Uuid)
			} else {
				fmt.Fprintf(os.Stderr, "%s %s\n", r.Uuid, r.FailureReason[0].Message)
			}
		}
	}

	return nil
}
