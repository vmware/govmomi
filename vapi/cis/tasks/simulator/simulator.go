/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"net/http"
	"net/url"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/cis/tasks"
	vapi "github.com/vmware/govmomi/vapi/simulator"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

type Handler struct {
	URL *url.URL
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL: u,
	}
}

func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(tasks.TasksPath+"/", h.depotsOffline)
	}
}

func (h *Handler) depotsOffline(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		task := make(map[string]string)
		task["status"] = "SUCCEEDED"
		vapi.StatusOK(w, task)
	}
}
