// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type rm struct {
	*flags.DatastoreFlag

	kind        bool
	force       bool
	isNamespace bool
}

func init() {
	cli.Register("datastore.rm", &rm{})
	cli.Alias("datastore.rm", "datastore.delete")
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.kind, "t", true, "Use file type to choose disk or file manager")
	f.BoolVar(&cmd.force, "f", false, "Force; ignore nonexistent files and arguments")
	f.BoolVar(&cmd.isNamespace, "namespace", false, "Path is uuid of namespace on vsan datastore")
}

func (cmd *rm) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rm) Usage() string {
	return "FILE"
}

func (cmd *rm) Description() string {
	return `Remove FILE from DATASTORE.

Examples:
  govc datastore.rm vm/vmware.log
  govc datastore.rm vm
  govc datastore.rm -f images/base.vmdk`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	args := f.Args()
	if len(args) == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	var dc *object.Datacenter
	dc, err = cmd.Datacenter()
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	if cmd.isNamespace {
		path := args[0]

		nm := object.NewDatastoreNamespaceManager(c)
		err = nm.DeleteDirectory(ctx, dc, path)
	} else {
		fm := ds.NewFileManager(dc, cmd.force)

		remove := fm.DeleteFile // File delete
		if cmd.kind {
			remove = fm.Delete // VirtualDisk or File delete
		}

		err = remove(ctx, args[0])
	}

	if err != nil {
		if types.IsFileNotFound(err) && cmd.force {
			// Ignore error
			return nil
		}
	}

	return err
}
