// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func TestKeepAlive(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		m := session.NewManager(c)

		err := m.Logout(ctx)
		if err != nil {
			t.Fatal(err)
		}

		var mu sync.Mutex
		n := 0
		c.RoundTripper = vim25.Retry(c.Client, vim25.RetryTemporaryNetworkError, 3)
		c.RoundTripper = session.KeepAliveHandler(c.RoundTripper, time.Millisecond, func(soap.RoundTripper) error {
			mu.Lock()
			n++
			mu.Unlock()
			return errors.New("stop") // stops the keep alive routine
		})

		err = m.Login(ctx, simulator.DefaultLogin) // starts the keep alive routine
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(time.Millisecond * 10)
		mu.Lock()
		if n != 1 {
			t.Errorf("handler called %d times", n)
		}
		mu.Unlock()
	})
}
