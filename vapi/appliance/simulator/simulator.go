// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/appliance/access/consolecli"
	"github.com/vmware/govmomi/vapi/appliance/access/dcui"
	"github.com/vmware/govmomi/vapi/appliance/access/shell"
	"github.com/vmware/govmomi/vapi/appliance/access/ssh"
	"github.com/vmware/govmomi/vapi/appliance/shutdown"
	vapi "github.com/vmware/govmomi/vapi/simulator"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// Handler implements the Appliance API simulator
type Handler struct {
	URL            *url.URL
	consolecli     consolecli.Access
	dcui           dcui.Access
	ssh            ssh.Access
	shell          shell.Access
	shutdownConfig shutdown.Config
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL:            nil,
		consolecli:     consolecli.Access{Enabled: false},
		dcui:           dcui.Access{Enabled: false},
		ssh:            ssh.Access{Enabled: false},
		shell:          shell.Access{Enabled: false, Timeout: 0},
		shutdownConfig: shutdown.Config{},
	}
}

// Register Appliance Management API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	s.HandleFunc(consolecli.Path, h.consoleCLIAccess)
	s.HandleFunc(dcui.Path, h.dcuiAccess)
	s.HandleFunc(ssh.Path, h.sshAccess)
	s.HandleFunc(shell.Path, h.shellAccess)
	s.HandleFunc(shutdown.Path, h.shutdown)
}

func (h *Handler) decode(r *http.Request, w http.ResponseWriter, val any) bool {
	return Decode(r, w, val)
}

// Decode decodes the request Body into val, returns true on success, otherwise false.
func Decode(request *http.Request, writer http.ResponseWriter, val any) bool {
	defer request.Body.Close()
	err := json.NewDecoder(request.Body).Decode(val)
	if err != nil {
		log.Printf("%s %s: %s", request.Method, request.RequestURI, err)
		return false
	}
	return true
}

func (h *Handler) consoleCLIAccess(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vapi.StatusOK(writer, h.consolecli.Enabled)
	case http.MethodPut:
		var input consolecli.Access
		if h.decode(request, writer, &input) {
			h.consolecli.Enabled = input.Enabled
			writer.WriteHeader(http.StatusNoContent)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		http.NotFound(writer, request)
	}
}

func (h *Handler) dcuiAccess(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vapi.StatusOK(writer, h.dcui.Enabled)
	case http.MethodPut:
		var input dcui.Access
		if h.decode(request, writer, &input) {
			h.dcui.Enabled = input.Enabled
			writer.WriteHeader(http.StatusNoContent)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		http.NotFound(writer, request)
	}
}

func (h *Handler) sshAccess(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vapi.StatusOK(writer, h.ssh.Enabled)
	case http.MethodPut:
		var input ssh.Access
		if h.decode(request, writer, &input) {
			h.ssh.Enabled = input.Enabled
			writer.WriteHeader(http.StatusNoContent)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		http.NotFound(writer, request)
	}
}

func (h *Handler) shellAccess(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vapi.StatusOK(writer, h.shell)
	case http.MethodPut:
		var input shell.Access
		if h.decode(request, writer, &input) {
			h.shell.Enabled = input.Enabled
			h.shell.Timeout = input.Timeout
			writer.WriteHeader(http.StatusNoContent)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		http.NotFound(writer, request)
	}
}

func (h *Handler) shutdown(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.shutdownConfig)
	case http.MethodPost:
		switch r.URL.Query().Get(shutdown.Action) {
		case shutdown.Cancel:
			h.shutdownConfig.ShutdownTime = ""
			h.shutdownConfig.Action = ""
			h.shutdownConfig.Reason = ""
			w.WriteHeader(http.StatusNoContent)
		case shutdown.Reboot:
			var spec shutdown.Spec
			if h.decode(r, w, &spec) {
				h.shutdownConfig.ShutdownTime = time.Now().UTC().Add(time.Duration(spec.Delay) * time.Minute).String()
				h.shutdownConfig.Reason = spec.Reason
				h.shutdownConfig.Action = shutdown.Reboot
				w.WriteHeader(http.StatusNoContent)
			}
		case shutdown.PowerOff:
			var spec shutdown.Spec
			if h.decode(r, w, &spec) {
				h.shutdownConfig.ShutdownTime = time.Now().UTC().Add(time.Duration(spec.Delay) * time.Minute).String()
				h.shutdownConfig.Reason = spec.Reason
				h.shutdownConfig.Action = shutdown.PowerOff
				w.WriteHeader(http.StatusNoContent)
			}
		default:
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}
