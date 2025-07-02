// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"testing"

	"github.com/vmware/govmomi/cli"
)

func TestMain(t *testing.T) {
	// Execute flag registration for every command to verify there are no
	// commands with flag name collisions
	for _, cmd := range cli.Commands() {
		fs := flag.NewFlagSet("", flag.ContinueOnError)

		// Use fresh context for every command
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cmd.Register(ctx, fs)
	}
}
