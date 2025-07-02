// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package kms

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
	vapicrypto "github.com/vmware/govmomi/vapi/crypto"
)

type rm struct {
	*flags.ClientFlag

	server string
}

func init() {
	cli.Register("kms.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.server, "s", "", "Server name")
}

func (cmd *rm) Usage() string {
	return "NAME"
}

func (cmd *rm) Description() string {
	return `Remove KMS server or cluster.

Examples:
  govc kms.rm my-kp
  govc kms.rm -s my-server my-kp`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
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

	native, err := m.IsNativeProvider(ctx, id)
	if err != nil {
		return err
	}
	if native {
		rc, err := cmd.RestClient()
		if err != nil {
			return err
		}
		return vapicrypto.NewManager(rc).KmsProviderDelete(ctx, id)
	}

	if cmd.server != "" {
		return m.RemoveKmipServer(ctx, id, cmd.server)
	}

	return m.UnregisterKmsCluster(ctx, id)
}
