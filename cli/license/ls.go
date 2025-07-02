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

var featureUsage = "List licenses with given feature"

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	feature string
}

func init() {
	cli.Register("license.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.feature, "feature", "", featureUsage)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m := license.NewManager(client)
	result, err := m.List(ctx)
	if err != nil {
		return err
	}

	if cmd.feature != "" {
		result = result.WithFeature(cmd.feature)
	}

	return cmd.WriteResult(licenseOutput(result))
}
