// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/crypto"
	"github.com/vmware/govmomi/vapi/crypto/internal"
	vapi "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	typeNativeProvider = string(types.KmipClusterInfoKmsManagementTypeNativeProvider)
	backupPath         = "/cryptomanager/kms/"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		New(r, s.Listen).Register(s, r)
	})
}

// Handler implements the Cluster Modules API simulator
type Handler struct {
	URL *url.URL
	Map *simulator.Registry
}

// New creates a Handler instance
func New(r *simulator.Registry, u *url.URL) *Handler {
	return &Handler{
		Map: r,
		URL: u,
	}
}

// Register Namespace Management API paths with the vapi simulator's http.ServeMux
func (h *Handler) Register(s *simulator.Service, r *simulator.Registry) {
	if r.IsVPX() {
		s.HandleFunc(internal.KmsProvidersPath, h.providers)
		s.HandleFunc(internal.KmsProvidersPath+"/", h.providersID)
		s.HandleFunc(backupPath, h.backup)
	}
}

// We need to use the simulator objects directly when updating fields (e.g. HasBackup)
// Skipping the trouble of locking for now, as existing use-cases would not race.
func (h *Handler) find(id string) *types.KmipClusterInfo {
	m := h.Map.CryptoManager()
	for i := range m.KmipServers {
		p := &m.KmipServers[i]
		if p.ClusterId.Id == id {
			return p
		}
	}
	return nil
}

func (h *Handler) backup(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.RequestURI)
	p := h.find(id)
	if p == nil {
		vapi.ApiErrorNotFound(w)
		return
	}

	// Content of the simulated backup does not matter for the use-case we're covering:
	// Export sets HasBackup=true, which sets CryptoManagerKmipClusterStatus.OverallStatus=green
	p.HasBackup = types.NewBool(true)

	name := fmt.Sprintf("%s%s.p12", id, time.Now().Format(time.RFC3339))

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	_ = json.NewEncoder(w).Encode(p)
}

func (h *Handler) providers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		switch r.URL.Query().Get("action") {
		case "":
			var spec crypto.KmsProviderCreateSpec

			if vapi.Decode(r, w, &spec) {
				if h.find(spec.Provider) != nil {
					vapi.ApiErrorAlreadyExists(w)
					return
				}
			}

			m := h.Map.CryptoManager()
			m.KmipServers = append(m.KmipServers, types.KmipClusterInfo{
				ClusterId:      types.KeyProviderId{Id: spec.Provider},
				ManagementType: typeNativeProvider,
				TpmRequired:    &spec.Constraints.TpmRequired,
				HasBackup:      types.NewBool(false),
			})
		case "export":
			var spec crypto.KmsProviderExportSpec
			var p *types.KmipClusterInfo

			if vapi.Decode(r, w, &spec) {
				if p = h.find(spec.Provider); p == nil {
					vapi.ApiErrorNotFound(w)
					return
				}
				if p.ManagementType != typeNativeProvider {
					vapi.ApiErrorUnsupported(w)
					return
				}

				u := url.URL{
					Scheme: h.URL.Scheme,
					Host:   h.URL.Host,
					Path:   backupPath + spec.Provider,
				}

				res := crypto.KmsProviderExport{
					Type: "LOCATION",
					Location: &crypto.KmsProviderExportLocation{
						URL: u.String(),
						DownloadToken: crypto.KmsProviderDownloadToken{
							Token:  uuid.NewString(),
							Expiry: time.Now().Add(time.Minute).Format(time.RFC3339),
						},
					},
				}

				vapi.StatusOK(w, res)
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) providersID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		id := path.Base(r.RequestURI)
		p := h.find(id)
		if p == nil {
			vapi.ApiErrorNotFound(w)
			return
		}
		if p.ManagementType != typeNativeProvider {
			vapi.ApiErrorUnsupported(w)
			return
		}
		m := h.Map.CryptoManager()
		ctx := &simulator.Context{Map: h.Map}
		_ = m.UnregisterKmsCluster(ctx, &types.UnregisterKmsCluster{
			This:      m.Self,
			ClusterId: types.KeyProviderId{Id: id},
		})
		vapi.StatusOK(w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
