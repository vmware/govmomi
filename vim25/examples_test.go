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

package vim25_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func ExampleTemporaryNetworkError() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		// Configure retry handler
		c.RoundTripper = vim25.Retry(c.Client, vim25.TemporaryNetworkError(3))

		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		// Tell vcsim to respond with 502 on the 1st request
		simulator.StatusSDK = http.StatusBadGateway

		state, err := vm.PowerState(ctx)
		if err != nil {
			return err
		}

		fmt.Println(state)

		return nil
	})
	// Output: poweredOn
}
