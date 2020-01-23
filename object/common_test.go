/*
Copyright (c) 2016 VMware, Inc. All Rights Reserved.

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

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestCommonName(t *testing.T) {
	c := &object.Common{}

	name := c.Name()
	if name != "" {
		t.Errorf("Name=%s", name)
	}

	c.InventoryPath = "/foo/bar"
	name = c.Name()
	if name != "bar" {
		t.Errorf("Name=%s", name)
	}
}

func TestObjectName(t *testing.T) {
	type common interface {
		ObjectName(context.Context) (string, error)
	}

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		kinds := []string{"VirtualMachine", "Network", "DistributedVirtualPortgroup"}

		for _, kind := range kinds {
			ref := simulator.Map.Any(kind)
			obj := object.NewReference(c, ref.Reference())

			name, err := obj.(common).ObjectName(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if name == "" {
				t.Errorf("empty name for %s", ref.Reference())
			}
		}
	})
}
