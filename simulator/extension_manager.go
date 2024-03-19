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
	"time"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var ExtensionList = []types.Extension{
	{
		Description: &types.Description{
			Label:   "vcsim",
			Summary: "Go vCenter simulator",
		},
		Key:         "com.vmware.govmomi.simulator",
		Company:     "VMware, Inc.",
		Type:        "",
		Version:     "0.37.0",
		SubjectName: "",
		Server:      nil,
		Client:      nil,
		TaskList: []types.ExtensionTaskTypeInfo{
			{
				TaskID: "com.vmware.govmomi.simulator.test",
			},
		},
		EventList:              nil,
		FaultList:              nil,
		PrivilegeList:          nil,
		ResourceList:           nil,
		LastHeartbeatTime:      time.Now(),
		HealthInfo:             (*types.ExtensionHealthInfo)(nil),
		OvfConsumerInfo:        (*types.ExtensionOvfConsumerInfo)(nil),
		ExtendedProductInfo:    (*types.ExtExtendedProductInfo)(nil),
		ManagedEntityInfo:      nil,
		ShownInSolutionManager: types.NewBool(false),
		SolutionManagerInfo:    (*types.ExtSolutionManagerInfo)(nil),
	},
}

type ExtensionManager struct {
	mo.ExtensionManager
}

func (m *ExtensionManager) init(r *Registry) {
	if r.IsVPX() && len(m.ExtensionList) == 0 {
		m.ExtensionList = ExtensionList
	}
}

func (m *ExtensionManager) FindExtension(ctx *Context, req *types.FindExtension) soap.HasFault {
	body := &methods.FindExtensionBody{
		Res: new(types.FindExtensionResponse),
	}

	for _, x := range m.ExtensionList {
		if x.Key == req.ExtensionKey {
			body.Res.Returnval = &x
			break
		}
	}

	return body
}

func (m *ExtensionManager) RegisterExtension(ctx *Context, req *types.RegisterExtension) soap.HasFault {
	body := &methods.RegisterExtensionBody{}

	for _, x := range m.ExtensionList {
		if x.Key == req.Extension.Key {
			body.Fault_ = Fault("", &types.InvalidArgument{
				InvalidProperty: "extension.key",
			})
			return body
		}
	}

	body.Res = new(types.RegisterExtensionResponse)
	m.ExtensionList = append(m.ExtensionList, req.Extension)

	return body
}

func (m *ExtensionManager) UnregisterExtension(ctx *Context, req *types.UnregisterExtension) soap.HasFault {
	body := &methods.UnregisterExtensionBody{}

	for i, x := range m.ExtensionList {
		if x.Key == req.ExtensionKey {
			m.ExtensionList = append(m.ExtensionList[:i], m.ExtensionList[i+1:]...)

			body.Res = new(types.UnregisterExtensionResponse)
			return body
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}

func (m *ExtensionManager) UpdateExtension(ctx *Context, req *types.UpdateExtension) soap.HasFault {
	body := &methods.UpdateExtensionBody{}

	for i, x := range m.ExtensionList {
		if x.Key == req.Extension.Key {
			m.ExtensionList[i] = req.Extension

			body.Res = new(types.UpdateExtensionResponse)
			return body
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}

func (m *ExtensionManager) SetExtensionCertificate(ctx *Context, req *types.SetExtensionCertificate) soap.HasFault {
	body := &methods.SetExtensionCertificateBody{}

	for _, x := range m.ExtensionList {
		if x.Key == req.ExtensionKey {
			// TODO: save req.CertificatePem for use with SessionManager.LoginExtensionByCertificate()

			body.Res = new(types.SetExtensionCertificateResponse)
			return body
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}
