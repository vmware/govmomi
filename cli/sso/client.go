// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package sso

import (
	"context"
	"log"
	"os"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/ssoadmin"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vim25/soap"
)

func WithClient(ctx context.Context, cmd *flags.ClientFlag, f func(*ssoadmin.Client) error) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	c, err := ssoadmin.NewClient(ctx, vc)
	if err != nil {
		return err
	}
	c.RoundTripper = cmd.RoundTripper(c.Client)

	// SSO admin server has its own session manager, so the govc persisted session cookies cannot
	// be used to authenticate.  There is no SSO token persistence in govc yet, so just use an env
	// var for now.  If no GOVC_LOGIN_TOKEN is set, issue a new token.
	token := os.Getenv("GOVC_LOGIN_TOKEN")
	header := soap.Header{
		Security: &sts.Signer{
			Certificate: vc.Certificate(),
			Token:       token,
		},
	}

	if token == "" {
		tokens, cerr := sts.NewClient(ctx, vc)
		if cerr != nil {
			return cerr
		}

		req := sts.TokenRequest{
			Certificate: vc.Certificate(),
			Userinfo:    cmd.Session.URL.User,
		}

		header.Security, cerr = tokens.Issue(ctx, req)
		if cerr != nil {
			return cerr
		}
	}

	if err = c.Login(c.WithHeader(ctx, header)); err != nil {
		return err
	}

	defer func() {
		if err := c.Logout(ctx); err != nil {
			log.Printf("user logout error: %v", err)
		}
	}()

	return f(c)
}
