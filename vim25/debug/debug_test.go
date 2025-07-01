// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package debug_test

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/debug"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func TestSetProvider(t *testing.T) {
	p := debug.FileProvider{
		Path: t.TempDir(),
	}
	debug.SetProvider(&p)

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		var wg sync.WaitGroup
		rc := rest.NewClient(c)

		// hit the debug package with some concurrency (see PR #2469)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				finder := find.NewFinder(c)

				_, err := finder.VirtualMachineList(ctx, "*")
				if err != nil {
					t.Error(err)
				}
			}()
		}

		wg.Wait()

		// send an http request with a nil Body to ensure debug trace doesn't panic in this case
		u := rc.URL().String() + "/com/vmware/cis/session"

		req, err := http.NewRequest(http.MethodPost, u, nil)
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("user", "pass")
		var id string
		if err = rc.Do(ctx, req, &id); err != nil {
			t.Fatal(err)
		}
	})
}
