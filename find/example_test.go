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

package find_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func ExampleMultipleFoundError() {
	model := simulator.VPX()
	model.Portgroup = 2

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)

		_, err := finder.Network(ctx, "DC0_DVPG*")
		_, ok := err.(*find.MultipleFoundError) // returns DC0_DVPG{0,1}
		if !ok {
			return errors.New("expected error")
		}

		net0, err := finder.Network(ctx, "DC0_DVPG0")
		if err != nil {
			return err
		}
		fmt.Println(net0.GetInventoryPath())

		net1, err := finder.Network(ctx, "DC0_DVPG1")
		if err != nil {
			return err
		}
		fmt.Println(net1.GetInventoryPath())

		return nil
	}, model)
	// Output:
	// /DC0/network/DC0_DVPG0
	// /DC0/network/DC0_DVPG1
}
