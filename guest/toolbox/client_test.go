// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox_test

import (
	"context"
	"errors"
	"os/exec"
	"slices"
	"testing"
	"time"

	"github.com/vmware/govmomi/guest/toolbox"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

const testPid = 42

// guestOps fakes the StartProgramInGuest and ListProcessesInGuest responses,
// as vcsim can only start guest processes in a docker container.
// Each ListProcessesInGuest call pops the next entry from lists,
// and returns an empty list once lists is exhausted.
// The pids of each ListProcessesInGuest request are recorded in polled.
type guestOps struct {
	soap.RoundTripper

	lists  [][]types.GuestProcessInfo
	polled [][]int64
}

func (g *guestOps) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	switch req := req.(type) {
	case *methods.StartProgramInGuestBody:
		res.(*methods.StartProgramInGuestBody).Res = &types.StartProgramInGuestResponse{
			Returnval: testPid,
		}
		return nil
	case *methods.ListProcessesInGuestBody:
		g.polled = append(g.polled, req.Req.Pids)

		var procs []types.GuestProcessInfo
		if len(g.lists) != 0 {
			procs, g.lists = g.lists[0], g.lists[1:]
		}

		res.(*methods.ListProcessesInGuestBody).Res = &types.ListProcessesInGuestResponse{
			Returnval: procs,
		}
		return nil
	}

	return g.RoundTripper.RoundTrip(ctx, req, res)
}

func TestClientRun(t *testing.T) {
	end := types.NewTime(time.Now())

	tests := []struct {
		name     string
		lists    [][]types.GuestProcessInfo
		polls    int   // expected number of ListProcessesInGuest calls
		wantErr  error // expected error, matched with errors.Is
		wantCode int   // expected exit code, if non-zero
	}{
		{
			name:  "exit 0",
			lists: [][]types.GuestProcessInfo{{{Pid: testPid, EndTime: end}}},
			polls: 1,
		},
		{
			name:     "exit 3",
			lists:    [][]types.GuestProcessInfo{{{Pid: testPid, EndTime: end, ExitCode: 3}}},
			polls:    1,
			wantCode: 3,
		},
		{
			name: "poll until exit",
			lists: [][]types.GuestProcessInfo{
				{{Pid: testPid}},
				{{Pid: testPid, EndTime: end}},
			},
			polls: 2,
		},
		{
			// The guest no longer knows the pid, such as after a vmtoolsd restart.
			// ListProcessesInGuest succeeds with an empty list in this case.
			name:    "process not found",
			polls:   1,
			wantErr: toolbox.ErrProcessNotFound,
		},
		{
			name: "process not found while polling",
			lists: [][]types.GuestProcessInfo{
				{{Pid: testPid}},
			},
			polls:   2,
			wantErr: toolbox.ErrProcessNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			simulator.Test(func(ctx context.Context, c *vim25.Client) {
				vm := simulator.Map(ctx).Any("VirtualMachine").(*simulator.VirtualMachine)

				ops := &guestOps{RoundTripper: c.RoundTripper, lists: test.lists}
				c.RoundTripper = ops

				tools, err := toolbox.NewClient(ctx, c, vm, &types.NamePasswordAuthentication{
					Username: "user",
					Password: "pass",
				})
				if err != nil {
					t.Fatal(err)
				}

				err = tools.Run(ctx, &exec.Cmd{Path: "/bin/true"})

				if len(ops.polled) != test.polls {
					t.Errorf("ListProcessesInGuest called %d times, expected %d", len(ops.polled), test.polls)
				}
				for _, pids := range ops.polled {
					if !slices.Equal(pids, []int64{testPid}) {
						t.Errorf("ListProcessesInGuest pids=%v, expected [%d]", pids, testPid)
					}
				}

				var ee interface{ ExitCode() int }
				isExit := errors.As(err, &ee)

				switch {
				case test.wantErr != nil:
					if !errors.Is(err, test.wantErr) {
						t.Fatalf("error=%v, expected %v", err, test.wantErr)
					}
					if isExit {
						t.Errorf("unexpected exit error: %s", err)
					}
				case test.wantCode != 0:
					if !isExit {
						t.Fatalf("expected an exit error, got: %v", err)
					}
					if ee.ExitCode() != test.wantCode {
						t.Errorf("exit code=%d, expected %d", ee.ExitCode(), test.wantCode)
					}
				default:
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}
				}
			})
		})
	}
}
