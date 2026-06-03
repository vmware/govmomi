// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tasks_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

// newTestManager builds a Manager backed by a fake HTTP server.
// The server returns each status in seq on successive GET calls.
func newTestManager(t *testing.T, seq []tasks.Status) (*tasks.Manager, func()) {
	t.Helper()
	var call atomic.Int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idx := int(call.Add(1)) - 1
		if idx >= len(seq) {
			t.Errorf("unexpected GET call #%d; only %d responses configured", idx+1, len(seq))
			http.Error(w, "too many calls", http.StatusInternalServerError)
			return
		}
		status := seq[idx]
		info := tasks.Info{Status: status}
		if status == tasks.Failed {
			info.Error = rest.Error{ErrorType: rest.ErrInternalServer}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	}))

	u, _ := url.Parse(ts.URL)
	soapClient := soap.NewClient(u, true)
	vc := &vim25.Client{Client: soapClient}
	rc := rest.NewClient(vc)
	mgr := tasks.NewManagerWithCustomInterval(rc, 1)
	return mgr, ts.Close
}

func TestWaitForRunningOrTerminalState(t *testing.T) {
	testcases := []struct {
		name       string
		seq        []tasks.Status
		ctx        func() (context.Context, context.CancelFunc)
		wantStatus tasks.Status
		wantErr    bool
	}{
		{
			name:       "fast task already succeeded",
			seq:        []tasks.Status{tasks.Succeeded},
			wantStatus: tasks.Succeeded,
		},
		{
			name:       "fast task already failed",
			seq:        []tasks.Status{tasks.Failed},
			wantStatus: tasks.Failed,
			wantErr:    true,
		},
		{
			name:       "fast task already blocked",
			seq:        []tasks.Status{tasks.Blocked},
			wantStatus: tasks.Blocked,
		},
		{
			name:       "pending then running",
			seq:        []tasks.Status{tasks.Pending, tasks.Running},
			wantStatus: tasks.Running,
		},
		{
			name: "context cancelled while pending",
			seq:  []tasks.Status{tasks.Pending, tasks.Pending, tasks.Pending},
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mgr, cleanup := newTestManager(t, tc.seq)
			defer cleanup()

			ctx := context.Background()
			var cancel context.CancelFunc
			if tc.ctx != nil {
				ctx, cancel = tc.ctx()
				defer cancel()
			}

			info, err := mgr.WaitForRunningOrTerminalState(ctx, "task-1")
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v, wantErr=%v", err, tc.wantErr)
			}
			if !tc.wantErr && info.Status != tc.wantStatus {
				t.Errorf("status: got %q, want %q", info.Status, tc.wantStatus)
			}
		})
	}
}

// TestWaitForRunningOrError documents the existing limitation: a task that
// reaches SUCCEEDED before the first poll causes the method to block
// indefinitely, because SUCCEEDED satisfies neither Running nor Failed.
func TestWaitForRunningOrError(t *testing.T) {
	testcases := []struct {
		name       string
		seq        []tasks.Status
		ctx        func() (context.Context, context.CancelFunc)
		wantStatus tasks.Status
		wantErr    bool
	}{
		{
			name:       "running",
			seq:        []tasks.Status{tasks.Running},
			wantStatus: tasks.Running,
		},
		{
			name:       "already failed",
			seq:        []tasks.Status{tasks.Failed},
			wantStatus: tasks.Failed,
			wantErr:    true,
		},
		{
			name: "succeeded causes hang until context expires",
			seq:  []tasks.Status{tasks.Succeeded, tasks.Succeeded, tasks.Succeeded},
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mgr, cleanup := newTestManager(t, tc.seq)
			defer cleanup()

			ctx := context.Background()
			var cancel context.CancelFunc
			if tc.ctx != nil {
				ctx, cancel = tc.ctx()
				defer cancel()
			}

			info, err := mgr.WaitForRunningOrError(ctx, "task-1")
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v, wantErr=%v", err, tc.wantErr)
			}
			if !tc.wantErr && info.Status != tc.wantStatus {
				t.Errorf("status: got %q, want %q", info.Status, tc.wantStatus)
			}
		})
	}
}

func TestWaitForCompletion(t *testing.T) {
	testcases := []struct {
		name       string
		seq        []tasks.Status
		wantStatus tasks.Status
		wantErr    bool
	}{
		{
			name:       "succeeds immediately",
			seq:        []tasks.Status{tasks.Succeeded},
			wantStatus: tasks.Succeeded,
		},
		{
			name:       "running then succeeded",
			seq:        []tasks.Status{tasks.Running, tasks.Succeeded},
			wantStatus: tasks.Succeeded,
		},
		{
			name:    "fails immediately",
			seq:     []tasks.Status{tasks.Failed},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mgr, cleanup := newTestManager(t, tc.seq)
			defer cleanup()

			info, err := mgr.WaitForCompletion(context.Background(), "task-1")
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v, wantErr=%v", err, tc.wantErr)
			}
			if !tc.wantErr && info.Status != tc.wantStatus {
				t.Errorf("status: got %q, want %q", info.Status, tc.wantStatus)
			}
		})
	}
}
