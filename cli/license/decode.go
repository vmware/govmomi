// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/license"
)

type decode struct {
	*flags.ClientFlag
	*flags.OutputFlag

	feature string
}

func init() {
	cli.Register("license.decode", &decode{})
}

func (cmd *decode) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.feature, "feature", "", featureUsage)
}

func (cmd *decode) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *decode) Usage() string {
	return "KEY..."
}

func (cmd *decode) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	var result license.InfoList
	m := license.NewManager(client)
	for _, v := range f.Args() {
		license, err := m.Decode(ctx, v)
		if err != nil {
			return err
		}

		result = append(result, license)
	}

	if cmd.feature != "" {
		result = result.WithFeature(cmd.feature)
	}

	return cmd.WriteResult(licenseOutput(result))
}
