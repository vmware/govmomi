/*
Copyright (c) 2014-2022 VMware, Inc. All Rights Reserved.

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

package tags

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

func TestManager_GetCategories(t *testing.T) {
	simulator.Run(func(ctx context.Context, client *vim25.Client) error {
		rc := rest.NewClient(client)

		tr := testRoundTripper{
			t:         t,
			transport: rc.Client.Client.Transport,
		}
		rc.Client.Client.Transport = &tr

		err := rc.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatalf("VAPI login: %v", err)
		}

		tm := NewManager(rc)

		// categories to createCount and (concurrently) delete while retrieving categories
		const (
			createCount = 5
			deleteCount = 2
		)

		var created []string
		for i := 1; i <= createCount; i++ {
			cat, err := tm.CreateCategory(ctx, &Category{
				Name: fmt.Sprintf("testcat-%d", i),
			})
			if err != nil {
				t.Fatalf("createCount category: %v", err)
			}
			created = append(created, cat)
		}

		for i := 0; i < deleteCount; i++ {
			tr.deleted = append(tr.deleted, created[i])
		}

		got, err := tm.GetCategories(ctx)
		if err != nil {
			t.Fatalf("get categories: %v", err)
		}

		if len(got) != createCount-deleteCount {
			t.Errorf("category count mismatch: got=%d, want=%d", len(got), createCount-deleteCount)
		}

		return nil
	})
}

// testRoundTripper returns 404 for all categories in deleted
type testRoundTripper struct {
	t *testing.T

	deleted   []string
	transport http.RoundTripper
}

func (tr *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, id := range tr.deleted {
		if strings.Contains(req.URL.Path, id) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Status:     http.StatusText(http.StatusNotFound),
				Body:       nil,
				Request:    req.Clone(context.Background()),
			}, nil
		}
	}

	return tr.transport.RoundTrip(req)
}
