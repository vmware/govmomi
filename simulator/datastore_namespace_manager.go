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
	"strings"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type DatastoreNamespaceManager struct {
	mo.DatastoreNamespaceManager
}

func (m *DatastoreNamespaceManager) ConvertNamespacePathToUuidPath(ctx *Context, req *types.ConvertNamespacePathToUuidPath) soap.HasFault {
	body := new(methods.ConvertNamespacePathToUuidPathBody)

	if req.Datacenter == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "datacenterRef"})
		return body
	}

	dc := ctx.Map.Get(*req.Datacenter).(*Datacenter)

	var ds *Datastore
	for _, ref := range dc.Datastore {
		ds = ctx.Map.Get(ref).(*Datastore)
		if strings.HasPrefix(req.NamespaceUrl, ds.Summary.Url) {
			break
		}
		ds = nil
	}

	if ds == nil {
		body.Fault_ = Fault("", &types.InvalidDatastorePath{DatastorePath: req.NamespaceUrl})
		return body
	}

	if !internal.IsDatastoreVSAN(ds.Datastore) {
		body.Fault_ = Fault("", &types.InvalidDatastore{Datastore: &ds.Self})
		return body
	}

	body.Res = &types.ConvertNamespacePathToUuidPathResponse{
		Returnval: req.NamespaceUrl,
	}

	return body
}
