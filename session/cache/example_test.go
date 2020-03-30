/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
