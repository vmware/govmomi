/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package session_test

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type testKeepAlive struct {
	val int32
}

func (t *testKeepAlive) Func(soap.RoundTripper) error {
	atomic.AddInt32(&t.val, 1)
	return nil
}

func (t *testKeepAlive) Value() int {
	n := atomic.LoadInt32(&t.val)
	return int(n)
}

type manager struct {
	*session.Manager
	rt soap.RoundTripper
}

func newManager(u *url.URL, idle time.Duration, handler func(soap.RoundTripper) error) manager {
	sc := soap.NewClient(u, true)
	vc, err := vim25.NewClient(context.Background(), sc)
	if err != nil {
		panic(err)
	}

	if idle != 0 {
		if handler == nil {
			vc.RoundTripper = session.KeepAlive(vc.RoundTripper, idle)
		} else {
			vc.RoundTripper = session.KeepAliveHandler(vc.RoundTripper, idle, handler)
		}
	}

	return manager{session.NewManager(vc), vc.RoundTripper}
}

func TestKeepAlive(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		var i testKeepAlive
		var j int

		m := newManager(c.URL(), time.Millisecond, i.Func)

		// Expect keep alive to not have triggered yet
		if i.Value() != 0 {
			t.Errorf("Expected i == 0, got i: %d", i)
		}

		// Logging in starts keep alive
		err := m.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(2 * time.Millisecond)

		// Expect keep alive to triggered at least once
		if i.Value() == 0 {
			t.Errorf("Expected i != 0, got i: %d", i)
		}

		j = i.Value()
		time.Sleep(2 * time.Millisecond)

		// Expect keep alive to triggered at least once more
		if i.Value() <= j {
			t.Errorf("Expected i > j, got i: %d, j: %d", i, j)
		}

		// Logging out stops keep alive
		err = m.Logout(context.Background())
		if err != nil {
			t.Error(err)
		}

		j = i.Value()
		time.Sleep(2 * time.Millisecond)

		// Expect keep alive to have stopped
		if i.Value() != j {
			t.Errorf("Expected i == j, got i: %d, j: %d", i, j)
		}
	})
}

func testSessionOK(t *testing.T, m manager, ok bool) {
	s, err := m.UserSession(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%s:%d", file, line)

	if ok && s == nil {
		t.Fatalf("%s: Expected session to be OK, but is invalid", prefix)
	}

	if !ok && s != nil {
		t.Fatalf("%s: Expected session to be invalid, but is OK", prefix)
	}
}

// Run with:
//
//   env GOVMOMI_KEEPALIVE_TEST=1 go test -timeout=60m -run TestRealKeepAlive
//
func TestRealKeepAlive(t *testing.T) {
	if os.Getenv("GOVMOMI_KEEPALIVE_TEST") != "1" {
		t.SkipNow()
	}
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}
	var m1, m2 manager
	m1 = newManager(u, 0, nil)
	// Enable keepalive on m2
	m2 = newManager(u, 10*time.Minute, nil)

	// Expect both sessions to be invalid
	testSessionOK(t, m1, false)
	testSessionOK(t, m2, false)

	// Logging in starts keep alive
	if err := m1.Login(context.Background(), u.User); err != nil {
		t.Error(err)
	}
	if err := m2.Login(context.Background(), u.User); err != nil {
		t.Error(err)
	}

	// Expect both sessions to be valid
	testSessionOK(t, m1, true)
	testSessionOK(t, m2, true)

	// Wait for m1 to time out
	delay := 31 * time.Minute
	fmt.Printf("%s: Waiting %d minutes for session to time out...\n", time.Now(), int(delay.Minutes()))
	time.Sleep(delay)

	// Expect m1's session to be invalid, m2's session to be valid
	testSessionOK(t, m1, false)
	testSessionOK(t, m2, true)
}

func isNotAuthenticated(err error) bool {
	if soap.IsSoapFault(err) {
		switch soap.ToSoapFault(err).VimFault().(type) {
		case types.NotAuthenticated:
			return true
		}
	}
	return false
}

func isInvalidLogin(err error) bool {
	if soap.IsSoapFault(err) {
		switch soap.ToSoapFault(err).VimFault().(type) {
		case types.InvalidLogin:
			return true
		}
	}
	return false
}

func TestKeepAliveHandler(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		reauth := make(chan bool)
		login := simulator.DefaultLogin
		var m1, m2 manager
		m1 = newManager(c.URL(), 0, nil)
		// Keep alive handler that will re-login.
		// Real-world case: connectivity to ESX/VC is down long enough for the session to expire
		// Test-world case: we call TerminateSession below
		m2 = newManager(c.URL(), time.Second, func(roundTripper soap.RoundTripper) error {
			_, err := methods.GetCurrentTime(ctx, roundTripper)
			if err != nil {
				if isNotAuthenticated(err) {
					err = m2.Login(ctx, login)
					if err != nil {
						if isInvalidLogin(err) {
							reauth <- false
							t.Log("failed to re-authenticate, quitting keep alive handler")
							return err
						}
					} else {
						reauth <- true
					}
				}
			}

			return nil
		})

		// Logging in starts keep alive
		if err := m1.Login(ctx, simulator.DefaultLogin); err != nil {
			t.Error(err)
		}
		defer m1.Logout(ctx)

		if err := m2.Login(ctx, simulator.DefaultLogin); err != nil {
			t.Error(err)
		}
		defer m2.Logout(ctx)

		// Terminate session for m2.  Note that self terminate fails, so we need 2 sessions for this test.
		s, err := m2.UserSession(ctx)
		if err != nil {
			t.Fatal(err)
		}

		err = m1.TerminateSession(ctx, []string{s.Key})
		if err != nil {
			t.Fatal(err)
		}

		_, err = methods.GetCurrentTime(ctx, m2.rt)
		if err == nil {
			t.Error("expected to fail")
		}

		// Wait for keepalive to re-authenticate
		<-reauth

		_, err = methods.GetCurrentTime(ctx, m2.rt)
		if err != nil {
			t.Fatal(err)
		}

		// Clear credentials to test re-authentication failure
		login = nil

		s, err = m2.UserSession(ctx)
		if err != nil {
			t.Fatal(err)
		}

		err = m1.TerminateSession(ctx, []string{s.Key})
		if err != nil {
			t.Fatal(err)
		}

		// Wait for keepalive re-authenticate attempt
		result := <-reauth

		_, err = methods.GetCurrentTime(ctx, m2.rt)
		if err == nil {
			t.Error("expected to fail")
		}

		if result {
			t.Errorf("expected reauth to fail")
		}
	})
}
