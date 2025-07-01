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
	"github.com/vmware/govmomi/vim25/types"
)

type setdefault struct {
	*flags.DatacenterFlag

	entity string
}

func init() {
	cli.Register("kms.default", &setdefault{})
}

func (cmd *setdefault) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.entity, "e", "", "Set entity default KMS cluster (cluster or host folder)")
}

func (cmd *setdefault) Usage() string {
	return "NAME"
}

func (cmd *setdefault) Description() string {
	return `Set default KMS cluster.

Examples:
  govc kms.default my-kp
  govc kms.default - # clear default
  govc kms.default -e /dc/host/cluster my-kp
  govc kms.default -e /dc/host/cluster my-kp - # clear default`
}

func (cmd *setdefault) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
		return flag.ErrHelp
	}

	if id == "-" {
		id = "" // clear default
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	var entity *types.ManagedObjectReference

	if cmd.entity != "" {
		obj, err := cmd.ManagedObject(ctx, cmd.entity)
		if err != nil {
			return err
		}

		entity = types.NewReference(obj.Reference())
	}

	return m.SetDefaultKmsClusterId(ctx, id, entity)
}
