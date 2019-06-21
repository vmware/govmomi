/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sso

import (
	"context"
	"log"
	"os"

	"github.com/vmware/govmomi/govc/flags"
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
			Userinfo:    cmd.Userinfo(),
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
