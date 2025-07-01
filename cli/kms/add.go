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
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.ClientFlag

	types.KmipServerSpec
	native vapicrypto.KmsProviderCreateSpec
	nkp    bool
}

func init() {
	cli.Register("kms.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.Info.Name, "n", "", "Server name")
	f.StringVar(&cmd.Info.Address, "a", "", "Server address")
	cmd.Info.Port = 5696 // default
	f.Var(flags.NewInt32(&cmd.Info.Port), "p", "Server port")

	f.BoolVar(&cmd.nkp, "N", false, "Add native key provider")
	f.BoolVar(&cmd.native.Constraints.TpmRequired, "tpm", true, "Use only with TPM protected ESXi hosts (native only)")
}

func (cmd *add) Usage() string {
	return "NAME"
}

func (cmd *add) Description() string {
	return `Add KMS cluster.

Server name and address are required, port defaults to 5696.

Examples:
  govc kms.add -N knp
  govc kms.add -n my-server -a kms.example.com my-kp`
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
		return flag.ErrHelp
	}

	if cmd.nkp {
		rc, err := cmd.RestClient()
		if err != nil {
			return err
		}
		cmd.native.Provider = id
		return vapicrypto.NewManager(rc).KmsProviderCreate(ctx, cmd.native)
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	spec := types.KmipServerSpec{
		ClusterId: types.KeyProviderId{Id: id},
		Info:      cmd.Info,
	}

	return m.RegisterKmipServer(ctx, spec)
}
