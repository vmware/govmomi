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

package keepalive_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/session/keepalive"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

var (
	sessionCheckPause  = time.Second / 2
	sessionIdleTimeout = sessionCheckPause / 2
	keepAliveIdle      = sessionIdleTimeout / 2
)

func init() {
	simulator.SessionIdleTimeout = sessionIdleTimeout
}

func ExampleHandlerSOAP() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		// No need for initial Login() here as simulator.Run already has
		m := session.NewManager(c)

		// check twice if session is valid, sleeping > SessionIdleTimeout in between
		check := func() {
			for i := 0; i < 2; i++ {
				s, err := m.UserSession(ctx)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("session valid=%t\n", s != nil)
				if i == 0 {
					time.Sleep(sessionCheckPause)
				}
			}
		}

		// session will expire here
		check()

		// this starts the keep alive handler when Login is called, and stops the handler when Logout is called
		c.RoundTripper = keepalive.NewHandlerSOAP(c.RoundTripper, keepAliveIdle, nil)

		err := m.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		// session will not expire here, with the keep alive handler in place.
		check()

		err = m.Logout(ctx)
		if err != nil {
			return err
		}

		// Logout() also stops the keep alive handler, session is no longer valid.
		check()

		return nil
	})
	// Output:
	// session valid=true
	// session valid=false
	// session valid=true
	// session valid=true
	// session valid=false
	// session valid=false
}

func ExampleHandlerREST() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)
		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		// check twice if session is valid, sleeping > SessionIdleTimeout in between.
		check := func() {
			for i := 0; i < 2; i++ {
				s, err := c.Session(ctx)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("session valid=%t\n", s != nil)
				if i == 0 {
					time.Sleep(sessionCheckPause)
				}
			}
		}

		// session will expire here
		check()

		// this starts the keep alive handler when Login is called, and stops the handler when Logout is called
		c.Transport = keepalive.NewHandlerREST(c, keepAliveIdle, nil)

		err = c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		// session will not expire here, with the keep alive handler in place.
		check()

		err = c.Logout(ctx)
		if err != nil {
			return err
		}

		// Logout() also stops the keep alive handler, session is no longer valid.
		check()

		return nil
	})
	// Output:
	// session valid=true
	// session valid=false
	// session valid=true
	// session valid=true
	// session valid=false
	// session valid=false
}
