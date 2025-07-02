// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest_test

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/sts"
	_ "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

func ExampleClient_LoginByToken() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c, err := sts.NewClient(ctx, vc)
		if err != nil {
			return err
		}

		// Issue a bearer token
		req := sts.TokenRequest{
			Userinfo: url.UserPassword("Administrator@VSPHERE.LOCAL", "password"),
		}

		signer, err := c.Issue(ctx, req)
		if err != nil {
			return err
		}

		rc := rest.NewClient(vc)

		err = rc.LoginByToken(rc.WithSigner(ctx, signer))
		if err != nil {
			return err
		}

		session, err := rc.Session(ctx)
		if err != nil {
			return err
		}

		// Note: vcsim does not currently parse the token NameID for rest as it does for soap
		fmt.Println(session.User)

		return nil
	})
	// Output: TODO
}
