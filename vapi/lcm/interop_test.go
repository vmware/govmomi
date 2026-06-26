// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package lcm_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/lcm"
)

// newTestManager builds a Manager pointed at ts with no polling delay.
func newTestManager(ts *httptest.Server) *lcm.Manager {
	return lcm.NewManager(ts.URL, "user", "pass", ts.Client()).
		WithPollingInterval(0)
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	return b
}

func TestCreateInteropTask(t *testing.T) {
	const wantTaskID = "interop-task-abc"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("got method %s, want POST", r.Method)
		}
		if !strings.HasPrefix(r.URL.Path, "/lcm/rest/vcenter/lcm/interop") {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		if r.URL.Query().Get("vmw-task") != "true" {
			t.Errorf("missing vmw-task=true query param")
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != "user" || pass != "pass" {
			t.Errorf("unexpected auth: user=%q pass=%q ok=%v", user, pass, ok)
		}

		var body struct {
			Spec struct {
				Components []struct {
					ProductName string  `json:"product_name"`
					Version     *string `json:"version,omitempty"`
					Identifier  *string `json:"identifier,omitempty"`
				} `json:"components"`
			} `json:"spec"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decoding request body: %v", err)
		}
		if len(body.Spec.Components) != 1 {
			t.Errorf("expected 1 component, got %d", len(body.Spec.Components))
		}
		if body.Spec.Components[0].ProductName != "com.vmware.vcenter" {
			t.Errorf("unexpected product_name %q", body.Spec.Components[0].ProductName)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(wantTaskID)
	}))
	defer ts.Close()

	m := newTestManager(ts)
	taskID, err := m.CreateInteropTask(context.Background(), lcm.InteropComponent{
		ProductName: "com.vmware.vcenter",
		Version:     "8.0.3",
	})
	if err != nil {
		t.Fatalf("CreateInteropTask: %v", err)
	}
	if taskID != wantTaskID {
		t.Errorf("got taskID %q, want %q", taskID, wantTaskID)
	}
}

func TestGetTask(t *testing.T) {
	const taskID = "interop-task-xyz"

	want := tasks.Info{
		Status:   tasks.Succeeded,
		Progress: tasks.Progress{Total: 100, Completed: 100},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("got method %s, want GET", r.Method)
		}
		wantPath := "/lcm/rest/cis/tasks/" + taskID
		if r.URL.Path != wantPath {
			t.Errorf("got path %q, want %q", r.URL.Path, wantPath)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	}))
	defer ts.Close()

	m := newTestManager(ts)
	got, err := m.GetTask(context.Background(), taskID)
	if err != nil {
		t.Fatalf("GetTask: %v", err)
	}
	if got.Status != tasks.Succeeded {
		t.Errorf("got status %q, want %q", got.Status, tasks.Succeeded)
	}
}

func TestWaitForCompletion(t *testing.T) {
	calls := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		info := tasks.Info{}
		if calls < 3 {
			info.Status = tasks.Running
			info.Progress = tasks.Progress{Total: 100, Completed: uint64(calls * 30)}
		} else {
			info.Status = tasks.Succeeded
			info.Progress = tasks.Progress{Total: 100, Completed: 100}
			info.Result = mustJSON(t, lcm.InteropResult{
				Report: lcm.InteropReport{
					DateCreated: time.Now(),
					Rows:        []lcm.InteropReportRow{},
				},
			})
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	}))
	defer ts.Close()

	m := newTestManager(ts)
	info, err := m.WaitForCompletion(context.Background(), "task-999")
	if err != nil {
		t.Fatalf("WaitForCompletion: %v", err)
	}
	if info.Status != tasks.Succeeded {
		t.Errorf("got status %q, want %q", info.Status, tasks.Succeeded)
	}
	if calls < 3 {
		t.Errorf("expected at least 3 polls, got %d", calls)
	}
}

func TestWaitForCompletionContextCancel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		info := tasks.Info{Status: tasks.Running}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	m := newTestManager(ts)
	_, err := m.WaitForCompletion(ctx, "task-999")
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}

func TestParseResult(t *testing.T) {
	note := "upgrade required"
	interopResult := lcm.InteropResult{
		Report: lcm.InteropReport{
			DateCreated: time.Date(2026, 6, 26, 0, 0, 0, 0, time.UTC),
			Rows: []lcm.InteropReportRow{
				{
					TargetComponent: lcm.InteropTargetComponent{
						Product: lcm.InteropProductInfo{DisplayName: "vCenter Server", Name: "vcenter"},
						Version: "8.0.3",
					},
					Summary: lcm.InteropSummary{CompatibleCount: 5, IncompatibleCount: 1},
				},
			},
		},
		Issues: &lcm.InteropNotifications{
			Warnings: []lcm.InteropNotification{
				{
					Id:      "warn-001",
					Message: lcm.InteropLocalizableMessage{DefaultMessage: "check compatibility"},
					Resolution: &lcm.InteropLocalizableMessage{
						DefaultMessage: note,
					},
				},
			},
		},
	}

	info := &tasks.Info{Result: mustJSON(t, interopResult)}

	got, err := lcm.ParseResult(info)
	if err != nil {
		t.Fatalf("ParseResult: %v", err)
	}
	if got == nil {
		t.Fatal("expected non-nil result")
	}
	if len(got.Report.Rows) != 1 {
		t.Errorf("got %d rows, want 1", len(got.Report.Rows))
	}
	if got.Report.Rows[0].Summary.CompatibleCount != 5 {
		t.Errorf("got compatible_count %d, want 5", got.Report.Rows[0].Summary.CompatibleCount)
	}
	if got.Issues == nil || len(got.Issues.Warnings) != 1 {
		t.Errorf("expected 1 warning")
	}
	if got.Issues.Warnings[0].Resolution.DefaultMessage != note {
		t.Errorf("got resolution %q, want %q", got.Issues.Warnings[0].Resolution.DefaultMessage, note)
	}
}

func TestParseResultEmpty(t *testing.T) {
	info := &tasks.Info{}
	got, err := lcm.ParseResult(info)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil result for empty task, got %+v", got)
	}
}

func TestCreateInteropTask_ServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	m := newTestManager(ts)
	_, err := m.CreateInteropTask(context.Background(), lcm.InteropComponent{
		ProductName: "com.vmware.vcenter",
		Version:     "8.0.3",
	})
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
}

func TestCreateInteropTaskWithIdentifier(t *testing.T) {
	const wantID = "inst-abc-123"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Spec struct {
				Components []struct {
					Identifier *string `json:"identifier,omitempty"`
				} `json:"components"`
			} `json:"spec"`
		}
		_ = json.NewDecoder(r.Body).Decode(&body)
		if len(body.Spec.Components) == 0 || body.Spec.Components[0].Identifier == nil {
			t.Errorf("expected identifier in request body")
		} else if *body.Spec.Components[0].Identifier != wantID {
			t.Errorf("got identifier %q, want %q", *body.Spec.Components[0].Identifier, wantID)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode("task-with-id")
	}))
	defer ts.Close()

	m := newTestManager(ts)
	_, err := m.CreateInteropTask(context.Background(), lcm.InteropComponent{
		ProductName: "com.vmware.vcenter",
		Version:     "8.0.3",
		Identifier:  wantID,
	})
	if err != nil {
		t.Fatalf("CreateInteropTask: %v", err)
	}
}

func TestNewInsecureManager(t *testing.T) {
	m := lcm.NewInsecureManager("https://vcenter.example.com", "admin", "secret")
	if m == nil {
		t.Fatal("expected non-nil manager from NewInsecureManager")
	}
}

func TestWaitForCompletionFailedTask(t *testing.T) {
	const failedJSON = `{"status":"FAILED","progress":{"total":100,"completed":50},` +
		`"error":{"error_type":"com.vmware.vapi.std.errors.error"}}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(failedJSON))
	}))
	defer ts.Close()

	m := newTestManager(ts)
	_, err := m.WaitForCompletion(context.Background(), "task-fail")
	if err == nil {
		t.Fatal("expected error for failed task, got nil")
	}
}

func TestWaitForCompletionContextCancelDuringPoll(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cancel() // cancel before returning so ctx.Done fires in the select
		info := tasks.Info{Status: tasks.Running}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	}))
	defer ts.Close()

	// Long polling interval so time.After loses to the already-closed ctx.Done.
	m := lcm.NewManager(ts.URL, "user", "pass", ts.Client()).
		WithPollingInterval(time.Hour)

	_, err := m.WaitForCompletion(ctx, "task-999")
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}

func TestGetTaskNetworkError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	m := newTestManager(ts)
	ts.Close() // close before the request is made

	_, err := m.GetTask(context.Background(), "task-123")
	if err == nil {
		t.Fatal("expected network error after server close, got nil")
	}
}

func TestGetTaskInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not-valid-json"))
	}))
	defer ts.Close()

	m := newTestManager(ts)
	_, err := m.GetTask(context.Background(), "task-123")
	if err == nil {
		t.Fatal("expected unmarshal error for invalid JSON, got nil")
	}
}

func TestParseResultInvalidJSON(t *testing.T) {
	info := &tasks.Info{Result: []byte("not-valid-json")}
	_, err := lcm.ParseResult(info)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
