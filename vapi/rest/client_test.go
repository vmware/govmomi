/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package rest_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func TestSession(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		session, err := c.Session(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if session != nil {
			t.Fatal("expected nil session")
		}

		err = c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		session, err = c.Session(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if session == nil {
			t.Fatal("expected non-nil session")
		}
	})
}
