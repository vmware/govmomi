// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
)

func TestClientCookie(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c, err := vslm.NewClient(ctx, vc)
		if err != nil {
			t.Fatal(err)
		}

		m := vslm.NewGlobalObjectManager(c)
		// Using default / expected Header.Cookie.XMLName = vcSessionCookie
		_, err = m.ListObjectsForSpec(ctx, nil, 1000)
		if err != nil {
			t.Fatal(err)
		}

		// Using invalid Header.Cookie.XMLName = myCookie
		myCookie := vc.SessionCookie()
		myCookie.XMLName.Local = "myCookie"

		c.Client.Cookie = func() *soap.HeaderElement {
			return myCookie
		}

		_, err = m.ListObjectsForSpec(ctx, nil, 1000)
		if !fault.Is(err, &types.NotAuthenticated{}) {
			t.Errorf("err=%#v", err)
		}
	})
}
