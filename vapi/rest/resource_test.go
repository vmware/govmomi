// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest_test

import (
	"context"
	"strings"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

func TestResource_WithParam(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		url := c.Resource("api/some/resource").
			WithParam("key1", "value1")
		expectedPath := "api/some/resource?key1=value1"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("Param incorrectly added to resource, URL %q, expected path %q", url.String(), expectedPath)
		}

		url = url.WithParam("key2", "value2")
		expectedPath = "api/some/resource?key1=value1&key2=value2"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("Param incorrectly updated on resource, URL %q, expected path %q", url.String(), expectedPath)
		}

		url = c.Resource("api/some/resource")
		url = url.WithParam("names", "foo").WithParam("names", "bar")
		expectedPath = "api/some/resource?names=foo&names=bar"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("Param incorrectly updated on resource, URL %q, expected path %q", url.String(), expectedPath)
		}
	})
}

func TestResource_WithPathEncodedParam(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		// path is correctly formatted when Path-Encoded param is first
		url := c.Resource("api/some/resource").
			WithPathEncodedParam("key1", "value 1")
		expectedPath := "api/some/resource?key1=value%201"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("First path-encoded param incorrectly added to resource, URL %q, expected path %q", url.String(), expectedPath)
		}

		// path is correctly formatted when Path-Encoded param is last
		url = c.Resource("api/some/resource").
			WithParam("key1", "value 1").
			WithPathEncodedParam("key2", "value 2")
		expectedPath = "api/some/resource?key1=value+1&key2=value%202"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("Last path-encoded param incorrectly added to resource, URL %q, expected path %q", url.String(), expectedPath)
		}

		// if WithParam is used again, it will re-encode the Path-Encoded value
		url = url.WithParam("key3", "value 3")
		expectedPath = "api/some/resource?key1=value+1&key2=value+2&key3=value+3"
		if !strings.Contains(url.String(), expectedPath) {
			t.Errorf("Middle path-encoded param not endcoded as expected, URL %q, expected path %q", url.String(), expectedPath)
		}

	})
}
