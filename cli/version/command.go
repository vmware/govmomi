// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type version struct {
	*flags.EmptyFlag

	require string
	long    bool // detailed govc version output
}

func init() {
	cli.Register("version", &version{})
}

func (cmd *version) Register(ctx context.Context, f *flag.FlagSet) {
	f.StringVar(&cmd.require, "require", "", "Require govc version >= this value")
	f.BoolVar(&cmd.long, "l", false, "Print detailed govc version information")
}

func (cmd *version) Run(ctx context.Context, f *flag.FlagSet) error {
	ver := strings.TrimPrefix(flags.BuildVersion, "v")
	if cmd.require != "" {
		v, err := flags.ParseVersion(ver)
		if err != nil {
			panic(err)
		}

		rv, err := flags.ParseVersion(cmd.require)
		if err != nil {
			return fmt.Errorf("failed to parse required version '%s': %s", cmd.require, err)
		}

		if !rv.Lte(v) {
			return fmt.Errorf("version %s or higher is required, this is version %s", cmd.require, ver)
		}
	}

	if cmd.long {
		fmt.Printf("Build Version: %s\n", flags.BuildVersion)
		fmt.Printf("Build Commit: %s\n", flags.BuildCommit)
		fmt.Printf("Build Date: %s\n", flags.BuildDate)
	} else {
		fmt.Printf("govc %s\n", ver)
	}

	return nil
}
