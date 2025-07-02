// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cache_test

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"

	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
)

func ExampleSession_Login() {
	simulator.Run(func(ctx context.Context, sim *vim25.Client) error {
		u := sim.URL()
		u.User = simulator.DefaultLogin

		// Login() each client twice
		// 1) creates a new authenticated session
		// 2) uses a cached session (proved by removing the password)
		for i := 0; i < 2; i++ {
			s := &cache.Session{
				URL:      u,
				Insecure: true,
			}

			vc := new(vim25.Client)
			err := s.Login(ctx, vc, nil)
			if err != nil {
				return err
			}

			rc := new(rest.Client)
			err = s.Login(ctx, rc, nil)
			if err != nil {
				return err
			}

			fmt.Printf("%s authenticated\n", u.User)
			u.User = url.User(u.User.Username()) // Remove password
		}
		return nil
	})
	// Output:
	// user:pass authenticated
	// user authenticated
}
