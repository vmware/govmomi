// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"os"
	"path"
	"strings"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type HostDatastoreBrowser struct {
	mo.HostDatastoreBrowser
}

type searchDatastore struct {
	*HostDatastoreBrowser

	DatastorePath string
	SearchSpec    *types.HostDatastoreBrowserSearchSpec

	res []types.HostDatastoreBrowserSearchResults

	recurse bool
}

func (s *searchDatastore) addFile(fname string, file os.FileInfo, res *types.HostDatastoreBrowserSearchResults) {
	details := s.SearchSpec.Details
	if details == nil {
		details = new(types.FileQueryFlags)
	}

	name := file.Name()

	info := types.FileInfo{
		Path:         name,
		FriendlyName: fname,
	}

	var finfo types.BaseFileInfo = &info

	if details.FileSize {
		info.FileSize = file.Size()
	}

	if details.Modification {
		mtime := file.ModTime()
		info.Modification = &mtime
	}

	if isTrue(details.FileOwner) {
		// Assume for now this process created all files in the datastore
		user := os.Getenv("USER")

		info.Owner = user
	}

	if file.IsDir() {
		finfo = &types.FolderFileInfo{FileInfo: info}
	} else if details.FileType {
		switch path.Ext(name) {
		case ".img":
			finfo = &types.FloppyImageFileInfo{FileInfo: info}
		case ".iso":
			finfo = &types.IsoImageFileInfo{FileInfo: info}
		case ".log":
			finfo = &types.VmLogFileInfo{FileInfo: info}
		case ".nvram":
			finfo = &types.VmNvramFileInfo{FileInfo: info}
		case ".vmdk":
			// TODO: lookup device to set other fields
			finfo = &types.VmDiskFileInfo{FileInfo: info}
		case ".vmx":
			finfo = &types.VmConfigFileInfo{FileInfo: info}
		}
	}

	res.File = append(res.File, finfo)
}

func (s *searchDatastore) queryMatch(file os.FileInfo) bool {
	if len(s.SearchSpec.Query) == 0 {
		return true
	}

	name := file.Name()
	ext := path.Ext(name)

	for _, q := range s.SearchSpec.Query {
		switch q.(type) {
		case *types.FileQuery:
			return true
		case *types.FolderFileQuery:
			if file.IsDir() {
				return true
			}
		case *types.FloppyImageFileQuery:
			if ext == ".img" {
				return true
			}
		case *types.IsoImageFileQuery:
			if ext == ".iso" {
				return true
			}
		case *types.VmConfigFileQuery:
			if ext == ".vmx" {
				// TODO: check Filter and Details fields
				return true
			}
		case *types.VmDiskFileQuery:
			if ext == ".vmdk" {
				// TODO: check Filter and Details fields
				return !strings.HasSuffix(name, "-flat.vmdk")
			}
		case *types.VmLogFileQuery:
			if ext == ".log" {
				return strings.HasPrefix(name, "vmware")
			}
		case *types.VmNvramFileQuery:
			if ext == ".nvram" {
				return true
			}
		case *types.VmSnapshotFileQuery:
			if ext == ".vmsn" {
				return true
			}
		}
	}

	return false
}

func friendlyName(ctx *Context, root bool, ds *Datastore, p string) string {
	if !root || p == "" || !internal.IsDatastoreVSAN(ds.Datastore) {
		return ""
	}

	unlock := ctx.Map.AcquireLock(ctx, ds.Self)
	defer unlock()

	if ds.namespace == nil {
		return ""
	}

	for name, id := range ds.namespace {
		if p == id {
			return name
		}
	}

	return ""
}

func (s *searchDatastore) search(ctx *Context, ds *Datastore, folder string, dir string, root bool) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		tracef("search %s: %s", dir, err)
		return err
	}

	res := types.HostDatastoreBrowserSearchResults{
		Datastore:  &ds.Self,
		FolderPath: folder,
	}

	for _, file := range files {
		name := file.Name()
		info, _ := file.Info()
		if s.queryMatch(info) {
			for _, m := range s.SearchSpec.MatchPattern {
				if ok, _ := path.Match(m, name); ok {
					s.addFile(friendlyName(ctx, root, ds, name), info, &res)
					break
				}
			}
		}

		if s.recurse && file.IsDir() {
			_ = s.search(ctx, ds, path.Join(folder, name), path.Join(dir, name), false)
		}
	}

	s.res = append(s.res, res)

	return nil
}

func (s *searchDatastore) Run(task *Task) (types.AnyType, types.BaseMethodFault) {
	p, fault := parseDatastorePath(s.DatastorePath)
	if fault != nil {
		return nil, fault
	}

	ref := task.ctx.Map.FindByName(p.Datastore, s.Datastore)
	if ref == nil {
		return nil, &types.InvalidDatastore{Name: p.Datastore}
	}

	ds := ref.(*Datastore)

	task.ctx.WithLock(task, func() {
		task.Info.Entity = &ds.Self // TODO: CreateTask() should require mo.Entity, rather than mo.Reference
		task.Info.EntityName = ds.Name
	})

	dir := ds.resolve(task.ctx, p.Path)

	err := s.search(task.ctx, ds, s.DatastorePath, dir, p.Path == "")
	if err != nil {
		ff := types.FileFault{
			File: p.Path,
		}

		if os.IsNotExist(err) {
			return nil, &types.FileNotFound{FileFault: ff}
		}

		return nil, &types.InvalidArgument{InvalidProperty: p.Path}
	}

	if s.recurse {
		return types.ArrayOfHostDatastoreBrowserSearchResults{
			HostDatastoreBrowserSearchResults: s.res,
		}, nil
	}

	return s.res[0], nil
}

func (b *HostDatastoreBrowser) SearchDatastoreTask(ctx *Context, s *types.SearchDatastore_Task) soap.HasFault {
	task := NewTask(&searchDatastore{
		HostDatastoreBrowser: b,
		DatastorePath:        s.DatastorePath,
		SearchSpec:           s.SearchSpec,
	})

	return &methods.SearchDatastore_TaskBody{
		Res: &types.SearchDatastore_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (b *HostDatastoreBrowser) SearchDatastoreSubFoldersTask(ctx *Context, s *types.SearchDatastoreSubFolders_Task) soap.HasFault {
	task := NewTask(&searchDatastore{
		HostDatastoreBrowser: b,
		DatastorePath:        s.DatastorePath,
		SearchSpec:           s.SearchSpec,
		recurse:              true,
	})

	return &methods.SearchDatastoreSubFolders_TaskBody{
		Res: &types.SearchDatastoreSubFolders_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
