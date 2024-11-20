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

package events

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/vim25/types"
)

type post struct {
	*flags.DatacenterFlag

	types.EventEx
}

func init() {
	cli.Register("event.post", &post{}, true)
}

func (cmd *post) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.EventTypeId, "i", "", "Event Type ID")
	f.StringVar(&cmd.Message, "m", "", "Event message")
	f.StringVar(&cmd.Severity, "s", string(types.EventEventSeverityInfo), "Event severity")
}

func (cmd *post) Usage() string {
	return "PATH"
}

func (cmd *post) Description() string {
	return `Post Event.

Examples:
  govc event.post -s warning -i com.vmware.wcp.RegisterVM.failure $vm
  govc event.post -s info -i com.vmware.wcp.RegisterVM.success $vm
  govc event.post -m "cluster degraded" /dc1/host/cluster1`
}

func (cmd *post) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	obj, err := cmd.ManagedObject(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	cmd.ObjectType = obj.Type
	cmd.ObjectId = obj.Value

	return event.NewManager(c).PostEvent(ctx, &cmd.EventEx)
}
