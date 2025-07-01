// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/google/uuid"

	"github.com/vmware/govmomi/lookup"
	"github.com/vmware/govmomi/lookup/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

var (
	siteID = "vcsim"
)

// registrationInfo returns a ServiceRegistration populated with vcsim's OptionManager settings.
// The complete list can be captured using: govc sso.service.ls -dump
func registrationInfo(ctx *simulator.Context) []types.LookupServiceRegistrationInfo {
	vctx := ctx.For(vim25.Path)
	vc := vctx.Map.Get(vim25.ServiceInstance).(*simulator.ServiceInstance)
	setting := vctx.Map.OptionManager().Setting
	sm := vctx.Map.SessionManager()
	opts := make(map[string]string, len(setting))

	for _, o := range setting {
		opt := o.GetOptionValue()
		if val, ok := opt.Value.(string); ok {
			opts[opt.Key] = val
		}
	}

	trust := []string{sm.TLSCert()}
	sdk := opts["vcsim.server.url"] + vim25.Path
	admin := opts["config.vpxd.sso.default.admin"]
	owner := opts["config.vpxd.sso.solutionUser.name"]
	instance := opts["VirtualCenter.InstanceName"]

	// Real PSC has 30+ services by default, we just provide a few that are useful for vmomi interaction..
	info := []types.LookupServiceRegistrationInfo{
		{
			LookupServiceRegistrationCommonServiceInfo: types.LookupServiceRegistrationCommonServiceInfo{
				LookupServiceRegistrationMutableServiceInfo: types.LookupServiceRegistrationMutableServiceInfo{
					ServiceVersion: lookup.Version,
					ServiceEndpoints: []types.LookupServiceRegistrationEndpoint{
						{
							Url: opts["config.vpxd.sso.sts.uri"],
							EndpointType: types.LookupServiceRegistrationEndpointType{
								Protocol: "wsTrust",
								Type:     "com.vmware.cis.cs.identity.sso",
							},
							SslTrust: trust,
						},
					},
				},
				OwnerId: admin,
				ServiceType: types.LookupServiceRegistrationServiceType{
					Product: "com.vmware.cis",
					Type:    "cs.identity",
				},
			},
			ServiceId: siteID + ":" + uuid.New().String(),
			SiteId:    siteID,
		},
		{
			LookupServiceRegistrationCommonServiceInfo: types.LookupServiceRegistrationCommonServiceInfo{
				LookupServiceRegistrationMutableServiceInfo: types.LookupServiceRegistrationMutableServiceInfo{
					ServiceVersion: lookup.Version,
					ServiceEndpoints: []types.LookupServiceRegistrationEndpoint{
						{
							Url: opts["config.vpxd.sso.admin.uri"],
							EndpointType: types.LookupServiceRegistrationEndpointType{
								Protocol: "vmomi",
								Type:     "com.vmware.cis.cs.identity.admin",
							},
							SslTrust: trust,
						},
					},
				},
				OwnerId: admin,
				ServiceType: types.LookupServiceRegistrationServiceType{
					Product: "com.vmware.cis",
					Type:    "cs.identity",
				},
			},
			ServiceId: siteID + ":" + uuid.New().String(),
			SiteId:    siteID,
		},
		{
			LookupServiceRegistrationCommonServiceInfo: types.LookupServiceRegistrationCommonServiceInfo{
				LookupServiceRegistrationMutableServiceInfo: types.LookupServiceRegistrationMutableServiceInfo{
					ServiceVersion: vim25.Version,
					ServiceEndpoints: []types.LookupServiceRegistrationEndpoint{
						{
							Url: sdk,
							EndpointType: types.LookupServiceRegistrationEndpointType{
								Protocol: "vmomi",
								Type:     "com.vmware.vim",
							},
							SslTrust: trust,
							EndpointAttributes: []types.LookupServiceRegistrationAttribute{
								{
									Key:   "cis.common.ep.localurl",
									Value: sdk,
								},
							},
						},
					},
					ServiceAttributes: []types.LookupServiceRegistrationAttribute{
						{
							Key:   "com.vmware.cis.cm.GroupInternalId",
							Value: "com.vmware.vim.vcenter",
						},
						{
							Key:   "com.vmware.vim.vcenter.instanceName",
							Value: instance,
						},
						{
							Key:   "com.vmware.cis.cm.ControlScript",
							Value: "service-control-default-vmon",
						},
						{
							Key:   "com.vmware.cis.cm.HostId",
							Value: uuid.New().String(),
						},
					},
					ServiceNameResourceKey:        "AboutInfo.vpx.name",
					ServiceDescriptionResourceKey: "AboutInfo.vpx.name",
				},
				OwnerId: owner,
				ServiceType: types.LookupServiceRegistrationServiceType{
					Product: "com.vmware.cis",
					Type:    "vcenterserver",
				},
				NodeId: uuid.New().String(),
			},
			ServiceId: vc.Content.About.InstanceUuid,
			SiteId:    siteID,
		},
	}

	sts := info[0]
	sts.ServiceType.Type = "sso:sts" // obsolete service type, but still used by PowerCLI

	return append(info, sts)
}
