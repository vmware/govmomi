// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/methods"
	"github.com/vmware/govmomi/cns/types"
	cnstypes "github.com/vmware/govmomi/cns/types"
	pbmtypes "github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	vim25methods "github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	vim25types "github.com/vmware/govmomi/vim25/types"
)

func init() {
	simulator.RegisterEndpoint(func(s *simulator.Service, r *simulator.Registry) {
		if r.IsVPX() {
			s.RegisterSDK(New())
		}
	})
}

func New() *simulator.Registry {
	r := simulator.NewRegistry()
	r.Namespace = cns.Namespace
	r.Path = cns.Path

	r.Put(&CnsVolumeManager{
		ManagedObjectReference: cns.CnsVolumeManagerInstance,
		volumes:                make(map[vim25types.ManagedObjectReference]map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume),
		attachments:            make(map[cnstypes.CnsVolumeId]vim25types.ManagedObjectReference),
		snapshots:              make(map[cnstypes.CnsVolumeId]map[cnstypes.CnsSnapshotId]*cnstypes.CnsSnapshot),
	})

	return r
}

type CnsVolumeManager struct {
	vim25types.ManagedObjectReference
	volumes     map[vim25types.ManagedObjectReference]map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume
	attachments map[cnstypes.CnsVolumeId]vim25types.ManagedObjectReference
	snapshots   map[cnstypes.CnsVolumeId]map[cnstypes.CnsSnapshotId]*cnstypes.CnsSnapshot
}

const simulatorDiskUUID = "6000c298595bf4575739e9105b2c0c2d"

func (m *CnsVolumeManager) findDisk(vctx *simulator.Context, createSpec cnstypes.CnsVolumeCreateSpec) (*vim25types.VStorageObject, vim25types.BaseMethodFault) {
	vsom := vctx.Map.VStorageObjectManager()
	details := createSpec.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails)

	for _, objs := range vsom.Catalog() {
		for id, val := range objs {
			if id.Id == details.BackingDiskId {
				return &val.VStorageObject, nil
			}
		}
	}

	return nil, &vim25types.InvalidArgument{InvalidProperty: "VolumeId"}
}

func (m *CnsVolumeManager) createDisk(vctx *simulator.Context, createSpec cnstypes.CnsVolumeCreateSpec) (*vim25types.VStorageObject, vim25types.BaseMethodFault) {
	vsom := vctx.Map.VStorageObjectManager()

	if len(createSpec.Datastores) == 0 {
		return nil, &vim25types.InvalidArgument{InvalidProperty: "createSpecs.datastores"}
	}

	// "Datastores to be considered for volume placement" - we'll just use the 1st for now
	datastoreRef := createSpec.Datastores[0]

	val := vsom.CreateDiskTask(vctx, &vim25types.CreateDisk_Task{
		This: vsom.Self,
		Spec: vim25types.VslmCreateSpec{
			Name:              createSpec.Name,
			KeepAfterDeleteVm: vim25types.NewBool(true),
			CapacityInMB:      createSpec.BackingObjectDetails.GetCnsBackingObjectDetails().CapacityInMb,
			Profile:           createSpec.Profile,
			BackingSpec: &vim25types.VslmCreateSpecDiskFileBackingSpec{
				VslmCreateSpecBackingSpec: vim25types.VslmCreateSpecBackingSpec{
					Datastore: datastoreRef,
				},
				ProvisioningType: string(vim25types.BaseConfigInfoDiskFileBackingInfoProvisioningTypeThin),
			},
		},
	})

	ref := val.(*vim25methods.CreateDisk_TaskBody).Res.Returnval
	task := vctx.Map.Get(ref).(*simulator.Task)
	task.Wait()
	if task.Info.Error != nil {
		return nil, task.Info.Error.Fault
	}

	return task.Info.Result.(*vim25types.VStorageObject), nil
}

func (m *CnsVolumeManager) CnsCreateVolume(ctx *simulator.Context, req *cnstypes.CnsCreateVolume) soap.HasFault {
	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	task := simulator.CreateTask(m, "CnsCreateVolume", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.CreateSpecs) != 1 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "InputSpec"} // Same as real VC, currently
		}

		var obj *vim25types.VStorageObject
		var fault vim25types.BaseMethodFault

		createSpec := req.CreateSpecs[0]

		switch details := createSpec.BackingObjectDetails.(type) {
		case *cnstypes.CnsBlockBackingDetails:
			if details.BackingDiskId == "" {
				obj, fault = m.createDisk(vctx, createSpec)
			} else {
				obj, fault = m.findDisk(vctx, createSpec)
			}
			if fault != nil {
				return nil, fault
			}
		default:
			return nil, &vim25types.InvalidArgument{InvalidProperty: "createSpecs.backingObjectDetails"}
		}

		datastoreRef := obj.Config.Backing.GetBaseConfigInfoBackingInfo().Datastore
		datastore := vctx.Map.Get(datastoreRef).(*simulator.Datastore)

		volumes, ok := m.volumes[datastore.Self]
		if !ok {
			volumes = make(map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume)
			m.volumes[datastore.Self] = volumes
		}

		policyId := ""
		if len(createSpec.Profile) != 0 {
			if profileSpec, ok := createSpec.Profile[0].(*vim25types.VirtualMachineDefinedProfileSpec); ok {
				policyId = profileSpec.ProfileId
			}
		}

		volume := &cnstypes.CnsVolume{
			VolumeId:     cnstypes.CnsVolumeId(obj.Config.Id),
			Name:         createSpec.Name,
			VolumeType:   createSpec.VolumeType,
			DatastoreUrl: datastore.Info.GetDatastoreInfo().Url,
			Metadata:     createSpec.Metadata,
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
					CapacityInMb: obj.Config.CapacityInMB,
				},
				BackingDiskId:   obj.Config.Id.Id,
				BackingDiskPath: obj.Config.Backing.(*vim25types.BaseConfigInfoDiskFileBackingInfo).FilePath,
			},
			ComplianceStatus:             string(pbmtypes.PbmComplianceStatusCompliant),
			DatastoreAccessibilityStatus: string(pbmtypes.PbmHealthStatusForEntityGreen),
			HealthStatus:                 string(pbmtypes.PbmHealthStatusForEntityGreen),
			StoragePolicyId:              policyId,
		}

		volumes[volume.VolumeId] = volume

		operationResult := []cnstypes.BaseCnsVolumeOperationResult{
			&cnstypes.CnsVolumeCreateResult{
				CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
					VolumeId: volume.VolumeId,
				},
				Name: volume.Name,
				PlacementResults: []cnstypes.CnsPlacementResult{{
					Datastore: datastore.Reference(),
				}},
			},
		}

		cc := createSpec.Metadata.ContainerCluster
		vsom.VCenterUpdateVStorageObjectMetadataExTask(vctx, &vim25types.VCenterUpdateVStorageObjectMetadataEx_Task{
			This:      vsom.Self,
			Id:        obj.Config.Id,
			Datastore: datastore.Self,
			Metadata: []vim25types.KeyValue{
				{Key: "cns.containerCluster.clusterDistribution", Value: cc.ClusterDistribution},
				{Key: "cns.containerCluster.clusterFlavor", Value: cc.ClusterFlavor},
				{Key: "cns.containerCluster.clusterId", Value: cc.ClusterId},
				{Key: "cns.containerCluster.vSphereUser", Value: cc.VSphereUser},
			},
		})

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsCreateVolumeBody{
		Res: &cnstypes.CnsCreateVolumeResponse{
			Returnval: task.Run(vctx),
		},
	}
}

func matchesDatastore(filter cnstypes.CnsQueryFilter, ds vim25types.ManagedObjectReference) bool {
	if len(filter.Datastores) == 0 {
		return true
	}
	return slices.Contains(filter.Datastores, ds)
}

func matchesID(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	if len(filter.VolumeIds) == 0 {
		return true
	}
	return slices.Contains(filter.VolumeIds, volume.VolumeId)
}

func matchesName(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	if len(filter.Names) == 0 {
		return true
	}
	return slices.Contains(filter.Names, volume.Name)
}

func matchesContainerCluster(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	if len(filter.ContainerClusterIds) == 0 {
		return true
	}
	return slices.Contains(filter.ContainerClusterIds, volume.Metadata.ContainerCluster.ClusterId)
}

func matchesStoragePolicy(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	return filter.StoragePolicyId == "" || filter.StoragePolicyId == volume.StoragePolicyId
}

func matchesLabel(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	if len(filter.Labels) == 0 {
		return true
	}

	for _, meta := range volume.Metadata.EntityMetadata {
		for _, label := range meta.GetCnsEntityMetadata().Labels {
			if slices.Contains(filter.Labels, label) {
				return true
			}
		}
	}

	return false
}

func matchesComplianceStatus(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	return filter.ComplianceStatus == "" || filter.ComplianceStatus == volume.ComplianceStatus
}

func matchesHealth(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	return filter.HealthStatus == "" || filter.HealthStatus == volume.HealthStatus
}

func matchesFilter(filter cnstypes.CnsQueryFilter, volume *types.CnsVolume) bool {
	matches := []func(cnstypes.CnsQueryFilter, *types.CnsVolume) bool{
		matchesID,
		matchesName,
		matchesContainerCluster,
		matchesStoragePolicy,
		matchesLabel,
		matchesComplianceStatus,
		matchesHealth,
	}

	for _, match := range matches {
		if !match(filter, volume) {
			return false
		}
	}

	return true
}

func (m *CnsVolumeManager) queryVolume(filter cnstypes.BaseCnsQueryFilter) []cnstypes.CnsVolume {
	var matches []cnstypes.CnsVolume

	for ds, volumes := range m.volumes {
		if !matchesDatastore(*filter.GetCnsQueryFilter(), ds) {
			continue
		}
		for _, volume := range volumes {
			if matchesFilter(*filter.GetCnsQueryFilter(), volume) {
				matches = append(matches, *volume)
			}
		}
	}

	return matches
}

// CnsQueryVolume simulates the query volumes implementation for CNSQuery API
func (m *CnsVolumeManager) CnsQueryVolume(ctx context.Context, req *cnstypes.CnsQueryVolume) soap.HasFault {
	// TODO: paginate results using CnsCursor
	return &methods.CnsQueryVolumeBody{
		Res: &cnstypes.CnsQueryVolumeResponse{
			Returnval: cnstypes.CnsQueryResult{
				Volumes: m.queryVolume(req.Filter),
			},
		},
	}
}

// CnsQueryAllVolume simulates the query volumes implementation for CNSQueryAll API
func (m *CnsVolumeManager) CnsQueryAllVolume(ctx context.Context, req *cnstypes.CnsQueryAllVolume) soap.HasFault {
	// Note we ignore req.Selection, which can limit fields to reduce response size
	// This method is internal and does not paginate results using CnsCursor.
	return &methods.CnsQueryAllVolumeBody{
		Res: &cnstypes.CnsQueryAllVolumeResponse{
			Returnval: cnstypes.CnsQueryResult{
				Volumes: m.queryVolume(req.Filter.GetCnsQueryFilter()),
			},
		},
	}
}

func (m *CnsVolumeManager) CnsDeleteVolume(ctx *simulator.Context, req *cnstypes.CnsDeleteVolume) soap.HasFault {
	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	task := simulator.CreateTask(m, "CnsDeleteVolume", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.VolumeIds) != 1 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "InputSpec"} // Same as real VC, currently
		}

		found := false
		volumeId := req.VolumeIds[0]
		res := &cnstypes.CnsVolumeOperationResult{
			VolumeId: volumeId,
		}

		for ds, volumes := range m.volumes {
			if _, ok := volumes[volumeId]; ok {
				found = true
				delete(m.volumes[ds], volumeId)

				if req.DeleteDisk {
					val := vsom.DeleteVStorageObjectTask(vctx, &vim25types.DeleteVStorageObject_Task{
						This:      vsom.Self,
						Id:        vim25types.ID(volumeId),
						Datastore: ds,
					})

					ref := val.(*vim25methods.DeleteVStorageObject_TaskBody).Res.Returnval
					task := vctx.Map.Get(ref).(*simulator.Task)
					task.Wait()
					if task.Info.Error != nil {
						res.Fault = task.Info.Error
					}
				}
			}
		}

		if !found {
			res.Fault = &vim25types.LocalizedMethodFault{Fault: new(vim25types.NotFound)}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: []cnstypes.BaseCnsVolumeOperationResult{res},
		}, nil
	})

	return &methods.CnsDeleteVolumeBody{
		Res: &cnstypes.CnsDeleteVolumeResponse{
			Returnval: task.Run(vctx),
		},
	}
}

// CnsUpdateVolumeMetadata simulates UpdateVolumeMetadata call for simulated vc
func (m *CnsVolumeManager) CnsUpdateVolumeMetadata(ctx *simulator.Context, req *cnstypes.CnsUpdateVolumeMetadata) soap.HasFault {
	task := simulator.CreateTask(m, "CnsUpdateVolumeMetadata", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.UpdateSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsUpdateVolumeMetadataSpec"}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, updateSpecs := range req.UpdateSpecs {
			for _, dsVolumes := range m.volumes {
				for id, volume := range dsVolumes {
					if id.Id == updateSpecs.VolumeId.Id {
						volume.Metadata.EntityMetadata = updateSpecs.Metadata.EntityMetadata
						operationResult = append(operationResult, &cnstypes.CnsVolumeOperationResult{
							VolumeId: volume.VolumeId,
						})
						break
					}
				}
			}

		}
		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})
	return &methods.CnsUpdateVolumeBody{
		Res: &cnstypes.CnsUpdateVolumeMetadataResponse{
			Returnval: task.Run(ctx),
		},
	}
}

// CnsAttachVolume simulates AttachVolume call for simulated vc
func (m *CnsVolumeManager) CnsAttachVolume(ctx *simulator.Context, req *cnstypes.CnsAttachVolume) soap.HasFault {
	vctx := ctx.For(vim25.Path)
	task := simulator.CreateTask(m, "CnsAttachVolume", func(task *simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.AttachSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsAttachVolumeSpec"}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, attachSpec := range req.AttachSpecs {
			node := vctx.Map.Get(attachSpec.Vm).(*simulator.VirtualMachine)
			if _, ok := m.attachments[attachSpec.VolumeId]; !ok {
				m.attachments[attachSpec.VolumeId] = node.Self
			} else {
				return nil, &vim25types.ResourceInUse{
					Name: attachSpec.VolumeId.Id,
				}
			}
			operationResult = append(operationResult, &cnstypes.CnsVolumeAttachResult{
				CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
					VolumeId: attachSpec.VolumeId,
				},
				DiskUUID: simulatorDiskUUID,
			})
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsAttachVolumeBody{
		Res: &cnstypes.CnsAttachVolumeResponse{
			Returnval: task.Run(vctx),
		},
	}
}

// CnsDetachVolume simulates DetachVolume call for simulated vc
func (m *CnsVolumeManager) CnsDetachVolume(ctx *simulator.Context, req *cnstypes.CnsDetachVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsDetachVolume", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.DetachSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsDetachVolumeSpec"}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, detachSpec := range req.DetachSpecs {
			if _, ok := m.attachments[detachSpec.VolumeId]; ok {
				delete(m.attachments, detachSpec.VolumeId)
				operationResult = append(operationResult, &cnstypes.CnsVolumeOperationResult{
					VolumeId: detachSpec.VolumeId,
				})
			} else {
				return nil, &vim25types.InvalidArgument{
					InvalidProperty: detachSpec.VolumeId.Id,
				}
			}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})
	return &methods.CnsDetachVolumeBody{
		Res: &cnstypes.CnsDetachVolumeResponse{
			Returnval: task.Run(ctx),
		},
	}
}

// CnsExtendVolume simulates ExtendVolume call for simulated vc
func (m *CnsVolumeManager) CnsExtendVolume(ctx *simulator.Context, req *cnstypes.CnsExtendVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsExtendVolume", func(task *simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.ExtendSpecs) != 1 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "InputSpec"} // Same as real VC, currently
		}

		found := false
		spec := req.ExtendSpecs[0]
		res := cnstypes.CnsVolumeOperationResult{
			VolumeId: spec.VolumeId,
		}

		for _, volumes := range m.volumes {
			if volume, ok := volumes[spec.VolumeId]; ok {
				found = true
				volume.BackingObjectDetails.GetCnsBackingObjectDetails().CapacityInMb = spec.CapacityInMb
				break
			}
		}

		if !found {
			res.Fault = &vim25types.LocalizedMethodFault{Fault: new(vim25types.NotFound)}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: []cnstypes.BaseCnsVolumeOperationResult{&res},
		}, nil
	})

	return &methods.CnsExtendVolumeBody{
		Res: &cnstypes.CnsExtendVolumeResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *CnsVolumeManager) CnsQueryVolumeInfo(ctx *simulator.Context, req *cnstypes.CnsQueryVolumeInfo) soap.HasFault {
	vctx := ctx.For(vim25.Path)
	vsom := vctx.Map.VStorageObjectManager()

	task := simulator.CreateTask(m, "CnsQueryVolumeInfo", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		var operationResult []cnstypes.BaseCnsVolumeOperationResult

		for _, objs := range vsom.Catalog() {
			for _, obj := range objs {
				for _, volumeId := range req.VolumeIds {
					if obj.Config.Id.Id == volumeId.Id {
						operationResult = append(operationResult, &cnstypes.CnsQueryVolumeInfoResult{
							CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
								VolumeId: volumeId,
							},
							VolumeInfo: &cnstypes.CnsBlockVolumeInfo{
								VStorageObject: obj.VStorageObject,
							},
						})
					}
				}
			}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsQueryVolumeInfoBody{
		Res: &cnstypes.CnsQueryVolumeInfoResponse{
			Returnval: task.Run(vctx),
		},
	}
}

func (m *CnsVolumeManager) CnsQueryAsync(ctx *simulator.Context, req *cnstypes.CnsQueryAsync) soap.HasFault {
	task := simulator.CreateTask(m, "QueryVolumeAsync", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		retVolumes := []cnstypes.CnsVolume{}
		reqVolumeIds := make(map[string]bool)
		isQueryFilter := false

		if req.Filter.VolumeIds != nil {
			isQueryFilter = true
		}
		// Create map of requested volume Ids in query request
		for _, volumeID := range req.Filter.VolumeIds {
			reqVolumeIds[volumeID.Id] = true
		}

		for _, dsVolumes := range m.volumes {
			for _, volume := range dsVolumes {
				if isQueryFilter {
					if _, ok := reqVolumeIds[volume.VolumeId.Id]; ok {
						retVolumes = append(retVolumes, *volume)
					}
				} else {
					retVolumes = append(retVolumes, *volume)
				}
			}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		operationResult = append(operationResult, &cnstypes.CnsAsyncQueryResult{
			QueryResult: cnstypes.CnsQueryResult{
				Volumes: retVolumes,
				Cursor:  cnstypes.CnsCursor{},
			},
		})

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsQueryAsyncBody{
		Res: &cnstypes.CnsQueryAsyncResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *CnsVolumeManager) CnsCreateSnapshots(ctx *simulator.Context, req *cnstypes.CnsCreateSnapshots) soap.HasFault {
	task := simulator.CreateTask(m, "CreateSnapshots", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.SnapshotSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsSnapshotCreateSpec"}
		}

		snapshotOperationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, snapshotCreateSpec := range req.SnapshotSpecs {
			for _, dsVolumes := range m.volumes {
				for id := range dsVolumes {
					if id.Id != snapshotCreateSpec.VolumeId.Id {
						continue
					}
					snapshots, ok := m.snapshots[snapshotCreateSpec.VolumeId]
					if !ok {
						snapshots = make(map[cnstypes.CnsSnapshotId]*cnstypes.CnsSnapshot)
						m.snapshots[snapshotCreateSpec.VolumeId] = snapshots
					}

					newSnapshot := &cnstypes.CnsSnapshot{
						SnapshotId: cnstypes.CnsSnapshotId{
							Id: uuid.New().String(),
						},
						VolumeId:    snapshotCreateSpec.VolumeId,
						Description: snapshotCreateSpec.Description,
						CreateTime:  time.Now(),
					}
					snapshots[newSnapshot.SnapshotId] = newSnapshot
					snapshotOperationResult = append(snapshotOperationResult, &cnstypes.CnsSnapshotCreateResult{
						CnsSnapshotOperationResult: cnstypes.CnsSnapshotOperationResult{
							CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
								VolumeId: newSnapshot.VolumeId,
							},
						},
						Snapshot: *newSnapshot,
					})
				}
			}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: snapshotOperationResult,
		}, nil
	})

	return &methods.CnsCreateSnapshotsBody{
		Res: &cnstypes.CnsCreateSnapshotsResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *CnsVolumeManager) CnsDeleteSnapshots(ctx *simulator.Context, req *cnstypes.CnsDeleteSnapshots) soap.HasFault {
	task := simulator.CreateTask(m, "DeleteSnapshots", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		snapshotOperationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, snapshotDeleteSpec := range req.SnapshotDeleteSpecs {
			for _, dsVolumes := range m.volumes {
				for id := range dsVolumes {
					if id.Id != snapshotDeleteSpec.VolumeId.Id {
						continue
					}
					snapshots := m.snapshots[snapshotDeleteSpec.VolumeId]
					snapshot, ok := snapshots[snapshotDeleteSpec.SnapshotId]
					if ok {
						delete(m.snapshots[snapshotDeleteSpec.VolumeId], snapshotDeleteSpec.SnapshotId)
						snapshotOperationResult = append(snapshotOperationResult, &cnstypes.CnsSnapshotDeleteResult{
							CnsSnapshotOperationResult: cnstypes.CnsSnapshotOperationResult{
								CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
									VolumeId: snapshot.VolumeId,
								},
							},
							SnapshotId: snapshot.SnapshotId,
						})
					}
				}
			}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: snapshotOperationResult,
		}, nil
	})

	return &methods.CnsDeleteSnapshotBody{
		Res: &cnstypes.CnsDeleteSnapshotsResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *CnsVolumeManager) CnsQuerySnapshots(ctx *simulator.Context, req *cnstypes.CnsQuerySnapshots) soap.HasFault {
	task := simulator.CreateTask(m, "QuerySnapshots", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.SnapshotQueryFilter.SnapshotQuerySpecs) > 1 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsSnapshotQuerySpec"}
		}

		snapshotQueryResultEntries := []cnstypes.CnsSnapshotQueryResultEntry{}
		checkVolumeExists := func(volumeId cnstypes.CnsVolumeId) bool {
			for _, dsVolumes := range m.volumes {
				for id := range dsVolumes {
					if id.Id == volumeId.Id {
						return true
					}
				}
			}
			return false
		}

		if req.SnapshotQueryFilter.SnapshotQuerySpecs == nil && len(req.SnapshotQueryFilter.SnapshotQuerySpecs) == 0 {
			// return all snapshots if snapshotQuerySpecs is empty
			for _, volSnapshots := range m.snapshots {
				for _, snapshot := range volSnapshots {
					snapshotQueryResultEntries = append(snapshotQueryResultEntries, cnstypes.CnsSnapshotQueryResultEntry{Snapshot: *snapshot})
				}
			}
		} else {
			// snapshotQuerySpecs is not empty
			isSnapshotQueryFilter := false
			snapshotQuerySpec := req.SnapshotQueryFilter.SnapshotQuerySpecs[0]
			if snapshotQuerySpec.SnapshotId != nil && (*snapshotQuerySpec.SnapshotId != cnstypes.CnsSnapshotId{}) {
				isSnapshotQueryFilter = true
			}

			if !checkVolumeExists(snapshotQuerySpec.VolumeId) {
				// volumeId in snapshotQuerySpecs does not exist
				snapshotQueryResultEntries = append(snapshotQueryResultEntries, cnstypes.CnsSnapshotQueryResultEntry{
					Error: &vim25types.LocalizedMethodFault{
						Fault: &cnstypes.CnsVolumeNotFoundFault{
							VolumeId: snapshotQuerySpec.VolumeId,
						},
					},
				})
			} else {
				// volumeId in snapshotQuerySpecs exists
				for _, snapshot := range m.snapshots[snapshotQuerySpec.VolumeId] {
					if isSnapshotQueryFilter && snapshot.SnapshotId.Id != (*snapshotQuerySpec.SnapshotId).Id {
						continue
					}

					snapshotQueryResultEntries = append(snapshotQueryResultEntries, cnstypes.CnsSnapshotQueryResultEntry{Snapshot: *snapshot})
				}

				if isSnapshotQueryFilter && len(snapshotQueryResultEntries) == 0 {
					snapshotQueryResultEntries = append(snapshotQueryResultEntries, cnstypes.CnsSnapshotQueryResultEntry{
						Error: &vim25types.LocalizedMethodFault{
							Fault: &cnstypes.CnsSnapshotNotFoundFault{
								VolumeId:   snapshotQuerySpec.VolumeId,
								SnapshotId: *snapshotQuerySpec.SnapshotId,
							},
						},
					})
				}
			}
		}

		return &cnstypes.CnsSnapshotQueryResult{
			Entries: snapshotQueryResultEntries,
		}, nil
	})

	return &methods.CnsQuerySnapshotsBody{
		Res: &cnstypes.CnsQuerySnapshotsResponse{
			Returnval: task.Run(ctx),
		},
	}
}
