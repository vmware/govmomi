// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type HostDatastoreSystem struct {
	mo.HostDatastoreSystem

	Host *mo.HostSystem
}

var defaultDatastoreCapability = types.DatastoreCapability{
	DirectoryHierarchySupported:      true,
	RawDiskMappingsSupported:         false,
	PerFileThinProvisioningSupported: true,
	StorageIORMSupported:             false,
	NativeSnapshotSupported:          false,
	SeSparseSupported:                types.NewBool(false),
	TopLevelDirectoryCreateSupported: types.NewBool(true),
}

func (dss *HostDatastoreSystem) add(ctx *Context, ds *Datastore) *soap.Fault {
	info := ds.Info.GetDatastoreInfo()

	info.Name = ds.Name

	if e := ctx.Map.FindByName(ds.Name, dss.Datastore); e != nil {
		return Fault(e.Reference().Value, &types.DuplicateName{
			Name:   ds.Name,
			Object: e.Reference(),
		})
	}

	fi, err := os.Stat(info.Url)
	if err == nil && !fi.IsDir() {
		err = os.ErrInvalid
	}

	if err != nil {
		switch {
		case os.IsNotExist(err):
			return Fault(err.Error(), &types.NotFound{})
		default:
			return Fault(err.Error(), &types.HostConfigFault{})
		}
	}

	folder := ctx.Map.getEntityFolder(dss.Host, "datastore")

	found := false
	if e := ctx.Map.FindByName(ds.Name, folder.ChildEntity); e != nil {
		if e.Reference().Type != "Datastore" {
			return Fault(e.Reference().Value, &types.DuplicateName{
				Name:   ds.Name,
				Object: e.Reference(),
			})
		}

		// if datastore already exists, use current reference
		found = true
		ds = e.(*Datastore)
	} else {
		ds.Summary.Datastore = &ds.Self
		ds.Summary.Name = ds.Name
		ds.Summary.Url = info.Url

		// put datastore to folder and generate reference
		folderPutChild(ctx, folder, ds)
	}

	ds.Host = append(ds.Host, types.DatastoreHostMount{
		Key: dss.Host.Reference(),
		MountInfo: types.HostMountInfo{
			AccessMode: string(types.HostMountModeReadWrite),
			Mounted:    types.NewBool(true),
			Accessible: types.NewBool(true),
		},
	})

	_ = ds.RefreshDatastore(ctx, &types.RefreshDatastore{This: ds.Self})

	dss.Datastore = append(dss.Datastore, ds.Self)
	dss.Host.Datastore = dss.Datastore
	parent := hostParent(ctx, dss.Host)
	ctx.Map.AddReference(ctx, parent, &parent.Datastore, ds.Self)

	// NOTE: browser must be created after ds is appended to dss.Datastore
	if !found {
		browser := &HostDatastoreBrowser{}
		browser.Datastore = dss.Datastore
		ds.Browser = ctx.Map.Put(browser).Reference()
	}

	return nil
}

func (dss *HostDatastoreSystem) CreateLocalDatastore(ctx *Context, c *types.CreateLocalDatastore) soap.HasFault {
	r := &methods.CreateLocalDatastoreBody{}

	ds := &Datastore{}
	ds.Name = c.Name

	ds.Info = &types.LocalDatastoreInfo{
		DatastoreInfo: types.DatastoreInfo{
			Name: c.Name,
			Url:  c.Path,
		},
		Path: c.Path,
	}

	ds.Summary.Type = string(types.HostFileSystemVolumeFileSystemTypeOTHER)
	ds.Summary.MaintenanceMode = string(types.DatastoreSummaryMaintenanceModeStateNormal)
	ds.Summary.Accessible = true
	ds.Capability = defaultDatastoreCapability

	if err := dss.add(ctx, ds); err != nil {
		r.Fault_ = err
		return r
	}

	r.Res = &types.CreateLocalDatastoreResponse{
		Returnval: ds.Self,
	}

	return r
}

func (dss *HostDatastoreSystem) CreateNasDatastore(ctx *Context, c *types.CreateNasDatastore) soap.HasFault {
	r := &methods.CreateNasDatastoreBody{}

	// validate RemoteHost and RemotePath are specified
	if c.Spec.RemoteHost == "" {
		r.Fault_ = Fault(
			"A specified parameter was not correct: Spec.RemoteHost",
			&types.InvalidArgument{InvalidProperty: "RemoteHost"},
		)
		return r
	}
	if c.Spec.RemotePath == "" {
		r.Fault_ = Fault(
			"A specified parameter was not correct: Spec.RemotePath",
			&types.InvalidArgument{InvalidProperty: "RemotePath"},
		)
		return r
	}

	ds := &Datastore{}
	ds.Name = path.Base(c.Spec.LocalPath)

	ds.Info = &types.NasDatastoreInfo{
		DatastoreInfo: types.DatastoreInfo{
			Url: c.Spec.LocalPath,
		},
		Nas: &types.HostNasVolume{
			HostFileSystemVolume: types.HostFileSystemVolume{
				Name: c.Spec.LocalPath,
				Type: c.Spec.Type,
			},
			RemoteHost: c.Spec.RemoteHost,
			RemotePath: c.Spec.RemotePath,
		},
	}

	ds.Summary.Type = c.Spec.Type
	ds.Summary.MaintenanceMode = string(types.DatastoreSummaryMaintenanceModeStateNormal)
	ds.Summary.Accessible = true
	ds.Capability = defaultDatastoreCapability

	if err := dss.add(ctx, ds); err != nil {
		r.Fault_ = err
		return r
	}

	r.Res = &types.CreateNasDatastoreResponse{
		Returnval: ds.Self,
	}

	return r
}

func (dss *HostDatastoreSystem) createVsanDatastore(ctx *Context) types.BaseMethodFault {
	ds := &Datastore{}
	ds.Name = "vsanDatastore"

	home := ctx.Map.OptionManager().find("vcsim.home").Value
	dc := ctx.Map.getEntityDatacenter(dss.Host)
	url := fmt.Sprintf("%s/%s-%s", home, dc.Name, ds.Name)
	if err := os.MkdirAll(url, 0700); err != nil {
		return ctx.Map.FileManager().fault(url, err, new(types.CannotAccessFile))
	}

	ds.Info = &types.VsanDatastoreInfo{
		DatastoreInfo: types.DatastoreInfo{
			Name: ds.Name,
			Url:  url,
		},
		MembershipUuid: uuid.NewString(),
	}

	ds.Summary.Type = string(types.HostFileSystemVolumeFileSystemTypeVsan)
	ds.Summary.MaintenanceMode = string(types.DatastoreSummaryMaintenanceModeStateNormal)
	ds.Summary.Accessible = true
	ds.Capability = defaultDatastoreCapability
	ds.Capability.TopLevelDirectoryCreateSupported = types.NewBool(false)

	if err := dss.add(ctx, ds); err != nil {
		return err.Detail.Fault.(types.BaseMethodFault)
	}

	return nil
}
