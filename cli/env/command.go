// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package env

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
)

type env struct {
	*flags.OutputFlag
	*flags.ClientFlag

	extra bool
}

func init() {
	cli.Register("env", &env{})
}

func (cmd *env) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.extra, "x", false, "Output variables for each GOVC_URL component")
}

func (cmd *env) Process(ctx context.Context) error {
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *env) Description() string {
	return `Output the environment variables for this client.

If credentials are included in the url, they are split into separate variables.
Useful as bash scripting helper to parse GOVC_URL.`
}

func (cmd *env) Run(ctx context.Context, f *flag.FlagSet) error {
	env := envResult(cmd.ClientFlag.Environ(cmd.extra))

	if f.NArg() > 1 {
		return flag.ErrHelp
	}

	// Option to just output the value, example use:
	// password=$(govc env GOVC_PASSWORD)
	if f.NArg() == 1 {
		var output []string

		prefix := fmt.Sprintf("%s=", f.Arg(0))

		for _, e := range env {
			if strings.HasPrefix(e, prefix) {
				output = append(output, e[len(prefix):])
				break
			}
		}

		return cmd.WriteResult(envResult(output))
	}

	return cmd.WriteResult(env)
}

type envResult []string

func (r envResult) Write(w io.Writer) error {
	for _, e := range r {
		fmt.Println(e)
	}

	return nil
}
