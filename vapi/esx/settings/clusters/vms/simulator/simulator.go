// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters/vms"
	vapi "github.com/vmware/govmomi/vapi/simulator"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// Handler implements the EAM vms API simulator.
type Handler struct {
	URL       *url.URL
	Solutions map[string]map[string]vms.SolutionInfo        // clusterID → solutionID → info
	Hooks     map[string]map[string][]vms.LifecycleHookInfo // clusterID → solutionID → hooks
	Tasks     map[string]json.RawMessage                    // taskID → task JSON
	taskCount int
}

// New creates a Handler instance.
func New(u *url.URL) *Handler {
	return &Handler{
		URL:       u,
		Solutions: make(map[string]map[string]vms.SolutionInfo),
		Hooks:     make(map[string]map[string][]vms.LifecycleHookInfo),
		Tasks:     make(map[string]json.RawMessage),
	}
}

// Register adds all VMS API paths to the simulator's HTTP mux.
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc("/api/esx/settings/clusters/{cluster}/vms/solutions", h.solutions)
		s.HandleFunc("/api/esx/settings/clusters/{cluster}/vms/solutions/{solution}", h.solution)
		s.HandleFunc("/api/esx/settings/clusters/{cluster}/vms/lifecycle-hooks", h.hooksCollection)
		s.HandleFunc("/api/esx/settings/clusters/{cluster}/vms/lifecycle-hooks/{solution}", h.hooksSolution)
		s.HandleFunc("/api/esx/settings/clusters/{cluster}/vms/transition/{solution}", h.transition)
		s.HandleFunc("/api/cis/tasks/{task}", h.taskGet)
	}
}

// newTask stores a SUCCEEDED task and returns its ID.
func (h *Handler) newTask() string {
	h.taskCount++
	id := fmt.Sprintf("vms-task-%d", h.taskCount)
	h.Tasks[id] = json.RawMessage(`{"status":"SUCCEEDED","result":null}`)
	return id
}

// clusterSolutions returns (and lazily initialises) the solution map for a cluster.
func (h *Handler) clusterSolutions(cluster string) map[string]vms.SolutionInfo {
	if h.Solutions[cluster] == nil {
		h.Solutions[cluster] = make(map[string]vms.SolutionInfo)
	}
	return h.Solutions[cluster]
}

// solutions handles the solutions collection endpoint.
//
//	GET  /api/esx/settings/clusters/{cluster}/vms/solutions
//	POST /api/esx/settings/clusters/{cluster}/vms/solutions?action=apply&vmw-task=true
//	POST /api/esx/settings/clusters/{cluster}/vms/solutions?action=check-compliance&vmw-task=true
func (h *Handler) solutions(w http.ResponseWriter, r *http.Request) {
	cluster := r.PathValue("cluster")
	action := r.URL.Query().Get("action")

	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, vms.ListResult{Solutions: h.clusterSolutions(cluster)})

	case http.MethodPost:
		switch action {
		case "apply":
			var spec vms.ApplySpec
			if vapi.Decode(r, w, &spec) {
				vapi.StatusOK(w, h.newTask())
			}
		case "check-compliance":
			var spec vms.CheckComplianceFilterSpec
			if vapi.Decode(r, w, &spec) {
				vapi.StatusOK(w, h.newTask())
			}
		default:
			vapi.ApiErrorInvalidArgument(w)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// solution handles the per-solution endpoint.
//
//	GET    /api/esx/settings/clusters/{cluster}/vms/solutions/{solution}
//	PUT    /api/esx/settings/clusters/{cluster}/vms/solutions/{solution}?vmw-task=true
//	DELETE /api/esx/settings/clusters/{cluster}/vms/solutions/{solution}?vmw-task=true
func (h *Handler) solution(w http.ResponseWriter, r *http.Request) {
	cluster := r.PathValue("cluster")
	solution := r.PathValue("solution")
	sols := h.clusterSolutions(cluster)

	switch r.Method {
	case http.MethodGet:
		info, ok := sols[solution]
		if !ok {
			vapi.ApiErrorNotFound(w)
			return
		}
		vapi.StatusOK(w, info)

	case http.MethodPut:
		var spec vms.SolutionSpec
		if !vapi.Decode(r, w, &spec) {
			return
		}
		// Round-trip through JSON to populate the overlapping SolutionInfo fields.
		data, _ := json.Marshal(spec)
		var info vms.SolutionInfo
		_ = json.Unmarshal(data, &info)
		sols[solution] = info
		vapi.StatusOK(w, h.newTask())

	case http.MethodDelete:
		if _, ok := sols[solution]; !ok {
			vapi.ApiErrorNotFound(w)
			return
		}
		delete(sols, solution)
		vapi.StatusOK(w, h.newTask())

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// hooksCollection handles POST actions on the lifecycle-hooks collection.
//
//	POST /api/esx/settings/clusters/{cluster}/vms/lifecycle-hooks?action=mark-as-processed
//	POST /api/esx/settings/clusters/{cluster}/vms/lifecycle-hooks?action=process-dynamic-update
func (h *Handler) hooksCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	action := r.URL.Query().Get("action")
	switch action {
	case "mark-as-processed":
		var spec vms.ProcessedHookSpec
		if vapi.Decode(r, w, &spec) {
			vapi.StatusOK(w, nil)
		}
	case "process-dynamic-update":
		var spec vms.DynamicUpdateSpec
		if vapi.Decode(r, w, &spec) {
			vapi.StatusOK(w, nil)
		}
	default:
		vapi.ApiErrorInvalidArgument(w)
	}
}

// hooksSolution handles GET on a per-solution hooks endpoint.
//
//	GET /api/esx/settings/clusters/{cluster}/vms/lifecycle-hooks/{solution}
func (h *Handler) hooksSolution(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	cluster := r.PathValue("cluster")
	solution := r.PathValue("solution")

	var hooks []vms.LifecycleHookInfo
	if clusterHooks, ok := h.Hooks[cluster]; ok {
		hooks = clusterHooks[solution]
	}
	if hooks == nil {
		hooks = []vms.LifecycleHookInfo{}
	}
	vapi.StatusOK(w, vms.HookListResult{Hooks: hooks})
}

// transition handles transition actions for a solution.
//
//	POST /api/esx/settings/clusters/{cluster}/vms/transition/{solution}?action=enable&vmw-task=true
//	POST /api/esx/settings/clusters/{cluster}/vms/transition/{solution}?action=multi-source-enable&vmw-task=true
//	POST /api/esx/settings/clusters/{cluster}/vms/transition/{solution}?action=transition&vmw-task=true
//	POST /api/esx/settings/clusters/{cluster}/vms/transition/{solution}?action=delete-solution-only
func (h *Handler) transition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cluster := r.PathValue("cluster")
	solution := r.PathValue("solution")
	action := r.URL.Query().Get("action")

	switch action {
	case "enable":
		var spec vms.EnableSpec
		if vapi.Decode(r, w, &spec) {
			vapi.StatusOK(w, h.newTask())
		}
	case "multi-source-enable":
		var spec vms.MultiSourceEnableSpec
		if vapi.Decode(r, w, &spec) {
			vapi.StatusOK(w, h.newTask())
		}
	case "transition":
		var spec vms.TransitionSpec
		if vapi.Decode(r, w, &spec) {
			vapi.StatusOK(w, h.newTask())
		}
	case "delete-solution-only":
		sols := h.clusterSolutions(cluster)
		if _, ok := sols[solution]; !ok {
			vapi.ApiErrorNotFound(w)
			return
		}
		delete(sols, solution)
		vapi.StatusOK(w)
	default:
		vapi.ApiErrorInvalidArgument(w)
	}
}

// taskGet handles task status polls from the task manager.
//
//	GET /api/cis/tasks/{task}
func (h *Handler) taskGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	result, ok := h.Tasks[r.PathValue("task")]
	if !ok {
		vapi.ApiErrorNotFound(w)
		return
	}
	vapi.StatusOK(w, result)
}
