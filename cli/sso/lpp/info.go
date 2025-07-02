// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package lpp

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/sso"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("sso.lpp.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `Get SSO local password policy.

Examples:
  govc sso.lpp.info
  govc sso.lpp.info -json`
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

type lppInfo struct {
	LocalPasswordPolicy *types.AdminPasswordPolicy
}

func (r *lppInfo) Write(w io.Writer) error {
	fmt.Fprintf(
		w,
		"Description: %s\n"+
			"MinLength: %d\n"+
			"MaxLength: %d\n"+
			"MinAlphabeticCount: %d\n"+
			"MinUppercaseCount: %d\n"+
			"MinLowercaseCount: %d\n"+
			"MinNumericCount: %d\n"+
			"MinSpecialCharCount: %d\n"+
			"MaxIdenticalAdjacentCharacters: %d\n"+
			"ProhibitedPreviousPasswordsCount: %d\n"+
			"PasswordLifetimeDays: %d\n",
		r.LocalPasswordPolicy.Description,
		r.LocalPasswordPolicy.PasswordFormat.LengthRestriction.MinLength,
		r.LocalPasswordPolicy.PasswordFormat.LengthRestriction.MaxLength,
		r.LocalPasswordPolicy.PasswordFormat.AlphabeticRestriction.MinAlphabeticCount,
		r.LocalPasswordPolicy.PasswordFormat.AlphabeticRestriction.MinUppercaseCount,
		r.LocalPasswordPolicy.PasswordFormat.AlphabeticRestriction.MinLowercaseCount,
		r.LocalPasswordPolicy.PasswordFormat.MinNumericCount,
		r.LocalPasswordPolicy.PasswordFormat.MinSpecialCharCount,
		r.LocalPasswordPolicy.PasswordFormat.MaxIdenticalAdjacentCharacters,
		r.LocalPasswordPolicy.ProhibitedPreviousPasswordsCount,
		r.LocalPasswordPolicy.PasswordLifetimeDays,
	)
	return nil
}

func (r *lppInfo) Dump() any {
	return r.LocalPasswordPolicy
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		var err error
		var pol lppInfo
		pol.LocalPasswordPolicy, err = c.GetLocalPasswordPolicy(ctx)
		if err != nil {
			return err
		}
		return cmd.WriteResult(&pol)
	})
}
