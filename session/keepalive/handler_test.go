// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package keepalive_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/session/keepalive"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

type count struct {
	sync.Mutex
	val int32
}

func (t *count) Send() error {
	t.Lock()
	defer t.Unlock()
	t.val++
	return nil
}

func (t *count) Value() int {
	t.Lock()
	defer t.Unlock()
	return int(t.val)
}

func TestHandlerSOAP(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		var i count

		sc := soap.NewClient(c.URL(), true)
		vc, err := vim25.NewClient(ctx, sc)
		if err != nil {
			t.Fatal(err)
		}

		vc.RoundTripper = keepalive.NewHandlerSOAP(sc, time.Millisecond, i.Send)

		m := session.NewManager(vc)

		// Expect keep alive to not have triggered yet
		v := i.Value()
		if v != 0 {
			t.Errorf("Expected i == 0, got i: %d", v)
		}

		// Logging in starts keep alive
		err = m.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to triggered at least once
		v = i.Value()
		if v == 0 {
			t.Errorf("Expected i != 0, got i: %d", v)
		}

		j := i.Value()
		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to triggered at least once more
		v = i.Value()
		if v <= j {
			t.Errorf("Expected i > j, got i: %d, j: %d", v, j)
		}

		// Logging out stops keep alive
		err = m.Logout(ctx)
		if err != nil {
			t.Error(err)
		}

		j = i.Value()
		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to have stopped
		v = i.Value()
		if v != j {
			t.Errorf("Expected i == j, got i: %d, j: %d", v, j)
		}
	})
}

func TestHandlerREST(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		var i count

		sc := soap.NewClient(c.URL(), true)
		vc, err := vim25.NewClient(ctx, sc)
		if err != nil {
			t.Fatal(err)
		}

		rc := rest.NewClient(vc)
		rc.Transport = keepalive.NewHandlerREST(rc, time.Millisecond, i.Send)
		if err != nil {
			t.Fatal(err)
		}

		// Expect keep alive to not have triggered yet
		v := i.Value()
		if v != 0 {
			t.Errorf("Expected i == 0, got i: %d", v)
		}

		// Logging in starts keep alive
		err = rc.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Error(err)
		}

		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to triggered at least once
		v = i.Value()
		if v == 0 {
			t.Errorf("Expected i != 0, got i: %d", v)
		}

		j := i.Value()
		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to triggered at least once more
		v = i.Value()
		if v <= j {
			t.Errorf("Expected i > j, got i: %d, j: %d", v, j)
		}

		// Logging out stops keep alive
		err = rc.Logout(ctx)
		if err != nil {
			t.Error(err)
		}

		j = i.Value()
		time.Sleep(10 * time.Millisecond)

		// Expect keep alive to have stopped
		v = i.Value()
		if v != j {
			t.Errorf("Expected i == j, got i: %d, j: %d", v, j)
		}
	})
}
