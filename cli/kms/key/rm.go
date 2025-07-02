// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
)

type rm struct {
	*flags.ClientFlag

	provider string
	force    bool
}

func init() {
	cli.Register("kms.key.rm", &rm{}, true)
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.provider, "p", "", "Provider ID")
	f.BoolVar(&cmd.force, "f", false, "Force")
}

func (cmd *rm) Usage() string {
	return "ID..."
}

func (cmd *rm) Description() string {
	return `Remove crypto keys.

Examples:
  govc kms.key.rm -p my-kp ID`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	n := f.NArg()
	if n == 0 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	ids := argsToKeys(cmd.provider, f.Args())

	return m.RemoveKeys(ctx, ids, cmd.force)
}
