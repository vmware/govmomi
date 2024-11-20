/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package task

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.ClientFlag

	obj string
}

func init() {
	cli.Register("task.create", &create{}, true)
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.obj, "o", "", "ManagedObject with which Task will be associated")
}

func (cmd *create) Description() string {
	return `Create task of type ID.

ID must be one of:
  govc extension.info -json | jq -r '.extensions[].taskList | select(. != null) | .[].taskID'

Examples:
  govc task.create $ID`
}

func (cmd *create) Usage() string {
	return "ID"
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	m := task.NewManager(c)

	req := types.CreateTask{
		This:       m.Reference(),
		Obj:        c.ServiceContent.RootFolder,
		TaskTypeId: f.Arg(0),
		Cancelable: true,
	}

	if cmd.obj != "" {
		req.Obj.FromString(cmd.obj)
	}

	s, err := session.NewManager(c).UserSession(ctx)
	if err != nil {
		return err
	}
	req.InitiatedBy = s.UserName

	res, err := methods.CreateTask(ctx, c, &req)
	if err != nil {
		return err
	}

	fmt.Println(res.Returnval.Task.Value)

	return nil
}
