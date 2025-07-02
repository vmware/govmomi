// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware/govmomi/simulator"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vapi/vcenter/consumptiondomains/zones"
)

const (
	zonesPath        = "/api/vcenter/consumption-domains/zones"
	associationsPath = "/api/vcenter/consumption-domains/zones/cluster"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(s.Listen).Register(s, r)
	})
}

// ZoneData Helper type to store simulated entries
type ZoneData struct {
	Name         string
	Id           string
	Description  string
	Associations []string
}

// Handler implements the Cluster Modules API simulator
type Handler struct {
	URL  *url.URL
	data map[string]*ZoneData
}

// New creates a Handler instance
func New(u *url.URL) *Handler {
	return &Handler{
		URL:  u,
		data: make(map[string]*ZoneData),
	}
}

// Register Consumption Domains API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(zonesPath, h.zones)
		s.HandleFunc(zonesPath+"/", h.zones)
		s.HandleFunc(associationsPath, h.associations)
		s.HandleFunc(associationsPath+"/", h.associations)
	}
}

func (h *Handler) zones(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(zonesPath):]
	zoneId := strings.Replace(subpath, "/", "", -1)

	switch r.Method {
	case http.MethodGet:
		if len(subpath) > 0 {
			if d, ok := h.data[zoneId]; ok {
				vapi.StatusOK(w, zones.ZoneInfo{Description: d.Description})
				return
			}
			vapi.ApiErrorNotFound(w)
		} else {
			items := make([]zones.ListItem, len(h.data))
			i := 0
			for _, d := range h.data {
				item := zones.ListItem{
					Zone: d.Name,
					Info: zones.ZoneInfo{
						Description: d.Description,
					},
				}
				items[i] = item
				i++
			}

			result := zones.ListResult{Items: items}
			vapi.StatusOK(w, result)
		}
	case http.MethodPost:
		var spec zones.CreateSpec
		if !vapi.Decode(r, w, &spec) {
			vapi.ApiErrorGeneral(w)
			return
		}

		newZone := ZoneData{
			Name:         spec.Zone,
			Description:  spec.Description,
			Id:           spec.Zone,
			Associations: make([]string, 0),
		}
		h.data[newZone.Id] = &newZone

		vapi.StatusOK(w, newZone.Id)
	case http.MethodDelete:
		if _, ok := h.data[zoneId]; ok {
			delete(h.data, zoneId)
			vapi.StatusOK(w)
			return
		}
		vapi.ApiErrorNotFound(w)
	}
}

func (h *Handler) associations(w http.ResponseWriter, r *http.Request) {
	subpath := r.URL.Path[len(associationsPath)+1:]
	pathParts := strings.Split(subpath, "/")

	if len(pathParts) != 2 || pathParts[1] != "associations" {
		vapi.ApiErrorNotFound(w)
		return
	}

	zoneId := pathParts[0]

	switch r.Method {
	case http.MethodGet:
		if d, ok := h.data[zoneId]; ok {
			vapi.StatusOK(w, d.Associations)
			return
		}
	case http.MethodPost:
		action := r.URL.Query().Get("action")

		var clusterIds []string
		if !vapi.Decode(r, w, &clusterIds) {
			vapi.ApiErrorGeneral(w)
			return
		}

		switch action {
		case "add":
			if d, ok := h.data[zoneId]; ok {
				associations := append(d.Associations, clusterIds...)
				d.Associations = associations
				res := make(map[string]any)
				res["success"] = true
				vapi.StatusOK(w, res)
				return
			}
			vapi.ApiErrorNotFound(w)
		case "remove":
			if d, ok := h.data[zoneId]; ok {
				associations := make([]string, 0)

				for _, a := range d.Associations {
					found := false
					for _, id := range clusterIds {
						if a == id {
							found = true
						}
					}

					if !found {
						associations = append(associations, a)
					}
				}

				d.Associations = associations

				vapi.StatusOK(w, nil)
				return
			}
			vapi.ApiErrorNotFound(w)
		default:
			vapi.ApiErrorGeneral(w)
		}
	}
}
