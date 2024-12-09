/*
Copyright (c) 2017-2024 VMware, Inc. All Rights Reserved.

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
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// EvalLicense is the default license
var EvalLicense = types.LicenseManagerLicenseInfo{
	LicenseKey: "00000-00000-00000-00000-00000",
	EditionKey: "eval",
	Name:       "Evaluation Mode",
	Properties: []types.KeyAnyValue{
		{
			Key: "feature",
			Value: types.KeyValue{
				Key:   "serialuri:2",
				Value: "Remote virtual Serial Port Concentrator",
			},
		},
		{
			Key: "feature",
			Value: types.KeyValue{
				Key:   "dvs",
				Value: "vSphere Distributed Switch",
			},
		},
	},
}

type LicenseManager struct {
	mo.LicenseManager
}

func (m *LicenseManager) init(r *Registry) {
	if len(m.Licenses) == 0 {
		about := r.content().About
		product := []types.KeyAnyValue{
			{
				Key:   "ProductName",
				Value: about.LicenseProductName,
			},
			{
				Key:   "ProductVersion",
				Value: about.LicenseProductVersion,
			},
		}

		EvalLicense.Properties = append(EvalLicense.Properties, product...)

		m.Licenses = []types.LicenseManagerLicenseInfo{EvalLicense}
	}

	if r.IsVPX() {
		if m.LicenseAssignmentManager == nil {
			m.LicenseAssignmentManager = &types.ManagedObjectReference{
				Type:  "LicenseAssignmentManager",
				Value: "LicenseAssignmentManager",
			}
		}
		lam := new(LicenseAssignmentManager)
		lam.Self = *m.LicenseAssignmentManager
		r.Put(lam)
	}
}

func (m *LicenseManager) AddLicense(ctx *Context, req *types.AddLicense) soap.HasFault {
	body := &methods.AddLicenseBody{
		Res: &types.AddLicenseResponse{},
	}

	for _, license := range m.Licenses {
		if license.LicenseKey == req.LicenseKey {
			body.Res.Returnval = licenseInfo(license.LicenseKey, license.Labels)
			return body
		}
	}

	m.Licenses = append(m.Licenses, types.LicenseManagerLicenseInfo{
		LicenseKey: req.LicenseKey,
		Labels:     req.Labels,
	})

	body.Res.Returnval = licenseInfo(req.LicenseKey, req.Labels)

	return body
}

func (m *LicenseManager) RemoveLicense(ctx *Context, req *types.RemoveLicense) soap.HasFault {
	body := &methods.RemoveLicenseBody{
		Res: &types.RemoveLicenseResponse{},
	}

	for i, license := range m.Licenses {
		if req.LicenseKey == license.LicenseKey {
			m.Licenses = append(m.Licenses[:i], m.Licenses[i+1:]...)
			return body
		}
	}
	return body
}

func (m *LicenseManager) UpdateLicenseLabel(ctx *Context, req *types.UpdateLicenseLabel) soap.HasFault {
	body := &methods.UpdateLicenseLabelBody{}

	for i := range m.Licenses {
		license := &m.Licenses[i]

		if req.LicenseKey != license.LicenseKey {
			continue
		}

		body.Res = new(types.UpdateLicenseLabelResponse)

		for j := range license.Labels {
			label := &license.Labels[j]

			if label.Key == req.LabelKey {
				if req.LabelValue == "" {
					license.Labels = append(license.Labels[:i], license.Labels[i+1:]...)
				} else {
					label.Value = req.LabelValue
				}
				return body
			}
		}

		license.Labels = append(license.Labels, types.KeyValue{
			Key:   req.LabelKey,
			Value: req.LabelValue,
		})

		return body
	}

	body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "licenseKey"})
	return body
}

type LicenseAssignmentManager struct {
	mo.LicenseAssignmentManager

	types.QueryAssignedLicensesResponse
}

func (m *LicenseAssignmentManager) QueryAssignedLicenses(ctx *Context, req *types.QueryAssignedLicenses) soap.HasFault {
	body := &methods.QueryAssignedLicensesBody{
		Res: &types.QueryAssignedLicensesResponse{},
	}

	if len(m.QueryAssignedLicensesResponse.Returnval) != 0 {
		// Using Returnval from govc object.save -l
		if req.EntityId == "" {
			body.Res = &m.QueryAssignedLicensesResponse
		} else {
			for _, r := range m.QueryAssignedLicensesResponse.Returnval {
				if r.EntityId == req.EntityId {
					body.Res.Returnval = append(body.Res.Returnval, r)
				}
			}
		}
		return body
	}

	c := ctx.Map.content()

	if req.EntityId == "" {
		var add func(child types.ManagedObjectReference)

		add = func(child types.ManagedObjectReference) {
			if child.Type == "HostSystem" || child.Type == "ClusterComputeResource" {
				la := types.LicenseAssignmentManagerLicenseAssignment{
					EntityId:          child.Value,
					Scope:             c.About.InstanceUuid,
					EntityDisplayName: ctx.Map.Get(child).(mo.Entity).Entity().Name,
					AssignedLicense:   EvalLicense,
				}
				body.Res.Returnval = append(body.Res.Returnval, la)
			}
			walk(ctx.Map.Get(child), add)
		}

		walk(ctx.Map.Get(c.RootFolder), add)

		la := types.LicenseAssignmentManagerLicenseAssignment{
			EntityId:          c.About.InstanceUuid,
			EntityDisplayName: ctx.svc.Listen.Hostname(),
			AssignedLicense:   EvalLicense,
		}
		body.Res.Returnval = append(body.Res.Returnval, la)
	} else {
		name := ctx.svc.Listen.Hostname()
		// EntityId can be a HostSystem or the vCenter InstanceUuid
		if req.EntityId != c.About.InstanceUuid {
			id := types.ManagedObjectReference{
				Type:  "HostSystem",
				Value: req.EntityId,
			}
			e := ctx.Map.Get(id)
			if e == nil {
				return body
			}
			name = e.(mo.Entity).Entity().Name
		}

		body.Res.Returnval = []types.LicenseAssignmentManagerLicenseAssignment{{
			EntityId:          req.EntityId,
			Scope:             c.About.InstanceUuid,
			EntityDisplayName: name,
			AssignedLicense:   EvalLicense,
		}}
	}

	return body
}

func (m *LicenseAssignmentManager) UpdateAssignedLicense(ctx *Context, req *types.UpdateAssignedLicense) soap.HasFault {
	body := &methods.UpdateAssignedLicenseBody{
		Res: &types.UpdateAssignedLicenseResponse{
			Returnval: licenseInfo(req.LicenseKey, nil),
		},
	}

	return body
}

func licenseInfo(key string, labels []types.KeyValue) types.LicenseManagerLicenseInfo {
	info := EvalLicense

	info.LicenseKey = key
	info.Labels = labels

	return info
}
