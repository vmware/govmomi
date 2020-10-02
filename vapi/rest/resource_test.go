/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
	})
}
