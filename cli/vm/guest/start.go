// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type start struct {
	*GuestFlag

	dir  string
	vars env
}

type env []string

func (e *env) String() string {
	return fmt.Sprint(*e)
}

func (e *env) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func init() {
	cli.Register("guest.start", &start{})
}

func (cmd *start) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestProcessFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.StringVar(&cmd.dir, "C", "", "The absolute path of the working directory for the program to start")
	f.Var(&cmd.vars, "e", "Set environment variable (key=val)")
}

func (cmd *start) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *start) Usage() string {
	return "PATH [ARG]..."
}

func (cmd *start) Description() string {
	return `Start program in VM.

The process can have its status queried with govc guest.ps.
When the process completes, its exit code and end time will be available for 5 minutes after completion.

Examples:
  govc guest.start -vm $name /bin/mount /dev/hdb1 /data
  pid=$(govc guest.start -vm $name /bin/long-running-thing)
  govc guest.ps -vm $name -p $pid -X`
}

func (cmd *start) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.ProcessManager()
	if err != nil {
		return err
	}

	spec := types.GuestProgramSpec{
		ProgramPath:      f.Arg(0),
		Arguments:        strings.Join(f.Args()[1:], " "),
		WorkingDirectory: cmd.dir,
		EnvVariables:     cmd.vars,
	}

	pid, err := m.StartProgram(ctx, cmd.Auth(), &spec)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", pid)

	return nil
}
