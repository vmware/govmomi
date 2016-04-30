/*
Copyright (c) 2016 VMware, Inc. All Rights Reserved.

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

package folder

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"golang.org/x/net/context"
)

type move struct {
	*flags.ClientFlag
	*flags.FolderFlag
}

func init() {
	cli.Register("folder.moveinto", &move{})
}

func (cmd *move) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *move) Usage() string {
	return "PATH..."
}

func (cmd *move) Description() string {
	return `Move managed entities into this folder.
Example:
govc folder.moveinto -folder /dc1/folder-foo /dc2/folder-bar/*
`
}

func (cmd *move) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *move) Run(ctx context.Context, f *flag.FlagSet) error {
	folder, err := cmd.Folder()
	if err != nil {
		return err
	}

	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	task, err := folder.MoveInto(ctx, objs)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
