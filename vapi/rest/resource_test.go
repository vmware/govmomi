package rest_test

import (
	"context"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"strings"
	"testing"
)

func TestResource_WithParam(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		url := c.Resource("api/some/resource").
			WithParam("key1", "value1")
		expectedPath := "api/some/resource?key1=value1"
		if ! strings.Contains(url.String(), expectedPath) {
			t.Errorf("Param incorrectly added to resource, URL %q, expected path %q", url.String(), expectedPath)
		}

		url = url.WithParam("key2", "value2")
		expectedPath = "api/some/resource?key1=value1&key2=value2"
		if ! strings.Contains(url.String(), expectedPath) {
			t.Errorf("Param incorrectly updated on resource, URL %q, expected path %q", url.String(), expectedPath)
		}
	})
}