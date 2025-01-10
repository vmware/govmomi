// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	vimx "github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
	"github.com/vmware/govmomi/vslm/methods"
	"github.com/vmware/govmomi/vslm/types"
)

var content = types.VslmServiceInstanceContent{
	AboutInfo: types.VslmAboutInfo{
		Name:         "VMware Virtual Storage Lifecycle Manager Service",
		FullName:     "VMware Virtual Storage Lifecycle Manager Service 1.0.0",
		Vendor:       "VMware, Inc.",
		ApiVersion:   "1.0.0",
		InstanceUuid: "31c68687-4f1e-4247-8158-f31d1ce95bbe",
	},
	SessionManager:          vim.ManagedObjectReference{Type: "VslmSessionManager", Value: "SessionManager"},
	VStorageObjectManager:   vim.ManagedObjectReference{Type: "VslmVStorageObjectManager", Value: "VStorageObjectManager"},
	StorageLifecycleManager: vim.ManagedObjectReference{Type: "VslmStorageLifecycleManager", Value: "StorageLifecycleManager"},
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
	r.Namespace = vslm.Namespace
	r.Path = vslm.Path
	r.Cookie = simulator.SOAPCookie

	r.Put(&ServiceInstance{
		ManagedObjectReference: vslm.ServiceInstance,
		Content:                content,
	})

	r.Put(&VStorageObjectManager{
		ManagedObjectReference: content.VStorageObjectManager,
	})

	return r
}

type ServiceInstance struct {
	vim.ManagedObjectReference

	Content types.VslmServiceInstanceContent
}

func (s *ServiceInstance) RetrieveContent(_ *types.RetrieveContent) soap.HasFault {
	return &methods.RetrieveContentBody{
		Res: &types.RetrieveContentResponse{
			Returnval: s.Content,
		},
	}
}

// VStorageObjectManager APIs manage First Class Disks (FCDs) using the "Global Catalog".
// The majority of methods in this API simulator dispatch to VcenterVStorageObjectManager methods,
// after looking up a disk's Datastore. Along with 'VslmTask', who's methods proxy to a vim25 'Task'.
type VStorageObjectManager struct {
	vim.ManagedObjectReference
}

func (m *VStorageObjectManager) VslmListVStorageObjectForSpec(ctx *simulator.Context, req *types.VslmListVStorageObjectForSpec) soap.HasFault {
	res := new(types.VslmVsoVStorageObjectQueryResult)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	res.AllRecordsReturned = true // TODO: req.MaxResult

	for _, objs := range vsom.Catalog() {
		// TODO: filter req.Query
		for id, obj := range objs {
			res.Id = append(res.Id, id)

			vso := types.VslmVsoVStorageObjectResult{
				Id:           id,
				Name:         obj.Config.Name,
				CapacityInMB: obj.Config.CapacityInMB,
				CreateTime:   &obj.Config.CreateTime,
			}
			res.QueryResults = append(res.QueryResults, vso)
		}
	}

	return &methods.VslmListVStorageObjectForSpecBody{
		Res: &types.VslmListVStorageObjectForSpecResponse{
			Returnval: res,
		},
	}
}

func (m *VStorageObjectManager) VslmRetrieveVStorageObject(ctx *simulator.Context, req *types.VslmRetrieveVStorageObject) soap.HasFault {
	body := new(methods.VslmRetrieveVStorageObjectBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	for _, objs := range vsom.Catalog() {
		for id, obj := range objs {
			if id == req.Id {
				body.Res = &types.VslmRetrieveVStorageObjectResponse{
					Returnval: obj.VStorageObject,
				}
				return body
			}
		}
	}

	body.Fault_ = simulator.Fault("", &vim.InvalidArgument{InvalidProperty: "VolumeId"})

	return body
}

func (m *VStorageObjectManager) VslmReconcileDatastoreInventoryTask(ctx *simulator.Context, req *types.VslmReconcileDatastoreInventory_Task) soap.HasFault {
	body := new(methods.VslmReconcileDatastoreInventory_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.ReconcileDatastoreInventoryTask(vctx, &vim.ReconcileDatastoreInventory_Task{
		This:      vsom.Self,
		Datastore: req.Datastore,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		ref := val.(*vimx.ReconcileDatastoreInventory_TaskBody).Res.Returnval

		body.Res = &types.VslmReconcileDatastoreInventory_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmRegisterDisk(ctx *simulator.Context, req *types.VslmRegisterDisk) soap.HasFault {
	body := new(methods.VslmRegisterDiskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.RegisterDisk(ctx, &vim.RegisterDisk{
		This: vsom.Self,
		Path: req.Path,
		Name: req.Name,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		body.Res = &types.VslmRegisterDiskResponse{
			Returnval: val.(*vimx.RegisterDiskBody).Res.Returnval,
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmCreateDiskTask(ctx *simulator.Context, req *types.VslmCreateDisk_Task) soap.HasFault {
	body := new(methods.VslmCreateDisk_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.CreateDiskTask(vctx, &vim.CreateDisk_Task{
		This: vsom.Self,
		Spec: req.Spec,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		ref := val.(*vimx.CreateDisk_TaskBody).Res.Returnval

		body.Res = &types.VslmCreateDisk_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmDeleteVStorageObjectTask(ctx *simulator.Context, req *types.VslmDeleteVStorageObject_Task) soap.HasFault {
	body := new(methods.VslmDeleteVStorageObject_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.DeleteVStorageObjectTask(vctx, &vim.DeleteVStorageObject_Task{
		This:      vsom.Self,
		Id:        req.Id,
		Datastore: m.ds(vsom, req.Id),
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
		return body
	} else {
		ref := val.(*vimx.DeleteVStorageObject_TaskBody).Res.Returnval

		body.Res = &types.VslmDeleteVStorageObject_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmRetrieveSnapshotInfo(ctx *simulator.Context, req *types.VslmRetrieveSnapshotInfo) soap.HasFault {
	body := new(methods.VslmRetrieveSnapshotInfoBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	var vso *simulator.VStorageObject
	for _, objs := range vsom.Catalog() {
		for id, obj := range objs {
			if id == req.Id {
				vso = obj
				break
			}
		}
	}

	if vso == nil {
		body.Fault_ = simulator.Fault("", &vim.InvalidArgument{InvalidProperty: "VolumeId"})
	} else {
		body.Res = &types.VslmRetrieveSnapshotInfoResponse{
			Returnval: vso.VStorageObjectSnapshotInfo,
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmCreateSnapshotTask(ctx *simulator.Context, req *types.VslmCreateSnapshot_Task) soap.HasFault {
	body := new(methods.VslmCreateSnapshot_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.VStorageObjectCreateSnapshotTask(vctx, &vim.VStorageObjectCreateSnapshot_Task{
		This:        vsom.Self,
		Id:          req.Id,
		Description: req.Description,
		Datastore:   m.ds(vsom, req.Id),
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		ref := val.(*vimx.VStorageObjectCreateSnapshot_TaskBody).Res.Returnval

		body.Res = &types.VslmCreateSnapshot_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmDeleteSnapshotTask(ctx *simulator.Context, req *types.VslmDeleteSnapshot_Task) soap.HasFault {
	body := new(methods.VslmDeleteSnapshot_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.DeleteSnapshotTask(vctx, &vim.DeleteSnapshot_Task{
		This:       vsom.Self,
		Id:         req.Id,
		SnapshotId: req.SnapshotId,
		Datastore:  m.ds(vsom, req.Id),
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		ref := val.(*vimx.DeleteSnapshot_TaskBody).Res.Returnval

		body.Res = &types.VslmDeleteSnapshot_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmAttachTagToVStorageObject(ctx *simulator.Context, req *types.VslmAttachTagToVStorageObject) soap.HasFault {
	body := new(methods.VslmAttachTagToVStorageObjectBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.AttachTagToVStorageObject(vctx, &vim.AttachTagToVStorageObject{
		This:     vsom.Self,
		Id:       req.Id,
		Category: req.Category,
		Tag:      req.Tag,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		body.Res = new(types.VslmAttachTagToVStorageObjectResponse)
	}

	return body
}

func (m *VStorageObjectManager) VslmDetachTagFromVStorageObject(ctx *simulator.Context, req *types.VslmDetachTagFromVStorageObject) soap.HasFault {
	body := new(methods.VslmDetachTagFromVStorageObjectBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.DetachTagFromVStorageObject(vctx, &vim.DetachTagFromVStorageObject{
		This:     vsom.Self,
		Id:       req.Id,
		Category: req.Category,
		Tag:      req.Tag,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		body.Res = new(types.VslmDetachTagFromVStorageObjectResponse)
	}

	return body
}

func (m *VStorageObjectManager) VslmListVStorageObjectsAttachedToTag(ctx *simulator.Context, req *types.VslmListVStorageObjectsAttachedToTag) soap.HasFault {
	body := new(methods.VslmListVStorageObjectsAttachedToTagBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.ListVStorageObjectsAttachedToTag(vctx, &vim.ListVStorageObjectsAttachedToTag{
		This:     vsom.Self,
		Category: req.Category,
		Tag:      req.Tag,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		body.Res = &types.VslmListVStorageObjectsAttachedToTagResponse{
			Returnval: val.(*vimx.ListVStorageObjectBody).Res.Returnval,
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmListTagsAttachedToVStorageObject(ctx *simulator.Context, req *types.VslmListTagsAttachedToVStorageObject) soap.HasFault {
	body := new(methods.VslmListTagsAttachedToVStorageObjectBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	val := vsom.ListTagsAttachedToVStorageObject(vctx, &vim.ListTagsAttachedToVStorageObject{
		This: vsom.Self,
		Id:   req.Id,
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		body.Res = &types.VslmListTagsAttachedToVStorageObjectResponse{
			Returnval: val.(*vimx.ListTagsAttachedToVStorageObjectBody).Res.Returnval,
		}
	}

	return body
}

func (m *VStorageObjectManager) VslmAttachDiskTask(ctx *simulator.Context, req *types.VslmAttachDisk_Task) soap.HasFault {
	body := new(methods.VslmAttachDisk_TaskBody)

	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	vm, ok := vctx.Map.Get(req.Vm).(*simulator.VirtualMachine)
	if !ok {
		body.Fault_ = simulator.Fault("", &vim.ManagedObjectNotFound{Obj: req.Vm})
		return body
	}

	var val soap.HasFault

	vctx.WithLock(vm, func() {
		val = vm.AttachDiskTask(vctx, &vim.AttachDisk_Task{
			This:          vm.Self,
			Datastore:     m.ds(vsom, req.Id),
			DiskId:        req.Id,
			ControllerKey: req.ControllerKey,
			UnitNumber:    req.UnitNumber,
		})
	})

	if val.Fault() != nil {
		body.Fault_ = val.Fault()
	} else {
		ref := val.(*vimx.AttachDisk_TaskBody).Res.Returnval

		body.Res = &types.VslmAttachDisk_TaskResponse{
			Returnval: newVslmTask(ctx, ref),
		}
	}

	return body
}

// VslmTask methods are just a proxy to vim25 Task methods
type VslmTask struct {
	vim.ManagedObjectReference
}

func newVslmTask(ctx *simulator.Context, ref vim.ManagedObjectReference) vim.ManagedObjectReference {
	task := &VslmTask{
		ManagedObjectReference: vim.ManagedObjectReference{
			Type:  "VslmTask",
			Value: ref.Value,
		},
	}

	return ctx.Map.Put(task).Reference()
}

func (p *VslmTask) VslmQueryInfo(ctx *simulator.Context, req *types.VslmQueryInfo) soap.HasFault {
	body := new(methods.VslmQueryInfoBody)

	task, fault := p.task(ctx, req.This)
	if fault != nil {
		body.Fault_ = fault
	} else {
		info := types.VslmTaskInfo{
			Key:           p.Value,
			Task:          p.ManagedObjectReference,
			DescriptionId: "com.vmware.cns.vslm.tasks.createDisk",
			State:         types.VslmTaskInfoState(task.State),
			Error:         task.Error,
			Result:        task.Result,
			QueueTime:     task.QueueTime,
			StartTime:     task.StartTime,
			CompleteTime:  task.CompleteTime,
		}

		body.Res = &types.VslmQueryInfoResponse{Returnval: info}
	}

	return body
}

func (p *VslmTask) VslmQueryTaskResult(ctx *simulator.Context, req *types.VslmQueryTaskResult) soap.HasFault {
	body := new(methods.VslmQueryTaskResultBody)

	task, fault := p.task(ctx, req.This)
	if fault != nil {
		body.Fault_ = fault
	} else {
		body.Res = &types.VslmQueryTaskResultResponse{Returnval: task.Result}
	}

	return body
}

func (*VslmTask) task(ctx *simulator.Context, ref vim.ManagedObjectReference) (vim.TaskInfo, *soap.Fault) {
	ctx = ctx.For(vim25.Path)

	ref.Type = "Task"

	if task, ok := ctx.Map.Get(ref).(*simulator.Task); ok {
		return task.Info, nil
	}

	return vim.TaskInfo{}, simulator.Fault("", &vim.ManagedObjectNotFound{Obj: ref})
}

func (*VStorageObjectManager) ds(vsom *simulator.VcenterVStorageObjectManager, reqID vim.ID) vim.ManagedObjectReference {
	for ds, objs := range vsom.Catalog() {
		for id := range objs {
			if id == reqID {
				return ds
			}
		}
	}

	// vsom calls will fault as they would when ID is NotFound
	return vim.ManagedObjectReference{Type: "Datastore"}
}
