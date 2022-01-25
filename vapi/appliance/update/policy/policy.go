/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package policy

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const Path = "/appliance/update/policy"

// Manager provides convenience methods to set/get background check for the new updates.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// CheckSchedule defines weekday and time the automatic check for new updates
// will be run.
type CheckSchedule struct {
	Day    string `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
}

// Info represents policy for the appliance update.
type Info struct {
	AutoStage        bool            `json:"auto_stage"`
	AutoUpdate       bool            `json:"auto_update"`
	CertificateCheck bool            `json:"certificate_check"`
	CheckSchedule    []CheckSchedule `json:"check_schedule"`
	CustomURL        string          `json:"custom_URL"`
	DefaultURL       string          `json:"default_URL"`
	ManualControl    bool            `json:"manual_control"`
	Username         string          `json:"username"`
}

//Get Gets the automatic update checking and staging policy.
func (m *Manager) Get(ctx context.Context) (Info, error) {
	r := m.Resource(Path)
	var info Info

	return info, m.Do(ctx, r.Request(http.MethodGet), &info)
}

type PolicyConfig struct {
	Config `json:"policy"`
}

// Config contains the policy for the appliance update.
type Config struct {
	AutoStage        bool            `json:"auto_stage"`
	CertificateCheck bool            `json:"certificate_check"`
	CheckSchedule    []CheckSchedule `json:"check_schedule"`
	CustomURL        string          `json:"custom_URL"`
	Password         string          `json:"password"`
	Username         string          `json:"username"`
}

// Set Sets the automatic update checking and staging policy.
func (m *Manager) Set(ctx context.Context, input Config) error {
	r := m.Resource(Path)
	p := PolicyConfig{
		Config: input,
	}

	return m.Do(ctx, r.Request(http.MethodPut, p), nil)
}
