// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package hgfs

import (
	"os"
	"syscall"
)

const attrMask = AttrValidAllocationSize |
	AttrValidAccessTime | AttrValidWriteTime | AttrValidCreateTime | AttrValidChangeTime |
	AttrValidSpecialPerms | AttrValidOwnerPerms | AttrValidGroupPerms | AttrValidOtherPerms | AttrValidEffectivePerms |
	AttrValidUserID | AttrValidGroupID | AttrValidFileID | AttrValidVolID

func (a *AttrV2) sysStat(info os.FileInfo) {
	sys, ok := info.Sys().(*syscall.Stat_t)

	if !ok {
		return
	}

	a.AllocationSize = uint64(sys.Blocks * 512)

	nt := func(t syscall.Timespec) uint64 {
		return uint64(t.Nano()) // TODO: this is supposed to be Windows NT system time, not needed atm
	}

	a.AccessTime = nt(sys.Atim)
	a.WriteTime = nt(sys.Mtim)
	a.CreationTime = a.WriteTime // see HgfsGetCreationTime
	a.AttrChangeTime = nt(sys.Ctim)

	a.SpecialPerms = uint8((sys.Mode & (syscall.S_ISUID | syscall.S_ISGID | syscall.S_ISVTX)) >> 9)
	a.OwnerPerms = uint8((sys.Mode & syscall.S_IRWXU) >> 6)
	a.GroupPerms = uint8((sys.Mode & syscall.S_IRWXG) >> 3)
	a.OtherPerms = uint8(sys.Mode & syscall.S_IRWXO)

	a.UserID = sys.Uid
	a.GroupID = sys.Gid
	a.HostFileID = sys.Ino
	a.VolumeID = uint32(sys.Dev)

	a.Mask |= attrMask
}
