// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	BaseImages      []depots.BaseImagesSummary

	SoftwareDrafts     map[string]clusters.SettingsClustersSoftwareDraftsMetadata
	SoftwareComponents map[string]clusters.SettingsComponentInfo
	ClusterImage       *clusters.SettingsBaseImageInfo

	depotCounter int
	draftCounter int

	vlcmEnabled bool
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL:                u,
		Depots:             make(map[string]depots.SettingsDepotsOfflineInfo),
		DepotComponents:    make(map[string]depots.SettingsDepotsOfflineContentInfo),
		BaseImages:         createMockBaseImages(),
		SoftwareDrafts:     make(map[string]clusters.SettingsClustersSoftwareDraftsMetadata),
		SoftwareComponents: make(map[string]clusters.SettingsComponentInfo),
		depotCounter:       0,
		vlcmEnabled:        false,
	}
}

func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(depots.DepotsOfflinePath, h.depotsOffline)
		s.HandleFunc(depots.DepotsOfflinePath+"/", h.depotsOffline)
		s.HandleFunc(depots.BaseImagesPath, h.baseImages)
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

func (h *Handler) baseImages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.BaseImages)
	}
}

func (h *Handler) clusters(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len("/api/esx/settings/clusters"):]
	segments := strings.Split(subpath, "/")

	if len(segments) > 3 && segments[2] == "software" && segments[3] == "drafts" {
		segments = segments[4:]
		if len(segments) > 2 && segments[1] == "software" && segments[2] == "components" {
			h.clustersSoftwareDraftsComponents(w, r, segments)
			return
		} else if len(segments) > 2 && segments[1] == "software" && segments[2] == "base-image" {
			h.clustersSoftwareDraftsBaseImage(w, r)
			return
		} else {
			h.clustersSoftwareDrafts(w, r, segments)
			return
		}
	} else if len(segments) > 3 && segments[2] == "enablement" && segments[3] == "software" {
		h.clustersSoftwareEnablement(w, r)
		return
	}

	vapi.ApiErrorUnsupported(w)
}

func (h *Handler) clustersSoftwareDrafts(w http.ResponseWriter, r *http.Request, subpath []string) {
	var draftId *string
	if len(subpath) > 0 {
		draftId = &subpath[0]
	}

	switch r.Method {
	case http.MethodGet:
		if draftId != nil {
			if draft, contains := h.SoftwareDrafts[*draftId]; !contains {
				vapi.ApiErrorNotFound(w)
				return
			} else {
				vapi.StatusOK(w, draft)
			}
		} else {
			vapi.StatusOK(w, h.SoftwareDrafts)
		}
	case http.MethodDelete:
		if draftId != nil {
			if _, contains := h.SoftwareDrafts[*draftId]; !contains {
				vapi.ApiErrorNotFound(w)
				return
			} else {
				delete(h.SoftwareDrafts, *draftId)
				vapi.StatusOK(w)
			}
		}
	case http.MethodPost:
		if strings.Contains(r.URL.RawQuery, "action=commit") {
			if draftId != nil {
				if _, contains := h.SoftwareDrafts[*draftId]; !contains {
					vapi.ApiErrorNotFound(w)
					return
				} else {
					delete(h.SoftwareDrafts, *draftId)
					vapi.StatusOK(w)
				}
			}
		}
		// Only one active draft is permitted
		if len(h.SoftwareDrafts) > 0 {
			vapi.ApiErrorNotAllowedInCurrentState(w)
			return
		}

		h.draftCounter += 1
		draft := clusters.SettingsClustersSoftwareDraftsMetadata{}
		newDraftId := fmt.Sprintf("%d", h.draftCounter)
		h.SoftwareDrafts[newDraftId] = draft
		vapi.StatusOK(w, newDraftId)
	}
}

func (h *Handler) clustersSoftwareDraftsComponents(w http.ResponseWriter, r *http.Request, subpath []string) {
	switch r.Method {
	case http.MethodGet:
		if len(subpath) > 3 {
			if comp, contains := h.SoftwareComponents[subpath[3]]; contains {
				vapi.StatusOK(w, comp)
			} else {
				vapi.ApiErrorNotFound(w)
			}
		} else {
			vapi.StatusOK(w, h.SoftwareComponents)
		}
	case http.MethodDelete:
		if len(subpath) > 3 {
			compId := subpath[3]
			if comp, contains := h.SoftwareComponents[compId]; contains {
				delete(h.SoftwareComponents, compId)
				vapi.StatusOK(w, comp)
			} else {
				vapi.ApiErrorNotFound(w)
			}
		}
	case http.MethodPatch:
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

func (h *Handler) clustersSoftwareDraftsBaseImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, h.ClusterImage)
	case http.MethodPut:
		var spec clusters.SettingsBaseImageSpec
		if vapi.Decode(r, w, &spec) {
			h.ClusterImage = &clusters.SettingsBaseImageInfo{Version: spec.Version}
			vapi.StatusOK(w)
		} else {
			vapi.ApiErrorInvalidArgument(w)
		}
	}
}

func (h *Handler) clustersSoftwareEnablement(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vapi.StatusOK(w, clusters.SoftwareManagementInfo{Enabled: h.vlcmEnabled})
	case http.MethodPut:
		h.vlcmEnabled = true
		vapi.StatusOK(w)
	}
}

func createMockBaseImages() []depots.BaseImagesSummary {
	baseImage := depots.BaseImagesSummary{
		DisplayName:    "DummyImage",
		DisplayVersion: "0.0.1",
		Version:        "0.0.1",
	}
	return []depots.BaseImagesSummary{baseImage}
}
