// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package lcm

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/vmware/govmomi/vapi/cis/tasks"
)

const (
	interopTasksPath = "/lcm/rest/cis/tasks/"
	interopPath      = "/lcm/rest/vcenter/lcm/interop"
)

// InteropComponent specifies the component to validate interoperability for.
type InteropComponent struct {
	ProductName string
	Version     string
	Identifier  string // optional instance identifier
}

// InteropResult is the decoded payload returned by a completed interop task.
type InteropResult struct {
	Report    InteropReport         `json:"report"`
	Issues    *InteropNotifications `json:"issues,omitempty"`
	CsvReport *string               `json:"csv_report,omitempty"`
}

// InteropNotifications groups info, warning, and error notifications from the API.
type InteropNotifications struct {
	Info     []InteropNotification `json:"info,omitempty"`
	Warnings []InteropNotification `json:"warnings,omitempty"`
	Errors   []InteropNotification `json:"errors,omitempty"`
}

// InteropNotification is a single notification entry returned alongside the report.
type InteropNotification struct {
	Id         string                     `json:"id"`
	Time       *time.Time                 `json:"time,omitempty"`
	Message    InteropLocalizableMessage  `json:"message"`
	Resolution *InteropLocalizableMessage `json:"resolution,omitempty"`
}

// InteropLocalizableMessage carries a human-readable message from the API.
type InteropLocalizableMessage struct {
	Id             string  `json:"id"`
	DefaultMessage string  `json:"default_message"`
	Localized      *string `json:"localized,omitempty"`
}

// InteropReport is the top-level report container.
type InteropReport struct {
	DateCreated time.Time          `json:"date_created"`
	Rows        []InteropReportRow `json:"rows"`
}

// InteropReportRow is one row of the interoperability report.
type InteropReportRow struct {
	TargetComponent      InteropTargetComponent       `json:"target_component"`
	Summary              InteropSummary               `json:"summary"`
	AssociatedComponents []InteropAssociatedComponent `json:"associated_components"`
}

// InteropTargetComponent describes the component being validated.
type InteropTargetComponent struct {
	Identifier *string            `json:"identifier,omitempty"`
	Product    InteropProductInfo `json:"product"`
	Version    string             `json:"version"`
	Name       *string            `json:"name,omitempty"`
}

// InteropProductInfo holds product metadata.
type InteropProductInfo struct {
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Category    string `json:"category"`
}

// InteropSummary holds compatibility counts for a row.
type InteropSummary struct {
	CompatibleCount   int64 `json:"compatible_count"`
	IncompatibleCount int64 `json:"incompatible_count"`
	UnknownCount      int64 `json:"unknown_count"`
}

// InteropAssociatedComponent describes a component's interoperability status.
type InteropAssociatedComponent struct {
	Component          InteropComponentInstance     `json:"component"`
	Status             string                       `json:"status"`
	CompatibleReleases []InteropReleaseInfo         `json:"compatible_releases"`
	LinkedComponents   []InteropAssociatedComponent `json:"linked_components"`
}

// InteropComponentInstance is an instantiated component in the environment.
type InteropComponentInstance struct {
	Identifier string             `json:"identifier"`
	Product    InteropProductInfo `json:"product"`
	Version    *string            `json:"version,omitempty"`
	Name       string             `json:"name"`
	Auto       bool               `json:"auto"`
}

// InteropReleaseInfo describes a compatible release for an associated component.
type InteropReleaseInfo struct {
	Version string  `json:"version"`
	Note    *string `json:"note,omitempty"`
}

// interopRequestBody is the JSON body for POST /lcm/rest/vcenter/lcm/interop.
type interopRequestBody struct {
	Spec interopSpecBody `json:"spec"`
}

type interopSpecBody struct {
	Components []interopComponentBody `json:"components"`
}

type interopComponentBody struct {
	ProductName string  `json:"product_name"`
	Version     *string `json:"version,omitempty"`
	Identifier  *string `json:"identifier,omitempty"`
}

// Manager provides access to the LCM interoperability REST API using a plain
// net/http client with HTTP Basic Auth. The LCM REST endpoint (/lcm/rest/...)
// is a separate service from the vSphere vAPI and does not accept VAPI session tokens.
type Manager struct {
	client          *http.Client
	baseURL         string
	username        string
	password        string
	pollingInterval time.Duration
}

// NewManager creates a Manager. client handles TLS and transport; baseURL is
// the scheme+host of the vCenter (e.g. "https://vcenter.example.com").
// Interop APIs support Basic Auth only.
func NewManager(baseURL, username, password string, client *http.Client) *Manager {
	return &Manager{
		client:          client,
		baseURL:         baseURL,
		username:        username,
		password:        password,
		pollingInterval: time.Second,
	}
}

// WithPollingInterval overrides the interval between WaitForCompletion polls.
func (m *Manager) WithPollingInterval(d time.Duration) *Manager {
	m.pollingInterval = d
	return m
}

// NewInsecureManager is a convenience constructor that creates its own
// http.Client with TLS verification disabled.
func NewInsecureManager(baseURL, username, password string) *Manager {
	return NewManager(baseURL, username, password, &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		},
	})
}

func (m *Manager) lcmURL(path string, params url.Values) string {
	u, _ := url.Parse(m.baseURL + path)
	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}
	return u.String()
}

func (m *Manager) do(ctx context.Context, method, rawURL string, reqBody, resBody any) error {
	var buf bytes.Buffer
	if reqBody != nil {
		if err := json.NewEncoder(&buf).Encode(reqBody); err != nil {
			return err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, rawURL, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(m.username, m.password)

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s: %s: %s", method, rawURL, resp.Status, body)
	}
	if resBody != nil {
		return json.Unmarshal(body, resBody)
	}
	return nil
}

// CreateInteropTask submits an interop validation request and returns the task ID.
func (m *Manager) CreateInteropTask(ctx context.Context, component InteropComponent) (string, error) {
	body := interopRequestBody{
		Spec: interopSpecBody{
			Components: []interopComponentBody{toComponentBody(component)},
		},
	}
	rawURL := m.lcmURL(interopPath, url.Values{"vmw-task": {"true"}})
	var taskID string
	if err := m.do(ctx, http.MethodPost, rawURL, body, &taskID); err != nil {
		return "", err
	}
	return taskID, nil
}

// GetTask fetches the current state of an interop task.
func (m *Manager) GetTask(ctx context.Context, taskID string) (*tasks.Info, error) {
	rawURL := m.lcmURL(interopTasksPath+taskID, nil)
	var info tasks.Info
	if err := m.do(ctx, http.MethodGet, rawURL, nil, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// WaitForCompletion polls GetTask until the task reaches a terminal state.
func (m *Manager) WaitForCompletion(ctx context.Context, taskID string) (*tasks.Info, error) {
	for {
		info, err := m.GetTask(ctx, taskID)
		if err != nil {
			return nil, err
		}
		if info.IsDone() {
			return info, info.Err()
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(m.pollingInterval):
		}
	}
}

// ParseResult extracts the InteropResult from a completed task's result field.
// Returns nil if the result field is empty.
func ParseResult(info *tasks.Info) (*InteropResult, error) {
	if len(info.Result) == 0 {
		return nil, nil
	}
	var r InteropResult
	if err := json.Unmarshal(info.Result, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func toComponentBody(c InteropComponent) interopComponentBody {
	body := interopComponentBody{ProductName: c.ProductName}
	if c.Version != "" {
		v := c.Version
		body.Version = &v
	}
	if c.Identifier != "" {
		id := c.Identifier
		body.Identifier = &id
	}
	return body
}
