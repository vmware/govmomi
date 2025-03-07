// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"path"
	"strings"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/object"
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

	dc, ok := ctx.Map.Get(*req.Datacenter).(*Datacenter)
	if !ok {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: *req.Datacenter})
		return body
	}

	var ds *Datastore
	ns := ""

	for _, ref := range dc.Datastore {
		ds = ctx.Map.Get(ref).(*Datastore)
		ns = strings.TrimPrefix(req.NamespaceUrl, ds.Summary.Url)
		if ns != req.NamespaceUrl {
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
		Returnval: ds.resolve(ctx, ns),
	}

	return body
}

func (m *DatastoreNamespaceManager) CreateDirectory(ctx *Context, req *types.CreateDirectory) soap.HasFault {
	body := new(methods.CreateDirectoryBody)

	ds, ok := ctx.Map.Get(req.Datastore).(*Datastore)
	if !ok {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: req.Datastore})
		return body
	}

	if !internal.IsDatastoreVSAN(ds.Datastore) {
		body.Fault_ = Fault("", &types.CannotCreateFile{
			FileFault: types.FileFault{
				File: "Datastore not supported for directory creation by DatastoreNamespaceManager",
			},
		})
		return body
	}

	if !isValidFileName(req.DisplayName) {
		body.Fault_ = Fault("", &types.InvalidDatastorePath{DatastorePath: req.DisplayName})
		return body
	}

	p := object.DatastorePath{
		Datastore: ds.Name,
		Path:      req.DisplayName,
	}

	dc := ctx.Map.getEntityDatacenter(ds)

	fm := ctx.Map.FileManager()

	fault := fm.MakeDirectory(ctx, &types.MakeDirectory{
		This:       fm.Self,
		Name:       p.String(),
		Datacenter: &dc.Self,
	})

	if fault.Fault() != nil {
		body.Fault_ = fault.Fault()
	} else {
		body.Res = &types.CreateDirectoryResponse{
			Returnval: ds.resolve(ctx, req.DisplayName),
		}
	}

	return body
}

func (m *DatastoreNamespaceManager) DeleteDirectory(ctx *Context, req *types.DeleteDirectory) soap.HasFault {
	body := new(methods.DeleteDirectoryBody)

	if req.Datacenter == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "datacenterRef"})
		return body
	}

	dc, ok := ctx.Map.Get(*req.Datacenter).(*Datacenter)
	if !ok {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: *req.Datacenter})
		return body
	}

	var ds *Datastore
	for _, ref := range dc.Datastore {
		ds = ctx.Map.Get(ref).(*Datastore)
		if strings.HasPrefix(req.DatastorePath, ds.Summary.Url) {
			break
		}
		ds = nil
	}

	if ds == nil || strings.Contains(req.DatastorePath, "..") {
		body.Fault_ = Fault("", &types.InvalidDatastorePath{DatastorePath: req.DatastorePath})
		return body
	}

	if !internal.IsDatastoreVSAN(ds.Datastore) {
		body.Fault_ = Fault("", &types.FileFault{
			File: "Datastore not supported for directory deletion by DatastoreNamespaceManager",
		})
		return body
	}

	name := &object.DatastorePath{
		Datastore: ds.Name,
		Path:      path.Base(req.DatastorePath),
	}

	fm := ctx.Map.FileManager()

	fault := fm.deleteDatastoreFile(ctx, &types.DeleteDatastoreFile_Task{
		Name:       name.String(),
		Datacenter: req.Datacenter,
	})

	if fault != nil {
		body.Fault_ = Fault("", fault)
	} else {
		body.Res = new(types.DeleteDirectoryResponse)
	}

	return body
}

func isValidFileName(s string) bool {
	return !strings.Contains(s, "/") &&
		!strings.Contains(s, "\\") &&
		!strings.Contains(s, "..")
}
