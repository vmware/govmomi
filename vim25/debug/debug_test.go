/*
   Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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
