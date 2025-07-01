// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session_test

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/sts"
	_ "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func ExampleManager_LoginByToken() {
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

		// Create a new un-authenticated client and LoginByToken
		vc2, err := vim25.NewClient(ctx, soap.NewClient(vc.URL(), true))
		if err != nil {
			return err
		}

		m := session.NewManager(vc2)
		header := soap.Header{Security: signer}

		err = m.LoginByToken(vc2.WithHeader(ctx, header))
		if err != nil {
			return err
		}

		session, err := m.UserSession(ctx)
		if err != nil {
			return err
		}

		fmt.Println(session.UserName)

		return nil
	})
	// Output: Administrator@VSPHERE.LOCAL
}
