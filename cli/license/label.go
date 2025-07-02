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
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type label struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("license.label.set", &label{})
}

func (cmd *label) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *label) Usage() string {
	return "LICENSE KEY VAL"
}

func (cmd *label) Description() string {
	return `Set license labels.

Examples:
  govc license.label.set 00000-00000-00000-00000-00000 team cnx # add/set label
  govc license.label.set 00000-00000-00000-00000-00000 team ""  # remove label
  govc license.ls -json | jq '.[] | select(.labels[].key == "team") | .licenseKey'`
}

func (cmd *label) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m := license.NewManager(client)

	if f.NArg() != 3 {
		return flag.ErrHelp
	}

	req := types.UpdateLicenseLabel{
		This:       m.Reference(),
		LicenseKey: f.Arg(0),
		LabelKey:   f.Arg(1),
		LabelValue: f.Arg(2),
	}

	_, err = methods.UpdateLicenseLabel(ctx, m.Client(), &req)
	return err
}
