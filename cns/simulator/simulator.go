/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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
	"context"
	"reflect"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/methods"
	cnstypes "github.com/vmware/govmomi/cns/types"
	pbmtypes "github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/simulator"
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
	})

	return r
}

type CnsVolumeManager struct {
	vim25types.ManagedObjectReference
	volumes     map[vim25types.ManagedObjectReference]map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume
	attachments map[cnstypes.CnsVolumeId]vim25types.ManagedObjectReference
}

const simulatorDiskUUID = "6000c298595bf4575739e9105b2c0c2d"

func (m *CnsVolumeManager) CnsCreateVolume(ctx context.Context, req *cnstypes.CnsCreateVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsCreateVolume", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.CreateSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsVolumeCreateSpec"}
		}

		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, createSpec := range req.CreateSpecs {
			staticProvisionedSpec, ok := interface{}(createSpec.BackingObjectDetails).(*cnstypes.CnsBlockBackingDetails)
			if ok && staticProvisionedSpec.BackingDiskId != "" {
				datastore := simulator.Map.Any("Datastore").(*simulator.Datastore)
				volumes, ok := m.volumes[datastore.Self]
				if !ok {
					volumes = make(map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume)
					m.volumes[datastore.Self] = volumes
				}
				newVolume := &cnstypes.CnsVolume{
					VolumeId: cnstypes.CnsVolumeId{
						Id: interface{}(createSpec.BackingObjectDetails).(*cnstypes.CnsBlockBackingDetails).BackingDiskId,
					},
					Name:                         createSpec.Name,
					VolumeType:                   createSpec.VolumeType,
					DatastoreUrl:                 datastore.Info.GetDatastoreInfo().Url,
					Metadata:                     createSpec.Metadata,
					BackingObjectDetails:         createSpec.BackingObjectDetails.(cnstypes.BaseCnsBackingObjectDetails).GetCnsBackingObjectDetails(),
					ComplianceStatus:             "Simulator Compliance Status",
					DatastoreAccessibilityStatus: "Simulator Datastore Accessibility Status",
					HealthStatus:                 string(pbmtypes.PbmHealthStatusForEntityGreen),
				}

				volumes[newVolume.VolumeId] = newVolume
				placementResults := []cnstypes.CnsPlacementResult{}
				placementResults = append(placementResults, cnstypes.CnsPlacementResult{
					Datastore: datastore.Reference(),
				})
				operationResult = append(operationResult, &cnstypes.CnsVolumeCreateResult{
					CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
						VolumeId: newVolume.VolumeId,
					},
					Name:             createSpec.Name,
					PlacementResults: placementResults,
				})

			} else {
				for _, datastoreRef := range createSpec.Datastores {
					datastore := simulator.Map.Get(datastoreRef).(*simulator.Datastore)

					volumes, ok := m.volumes[datastore.Self]
					if !ok {
						volumes = make(map[cnstypes.CnsVolumeId]*cnstypes.CnsVolume)
						m.volumes[datastore.Self] = volumes

					}

					var policyId string
					if createSpec.Profile != nil && createSpec.Profile[0] != nil &&
						reflect.TypeOf(createSpec.Profile[0]) == reflect.TypeOf(&vim25types.VirtualMachineDefinedProfileSpec{}) {
						policyId = interface{}(createSpec.Profile[0]).(*vim25types.VirtualMachineDefinedProfileSpec).ProfileId
					}

					newVolume := &cnstypes.CnsVolume{
						VolumeId: cnstypes.CnsVolumeId{
							Id: uuid.New().String(),
						},
						Name:                         createSpec.Name,
						VolumeType:                   createSpec.VolumeType,
						DatastoreUrl:                 datastore.Info.GetDatastoreInfo().Url,
						Metadata:                     createSpec.Metadata,
						BackingObjectDetails:         createSpec.BackingObjectDetails.(cnstypes.BaseCnsBackingObjectDetails).GetCnsBackingObjectDetails(),
						ComplianceStatus:             "Simulator Compliance Status",
						DatastoreAccessibilityStatus: "Simulator Datastore Accessibility Status",
						HealthStatus:                 string(pbmtypes.PbmHealthStatusForEntityGreen),
						StoragePolicyId:              policyId,
					}

					volumes[newVolume.VolumeId] = newVolume
					placementResults := []cnstypes.CnsPlacementResult{}
					placementResults = append(placementResults, cnstypes.CnsPlacementResult{
						Datastore: datastore.Reference(),
					})
					operationResult = append(operationResult, &cnstypes.CnsVolumeCreateResult{
						CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
							VolumeId: newVolume.VolumeId,
						},
						Name:             createSpec.Name,
						PlacementResults: placementResults,
					})
				}
			}
		}

		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsCreateVolumeBody{
		Res: &cnstypes.CnsCreateVolumeResponse{
			Returnval: task.Run(),
		},
	}
}

// CnsQueryVolume simulates the query volumes implementation for CNSQuery API
func (m *CnsVolumeManager) CnsQueryVolume(ctx context.Context, req *cnstypes.CnsQueryVolume) soap.HasFault {
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

	return &methods.CnsQueryVolumeBody{
		Res: &cnstypes.CnsQueryVolumeResponse{
			Returnval: cnstypes.CnsQueryResult{
				Volumes: retVolumes,
				Cursor:  cnstypes.CnsCursor{},
			},
		},
	}
}

// CnsQueryAllVolume simulates the query volumes implementation for CNSQueryAll API
func (m *CnsVolumeManager) CnsQueryAllVolume(ctx context.Context, req *cnstypes.CnsQueryAllVolume) soap.HasFault {
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

	return &methods.CnsQueryAllVolumeBody{
		Res: &cnstypes.CnsQueryAllVolumeResponse{
			Returnval: cnstypes.CnsQueryResult{
				Volumes: retVolumes,
				Cursor:  cnstypes.CnsCursor{},
			},
		},
	}
}

func (m *CnsVolumeManager) CnsDeleteVolume(ctx context.Context, req *cnstypes.CnsDeleteVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsDeleteVolume", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, volumeId := range req.VolumeIds {
			for ds, dsVolumes := range m.volumes {
				volume := dsVolumes[volumeId]
				if volume != nil {
					delete(m.volumes[ds], volumeId)
					operationResult = append(operationResult, &cnstypes.CnsVolumeOperationResult{
						VolumeId: volumeId,
					})

				}
			}
		}
		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsDeleteVolumeBody{
		Res: &cnstypes.CnsDeleteVolumeResponse{
			Returnval: task.Run(),
		},
	}
}

// CnsUpdateVolumeMetadata simulates UpdateVolumeMetadata call for simulated vc
func (m *CnsVolumeManager) CnsUpdateVolumeMetadata(ctx context.Context, req *cnstypes.CnsUpdateVolumeMetadata) soap.HasFault {
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
			Returnval: task.Run(),
		},
	}
}

// CnsAttachVolume simulates AttachVolume call for simulated vc
func (m *CnsVolumeManager) CnsAttachVolume(ctx context.Context, req *cnstypes.CnsAttachVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsAttachVolume", func(task *simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.AttachSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsAttachVolumeSpec"}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, attachSpec := range req.AttachSpecs {
			node := simulator.Map.Get(attachSpec.Vm).(*simulator.VirtualMachine)
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
			Returnval: task.Run(),
		},
	}
}

// CnsDetachVolume simulates DetachVolume call for simulated vc
func (m *CnsVolumeManager) CnsDetachVolume(ctx context.Context, req *cnstypes.CnsDetachVolume) soap.HasFault {
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
			Returnval: task.Run(),
		},
	}
}

// CnsExtendVolume simulates ExtendVolume call for simulated vc
func (m *CnsVolumeManager) CnsExtendVolume(ctx context.Context, req *cnstypes.CnsExtendVolume) soap.HasFault {
	task := simulator.CreateTask(m, "CnsExtendVolume", func(task *simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		if len(req.ExtendSpecs) == 0 {
			return nil, &vim25types.InvalidArgument{InvalidProperty: "CnsExtendVolumeSpec"}
		}
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}

		for _, extendSpecs := range req.ExtendSpecs {
			for _, dsVolumes := range m.volumes {
				for id, volume := range dsVolumes {
					if id.Id == extendSpecs.VolumeId.Id {
						volume.BackingObjectDetails = &cnstypes.CnsBackingObjectDetails{
							CapacityInMb: extendSpecs.CapacityInMb,
						}
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

	return &methods.CnsExtendVolumeBody{
		Res: &cnstypes.CnsExtendVolumeResponse{
			Returnval: task.Run(),
		},
	}
}

func (m *CnsVolumeManager) CnsQueryVolumeInfo(ctx context.Context, req *cnstypes.CnsQueryVolumeInfo) soap.HasFault {
	task := simulator.CreateTask(m, "CnsQueryVolumeInfo", func(*simulator.Task) (vim25types.AnyType, vim25types.BaseMethodFault) {
		operationResult := []cnstypes.BaseCnsVolumeOperationResult{}
		for _, volumeId := range req.VolumeIds {
			vstorageObject := vim25types.VStorageObject{
				Config: vim25types.VStorageObjectConfigInfo{
					BaseConfigInfo: vim25types.BaseConfigInfo{
						Id: vim25types.ID{
							Id: uuid.New().String(),
						},
						Name:                        "name",
						CreateTime:                  time.Now(),
						KeepAfterDeleteVm:           vim25types.NewBool(true),
						RelocationDisabled:          vim25types.NewBool(false),
						NativeSnapshotSupported:     vim25types.NewBool(false),
						ChangedBlockTrackingEnabled: vim25types.NewBool(false),
						Iofilter:                    nil,
					},
					CapacityInMB:    1024,
					ConsumptionType: []string{"disk"},
					ConsumerId:      nil,
				},
			}
			vstorageObject.Config.Backing = &vim25types.BaseConfigInfoDiskFileBackingInfo{
				BaseConfigInfoFileBackingInfo: vim25types.BaseConfigInfoFileBackingInfo{
					BaseConfigInfoBackingInfo: vim25types.BaseConfigInfoBackingInfo{
						Datastore: simulator.Map.Any("Datastore").(*simulator.Datastore).Self,
					},
					FilePath:        "[vsanDatastore] 6785a85e-268e-6352-a2e8-02008b7afadd/kubernetes-dynamic-pvc-68734c9f-a679-42e6-a694-39632c51e31f.vmdk",
					BackingObjectId: volumeId.Id,
					Parent:          nil,
					DeltaSizeInMB:   0,
				},
			}

			operationResult = append(operationResult, &cnstypes.CnsQueryVolumeInfoResult{
				CnsVolumeOperationResult: cnstypes.CnsVolumeOperationResult{
					VolumeId: volumeId,
				},
				VolumeInfo: &cnstypes.CnsBlockVolumeInfo{
					CnsVolumeInfo:  cnstypes.CnsVolumeInfo{},
					VStorageObject: vstorageObject,
				},
			})

		}
		return &cnstypes.CnsVolumeOperationBatchResult{
			VolumeResults: operationResult,
		}, nil
	})

	return &methods.CnsQueryVolumeInfoBody{
		Res: &cnstypes.CnsQueryVolumeInfoResponse{
			Returnval: task.Run(),
		},
	}
}
