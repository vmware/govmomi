// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"slices"

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
		lam.QueryAssignedLicensesResponse.Returnval = []types.LicenseAssignmentManagerLicenseAssignment{{
			EntityId:          r.content().About.InstanceUuid,
			EntityDisplayName: "vcsim",
			AssignedLicense:   EvalLicense,
		}}
		r.Put(lam)
		r.AddHandler(lam)
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

func (m *LicenseManager) DecodeLicense(ctx *Context, req *types.DecodeLicense) soap.HasFault {
	body := &methods.DecodeLicenseBody{
		Res: &types.DecodeLicenseResponse{},
	}

	for _, license := range m.Licenses {
		if req.LicenseKey == license.LicenseKey {
			body.Res.Returnval = license
			break
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

var licensedTypes = []string{"HostSystem", "ClusterComputeResource"}

// PutObject assigns a license when a host or cluster is created.
func (m *LicenseAssignmentManager) PutObject(ctx *Context, obj mo.Reference) {
	ref := obj.Reference()

	if !slices.Contains(licensedTypes, ref.Type) {
		return
	}

	if slices.ContainsFunc(m.QueryAssignedLicensesResponse.Returnval,
		func(am types.LicenseAssignmentManagerLicenseAssignment) bool {
			return am.EntityId == ref.Value
		}) {
		return // via vcsim -load
	}

	la := types.LicenseAssignmentManagerLicenseAssignment{
		EntityId:          ref.Value,
		Scope:             ctx.Map.content().About.InstanceUuid,
		EntityDisplayName: obj.(mo.Entity).Entity().Name,
		AssignedLicense:   EvalLicense,
	}

	m.QueryAssignedLicensesResponse.Returnval =
		append(m.QueryAssignedLicensesResponse.Returnval, la)
}

// RemoveObject removes the license assignment when a host or cluster is removed.
func (m *LicenseAssignmentManager) RemoveObject(ctx *Context, ref types.ManagedObjectReference) {
	if !slices.Contains(licensedTypes, ref.Type) {
		return
	}

	m.QueryAssignedLicensesResponse.Returnval =
		slices.DeleteFunc(m.QueryAssignedLicensesResponse.Returnval,
			func(am types.LicenseAssignmentManagerLicenseAssignment) bool {
				return am.EntityId == ref.Value
			})
}

func (*LicenseAssignmentManager) UpdateObject(*Context, mo.Reference, []types.PropertyChange) {}

func (m *LicenseAssignmentManager) init(r *Registry) {
	r.AddHandler(m)
}

func (m *LicenseAssignmentManager) QueryAssignedLicenses(ctx *Context, req *types.QueryAssignedLicenses) soap.HasFault {
	body := &methods.QueryAssignedLicensesBody{
		Res: &types.QueryAssignedLicensesResponse{},
	}

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

func (m *LicenseAssignmentManager) UpdateAssignedLicense(ctx *Context, req *types.UpdateAssignedLicense) soap.HasFault {
	body := new(methods.UpdateAssignedLicenseBody)

	var license *types.LicenseManagerLicenseInfo
	lm := ctx.Map.Get(*ctx.Map.content().LicenseManager).(*LicenseManager)

	for i, l := range lm.Licenses {
		if l.LicenseKey == req.LicenseKey {
			license = &lm.Licenses[i]
		}
	}

	if license == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "entityId"})
		return body
	}

	for i, r := range m.QueryAssignedLicensesResponse.Returnval {
		if r.EntityId == req.Entity {
			r.AssignedLicense = *license

			if req.EntityDisplayName != "" {
				r.EntityDisplayName = req.EntityDisplayName
			}

			m.QueryAssignedLicensesResponse.Returnval[i] = r

			body.Res = &types.UpdateAssignedLicenseResponse{
				Returnval: r.AssignedLicense,
			}

			break
		}
	}

	return body
}

func licenseInfo(key string, labels []types.KeyValue) types.LicenseManagerLicenseInfo {
	info := EvalLicense

	info.LicenseKey = key
	info.Labels = labels

	return info
}
