// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
	vslmtypes "github.com/vmware/govmomi/vslm/types"
)

// Manager provides a layer for switching between Virtual Storage Object Manager (VSOM)
// and Virtual Storage Lifecycle Manager (VSLM). The majority of VSOM methods require a
// Datastore param, and most VSLM methods do not as it uses the "Global Catalog".
// VSOM was introduced in vSphere 6.5 (11/2016) and VSLM in vSphere 6.7 U2 (04/2019).
// The govc disk commands were introduced prior to 6.7 U2 and continue to use VSOM
// when an optional Datastore (-ds flag) is provided. Otherwise, VSLM global methods
// are used when a Datastore is not specified.
// Also note that VSOM methods can be used when connected directly to an ESX hosts,
// but VSLM global methods require vCenter.
// A disk managed by these methods are also known as a "First Class Disk" (FCD).
type Manager struct {
	Client              *vim25.Client
	Datastore           *object.Datastore
	ObjectManager       *vslm.ObjectManager
	GlobalObjectManager *vslm.GlobalObjectManager
}

var (
	// vslm does not have a PropertyCollector, the Wait func polls via VslmQueryInfo
	taskTimeout = time.Hour

	// Some methods require a Datastore param regardless using vsom or vslm.
	// By default we use Global vslm for those methods, use vsom when this is "false".
	useGlobal = os.Getenv("GOVC_GLOBAL_VSLM") != "false"
)

func NewManagerFromFlag(ctx context.Context, cmd *flags.DatastoreFlag) (*Manager, error) {
	c, err := cmd.Client()
	if err != nil {
		return nil, err
	}

	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return nil, err
	}

	return NewManager(ctx, c, ds)
}

func NewManager(ctx context.Context, c *vim25.Client, ds *object.Datastore) (*Manager, error) {
	m := &Manager{
		Client:        c,
		Datastore:     ds,
		ObjectManager: vslm.NewObjectManager(c),
	}

	if ds == nil {
		if err := m.initGlobalManager(ctx); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Manager) initGlobalManager(ctx context.Context) error {
	if m.GlobalObjectManager == nil {
		vc, err := vslm.NewClient(ctx, m.Client)
		if err != nil {
			return err
		}
		m.GlobalObjectManager = vslm.NewGlobalObjectManager(vc)
	}
	return nil
}

func (m *Manager) CreateDisk(ctx context.Context, spec types.VslmCreateSpec) (*types.VStorageObject, error) {
	if useGlobal && m.Client.IsVC() {
		if err := m.initGlobalManager(ctx); err != nil {
			return nil, err
		}
		task, err := m.GlobalObjectManager.CreateDisk(ctx, spec)
		if err != nil {
			return nil, err
		}
		res, err := task.Wait(ctx, taskTimeout)
		if err != nil {
			return nil, err
		}

		obj := res.(types.VStorageObject)
		return &obj, nil
	}

	task, err := m.ObjectManager.CreateDisk(ctx, spec)
	if err != nil {
		return nil, err
	}
	res, err := task.WaitForResult(ctx)
	if err != nil {
		return nil, err
	}
	obj := res.Result.(types.VStorageObject)
	return &obj, nil
}

func (m *Manager) Delete(ctx context.Context, id string) error {
	if m.Datastore != nil {
		task, err := m.ObjectManager.Delete(ctx, m.Datastore, id)
		if err != nil {
			return err
		}
		return task.Wait(ctx)
	}

	task, err := m.GlobalObjectManager.Delete(ctx, types.ID{Id: id})
	if err != nil {
		return err
	}
	_, err = task.Wait(ctx, taskTimeout)
	return err
}

func (m *Manager) Retrieve(ctx context.Context, id string) (*types.VStorageObject, error) {
	if m.Datastore != nil {
		return m.ObjectManager.Retrieve(ctx, m.Datastore, id)
	}
	return m.GlobalObjectManager.Retrieve(ctx, types.ID{Id: id})
}

func (m *Manager) List(ctx context.Context, qs ...vslmtypes.VslmVsoVStorageObjectQuerySpec) ([]types.ID, error) {
	if m.Datastore != nil {
		return m.ObjectManager.List(ctx, m.Datastore)
	}

	res, err := m.GlobalObjectManager.List(ctx, qs...)
	if err != nil {
		return nil, err
	}

	return res.Id, nil
}

func (m *Manager) RegisterDisk(ctx context.Context, path, name string) (*types.VStorageObject, error) {
	if useGlobal && m.Client.IsVC() {
		if err := m.initGlobalManager(ctx); err != nil {
			return nil, err
		}
		return m.GlobalObjectManager.RegisterDisk(ctx, path, name) // VslmRegisterDisk
	}
	return m.ObjectManager.RegisterDisk(ctx, path, name) // RegisterDisk
}

func (m *Manager) CreateSnapshot(ctx context.Context, id, desc string) (types.ID, error) {
	if m.Datastore != nil {
		task, err := m.ObjectManager.CreateSnapshot(ctx, m.Datastore, id, desc)
		if err != nil {
			return types.ID{}, err
		}
		res, err := task.WaitForResult(ctx, nil)
		if err != nil {
			return types.ID{}, err
		}
		return res.Result.(types.ID), nil
	}

	task, err := m.GlobalObjectManager.CreateSnapshot(ctx, types.ID{Id: id}, desc)
	if err != nil {
		return types.ID{}, err
	}
	res, err := task.Wait(ctx, taskTimeout)
	if err != nil {
		return types.ID{}, err
	}
	return res.(types.ID), err
}

func (m *Manager) DeleteSnapshot(ctx context.Context, id, sid string) error {
	if m.Datastore != nil {
		task, err := m.ObjectManager.DeleteSnapshot(ctx, m.Datastore, id, sid)
		if err != nil {
			return err
		}
		return task.Wait(ctx)
	}

	task, err := m.GlobalObjectManager.DeleteSnapshot(ctx, types.ID{Id: id}, types.ID{Id: sid})
	if err != nil {
		return err
	}
	_, err = task.Wait(ctx, taskTimeout)
	return err
}

func (m *Manager) RetrieveSnapshotInfo(ctx context.Context, id string) ([]types.VStorageObjectSnapshotInfoVStorageObjectSnapshot, error) {
	if m.Datastore != nil {
		info, err := m.ObjectManager.RetrieveSnapshotInfo(ctx, m.Datastore, id)
		if err != nil {
			return nil, err
		}
		return info.Snapshots, nil
	}

	return m.GlobalObjectManager.RetrieveSnapshotInfo(ctx, types.ID{Id: id})
}

func (m *Manager) AttachTag(ctx context.Context, id string, tag types.VslmTagEntry) error {
	if useGlobal && m.Client.IsVC() {
		if err := m.initGlobalManager(ctx); err != nil {
			return err
		}
		// TODO: use types.VslmTagEntry
		return m.GlobalObjectManager.AttachTag(ctx, types.ID{Id: id}, tag.ParentCategoryName, tag.TagName)
	}
	return m.ObjectManager.AttachTag(ctx, id, tag)
}

func (m *Manager) DetachTag(ctx context.Context, id string, tag types.VslmTagEntry) error {
	if useGlobal && m.Client.IsVC() {
		if err := m.initGlobalManager(ctx); err != nil {
			return err
		}
		// TODO: use types.VslmTagEntry
		return m.GlobalObjectManager.DetachTag(ctx, types.ID{Id: id}, tag.ParentCategoryName, tag.TagName)
	}
	return m.ObjectManager.DetachTag(ctx, id, tag)
}

func (m *Manager) ListAttachedObjects(ctx context.Context, category, tag string) ([]types.ID, error) {
	if m.Datastore != nil {
		// ListVStorageObjectsAttachedToTag
		return m.ObjectManager.ListAttachedObjects(ctx, category, tag)
	}

	// VslmListVStorageObjectsAttachedToTag
	return m.GlobalObjectManager.ListAttachedObjects(ctx, category, tag)
}

func (m *Manager) ListAttachedTags(ctx context.Context, id string) ([]types.VslmTagEntry, error) {
	if m.Datastore != nil {
		// ListTagsAttachedToVStorageObject
		return m.ObjectManager.ListAttachedTags(ctx, id)
	}

	// VslmListTagsAttachedToVStorageObject
	return m.GlobalObjectManager.ListAttachedTags(ctx, types.ID{Id: id})
}

func (m *Manager) ReconcileDatastoreInventory(ctx context.Context) error {
	if m.Datastore != nil {
		// ReconcileDatastoreInventory_Task
		task, err := m.ObjectManager.ReconcileDatastoreInventory(ctx, m.Datastore)
		if err != nil {
			return err
		}
		return task.Wait(ctx)
	}

	// VslmReconcileDatastoreInventory_Task also requires a Datastore
	return errors.New("-R requires -ds") // TODO
}

func (m *Manager) AttachDisk(ctx context.Context, vm *object.VirtualMachine, id string) error {
	if m.Datastore != nil {
		return vm.AttachDisk(ctx, id, m.Datastore, 0, nil)
	}

	task, err := m.GlobalObjectManager.AttachDisk(ctx, types.ID{Id: id}, vm, 0, nil)
	if err != nil {
		return err
	}
	_, err = task.Wait(ctx, taskTimeout)
	return err
}
