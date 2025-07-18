// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type mkdir struct {
	*flags.DatastoreFlag

	createParents bool
	isNamespace   bool
}

func init() {
	cli.Register("datastore.mkdir", &mkdir{})
}

func (cmd *mkdir) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.createParents, "p", false, "Create intermediate directories as needed")
	f.BoolVar(&cmd.isNamespace, "namespace", false, "Return uuid of namespace created on vsan datastore")
}

func (cmd *mkdir) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *mkdir) Usage() string {
	return "DIRECTORY"
}

func (cmd *mkdir) Run(ctx context.Context, f *flag.FlagSet) error {
	args := f.Args()
	if len(args) == 0 {
		return errors.New("missing operand")
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	if cmd.isNamespace {
		var uuid string
		var ds *object.Datastore

		if ds, err = cmd.Datastore(); err != nil {
			return err
		}

		path := args[0]

		nm := object.NewDatastoreNamespaceManager(c)
		if uuid, err = nm.CreateDirectory(ctx, ds, path, ""); err != nil {
			return err
		}

		fmt.Println(uuid)
	} else {
		var dc *object.Datacenter
		var path string

		dc, err = cmd.Datacenter()
		if err != nil {
			return err
		}

		path, err = cmd.DatastorePath(args[0])
		if err != nil {
			return err
		}

		m := object.NewFileManager(c)
		err = m.MakeDirectory(ctx, path, dc, cmd.createParents)

		// ignore EEXIST if -p flag is given
		if err != nil && cmd.createParents {
			if fault.Is(err, &types.FileAlreadyExists{}) {
				return nil
			}
		}
	}

	return err
}
