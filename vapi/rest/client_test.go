// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
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

		path := c.Resource("/xpto/bla")
		err = c.Do(ctx, path.Request(http.MethodGet), nil)
		if !rest.IsStatusError(err, http.StatusNotFound) {
			t.Fatal("expecting error to be 'StatusNotFound'", err)
		}
	})
}

func TestWithHeaders(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)
		if err := c.Login(ctx, simulator.DefaultLogin); err != nil {
			t.Fatal(err)
		}

		// Assign the headers.
		ctx = c.WithHeader(ctx, http.Header{
			"client-token": []string{"unique-token"},
			"pid":          []string{"2", "3", "4"},
		})

		req, err := http.NewRequest(
			http.MethodPost,
			c.Resource(internal.DebugEcho).String(),
			strings.NewReader("Hello, world."))
		if err != nil {
			t.Fatal(err)
		}

		// Send a rest.RawResponse into the decoder to capture the raw
		// response data.
		var res rest.RawResponse
		if err := c.Do(ctx, req, &res); err != nil {
			t.Fatal(err)
		}

		// Read the raw response.
		data, err := io.ReadAll(&res.Buffer)
		if err != nil {
			t.Fatal(err)
		}

		// Assert all of the request headers were echoed back to the client.
		if !bytes.Contains(data, []byte("Client-Token: unique-token")) {
			t.Fatal("missing Client-Token: unique-token")
		}
		if !bytes.Contains(data, []byte("Pid: 2")) {
			t.Fatal("missing Pid: 2")
		}
		if !bytes.Contains(data, []byte("Pid: 3")) {
			t.Fatal("missing Pid: 3")
		}
		if !bytes.Contains(data, []byte("Pid: 4")) {
			t.Fatal("missing Pid: 4")
		}
	})
}
