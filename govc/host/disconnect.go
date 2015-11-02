/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package host

import (
	"flag"
	"fmt"

	"golang.org/x/net/context"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
)

type disconnect struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.disconnect", &disconnect{})
}

func (cmd *disconnect) Register(f *flag.FlagSet) {}

func (cmd *disconnect) Process() error { return nil }

func (cmd *disconnect) Description() string {
	return `Disconnect host from vCenter and instruct the host to stop sending heartbeats.`
}

func (cmd *disconnect) Disconnect(ctx context.Context, host *object.HostSystem) error {
	task, err := host.Disconnect(ctx)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s disconnecting... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *disconnect) Run(f *flag.FlagSet) error {
	ctx := context.TODO()

	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.Disconnect(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
