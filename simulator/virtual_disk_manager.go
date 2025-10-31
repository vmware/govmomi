// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vmdk"
)

type VirtualDiskManager struct {
	mo.VirtualDiskManager
}

func (m *VirtualDiskManager) MO() mo.VirtualDiskManager {
	return m.VirtualDiskManager
}

func VirtualDiskBackingFileName(name string) string {
	return strings.Replace(name, ".vmdk", "-flat.vmdk", 1)
}

func vdmNames(name string) []string {
	return []string{
		VirtualDiskBackingFileName(name),
		name,
	}
}

func vdmCreateVirtualDisk(
	ctx *Context,
	op types.VirtualDeviceConfigSpecFileOperation,
	req *types.CreateVirtualDisk_Task) (*vmdk.Descriptor, types.BaseMethodFault) {

	fm := ctx.Map.FileManager()

	file, fault := fm.resolve(ctx, req.Datacenter, req.Name)
	if fault != nil {
		return nil, fault
	}

	shouldReplace := op == types.VirtualDeviceConfigSpecFileOperationReplace
	shouldExist := op == ""

	_, err := os.Stat(file)
	if err == nil {
		if shouldExist {
			return nil, nil // TODO
		}
		if !shouldReplace {
			return nil, fm.fault(file, nil, new(types.FileAlreadyExists))
		}
	} else if shouldExist {
		return nil, fm.fault(file, nil, new(types.FileNotFound))
	}

	backing := VirtualDiskBackingFileName(file)

	extent := vmdk.Extent{
		Info: filepath.Base(backing),
	}

	f, err := os.Create(file)
	if err != nil {
		return nil, fm.fault(file, err, new(types.CannotCreateFile))
	}

	defer f.Close()

	if req.Spec != nil {
		spec, ok := req.Spec.(types.BaseFileBackedVirtualDiskSpec)
		if !ok {
			return nil, fm.fault(file, nil, new(types.FileFault))
		}

		fileSpec := spec.GetFileBackedVirtualDiskSpec()
		extent.Size = fileSpec.CapacityKb * 1024 / vmdk.SectorSize
	}

	desc := vmdk.NewDescriptor(extent)

	if req != nil && req.Spec != nil {
		if s, ok := req.Spec.(*types.FileBackedVirtualDiskSpec); ok {
			var (
				providerID string
				keyID      string
			)

			switch c := s.Crypto.(type) {
			case *types.CryptoSpecEncrypt:
				keyID = c.CryptoKeyId.KeyId
				if c.CryptoKeyId.ProviderId != nil {
					providerID = c.CryptoKeyId.ProviderId.Id
				}

			case *types.CryptoSpecDeepRecrypt:
				keyID = c.NewKeyId.KeyId
				if c.NewKeyId.ProviderId != nil {
					providerID = c.NewKeyId.ProviderId.Id
				}
			case *types.CryptoSpecShallowRecrypt:
				keyID = c.NewKeyId.KeyId
				if c.NewKeyId.ProviderId != nil {
					providerID = c.NewKeyId.ProviderId.Id
				}
			}

			if providerID != "" || keyID != "" {
				desc.EncryptionKeys = &crypto.KeyLocator{
					Type: crypto.KeyLocatorTypeList,
					List: []*crypto.KeyLocator{
						{
							Type: crypto.KeyLocatorTypePair,
							Pair: &crypto.KeyLocatorPair{
								CryptoMAC:  "HMAC-SHA-256",
								LockedData: []byte("Cg=="),
								Locker: &crypto.KeyLocator{
									Type: crypto.KeyLocatorTypeFQID,
									Indirect: &crypto.KeyLocatorIndirect{
										Type: crypto.KeyLocatorTypeFQID,
										FQID: crypto.KeyLocatorFQIDParams{
											KeyServerID: providerID,
											KeyID:       keyID,
										},
									},
								},
							},
						},
					},
				}
			}
		}
	}
	if err := desc.Write(f); err != nil {
		return nil, fm.fault(file, err, new(types.FileFault))
	}

	b, err := os.Create(backing)
	if err != nil {
		return nil, fm.fault(backing, err, new(types.CannotCreateFile))
	}
	_ = b.Close()

	return desc, nil
}

func vdmExtendVirtualDisk(ctx *Context, req *types.ExtendVirtualDisk_Task) types.BaseMethodFault {
	fm := ctx.Map.FileManager()

	desc, file, fault := fm.DiskDescriptor(ctx, req.Datacenter, req.Name)
	if fault != nil {
		return fault
	}

	newCapacity := req.NewCapacityKb * 1024
	if desc.Capacity() > newCapacity {
		return fm.fault(req.Name, errors.New("cannot shrink disk"), new(types.FileFault))
	}

	desc.Extent[0].Size = newCapacity / vmdk.SectorSize

	return fm.SaveDiskDescriptor(ctx, desc, file)
}

func (m *VirtualDiskManager) CreateVirtualDiskTask(ctx *Context, req *types.CreateVirtualDisk_Task) soap.HasFault {
	task := CreateTask(m, "createVirtualDisk", func(*Task) (types.AnyType, types.BaseMethodFault) {
		if _, err := vdmCreateVirtualDisk(ctx, types.VirtualDeviceConfigSpecFileOperationCreate, req); err != nil {
			return "", err
		}
		return req.Name, nil
	})

	return &methods.CreateVirtualDisk_TaskBody{
		Res: &types.CreateVirtualDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *VirtualDiskManager) ExtendVirtualDiskTask(ctx *Context, req *types.ExtendVirtualDisk_Task) soap.HasFault {
	task := CreateTask(m, "extendVirtualDisk", func(*Task) (types.AnyType, types.BaseMethodFault) {
		if err := vdmExtendVirtualDisk(ctx, req); err != nil {
			return "", err
		}
		return req.Name, nil
	})

	return &methods.ExtendVirtualDisk_TaskBody{
		Res: &types.ExtendVirtualDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *VirtualDiskManager) DeleteVirtualDiskTask(ctx *Context, req *types.DeleteVirtualDisk_Task) soap.HasFault {
	task := CreateTask(m, "deleteVirtualDisk", func(*Task) (types.AnyType, types.BaseMethodFault) {
		fm := ctx.Map.FileManager()

		for _, name := range vdmNames(req.Name) {
			err := fm.deleteDatastoreFile(ctx, &types.DeleteDatastoreFile_Task{
				Name:       name,
				Datacenter: req.Datacenter,
			})

			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return &methods.DeleteVirtualDisk_TaskBody{
		Res: &types.DeleteVirtualDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *VirtualDiskManager) MoveVirtualDiskTask(ctx *Context, req *types.MoveVirtualDisk_Task) soap.HasFault {
	task := CreateTask(m, "moveVirtualDisk", func(*Task) (types.AnyType, types.BaseMethodFault) {
		fm := ctx.Map.FileManager()

		dest := vdmNames(req.DestName)

		for i, name := range vdmNames(req.SourceName) {
			err := fm.moveDatastoreFile(ctx, &types.MoveDatastoreFile_Task{
				SourceName:            name,
				SourceDatacenter:      req.SourceDatacenter,
				DestinationName:       dest[i],
				DestinationDatacenter: req.DestDatacenter,
				Force:                 req.Force,
			})

			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return &methods.MoveVirtualDisk_TaskBody{
		Res: &types.MoveVirtualDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (m *VirtualDiskManager) CopyVirtualDiskTask(ctx *Context, req *types.CopyVirtualDisk_Task) soap.HasFault {
	task := CreateTask(m, "copyVirtualDisk", func(*Task) (types.AnyType, types.BaseMethodFault) {

		fm := ctx.Map.FileManager()

		dest := vdmNames(req.DestName)

		for i, name := range vdmNames(req.SourceName) {

			if fault := fm.copyDatastoreFile(ctx, &types.CopyDatastoreFile_Task{
				SourceName:            name,
				SourceDatacenter:      req.SourceDatacenter,
				DestinationName:       dest[i],
				DestinationDatacenter: req.DestDatacenter,
				Force:                 req.Force,
			}); fault != nil {
				return nil, fault
			}

			if req.DestSpec != nil {
				desc, descPath, fault := fm.DiskDescriptor(ctx, req.DestDatacenter, dest[i])
				if fault != nil {
					return nil, fault
				}
				if s, ok := req.DestSpec.(*types.FileBackedVirtualDiskSpec); ok {

					var (
						keyID      string
						providerID string
					)
					switch c := s.Crypto.(type) {
					case *types.CryptoSpecEncrypt:
						keyID = c.CryptoKeyId.KeyId
						if c.CryptoKeyId.ProviderId != nil {
							providerID = c.CryptoKeyId.ProviderId.Id
						}
					case *types.CryptoSpecDeepRecrypt:
						keyID = c.NewKeyId.KeyId
						if c.NewKeyId.ProviderId != nil {
							providerID = c.NewKeyId.ProviderId.Id
						}
					case *types.CryptoSpecShallowRecrypt:
						keyID = c.NewKeyId.KeyId
						if c.NewKeyId.ProviderId != nil {
							providerID = c.NewKeyId.ProviderId.Id
						}
					}

					if providerID != "" || keyID != "" {
						desc.EncryptionKeys = &crypto.KeyLocator{
							Type: crypto.KeyLocatorTypeList,
							List: []*crypto.KeyLocator{
								{
									Type: crypto.KeyLocatorTypePair,
									Pair: &crypto.KeyLocatorPair{
										CryptoMAC:  "HMAC-SHA-256",
										LockedData: []byte("Cg=="),
										Locker: &crypto.KeyLocator{
											Type: crypto.KeyLocatorTypeFQID,
											Indirect: &crypto.KeyLocatorIndirect{
												Type: crypto.KeyLocatorTypeFQID,
												FQID: crypto.KeyLocatorFQIDParams{
													KeyServerID: providerID,
													KeyID:       keyID,
												},
											},
										},
									},
								},
							},
						}
					}
				}

				if fault := fm.SaveDiskDescriptor(ctx, desc, descPath); fault != nil {
					return nil, fault
				}
			}
		}

		return nil, nil
	})

	return &methods.CopyVirtualDisk_TaskBody{
		Res: &types.CopyVirtualDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func virtualDiskUUID(dc *types.ManagedObjectReference, file string) string {
	if dc != nil {
		file = dc.String() + file
	}
	return newUUID(file)
}

func (m *VirtualDiskManager) QueryVirtualDiskUuid(ctx *Context, req *types.QueryVirtualDiskUuid) soap.HasFault {
	body := new(methods.QueryVirtualDiskUuidBody)

	fm := ctx.Map.FileManager()

	file, fault := fm.resolve(ctx, req.Datacenter, req.Name)
	if fault != nil {
		body.Fault_ = Fault("", fault)
		return body
	}

	_, err := os.Stat(file)
	if err != nil {
		fault = fm.fault(req.Name, err, new(types.CannotAccessFile))
		body.Fault_ = Fault(fmt.Sprintf("File %s was not found", req.Name), fault)
		return body
	}

	body.Res = &types.QueryVirtualDiskUuidResponse{
		Returnval: virtualDiskUUID(req.Datacenter, req.Name),
	}

	return body
}

func (m *VirtualDiskManager) SetVirtualDiskUuid(_ *Context, req *types.SetVirtualDiskUuid) soap.HasFault {
	body := new(methods.SetVirtualDiskUuidBody)
	// TODO: validate uuid format and persist
	body.Res = new(types.SetVirtualDiskUuidResponse)
	return body
}

func (m *VirtualDiskManager) QueryVirtualDiskInfoTask(ctx *Context, req *internal.QueryVirtualDiskInfoTaskRequest) soap.HasFault {
	task := CreateTask(m, "queryVirtualDiskInfo", func(*Task) (types.AnyType, types.BaseMethodFault) {
		var res []internal.VirtualDiskInfo

		fm := ctx.Map.FileManager()

		_, fault := fm.resolve(ctx, req.Datacenter, req.Name)
		if fault != nil {
			return nil, fault
		}

		res = append(res, internal.VirtualDiskInfo{Name: req.Name, DiskType: "thin"})

		if req.IncludeParents {
			// TODO
		}

		return res, nil
	})

	return &internal.QueryVirtualDiskInfoTaskBody{
		Res: &internal.QueryVirtualDiskInfo_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}
