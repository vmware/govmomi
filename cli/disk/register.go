// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type register struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register("disk.register", &register{})
}

func (cmd *register) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *register) Usage() string {
	return "PATH [NAME]"
}

func (cmd *register) Description() string {
	return `Register existing disk on DS.

Examples:
  govc disk.register disks/disk1.vmdk my-disk`
}

func (cmd *register) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	path := ds.NewURL(f.Arg(0)).String()

	obj, err := m.RegisterDisk(ctx, path, f.Arg(1))
	if err != nil {
		return err
	}

	fmt.Println(obj.Config.Id.Id)

	return nil
}
