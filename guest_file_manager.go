/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type GuestFileManager struct {
	types.ManagedObjectReference

	c *Client
}

func (m GuestFileManager) Reference() types.ManagedObjectReference {
	return m.ManagedObjectReference
}

func (m GuestFileManager) ChangeFileAttributesInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, guestFilePath string, fileAttributes types.BaseGuestFileAttributes) error {
	req := types.ChangeFileAttributesInGuest{
		This:           m.Reference(),
		Vm:             vm.Reference(),
		Auth:           auth,
		GuestFilePath:  guestFilePath,
		FileAttributes: fileAttributes,
	}

	_, err := methods.ChangeFileAttributesInGuest(m.c, &req)
	return err
}

func (m GuestFileManager) CreateTemporaryDirectoryInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, prefix, suffix string) (string, error) {
	req := types.CreateTemporaryDirectoryInGuest{
		This:   m.Reference(),
		Vm:     vm.Reference(),
		Auth:   auth,
		Prefix: prefix,
		Suffix: suffix,
	}

	res, err := methods.CreateTemporaryDirectoryInGuest(m.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (m GuestFileManager) CreateTemporaryFileInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, prefix, suffix string) (string, error) {
	req := types.CreateTemporaryFileInGuest{
		This:   m.Reference(),
		Vm:     vm.Reference(),
		Auth:   auth,
		Prefix: prefix,
		Suffix: suffix,
	}

	res, err := methods.CreateTemporaryFileInGuest(m.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (m GuestFileManager) DeleteDirectoryInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, directoryPath string, recursive bool) error {
	req := types.DeleteDirectoryInGuest{
		This:          m.Reference(),
		Vm:            vm.Reference(),
		Auth:          auth,
		DirectoryPath: directoryPath,
		Recursive:     recursive,
	}

	_, err := methods.DeleteDirectoryInGuest(m.c, &req)
	return err
}

func (m GuestFileManager) DeleteFileInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, filePath string) error {
	req := types.DeleteFileInGuest{
		This:     m.Reference(),
		Vm:       vm.Reference(),
		Auth:     auth,
		FilePath: filePath,
	}

	_, err := methods.DeleteFileInGuest(m.c, &req)
	return err
}

func (m GuestFileManager) InitiateFileTransferFromGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, guestFilePath string) (*types.FileTransferInformation, error) {
	req := types.InitiateFileTransferFromGuest{
		This:          m.Reference(),
		Vm:            vm.Reference(),
		Auth:          auth,
		GuestFilePath: guestFilePath,
	}

	res, err := methods.InitiateFileTransferFromGuest(m.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (m GuestFileManager) InitiateFileTransferToGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, guestFilePath string, fileAttributes types.BaseGuestFileAttributes, fileSize int64, overwrite bool) (string, error) {
	req := types.InitiateFileTransferToGuest{
		This:           m.Reference(),
		Vm:             vm.Reference(),
		Auth:           auth,
		GuestFilePath:  guestFilePath,
		FileAttributes: fileAttributes,
		FileSize:       fileSize,
		Overwrite:      overwrite,
	}

	res, err := methods.InitiateFileTransferToGuest(m.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

func (m GuestFileManager) ListFilesInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, filePath string, index int, maxResults int, matchPattern string) (*types.GuestListFileInfo, error) {
	req := types.ListFilesInGuest{
		This:         m.Reference(),
		Vm:           vm.Reference(),
		Auth:         auth,
		FilePath:     filePath,
		Index:        index,
		MaxResults:   maxResults,
		MatchPattern: matchPattern,
	}

	res, err := methods.ListFilesInGuest(m.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (m GuestFileManager) MakeDirectoryInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, directoryPath string, createParentDirectories bool) error {
	req := types.MakeDirectoryInGuest{
		This:                    m.Reference(),
		Vm:                      vm.Reference(),
		Auth:                    auth,
		DirectoryPath:           directoryPath,
		CreateParentDirectories: createParentDirectories,
	}

	_, err := methods.MakeDirectoryInGuest(m.c, &req)
	return err
}

func (m GuestFileManager) MoveDirectoryInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, srcDirectoryPath string, dstDirectoryPath string) error {
	req := types.MoveDirectoryInGuest{
		This:             m.Reference(),
		Vm:               vm.Reference(),
		Auth:             auth,
		SrcDirectoryPath: srcDirectoryPath,
		DstDirectoryPath: dstDirectoryPath,
	}

	_, err := methods.MoveDirectoryInGuest(m.c, &req)
	return err
}

func (m GuestFileManager) MoveFileInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, srcFilePath string, dstFilePath string, overwrite bool) error {
	req := types.MoveFileInGuest{
		This:        m.Reference(),
		Vm:          vm.Reference(),
		Auth:        auth,
		SrcFilePath: srcFilePath,
		DstFilePath: dstFilePath,
		Overwrite:   overwrite,
	}

	_, err := methods.MoveFileInGuest(m.c, &req)
	return err
}
