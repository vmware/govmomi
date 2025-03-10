// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type Datastore struct {
	mo.Datastore

	namespace map[string]string // TODO: make thread safe
}

func (ds *Datastore) eventArgument() *types.DatastoreEventArgument {
	return &types.DatastoreEventArgument{
		Datastore:           ds.Self,
		EntityEventArgument: types.EntityEventArgument{Name: ds.Name},
	}
}

func (ds *Datastore) model(m *Model) error {
	info := ds.Info.GetDatastoreInfo()
	u, _ := url.Parse(info.Url)
	if u.Scheme == "ds" {
		// rewrite saved vmfs path to a local temp dir
		u.Path = path.Clean(u.Path)
		parent := strings.ReplaceAll(path.Dir(u.Path), "/", "_")
		name := strings.ReplaceAll(path.Base(u.Path), ":", "_")

		dir, err := m.createTempDir(parent, name)
		if err != nil {
			return err
		}

		info.Url = dir
	}
	return nil
}

// resolve Datastore relative file path to absolute path.
// vSAN top-level directories are named with its vSAN object uuid.
// The directory uuid or friendlyName can be used the various FileManager,
// VirtualDiskManager, etc. methods that have a Datastore path param.
// Note that VirtualDevice file backing paths must use the vSAN uuid.
func (ds *Datastore) resolve(ctx *Context, p string, remove ...bool) string {
	if p == "" || !internal.IsDatastoreVSAN(ds.Datastore) {
		return path.Join(ds.Summary.Url, p)
	}

	rm := len(remove) != 0 && remove[0]
	unlock := ctx.Map.AcquireLock(ctx, ds.Self)
	defer unlock()

	if ds.namespace == nil {
		ds.namespace = make(map[string]string)
	}

	elem := strings.Split(p, "/")
	dir := elem[0]

	_, err := uuid.Parse(dir)
	if err != nil {
		// Translate friendlyName to UUID
		id, ok := ds.namespace[dir]
		if !ok {
			id = uuid.NewString()
			ds.namespace[dir] = id
		}

		elem[0] = id
		p = path.Join(elem...)
		if rm {
			delete(ds.namespace, id)
		}
	} else if rm {
		// UUID was given
		for name, id := range ds.namespace {
			if p == id {
				delete(ds.namespace, name)
				break
			}
		}
	}

	return path.Join(ds.Summary.Url, p)
}

func parseDatastorePath(dsPath string) (*object.DatastorePath, types.BaseMethodFault) {
	var p object.DatastorePath

	if p.FromString(dsPath) {
		return &p, nil
	}

	return nil, &types.InvalidDatastorePath{DatastorePath: dsPath}
}

func (ds *Datastore) RefreshDatastore(*Context, *types.RefreshDatastore) soap.HasFault {
	r := &methods.RefreshDatastoreBody{}

	_, err := os.Stat(ds.Info.GetDatastoreInfo().Url)
	if err != nil {
		r.Fault_ = Fault(err.Error(), &types.HostConfigFault{})
		return r
	}

	ds.Summary.Capacity = int64(units.ByteSize(units.TB)) * int64(len(ds.Host))
	ds.Summary.FreeSpace = ds.Summary.Capacity

	info := ds.Info.GetDatastoreInfo()

	info.FreeSpace = ds.Summary.FreeSpace
	info.MaxMemoryFileSize = ds.Summary.Capacity
	info.MaxFileSize = ds.Summary.Capacity
	info.Timestamp = types.NewTime(time.Now())

	r.Res = &types.RefreshDatastoreResponse{}
	return r
}

func (ds *Datastore) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
	task := CreateTask(ds, "destroy", func(*Task) (types.AnyType, types.BaseMethodFault) {
		if len(ds.Vm) != 0 {
			return nil, &types.ResourceInUse{
				Type: ds.Self.Type,
				Name: ds.Name,
			}
		}

		for _, mount := range ds.Host {
			host := ctx.Map.Get(mount.Key).(*HostSystem)
			ctx.Map.RemoveReference(ctx, host, &host.Datastore, ds.Self)
			parent := hostParent(ctx, &host.HostSystem)
			ctx.Map.RemoveReference(ctx, parent, &parent.Datastore, ds.Self)
		}

		p, _ := asFolderMO(ctx.Map.Get(*ds.Parent))
		folderRemoveChild(ctx, p, ds.Self)

		return nil, nil
	})

	return &methods.Destroy_TaskBody{
		Res: &types.Destroy_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
