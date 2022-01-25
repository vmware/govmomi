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

package pending

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	Path            = "/appliance/update/pending"
	Action          = "action"
	Precheck        = "precheck"
	Stage           = "stage"
	Validate        = "validate"
	Install         = "install"
	StageAndInstall = "stage-and-install"
)

// Manager provides convenience methods to configure appliance logging forwarding.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager with the given client
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Description represents description of the update.
type Description struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// Summary contains the essential information about the update.
type Summary struct {
	Description    Description `json:"description"`
	Name           string      `json:"name"`
	Priority       Priority    `json:"priority"`
	RebootRequired bool        `json:"reboot_required"`
	ReleaseDate    string      `json:"release_date"`
	Severity       Severity    `json:"severity"`
	Size           int         `json:"size"`
	UpdateType     UpdateType  `json:"update_type"`
	Version        string      `json:"version"`
}

// List Checks if new updates are available.
func (m *Manager) List(ctx context.Context, sourceType string) ([]Summary, error) {
	r := m.Resource(Path)

	r.WithParam("source_type", sourceType)
	var s []Summary
	return s, m.Do(ctx, r.Request(http.MethodGet), &s)
}

// Content represents list of the issues addressed since previous/current version
// and also new features/improvements
type Content struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// Eula represents EULA based on what we are actually installing
type Eula struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// UpdateServiceInfo describes a service to be stopped and started
// during the update installation.
type UpdateServiceInfo struct {
	Description Description `json:"description"`
	Service     string      `json:"service"`
}

// Info contains the extended information about the update.
type Info struct {
	Contents       []Content           `json:"contents"`
	Description    Description         `json:"description"`
	Eulas          []Eula              `json:"eulas"`
	KnowledgeBase  string              `json:"knowledge_base"`
	Name           string              `json:"name"`
	Priority       Priority            `json:"priority"`
	RebootRequired bool                `json:"reboot_required"`
	ReleaseDate    string              `json:"release_date"`
	ServicesInfo   []UpdateServiceInfo `json:"services_will_be_stopped"`
	Severity       Severity            `json:"severity"`
	Size           int                 `json:"size"`
	Staged         bool                `json:"staged"`
	UpdateType     UpdateType          `json:"update_type"`
}

// Get Gets update information
func (m *Manager) Get(ctx context.Context, version string) (Info, error) {
	path := Path + "/" + version
	r := m.Resource(path)

	var info Info
	return info, m.Do(ctx, r.Request(http.MethodGet), &info)
}

// Text Label for the item to be used in GUI/CLI
type Text struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// Question describes a item of information that must be provided by the user
// in order to install the update.
type Question struct {
	DataItem    string       `json:"data_item"`
	Description Description  `json:"description"`
	Text        Text         `json:"text"`
	Type        QuestionType `json:"type"`
}

// PrecheckResult structure contains estimates of how long it will take install
// and rollback an update as well as a list of possible warnings and problems
//with installing the update.
type PrecheckResult struct {
	CheckTime      string     `json:"check_time"`
	Questions      []Question `json:"questions"`
	RebootRequired bool       `json:"reboot_required"`
}

// Precheck Runs update precheck
func (m *Manager) Precheck(ctx context.Context, version string) (PrecheckResult, error) {
	path := Path + "/" + version
	r := m.Resource(path)
	r.WithParam(Action, Precheck)

	var result PrecheckResult

	return result, m.Do(ctx, r.Request(http.MethodPost), &result)
}

// Stage Starts staging the appliance update.
// The updates are searched for in the following order: staged, CDROM, URL
func (m *Manager) Stage(ctx context.Context, version string) error {
	path := Path + "/" + version
	r := m.Resource(path)
	r.WithParam(Action, Stage)

	return m.Do(ctx, r.Request(http.MethodPost), nil)
}

// Message represents notification message.
type Message struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// Resolution represents resolution message, if any. Only set for warnings and errors.
type Resolution struct {
	Args           []string `json:"args"`
	DefaultMessage string   `json:"default_message"`
	ID             string   `json:"id"`
}

// Notification structure describes a notification that can be reported by
// the appliance task
type Notification struct {
	ID         string     `json:"id"`
	Message    Message    `json:"message"`
	Resolution Resolution `json:"resolution"`
}

// Notifications structure contains info/warning/error messages that can be
//reported by the appliance task.
type Notifications struct {
	Errors   []Notification
	Warnings []Notification
	Info     []Notification
}

// Validate Validates the user provided data before the update installation.
func (m *Manager) Validate(ctx context.Context, version string, inp map[string]string) (Notifications, error) {
	var userData map[string]map[string]string

	userData = make(map[string]map[string]string)
	userData["user_data"] = inp

	path := Path + "/" + version
	r := m.Resource(path)
	r.WithParam(Action, Validate)

	var notifications Notifications

	return notifications, m.Do(ctx, r.Request(http.MethodPost, userData), notifications)
}

// Install Starts operation of installing the appliance update. Will fail is the update is not staged
func (m *Manager) Install(ctx context.Context, version string, inp map[string]string) error {
	var userData map[string]map[string]string

	userData = make(map[string]map[string]string)
	userData["user_data"] = inp

	path := Path + "/" + version
	r := m.Resource(path)
	r.WithParam(Action, Install)

	return m.Do(ctx, r.Request(http.MethodPost, userData), nil)
}

// StageAndInstall Starts operation of installing the appliance update.
// Will stage update if not already staged The updates are searched for in the following order: staged, CDROM, URL
func (m *Manager) StageAndInstall(ctx context.Context, version string, inp map[string]string) error {
	var userData map[string]map[string]string

	userData = make(map[string]map[string]string)
	userData["user_data"] = inp

	path := Path + "/" + version
	r := m.Resource(path)
	r.WithParam(Action, StageAndInstall)

	return m.Do(ctx, r.Request(http.MethodPost, userData), nil)
}
