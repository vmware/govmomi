// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"net/url"
	"strings"
	"sync"

	"github.com/vmware/govmomi/lookup"
	"github.com/vmware/govmomi/lookup/methods"
	"github.com/vmware/govmomi/lookup/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

var content = types.LookupServiceContent{
	LookupService:                vim.ManagedObjectReference{Type: "LookupLookupService", Value: "lookupService"},
	ServiceRegistration:          &vim.ManagedObjectReference{Type: "LookupServiceRegistration", Value: "ServiceRegistration"},
	DeploymentInformationService: vim.ManagedObjectReference{Type: "LookupDeploymentInformationService", Value: "deploymentInformationService"},
	L10n:                         vim.ManagedObjectReference{Type: "LookupL10n", Value: "l10n"},
}

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New())
		}
	})
}

func New() *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = lookup.Namespace
	r.Path = lookup.Path

	r.Put(&ServiceInstance{
		ManagedObjectReference: lookup.ServiceInstance,
		Content:                content,
	})

	r.Put(&ServiceRegistration{
		ManagedObjectReference: *content.ServiceRegistration,
	})

	return r
}

type ServiceInstance struct {
	vim.ManagedObjectReference

	Content types.LookupServiceContent
}

func (s *ServiceInstance) RetrieveServiceContent(ctx *simulator.Context, _ *types.RetrieveServiceContent) soap.HasFault {
	// Initialize prior to List() being called (see ExampleServiceRegistration)
	ctx.Map.Get(*content.ServiceRegistration).(*ServiceRegistration).info(ctx)

	return &methods.RetrieveServiceContentBody{
		Res: &types.RetrieveServiceContentResponse{
			Returnval: s.Content,
		},
	}
}

type ServiceRegistration struct {
	vim.ManagedObjectReference

	Info []types.LookupServiceRegistrationInfo

	register sync.Once
	mu       sync.RWMutex
}

func (s *ServiceRegistration) GetSiteId(_ *types.GetSiteId) soap.HasFault {
	return &methods.GetSiteIdBody{
		Res: &types.GetSiteIdResponse{
			Returnval: siteID,
		},
	}
}

func matchServiceType(filter, info *types.LookupServiceRegistrationServiceType) bool {
	if filter.Product != "" {
		if filter.Product != info.Product {
			return false
		}
	}

	if filter.Type != "" {
		if filter.Type != info.Type {
			return false
		}
	}

	return true
}

func matchEndpointType(filter, info *types.LookupServiceRegistrationEndpointType) bool {
	if filter.Protocol != "" {
		if filter.Protocol != info.Protocol {
			return false
		}
	}

	if filter.Type != "" {
		if filter.Type != info.Type {
			return false
		}
	}

	return true
}

// defer register to this point to ensure we can include vcsim's cert in ServiceEndpoints.SslTrust
// TODO: we should be able to register within New(), but this is the only place that currently depends on vcsim's cert
func (s *ServiceRegistration) info(ctx *simulator.Context) {
	s.register.Do(func() {
		s.Info = registrationInfo(ctx)
	})
}

func (s *ServiceRegistration) Get(ctx *simulator.Context, req *types.Get) soap.HasFault {
	s.info(ctx)

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, entry := range s.Info {
		if entry.ServiceId == req.ServiceId {
			return &methods.GetBody{
				Res: &types.GetResponse{Returnval: entry},
			}
		}
	}

	body := new(methods.GetBody)
	body.Fault_ = simulator.Fault("LookupFaultServiceFault", &vim.SystemError{
		Reason: "Service not found: " + req.ServiceId,
	})
	return body
}

// Create registers a new service entry in the Lookup Service.
// It is idempotent with respect to ServiceId: if an entry already exists it is overwritten.
func (s *ServiceRegistration) Create(ctx *simulator.Context, req *types.Create) soap.HasFault {
	s.info(ctx)

	entry := types.LookupServiceRegistrationInfo{
		LookupServiceRegistrationCommonServiceInfo: req.CreateSpec.LookupServiceRegistrationCommonServiceInfo,
		ServiceId: req.ServiceId,
		SiteId:    siteID,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, existing := range s.Info {
		if existing.ServiceId == req.ServiceId {
			s.Info[i] = entry
			return &methods.CreateBody{Res: new(types.CreateResponse)}
		}
	}
	s.Info = append(s.Info, entry)

	return &methods.CreateBody{Res: new(types.CreateResponse)}
}

// Delete removes a service entry from the Lookup Service.
// Returns a fault if the entry isn't found.
func (s *ServiceRegistration) Delete(ctx *simulator.Context, req *types.Delete) soap.HasFault {
	s.info(ctx)

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, entry := range s.Info {
		if entry.ServiceId == req.ServiceId {
			s.Info = append(s.Info[:i], s.Info[i+1:]...)
			return &methods.DeleteBody{Res: new(types.DeleteResponse)}
		}
	}

	body := new(methods.DeleteBody)
	body.Fault_ = simulator.Fault("LookupFaultServiceFault", &vim.SystemError{
		Reason: "Service not found: " + req.ServiceId,
	})
	return body
}

func (s *ServiceRegistration) List(ctx *simulator.Context, req *types.List) soap.HasFault {
	body := new(methods.ListBody)
	filter := req.FilterCriteria

	if filter == nil {
		// This is what a real PSC returns if FilterCriteria is nil.
		body.Fault_ = simulator.Fault("LookupFaultServiceFault", &vim.SystemError{
			Reason: "Invalid fault",
		})
		return body
	}
	body.Res = new(types.ListResponse)

	s.info(ctx)

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, info := range s.Info {
		if filter.SiteId != "" {
			if filter.SiteId != info.SiteId {
				continue
			}
		}
		if filter.NodeId != "" {
			if filter.NodeId != info.NodeId {
				continue
			}
		}
		if filter.ServiceType != nil {
			if !matchServiceType(filter.ServiceType, &info.ServiceType) {
				continue
			}
		}
		if filter.EndpointType != nil {
			services := info.ServiceEndpoints
			info.ServiceEndpoints = nil
			for _, service := range services {
				if !matchEndpointType(filter.EndpointType, &service.EndpointType) {
					continue
				}
				info.ServiceEndpoints = append(info.ServiceEndpoints, service)
			}
			if len(info.ServiceEndpoints) == 0 {
				continue
			}
		}
		body.Res.Returnval = append(body.Res.Returnval, info)
	}

	return body
}

// BreakLookupServiceURLs makes the path of all lookup service urls invalid
func BreakLookupServiceURLs(ctx context.Context) {
	setting := simulator.Map(ctx).OptionManager().Setting

	for _, s := range setting {
		o := s.GetOptionValue()
		if strings.HasSuffix(o.Key, ".uri") {
			val := o.Value.(string)
			u, _ := url.Parse(val)
			u.Path = "/enoent" + u.Path
			o.Value = u.String()
		}
	}
}

// UnresolveLookupServiceURLs makes the path of all lookup service urls invalid
func UnresolveLookupServiceURLs(ctx context.Context) {
	setting := simulator.Map(ctx).OptionManager().Setting

	for _, s := range setting {
		o := s.GetOptionValue()
		if strings.HasSuffix(o.Key, ".uri") {
			val := o.Value.(string)
			u, _ := url.Parse(val)
			port := u.Port()
			u.Host = "fake-name-will-not-dns-resolve:" + port
			o.Value = u.String()
		}
	}
}
