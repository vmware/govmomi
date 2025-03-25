// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type ps struct {
	*flags.OutputFlag
	*GuestFlag

	every bool
	exit  bool
	wait  bool

	pids pidSelector
	uids uidSelector
}

type pidSelector []int64

func (s *pidSelector) String() string {
	return fmt.Sprint(*s)
}

func (s *pidSelector) Set(value string) error {
	v, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}
	*s = append(*s, v)
	return nil
}

type uidSelector map[string]bool

func (s uidSelector) String() string {
	return ""
}

func (s uidSelector) Set(value string) error {
	s[value] = true
	return nil
}

func init() {
	cli.Register("guest.ps", &ps{})
}

func (cmd *ps) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.GuestFlag, ctx = newGuestProcessFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	cmd.uids = make(map[string]bool)
	f.BoolVar(&cmd.every, "e", false, "Select all processes")
	f.BoolVar(&cmd.exit, "x", false, "Output exit time and code")
	f.BoolVar(&cmd.wait, "X", false, "Wait for process to exit")
	f.Var(&cmd.pids, "p", "Select by process ID")
	f.Var(&cmd.uids, "U", "Select by process UID")
}

func (cmd *ps) Process(ctx context.Context) error {
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ps) Description() string {
	return `List processes in VM.

By default, unless the '-e', '-p' or '-U' flag is specified, only processes owned
by the '-l' flag user are displayed.

The '-x' and '-X' flags only apply to processes started by vmware-tools,
such as those started with the govc guest.start command.

Examples:
  govc guest.ps -vm $name
  govc guest.ps -vm $name -e
  govc guest.ps -vm $name -p 12345
  govc guest.ps -vm $name -U root`
}

func running(procs []types.GuestProcessInfo) bool {
	for _, p := range procs {
		if p.EndTime == nil {
			return true
		}
	}
	return false
}

func (cmd *ps) list(ctx context.Context) ([]types.GuestProcessInfo, error) {
	m, err := cmd.ProcessManager()
	if err != nil {
		return nil, err
	}

	auth := cmd.Auth()

	for {
		procs, err := m.ListProcesses(ctx, auth, cmd.pids)
		if err != nil {
			return nil, err
		}

		if cmd.wait && running(procs) {
			<-time.After(time.Second)
			continue
		}

		return procs, nil
	}
}

func (cmd *ps) Run(ctx context.Context, f *flag.FlagSet) error {
	procs, err := cmd.list(ctx)
	if err != nil {
		return err
	}

	r := &psResult{cmd, procs}

	return cmd.WriteResult(r)
}

type psResult struct {
	cmd         *ps
	ProcessInfo []types.GuestProcessInfo `json:"processInfo"`
}

func (r *psResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%s\t%s\t%s", "UID", "PID", "STIME")
	if r.cmd.exit {
		fmt.Fprintf(tw, "\t%s\t%s", "XTIME", "XCODE")
	}
	fmt.Fprint(tw, "\tCMD\n")

	if len(r.cmd.pids) != 0 {
		r.cmd.every = true
	}

	if !r.cmd.every && len(r.cmd.uids) == 0 {
		r.cmd.uids[r.cmd.auth.Username] = true
	}

	for _, p := range r.ProcessInfo {
		if r.cmd.every || r.cmd.uids[p.Owner] {
			fmt.Fprintf(tw, "%s\t%d\t%s", p.Owner, p.Pid, p.StartTime.Format("15:04"))
			if r.cmd.exit {
				etime := "-"
				ecode := "-"
				if p.EndTime != nil {
					etime = p.EndTime.Format("15:04")
					ecode = strconv.Itoa(int(p.ExitCode))
				}
				fmt.Fprintf(tw, "\t%s\t%s", etime, ecode)
			}
			fmt.Fprintf(tw, "\t%s\n", strings.TrimSpace(p.CmdLine))
		}
	}

	return tw.Flush()
}
