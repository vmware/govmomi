// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type FileAttrFlag struct {
	types.GuestPosixFileAttributes
}

func newFileAttrFlag(ctx context.Context) (*FileAttrFlag, context.Context) {
	return &FileAttrFlag{}, ctx
}

func (flag *FileAttrFlag) Register(ctx context.Context, f *flag.FlagSet) {
	f.Var(flags.NewOptionalInt32(&flag.OwnerId), "uid", "User ID")
	f.Var(flags.NewOptionalInt32(&flag.GroupId), "gid", "Group ID")
	f.Int64Var(&flag.Permissions, "perm", 0, "File permissions")
}

func (flag *FileAttrFlag) Process(ctx context.Context) error {
	return nil
}

func (flag *FileAttrFlag) Attr() types.BaseGuestFileAttributes {
	return &flag.GuestPosixFileAttributes
}
