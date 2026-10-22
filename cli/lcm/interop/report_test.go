// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interop

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/cis/tasks"
	"github.com/vmware/govmomi/vapi/lcm"
)

func TestPrintReportNil(t *testing.T) {
	var buf bytes.Buffer
	if err := printReport(&buf, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "No result payload") {
		t.Errorf("expected 'No result payload' message, got: %q", buf.String())
	}
}

func TestPrintReport(t *testing.T) {
	r := &lcm.InteropResult{
		Report: lcm.InteropReport{
			DateCreated: time.Date(2026, 6, 26, 10, 0, 0, 0, time.UTC),
			Rows: []lcm.InteropReportRow{
				{
					TargetComponent: lcm.InteropTargetComponent{
						Product: lcm.InteropProductInfo{DisplayName: "vCenter Server"},
						Version: "8.0.3",
					},
					Summary: lcm.InteropSummary{
						CompatibleCount:   3,
						IncompatibleCount: 1,
						UnknownCount:      2,
					},
					AssociatedComponents: []lcm.InteropAssociatedComponent{
						{
							Component: lcm.InteropComponentInstance{
								Product: lcm.InteropProductInfo{DisplayName: "NSX"},
								Name:    "nsx",
							},
							Status:             "COMPATIBLE",
							CompatibleReleases: []lcm.InteropReleaseInfo{{Version: "4.1"}},
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := printReport(&buf, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()

	for _, want := range []string{
		"2026-06-26",
		"vCenter Server",
		"8.0.3",
		"NSX",
		"COMPATIBLE",
		"4.1",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestPrintReportWithIssues(t *testing.T) {
	resolution := "Upgrade to version 2.0"
	r := &lcm.InteropResult{
		Report: lcm.InteropReport{
			DateCreated: time.Now(),
			Rows:        []lcm.InteropReportRow{},
		},
		Issues: &lcm.InteropNotifications{
			Errors: []lcm.InteropNotification{
				{
					Id:      "err-001",
					Message: lcm.InteropLocalizableMessage{DefaultMessage: "Component X is incompatible"},
				},
			},
			Warnings: []lcm.InteropNotification{
				{
					Id:      "warn-001",
					Message: lcm.InteropLocalizableMessage{DefaultMessage: "Component Y status unknown"},
					Resolution: &lcm.InteropLocalizableMessage{
						DefaultMessage: resolution,
					},
				},
			},
			Info: []lcm.InteropNotification{
				{
					Id:      "info-001",
					Message: lcm.InteropLocalizableMessage{DefaultMessage: "Validation completed"},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := printReport(&buf, r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()

	for _, want := range []string{
		"[ERROR]", "Component X is incompatible", "err-001",
		"[WARNING]", "Component Y status unknown", "warn-001", resolution,
		"[INFO]", "Validation completed", "info-001",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestPrintIssuesNoResolution(t *testing.T) {
	notifications := []lcm.InteropNotification{
		{
			Id:      "n-001",
			Message: lcm.InteropLocalizableMessage{DefaultMessage: "something happened"},
		},
	}

	var buf bytes.Buffer
	printIssues(&buf, "WARNING", notifications)
	out := buf.String()

	if !strings.Contains(out, "[WARNING]") {
		t.Errorf("missing level prefix: %q", out)
	}
	if !strings.Contains(out, "something happened") {
		t.Errorf("missing message: %q", out)
	}
	if strings.Contains(out, "Resolution") {
		t.Errorf("unexpected Resolution line when none set: %q", out)
	}
}

func TestReportOutputMarshalJSON(t *testing.T) {
	csvReport := "col1,col2\nval1,val2"
	r := &lcm.InteropResult{
		Report: lcm.InteropReport{
			DateCreated: time.Date(2026, 6, 26, 0, 0, 0, 0, time.UTC),
			Rows:        []lcm.InteropReportRow{},
		},
		CsvReport: &csvReport,
		Issues: &lcm.InteropNotifications{
			Errors: []lcm.InteropNotification{
				{Id: "e1", Message: lcm.InteropLocalizableMessage{DefaultMessage: "bad"}},
			},
		},
	}

	o := &reportOutput{result: r}
	data, err := o.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}

	var decoded lcm.InteropResult
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if decoded.CsvReport == nil || *decoded.CsvReport != csvReport {
		t.Errorf("csv_report not preserved: got %v", decoded.CsvReport)
	}
	if decoded.Issues == nil || len(decoded.Issues.Errors) != 1 {
		t.Errorf("issues not preserved: got %+v", decoded.Issues)
	}
	if decoded.Issues.Errors[0].Id != "e1" {
		t.Errorf("issue id not preserved: got %q", decoded.Issues.Errors[0].Id)
	}
	if len(decoded.Report.Rows) != 0 {
		t.Errorf("expected empty rows, got %d", len(decoded.Report.Rows))
	}
}

func TestReportOutputMarshalJSONNilResult(t *testing.T) {
	o := &reportOutput{result: nil}
	data, err := o.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if string(data) != "null" {
		t.Errorf("expected null for nil result, got %q", data)
	}
}

func TestCheckDescription(t *testing.T) {
	cmd := &check{}
	if !strings.Contains(cmd.Description(), "Submit") {
		t.Errorf("unexpected description: %q", cmd.Description())
	}
}

func TestCheckRunMissingVersion(t *testing.T) {
	cmd := &check{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	err := cmd.Run(context.Background(), f)
	if err == nil || err.Error() != "-version is required" {
		t.Errorf("got err %v, want '-version is required'", err)
	}
}

func TestReportDescription(t *testing.T) {
	cmd := &report{}
	if !strings.Contains(cmd.Description(), "Fetch") {
		t.Errorf("unexpected description: %q", cmd.Description())
	}
}

func TestReportUsage(t *testing.T) {
	cmd := &report{}
	if cmd.Usage() != "TASK_ID" {
		t.Errorf("got usage %q, want 'TASK_ID'", cmd.Usage())
	}
}

func TestReportRunMissingArgs(t *testing.T) {
	cmd := &report{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	err := cmd.Run(context.Background(), f)
	if err != flag.ErrHelp {
		t.Errorf("got err %v, want flag.ErrHelp", err)
	}
}

func TestReportWrite(t *testing.T) {
	r := &lcm.InteropResult{
		Report: lcm.InteropReport{
			DateCreated: time.Date(2026, 6, 26, 0, 0, 0, 0, time.UTC),
			Rows:        []lcm.InteropReportRow{},
		},
	}
	o := &reportOutput{result: r}
	var buf bytes.Buffer
	if err := o.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if !strings.Contains(buf.String(), "2026-06-26") {
		t.Errorf("output missing date: %q", buf.String())
	}
}

func TestPrintAssociatedComponentsLinked(t *testing.T) {
	version := "3.0"
	components := []lcm.InteropAssociatedComponent{
		{
			Component: lcm.InteropComponentInstance{
				Product: lcm.InteropProductInfo{DisplayName: "NSX"},
				Version: &version,
			},
			Status: "COMPATIBLE",
			LinkedComponents: []lcm.InteropAssociatedComponent{
				{
					Component: lcm.InteropComponentInstance{
						Product: lcm.InteropProductInfo{DisplayName: "NSX Edge"},
					},
					Status: "COMPATIBLE",
				},
			},
		},
	}

	var buf bytes.Buffer
	printAssociatedComponents(&buf, components, "  ")
	out := buf.String()

	if !strings.Contains(out, "NSX 3.0") {
		t.Errorf("missing parent component: %q", out)
	}
	if !strings.Contains(out, "NSX Edge") {
		t.Errorf("missing linked component: %q", out)
	}
}

func TestNewManagerHelper(t *testing.T) {
	f := &flags.ClientFlag{}
	f.Session.URL = &url.URL{
		Host: "vcenter.example.com",
		User: url.UserPassword("admin", "secret"),
	}
	f.Session.Insecure = true
	m := newManager(f)
	if m == nil {
		t.Fatal("expected non-nil manager from newManager")
	}
}

func TestCheckRegister(t *testing.T) {
	cmd := &check{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	cmd.Register(context.Background(), f)
	if cmd.ClientFlag == nil {
		t.Fatal("ClientFlag should be non-nil after Register")
	}
	if cmd.OutputFlag == nil {
		t.Fatal("OutputFlag should be non-nil after Register")
	}
	if f.Lookup("version") == nil {
		t.Fatal("expected -version flag to be registered")
	}
	if f.Lookup("product") == nil {
		t.Fatal("expected -product flag to be registered")
	}
}

func TestReportRegister(t *testing.T) {
	cmd := &report{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	cmd.Register(context.Background(), f)
	if cmd.ClientFlag == nil {
		t.Fatal("ClientFlag should be non-nil after Register")
	}
	if cmd.OutputFlag == nil {
		t.Fatal("OutputFlag should be non-nil after Register")
	}
}

// newTLSTestServer creates an HTTPS test server. newManager always builds
// "https://host", so we use TLS + Insecure to let the test client through.
func newTLSTestServer(t *testing.T, h http.HandlerFunc) (*httptest.Server, *url.URL) {
	t.Helper()
	ts := httptest.NewTLSServer(h)
	t.Cleanup(ts.Close)
	u, _ := url.Parse(ts.URL)
	u.User = url.UserPassword("user", "pass")
	return ts, u
}

func registeredCheck(t *testing.T, tsURL *url.URL) (*check, *flag.FlagSet) {
	t.Helper()
	cmd := &check{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	cmd.Register(context.Background(), f)
	cmd.ClientFlag.Session.URL = tsURL
	cmd.ClientFlag.Session.Insecure = true
	return cmd, f
}

func registeredReport(t *testing.T, tsURL *url.URL) (*report, *flag.FlagSet) {
	t.Helper()
	cmd := &report{}
	f := flag.NewFlagSet("", flag.ContinueOnError)
	cmd.Register(context.Background(), f)
	cmd.ClientFlag.Session.URL = tsURL
	cmd.ClientFlag.Session.Insecure = true
	return cmd, f
}

func TestCheckRunPrintsTaskID(t *testing.T) {
	const wantTaskID = "task-check-run"
	_, tsURL := newTLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(wantTaskID)
	})

	cmd, f := registeredCheck(t, tsURL)
	cmd.version = "8.0.3"

	if err := cmd.Run(context.Background(), f); err != nil {
		t.Fatalf("Run: %v", err)
	}
}

func TestCheckRunWait(t *testing.T) {
	const wantTaskID = "task-wait-run"
	calls := 0
	_, tsURL := newTLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			_ = json.NewEncoder(w).Encode(wantTaskID)
			return
		}
		calls++
		info := tasks.Info{}
		if calls < 2 {
			info.Status = tasks.Running
		} else {
			info.Status = tasks.Succeeded
			result, _ := json.Marshal(lcm.InteropResult{
				Report: lcm.InteropReport{DateCreated: time.Now(), Rows: []lcm.InteropReportRow{}},
			})
			info.Result = result
		}
		_ = json.NewEncoder(w).Encode(info)
	})

	cmd, f := registeredCheck(t, tsURL)
	cmd.version = "8.0.3"
	cmd.wait = true
	cmd.ClientFlag.Session.Insecure = true

	if err := cmd.Run(context.Background(), f); err != nil {
		t.Fatalf("Run -wait: %v", err)
	}
}

func TestReportRunFetchesTask(t *testing.T) {
	result, _ := json.Marshal(lcm.InteropResult{
		Report: lcm.InteropReport{DateCreated: time.Now(), Rows: []lcm.InteropReportRow{}},
	})
	info := tasks.Info{Status: tasks.Succeeded, Result: result}

	_, tsURL := newTLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	})

	cmd, f := registeredReport(t, tsURL)
	if err := f.Parse([]string{"task-123"}); err != nil {
		t.Fatal(err)
	}

	if err := cmd.Run(context.Background(), f); err != nil {
		t.Fatalf("Run: %v", err)
	}
}
