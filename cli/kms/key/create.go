// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
)

type create struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("kms.key.create", &create{}, true)
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *create) Usage() string {
	return "ID"
}

func (cmd *create) Description() string {
	return `Generate crypto key.

Examples:
  govc kms.key.create my-kp`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
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

	key, err := m.GenerateKey(ctx, id)
	if err != nil {
		return err
	}

	fmt.Println(key)

	return nil
}
