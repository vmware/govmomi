// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"bytes"
	"context"
	"flag"
	"os"
	"os/exec"

	"github.com/vmware/govmomi/cli"
)

type run struct {
	*GuestFlag

	data string
	dir  string
	vars env
}

func init() {
	cli.Register("guest.run", &run{})
}

func (cmd *run) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestProcessFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.StringVar(&cmd.data, "d", "", "Input data string. A value of '-' reads from OS stdin")
	f.StringVar(&cmd.dir, "C", "", "The absolute path of the working directory for the program to start")
	f.Var(&cmd.vars, "e", "Set environment variables")
}

func (cmd *run) Usage() string {
	return "PATH [ARG]..."
}

func (cmd *run) Description() string {
	return `Run program PATH in VM and display output.

The guest.run command starts a program in the VM with i/o redirected, waits for the process to exit and
propagates the exit code to the govc process exit code.  Note that stdout and stderr are redirected by default,
stdin is only redirected when the '-d' flag is specified.

Note that vmware-tools requires program PATH to be absolute.
If PATH is not absolute and vm guest family is Windows,
guest.run changes the command to: 'c:\\Windows\\System32\\cmd.exe /c "PATH [ARG]..."'
Otherwise the command is changed to: '/bin/bash -c "PATH [ARG]..."'

Examples:
  govc guest.run -vm $name ifconfig
  govc guest.run -vm $name ifconfig eth0
  cal | govc guest.run -vm $name -d - cat
  govc guest.run -vm $name -d "hello $USER" cat
  govc guest.run -vm $name curl -s :invalid: || echo $? # exit code 6
  govc guest.run -vm $name -e FOO=bar -e BIZ=baz -C /tmp env
  govc guest.run -vm $name -l root:mypassword ntpdate -u pool.ntp.org
  govc guest.run -vm $name powershell C:\\network_refresh.ps1`
}

func (cmd *run) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}
	name := f.Arg(0)

	c, err := cmd.Toolbox(ctx)
	if err != nil {
		return err
	}

	ecmd := &exec.Cmd{
		Path:   name,
		Args:   f.Args()[1:],
		Env:    cmd.vars,
		Dir:    cmd.dir,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	switch cmd.data {
	case "":
	case "-":
		ecmd.Stdin = os.Stdin
	default:
		ecmd.Stdin = bytes.NewBuffer([]byte(cmd.data))
	}

	return c.Run(ctx, ecmd)
}
