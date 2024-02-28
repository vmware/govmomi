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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
	"github.com/vmware/govmomi/vapi/esx/settings/depots"
	vapi "github.com/vmware/govmomi/vapi/simulator"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

type Handler struct {
	URL *url.URL

	Depots          map[string]depots.SettingsDepotsOfflineInfo
	DepotComponents map[string]depots.SettingsDepotsOfflineContentInfo

	SoftwareDrafts     map[string]clusters.SettingsClustersSoftwareDraftsMetadata
	SoftwareComponents map[string]clusters.SettingsComponentInfo

	depotCounter int
	draftCounter int
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL:                u,
		Depots:             make(map[string]depots.SettingsDepotsOfflineInfo),
		DepotComponents:    make(map[string]depots.SettingsDepotsOfflineContentInfo),
		SoftwareDrafts:     make(map[string]clusters.SettingsClustersSoftwareDraftsMetadata),
		SoftwareComponents: make(map[string]clusters.SettingsComponentInfo),
		depotCounter:       0,
	}
}

func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(depots.DepotsOfflinePath, h.depotsOffline)
		s.HandleFunc(depots.DepotsOfflinePath+"/", h.depotsOffline)
		s.HandleFunc("/api/esx/settings/clusters/", h.clusters)
	}
}

func (h *Handler) depotsOffline(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(depots.DepotsOfflinePath):]
	segments := strings.Split(subpath, "/")

	switch r.Method {
	case http.MethodGet:
		if len(segments) > 1 {
			if res, contains := h.DepotComponents[segments[1]]; !contains {
				vapi.ApiErrorNotFound(w)
			} else if len(segments) > 2 && segments[2] == "content" {
				vapi.StatusOK(w, res)
			} else {
				vapi.ApiErrorUnsupported(w)
			}
		}
		vapi.StatusOK(w, h.Depots)
	case http.MethodDelete:
		if _, contains := h.Depots[segments[1]]; !contains {
			vapi.ApiErrorNotFound(w)
		} else {
			delete(h.Depots, segments[1])
			delete(h.DepotComponents, segments[1])
			vapi.StatusOK(w, "")
		}
	case http.MethodPost:
		var spec depots.SettingsDepotsOfflineCreateSpec
		if vapi.Decode(r, w, &spec) {
			// Create depot
			depot := depots.SettingsDepotsOfflineInfo{}
			depot.SourceType = spec.SourceType
			depot.FileId = spec.FileId
			depot.Location = spec.Location
			depot.OwnerData = spec.OwnerData
			depot.Description = spec.Description

			h.depotCounter += 1
			depotId := fmt.Sprintf("depot-%d", h.depotCounter)
			h.Depots[depotId] = depot

			// Generate content
			content := depots.SettingsDepotsOfflineContentInfo{}
			content.MetadataBundles = make(map[string][]depots.SettingsDepotsMetadataInfo)
			content.MetadataBundles["dummy-content"] = make([]depots.SettingsDepotsMetadataInfo, 1)
			bundle := depots.SettingsDepotsMetadataInfo{}
			bundle.IndependentComponents = make(map[string]depots.SettingsDepotsComponentSummary)
			independentComp := depots.SettingsDepotsComponentSummary{}
			independentComp.Versions = make([]depots.ComponentVersion, 1)
			independentComp.Versions[0].Version = "1.0.0"
			independentComp.Versions[0].DisplayVersion = "1.0.0"
			independentComp.DisplayName = "DummyComponent"
			bundle.IndependentComponents["dummy-component"] = independentComp
			content.MetadataBundles["dummy-content"][0] = bundle
			h.DepotComponents[depotId] = content

			vapi.StatusOK(w, fmt.Sprintf("depot-task-%d", h.depotCounter))
		}
	}
}

func (h *Handler) clusters(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len("/api/esx/settings/clusters"):]
	segments := strings.Split(subpath, "/")

	if len(segments) < 4 || segments[2] != "software" || segments[3] != "drafts" {
		vapi.ApiErrorUnsupported(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		id := fmt.Sprintf("%d", h.draftCounter)
		if isEndpointSoftwareComponents(segments, id) {
			if len(segments) > 7 {
				if comp, contains := h.SoftwareComponents[segments[7]]; contains {
					vapi.StatusOK(w, comp)
				} else {
					vapi.ApiErrorNotFound(w)
				}
			} else {
				vapi.StatusOK(w, h.SoftwareComponents)
			}
			return
		}

		if len(segments) > 4 {
			if draft, contains := h.SoftwareDrafts[id]; contains {
				vapi.StatusOK(w, draft)
			} else {
				vapi.ApiErrorNotFound(w)
			}
		} else {
			vapi.StatusOK(w, h.SoftwareDrafts)
		}
	case http.MethodDelete:
		id := fmt.Sprintf("%d", h.draftCounter)
		// component delete
		if isEndpointSoftwareComponents(segments, id) {
			if len(segments) > 7 {
				delete(h.SoftwareComponents, segments[7])
				vapi.StatusOK(w, "")
				return
			}
		}

		// draft delete
		if len(segments) > 4 && segments[4] == id {
			delete(h.SoftwareDrafts, id)
			vapi.StatusOK(w, "")
			return
		}

		vapi.ApiErrorNotFound(w)
	case http.MethodPost:
		if strings.Contains(r.URL.RawQuery, "action=commit") {
			id := fmt.Sprintf("%d", h.draftCounter)
			if len(segments) > 4 && segments[4] == id {
				delete(h.SoftwareDrafts, id)
				vapi.StatusOK(w, "")
			} else {
				vapi.ApiErrorNotFound(w)
			}
			return
		}
		// Only one active draft is permitted
		if len(h.SoftwareDrafts) > 0 {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}

		h.draftCounter += 1
		draft := clusters.SettingsClustersSoftwareDraftsMetadata{}
		id := fmt.Sprintf("%d", h.draftCounter)
		h.SoftwareDrafts[id] = draft
		vapi.StatusOK(w, id)
	case http.MethodPatch:
		id := fmt.Sprintf("%d", h.draftCounter)
		if !isEndpointSoftwareComponents(segments, id) {
			vapi.ApiErrorUnsupported(w)
			return
		}

		var spec clusters.SoftwareComponentsUpdateSpec
		if vapi.Decode(r, w, &spec) {
			for k, v := range spec.ComponentsToSet {
				h.SoftwareComponents[k] = clusters.SettingsComponentInfo{
					Version: v,
					Details: clusters.SettingsComponentDetails{DisplayName: "DummyComponent"},
				}
			}
		}
	}
}

func isEndpointSoftwareComponents(subpathSegments []string, draftId string) bool {
	return len(subpathSegments) > 6 &&
		subpathSegments[4] == draftId &&
		subpathSegments[5] == "software" &&
		subpathSegments[6] == "components"
}
