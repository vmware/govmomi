// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vim25_test

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func ExampleTemporaryNetworkError() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		// Configure retry handler
		delay := time.Millisecond * 100
		retry := func(err error) (bool, time.Duration) {
			return vim25.IsTemporaryNetworkError(err), delay
		}
		c.RoundTripper = vim25.Retry(c.Client, retry, 3)

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
