// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
