// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"io"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vmdk"
)

type FileManager struct {
	mo.FileManager
}

func (f *FileManager) findDatastore(ctx *Context, ref mo.Reference, name string) (*Datastore, types.BaseMethodFault) {
	var refs []types.ManagedObjectReference

	if d, ok := asFolderMO(ref); ok {
		refs = d.ChildEntity
	}
	if p, ok := ref.(*StoragePod); ok {
		refs = p.ChildEntity
	}

	for _, ref := range refs {
		obj := ctx.Map.Get(ref)

		if ds, ok := obj.(*Datastore); ok && ds.Name == name {
			return ds, nil
		}
		if p, ok := obj.(*StoragePod); ok {
			ds, _ := f.findDatastore(ctx, p, name)
			if ds != nil {
				return ds, nil
			}
		}
		if d, ok := asFolderMO(obj); ok {
			ds, _ := f.findDatastore(ctx, d, name)
			if ds != nil {
				return ds, nil
			}
		}
	}

	return nil, &types.InvalidDatastore{Name: name}
}

func (f *FileManager) resolve(ctx *Context, dc *types.ManagedObjectReference, name string, remove ...bool) (string, types.BaseMethodFault) {
	p, fault := parseDatastorePath(name)
	if fault != nil {
		return "", fault
	}

	if dc == nil {
		if ctx.Map.IsESX() {
			dc = &esx.Datacenter.Self
		} else {
			return "", &types.InvalidArgument{InvalidProperty: "dc"}
		}
	}

	folder := ctx.Map.Get(*dc).(*Datacenter).DatastoreFolder

	ds, fault := f.findDatastore(ctx, ctx.Map.Get(folder), p.Datastore)
	if fault != nil {
		return "", fault
	}

	return ds.resolve(ctx, p.Path, remove...), nil
}

func (f *FileManager) fault(name string, err error, fault types.BaseFileFault) types.BaseMethodFault {
	switch {
	case os.IsNotExist(err):
		fault = new(types.FileNotFound)
	case os.IsExist(err):
		fault = new(types.FileAlreadyExists)
	}

	fault.GetFileFault().File = name

	return fault.(types.BaseMethodFault)
}

func (f *FileManager) deleteDatastoreFile(ctx *Context, req *types.DeleteDatastoreFile_Task) types.BaseMethodFault {
	file, fault := f.resolve(ctx, req.Datacenter, req.Name, true)
	if fault != nil {
		return fault
	}

	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return f.fault(file, err, new(types.CannotDeleteFile))
		}
	}

	err = os.RemoveAll(file)
	if err != nil {
		return f.fault(file, err, new(types.CannotDeleteFile))
	}

	return nil
}

func (f *FileManager) DiskDescriptor(ctx *Context, dc *types.ManagedObjectReference, name string) (*vmdk.Descriptor, string, types.BaseMethodFault) {
	path, fault := f.resolve(ctx, dc, name)
	if fault != nil {
		return nil, "", fault
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, "", f.fault(name, err, new(types.FileFault))
	}

	defer file.Close()

	desc, err := vmdk.ParseDescriptor(file)
	if err != nil {
		return nil, "", f.fault(name, err, new(types.FileFault))
	}

	return desc, path, nil
}

func (f *FileManager) SaveDiskDescriptor(ctx *Context, desc *vmdk.Descriptor, path string) types.BaseMethodFault {
	file, err := os.Create(path)
	if err != nil {
		return f.fault(path, err, new(types.FileFault))
	}

	if err = desc.Write(file); err != nil {
		_ = file.Close()
		return f.fault(path, err, new(types.FileFault))
	}

	if err = file.Close(); err != nil {
		return f.fault(path, err, new(types.FileFault))
	}

	return nil
}

func (f *FileManager) DeleteDatastoreFileTask(ctx *Context, req *types.DeleteDatastoreFile_Task) soap.HasFault {
	task := CreateTask(f, "deleteDatastoreFile", func(*Task) (types.AnyType, types.BaseMethodFault) {
		return nil, f.deleteDatastoreFile(ctx, req)
	})

	return &methods.DeleteDatastoreFile_TaskBody{
		Res: &types.DeleteDatastoreFile_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (f *FileManager) MakeDirectory(ctx *Context, req *types.MakeDirectory) soap.HasFault {
	body := &methods.MakeDirectoryBody{}

	name, fault := f.resolve(ctx, req.Datacenter, req.Name)
	if fault != nil {
		body.Fault_ = Fault("", fault)
		return body
	}

	mkdir := os.Mkdir

	if isTrue(req.CreateParentDirectories) {
		mkdir = os.MkdirAll
	}

	err := mkdir(name, 0700)
	if err != nil {
		fault = f.fault(req.Name, err, new(types.CannotCreateFile))
		body.Fault_ = Fault(err.Error(), fault)
		return body
	}

	body.Res = new(types.MakeDirectoryResponse)
	return body
}

func (f *FileManager) moveDatastoreFile(ctx *Context, req *types.MoveDatastoreFile_Task) types.BaseMethodFault {
	src, fault := f.resolve(ctx, req.SourceDatacenter, req.SourceName)
	if fault != nil {
		return fault
	}

	dst, fault := f.resolve(ctx, req.DestinationDatacenter, req.DestinationName)
	if fault != nil {
		return fault
	}

	if !isTrue(req.Force) {
		_, err := os.Stat(dst)
		if err == nil {
			return f.fault(dst, nil, new(types.FileAlreadyExists))
		}
	}

	err := os.Rename(src, dst)
	if err != nil {
		return f.fault(src, err, new(types.CannotAccessFile))
	}

	return nil
}

func (f *FileManager) MoveDatastoreFileTask(ctx *Context, req *types.MoveDatastoreFile_Task) soap.HasFault {
	task := CreateTask(f, "moveDatastoreFile", func(*Task) (types.AnyType, types.BaseMethodFault) {
		return nil, f.moveDatastoreFile(ctx, req)
	})

	return &methods.MoveDatastoreFile_TaskBody{
		Res: &types.MoveDatastoreFile_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (f *FileManager) copyDatastoreFile(ctx *Context, req *types.CopyDatastoreFile_Task) types.BaseMethodFault {
	src, fault := f.resolve(ctx, req.SourceDatacenter, req.SourceName)
	if fault != nil {
		return fault
	}

	dst, fault := f.resolve(ctx, req.DestinationDatacenter, req.DestinationName)
	if fault != nil {
		return fault
	}

	if !isTrue(req.Force) {
		_, err := os.Stat(dst)
		if err == nil {
			return f.fault(dst, nil, new(types.FileAlreadyExists))
		}
	}

	r, err := os.Open(filepath.Clean(src))
	if err != nil {
		return f.fault(dst, err, new(types.CannotAccessFile))
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return f.fault(dst, err, new(types.CannotCreateFile))
	}
	defer w.Close()

	if _, err = io.Copy(w, r); err != nil {
		return f.fault(dst, err, new(types.CannotCreateFile))
	}

	return nil
}

func (f *FileManager) CopyDatastoreFileTask(ctx *Context, req *types.CopyDatastoreFile_Task) soap.HasFault {
	task := CreateTask(f, "copyDatastoreFile", func(*Task) (types.AnyType, types.BaseMethodFault) {
		return nil, f.copyDatastoreFile(ctx, req)
	})

	return &methods.CopyDatastoreFile_TaskBody{
		Res: &types.CopyDatastoreFile_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
