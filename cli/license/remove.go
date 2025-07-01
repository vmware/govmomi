// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/license"
)

type remove struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("license.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "KEY..."
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m := license.NewManager(client)
	for _, v := range f.Args() {
		err = m.Remove(ctx, v)
		if err != nil {
			return err
		}
	}

	list, err := m.List(ctx)
	if err != nil {
		return err
	}

	for _, v := range f.Args() {
		for _, l := range list {
			if v == l.LicenseKey {
				return fmt.Errorf("cannot remove license %q", v)
			}
		}
	}

	return nil
}
