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
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type rename struct {
	*flags.ClientFlag
	*flags.FolderFlag
}

func init() {
	cli.Register("folder.rename", &rename{})
}

func (cmd *rename) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *rename) Usage() string {
	return "NAME"
}

func (cmd *rename) Description() string {
	return `Rename an existing folder with NAME.
Example:
govc folder.rename -folder /dc1/vm/folder-foo folder-bar
`
}

func (cmd *rename) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rename) Run(ctx context.Context, f *flag.FlagSet) error {
	folder, err := cmd.Folder()
	if err != nil {
		return err
	}

	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	task, err := folder.Rename(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}
