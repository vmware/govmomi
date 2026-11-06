// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package simulator implements the vAPI com.vmware.cis.authz.permission service for Global Permissions.
//
// Only the Global Permissions resource (authz.GlobalPermission) is supported. A principal holds at most one Global
// Permission: creating a second one fails with authz.AlreadyExistsGlobal, as vCenter does. Principals are not
// checked against an identity source. Role ids are checked against the simulator's AuthorizationManager.
package simulator

import (
	"context"
	"encoding/json"
	"log"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/cis/authz"
	"github.com/vmware/govmomi/vapi/rest"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	permissionPath = rest.Path + authz.PermissionPath

	errAlreadyExists   = "com.vmware.vapi.std.errors.already_exists"
	errInvalidArgument = "com.vmware.vapi.std.errors.invalid_argument"
	errNotFound        = "com.vmware.vapi.std.errors.not_found"
	errUnsupported     = "com.vmware.vapi.std.errors.unsupported"

	msgInvalidRole = "com.vmware.vcenter.authorization.permissions.invalid_role"
	msgNotFound    = "com.vmware.vcenter.authorization.permissions.not_found"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(r).Register(s, r)
	})
}

// Handler implements the permission service.
// The handlers never log request or response bodies.
type Handler struct {
	Map *simulator.Registry

	mu          sync.Mutex
	permissions map[string]authz.Info
}

// New creates a Handler instance
func New(r *simulator.Registry) *Handler {
	return &Handler{
		Map:         r,
		permissions: make(map[string]authz.Info),
	}
}

// Register the permission service paths with the vapi simulator's http.ServeMux.
// The vapi simulator requires a valid session for these paths, returning 401 otherwise.
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(permissionPath, h.permission)
		s.HandleFunc(permissionPath+"/", h.permissionID)
	}
}

// restError responds with a "/rest" style vAPI error of the given kind.
func restError(w http.ResponseWriter, status int, kind string, messages ...rest.LocalizableMessage) {
	if messages == nil {
		messages = []rest.LocalizableMessage{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	type value struct {
		Messages []rest.LocalizableMessage `json:"messages"`
	}

	err := json.NewEncoder(w).Encode(struct {
		Type  string `json:"type"`
		Value value  `json:"value"`
	}{kind, value{messages}})

	if err != nil {
		log.Panic(err)
	}
}

func action(r *http.Request) string {
	q := r.URL.Query()
	if a := q.Get("~action"); a != "" {
		return a
	}
	return q.Get("action")
}

func (h *Handler) permission(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.mu.Lock()
		ids := slices.Sorted(maps.Keys(h.permissions))
		if ids == nil {
			ids = []string{} // vAPI encodes an empty list as [], not null
		}
		h.mu.Unlock()
		vapi.OK(w, ids)
	case http.MethodPost:
		switch action(r) {
		case "":
			h.create(w, r)
		case "list-detail":
			h.mu.Lock()
			infos := make([]authz.Info, 0, len(h.permissions))
			for _, info := range h.permissions {
				infos = append(infos, info)
			}
			h.mu.Unlock()
			slices.SortFunc(infos, func(a, b authz.Info) int { return strings.Compare(a.ID, b.ID) })
			vapi.OK(w, infos)
		default:
			http.NotFound(w, r)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Spec authz.CreateSpec `json:"create_spec"`
	}
	if !vapi.Decode(r, w, &req) {
		return
	}
	spec := req.Spec

	if spec.ResourceID != authz.GlobalPermission {
		restError(w, http.StatusBadRequest, errUnsupported)
		return
	}

	switch spec.Principal.Type {
	case authz.PrincipalUser, authz.PrincipalGroup:
	default:
		restError(w, http.StatusBadRequest, errInvalidArgument)
		return
	}
	name := spec.Principal.Name()
	if name == "" {
		restError(w, http.StatusBadRequest, errInvalidArgument)
		return
	}

	// vCenter requires exactly one role per permission.
	if len(spec.RoleID) != 1 {
		restError(w, http.StatusBadRequest, errInvalidArgument)
		return
	}
	if !h.roleExists(spec.RoleID[0]) {
		restError(w, http.StatusBadRequest, errInvalidArgument, rest.LocalizableMessage{
			Args:           []string{spec.RoleID[0]},
			DefaultMessage: "Role " + spec.RoleID[0] + " does not exist.",
			ID:             msgInvalidRole,
		})
		return
	}

	id := authz.PermissionID(authz.GlobalPermissionsDoc, spec.Principal)

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.permissions[id]; exists {
		restError(w, http.StatusBadRequest, errAlreadyExists, rest.LocalizableMessage{
			Args:           []string{name},
			DefaultMessage: "Principal " + name + " already has global permission.",
			ID:             authz.AlreadyExistsGlobal,
		})
		return
	}

	h.permissions[id] = authz.Info{
		ID:         id,
		ResourceID: spec.ResourceID,
		Principal:  spec.Principal,
		RoleID:     slices.Clone(spec.RoleID),
		Propagate:  spec.Propagate,
	}

	vapi.OK(w, id)
}

func (h *Handler) permissionID(w http.ResponseWriter, r *http.Request) {
	id, ok := strings.CutPrefix(strings.TrimPrefix(r.URL.Path, permissionPath+"/"), "id:")
	if !ok || id == "" {
		http.NotFound(w, r)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	info, exists := h.permissions[id]
	if !exists {
		restError(w, http.StatusNotFound, errNotFound, rest.LocalizableMessage{
			Args:           []string{id},
			DefaultMessage: "Permission " + id + " not found.",
			ID:             msgNotFound,
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		vapi.OK(w, info)
	case http.MethodDelete:
		delete(h.permissions, id)
		vapi.OK(w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// roleExists reports whether id names a role in the simulator's AuthorizationManager.
func (h *Handler) roleExists(id string) bool {
	roleID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return false
	}

	si := h.Map.Get(vim25.ServiceInstance).(*simulator.ServiceInstance)
	am := h.Map.Get(*si.Content.AuthorizationManager).(*simulator.AuthorizationManager)

	ctx := &simulator.Context{Context: context.Background(), Map: h.Map}
	found := false
	h.Map.WithLock(ctx, am, func() {
		found = slices.ContainsFunc(am.RoleList, func(role types.AuthorizationRole) bool {
			return role.RoleId == int32(roleID)
		})
	})
	return found
}
