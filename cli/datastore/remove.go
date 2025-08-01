// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type remove struct {
	*flags.HostSystemFlag
	*flags.DatastoreFlag
}

func init() {
	cli.Register("datastore.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "HOST..."
}

func (cmd *remove) Description() string {
	return `Remove datastore from HOST.

Examples:
  govc datastore.remove -ds nfsDatastore cluster1
  govc datastore.remove -ds nasDatastore host1 host2 host3`
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		hds, err := host.ConfigManager().DatastoreSystem(ctx)
		if err != nil {
			return err
		}

		err = hds.Remove(ctx, ds)
		if err != nil {
			return err
		}
	}

	return nil
}
