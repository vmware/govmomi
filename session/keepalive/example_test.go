// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package keepalive_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/session/keepalive"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/simulator/sim25"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

var (
	sessionCheckPause  = time.Second / 2
	sessionIdleTimeout = sessionCheckPause / 2
	keepAliveIdle      = sessionIdleTimeout / 2
)

func ExampleHandlerSOAP() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		// Using the authenticated vc client, timeout will apply to new sessions.
		sim25.SetSessionTimeout(ctx, vc, sessionIdleTimeout)

		c, err := govmomi.NewClient(ctx, vc.URL(), true)
		if err != nil {
			return err
		}
		m := c.SessionManager

		err = m.Login(ctx, simulator.DefaultLogin) // New session
		if err != nil {
			return err
		}

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

		err = m.Login(ctx, simulator.DefaultLogin)
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
		// Using the authenticated vc client, timeout will apply to new sessions.
		sim25.SetSessionTimeout(ctx, vc, sessionIdleTimeout)

		c := rest.NewClient(vc)
		err := c.Login(ctx, simulator.DefaultLogin) // New session
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
