// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package account

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type AccountFlag struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.HostSystemFlag

	types.HostAccountSpec
}

func newAccountFlag(ctx context.Context) (*AccountFlag, context.Context) {
	f := &AccountFlag{}
	f.ClientFlag, ctx = flags.NewClientFlag(ctx)
	f.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	f.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	return f, ctx
}

func (f *AccountFlag) Register(ctx context.Context, fs *flag.FlagSet) {
	f.ClientFlag.Register(ctx, fs)
	f.DatacenterFlag.Register(ctx, fs)
	f.HostSystemFlag.Register(ctx, fs)

	fs.StringVar(&f.Id, "id", "", "The ID of the specified account")
	fs.StringVar(&f.Password, "password", "", "The password for the specified account id")
	fs.StringVar(&f.Description, "description", "", "The description of the specified account")
}

func (f *AccountFlag) Process(ctx context.Context) error {
	if err := f.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := f.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := f.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (f *AccountFlag) HostAccountManager(ctx context.Context) (*object.HostAccountManager, error) {
	h, err := f.HostSystem()
	if err != nil {
		return nil, err
	}

	return h.ConfigManager().AccountManager(ctx)
}
