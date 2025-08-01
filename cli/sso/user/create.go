// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package user

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/sso"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/ssoadmin/types"
)

type userDetails struct {
	*flags.ClientFlag

	types.AdminPersonDetails
	password string
	solution types.AdminSolutionDetails
	actas    *bool
	role     string
}

func (cmd *userDetails) Usage() string {
	return "NAME"
}

func (cmd *userDetails) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.Description, "d", "", "User description")
	f.StringVar(&cmd.EmailAddress, "m", "", "Email address")
	f.StringVar(&cmd.FirstName, "f", "", "First name")
	f.StringVar(&cmd.LastName, "l", "", "Last name")
	f.StringVar(&cmd.password, "p", "", "Password")
	f.StringVar(&cmd.solution.Certificate, "C", "", "Certificate for solution user")
	f.Var(flags.NewOptionalBool(&cmd.actas), "A", "ActAsUser role for solution user WSTrust")
	f.StringVar(&cmd.role, "R", "", "Role for solution user (RegularUser|Administrator)")
}

func (cmd *userDetails) Certificate() string {
	block, _ := pem.Decode([]byte(cmd.solution.Certificate))
	if block != nil {
		return base64.StdEncoding.EncodeToString(block.Bytes)
	}
	return cmd.solution.Certificate
}

type create struct {
	userDetails
}

func init() {
	cli.Register("sso.user.create", &create{})
}

func (cmd *create) Description() string {
	return `Create SSO users.

Examples:
  govc sso.user.create -C "$(cat cert.pem)" -A -R Administrator NAME # solution user
  govc sso.user.create -p password NAME # person user`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	id := f.Arg(0)
	person := cmd.solution.Certificate == ""
	if person {
		if cmd.password == "" {
			return flag.ErrHelp
		}
	} else {
		if cmd.password != "" {
			return flag.ErrHelp
		}
	}

	return sso.WithClient(ctx, cmd.ClientFlag, func(c *ssoadmin.Client) error {
		if person {
			return c.CreatePersonUser(ctx, id, cmd.AdminPersonDetails, cmd.password)
		}

		cmd.solution.Certificate = cmd.Certificate()
		cmd.solution.Description = cmd.AdminPersonDetails.Description

		if err := c.CreateSolutionUser(ctx, id, cmd.solution); err != nil {
			return err
		}

		p := types.PrincipalId{Name: id, Domain: c.Domain}

		if cmd.role != "" {
			if _, err := c.SetRole(ctx, p, cmd.role); err != nil {
				return err
			}
		}

		if cmd.actas != nil && *cmd.actas {
			if _, err := c.GrantWSTrustRole(ctx, p, types.RoleActAsUser); err != nil {
				return err
			}
		}

		return nil
	})
}
