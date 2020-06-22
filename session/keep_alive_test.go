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
		c.RoundTripper = vim25.Retry(c.Client, vim25.TemporaryNetworkError(3))
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
